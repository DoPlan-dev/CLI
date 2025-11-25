<div align="center">


 ██████╗░░█████╗░██████╗░██╗░░░░░░█████╗░███╗░░██╗
██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔══██╗████╗░██║
 ██║░░██║██║░░██║██████╔╝██║░░░░░███████║██╔██╗██║
██║░░██║██║░░██║██╔═══╝░██║░░░░░██╔══██║██║╚████║
 ██████╔╝╚█████╔╝██║░░░░░███████╗██║░░██║██║░╚███║
╚═════╝░░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚══╝


**Zero-install AI Project Director** - Bootstrap production-ready projects with a complete hierarchical AI agency system in seconds.

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


# 🚀 DoPlan CLI

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
- 🤖 **18 AI Agents**: Complete hierarchical agency (Product Manager, Engineers, Designers, QA, etc.)
- 📚 **1000+ Rules Library**: Embedded best practices for all major tech stacks
- 🎨 **Interactive TUI**: Beautiful terminal interface built with Bubbletea
- 🔌 **IDE-Agnostic**: Supports 6 AI-powered IDEs (Cursor, Claude Code, Antigravity, Windsurf, Cline, OpenCode)
- 🚀 **Complete Automation**: Project structure, agents, commands, rules, CI/CD, and boilerplate
- 📦 **Offline-First**: Works completely offline after first run
- 🔓 **Transparent**: All AI logic lives in markdown files - see and modify everything

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

### Platform-Specific Installation

<details>
<summary><b>🍎 macOS</b></summary>

#### Option 1: Using Homebrew (Recommended)

```bash
# Add tap (if needed)
brew tap doplan-dev/cli

# Install
brew install doplan
```

#### Option 2: Using npx (No Installation)

```bash
npx @doplan-dev/cli
```

#### Option 3: Direct Binary Download

1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-darwin-amd64` (Intel) or `doplan-darwin-arm64` (Apple Silicon)
3. Make it executable:
   ```bash
   chmod +x doplan-darwin-amd64
   mv doplan-darwin-amd64 /usr/local/bin/doplan
   ```

#### Option 4: Build from Source

```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan ./cmd/doplan
sudo mv doplan /usr/local/bin/
```

</details>

<details>
<summary><b>🪟 Windows</b></summary>

#### Option 1: Using Scoop (Recommended)

```bash
scoop bucket add doplan https://github.com/DoPlan-dev/scoop-bucket.git
scoop install doplan
```

#### Option 2: Using npx (No Installation)

```bash
npx @doplan-dev/cli
```

#### Option 3: Direct Binary Download

1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-windows-amd64.exe`
3. Rename to `doplan.exe` and add to your PATH

#### Option 4: Build from Source

```powershell
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan.exe ./cmd/doplan
# Add to PATH or use from current directory
```

</details>

<details>
<summary><b>🐧 Linux</b></summary>

#### Option 1: Using npx (No Installation - Recommended)

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

#### Option 4: Using Package Managers

**Debian/Ubuntu:**
```bash
# Download .deb package from releases (when available)
wget https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan_amd64.deb
sudo dpkg -i doplan_amd64.deb
```

**Arch Linux:**
```bash
# Using AUR (when available)
yay -S doplan-cli
```

</details>

<details>
<summary><b>🐳 Docker</b></summary>

```bash
docker run --rm -it -v $(pwd):/workspace doplan/cli
```

</details>

### Verify Installation

After installation, verify it works:

```bash
doplan --version
```

You should see the version number (e.g., `doplan version 1.0.4`).

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

```
/tell    → Capture your project idea
/improve → Team brainstorm session
/write   → Generate PRD + Architecture + Design System
/good    → Approve & lock the plan
/plan    → Generate execution plan + task hierarchy
/build   → Start coding
/report  → Generate scan metadata + diffs
/feedback → Log structured feedback entries
/state   → Snapshot or restore .plan/active_state
```

---

## 📖 How to Use

### Core Commands

DoPlan uses intuitive slash commands that work directly in your AI-powered IDE:

#### Project Planning Commands

- **`/tell`** - Capture your project idea
  ```
  /tell I want to build a task management app with React and Node.js
  ```

- **`/improve`** - Team brainstorm session
  ```
  /improve
  ```
  Activates Product Manager, Engineering Lead, and Design Manager for interactive brainstorming.

- **`/write`** - Generate planning documents
  ```
  /write
  ```
  Generates PRD, Architecture, and Design System documents.

- **`/change`** - Edit any document
  ```
  /change prd Add dark mode support
  /change architecture Use PostgreSQL instead of MongoDB
  ```

