//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// These encode the sensors whose status is sent as part of the SYS_STATUS message in the extended fields.
type MAV_SYS_STATUS_SENSOR_EXTENDED int

const (
	// 0x01 Recovery system (parachute, balloon, retracts etc)
	MAV_SYS_STATUS_RECOVERY_SYSTEM MAV_SYS_STATUS_SENSOR_EXTENDED = 1
)

var labels_MAV_SYS_STATUS_SENSOR_EXTENDED = map[MAV_SYS_STATUS_SENSOR_EXTENDED]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_SYS_STATUS_SENSOR_EXTENDED) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_SYS_STATUS_SENSOR_EXTENDED[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_SYS_STATUS_SENSOR_EXTENDED = map[string]MAV_SYS_STATUS_SENSOR_EXTENDED{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_SYS_STATUS_SENSOR_EXTENDED) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_SYS_STATUS_SENSOR_EXTENDED[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_SYS_STATUS_SENSOR_EXTENDED) String() string {
	if l, ok := labels_MAV_SYS_STATUS_SENSOR_EXTENDED[e]; ok {
		return l
	}
	return "invalid value"
}
