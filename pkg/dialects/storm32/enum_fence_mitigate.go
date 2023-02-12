//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Actions being taken to mitigate/prevent fence breach
type FENCE_MITIGATE = common.FENCE_MITIGATE

const (
	// Unknown
	FENCE_MITIGATE_UNKNOWN FENCE_MITIGATE = common.FENCE_MITIGATE_UNKNOWN
	// No actions being taken
	FENCE_MITIGATE_NONE FENCE_MITIGATE = common.FENCE_MITIGATE_NONE
	// Velocity limiting active to prevent breach
	FENCE_MITIGATE_VEL_LIMIT FENCE_MITIGATE = common.FENCE_MITIGATE_VEL_LIMIT
)
