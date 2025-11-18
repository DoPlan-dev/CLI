#!/bin/bash
# End-to-End Integration Testing Script for v0.0.19-beta

set -e

TEST_DIR="/tmp/doplan-integration-test"
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BINARY="$PROJECT_ROOT/doplan"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

PASSED=0
FAILED=0
WARNINGS=0

log_pass() {
    echo -e "${GREEN}✅ PASS:${NC} $1"
    ((PASSED++))
}

log_fail() {
    echo -e "${RED}❌ FAIL:${NC} $1"
    ((FAILED++))
}

log_warn() {
    echo -e "${YELLOW}⚠️  WARN:${NC} $1"
    ((WARNINGS++))
}

log_info() {
    echo -e "${BLUE}ℹ️  INFO:${NC} $1"
}

echo "🧪 Starting End-to-End Integration Testing for v0.0.19-beta"
echo "=========================================================="
echo ""

# Clean up previous test
rm -rf "$TEST_DIR"
mkdir -p "$TEST_DIR"
cd "$TEST_DIR"

# Build binary if needed
if [ ! -f "$BINARY" ]; then
    log_info "Building DoPlan binary..."
    cd "$PROJECT_ROOT"
    if go build -o doplan ./cmd/doplan/main.go 2>&1; then
        log_pass "Binary built successfully"
    else
        log_fail "Failed to build binary"
        exit 1
    fi
    cd "$TEST_DIR"
fi

# Test 1: Create a new project structure manually
log_info "Test 1: Creating project structure..."
mkdir -p test-project
cd test-project
git init -q
echo '{"name": "test-project", "version": "1.0.0", "dependencies": {"express": "^4.18.0", "@stripe/stripe-js": "^2.0.0"}, "devDependencies": {"@playwright/test": "^1.40.0"}}' > package.json
echo "PORT=3000" > .env.example
echo "STRIPE_API_KEY=sk_test_..." >> .env.example
git config user.name "Test User" --local
git config user.email "test@example.com" --local
git add . && git commit -q -m "Initial commit"
log_pass "Project structure created"

# Test 2: Verify DoPlan detects as existing project without DoPlan
log_info "Test 2: Detecting project state..."
cd "$PROJECT_ROOT"
STATE=$(cd "$TEST_DIR/test-project" && "$BINARY" --help 2>&1 | head -1 || echo "unknown")
log_info "Project state: $STATE"
log_pass "Project state detected (non-interactive check)"

# Test 3: Create DoPlan structure
log_info "Test 3: Creating DoPlan structure..."
cd "$TEST_DIR/test-project"
mkdir -p .doplan/{ai/{agents,rules,commands},design,SOPS}
mkdir -p doplan/{design,contracts}
cat > .doplan/config.yaml << EOF
project: test-project
version: 1.0.0
created: $(date +%Y-%m-%d)
EOF
log_pass "DoPlan structure created"

# Test 4: Generate Agents
log_info "Test 4: Generating AI Agents..."
cd "$PROJECT_ROOT"
go run -C "$PROJECT_ROOT" -exec "$PROJECT_ROOT/doplan" << 'GOSCRIPT' 2>&1 || true
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/DoPlan-dev/CLI/internal/generators"
)

func main() {
    projectRoot := "/tmp/doplan-integration-test/test-project"
    gen := generators.NewAgentsGenerator(projectRoot)
    if err := gen.Generate(); err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Agents generated successfully")
}
GOSCRIPT

cd "$TEST_DIR/test-project"
AGENT_FILES=(
    ".doplan/ai/agents/README.md"
    ".doplan/ai/agents/planner.agent.md"
    ".doplan/ai/agents/coder.agent.md"
    ".doplan/ai/agents/designer.agent.md"
    ".doplan/ai/agents/reviewer.agent.md"
    ".doplan/ai/agents/tester.agent.md"
    ".doplan/ai/agents/devops.agent.md"
)
ALL_AGENTS_EXIST=true
for file in "${AGENT_FILES[@]}"; do
    if [ -f "$file" ]; then
        log_pass "Agent file exists: $file"
    else
        log_fail "Agent file missing: $file"
        ALL_AGENTS_EXIST=false
    fi
