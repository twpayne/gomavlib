//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type LIMIT_MODULE uint32

const (
	// Pre-initialization.
	LIMIT_GPSLOCK LIMIT_MODULE = 1
	// Disabled.
	LIMIT_GEOFENCE LIMIT_MODULE = 2
	// Checking limits.
	LIMIT_ALTITUDE LIMIT_MODULE = 4
)

var labels_LIMIT_MODULE = map[LIMIT_MODULE]string{
	LIMIT_GPSLOCK:  "LIMIT_GPSLOCK",
	LIMIT_GEOFENCE: "LIMIT_GEOFENCE",
	LIMIT_ALTITUDE: "LIMIT_ALTITUDE",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e LIMIT_MODULE) MarshalText() ([]byte, error) {
	if l, ok := labels_LIMIT_MODULE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_LIMIT_MODULE = map[string]LIMIT_MODULE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *LIMIT_MODULE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_LIMIT_MODULE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e LIMIT_MODULE) String() string {
	if l, ok := labels_LIMIT_MODULE[e]; ok {
		return l
	}
	return "invalid value"
}
