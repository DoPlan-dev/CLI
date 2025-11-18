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

	// Create directories
	cursorDirs := []string{
		filepath.Join(cursorDir, "agents"),
		filepath.Join(cursorDir, "rules"),
		filepath.Join(cursorDir, "commands"),
	}

	for _, dir := range cursorDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create symlinks
	symlinks := map[string]string{
		filepath.Join(cursorDir, "agents"):   filepath.Join(doplanAIDir, "agents"),
		filepath.Join(cursorDir, "rules"):     filepath.Join(doplanAIDir, "rules"),
		filepath.Join(cursorDir, "commands"): filepath.Join(doplanAIDir, "commands"),
	}

	for link, target := range symlinks {
		if err := os.RemoveAll(link); err != nil {
			return err
		}

		if runtime.GOOS == "windows" {
			// Windows doesn't support symlinks well, copy instead
			if err := copyDir(target, link); err != nil {
				return err
			}
		} else {
			if err := os.Symlink(target, link); err != nil {
				// Fallback to copy
				if err := copyDir(target, link); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// VerifyCursor verifies Cursor integration
func VerifyCursor(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".cursor", "agents"),
		filepath.Join(projectRoot, ".cursor", "rules"),
		filepath.Join(projectRoot, ".cursor", "commands"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found", check)
		}
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
DoPlan stores all AI-related files in \`.doplan/ai/\` directory.

## Directory Structure
\`\`\`
.doplan/ai/
├── agents/      # AI agent definitions
├── rules/       # Workflow and design rules
└── commands/    # Command definitions
\`\`\`

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
	// Add more IDEs here as needed
}

