# DoPlan CLI

[![Test](https://github.com/DoPlan-dev/CLI/actions/workflows/test.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/test.yml)
[![Lint](https://github.com/DoPlan-dev/CLI/actions/workflows/lint.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/lint.yml)
[![Build](https://github.com/DoPlan-dev/CLI/actions/workflows/build.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/build.yml)
[![Release](https://github.com/DoPlan-dev/CLI/actions/workflows/release.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/release.yml)
[![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

**DoPlan** is a comprehensive project workflow automation tool that transforms your app ideas into well-structured, documented, and trackable development projects. Built with Go, DoPlan provides a complete command-line interface and integrates seamlessly with popular AI-powered IDEs to guide you through the entire development lifecycle—from initial concept to deployment.

DoPlan automates the tedious aspects of project management by generating documentation, creating project structures, managing Git workflows, tracking progress, and providing visual dashboards. It enforces best practices through automated rules, templates, and checkpoints, ensuring your project stays organized and maintainable throughout its development.

Whether you're building a web application, API service, mobile app, or any software project, DoPlan helps you break down your idea into manageable phases and features, track progress visually, automate GitHub operations, and maintain comprehensive documentation—all while working naturally within your preferred development environment.

## Features

### 🚀 **Workflow Automation**
- **Automated Project Setup**: Initialize your project with a structured workflow in seconds
- **Phase & Feature Management**: Organize your project into phases and features with automatic directory structure
- **Progress Tracking**: Real-time progress monitoring with visual dashboards (markdown and HTML)
- **Document Generation**: Auto-generate PRD, project structure, API contracts, and planning documents

### 🎨 **Interactive Interface**
- **Fullscreen TUI**: Beautiful terminal user interface with interactive dashboard
- **Visual Progress Bars**: See project, phase, and feature progress at a glance
- **Multi-view Dashboard**: Switch between project overview, phases, features, GitHub activity, and configuration

### 🔧 **GitHub Integration**
- **Automatic Branching**: Create feature branches automatically
- **Auto-PR Creation**: Automatically create pull requests when features are complete
- **Commit Management**: Track commits, pushes, and branch status
- **PR Tracking**: Monitor pull request status and URLs

### 📝 **Documentation & Templates**
- **Template System**: Customizable templates for plans, designs, and tasks
- **Template Management**: Add, edit, remove, and set default templates
- **Context Generation**: Auto-generate `CONTEXT.md` with tech stack and documentation links
- **Documentation Rules**: Automated rules for documentation organization and naming

### ⚙️ **Configuration & Validation**
- **Flexible Configuration**: Manage settings for GitHub, checkpoints, and workflow
- **Project Validation**: Validate project structure and configuration with auto-fix
- **Checkpoint System**: Create, list, and restore project checkpoints (Time Machine)
- **Auto-Checkpointing**: Automatic checkpoints for features and phases

### 🔌 **IDE Integration**
- **Multi-IDE Support**: Works with Cursor, Gemini CLI, Claude CLI, Codex CLI, OpenCode, and Qwen Code
- **Custom Commands**: IDE-specific commands for seamless workflow integration
- **Workflow Rules**: Automated rules and conventions for your development process

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

### Homebrew (Recommended for macOS/Linux)

```bash
# Install from Homebrew tap (coming soon)
brew install DoPlan-dev/doplan/doplan

# Or if tap is added
brew tap DoPlan-dev/doplan
brew install doplan
```

**Note:** Homebrew tap repository setup is in progress. Binary releases are available below.

### Binary Release

Download the latest release for your platform from [releases](https://github.com/DoPlan-dev/CLI/releases):

```bash
# Linux/macOS
curl -L https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan_<version>_<platform>.tar.gz | tar -xz
sudo mv doplan /usr/local/bin/

# Windows
# Download doplan_<version>_windows_amd64.tar.gz
# Extract and add to PATH
```

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
| `doplan install` | Install DoPlan in your project (interactive IDE selection) |
| `doplan dashboard` | View project dashboard with progress and GitHub activity |
| `doplan --tui` | Launch fullscreen interactive TUI dashboard |
| `doplan github` | Sync GitHub data (branches, commits, PRs) and update dashboard |
| `doplan progress` | Update all progress tracking files and regenerate dashboard |
| `doplan validate` | Validate project structure, configuration, and state consistency |

### Configuration Commands

| Command | Description |
|---------|-------------|
| `doplan config show` | Show current configuration (table or JSON format) |
| `doplan config set <key> <value>` | Set a configuration value |
| `doplan config reset` | Reset configuration to defaults |
| `doplan config validate` | Validate configuration settings |

**Configuration Keys:**
- `github.enabled` - Enable/disable GitHub integration
- `github.autoBranch` - Auto-create branches for features
- `github.autoPR` - Auto-create PRs when features complete
- `checkpoint.autoFeature` - Auto-checkpoint when feature starts
- `checkpoint.autoPhase` - Auto-checkpoint when phase starts
- `checkpoint.autoComplete` - Auto-checkpoint when feature/phase completes

### Checkpoint Commands

| Command | Description |
|---------|-------------|
| `doplan checkpoint create [name]` | Create a manual checkpoint |
| `doplan checkpoint list` | List all checkpoints |
| `doplan checkpoint restore <id>` | Restore a checkpoint |

**Checkpoint Options:**
- `--type <type>` - Checkpoint type: `manual`, `feature`, `phase`
- `--description <text>` - Add description to checkpoint

### Template Commands

| Command | Description |
|---------|-------------|
| `doplan templates list` | List all available templates |
| `doplan templates show <name>` | Show template content |
| `doplan templates add <name> <file>` | Add a template from file |
| `doplan templates edit <name>` | Edit template (opens in default editor) |
| `doplan templates use <name> [--for type]` | Set default template (plan/design/tasks) |
| `doplan templates remove <name>` | Remove a template |

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

### Step 7: Update Progress

As you complete tasks, update progress:

**In your IDE:**
```
/Progress
```

**Or via CLI:**
```bash
doplan progress
```

This updates:
- Feature progress percentages
- Task completion status
- Phase progress
- Overall project progress
- Regenerates the dashboard

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

When you're ready for the next step:

**In your IDE:**
```
/Next
```

This suggests:
- Next feature to work on
- Tasks to complete
- Documentation to update
- Pull requests to review

### Step 10: Create Checkpoints (Optional)

Save your project state at important milestones:

**Via CLI:**
```bash
# Create manual checkpoint
doplan checkpoint create "Before major refactor"

# List checkpoints
doplan checkpoint list

# Restore a checkpoint
doplan checkpoint restore <checkpoint-id>
```

### Step 11: Customize Templates

Customize your document templates:

**Via CLI:**
```bash
# List templates
doplan templates list

# Show a template
doplan templates show plan-template.md

# Add custom template
doplan templates add my-plan.md /path/to/template.md

# Set as default
doplan templates use my-plan.md --for plan

# Edit template
doplan templates edit plan-template.md
```

### Step 12: Configure DoPlan

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

## Project Structure

After installation, your project will have this structure:

```
project-root/
├── .cursor/              # Cursor IDE integration (or .gemini/, .claude/, etc.)
│   ├── commands/         # DoPlan command definitions
│   ├── rules/            # Workflow rules and policies
│   └── config/           # Configuration and state
│       ├── doplan-config.json
│       └── doplan-state.json
├── doplan/               # Planning directory
│   ├── dashboard.md      # Visual progress dashboard
│   ├── dashboard.html    # HTML version of dashboard
│   ├── PRD.md            # Product Requirements Document
│   ├── structure.md      # Project structure
│   ├── CONTEXT.md        # Tech stack and documentation
│   ├── contracts/        # API contracts
│   │   ├── api-spec.json
│   │   └── data-model.md
│   ├── templates/        # Reusable templates
│   │   ├── plan-template.md
│   │   ├── design-template.md
│   │   └── tasks-template.md
│   ├── 01-phase/         # Phase 1
│   │   ├── phase-plan.md
│   │   ├── phase-progress.json
│   │   ├── 01-Feature/
│   │   │   ├── plan.md
│   │   │   ├── design.md
│   │   │   ├── tasks.md
│   │   │   └── progress.json
│   │   └── 02-Feature/
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

Releases are automated using GitHub Actions and GoReleaser. See [docs/development/RELEASE.md](docs/development/RELEASE.md) for details on:
- Versioning strategy
- Release process
- Homebrew update process
- Docker release process (if enabled)

## License

MIT License - see LICENSE file for details
