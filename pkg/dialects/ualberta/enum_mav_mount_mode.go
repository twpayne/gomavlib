//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Enumeration of possible mount operation modes. This message is used by obsolete/deprecated gimbal messages.
type MAV_MOUNT_MODE = common.MAV_MOUNT_MODE

const (
	// Load and keep safe position (Roll,Pitch,Yaw) from permant memory and stop stabilization
	MAV_MOUNT_MODE_RETRACT MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_RETRACT
	// Load and keep neutral position (Roll,Pitch,Yaw) from permanent memory.
	MAV_MOUNT_MODE_NEUTRAL MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_NEUTRAL
	// Load neutral position and start MAVLink Roll,Pitch,Yaw control with stabilization
	MAV_MOUNT_MODE_MAVLINK_TARGETING MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_MAVLINK_TARGETING
	// Load neutral position and start RC Roll,Pitch,Yaw control with stabilization
	MAV_MOUNT_MODE_RC_TARGETING MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_RC_TARGETING
	// Load neutral position and start to point to Lat,Lon,Alt
	MAV_MOUNT_MODE_GPS_POINT MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_GPS_POINT
	// Gimbal tracks system with specified system ID
	MAV_MOUNT_MODE_SYSID_TARGET MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_SYSID_TARGET
	// Gimbal tracks home position
	MAV_MOUNT_MODE_HOME_LOCATION MAV_MOUNT_MODE = common.MAV_MOUNT_MODE_HOME_LOCATION
)
