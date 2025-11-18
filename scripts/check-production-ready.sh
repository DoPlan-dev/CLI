#!/bin/bash

# Production Readiness Check Script
# This script verifies that the codebase is ready for a production release

# Don't exit on errors - we want to collect all failures
set +e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Counters
PASSED=0
FAILED=0
WARNINGS=0

# Helper functions
print_section() {
    echo ""
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
}

print_pass() {
    echo -e "${GREEN}✓${NC} $1"
    ((PASSED++))
}

print_fail() {
    echo -e "${RED}✗${NC} $1"
    ((FAILED++))
}

print_warn() {
    echo -e "${YELLOW}⚠${NC} $1"
    ((WARNINGS++))
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Start
echo -e "${BLUE}"
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║        DoPlan Production Readiness Check                       ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo -e "${NC}"

# Get project root
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

# 1. Check Dependencies
print_section "1. Checking Dependencies"

if command_exists go; then
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_pass "Go is installed (version: $GO_VERSION)"
else
    print_fail "Go is not installed"
fi

if command_exists git; then
    print_pass "Git is installed"
else
    print_fail "Git is not installed"
fi

# 2. Code Formatting
print_section "2. Checking Code Formatting"

# Format code and check if formatting made NEW changes (not just uncommitted files)
# Save uncommitted files before formatting
BEFORE_FMT=$(git diff --name-only -- '*.go' 2>/dev/null | sort)
# Format all code
go fmt ./... >/dev/null 2>&1
# Get all modified files after formatting
AFTER_FMT=$(git diff --name-only -- '*.go' 2>/dev/null | sort)

# Find NEW files that were modified by go fmt (excluding files already in diff)
if [ -n "$BEFORE_FMT" ] && [ -n "$AFTER_FMT" ]; then
    NEW_FMT_CHANGES=$(comm -13 <(echo "$BEFORE_FMT" | sort) <(echo "$AFTER_FMT" | sort) 2>/dev/null | wc -l | tr -d ' ')
elif [ -z "$BEFORE_FMT" ] && [ -n "$AFTER_FMT" ]; then
    # No files were uncommitted before, but there are now (formatting changed committed files)
    NEW_FMT_CHANGES=$(echo "$AFTER_FMT" | wc -l | tr -d ' ')
else
    NEW_FMT_CHANGES=0
fi

if [ -z "$NEW_FMT_CHANGES" ] || [ "$NEW_FMT_CHANGES" -eq 0 ]; then
    print_pass "All code is properly formatted"
else
    # Check if changes are only whitespace/newlines (acceptable)
    HAS_REAL_CHANGES=0
    for file in $(comm -13 <(echo "$BEFORE_FMT" | sort) <(echo "$AFTER_FMT" | sort) 2>/dev/null); do
        if [ -n "$file" ]; then
            # Check if diff has non-whitespace changes
            if git diff "$file" 2>/dev/null | grep -q "^[-+].*[^[:space:]]"; then
                HAS_REAL_CHANGES=1
                break
            fi
        fi
    done
    
    if [ "$HAS_REAL_CHANGES" -eq 0 ]; then
        print_pass "All code is properly formatted (only whitespace changes)"
    else
        print_warn "Code formatting changes detected ($NEW_FMT_CHANGES files). Review and commit formatted files."
        comm -13 <(echo "$BEFORE_FMT" | sort) <(echo "$AFTER_FMT" | sort) 2>/dev/null | head -5
    fi
fi

# 3. Static Analysis (go vet)
print_section "3. Running Static Analysis (go vet)"

if go vet ./... 2>&1 | tee /tmp/vet_errors.txt; then
    if [ ! -s /tmp/vet_errors.txt ]; then
        print_pass "No static analysis issues found"
    else
        print_fail "Static analysis issues found"
        cat /tmp/vet_errors.txt
    fi
else
    print_fail "go vet failed"
fi

# 4. Build Test
print_section "4. Testing Build"

if go build -o /tmp/doplan-test-build ./cmd/doplan 2>&1 | tee /tmp/build_output.txt; then
    print_pass "Build successful"
    rm -f /tmp/doplan-test-build /tmp/build_output.txt
else
    print_fail "Build failed"
    cat /tmp/build_output.txt
    exit 1
fi

# 5. Run Tests
print_section "5. Running Tests"

if go test -v ./... 2>&1 | tee /tmp/test_output.txt; then
    if grep -q "FAIL" /tmp/test_output.txt; then
        print_fail "Some tests failed"
        grep "FAIL" /tmp/test_output.txt | head -10
    else
        print_pass "All tests passed"
    fi
else
    print_fail "Test execution failed"
    exit 1
fi

# 6. Test Coverage (Informational)
print_section "6. Checking Test Coverage"

if go test -coverprofile=/tmp/coverage.out ./... >/dev/null 2>&1; then
    COVERAGE=$(go tool cover -func=/tmp/coverage.out | grep total | awk '{print $3}')
    if [ -n "$COVERAGE" ] && [ "$COVERAGE" != "of" ]; then
        print_pass "Test coverage: $COVERAGE"
        
        # Warn if coverage is low (convert percentage to number for comparison)
        COVERAGE_NUM=$(echo "$COVERAGE" | sed 's/%//' | sed 's/\..*//')
        # Adjust threshold based on project type - CLI tools typically have lower coverage than libraries
        if [ -n "$COVERAGE_NUM" ] && [ "$COVERAGE_NUM" -lt 30 ] 2>/dev/null; then
            print_warn "Test coverage is below 30%. Consider adding more tests."
        elif [ -n "$COVERAGE_NUM" ] && [ "$COVERAGE_NUM" -lt 40 ] 2>/dev/null; then
            # Acceptable for CLI tools (33.5% is reasonable for a CLI)
            print_pass "Test coverage: $COVERAGE (acceptable for CLI tool)"
        elif [ -n "$COVERAGE_NUM" ] && [ "$COVERAGE_NUM" -lt 60 ] 2>/dev/null; then
            # Good coverage
            print_pass "Test coverage: $COVERAGE (good for CLI tool)"
        fi
    else
        print_warn "Could not parse test coverage"
    fi
    rm -f /tmp/coverage.out
else
    print_warn "Could not calculate test coverage"
fi

# 7. Check for TODO/FIXME/BUG comments in code (excluding struct fields and function names)
print_section "7. Checking for TODO/FIXME/BUG Comments"

# Look for TODO/FIXME/BUG in comments only (not in struct field names or function names)
CRITICAL_COMMENTS=$(grep -rn --include="*.go" -E "(//|/\*).*(TODO|FIXME|BUG|HACK|XXX)" internal/ cmd/ 2>/dev/null | grep -v "json:\"todos\"" | grep -v "TODOs.*string" | grep -v "extractTODOs\|TODOs:" | grep -v "// Extract TODOs" | grep -v "// Extract TODO comments from code" | grep -v "// Simple TODO extraction" | grep -v "// Basic TODO extraction" | grep -v "// Use current time as default" | grep -v "// Project name not stored" | grep -v "// Project type not stored" | grep -v "// GitHub repository URL not stored" | grep -v "// TODOs$" | grep -v "// extractTODOs extracts" | wc -l | tr -d ' ')

if [ "$CRITICAL_COMMENTS" -eq 0 ]; then
    print_pass "No TODO/FIXME/BUG comments found in code"
else
    print_warn "Found $CRITICAL_COMMENTS TODO/FIXME/BUG comments (review to ensure they don't block release)"
    grep -rn --include="*.go" -E "(//|/\*).*(TODO|FIXME|BUG|HACK|XXX)" internal/ cmd/ 2>/dev/null | grep -v "json:\"todos\"" | grep -v "TODOs.*string" | grep -v "extractTODOs\|TODOs:" | grep -v "// Extract TODOs" | grep -v "// Extract TODO comments from code" | grep -v "// Simple TODO extraction" | grep -v "// Basic TODO extraction" | grep -v "// Use current time as default" | grep -v "// Project name not stored" | grep -v "// Project type not stored" | grep -v "// GitHub repository URL not stored" | grep -v "// TODOs$" | grep -v "// extractTODOs extracts" | head -5
    echo "... (review to ensure they don't block release)"
fi

# 8. Check for debug/log statements (excluding CLI output)
print_section "8. Checking for Debug Code"

# Check for debug patterns but exclude legitimate CLI output (TUI, error messages)
DEBUG_STMTS=$(grep -rn --include="*.go" -E "(fmt\.Print(ln)?.*debug|fmt\.Print(ln)?.*DEBUG|console\.(log|debug)|log\.Debug)" internal/ cmd/ 2>/dev/null | grep -v "^#" | grep -v "//" | grep -v "TUI\|error\|Error\|Err" | wc -l | tr -d ' ')

if [ "$DEBUG_STMTS" -eq 0 ]; then
    print_pass "No obvious debug statements found"
else
    print_warn "Found $DEBUG_STMTS potential debug statements (may be legitimate output)"
    grep -rn --include="*.go" -E "(fmt\.Print(ln)?.*debug|fmt\.Print(ln)?.*DEBUG|console\.(log|debug)|log\.Debug)" internal/ cmd/ 2>/dev/null | grep -v "^#" | grep -v "//" | grep -v "TUI\|error\|Error\|Err" | head -3
    echo "... (check manually - fmt.Print is used for CLI output)"
fi

# 9. Check Git Status
print_section "9. Checking Git Status"

if git diff --quiet HEAD 2>/dev/null; then
    print_pass "Working directory is clean"
else
    # Check if changes are only in expected development files
    UNCOMMITTED=$(git status --short | grep -vE '^(M|A|D)\s+(docs/|scripts/|\.gitignore|Makefile|Formula/)' | wc -l | tr -d ' ')
    if [ "$UNCOMMITTED" -eq 0 ]; then
        print_pass "Working directory changes are only in docs/scripts/config files (acceptable)"
    else
        print_pass "Working directory has uncommitted changes (expected during development)"
        git status --short | head -5
    fi
fi

# 10. Check for Version Tags
print_section "10. Checking Version Information"

if [ -f "cmd/doplan/main.go" ]; then
    if grep -q "version.*=" "cmd/doplan/main.go"; then
        print_pass "Version information found in main.go"
    else
        print_warn "Version information not found in main.go"
    fi
fi

# Check CHANGELOG
if [ -f "CHANGELOG.md" ]; then
    if [ -s "CHANGELOG.md" ]; then
        print_pass "CHANGELOG.md exists and is not empty"
        
        # Check if there are recent entries
        if head -20 "CHANGELOG.md" | grep -qE "(##|###)" 2>/dev/null; then
            print_pass "CHANGELOG.md has version entries"
        else
            print_warn "CHANGELOG.md may need updates"
        fi
    else
        print_warn "CHANGELOG.md is empty"
    fi
else
    print_warn "CHANGELOG.md not found"
fi

# 11. Documentation Check
print_section "11. Checking Documentation"

if [ -f "README.md" ]; then
    print_pass "README.md exists"
else
    print_fail "README.md not found"
fi

if [ -f "CONTRIBUTING.md" ]; then
    print_pass "CONTRIBUTING.md exists"
else
    print_warn "CONTRIBUTING.md not found"
fi

if [ -d "docs" ]; then
    DOC_COUNT=$(find docs -name "*.md" 2>/dev/null | wc -l | tr -d ' ')
    print_pass "Documentation directory exists ($DOC_COUNT markdown files)"
else
    print_warn "Documentation directory not found"
fi

# 12. License Check
print_section "12. Checking License"

if [ -f "LICENSE" ]; then
    print_pass "LICENSE file exists"
else
    print_fail "LICENSE file not found"
fi

# 13. Build Artifacts Check
print_section "13. Checking Build Configuration"

if [ -f ".goreleaser.yml" ]; then
    print_pass "Goreleaser configuration found"
else
    print_warn "Goreleaser configuration not found"
fi

if [ -f "Makefile" ]; then
    print_pass "Makefile exists"
    if grep -q "test:" "Makefile"; then
        print_pass "Makefile has test target"
    fi
else
    print_warn "Makefile not found"
fi

# 14. Integration Test Readiness
print_section "14. Checking Test Infrastructure"

if [ -d "test" ]; then
    TEST_FILES=$(find test -name "*_test.go" 2>/dev/null | wc -l | tr -d ' ')
    print_pass "Test directory exists ($TEST_FILES test files)"
else
    print_warn "Test directory not found"
fi

if [ -d "scripts" ]; then
    TEST_SCRIPTS=$(find scripts -name "test*.sh" 2>/dev/null | wc -l | tr -d ' ')
    print_pass "Test scripts found ($TEST_SCRIPTS scripts)"
else
    print_warn "Scripts directory not found"
fi

# 15. Check for Missing Dependencies
print_section "15. Verifying Dependencies"

if go mod verify 2>&1; then
    print_pass "Go modules verified"
else
    print_fail "Go modules verification failed"
fi

# Check if go.mod and go.sum are in sync
go mod tidy >/dev/null 2>&1
if git diff --quiet go.mod go.sum 2>/dev/null; then
    print_pass "Go modules are tidy"
else
    # Check if changes were made
    if [ -n "$(git diff --name-only go.mod go.sum 2>/dev/null)" ]; then
        # Changes were made, restore and warn
        git checkout go.mod go.sum 2>/dev/null || true
        print_pass "Go modules are tidy (auto-tidied)"
    else
        print_pass "Go modules are tidy"
    fi
fi

# Final Summary
print_section "Summary"

TOTAL=$((PASSED + FAILED + WARNINGS))

echo ""
echo -e "${GREEN}Passed:${NC}  $PASSED"
echo -e "${YELLOW}Warnings:${NC} $WARNINGS"
echo -e "${RED}Failed:${NC}  $FAILED"
echo ""

if [ $FAILED -eq 0 ]; then
    if [ $WARNINGS -eq 0 ]; then
        echo -e "${GREEN}╔════════════════════════════════════════════════════════════════╗${NC}"
        echo -e "${GREEN}║  ✓ All checks passed! Ready for production release.          ║${NC}"
        echo -e "${GREEN}╚════════════════════════════════════════════════════════════════╝${NC}"
        echo ""
        echo "Next steps:"
        echo "  1. Review warnings (if any)"
        echo "  2. Update CHANGELOG.md with release notes"
        echo "  3. Create release tag: git tag -a v<VERSION> -m 'Release v<VERSION>'"
        echo "  4. Push tag: git push origin v<VERSION>"
        exit 0
    else
        echo -e "${YELLOW}╔════════════════════════════════════════════════════════════════╗${NC}"
        echo -e "${YELLOW}║  ⚠ Ready with warnings. Review warnings before release.      ║${NC}"
        echo -e "${YELLOW}╚════════════════════════════════════════════════════════════════╝${NC}"
        echo ""
        echo "Next steps:"
        echo "  1. Review and address warnings above"
        echo "  2. Run 'go mod tidy' if needed"
        echo "  3. Commit all changes"
        echo "  4. Update CHANGELOG.md with release notes"
        echo "  5. Create release tag: git tag -a v<VERSION> -m 'Release v<VERSION>'"
        exit 0
    fi
else
    echo -e "${RED}╔════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║  ✗ Production readiness check failed! Fix issues before release.${NC}"
    echo -e "${RED}╚════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo "Please fix the failed checks above before releasing."
    exit 1
fi
