package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/generators"
)

func main() {
	// Create a temporary test directory
	testDir := "/tmp/doplan-test-agents"
	if err := os.MkdirAll(testDir, 0755); err != nil {
		fmt.Printf("Error creating test directory: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(testDir)

	fmt.Println("Testing Agent Generation...")
	fmt.Println("===========================")

	// Test agents generator
	agentsGen := generators.NewAgentsGenerator(testDir)
	if err := agentsGen.Generate(); err != nil {
		fmt.Printf("❌ Error generating agents: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Agents generated successfully!")

	// Verify files were created
	agentsDir := filepath.Join(testDir, ".doplan", "ai", "agents")
	expectedFiles := []string{
		"README.md",
		"planner.agent.md",
		"coder.agent.md",
		"designer.agent.md",
		"reviewer.agent.md",
		"tester.agent.md",
		"devops.agent.md",
	}

	fmt.Println("\nVerifying generated files...")
	allFilesExist := true
	for _, filename := range expectedFiles {
		filePath := filepath.Join(agentsDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("❌ Missing file: %s\n", filename)
			allFilesExist = false
		} else {
			// Check file size
			info, _ := os.Stat(filePath)
			fmt.Printf("✅ %s (%d bytes)\n", filename, info.Size())
		}
	}

	if !allFilesExist {
		fmt.Println("\n❌ Some files are missing!")
		os.Exit(1)
	}

	// Test rules generator
	fmt.Println("\nTesting Rules Generation...")
	fmt.Println("===========================")

	rulesGen := generators.NewRulesGenerator(testDir)
	if err := rulesGen.Generate(); err != nil {
		fmt.Printf("❌ Error generating rules: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Rules generated successfully!")

	// Verify workflow.mdc and communication.mdc exist
	rulesDir := filepath.Join(testDir, ".doplan", "ai", "rules")
	expectedRules := []string{
		"workflow.mdc",
		"communication.mdc",
	}

	fmt.Println("\nVerifying generated rules...")
	allRulesExist := true
	for _, filename := range expectedRules {
		filePath := filepath.Join(rulesDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("❌ Missing file: %s\n", filename)
			allRulesExist = false
		} else {
			info, _ := os.Stat(filePath)
			fmt.Printf("✅ %s (%d bytes)\n", filename, info.Size())
		}
	}

	if !allRulesExist {
		fmt.Println("\n❌ Some rules files are missing!")
		os.Exit(1)
	}

	fmt.Println("\n✅ All tests passed!")
	fmt.Println("\nGenerated files:")
	fmt.Printf("  - Agents: %s\n", agentsDir)
	fmt.Printf("  - Rules: %s\n", rulesDir)
}

