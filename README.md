<div align="center">

<h1>DoPlan CLI</h1>

**Zero-install AI Project Director** – Bootstrap production-ready projects with a complete hierarchical AI agency system in seconds.

[![Version](https://img.shields.io/npm/v/@doplan-dev/cli?style=for-the-badge&color=blue)](https://www.npmjs.com/package/@doplan-dev/cli)
[![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)](LICENSE)
[![Node.js](https://img.shields.io/badge/node-%3E%3D14.0.0-brightgreen?style=for-the-badge&logo=node.js)](https://nodejs.org/)
[![Go](https://img.shields.io/badge/go-1.23+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![CI](https://img.shields.io/github/actions/workflow/status/DoPlan-dev/CLI/ci.yml?style=for-the-badge&label=CI)](https://github.com/DoPlan-dev/CLI/actions/workflows/ci.yml)
[![Branch Policy](https://img.shields.io/github/actions/workflow/status/DoPlan-dev/CLI/branch-protection.yml?style=for-the-badge&label=Branch%20Policy)](https://github.com/DoPlan-dev/CLI/actions/workflows/branch-protection.yml)
[![NPM Downloads](https://img.shields.io/npm/dm/@doplan-dev/cli?style=for-the-badge&color=orange)](https://www.npmjs.com/package/@doplan-dev/cli)
[![GitHub Stars](https://img.shields.io/github/stars/DoPlan-dev/CLI?style=for-the-badge&logo=github)](https://github.com/DoPlan-dev/CLI)
[![GitHub Issues](https://img.shields.io/github/issues/DoPlan-dev/CLI?style=for-the-badge&logo=github)](https://github.com/DoPlan-dev/CLI/issues)

[Installation](#-installation) • [Quick Start](#-quick-start) • [Documentation](#-documentation) • [Features](#-features) • [Contributing](#-contributing)

</div>


## 🚀 Overview

---

## ✨ What is DoPlan CLI?

**DoPlan CLI** is a revolutionary command-line tool that transforms how you start new projects. Instead of spending hours setting up project structure, configuring IDEs, writing boilerplate, and setting up CI/CD, DoPlan generates a **complete, production-ready project** with a full **hierarchical AI agency system** in under 5 seconds.

### 🎯 Perfect For

- **Solo Developers** who want to focus on building, not configuration
- **Small Teams** looking to standardize their development workflow
- **Professionals** who need production-ready project structures from day one
- **Anyone** who wants to leverage AI agents for faster development

### 🌟 Key Features

- ⚡ **Zero-Install**: Run with `npx` - no global installation needed
- 🚀 **Lightning Fast**: 80-90% faster for new projects with optimized performance
- 🤖 **18 AI Agents**: Complete hierarchical agency (Product Manager, Engineers, Designers, QA, etc.)
- 📚 **1000+ Rules Library**: Embedded best practices for all major tech stacks
- 🎨 **Interactive TUI**: Beautiful terminal interface built with Bubbletea
- 🔌 **IDE-Agnostic**: Supports 6 AI-powered IDEs (Cursor, Claude Code, Antigravity, Windsurf, Cline, OpenCode)
- 🚀 **Complete Automation**: Project structure, agents, commands, rules, CI/CD, and boilerplate
- 📦 **Offline-First**: Works completely offline after first run
- 🔓 **Transparent**: All AI logic lives in markdown files - see and modify everything
- 🧠 **Memory & Brain**: Personalizes every interaction, learns your preferences, adapts to your style
- 🏆 **Engagement System**: 200+ achievements, 30+ challenges, score tracking, and rewards
- 💾 **Backup & Restore**: Full project backup and restore functionality
- 📊 **Performance Monitoring**: Built-in performance metrics and cache statistics

## 📈 KPIs & Targets
<!-- KPIS:START -->
- **Adoption**: 10,000+ projects created in first 6 months
- **Engagement**: Average 5+ commands used per project
- **Retention**: 30%+ users create second project
- **Community**: 100+ GitHub stars, active discussions
- **Quality**: < 1% bug reports, 4.5+ star rating
- **Performance**: 95%+ of projects generated in < 5 seconds
---
<!-- KPIS:END -->
> Generated via `/github info`. The helper caches metadata in `Docs/history/github-meta.json`, so KPI data stays available even when you're offline.


---

## 📦 Installation

### Prerequisites

- **Node.js** >= 14.0.0 (for npx wrapper)
- **Go** >= 1.23.0 (only if building from source)

### Quick Install (Recommended)

The easiest way to use DoPlan CLI is via `npx` - no installation required!

```bash
npx @doplan-dev/cli
```

This will automatically download the correct binary for your platform and run it.

### Platform-Specific Installation (Optional)

<details>
<summary><b>🍎 macOS</b></summary>

#### Option 1: Using npx (No Installation – Recommended)

```bash
npx @doplan-dev/cli
```

#### Option 2: Direct Binary Download

1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-darwin-amd64` (Intel) or `doplan-darwin-arm64` (Apple Silicon)
3. Make it executable:
   ```bash
   chmod +x doplan-darwin-amd64
   mv doplan-darwin-amd64 /usr/local/bin/doplan
   ```

#### Option 3: Build from Source

```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan ./cmd/doplan
sudo mv doplan /usr/local/bin/
```

</details>

<details>
<summary><b>🪟 Windows</b></summary>

#### Option 1: Using npx (No Installation – Recommended)

```bash
npx @doplan-dev/cli
```

#### Option 2: Direct Binary Download

1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-windows-amd64.exe`
3. Rename to `doplan.exe` and add to your PATH

#### Option 3: Build from Source

```powershell
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan.exe ./cmd/doplan
# Add to PATH or use from current directory
```

</details>

<details>
<summary><b>🐧 Linux</b></summary>

#### Option 1: Using npx (No Installation – Recommended)

```bash
npx @doplan-dev/cli
```

#### Option 2: Direct Binary Download

```bash
# Download latest release
curl -L https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan-linux-amd64 -o doplan

# Make executable
chmod +x doplan

# Move to PATH
sudo mv doplan /usr/local/bin/
```

#### Option 3: Build from Source

```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan ./cmd/doplan
sudo mv doplan /usr/local/bin/
```

</details>

### Verify Installation

After installation, verify it works:

```bash
doplan --version
```

You should see the version number (e.g., `doplan version v1.3.0`).

---

## 🚀 Quick Start

### 1. Create Your First Project

```bash
npx @doplan-dev/cli
```

This launches an interactive wizard that will:
1. Ask for your project name
2. Let you choose your preferred IDE
3. Generate a complete project structure

### 2. Open Your Project

```bash
cd your-project-name
code .  # or your preferred IDE
```

### 3. Start Building

Once in your IDE, start using DoPlan commands:

```bash
/hey       # Welcome, tutorial, and command introductions
/do        # Capture project idea, conduct meeting, and refine
/plan      # Generate execution plan + task hierarchy
/dev       # Start development workflow for a feature
/done      # Mark current task complete and auto-commit/push
/sys       # System control panel (engagement, performance, backup, etc.)
```

## 📖 How to Use

### Core Commands

DoPlan uses intuitive slash commands that work directly in your AI-powered IDE:

#### Project Planning Commands

- **`/hey`** – Welcome, tutorial, and command introductions
  ```bash
  /hey
  ```
  Interactive onboarding experience for first-time users or when you need a refresher. Provides tutorial, system overview, and creates reference materials.

- **`/do`** – Capture project idea, conduct meeting, and refine
  ```bash
  /do                    # Full ideation workflow
  /do feature            # Add single feature idea
  /do now                # Fast-track with detailed prompt/PRD
  /do i'm lucky          # Get AI-suggested ideas
  ```
  Captures your project idea through iterative conversation, conducts discovery meeting, and refines suggestions. Generates IDEA.md, BRAINSTORM.md, and REFINEMENTS.md.

- **`/plan`** – Generate execution plan + tasks
  ```bash
  /plan
  ```
  Reads IDEA.md and BRAINSTORM.md, then generates TASKS.md with organized phases and feature folders.

#### Development Commands

- **`/dev`** – Start development workflow
  ```bash
  /dev                    # Start next task
  /dev --feature "auth"   # Start specific feature
  ```
  Finds next available task (or specific task), creates/checks out Git branch, syncs documentation, and starts time tracking.

#### System Commands

- **`/sys`** – System management and monitoring
  ```bash
  /sys performance    # View performance metrics and cache statistics
  /sys backup         # Create compressed project backups
  /sys restore        # Restore from backup
  /sys memory         # Export/import memory card
  /sys engagement     # View engagement dashboard
  ```

#### Team & Information Commands
  ```
  /load
  ```

#### Specialized Commands

- **`/ship`** - Release management
- **`/safe`** - Security audit
- **`/cheap`** - Cost optimization

### Complete Workflow Example

```bash
# 1. Create project
npx @doplan-dev/cli

# 2. Open in IDE
cd my-awesome-project
code .

# 3. In your IDE, start with onboarding (first time):
/hey

# 4. Capture your project idea
/do

# 5. Generate execution plan + tasks
/plan

# 6. Start development
/dev

# 7. Mark task complete when done
/done
```

---

## 🧠 Command Workflow

1. **Onboard & Capture** – Start with `/hey` for onboarding (first time), then `/do` to capture your project idea through iterative conversation, which automatically conducts a discovery meeting and refinement phase.
2. **Plan** – `/plan` reads IDEA.md and BRAINSTORM.md, then generates phased TASKS.md with organized phases and feature folders.
3. **Develop** – For each task: `/dev` (starts task, creates branch) → code → repeat.
4. **Manage & Monitor** – Use `/sys` commands to view engagement dashboard, performance metrics, create backups, manage memory card, and control system settings.

This end-to-end loop is generated with every project, so the same commands are available in Cursor, Claude Code, Windsurf, Antigravity, Cline, and OpenCode without extra setup.

## 📟 Command Catalog

| Command | Phase | What it unlocks |
| --- | --- | --- |
| `/hey` | Onboarding | Welcome, tutorial, and command introductions |
| `/do` | Strategy | Capture project idea, conduct meeting, refine suggestions → generates IDEA.md, BRAINSTORM.md, REFINEMENTS.md |
| `/plan` | Delivery | Expand planning docs into phased `TASKS.md` |
| `/dev [feature]` | Delivery | Start the next (or specific) implementation task |
| `/sys` | Operations | System control panel |
| `/sys engagement` | Context | View engagement dashboard and statistics |
| `/sys performance` | Operations | View performance metrics and cache statistics |
| `/sys backup` | Operations | Create compressed project backups |
| `/sys restore` | Operations | Restore project from backup |
| `/sys memory` | Operations | Export/import memory card |
| `/sys role` | Operations | Manage roles and permissions |
| `/sys security` | Operations | Security settings and tests |
| `/sys control` | Operations | System control panel |

👉 Looking for deeper explanations? See `docs/foundation/the-guide.md` or the [Complete Wiki](https://github.com/DoPlan-dev/CLI/tree/main/wiki) for [Commands](https://github.com/DoPlan-dev/CLI/blob/main/wiki/02-Commands/01-Command-Overview.md) and [Workflow](https://github.com/DoPlan-dev/CLI/blob/main/wiki/05-Workflow/01-Complete-Workflow.md).

### Project Structure

When you create a project, DoPlan generates:

```
my-project/
├── .cursor/
│   ├── agents/              # 18 AI agent personas
│   ├── commands/            # Command definitions
│   └── rules/               # 1000+ rules library
│       └── library/         # Tech stack rules
├── .do/
│   ├── 00_System/          # IDEA.md, PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md
│   ├── TASKS.md            # Implementation tasks
│   ├── active_state.json   # Project state
│   └── history/            # Time-stamped snapshots for rollback + reports
├── Docs/                   # Optional capitalized docs (see test fixtures)
├── .github/
│   └── workflows/          # CI/CD automation
├── src/                    # Your source code
├── STANDUP.md             # Daily standup notes
└── README.md              # Project documentation
```

---

## 📑 Docs, Changelog & Wiki

- `CHANGELOG.md` follows Keep a Changelog + SemVer. Check the **[latest entry](CHANGELOG.md)** (v1.3.0) for performance optimizations and new features.
- The `Docs/` tree mirrors what every generated project should publish (foundation, features, release, history). Use it as the canonical structure reference.
- The **[Complete Wiki](https://github.com/DoPlan-dev/CLI/tree/main/wiki)** is now organized in the repository with 9 sections and 52 files covering all features. The [GitHub Wiki](https://github.com/DoPlan-dev/CLI/wiki) serves as a gateway to the complete documentation.
- Automation helpers such as `/state`, `/sys performance`, and `/github info` keep each of those artifacts aligned (KPI block, scan diffs, state history, and performance metrics).

---

## 🎯 Features

### 🤖 Hierarchical AI Agency

DoPlan includes 18 specialized AI agents:

- **Project Orchestrator** (CEO/Engineering Manager)
- **Product Manager**
- **Engineering Lead**
- **System Architect**
- **Frontend Lead** & **Backend Lead**
- **DevOps Engineer**
- **Security Lead**
- **Design & UX Manager** & **UI/UX Designer**
- **QA & Reliability Manager** & **QA Engineer**
- **Release & Growth Manager** & **Release Captain** & **Growth Coach**
- **Documentation Lead** & **Documentation Writer**
- **Performance Engineer**

Each agent has a specific role and expertise, working together to guide your project from idea to production.

### 📚 Comprehensive Rules Library

1000+ embedded rules covering:

- Core workflows and best practices
- AI agents and orchestration
- Programming languages (Go, JavaScript, TypeScript, Python)
- Frameworks (Next.js, React, Express)
- UI libraries and design systems
- Cloud infrastructure
- Databases (PostgreSQL, MongoDB)
- Testing (Jest, Vitest, Go testing)
- DevOps and CI/CD
- Code quality and linting
- Documentation standards
- Security practices
- MCP tools integration

### 🕒 State History & Rollback

- `.do/system/history/state-*.json` stores every update to `active_state.json`, captured automatically around `/build` and `/finished`
- `/state` (backed by `go run scripts/statehistory/main.go`) lets you snapshot, list, diff, or restore with confirmation guardrails
- `/progress` surfaces the latest history diff so stakeholders always know *what* changed (phase, task, branch, completed tasks)

### ⚡ Performance Optimizations (v1.3.0)

- **80-90% faster** for new projects with fast path optimization
- **40-50% faster** for existing projects with intelligent caching
- **Lazy loading** for rules and agents - resources loaded only when needed
- **TTL-based caching** with automatic cleanup for optimal memory usage
- **Performance monitoring** via `/sys performance` command
- **60-70% reduction** in file I/O operations

### 💾 Backup & Restore

- Multiple backup types: project, plan, project-plan, or full backup
- Compressed backups with automatic naming
- Safe restore with dry-run mode and version compatibility checks
- Memory card export/import for easy migration
- Migration assistant for project upgrades

### 🧠 Memory & Brain System

DoPlan learns about you and personalizes every interaction:

- **Memory Card** - Your personal relationship file stored at `~/.doplan/memory_card.json`
  - Remembers your identity, preferences, and communication style
  - Tracks your learning goals, tech stack, and pain points
  - Builds relationship history across all projects
  - Works globally across all your DoPlan projects

- **Brain System** - Intelligent personalization layer
  - Personalizes agent responses to match your style
  - Adjusts tone based on relationship level
  - Provides context-aware guidance
  - Matches responses to your experience level
  - Offers personalized encouragement and motivation

- **View & Manage**: Use `/sys memory` to export/import your memory card or `/sys engagement` to see your relationship stats

### 🏆 Engagement System

Make development fun and motivating with achievements, challenges, and rewards:

- **Achievement System** - 200+ achievements for milestones
  - Score milestones (100, 250, 1,000, 10,000+ points)
  - Project achievements (First Steps, Serial Builder, Project Master)
  - Command usage achievements (Hello There!, Code Machine)
  - Learning achievements (Student, Tech Master)
  - Relationship achievements (Trust Builder, Best Friend)

- **Challenge System** - 30+ high-scoring challenges
  - First-time task challenges (300-2000 points each)
  - Milestone challenges for major accomplishments
  - Special event challenges

- **Score System** - Points-based progression (0-100,000+)
  - Earn points from achievements and challenges
  - Track progress through score ranges
  - Unlock new achievements as you progress
  - View your score via `/sys engagement`

- **Reward System** - Strategic reward scheduling
  - Celebrations at key milestones
  - Personalized encouragement
  - Dopamine timing for optimal motivation

### 🎨 Beautiful Interactive TUI

Built with [Bubbletea](https://github.com/charmbracelet/bubbletea), DoPlan's terminal interface is:

- Fast and responsive
- Visually appealing
- Keyboard-friendly
- Accessible

### 🔌 Multi-IDE Support

Works seamlessly with:

- **Cursor** (Recommended)
- **Claude Code**
- **Antigravity**
- **Windsurf**
- **Cline**
- **OpenCode**

### 🚀 Complete Automation

DoPlan generates:

- ✅ Project structure
- ✅ AI agent system
- ✅ Command definitions
- ✅ Rules library
- ✅ GitHub Actions workflows (CI/CD, releases, changelog)
- ✅ IDE configuration files
- ✅ Boilerplate code
- ✅ Documentation templates

---

## 📚 Documentation

### Complete Wiki

The comprehensive DoPlan CLI documentation is organized in the repository's **[wiki directory](https://github.com/DoPlan-dev/CLI/tree/main/wiki)** with 9 main sections:

- **[01. Getting Started](https://github.com/DoPlan-dev/CLI/tree/main/wiki/01-Getting-Started)** - Installation, first project, quick tour
- **[02. Commands](https://github.com/DoPlan-dev/CLI/tree/main/wiki/02-Commands)** - Complete command documentation
- **[03. Engagement System](https://github.com/DoPlan-dev/CLI/tree/main/wiki/03-Engagement-System)** - Achievements, challenges, rewards
- **[04. Memory and Brain](https://github.com/DoPlan-dev/CLI/tree/main/wiki/04-Memory-and-Brain)** - Personalization features
- **[05. Workflow](https://github.com/DoPlan-dev/CLI/tree/main/wiki/05-Workflow)** - Complete workflow guides
- **[06. Features](https://github.com/DoPlan-dev/CLI/tree/main/wiki/06-Features)** - Time tracking, state management, etc.
- **[07. Learning & Education](https://github.com/DoPlan-dev/CLI/tree/main/wiki/07-Learning-Education)** - Educational content
- **[08. Advanced Topics](https://github.com/DoPlan-dev/CLI/tree/main/wiki/08-Advanced)** - Power user features
- **[09. Reference](https://github.com/DoPlan-dev/CLI/tree/main/wiki/09-Reference)** - Quick reference guides

**[View Complete Wiki Index](https://github.com/DoPlan-dev/CLI/blob/main/wiki/INDEX.md)** | **[Wiki README](https://github.com/DoPlan-dev/CLI/blob/main/wiki/README.md)** | **[GitHub Wiki Gateway](https://github.com/DoPlan-dev/CLI/wiki)**

### Quick Links

**Getting Started:**
- [Installation Guide](https://github.com/DoPlan-dev/CLI/blob/main/wiki/01-Getting-Started/02-Installation.md)
- [First Project Tutorial](https://github.com/DoPlan-dev/CLI/blob/main/wiki/01-Getting-Started/03-First-Project.md)
- [Quick Tour](https://github.com/DoPlan-dev/CLI/blob/main/wiki/01-Getting-Started/04-Quick-Tour.md)

**Essential Guides:**
- [Command Overview](https://github.com/DoPlan-dev/CLI/blob/main/wiki/02-Commands/01-Command-Overview.md)
- [Complete Workflow](https://github.com/DoPlan-dev/CLI/blob/main/wiki/05-Workflow/01-Complete-Workflow.md)
- [Engagement System](https://github.com/DoPlan-dev/CLI/blob/main/wiki/03-Engagement-System/01-Overview.md)

**Reference:**
- [Troubleshooting](https://github.com/DoPlan-dev/CLI/blob/main/wiki/08-Advanced/04-Troubleshooting.md)
- [Contributing Guide](https://github.com/DoPlan-dev/CLI/blob/main/wiki/08-Advanced/05-Contributing.md)

---

## 🤝 Contributing

We welcome contributions! Whether it's:

- 🐛 Reporting bugs
- 💡 Suggesting features
- 📝 Improving documentation
- 🔧 Submitting pull requests
- ⭐ Giving us a star

Every contribution helps make DoPlan better for everyone.

See our [Contributing Guide](https://github.com/DoPlan-dev/CLI/blob/main/wiki/08-Advanced/05-Contributing.md) for details.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - Beautiful TUI framework
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling library
- All our amazing contributors and users

---

## 🔗 Links

- **GitHub**: [https://github.com/DoPlan-dev/CLI](https://github.com/DoPlan-dev/CLI)
- **NPM Package**: [https://www.npmjs.com/package/@doplan-dev/cli](https://www.npmjs.com/package/@doplan-dev/cli)
- **Issues**: [https://github.com/DoPlan-dev/CLI/issues](https://github.com/DoPlan-dev/CLI/issues)
- **Discussions**: [https://github.com/DoPlan-dev/CLI/discussions](https://github.com/DoPlan-dev/CLI/discussions)
- **Complete Wiki**: [https://github.com/DoPlan-dev/CLI/tree/main/wiki](https://github.com/DoPlan-dev/CLI/tree/main/wiki)
- **GitHub Wiki**: [https://github.com/DoPlan-dev/CLI/wiki](https://github.com/DoPlan-dev/CLI/wiki)

---

<div align="center">

**Made with ❤️ by the DoPlan Team**

[⭐ Star us on GitHub](https://github.com/DoPlan-dev/CLI) • [🐛 Report Bug](https://github.com/DoPlan-dev/CLI/issues) • [💡 Request Feature](https://github.com/DoPlan-dev/CLI/issues) • [📖 Complete Wiki](https://github.com/DoPlan-dev/CLI/tree/main/wiki)

</div>
