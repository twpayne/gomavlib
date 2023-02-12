//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Possible responses from a CELLULAR_CONFIG message.
type CELLULAR_CONFIG_RESPONSE = common.CELLULAR_CONFIG_RESPONSE

const (
	// Changes accepted.
	CELLULAR_CONFIG_RESPONSE_ACCEPTED CELLULAR_CONFIG_RESPONSE = common.CELLULAR_CONFIG_RESPONSE_ACCEPTED
	// Invalid APN.
	CELLULAR_CONFIG_RESPONSE_APN_ERROR CELLULAR_CONFIG_RESPONSE = common.CELLULAR_CONFIG_RESPONSE_APN_ERROR
	// Invalid PIN.
	CELLULAR_CONFIG_RESPONSE_PIN_ERROR CELLULAR_CONFIG_RESPONSE = common.CELLULAR_CONFIG_RESPONSE_PIN_ERROR
	// Changes rejected.
	CELLULAR_CONFIG_RESPONSE_REJECTED CELLULAR_CONFIG_RESPONSE = common.CELLULAR_CONFIG_RESPONSE_REJECTED
	// PUK is required to unblock SIM card.
	CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED CELLULAR_CONFIG_RESPONSE = common.CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED
)
