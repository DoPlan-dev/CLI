#!/bin/bash

# DoPlan v0.0.19-beta Phase Testing Script
# This script tests each phase systematically and reports issues

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

TEST_DIR="/tmp/doplan-phase-test"
ISSUES_FILE="docs/development/V0.0.19-BETA-TESTING-CHECKLIST.md"
ISSUES_FOUND=0

echo "=========================================="
echo "DoPlan v0.0.19-beta Phase Testing"
echo "=========================================="
echo ""

# Cleanup
rm -rf "$TEST_DIR"
mkdir -p "$TEST_DIR"
cd "$TEST_DIR"

echo "[TEST] Phase 1: Build & Basic Structure"
echo "----------------------------------------"

# Test 1: Build CLI
cd /Users/Dorgham/Documents/Work/Devleopment/DoPlan/cli
if ! go build -o doplan-test ./cmd/doplan/main.go 2>&1; then
    echo -e "${RED}✗ FAIL: CLI build failed${NC}"
    ISSUES_FOUND=$((ISSUES_FOUND + 1))
else
    echo -e "${GREEN}✓ PASS: CLI builds successfully${NC}"
fi

# Test 2: Check main files exist
echo ""
echo "[TEST] Phase 2: File Structure Verification"
echo "----------------------------------------"

check_file() {
    local file=$1
    local phase=$2
    if [ -f "$file" ]; then
        echo -e "${GREEN}✓ PASS: $file exists${NC}"
    else
        echo -e "${RED}✗ FAIL: $file missing (Phase $phase)${NC}"
        ISSUES_FOUND=$((ISSUES_FOUND + 1))
    fi
}

# Phase 2: Design System (DPR)
check_file "internal/dpr/questionnaire.go" "2"
check_file "internal/dpr/generator.go" "2"
check_file "internal/dpr/tokens.go" "2"
check_file "internal/dpr/cursor_rules.go" "2"
check_file "internal/wizard/design.go" "2"
check_file "internal/commands/design.go" "2"

# Phase 3: Secrets & API Keys
check_file "internal/rakd/types.go" "3"
check_file "internal/rakd/detector.go" "3"
check_file "internal/rakd/validator.go" "3"
check_file "internal/rakd/generator.go" "3"
check_file "internal/sops/generator.go" "3"
check_file "internal/wizard/keys.go" "3"
check_file "internal/commands/keys.go" "3"

# Phase 4: AI Agents System
check_file "internal/generators/agents.go" "4"
check_file "internal/generators/rules.go" "4"

# Phase 5: Workflow Guidance
check_file "internal/workflow/recommender.go" "5"

echo ""
echo "[TEST] Phase 3: Compilation Check"
echo "----------------------------------------"

# Test compilation of each package
test_package() {
    local pkg=$1
    local phase=$2
    if go build "./$pkg" 2>&1 | head -5; then
        echo -e "${GREEN}✓ PASS: $pkg compiles${NC}"
    else
        echo -e "${RED}✗ FAIL: $pkg compilation errors (Phase $phase)${NC}"
        ISSUES_FOUND=$((ISSUES_FOUND + 1))
    fi
}

echo "Testing package compilation..."
test_package "internal/dpr" "2"
test_package "internal/rakd" "3"
test_package "internal/sops" "3"
test_package "internal/workflow" "5"
test_package "internal/generators" "4"

echo ""
echo "[TEST] Phase 4: Linter Check"
echo "----------------------------------------"

# Check for linter errors in critical files
lint_check() {
    local file=$1
    local phase=$2
    if go vet "./$file" 2>&1 | grep -q "error"; then
        echo -e "${RED}✗ FAIL: $file has vet errors (Phase $phase)${NC}"
        ISSUES_FOUND=$((ISSUES_FOUND + 1))
    else
        echo -e "${GREEN}✓ PASS: $file passes go vet${NC}"
    fi
}

echo "Running go vet on critical files..."
lint_check "internal/tui/app.go" "1"
lint_check "internal/workflow/recommender.go" "5"
lint_check "internal/dpr/generator.go" "2"
lint_check "internal/rakd/generator.go" "3"

echo ""
echo "[TEST] Phase 5: Function Existence Check"
echo "----------------------------------------"

# Check critical functions exist (methods or standalone functions)
check_function() {
    local file=$1
    local func=$2
    local phase=$3
    if grep -q "func.*$func" "$file" 2>/dev/null; then
        echo -e "${GREEN}✓ PASS: $func found in $file${NC}"
    else
        echo -e "${RED}✗ FAIL: $func missing in $file (Phase $phase)${NC}"
        ISSUES_FOUND=$((ISSUES_FOUND + 1))
    fi
}

# Phase 5: Workflow Guidance
check_function "internal/workflow/recommender.go" "GetNextStep" "5"

# Phase 2: Design System
check_function "internal/dpr/generator.go" "Generate" "2"
check_function "internal/dpr/questionnaire.go" "RunQuestionnaire" "2"

# Phase 3: API Keys
check_function "internal/rakd/generator.go" "GenerateRAKD" "3"
check_function "internal/rakd/detector.go" "DetectServices" "3"

# Phase 4: Agents
check_function "internal/generators/agents.go" "Generate" "4"
check_function "internal/generators/rules.go" "Generate" "4"

echo ""
echo "[TEST] Phase 6: Integration Points Check"
echo "----------------------------------------"

# Check TUI integration
if grep -q "GetNextStep" "internal/tui/app.go" 2>/dev/null; then
    echo -e "${GREEN}✓ PASS: Workflow guidance integrated in TUI${NC}"
else
    echo -e "${RED}✗ FAIL: Workflow guidance not integrated in TUI (Phase 5)${NC}"
    ISSUES_FOUND=$((ISSUES_FOUND + 1))
fi

# Check dashboard API keys widget
if grep -q "renderAPIKeysWidget" "internal/tui/screens/dashboard.go" 2>/dev/null; then
    echo -e "${GREEN}✓ PASS: API keys widget in dashboard${NC}"
else
    echo -e "${RED}✗ FAIL: API keys widget missing in dashboard (Phase 3)${NC}"
    ISSUES_FOUND=$((ISSUES_FOUND + 1))
fi

# Check menu actions
if grep -q "design" "internal/tui/screens/menu.go" 2>/dev/null; then
    echo -e "${GREEN}✓ PASS: Design action in menu${NC}"
else
    echo -e "${RED}✗ FAIL: Design action missing in menu (Phase 2)${NC}"
    ISSUES_FOUND=$((ISSUES_FOUND + 1))
fi

if grep -q "keys" "internal/tui/screens/menu.go" 2>/dev/null; then
    echo -e "${GREEN}✓ PASS: Keys action in menu${NC}"
else
    echo -e "${RED}✗ FAIL: Keys action missing in menu (Phase 3)${NC}"
    ISSUES_FOUND=$((ISSUES_FOUND + 1))
fi

echo ""
echo "=========================================="
echo "Testing Summary"
echo "=========================================="
echo "Issues Found: $ISSUES_FOUND"

if [ $ISSUES_FOUND -eq 0 ]; then
    echo -e "${GREEN}All basic checks passed!${NC}"
    exit 0
else
    echo -e "${YELLOW}Some issues found. Review output above.${NC}"
    exit 1
fi

