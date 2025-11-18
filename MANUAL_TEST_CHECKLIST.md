# Manual Test Checklist - DoPlan v0.0.18-beta

Use this checklist to track your progress through manual testing.

## Quick Start

1. **Build the binary:**
   ```bash
   go build -o doplan ./cmd/doplan/main.go
   ```

2. **Set up test environment:**
   ```bash
   mkdir -p /tmp/doplan-manual-test/{empty,existing,old,new}
   ```

3. **Follow the detailed guide:**
   - See `docs/development/MANUAL_TESTING_GUIDE.md` for step-by-step instructions

4. **Verify results:**
   ```bash
   ./scripts/verify-test-results.sh
   ```

---

## Test Suite 2: New Project Wizard

- [ ] **Test 2.1:** Complete Wizard Flow
  - [ ] Welcome screen appears
  - [ ] Project name input works
  - [ ] Template selection works
  - [ ] GitHub setup completes
  - [ ] IDE selection works
  - [ ] Installation completes
  - [ ] Dashboard opens automatically

- [ ] **Test 2.2:** Project Name Validation
  - [ ] `test-project` → Valid
  - [ ] `Test Project` → Invalid
  - [ ] `test_project` → Invalid
  - [ ] `test` → Invalid (too short)
  - [ ] `test-project-123` → Valid

- [ ] **Test 2.3:** Template Selection
  - [ ] All templates selectable
  - [ ] Preview updates
  - [ ] Descriptions accurate
  - [ ] Selection works

- [ ] **Test 2.4:** GitHub Repository Validation
  - [ ] `https://github.com/user/repo` → Valid
  - [ ] `git@github.com:user/repo.git` → Valid
  - [ ] `user/repo` → Valid
  - [ ] `invalid-url` → Invalid
  - [ ] `skip` → Valid

- [ ] **Test 2.5:** IDE Selection
  - [ ] All IDE options available
  - [ ] Selection works
  - [ ] Integration files created

---

## Test Suite 3: Project Adoption Wizard

- [ ] **Test 3.1:** Project Analysis
  - [ ] Detects tech stack
  - [ ] Finds project files
  - [ ] Extracts documentation
  - [ ] Identifies potential features
  - [ ] Shows analysis results

- [ ] **Test 3.2:** Adoption Options
  - [ ] Three options available
  - [ ] Each option works
  - [ ] Can navigate back

- [ ] **Test 3.3:** Auto-Plan Generation
  - [ ] Plan generated from code
  - [ ] Phases created correctly
  - [ ] Features identified
  - [ ] Tasks extracted
  - [ ] Plan is accurate

---

## Test Suite 4: Migration Wizard

- [ ] **Test 4.1:** Migration Detection
  - [ ] Detects old structure
  - [ ] Shows migration screen
  - [ ] Explains what will be migrated
  - [ ] Offers to proceed or skip

- [ ] **Test 4.2:** Backup Creation
  - [ ] Backup created in `.doplan/backup/TIMESTAMP/`
  - [ ] All files backed up
  - [ ] Backup is complete

- [ ] **Test 4.3:** Config Migration
  - [ ] `doplan-config.json` → `config.yaml`
  - [ ] All fields migrated correctly
  - [ ] New fields added with defaults
  - [ ] Config is valid YAML
  - [ ] Config validation passes

- [ ] **Test 4.4:** Folder Renaming
  - [ ] Old folders renamed
  - [ ] All files copied correctly
  - [ ] References updated
  - [ ] Old folders removed (after verification)

- [ ] **Test 4.5:** Migration Rollback
  - [ ] Rollback option available
  - [ ] Old structure restored
  - [ ] New structure removed
  - [ ] All files restored correctly

---

## Test Suite 5: Dashboard TUI

- [ ] **Test 5.1:** Dashboard Loading
  - [ ] Loads in <100ms (target)
  - [ ] Shows project name
  - [ ] Shows progress bars
  - [ ] Shows phase list
  - [ ] Shows feature list
  - [ ] Shows GitHub activity
  - [ ] No errors

- [ ] **Test 5.2:** Dashboard Navigation
  - [ ] Press `1` → Dashboard view
  - [ ] Press `2` → Phases view
  - [ ] Press `3` → Features view
  - [ ] Press `4` → GitHub view
  - [ ] Press `5` → Config view
  - [ ] Press `6` → Stats view
  - [ ] Navigation is smooth

- [ ] **Test 5.3:** Progress Bar Accuracy
  - [ ] Overall progress is accurate
  - [ ] Phase progress is accurate
  - [ ] Feature progress is accurate
  - [ ] Progress bars render correctly
  - [ ] Colors match status

- [ ] **Test 5.4:** Real-time Updates
  - [ ] Refresh works (press `r`)
  - [ ] Data updates correctly
  - [ ] Progress recalculated
  - [ ] GitHub data refreshed
  - [ ] No flickering

---

## Test Suite 6: IDE Integration

- [ ] **Test 6.1:** Cursor Integration
  - [ ] `.cursor/agents/` → symlink to `.doplan/ai/agents/`
  - [ ] `.cursor/rules/` → symlink to `.doplan/ai/rules/`
  - [ ] `.cursor/commands/` → symlink to `.doplan/ai/commands/`
  - [ ] Symlinks work correctly
  - [ ] Files are accessible

- [ ] **Test 6.2:** VS Code Integration
  - [ ] `.vscode/tasks.json` created
  - [ ] `.vscode/settings.json` created
  - [ ] `.vscode/prompts/` directory created
  - [ ] Files contain correct content
  - [ ] Tasks are executable

- [ ] **Test 6.3:** Generic IDE Integration
  - [ ] `.doplan/guides/generic_ide_setup.md` created
  - [ ] Guide is comprehensive
  - [ ] Instructions are clear

---

## Test Suite 7: Error Handling

- [ ] **Test 7.1:** Invalid Project Name
  - [ ] Error message displayed
  - [ ] Error is helpful
  - [ ] Can correct and retry
  - [ ] No crash

- [ ] **Test 7.2:** GitHub API Failure
  - [ ] Error message displayed
  - [ ] Suggests checking URL
  - [ ] Can retry or skip
  - [ ] No crash

- [ ] **Test 7.3:** Migration Failure
  - [ ] Error caught gracefully
  - [ ] Rollback option offered
  - [ ] Backup preserved
  - [ ] Clear error message

- [ ] **Test 7.4:** Dashboard Load Failure
  - [ ] Error caught gracefully
  - [ ] Fallback to markdown dashboard
  - [ ] Error message displayed
  - [ ] Can retry

---

## Test Suite 8: Performance

- [ ] **Test 8.1:** Dashboard Load Time
  - [ ] Loads in <100ms (target)
  - [ ] No blocking operations
  - [ ] Smooth rendering
  - [ ] **Measured Time:** _____ ms

- [ ] **Test 8.2:** Large Project Handling
  - [ ] Dashboard loads in reasonable time
  - [ ] Navigation is smooth
  - [ ] No memory issues
  - [ ] Pagination works if needed

---

## Test Summary

**Date Completed:** _____
**Tester:** _____

**Total Tests:** 40
**Passed:** _____
**Failed:** _____
**Skipped:** _____

**Overall Status:**
- [ ] ✅ READY FOR RELEASE
- [ ] ⚠️ NEEDS FIXES
- [ ] ❌ NOT READY

**Critical Issues:**
1. 
2. 
3. 

**Notes:**
_____
_____
_____

