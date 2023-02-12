//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Message encoding a command with parameters as scaled integers. Scaling depends on the actual command value. NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandInt = common.MessageCommandInt
