package generators

import (
	"os"
	"path/filepath"
)

// RulesGenerator generates workflow rules
type RulesGenerator struct {
	projectRoot string
}

// NewRulesGenerator creates a new rules generator
func NewRulesGenerator(projectRoot string) *RulesGenerator {
	return &RulesGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates all rule files
func (g *RulesGenerator) Generate() error {
	rulesDir := filepath.Join(g.projectRoot, ".cursor", "rules")

	if err := os.MkdirAll(rulesDir, 0755); err != nil {
		return err
	}

	rules := map[string]string{
		"workflow-rules.md": generateWorkflowRules(),
		"github-rules.md":   generateGitHubRules(),
		"command-rules.md":  generateCommandRules(),
		"branch-rules.md":   generateBranchRules(),
		"commit-rules.md":   generateCommitRules(),
		"docs-rules.md":     generateDocsRules(),
	}

	for filename, content := range rules {
		path := filepath.Join(rulesDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func generateWorkflowRules() string {
	backtick := "`"
	return `# DoPlan Workflow Rules

These rules govern how Cursor should operate within the DoPlan workflow.

## Core Principles

- Follow Spec-Kit and BMAD-METHOD methodologies
- Generate contracts before implementation
- Track progress in dashboard
- Use feature branches for each feature
- All commands are available as slash commands in Cursor

## Workflow Stages

### 1. Idea Discussion (/Discuss)
- Ask comprehensive questions about the idea
- Suggest improvements and enhancements
- Help organize features into phases
- Recommend tech stack based on requirements
- Save results to ` + backtick + `.cursor/config/doplan-state.json` + backtick + ` and ` + backtick + `doplan/idea-notes.md` + backtick + `

### 2. Idea Refinement (/Refine)
- Review existing idea notes
- Suggest additional features
- Identify gaps in the plan
- Enhance technical specifications
- Update idea documentation

### 3. Document Generation (/Generate)
- Create ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
- Create ` + backtick + `doplan/structure.md` + backtick + ` - Project structure
- Create ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification
- Create ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models
- Use templates from ` + backtick + `doplan/templates/` + backtick + `

### 4. Planning (/Plan)
- Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
- Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
- Generate ` + backtick + `plan.md` + backtick + `, ` + backtick + `design.md` + backtick + `, ` + backtick + `tasks.md` + backtick + ` for each feature
- Create ` + backtick + `phase-plan.md` + backtick + ` and ` + backtick + `phase-progress.json` + backtick + ` for each phase
- Update dashboard with new structure

### 5. Implementation (/Implement)
- Check current feature context from state
- Automatically create GitHub branch: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
- Initialize feature branch
- Guide implementation based on plan.md, design.md, tasks.md
- Update progress as tasks complete

### 6. Progress Tracking (/Progress)
- Scan all feature directories
- Read tasks.md files
- Calculate completion percentages
- Update progress.json files
- Regenerate dashboard
- Sync GitHub data

## State Management

- State lives in ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
- Progress tracked in ` + backtick + `doplan/**/progress.json` + backtick + ` files
- Dashboard in ` + backtick + `doplan/dashboard.md` + backtick + `
- Always update state after command execution

## Best Practices

- Always start with /Discuss before /Plan
- Generate PRD before creating phases
- Follow the phase → feature hierarchy
- Update tasks.md as you work
- Run /Progress regularly
- Use /Next to get recommendations
`
}

func generateGitHubRules() string {
	backtick := "`"
	codeBlock := "```"
	return `# GitHub Integration Rules

## Branch Naming Convention

- Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
- Example: ` + backtick + `feature/01-phase-01-user-authentication` + backtick + `
- Use kebab-case for feature names
- Always prefix with ` + backtick + `feature/` + backtick + `

## Automatic Branch Creation

When /Implement command is used:

1. **Check current feature:**
   - Read from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - Get current phase and feature

2. **Generate branch name:**
   - Format: ` + backtick + `feature/{phase-id}-{feature-id}-{feature-name}` + backtick + `
   - Convert to kebab-case
   - Example: ` + backtick + `feature/01-phase-01-user-authentication` + backtick + `

3. **Create branch:**
   ` + codeBlock + `bash
   git checkout -b feature/01-phase-01-user-authentication
   ` + codeBlock + `

4. **Initial commit:**
   - Add feature planning files (plan.md, design.md, tasks.md)
   - Commit message: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push to remote

5. **Update state:**
   - Save branch name to state
   - Update dashboard

## Automatic Commit Rules

### Commit Message Format

Follow conventional commits:
- ` + backtick + `feat:` + backtick + ` - New feature
- ` + backtick + `fix:` + backtick + ` - Bug fix
- ` + backtick + `docs:` + backtick + ` - Documentation
- ` + backtick + `refactor:` + backtick + ` - Code refactoring
- ` + backtick + `test:` + backtick + ` - Tests
- ` + backtick + `chore:` + backtick + ` - Maintenance

### Commit Workflow

1. **During development:**
   - Commit frequently with clear messages
   - Reference task numbers when possible
   - Format: ` + backtick + `feat: implement {task-description}` + backtick + `

2. **Task completion:**
   - Update tasks.md (check off completed tasks)
   - Commit: ` + backtick + `feat: complete {task-name}` + backtick + `

3. **Feature completion:**
   - Mark all tasks complete in tasks.md
   - Update progress.json: ` + backtick + `"status": "complete"` + backtick + `
   - Commit: ` + backtick + `feat: complete {feature-name}` + backtick + `
   - Push branch

## Automatic Push Rules

### Push on Commit

After each commit:
` + codeBlock + `bash
git push origin feature/XX-phase-XX-feature-name
` + codeBlock + `

### Push Status Tracking

- Track push success/failure
- Update dashboard with push status
- Log push history in ` + backtick + `doplan/github-data.json` + backtick + `

## Pull Request Automation

### When Feature is Complete

1. **Check completion:**
   - All tasks in tasks.md are checked
   - progress.json shows ` + backtick + `"status": "complete"` + backtick + `

2. **Create PR automatically:**
   - Use GitHub CLI: ` + backtick + `gh pr create` + backtick + `
   - Title: ` + backtick + `Feature: {feature-name}` + backtick + `
   - Body includes links to plan.md, design.md, tasks.md
   - Base branch: ` + backtick + `main` + backtick + `
   - Head branch: feature branch

3. **Update dashboard:**
   - Add PR link to dashboard
   - Update state with PR number

## Merge Automation

After PR approval:
1. Merge PR: ` + backtick + `gh pr merge {pr-number} --merge` + backtick + `
2. Delete branch: ` + backtick + `gh pr merge {pr-number} --delete-branch` + backtick + `
3. Update progress
4. Update dashboard
5. Sync with main branch
`
}

func generateCommandRules() string {
	backtick := "`"
	return `# DoPlan Command Rules

## /Discuss Command

**Purpose:** Start idea discussion and refinement

**Workflow:**
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
3. Help organize features into phases
4. Recommend tech stack based on requirements
5. Save results to:
   - ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - ` + backtick + `doplan/idea-notes.md` + backtick + `

**Output:**
- Idea notes document
- Updated state file
- Tech stack recommendations

## /Refine Command

**Purpose:** Enhance and improve existing idea

**Workflow:**
1. Review existing idea notes
2. Suggest additional features
3. Identify gaps in the plan
4. Enhance technical specifications
5. Update idea documentation

## /Generate Command

**Purpose:** Generate PRD, Structure, and API contracts

**Workflow:**
1. Read idea notes and state
2. Generate ` + backtick + `doplan/PRD.md` + backtick + `
3. Generate ` + backtick + `doplan/structure.md` + backtick + `
4. Generate ` + backtick + `doplan/contracts/api-spec.json` + backtick + `
5. Generate ` + backtick + `doplan/contracts/data-model.md` + backtick + `
6. Use templates from ` + backtick + `doplan/templates/` + backtick + `

## /Plan Command

**Purpose:** Generate phase and feature structure

**Workflow:**
1. Read PRD and contracts
2. Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, etc.
3. Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
4. Generate for each phase:
   - ` + backtick + `phase-plan.md` + backtick + `
   - ` + backtick + `phase-progress.json` + backtick + `
5. Generate for each feature:
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `
6. Update dashboard

## /Dashboard Command

**Purpose:** Show project dashboard with progress

**Workflow:**
1. Read all progress.json files
2. Calculate overall and phase progress
3. Check GitHub for active PRs
4. Generate visual progress bars
5. Update ` + backtick + `doplan/dashboard.md` + backtick + `
6. Update ` + backtick + `doplan/dashboard.html` + backtick + `

## /Implement Command

**Purpose:** Start implementing a feature

**Workflow:**
1. Check current feature context from state
2. **Automatically create GitHub branch:**
   - Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Create branch: ` + backtick + `git checkout -b {branch-name}` + backtick + `
3. **Initial commit:**
   - Add plan.md, design.md, tasks.md
   - Commit: ` + backtick + `docs: add planning docs for {feature-name}` + backtick + `
   - Push: ` + backtick + `git push origin {branch-name}` + backtick + `
4. Update state with branch name
5. Update dashboard
6. Guide implementation based on plan.md, design.md, tasks.md

## /Next Command

**Purpose:** Get recommendation for next action

**Workflow:**
1. Analyze current state
2. Check incomplete tasks
3. Consider dependencies
4. Recommend highest priority action
5. Display in dashboard format

## /Progress Command

**Purpose:** Update all progress tracking

**Workflow:**
1. Scan all feature directories
2. Read tasks.md files
3. Calculate completion percentages
4. Update progress.json files
5. Regenerate dashboard
6. Sync GitHub data
7. Update state file
`
}

func generateBranchRules() string {
	backtick := "`"
	codeBlock := "```"
	return `# Branch Management Rules

## Branch Naming Convention

**Format:** ` + backtick + `feature/{phase-id}-{feature-id}-{feature-name}` + backtick + `

**Examples:**
- ` + backtick + `feature/01-phase-01-user-authentication` + backtick + `
- ` + backtick + `feature/01-phase-02-database-setup` + backtick + `
- ` + backtick + `feature/02-phase-01-task-creation` + backtick + `

**Rules:**
- Always prefix with ` + backtick + `feature/` + backtick + `
- Use phase and feature IDs from plan structure
- Convert feature name to kebab-case
- No spaces or special characters (except hyphens)

## Automatic Branch Creation

### Trigger: /Implement Command

When /Implement is executed:

1. **Get feature information:**
   - Read from ` + backtick + `.cursor/config/doplan-state.json` + backtick + `
   - Current phase: ` + backtick + `state.currentPhase` + backtick + `
   - Current feature: ` + backtick + `state.currentFeature` + backtick + `

2. **Generate branch name:**
   - Format: ` + backtick + `feature/{phase-id}-{feature-id}-{feature-name}` + backtick + `
   - Convert to kebab-case

3. **Create and switch to branch:**
   ` + codeBlock + `bash
   git checkout -b {branchName}
   ` + codeBlock + `

4. **Initial commit:**
   ` + codeBlock + `bash
   git add doplan/{phase}/{feature}/*
   git commit -m "docs: add planning docs for {feature-name}"
   git push origin {branchName}
   ` + codeBlock + `

5. **Update state:**
   - Save branch name to feature in state
   - Update dashboard

## Branch Workflow

### During Development

1. **Work on feature branch:**
   - All commits go to feature branch
   - Regular commits with clear messages
   - Push frequently

2. **Update tasks:**
   - Mark tasks complete in tasks.md
   - Commit: ` + backtick + `feat: complete {task-name}` + backtick + `

3. **Keep branch updated:**
   - Regularly merge main into feature branch
   - Resolve conflicts early

### Feature Completion

1. **Final checks:**
   - All tasks in tasks.md are checked
   - progress.json shows complete
   - All tests pass

2. **Final commit:**
   ` + codeBlock + `bash
   git add .
   git commit -m "feat: complete {feature-name}"
   git push origin {branchName}
   ` + codeBlock + `

3. **Create PR:**
   - Automatically via /Implement or manual
   - PR will be created when feature marked complete

## Branch Cleanup

After PR merge:
1. Switch to main: ` + backtick + `git checkout main` + backtick + `
2. Pull latest: ` + backtick + `git pull origin main` + backtick + `
3. Delete local branch: ` + backtick + `git branch -d {branchName}` + backtick + `
4. Delete remote branch: ` + backtick + `git push origin --delete {branchName}` + backtick + `
`
}

func generateCommitRules() string {
	backtick := "`"
	codeBlock := "```"
	return `# Commit Rules

## Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/) specification:

` + codeBlock + `
<type>(<scope>): <subject>

<body>

<footer>
` + codeBlock + `

### Types

- ` + backtick + `feat:` + backtick + ` - New feature
- ` + backtick + `fix:` + backtick + ` - Bug fix
- ` + backtick + `docs:` + backtick + ` - Documentation changes
- ` + backtick + `refactor:` + backtick + ` - Code refactoring
- ` + backtick + `test:` + backtick + ` - Adding or updating tests
- ` + backtick + `chore:` + backtick + ` - Maintenance tasks
- ` + backtick + `style:` + backtick + ` - Code style changes (formatting, etc.)
- ` + backtick + `perf:` + backtick + ` - Performance improvements

### Examples

` + codeBlock + `
feat: implement user authentication endpoint
fix: resolve login validation issue
docs: update API documentation
refactor: simplify authentication logic
test: add unit tests for auth service
` + codeBlock + `

## Automatic Commit Workflow

### During Feature Development

1. **Task completion:**
   - Update tasks.md (check off task)
   - Commit: ` + backtick + `feat: complete {task-description}` + backtick + `
   - Push: ` + backtick + `git push origin {branch-name}` + backtick + `

2. **Implementation milestone:**
   - Commit: ` + backtick + `feat: implement {component-name}` + backtick + `
   - Push immediately

3. **Bug fixes:**
   - Commit: ` + backtick + `fix: resolve {issue-description}` + backtick + `
   - Push immediately

### Feature Planning Phase

1. **Initial planning docs:**
   ` + codeBlock + `bash
   git add doplan/{phase}/{feature}/*
   git commit -m "docs: add planning docs for {feature-name}"
   ` + codeBlock + `

2. **Plan updates:**
   ` + codeBlock + `bash
   git commit -m "docs: update {feature-name} plan"
   ` + codeBlock + `

### Feature Completion

1. **Final commit:**
   ` + codeBlock + `bash
   git add .
   git commit -m "feat: complete {feature-name}"
   git push origin {branch-name}
   ` + codeBlock + `

## Automatic Push Rules

### Always Push After Commit

After every commit, automatically push:
` + codeBlock + `bash
git push origin {current-branch}
` + codeBlock + `

### Push Status Tracking

- Track push success/failure
- Update dashboard with push status
- Log in ` + backtick + `doplan/github-data.json` + backtick + `

### Error Handling

If push fails:
1. Show error message
2. Suggest: ` + backtick + `git pull origin {branch-name}` + backtick + `
3. Retry push after pull

## Commit Frequency

- Commit after each logical unit of work
- Commit at least once per day
- Commit before switching tasks
- Commit before leaving for the day

## Commit Best Practices

- Write clear, descriptive commit messages
- Reference task numbers when possible
- Keep commits focused (one logical change)
- Don't commit broken code
- Test before committing
- Review changes before committing
`
}

func generateDocsRules() string {
	backtick := "`"
	codeBlock := "```"
	return `# Documentation Organization and Naming Rules

## Documentation Structure

### Root Documentation Files

All project documentation follows a consistent structure:

` + codeBlock + `
project-root/
├── README.md                    # Main project documentation
├── CONTEXT.md                   # Tech stack context (for IDEs/CLIs)
├── CHANGELOG.md                 # Version history
├── CONTRIBUTING.md              # Contribution guidelines
├── LICENSE                      # License file
├── doplan/                      # DoPlan planning documents
│   ├── PRD.md                  # Product Requirements Document
│   ├── structure.md            # Project structure
│   ├── dashboard.md            # Progress dashboard (markdown)
│   ├── dashboard.html          # Progress dashboard (HTML)
│   ├── idea-notes.md            # Idea discussion notes
│   ├── contracts/               # API contracts and specifications
│   │   ├── api-spec.json       # OpenAPI/Swagger specification
│   │   └── data-model.md       # Data models documentation
│   ├── templates/              # Document templates
│   │   ├── plan.md             # Feature plan template
│   │   ├── design.md          # Design specification template
│   │   └── tasks.md            # Task list template
│   └── XX-phase/               # Phase directories (01-phase, 02-phase, etc.)
│       ├── phase-plan.md       # Phase planning document
│       ├── phase-progress.json # Phase progress tracking
│       └── XX-Feature/         # Feature directories (01-Feature, 02-Feature, etc.)
│           ├── plan.md         # Feature plan
│           ├── design.md       # Feature design
│           ├── tasks.md        # Feature tasks
│           └── progress.json   # Feature progress tracking
` + codeBlock + `

## File Naming Conventions

### General Rules

1. **Use lowercase with hyphens** for file names:
   - ` + backtick + `api-spec.json` + backtick + `
   - ` + backtick + `data-model.md` + backtick + `
   - NOT: ` + backtick + `ApiSpec.json` + backtick + `
   - NOT: ` + backtick + `data_model.md` + backtick + `

2. **Use descriptive names**:
   - ` + backtick + `user-authentication-plan.md` + backtick + `
   - NOT: ` + backtick + `plan1.md` + backtick + `

3. **Use consistent extensions**:
   - Markdown: ` + backtick + `*.md` + backtick + `
   - JSON: ` + backtick + `*.json` + backtick + `
   - YAML: ` + backtick + `*.yml` + backtick + ` or ` + backtick + `*.yaml` + backtick + `
   - HTML: ` + backtick + `*.html` + backtick + `

### Phase and Feature Naming

1. **Phase directories**: ` + backtick + `XX-phase` + backtick + ` (e.g., ` + backtick + `01-phase` + backtick + `, ` + backtick + `02-phase` + backtick + `)
   - Always use two-digit numbers with leading zeros
   - Always use lowercase ` + backtick + `phase` + backtick + `

2. **Feature directories**: ` + backtick + `XX-Feature` + backtick + ` (e.g., ` + backtick + `01-Feature` + backtick + `, ` + backtick + `02-Feature` + backtick + `)
   - Always use two-digit numbers with leading zeros
   - Capitalize ` + backtick + `Feature` + backtick + `

3. **Feature-specific files**: Always lowercase with hyphens
   - ` + backtick + `plan.md` + backtick + `
   - ` + backtick + `design.md` + backtick + `
   - ` + backtick + `tasks.md` + backtick + `
   - ` + backtick + `progress.json` + backtick + `

### Document File Naming

| Document Type | File Name | Location | Description |
|--------------|-----------|----------|-------------|
| Product Requirements | ` + backtick + `PRD.md` + backtick + ` | ` + backtick + `doplan/` + backtick + ` | Product Requirements Document |
| Project Structure | ` + backtick + `structure.md` + backtick + ` | ` + backtick + `doplan/` + backtick + ` | Architecture and structure |
| Tech Stack Context | ` + backtick + `CONTEXT.md` + backtick + ` | Root | Technology stack with docs links |
| Dashboard | ` + backtick + `dashboard.md` + backtick + ` / ` + backtick + `dashboard.html` + backtick + ` | ` + backtick + `doplan/` + backtick + ` | Progress visualization |
| Idea Notes | ` + backtick + `idea-notes.md` + backtick + ` | ` + backtick + `doplan/` + backtick + ` | Initial idea discussion |
| API Specification | ` + backtick + `api-spec.json` + backtick + ` | ` + backtick + `doplan/contracts/` + backtick + ` | OpenAPI/Swagger spec |
| Data Models | ` + backtick + `data-model.md` + backtick + ` | ` + backtick + `doplan/contracts/` + backtick + ` | Data structure documentation |
| Phase Plan | ` + backtick + `phase-plan.md` + backtick + ` | ` + backtick + `doplan/XX-phase/` + backtick + ` | Phase planning document |
| Phase Progress | ` + backtick + `phase-progress.json` + backtick + ` | ` + backtick + `doplan/XX-phase/` + backtick + ` | Phase progress tracking |
| Feature Plan | ` + backtick + `plan.md` + backtick + ` | ` + backtick + `doplan/XX-phase/XX-Feature/` + backtick + ` | Feature planning |
| Feature Design | ` + backtick + `design.md` + backtick + ` | ` + backtick + `doplan/XX-phase/XX-Feature/` + backtick + ` | Feature design specs |
| Feature Tasks | ` + backtick + `tasks.md` + backtick + ` | ` + backtick + `doplan/XX-phase/XX-Feature/` + backtick + ` | Task breakdown |
| Feature Progress | ` + backtick + `progress.json` + backtick + ` | ` + backtick + `doplan/XX-phase/XX-Feature/` + backtick + ` | Feature progress |

## Documentation Content Standards

### Markdown Files

1. **Always include frontmatter** (for IDE context):
   ` + codeBlock + `yaml
   ---
   title: Document Title
   description: Brief description
   last_updated: YYYY-MM-DD
   ---
   ` + codeBlock + `

2. **Use consistent heading hierarchy**:
   - ` + backtick + `#` + backtick + ` - Document title
   - ` + backtick + `##` + backtick + ` - Main sections
   - ` + backtick + `###` + backtick + ` - Subsections
   - ` + backtick + `####` + backtick + ` - Sub-subsections

3. **Include table of contents** for long documents:
   ` + codeBlock + `markdown
   ## Table of Contents
   - [Section 1](#section-1)
   - [Section 2](#section-2)
   ` + codeBlock + `

### JSON Files

1. **Use consistent structure**:
   - Always include ` + backtick + `version` + backtick + ` field
   - Use descriptive property names
   - Include ` + backtick + `last_updated` + backtick + ` timestamp

2. **Progress JSON structure**:
   ` + codeBlock + `json
   {
     "version": "1.0.0",
     "status": "in-progress",
     "progress": 45,
     "completed_tasks": 9,
     "total_tasks": 20,
     "last_updated": "2024-01-15T10:30:00Z"
   }
   ` + codeBlock + `

## CONTEXT.md Structure

The ` + backtick + `CONTEXT.md` + backtick + ` file provides comprehensive tech stack information for IDEs and CLIs.

### Required Sections

1. **Programming Languages**
   - Language name
   - Version
   - Official documentation link
   - Key features used

2. **Frameworks**
   - Framework name
   - Version
   - Official documentation link
   - Purpose in project

3. **CLIs and Tools**
   - Tool name
   - Version
   - Installation command
   - Documentation link
   - Usage examples

4. **Services**
   - Service name
   - Provider
   - Documentation link
   - Configuration details

5. **Databases**
   - Database type
   - Version
   - Documentation link
   - Connection details

6. **Development Tools**
   - IDE/Editor
   - Extensions/Plugins
   - Configuration files

### CONTEXT.md Format

` + codeBlock + `markdown
# Project Technology Stack

## Programming Languages

### Go
- **Version:** 1.24.0
- **Documentation:** https://go.dev/doc/
- **Usage:** Backend CLI development
- **Key Features:** Goroutines, Channels, Interfaces

## Frameworks

### Cobra
- **Version:** v1.8.0
- **Documentation:** https://github.com/spf13/cobra
- **Purpose:** CLI command structure

## CLIs and Tools

### DoPlan CLI
- **Version:** 1.0.0
- **Installation:** ` + backtick + `go install github.com/DoPlan-dev/CLI@latest` + backtick + `
- **Documentation:** https://github.com/DoPlan-dev/CLI
- **Usage:** ` + backtick + `doplan install` + backtick + `

## Services

### GitHub
- **Provider:** GitHub
- **Documentation:** https://docs.github.com/
- **Purpose:** Version control and CI/CD

## Databases

### SQLite (if applicable)
- **Version:** 3.x
- **Documentation:** https://www.sqlite.org/docs.html
- **Usage:** Local development database

## Development Tools

### Cursor IDE
- **Extensions:** DoPlan commands
- **Configuration:** ` + backtick + `.cursor/rules/` + backtick + `, ` + backtick + `.cursor/commands/` + backtick + `
` + codeBlock + `

## Documentation Generation Rules

### When to Generate CONTEXT.md

1. **During installation** (` + backtick + `doplan install` + backtick + `):
   - Automatically generate initial ` + backtick + `CONTEXT.md` + backtick + `
   - Include detected technologies from project files

2. **During /Discuss command**:
   - Update ` + backtick + `CONTEXT.md` + backtick + ` with recommended tech stack
   - Add documentation links for selected technologies

3. **During /Generate command**:
   - Update ` + backtick + `CONTEXT.md` + backtick + ` with finalized tech stack
   - Include all technologies mentioned in PRD

4. **Manual updates**:
   - Update when adding new dependencies
   - Update when changing tech stack
   - Keep documentation links current

### Auto-Detection Rules

When generating ` + backtick + `CONTEXT.md` + backtick + `, detect technologies from:

1. **Package managers**:
   - ` + backtick + `go.mod` + backtick + ` → Go, Go modules
   - ` + backtick + `package.json` + backtick + ` → Node.js, npm packages
   - ` + backtick + `requirements.txt` + backtick + ` → Python, pip packages
   - ` + backtick + `Cargo.toml` + backtick + ` → Rust, Cargo packages
   - ` + backtick + `pom.xml` + backtick + ` → Java, Maven
   - ` + backtick + `Gemfile` + backtick + ` → Ruby, Bundler

2. **Configuration files**:
   - ` + backtick + `.git/config` + backtick + ` → Git version
   - ` + backtick + `docker-compose.yml` + backtick + ` → Docker, services
   - ` + backtick + `.github/workflows/` + backtick + ` → GitHub Actions
   - ` + backtick + `tsconfig.json` + backtick + ` → TypeScript
   - ` + backtick + `webpack.config.js` + backtick + ` → Webpack

3. **Project structure**:
   - ` + backtick + `src/` + backtick + ` → Common source structure
   - ` + backtick + `api/` + backtick + ` → API-related code
   - ` + backtick + `frontend/` + backtick + ` → Frontend framework
   - ` + backtick + `backend/` + backtick + ` → Backend framework

## Documentation Maintenance

### Update Frequency

- **CONTEXT.md**: Update when tech stack changes
- **PRD.md**: Update when requirements change
- **structure.md**: Update when architecture changes
- **Progress files**: Update automatically via /Progress command
- **Dashboard**: Update automatically via /Dashboard command

### Version Control

- All documentation files should be committed to Git
- Use conventional commit messages: ` + backtick + `docs: update CONTEXT.md` + backtick + `
- Keep documentation in sync with code changes
- Review documentation during PR reviews

## IDE/CLI Integration

### Cursor IDE

- Rules in: ` + backtick + `.cursor/rules/` + backtick + `
- Commands in: ` + backtick + `.cursor/commands/` + backtick + `
- Context file: ` + backtick + `CONTEXT.md` + backtick + ` (auto-loaded)

### Gemini CLI

- Commands in: ` + backtick + `.gemini/commands/` + backtick + `
- Context file: ` + backtick + `CONTEXT.md` + backtick + ` (reference in prompts)

### Claude Code

- Commands in: ` + backtick + `.claude/commands/` + backtick + `
- Context file: ` + backtick + `CONTEXT.md` + backtick + ` (reference with @CONTEXT.md)

### Codex CLI

- Prompts in: ` + backtick + `.codex/prompts/` + backtick + `
- Context file: ` + backtick + `CONTEXT.md` + backtick + ` (reference in prompts)

### OpenCode

- Commands in: ` + backtick + `.opencode/command/` + backtick + `
- Context file: ` + backtick + `CONTEXT.md` + backtick + ` (reference with @CONTEXT.md)

### Qwen Code

- Commands in: ` + backtick + `.qwen/commands/` + backtick + `
- Context file: ` + backtick + `CONTEXT.md` + backtick + ` (reference in prompts)

## Best Practices

1. **Keep documentation up-to-date**:
   - Update when code changes
   - Review during PR process
   - Automate updates where possible

2. **Use consistent formatting**:
   - Follow markdown best practices
   - Use consistent heading levels
   - Include code examples

3. **Link to official documentation**:
   - Always include official docs links
   - Keep links current
   - Verify links periodically

4. **Make documentation discoverable**:
   - Use clear file names
   - Include in README.md
   - Reference in IDE commands

5. **Automate documentation**:
   - Generate from code when possible
   - Use templates for consistency
   - Update via CLI commands
`
}

