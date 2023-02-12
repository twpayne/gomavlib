//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Aircraft-rated danger from this threat.
type MAV_COLLISION_THREAT_LEVEL = common.MAV_COLLISION_THREAT_LEVEL

const (
	// Not a threat
	MAV_COLLISION_THREAT_LEVEL_NONE MAV_COLLISION_THREAT_LEVEL = common.MAV_COLLISION_THREAT_LEVEL_NONE
	// Craft is mildly concerned about this threat
	MAV_COLLISION_THREAT_LEVEL_LOW MAV_COLLISION_THREAT_LEVEL = common.MAV_COLLISION_THREAT_LEVEL_LOW
	// Craft is panicking, and may take actions to avoid threat
	MAV_COLLISION_THREAT_LEVEL_HIGH MAV_COLLISION_THREAT_LEVEL = common.MAV_COLLISION_THREAT_LEVEL_HIGH
)
