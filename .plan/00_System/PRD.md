# Product Requirements Document (PRD)
## DoPlan CLI v1.0

**Version**: 1.0  
**Status**: Draft  
**Last Updated**: November 2025  
**Owner**: Product Manager

---

## 📋 Executive Summary

**DoPlan CLI** is a zero-install, pure-Go command-line tool that instantly generates professional project structures with a complete hierarchical AI agency system. Users can bootstrap production-ready projects in seconds with full automation, intelligent agents, and comprehensive rules libraries.

**Key Value Propositions**:
- Zero-install: `npx doplan@latest` - no global installation required
- Offline-first: Works completely offline after first run
- Intelligence in files: All AI logic lives in transparent markdown files
- IDE-agnostic: Supports 6 AI-powered IDEs
- Complete automation: Project structure, agents, commands, rules, CI/CD, and boilerplate

---

## 👥 User Personas

### Primary Persona: Solo Developer (Alex)
- **Age**: 28-35
- **Role**: Full-stack developer, freelancer or startup founder
- **Tech Stack**: Modern web technologies (Next.js, React, TypeScript)
- **Pain Points**:
  - Wastes 2-4 hours setting up new projects
  - Struggles with consistent project structure
  - Wants to focus on building, not configuration
  - Needs professional CI/CD but doesn't know how to set it up
- **Goals**:
  - Bootstrap projects in < 5 minutes
  - Have production-ready structure from day one
  - Use AI agents to accelerate development
- **Tech Comfort**: High - comfortable with CLI tools

### Secondary Persona: Small Team Lead (Sam)
- **Age**: 30-40
- **Role**: Engineering lead or CTO of 5-20 person team
- **Tech Stack**: Multiple projects, various frameworks
- **Pain Points**:
  - Team members create inconsistent project structures
  - Onboarding new developers takes too long
  - No standardized development workflow
  - CI/CD setup is repetitive and error-prone
- **Goals**:
  - Standardize project structure across team
  - Reduce onboarding time
  - Ensure all projects have best practices built-in
- **Tech Comfort**: Very High - makes architectural decisions

### Tertiary Persona: Enterprise Architect (Jordan)
- **Age**: 35-45
- **Role**: Enterprise architect or platform engineer
- **Tech Stack**: Multiple teams, various languages/frameworks
- **Pain Points**:
  - Need consistent project structures across 50+ projects
  - Compliance and security requirements
  - Documentation standards
  - Governance and best practices enforcement
- **Goals**:
  - Enforce organizational standards
  - Reduce security vulnerabilities
  - Improve developer productivity at scale
- **Tech Comfort**: Expert - designs systems for large organizations

---

## 🎯 Product Goals

### Primary Goals
1. **Speed**: Generate complete project in < 5 seconds
2. **Quality**: Production-ready structure with best practices
3. **Simplicity**: 2-question wizard (name, IDE choice)
4. **Reliability**: < 1% error rate, works offline
5. **Adoption**: 10,000+ projects created in first 6 months

### Success Metrics
- **Adoption**: 10,000+ projects created in first 6 months
- **Engagement**: Average 5+ commands used per project
- **Retention**: 30%+ users create second project
- **Community**: 100+ GitHub stars, active discussions
- **Quality**: < 1% bug reports, 4.5+ star rating
- **Performance**: 95%+ of projects generated in < 5 seconds

---

## 🚀 Features

### MVP Features (v1.0)

#### 1. Interactive TUI Wizard
**Priority**: P0 (Must Have)  
**Description**: Beautiful, keyboard-navigable wizard built with Bubbletea

**User Stories**:
- As a developer, I want to see a beautiful welcome screen when I run DoPlan
- As a developer, I want to enter my project name with real-time validation
- As a developer, I want to select my IDE from a visual menu
- As a developer, I want to see progress indicators during generation

**Acceptance Criteria**:
- ✅ TUI displays ASCII art header with emojis
- ✅ Project name input validates in real-time (alphanumeric + hyphens/underscores)
- ✅ IDE selection menu shows all 6 IDEs with descriptions
- ✅ Progress screen shows status for each generation step
- ✅ Success screen displays clear next steps
- ✅ Full keyboard navigation (↑/↓, Enter, q to quit)
- ✅ Color-coded output (purple/pink primary, green success, red errors)

