package generator

import (
	"fmt"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// IDEGenerator generates IDE-specific configuration files
type IDEGenerator struct{}

// Name returns the name of the generator
func (g *IDEGenerator) Name() string {
	return "IDE Configs"
}

// Generate creates IDE configuration files
func (g *IDEGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// Generate CLAUDE.md for all projects so Claude users have quick-start instructions
	if err := generateClaudeConfig(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate CLAUDE.md: %w", err)
	}
	return nil
}

// generateClaudeConfig generates CLAUDE.md file in docs/ directory
func generateClaudeConfig(projectPath string, request *models.ProjectRequest) error {
	content := `# Claude Code Project Configuration

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .cursor/agents/

## Commands
All commands are defined in .cursor/commands/. Type any command (e.g., /tell, /write) to activate it.

## Rules
Stack-specific rules are in .cursor/rules/

## Project State
Current project state is tracked in .plan/active_state.json

## Usage
1. Type /tell to capture your idea
2. Type /improve to brainstorm
3. Type /write to generate documents
4. Type /good to approve
5. Type /plan to generate the execution plan
6. Type /build to start coding

For full command list, see README.md
`
	// Create docs directory if it doesn't exist
	docsDir := filepath.Join(projectPath, "docs")
	if err := utils.CreateDirectory(docsDir); err != nil {
		return fmt.Errorf("failed to create docs directory: %w", err)
	}

	path := filepath.Join(docsDir, "CLAUDE.md")
	return utils.WriteFile(path, []byte(content))
}

// GenerateIDEConfigs is a convenience function that creates an IDEGenerator and generates IDE configs
func GenerateIDEConfigs(request *models.ProjectRequest, projectPath string) error {
	generator := &IDEGenerator{}
	return generator.Generate(request, projectPath)
}
