# Production Readiness Checklist

This guide explains how to verify that the DoPlan CLI is ready for a production release.

## Quick Start

Run the production readiness check:

```bash
# Using Make
make check-production

# Or directly
./scripts/check-production-ready.sh
```

## What Gets Checked

The production readiness check verifies:

### 1. Dependencies
- ✅ Go is installed and accessible
- ✅ Git is installed and accessible

### 2. Code Quality
- ✅ **Code Formatting**: All code is properly formatted (`go fmt`)
- ✅ **Static Analysis**: No issues found by `go vet`
- ✅ **Build Test**: Project builds successfully
- ✅ **Tests**: All tests pass
- ✅ **Test Coverage**: Coverage percentage reported

### 3. Code Health
- ⚠️ **TODO/FIXME Comments**: Warns if found (should be addressed before release)
- ⚠️ **Debug Code**: Checks for debug statements that might have been left behind

### 4. Version Control
- ✅ **Git Status**: Working directory is clean (or warns about changes)
- ✅ **Version Info**: Version information is present in code

### 5. Documentation
- ✅ **CHANGELOG**: Exists and has entries
- ✅ **README**: Main documentation exists
- ✅ **CONTRIBUTING**: Contributing guide exists (warns if missing)
- ✅ **Docs Directory**: Documentation is organized

### 6. Legal & Compliance
- ✅ **LICENSE**: License file exists

### 7. Build Configuration
- ✅ **Goreleaser**: Release configuration present
- ✅ **Makefile**: Build commands available

### 8. Test Infrastructure
- ✅ **Test Directory**: Test files exist
- ✅ **Test Scripts**: Test automation scripts available

### 9. Dependencies
- ✅ **Go Modules**: Dependencies are verified
- ✅ **Go Modules Tidy**: No unused dependencies

## Exit Codes

- **0**: Ready for production (may have warnings)
- **1**: Not ready - critical issues found

## Making It Pass

### Critical Issues (Must Fix)

1. **Build Failures**: Fix compilation errors
2. **Test Failures**: Fix failing tests
3. **Missing LICENSE**: Add license file
4. **Missing README**: Add main documentation
5. **Dependency Issues**: Run `go mod tidy` and verify modules

### Warnings (Should Review)

1. **TODO/FIXME Comments**: Review and either address or move to issue tracker
2. **Debug Statements**: Ensure no debug code is left in production
3. **Low Test Coverage**: Consider adding more tests (aim for >60%)
4. **Uncommitted Changes**: Commit or stash before release
5. **Missing CHANGELOG Entry**: Add release notes

## Example Output

```
╔════════════════════════════════════════════════════════════════╗
║        DoPlan Production Readiness Check                       ║
╚════════════════════════════════════════════════════════════════╝

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Checking Dependencies
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
✓ Go is installed (version: 1.25.4)
✓ Git is installed

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
2. Checking Code Formatting
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
✓ All code is properly formatted

... (more checks) ...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Summary
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Passed:  35
Warnings: 2
Failed:  0

╔════════════════════════════════════════════════════════════════╗
║  ⚠ Ready with warnings. Review warnings before release.      ║
╚════════════════════════════════════════════════════════════════╝
```

## Integration with CI/CD

You can integrate this check into your CI/CD pipeline:

```yaml
# .github/workflows/release.yml
- name: Check Production Readiness
  run: make check-production
```

## Manual Checklist

Before a production release, also manually verify:

- [ ] Version number is updated in code
- [ ] CHANGELOG.md has entry for this version
- [ ] All features are documented
- [ ] Breaking changes are clearly marked
- [ ] Release notes are prepared
- [ ] Build artifacts are tested
- [ ] Installation instructions are tested
- [ ] Performance benchmarks are acceptable

## Related Documentation

- [Testing Guide](./TESTING.md)
- [Release Process](../releases/RELEASE.md)
- [Troubleshooting Guide](./TROUBLESHOOTING.md)

