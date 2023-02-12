//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/storm32"
)

// Flags for gimbal manager operation. Used for setting and reporting, unless specified otherwise. If a setting has been accepted by the gimbal manager is reported in the STORM32_GIMBAL_MANAGER_STATUS message.
type MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS

const (
	// 0 = ignore.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_NONE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_NONE
	// Request to set RC input to active, or report RC input is active. Implies RC mixed. RC exclusive is achieved by setting all clients to inactive.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_RC_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_RC_ACTIVE
	// Request to set onboard/companion computer client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_ONBOARD_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_ONBOARD_ACTIVE
	// Request to set autopliot client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_AUTOPILOT_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_AUTOPILOT_ACTIVE
	// Request to set GCS client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS_ACTIVE
	// Request to set camera client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA_ACTIVE
	// Request to set GCS2 client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS2_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS2_ACTIVE
	// Request to set camera2 client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA2_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA2_ACTIVE
	// Request to set custom client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM_ACTIVE
	// Request to set custom2 client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM2_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM2_ACTIVE
	// Request supervision. This flag is only for setting, it is not reported.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_SUPERVISON MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_SUPERVISON
	// Release supervision. This flag is only for setting, it is not reported.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_RELEASE MAV_STORM32_GIMBAL_MANAGER_FLAGS = storm32.MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_RELEASE
)
