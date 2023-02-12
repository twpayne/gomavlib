//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Request to read the onboard parameter with the param_id string id. Onboard parameters are stored as key[const char*] -> value[float]. This allows to send a parameter to any other component (such as the GCS) without the need of previous knowledge of possible parameter names. Thus the same GCS can store different parameters for different autopilots. See also https://mavlink.io/en/services/parameter.html for a full documentation of QGroundControl and IMU code.
type MessageParamRequestRead = common.MessageParamRequestRead
