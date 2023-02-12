//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Actions for reading and writing plan information (mission, rally points, geofence) between persistent and volatile storage when using MAV_CMD_PREFLIGHT_STORAGE.
// (Commonly missions are loaded from persistent storage (flash/EEPROM) into volatile storage (RAM) on startup and written back when they are changed.)
type PREFLIGHT_STORAGE_MISSION_ACTION = common.PREFLIGHT_STORAGE_MISSION_ACTION

const (
	// Read current mission data from persistent storage
	MISSION_READ_PERSISTENT PREFLIGHT_STORAGE_MISSION_ACTION = common.MISSION_READ_PERSISTENT
	// Write current mission data to persistent storage
	MISSION_WRITE_PERSISTENT PREFLIGHT_STORAGE_MISSION_ACTION = common.MISSION_WRITE_PERSISTENT
	// Erase all mission data stored on the vehicle (both persistent and volatile storage)
	MISSION_RESET_DEFAULT PREFLIGHT_STORAGE_MISSION_ACTION = common.MISSION_RESET_DEFAULT
)
