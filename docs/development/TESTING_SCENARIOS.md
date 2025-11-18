# Testing Scenarios - Detailed Test Cases & Expected Outcomes

This document provides comprehensive test cases for all v0.0.18-beta features with expected outcomes.

## 📋 Test Environment Setup

### Prerequisites
```bash
# Test projects
- Empty directory: /tmp/doplan-test-empty
- Existing project: /tmp/doplan-test-existing
- Old DoPlan project: /tmp/doplan-test-old
- New DoPlan project: /tmp/doplan-test-new

# Test tools
- Go 1.24+
- Git installed
- GitHub CLI (gh) installed
- Terminal with TUI support
```

## 🧪 Test Suite 1: Context Detection

### Test 1.1: Empty Folder Detection

**Setup**:
```bash
cd /tmp/doplan-test-empty
rm -rf .doplan .cursor doplan
```

**Test Steps**:
1. Run `doplan`
2. Observe behavior

**Expected Outcome**:
- ✅ Detects `StateEmptyFolder`
- ✅ Launches New Project Wizard
- ✅ Shows welcome screen
- ✅ No errors

### Test 1.2: Existing Code Without DoPlan

**Expected Outcome**:
- ✅ Detects `StateExistingCodeNoDoPlan`
- ✅ Launches Adoption Wizard
- ✅ Shows "Found existing project!" message
- ✅ Analyzes project structure

### Test 1.3: Old DoPlan Structure Detection

**Expected Outcome**:
- ✅ Detects `StateOldDoPlanStructure`
- ✅ Launches Migration Wizard
- ✅ Shows migration detection screen
- ✅ Offers to migrate

### Test 1.4: New DoPlan Structure Detection

**Expected Outcome**:
- ✅ Detects `StateNewDoPlanStructure`
- ✅ Opens main TUI dashboard
- ✅ Loads project data
- ✅ No migration prompt

### Test 1.5: Inside Feature Directory Detection

**Expected Outcome**:
- ✅ Detects `StateInsideFeature`
- ✅ Shows feature-specific view
- ✅ Displays feature progress
- ✅ Shows feature tasks

## 🧪 Test Suite 2: New Project Wizard

### Test 2.1: Complete Wizard Flow

**Expected Outcome**:
- ✅ Welcome screen appears
- ✅ Project name input works
- ✅ Template selection works
- ✅ GitHub setup completes
- ✅ IDE selection works
- ✅ Installation completes
- ✅ Dashboard opens automatically

### Test 2.2: Project Name Validation

**Test Cases**:
| Input | Expected | Reason |
|-------|----------|--------|
| "test-project" | ✅ Valid | Valid format |
| "Test Project" | ❌ Invalid | Contains space |
| "test_project" | ❌ Invalid | Contains underscore |
| "test" | ❌ Invalid | Too short |
| "test-project-123" | ✅ Valid | Valid format |

### Test 2.3: Template Selection

**Expected Outcome**:
- ✅ All templates are selectable
- ✅ Preview updates for each template
- ✅ Template description is accurate
- ✅ Selection works with Enter key

### Test 2.4: GitHub Repository Validation

**Test Cases**:
| Input | Expected | Reason |
|-------|----------|--------|
| "https://github.com/user/repo" | ✅ Valid | Valid URL |
| "git@github.com:user/repo.git" | ✅ Valid | SSH URL |
| "user/repo" | ✅ Valid | Short format |
| "invalid-url" | ❌ Invalid | Not a URL |

### Test 2.5: IDE Selection

**Expected Outcome**:
- ✅ All IDE options available
- ✅ Selection works correctly
- ✅ Installation uses selected IDE
- ✅ IDE integration files created

## 🧪 Test Suite 3: Project Adoption Wizard

### Test 3.1: Project Analysis

**Expected Outcome**:
- ✅ Detects tech stack (Next.js, TypeScript)
- ✅ Finds project files
- ✅ Extracts documentation
- ✅ Identifies potential features
- ✅ Shows analysis results

### Test 3.2: Adoption Options

**Expected Outcome**:
- ✅ Three options available
- ✅ Each option works correctly
- ✅ Can navigate back

### Test 3.3: Auto-Plan Generation

**Expected Outcome**:
- ✅ Plan generated from code structure
- ✅ Phases created correctly
- ✅ Features identified
- ✅ Tasks extracted from TODOs
- ✅ Plan is accurate

## 🧪 Test Suite 4: Migration Wizard

### Test 4.1: Migration Detection

**Expected Outcome**:
- ✅ Detects old structure
- ✅ Shows migration screen
- ✅ Explains what will be migrated
- ✅ Offers to proceed or skip

### Test 4.2: Backup Creation

**Expected Outcome**:
- ✅ Backup created in `.doplan/backup/TIMESTAMP/`
- ✅ All files backed up
- ✅ Backup is complete
- ✅ Can restore from backup

### Test 4.3: Config Migration

**Expected Outcome**:
- ✅ `doplan-config.json` → `config.yaml`
- ✅ All fields migrated correctly
- ✅ New fields added with defaults
- ✅ Config is valid YAML
- ✅ Config validation passes

