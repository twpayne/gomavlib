/*
Package gomavlib is a library that implements Mavlink 2.0 and 1.0 in the Go
programming language. It can power UGVs, UAVs, ground stations, monitoring
systems or routers acting in a Mavlink network.

Mavlink is a lighweight and endpoint-independent protocol that is mostly
used to communicate with unmanned ground vehicles (UGV) and unmanned aerial
vehicles (UAV, drones, quadcopters, multirotors). It is supported by both
of the most common open-source flight controllers (Ardupilot and PX4).

Basic example (more are available at https://github.com/gswly/gomavlib/tree/master/example)

  package main

  import (
  	"fmt"
  	"github.com/gswly/gomavlib"
  	"github.com/gswly/gomavlib/dialects/ardupilotmega"
  )

  func main() {
  	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointSerial{"/dev/ttyUSB0:57600"},
		},
  		Dialect:     ardupilotmega.Dialect,
  		SystemId:    10,
  		ComponentId: 1,
  	})
  	if err != nil {
  		panic(err)
  	}
  	defer node.Close()

  	for {
  		res, ok := node.Read()
  		if ok == false {
  			break
  		}

  		fmt.Printf("received: id=%d, %+v\n", res.Message().GetId(), res.Message())
  	}
  }

*/
package gomavlib

import (
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	// constant for ip-based endpoints
	netBufferSize      = 512 // frames cannot go beyond len(header) + 255 + len(check) + len(sig)
	netConnectTimeout  = 10 * time.Second
	netReconnectPeriod = 2 * time.Second
	netReadTimeout     = 60 * time.Second
	netWriteTimeout    = 10 * time.Second
)

// NodeVersion allows to set the frame version used in a Node to wrap outgoing messages.
type NodeVersion int

const (
	// V2 wrap outgoing messages in v2 frames.
	V2 NodeVersion = iota
	// V1 wrap outgoing messages in v1 frames.
	V1
)

// NodeConf allows to configure a Node.
type NodeConf struct {
	// contains the endpoint with which this node will
	// communicate. Each endpoint contains one or more channels.
	Endpoints []EndpointConf

	// contains the messages which will be automatically decoded and
	// encoded. If not provided, messages are decoded in the MessageRaw struct.
	Dialect *Dialect

	// Mavlink version used to encode frames. See Version
	// for the available options.
	Version NodeVersion

	// these are used to identify this node in the network.
	// They are added to every outgoing message.
	SystemId    byte
	ComponentId byte

	// (optional) the secret key used to validate incoming frames.
	// Non signed frames are discarded. This feature requires Mavlink v2.
	SignatureInKey *FrameSignatureKey
	// (optional) the secret key used to sign outgoing frames.
	// This feature requires Mavlink v2.
	SignatureOutKey *FrameSignatureKey

	// (optional) disables the periodic sending of heartbeats to
	// open channels.
	HeartbeatDisable bool
	// (optional) set the period between heartbeats.
	// It defaults to 5 seconds.
	HeartbeatPeriod time.Duration

	// (optional) return parse errors when calling Read()
	// normally parse errors are ignored.
	ReturnParseErrors bool

	// (optional) disables checksum validation of incoming frames.
	// Not recommended, useful only for debugging purposes.
	ChecksumDisable bool
}

// Node is a high-level Mavlink encoder and decoder that works with endpoints.
// See NodeConf for the options.
type Node struct {
	conf          NodeConf
	wg            sync.WaitGroup
	chanAccepters map[endpointChannelAccepter]struct{}
	channelsMutex sync.Mutex
	channels      map[*EndpointChannel]struct{}
	frameQueue    chan *NodeReadResult
	writeDone     chan struct{}
	nodeHeartbeat *nodeHeartbeat
}

