//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// The PPM values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.  A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type MessageRcChannels = common.MessageRcChannels
