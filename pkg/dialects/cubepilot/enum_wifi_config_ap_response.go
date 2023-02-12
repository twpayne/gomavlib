//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Possible responses from a WIFI_CONFIG_AP message.
type WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE

const (
	// Undefined response. Likely an indicative of a system that doesn't support this request.
	WIFI_CONFIG_AP_RESPONSE_UNDEFINED WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE_UNDEFINED
	// Changes accepted.
	WIFI_CONFIG_AP_RESPONSE_ACCEPTED WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE_ACCEPTED
	// Changes rejected.
	WIFI_CONFIG_AP_RESPONSE_REJECTED WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE_REJECTED
	// Invalid Mode.
	WIFI_CONFIG_AP_RESPONSE_MODE_ERROR WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE_MODE_ERROR
	// Invalid SSID.
	WIFI_CONFIG_AP_RESPONSE_SSID_ERROR WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE_SSID_ERROR
	// Invalid Password.
	WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR WIFI_CONFIG_AP_RESPONSE = common.WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR
)
