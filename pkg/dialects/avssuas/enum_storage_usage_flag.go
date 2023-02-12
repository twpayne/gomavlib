//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Flags to indicate usage for a particular storage (see STORAGE_INFORMATION.storage_usage and MAV_CMD_SET_STORAGE_USAGE).
type STORAGE_USAGE_FLAG = common.STORAGE_USAGE_FLAG

const (
	// Always set to 1 (indicates STORAGE_INFORMATION.storage_usage is supported).
	STORAGE_USAGE_FLAG_SET STORAGE_USAGE_FLAG = common.STORAGE_USAGE_FLAG_SET
	// Storage for saving photos.
	STORAGE_USAGE_FLAG_PHOTO STORAGE_USAGE_FLAG = common.STORAGE_USAGE_FLAG_PHOTO
	// Storage for saving videos.
	STORAGE_USAGE_FLAG_VIDEO STORAGE_USAGE_FLAG = common.STORAGE_USAGE_FLAG_VIDEO
	// Storage for saving logs.
	STORAGE_USAGE_FLAG_LOGS STORAGE_USAGE_FLAG = common.STORAGE_USAGE_FLAG_LOGS
)
