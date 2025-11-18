package commands

import (
	"github.com/DoPlan-dev/CLI/internal/wizard"
)

// PublishPackage launches the package publishing wizard
func PublishPackage() error {
	return wizard.RunPublishWizard()
}

