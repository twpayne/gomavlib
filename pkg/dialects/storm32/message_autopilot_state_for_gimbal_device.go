//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Low level message containing autopilot state relevant for a gimbal device. This message is to be sent from the autopilot to the gimbal device component. The data of this message are for the gimbal device's estimator corrections, in particular horizon compensation, as well as indicates autopilot control intentions, e.g. feed forward angular control in the z-axis.
type MessageAutopilotStateForGimbalDevice = common.MessageAutopilotStateForGimbalDevice
