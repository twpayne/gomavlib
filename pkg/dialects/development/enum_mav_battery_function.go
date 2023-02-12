//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Enumeration of battery functions
type MAV_BATTERY_FUNCTION = common.MAV_BATTERY_FUNCTION

const (
	// Battery function is unknown
	MAV_BATTERY_FUNCTION_UNKNOWN MAV_BATTERY_FUNCTION = common.MAV_BATTERY_FUNCTION_UNKNOWN
	// Battery supports all flight systems
	MAV_BATTERY_FUNCTION_ALL MAV_BATTERY_FUNCTION = common.MAV_BATTERY_FUNCTION_ALL
	// Battery for the propulsion system
	MAV_BATTERY_FUNCTION_PROPULSION MAV_BATTERY_FUNCTION = common.MAV_BATTERY_FUNCTION_PROPULSION
	// Avionics battery
	MAV_BATTERY_FUNCTION_AVIONICS MAV_BATTERY_FUNCTION = common.MAV_BATTERY_FUNCTION_AVIONICS
	// Payload battery
	MAV_BATTERY_TYPE_PAYLOAD MAV_BATTERY_FUNCTION = common.MAV_BATTERY_TYPE_PAYLOAD
)
