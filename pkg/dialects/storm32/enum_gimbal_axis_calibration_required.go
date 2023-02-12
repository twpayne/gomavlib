//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
)

type GIMBAL_AXIS_CALIBRATION_REQUIRED = ardupilotmega.GIMBAL_AXIS_CALIBRATION_REQUIRED

const (
	// Whether or not this axis requires calibration is unknown at this time.
	GIMBAL_AXIS_CALIBRATION_REQUIRED_UNKNOWN GIMBAL_AXIS_CALIBRATION_REQUIRED = ardupilotmega.GIMBAL_AXIS_CALIBRATION_REQUIRED_UNKNOWN
	// This axis requires calibration.
	GIMBAL_AXIS_CALIBRATION_REQUIRED_TRUE GIMBAL_AXIS_CALIBRATION_REQUIRED = ardupilotmega.GIMBAL_AXIS_CALIBRATION_REQUIRED_TRUE
	// This axis does not require calibration.
	GIMBAL_AXIS_CALIBRATION_REQUIRED_FALSE GIMBAL_AXIS_CALIBRATION_REQUIRED = ardupilotmega.GIMBAL_AXIS_CALIBRATION_REQUIRED_FALSE
)
