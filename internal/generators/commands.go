package generators

import (
	"os"
	"path/filepath"
)

// CommandsGenerator generates command definitions
type CommandsGenerator struct {
	projectRoot string
}

// NewCommandsGenerator creates a new commands generator
func NewCommandsGenerator(projectRoot string) *CommandsGenerator {
	return &CommandsGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates all command files in .doplan/ai/commands/
func (g *CommandsGenerator) Generate() error {
	commandsDir := filepath.Join(g.projectRoot, ".doplan", "ai", "commands")

	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		return err
	}

	backtick := "`"
	commands := map[string]string{
		"run.md": `# Run

## Overview
Auto-detect and run the development server for your project. Supports multiple frameworks and package managers.

## Workflow
1. Detect project type (Next.js, React, Vue, Go, Python, etc.)
2. Detect package manager (npm, yarn, pnpm, go, python, etc.)
3. Run appropriate dev command:
   - Next.js/React: ` + backtick + `npm run dev` + backtick + `
   - Go: ` + backtick + `go run .` + backtick + `
   - Python: ` + backtick + `python -m uvicorn main:app --reload` + backtick + `
   - etc.
4. Display server URL and status
5. Monitor for errors and restart if needed

## Supported Frameworks
- Next.js, React, Vue, Svelte
- Express, FastAPI, Flask
- Go applications
- Python applications
- And more (auto-detected)
`,
		"undo.md": `# Undo

## Overview
Revert the last DoPlan action using the action history stored in ` + backtick + `.doplan/state.json` + backtick + `.

## Workflow
1. Read action history from ` + backtick + `.doplan/state.json` + backtick + `
2. Display last 10 actions with timestamps
3. Allow user to select which action to undo
4. Revert changes:
   - Restore files from checkpoint
   - Revert git commits
   - Restore configuration
   - Update state and progress
5. Confirm undo completion

## Supported Actions
- File creation/modification
- Git operations (commits, branches)
- Configuration changes
- Progress updates
`,
		"deploy.md": `# Deploy

## Overview
Launch a multi-platform deployment wizard to deploy your project to various hosting platforms.

## Workflow
1. Detect project type and build configuration
2. Show deployment platform options:
   - Vercel (Next.js, React, etc.)
   - Netlify (Static sites, JAMstack)
   - Railway (Full-stack apps)
   - Render (Docker, static sites)
   - Coolify (Self-hosted)
   - Custom (Docker, SSH, etc.)
3. Guide through platform-specific setup
4. Configure environment variables
5. Deploy and provide deployment URL
6. Set up auto-deployment (optional)

## Platform Support
- Vercel, Netlify, Railway, Render, Coolify
- Custom deployments via Docker or SSH
`,
		"publish.md": `# Publish

## Overview
Launch a package publishing wizard to publish your project to package registries.

## Workflow
1. Detect package type (npm, Homebrew, Scoop, Winget, etc.)
2. Show publishing options:
   - npm (for Node.js packages)
   - Homebrew (for macOS/Linux CLI tools)
   - Scoop (for Windows CLI tools)
   - Winget (for Windows apps)
3. Validate package configuration
4. Build package
5. Publish to registry
6. Provide installation instructions

## Package Types
- npm packages
- Homebrew formulas
- Scoop manifests
- Winget manifests
`,
		"create.md": `# Create

## Overview
Launch the new project wizard to create a new DoPlan project with template selection.

## Workflow
1. Show template gallery
2. Allow template selection or start without template
3. Configure project settings
4. Set up GitHub repository
5. Configure IDE integration
6. Initialize project structure
7. Generate initial files

## Templates Available
- Website (Frontend)
- Website & Admin Dashboard
- Web Application
- Mobile Application
- Micro SaaS
- SaaS
- Web Game
- CLI
- Chrome Extension
- AI Agent
- And more
`,
		"security.md": `# Security

## Overview
Run comprehensive security scans on your project to detect vulnerabilities and security issues.

## Workflow
1. Detect project type and dependencies
2. Run security scans:
   - npm audit (for Node.js projects)
   - trufflehog (secrets scanning)
   - git-secrets (Git history scanning)
   - gosec (Go security scanner)
   - dive (Docker image analysis)
3. Display security report with:
   - Vulnerabilities found
   - Severity levels
   - Affected packages/files
   - Fix recommendations
4. Offer to auto-fix issues (if available)

## Scanners Used
- npm audit, trufflehog, git-secrets
- gosec, dive, and more
`,
		"fix.md": `# Fix

## Overview
Run AI-powered auto-fix to automatically resolve common issues and errors in your project.

## Workflow
1. Scan project for issues:
   - Linter errors
   - Security vulnerabilities
   - Code quality issues
   - Configuration problems
2. Use AI to suggest fixes
3. Show fix preview
4. Apply fixes with user confirmation
5. Verify fixes work
6. Update progress and state

## Fix Types
- Code linting errors
- Security vulnerabilities
- Dependency updates
- Configuration errors
- And more
`,
		"design.md": `# Design

## Overview
Launch the Design Preferences & Requirements (DPR) questionnaire to generate a complete design system.

## Workflow
1. Launch interactive questionnaire (20-30 questions)
2. Collect design preferences:
   - Audience analysis
   - Emotional design
   - Style preferences
   - Colors and typography
   - Layout and components
   - Animation preferences
3. Generate DPR.md document
4. Generate design-tokens.json
5. Generate .doplan/ai/rules/design_rules.mdc
6. Provide design system overview

## Output Files
- doplan/design/DPR.md
- doplan/design/design-tokens.json
- .doplan/ai/rules/design_rules.mdc
`,
		"keys.md": `# Keys

## Overview
Manage API keys and service configuration. Detect required keys, validate them, and manage service setup guides.

## Workflow
1. Detect required services from dependencies
2. Check .env file for configured keys
3. Show RAKD status:
   - Configured keys
   - Pending keys
   - Optional keys
4. Allow actions:
   - Validate all keys
   - Check for missing keys
   - Sync .env.example
   - Launch setup wizard for a service
   - Test API connections
5. Generate/update RAKD.md
6. Generate SOPS guides for services

## Features
- Auto-detect required services
- Validate API keys
- Generate service setup guides
- Manage .env files
`,
		"discuss.md": `# Discuss

## Overview
Start idea discussion and refinement workflow. This command helps refine your project idea, suggest improvements, organize features, and select the best tech stack.

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into logical phases
4. Recommend the best tech stack for your project
5. Save results to:
   - ` + backtick + `.doplan/state.json` + backtick + `
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
2. Review current state from ` + backtick + `.doplan/state.json` + backtick + `
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
2. Read state from ` + backtick + `.doplan/state.json` + backtick + `
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
Generate the project plan with phases and features. Create the directory structure following DoPlan workflow.

## Workflow
1. Read PRD from ` + backtick + `doplan/PRD.md` + backtick + `
2. Read contracts from ` + backtick + `doplan/contracts/` + backtick + `
3. Create phase directories: ` + backtick + `doplan/01-phase-name/` + backtick + `, ` + backtick + `doplan/02-phase-name/` + backtick + `, etc.
4. Create feature directories: ` + backtick + `doplan/01-phase-name/01-feature-name/` + backtick + `, etc.
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
├── 01-phase-name/
│   ├── phase-plan.md
│   ├── phase-progress.json
│   ├── 01-feature-name/
│   │   ├── plan.md
│   │   ├── design.md
│   │   ├── tasks.md
│   │   └── progress.json
│   └── 02-feature-name/
└── 02-phase-name/
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
5. Update ` + backtick + `.doplan/dashboard.json` + backtick + ` (machine-readable)
6. Update ` + backtick + `.doplan/dashboard.md` + backtick + ` (human-readable)

## Dashboard Sections
- Overall project progress
- Phase-by-phase progress
- Feature progress within phases
- Active pull requests
- Recent GitHub activity (commits, branches, pushes)
- Next recommended actions

## Usage
After running this command, view the dashboard:
- Markdown: Open ` + backtick + `.doplan/dashboard.md` + backtick + `
- CLI: Run ` + backtick + `doplan` + backtick + ` in terminal
`,
		"implement.md": `# Implement

## Overview
Start implementing a feature. This command helps guide implementation based on the feature's planning documents and automatically creates a GitHub branch.

## Workflow
1. Check current feature context from ` + backtick + `.doplan/state.json` + backtick + `
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
1. Read current state from ` + backtick + `.doplan/state.json` + backtick + `
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
   - Feature-level: ` + backtick + `doplan/XX-phase/XX-feature/progress.json` + backtick + `
   - Phase-level: ` + backtick + `doplan/XX-phase/phase-progress.json` + backtick + `
6. Regenerate dashboard:
   - ` + backtick + `.doplan/dashboard.json` + backtick + `
   - ` + backtick + `.doplan/dashboard.md` + backtick + `
7. Sync GitHub data (if enabled)
8. Update state file: ` + backtick + `.doplan/state.json` + backtick + `

## Progress Calculation
- Feature progress: (completed tasks / total tasks) * 100
- Phase progress: Average of all feature progress in phase
- Overall progress: Average of all phase progress
`,
	}

	for filename, content := range commands {
		path := filepath.Join(commandsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}
