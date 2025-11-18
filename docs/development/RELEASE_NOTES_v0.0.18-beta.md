# Release Notes: DoPlan CLI v0.0.18-beta

**Release Date:** January 2025  
**Status:** ✅ Complete - Ready for Testing

---

## 🎉 Overview

v0.0.18-beta completes the foundational features of DoPlan CLI, establishing a solid base for workflow automation and project management. This release focuses on core functionality, smart context detection, dashboard visualization, and project-first documentation.

---

## ✨ What's New

### 🏗️ Phase 1: Architecture Setup

**Complete project structure and foundation**
- ✅ Centralized configuration system (`.doplan/config.yaml` and `.cursor/config/doplan-config.json`)
- ✅ State management (`.doplan/state.json`)
- ✅ Dashboard JSON format (`.doplan/dashboard.json`)
- ✅ Comprehensive error handling framework with structured error types
- ✅ Beautiful logging system with rotation
- ✅ Smooth animations and loading indicators
- ✅ Consistent theme system using Lipgloss

**Key Files:**
- `internal/config/` - Configuration management
- `internal/error/` - Error handling framework
- `pkg/theme/` - Theme system
- `pkg/animations/` - Animation utilities

---

### 🧠 Phase 2: Smart Root Command & Context Detection

**Intelligent project state detection**
- ✅ Detects 5 project states:
  - Empty directory → New Project Wizard
  - Existing code (no DoPlan) → Adopt Project Wizard
  - Old DoPlan structure → Migration Wizard
  - New DoPlan structure → Dashboard TUI
  - Inside feature/phase → Feature/Phase View
- ✅ Smart root command routing based on context
- ✅ Dashboard aliases: `.`, `dash`, `d`
- ✅ Project analyzer detects tech stack automatically

**Key Features:**
- Context-aware command behavior
- Automatic wizard selection
- Tech stack detection from project files

---

### 📊 Phase 3: Dashboard Supercharge

**Enhanced dashboard with real-time data**
- ✅ Dashboard JSON format with comprehensive project data
- ✅ Auto-updating dashboard generation
- ✅ TUI dashboard with multiple views:
  - Dashboard overview
  - Phases view
  - Features view
  - GitHub activity
  - Configuration
  - Statistics
- ✅ Progress bars with color-coded status
- ✅ Activity feed with recent changes
- ✅ Velocity metrics and trends
- ✅ Sparkline visualizations

**Performance Optimizations:**
- Deferred statistics loading
- Prioritized `dashboard.json` loading
- Caching for faster updates

---

### 📝 Phase 4: Project-First Documentation

**AI-ready, project-focused documentation**

#### CONTEXT.md Improvements
- ✅ Project-specific header: "Project Context: [Project Name]"
- ✅ Project Overview section (auto-populated from state/idea)
- ✅ Technology Stack categorized by:
  - Frontend technologies
  - Backend technologies
  - Services & APIs (with SOPS links)
- ✅ Project-Specific Documentation links
- ✅ Development Guidelines section
- ✅ DoPlan Resources in collapsible `<details>` section

#### README.md Improvements
- ✅ Project-first structure:
  - Project name and description
  - Quick Start (project-specific)
  - Features (project features, not DoPlan features)
  - Tech Stack
  - Project Structure (reflecting `##-phase-name/##-feature-name`)
  - Environment Variables (links to RAKD.md)
  - Documentation links
- ✅ DoPlan information moved to collapsible section at bottom
- ✅ Auto-populated from project state and config

**Key Files:**
- `internal/generators/context.go` - CONTEXT.md generator
- `internal/generators/readme.go` - README.md generator

---

### 🔗 Phase 5: GitHub & IDE Integration

**Mandatory GitHub integration and IDE support**

#### GitHub Repository Requirement
- ✅ `RequireGitHubRepo()` validator function
- ✅ Validates repository format (user/repo, URL, SSH)
- ✅ Checks GitHub CLI access if available
- ✅ Structured error messages with suggestions
- ✅ Actions requiring GitHub: discuss, plan, implement, feature, progress, deploy

#### GitHub Badge on Dashboard
- ✅ Permanent repository badge at top of dashboard
- ✅ Shows commit count and last commit time
- ✅ Styled with rounded border and primary color
- ✅ Warning badge if GitHub enabled but repo not configured

#### IDE Integration
- ✅ Support for multiple IDEs:
  - Cursor IDE (symlinks from `.doplan/ai/`)
  - VS Code + Copilot
  - Gemini CLI
  - Claude Code
  - Codex CLI
  - OpenCode
  - Qwen Code
  - Generic IDE (setup guides)
