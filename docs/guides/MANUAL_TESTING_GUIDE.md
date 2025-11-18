# Manual Testing Guide - DoPlan v0.0.18-beta

This guide provides step-by-step instructions for executing all manual test scenarios.

## Prerequisites

1. Build the DoPlan binary:
   ```bash
   # From the DoPlan CLI directory
   go build -o doplan ./cmd/doplan/main.go
   
   # Copy binary to test directories (or ensure it's in PATH)
   cp doplan /tmp/doplan-manual-test/empty/
   cp doplan /tmp/doplan-manual-test/existing/
   cp doplan /tmp/doplan-manual-test/old/
   cp doplan /tmp/doplan-manual-test/new/
   ```

2. Set up test directories:
   ```bash
   mkdir -p /tmp/doplan-manual-test/{empty,existing,old,new}
   ```

3. Ensure you have:
   - Terminal with TUI support
   - Git installed
   - GitHub CLI (gh) installed (optional, for GitHub tests)

**Note:** Throughout this guide, `./doplan` refers to the binary in the current test directory. If you have `doplan` in your PATH, you can use `doplan` instead.

## Test Suite 2: New Project Wizard

### Test 2.1: Complete Wizard Flow

**Setup:**
```bash
cd /tmp/doplan-manual-test/empty
rm -rf .doplan .cursor doplan 2>/dev/null || true
```

**Steps:**
1. Run: `./doplan` (or `doplan` if in PATH)
2. **Expected:** Welcome screen appears with DoPlan logo
3. Press Enter to continue
4. **Expected:** Project name input screen
5. Enter: `test-project-manual`
6. **Expected:** Name accepted, moves to template selection
7. Use arrow keys to navigate templates
8. Select: `web-app` (or any template)
9. Press Enter
10. **Expected:** GitHub repository input screen
11. Enter: `skip` (or a valid GitHub repo)
12. Press Enter
13. **Expected:** IDE selection screen
14. Select: `cursor` (or any IDE)
15. Press Enter
16. **Expected:** Installation progress screen
17. Wait for completion
18. **Expected:** Dashboard opens automatically

**Verification:**
```bash
# After wizard completes, verify:
ls -la .doplan/config.yaml  # Should exist
ls -la doplan/               # Should exist
cat .doplan/config.yaml | grep "test-project-manual"  # Should contain project name
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 2.2: Project Name Validation

**Steps:**
1. Run wizard again in empty directory
2. Try entering each name below and observe behavior:

| Input | Expected Result | Actual Result |
|-------|----------------|---------------|
| `test-project` | ✅ Valid - proceeds | |
| `Test Project` | ❌ Invalid - shows error | |
| `test_project` | ❌ Invalid - shows error | |
| `test` | ❌ Invalid - shows error (too short) | |
| `test-project-123` | ✅ Valid - proceeds | |

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 2.3: Template Selection

**Steps:**
1. Run wizard, enter valid project name
2. Navigate through templates using arrow keys
3. For each template:
   - Verify description appears
   - Verify preview updates
   - Press Enter to select
   - Verify selection works

**Templates to test:**
- web-app
- api-service
- mobile-app
- cli-tool
- library

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 2.4: GitHub Repository Validation

**Steps:**
1. Run wizard, complete project name and template
2. Test each GitHub input:

| Input | Expected Result | Actual Result |
|-------|----------------|---------------|
| `https://github.com/user/repo` | ✅ Valid | |
| `git@github.com:user/repo.git` | ✅ Valid | |
| `user/repo` | ✅ Valid | |
| `invalid-url` | ❌ Invalid - shows error | |
| `skip` | ✅ Valid - skips GitHub | |

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 2.5: IDE Selection

**Steps:**
1. Run wizard, complete previous steps
2. Navigate through IDE options
3. Select each IDE and verify:
   - Selection works
   - Installation completes
   - Integration files created

**IDEs to test:**
- cursor
- vscode
- generic