// NewNode allocates a Node. See NodeConf for the options.
func NewNode(conf NodeConf) (*Node, error) {
	if conf.SystemId < 1 {
		return nil, fmt.Errorf("SystemId must be >= 1")
	}
	if conf.ComponentId < 1 {
		return nil, fmt.Errorf("ComponentId must be >= 1")
	}
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("at least one endpoint must be provided")
	}
	if conf.SignatureInKey != nil && conf.Version != V2 {
		return nil, fmt.Errorf("SignatureInKey requires V2 frames")
	}
	if conf.SignatureOutKey != nil && conf.Version != V2 {
		return nil, fmt.Errorf("SignatureOutKey requires V2 frames")
	}
	if conf.HeartbeatPeriod == 0 {
		conf.HeartbeatPeriod = 5 * time.Second
	}

	n := &Node{
		conf:          conf,
		chanAccepters: make(map[endpointChannelAccepter]struct{}),
		channels:      make(map[*EndpointChannel]struct{}),
		frameQueue:    make(chan *NodeReadResult),
		writeDone:     make(chan struct{}),
	}

	// init endpoints
	for _, tconf := range conf.Endpoints {
		tp, err := tconf.init()
		if err != nil {
			n.Close()
			return nil, err
		}

		if tm, ok := tp.(endpointChannelAccepter); ok {
			n.startChannelAccepter(tm)

		} else if ts, ok := tp.(endpointChannelSingle); ok {
			n.startChannel(ts)

		} else {
			panic(fmt.Errorf("endpoint %d does not implement required interface", tp))
		}
	}

	// start heartbeat
	if n.conf.HeartbeatDisable == false {
		n.nodeHeartbeat = newNodeHeartbeat(n, n.conf.HeartbeatPeriod)
	}
	return n, nil
}

// Close stops node operations and wait for all routines to return.
func (n *Node) Close() {
	func() {
		if n.nodeHeartbeat != nil {
			n.nodeHeartbeat.close()
		}

		for mc := range n.chanAccepters {
			mc.Close()
		}

		n.channelsMutex.Lock()
		defer n.channelsMutex.Unlock()

		for ch := range n.channels {
			ch.rwc.Close()
		}
	}()

	// consume queued frames up to close(n.frameQueue)
	// in case the user is not calling Read() in a loop.
	go func() {
		for {
			_, ok := n.Read()
			if ok == false {
				break
			}
		}
	}()

	n.wg.Wait()

	// close queue after ensuring no one will write to it
	close(n.frameQueue)
}

func (n *Node) startChannelAccepter(tm endpointChannelAccepter) {
	n.chanAccepters[tm] = struct{}{}

	n.wg.Add(1)
	go func() {
		defer n.wg.Done()

		for {
			ch, err := tm.Accept()
			if err != nil {
				if err != errorTerminated {
					panic("errorTerminated is the only error allowed here")
				}
				break
			}

			n.startChannel(ch)
		}
	}()
}

func (n *Node) startChannel(rwc io.ReadWriteCloser) {
	n.channelsMutex.Lock()
	defer n.channelsMutex.Unlock()

	conn := &EndpointChannel{
		rwc:       rwc,
		writeChan: make(chan interface{}),
	}
	n.channels[conn] = struct{}{}

	parser := NewParser(ParserConf{
		Reader:          conn.rwc,
		Writer:          conn.rwc,
		Dialect:         n.conf.Dialect,
		SystemId:        n.conf.SystemId,
		ComponentId:     n.conf.ComponentId,
		SignatureInKey:  n.conf.SignatureInKey,
		SignatureLinkId: randomByte(),
		SignatureOutKey: n.conf.SignatureOutKey,
		ChecksumDisable: n.conf.ChecksumDisable,
	})

	// reader
	n.wg.Add(1)
	go func() {
		defer n.wg.Done()
		defer func() {
			n.channelsMutex.Lock()
			defer n.channelsMutex.Unlock()
			delete(n.channels, conn)
			close(conn.writeChan)
		}()

		for {
			frame, err := parser.Read()
			if err != nil {
				// continue in case of parse errors
				if _, ok := err.(*ParserError); ok {
					if n.conf.ReturnParseErrors == true {
						n.frameQueue <- &NodeReadResult{nil, err, conn}
					}
					continue
				}
				// avoid calling twice Close()
				if err != errorTerminated {
					conn.rwc.Close()
				}
				return
			}

			n.frameQueue <- &NodeReadResult{frame, nil, conn}
		}
	}()

	// writer
	n.wg.Add(1)
	go func() {
		defer n.wg.Done()

		for {
			what, ok := <-conn.writeChan
			if ok == false {
				return
			}

			switch wh := what.(type) {
			case Message:
				if n.conf.Version == V1 {
					parser.Write(&FrameV1{Message: wh}, false)
				} else {
					parser.Write(&FrameV2{Message: wh}, false)
				}

			case Frame:
				parser.Write(wh, true)
			}

			n.writeDone <- struct{}{}
		}
	}()
}

