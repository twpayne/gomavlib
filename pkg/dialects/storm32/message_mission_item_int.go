//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Message encoding a mission item. This message is emitted to announce
// the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MessageMissionItemInt = common.MessageMissionItemInt