**After installation, verify:**
```bash
# For Cursor:
ls -la .cursor/agents/  # Should be symlink or directory
ls -la .cursor/rules/   # Should be symlink or directory

# For VS Code:
ls -la .vscode/tasks.json  # Should exist
ls -la .vscode/settings.json  # Should exist
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Suite 3: Project Adoption Wizard

### Test 3.1: Project Analysis

**Setup:**
```bash
cd /tmp/doplan-manual-test/existing
rm -rf .doplan .cursor doplan 2>/dev/null || true

# Create a sample project
mkdir -p src/components src/pages
echo "import React from 'react'" > src/components/App.tsx
echo "export default function App() {}" > src/components/App.tsx
echo '{"name": "test-app", "version": "1.0.0"}' > package.json
echo "# Test Project" > README.md
```

**Steps:**
1. Run: `./doplan` (or `doplan` if in PATH)
2. **Expected:** Adoption wizard launches
3. **Expected:** Shows "Found existing project!" message
4. Wait for analysis to complete
5. **Expected:** Analysis results show:
   - Tech stack detected (TypeScript, React)
   - Project files found
   - Documentation extracted
   - Potential features identified

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 3.2: Adoption Options

**Steps:**
1. After analysis completes, verify three options:
   - Option 1: Auto-generate plan from code
   - Option 2: Manual plan creation
   - Option 3: Skip adoption
2. Navigate between options
3. Select each option and verify behavior

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 3.3: Auto-Plan Generation

**Steps:**
1. Select "Auto-generate plan from code"
2. Wait for plan generation
3. **Expected:** Plan shows:
   - Phases created from code structure
   - Features identified
   - Tasks extracted from TODOs/comments
4. Verify plan accuracy

**Verification:**
```bash
# After generation:
cat .cursor/config/doplan-state.json | jq '.phases'  # Should have phases
cat .cursor/config/doplan-state.json | jq '.features'  # Should have features
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Suite 4: Migration Wizard

### Test 4.1: Migration Detection

**Setup:**
```bash
cd /tmp/doplan-manual-test/old
rm -rf .doplan .cursor doplan 2>/dev/null || true

# Create old structure
mkdir -p .cursor/config doplan/01-phase/01-Feature
echo '{"installed": true, "ide": "cursor"}' > .cursor/config/doplan-config.json
echo "# Feature 1" > doplan/01-phase/01-Feature/README.md
```

**Steps:**
1. Run: `./doplan` (or `doplan` if in PATH)
2. **Expected:** Migration wizard launches
3. **Expected:** Shows migration detection screen
4. **Expected:** Explains what will be migrated:
   - Config file migration
   - Folder renaming
   - Structure updates
5. **Expected:** Offers to proceed or skip

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 4.2: Backup Creation

**Steps:**
1. Select "Proceed with migration"
2. **Expected:** Backup creation starts
3. **Expected:** Shows backup location
4. Wait for backup to complete

**Verification:**
```bash
# After migration starts:
ls -la .doplan/backup/  # Should have timestamp directory
ls -la .doplan/backup/*/  # Should contain backed up files
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 4.3: Config Migration

**Steps:**
1. After backup, migration proceeds
2. **Expected:** Config migration completes
3. **Expected:** Shows migration summary

**Verification:**
```bash
# After migration:
cat .doplan/config.yaml  # Should exist and be valid YAML
cat .doplan/config.yaml | grep "cursor"  # Should contain IDE setting
# Old config should still exist in backup
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 4.4: Folder Renaming

**Steps:**
1. Migration continues after config
2. **Expected:** Folder renaming occurs
3. **Expected:** Shows progress for each folder

**Verification:**
```bash
# After migration:
ls -la doplan/  # Should show renamed folders (not 01-phase, 01-Feature)
# Old folders should be in backup
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 4.5: Migration Rollback

**Steps:**
1. If migration fails or you want to test rollback:
2. **Expected:** Rollback option available
3. Select rollback
4. **Expected:** Old structure restored
5. **Expected:** New structure removed

**Verification:**
```bash
# After rollback:
ls -la .cursor/config/doplan-config.json  # Should exist (old structure)
ls -la .doplan/config.yaml  # Should NOT exist (new structure removed)
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Suite 5: Dashboard TUI

