//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

type NAV_VTOL_LAND_OPTIONS = common.NAV_VTOL_LAND_OPTIONS

const (
	// Default autopilot landing behaviour.
	NAV_VTOL_LAND_OPTIONS_DEFAULT NAV_VTOL_LAND_OPTIONS = common.NAV_VTOL_LAND_OPTIONS_DEFAULT
	// Descend in fixed wing mode, transitioning to multicopter mode for vertical landing when close to the ground.
	// The fixed wing descent pattern is at the discretion of the vehicle (e.g. transition altitude, loiter direction, radius, and speed, etc.).
	NAV_VTOL_LAND_OPTIONS_FW_DESCENT NAV_VTOL_LAND_OPTIONS = common.NAV_VTOL_LAND_OPTIONS_FW_DESCENT
	// Land in multicopter mode on reaching the landing coordinates (the whole landing is by "hover descent").
	NAV_VTOL_LAND_OPTIONS_HOVER_DESCENT NAV_VTOL_LAND_OPTIONS = common.NAV_VTOL_LAND_OPTIONS_HOVER_DESCENT
)
