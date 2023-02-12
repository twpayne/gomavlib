//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package matrixpilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Zoom types for MAV_CMD_SET_CAMERA_ZOOM
type CAMERA_ZOOM_TYPE = common.CAMERA_ZOOM_TYPE

const (
	// Zoom one step increment (-1 for wide, 1 for tele)
	ZOOM_TYPE_STEP CAMERA_ZOOM_TYPE = common.ZOOM_TYPE_STEP
	// Continuous zoom up/down until stopped (-1 for wide, 1 for tele, 0 to stop zooming)
	ZOOM_TYPE_CONTINUOUS CAMERA_ZOOM_TYPE = common.ZOOM_TYPE_CONTINUOUS
	// Zoom value as proportion of full camera range (a percentage value between 0.0 and 100.0)
	ZOOM_TYPE_RANGE CAMERA_ZOOM_TYPE = common.ZOOM_TYPE_RANGE
	// Zoom value/variable focal length in millimetres. Note that there is no message to get the valid zoom range of the camera, so this can type can only be used for cameras where the zoom range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera)
	ZOOM_TYPE_FOCAL_LENGTH CAMERA_ZOOM_TYPE = common.ZOOM_TYPE_FOCAL_LENGTH
)
