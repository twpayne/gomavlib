//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b00000000 indicates that none of the setpoint dimensions should be ignored.
type ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK

const (
	// Ignore body roll rate
	ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE
	// Ignore body pitch rate
	ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE
	// Ignore body yaw rate
	ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE
	// Use 3D body thrust setpoint instead of throttle
	ATTITUDE_TARGET_TYPEMASK_THRUST_BODY_SET ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK_THRUST_BODY_SET
	// Ignore throttle
	ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE
	// Ignore attitude
	ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE ATTITUDE_TARGET_TYPEMASK = common.ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE
)
