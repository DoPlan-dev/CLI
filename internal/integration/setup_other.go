package integration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// SetupKiro sets up Kiro integration
func SetupKiro(projectRoot string) error {
	kiroDir := filepath.Join(projectRoot, ".kiro")
	kiroConfig := filepath.Join(projectRoot, "kiro.config.json")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create setup guide
	if err := createKiroGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create Kiro guide: %w", err)
	}

	// If .kiro directory exists, copy files
	if _, err := os.Stat(kiroDir); err == nil {
		// Copy agents, rules, and commands to .kiro/
		if err := copyToKiro(kiroDir, doplanAIDir); err != nil {
			return fmt.Errorf("failed to copy files to .kiro: %w", err)
		}
	}

	// If kiro.config.json exists, update it
	if _, err := os.Stat(kiroConfig); err == nil {
		// Note: We don't modify existing config, just document it
	}

	return nil
}

// VerifyKiro verifies Kiro integration
func VerifyKiro(projectRoot string) error {
	guidePath := filepath.Join(projectRoot, ".doplan", "guides", "kiro_setup.md")
	if _, err := os.Stat(guidePath); err != nil {
		return fmt.Errorf("Kiro guide not found: %w", err)
	}

	return nil
}

// SetupWindsurf sets up Windsurf integration
func SetupWindsurf(projectRoot string) error {
	windsurfDir := filepath.Join(projectRoot, ".windsurf")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create Windsurf directories
	windsurfDirs := []string{
		filepath.Join(windsurfDir, "agents"),
		filepath.Join(windsurfDir, "rules"),
		filepath.Join(windsurfDir, "commands"),
	}

	for _, dir := range windsurfDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create symlinks from .windsurf/ to .doplan/ai/
	symlinks := map[string]string{
		filepath.Join(windsurfDir, "agents"):   filepath.Join(doplanAIDir, "agents"),
		filepath.Join(windsurfDir, "rules"):     filepath.Join(doplanAIDir, "rules"),
		filepath.Join(windsurfDir, "commands"): filepath.Join(doplanAIDir, "commands"),
	}

	for link, target := range symlinks {
		// Remove existing link/directory if it exists
		if err := os.RemoveAll(link); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove %s: %w", link, err)
		}

		// Ensure target exists before creating symlink
		if _, err := os.Stat(target); os.IsNotExist(err) {
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

	return nil
}

// VerifyWindsurf verifies Windsurf integration
func VerifyWindsurf(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".windsurf", "agents"),
		filepath.Join(projectRoot, ".windsurf", "rules"),
		filepath.Join(projectRoot, ".windsurf", "commands"),
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

// SetupQoder sets up Qoder integration
func SetupQoder(projectRoot string) error {
	qoderDir := filepath.Join(projectRoot, ".qoder")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create .qoder directory
	if err := os.MkdirAll(qoderDir, 0755); err != nil {
		return err
	}

	// Create doplan.json config file
	config := map[string]interface{}{
		"doplan": map[string]interface{}{
			"agents":  filepath.Join(doplanAIDir, "agents"),
			"rules":   filepath.Join(doplanAIDir, "rules"),
			"commands": filepath.Join(doplanAIDir, "commands"),
		},
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	configPath := filepath.Join(qoderDir, "doplan.json")
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	// Create usage guide
	if err := createQoderGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create Qoder guide: %w", err)
	}

	return nil
}

// VerifyQoder verifies Qoder integration
func VerifyQoder(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".qoder", "doplan.json"),
		filepath.Join(projectRoot, ".doplan", "guides", "qoder_setup.md"),
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

func copyToKiro(kiroDir, doplanAIDir string) error {
	// Copy agents
	agentsSrc := filepath.Join(doplanAIDir, "agents")
	agentsDst := filepath.Join(kiroDir, "agents")
	if _, err := os.Stat(agentsSrc); err == nil {
		if err := copyDir(agentsSrc, agentsDst); err != nil {
			return fmt.Errorf("failed to copy agents: %w", err)
		}
	}

	// Copy rules
	rulesSrc := filepath.Join(doplanAIDir, "rules")
	rulesDst := filepath.Join(kiroDir, "rules")
	if _, err := os.Stat(rulesSrc); err == nil {
		if err := copyDir(rulesSrc, rulesDst); err != nil {
			return fmt.Errorf("failed to copy rules: %w", err)
		}
	}

	// Copy commands
	commandsSrc := filepath.Join(doplanAIDir, "commands")
	commandsDst := filepath.Join(kiroDir, "commands")
	if _, err := os.Stat(commandsSrc); err == nil {
		if err := copyDir(commandsSrc, commandsDst); err != nil {
			return fmt.Errorf("failed to copy commands: %w", err)
		}
	}

	return nil
}

func createKiroGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "kiro_setup.md")
	backtick := "`"
	content := `# Kiro Integration Guide

## Overview
DoPlan stores all AI-related files in ` + backtick + `.doplan/ai/` + backtick + ` directory. To use DoPlan with Kiro, you need to configure Kiro to read from these directories.

## Manual Setup

### Step 1: Configure Kiro
Configure Kiro to read agent prompts from ` + backtick + `.doplan/ai/agents/` + backtick + ` and rules from ` + backtick + `.doplan/ai/rules/` + backtick + `.

### Step 2: Copy Files (Optional)
If your Kiro configuration requires files in a specific location, you can copy them:
` + "```" + `bash
# Copy agents
cp -r .doplan/ai/agents/* .kiro/agents/

# Copy rules
cp -r .doplan/ai/rules/* .kiro/rules/

# Copy commands
cp -r .doplan/ai/commands/* .kiro/commands/
` + "```" + `

### Step 3: Update Kiro Config
Update your ` + backtick + `kiro.config.json` + backtick + ` to point to DoPlan directories:
` + "```" + `json
{
  "agents": ".doplan/ai/agents",
  "rules": ".doplan/ai/rules",
  "commands": ".doplan/ai/commands"
}
` + "```" + `

## Project Structure
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Commands: ` + backtick + `.doplan/ai/commands/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `

## Using DoPlan with Kiro
1. Reference agents from ` + backtick + `.doplan/ai/agents/` + backtick + `
2. Follow rules from ` + backtick + `.doplan/ai/rules/` + backtick + `
3. Use commands as defined in ` + backtick + `.doplan/ai/commands/` + backtick + `
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

func createQoderGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "qoder_setup.md")
	backtick := "`"
	content := `# Qoder Integration Guide

## Overview
DoPlan integration for Qoder is configured via ` + backtick + `.qoder/doplan.json` + backtick + `.

## Configuration

The ` + backtick + `.qoder/doplan.json` + backtick + ` file points to DoPlan directories:
` + "```" + `json
{
  "doplan": {
    "agents": ".doplan/ai/agents",
    "rules": ".doplan/ai/rules",
    "commands": ".doplan/ai/commands"
  }
}
` + "```" + `

## Usage

1. Qoder will read the configuration from ` + backtick + `.qoder/doplan.json` + backtick + `
2. Agents are available from ` + backtick + `.doplan/ai/agents/` + backtick + `
3. Rules are applied from ` + backtick + `.doplan/ai/rules/` + backtick + `
4. Commands are available from ` + backtick + `.doplan/ai/commands/` + backtick + `

## Project Structure
- Config: ` + backtick + `.qoder/doplan.json` + backtick + `
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Commands: ` + backtick + `.doplan/ai/commands/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

