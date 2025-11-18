# DoPlan v0.0.18-beta Troubleshooting Guide

Comprehensive troubleshooting guide for common issues during development and implementation.

## 🔍 Diagnostic Tools

### Quick Diagnostics

```bash
# Run full diagnostic
./scripts/check-dependencies.sh
./scripts/validate-config.sh
./scripts/verify-ide-integration.sh
./scripts/analyze-folder-structure.sh
```

### Debug Mode

```bash
# Enable debug logging
export DOPLAN_DEBUG=1
doplan

# Check debug logs
tail -f .doplan/logs/doplan.log
```

## 🐛 Common Issues

### Issue 1: Test Environment Not Found

**Symptoms:**
```
❌ Directory not found: /tmp/doplan-test/empty
```

**Solutions:**
```bash
# Set up test environment
make setup
# or
./scripts/setup-test-env.sh

# Verify setup
ls -la /tmp/doplan-test/
```

**Prevention:**
- Always run `make setup` before testing
- Use `./scripts/backup-test-projects.sh` before major changes

### Issue 2: Migration Fails

**Symptoms:**
```
❌ Migration failed: Config migration failed
```

**Diagnosis:**
```bash
# Check old config
cat .cursor/config/doplan-config.json | jq .

# Validate JSON
jq '.' .cursor/config/doplan-config.json

# Check permissions
ls -la .cursor/config/
```

**Solutions:**

1. **Corrupted old config:**
   ```bash
   # Restore from backup
   ./scripts/restore-test-projects.sh <timestamp>
   ```

2. **Permission issues:**
   ```bash
   # Check permissions
   ls -la .doplan/
   
   # Fix permissions
   chmod -R 755 .doplan/
   ```

3. **Disk space:**
   ```bash
   # Check disk space
   df -h .
   ```

**Prevention:**
- Always create backup before migration
- Validate old config before migration
- Check disk space before migration

### Issue 3: IDE Integration Not Working

**Symptoms:**
```
⚠️ Cursor integration incomplete
```

**Diagnosis:**
```bash
# Verify integration
./scripts/verify-ide-integration.sh

# Check symlinks
ls -la .cursor/

# Check target exists
readlink -f .cursor/agents
```

**Solutions:**

1. **Symlinks broken:**
   ```bash
   # Recreate symlinks
   doplan setup-ide cursor
   ```

2. **Target doesn't exist:**
   ```bash
   # Check .doplan/ai/ exists
   ls -la .doplan/ai/
   
   # Recreate if missing
   mkdir -p .doplan/ai/{agents,rules,commands}
   ```

3. **Windows symlink issues:**
   - On Windows, use file copy instead
   - Edit integration code to use copyDir instead of symlink

**Prevention:**
- Verify integration after setup
- Test on target platform
- Handle Windows separately

### Issue 4: Dashboard Not Loading

**Symptoms:**
```
❌ Dashboard load failed
```

**Diagnosis:**
```bash
# Check dashboard.json
cat .doplan/dashboard.json | jq .

# Validate JSON
jq '.' .doplan/dashboard.json

# Check file permissions
ls -la .doplan/dashboard.json
```

**Solutions:**

1. **Invalid JSON:**
   ```bash
   # Regenerate dashboard
   doplan dashboard --json
   ```

2. **Missing data:**
   ```bash
   # Check progress files
   find doplan -name "progress.json"
   
   # Regenerate from progress files
   doplan dashboard --regenerate
   ```

3. **Performance issues:**
   ```bash
   # Check project size
   find doplan -type f | wc -l
   ```

**Prevention:**
- Validate JSON after generation
- Test with various project sizes
- Add performance monitoring

### Issue 5: TUI Not Rendering

**Symptoms:**
```
Terminal too small
TUI not displaying correctly
```

**Diagnosis:**
```bash
# Check terminal size
echo $COLUMNS x $LINES

# Test TUI
./scripts/test-tui-interactive.sh
```

**Solutions:**

1. **Terminal too small:**
   - Resize terminal
   - Minimum: 80x24
   - Recommended: 120x30

2. **Terminal compatibility:**
   ```bash
   # Check terminal type
   echo $TERM
   
   # Try different terminal
   # Use xterm-256color or similar
   ```

3. **Encoding issues:**
   ```bash
   # Set UTF-8
   export LANG=en_US.UTF-8
   export LC_ALL=en_US.UTF-8
   ```

**Prevention:**
- Check terminal size before rendering
- Handle small terminals gracefully
- Test on multiple terminals

### Issue 6: Tests Failing

**Symptoms:**
```
FAIL: TestMigration
FAIL: TestWizard
```

