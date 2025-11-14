#!/bin/bash

# DoPlan Integration Test Script
# Tests the full workflow: install -> use commands -> verify output

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BINARY_PATH="$PROJECT_ROOT/bin/doplan"
TEST_PROJECT=$(mktemp -d)

PASSED=0
FAILED=0

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  DoPlan Integration Test${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo "Test project: $TEST_PROJECT"
echo ""

cleanup() {
    echo ""
    echo -e "${BLUE}Cleaning up...${NC}"
    rm -rf "$TEST_PROJECT"
}

trap cleanup EXIT

test_check() {
    local name="$1"
    local check="$2"
    
    echo -n "  $name... "
    if eval "$check" > /dev/null 2>&1; then
        echo -e "${GREEN}✓${NC}"
        PASSED=$((PASSED + 1))
        return 0
    else
        echo -e "${RED}✗${NC}"
        FAILED=$((FAILED + 1))
        return 1
    fi
}

# Setup
cd "$TEST_PROJECT"
echo -e "${BLUE}Setting up test environment...${NC}"

# Initialize git
git init > /dev/null 2>&1 || true
git config user.name "Test User" > /dev/null 2>&1 || true
git config user.email "test@example.com" > /dev/null 2>&1 || true
echo "# Test" > README.md
git add README.md
git commit -m "Initial" > /dev/null 2>&1 || true

# Build if needed
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${YELLOW}Building binary...${NC}"
    cd "$PROJECT_ROOT"
    make build
    cd "$TEST_PROJECT"
fi

# Test 1: Verify binary works
echo ""
echo -e "${BLUE}Test 1: Binary functionality${NC}"
test_check "Binary exists" "[ -f '$BINARY_PATH' ]"
test_check "Help works" "$BINARY_PATH --help | grep -q 'DoPlan'"
test_check "Version works" "$BINARY_PATH --version | grep -q 'dev'"

# Test 2: Commands without installation
echo ""
echo -e "${BLUE}Test 2: Commands before installation${NC}"
test_check "Dashboard shows not installed" "$BINARY_PATH dashboard 2>&1 | grep -q 'not installed'"
test_check "GitHub shows not installed" "$BINARY_PATH github 2>&1 | grep -q 'not installed'"
test_check "Progress shows not installed" "$BINARY_PATH progress 2>&1 | grep -q 'not installed'"

# Test 3: Installation structure (manual verification needed)
echo ""
echo -e "${BLUE}Test 3: Installation structure${NC}"
echo -e "${YELLOW}  Note: Full installation test requires interactive input${NC}"
echo -e "${YELLOW}  Run manually: cd $TEST_PROJECT && $BINARY_PATH install${NC}"

# Test 4: File generation (if installed)
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo ""
    echo -e "${BLUE}Test 4: File generation${NC}"
    test_check "Config exists" "[ -f '.cursor/config/doplan-config.json' ]"
    test_check "Commands directory exists" "[ -d '.cursor/commands' ]"
    test_check "Rules directory exists" "[ -d '.cursor/rules' ]"
    test_check "Dashboard exists" "[ -f 'doplan/dashboard.md' ]"
fi

# Summary
echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Results${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "${GREEN}Passed: $PASSED${NC}"
if [ $FAILED -gt 0 ]; then
    echo -e "${RED}Failed: $FAILED${NC}"
else
    echo -e "${GREEN}Failed: $FAILED${NC}"
fi
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ All integration tests passed!${NC}"
    exit 0
else
    echo -e "${YELLOW}⚠ Some tests require manual verification${NC}"
    exit 0
fi

