//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
)

// A mapping of plane flight modes for custom_mode field of heartbeat.
type PLANE_MODE = ardupilotmega.PLANE_MODE

const (
	PLANE_MODE_MANUAL        PLANE_MODE = ardupilotmega.PLANE_MODE_MANUAL
	PLANE_MODE_CIRCLE        PLANE_MODE = ardupilotmega.PLANE_MODE_CIRCLE
	PLANE_MODE_STABILIZE     PLANE_MODE = ardupilotmega.PLANE_MODE_STABILIZE
	PLANE_MODE_TRAINING      PLANE_MODE = ardupilotmega.PLANE_MODE_TRAINING
	PLANE_MODE_ACRO          PLANE_MODE = ardupilotmega.PLANE_MODE_ACRO
	PLANE_MODE_FLY_BY_WIRE_A PLANE_MODE = ardupilotmega.PLANE_MODE_FLY_BY_WIRE_A
	PLANE_MODE_FLY_BY_WIRE_B PLANE_MODE = ardupilotmega.PLANE_MODE_FLY_BY_WIRE_B
	PLANE_MODE_CRUISE        PLANE_MODE = ardupilotmega.PLANE_MODE_CRUISE
	PLANE_MODE_AUTOTUNE      PLANE_MODE = ardupilotmega.PLANE_MODE_AUTOTUNE
	PLANE_MODE_AUTO          PLANE_MODE = ardupilotmega.PLANE_MODE_AUTO
	PLANE_MODE_RTL           PLANE_MODE = ardupilotmega.PLANE_MODE_RTL
	PLANE_MODE_LOITER        PLANE_MODE = ardupilotmega.PLANE_MODE_LOITER
	PLANE_MODE_TAKEOFF       PLANE_MODE = ardupilotmega.PLANE_MODE_TAKEOFF
	PLANE_MODE_AVOID_ADSB    PLANE_MODE = ardupilotmega.PLANE_MODE_AVOID_ADSB
	PLANE_MODE_GUIDED        PLANE_MODE = ardupilotmega.PLANE_MODE_GUIDED
	PLANE_MODE_INITIALIZING  PLANE_MODE = ardupilotmega.PLANE_MODE_INITIALIZING
	PLANE_MODE_QSTABILIZE    PLANE_MODE = ardupilotmega.PLANE_MODE_QSTABILIZE
	PLANE_MODE_QHOVER        PLANE_MODE = ardupilotmega.PLANE_MODE_QHOVER
	PLANE_MODE_QLOITER       PLANE_MODE = ardupilotmega.PLANE_MODE_QLOITER
	PLANE_MODE_QLAND         PLANE_MODE = ardupilotmega.PLANE_MODE_QLAND
	PLANE_MODE_QRTL          PLANE_MODE = ardupilotmega.PLANE_MODE_QRTL
	PLANE_MODE_QAUTOTUNE     PLANE_MODE = ardupilotmega.PLANE_MODE_QAUTOTUNE
	PLANE_MODE_QACRO         PLANE_MODE = ardupilotmega.PLANE_MODE_QACRO
	PLANE_MODE_THERMAL       PLANE_MODE = ardupilotmega.PLANE_MODE_THERMAL
)
