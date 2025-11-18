package integration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SetupVSCode sets up VS Code + GitHub Copilot integration
func SetupVSCode(projectRoot string) error {
	vscodeDir := filepath.Join(projectRoot, ".vscode")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	dirs := []string{
		filepath.Join(doplanAIDir, "agents"),
		filepath.Join(doplanAIDir, "rules"),
		filepath.Join(doplanAIDir, "commands"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create .vscode directory
	if err := os.MkdirAll(vscodeDir, 0755); err != nil {
		return err
	}

	// Create tasks.json
	if err := createVSCodeTasks(vscodeDir); err != nil {
		return fmt.Errorf("failed to create tasks.json: %w", err)
	}

	// Create settings.json
	if err := createVSCodeSettings(vscodeDir); err != nil {
		return fmt.Errorf("failed to create settings.json: %w", err)
	}

	// Create prompts directory
	promptsDir := filepath.Join(vscodeDir, "prompts")
	if err := os.MkdirAll(promptsDir, 0755); err != nil {
		return err
	}

	// Copy agent files to prompts with Copilot-specific headers
	if err := createVSCodePrompts(promptsDir, doplanAIDir); err != nil {
		return fmt.Errorf("failed to create prompts: %w", err)
	}

	// Create doplan-context.md
	if err := createVSCodeContext(promptsDir, projectRoot); err != nil {
		return fmt.Errorf("failed to create context file: %w", err)
	}

	return nil
}

// VerifyVSCode verifies VS Code integration
func VerifyVSCode(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".vscode", "tasks.json"),
		filepath.Join(projectRoot, ".vscode", "settings.json"),
		filepath.Join(projectRoot, ".vscode", "prompts"),
		filepath.Join(projectRoot, ".vscode", "prompts", "doplan-context.md"),
		filepath.Join(projectRoot, ".doplan", "ai", "agents"),
		filepath.Join(projectRoot, ".doplan", "ai", "rules"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	return nil
}

func createVSCodeTasks(vscodeDir string) error {
	tasks := map[string]interface{}{
		"version": "2.0.0",
		"tasks": []map[string]interface{}{
			{
				"label":    "DoPlan: Dashboard",
				"type":     "shell",
				"command":  "doplan",
				"args":     []string{"dashboard"},
				"problemMatcher": []string{},
			},
			{
				"label":    "DoPlan: Discuss Idea",
				"type":     "shell",
				"command":  "doplan",
				"args":     []string{"discuss"},
				"problemMatcher": []string{},
			},
			{
				"label":    "DoPlan: Generate Docs",
				"type":     "shell",
				"command":  "doplan",
				"args":     []string{"generate"},
				"problemMatcher": []string{},
			},
			{
				"label":    "DoPlan: Update Progress",
				"type":     "shell",
				"command":  "doplan",
				"args":     []string{"progress"},
				"problemMatcher": []string{},
			},
		},
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(vscodeDir, "tasks.json")
	return os.WriteFile(path, data, 0644)
}

func createVSCodeSettings(vscodeDir string) error {
	settings := map[string]interface{}{
		"github.copilot.enable": map[string]bool{
			"*": true,
		},
		"files.associations": map[string]string{
			"*.mdc": "markdown",
		},
		"files.exclude": map[string]bool{
			".doplan": false,
		},
		"recommendations": []string{
			"github.copilot",
			"github.copilot-chat",
		},
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(vscodeDir, "settings.json")
	return os.WriteFile(path, data, 0644)
}

func createVSCodePrompts(promptsDir, doplanAIDir string) error {
	agentsDir := filepath.Join(doplanAIDir, "agents")
	
	// List of agent files to copy
	agentFiles := []string{
		"planner.agent.md",
		"coder.agent.md",
		"designer.agent.md",
		"reviewer.agent.md",
		"tester.agent.md",
		"devops.agent.md",
	}

	for _, agentFile := range agentFiles {
		srcPath := filepath.Join(agentsDir, agentFile)
		dstPath := filepath.Join(promptsDir, agentFile)

		// Read source file if it exists
		content, err := os.ReadFile(srcPath)
		if err != nil {
			// If file doesn't exist yet, create a placeholder
			content = []byte(fmt.Sprintf("# %s\n\nThis agent will be available after running doplan install.\n", agentFile))
		}

		// Add Copilot-specific header
		copilotContent := fmt.Sprintf("---\n# Copilot Chat Prompt\n# Use this agent in GitHub Copilot Chat\n---\n\n%s", string(content))

		if err := os.WriteFile(dstPath, []byte(copilotContent), 0644); err != nil {
			return fmt.Errorf("failed to create %s: %w", agentFile, err)
		}
	}

	return nil
}

func createVSCodeContext(promptsDir, projectRoot string) error {
	contextPath := filepath.Join(promptsDir, "doplan-context.md")
	contextContent := `# DoPlan Project Context

This file provides context about the DoPlan project structure and workflow.

## Project Structure

- **Phases**: ` + "`doplan/##-phase-name/`" + `
- **Features**: ` + "`doplan/##-phase-name/##-feature-name/`" + `
- **Documentation**: ` + "`doplan/`" + ` root
- **Rules**: ` + "`.doplan/ai/rules/`" + `
- **Agents**: ` + "`.doplan/ai/agents/`" + `
- **Commands**: ` + "`.doplan/ai/commands/`" + `

## Workflow

1. Use DoPlan commands via VS Code tasks or Copilot Chat
2. Follow the workflow defined in ` + "`.doplan/ai/rules/workflow.mdc`" + `
3. Update progress files after completing tasks
4. Use agents from ` + "`.vscode/prompts/`" + ` in Copilot Chat

## Available Tasks

- **DoPlan: Dashboard** - View project dashboard
- **DoPlan: Discuss Idea** - Start idea discussion
- **DoPlan: Generate Docs** - Generate PRD and contracts
- **DoPlan: Update Progress** - Update progress tracking

## Using Agents in Copilot Chat

Reference agent files in ` + "`.vscode/prompts/`" + ` when using GitHub Copilot Chat:
- ` + "`@planner.agent.md`" + ` - For planning tasks
- ` + "`@coder.agent.md`" + ` - For implementation tasks
- ` + "`@designer.agent.md`" + ` - For design tasks
- ` + "`@reviewer.agent.md`" + ` - For code review
- ` + "`@tester.agent.md`" + ` - For testing tasks
- ` + "`@devops.agent.md`" + ` - For deployment tasks

## Rules

All agents must follow:
- ` + "`.doplan/ai/rules/workflow.mdc`" + ` - Main workflow rules
- ` + "`.doplan/ai/rules/communication.mdc`" + ` - Communication rules
`

	return os.WriteFile(contextPath, []byte(contextContent), 0644)
}

