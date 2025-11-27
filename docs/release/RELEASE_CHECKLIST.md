# Release Checklist for v1.0.0

## Pre-Release Verification

### Code Quality
- [x] All tests passing (200+ tests)
- [x] Code coverage > 80%
- [x] No linter errors
- [x] Security audit passed
- [x] Documentation reviewed

### Build & Distribution
- [x] Build scripts tested (Unix + Windows)
- [x] Multi-platform builds working
- [x] Version system tested
- [x] Binary size verified (< 15MB)
- [x] Generation time verified (< 5 seconds)

### Documentation
- [x] README.md complete and accurate
- [x] CHANGELOG.md updated with v1.0.0
- [x] BUILD.md created
- [x] TESTING.md created
- [x] SECURITY_AUDIT.md created
- [x] RELEASE_NOTES_v1.0.0.md created

### GitHub Release
- [x] Release notes prepared
- [x] CHANGELOG.md updated
- [ ] **Tag v1.0.0** (to be done when ready)
- [ ] **Push tag to trigger release workflow**
- [ ] **Verify GitHub Release created**
- [ ] **Verify binaries uploaded**

## Release Steps

### 1. Final Verification
```bash
# Run full pre-release checks (tests, coverage, lint, vet)
make pre-release

# Or use the alias
make release-check

# Build and test
make build
./doplan --version

# Verify binary size
ls -lh doplan
```

**Note**: The `pre-release` command will:
- Format code (`make fmt`)
- Run go vet (`make vet`)
- Run linter (`make lint`)
- Run tests with coverage check (`make test-coverage-check`)
  - Generates `coverage.out` and `coverage.html`
  - Validates coverage meets minimum threshold (default: 80%)
  - Fails if coverage is below threshold

**Custom Coverage Threshold**:
```bash
# Use a custom coverage threshold (e.g., 85%)
make pre-release COVERAGE_THRESHOLD=85
```

### 2. Update Version
- [x] Version set to 1.0.0 in release notes
- [x] CHANGELOG.md updated
- [x] Version system ready

### 3. Create Git Tag
```bash
# Create annotated tag
git tag -a v1.0.0 -m "Release v1.0.0: Initial release of DoPlan CLI"

# Push tag (this will trigger GitHub Actions release workflow)
git push origin v1.0.0
```

### 4. Verify Release
- [ ] Check GitHub Actions workflow completed successfully
- [ ] Verify GitHub Release created
- [ ] Verify binaries uploaded for all platforms
- [ ] Verify checksums included
- [ ] Test binary download from release page

### 5. Post-Release
- [ ] Announce release on social media
- [ ] Update website/docs if applicable
- [ ] Monitor for issues
- [ ] Plan v1.1.0 features

## Release Notes Summary

**Version**: 1.0.0  
**Release Date**: November 23, 2024  
**Status**: Initial Release

**Key Features**:
- Interactive TUI wizard
- 18 hierarchical AI agents
- 19 commands (11 core + 8 squad)
- 500+ rules library
- Complete project generation
- Cross-platform support
- GitHub automation
- Security hardened

**Installation**:
```bash
npx @doplan-dev/cli
```

**Documentation**: See RELEASE_NOTES_v1.0.0.md for full details.

---

**Note**: This checklist should be completed before creating the git tag.

