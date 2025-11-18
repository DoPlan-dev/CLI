package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/context"
	"github.com/DoPlan-dev/CLI/internal/tui"
	"github.com/DoPlan-dev/CLI/internal/wizard"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "doplan",
		Aliases: []string{".", "dash", "d"},
		Short:   "DoPlan - Project Workflow Manager",
		Long: `DoPlan automates your project workflow from idea to deployment.
Combines Spec-Kit and BMAD-METHOD methodologies.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
		RunE:    executeRoot,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// executeRoot is the context-aware root command handler
func executeRoot(cmd *cobra.Command, args []string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	detector := context.NewDetector(projectRoot)
	state, err := detector.DetectProjectState()
	if err != nil {
		return fmt.Errorf("failed to detect project state: %w", err)
	}

	switch state {
	case context.StateEmptyFolder:
		return launchNewProjectWizard()
	case context.StateExistingCodeNoDoPlan:
		return launchAdoptProjectWizard()
	case context.StateOldDoPlanStructure:
		return launchMigrationWizard()
	case context.StateNewDoPlanStructure:
		return tui.Run() // Open dashboard
	case context.StateInsideFeature:
		return showFeatureView(detector)
	case context.StateInsidePhase:
		return showPhaseView(detector)
	default:
		// Fallback to dashboard
		return tui.Run()
	}
}

// launchNewProjectWizard launches the new project creation wizard
func launchNewProjectWizard() error {
	return wizard.RunNewProjectWizard()
}

// launchAdoptProjectWizard launches the project adoption wizard
func launchAdoptProjectWizard() error {
	return wizard.RunAdoptProjectWizard()
}

// launchMigrationWizard launches the migration wizard
func launchMigrationWizard() error {
	return wizard.RunMigrationWizard()
}

// showFeatureView shows the feature-specific view
func showFeatureView(detector *context.Detector) error {
	details, err := detector.DetectContextDetails()
	if err != nil {
		// Fallback to dashboard if context detection fails
		return tui.Run()
	}
	return tui.RunFeatureView(details)
}

// showPhaseView shows the phase-specific view
func showPhaseView(detector *context.Detector) error {
	details, err := detector.DetectContextDetails()
	if err != nil {
		// Fallback to dashboard if context detection fails
		return tui.Run()
	}
	return tui.RunPhaseView(details)
}
