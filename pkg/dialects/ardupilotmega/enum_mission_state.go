//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// States of the mission state machine.
// Note that these states are independent of whether the mission is in a mode that can execute mission items or not (is suspended).
// They may not all be relevant on all vehicles.
type MISSION_STATE = common.MISSION_STATE

const (
	// The mission status reporting is not supported.
	MISSION_STATE_UNKNOWN MISSION_STATE = common.MISSION_STATE_UNKNOWN
	// No mission on the vehicle.
	MISSION_STATE_NO_MISSION MISSION_STATE = common.MISSION_STATE_NO_MISSION
	// Mission has not started. This is the case after a mission has uploaded but not yet started executing.
	MISSION_STATE_NOT_STARTED MISSION_STATE = common.MISSION_STATE_NOT_STARTED
	// Mission is active, and will execute mission items when in auto mode.
	MISSION_STATE_ACTIVE MISSION_STATE = common.MISSION_STATE_ACTIVE
	// Mission is paused when in auto mode.
	MISSION_STATE_PAUSED MISSION_STATE = common.MISSION_STATE_PAUSED
	// Mission has executed all mission items.
	MISSION_STATE_COMPLETE MISSION_STATE = common.MISSION_STATE_COMPLETE
)
