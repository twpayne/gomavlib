//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Message for transporting "arbitrary" variable-length data from one component to another (broadcast is not forbidden, but discouraged). The encoding of the data is usually extension specific, i.e. determined by the source, and is usually not documented as part of the MAVLink specification.
type MessageTunnel = common.MessageTunnel
