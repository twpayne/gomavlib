//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

type MAV_ODID_STATUS = common.MAV_ODID_STATUS

const (
	// The status of the (UA) Unmanned Aircraft is undefined.
	MAV_ODID_STATUS_UNDECLARED MAV_ODID_STATUS = common.MAV_ODID_STATUS_UNDECLARED
	// The UA is on the ground.
	MAV_ODID_STATUS_GROUND MAV_ODID_STATUS = common.MAV_ODID_STATUS_GROUND
	// The UA is in the air.
	MAV_ODID_STATUS_AIRBORNE MAV_ODID_STATUS = common.MAV_ODID_STATUS_AIRBORNE
	// The UA is having an emergency.
	MAV_ODID_STATUS_EMERGENCY MAV_ODID_STATUS = common.MAV_ODID_STATUS_EMERGENCY
	// The remote ID system is failing or unreliable in some way.
	MAV_ODID_STATUS_REMOTE_ID_SYSTEM_FAILURE MAV_ODID_STATUS = common.MAV_ODID_STATUS_REMOTE_ID_SYSTEM_FAILURE
)