#### 2. Project Structure Generation
**Priority**: P0 (Must Have)  
**Description**: Generate complete directory structure with all required folders

**User Stories**:
- As a developer, I want a complete project structure generated automatically
- As a developer, I want the structure to follow best practices
- As a developer, I want all necessary directories created

**Acceptance Criteria**:
- ✅ Creates `.cursor/agents/` directory with 18 agent files
- ✅ Creates `.cursor/commands/` directory with 11+ command files
- ✅ Creates `.cursor/rules/library/` with 15 category directories
- ✅ Creates `.plan/00_System/` with all planning documents
- ✅ Creates `.github/workflows/` with 4 workflow files
- ✅ Creates `src/` directory with boilerplate code
- ✅ Creates all IDE config files based on selection

#### 3. AI Agent System
**Priority**: P0 (Must Have)  
**Description**: Generate 18 hierarchical AI agent persona files

**User Stories**:
- As a developer, I want AI agents that understand my project context
- As a developer, I want agents organized in a clear hierarchy
- As a developer, I want each agent to have clear responsibilities

**Acceptance Criteria**:
- ✅ Generates 18 agent markdown files in `.cursor/agents/`
- ✅ Each agent file includes: Role, System Prompt, Responsibilities, Reports To, Manages
- ✅ Agents follow hierarchical structure (Project Orchestrator at top)
- ✅ Agent prompts are comprehensive and actionable
- ✅ All agents reference the rules library

#### 4. Command System
**Priority**: P0 (Must Have)  
**Description**: Generate 11 core commands + squad-specific commands

**User Stories**:
- As a developer, I want simple slash commands to control my project
- As a developer, I want commands to activate appropriate agents
- As a developer, I want commands to be well-documented

**Acceptance Criteria**:
- ✅ Generates 11 core command files in `.cursor/commands/`
- ✅ Each command file includes: Trigger, Action, Agent Involvement, Files Read, Files Modified
- ✅ Commands: `/tell`, `/improve`, `/team`, `/write`, `/change`, `/good`, `/tasks`, `/load`, `/build`, `/progress`, `/finished`
- ✅ Squad-specific commands: `/secure`, `/roles`, `/money`, `/pretty`, `/seo`, `/ship`, `/safe`, `/cheap`
- ✅ All commands have clear documentation

#### 5. Rules Library
**Priority**: P0 (Must Have)  
**Description**: Extract 500+ embedded rules files to project

**User Stories**:
- As a developer, I want comprehensive rules covering all major tech stacks
- As a developer, I want rules organized by category
- As a developer, I want rules to be easily discoverable

**Acceptance Criteria**:
- ✅ Extracts 500+ rules files to `.cursor/rules/library/`
- ✅ Rules organized in 15 categories (01-core-workflow through 15-project-specific)
- ✅ Each category has README.md explaining purpose
- ✅ Rules cover: Go, Python, TypeScript, Next.js, React, Express, databases, testing, CI/CD, security
- ✅ Rules are properly formatted markdown

#### 6. GitHub Workflows
**Priority**: P0 (Must Have)  
**Description**: Generate 4 GitHub Actions workflows

**User Stories**:
- As a developer, I want CI/CD set up automatically
- As a developer, I want automated releases
- As a developer, I want changelog management

**Acceptance Criteria**:
- ✅ Generates `ci.yml` - runs on all branches, tests + lints
- ✅ Generates `release.yml` - automated releases on version tags
- ✅ Generates `changelog.yml` - auto-updates changelog
- ✅ Generates `branch-protection.yml` - PR requirements
- ✅ All workflows use best practices and are production-ready

#### 7. IDE Configuration
**Priority**: P0 (Must Have)  
**Description**: Generate IDE-specific config files

**User Stories**:
- As a developer, I want my IDE configured for AI agents
- As a developer, I want IDE config to reference agents and commands
- As a developer, I want to use my preferred IDE

