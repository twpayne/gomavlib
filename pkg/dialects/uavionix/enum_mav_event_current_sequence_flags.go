//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Flags for CURRENT_EVENT_SEQUENCE.
type MAV_EVENT_CURRENT_SEQUENCE_FLAGS = common.MAV_EVENT_CURRENT_SEQUENCE_FLAGS

const (
	// A sequence reset has happened (e.g. vehicle reboot).
	MAV_EVENT_CURRENT_SEQUENCE_FLAGS_RESET MAV_EVENT_CURRENT_SEQUENCE_FLAGS = common.MAV_EVENT_CURRENT_SEQUENCE_FLAGS_RESET
)
