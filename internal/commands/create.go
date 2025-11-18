package commands

import (
	"github.com/DoPlan-dev/CLI/internal/wizard"
)

// CreateNewProject launches the new project wizard
func CreateNewProject() error {
	return wizard.RunNewProjectWizard()
}
