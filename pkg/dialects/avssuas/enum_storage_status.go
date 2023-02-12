//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Flags to indicate the status of camera storage.
type STORAGE_STATUS = common.STORAGE_STATUS

const (
	// Storage is missing (no microSD card loaded for example.)
	STORAGE_STATUS_EMPTY STORAGE_STATUS = common.STORAGE_STATUS_EMPTY
	// Storage present but unformatted.
	STORAGE_STATUS_UNFORMATTED STORAGE_STATUS = common.STORAGE_STATUS_UNFORMATTED
	// Storage present and ready.
	STORAGE_STATUS_READY STORAGE_STATUS = common.STORAGE_STATUS_READY
	// Camera does not supply storage status information. Capacity information in STORAGE_INFORMATION fields will be ignored.
	STORAGE_STATUS_NOT_SUPPORTED STORAGE_STATUS = common.STORAGE_STATUS_NOT_SUPPORTED
)
