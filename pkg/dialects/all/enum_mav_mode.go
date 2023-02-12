//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// These defines are predefined OR-combined mode flags. There is no need to use values from this enum, but it
// simplifies the use of the mode flags. Note that manual input is enabled in all modes as a safety override.
type MAV_MODE = common.MAV_MODE

const (
	// System is not ready to fly, booting, calibrating, etc. No flag is set.
	MAV_MODE_PREFLIGHT MAV_MODE = common.MAV_MODE_PREFLIGHT
	// System is allowed to be active, under assisted RC control.
	MAV_MODE_STABILIZE_DISARMED MAV_MODE = common.MAV_MODE_STABILIZE_DISARMED
	// System is allowed to be active, under assisted RC control.
	MAV_MODE_STABILIZE_ARMED MAV_MODE = common.MAV_MODE_STABILIZE_ARMED
	// System is allowed to be active, under manual (RC) control, no stabilization
	MAV_MODE_MANUAL_DISARMED MAV_MODE = common.MAV_MODE_MANUAL_DISARMED
	// System is allowed to be active, under manual (RC) control, no stabilization
	MAV_MODE_MANUAL_ARMED MAV_MODE = common.MAV_MODE_MANUAL_ARMED
	// System is allowed to be active, under autonomous control, manual setpoint
	MAV_MODE_GUIDED_DISARMED MAV_MODE = common.MAV_MODE_GUIDED_DISARMED
	// System is allowed to be active, under autonomous control, manual setpoint
	MAV_MODE_GUIDED_ARMED MAV_MODE = common.MAV_MODE_GUIDED_ARMED
	// System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
	MAV_MODE_AUTO_DISARMED MAV_MODE = common.MAV_MODE_AUTO_DISARMED
	// System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
	MAV_MODE_AUTO_ARMED MAV_MODE = common.MAV_MODE_AUTO_ARMED
	// UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only.
	MAV_MODE_TEST_DISARMED MAV_MODE = common.MAV_MODE_TEST_DISARMED
	// UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only.
	MAV_MODE_TEST_ARMED MAV_MODE = common.MAV_MODE_TEST_ARMED
)
