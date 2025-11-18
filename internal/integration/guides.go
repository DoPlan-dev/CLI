package integration

import (
	"os"
	"path/filepath"
)

// GenerateIntegrationGuides generates integration guides for all supported IDEs
func GenerateIntegrationGuides(projectRoot string) error {
	guidesDir := filepath.Join(projectRoot, ".doplan", "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return err
	}

	// Guides are already created by individual setup functions
	// This function can be used to generate a master guide or update existing ones

	// Create master integration guide
	if err := createMasterGuide(projectRoot, guidesDir); err != nil {
		return err
	}

	return nil
}

func createMasterGuide(projectRoot, guidesDir string) error {
	masterGuidePath := filepath.Join(guidesDir, "IDE_INTEGRATION.md")
	backtick := "`"
	content := `# DoPlan IDE Integration Guide

This guide provides an overview of DoPlan integration with various IDEs and AI tools.

## Supported IDEs

### Cursor
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: See ` + backtick + `.doplan/guides/cursor_setup.md` + backtick + ` (if exists)
- **Location**: ` + backtick + `.cursor/` + backtick + ` directory with symlinks to ` + backtick + `.doplan/ai/` + backtick + `

### VS Code + GitHub Copilot
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: See ` + backtick + `.doplan/guides/vscode_setup.md` + backtick + ` (if exists)
- **Location**: ` + backtick + `.vscode/` + backtick + ` directory with tasks and prompts

### Gemini CLI
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: ` + backtick + `.doplan/guides/gemini_setup.md` + backtick + `
- **Location**: ` + backtick + `.gemini/commands/` + backtick + `

### Claude Code
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: ` + backtick + `.doplan/guides/claude_setup.md` + backtick + `
- **Location**: ` + backtick + `.claude/commands/` + backtick + `

### Codex CLI
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: ` + backtick + `.doplan/guides/codex_setup.md` + backtick + `
- **Location**: ` + backtick + `.codex/prompts/` + backtick + `

### OpenCode
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: ` + backtick + `.doplan/guides/opencode_setup.md` + backtick + `
- **Location**: ` + backtick + `.opencode/command/` + backtick + `

### Qwen Code
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: ` + backtick + `.doplan/guides/qwen_setup.md` + backtick + `
- **Location**: ` + backtick + `.qwen/commands/` + backtick + `

### Kiro
- **Setup**: Manual (see guide)
- **Guide**: ` + backtick + `.doplan/guides/kiro_setup.md` + backtick + `
- **Location**: Manual configuration required

### Windsurf
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Location**: ` + backtick + `.windsurf/` + backtick + ` directory with symlinks to ` + backtick + `.doplan/ai/` + backtick + `

### Qoder
- **Setup**: Automatic via ` + backtick + `doplan install` + backtick + `
- **Guide**: ` + backtick + `.doplan/guides/qoder_setup.md` + backtick + `
- **Location**: ` + backtick + `.qoder/doplan.json` + backtick + ` config file

## Central Storage

All AI-related files are stored in ` + backtick + `.doplan/ai/` + backtick + `:
- **Agents**: ` + backtick + `.doplan/ai/agents/` + backtick + ` - AI agent definitions
- **Rules**: ` + backtick + `.doplan/ai/rules/` + backtick + ` - Workflow and design rules
- **Commands**: ` + backtick + `.doplan/ai/commands/` + backtick + ` - Command definitions

## Integration Methods

### Symlinks (Unix/Mac)
Most IDEs use symlinks from their config directory to ` + backtick + `.doplan/ai/` + backtick + `:
- Cursor: ` + backtick + `.cursor/agents/` + backtick + ` → ` + backtick + `.doplan/ai/agents/` + backtick + `
- Windsurf: ` + backtick + `.windsurf/agents/` + backtick + ` → ` + backtick + `.doplan/ai/agents/` + backtick + `

### File Copy (Windows)
On Windows, files are copied instead of symlinked due to limitations.

### Configuration Files
Some IDEs use configuration files:
- Qoder: ` + backtick + `.qoder/doplan.json` + backtick + ` points to DoPlan directories

### Direct Integration
Some IDEs have commands installed directly:
- Gemini: ` + backtick + `.gemini/commands/` + backtick + `
- Claude: ` + backtick + `.claude/commands/` + backtick + `
- Codex: ` + backtick + `.codex/prompts/` + backtick + `
- OpenCode: ` + backtick + `.opencode/command/` + backtick + `
- Qwen: ` + backtick + `.qwen/commands/` + backtick + `

## Verification

To verify integration, run:
` + "```" + `bash
doplan verify
` + "```" + `

Or check manually:
- Verify required directories exist
- Check symlinks/copies are accessible
- Confirm configuration files are present

## Troubleshooting

### Symlinks Not Working
- On Windows, files are copied instead of symlinked
- On Unix/Mac, ensure you have permissions to create symlinks

### IDE Not Detecting Files
1. Verify file paths are correct
2. Check file permissions
3. Restart IDE
4. Check IDE logs for errors

### Agents Not Loading
1. Verify agent file format matches IDE requirements
2. Check file encoding (UTF-8)
3. Verify file extensions (.md, .mdc, etc.)

## Switching IDEs

To switch to a different IDE:
1. Run ` + backtick + `doplan install` + backtick + ` again
2. Select the new IDE
3. DoPlan will set up the new integration
4. Old integration files remain but won't interfere

## More Information

For IDE-specific setup instructions, see the individual guides in ` + backtick + `.doplan/guides/` + backtick + `.
`

	return os.WriteFile(masterGuidePath, []byte(content), 0644)
}
