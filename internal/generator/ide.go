package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/doplan/cli/internal/utils"
	"github.com/doplan/cli/pkg/models"
)

// IDEGenerator generates IDE-specific configuration files
type IDEGenerator struct{}

// Name returns the name of the generator
func (g *IDEGenerator) Name() string {
	return "IDE Configs"
}

// Generate creates IDE configuration files based on the selected IDE
func (g *IDEGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// Generate configs based on IDE selection
	switch strings.ToLower(request.IDE) {
	case "cursor":
		if err := generateCursorRules(projectPath, request); err != nil {
			return fmt.Errorf("failed to generate .cursorrules: %w", err)
		}
	case "claude code", "claudecode":
		if err := generateClaudeConfig(projectPath, request); err != nil {
			return fmt.Errorf("failed to generate CLAUDE.md: %w", err)
		}
	case "antigravity", "windsurf", "cline", "opencode":
		// These IDEs may use similar configs, generate .cursorrules as default
		if err := generateCursorRules(projectPath, request); err != nil {
			return fmt.Errorf("failed to generate IDE config: %w", err)
		}
	default:
		// Default to Cursor rules for unknown IDEs
		if err := generateCursorRules(projectPath, request); err != nil {
			return fmt.Errorf("failed to generate .cursorrules: %w", err)
		}
	}

	// Generate CLAUDE.md for all projects (some IDEs support it)
	if err := generateClaudeConfig(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate CLAUDE.md: %w", err)
	}

	return nil
}

// generateCursorRules generates .cursorrules file
func generateCursorRules(projectPath string, request *models.ProjectRequest) error {
	content := `# DoPlan Project Configuration

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
5. Type /tasks to generate tasks
6. Type /build to start coding

For full command list, see README.md
`
	path := filepath.Join(projectPath, ".cursorrules")
	return utils.WriteFile(path, []byte(content))
}

// generateClaudeConfig generates CLAUDE.md file
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
5. Type /tasks to generate tasks
6. Type /build to start coding

For full command list, see README.md
`
	path := filepath.Join(projectPath, "CLAUDE.md")
	return utils.WriteFile(path, []byte(content))
}

// GenerateIDEConfigs is a convenience function that creates an IDEGenerator and generates IDE configs
func GenerateIDEConfigs(request *models.ProjectRequest, projectPath string) error {
	generator := &IDEGenerator{}
	return generator.Generate(request, projectPath)
}

