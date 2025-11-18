package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/integration"
	"github.com/DoPlan-dev/CLI/internal/tui"
	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install DoPlan in current project",
		Long:  "Install DoPlan and set up the project structure",
		RunE:  runInstall,
	}

	return cmd
}

func runInstall(cmd *cobra.Command, args []string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return doplanerror.NewIOError("IO001", "Failed to get current directory").WithCause(err)
	}

	// Initialize error handler
	errLogger := doplanerror.NewLogger(projectRoot, doplanerror.LogLevelInfo)
	errHandler := doplanerror.NewHandler(errLogger)

	// Check if already installed
	if config.IsInstalled(projectRoot) {
		confirmed, err := tui.ConfirmReinstall()
		if err != nil {
			return err
		}
		if !confirmed {
			color.Yellow("Installation cancelled.")
			return nil
		}
	}

	// Show installation menu
	ide, err := tui.ShowInstallMenu()
	if err != nil {
		return errHandler.Handle(doplanerror.NewIOError("IO001", "Failed to show installation menu").WithCause(err))
	}

	if ide == "back" {
		color.Yellow("Installation cancelled.")
		return nil
	}

	// Start installation
	color.Blue("\nInstalling DoPlan for %s...\n", ide)

	installer := NewInstaller(projectRoot, ide)
	if err := installer.Install(); err != nil {
		return errHandler.Handle(doplanerror.NewValidationError("VAL001", "Installation failed").WithCause(err))
	}

	color.Green("\n✅ DoPlan installed successfully for %s!", ide)
	color.Yellow("\nYou can now use DoPlan commands in your IDE/CLI.")
	color.Yellow("Run 'doplan dashboard' to view progress, or use slash commands in your IDE.\n")
	
	// Show IDE-specific instructions
	switch ide {
	case "vscode":
		color.Cyan("📝 VS Code: Use tasks from Command Palette (Ctrl+Shift+P) or Copilot Chat")
	case "gemini", "claude", "codex", "opencode", "qwen":
		color.Cyan("📝 %s: Type '/' to see available commands", ide)
	case "cursor", "windsurf":
		color.Cyan("📝 %s: Type '/' in chat to see available commands", ide)
	}
	
	color.Cyan("📚 Integration guide: .doplan/guides/IDE_INTEGRATION.md\n")

	return nil
}

// Installer handles the installation process
type Installer struct {
	projectRoot string
	ide         string
}

// NewInstaller creates a new installer
func NewInstaller(projectRoot, ide string) *Installer {
	return &Installer{
		projectRoot: projectRoot,
		ide:         ide,
	}
}

// Install performs the installation
func (i *Installer) Install() error {
	steps := []struct {
		name string
		fn   func() error
	}{
		{"Creating directory structure", i.createDirectories},
		{"Installing IDE commands", i.installCommands},
		{"Generating AI agents", i.generateAgents},
		{"Generating workflow rules", i.generateRules},
		{"Setting up IDE integration", i.setupIDEIntegration},
		{"Creating templates", i.createTemplates},
		{"Generating tech stack context", i.generateContext},
		{"Generating configuration", i.generateConfig},
		{"Creating README", i.createREADME},
		{"Creating initial dashboard", i.createDashboard},
		{"Verifying installation", i.verifyInstallation},
	}

	useAnimations := utils.AnimationsEnabled()

	for _, step := range steps {
		if useAnimations {
			if err := runInstallerAnimation(step.name, step.fn); err != nil {
				return fmt.Errorf("failed to %s: %w", step.name, err)
			}
			continue
		}

		color.Cyan("✓ %s...\n", step.name)
		if err := step.fn(); err != nil {
			return fmt.Errorf("failed to %s: %w", step.name, err)
		}
	}

	return nil
}