- **`/good`** - Approve & lock the plan
  ```
  /good
  ```
  Locks the current plan and enables task generation.

#### Development Commands

- **`/plan`** - Generate execution plan + tasks
  ```
  /plan
  ```
  Synthesizes TASKS.md from the approved plan and scaffolds phase folders.

- **`/build`** - Start coding
  ```
  /build        # Start next task
  /build 3      # Start specific task
  ```

- **`/progress`** - Show current progress
  ```
  /progress
  go run scripts/progress/main.go --root .
  ```
  Runs the Go helper at `scripts/progress/` to display totals plus the most recent `.plan/history` delta (phase/task/branch/completed changes). Pass `--json` for machine-readable output.

- **`/state`** - Manage state snapshots & rollbacks
  ```
  /state list
  /state diff --json
  /state restore --file state-20251124T120000Z.json --yes
  ```
  Wraps `go run scripts/statehistory/main.go` so you can inspect history, capture snapshots (before/after `/build` and `/finished`), and safely restore `.plan/active_state.json`.

- **`/finished`** - Mark task complete
  ```
  /finished
  ```

#### Feedback & Reporting Commands

- **`/feedback`** - Log structured product/bug feedback
  ```
  /feedback bug "QR download fails" "API returns 500 when Accept header missing" --author QA
  ```
  Saves to `Docs/history/feedback.md` (human readable) and `Docs/history/feedback.json` (consumed by automation).

- **`/report`** - Generate scan metadata + diffs
  ```
  /report                       # current project
  /report ./test/qr-generator/test-no01
  ```
  Runs `go run scripts/scanreport/main.go` to update `SCAN_REPORT_*.json`, create `SCAN_DIFF_<date>.md`, and append both the latest state-history summary and an embedded `/progress` snapshot (phase, completion %, upcoming tasks). Use `--preset exec` or `--preset detailed` for alternate templates (exec view, detailed visuals + dependency audit).
  Customization:
  - Create `.plan/reports/config.json` to set defaults:
    ```json
    {
      "preset": "exec",
      "sections": ["executive", "progress", "visuals", "state", "feedback"]
    }
    ```
  - CLI flags override config; custom `sections` let you reorder or omit report blocks.

#### Team & Information Commands

- **`/team`** - Show active agents and hierarchy
  ```
  /team
  ```

- **`/load`** - Inject context into AI agents
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

# 3. In your IDE, start planning:
/tell I want to build a social media dashboard with real-time analytics

# 4. Brainstorm with the team
/improve

# 5. Generate planning documents
/write

# 6. Review and approve
/good

# 7. Generate execution plan + tasks
/plan

# 8. Start building
/build

# 9. Track progress
/progress
```

---

## 🧠 Command Workflow

```
/tell → /improve → /write → /change → /good
                    │
                    ▼
                 /plan → /build ⇆ /progress ⇆ /state
                               │
                               ▼
                        /finished → /report → /ship
                                      │
                                      └── /feedback, /safe, /cheap, /branchci
```

- **Plan**: `/tell` through `/good` capture and approve strategy.
- **Execute**: `/plan`, `/build`, `/progress`, `/state`, and `/finished` keep delivery disciplined and auditable.
- **Operate**: `/report`, `/feedback`, `/ship`, `/safe`, `/cheap`, `/branchci`, `/github`, `/team`, and `/load` keep stakeholders aligned, secure, and informed.

Every generated project ships with this workflow baked into `.plan`, `Docs/`, `.github/`, and the wiki so cursor-based IDEs can enforce it automatically.

## 📟 Command Catalog

| Command | Phase | What it unlocks |
| --- | --- | --- |
| `/tell` | Strategy | Capture project intent into `.plan/00_System/IDEA.md` |
| `/improve` | Strategy | Brainstorm with all Level 1 managers |
| `/write` | Strategy | Generate PRD, Architecture, Design System |
| `/change` | Strategy | Patch any planning doc with natural language |
| `/good` | Strategy | Lock planning set and advance to tasks |
| `/plan` | Delivery | Expand planning docs into phased TASKS.md |
| `/build [id]` | Delivery | Start next (or specific) implementation task |
| `/progress` | Delivery | Summaries for total/completed tasks + upcoming work |
| `/state <subcommand>` | Delivery | Snapshot, diff, and restore `.plan/active_state.json` |
| `/finished` | Delivery | Mark tasks complete, auto-commit, and push |
| `/feedback <type>` | Operations | Log bugs/features/questions into `Docs/history/feedback.*` |
| `/report [path]` | Operations | Generate SCAN_REPORT metadata + diffs |
| `/ship` | Operations | Release orchestration + versioning checklist |
| `/safe` | Operations | Security review + dependency risk scan |
| `/cheap` | Operations | Cost optimization playbook |
| `/team` | Context | Display the 18-agent hierarchy |
| `/load <context>` | Context | Inject extra domain knowledge for agents |
| `/github <subcommand>` | Integrations | Sync KPIs, prep issues/milestones, update cache |
| `/branchci` | Integrations | Regenerate per-branch workflow guardrails |

👉 Looking for deeper explanations? See `Docs/foundation/the-guide.md` or the wiki pages for [Commands](https://github.com/DoPlan-dev/CLI/wiki/Commands) and [Workflow](https://github.com/DoPlan-dev/CLI/wiki/Workflow).

### Project Structure

When you create a project, DoPlan generates:

```
my-project/
├── .cursor/
│   ├── agents/              # 18 AI agent personas
│   ├── commands/            # Command definitions
│   └── rules/               # 1000+ rules library
│       └── library/         # Tech stack rules
├── .plan/
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

