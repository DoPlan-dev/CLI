# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Task Progress Automation**: Automatic task completion with dependency checking
  - `/finished` command now automatically updates TASKS.md task status and checklist items
  - Dependency validation blocks task completion if required dependencies are unfinished
  - Integrated with progress reporting for automatic percentage recalculation
- **`/plan` Command**: Scaffold structured planning hierarchy with phase folders, feature folders, and contract directories
  - Automatically parses TASKS.md to create phase folders (e.g., `01-foundation`, `02-core_features`)
  - Generates feature folders for each task with templates: `design.md`, `plan.md`, `tasks.md`, `prompts.md`, `github.md`
  - Creates `_contracts/` directory in each phase for shared API/data schemas
- **Clean Root Enforcement Rules**: Automated documentation organization validation
  - Enhanced `scripts/check-docs-organization.sh` to validate `Docs/` structure (capital D)
  - Added `docs-check` job to CI/CD workflow to block non-compliant PRs
  - Documented clean root policy in PRD, ARCHITECTURE, and contributor guidelines
  - Enforces that all documentation lives under `Docs/` with canonical subdirectories
- **GitHub Wiki Documentation**: Complete wiki with 20 comprehensive pages covering installation, usage, workflow, agents, rules, examples, and more
- **Wiki Pages**: Home, Installation, Quick Start, Commands, FAQ, Troubleshooting, Workflow, Agents, Rules, First Project Tutorial, Contributing, Project Structure, Configuration, API Reference, Examples, Migration Guide, Development, Code of Conduct, Release Notes, and Wiki Maintenance Plan

### Planned
- Additional project types (Tauri, Expo, etc.)
- Extended rules library (1000+ files)
- Custom agent templates
- Project templates marketplace

## [1.0.0] - 2024-11-23

### Added
- **Interactive TUI Wizard**: Beautiful step-by-step project creation with Bubbletea
- **18 Hierarchical AI Agents**: Complete agent system from Project Orchestrator to Documentation Writer
- **19 Commands**: 11 core commands + 8 squad-specific commands
- **Rules Library**: 15 categories with 500+ embedded rules for all major tech stacks
- **GitHub Workflows**: Automated CI/CD, releases, changelog management, and branch protection
- **IDE Support**: Configuration for 6 AI-powered IDEs (Cursor, Claude Code, Antigravity, Windsurf, Cline, OpenCode)
- **Project Generation**: Complete project structure with boilerplate code
- **Next.js Boilerplate**: Full Next.js 15.2.1 setup with React 19, TypeScript 5.6, Tailwind CSS
- **Documentation Generation**: README, CHANGELOG, STANDUP, and rules README templates
- **Planning Documents**: Templates for IDEA, PRD, ARCHITECTURE, and DESIGN_SYSTEM
- **Cross-Platform Support**: macOS (Intel + Apple Silicon), Linux (amd64 + arm64), Windows (amd64)
- **Build System**: Multi-platform build scripts and Makefile
- **Version System**: Build-time version injection with ldflags
- **npm/npx Support**: Package for easy installation via npm
- **Security Features**: Input validation, path sanitization, permission checks
- **Comprehensive Testing**: 200+ tests with 80%+ coverage
- **End-to-End Tests**: Complete integration test suite
- **Performance**: Generation time < 50ms, binary size 8.2MB

### Security
- Comprehensive security audit completed
- No hardcoded secrets
- Path traversal protection
- Input validation and sanitization
- Proper file permissions

### Documentation
- Complete README with quick start guide
- BUILD.md with build and distribution instructions
- TESTING.md with testing documentation
- SECURITY_AUDIT.md with security review
- DOCUMENTATION_REVIEW.md with documentation review

### Technical
- Pure Go implementation (no external runtime dependencies)
- Offline-capable after first run
- Atomic file writes for data integrity
- Error handling with rollback support
- Cross-platform path handling

## [0.1.0] - 2024-11-23

### Added
- Initial project structure
- DoPlan AI agency setup
- Core commands and agents
- Rules library integration