// NodeReadResult contains the result of node.Read()
type NodeReadResult struct {
	// the just read frame
	Frame Frame

	// a parse error returned instead of frame
	// This is used only when ReturnParseErrors is true
	Error error

	// the channel used to send the frame
	Channel *EndpointChannel
}

// Message returns the message inside the frame.
func (res *NodeReadResult) Message() Message {
	return res.Frame.GetMessage()
}

// SystemId returns the sender system id.
func (res *NodeReadResult) SystemId() byte {
	return res.Frame.GetSystemId()
}

// ComponentId returns the sender component id.
func (res *NodeReadResult) ComponentId() byte {
	return res.Frame.GetComponentId()
}

// Read reads a single message from available channels.
// NodeReadResult contains all the properties of the received message
// (see NodeReadResult for details).
// bool is true whenever the node is still open.
func (n *Node) Read() (*NodeReadResult, bool) {
	res, ok := <-n.frameQueue
	return res, ok
}

// WriteMessageTo write a message to given channel.
func (n *Node) WriteMessageTo(channel *EndpointChannel, message Message) {
	n.writeTo(channel, message)
}

// WriteMessageAll write a message to all channels.
func (n *Node) WriteMessageAll(message Message) {
	n.writeAll(message)
}

// WriteMessageExcept write a message to all channels except specified channel.
func (n *Node) WriteMessageExcept(exceptChannel *EndpointChannel, message Message) {
	n.writeExcept(exceptChannel, message)
}

// WriteFrameTo write a frame to given channel.
// This function is intended for routing frames to other nodes, since all
// the fields must be filled manually.
func (n *Node) WriteFrameTo(channel *EndpointChannel, frame Frame) {
	n.writeTo(channel, frame)
}

// WriteFrameAll write a frame to all channels.
// This function is intended for routing frames to other nodes, since all
// the fields must be filled manually.
func (n *Node) WriteFrameAll(frame Frame) {
	n.writeAll(frame)
}

// WriteFrameExcept write a frame to all channels except specified channel.
// This function is intended for routing frames to other nodes, since all
// the fields must be filled manually.
func (n *Node) WriteFrameExcept(exceptChannel *EndpointChannel, frame Frame) {
	n.writeExcept(exceptChannel, frame)
}

func (n *Node) writeTo(channel *EndpointChannel, what interface{}) {
	n.channelsMutex.Lock()
	defer n.channelsMutex.Unlock()

	if _, ok := n.channels[channel]; ok == false {
		return
	}

	// route to channels
	// wait for responses (otherwise endpoints can be removed before writing)
	channel.writeChan <- what
	<-n.writeDone
}

func (n *Node) writeAll(what interface{}) {
	n.channelsMutex.Lock()
	defer n.channelsMutex.Unlock()

	for conn := range n.channels {
		conn.writeChan <- what
		defer func() { <-n.writeDone }()
	}
}

func (n *Node) writeExcept(exceptChannel *EndpointChannel, what interface{}) {
	n.channelsMutex.Lock()
	defer n.channelsMutex.Unlock()

	for conn := range n.channels {
		if conn != exceptChannel {
			conn.writeChan <- what
			defer func() { <-n.writeDone }()
		}
	}
}
