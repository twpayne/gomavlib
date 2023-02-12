//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Reports the current commanded attitude of the vehicle as specified by the autopilot. This should match the commands sent in a SET_ATTITUDE_TARGET message if the vehicle is being controlled this way.
type MessageAttitudeTarget = common.MessageAttitudeTarget
