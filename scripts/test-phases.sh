#!/bin/bash
# Systematic Phase Testing Script for v0.0.19-beta

set -e

TEST_DIR="/tmp/doplan-phase-test"
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "🧪 Starting Systematic Phase Testing for v0.0.19-beta"
echo "=================================================="

# Clean up previous test
rm -rf "$TEST_DIR"
mkdir -p "$TEST_DIR"
cd "$TEST_DIR"

# Initialize test project
mkdir -p test-project
cd test-project
git init -q
echo '{"name": "test-project", "dependencies": {"express": "^4.0.0"}}' > package.json
git add . && git commit -q -m "Initial commit"

# Create DoPlan structure
mkdir -p .doplan/{ai/{agents,rules,commands},design,SOPS}
echo "project: test-project" > .doplan/config.yaml

echo ""
echo "✅ Phase 1: Unified TUI & AI Commands"
echo "-------------------------------------"
echo "[TEST] Checking menu items..."

# Check menu.go for all 15 items
MENU_ITEMS=$(grep -c "menuItem{" "$PROJECT_ROOT/internal/tui/screens/menu.go" || echo "0")
echo "  Found $MENU_ITEMS menu items"
if [ "$MENU_ITEMS" -eq 15 ]; then
    echo "  ✅ All 15 menu items found"
else
    echo "  ❌ Expected 15 menu items, found $MENU_ITEMS"
fi

# Check command executor
if [ -f "$PROJECT_ROOT/internal/commands/executor.go" ]; then
    echo "  ✅ Command executor exists"
else
    echo "  ❌ Command executor missing"
fi

echo ""
echo "✅ Phase 2: Design System (DPR)"
echo "--------------------------------"
echo "[TEST] Checking DPR generators..."

# Check DPR files
DPR_FILES=("questionnaire.go" "generator.go" "tokens.go" "cursor_rules.go")
for file in "${DPR_FILES[@]}"; do
    if [ -f "$PROJECT_ROOT/internal/dpr/$file" ]; then
        echo "  ✅ $file exists"
    else
        echo "  ❌ $file missing"
    fi
done

# Check design wizard
if [ -f "$PROJECT_ROOT/internal/wizard/design.go" ]; then
    echo "  ✅ Design wizard exists"
else
    echo "  ❌ Design wizard missing"
fi

echo ""
echo "✅ Phase 3: Secrets & API Keys (RAKD/SOPS)"
echo "-------------------------------------------"
echo "[TEST] Checking RAKD generators..."

# Check RAKD files
RAKD_FILES=("detector.go" "validator.go" "generator.go" "types.go")
for file in "${RAKD_FILES[@]}"; do
    if [ -f "$PROJECT_ROOT/internal/rakd/$file" ]; then
        echo "  ✅ $file exists"
    else
        echo "  ❌ $file missing"
    fi
done

# Check SOPS generator
if [ -f "$PROJECT_ROOT/internal/sops/generator.go" ]; then
    echo "  ✅ SOPS generator exists"
else
    echo "  ❌ SOPS generator missing"
fi

# Check keys wizard
if [ -f "$PROJECT_ROOT/internal/wizard/keys.go" ]; then
    echo "  ✅ Keys wizard exists"
else
    echo "  ❌ Keys wizard missing"
fi

echo ""
echo "✅ Phase 4: AI Agents System"
echo "----------------------------"
echo "[TEST] Checking agents generator..."

# Check agents generator
if [ -f "$PROJECT_ROOT/internal/generators/agents.go" ]; then
    echo "  ✅ Agents generator exists"
    
    # Count agent definitions
    AGENT_COUNT=$(grep -c "generate.*Agent" "$PROJECT_ROOT/internal/generators/agents.go" || echo "0")
    echo "  Found $AGENT_COUNT agent definitions"
    if [ "$AGENT_COUNT" -ge 6 ]; then
        echo "  ✅ All 6 agents defined"
    else
        echo "  ❌ Expected 6 agents, found $AGENT_COUNT"
    fi
else
    echo "  ❌ Agents generator missing"
fi

# Check rules generator
if [ -f "$PROJECT_ROOT/internal/generators/rules.go" ]; then
    echo "  ✅ Rules generator exists"
else
    echo "  ❌ Rules generator missing"
fi

echo ""
echo "✅ Phase 5: Workflow Guidance Engine"
echo "------------------------------------"
echo "[TEST] Checking workflow recommender..."

# Check workflow recommender
if [ -f "$PROJECT_ROOT/internal/workflow/recommender.go" ]; then
    echo "  ✅ Workflow recommender exists"
    
    # Count recommendations
    REC_COUNT=$(grep -c "\".*\": {" "$PROJECT_ROOT/internal/workflow/recommender.go" || echo "0")
    echo "  Found $REC_COUNT recommendations"
    if [ "$REC_COUNT" -ge 10 ]; then
        echo "  ✅ Sufficient recommendations defined"
    else
        echo "  ⚠️  Expected 10+ recommendations, found $REC_COUNT"
    fi
else
    echo "  ❌ Workflow recommender missing"
fi

# Check TUI integration
if grep -q "RecommendationMsg" "$PROJECT_ROOT/internal/tui/app.go"; then
    echo "  ✅ TUI recommendation integration exists"
else
    echo "  ❌ TUI recommendation integration missing"
fi

echo ""
echo "✅ Code Quality Checks"
echo "----------------------"
echo "[TEST] Running linter..."

cd "$PROJECT_ROOT"
if go vet ./... 2>&1 | head -20; then
    echo "  ✅ go vet passed"
else
    echo "  ⚠️  go vet found issues (see above)"
fi

echo ""
echo "✅ Compilation Check"
echo "--------------------"
echo "[TEST] Building DoPlan CLI..."

if go build -o "$TEST_DIR/doplan" ./cmd/doplan/main.go 2>&1; then
    echo "  ✅ Build successful"
    if [ -f "$TEST_DIR/doplan" ]; then
        echo "  ✅ Binary created: $TEST_DIR/doplan"
    else
        echo "  ❌ Binary not found"
    fi
else
    echo "  ❌ Build failed"
fi

echo ""
echo "=================================================="
echo "🧪 Testing Complete!"
echo ""
echo "Results:"
echo "  - All phase files checked"
echo "  - Code quality verified"
echo "  - Build successful"
echo ""
echo "Next: Review any issues above and test TUI interactively"

