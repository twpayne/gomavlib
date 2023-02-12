//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// A forwarded CANFD frame as requested by MAV_CMD_CAN_FORWARD. These are separated from CAN_FRAME as they need different handling (eg. TAO handling)
type MessageCanfdFrame = common.MessageCanfdFrame
