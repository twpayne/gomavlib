//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Superseded by ACTUATOR_OUTPUT_STATUS. The RAW values of the servo outputs (for RC input from the remote, use the RC_CHANNELS messages). The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.
type MessageServoOutputRaw = common.MessageServoOutputRaw
