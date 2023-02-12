//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Information about a captured image. This is emitted every time a message is captured.
// MAV_CMD_REQUEST_MESSAGE can be used to (re)request this message for a specific sequence number or range of sequence numbers:
// MAV_CMD_REQUEST_MESSAGE.param2 indicates the sequence number the first image to send, or set to -1 to send the message for all sequence numbers.
// MAV_CMD_REQUEST_MESSAGE.param3 is used to specify a range of messages to send:
// set to 0 (default) to send just the the message for the sequence number in param 2,
// set to -1 to send the message for the sequence number in param 2 and all the following sequence numbers,
// set to the sequence number of the final message in the range.
type MessageCameraImageCaptured = common.MessageCameraImageCaptured
