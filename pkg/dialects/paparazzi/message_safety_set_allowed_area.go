//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Set a safety zone (volume), which is defined by two corners of a cube. This message can be used to tell the MAV which setpoints/waypoints to accept and which to reject. Safety areas are often enforced by national or competition regulations.
type MessageSafetySetAllowedArea = common.MessageSafetySetAllowedArea
