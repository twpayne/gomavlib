//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Flags for high level gimbal manager operation The first 16 bits are identical to the GIMBAL_DEVICE_FLAGS.
type GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS

const (
	// Based on GIMBAL_DEVICE_FLAGS_RETRACT.
	GIMBAL_MANAGER_FLAGS_RETRACT GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_RETRACT
	// Based on GIMBAL_DEVICE_FLAGS_NEUTRAL.
	GIMBAL_MANAGER_FLAGS_NEUTRAL GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_NEUTRAL
	// Based on GIMBAL_DEVICE_FLAGS_ROLL_LOCK.
	GIMBAL_MANAGER_FLAGS_ROLL_LOCK GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_ROLL_LOCK
	// Based on GIMBAL_DEVICE_FLAGS_PITCH_LOCK.
	GIMBAL_MANAGER_FLAGS_PITCH_LOCK GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_PITCH_LOCK
	// Based on GIMBAL_DEVICE_FLAGS_YAW_LOCK.
	GIMBAL_MANAGER_FLAGS_YAW_LOCK GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_YAW_LOCK
	// Based on GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME.
	GIMBAL_MANAGER_FLAGS_YAW_IN_VEHICLE_FRAME GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_YAW_IN_VEHICLE_FRAME
	// Based on GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME.
	GIMBAL_MANAGER_FLAGS_YAW_IN_EARTH_FRAME GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_YAW_IN_EARTH_FRAME
	// Based on GIMBAL_DEVICE_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME.
	GIMBAL_MANAGER_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME
	// Based on GIMBAL_DEVICE_FLAGS_RC_EXCLUSIVE.
	GIMBAL_MANAGER_FLAGS_RC_EXCLUSIVE GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_RC_EXCLUSIVE
	// Based on GIMBAL_DEVICE_FLAGS_RC_MIXED.
	GIMBAL_MANAGER_FLAGS_RC_MIXED GIMBAL_MANAGER_FLAGS = common.GIMBAL_MANAGER_FLAGS_RC_MIXED
)