### Test 4.4: Folder Renaming

**Expected Outcome**:
- ✅ `01-phase` → `01-user-authentication` (or similar)
- ✅ `01-Feature` → `01-login-with-email` (or similar)
- ✅ All files copied correctly
- ✅ References updated in files
- ✅ Old folders removed (after verification)

### Test 4.5: Migration Rollback

**Expected Outcome**:
- ✅ Old structure restored
- ✅ New structure removed
- ✅ All files restored correctly
- ✅ Project works as before

## 🧪 Test Suite 5: Dashboard TUI

### Test 5.1: Dashboard Loading

**Expected Outcome**:
- ✅ Dashboard loads in <100ms
- ✅ Shows project name
- ✅ Shows progress bars
- ✅ Shows phase list
- ✅ Shows feature list
- ✅ Shows GitHub activity
- ✅ No errors

### Test 5.2: Dashboard Navigation

**Expected Outcome**:
- ✅ Press `1` → Dashboard view
- ✅ Press `2` → Phases view
- ✅ Press `3` → Features view
- ✅ Press `4` → GitHub view
- ✅ Press `5` → Config view
- ✅ Press `6` → Stats view
- ✅ Navigation is smooth

### Test 5.3: Progress Bar Accuracy

**Expected Outcome**:
- ✅ Overall progress is accurate
- ✅ Phase progress is accurate
- ✅ Feature progress is accurate
- ✅ Progress bars render correctly
- ✅ Colors match status

### Test 5.4: Real-time Updates

**Expected Outcome**:
- ✅ Refresh works
- ✅ Data updates correctly
- ✅ Progress recalculated
- ✅ GitHub data refreshed
- ✅ No flickering

## 🧪 Test Suite 6: IDE Integration

### Test 6.1: Cursor Integration

**Expected Outcome**:
- ✅ `.cursor/agents/` → symlink to `.doplan/ai/agents/`
- ✅ `.cursor/rules/` → symlink to `.doplan/ai/rules/`
- ✅ `.cursor/commands/` → symlink to `.doplan/ai/commands/`
- ✅ Symlinks work correctly
- ✅ Files are accessible

### Test 6.2: VS Code Integration

**Expected Outcome**:
- ✅ `.vscode/tasks.json` created
- ✅ `.vscode/settings.json` created
- ✅ `.vscode/prompts/` directory created
- ✅ Files contain correct content
- ✅ Tasks are executable

### Test 6.3: Generic IDE Integration

**Expected Outcome**:
- ✅ `.doplan/guides/generic_ide_setup.md` created
- ✅ Guide is comprehensive
- ✅ Instructions are clear

## 🧪 Test Suite 7: Error Handling

### Test 7.1: Invalid Project Name

**Expected Outcome**:
- ✅ Error message displayed
- ✅ Error is helpful
- ✅ Can correct and retry
- ✅ No crash

### Test 7.2: GitHub API Failure

**Expected Outcome**:
- ✅ Error message displayed
- ✅ Suggests checking URL
- ✅ Can retry or skip
- ✅ No crash

### Test 7.3: Migration Failure

**Expected Outcome**:
- ✅ Error caught gracefully
- ✅ Rollback option offered
- ✅ Backup preserved
- ✅ Clear error message

### Test 7.4: Dashboard Load Failure

**Expected Outcome**:
- ✅ Error caught gracefully
- ✅ Fallback to markdown dashboard
- ✅ Error message displayed
- ✅ Can retry

## 🧪 Test Suite 8: Performance

### Test 8.1: Dashboard Load Time

**Expected Outcome**:
- ✅ Loads in <100ms
- ✅ No blocking operations
- ✅ Smooth rendering

### Test 8.2: Large Project Handling

**Expected Outcome**:
- ✅ Dashboard loads in reasonable time
- ✅ Navigation is smooth
- ✅ No memory issues
- ✅ Pagination works if needed

## 📊 Test Results Template

```markdown
## Test Results - [Date]

### Test Suite 1: Context Detection
- [ ] Test 1.1: Empty Folder Detection - ✅ PASS
- [ ] Test 1.2: Existing Code Detection - ✅ PASS
- [ ] Test 1.3: Old Structure Detection - ✅ PASS
- [ ] Test 1.4: New Structure Detection - ✅ PASS
- [ ] Test 1.5: Inside Feature Detection - ✅ PASS

**Summary**: X/Y tests passing
```

## 🔄 Continuous Testing

### Automated Test Script
```bash
#!/bin/bash
# run-all-tests.sh

echo "Running DoPlan v0.0.18-beta Test Suite"

# Setup test environment
./scripts/setup-test-env.sh

# Run test suites
go test ./internal/migration/... -v
go test ./internal/context/... -v
go test ./internal/wizard/... -v
go test ./internal/integration/... -v
go test ./internal/tui/... -v

# Integration tests
./test/integration/test-wizard-flow.sh
./test/integration/test-migration.sh
./test/integration/test-ide-integration.sh

echo "Test suite complete"
```

