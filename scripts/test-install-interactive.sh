#!/bin/bash

# Interactive Installation Test Script
# This script helps test the interactive installation flow

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BINARY_PATH="$PROJECT_ROOT/bin/doplan"
TEST_PROJECT="$PROJECT_ROOT/../test-installation"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  DoPlan Interactive Installation Test${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Check if binary exists
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${YELLOW}Binary not found. Building...${NC}"
    cd "$PROJECT_ROOT"
    make build
fi

# Check if test project exists
if [ ! -d "$TEST_PROJECT" ]; then
    echo -e "${YELLOW}Creating test project...${NC}"
    mkdir -p "$TEST_PROJECT"
    cd "$TEST_PROJECT"
    git init > /dev/null 2>&1 || true
    git config user.name "Test User" > /dev/null 2>&1 || true
    git config user.email "test@example.com" > /dev/null 2>&1 || true
    echo "# Test Project" > README.md
    git add README.md
    git commit -m "Initial commit" > /dev/null 2>&1 || true
fi

cd "$TEST_PROJECT"

echo -e "${CYAN}Test Project: $TEST_PROJECT${NC}"
echo ""

# Check if already installed
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo -e "${YELLOW}DoPlan is already installed in this project.${NC}"
    echo -e "${YELLOW}Would you like to test reinstallation? (y/n)${NC}"
    read -r response
    if [ "$response" != "y" ] && [ "$response" != "Y" ]; then
        echo -e "${BLUE}Exiting. To test fresh installation, remove:${NC}"
        echo "  rm -rf $TEST_PROJECT/.cursor"
        exit 0
    fi
fi

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Installation Instructions${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo -e "${CYAN}1. You will see a menu asking:${NC}"
echo -e "   ${YELLOW}What are you using to develop your application?${NC}"
echo ""
echo -e "${CYAN}2. Select an option:${NC}"
echo -e "   ${GREEN}[1] Cursor${NC}"
echo -e "   ${GREEN}[2] Gemini CLI${NC}"
echo -e "   ${GREEN}[3] Claude CLI${NC}"
echo -e "   ${GREEN}[4] Codex${NC}"
echo -e "   ${GREEN}[5] Back${NC}"
echo ""
echo -e "${CYAN}3. After selection, DoPlan will:${NC}"
echo -e "   • Create directory structure"
echo -e "   • Install IDE commands"
echo -e "   • Create templates"
echo -e "   • Generate workflow rules"
echo -e "   • Generate configuration"
echo -e "   • Create README"
echo -e "   • Create initial dashboard"
echo ""
echo -e "${BLUE}========================================${NC}"
echo ""
echo -e "${YELLOW}Press Enter to start installation...${NC}"
read -r

echo ""
echo -e "${BLUE}Running: doplan install${NC}"
echo ""

# Run installation
"$BINARY_PATH" install

echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Verification${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Verify installation
VERIFY_PASSED=0
VERIFY_FAILED=0

check_file() {
    local file="$1"
    local name="$2"
    echo -n "Checking $name... "
    if [ -f "$file" ] || [ -d "$file" ]; then
        echo -e "${GREEN}✓${NC}"
        VERIFY_PASSED=$((VERIFY_PASSED + 1))
    else
        echo -e "${RED}✗${NC}"
        VERIFY_FAILED=$((VERIFY_FAILED + 1))
    fi
}

check_dir() {
    local dir="$1"
    local name="$2"
    echo -n "Checking $name... "
    if [ -d "$dir" ]; then
        echo -e "${GREEN}✓${NC}"
        VERIFY_PASSED=$((VERIFY_PASSED + 1))
    else
        echo -e "${RED}✗${NC}"
        VERIFY_FAILED=$((VERIFY_FAILED + 1))
    fi
}

echo -e "${CYAN}Verifying installation structure...${NC}"
echo ""

# Check directories
check_dir ".cursor/commands" "Commands directory"
check_dir ".cursor/rules" "Rules directory"
check_dir ".cursor/config" "Config directory"
check_dir "doplan/contracts" "Contracts directory"
check_dir "doplan/templates" "Templates directory"

# Check files
check_file ".cursor/config/doplan-config.json" "Config file"
check_file "doplan/dashboard.md" "Dashboard file"
check_file "README.md" "README file"

# Check command files
check_file ".cursor/commands/discuss.json" "Discuss command"
check_file ".cursor/commands/refine.json" "Refine command"
check_file ".cursor/commands/generate.json" "Generate command"
check_file ".cursor/commands/plan.json" "Plan command"
check_file ".cursor/commands/dashboard.json" "Dashboard command"
check_file ".cursor/commands/implement.json" "Implement command"
check_file ".cursor/commands/next.json" "Next command"
check_file ".cursor/commands/progress.json" "Progress command"

# Check rule files
check_file ".cursor/rules/workflow-rules.md" "Workflow rules"
check_file ".cursor/rules/github-rules.md" "GitHub rules"
check_file ".cursor/rules/command-rules.md" "Command rules"
check_file ".cursor/rules/branch-rules.md" "Branch rules"
check_file ".cursor/rules/commit-rules.md" "Commit rules"

# Check template files
check_file "doplan/templates/plan-template.md" "Plan template"
check_file "doplan/templates/design-template.md" "Design template"
check_file "doplan/templates/tasks-template.md" "Tasks template"

echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Verification Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "${GREEN}Passed: $VERIFY_PASSED${NC}"
if [ $VERIFY_FAILED -gt 0 ]; then
    echo -e "${RED}Failed: $VERIFY_FAILED${NC}"
else
    echo -e "${GREEN}Failed: $VERIFY_FAILED${NC}"
fi
echo ""

# Show config content
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo -e "${CYAN}Configuration file content:${NC}"
    cat .cursor/config/doplan-config.json | head -20
    echo ""
fi

# Show dashboard
if [ -f "doplan/dashboard.md" ]; then
    echo -e "${CYAN}Dashboard preview:${NC}"
    head -20 doplan/dashboard.md
    echo ""
fi

# List commands
if [ -d ".cursor/commands" ]; then
    echo -e "${CYAN}Installed commands:${NC}"
    ls -1 .cursor/commands/ | sed 's/^/  - /'
    echo ""
fi

if [ $VERIFY_FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ Installation verified successfully!${NC}"
    echo ""
    echo -e "${CYAN}Next steps:${NC}"
    echo "  1. Try: doplan dashboard"
    echo "  2. Try: doplan github"
    echo "  3. Try: doplan progress"
    echo "  4. Check Cursor commands in .cursor/commands/"
    exit 0
else
    echo -e "${YELLOW}⚠ Some files/directories are missing${NC}"
    echo -e "${YELLOW}Check the installation output above for errors${NC}"
    exit 1
fi

