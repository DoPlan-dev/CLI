# DoPlan v0.0.18-beta Preparation Checklist

This checklist breaks down all preparation work needed before implementing v0.0.18-beta changes.

## 📋 Pre-Implementation Assessment

### Current State Analysis
- [ ] **Document current architecture**
  - [ ] Map all existing `internal/` directories
  - [ ] List all current sub-commands (`install`, `dashboard`, `github`, `progress`, etc.)
  - [ ] Document current folder structure (`01-phase/01-Feature/`)
  - [ ] List all configuration files and their locations
  - [ ] Document current TUI implementation (`internal/ui/`)

- [ ] **Identify breaking changes**
  - [ ] List all sub-commands that will be removed
  - [ ] Document current folder naming convention
  - [ ] List all config file locations that will change
  - [ ] Identify all code that depends on old structure

- [ ] **Test project analysis**
  - [ ] Document structure of `Club Website` project
  - [ ] Document structure of `mediabubble` project
  - [ ] List all DoPlan-generated files in test projects
  - [ ] Identify custom configurations in test projects

## 🏗️ Phase 1: Architecture Setup Preparation

### 1.1 Directory Structure Planning
- [ ] **Create directory structure map**
  ```
  internal/
  ├── tui/              [NEW - Main TUI app]
  │   ├── app.go
  │   ├── navigation.go
  │   ├── components/
  │   └── screens/
  ├── context/          [NEW - Context detection]
  │   ├── detector.go
  │   └── analyzer.go
  ├── wizard/           [NEW - Wizards]
  │   ├── new_project.go
  │   └── adopt_project.go
  ├── integration/     [NEW - IDE integration]
  │   ├── wizard.go
  │   └── setup.go
  ├── commands/         [REFACTOR - Action logic]
  ├── deployment/       [NEW]
  ├── publisher/        [NEW]
  ├── security/         [NEW]
  ├── fixer/            [NEW]
  ├── sops/             [NEW]
  ├── rakd/             [NEW]
  ├── dpr/              [NEW]
  ├── agents/           [NEW]
  └── workflow/         [NEW]
  ```

- [ ] **Plan migration of existing code**
  - [ ] Map `internal/ui/` → `internal/tui/`
  - [ ] Identify reusable code from `internal/commands/`
  - [ ] Plan extraction of wizard logic
  - [ ] Plan context detection from existing install logic

### 1.2 Package Structure Planning
- [ ] **Create pkg/ utilities map**
  ```
  pkg/
  ├── theme/            [NEW - Lipgloss theme]
  ├── errors/           [REFACTOR from internal/error/]
  ├── logger/           [REFACTOR from internal/error/logger.go]
  └── animations/       [NEW]
  ```

- [ ] **Plan theme system**
  - [ ] Define color palette constants
  - [ ] Design reusable style components
  - [ ] Plan migration from current inline styles

### 1.3 Configuration System Planning
- [ ] **Design new config.yaml schema**
  - [ ] Document all new fields (ide, github.repository, etc.)
  - [ ] Plan migration from `.cursor/config/doplan-config.json`
  - [ ] Design backward compatibility layer
  - [ ] Plan Viper integration

- [ ] **State management planning**
  - [ ] Design `.doplan/state.json` schema
  - [ ] Plan action history structure
  - [ ] Design checkpoint system integration

- [ ] **Dashboard JSON planning**
  - [ ] Design `.doplan/dashboard.json` schema
  - [ ] Plan migration from current dashboard.md generation
  - [ ] Design real-time update triggers

## 🔄 Phase 2: Migration Strategy Preparation

### 2.1 Backward Compatibility Planning
- [ ] **Design compatibility layer**
  - [ ] Plan detection of old folder structure (`01-phase/`)
  - [ ] Design auto-migration on first run
  - [ ] Plan support for both old and new structures during transition
  - [ ] Design migration wizard

- [ ] **Config migration planning**
  - [ ] Map `.cursor/config/doplan-config.json` → `.doplan/config.yaml`
  - [ ] Plan automatic migration on upgrade
  - [ ] Design validation for migrated config

