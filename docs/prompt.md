# DoPlan CLI - Complete Recreation Prompt

You are an elite Go engineer in November 2025. Your task is to build **DoPlan CLI** - a zero-install, pure-Go CLI that instantly creates professional 2025 projects with a full hierarchical AI agency controlled by real-English slash commands.

## 🎯 Project Overview

**DoPlan CLI** is a single-binary Go application (< 15MB) that generates complete project structures with:
- 18+ hierarchical AI agents (Project Orchestrator, Product Manager, Engineers, Designers, QA, etc.)
- 1000+ embedded rules library covering all major tech stacks
- Real-English slash commands (`/tell`, `/write`, `/build`, etc.)
- Automated GitHub workflows (CI/CD, releases, changelog)
- Support for 6 AI-powered IDEs (Cursor, Claude Code, Antigravity, Windsurf, Cline, OpenCode)
- Beautiful interactive TUI built with Bubbletea

## 📋 Core Requirements

### Language & Distribution
- **100% Go** - No Node.js, no npm dependencies (except npx wrapper for distribution)
- **Single binary** - Must be < 15MB when compiled
- **Zero LLM in CLI** - All intelligence lives in generated `.cursor/commands/` markdown files
- **Offline capable** - Works completely offline after first run
- **Distribution**: Via `npx doplan@latest` (npx wrapper downloads Go binary)

### Project Generation
When run, DoPlan must:
1. Ask user: project name, IDE choice (project type defaults to "Fullstack")
2. Generate complete project structure (see Project Structure below)
3. Extract embedded rules library to `.cursor/rules/library/`
4. Generate all agent personas in `.cursor/agents/`
5. Generate all command definitions in `.cursor/commands/`
6. Generate GitHub Actions workflows in `.github/workflows/`
7. Generate boilerplate code in `src/` based on project type
8. Generate IDE-specific configuration files
9. Print: "Open with: code ./project-name" and "Then type /tell to begin"

## 🏗️ Project Structure to Generate

```
my-project/
├── .cursor/
│   ├── agents/              # 18+ agent persona markdown files
│   ├── commands/            # Command definition markdown files
│   └── rules/              # Rules library
│       ├── README.md       # Rules overview
│       └── library/        # Embedded rules (15 categories, 1000+ files)
│           ├── 01-core-workflow/
│           ├── 02-ai-agents/
│           ├── 03-languages/
│           ├── 04-frameworks/
│           ├── 05-ui-libraries/
│           ├── 06-cloud-infrastructure/
│           ├── 07-databases/
│           ├── 08-testing/
│           ├── 09-devops-ci-cd/
│           ├── 10-code-quality/
│           ├── 11-documentation/
│           ├── 12-security/
│           ├── 13-development-practices/
│           ├── 14-mcp-tools/
│           └── 15-project-specific/
├── .plan/
│   ├── 00_System/
│   │   ├── IDEA.md
│   │   ├── BRAINSTORM.md
│   │   ├── PRD.md
│   │   ├── ARCHITECTURE.md
│   │   └── DESIGN_SYSTEM.md
│   ├── TASKS.md
│   └── active_state.json
├── .github/
│   └── workflows/
│       ├── ci.yml          # Runs on all branches
│       ├── release.yml      # Automated releases
│       ├── changelog.yml   # Auto-updates changelog
│       └── branch-protection.yml
├── src/                    # Boilerplate code (Next.js, Tauri, Expo, etc.)
├── CHANGELOG.md
├── STANDUP.md
├── README.md
└── [IDE config files]     # .cursorrules, CLAUDE.md, etc.
```

## 🤖 AI Agent System

### 18 Base Agents (Hierarchical Structure)

```
Project Orchestrator (CEO/Engineering Manager)
├── Product Manager
├── Engineering Lead
│   ├── System Architect
│   ├── Frontend Lead
│   ├── Backend Lead
│   ├── DevOps Engineer
│   ├── Security Lead
│   └── Performance Engineer
├── Design & UX Manager
│   └── UI/UX Designer
├── QA & Reliability Manager
│   └── QA Engineer
├── Release & Growth Manager
│   ├── Release Captain
│   └── Growth Coach
└── Documentation Lead
    └── Documentation Writer
```