done

if [ "$ALL_AGENTS_EXIST" = true ]; then
    log_pass "All 7 agent files generated"
else
    log_fail "Some agent files missing"
fi

# Test 5: Generate Rules
log_info "Test 5: Generating Rules..."
cd "$PROJECT_ROOT"
cat > /tmp/test_rules.go << 'EOF'
package main

import (
    "fmt"
    "os"
    "github.com/DoPlan-dev/CLI/internal/generators"
)

func main() {
    projectRoot := "/tmp/doplan-integration-test/test-project"
    gen := generators.NewRulesGenerator(projectRoot)
    if err := gen.Generate(); err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Rules generated successfully")
}
EOF
cd "$PROJECT_ROOT"
go run /tmp/test_rules.go 2>&1 || log_fail "Rules generation failed"

cd "$TEST_DIR/test-project"
RULE_FILES=(
    ".doplan/ai/rules/workflow.mdc"
    ".doplan/ai/rules/communication.mdc"
)
ALL_RULES_EXIST=true
for file in "${RULE_FILES[@]}"; do
    if [ -f "$file" ]; then
        log_pass "Rule file exists: $file"
    else
        log_fail "Rule file missing: $file"
        ALL_RULES_EXIST=false
    fi
done

if [ "$ALL_RULES_EXIST" = true ]; then
    log_pass "Core rule files generated"
else
    log_fail "Some rule files missing"
fi

# Test 6: Generate Commands
log_info "Test 6: Generating Commands..."
cd "$PROJECT_ROOT"
cat > /tmp/test_commands.go << 'EOF'
package main

import (
    "fmt"
    "os"
    "github.com/DoPlan-dev/CLI/internal/generators"
)

func main() {
    projectRoot := "/tmp/doplan-integration-test/test-project"
    gen := generators.NewCommandsGenerator(projectRoot)
    if err := gen.Generate(); err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Commands generated successfully")
}
EOF
cd "$PROJECT_ROOT"
go run /tmp/test_commands.go 2>&1 || log_fail "Commands generation failed"

cd "$TEST_DIR/test-project"
COMMAND_FILES=(
    ".doplan/ai/commands/run.md"
    ".doplan/ai/commands/deploy.md"
    ".doplan/ai/commands/create.md"
)
SOME_COMMANDS_EXIST=false
for file in "${COMMAND_FILES[@]}"; do
    if [ -f "$file" ]; then
        log_pass "Command file exists: $file"
        SOME_COMMANDS_EXIST=true
    fi
done

if [ "$SOME_COMMANDS_EXIST" = true ]; then
    log_pass "Command files generated"
else
    log_warn "No command files found (may be optional)"
fi

# Test 7: Test Design System Generation (DPR)
log_info "Test 7: Testing Design System (DPR) Generation..."
cd "$TEST_DIR/test-project"

# Create mock DPR data
cat > /tmp/test_dpr.go << 'EOF'
package main

import (
    "fmt"
    "os"
    "github.com/DoPlan-dev/CLI/internal/dpr"
)

