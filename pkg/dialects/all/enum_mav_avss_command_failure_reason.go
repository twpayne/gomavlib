//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/avssuas"
)

type MAV_AVSS_COMMAND_FAILURE_REASON = avssuas.MAV_AVSS_COMMAND_FAILURE_REASON

const (
	// AVSS defined command failure reason. PRS not steady.
	PRS_NOT_STEADY MAV_AVSS_COMMAND_FAILURE_REASON = avssuas.PRS_NOT_STEADY
	// AVSS defined command failure reason. PRS DTM not armed.
	PRS_DTM_NOT_ARMED MAV_AVSS_COMMAND_FAILURE_REASON = avssuas.PRS_DTM_NOT_ARMED
	// AVSS defined command failure reason. PRS OTM not armed.
	PRS_OTM_NOT_ARMED MAV_AVSS_COMMAND_FAILURE_REASON = avssuas.PRS_OTM_NOT_ARMED
)
