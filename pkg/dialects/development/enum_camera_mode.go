//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Camera Modes.
type CAMERA_MODE = common.CAMERA_MODE

const (
	// Camera is in image/photo capture mode.
	CAMERA_MODE_IMAGE CAMERA_MODE = common.CAMERA_MODE_IMAGE
	// Camera is in video capture mode.
	CAMERA_MODE_VIDEO CAMERA_MODE = common.CAMERA_MODE_VIDEO
	// Camera is in image survey capture mode. It allows for camera controller to do specific settings for surveys.
	CAMERA_MODE_IMAGE_SURVEY CAMERA_MODE = common.CAMERA_MODE_IMAGE_SURVEY
)