func (i *Installer) createDirectories() error {
	dirs := []string{
		filepath.Join(i.projectRoot, ".cursor", "commands"),
		filepath.Join(i.projectRoot, ".cursor", "rules"),
		filepath.Join(i.projectRoot, ".cursor", "config"),
		filepath.Join(i.projectRoot, "doplan", "contracts"),
		filepath.Join(i.projectRoot, "doplan", "templates"),
	}

	// Create IDE-specific directories
	switch i.ide {
	case "gemini":
		dirs = append(dirs, filepath.Join(i.projectRoot, ".gemini", "commands"))
	case "claude":
		// Claude CLI uses .claude/commands/ for project-specific commands
		// Global commands go in ~/.claude/commands/
		dirs = append(dirs, filepath.Join(i.projectRoot, ".claude", "commands"))
	case "codex":
		// Codex uses ~/.codex/prompts/ for global commands
		// Project-specific commands go in .codex/prompts/ in project root
		dirs = append(dirs, filepath.Join(i.projectRoot, ".codex", "prompts"))
	case "opencode":
		// OpenCode uses .opencode/command/ for project-specific commands
		// Global commands go in ~/.config/opencode/command/
		dirs = append(dirs, filepath.Join(i.projectRoot, ".opencode", "command"))
	case "qwen":
		// Qwen Code uses .qwen/commands/ for project-specific commands
		// Global commands go in ~/.qwen/commands/
		dirs = append(dirs, filepath.Join(i.projectRoot, ".qwen", "commands"))
	}

	for _, dir := range dirs {
		if err := utils.EnsureDir(dir); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) installCommands() error {
	switch i.ide {
	case "cursor":
		return i.installCursorCommands()
	case "gemini":
		return i.installGeminiCommands()
	case "claude":
		return i.installClaudeCommands()
	case "codex":
		return i.installCodexCommands()
	case "opencode":
		return i.installOpenCodeCommands()
	case "qwen":
		return i.installQwenCommands()
	default:
		return i.installCursorCommands() // Default to Cursor
	}
}

func (i *Installer) installCursorCommands() error {
	// Cursor commands are Markdown files
	// See: https://cursor.com/docs/agent/chat/commands
	backtick := "`"
	commands := map[string]string{
		"discuss.md": `# Discuss

## Overview
Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

## Output
- Idea notes document
- Updated state file
- Tech stack recommendations
- Feature organization suggestions
`,
		"refine.md": `# Refine

## Overview
Refine and enhance the current idea. Review existing idea notes and suggest improvements, additional features, and better organization.

## Workflow
1. Review existing idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Review current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Suggest additional features
4. Identify gaps in the plan
5. Enhance technical specifications
6. Update idea documentation

## Focus Areas
- Feature completeness
- Technical feasibility
- User experience improvements
- Architecture enhancements
- Risk mitigation
`,
		"generate.md": `# Generate

## Overview
Generate Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea.

## Workflow
1. Read idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Read state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Generate ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
4. Generate ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
5. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
6. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas
7. Use templates from ` + backtick + `doplan/templates/` + backtick + ` directory

## Documents Created
- PRD.md - Complete product requirements
- structure.md - Project architecture
- api-spec.json - API contracts
- data-model.md - Data models
`,
		"plan.md": `# Plan

## Overview
Generate the project plan with phases and features. Create the directory structure following BMAD-METHOD approach.

## Workflow
1. Read PRD from ` + backtick + `doplan/PRD.md` + backtick + `
2. Read contracts from ` + backtick + `doplan/contracts/` + backtick + `
3. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
4. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
5. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
6. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
7. Update dashboard with new structure

## Structure
` + "```" + `
doplan/
├── 01-phase/
│   ├── phase-plan.md
│   ├── phase-progress.json
│   ├── 01-Feature/
│   │   ├── plan.md
│   │   ├── design.md
│   │   ├── tasks.md
│   │   └── progress.json
│   └── 02-Feature/
└── 02-phase/
` + "```" + `
`,
		"dashboard.md": `# Dashboard

## Overview
Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests.

## Workflow
1. Read all progress.json files from feature directories
2. Calculate overall and phase progress percentages
3. Check GitHub for active PRs (if GitHub integration enabled)
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + ` (markdown)
6. Update ` + backtick + `doplan/dashboard.html` + backtick + ` (visual HTML dashboard)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions

## Usage
After running this command, view the dashboard:
- Markdown: Open ` + backtick + `doplan/dashboard.md` + backtick + `
- HTML: Open ` + backtick + `doplan/dashboard.html` + backtick + ` in browser
- CLI: Run ` + backtick + `doplan dashboard` + backtick + ` in terminal
`,
		"implement.md": `# Implement

## Overview
Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch.

## Workflow
1. Check current feature context from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. **Automatically create GitHub branch:**
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: ` + backtick + `git checkout -b {branch-name}` + backtick + `
3. **Initial commit:**
   - Add plan.md, design.md, tasks.md files
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: ` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on:
   - ` + backtick + `plan.md` + backtick + ` - Feature plan
   - ` + backtick + `design.md` + backtick + ` - Design specifications
   - ` + backtick + `tasks.md` + backtick + ` - Task breakdown

## Implementation Guidance
- Follow the feature's plan.md and design.md
- Check off tasks in tasks.md as you complete them
- Commit regularly with clear messages
- Update progress as you work
`,
		"next.md": `# Next

## Overview
Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.

## Workflow
1. Read current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Scan all feature directories for incomplete tasks
3. Check progress.json files
4. Consider dependencies between features
5. Recommend highest priority action
6. Display recommendation in dashboard format

## Analysis Factors
- Task completion status
- Feature dependencies
- Phase priorities
- Blocked items
- Progress percentages
- GitHub branch status

## Output
- Next recommended action
- Priority level
- Estimated effort
- Dependencies to consider
`,
		"progress.md": `# Progress

## Overview
Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.

## Workflow
1. Scan all feature directories in ` + backtick + `doplan/` + backtick + `
2. Read tasks.md files from each feature
3. Count completed tasks (marked with [x])
4. Calculate completion percentages
5. Update progress.json files:
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `doplan/dashboard.md` + backtick + `
   - ` + backtick + `doplan/dashboard.html` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress
`,
	}

	commandsDir := filepath.Join(i.projectRoot, ".cursor", "commands")

	for filename, content := range commands {
		path := filepath.Join(commandsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) installGeminiCommands() error {
	// Gemini CLI commands are TOML files
	// See: https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/custom-commands.md
	backtick := "`"
	commands := map[string]string{
		"discuss.toml": `description = "Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack."

prompt = """
Start the DoPlan idea discussion workflow. This will help refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

## Output
- Idea notes document
- Updated state file
- Tech stack recommendations
- Feature organization suggestions
"""
`,
		"refine.toml": `description = "Refine and enhance the current idea. Review existing idea notes and suggest improvements, additional features, and better organization."

prompt = """
Review the current idea and suggest improvements, additional features, and better organization. Use the idea notes and current state to provide actionable enhancements.

## Workflow
1. Review existing idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Review current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Suggest additional features
4. Identify gaps in the plan
5. Enhance technical specifications
6. Update idea documentation

## Focus Areas
- Feature completeness
- Technical feasibility
- User experience improvements
- Architecture enhancements
- Risk mitigation
"""
`,
		"generate.toml": `description = "Generate Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea."

prompt = """
Generate the Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea. Create files in doplan/ directory.

## Workflow
1. Read idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Read state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Generate ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
4. Generate ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
5. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
6. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas
7. Use templates from ` + backtick + `doplan/templates/` + backtick + ` directory

## Documents Created
- PRD.md - Complete product requirements
- structure.md - Project architecture
- api-spec.json - API contracts
- data-model.md - Data models
"""
`,
		"plan.toml": `description = "Generate the project plan with phases and features. Create the directory structure following BMAD-METHOD approach."

prompt = """
Generate the project plan with phases and features. Create the directory structure: doplan/01-phase/01-Feature/, etc. Each feature should have plan.md, design.md, and tasks.md files.

## Workflow
1. Read PRD from ` + backtick + `doplan/PRD.md` + backtick + `
2. Read contracts from ` + backtick + `doplan/contracts/` + backtick + `
3. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
4. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
5. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
6. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
7. Update dashboard with new structure
"""
`,
		"dashboard.toml": `description = "Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests."

prompt = """
Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests. Update doplan/dashboard.md file.

## Workflow
1. Read all progress.json files from feature directories
2. Calculate overall and phase progress percentages
3. Check GitHub for active PRs (if GitHub integration enabled)
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + ` (markdown)
6. Update ` + backtick + `doplan/dashboard.html` + backtick + ` (visual HTML dashboard)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions
"""
`,
		"implement.toml": `description = "Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch."

prompt = """
Help implement the current feature. Check the feature's plan.md, design.md, and tasks.md files. Create a GitHub branch for this feature and start implementation.

## Workflow
1. Check current feature context from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Automatically create GitHub branch:
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: ` + backtick + `git checkout -b {branch-name}` + backtick + `
3. Initial commit:
   - Add plan.md, design.md, tasks.md files
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: ` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on:
   - ` + backtick + `plan.md` + backtick + ` - Feature plan
   - ` + backtick + `design.md` + backtick + ` - Design specifications
   - ` + backtick + `tasks.md` + backtick + ` - Task breakdown
"""
`,
		"next.toml": `description = "Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next."

prompt = """
Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.

## Workflow
1. Read current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Scan all feature directories for incomplete tasks
3. Check progress.json files
4. Consider dependencies between features
5. Recommend highest priority action
6. Display recommendation in dashboard format

## Analysis Factors
- Task completion status
- Feature dependencies
- Phase priorities
- Blocked items
- Progress percentages
- GitHub branch status
"""
`,
		"progress.toml": `description = "Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files."

prompt = """
Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.

## Workflow
1. Scan all feature directories in ` + backtick + `doplan/` + backtick + `
2. Read tasks.md files from each feature
3. Count completed tasks (marked with [x])
4. Calculate completion percentages
5. Update progress.json files:
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `doplan/dashboard.md` + backtick + `
   - ` + backtick + `doplan/dashboard.html` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress
"""
`,
	}

	commandsDir := filepath.Join(i.projectRoot, ".gemini", "commands")

	for filename, content := range commands {
		path := filepath.Join(commandsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) installClaudeCommands() error {
	// Claude Code commands are Markdown files with optional YAML frontmatter
	// See: https://code.claude.com/docs/en/slash-commands
	// Commands are invoked with /command-name
	// Supports: $ARGUMENTS, $1-$9, !bash commands, @file references
	backtick := "`"
	commands := map[string]string{
		"discuss.md": `---
description: Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack.
argument-hint: [TOPIC="<idea description>"]
---

Start the DoPlan idea discussion workflow. This will help refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Context
- Current project state: @.cursor/config/doplan-state.json
- Existing idea notes: @doplan/idea-notes.md

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

## Output
- Idea notes document
- Updated state file
- Tech stack recommendations
- Feature organization suggestions

$ARGUMENTS
`,
		"refine.md": `---
description: Refine and enhance the current idea. Review existing idea notes and suggest improvements, additional features, and better organization.
argument-hint: [FOCUS="<specific area to refine>"]
---

Review the current idea and suggest improvements, additional features, and better organization. Use the idea notes and current state to provide actionable enhancements.

## Context
- Existing idea notes: @doplan/idea-notes.md
- Current state: @.cursor/config/doplan-state.json

## Workflow
1. Review existing idea notes
2. Review current state
3. Suggest additional features
4. Identify gaps in the plan
5. Enhance technical specifications
6. Update idea documentation

## Focus Areas
- Feature completeness
- Technical feasibility
- User experience improvements
- Architecture enhancements
- Risk mitigation

$ARGUMENTS
`,
		"generate.md": `---
description: Generate Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea.
argument-hint: [OUTPUT_DIR="<directory>"]
---

Generate the Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea. Create files in doplan/ directory.

## Context
- Idea notes: @doplan/idea-notes.md
- Current state: @.cursor/config/doplan-state.json
- Templates: @doplan/templates/

## Workflow
1. Read idea notes and state
2. Generate ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
3. Generate ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
4. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
5. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas
6. Use templates from ` + backtick + `doplan/templates/` + backtick + ` directory

## Documents Created
- PRD.md - Complete product requirements
- structure.md - Project architecture
- api-spec.json - API contracts
- data-model.md - Data models

$ARGUMENTS
`,
		"plan.md": `---
description: Generate the project plan with phases and features. Create the directory structure following BMAD-METHOD approach.
argument-hint: [PHASES=<number>] [FEATURES_PER_PHASE=<number>]
---

Generate the project plan with phases and features. Create the directory structure: doplan/01-phase/01-Feature/, etc. Each feature should have plan.md, design.md, and tasks.md files.

## Context
- PRD: @doplan/PRD.md
- Contracts: @doplan/contracts/
- Templates: @doplan/templates/

## Workflow
1. Read PRD and contracts
2. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
3. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
4. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
5. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
6. Update dashboard with new structure

$ARGUMENTS
`,
		"dashboard.md": `---
description: Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests.
argument-hint: [FORMAT="<markdown|html|both>"]
allowed-tools: Bash(git branch:*), Bash(git log:*), Bash(git status:*)
---

Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests. Update doplan/dashboard.md file.

## Context
- Current git status: !` + backtick + `git status` + backtick + `
- Current branch: !` + backtick + `git branch --show-current` + backtick + `
- Recent commits: !` + backtick + `git log --oneline -10` + backtick + `

## Workflow
1. Read all progress.json files from feature directories
2. Calculate overall and phase progress percentages
3. Check GitHub for active PRs (if GitHub integration enabled)
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + ` (markdown)
6. Update ` + backtick + `doplan/dashboard.html` + backtick + ` (visual HTML dashboard)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions

$ARGUMENTS
`,
		"implement.md": `---
description: Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch.
argument-hint: [FEATURE="<feature name>"] [BRANCH="<branch name>"]
allowed-tools: Bash(git checkout:*), Bash(git add:*), Bash(git commit:*), Bash(git push:*), Bash(git branch:*)
---

Help implement the current feature. Check the feature's plan.md, design.md, and tasks.md files. Create a GitHub branch for this feature and start implementation.

## Context
- Current state: @.cursor/config/doplan-state.json
- Current git status: !` + backtick + `git status` + backtick + `
- Current branch: !` + backtick + `git branch --show-current` + backtick + `

## Workflow
1. Check current feature context
2. Automatically create GitHub branch:
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: !` + backtick + `git checkout -b {branch-name}` + backtick + `
3. Initial commit:
   - Add plan.md, design.md, tasks.md files
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: !` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on:
   - ` + backtick + `plan.md` + backtick + ` - Feature plan
   - ` + backtick + `design.md` + backtick + ` - Design specifications
   - ` + backtick + `tasks.md` + backtick + ` - Task breakdown

$ARGUMENTS
`,
		"next.md": `---
description: Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.
argument-hint: [PRIORITY="<high|medium|low>"]
---

Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.

## Context
- Current state: @.cursor/config/doplan-state.json
- Dashboard: @doplan/dashboard.md

## Workflow
1. Read current state
2. Scan all feature directories for incomplete tasks
3. Check progress.json files
4. Consider dependencies between features
5. Recommend highest priority action
6. Display recommendation in dashboard format

## Analysis Factors
- Task completion status
- Feature dependencies
- Phase priorities
- Blocked items
- Progress percentages
- GitHub branch status

$ARGUMENTS
`,
		"progress.md": `---
description: Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.
argument-hint: [FEATURE="<feature path>"] [PHASE="<phase number>"]
---

Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.

## Context
- Current state: @.cursor/config/doplan-state.json
- Dashboard: @doplan/dashboard.md

## Workflow
1. Scan all feature directories in ` + backtick + `doplan/` + backtick + `
2. Read tasks.md files from each feature
3. Count completed tasks (marked with [x])
4. Calculate completion percentages
5. Update progress.json files:
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `doplan/dashboard.md` + backtick + `
   - ` + backtick + `doplan/dashboard.html` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress

$ARGUMENTS
`,
	}

	commandsDir := filepath.Join(i.projectRoot, ".claude", "commands")

	for filename, content := range commands {
		path := filepath.Join(commandsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) installCodexCommands() error {
	// Codex CLI commands are Markdown files with YAML frontmatter
	// See: https://developers.openai.com/codex/guides/slash-commands
	// Commands are invoked with /prompts:<name>
	backtick := "`"
	commands := map[string]string{
		"discuss.md": `---
description: Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack.
argument-hint: [TOPIC="<idea description>"]
---

Start the DoPlan idea discussion workflow. This will help refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

## Output
- Idea notes document
- Updated state file
- Tech stack recommendations
- Feature organization suggestions

$ARGUMENTS
`,
		"refine.md": `---
description: Refine and enhance the current idea. Review existing idea notes and suggest improvements, additional features, and better organization.
argument-hint: [FOCUS="<specific area to refine>"]
---

Review the current idea and suggest improvements, additional features, and better organization. Use the idea notes and current state to provide actionable enhancements.

## Workflow
1. Review existing idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Review current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Suggest additional features
4. Identify gaps in the plan
5. Enhance technical specifications
6. Update idea documentation

## Focus Areas
- Feature completeness
- Technical feasibility
- User experience improvements
- Architecture enhancements
- Risk mitigation

$ARGUMENTS
`,
		"generate.md": `---
description: Generate Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea.
argument-hint: [OUTPUT_DIR="<directory>"]
---

Generate the Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea. Create files in doplan/ directory.

## Workflow
1. Read idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Read state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Generate ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
4. Generate ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
5. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
6. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas
7. Use templates from ` + backtick + `doplan/templates/` + backtick + ` directory

## Documents Created
- PRD.md - Complete product requirements
- structure.md - Project architecture
- api-spec.json - API contracts
- data-model.md - Data models

$ARGUMENTS
`,
		"plan.md": `---
description: Generate the project plan with phases and features. Create the directory structure following BMAD-METHOD approach.
argument-hint: [PHASES=<number>] [FEATURES_PER_PHASE=<number>]
---

Generate the project plan with phases and features. Create the directory structure: doplan/01-phase/01-Feature/, etc. Each feature should have plan.md, design.md, and tasks.md files.

## Workflow
1. Read PRD from ` + backtick + `doplan/PRD.md` + backtick + `
2. Read contracts from ` + backtick + `doplan/contracts/` + backtick + `
3. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
4. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
5. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
6. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
7. Update dashboard with new structure

$ARGUMENTS
`,
		"dashboard.md": `---
description: Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests.
argument-hint: [FORMAT="<markdown|html|both>"]
---

Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests. Update doplan/dashboard.md file.

## Workflow
1. Read all progress.json files from feature directories
2. Calculate overall and phase progress percentages
3. Check GitHub for active PRs (if GitHub integration enabled)
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + ` (markdown)
6. Update ` + backtick + `doplan/dashboard.html` + backtick + ` (visual HTML dashboard)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions

$ARGUMENTS
`,
		"implement.md": `---
description: Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch.
argument-hint: [FEATURE="<feature name>"] [BRANCH="<branch name>"]
---

Help implement the current feature. Check the feature's plan.md, design.md, and tasks.md files. Create a GitHub branch for this feature and start implementation.

## Workflow
1. Check current feature context from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Automatically create GitHub branch:
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: ` + backtick + `git checkout -b {branch-name}` + backtick + `
3. Initial commit:
   - Add plan.md, design.md, tasks.md files
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: ` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on:
   - ` + backtick + `plan.md` + backtick + ` - Feature plan
   - ` + backtick + `design.md` + backtick + ` - Design specifications
   - ` + backtick + `tasks.md` + backtick + ` - Task breakdown

$ARGUMENTS
`,
		"next.md": `---
description: Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.
argument-hint: [PRIORITY="<high|medium|low>"]
---

Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.

## Workflow
1. Read current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Scan all feature directories for incomplete tasks
3. Check progress.json files
4. Consider dependencies between features
5. Recommend highest priority action
6. Display recommendation in dashboard format

## Analysis Factors
- Task completion status
- Feature dependencies
- Phase priorities
- Blocked items
- Progress percentages
- GitHub branch status

$ARGUMENTS
`,
		"progress.md": `---
description: Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.
argument-hint: [FEATURE="<feature path>"] [PHASE="<phase number>"]
---

Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.

## Workflow
1. Scan all feature directories in ` + backtick + `doplan/` + backtick + `
2. Read tasks.md files from each feature
3. Count completed tasks (marked with [x])
4. Calculate completion percentages
5. Update progress.json files:
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `doplan/dashboard.md` + backtick + `
   - ` + backtick + `doplan/dashboard.html` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress

$ARGUMENTS
`,
	}

	// Codex uses .codex/prompts/ in project root for project-specific commands
	// Global commands would go in ~/.codex/prompts/ but we're installing project-specific ones
	promptsDir := filepath.Join(i.projectRoot, ".codex", "prompts")

	for filename, content := range commands {
		path := filepath.Join(promptsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) installOpenCodeCommands() error {
	// OpenCode commands are Markdown files with YAML frontmatter
	// See: https://opencode.ai/docs/commands/
	// Commands are invoked with /command-name
	// Supports: $ARGUMENTS, $1-$9, !`command`, @file references
	backtick := "`"
	commands := map[string]string{
		"discuss.md": `---
description: Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack.
---

Start the DoPlan idea discussion workflow. This will help refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Context
- Current project state: @.cursor/config/doplan-state.json
- Existing idea notes: @doplan/idea-notes.md

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

## Output
- Idea notes document
- Updated state file
- Tech stack recommendations
- Feature organization suggestions

$ARGUMENTS
`,
		"refine.md": `---
description: Refine and enhance the current idea. Review existing idea notes and suggest improvements, additional features, and better organization.
---

Review the current idea and suggest improvements, additional features, and better organization. Use the idea notes and current state to provide actionable enhancements.

## Context
- Existing idea notes: @doplan/idea-notes.md
- Current state: @.cursor/config/doplan-state.json

## Workflow
1. Review existing idea notes
2. Review current state
3. Suggest additional features
4. Identify gaps in the plan
5. Enhance technical specifications
6. Update idea documentation

## Focus Areas
- Feature completeness
- Technical feasibility
- User experience improvements
- Architecture enhancements
- Risk mitigation

$ARGUMENTS
`,
		"generate.md": `---
description: Generate Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea.
---

Generate the Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea. Create files in doplan/ directory.

## Context
- Idea notes: @doplan/idea-notes.md
- Current state: @.cursor/config/doplan-state.json
- Templates: @doplan/templates/

## Workflow
1. Read idea notes and state
2. Generate ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
3. Generate ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
4. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
5. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas
6. Use templates from ` + backtick + `doplan/templates/` + backtick + ` directory

## Documents Created
- PRD.md - Complete product requirements
- structure.md - Project architecture
- api-spec.json - API contracts
- data-model.md - Data models

$ARGUMENTS
`,
		"plan.md": `---
description: Generate the project plan with phases and features. Create the directory structure following BMAD-METHOD approach.
---

Generate the project plan with phases and features. Create the directory structure: doplan/01-phase/01-Feature/, etc. Each feature should have plan.md, design.md, and tasks.md files.

## Context
- PRD: @doplan/PRD.md
- Contracts: @doplan/contracts/
- Templates: @doplan/templates/

## Workflow
1. Read PRD and contracts
2. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
3. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
4. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
5. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
6. Update dashboard with new structure

$ARGUMENTS
`,
		"dashboard.md": `---
description: Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests.
---

Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests. Update doplan/dashboard.md file.

## Context
- Current git status: !` + backtick + `git status` + backtick + `
- Current branch: !` + backtick + `git branch --show-current` + backtick + `
- Recent commits: !` + backtick + `git log --oneline -10` + backtick + `

## Workflow
1. Read all progress.json files from feature directories
2. Calculate overall and phase progress percentages
3. Check GitHub for active PRs (if GitHub integration enabled)
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + ` (markdown)
6. Update ` + backtick + `doplan/dashboard.html` + backtick + ` (visual HTML dashboard)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions

$ARGUMENTS
`,
		"implement.md": `---
description: Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch.
---

Help implement the current feature. Check the feature's plan.md, design.md, and tasks.md files. Create a GitHub branch for this feature and start implementation.

## Context
- Current state: @.cursor/config/doplan-state.json
- Current git status: !` + backtick + `git status` + backtick + `
- Current branch: !` + backtick + `git branch --show-current` + backtick + `

## Workflow
1. Check current feature context
2. Automatically create GitHub branch:
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: !` + backtick + `git checkout -b {branch-name}` + backtick + `
3. Initial commit:
   - Add plan.md, design.md, tasks.md files
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: !` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on:
   - ` + backtick + `plan.md` + backtick + ` - Feature plan
   - ` + backtick + `design.md` + backtick + ` - Design specifications
   - ` + backtick + `tasks.md` + backtick + ` - Task breakdown

$ARGUMENTS
`,
		"next.md": `---
description: Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.
---

Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.

## Context
- Current state: @.cursor/config/doplan-state.json
- Dashboard: @doplan/dashboard.md

## Workflow
1. Read current state
2. Scan all feature directories for incomplete tasks
3. Check progress.json files
4. Consider dependencies between features
5. Recommend highest priority action
6. Display recommendation in dashboard format

## Analysis Factors
- Task completion status
- Feature dependencies
- Phase priorities
- Blocked items
- Progress percentages
- GitHub branch status

$ARGUMENTS
`,
		"progress.md": `---
description: Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.
---

Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.

## Context
- Current state: @.cursor/config/doplan-state.json
- Dashboard: @doplan/dashboard.md

## Workflow
1. Scan all feature directories in ` + backtick + `doplan/` + backtick + `
2. Read tasks.md files from each feature
3. Count completed tasks (marked with [x])
4. Calculate completion percentages
5. Update progress.json files:
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `doplan/dashboard.md` + backtick + `
   - ` + backtick + `doplan/dashboard.html` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress

$ARGUMENTS
`,
	}

	commandsDir := filepath.Join(i.projectRoot, ".opencode", "command")

	for filename, content := range commands {
		path := filepath.Join(commandsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) installQwenCommands() error {
	// Qwen Code commands are TOML files with [command] section
	// Commands are invoked with /command-name
	// Supports namespacing: subdirectories become /namespace:command
	backtick := "`"
	commands := map[string]string{
		"discuss.toml": `[command]
name = "discuss"
description = "Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack."
prompt = """Start the DoPlan idea discussion workflow. This will help refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

## Output
- Idea notes document
- Updated state file
- Tech stack recommendations
- Feature organization suggestions"""
`,
		"refine.toml": `[command]
name = "refine"
description = "Refine and enhance the current idea. Review existing idea notes and suggest improvements, additional features, and better organization."
prompt = """Review the current idea and suggest improvements, additional features, and better organization. Use the idea notes and current state to provide actionable enhancements.

## Workflow
1. Review existing idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Review current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Suggest additional features
4. Identify gaps in the plan
5. Enhance technical specifications
6. Update idea documentation

## Focus Areas
- Feature completeness
- Technical feasibility
- User experience improvements
- Architecture enhancements
- Risk mitigation"""
`,
		"generate.toml": `[command]
name = "generate"
description = "Generate Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea."
prompt = """Generate the Product Requirements Document (PRD), project structure document, and API contracts based on the refined idea. Create files in doplan/ directory.

## Workflow
1. Read idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
2. Read state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
3. Generate ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
4. Generate ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
5. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
6. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas
7. Use templates from ` + backtick + `doplan/templates/` + backtick + ` directory

## Documents Created
- PRD.md - Complete product requirements
- structure.md - Project architecture
- api-spec.json - API contracts
- data-model.md - Data models"""
`,
		"plan.toml": `[command]
name = "plan"
description = "Generate the project plan with phases and features. Create the directory structure following BMAD-METHOD approach."
prompt = """Generate the project plan with phases and features. Create the directory structure: doplan/01-phase/01-Feature/, etc. Each feature should have plan.md, design.md, and tasks.md files.

## Workflow
1. Read PRD from ` + backtick + `doplan/PRD.md` + backtick + `
2. Read contracts from ` + backtick + `doplan/contracts/` + backtick + `
3. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
4. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
5. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
6. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
7. Update dashboard with new structure"""
`,
		"dashboard.toml": `[command]
name = "dashboard"
description = "Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests."
prompt = """Generate and display the project dashboard showing overall progress, phase progress, feature progress, and active pull requests. Update doplan/dashboard.md file.

## Workflow
1. Read all progress.json files from feature directories
2. Calculate overall and phase progress percentages
3. Check GitHub for active PRs (if GitHub integration enabled)
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + ` (markdown)
6. Update ` + backtick + `doplan/dashboard.html` + backtick + ` (visual HTML dashboard)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions"""
`,
		"implement.toml": `[command]
name = "implement"
description = "Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch."
prompt = """Help implement the current feature. Check the feature's plan.md, design.md, and tasks.md files. Create a GitHub branch for this feature and start implementation.

## Workflow
1. Check current feature context from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Automatically create GitHub branch:
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: ` + backtick + `git checkout -b {branch-name}` + backtick + `
3. Initial commit:
   - Add plan.md, design.md, tasks.md files
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: ` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on:
   - ` + backtick + `plan.md` + backtick + ` - Feature plan
   - ` + backtick + `design.md` + backtick + ` - Design specifications
   - ` + backtick + `tasks.md` + backtick + ` - Task breakdown"""
`,
		"next.toml": `[command]
name = "next"
description = "Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next."
prompt = """Analyze the current project state and recommend the next best action. Check progress, incomplete tasks, and suggest what to work on next.

## Workflow
1. Read current state from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
2. Scan all feature directories for incomplete tasks
3. Check progress.json files
4. Consider dependencies between features
5. Recommend highest priority action
6. Display recommendation in dashboard format

## Analysis Factors
- Task completion status
- Feature dependencies
- Phase priorities
- Blocked items
- Progress percentages
- GitHub branch status"""
`,
		"progress.toml": `[command]
name = "progress"
description = "Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files."
prompt = """Update all progress tracking files. Recalculate progress bars for phases and features, update dashboard, and sync progress.json files.

## Workflow
1. Scan all feature directories in ` + backtick + `doplan/` + backtick + `
2. Read tasks.md files from each feature
3. Count completed tasks (marked with [x])
4. Calculate completion percentages
5. Update progress.json files:
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `doplan/dashboard.md` + backtick + `
   - ` + backtick + `doplan/dashboard.html` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress"""
`,
	}

	commandsDir := filepath.Join(i.projectRoot, ".qwen", "commands")

	for filename, content := range commands {
		path := filepath.Join(commandsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) createTemplates() error {
	templatesGen := generators.NewTemplatesGenerator(i.projectRoot)
	return templatesGen.Generate()
}

func (i *Installer) generateAgents() error {
	agentsGen := generators.NewAgentsGenerator(i.projectRoot)
	return agentsGen.Generate()
}

func (i *Installer) generateRules() error {
	rulesGen := generators.NewRulesGenerator(i.projectRoot)
	return rulesGen.Generate()
}

func (i *Installer) setupIDEIntegration() error {
	// Setup IDE-specific integration (symlinks, config files, etc.)
	return integration.SetupIDE(i.projectRoot, i.ide)
}

func (i *Installer) verifyInstallation() error {
	// Verify the installation was successful
	return integration.VerifyIDE(i.projectRoot, i.ide)
}

func (i *Installer) generateContext() error {
	contextGen := generators.NewContextGenerator(i.projectRoot)
	return contextGen.Generate()
}

func (i *Installer) generateConfig() error {
	cfg := config.NewConfig(i.ide)
	cfgMgr := config.NewManager(i.projectRoot)
	return cfgMgr.SaveConfig(cfg)
}

func (i *Installer) createREADME() error {
	readmeGen := generators.NewREADMEGenerator(i.projectRoot)
	content := readmeGen.Generate()
	path := filepath.Join(i.projectRoot, "README.md")
	return os.WriteFile(path, []byte(content), 0644)
}

func (i *Installer) createDashboard() error {
	initialDashboard := generators.GenerateInitialDashboard()
	path := filepath.Join(i.projectRoot, "doplan", "dashboard.md")
	return os.WriteFile(path, []byte(initialDashboard), 0644)
}

func runInstallerAnimation(name string, fn func() error) error {
	spinnerFrames := []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
	done := make(chan struct{})

	go func() {
		idx := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r%s %c", name, spinnerFrames[idx%len(spinnerFrames)])
				time.Sleep(90 * time.Millisecond)
				idx++
			}
		}
	}()

	err := fn()
	close(done)

	status := "✓"
	if err != nil {
		status = "✗"
	}

	fmt.Printf("\r%s %s\n", name, status)
	return err
}
