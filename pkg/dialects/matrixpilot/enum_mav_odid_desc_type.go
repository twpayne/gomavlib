//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package matrixpilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

type MAV_ODID_DESC_TYPE = common.MAV_ODID_DESC_TYPE

const (
	// Optional free-form text description of the purpose of the flight.
	MAV_ODID_DESC_TYPE_TEXT MAV_ODID_DESC_TYPE = common.MAV_ODID_DESC_TYPE_TEXT
	// Optional additional clarification when status == MAV_ODID_STATUS_EMERGENCY.
	MAV_ODID_DESC_TYPE_EMERGENCY MAV_ODID_DESC_TYPE = common.MAV_ODID_DESC_TYPE_EMERGENCY
	// Optional additional clarification when status != MAV_ODID_STATUS_EMERGENCY.
	MAV_ODID_DESC_TYPE_EXTENDED_STATUS MAV_ODID_DESC_TYPE = common.MAV_ODID_DESC_TYPE_EXTENDED_STATUS
)
