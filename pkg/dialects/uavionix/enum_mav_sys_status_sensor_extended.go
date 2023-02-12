//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// These encode the sensors whose status is sent as part of the SYS_STATUS message in the extended fields.
type MAV_SYS_STATUS_SENSOR_EXTENDED = common.MAV_SYS_STATUS_SENSOR_EXTENDED

const (
	// 0x01 Recovery system (parachute, balloon, retracts etc)
	MAV_SYS_STATUS_RECOVERY_SYSTEM MAV_SYS_STATUS_SENSOR_EXTENDED = common.MAV_SYS_STATUS_RECOVERY_SYSTEM
)
