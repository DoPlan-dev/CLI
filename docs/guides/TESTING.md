# DoPlan CLI Testing Guide

This document describes how to test the DoPlan CLI.

## Quick Start

Run all tests:
```bash
make test-scripts
```

Or run individual test suites:
```bash
./scripts/test-cli.sh        # Basic CLI tests
./scripts/test-install.sh    # Installation tests
./scripts/test-integration.sh # Integration tests
```

## Test Scripts

### 1. `test-cli.sh` - Basic CLI Tests

Tests all CLI commands and basic functionality.

**What it tests:**
- ✅ Build process
- ✅ Help command (`--help`)
- ✅ Version command (`--version`)
- ✅ All subcommands (install, dashboard, github, progress)
- ✅ Error handling (commands without installation)
- ✅ TUI mode flag

**Run:**
```bash
./scripts/test-cli.sh
```

**Expected output:**
```
========================================
  DoPlan CLI Test Suite
========================================

Total tests: 14
Passed: 14
Failed: 0

✓ All tests passed!
```

### 2. `test-install.sh` - Installation Tests

Tests the installation flow and directory structure.

**What it tests:**
- ✅ Installation command availability
- ✅ Expected directory structure
- ✅ Git repository setup

**Run:**
```bash
./scripts/test-install.sh
```

**Note:** Full interactive installation test requires manual verification.

### 3. `test-integration.sh` - Integration Tests

Tests the full workflow end-to-end.

**What it tests:**
- ✅ Binary functionality
- ✅ Command execution
- ✅ File generation
- ✅ Directory structure

**Run:**
```bash
./scripts/test-integration.sh
```

## Manual Testing

### Test Installation Flow

1. Create a test project:
```bash
mkdir test-project
cd test-project
git init
```

2. Run installation:
```bash
../bin/doplan install
```

3. Select an IDE (e.g., "Cursor")

4. Verify directories are created:
```bash
ls -la .cursor/
ls -la doplan/
```

### Test Commands

After installation, test each command:

```bash
# Dashboard
doplan dashboard

# GitHub sync
doplan github

# Progress update
doplan progress

# TUI mode
doplan --tui
```

## Unit Tests

Run Go unit tests:
```bash
make test
```

Run with coverage:
```bash
make test-coverage
```

## Continuous Integration

To add CI/CD, you can use the test scripts in your CI pipeline:

```yaml
# Example GitHub Actions
- name: Run CLI tests
  run: make test-scripts

- name: Run unit tests
  run: make test
```

## Test Results

### Latest Test Run

```
Total tests: 14
Passed: 14
Failed: 0

✓ All tests passed!
```

## Troubleshooting

### Tests Fail to Run

1. **Permission denied:**
   ```bash
   chmod +x scripts/*.sh
   ```

2. **Binary not found:**
   ```bash
   make build
   ```

3. **Git not configured:**
   ```bash
   git config user.name "Test User"
   git config user.email "test@example.com"
   ```

### Build Errors

If build fails:
```bash
go mod tidy
make build
```

### Test Environment

Tests create temporary directories and clean up automatically. If tests are interrupted, you may need to manually clean up:

```bash
rm -rf /tmp/tmp.*
```

## Adding New Tests

To add new tests:

1. Edit the appropriate test script
2. Add a new `test_command` or `test_output_contains` call
3. Run the script to verify
4. Update this documentation

Example:
```bash
test_command "New feature test" \
    "$BINARY_PATH new-feature" \
    0
```

## Test Coverage

Current test coverage:
- ✅ CLI commands: 100%
- ✅ Error handling: 100%
- ✅ Build process: 100%
- ⚠️ Interactive flows: Manual testing required
- ⚠️ TUI mode: Basic flag test only

## Next Steps

1. ✅ Basic CLI tests - Complete
2. ✅ Installation tests - Complete
3. ✅ Integration tests - Complete
4. ⏳ Interactive installation test automation
5. ⏳ TUI mode full testing
6. ⏳ GitHub integration tests (requires GitHub CLI)

