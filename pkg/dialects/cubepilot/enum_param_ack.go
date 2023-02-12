//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Result from PARAM_EXT_SET message (or a PARAM_SET within a transaction).
type PARAM_ACK = common.PARAM_ACK

const (
	// Parameter value ACCEPTED and SET
	PARAM_ACK_ACCEPTED PARAM_ACK = common.PARAM_ACK_ACCEPTED
	// Parameter value UNKNOWN/UNSUPPORTED
	PARAM_ACK_VALUE_UNSUPPORTED PARAM_ACK = common.PARAM_ACK_VALUE_UNSUPPORTED
	// Parameter failed to set
	PARAM_ACK_FAILED PARAM_ACK = common.PARAM_ACK_FAILED
	// Parameter value received but not yet set/accepted. A subsequent PARAM_ACK_TRANSACTION or PARAM_EXT_ACK with the final result will follow once operation is completed. This is returned immediately for parameters that take longer to set, indicating that the the parameter was received and does not need to be resent.
	PARAM_ACK_IN_PROGRESS PARAM_ACK = common.PARAM_ACK_IN_PROGRESS
)
