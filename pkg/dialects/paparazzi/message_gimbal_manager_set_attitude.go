//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// High level message to control a gimbal's attitude. This message is to be sent to the gimbal manager (e.g. from a ground station). Angles and rates can be set to NaN according to use case.
type MessageGimbalManagerSetAttitude = common.MessageGimbalManagerSetAttitude
