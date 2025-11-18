package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/generators"
)

func main() {
	// Create a temporary test directory
	testDir := "/tmp/doplan-test-agents-content"
	if err := os.MkdirAll(testDir, 0755); err != nil {
		fmt.Printf("Error creating test directory: %v\n", err)
		os.Exit(1)
	}
	// Don't remove - keep for inspection

	fmt.Println("Testing Agent Generation with Content Verification...")
	fmt.Println("====================================================\n")

	// Test agents generator
	agentsGen := generators.NewAgentsGenerator(testDir)
	if err := agentsGen.Generate(); err != nil {
		fmt.Printf("❌ Error generating agents: %v\n", err)
		os.Exit(1)
	}

	// Test rules generator
	rulesGen := generators.NewRulesGenerator(testDir)
	if err := rulesGen.Generate(); err != nil {
		fmt.Printf("❌ Error generating rules: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ All files generated successfully!\n")

	// Verify content
	agentsDir := filepath.Join(testDir, ".doplan", "ai", "agents")
	rulesDir := filepath.Join(testDir, ".doplan", "ai", "rules")

	// Check README.md content
	readmePath := filepath.Join(agentsDir, "README.md")
	readmeContent, _ := os.ReadFile(readmePath)
	readmeStr := string(readmeContent)

	checks := []struct {
		name     string
		file     string
		content  string
		expected string
	}{
		{"README has workflow sequence", readmePath, readmeStr, "Perfect Workflow Sequence"},
		{"README has folder structure requirements", readmePath, readmeStr, "numbered and slugified"},
		{"README has screenshot info", readmePath, readmeStr, "screenshots"},
		{"Planner has folder structure", filepath.Join(agentsDir, "planner.agent.md"), "", "numbered and slugified"},
		{"Tester has screenshot capture", filepath.Join(agentsDir, "tester.agent.md"), "", "Screenshot Location"},
		{"Workflow has perfect sequence", filepath.Join(rulesDir, "workflow.mdc"), "", "Perfect Workflow Sequence"},
		{"Communication has handoff protocol", filepath.Join(rulesDir, "communication.mdc"), "", "Agent Handoff Protocol"},
	}

	fmt.Println("Content Verification:")
	fmt.Println("--------------------")
	allPassed := true
	for _, check := range checks {
		if check.content == "" {
			content, err := os.ReadFile(check.file)
			if err != nil {
				fmt.Printf("❌ Cannot read %s: %v\n", check.name, err)
				allPassed = false
				continue
			}
			check.content = string(content)
		}

		if strings.Contains(check.content, check.expected) {
			fmt.Printf("✅ %s\n", check.name)
		} else {
			fmt.Printf("❌ %s - Missing: %s\n", check.name, check.expected)
			allPassed = false
		}
	}

	if !allPassed {
		fmt.Println("\n❌ Some content checks failed!")
		os.Exit(1)
	}

	fmt.Println("\n✅ All content checks passed!")
	fmt.Printf("\nFiles available at:\n")
	fmt.Printf("  - Agents: %s\n", agentsDir)
	fmt.Printf("  - Rules: %s\n", rulesDir)
	fmt.Println("\nYou can inspect the files manually if needed.")
}

