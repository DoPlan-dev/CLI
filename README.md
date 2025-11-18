# DoPlan CLI

[![Test](https://github.com/DoPlan-dev/CLI/actions/workflows/test.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/test.yml)
[![Lint](https://github.com/DoPlan-dev/CLI/actions/workflows/lint.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/lint.yml)
[![Build](https://github.com/DoPlan-dev/CLI/actions/workflows/build.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/build.yml)
[![Release](https://github.com/DoPlan-dev/CLI/actions/workflows/release.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/release.yml)
[![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/DoPlan-dev/CLI?label=latest%20release)](https://github.com/DoPlan-dev/CLI/releases/latest)
[![npm version](https://img.shields.io/npm/v/@doplan-dev/cli)](https://www.npmjs.com/package/@doplan-dev/cli)

**DoPlan** is a comprehensive project workflow automation tool that transforms your app ideas into well-structured, documented, and trackable development projects. Built with Go, DoPlan provides a complete command-line interface and integrates seamlessly with popular AI-powered IDEs to guide you through the entire development lifecycle—from initial concept to deployment.

DoPlan automates the tedious aspects of project management by generating documentation, creating project structures, managing Git workflows, tracking progress, and providing visual dashboards. It enforces best practices through automated rules, templates, and checkpoints, ensuring your project stays organized and maintainable throughout its development.

Whether you're building a web application, API service, mobile app, or any software project, DoPlan helps you break down your idea into manageable phases and features, track progress visually, automate GitHub operations, and maintain comprehensive documentation—all while working naturally within your preferred development environment.

## Features

### 🚀 **Workflow Automation**
- **Automated Project Setup**: Initialize your project with a structured workflow in seconds
- **Phase & Feature Management**: Organize your project into phases and features with automatic directory structure
- **Progress Tracking**: Real-time progress monitoring with visual dashboards (markdown and HTML)
- **Document Generation**: Auto-generate PRD, project structure, API contracts, and planning documents
- **Workflow Guidance**: Intelligent recommendations for next steps based on your current progress

### 🎨 **Interactive TUI Interface** *(Enhanced in v0.0.20-beta)*
- **Context-Aware Menu**: Menu adapts to project state (installed vs new projects)
- **Fixed-Width Header**: Clean, centered header design (70 chars)
- **Visual Progress Bars**: See project, phase, and feature progress at a glance
- **Multi-view Dashboard**: Switch between project overview, phases, features, GitHub activity, configuration, and statistics
- **Real-time Updates**: Live refresh of progress, GitHub data, and recommendations
- **API Keys Widget**: Dashboard integration for API key status monitoring

### 🤖 **AI Agents System** *(NEW in v0.0.19-beta)*
- **6 Specialized Agents**: Planner, Designer, Coder, Reviewer, Tester, DevOps
- **Workflow Enforcement**: Agents follow strict workflow sequence (Plan → Design → Code → Test → Review → Deploy)
- **Agent Definitions**: Comprehensive agent files with roles, responsibilities, and communication protocols
- **Workflow Rules**: Automated rules for agent handoffs and task sequencing
- **Communication Protocols**: Tag-based communication system for agent collaboration

### 🎨 **Design System (DPR)** *(NEW in v0.0.19-beta)*
- **Interactive Questionnaire**: Guided questionnaire for design preferences and requirements
- **DPR Document**: Complete Design Preferences & Requirements document generation
- **Design Tokens**: Auto-generated `design-tokens.json` with colors, typography, spacing, and more
- **AI Agent Rules**: `design_rules.mdc` for AI agents to follow design system guidelines
- **Design System Maintenance**: Guidelines for keeping design system current

### 🔐 **Secrets & API Keys Management** *(NEW in v0.0.19-beta)*
- **Service Detection**: Automatically detects required services from project files (package.json, go.mod, etc.)
- **API Key Validation**: Validates API keys in `.env` files with format checking
- **RAKD Document**: Required API Keys Document (RAKD.md) with status and validation results
- **SOPS Guides**: Service Operating Procedures guides for each detected service
- **TUI Management**: Interactive TUI for managing API keys and viewing status

### 🚀 **Deployment & Publishing** *(NEW in v0.0.19-beta)*
- **Multi-Platform Deployment**: Deploy to Vercel, Netlify, Railway, Render, Coolify, Docker
- **Package Publishing**: Publish to npm, Homebrew, Scoop, Winget
- **Auto-Detection**: Automatically detects and recommends deployment platforms
- **Interactive Wizards**: Beautiful TUI wizards for deployment and publishing configuration

### 🛡️ **Security & Auto-Fix** *(NEW in v0.0.19-beta)*
- **Security Scanning**: Comprehensive security scans (npm audit, gosec, git-secrets, trufflehog)
- **Auto-Fix**: Automated fixes for common issues (npm audit fix, go mod tidy, go fmt, ESLint)
- **AI Suggestions**: Placeholder for AI-powered fix suggestions
- **Severity Reporting**: Reports issues by severity level with actionable recommendations

### 🔧 **GitHub Integration**
- **Automatic Branching**: Create feature branches automatically
- **Auto-PR Creation**: Automatically create pull requests when features are complete
- **Commit Management**: Track commits, pushes, and branch status
- **PR Tracking**: Monitor pull request status and URLs

### 📝 **Documentation & Templates** *(Enhanced in v0.0.20-beta)*
- **Organized Documentation**: Restructured docs into guides, implementation, references, and releases
- **Template System**: Customizable templates for plans, designs, and tasks
- **Template Management**: Add, edit, remove, and set default templates
- **Context Generation**: Auto-generate `CONTEXT.md` with tech stack and documentation links
- **Documentation Rules**: Automated rules for documentation organization and naming

### ⚙️ **Configuration & Validation** *(Enhanced in v0.0.20-beta)*
- **Production Readiness Check**: Comprehensive verification script (`make check-production`)
  - 15 categories of checks: dependencies, formatting, static analysis, build, tests, coverage, TODO comments, debug code, git status, version info, documentation, license, build config, test infrastructure, and dependencies
  - Automatic test coverage reporting with CLI-appropriate thresholds
  - Smart code formatting detection
- **Flexible Configuration**: Manage settings for GitHub, checkpoints, and workflow
- **Project Validation**: Validate project structure and configuration with auto-fix
- **Minimal CLI**: Streamlined command set (removed deprecated commands: checkpoint, completion, templates, progress, validate, stats)

### 🔌 **IDE Integration**
- **Multi-IDE Support**: Works with Cursor, VS Code, Gemini CLI, Claude CLI, Codex CLI, OpenCode, Qwen Code, and more
- **Custom Commands**: IDE-specific commands for seamless workflow integration
- **Workflow Rules**: Automated rules and conventions for your development process
- **Setup Wizard**: Interactive TUI wizard for IDE integration setup

## Installation

### Prerequisites

- **Go 1.21+** (for building from source)
- **Git** (for version control)
- **GitHub CLI** (`gh`) - Optional, for GitHub automation features

### npm (Recommended for Node.js users)

```bash
# Install globally
npm install -g @doplan-dev/cli

# Verify installation
doplan --version
```

**Note:** The npm package automatically downloads the correct binary for your platform during installation.

### Homebrew (macOS/Linux)

**Status:** Homebrew tap repository is not yet available. Please use npm or binary releases.

```bash
# Homebrew installation will be available once the tap repository is set up
# For now, please use:
# - npm: npm install -g @doplan-dev/cli
# - Binary releases: See below
```

**Note:** Homebrew tap repository setup is in progress. Once available, installation will be:
```bash
brew tap DoPlan-dev/doplan
brew install doplan
```

### Binary Release

Download the latest release for your platform from [releases](https://github.com/DoPlan-dev/CLI/releases/latest):

```bash
# Linux/macOS - Download and extract
# Visit https://github.com/DoPlan-dev/CLI/releases/latest
# Download the appropriate tar.gz file for your platform
# Then extract and install:
tar -xzf doplan_*_linux_amd64.tar.gz  # or darwin_arm64, etc.
sudo mv doplan /usr/local/bin/

# Windows
# Download doplan_*_windows_amd64.tar.gz from releases page
# Extract and add to PATH
```

**Latest Release:** [v0.0.20-beta](https://github.com/DoPlan-dev/CLI/releases/tag/v0.0.20-beta) | [View all releases](https://github.com/DoPlan-dev/CLI/releases)

### From Source

```bash
# Clone the repository
git clone https://github.com/DoPlan-dev/CLI.git
cd cli

# Build and install
make install

# Or build manually
go build -o bin/doplan ./cmd/doplan
sudo mv bin/doplan /usr/local/bin/
```

### Verify Installation

```bash
doplan --version
```

## CLI Commands

### Main Commands

| Command | Description |
|---------|-------------|
| `doplan` | Launch interactive TUI menu (context-aware: shows Install/Exit for new projects, Development Server/Dashboard/Configuration for installed projects) |
| `doplan install` | Install DoPlan in your project (interactive IDE selection) |
| `doplan dashboard` | View project dashboard with progress and GitHub activity |
| `doplan --tui` | Launch fullscreen interactive TUI dashboard |
| `doplan github` | Sync GitHub data (branches, commits, PRs) and update dashboard |

**Note:** Commands like `progress`, `validate`, `checkpoint`, `completion`, `templates`, and `stats` have been removed in v0.0.20-beta for a more focused, minimal CLI.

### TUI Menu Actions *(Enhanced in v0.0.20-beta)*

The TUI menu is now **context-aware** and adapts based on your project state:

**For New Projects** (not yet installed):
| Action | Description |
|--------|-------------|
| 📦 Install | Install DoPlan in this project |
| ❌ Exit | Exit DoPlan |

**For Installed Projects**:
| Action | Description |
|--------|-------------|
| ▶️ Development Server | Run dev server for testing and developing |
| 📊 Dashboard | Show development progress |
| ⚙️ Configuration | Install files for other IDEs or CLIs |

**Note:** Additional actions (Deploy, Publish, Security Scan, etc.) are available through IDE commands after installation.

### Configuration Commands

| Command | Description |
|---------|-------------|
| `doplan config show` | Show current configuration (table or JSON format) |
| `doplan config set <key> <value>` | Set a configuration value |
| `doplan config reset` | Reset configuration to defaults |

**Configuration Keys:**
- `github.enabled` - Enable/disable GitHub integration
- `github.autoBranch` - Auto-create branches for features
- `github.autoPR` - Auto-create PRs when features complete

### Production Readiness Check *(NEW in v0.0.20-beta)*

Before releasing, verify your codebase is ready:

```bash
# Run comprehensive production readiness check
make check-production

# Or directly
./scripts/check-production-ready.sh
```

The check verifies:
- ✅ Dependencies (Go, Git)
- ✅ Code formatting and static analysis
- ✅ Build and test execution
- ✅ Test coverage (with CLI-appropriate thresholds)
- ✅ TODO/FIXME comments review
- ✅ Debug code detection
- ✅ Git status
- ✅ Version information
- ✅ Documentation completeness
- ✅ License file
- ✅ Build configuration
- ✅ Test infrastructure
- ✅ Dependency verification

See [docs/guides/PRODUCTION_READINESS.md](docs/guides/PRODUCTION_READINESS.md) for details.

## How to Start: Using Commands in Your IDE to Develop Your App Idea

### Step 1: Install DoPlan in Your Project

Navigate to your project directory and run:

```bash
doplan install
```

You'll be prompted to select your IDE:
- **Cursor** - AI-powered code editor
- **Gemini CLI** - Google's Gemini command-line interface
- **Claude CLI** - Anthropic's Claude command-line interface
- **Codex CLI** - OpenAI Codex command-line interface
- **OpenCode** - OpenCode AI development environment
- **Qwen Code** - Qwen3-Coder command-line interface

After installation, DoPlan will:
- Create `.cursor/`, `.gemini/`, `.claude/`, `.codex/`, `.opencode/`, or `.qwen/` directories
- Generate IDE-specific commands
- Create workflow rules and templates
- Set up the project structure

### Step 2: Start Your Idea Discussion

Open your IDE and use the DoPlan commands:

**In Cursor:**
```
/Discuss
```

**In Gemini CLI:**
```
/discuss
```

**In Claude CLI:**
```
/discuss
```

This command will help you:
- Refine your app idea
- Identify key features
- Suggest improvements
- Organize your thoughts

### Step 3: Generate Project Documentation

Once you've discussed and refined your idea, generate the foundational documents:

**In your IDE, run:**
```
/Generate
```

This will create:
- **PRD.md** - Product Requirements Document
- **structure.md** - Project structure overview
- **api-spec.json** - API specifications (OpenAPI format)
- **data-model.md** - Data model documentation
- **CONTEXT.md** - Technology stack and documentation links

### Step 4: Create Your Project Plan

Generate the phase and feature structure:

**In your IDE, run:**
```
/Plan
```

This creates:
- Phase directories (`01-phase/`, `02-phase/`, etc.)
- Feature directories within each phase (`01-Feature/`, `02-Feature/`, etc.)
- Planning documents for each feature:
  - `plan.md` - Feature plan
  - `design.md` - Feature design
  - `tasks.md` - Task breakdown

### Step 5: View Your Dashboard

See your project progress:

**In your IDE:**
```
/Dashboard
```

**Or via CLI:**
```bash
doplan dashboard
```

**Or launch interactive TUI:**
```bash
doplan --tui
```

The dashboard shows:
- Overall project progress
- Phase progress bars
- Feature status and progress
- GitHub branches and commits
- Active pull requests
- Next recommended actions

### Step 6: Start Implementing a Feature

Begin working on a feature:

**In your IDE:**
```
/Implement
```

This command will:
- Create a Git branch for the feature (if GitHub automation is enabled)
- Set up the development environment
- Guide you through the implementation process

### Step 7: Workflow Automation

DoPlan automatically tracks progress through:
- Task completion in `tasks.md` files
- Dashboard auto-updates
- GitHub activity tracking
- Phase and feature status monitoring

Progress is updated automatically when you mark tasks as complete in your feature `tasks.md` files.

### Step 8: Sync GitHub Data

Keep GitHub data up to date:

**Via CLI:**
```bash
doplan github
```

This syncs:
- Branch information
- Commit history
- Pull request status
- Push events

### Step 9: Get Next Action Recommendation

When you're ready for the next step, DoPlan automatically suggests:
- Next feature to work on
- Tasks to complete
- Documentation to update
- Pull requests to review

These recommendations appear after completing actions in the TUI.

### Step 10: Configure DoPlan

Adjust settings to match your workflow:

**Via CLI:**
```bash
# Show current config
doplan config show

# Enable auto-PR creation
doplan config set github.autoPR true

# Enable auto-checkpointing
doplan config set checkpoint.autoFeature true

# Validate configuration
doplan config validate
```

### Step 11: Generate Design System (DPR)

Create a comprehensive design system for your project:

**Via TUI:**
```bash
doplan
# Select "🎨 Apply Design / DPR"
```

This will:
- Guide you through an interactive questionnaire
- Generate `doplan/design/DPR.md` - Complete design specifications
- Generate `doplan/design/design-tokens.json` - Design tokens (colors, typography, spacing)
- Generate `.doplan/ai/rules/design_rules.mdc` - Rules for AI agents to follow design system

### Step 12: Manage API Keys

Detect and manage API keys for your project services:

**Via TUI:**
```bash
doplan
# Select "🔑 Manage API Keys"
```

This will:
- Detect required services from your project files (package.json, go.mod, etc.)
- Validate API keys in `.env` files
- Generate `doplan/RAKD.md` - Required API Keys Document
- Generate `.doplan/SOPS/` - Service Operating Procedures guides for each service

### Step 13: Deploy Your Project

Deploy to your preferred platform:

**Via TUI:**
```bash
doplan
# Select "🚀 Deploy Project"
```

Supported platforms:
- Vercel
- Netlify
- Railway
- Render
- Coolify
- Docker

### Step 14: Publish Your Package

Publish to package registries:

**Via TUI:**
```bash
doplan
# Select "📦 Publish Package"
```

Supported registries:
- npm
- Homebrew
- Scoop
- Winget

### Step 15: Use AI Agents

Work with specialized AI agents in your IDE:

**Available Agents:**
- **@planner** - Project planning, PRD generation, idea refinement
- **@designer** - Design specifications following DPR
- **@coder** - Implementation based on plans and designs
- **@tester** - Test creation and execution with screenshot capture
- **@reviewer** - Code review and quality assurance
- **@devops** - Deployment and infrastructure management

**Workflow Sequence:**
```
Plan → Design → Code → Test → Review → Deploy
```

**In Cursor:**
```
@planner /Plan          # Start planning
@designer /Design       # Create design specs
@coder /Implement       # Implement feature
@tester /Test           # Run tests
@reviewer /Review       # Review code
@devops /Deploy         # Deploy feature
```

### Step 16: Get Workflow Guidance

Receive intelligent recommendations after each action:

After completing any action, DoPlan will:
- Display a "Recommended Next Step" box
- Suggest the next action based on your current progress
- Guide you through the optimal workflow sequence

**Example:**
```
✅ Plan created!

💡 Recommended Next Step: Create Design Specifications
Use @designer /Design to create design specifications for your first feature,
or use the TUI menu: [d]esign
```

## Project Structure

After installation, your project will have this structure:

```
project-root/
├── .cursor/              # Cursor IDE integration (or .gemini/, .claude/, etc.)
│   ├── agents/           # AI agent definitions (symlinked from .doplan/ai/agents/)
│   ├── rules/            # Workflow rules (symlinked from .doplan/ai/rules/)
│   ├── commands/         # DoPlan command definitions (symlinked from .doplan/ai/commands/)
│   └── config/           # Configuration and state
│       ├── doplan-config.json
│       └── doplan-state.json
├── .doplan/              # DoPlan configuration directory
│   ├── ai/               # AI integration files
│   │   ├── agents/       # AI agent definitions (6 agents)
│   │   │   ├── README.md
│   │   │   ├── planner.agent.md
│   │   │   ├── designer.agent.md
│   │   │   ├── coder.agent.md
│   │   │   ├── reviewer.agent.md
│   │   │   ├── tester.agent.md
│   │   │   └── devops.agent.md
│   │   ├── rules/        # Workflow and communication rules
│   │   │   ├── workflow.mdc
│   │   │   ├── communication.mdc
│   │   │   └── design_rules.mdc (if DPR applied)
│   │   └── commands/     # IDE command definitions
│   │       ├── run.md
│   │       ├── deploy.md
│   │       ├── create.md
│   │       └── ...
│   ├── SOPS/             # Service Operating Procedures guides
│   │   ├── stripe.md
│   │   ├── sendgrid.md
│   │   ├── aws-s3.md
│   │   └── ...
│   └── config.yaml       # DoPlan configuration
├── doplan/               # Planning directory
│   ├── dashboard.md      # Visual progress dashboard
│   ├── dashboard.html    # HTML version of dashboard
│   ├── PRD.md            # Product Requirements Document
│   ├── structure.md      # Project structure
│   ├── CONTEXT.md        # Tech stack and documentation
│   ├── RAKD.md           # Required API Keys Document (if services detected)
│   ├── design/           # Design system files (if DPR applied)
│   │   ├── DPR.md        # Design Preferences & Requirements
│   │   └── design-tokens.json
│   ├── contracts/        # API contracts
│   │   ├── api-spec.json
│   │   └── data-model.md
│   ├── templates/        # Reusable templates
│   │   ├── plan-template.md
│   │   ├── design-template.md
│   │   └── tasks-template.md
│   ├── 01-phase/         # Phase 1 (numbered format)
│   │   ├── phase-plan.md
│   │   ├── phase-progress.json
│   │   ├── 01-feature/   # Feature 1 (numbered format)
│   │   │   ├── plan.md
│   │   │   ├── design.md
│   │   │   ├── tasks.md
│   │   │   └── progress.json
│   │   └── 02-feature/   # Feature 2
│   │       └── ...
│   └── 02-phase/         # Phase 2
│       └── ...
└── README.md
```

## Development

### Building from Source

```bash
# Build
make build

# Install
make install

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean

# Production readiness check
make check-production
```

### Running in Development

```bash
# Run with development flags
make dev

# Or manually
go run ./cmd/doplan
```

## Documentation

Comprehensive documentation is available in a separate repository:
- **Documentation Repository**: [github.com/DoPlan-dev/docs](https://github.com/DoPlan-dev/docs)
- **Documentation Site**: [doplan.dev](https://doplan.dev) (if deployed)

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on:
- Development setup
- Code standards
- Testing guidelines
- Pull request process
- Release process

## Release Process

Releases are automated using GitHub Actions and GoReleaser. Before releasing:

```bash
# Run production readiness check
make check-production

# This checks 15 categories including:
# - Code formatting and static analysis
# - Tests and coverage
# - Documentation completeness
# - Version information
# - And more...
```

See [docs/guides/PRODUCTION_READINESS.md](docs/guides/PRODUCTION_READINESS.md) for the complete checklist.

For release details, see [docs/releases/RELEASE.md](docs/releases/RELEASE.md):
- Versioning strategy
- Release process
- Homebrew update process
- Docker release process (if enabled)

## License

MIT License - see LICENSE file for details
