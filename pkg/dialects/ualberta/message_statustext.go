//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Status text message. These messages are printed in yellow in the COMM console of QGroundControl. WARNING: They consume quite some bandwidth, so use only for important status and error messages. If implemented wisely, these messages are buffered on the MCU and sent only at a limited rate (e.g. 10 Hz).
type MessageStatustext = common.MessageStatustext
