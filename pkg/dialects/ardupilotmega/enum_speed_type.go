//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type SPEED_TYPE uint32

const (
	SPEED_TYPE_AIRSPEED    SPEED_TYPE = 0
	SPEED_TYPE_GROUNDSPEED SPEED_TYPE = 1
)

var labels_SPEED_TYPE = map[SPEED_TYPE]string{
	SPEED_TYPE_AIRSPEED:    "SPEED_TYPE_AIRSPEED",
	SPEED_TYPE_GROUNDSPEED: "SPEED_TYPE_GROUNDSPEED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SPEED_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_SPEED_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_SPEED_TYPE = map[string]SPEED_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SPEED_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_SPEED_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e SPEED_TYPE) String() string {
	if l, ok := labels_SPEED_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
