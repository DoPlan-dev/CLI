package integration

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// SetupIDE sets up IDE integration
func SetupIDE(projectRoot, ideName string) error {
	ide, ok := supportedIDEs[ideName]
	if !ok {
		return createGenericGuide(projectRoot)
	}

	return ide.SetupFunc(projectRoot)
}

// VerifyIDE verifies IDE integration
func VerifyIDE(projectRoot, ideName string) error {
	ide, ok := supportedIDEs[ideName]
	if !ok {
		return verifyGenericIntegration(projectRoot)
	}

	return ide.VerifyFunc(projectRoot)
}

// SetupCursor sets up Cursor integration
func SetupCursor(projectRoot string) error {
	cursorDir := filepath.Join(projectRoot, ".cursor")
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

	// Create Cursor-specific directories
	cursorDirs := []string{
		filepath.Join(cursorDir, "agents"),
		filepath.Join(cursorDir, "rules"),
		filepath.Join(cursorDir, "commands"),
		filepath.Join(cursorDir, "config"),
	}

	for _, dir := range cursorDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create symlinks from .cursor/ to .doplan/ai/
	symlinks := map[string]string{
		filepath.Join(cursorDir, "agents"):   filepath.Join(doplanAIDir, "agents"),
		filepath.Join(cursorDir, "rules"):     filepath.Join(doplanAIDir, "rules"),
		filepath.Join(cursorDir, "commands"): filepath.Join(doplanAIDir, "commands"),
	}

	for link, target := range symlinks {
		// Remove existing link/directory if it exists
		if err := os.RemoveAll(link); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove %s: %w", link, err)
		}

		// Ensure target exists before creating symlink
		if _, err := os.Stat(target); os.IsNotExist(err) {
			// Target doesn't exist yet, create empty directory
			if err := os.MkdirAll(target, 0755); err != nil {
				return fmt.Errorf("failed to create target %s: %w", target, err)
			}
		}

		if runtime.GOOS == "windows" {
			// Windows doesn't support symlinks well, copy instead
			if err := copyDir(target, link); err != nil {
				return fmt.Errorf("failed to copy %s to %s: %w", target, link, err)
			}
		} else {
			// Try to create symlink
			if err := os.Symlink(target, link); err != nil {
				// Fallback to copy if symlink fails
				if err := copyDir(target, link); err != nil {
					return fmt.Errorf("failed to create symlink or copy %s to %s: %w", target, link, err)
				}
			}
		}
	}

	// Create Cursor-specific rules file
	rulesFile := filepath.Join(cursorDir, "rules", "doplan.mdc")
	rulesContent := `# DoPlan Workflow Rules

You are working on a project managed by DoPlan CLI.

## Project Structure
- All phases are in ` + "`doplan/##-phase-name/`" + `
- All features are in ` + "`doplan/##-phase-name/##-feature-name/`" + `
- Documentation is in ` + "`doplan/`" + ` root

## Workflow
1. Always read the phase plan before implementing features
2. Follow the feature design specifications
3. Update progress.json after completing tasks
4. Use the DoPlan commands for project management

## Rules Location
All workflow rules are in ` + "`.doplan/ai/rules/`" + `:
- ` + "`workflow.mdc`" + ` - Main workflow rules
- ` + "`communication.mdc`" + ` - Communication rules
- Additional rule files for specific areas

## Agents
AI agents are defined in ` + "`.doplan/ai/agents/`" + `:
- ` + "`planner.agent.md`" + ` - Planning agent
- ` + "`coder.agent.md`" + ` - Implementation agent
- ` + "`designer.agent.md`" + ` - Design agent
- ` + "`reviewer.agent.md`" + ` - Review agent
- ` + "`tester.agent.md`" + ` - Testing agent
- ` + "`devops.agent.md`" + ` - DevOps agent

## Commands
DoPlan commands are available in ` + "`.cursor/commands/`" + `:
- ` + "`/discuss`" + ` - Start idea discussion
- ` + "`/refine`" + ` - Refine idea
- ` + "`/generate`" + ` - Generate PRD and contracts
- ` + "`/plan`" + ` - Create phase structure
- ` + "`/dashboard`" + ` - Show dashboard
- ` + "`/implement`" + ` - Start feature implementation
- ` + "`/next`" + ` - Get next action recommendation
- ` + "`/progress`" + ` - Update progress tracking
`

	if err := os.WriteFile(rulesFile, []byte(rulesContent), 0644); err != nil {
		return fmt.Errorf("failed to create Cursor rules file: %w", err)
	}

	return nil
}

