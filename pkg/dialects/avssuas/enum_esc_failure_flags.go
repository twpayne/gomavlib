//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Flags to report ESC failures.
type ESC_FAILURE_FLAGS = common.ESC_FAILURE_FLAGS

const (
	// No ESC failure.
	ESC_FAILURE_NONE ESC_FAILURE_FLAGS = common.ESC_FAILURE_NONE
	// Over current failure.
	ESC_FAILURE_OVER_CURRENT ESC_FAILURE_FLAGS = common.ESC_FAILURE_OVER_CURRENT
	// Over voltage failure.
	ESC_FAILURE_OVER_VOLTAGE ESC_FAILURE_FLAGS = common.ESC_FAILURE_OVER_VOLTAGE
	// Over temperature failure.
	ESC_FAILURE_OVER_TEMPERATURE ESC_FAILURE_FLAGS = common.ESC_FAILURE_OVER_TEMPERATURE
	// Over RPM failure.
	ESC_FAILURE_OVER_RPM ESC_FAILURE_FLAGS = common.ESC_FAILURE_OVER_RPM
	// Inconsistent command failure i.e. out of bounds.
	ESC_FAILURE_INCONSISTENT_CMD ESC_FAILURE_FLAGS = common.ESC_FAILURE_INCONSISTENT_CMD
	// Motor stuck failure.
	ESC_FAILURE_MOTOR_STUCK ESC_FAILURE_FLAGS = common.ESC_FAILURE_MOTOR_STUCK
	// Generic ESC failure.
	ESC_FAILURE_GENERIC ESC_FAILURE_FLAGS = common.ESC_FAILURE_GENERIC
)
