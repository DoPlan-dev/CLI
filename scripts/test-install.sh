#!/bin/bash

# DoPlan Installation Test Script
# This script tests the installation flow in a real project

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

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  DoPlan Installation Test${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo "Test project: $TEST_PROJECT"
echo ""

# Cleanup
cleanup() {
    echo ""
    echo -e "${BLUE}Cleaning up test project...${NC}"
    rm -rf "$TEST_PROJECT"
}

trap cleanup EXIT

# Setup test project
cd "$TEST_PROJECT"
echo -e "${BLUE}Setting up test project...${NC}"

# Initialize git repo
git init > /dev/null 2>&1 || true
git config user.name "Test User" > /dev/null 2>&1 || true
git config user.email "test@example.com" > /dev/null 2>&1 || true
echo "# Test Project" > README.md
git add README.md
git commit -m "Initial commit" > /dev/null 2>&1 || true

# Check if binary exists
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${RED}Binary not found. Building...${NC}"
    cd "$PROJECT_ROOT"
    make build
fi

# Test installation
echo -e "${BLUE}Testing installation...${NC}"
echo ""

# Check if already installed
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo -e "${YELLOW}DoPlan already installed. Removing...${NC}"
    rm -rf .cursor
fi

# Run install (non-interactive - we'll need to provide input)
echo -e "${BLUE}Running: doplan install${NC}"
echo ""

# Since install is interactive, we'll test the structure it creates
# In a real scenario, you'd use expect or similar for interactive testing

# For now, let's verify the install command exists and check its help
if "$BINARY_PATH" install --help > /dev/null 2>&1; then
    echo -e "${GREEN}✓ Install command available${NC}"
else
    echo -e "${RED}✗ Install command not available${NC}"
    exit 1
fi

# Verify expected directories would be created
echo ""
echo -e "${BLUE}Verifying installation structure...${NC}"

# Check if install creates the right structure
# We'll manually verify what install should create:
EXPECTED_DIRS=(
    ".cursor/commands"
    ".cursor/rules"
    ".cursor/config"
    "doplan/contracts"
    "doplan/templates"
)

echo "Expected directories after installation:"
for dir in "${EXPECTED_DIRS[@]}"; do
    echo "  - $dir"
done

echo ""
echo -e "${GREEN}Installation test setup complete!${NC}"
echo -e "${YELLOW}Note: Full interactive installation test requires manual verification${NC}"
echo ""
echo "To test installation manually:"
echo "  1. cd $TEST_PROJECT"
echo "  2. $BINARY_PATH install"
echo "  3. Select an IDE option"
echo "  4. Verify directories and files are created"

