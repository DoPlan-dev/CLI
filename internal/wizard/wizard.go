package wizard

import (
	"os"
)

// RunNewProjectWizard runs the new project creation wizard
func RunNewProjectWizard() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}

	wizard := NewNewProjectWizard(projectRoot)
	return wizard.Run()
}

// RunAdoptProjectWizard runs the project adoption wizard
func RunAdoptProjectWizard() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}

	wizard := NewAdoptProjectWizard(projectRoot)
	return wizard.Run()
}

// RunMigrationWizard runs the migration wizard for old DoPlan projects
// Future enhancement: Implement migration wizard in future phase
func RunMigrationWizard() error {
	// For now, return nil to allow fallback behavior
	return nil
}
