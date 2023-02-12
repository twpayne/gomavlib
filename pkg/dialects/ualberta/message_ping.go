//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// A ping message either requesting or responding to a ping. This allows to measure the system latencies, including serial port, radio modem and UDP connections. The ping microservice is documented at https://mavlink.io/en/services/ping.html
type MessagePing = common.MessagePing
