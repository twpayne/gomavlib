//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"github.com/aler9/gomavlib/pkg/dialects/uavionix"
)

// Emergency status encoding
type UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_EMERGENCY_STATUS

const (
	UAVIONIX_ADSB_OUT_NO_EMERGENCY                    UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_NO_EMERGENCY
	UAVIONIX_ADSB_OUT_GENERAL_EMERGENCY               UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_GENERAL_EMERGENCY
	UAVIONIX_ADSB_OUT_LIFEGUARD_EMERGENCY             UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_LIFEGUARD_EMERGENCY
	UAVIONIX_ADSB_OUT_MINIMUM_FUEL_EMERGENCY          UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_MINIMUM_FUEL_EMERGENCY
	UAVIONIX_ADSB_OUT_NO_COMM_EMERGENCY               UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_NO_COMM_EMERGENCY
	UAVIONIX_ADSB_OUT_UNLAWFUL_INTERFERANCE_EMERGENCY UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_UNLAWFUL_INTERFERANCE_EMERGENCY
	UAVIONIX_ADSB_OUT_DOWNED_AIRCRAFT_EMERGENCY       UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_DOWNED_AIRCRAFT_EMERGENCY
	UAVIONIX_ADSB_OUT_RESERVED                        UAVIONIX_ADSB_EMERGENCY_STATUS = uavionix.UAVIONIX_ADSB_OUT_RESERVED
)
