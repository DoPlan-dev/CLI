#!/bin/bash

# Non-Interactive Installation Test
# This simulates installation by checking what would be created

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BINARY_PATH="$PROJECT_ROOT/bin/doplan"
TEST_PROJECT="$PROJECT_ROOT/../test-installation"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  DoPlan Installation Test${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

cd "$TEST_PROJECT"

# Check if already installed
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo -e "${YELLOW}DoPlan is already installed.${NC}"
    echo -e "${CYAN}Verifying installation...${NC}"
    echo ""
else
    echo -e "${CYAN}DoPlan is not installed yet.${NC}"
    echo -e "${YELLOW}To install manually, run:${NC}"
    echo "  cd $TEST_PROJECT"
    echo "  $BINARY_PATH install"
    echo ""
    echo -e "${CYAN}Then select option 1 (Cursor) when prompted.${NC}"
    echo ""
    exit 0
fi

# Verification
PASSED=0
FAILED=0

check() {
    local path="$1"
    local name="$2"
    echo -n "  $name... "
    if [ -e "$path" ]; then
        echo -e "${GREEN}✓${NC}"
        PASSED=$((PASSED + 1))
    else
        echo -e "${YELLOW}✗${NC}"
        FAILED=$((FAILED + 1))
    fi
}

echo -e "${BLUE}Verifying installation structure...${NC}"
echo ""

# Directories
echo -e "${CYAN}Directories:${NC}"
check ".cursor/commands" "Commands"
check ".cursor/rules" "Rules"
check ".cursor/config" "Config"
check "doplan/contracts" "Contracts"
check "doplan/templates" "Templates"
echo ""

# Config file
echo -e "${CYAN}Configuration:${NC}"
check ".cursor/config/doplan-config.json" "Config file"
echo ""

# Command files
echo -e "${CYAN}Command files:${NC}"
check ".cursor/commands/discuss.json" "Discuss"
check ".cursor/commands/refine.json" "Refine"
check ".cursor/commands/generate.json" "Generate"
check ".cursor/commands/plan.json" "Plan"
check ".cursor/commands/dashboard.json" "Dashboard"
check ".cursor/commands/implement.json" "Implement"
check ".cursor/commands/next.json" "Next"
check ".cursor/commands/progress.json" "Progress"
echo ""

# Rule files
echo -e "${CYAN}Rule files:${NC}"
check ".cursor/rules/workflow-rules.md" "Workflow"
check ".cursor/rules/github-rules.md" "GitHub"
check ".cursor/rules/command-rules.md" "Commands"
check ".cursor/rules/branch-rules.md" "Branches"
check ".cursor/rules/commit-rules.md" "Commits"
echo ""

# Template files
echo -e "${CYAN}Template files:${NC}"
check "doplan/templates/plan-template.md" "Plan template"
check "doplan/templates/design-template.md" "Design template"
check "doplan/templates/tasks-template.md" "Tasks template"
echo ""

# Other files
echo -e "${CYAN}Other files:${NC}"
check "doplan/dashboard.md" "Dashboard"
check "README.md" "README"
echo ""

# Show summary
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "${GREEN}Passed: $PASSED${NC}"
if [ $FAILED -gt 0 ]; then
    echo -e "${YELLOW}Missing: $FAILED${NC}"
else
    echo -e "${GREEN}Missing: $FAILED${NC}"
fi
echo ""

# Show config if exists
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo -e "${CYAN}Configuration:${NC}"
    cat .cursor/config/doplan-config.json | python3 -m json.tool 2>/dev/null || cat .cursor/config/doplan-config.json
    echo ""
fi

# Show command count
if [ -d ".cursor/commands" ]; then
    CMD_COUNT=$(ls -1 .cursor/commands/*.json 2>/dev/null | wc -l | tr -d ' ')
    echo -e "${CYAN}Commands installed: $CMD_COUNT${NC}"
    echo ""
fi

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ Installation verified!${NC}"
    echo ""
    echo -e "${CYAN}Test commands:${NC}"
    echo "  cd $TEST_PROJECT"
    echo "  $BINARY_PATH dashboard"
    echo "  $BINARY_PATH github"
    echo "  $BINARY_PATH progress"
    exit 0
else
    echo -e "${YELLOW}⚠ Some files are missing${NC}"
    echo -e "${YELLOW}Run installation: $BINARY_PATH install${NC}"
    exit 1
fi

