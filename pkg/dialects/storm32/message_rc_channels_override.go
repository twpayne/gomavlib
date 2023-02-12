//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// The RAW values of the RC channels sent to the MAV to override info received from the RC radio. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.  Note carefully the semantic differences between the first 8 channels and the subsequent channels
type MessageRcChannelsOverride = common.MessageRcChannelsOverride
