# Changelog

All notable changes to DoPlan CLI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.4-alpha] - 2025-11-14

### Fixed
- **Homebrew v2 Syntax**: Updated Homebrew formula to use v2 syntax (removed deprecated `bottle :unneeded`)
  - Platform-specific `install` methods inside each `on_macos` and `on_linux` block
  - Added `Hardware::CPU.is_64_bit?` checks for Linux (v2 requirement)
  - Updated formula templates in `Formula/doplan.rb` and setup scripts
  - GoReleaser automatically generates v2-compliant formulas

### Changed
- **npm Publishing**: Enabled npm publish job in release workflow
  - `bin/doplan.js` wrapper script is ready and verified
  - npm package structure complete and ready for publishing
  - Automated npm publishing will trigger on tag pushes (requires `NPM_TOKEN` secret)

### Documentation
- Updated `docs/development/HOMEBREW_SETUP.md` with Homebrew v2 syntax examples and notes
- Added notes about GoReleaser's automatic v2 syntax generation

## [1.0.0] - 2025-11-14

### Added

#### Core Features
- **Initial release** of DoPlan CLI - comprehensive project workflow automation tool
- **Automated Project Setup** with IDE integration in seconds
- **Phase & Feature Management** - organize projects into phases and features with automatic directory structure
- **Progress Tracking** - real-time progress monitoring with visual dashboards (markdown and HTML)
- **Document Generation** - auto-generate PRD, project structure, API contracts, and planning documents

#### Interactive Interface
- **Fullscreen TUI** - beautiful terminal user interface with interactive dashboard
- **Visual Progress Bars** - see project, phase, and feature progress at a glance
- **Multi-view Dashboard** - switch between project overview, phases, features, GitHub activity, configuration, and statistics
- **Real-time Refresh** - update dashboard data with single key press

#### GitHub Integration
- **Automatic Branching** - create feature branches automatically
- **Auto-PR Creation** - automatically create pull requests when features complete
- **Commit Management** - track commits, pushes, and branch status
- **PR Tracking** - monitor pull request status and URLs
- **GitHub Data Sync** - sync branches, commits, and PRs from GitHub

#### Documentation & Templates
- **Template System** - customizable templates for plans, designs, and tasks
- **Template Management** - add, edit, remove, and set default templates
- **Context Generation** - auto-generate `CONTEXT.md` with tech stack and documentation links
- **Documentation Rules** - automated rules for documentation organization and naming

#### Configuration & Validation
- **Flexible Configuration** - manage settings for GitHub, checkpoints, and workflow
- **Project Validation** - validate project structure and configuration with auto-fix
- **Checkpoint System** - create, list, and restore project checkpoints (Time Machine)
- **Auto-Checkpointing** - automatic checkpoints for features and phases

#### Statistics & Analytics
- **Statistics Feature** - comprehensive project insights with historical tracking
- **Metrics Calculation** - velocity, completion, time, and quality metrics
- **Historical Data Storage** - automatic tracking of statistics over time
- **Trend Analysis** - identify improving/declining metrics with percentage changes
- **Multiple Output Formats** - CLI table, JSON, HTML, and Markdown formats
- **Time Range Filtering** - filter statistics by date range (`--since`, `--range`)
- **Metrics Filtering** - filter specific metrics (`--metrics`)

#### Error Handling
- **Comprehensive Error System** - user-friendly error messages with actionable suggestions
- **Error Logging** - automatic error logging to `.doplan/errors.json`
- **Error Recovery** - automatic fixing of common issues (AutoFix, Prompt, Rollback, Skip)
- **Context-Aware Errors** - detailed error messages with file paths and error codes

#### IDE Integration
- **Multi-IDE Support** - works with Cursor, Gemini CLI, Claude CLI, Codex CLI, OpenCode, and Qwen Code
- **Custom Commands** - IDE-specific commands for seamless workflow integration
- **Workflow Rules** - automated rules and conventions for development process

#### CLI Commands
- `doplan install` - Install DoPlan in your project (interactive IDE selection)
- `doplan dashboard` - View project dashboard with progress and GitHub activity
- `doplan --tui` - Launch fullscreen interactive TUI dashboard
- `doplan github` - Sync GitHub data (branches, commits, PRs) and update dashboard
- `doplan progress` - Update all progress tracking files and regenerate dashboard
- `doplan validate` - Validate project structure, configuration, and state consistency
- `doplan stats` - Display comprehensive project statistics with trends and filtering
- `doplan config` - Manage configuration settings (show, set, reset, validate)
- `doplan checkpoint` - Create, list, and restore project checkpoints
- `doplan templates` - Manage document templates (list, show, add, edit, use, remove)

### Infrastructure
- Multi-platform builds (Linux, macOS, Windows) for amd64 and arm64 architectures
- Automated release process with GoReleaser
- Debian (.deb) and RPM (.rpm) packages
- Checksum verification for all releases
- GitHub Actions release workflow
- Comprehensive test coverage
- Error handling and logging infrastructure

### Technical Details
- Built with Go 1.21+
- Zero external runtime dependencies
- Statically linked binaries
- Cross-platform support (Linux, macOS, Windows)
- Support for both amd64 and arm64 architectures

### Documentation
- Comprehensive README with installation instructions
- API documentation
- User guides and tutorials
- Development documentation
- Contribution guidelines
- Release process documentation

[Unreleased]: https://github.com/DoPlan-dev/CLI/compare/v1.0.4-alpha...HEAD
[1.0.4-alpha]: https://github.com/DoPlan-dev/CLI/releases/tag/v1.0.4-alpha
[1.0.0]: https://github.com/DoPlan-dev/CLI/releases/tag/v1.0.0

