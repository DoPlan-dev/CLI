package commands

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/generators"
)

// GenerateDocuments generates PRD, contracts, and documentation
func GenerateDocuments() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	fmt.Println("📝 Generating documents...")
	fmt.Println()

	// Load state
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	// Generate PRD
	fmt.Println("📄 Generating PRD...")
	prdGenerator := generators.NewPRDGenerator(projectRoot, state)
	if err := prdGenerator.Generate(); err != nil {
		return fmt.Errorf("failed to generate PRD: %w", err)
	}
	fmt.Println("   ✅ PRD.md created")

	// Generate contracts
	fmt.Println("📋 Generating contracts...")
	contractsGenerator := generators.NewContractsGenerator(projectRoot, state)
	if err := contractsGenerator.Generate(); err != nil {
		return fmt.Errorf("failed to generate contracts: %w", err)
	}
	fmt.Println("   ✅ Contracts created")

	fmt.Println()
	fmt.Println("✅ Documents generated successfully!")
	fmt.Println("📁 Files created:")
	fmt.Println("   • doplan/PRD.md")
	fmt.Println("   • doplan/contracts/api-spec.json")
	fmt.Println("   • doplan/contracts/data-model.md")
	fmt.Println()
	fmt.Println("💡 Next step: Run 'doplan plan' to create phase structure")

	return nil
}