- ✅ Automatic symlink/copy creation
- ✅ IDE-specific command formats

**Key Files:**
- `internal/github/validator.go` - GitHub requirement validation
- `internal/integration/` - IDE integration logic
- `internal/tui/screens/dashboard.go` - GitHub badge display

---

### 🎨 Phase 6: Foundational Polish

**Consistent styling and user experience**
- ✅ Lipgloss theme applied throughout TUI
- ✅ Consistent color palette and styling
- ✅ GitHub badge with rounded borders
- ✅ Progress bars with color-coded states
- ✅ Smooth animations and transitions
- ✅ Beautiful error messages with fix suggestions

---

## 🔧 Technical Improvements

### Bug Fixes
- ✅ Fixed import cycles between `dashboard` and `generators` packages
- ✅ Resolved template compilation issues with build tags
- ✅ Fixed test assertions for new documentation structure
- ✅ Fixed symlink creation to preserve existing files
- ✅ Enhanced directory creation with proper error handling

### Code Quality
- ✅ All tests passing
- ✅ Consistent error handling across all packages
- ✅ Proper package organization (moved shared types to `pkg/models`)
- ✅ Test packages separated to avoid import cycles

---

## 📦 Installation

```bash
# Install DoPlan CLI
go install github.com/DoPlan-dev/CLI/cmd/doplan@latest

# Or build from source
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI/cli
go build -o doplan ./cmd/doplan
```

---

## 🚀 Quick Start

### New Project
```bash
# In an empty directory
doplan
# → Launches New Project Wizard
```

### Existing Project
```bash
# In a project with code but no DoPlan
doplan
# → Launches Adopt Project Wizard
```

### View Dashboard
```bash
# In a DoPlan-managed project
doplan
# or
doplan dashboard
# or
doplan .
```

---

## 📚 Documentation

- **Implementation Guide:** `docs/development/V0.0.18_IMPLEMENTATION_GUIDE.md`
- **Testing Scenarios:** `docs/development/TESTING_SCENARIOS.md`
- **Manual Testing Guide:** `docs/development/MANUAL_TESTING_GUIDE.md`
- **TUI Troubleshooting:** `docs/development/TUI_TROUBLESHOOTING.md`
- **Development Status:** `docs/development/DEVELOPMENT_STATUS.md`

---

## 🧪 Testing

All automated tests are passing:
- ✅ Unit tests
- ✅ Integration tests
- ✅ CLI command tests
- ✅ Generator tests

**Manual Testing:**
See `MANUAL_TEST_CHECKLIST.md` for comprehensive manual testing scenarios.

---

## 🔄 Migration Notes

### From Previous Versions

If you have an existing DoPlan installation:
- Run `doplan` in your project directory
- The Migration Wizard will automatically detect old structure
- Follow the prompts to migrate to new structure

### Configuration Changes

- New YAML config format: `.doplan/config.yaml`
- Old JSON format still supported: `.cursor/config/doplan-config.json`
- State moved to: `.doplan/state.json`
- Dashboard JSON: `.doplan/dashboard.json`

---

## 🐛 Known Issues

None at this time. All planned features for v0.0.18-beta are complete.

---

## 🙏 Acknowledgments

This release represents a complete rewrite and restructuring of DoPlan CLI, establishing a solid foundation for future enhancements.

---

## 📋 What's Next: v0.0.19-beta

The next release will focus on advanced features:
- Advanced TUI actions (Run Dev Server, Undo, Deploy, Publish, Security, Fix)
- Design System (DPR) generation
- Secrets Management (RAKD/SOPS)
- Complete AI Agents system
- Workflow Guidance Engine

---

## 📝 Changelog

### Added
- Project-first documentation structure
- GitHub requirement enforcement
- GitHub badge on dashboard
- IDE integration for 7+ IDEs
- Smart context detection
- Dashboard JSON format
- Comprehensive error handling
- Theme system

### Changed
- CONTEXT.md now project-first instead of DoPlan-focused
- README.md restructured with DoPlan in collapsible section
- Rules generated to `.doplan/ai/rules/` instead of `.cursor/rules/`
- Commands symlinked from `.doplan/ai/commands/`

### Fixed
- Import cycles resolved
- Template compilation issues
- Test assertions updated
- Symlink creation preserves files

---

**Ready for Production Testing** ✅

All features planned for v0.0.18-beta are complete and tested. The codebase is stable and ready for comprehensive manual testing before release.

