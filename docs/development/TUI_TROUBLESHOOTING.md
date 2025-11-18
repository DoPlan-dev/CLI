# TUI Troubleshooting Guide

## Why Can't I Open the TUI?

The DoPlan TUI (Terminal User Interface) only opens when:
1. **DoPlan is already installed** in the current project
2. The project state is detected as `StateNewDoPlanStructure`

### Current Behavior by Project State

| Project State | What Opens |
|--------------|------------|
| `StateEmptyFolder` | **New Project Wizard** (not TUI) |
| `StateExistingCodeNoDoPlan` | **Adoption Wizard** (not TUI) |
| `StateOldDoPlanStructure` | **Migration Wizard** (not TUI) |
| `StateNewDoPlanStructure` | **TUI Dashboard** ✅ |
| `StateInsideFeature` | **TUI Dashboard** ✅ |
| `StateInsidePhase` | **TUI Dashboard** ✅ |

## How to Open the TUI

### Option 1: Install DoPlan First (Recommended)

1. **Run the New Project Wizard:**
   ```bash
   cd /tmp/doplan-manual-test/empty
   doplan
   # Follow the wizard to install DoPlan
   ```

2. **After installation completes**, the TUI will open automatically

3. **To open TUI again later:**
   ```bash
   cd /path/to/installed/project
   doplan  # TUI opens automatically
   ```

### Option 2: Navigate to an Installed Project

```bash
# Find a project with DoPlan installed
cd /path/to/project/with/.doplan/config.yaml

# Run doplan
doplan  # TUI opens
```

### Option 3: Create a Test Project with DoPlan

```bash
# Create test project
mkdir -p /tmp/doplan-test-project
cd /tmp/doplan-test-project

# Install DoPlan
doplan
# Complete the wizard

# TUI will open automatically after installation
```

## Common Issues

### Issue 1: "TUI doesn't open in empty folder"

**Solution:** This is expected behavior. Empty folders trigger the New Project Wizard, not the TUI. Complete the wizard first.

### Issue 2: "Terminal doesn't support TUI"

**Symptoms:**
- Error about terminal not supporting TUI
- Blank screen
- Immediate exit

**Check:**
```bash
echo $TERM  # Should show something like: xterm-256color, screen-256color, etc.
[ -t 0 ] && echo "Interactive" || echo "Not interactive"
```

**Solutions:**
- Use a proper terminal emulator (iTerm2, Terminal.app, Alacritty, etc.)
- Ensure `TERM` environment variable is set
- Don't pipe output: `doplan | cat` won't work

### Issue 3: "Wizard opens instead of TUI"

**Cause:** Project state detection routes to wizard instead of TUI.

**Check project state:**
```bash
# Check if DoPlan is installed
ls -la .doplan/config.yaml  # Should exist for TUI to open
ls -la .cursor/config/doplan-config.json  # Old structure (triggers migration)
```

**Solution:** Complete the wizard/migration first, then TUI will open.

### Issue 4: "TUI crashes or exits immediately"

**Possible causes:**
1. Missing state file
2. Corrupted config
3. Terminal compatibility issue

**Debug:**
```bash
# Run with error output
doplan 2>&1 | tee doplan-error.log

# Check for errors
cat doplan-error.log
```

**Fix:**
```bash
# Reinstall DoPlan
rm -rf .doplan .cursor/config/doplan-*
doplan  # Run wizard again
```

### Issue 5: "TUI shows blank screen"

**Possible causes:**
1. Terminal doesn't support colors/ANSI
2. Window size detection failed
3. Loading error

**Check:**
```bash
# Test terminal capabilities
tput colors  # Should show 256 or more
tput cols    # Should show terminal width
tput lines   # Should show terminal height
```

**Solution:** Use a modern terminal emulator with full color support.

## Testing TUI Launch

### Quick Test Script

```bash
#!/bin/bash
# test-tui.sh

echo "Testing TUI launch conditions..."

# Check terminal
echo "TERM: $TERM"
echo "Interactive: $([ -t 0 ] && echo yes || echo no)"

# Check if DoPlan is installed
if [ -f ".doplan/config.yaml" ]; then
    echo "✅ DoPlan installed - TUI should open"
    doplan
elif [ -f ".cursor/config/doplan-config.json" ]; then
    echo "⚠️  Old DoPlan structure - Migration wizard will open"
    doplan
else
    echo "❌ DoPlan not installed - New Project Wizard will open"
    echo "Run 'doplan' to install, then TUI will open automatically"
fi
```

## Expected Behavior

### Empty Folder
```bash
cd /tmp/empty-folder
doplan
# → Opens New Project Wizard (not TUI)
```

### After Installation
```bash
cd /path/to/installed/project
doplan
# → Opens TUI Dashboard ✅
```

### Inside Feature Directory
```bash
cd /path/to/project/doplan/01-phase-1/01-feature-1
doplan
# → Opens TUI Dashboard ✅
```

## Manual TUI Launch (Advanced)

If you need to force TUI launch for testing:

```bash
# Create minimal DoPlan structure
mkdir -p .doplan
cat > .doplan/config.yaml <<EOF
project:
  name: test
  ide: cursor
EOF

# Now TUI should open
doplan
```

## Getting Help

If TUI still doesn't open:

1. **Check error messages:**
   ```bash
   doplan 2>&1
   ```

2. **Verify installation:**
   ```bash
   ls -la .doplan/config.yaml
   ```

3. **Test terminal:**
   ```bash
   echo $TERM
   [ -t 0 ] && echo "OK" || echo "Not interactive"
   ```

4. **Check Go version:**
   ```bash
   go version  # Should be 1.21+
   ```

5. **Rebuild binary:**
   ```bash
   go build -o doplan ./cmd/doplan/main.go
   ```

