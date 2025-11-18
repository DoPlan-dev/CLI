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

// Generate creates all rule files in .doplan/ai/rules/
func (g *RulesGenerator) Generate() error {
	// Generate to central location .doplan/ai/rules/
	rulesDir := filepath.Join(g.projectRoot, ".doplan", "ai", "rules")

	if err := os.MkdirAll(rulesDir, 0755); err != nil {
		return err
	}

	rules := map[string]string{
		"workflow.mdc":      generateWorkflowRules(),
		"communication.mdc": generateCommunicationRules(),
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

func generateCommunicationRules() string {
	backtick := "`"
	return `# DoPlan Communication Rules

These rules govern how AI agents **MUST** communicate and collaborate within the DoPlan workflow.

## ⚠️ CRITICAL: Agent Handoff Protocol

**When your work is complete, you MUST tag the next agent.** This ensures the workflow continues smoothly.

## Communication Principles

- Be clear and concise in all communications
- Reference specific files and line numbers when discussing code
- Use the DoPlan workflow stages appropriately
- Always update state and progress files after actions
- Follow the established naming conventions
- **Tag the next agent when work is complete**
- **Wait for previous agents to finish before starting**

## Agent Collaboration & Handoff

### Workflow Sequence & Handoffs

1. **@planner** → **@designer**
   - @planner completes planning
   - Tags @designer: "Planning complete. @designer please create design specifications."
   - @designer waits for @planner to finish

2. **@designer** → **@coder**
   - @designer completes design
   - Tags @coder: "Design complete. @coder please begin implementation."
   - @coder waits for @planner AND @designer to finish

3. **@coder** → **@tester**
   - @coder completes implementation
   - Tags @tester: "Implementation complete. @tester please run tests."
   - @tester ONLY begins when tagged by @coder

4. **@tester** → **@reviewer**
   - @tester completes testing
   - Tags @reviewer: "All tests passed. Screenshot saved. @reviewer please review."
   - @reviewer ONLY reviews AFTER @tester has successfully run all tests

5. **@reviewer** → **@devops** OR **@coder**
   - If approved: Tags @devops: "Code review passed. @devops please deploy."
   - If changes needed: Tags @coder: "Changes needed. @coder please address."
   - @devops ONLY begins AFTER @reviewer has approved

6. **@devops** → **Team**
   - @devops completes deployment
   - Reports to team: "Deployment successful" or "Deployment failed"

## Agent Roles & Responsibilities

### @planner
- **Role:** Senior project planner
- **Starts:** First (project start)
- **Hands off to:** @designer
- **Must complete:** Planning, PRD, folder structure

### @designer
- **Role:** UI/UX specialist
- **Starts:** After @planner completes
- **Hands off to:** @coder
- **Must complete:** Design specifications following design_rules.mdc

### @coder
- **Role:** Implementation specialist
- **Starts:** After @planner AND @designer complete
- **Hands off to:** @tester
- **Must complete:** Feature implementation

### @tester
- **Role:** QA & Test Automation Specialist
- **Starts:** When tagged by @coder
- **Hands off to:** @reviewer
- **Must complete:** Tests, screenshots, test report

### @reviewer
- **Role:** Quality assurance
- **Starts:** AFTER @tester has successfully run all tests
- **Hands off to:** @devops (if approved) OR @coder (if changes needed)
- **Must complete:** Code review

### @devops
- **Role:** Deployment and infrastructure specialist
- **Starts:** AFTER @reviewer has approved
- **Hands off to:** Team (reports status)
- **Must complete:** Deployment

## Tagging Format

**When tagging the next agent:**

` + backtick + backtick + backtick + `
✅ [Your work] complete for {feature-name}. [Brief summary]. @{next-agent} please [next action].
` + backtick + backtick + backtick + `

**Examples:**
- "✅ Planning complete. Created 3 phases with 8 features. @designer please create design specifications."
- "✅ Design complete. Design follows DPR and design_rules.mdc. @coder please begin implementation."
- "✅ Implementation complete. All tasks done. @tester please run tests."
- "✅ All tests passed. Screenshot saved to .doplan/artifacts/screenshots/01-auth/01-login.png. @reviewer please review."
- "✅ Code review passed. All requirements met. @devops please deploy."
- "❌ Code review: Changes needed. Issues: [list]. @coder please address."

## File References

When discussing code or files:
- Use absolute paths: ` + backtick + `doplan/01-user-authentication/01-login-with-email/plan.md` + backtick + `
- Reference line numbers when needed
- Include file context in discussions
- Use numbered folder structure

## State Updates

After any action:
- Update relevant progress.json files
- Update state file if applicable
- Regenerate dashboard if progress changed
- Commit changes with clear messages
- **Tag the next agent**

## Error Handling

**If you encounter issues:**
- Tag the appropriate agent for clarification
- Document the issue
- Provide context and error details
- Request specific help

**Examples:**
- "@planner: Need clarification on {feature}. Question: {question}."
- "@designer: Design clarification needed. Issue: {issue}."
- "@coder: Implementation issue. Error: {error}. @coder please help."

## Multi-Agent Conversations

Agents can work together in conversations. Always:
1. Tag agents explicitly
2. Provide context
3. Reference relevant files
4. Follow the workflow sequence
5. Wait for responses before proceeding

## Best Practices

- Always tag the next agent when work is complete
- Wait for previous agents to finish
- Provide clear context when tagging
- Reference specific files and line numbers
- Update progress files after actions
- Follow the workflow sequence strictly
- Be respectful and constructive in feedback
`
}

func generateWorkflowRules() string {
	backtick := "`"
	return `# DoPlan Workflow Rules

These rules define the **perfect workflow sequence** that all AI agents MUST follow.

## ⚠️ CRITICAL: Perfect Workflow Sequence

The DoPlan workflow follows this **exact sequence**. **DO NOT skip steps or work out of order:**

1. **Plan** → @planner creates project structure
2. **Design** → @designer creates design specifications
3. **Code** → @coder implements features
4. **Test** → @tester runs tests and captures screenshots
5. **Review** → @reviewer reviews code quality
6. **Deploy** → @devops handles deployment

## Core Principles

- Follow Spec-Kit and BMAD-METHOD methodologies
- Generate contracts before implementation
- Track progress in dashboard
- Use feature branches for each feature
- All commands are available as slash commands in your IDE/CLI
- **Agents MUST wait for previous agents to complete**
- **Agents MUST tag the next agent when work is complete**

## Workflow Stages

### 1. Plan → @planner (FIRST STEP)

**Agent:** @planner  
**When:** Project start or new feature  
**Commands:** /Discuss, /Refine, /Generate, /Plan

**Process:**
1. **Idea Discussion (/Discuss)**
   - Ask comprehensive questions about the idea
   - Suggest improvements and enhancements
   - Help organize features into phases
   - Recommend tech stack based on requirements
   - Save results to state file and ` + backtick + `doplan/idea-notes.md` + backtick + `

2. **Idea Refinement (/Refine)**
   - Review existing idea notes
   - Suggest additional features
   - Identify gaps in the plan
   - Enhance technical specifications
   - Update idea documentation

3. **Document Generation (/Generate)**
   - Create ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
   - Create ` + backtick + `doplan/structure.md` + backtick + ` - Project structure
   - Create ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification
   - Create ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models
   - Use templates from ` + backtick + `doplan/templates/` + backtick + `

4. **Planning (/Plan)**
   - **CRITICAL:** Create phase directories using numbered and slugified names: ` + backtick + `doplan/01-{phase-name}/` + backtick + `
   - **CRITICAL:** Create feature directories: ` + backtick + `doplan/01-{phase-name}/01-{feature-name}/` + backtick + `
   - Generate ` + backtick + `plan.md` + backtick + `, ` + backtick + `design.md` + backtick + ` (placeholder), ` + backtick + `tasks.md` + backtick + ` for each feature
   - Create ` + backtick + `phase-plan.md` + backtick + ` and ` + backtick + `phase-progress.json` + backtick + ` for each phase
   - Update dashboard with new structure
   - **Tag @designer** to begin design work

### 2. Design → @designer

**Agent:** @designer  
**When:** After @planner completes planning  
**Commands:** /Design, /Design:Review

**Process:**
1. Read ` + backtick + `plan.md` + backtick + ` and PRD
2. Read ` + backtick + `doplan/design/DPR.md` + backtick + ` - Design Preferences & Requirements
3. Read ` + backtick + `.doplan/ai/rules/design_rules.mdc` + backtick + ` - Design rules
4. Create comprehensive ` + backtick + `design.md` + backtick + ` following design system
5. Use design tokens from ` + backtick + `doplan/design/design-tokens.json` + backtick + `
6. **Tag @coder** to begin implementation

### 3. Code → @coder

**Agent:** @coder  
**When:** After @planner AND @designer complete  
**Commands:** /Implement

**Process:**
1. Verify @planner and @designer have completed
2. Read ` + backtick + `plan.md` + backtick + `, ` + backtick + `design.md` + backtick + `, ` + backtick + `tasks.md` + backtick + `
3. Automatically create GitHub branch: ` + backtick + `feature/{##}-{phase-name}-{##}-{feature-name}` + backtick + `
4. Implement features following plan and design
5. Follow ` + backtick + `.doplan/ai/rules/design_rules.mdc` + backtick + ` for UI work
6. Update progress as tasks complete
7. **Tag @tester** when work is ready for testing

### 4. Test → @tester

**Agent:** @tester  
**When:** Tagged by @coder  
**Commands:** /Test, /Test:Visual

**Process:**
1. Read ` + backtick + `plan.md` + backtick + ` and ` + backtick + `design.md` + backtick + `
2. Generate test scenarios from acceptance criteria
3. Write and execute automated tests using Playwright (MCP)
4. **CRITICAL:** Capture screenshots to ` + backtick + `.doplan/artifacts/screenshots/{phase-name}/{feature-name}.png` + backtick + `
5. Perform visual regression checks
6. Report bugs with screenshots if found
7. **Tag @reviewer** with test report (pass or fail)

### 5. Review → @reviewer

**Agent:** @reviewer  
**When:** AFTER @tester has successfully run all tests  
**Commands:** /Review

**Process:**
1. Verify @tester has run all tests and they pass
2. Review implementation against ` + backtick + `plan.md` + backtick + ` and ` + backtick + `design.md` + backtick + `
3. Check code quality and standards
4. Verify design system compliance
5. **If approved:** Tag @devops for deployment
6. **If changes needed:** Tag @coder with feedback

### 6. Deploy → @devops

**Agent:** @devops  
**When:** AFTER @reviewer has approved  
**Commands:** /Deploy, /Deploy:Configure

**Process:**
1. Verify @reviewer has approved
2. Check ` + backtick + `doplan/RAKD.md` + backtick + ` for required API keys
3. Configure deployment pipelines
4. Deploy to staging, then production
5. Monitor deployment status
6. Report deployment status to team

## State Management

- State lives in IDE-specific config directory (e.g., ` + backtick + `.cursor/config/doplan-state.json` + backtick + `)
- Progress tracked in ` + backtick + `doplan/**/progress.json` + backtick + ` files
- Dashboard in ` + backtick + `doplan/dashboard.md` + backtick + ` and ` + backtick + `.doplan/dashboard.json` + backtick + `
- Always update state after command execution

## Agent Handoff Protocol

**When your work is complete:**
1. Update progress files
2. Update state file
3. **Tag the next agent** as defined in communication.mdc
4. Provide context about what was completed

**Example:**
- @planner: "✅ Planning complete. Created 3 phases with 8 features. @designer please create design specifications."
- @designer: "✅ Design complete. @coder please begin implementation."
- @coder: "✅ Implementation complete. @tester please run tests."
- @tester: "✅ All tests passed. Screenshot saved. @reviewer please review."
- @reviewer: "✅ Code review passed. @devops please deploy."
- @devops: "✅ Deployment successful."

## Best Practices

- Always start with /Discuss before /Plan
- Generate PRD before creating phases
- Follow the phase → feature hierarchy
- Use numbered and slugified folder names
- Update tasks.md as you work
- Run /Progress regularly
- Use /Next to get recommendations
- **Wait for previous agents to complete**
- **Tag next agent when work is done**
- **Follow the workflow sequence strictly**
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