**Diagnosis:**
```bash
# Run specific test
go test -v ./internal/migration/... -run TestMigration

# Run with race detector
go test -race ./internal/...

# Check test coverage
go test -cover ./internal/...
```

**Solutions:**

1. **Test data issues:**
   ```bash
   # Regenerate test data
   ./scripts/generate-test-data.sh
   
   # Clean and reset
   make clean
   make setup
   ```

2. **Race conditions:**
   ```bash
   # Run with race detector
   go test -race ./...
   
   # Fix race conditions
   # Use mutexes or channels
   ```

3. **Timing issues:**
   - Add timeouts to async operations
   - Use context with timeout

**Prevention:**
- Use table-driven tests
- Mock external dependencies
- Add timeouts to async operations

## 🔧 Advanced Troubleshooting

### Debugging TUI Issues

```go
// Add debug logging
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if os.Getenv("DOPLAN_DEBUG") == "1" {
        log.Printf("Received message: %T %+v", msg, msg)
    }
    // ...
}

// Dump state
func (m *Model) dumpState() {
    stateJSON, _ := json.MarshalIndent(m, "", "  ")
    os.WriteFile("debug-state.json", stateJSON, 0644)
}
```

### Debugging Migration Issues

```bash
# Enable verbose migration
DOPLAN_VERBOSE=1 doplan migrate

# Check migration log
cat .doplan/migration.log

# Validate step by step
./scripts/validate-migration.sh --verbose
```

### Debugging Performance Issues

```bash
# Profile CPU
go test -cpuprofile=cpu.prof ./...
go tool pprof cpu.prof

# Profile memory
go test -memprofile=mem.prof ./...
go tool pprof mem.prof

# Benchmark
go test -bench=. ./...
```

## 📊 Error Code Reference

### Quick Lookup

```bash
# Search error codes
grep -r "USR1001" internal/
grep -r "SYS2001" internal/

# Find error definitions
grep -r "ErrInvalidProjectName" internal/
```

### Error Code Meanings

| Code | Category | Common Causes | Solutions |
|------|----------|---------------|-----------|
| USR1001 | User Input | Invalid project name | Check format requirements |
| USR1003 | User Input | Invalid GitHub URL | Verify URL format |
| SYS2001 | System | Permission denied | Check file permissions |
| SYS2002 | System | Disk full | Free up disk space |
| CFG3001 | Config | Config not found | Run initialization |
| CFG3002 | Config | Invalid config | Fix YAML syntax |
| MIG4002 | Migration | Backup failed | Check disk space |
| MIG4004 | Migration | Folder migration failed | Check permissions |
| INT5002 | Integration | IDE setup failed | Re-run setup |
| TUI6001 | TUI | Terminal too small | Resize terminal |

## 🆘 Getting Help

### Before Asking for Help

1. **Collect information:**
   ```bash
   # System info
   uname -a
   go version
   doplan --version
   
   # Error logs
   cat .doplan/logs/doplan.log
   
   # Config
   cat .doplan/config.yaml
   ```

2. **Reproduce issue:**
   - Minimal reproduction
   - Document steps
   - Note expected vs actual behavior

3. **Check documentation:**
   - Error Handling Guide
   - Testing Scenarios
   - Migration Strategy

### Reporting Issues

Include:
- Error code (if available)
- Error message
- Steps to reproduce
- System information
- Relevant logs
- Expected behavior

## 🔄 Recovery Procedures

### Recovery 1: Rollback Migration

```bash
# Find backup
ls -la .doplan/backup/

# Restore
./scripts/restore-test-projects.sh <timestamp>

# Or manual
cp -r .doplan/backup/<timestamp>/.cursor .cursor
cp -r .doplan/backup/<timestamp>/doplan doplan
rm -rf .doplan
```

### Recovery 2: Reset Configuration

```bash
# Backup current config
cp .doplan/config.yaml .doplan/config.yaml.backup

# Regenerate
doplan init

# Or restore
cp .doplan/config.yaml.backup .doplan/config.yaml
```

### Recovery 3: Rebuild Test Environment

```bash
# Clean everything
make clean

# Rebuild
make setup

# Verify
./scripts/check-dependencies.sh
```

## 📝 Prevention Checklist

Before starting work:
- [ ] Dependencies installed (`make deps`)
- [ ] Test environment set up (`make setup`)
- [ ] Backups created (`./scripts/backup-test-projects.sh`)
- [ ] Documentation reviewed

During development:
- [ ] Tests passing (`make test`)
- [ ] No race conditions (`go test -race`)
- [ ] Performance acceptable (`make benchmark`)
- [ ] Error handling tested

Before committing:
- [ ] All tests pass
- [ ] Code formatted (`go fmt`)
- [ ] Documentation updated
- [ ] Migration validated (if applicable)

