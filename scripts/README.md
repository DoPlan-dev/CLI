# DoPlan Test Scripts

This directory contains test scripts to verify DoPlan CLI functionality.

## Available Scripts

### `test-cli.sh`
Comprehensive test suite for all CLI commands.

**Usage:**
```bash
./scripts/test-cli.sh
```

**Tests:**
- Build verification
- Help and version commands
- All subcommands (install, dashboard, github, progress)
- Error handling (commands without installation)
- TUI mode flag

### `test-install.sh`
Tests the installation flow in a real project environment.

**Usage:**
```bash
./scripts/test-install.sh
```

**Tests:**
- Installation command availability
- Expected directory structure
- Git repository setup

### `test-integration.sh`
Integration tests for the full workflow.

**Usage:**
```bash
./scripts/test-integration.sh
```

**Tests:**
- Binary functionality
- Command execution
- File generation
- Directory structure

## Running All Tests

Run all test scripts:

```bash
cd cli
./scripts/test-cli.sh
./scripts/test-install.sh
./scripts/test-integration.sh
```

Or use Make:

```bash
make test  # If you add this to Makefile
```

## Manual Testing

Some tests require manual interaction:

1. **Installation Test:**
   ```bash
   mkdir test-project
   cd test-project
   git init
   ../bin/doplan install
   # Select an IDE option
   # Verify directories are created
   ```

2. **Dashboard Test:**
   ```bash
   doplan dashboard
   # Should show dashboard or "not installed" message
   ```

3. **TUI Test:**
   ```bash
   doplan --tui
   # Should open fullscreen TUI
   ```

## Expected Results

All automated tests should pass. Manual tests require:
- Git repository initialized
- GitHub CLI installed (for PR tests)
- Terminal with color support

## Troubleshooting

If tests fail:

1. **Build errors:** Run `make build` manually
2. **Permission errors:** Ensure scripts are executable (`chmod +x scripts/*.sh`)
3. **Binary not found:** Build the binary first (`make build`)
4. **Git errors:** Ensure git is installed and configured

