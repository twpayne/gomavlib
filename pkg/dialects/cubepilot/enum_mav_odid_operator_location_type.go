//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

type MAV_ODID_OPERATOR_LOCATION_TYPE = common.MAV_ODID_OPERATOR_LOCATION_TYPE

const (
	// The location/altitude of the operator is the same as the take-off location.
	MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF MAV_ODID_OPERATOR_LOCATION_TYPE = common.MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF
	// The location/altitude of the operator is dynamic. E.g. based on live GNSS data.
	MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS MAV_ODID_OPERATOR_LOCATION_TYPE = common.MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS
	// The location/altitude of the operator are fixed values.
	MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED MAV_ODID_OPERATOR_LOCATION_TYPE = common.MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED
)
