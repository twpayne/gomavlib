//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Sent from simulation to autopilot. The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.
type MessageHilRcInputsRaw = common.MessageHilRcInputsRaw
