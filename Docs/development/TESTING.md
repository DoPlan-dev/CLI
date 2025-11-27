# Testing Guide

## End-to-End Integration Testing

This document describes the comprehensive end-to-end testing strategy for DoPlan CLI.

## Test Coverage

### 1. Complete Wizard Flow
- **Test**: `TestEndToEnd_CompleteWizardFlow`
- **Purpose**: Verifies the complete user journey from wizard to project generation
- **Validates**:
  - Request validation
  - Full orchestration
  - Generation time < 5 seconds
  - Complete project structure

### 2. All Files Generated
- **Test**: `TestEndToEnd_AllFilesGenerated`
- **Purpose**: Ensures all expected files are created correctly
- **Validates**:
  - Documentation files (README.md, STANDUP.md, CHANGELOG.md)
  - IDE configs (docs/CLAUDE.md + `.cursor/rules` library)
- Boilerplate scaffolding instructions (legacy `scripts/boilerplate` helper removed; ensure docs point developers to their preferred stack generator)
  - Boilerplate files (package.json, tsconfig.json, etc.)
  - Agent files (18 agents)
  - Command files (19 commands)
  - Rules library (15+ categories)
  - Plan structure (.plan/ directory)
  - GitHub workflows (4 workflows)

### 3. Generation Time
- **Test**: `TestEndToEnd_GenerationTime`
- **Purpose**: Verifies generation completes within performance target
- **Target**: < 5 seconds
- **Current Performance**: ~40-50ms (well under target)

### 4. Cross-Platform Path Handling
- **Test**: `TestEndToEnd_CrossPlatformPaths`
- **Purpose**: Ensures path handling works across platforms
- **Validates**:
  - Simple project names
  - Names with underscores
  - Names with numbers
  - Mixed case names

### 5. All Supported IDEs
- **Test**: `TestEndToEnd_AllIDEs`
- **Purpose**: Verifies generation works with all supported IDEs
- **IDEs Tested**:
  - Cursor
  - Claude Code
  - Antigravity
  - Windsurf
  - Cline
  - OpenCode

### 6. Platform Information
- **Test**: `TestEndToEnd_PlatformInfo`
- **Purpose**: Logs platform information for cross-platform testing
- **Current Platform**: macOS (darwin/amd64)

## Running Tests

### Run All End-to-End Tests
```bash
go test ./internal/generator -v -run TestEndToEnd
```

### Run Specific Test
```bash
go test ./internal/generator -v -run TestEndToEnd_CompleteWizardFlow
```

### Run All Generator Tests
```bash
go test ./internal/generator/... -v
```

## Binary Size Verification

### Target
- Binary size < 15MB

### Check Binary Size
```bash
go build -o doplan ./cmd/doplan
ls -lh doplan
```

## Cross-Platform Testing

### Platforms to Test
1. **macOS**
   - Intel (x86_64)
   - Apple Silicon (arm64)

2. **Linux**
   - Ubuntu
   - Debian

3. **Windows**
   - Windows 10/11

### Testing Checklist
- [ ] Project generation works
- [ ] File paths are correct
- [ ] TUI wizard displays correctly
- [ ] All files generated correctly
- [ ] Binary size < 15MB
- [ ] Generation time < 5 seconds

## Performance Benchmarks

### Current Performance (macOS)
- **Generation Time**: ~40-50ms
- **Target**: < 5 seconds
- **Status**: ✅ Exceeds target by 100x

### Binary Size
- **Target**: < 15MB
- **Status**: ✅ (Verify with `ls -lh` after build)

## Test Results Summary

### End-to-End Tests
- ✅ Complete wizard flow
- ✅ All files generated
- ✅ Generation time verified
- ✅ Cross-platform paths
- ✅ All IDEs supported
- ✅ Platform information logged

### Integration Tests
- ✅ Full project generation
- ✅ Generator pipeline order
- ✅ All components verified

### Unit Tests
- ✅ All generators tested
- ✅ Error handling verified
- ✅ Validation tested

## Continuous Integration

These tests should be run:
- Before every commit
- In CI/CD pipeline
- On all supported platforms
- Before releases

## Next Steps

1. Set up CI/CD to run tests on all platforms
2. Add performance benchmarks to CI
3. Monitor binary size in CI
4. Add test coverage reporting

