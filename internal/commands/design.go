package commands

import (
	"github.com/DoPlan-dev/CLI/internal/wizard"
)

// ApplyDesign launches the design system (DPR) wizard
func ApplyDesign() error {
	return wizard.RunDesignWizard()
}