### Test 5.1: Dashboard Loading

**Setup:**
```bash
cd /tmp/doplan-manual-test/new
# Use a project with DoPlan installed (from Test 2.1)
```

**Steps:**
1. Run: `./doplan` (or `doplan` if in PATH)
2. **Expected:** Dashboard loads quickly (<100ms ideally)
3. **Expected:** Shows:
   - Project name
   - Overall progress bar
   - Phase list
   - Feature list
   - GitHub activity (if configured)
4. **Expected:** No errors

**Result:** ☐ PASS ☐ FAIL
**Load Time:** _____ ms
**Notes:** 

---

### Test 5.2: Dashboard Navigation

**Steps:**
1. Dashboard is open
2. Press `1` → **Expected:** Dashboard view (overview)
3. Press `2` → **Expected:** Phases view
4. Press `3` → **Expected:** Features view
5. Press `4` → **Expected:** GitHub view
6. Press `5` → **Expected:** Config view
7. Press `6` → **Expected:** Stats view
8. **Expected:** Navigation is smooth, no flickering

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 5.3: Progress Bar Accuracy

**Steps:**
1. Navigate to Dashboard view (press `1`)
2. Verify overall progress bar:
   - Shows correct percentage
   - Bar length matches percentage
   - Color matches status (green=complete, yellow=in-progress, gray=todo)
3. Navigate to Phases view (press `2`)
4. Verify phase progress bars are accurate
5. Navigate to Features view (press `3`)
6. Verify feature progress bars are accurate

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 5.4: Real-time Updates

**Steps:**
1. Dashboard is open
2. Press `r` to refresh
3. **Expected:** Data updates
4. **Expected:** Progress recalculated
5. **Expected:** GitHub data refreshed (if configured)
6. **Expected:** No flickering or UI glitches

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Suite 6: IDE Integration

### Test 6.1: Cursor Integration

**Setup:**
```bash
cd /tmp/doplan-manual-test/new
# Use project installed with Cursor IDE
```

**Verification:**
```bash
# Check symlinks:
ls -la .cursor/agents/  # Should be symlink to .doplan/ai/agents/
ls -la .cursor/rules/   # Should be symlink to .doplan/ai/rules/
ls -la .cursor/commands/  # Should be symlink to .doplan/ai/commands/

# Verify symlinks work:
readlink .cursor/agents/  # Should show .doplan/ai/agents/
cat .cursor/agents/*.md  # Should be readable
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 6.2: VS Code Integration

**Setup:**
```bash
cd /tmp/doplan-manual-test/new
# Use project installed with VS Code IDE
```

**Verification:**
```bash
# Check files exist:
ls -la .vscode/tasks.json  # Should exist
ls -la .vscode/settings.json  # Should exist
ls -la .vscode/prompts/  # Should be directory

# Verify content:
cat .vscode/tasks.json | jq .  # Should be valid JSON
cat .vscode/settings.json | jq .  # Should be valid JSON
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 6.3: Generic IDE Integration

**Setup:**
```bash
cd /tmp/doplan-manual-test/new
# Use project installed with Generic IDE
```

**Verification:**
```bash
# Check guide exists:
ls -la .doplan/guides/generic_ide_setup.md  # Should exist
cat .doplan/guides/generic_ide_setup.md  # Should have content
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Suite 7: Error Handling

### Test 7.1: Invalid Project Name

**Steps:**
1. Run new project wizard
2. Enter invalid name: `Test Project` (with space)
3. **Expected:** Error message displayed
4. **Expected:** Error is helpful (explains what's wrong)
5. **Expected:** Can correct and retry
6. **Expected:** No crash

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 7.2: GitHub API Failure

**Steps:**
1. Run new project wizard
2. Enter invalid GitHub URL: `invalid-url`
3. **Expected:** Error message displayed
4. **Expected:** Suggests checking URL format
5. **Expected:** Can retry or skip
6. **Expected:** No crash

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 7.3: Migration Failure

**Setup:**
```bash
cd /tmp/doplan-manual-test/old
rm -rf .doplan .cursor doplan 2>/dev/null || true