### Agent File Format
Each agent file (`.cursor/agents/{name}.md`) contains:
- **Role**: Brief role description
- **System Prompt**: Detailed persona and responsibilities
- **Responsibilities**: List of duties
- **Reports To**: Manager in hierarchy
- **Manages**: Subordinates (if any)

### Agent Responsibilities

**Project Orchestrator**: Ultimate decision maker, coordinates all teams, strategic vision

**Product Manager**: Defines product requirements, creates PRD.md, prioritizes features

**Engineering Lead**: Technical architecture, code quality, team management

**System Architect**: System design, technology selection, scalability planning

**Frontend Lead**: Frontend architecture, component design, state management

**Backend Lead**: Backend architecture, API design, database design

**DevOps Engineer**: Infrastructure, deployment, CI/CD setup

**Security Lead**: Security reviews, vulnerability assessments

**Design & UX Manager**: Design strategy, UX planning

**UI/UX Designer**: Interface design, component library

**QA & Reliability Manager**: Testing strategy, quality assurance

**QA Engineer**: Test execution, bug finding

**Release & Growth Manager**: Release planning, growth strategy

**Release Captain**: Release coordination, version control, GitHub automation, changelog management

**Growth Coach**: User onboarding, growth optimization

**Documentation Lead**: Documentation strategy, standards

**Documentation Writer**: Technical writing, API docs

**Performance Engineer**: Performance optimization, profiling

## 📝 Command System

### Core Commands (11 commands)

1. **`/tell`** - Capture project idea
   - Saves to `.plan/00_System/IDEA.md`
   - Activates Project Orchestrator

2. **`/improve`** - Team brainstorm session
   - All Level 1 managers brainstorm
   - Updates `.plan/00_System/BRAINSTORM.md`

3. **`/team`** - Show active agents and hierarchy
   - Lists all agents with their roles

4. **`/write`** - Generate PRD + ARCHITECTURE + DESIGN_SYSTEM
   - Product Manager → PRD.md
   - Engineering Lead → ARCHITECTURE.md
   - Design Manager → DESIGN_SYSTEM.md

5. **`/change`** - Edit any document
   - Example: `/change prd Add dark mode`
   - Updates specified document

6. **`/good`** - Approve & lock plan
   - Sets `locked: true` in active_state.json
   - Prepares for task generation

7. **`/tasks`** - Generate TASKS.md
   - Creates implementation tasks from approved plan
   - Organized by phases

8. **`/load`** - Inject context into AI agents
   - Loads rules or files into agent context
   - Example: `/load @library/04-frameworks/frontend/nextjs.md`

9. **`/build`** - Start coding next task
   - `/build` - Next uncompleted task
   - `/build 3` - Specific task
   - Activates Engineering Lead + relevant team leads
   - Updates active_state.json

10. **`/progress`** - Show current progress
    - Reads TASKS.md and active_state.json
    - Shows completion percentage, current task

11. **`/finished`** - Mark current task done
    - Marks task complete in TASKS.md
    - Updates active_state.json
    - **Auto-commits** with conventional commit format
    - **Auto-pushes** to current branch
    - Updates CHANGELOG.md if significant
    - Triggers CI workflow

### Squad-Specific Commands
- `/secure` - Security review (Security Lead)
- `/roles` - Design RBAC system
- `/money` - Billing & payment setup
- `/pretty` - UI/UX improvements
- `/seo` - SEO optimization
- `/ship` - Release management (Release Captain)
- `/safe` - Security audit
- `/cheap` - Cost optimization

### Command File Format
Each command file (`.cursor/commands/{name}.md`) contains:
- **Trigger**: Exact match pattern
- **Action**: Step-by-step what happens
- **Agent Involvement**: Which agents are activated
- **Files Read**: What files are read
- **Files Modified**: What gets changed
- **GitHub Automation**: Auto-commit/push behavior (if applicable)

## 🔄 GitHub Workflows

### CI Workflow (`ci.yml`)
- **Triggers**: Push to any branch (`branches: ['*']`)
- **Jobs**: 
  - Test: Install deps, lint, test, build
  - Lint: Separate linting job
- **Must pass** before merging to main/develop

