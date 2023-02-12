//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Event message. Each new event from a particular component gets a new sequence number. The same message might be sent multiple times if (re-)requested. Most events are broadcast, some can be specific to a target component (as receivers keep track of the sequence for missed events, all events need to be broadcast. Thus we use destination_component instead of target_component).
type MessageEvent = common.MessageEvent