### 2.2 Test Project Migration Plan
- [ ] **Club Website migration plan**
  - [ ] Document current structure
  - [ ] Plan folder rename strategy (`01-phase/` → `01-phase-name/`)
  - [ ] Plan config file migration
  - [ ] Design rollback strategy

- [ ] **mediabubble migration plan**
  - [ ] Document current structure
  - [ ] Plan folder rename strategy
  - [ ] Plan config file migration
  - [ ] Design rollback strategy

### 2.3 Data Migration Scripts
- [ ] **Design migration utilities**
  - [ ] Plan `internal/migration/` package
  - [ ] Design folder structure migrator
  - [ ] Design config file migrator
  - [ ] Design progress.json migrator
  - [ ] Design validation after migration

## 🧪 Phase 3: Testing Strategy Preparation

### 3.1 Test Environment Setup
- [ ] **Create test projects**
  - [ ] Clone test projects to separate directories
  - [ ] Create backup of test projects
  - [ ] Document test project states

- [ ] **Test scenarios planning**
  - [ ] Empty folder → new project wizard
  - [ ] Existing project → adoption wizard
  - [ ] Old DoPlan project → migration
  - [ ] New DoPlan project → full flow

### 3.2 Integration Testing Plan
- [ ] **IDE integration testing**
  - [ ] Plan Cursor integration tests
  - [ ] Plan VS Code integration tests
  - [ ] Plan generic IDE tests

- [ ] **TUI testing plan**
  - [ ] Plan dashboard loading tests
  - [ ] Plan wizard flow tests
  - [ ] Plan navigation tests

## 📚 Phase 4: Documentation Preparation

### 4.1 User Documentation
- [ ] **Migration guide**
  - [ ] Write upgrade guide for v0.0.17 → v0.0.18
  - [ ] Document breaking changes
  - [ ] Create migration examples

- [ ] **New feature documentation**
  - [ ] Document new TUI-first workflow
  - [ ] Document IDE integration setup
  - [ ] Document new folder structure

### 4.2 Developer Documentation
- [ ] **Architecture documentation**
  - [ ] Document new directory structure
  - [ ] Document new config system
  - [ ] Document migration utilities

- [ ] **API documentation**
  - [ ] Document new internal APIs
  - [ ] Document context detection API
  - [ ] Document wizard APIs

## 🛠️ Phase 5: Implementation Readiness

### 5.1 Code Organization
- [ ] **Create feature branches**
  - [ ] `feature/v0.0.18-architecture`
  - [ ] `feature/v0.0.18-migration`
  - [ ] `feature/v0.0.18-tui`

- [ ] **Set up development environment**
  - [ ] Ensure Go 1.24+ installed
  - [ ] Verify Bubble Tea dependencies
  - [ ] Verify Lipgloss dependencies
  - [ ] Set up test environment

### 5.2 Dependency Management
- [ ] **Review dependencies**
  - [ ] Verify Bubble Tea version compatibility
  - [ ] Verify Lipgloss version compatibility
  - [ ] Check for new dependencies needed
  - [ ] Plan dependency updates

### 5.3 Version Management
- [ ] **Version planning**
  - [ ] Update version to v0.0.18-beta
  - [ ] Plan release notes
  - [ ] Plan changelog updates

## ✅ Pre-Implementation Sign-off

Before starting implementation, ensure:
- [ ] All architecture decisions documented
- [ ] Migration strategy approved
- [ ] Test projects backed up
- [ ] Development environment ready
- [ ] Team aligned on approach
- [ ] Rollback plan in place

---

## 📝 Notes

### Key Decisions Needed
1. **Folder naming migration**: Automatic or manual?
2. **Config migration**: Silent or with user prompt?
3. **Backward compatibility**: How long to support old structure?
4. **Breaking changes**: How to communicate to users?

### Risks Identified
1. **Data loss risk** during folder migration
2. **Config corruption** during migration
3. **Breaking existing workflows** with sub-command removal
4. **IDE integration complexity** for multiple IDEs

### Mitigation Strategies
1. Always create backups before migration
2. Validate config after migration
3. Provide clear migration guide
4. Test IDE integrations thoroughly

