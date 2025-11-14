# Changelog

All notable changes to DoPlan CLI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.14-beta] - 2025-11-14

### Changed
- Updated npm registry URLs to HTTPS everywhere for TLS 1.2+ compliance
- Fixed npm registry configuration for secure publishing

## [0.0.13-beta] - 2025-11-14

### Changed
- Version bump to v0.0.13-beta

## [0.0.12-beta] - 2025-11-14

### Initial Release

This is the first beta release of DoPlan CLI, a comprehensive project workflow automation tool that transforms app ideas into structured development projects.

### Features
- **Project Workflow Automation**: Transform app ideas into structured development projects
- **IDE Integration**: Support for Cursor, Gemini CLI, Claude CLI, Codex CLI, OpenCode, and Qwen Code
- **GitHub Integration**: Automatic branch creation, commit tracking, and PR management
- **Progress Tracking**: Real-time progress monitoring with visual dashboards
- **Document Generation**: Auto-generate PRD, project structure, API contracts, and planning documents
- **Interactive TUI**: Fullscreen terminal user interface with interactive dashboard
- **Statistics & Analytics**: Comprehensive project insights with historical tracking
- **Checkpoint System**: Create, list, and restore project checkpoints (Time Machine)

### Infrastructure
- Multi-platform builds (Linux, macOS, Windows) for amd64 and arm64 architectures
- Automated release process with GoReleaser
- Debian (.deb) and RPM (.rpm) packages
- Homebrew v2 syntax support (ready for tap repository)
- npm publishing support with package `doplan-cli`
- Comprehensive CI/CD workflows
- npm token verification and maintainer checks

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

[Unreleased]: https://github.com/DoPlan-dev/CLI/compare/v0.0.14-beta...HEAD
[0.0.14-beta]: https://github.com/DoPlan-dev/CLI/releases/tag/v0.0.14-beta
[0.0.13-beta]: https://github.com/DoPlan-dev/CLI/releases/tag/v0.0.13-beta
[0.0.12-beta]: https://github.com/DoPlan-dev/CLI/releases/tag/v0.0.12-beta