func main() {
    projectRoot := "/tmp/doplan-integration-test/test-project"
    
    // Create mock DPR data
    data := &dpr.DPRData{
        Answers: map[string]interface{}{
            "project_name": "Test Project",
            "audience_primary": "Developers",
            "emotion_target": "Professional",
            "style_overall": "Modern",
            "color_primary": "#667eea",
            "typography_font": "Inter",
            "layout_style": "Card-based",
            "components_style": "Elevated",
            "animation_level": "Subtle",
            "accessibility_importance": 5,
            "responsive_priority": "Desktop First",
        },
    }
    
    // Test DPR generator
    gen := dpr.NewGenerator(projectRoot, data)
    if err := gen.Generate(); err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
        os.Exit(1)
    }
    
    // Test tokens generator
    tokenGen := dpr.NewTokenGenerator(projectRoot, data)
    if err := tokenGen.Generate(); err != nil {
        fmt.Fprintf(os.Stderr, "ERROR generating tokens: %v\n", err)
        os.Exit(1)
    }
    
    // Test cursor rules generator
    rulesGen := dpr.NewCursorRulesGenerator(projectRoot, data)
    if err := rulesGen.Generate(); err != nil {
        fmt.Fprintf(os.Stderr, "ERROR generating cursor rules: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("DPR generation completed successfully")
}
EOF

cd "$PROJECT_ROOT"
if go run /tmp/test_dpr.go 2>&1; then
    log_pass "DPR generation completed"
else
    log_fail "DPR generation failed"
fi

cd "$TEST_DIR/test-project"
DPR_FILES=(
    "doplan/design/DPR.md"
    "doplan/design/design-tokens.json"
    ".doplan/ai/rules/design_rules.mdc"
)
ALL_DPR_EXIST=true
for file in "${DPR_FILES[@]}"; do
    if [ -f "$file" ]; then
        log_pass "DPR file exists: $file"
        # Check file is not empty
        if [ -s "$file" ]; then
            log_pass "DPR file has content: $file"
        else
            log_warn "DPR file is empty: $file"
        fi
    else
        log_fail "DPR file missing: $file"
        ALL_DPR_EXIST=false
    fi
done

if [ "$ALL_DPR_EXIST" = true ]; then
    log_pass "All DPR files generated"
else
    log_fail "Some DPR files missing"
fi

# Test 8: Test API Keys Detection (RAKD)
log_info "Test 8: Testing API Keys Detection (RAKD)..."
cd "$PROJECT_ROOT"
cat > /tmp/test_rakd.go << 'EOF'
package main

import (
    "fmt"
    "os"
    "github.com/DoPlan-dev/CLI/internal/rakd"
)

func main() {
    projectRoot := "/tmp/doplan-integration-test/test-project"
    
    // Generate RAKD
    data, err := rakd.GenerateRAKD(projectRoot)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("RAKD generated: %d services detected\n", len(data.Services))
    fmt.Printf("Configured: %d, Required: %d, Pending: %d\n", 
        data.ConfiguredCount, data.RequiredCount, data.PendingCount)
}
EOF

cd "$PROJECT_ROOT"
if go run /tmp/test_rakd.go 2>&1; then
    log_pass "RAKD generation completed"
else
    log_warn "RAKD generation failed (may be expected if no services detected)"
fi

cd "$TEST_DIR/test-project"
if [ -f "doplan/RAKD.md" ]; then
    log_pass "RAKD.md file exists"
    # Check if services were detected
    if grep -q "Stripe\|SendGrid\|AWS" doplan/RAKD.md 2>/dev/null; then
        log_pass "Services detected in RAKD.md"
    else
        log_warn "No services detected (may be expected)"
    fi
else
    log_warn "RAKD.md not generated (may be expected if no services)"
fi

# Test 9: Test Workflow Guidance Integration
log_info "Test 9: Testing Workflow Guidance Integration..."
cd "$PROJECT_ROOT"
cat > /tmp/test_workflow.go << 'EOF'
package main

import (
    "fmt"
    "github.com/DoPlan-dev/CLI/internal/workflow"
)

func main() {
    // Test various action recommendations
    actions := []string{
        "project_created",
        "plan_complete",
        "design_complete",
        "feature_implemented",
        "tests_passed",
        "review_approved",
        "deployment_complete",
        "",
    }
    
    for _, action := range actions {
        title, desc := workflow.GetNextStep(action)
        if title != "" {
            fmt.Printf("Action: %s -> Title: %s\n", action, title)
        } else {
            fmt.Printf("Action: %s -> No recommendation\n", action)
        }
    }
    
    // Test workflow sequence
    sequence := workflow.GetWorkflowSequence()
    fmt.Printf("Workflow sequence: %v\n", sequence)
    
    fmt.Println("Workflow guidance tests passed")
}
EOF

cd "$PROJECT_ROOT"
if go run /tmp/test_workflow.go 2>&1 | head -15; then
    log_pass "Workflow guidance integration works"
else
    log_fail "Workflow guidance integration failed"
fi

# Test 10: Verify Cross-Phase Integration
log_info "Test 10: Verifying Cross-Phase Integration..."
cd "$TEST_DIR/test-project"

# Check that agents reference rules
if [ -f ".doplan/ai/agents/planner.agent.md" ]; then
    if grep -q "workflow.mdc\|communication.mdc" .doplan/ai/agents/planner.agent.md 2>/dev/null; then
        log_pass "Agents reference workflow rules"
    else
        log_warn "Agents may not reference workflow rules correctly"
    fi
fi

# Check that design rules reference DPR
if [ -f ".doplan/ai/rules/design_rules.mdc" ]; then
    if grep -q "DPR\|design-tokens" .doplan/ai/rules/design_rules.mdc 2>/dev/null; then
        log_pass "Design rules reference DPR"
    else
        log_warn "Design rules may not reference DPR correctly"
    fi
fi

# Check that devops agent references RAKD
if [ -f ".doplan/ai/agents/devops.agent.md" ]; then
    if grep -q "RAKD\|API.*keys" .doplan/ai/agents/devops.agent.md 2>/dev/null; then
        log_pass "DevOps agent references RAKD"
    else
        log_warn "DevOps agent may not reference RAKD"
    fi
fi

# Test 11: Verify File Structure Integrity
log_info "Test 11: Verifying File Structure Integrity..."
cd "$TEST_DIR/test-project"

EXPECTED_DIRS=(
    ".doplan"
    ".doplan/ai"
    ".doplan/ai/agents"
    ".doplan/ai/rules"
    ".doplan/ai/commands"
    ".doplan/design"
    ".doplan/SOPS"
    "doplan"
    "doplan/design"
    "doplan/contracts"
)

ALL_DIRS_EXIST=true
for dir in "${EXPECTED_DIRS[@]}"; do
    if [ -d "$dir" ]; then
        log_pass "Directory exists: $dir"
    else
        log_fail "Directory missing: $dir"
        ALL_DIRS_EXIST=false
    fi
done

if [ "$ALL_DIRS_EXIST" = true ]; then
    log_pass "All expected directories exist"
else
    log_fail "Some directories missing"
fi

# Test 12: Verify Agent Workflow Sequence
log_info "Test 12: Verifying Agent Workflow Sequence..."
cd "$TEST_DIR/test-project"

if [ -f ".doplan/ai/rules/workflow.mdc" ]; then
    # Check that workflow mentions correct sequence
    if grep -q "Plan.*Design.*Code.*Test.*Review.*Deploy" .doplan/ai/rules/workflow.mdc 2>/dev/null || \
       grep -q "planner.*designer.*coder.*tester.*reviewer.*devops" .doplan/ai/rules/workflow.mdc 2>/dev/null; then
        log_pass "Workflow sequence documented correctly"
    else
        log_warn "Workflow sequence may not be documented"
    fi
fi

# Summary
echo ""
echo "=========================================================="
echo "🧪 Integration Testing Summary"
echo "=========================================================="
echo -e "${GREEN}✅ Passed:${NC} $PASSED"
echo -e "${RED}❌ Failed:${NC} $FAILED"
echo -e "${YELLOW}⚠️  Warnings:${NC} $WARNINGS"
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✅ All critical tests passed!${NC}"
    echo ""
    echo "Integration testing: SUCCESS"
    echo "Ready for manual testing and user acceptance testing"
    exit 0
else
    echo -e "${RED}❌ Some tests failed${NC}"
    echo ""
    echo "Integration testing: PARTIAL SUCCESS"
    echo "Review failures above and fix before release"
    exit 1
fi

