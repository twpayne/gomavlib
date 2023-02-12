//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// SERIAL_CONTROL flags (bitmask)
type SERIAL_CONTROL_FLAG = common.SERIAL_CONTROL_FLAG

const (
	// Set if this is a reply
	SERIAL_CONTROL_FLAG_REPLY SERIAL_CONTROL_FLAG = common.SERIAL_CONTROL_FLAG_REPLY
	// Set if the sender wants the receiver to send a response as another SERIAL_CONTROL message
	SERIAL_CONTROL_FLAG_RESPOND SERIAL_CONTROL_FLAG = common.SERIAL_CONTROL_FLAG_RESPOND
	// Set if access to the serial port should be removed from whatever driver is currently using it, giving exclusive access to the SERIAL_CONTROL protocol. The port can be handed back by sending a request without this flag set
	SERIAL_CONTROL_FLAG_EXCLUSIVE SERIAL_CONTROL_FLAG = common.SERIAL_CONTROL_FLAG_EXCLUSIVE
	// Block on writes to the serial port
	SERIAL_CONTROL_FLAG_BLOCKING SERIAL_CONTROL_FLAG = common.SERIAL_CONTROL_FLAG_BLOCKING
	// Send multiple replies until port is drained
	SERIAL_CONTROL_FLAG_MULTI SERIAL_CONTROL_FLAG = common.SERIAL_CONTROL_FLAG_MULTI
)