// VerifyCursor verifies Cursor integration
func VerifyCursor(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".cursor", "agents"),
		filepath.Join(projectRoot, ".cursor", "rules"),
		filepath.Join(projectRoot, ".cursor", "commands"),
		filepath.Join(projectRoot, ".cursor", "rules", "doplan.mdc"),
		filepath.Join(projectRoot, ".doplan", "ai", "agents"),
		filepath.Join(projectRoot, ".doplan", "ai", "rules"),
		filepath.Join(projectRoot, ".doplan", "ai", "commands"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	// Verify symlinks/copies are accessible
	agentsLink := filepath.Join(projectRoot, ".cursor", "agents")
	if entries, err := os.ReadDir(agentsLink); err != nil {
		return fmt.Errorf("cannot read .cursor/agents: %w", err)
	} else if len(entries) == 0 {
		// Empty is okay, but warn
	}

	return nil
}

// copyDir recursively copies a directory
func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			if err := os.WriteFile(dstPath, data, entry.Type().Perm()); err != nil {
				return err
			}
		}
	}

	return nil
}

// createGenericGuide creates a generic IDE setup guide
func createGenericGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "generic_ide_setup.md")
	content := `# Generic IDE Integration Guide

## Overview
DoPlan stores all AI-related files in .doplan/ai/ directory.

## Directory Structure
` + "```" + `
.doplan/ai/
├── agents/      # AI agent definitions
├── rules/       # Workflow and design rules
└── commands/    # Command definitions
` + "```" + `

## Integration Steps
1. Locate IDE configuration directory
2. Point to DoPlan files
3. Use agents in IDE's AI chat
4. Follow workflow rules
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

// verifyGenericIntegration verifies generic IDE integration
func verifyGenericIntegration(projectRoot string) error {
	guidePath := filepath.Join(projectRoot, ".doplan", "guides", "generic_ide_setup.md")
	if _, err := os.Stat(guidePath); err != nil {
		return fmt.Errorf("generic guide not found")
	}
	return nil
}

// IDE represents an IDE integration
type IDE struct {
	Name       string
	ConfigDir  string
	SetupFunc  func(string) error
	VerifyFunc func(string) error
}

// supportedIDEs maps IDE names to their setup functions
var supportedIDEs = map[string]*IDE{
	"cursor": {
		Name:       "Cursor",
		ConfigDir:  ".cursor",
		SetupFunc:  SetupCursor,
		VerifyFunc: VerifyCursor,
	},
	"vscode": {
		Name:       "VS Code",
		ConfigDir:  ".vscode",
		SetupFunc:  SetupVSCode,
		VerifyFunc: VerifyVSCode,
	},
	"gemini": {
		Name:       "Gemini CLI",
		ConfigDir:  ".gemini",
		SetupFunc:  SetupGemini,
		VerifyFunc: VerifyGemini,
	},
	"claude": {
		Name:       "Claude Code",
		ConfigDir:  ".claude",
		SetupFunc:  SetupClaude,
		VerifyFunc: VerifyClaude,
	},
	"codex": {
		Name:       "Codex",
		ConfigDir:  ".codex",
		SetupFunc:  SetupCodex,
		VerifyFunc: VerifyCodex,
	},
	"opencode": {
		Name:       "OpenCode",
		ConfigDir:  ".opencode",
		SetupFunc:  SetupOpenCode,
		VerifyFunc: VerifyOpenCode,
	},
	"qwen": {
		Name:       "Qwen Code",
		ConfigDir:  ".qwen",
		SetupFunc:  SetupQwen,
		VerifyFunc: VerifyQwen,
	},
	"kiro": {
		Name:       "Kiro",
		ConfigDir:  ".kiro",
		SetupFunc:  SetupKiro,
		VerifyFunc: VerifyKiro,
	},
	"windsurf": {
		Name:       "Windsurf",
		ConfigDir:  ".windsurf",
		SetupFunc:  SetupWindsurf,
		VerifyFunc: VerifyWindsurf,
	},
	"qoder": {
		Name:       "Qoder",
		ConfigDir:  ".qoder",
		SetupFunc:  SetupQoder,
		VerifyFunc: VerifyQoder,
	},
}

