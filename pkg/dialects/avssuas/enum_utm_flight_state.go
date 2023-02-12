//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Airborne status of UAS.
type UTM_FLIGHT_STATE = common.UTM_FLIGHT_STATE

const (
	// The flight state can't be determined.
	UTM_FLIGHT_STATE_UNKNOWN UTM_FLIGHT_STATE = common.UTM_FLIGHT_STATE_UNKNOWN
	// UAS on ground.
	UTM_FLIGHT_STATE_GROUND UTM_FLIGHT_STATE = common.UTM_FLIGHT_STATE_GROUND
	// UAS airborne.
	UTM_FLIGHT_STATE_AIRBORNE UTM_FLIGHT_STATE = common.UTM_FLIGHT_STATE_AIRBORNE
	// UAS is in an emergency flight state.
	UTM_FLIGHT_STATE_EMERGENCY UTM_FLIGHT_STATE = common.UTM_FLIGHT_STATE_EMERGENCY
	// UAS has no active controls.
	UTM_FLIGHT_STATE_NOCTRL UTM_FLIGHT_STATE = common.UTM_FLIGHT_STATE_NOCTRL
)