**Acceptance Criteria**:
- ✅ Generates `.cursorrules` for Cursor
- ✅ Generates `CLAUDE.md` for Claude Code
- ✅ Config files reference agent hierarchy
- ✅ Config files reference command locations
- ✅ Config files reference rules library

#### 8. Boilerplate Code
**Priority**: P0 (Must Have)  
**Description**: Generate Next.js 15.2.1 + React 19 + TypeScript boilerplate

**User Stories**:
- As a developer, I want production-ready boilerplate code
- As a developer, I want latest framework versions
- As a developer, I want proper TypeScript configuration

**Acceptance Criteria**:
- ✅ Generates `package.json` with Next.js 15.2.1, React 19, TypeScript 5.6.0
- ✅ Generates `tsconfig.json` with proper configuration
- ✅ Generates `tailwind.config.ts` configured
- ✅ Generates basic Next.js app structure in `src/`
- ✅ Generates ESLint configuration
- ✅ All dependencies are latest stable versions

#### 9. Binary Size & Performance
**Priority**: P0 (Must Have)  
**Description**: Keep binary < 15MB, generation < 5 seconds

**User Stories**:
- As a developer, I want a small, fast CLI tool
- As a developer, I don't want to wait long for project generation

**Acceptance Criteria**:
- ✅ Binary size < 15MB (target: 10-12MB)
- ✅ Project generation completes in < 5 seconds
- ✅ TUI responds to keypresses in < 100ms
- ✅ Memory usage < 100MB during generation

#### 10. Offline Capability
**Priority**: P0 (Must Have)  
**Description**: Works completely offline after first run

**User Stories**:
- As a developer, I want to use DoPlan without internet
- As a developer, I want reliable tooling that doesn't depend on external services

**Acceptance Criteria**:
- ✅ All resources embedded in binary
- ✅ No network calls after initial download
- ✅ Works in airplane mode, remote locations
- ✅ Rules library embedded, not downloaded

### Post-MVP Features (v1.1+)

#### 11. Additional Project Types
- Tauri (desktop apps)
- Expo (mobile apps)
- Express (API-only)
- Remix
- SvelteKit

#### 12. Extended IDE Support
- Antigravity
- Windsurf
- Cline
- OpenCode

#### 13. Extended Rules Library
- Expand from 500+ to 1000+ rules
- Add more framework-specific rules
- Add more language-specific rules

#### 14. Custom Agent Templates
- Allow users to customize agent prompts
- Community-contributed agent templates

#### 15. Project Templates Marketplace
- Community-contributed project templates
- Template discovery and sharing

#### 16. `/plan` Command & Hierarchical Planning
- Scaffold multi-phase planning directories automatically
- Generate feature folders with design/plan/tasks/prompts/github templates
- Keep contracts/API schemas synchronized with planning docs

#### 17. Branch Automation in `/build` & `/finished`
- Auto-create/switch `task/TASK-###` branches when work starts or ends
- Enforce naming/pushing rules so work never lands on the wrong branch
- Surface branch metadata in `active_state.json` and scan reports

#### 18. Task Progress Automation
- Update `TASKS.md` status + completion metrics when `/finished` runs
- Reflect progress deltas inside `/progress`
- Block completion when required dependencies remain unfinished

#### 19. Scan Report Diffing
- Compare each `/scan` with the previous report automatically
- Highlight new/changed files, dependencies, and metrics deltas
- Attach delta tables directly inside scan outputs

#### 20. `/feedback` Command & Sync
- Collect structured feedback from within the CLI
- Sync feedback entries to the DoPlan CLI GitHub repo/issues
- Track feedback history per project

#### 21. Advanced GitHub Integration
- Auto-detect git remote + default branch configuration
- Open GitHub issues/milestones directly from feedback or scans
- Keep repo descriptions aligned with PRD/ROADMAP content

#### 22. CI/CD Workflow Generator for Task Branches
- Emit GitHub Actions tailored to task branches automatically
- Recommend branch protection rules + PR templates
- Provide reusable job presets per tech stack

