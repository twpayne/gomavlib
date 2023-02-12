//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Streamed from drone to report progress of terrain map download (initiated by TERRAIN_REQUEST), or sent as a response to a TERRAIN_CHECK request. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type MessageTerrainReport = common.MessageTerrainReport
