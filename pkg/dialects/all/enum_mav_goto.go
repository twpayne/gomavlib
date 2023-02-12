//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Actions that may be specified in MAV_CMD_OVERRIDE_GOTO to override mission execution.
type MAV_GOTO = common.MAV_GOTO

const (
	// Hold at the current position.
	MAV_GOTO_DO_HOLD MAV_GOTO = common.MAV_GOTO_DO_HOLD
	// Continue with the next item in mission execution.
	MAV_GOTO_DO_CONTINUE MAV_GOTO = common.MAV_GOTO_DO_CONTINUE
	// Hold at the current position of the system
	MAV_GOTO_HOLD_AT_CURRENT_POSITION MAV_GOTO = common.MAV_GOTO_HOLD_AT_CURRENT_POSITION
	// Hold at the position specified in the parameters of the DO_HOLD action
	MAV_GOTO_HOLD_AT_SPECIFIED_POSITION MAV_GOTO = common.MAV_GOTO_HOLD_AT_SPECIFIED_POSITION
)
