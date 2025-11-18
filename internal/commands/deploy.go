package commands

import (
	"github.com/DoPlan-dev/CLI/internal/wizard"
)

// DeployProject launches the deployment wizard
func DeployProject() error {
	return wizard.RunDeployWizard()
}

