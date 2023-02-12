//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// WiFi Mode.
type WIFI_CONFIG_AP_MODE = common.WIFI_CONFIG_AP_MODE

const (
	// WiFi mode is undefined.
	WIFI_CONFIG_AP_MODE_UNDEFINED WIFI_CONFIG_AP_MODE = common.WIFI_CONFIG_AP_MODE_UNDEFINED
	// WiFi configured as an access point.
	WIFI_CONFIG_AP_MODE_AP WIFI_CONFIG_AP_MODE = common.WIFI_CONFIG_AP_MODE_AP
	// WiFi configured as a station connected to an existing local WiFi network.
	WIFI_CONFIG_AP_MODE_STATION WIFI_CONFIG_AP_MODE = common.WIFI_CONFIG_AP_MODE_STATION
	// WiFi disabled.
	WIFI_CONFIG_AP_MODE_DISABLED WIFI_CONFIG_AP_MODE = common.WIFI_CONFIG_AP_MODE_DISABLED
)
