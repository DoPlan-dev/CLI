# DoPlan v0.0.18-beta Workflow Examples

Real-world workflow examples for common development tasks.

## 📋 Example 1: Complete Feature Implementation

**Goal:** Implement the new project wizard from scratch

### Day 1: Planning and Setup

```bash
# 1. Review documentation
cat docs/development/TUI_WIZARD_FLOW.md | grep -A 20 "Wizard 1"

# 2. Set up test environment
make setup

# 3. Create feature branch
git checkout -b feature/new-project-wizard

# 4. Generate base code
go run scripts/generate-wizard.go NewProject

# 5. Review generated code
cat internal/wizard/new_project.go
```

### Day 2: Implement Welcome Screen

```bash
# 1. Implement welcome screen
# Edit internal/wizard/new_project.go

# 2. Test interactively
./scripts/test-tui-interactive.sh
# Select option 1

# 3. Fix issues
# Repeat until working

# 4. Add unit tests
# Edit internal/wizard/new_project_test.go
```

### Day 3: Implement Input Screens

```bash
# 1. Implement project name input
# Add validation

# 2. Test validation
./scripts/test-error-scenarios.sh
# Check USR1001 errors

# 3. Implement template selection
# Add preview functionality

# 4. Test all inputs
./scripts/test-tui-interactive.sh
```

### Day 4: Implement GitHub and IDE Setup

```bash
# 1. Implement GitHub URL input
# Add validation

# 2. Test GitHub validation
# Use test-error-scenarios.sh

# 3. Implement IDE selection
# Add integration setup

# 4. Test IDE integration
./scripts/test-ide-integration.sh
```

### Day 5: Implement Installation and Testing

```bash
# 1. Implement installation progress
# Add async operations

# 2. Implement success screen
# Add dashboard auto-open

# 3. Full integration test
make test

# 4. Performance test
make benchmark

# 5. Documentation
# Update TUI_WIZARD_FLOW.md with actual implementation
```

## 📋 Example 2: Migration Implementation

**Goal:** Implement migration from v0.0.17 to v0.0.18

### Week 1: Core Migration Logic

```bash
# Day 1-2: Detection
# Implement DetectOldStructure()
# Test with test projects

# Day 3-4: Backup
# Implement CreateBackup()
# Test backup/restore

# Day 5: Config Migration
# Implement MigrateConfig()
# Test with various configs
```

### Week 2: Folder Migration

```bash
# Day 1-2: Folder Detection
# Implement DetectOldFolders()
# Test detection accuracy

# Day 3-4: Name Generation
# Implement GenerateSlugName()
# Test name extraction

# Day 5: Folder Renaming
# Implement MigrateFolders()
# Test with real projects
```

### Week 3: Integration and Testing

```bash
# Day 1-2: Integration
# Integrate all migration steps
# Test complete flow

# Day 3: Error Handling
# Add error recovery
# Test error scenarios

# Day 4-5: Testing
# Run all migration tests
# Test with real projects
```

## 📋 Example 3: IDE Integration

**Goal:** Add support for a new IDE

### Step 1: Research

```bash
# 1. Study IDE documentation
# Find configuration locations
# Understand file formats

# 2. Review existing integrations
cat internal/integration/cursor.go
cat internal/integration/vscode.go

# 3. Check integration docs
cat docs/development/IDE_INTEGRATION_SPECIFICS.md
```

### Step 2: Implementation

```bash
# 1. Create integration function
# Edit internal/integration/newide.go

# 2. Add to supported IDEs
# Edit internal/integration/ide.go

# 3. Create setup guide
# Create .doplan/guides/newide_setup.md
```

### Step 3: Testing

```bash
# 1. Test setup
doplan setup-ide newide

# 2. Verify integration
./scripts/verify-ide-integration.sh

# 3. Test in actual IDE
# Open IDE and verify files are accessible

# 4. Update tests
# Add to test-ide-integration.sh
```

