//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Message reporting the status of a gimbal device. This message should be broadcast by a gimbal device component at a low regular rate (e.g. 5 Hz). The angles encoded in the quaternion and the angular velocities are relative to North if the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set or relative to the vehicle heading if the flag is not set.
type MessageGimbalDeviceAttitudeStatus struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Current gimbal flags set.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation). The frame depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set.
	Q [4]float32
	// X component of angular velocity (positive: rolling to the right). The frame is as for the quaternion. NaN if unknown.
	AngularVelocityX float32
	// Y component of angular velocity (positive: pitching up). The frame is as for the quaternion. NaN if unknown.
	AngularVelocityY float32
	// Z component of angular velocity (positive: yawing to the right). The frame is as for the quaternion. NaN if unknown.
	AngularVelocityZ float32
	// Failure flags (0 for no failure).
	FailureFlags GIMBAL_DEVICE_ERROR_FLAGS `mavenum:"uint32"`
}

// GetID implements the message.Message interface.
func (*MessageGimbalDeviceAttitudeStatus) GetID() uint32 {
	return 285
}
