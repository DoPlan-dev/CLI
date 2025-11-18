package integration

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupGemini sets up Gemini CLI integration
func SetupGemini(projectRoot string) error {
	geminiDir := filepath.Join(projectRoot, ".gemini")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create .gemini/commands directory (commands already installed by install command)
	commandsDir := filepath.Join(geminiDir, "commands")
	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		return err
	}

	// Create usage guide
	if err := createGeminiGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create Gemini guide: %w", err)
	}

	return nil
}

// VerifyGemini verifies Gemini integration
func VerifyGemini(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".gemini", "commands"),
		filepath.Join(projectRoot, ".doplan", "guides", "gemini_setup.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	return nil
}

// SetupClaude sets up Claude Code integration
func SetupClaude(projectRoot string) error {
	claudeDir := filepath.Join(projectRoot, ".claude")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create .claude/commands directory (commands already installed by install command)
	commandsDir := filepath.Join(claudeDir, "commands")
	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		return err
	}

	// Create usage guide
	if err := createClaudeGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create Claude guide: %w", err)
	}

	return nil
}

// VerifyClaude verifies Claude integration
func VerifyClaude(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".claude", "commands"),
		filepath.Join(projectRoot, ".doplan", "guides", "claude_setup.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	return nil
}

// SetupCodex sets up Codex CLI integration
func SetupCodex(projectRoot string) error {
	codexDir := filepath.Join(projectRoot, ".codex")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create .codex/prompts directory (commands already installed by install command)
	promptsDir := filepath.Join(codexDir, "prompts")
	if err := os.MkdirAll(promptsDir, 0755); err != nil {
		return err
	}

	// Create usage guide
	if err := createCodexGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create Codex guide: %w", err)
	}

	return nil
}

// VerifyCodex verifies Codex integration
func VerifyCodex(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".codex", "prompts"),
		filepath.Join(projectRoot, ".doplan", "guides", "codex_setup.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	return nil
}

// SetupOpenCode sets up OpenCode integration
func SetupOpenCode(projectRoot string) error {
	opencodeDir := filepath.Join(projectRoot, ".opencode")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create .opencode/command directory (commands already installed by install command)
	commandDir := filepath.Join(opencodeDir, "command")
	if err := os.MkdirAll(commandDir, 0755); err != nil {
		return err
	}

	// Create usage guide
	if err := createOpenCodeGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create OpenCode guide: %w", err)
	}

	return nil
}

// VerifyOpenCode verifies OpenCode integration
func VerifyOpenCode(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".opencode", "command"),
		filepath.Join(projectRoot, ".doplan", "guides", "opencode_setup.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	return nil
}

// SetupQwen sets up Qwen Code integration
func SetupQwen(projectRoot string) error {
	qwenDir := filepath.Join(projectRoot, ".qwen")
	doplanAIDir := filepath.Join(projectRoot, ".doplan", "ai")

	// Ensure .doplan/ai directories exist
	if err := ensureDoplanAIDirs(doplanAIDir); err != nil {
		return err
	}

	// Create .qwen/commands directory (commands already installed by install command)
	commandsDir := filepath.Join(qwenDir, "commands")
	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		return err
	}

	// Create usage guide
	if err := createQwenGuide(projectRoot); err != nil {
		return fmt.Errorf("failed to create Qwen guide: %w", err)
	}

	return nil
}

// VerifyQwen verifies Qwen integration
func VerifyQwen(projectRoot string) error {
	checks := []string{
		filepath.Join(projectRoot, ".qwen", "commands"),
		filepath.Join(projectRoot, ".doplan", "guides", "qwen_setup.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			return fmt.Errorf("%s not found: %w", check, err)
		}
	}

	return nil
}

// Helper function to ensure .doplan/ai directories exist
func ensureDoplanAIDirs(doplanAIDir string) error {
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

	return nil
}

func createGeminiGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "gemini_setup.md")
	backtick := "`"
	content := `# Gemini CLI Integration Guide

## Overview
DoPlan commands are installed in ` + backtick + `.gemini/commands/` + backtick + ` and can be used directly in Gemini CLI.

## Usage

### Available Commands
Type ` + backtick + `/` + backtick + ` in Gemini CLI to see available commands:
- ` + backtick + `/discuss` + backtick + ` - Start idea discussion
- ` + backtick + `/refine` + backtick + ` - Refine idea
- ` + backtick + `/generate` + backtick + ` - Generate PRD and contracts
- ` + backtick + `/plan` + backtick + ` - Create phase structure
- ` + backtick + `/dashboard` + backtick + ` - Show dashboard
- ` + backtick + `/implement` + backtick + ` - Start feature implementation
- ` + backtick + `/next` + backtick + ` - Get next action recommendation
- ` + backtick + `/progress` + backtick + ` - Update progress tracking

### Using Commands with Arguments
Commands support arguments using ` + backtick + `{{args}}` + backtick + `:
` + "```" + `
/discuss "my project idea"
` + "```" + `

### Context Files
Reference context files in commands:
- ` + backtick + `@{CONTEXT.md}` + backtick + ` - Project tech stack
- ` + backtick + `@{doplan/PRD.md}` + backtick + ` - Product requirements
- ` + backtick + `@{doplan/XX-phase/XX-Feature/plan.md}` + backtick + ` - Feature plan

## Project Structure
- Commands: ` + backtick + `.gemini/commands/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `

## Documentation
For more information, see: https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/custom-commands.md
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

func createClaudeGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "claude_setup.md")
	backtick := "`"
	content := `# Claude Code Integration Guide

## Overview
DoPlan commands are installed in ` + backtick + `.claude/commands/` + backtick + ` and can be used directly in Claude Code.

## Usage

### Available Commands
Type ` + backtick + `/` + backtick + ` in Claude Code to see available commands:
- ` + backtick + `/discuss` + backtick + ` - Start idea discussion
- ` + backtick + `/refine` + backtick + ` - Refine idea
- ` + backtick + `/generate` + backtick + ` - Generate PRD and contracts
- ` + backtick + `/plan` + backtick + ` - Create phase structure
- ` + backtick + `/dashboard` + backtick + ` - Show dashboard
- ` + backtick + `/implement` + backtick + ` - Start feature implementation
- ` + backtick + `/next` + backtick + ` - Get next action recommendation
- ` + backtick + `/progress` + backtick + ` - Update progress tracking

### Using Commands with Arguments
Commands support arguments using ` + backtick + `$ARGUMENTS` + backtick + `, ` + backtick + `$1-$9` + backtick + `:
` + "```" + `
/discuss TOPIC="my project idea"
` + "```" + `

### File References
Reference files in commands:
- ` + backtick + `@CONTEXT.md` + backtick + ` - Project tech stack
- ` + backtick + `@doplan/PRD.md` + backtick + ` - Product requirements
- ` + backtick + `@doplan/XX-phase/XX-Feature/plan.md` + backtick + ` - Feature plan

### Shell Commands
Execute shell commands with ` + backtick + `!command` + backtick + `:
` + "```" + `
!git status
!git branch --show-current
` + "```" + `

## Project Structure
- Commands: ` + backtick + `.claude/commands/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `

## Documentation
For more information, see: https://code.claude.com/docs/en/slash-commands
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

func createCodexGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "codex_setup.md")
	backtick := "`"
	content := `# Codex CLI Integration Guide

## Overview
DoPlan prompts are installed in ` + backtick + `.codex/prompts/` + backtick + ` and can be used in Codex CLI.

## Usage

### Available Prompts
Use prompts with ` + backtick + `/prompts:name` + backtick + `:
- ` + backtick + `/prompts:discuss` + backtick + ` - Start idea discussion
- ` + backtick + `/prompts:refine` + backtick + ` - Refine idea
- ` + backtick + `/prompts:generate` + backtick + ` - Generate PRD and contracts
- ` + backtick + `/prompts:plan` + backtick + ` - Create phase structure
- ` + backtick + `/prompts:dashboard` + backtick + ` - Show dashboard
- ` + backtick + `/prompts:implement` + backtick + ` - Start feature implementation
- ` + backtick + `/prompts:next` + backtick + ` - Get next action recommendation
- ` + backtick + `/prompts:progress` + backtick + ` - Update progress tracking

### Using Prompts with Arguments
Prompts support arguments using ` + backtick + `$ARGUMENTS` + backtick + `, ` + backtick + `$KEY` + backtick + `:
` + "```" + `
/prompts:discuss TOPIC="my project idea"
` + "```" + `

## Project Structure
- Prompts: ` + backtick + `.codex/prompts/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `

