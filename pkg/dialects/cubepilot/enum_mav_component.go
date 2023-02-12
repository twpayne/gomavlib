//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/aler9/gomavlib/pkg/dialects/minimal"
)

// Component ids (values) for the different types and instances of onboard hardware/software that might make up a MAVLink system (autopilot, cameras, servos, GPS systems, avoidance systems etc.).
// Components must use the appropriate ID in their source address when sending messages. Components can also use IDs to determine if they are the intended recipient of an incoming message. The MAV_COMP_ID_ALL value is used to indicate messages that must be processed by all components.
// When creating new entries, components that can have multiple instances (e.g. cameras, servos etc.) should be allocated sequential values. An appropriate number of values should be left free after these components to allow the number of instances to be expanded.
type MAV_COMPONENT = minimal.MAV_COMPONENT

const (
	// Target id (target_component) used to broadcast messages to all components of the receiving system. Components should attempt to process messages with this component ID and forward to components on any other interfaces. Note: This is not a valid *source* component id for a message.
	MAV_COMP_ID_ALL MAV_COMPONENT = minimal.MAV_COMP_ID_ALL
	// System flight controller component ("autopilot"). Only one autopilot is expected in a particular system.
	MAV_COMP_ID_AUTOPILOT1 MAV_COMPONENT = minimal.MAV_COMP_ID_AUTOPILOT1
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER1 MAV_COMPONENT = minimal.MAV_COMP_ID_USER1
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER2 MAV_COMPONENT = minimal.MAV_COMP_ID_USER2
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER3 MAV_COMPONENT = minimal.MAV_COMP_ID_USER3
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER4 MAV_COMPONENT = minimal.MAV_COMP_ID_USER4
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER5 MAV_COMPONENT = minimal.MAV_COMP_ID_USER5
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER6 MAV_COMPONENT = minimal.MAV_COMP_ID_USER6
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER7 MAV_COMPONENT = minimal.MAV_COMP_ID_USER7
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER8 MAV_COMPONENT = minimal.MAV_COMP_ID_USER8
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER9 MAV_COMPONENT = minimal.MAV_COMP_ID_USER9
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER10 MAV_COMPONENT = minimal.MAV_COMP_ID_USER10
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER11 MAV_COMPONENT = minimal.MAV_COMP_ID_USER11
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER12 MAV_COMPONENT = minimal.MAV_COMP_ID_USER12
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER13 MAV_COMPONENT = minimal.MAV_COMP_ID_USER13
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER14 MAV_COMPONENT = minimal.MAV_COMP_ID_USER14
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER15 MAV_COMPONENT = minimal.MAV_COMP_ID_USER15
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER16 MAV_COMPONENT = minimal.MAV_COMP_ID_USER16
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER17 MAV_COMPONENT = minimal.MAV_COMP_ID_USER17
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER18 MAV_COMPONENT = minimal.MAV_COMP_ID_USER18
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER19 MAV_COMPONENT = minimal.MAV_COMP_ID_USER19
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER20 MAV_COMPONENT = minimal.MAV_COMP_ID_USER20
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER21 MAV_COMPONENT = minimal.MAV_COMP_ID_USER21
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER22 MAV_COMPONENT = minimal.MAV_COMP_ID_USER22
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER23 MAV_COMPONENT = minimal.MAV_COMP_ID_USER23
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER24 MAV_COMPONENT = minimal.MAV_COMP_ID_USER24
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER25 MAV_COMPONENT = minimal.MAV_COMP_ID_USER25
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER26 MAV_COMPONENT = minimal.MAV_COMP_ID_USER26
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER27 MAV_COMPONENT = minimal.MAV_COMP_ID_USER27
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER28 MAV_COMPONENT = minimal.MAV_COMP_ID_USER28
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER29 MAV_COMPONENT = minimal.MAV_COMP_ID_USER29
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER30 MAV_COMPONENT = minimal.MAV_COMP_ID_USER30
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER31 MAV_COMPONENT = minimal.MAV_COMP_ID_USER31
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER32 MAV_COMPONENT = minimal.MAV_COMP_ID_USER32
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER33 MAV_COMPONENT = minimal.MAV_COMP_ID_USER33
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER34 MAV_COMPONENT = minimal.MAV_COMP_ID_USER34
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER35 MAV_COMPONENT = minimal.MAV_COMP_ID_USER35
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER36 MAV_COMPONENT = minimal.MAV_COMP_ID_USER36
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER37 MAV_COMPONENT = minimal.MAV_COMP_ID_USER37
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER38 MAV_COMPONENT = minimal.MAV_COMP_ID_USER38
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER39 MAV_COMPONENT = minimal.MAV_COMP_ID_USER39
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER40 MAV_COMPONENT = minimal.MAV_COMP_ID_USER40
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER41 MAV_COMPONENT = minimal.MAV_COMP_ID_USER41
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER42 MAV_COMPONENT = minimal.MAV_COMP_ID_USER42
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER43 MAV_COMPONENT = minimal.MAV_COMP_ID_USER43
	// Telemetry radio (e.g. SiK radio, or other component that emits RADIO_STATUS messages).
	MAV_COMP_ID_TELEMETRY_RADIO MAV_COMPONENT = minimal.MAV_COMP_ID_TELEMETRY_RADIO
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER45 MAV_COMPONENT = minimal.MAV_COMP_ID_USER45
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER46 MAV_COMPONENT = minimal.MAV_COMP_ID_USER46
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER47 MAV_COMPONENT = minimal.MAV_COMP_ID_USER47
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER48 MAV_COMPONENT = minimal.MAV_COMP_ID_USER48
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER49 MAV_COMPONENT = minimal.MAV_COMP_ID_USER49
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER50 MAV_COMPONENT = minimal.MAV_COMP_ID_USER50
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER51 MAV_COMPONENT = minimal.MAV_COMP_ID_USER51
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER52 MAV_COMPONENT = minimal.MAV_COMP_ID_USER52
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER53 MAV_COMPONENT = minimal.MAV_COMP_ID_USER53
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER54 MAV_COMPONENT = minimal.MAV_COMP_ID_USER54
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER55 MAV_COMPONENT = minimal.MAV_COMP_ID_USER55
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER56 MAV_COMPONENT = minimal.MAV_COMP_ID_USER56
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER57 MAV_COMPONENT = minimal.MAV_COMP_ID_USER57
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER58 MAV_COMPONENT = minimal.MAV_COMP_ID_USER58
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER59 MAV_COMPONENT = minimal.MAV_COMP_ID_USER59
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER60 MAV_COMPONENT = minimal.MAV_COMP_ID_USER60
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER61 MAV_COMPONENT = minimal.MAV_COMP_ID_USER61
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER62 MAV_COMPONENT = minimal.MAV_COMP_ID_USER62
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER63 MAV_COMPONENT = minimal.MAV_COMP_ID_USER63
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER64 MAV_COMPONENT = minimal.MAV_COMP_ID_USER64
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER65 MAV_COMPONENT = minimal.MAV_COMP_ID_USER65
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER66 MAV_COMPONENT = minimal.MAV_COMP_ID_USER66
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER67 MAV_COMPONENT = minimal.MAV_COMP_ID_USER67
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER68 MAV_COMPONENT = minimal.MAV_COMP_ID_USER68
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER69 MAV_COMPONENT = minimal.MAV_COMP_ID_USER69
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER70 MAV_COMPONENT = minimal.MAV_COMP_ID_USER70
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER71 MAV_COMPONENT = minimal.MAV_COMP_ID_USER71
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER72 MAV_COMPONENT = minimal.MAV_COMP_ID_USER72
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER73 MAV_COMPONENT = minimal.MAV_COMP_ID_USER73
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER74 MAV_COMPONENT = minimal.MAV_COMP_ID_USER74
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER75 MAV_COMPONENT = minimal.MAV_COMP_ID_USER75
	// Camera #1.
	MAV_COMP_ID_CAMERA MAV_COMPONENT = minimal.MAV_COMP_ID_CAMERA
	// Camera #2.
	MAV_COMP_ID_CAMERA2 MAV_COMPONENT = minimal.MAV_COMP_ID_CAMERA2
	// Camera #3.
	MAV_COMP_ID_CAMERA3 MAV_COMPONENT = minimal.MAV_COMP_ID_CAMERA3
	// Camera #4.
	MAV_COMP_ID_CAMERA4 MAV_COMPONENT = minimal.MAV_COMP_ID_CAMERA4
	// Camera #5.
	MAV_COMP_ID_CAMERA5 MAV_COMPONENT = minimal.MAV_COMP_ID_CAMERA5
	// Camera #6.
	MAV_COMP_ID_CAMERA6 MAV_COMPONENT = minimal.MAV_COMP_ID_CAMERA6
	// Servo #1.
	MAV_COMP_ID_SERVO1 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO1
	// Servo #2.
	MAV_COMP_ID_SERVO2 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO2
	// Servo #3.
	MAV_COMP_ID_SERVO3 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO3
	// Servo #4.
	MAV_COMP_ID_SERVO4 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO4
	// Servo #5.
	MAV_COMP_ID_SERVO5 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO5
	// Servo #6.
	MAV_COMP_ID_SERVO6 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO6
	// Servo #7.
	MAV_COMP_ID_SERVO7 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO7
	// Servo #8.
	MAV_COMP_ID_SERVO8 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO8
	// Servo #9.
	MAV_COMP_ID_SERVO9 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO9
	// Servo #10.
	MAV_COMP_ID_SERVO10 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO10
	// Servo #11.
	MAV_COMP_ID_SERVO11 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO11
	// Servo #12.
	MAV_COMP_ID_SERVO12 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO12
	// Servo #13.
	MAV_COMP_ID_SERVO13 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO13
	// Servo #14.
	MAV_COMP_ID_SERVO14 MAV_COMPONENT = minimal.MAV_COMP_ID_SERVO14
	// Gimbal #1.
	MAV_COMP_ID_GIMBAL MAV_COMPONENT = minimal.MAV_COMP_ID_GIMBAL
	// Logging component.
	MAV_COMP_ID_LOG MAV_COMPONENT = minimal.MAV_COMP_ID_LOG
	// Automatic Dependent Surveillance-Broadcast (ADS-B) component.
	MAV_COMP_ID_ADSB MAV_COMPONENT = minimal.MAV_COMP_ID_ADSB
	// On Screen Display (OSD) devices for video links.
	MAV_COMP_ID_OSD MAV_COMPONENT = minimal.MAV_COMP_ID_OSD
	// Generic autopilot peripheral component ID. Meant for devices that do not implement the parameter microservice.
	MAV_COMP_ID_PERIPHERAL MAV_COMPONENT = minimal.MAV_COMP_ID_PERIPHERAL
	// Gimbal ID for QX1.
	MAV_COMP_ID_QX1_GIMBAL MAV_COMPONENT = minimal.MAV_COMP_ID_QX1_GIMBAL
	// FLARM collision alert component.
	MAV_COMP_ID_FLARM MAV_COMPONENT = minimal.MAV_COMP_ID_FLARM
	// Parachute component.
	MAV_COMP_ID_PARACHUTE MAV_COMPONENT = minimal.MAV_COMP_ID_PARACHUTE
	// Winch component.
	MAV_COMP_ID_WINCH MAV_COMPONENT = minimal.MAV_COMP_ID_WINCH
	// Gimbal #2.
	MAV_COMP_ID_GIMBAL2 MAV_COMPONENT = minimal.MAV_COMP_ID_GIMBAL2
	// Gimbal #3.
	MAV_COMP_ID_GIMBAL3 MAV_COMPONENT = minimal.MAV_COMP_ID_GIMBAL3
	// Gimbal #4
	MAV_COMP_ID_GIMBAL4 MAV_COMPONENT = minimal.MAV_COMP_ID_GIMBAL4
	// Gimbal #5.
	MAV_COMP_ID_GIMBAL5 MAV_COMPONENT = minimal.MAV_COMP_ID_GIMBAL5
	// Gimbal #6.
	MAV_COMP_ID_GIMBAL6 MAV_COMPONENT = minimal.MAV_COMP_ID_GIMBAL6
	// Battery #1.
	MAV_COMP_ID_BATTERY MAV_COMPONENT = minimal.MAV_COMP_ID_BATTERY
	// Battery #2.
	MAV_COMP_ID_BATTERY2 MAV_COMPONENT = minimal.MAV_COMP_ID_BATTERY2
	// CAN over MAVLink client.
	MAV_COMP_ID_MAVCAN MAV_COMPONENT = minimal.MAV_COMP_ID_MAVCAN
	// Component that can generate/supply a mission flight plan (e.g. GCS or developer API).
	MAV_COMP_ID_MISSIONPLANNER MAV_COMPONENT = minimal.MAV_COMP_ID_MISSIONPLANNER
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER MAV_COMPONENT = minimal.MAV_COMP_ID_ONBOARD_COMPUTER
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER2 MAV_COMPONENT = minimal.MAV_COMP_ID_ONBOARD_COMPUTER2
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER3 MAV_COMPONENT = minimal.MAV_COMP_ID_ONBOARD_COMPUTER3
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER4 MAV_COMPONENT = minimal.MAV_COMP_ID_ONBOARD_COMPUTER4
	// Component that finds an optimal path between points based on a certain constraint (e.g. minimum snap, shortest path, cost, etc.).
	MAV_COMP_ID_PATHPLANNER MAV_COMPONENT = minimal.MAV_COMP_ID_PATHPLANNER
	// Component that plans a collision free path between two points.
	MAV_COMP_ID_OBSTACLE_AVOIDANCE MAV_COMPONENT = minimal.MAV_COMP_ID_OBSTACLE_AVOIDANCE
	// Component that provides position estimates using VIO techniques.
	MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY MAV_COMPONENT = minimal.MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY
	// Component that manages pairing of vehicle and GCS.
	MAV_COMP_ID_PAIRING_MANAGER MAV_COMPONENT = minimal.MAV_COMP_ID_PAIRING_MANAGER
	// Inertial Measurement Unit (IMU) #1.
	MAV_COMP_ID_IMU MAV_COMPONENT = minimal.MAV_COMP_ID_IMU
	// Inertial Measurement Unit (IMU) #2.
	MAV_COMP_ID_IMU_2 MAV_COMPONENT = minimal.MAV_COMP_ID_IMU_2
	// Inertial Measurement Unit (IMU) #3.
	MAV_COMP_ID_IMU_3 MAV_COMPONENT = minimal.MAV_COMP_ID_IMU_3
	// GPS #1.
	MAV_COMP_ID_GPS MAV_COMPONENT = minimal.MAV_COMP_ID_GPS
	// GPS #2.
	MAV_COMP_ID_GPS2 MAV_COMPONENT = minimal.MAV_COMP_ID_GPS2
	// Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet).
	MAV_COMP_ID_ODID_TXRX_1 MAV_COMPONENT = minimal.MAV_COMP_ID_ODID_TXRX_1
	// Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet).
	MAV_COMP_ID_ODID_TXRX_2 MAV_COMPONENT = minimal.MAV_COMP_ID_ODID_TXRX_2
	// Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet).
	MAV_COMP_ID_ODID_TXRX_3 MAV_COMPONENT = minimal.MAV_COMP_ID_ODID_TXRX_3
	// Component to bridge MAVLink to UDP (i.e. from a UART).
	MAV_COMP_ID_UDP_BRIDGE MAV_COMPONENT = minimal.MAV_COMP_ID_UDP_BRIDGE
	// Component to bridge to UART (i.e. from UDP).
	MAV_COMP_ID_UART_BRIDGE MAV_COMPONENT = minimal.MAV_COMP_ID_UART_BRIDGE
	// Component handling TUNNEL messages (e.g. vendor specific GUI of a component).
	MAV_COMP_ID_TUNNEL_NODE MAV_COMPONENT = minimal.MAV_COMP_ID_TUNNEL_NODE
	// Deprecated, don't use. Component for handling system messages (e.g. to ARM, takeoff, etc.).
	MAV_COMP_ID_SYSTEM_CONTROL MAV_COMPONENT = minimal.MAV_COMP_ID_SYSTEM_CONTROL
)