### Release Workflow (`release.yml`)
- **Triggers**: Version tags (`v*.*.*`) or manual dispatch
- **Actions**:
  - Extract version from tag
  - Update package.json version
  - Generate release notes from CHANGELOG.md
  - Create GitHub release
  - Commit version bump

### Changelog Workflow (`changelog.yml`)
- **Triggers**: Push to main/develop with CHANGELOG.md changes
- **Actions**: Auto-commit changelog updates

### Branch Protection Workflow (`branch-protection.yml`)
- **Triggers**: Pull requests
- **Actions**: Enforces PR requirements

## 📚 Rules Library

### Embedded Rules Structure
Rules are embedded using Go's `//go:embed` directive:
```go
//go:embed library/*
var rulesFS embed.FS
```

### Rules Categories (15 categories)
1. **01-core-workflow/** - Core development workflow rules
   - Includes: `github-workflow-automation.md`, `continuous-github-workflows.md`, `command-agent-integration.md`
2. **02-ai-agents/** - AI agent interaction guidelines
3. **03-languages/** - Language-specific rules (Go, Python, TypeScript, etc.)
4. **04-frameworks/** - Framework rules (Next.js, React, Express, etc.)
5. **05-ui-libraries/** - UI library rules
6. **06-cloud-infrastructure/** - Cloud platform rules
7. **07-databases/** - Database rules
8. **08-testing/** - Testing framework rules
9. **09-devops-ci-cd/** - DevOps and CI/CD rules
10. **10-code-quality/** - Code quality rules
11. **11-documentation/** - Documentation standards
12. **12-security/** - Security best practices
13. **13-development-practices/** - General development practices
14. **14-mcp-tools/** - MCP tool integration
15. **15-project-specific/** - Project-specific rules

### Rules Extraction
Rules are extracted to `.cursor/rules/library/` maintaining exact directory structure.

## 🎨 Tech Stack & Boilerplate

### Default Project Type: Fullstack
Generates Next.js 15.2.1 + React 19 + TypeScript 5.6.0 boilerplate.

### Package Versions (as of Nov 2025)
- Next.js: 15.2.1
- React: 19.0.0
- TypeScript: 5.6.0
- Tailwind CSS: 3.4.10
- Node.js: 22

### Boilerplate Includes
- `package.json` with latest versions
- `tsconfig.json` configured
- `tailwind.config.ts` configured
- Basic Next.js app structure in `src/`
- ESLint configuration

## 💻 IDE Support

### Supported IDEs
1. **Cursor** - `.cursorrules` file
2. **Claude Code** - `CLAUDE.md` file
3. **Antigravity** - `.antigravity/config.md` file
4. **Windsurf** - `.windsurfrules` file
5. **Cline** - `.cline/config.md` file
6. **OpenCode** - `CONTEXT.md` file

### IDE Config Content
Each config file should:
- Reference agent hierarchy
- Reference command locations
- Reference rules library
- Include workflow automation notes
- Include command-agent integration notes

## 🛠️ Implementation Details

### Go Module Structure
```
doplan/
├── cmd/doplan/
│   └── main.go              # Entry point
├── internal/
│   ├── cli/
│   │   └── root.go          # Cobra CLI setup
│   ├── tui/
│   │   └── wizard.go        # Bubbletea TUI wizard
│   ├── generator/
│   │   ├── generator.go     # Main orchestration
│   │   ├── agents.go         # Agent generation
│   │   ├── commands.go       # Command generation
│   │   ├── plan.go          # .plan/ structure
│   │   ├── ide.go           # IDE configs
│   │   ├── boilerplate.go   # Source code boilerplate
│   │   ├── github.go        # GitHub workflows
│   │   └── docs.go          # README, STANDUP, rules README
│   └── rules/
│       ├── rules.go         # Rules extraction (embed.FS)
│       └── library/         # Embedded rules (1000+ files)
└── pkg/models/
    └── project.go           # Data models (ProjectRequest, ProjectState)
```

### Key Dependencies
```go
require (
    github.com/charmbracelet/bubbletea v1.3.4
    github.com/charmbracelet/lipgloss v1.1.0
    github.com/spf13/cobra v1.8.0
)
```

### TUI Wizard Flow
1. **Step 0**: Ask project name (text input)
2. **Step 1**: Ask IDE choice (selection menu)
3. **Generate**: Create project with default type "Fullstack"

### Project Generation Flow
1. Create project directory
2. Change into directory
3. Generate directory structure
4. Generate agents (18 base agents)
5. Generate commands (11 core + squad-specific)
6. Generate .plan/ structure
7. Generate IDE config
8. Generate boilerplate
9. Extract rules library
9. Generate GitHub workflows
10. Generate README.md
11. Generate STANDUP.md
12. Generate CHANGELOG.md
13. Generate active_state.json

### State Management
`active_state.json` structure:
```json
{
  "phase": "idea|brainstorm|writing|approved|tasks|building",
  "active_task": null | task_id,
  "completed": [task_ids],
  "locked": false | true
}
```

## 📋 Key Rules to Implement

### GitHub Workflow Automation
- CI runs on every push to any branch
- Commits use conventional commit format
- Auto-commit and push on `/finished`
- Release automation on version tags
- Changelog auto-updates

### Command-Agent Integration
- Each command clearly defines agent involvement
- Agents coordinate through hierarchy
- Context loaded before agent activation
- State updates are atomic

### Continuous Development
- Workflows must pass before merging
- Clear error messages
- Status monitoring
- Pre-commit checks

## 🎯 Success Criteria

The CLI is complete when:
1. ✅ Single binary compiles to < 15MB
2. ✅ Interactive TUI wizard works (Bubbletea)
3. ✅ Generates complete project structure
4. ✅ All 18 agents generated with proper hierarchy
5. ✅ All 11 core commands generated
6. ✅ Rules library extracted correctly
7. ✅ GitHub workflows generated and functional
8. ✅ IDE configs generated for all 6 IDEs
9. ✅ Boilerplate code generated (Next.js default)
10. ✅ Works offline after first run
11. ✅ All files use proper markdown formatting
12. ✅ No hardcoded secrets or sensitive data

## 🚀 Build & Test

### Build Command
```bash
go build -o doplan cmd/doplan/main.go
```

### Test Command
```bash
go test ./...
```

### Binary Size Check
```bash
ls -lh doplan  # Should be < 15MB
```

## 📝 Important Notes

1. **No LLM in CLI**: All intelligence is in generated markdown files
2. **Offline First**: Everything embedded, no network calls after install
3. **Beginner Friendly**: Clear documentation, simple commands
4. **Production Ready**: Follows Go best practices, proper error handling
5. **Extensible**: Easy to add new agents, commands, or rules
6. **Well Documented**: Code comments, clear structure

## 🎨 UI/UX Requirements

### TUI Wizard
- Beautiful header art with emojis
- Color-coded prompts (purple, pink, green, yellow)
- Keyboard navigation (↑/↓, Enter, q to quit)
- Clear help text
- Smooth transitions

### Output Messages
- Success: Green checkmarks ✅
- Info: Blue icons ℹ️
- Warnings: Yellow ⚠️
- Errors: Red ❌

## 🔒 Security & Best Practices

- No secrets in code
- Input validation
- Proper error handling
- File permission checks
- Path sanitization
- No arbitrary code execution

## 📚 Reference Files

When implementing, reference:
- `internal/generator/agents.go` - Agent definitions and prompts
- `internal/generator/commands.go` - Command definitions
- `internal/generator/github.go` - Workflow templates
- `internal/rules/rules.go` - Rules extraction logic
- `internal/tui/wizard.go` - TUI implementation
- `pkg/models/project.go` - Data structures

## 🎯 Final Checklist

Before considering complete:
- [ ] All 18 agents generated with proper hierarchy
- [ ] All 11 core commands generated
- [ ] Rules library extracted (1000+ files)
- [ ] GitHub workflows generated (4 workflows)
- [ ] IDE configs for all 6 IDEs
- [ ] Boilerplate code generated
- [ ] TUI wizard works smoothly
- [ ] Binary size < 15MB
- [ ] Works offline
- [ ] No compilation errors
- [ ] All markdown files properly formatted
- [ ] README.md comprehensive and beginner-friendly

---

**Start building now. Create a production-ready, beautiful, magical CLI that developers will love to use.**

