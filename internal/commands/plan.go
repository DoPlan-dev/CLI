package commands

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/generators"
)

// CreatePlan generates phase and feature structure
func CreatePlan() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	fmt.Println("🗺️  Creating project plan...")
	fmt.Println()

	// Load state
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	// Generate plan structure
	planGenerator := generators.NewPlanGenerator(projectRoot, state)
	if err := planGenerator.Generate(); err != nil {
		return fmt.Errorf("failed to generate plan: %w", err)
	}

	fmt.Println("✅ Project plan created successfully!")
	fmt.Println("📁 Structure created:")
	fmt.Printf("   • %d phases\n", len(state.Phases))
	totalFeatures := 0
	for _, phase := range state.Phases {
		totalFeatures += len(phase.Features)
	}
	fmt.Printf("   • %d features\n", totalFeatures)
	fmt.Println()
	fmt.Println("💡 Next step: Run 'doplan progress' to update progress tracking")

	return nil
}
