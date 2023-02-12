//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Bitmask of (optional) autopilot capabilities (64 bit). If a bit is set, the autopilot supports this capability.
type MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY

const (
	// Autopilot supports the MISSION_ITEM float message type.
	// Note that MISSION_ITEM is deprecated, and autopilots should use MISSION_INT instead.
	MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT
	// Autopilot supports the new param float message type.
	MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT
	// Autopilot supports MISSION_ITEM_INT scaled integer message type.
	// Note that this flag must always be set if missions are supported, because missions must always use MISSION_ITEM_INT (rather than MISSION_ITEM, which is deprecated).
	MAV_PROTOCOL_CAPABILITY_MISSION_INT MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_MISSION_INT
	// Autopilot supports COMMAND_INT scaled integer message type.
	MAV_PROTOCOL_CAPABILITY_COMMAND_INT MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_COMMAND_INT
	// Parameter protocol uses byte-wise encoding of parameter values into param_value (float) fields: https://mavlink.io/en/services/parameter.html#parameter-encoding.
	// Note that either this flag or MAV_PROTOCOL_CAPABILITY_PARAM_ENCODE_C_CAST should be set if the parameter protocol is supported.
	MAV_PROTOCOL_CAPABILITY_PARAM_ENCODE_BYTEWISE MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_PARAM_ENCODE_BYTEWISE
	// Autopilot supports the File Transfer Protocol v1: https://mavlink.io/en/services/ftp.html.
	MAV_PROTOCOL_CAPABILITY_FTP MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_FTP
	// Autopilot supports commanding attitude offboard.
	MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET
	// Autopilot supports commanding position and velocity targets in local NED frame.
	MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED
	// Autopilot supports commanding position and velocity targets in global scaled integers.
	MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT
	// Autopilot supports terrain protocol / data handling.
	MAV_PROTOCOL_CAPABILITY_TERRAIN MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_TERRAIN
	// Autopilot supports direct actuator control.
	MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET
	// Autopilot supports the MAV_CMD_DO_FLIGHTTERMINATION command (flight termination).
	MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION
	// Autopilot supports onboard compass calibration.
	MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION
	// Autopilot supports MAVLink version 2.
	MAV_PROTOCOL_CAPABILITY_MAVLINK2 MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_MAVLINK2
	// Autopilot supports mission fence protocol.
	MAV_PROTOCOL_CAPABILITY_MISSION_FENCE MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_MISSION_FENCE
	// Autopilot supports mission rally point protocol.
	MAV_PROTOCOL_CAPABILITY_MISSION_RALLY MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_MISSION_RALLY
	// Reserved for future use.
	MAV_PROTOCOL_CAPABILITY_RESERVED2 MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_RESERVED2
	// Parameter protocol uses C-cast of parameter values to set the param_value (float) fields: https://mavlink.io/en/services/parameter.html#parameter-encoding.
	// Note that either this flag or MAV_PROTOCOL_CAPABILITY_PARAM_ENCODE_BYTEWISE should be set if the parameter protocol is supported.
	MAV_PROTOCOL_CAPABILITY_PARAM_ENCODE_C_CAST MAV_PROTOCOL_CAPABILITY = common.MAV_PROTOCOL_CAPABILITY_PARAM_ENCODE_C_CAST
)
