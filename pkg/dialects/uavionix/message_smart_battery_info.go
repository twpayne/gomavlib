//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Smart Battery information (static/infrequent update). Use for updates from: smart battery to flight stack, flight stack to GCS. Use BATTERY_STATUS for smart battery frequent updates.
type MessageSmartBatteryInfo = common.MessageSmartBatteryInfo