# Create old structure with corrupted config
mkdir -p .cursor/config doplan/01-phase/01-Feature
echo '{"installed": true, "ide": "cursor"' > .cursor/config/doplan-config.json
# Note: Missing closing brace makes it invalid JSON
echo "# Feature 1" > doplan/01-phase/01-Feature/README.md
```

**Steps:**
1. Run: `./doplan` (or `doplan` if in PATH)
2. **Expected:** Migration wizard launches
3. Select "Proceed with migration"
4. **Expected:** Error caught gracefully when parsing config
5. **Expected:** Rollback option offered
6. **Expected:** Backup preserved (if created before error)
7. **Expected:** Clear error message explaining the issue
8. **Expected:** Can retry after fixing config

**Verification:**
```bash
# After error, verify backup exists:
ls -la .doplan/backup/  # Should exist if backup was created
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

### Test 7.4: Dashboard Load Failure

**Setup:**
```bash
cd /tmp/doplan-manual-test/new
# Use a project with DoPlan installed (from Test 2.1)

# Corrupt the state file
mkdir -p .cursor/config
echo '{"phases": [{"id": "phase-1"' > .cursor/config/doplan-state.json
# Note: Incomplete JSON structure
```

**Steps:**
1. Run: `./doplan` (or `doplan` if in PATH)
2. **Expected:** Error caught gracefully when loading state
3. **Expected:** Fallback to markdown dashboard (if `doplan/dashboard.md` exists)
4. **Expected:** Error message displayed explaining state file issue
5. **Expected:** Can retry after fixing state file
6. **Expected:** No crash or panic

**Alternative Test (Missing State File):**
```bash
# Remove state file entirely
rm .cursor/config/doplan-state.json
./doplan  # or `doplan` if in PATH
# Expected: Creates new empty state or shows helpful error
```

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Suite 8: Performance

### Test 8.1: Dashboard Load Time

**Steps:**
1. Use a project with DoPlan installed (from Test 2.1)
2. Measure dashboard generation time:
   ```bash
   time ./doplan dashboard  # or `time doplan dashboard` if in PATH
   ```
   Note: This measures markdown dashboard generation. For TUI dashboard load time, see Test 5.1.
3. **Expected:** Generates in <500ms (markdown generation)
4. **Expected:** No blocking operations
5. **Expected:** Files created successfully

**For TUI Dashboard Load Time:**
1. Run: `./doplan` (or `doplan` if in PATH) - opens TUI dashboard
2. Observe load time visually
3. **Expected:** Loads in <100ms (target for TUI)
4. **Expected:** Smooth rendering, no flickering

**Result:** ☐ PASS ☐ FAIL
**Markdown Generation Time:** _____ ms
**TUI Load Time:** _____ ms (visual estimate)
**Notes:** 

---

### Test 8.2: Large Project Handling

**Setup:**
```bash
# Create a large project structure
cd /tmp/doplan-manual-test/new
# Add many phases and features
```

**Steps:**
1. Run dashboard
2. **Expected:** Dashboard loads in reasonable time (<2s for large projects)
3. **Expected:** Navigation is smooth
4. **Expected:** No memory issues
5. **Expected:** Pagination works if needed

**Result:** ☐ PASS ☐ FAIL
**Notes:** 

---

## Test Results Summary

After completing all tests, fill in the summary:

**Total Tests:** 40
**Passed:** _____
**Failed:** _____
**Skipped:** _____

**Date:** _____
**Tester:** _____

**Overall Status:** ☐ READY FOR RELEASE ☐ NEEDS FIXES

**Critical Issues Found:**
1. 
2. 
3. 

**Notes:**
_____
_____
_____