## Documentation
For more information, see: https://developers.openai.com/codex/guides/slash-commands
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

func createOpenCodeGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "opencode_setup.md")
	backtick := "`"
	content := `# OpenCode Integration Guide

## Overview
DoPlan commands are installed in ` + backtick + `.opencode/command/` + backtick + ` and can be used directly in OpenCode.

## Usage

### Available Commands
Type ` + backtick + `/` + backtick + ` in OpenCode to see available commands:
- ` + backtick + `/discuss` + backtick + ` - Start idea discussion
- ` + backtick + `/refine` + backtick + ` - Refine idea
- ` + backtick + `/generate` + backtick + ` - Generate PRD and contracts
- ` + backtick + `/plan` + backtick + ` - Create phase structure
- ` + backtick + `/dashboard` + backtick + ` - Show dashboard
- ` + backtick + `/implement` + backtick + ` - Start feature implementation
- ` + backtick + `/next` + backtick + ` - Get next action recommendation
- ` + backtick + `/progress` + backtick + ` - Update progress tracking

### Using Commands with Arguments
Commands support arguments using ` + backtick + `$ARGUMENTS` + backtick + `, ` + backtick + `$1-$9` + backtick + `:
` + "```" + `
/discuss "my project idea"
` + "```" + `

### File References
Reference files in commands:
- ` + backtick + `@CONTEXT.md` + backtick + ` - Project tech stack
- ` + backtick + `@doplan/PRD.md` + backtick + ` - Product requirements
- ` + backtick + `@doplan/XX-phase/XX-Feature/plan.md` + backtick + ` - Feature plan

### Shell Commands
Execute shell commands with ` + backtick + `!command` + backtick + `:
` + "```" + `
!git status
!git branch --show-current
` + "```" + `

## Project Structure
- Commands: ` + backtick + `.opencode/command/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `

## Documentation
For more information, see: https://opencode.ai/docs/commands/
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

func createQwenGuide(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	guidePath := filepath.Join(guidesDir, "qwen_setup.md")
	backtick := "`"
	content := `# Qwen Code Integration Guide

## Overview
DoPlan commands are installed in ` + backtick + `.qwen/commands/` + backtick + ` and can be used directly in Qwen Code.

## Usage

### Available Commands
Type ` + backtick + `/` + backtick + ` in Qwen Code to see available commands:
- ` + backtick + `/discuss` + backtick + ` - Start idea discussion
- ` + backtick + `/refine` + backtick + ` - Refine idea
- ` + backtick + `/generate` + backtick + ` - Generate PRD and contracts
- ` + backtick + `/plan` + backtick + ` - Create phase structure
- ` + backtick + `/dashboard` + backtick + ` - Show dashboard
- ` + backtick + `/implement` + backtick + ` - Start feature implementation
- ` + backtick + `/next` + backtick + ` - Get next action recommendation
- ` + backtick + `/progress` + backtick + ` - Update progress tracking

### Namespaced Commands
Commands can be organized in subdirectories:
- ` + backtick + `.qwen/commands/my_category/my_command.toml` + backtick + ` becomes ` + backtick + `/my_category:my_command` + backtick + `

## Project Structure
- Commands: ` + backtick + `.qwen/commands/` + backtick + `
- Rules: ` + backtick + `.doplan/ai/rules/` + backtick + `
- Agents: ` + backtick + `.doplan/ai/agents/` + backtick + `
- Project docs: ` + backtick + `doplan/` + backtick + `

## Documentation
For more information, see Qwen Code documentation.
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}
