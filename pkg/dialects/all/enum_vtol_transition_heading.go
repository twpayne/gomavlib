//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Direction of VTOL transition
type VTOL_TRANSITION_HEADING = common.VTOL_TRANSITION_HEADING

const (
	// Respect the heading configuration of the vehicle.
	VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT VTOL_TRANSITION_HEADING = common.VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT
	// Use the heading pointing towards the next waypoint.
	VTOL_TRANSITION_HEADING_NEXT_WAYPOINT VTOL_TRANSITION_HEADING = common.VTOL_TRANSITION_HEADING_NEXT_WAYPOINT
	// Use the heading on takeoff (while sitting on the ground).
	VTOL_TRANSITION_HEADING_TAKEOFF VTOL_TRANSITION_HEADING = common.VTOL_TRANSITION_HEADING_TAKEOFF
	// Use the specified heading in parameter 4.
	VTOL_TRANSITION_HEADING_SPECIFIED VTOL_TRANSITION_HEADING = common.VTOL_TRANSITION_HEADING_SPECIFIED
	// Use the current heading when reaching takeoff altitude (potentially facing the wind when weather-vaning is active).
	VTOL_TRANSITION_HEADING_ANY VTOL_TRANSITION_HEADING = common.VTOL_TRANSITION_HEADING_ANY
)
