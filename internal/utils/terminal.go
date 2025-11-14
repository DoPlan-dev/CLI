package utils

import (
	"os"
	"strings"
)

const (
	envDisableAnimation = "DOPLAN_NO_ANIMATION"
	envForceAnimation   = "DOPLAN_FORCE_ANIMATION"
)

// AnimationsEnabled returns true when CLI animations should run.
func AnimationsEnabled() bool {
	if val := strings.ToLower(os.Getenv(envForceAnimation)); val == "1" || val == "true" {
		return true
	}

	if val := strings.ToLower(os.Getenv(envDisableAnimation)); val == "1" || val == "true" {
		return false
	}

	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}

	return fi.Mode()&os.ModeCharDevice != 0
}