- `CHANGELOG.md` follows Keep a Changelog + SemVer. Check the **[latest entry](CHANGELOG.md)** before cutting a release or running `/ship`.
- The `Docs/` tree mirrors what every generated project should publish (foundation, features, release, history). Use it as the canonical structure reference.
- The **[GitHub wiki](https://github.com/DoPlan-dev/CLI/wiki)** stays in sync with this README—Commands, Workflow, Quick Start, and Troubleshooting are updated whenever the CLI changes.
- Automation helpers such as `/report`, `/feedback`, `/state`, and `/github info` keep each of those artifacts aligned (KPI block, scan diffs, feedback logs, and state history).

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

- `.plan/history/state-*.json` stores every update to `active_state.json`, captured automatically around `/build` and `/finished`
- `/state` (backed by `go run scripts/statehistory/main.go`) lets you snapshot, list, diff, or restore with confirmation guardrails
- `/progress` and `/report` surface the latest history diff so stakeholders always know *what* changed (phase, task, branch, completed tasks)

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

### Getting Started

- [Installation Guide](https://github.com/DoPlan-dev/CLI/wiki/Installation) - Detailed installation for all platforms
- [Quick Start Tutorial](https://github.com/DoPlan-dev/CLI/wiki/Quick-Start) - 5-minute tutorial
- [Command Reference](https://github.com/DoPlan-dev/CLI/wiki/Commands) - Complete command documentation

### Guides

- [Workflow Guide](https://github.com/DoPlan-dev/CLI/wiki/Workflow) - End-to-end project creation
- [Agent System](https://github.com/DoPlan-dev/CLI/wiki/Agents) - Understanding the AI agency
- [Rules Library](https://github.com/DoPlan-dev/CLI/wiki/Rules) - Using and customizing rules
- [Advanced Usage](https://github.com/DoPlan-dev/CLI/wiki/Advanced) - Customization and extensibility

### Reference

- [Architecture](https://github.com/DoPlan-dev/CLI/wiki/Architecture) - Technical deep dive
- [Troubleshooting](https://github.com/DoPlan-dev/CLI/wiki/Troubleshooting) - Common issues and solutions
- [FAQ](https://github.com/DoPlan-dev/CLI/wiki/FAQ) - Frequently asked questions

### Contributing

- [Contributing Guide](https://github.com/DoPlan-dev/CLI/wiki/Contributing) - How to contribute
- [Development Setup](https://github.com/DoPlan-dev/CLI/wiki/Development) - Building from source
- [Code of Conduct](https://github.com/DoPlan-dev/CLI/wiki/Code-of-Conduct) - Community guidelines

---

## 🤝 Contributing

We welcome contributions! Whether it's:

- 🐛 Reporting bugs
- 💡 Suggesting features
- 📝 Improving documentation
- 🔧 Submitting pull requests
- ⭐ Giving us a star

Every contribution helps make DoPlan better for everyone.

See our [Contributing Guide](https://github.com/DoPlan-dev/CLI/wiki/Contributing) for details.

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
- **Wiki**: [https://github.com/DoPlan-dev/CLI/wiki](https://github.com/DoPlan-dev/CLI/wiki)

---

<div align="center">

**Made with ❤️ by the DoPlan Team**

[⭐ Star us on GitHub](https://github.com/DoPlan-dev/CLI) • [🐛 Report Bug](https://github.com/DoPlan-dev/CLI/issues) • [💡 Request Feature](https://github.com/DoPlan-dev/CLI/issues) • [📖 Documentation](https://github.com/DoPlan-dev/CLI/wiki)

</div>