## 📋 Example 4: Bug Fix Workflow

**Goal:** Fix a bug in dashboard loading

### Step 1: Reproduce

```bash
# 1. Reproduce bug
cd /tmp/doplan-test/new
doplan
# Observe error

# 2. Enable debug
export DOPLAN_DEBUG=1
doplan > debug.log 2>&1

# 3. Check logs
cat .doplan/logs/doplan.log
cat debug.log
```

### Step 2: Diagnose

```bash
# 1. Check dashboard.json
cat .doplan/dashboard.json | jq .

# 2. Validate JSON
jq '.' .doplan/dashboard.json

# 3. Check related files
find doplan -name "progress.json" | head -5 | xargs cat
```

### Step 3: Fix

```bash
# 1. Create fix branch
git checkout -b fix/dashboard-loading

# 2. Implement fix
# Edit internal/dashboard/generator.go

# 3. Add test
# Edit internal/dashboard/generator_test.go
```

### Step 4: Verify

```bash
# 1. Test fix
make test

# 2. Test with real project
./scripts/test-dashboard-generation.sh

# 3. Performance test
make benchmark

# 4. Integration test
./scripts/test-tui-interactive.sh
# Select dashboard
```

## 📋 Example 5: Performance Optimization

**Goal:** Improve dashboard load time

### Step 1: Measure

```bash
# 1. Baseline
./scripts/benchmark-performance.sh

# 2. Profile
go test -cpuprofile=cpu.prof ./internal/dashboard/...
go tool pprof cpu.prof
```

### Step 2: Identify Bottlenecks

```go
// Add timing
start := time.Now()
// ... operation ...
duration := time.Since(start)
log.Printf("Operation took: %v", duration)
```

### Step 3: Optimize

```bash
# 1. Implement optimizations
# Add caching
# Add lazy loading
# Batch operations

# 2. Test
make benchmark

# 3. Verify improvement
./scripts/benchmark-performance.sh
```

## 🔄 Daily Development Routine

### Morning Routine

```bash
# 1. Check dependencies
make deps

# 2. Update test environment
make setup

# 3. Run tests
make test-unit

# 4. Check for issues
git status
```

### During Development

```bash
# 1. Run tests frequently
go test ./internal/... -short

# 2. Test interactively
./scripts/test-tui-interactive.sh

# 3. Validate as you go
./scripts/validate-config.sh
```

### End of Day

```bash
# 1. Run full test suite
make test

# 2. Check performance
make benchmark

# 3. Backup work
git commit -am "WIP: feature description"
git push

# 4. Clean up
# Remove temporary files
```

## 🎯 Sprint Planning Example

### Sprint Goal
Implement migration wizard with full error handling

### Tasks

**Week 1:**
- [ ] Day 1: Migration detection
- [ ] Day 2: Backup creation
- [ ] Day 3: Config migration
- [ ] Day 4: Folder detection
- [ ] Day 5: Testing and fixes

**Week 2:**
- [ ] Day 1: Folder renaming
- [ ] Day 2: Reference updates
- [ ] Day 3: Error handling
- [ ] Day 4: Integration testing
- [ ] Day 5: Documentation

### Daily Standup Template

```
Yesterday:
- Implemented X
- Fixed Y bug
- Tested Z feature

Today:
- Will implement A
- Will fix B issue
- Will test C scenario

Blockers:
- Need help with D
- Waiting on E
```

## 📊 Progress Tracking

### Track Implementation Progress

```bash
# Create progress file
cat > implementation-progress.md <<EOF
# Implementation Progress

## Completed
- [x] Context detection
- [x] New project wizard
- [ ] Adoption wizard (in progress)
- [ ] Migration wizard

## Next Up
- Adoption wizard completion
- Migration wizard start
EOF
```

### Track Test Coverage

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Check coverage
go tool cover -func=coverage.out | grep total
```

