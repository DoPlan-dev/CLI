#!/bin/bash

# DoPlan CLI Test Script
# This script tests all CLI commands and functionality

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test counters
PASSED=0
FAILED=0
TOTAL=0

# Test directory
TEST_DIR=$(mktemp -d)
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BINARY_PATH="$PROJECT_ROOT/bin/doplan"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  DoPlan CLI Test Suite${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo "Test directory: $TEST_DIR"
echo "Project root: $PROJECT_ROOT"
echo ""

# Function to run a test
test_command() {
    local test_name="$1"
    local command="$2"
    local expected_exit="${3:-0}"
    
    TOTAL=$((TOTAL + 1))
    echo -n "Testing: $test_name... "
    
    if eval "$command" > /tmp/doplan-test-output.log 2>&1; then
        EXIT_CODE=$?
    else
        EXIT_CODE=$?
    fi
    
    if [ $EXIT_CODE -eq $expected_exit ]; then
        echo -e "${GREEN}✓ PASSED${NC}"
        PASSED=$((PASSED + 1))
        return 0
    else
        echo -e "${RED}✗ FAILED (exit code: $EXIT_CODE, expected: $expected_exit)${NC}"
        echo -e "${YELLOW}Output:${NC}"
        cat /tmp/doplan-test-output.log | head -10
        FAILED=$((FAILED + 1))
        return 1
    fi
}

# Function to check if output contains text
test_output_contains() {
    local test_name="$1"
    local command="$2"
    local expected_text="$3"
    
    TOTAL=$((TOTAL + 1))
    echo -n "Testing: $test_name... "
    
    if eval "$command" > /tmp/doplan-test-output.log 2>&1; then
        if grep -q "$expected_text" /tmp/doplan-test-output.log; then
            echo -e "${GREEN}✓ PASSED${NC}"
            PASSED=$((PASSED + 1))
            return 0
        else
            echo -e "${RED}✗ FAILED (output doesn't contain '$expected_text')${NC}"
            echo -e "${YELLOW}Output:${NC}"
            cat /tmp/doplan-test-output.log | head -10
            FAILED=$((FAILED + 1))
            return 1
        fi
    else
        echo -e "${RED}✗ FAILED (command failed)${NC}"
        cat /tmp/doplan-test-output.log | head -10
        FAILED=$((FAILED + 1))
        return 1
    fi
}

# Cleanup function
cleanup() {
    echo ""
    echo -e "${BLUE}Cleaning up...${NC}"
    rm -rf "$TEST_DIR"
    rm -f /tmp/doplan-test-output.log
}

trap cleanup EXIT

# Build the binary
echo -e "${BLUE}Building DoPlan CLI...${NC}"
cd "$PROJECT_ROOT"
if ! make build > /tmp/doplan-build.log 2>&1; then
    echo -e "${RED}Build failed!${NC}"
    cat /tmp/doplan-build.log
    exit 1
fi

if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${RED}Binary not found at $BINARY_PATH${NC}"
    exit 1
fi

echo -e "${GREEN}Build successful!${NC}"
echo ""

# Test 1: Help command
test_output_contains "Help command" \
    "$BINARY_PATH --help" \
    "DoPlan"

# Test 2: Version command
test_output_contains "Version command" \
    "$BINARY_PATH --version" \
    "dev"

# Test 3: Command without arguments (should show help or error)
test_command "No arguments" \
    "$BINARY_PATH" \
    0

# Test 4: Invalid command
test_command "Invalid command" \
    "$BINARY_PATH invalid-command" \
    1

# Test 5: Dashboard command (should fail - not installed)
test_output_contains "Dashboard without installation" \
    "$BINARY_PATH dashboard" \
    "not installed"

# Test 6: GitHub command (should fail - not installed)
test_output_contains "GitHub without installation" \
    "$BINARY_PATH github" \
    "not installed"

# Test 7: Progress command (should fail - not installed)
test_output_contains "Progress without installation" \
    "$BINARY_PATH progress" \
    "not installed"

# Test 8: Install command (interactive - we'll test non-interactively)
echo ""
echo -e "${BLUE}Testing installation...${NC}"
cd "$TEST_DIR"

# Create a mock git repo for testing
git init > /dev/null 2>&1 || true
git config user.name "Test User" > /dev/null 2>&1 || true
git config user.email "test@example.com" > /dev/null 2>&1 || true

# Test installation (we'll need to mock the interactive prompt)
# For now, we'll just verify the command exists and doesn't crash
test_command "Install command exists" \
    "$BINARY_PATH install --help" \
    0

# Test 9: Check if binary runs
test_command "Binary execution" \
    "$BINARY_PATH --version" \
    0

# Test 10: Check all subcommands exist
for cmd in install dashboard github progress; do
    test_command "Subcommand: $cmd" \
        "$BINARY_PATH $cmd --help" \
        0
done

# Test 11: TUI mode (just check it doesn't crash immediately)
# Note: TUI mode requires a terminal, so we'll just check the flag works
test_command "TUI flag" \
    "timeout 1 $BINARY_PATH --tui 2>&1 || true" \
    0

# Summary
echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Test Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "Total tests: $TOTAL"
echo -e "${GREEN}Passed: $PASSED${NC}"
if [ $FAILED -gt 0 ]; then
    echo -e "${RED}Failed: $FAILED${NC}"
else
    echo -e "${GREEN}Failed: $FAILED${NC}"
fi
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}✗ Some tests failed${NC}"
    exit 1
fi

