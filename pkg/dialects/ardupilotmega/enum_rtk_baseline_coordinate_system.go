//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// RTK GPS baseline coordinate system, used for RTK corrections
type RTK_BASELINE_COORDINATE_SYSTEM = common.RTK_BASELINE_COORDINATE_SYSTEM

const (
	// Earth-centered, Earth-fixed
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF RTK_BASELINE_COORDINATE_SYSTEM = common.RTK_BASELINE_COORDINATE_SYSTEM_ECEF
	// RTK basestation centered, north, east, down
	RTK_BASELINE_COORDINATE_SYSTEM_NED RTK_BASELINE_COORDINATE_SYSTEM = common.RTK_BASELINE_COORDINATE_SYSTEM_NED
)