#### 23. Project State History & Rollback
- Version `.plan/active_state.json` with timestamped snapshots
- Offer safe rollback/restore commands
- Ship a `/state` helper (backed by `scripts/statehistory`) for snapshot/list/diff/restore with explicit confirmations
- Show state deltas in scan reports and `/progress`

#### 24. Rich Scan Report Customization
- Offer report templates per project type/audience
- Embed visuals, charts, and dependency vulnerability summaries

#### 25. Clean Root Enforcement Rules
- Enforce documentation organization via lint scripts
- Block root-level `.md` files (except README.md and CHANGELOG.md)
- Require all documentation to live under `Docs/` with canonical structure
- Integrate checks into CI/CD pipeline
- Document policy in contributor guidelines
- Allow teams to extend report sections via configuration

#### 25. `Docs/` Folder & Clean Root Policy
- Scaffold a top-level `Docs/` directory (mirroring `test-no01`) with `foundation/`, `features/`, `release/`, and `history/` categories.
- Ensure every autogenerated document (PRD updates, feature specs, retros, prompt logs) lands inside `Docs/`—the repo root must remain limited to code, `README.md`, and `CHANGELOG.md`.
- Update documentation linting/rules so PRs cannot add markdown outside `Docs/` and always refresh `Docs/README.md`.

---

## 📊 User Journey

### Happy Path
1. **Discovery**: Developer hears about DoPlan via Twitter/Reddit/HackerNews
2. **First Run**: `npx doplan@latest` → Beautiful TUI wizard appears
3. **Project Creation**: 
   - Enters project name: "my-awesome-app"
   - Selects IDE: Cursor
   - Watches progress indicators
4. **Project Generated**: Success screen shows "Open with: code ./my-awesome-app"
5. **First Command**: Opens project in IDE → Types `/tell` → Captures idea
6. **Workflow**: Uses `/improve`, `/write`, `/build`, `/finished` seamlessly
7. **Magic Moment**: Realizes they've built a production-ready project with full AI agency in minutes

### Error Scenarios
1. **Directory Exists**: Clear error message with recovery suggestion
2. **Invalid Project Name**: Real-time validation prevents errors
3. **Network Issues**: Works offline, no network required
4. **Permission Errors**: Clear error message with fix instructions

---

## 🎯 Success Criteria

### Launch Criteria (v1.0)
- ✅ All 10 MVP features implemented
- ✅ Binary size < 15MB
- ✅ Generation time < 5 seconds
- ✅ 80%+ test coverage
- ✅ Comprehensive documentation
- ✅ Works on macOS, Linux, Windows
- ✅ Zero critical bugs

### Post-Launch Goals (6 months)
- 10,000+ projects created
- 30%+ user retention (second project)
- 100+ GitHub stars
- 4.5+ star rating
- < 1% bug report rate
- Active community discussions

---

## 📅 Timeline

### Phase 1: Foundation (Week 1-2)
- Project setup and structure
- TUI wizard implementation
- Basic project generation

### Phase 2: Core Features (Week 3-4)
- Agent generation
- Command generation
- Rules library extraction
- GitHub workflows

### Phase 3: Polish (Week 5-6)
- Boilerplate generation
- IDE configs
- Testing and bug fixes
- Documentation

### Phase 4: Release (Week 7)
- Final testing
- Release preparation
- Launch

**Total Timeline**: 6-7 weeks for MVP

---

## 🔒 Non-Goals (Out of Scope for v1.0)

- Web UI or dashboard
- Cloud-hosted agents
- Real-time collaboration
- Project hosting
- Database integration
- User authentication
- Analytics/telemetry (opt-in only)
- Multi-language support (English only for MVP)

---

## 📝 Notes

- All features must work offline
- No external dependencies except Go standard library + Bubbletea + Cobra
- All intelligence lives in generated markdown files
- Binary must be distributable via npx wrapper
- Focus on developer experience above all else

---

**Document Status**: ✅ Complete  
**Next Step**: Review and approve, then type `/good` to lock plan and generate tasks.
