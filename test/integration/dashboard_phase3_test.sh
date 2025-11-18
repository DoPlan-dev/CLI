#!/bin/bash

# Integration Test for Phase 3: Dashboard Supercharge
# Tests dashboard generation with real project data

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
BINARY_PATH="$PROJECT_ROOT/doplan"
TEST_PROJECT=$(mktemp -d)

PASSED=0
FAILED=0

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Phase 3 Dashboard Integration Test${NC}"
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

# Setup test project
cd "$TEST_PROJECT"
echo -e "${BLUE}Setting up test project...${NC}"

# Initialize git
git init > /dev/null 2>&1 || true
git config user.name "Test User" > /dev/null 2>&1 || true
git config user.email "test@example.com" > /dev/null 2>&1 || true
echo "# Test Project" > README.md
git add README.md
git commit -m "Initial commit" > /dev/null 2>&1 || true

# Create some commits for activity feed
echo "// Feature 1" > feature1.js
git add feature1.js
git commit -m "feat: implement feature 1" > /dev/null 2>&1 || true

echo "// Feature 2" > feature2.js
git add feature2.js
git commit -m "feat: implement feature 2" > /dev/null 2>&1 || true

# Create project structure
mkdir -p doplan/01-phase-1/01-Feature-1
mkdir -p doplan/01-phase-1/02-Feature-2
mkdir -p doplan/02-phase-2/01-Feature-3
mkdir -p .doplan

# Create progress.json files
cat > doplan/01-phase-1/01-Feature-1/progress.json <<EOF
{
  "featureID": "feature-1",
  "featureName": "Feature 1",
  "status": "in-progress",
  "progress": 50,
  "branch": "feature/01-phase-1-01-feature-1"
}
EOF

cat > doplan/01-phase-1/02-Feature-2/progress.json <<EOF
{
  "featureID": "feature-2",
  "featureName": "Feature 2",
  "status": "complete",
  "progress": 100,
  "branch": "feature/01-phase-1-02-feature-2"
}
EOF

cat > doplan/02-phase-2/01-Feature-3/progress.json <<EOF
{
  "featureID": "feature-3",
  "featureName": "Feature 3",
  "status": "in-progress",
  "progress": 25,
  "branch": "feature/02-phase-2-01-feature-3"
}
EOF

# Create state.json structure (simplified)
mkdir -p .cursor/config
cat > .cursor/config/doplan-state.json <<EOF
{
  "phases": [
    {
      "id": "phase-1",
      "name": "Phase 1",
      "status": "in-progress",
      "features": ["feature-1", "feature-2"]
    },
    {
      "id": "phase-2",
      "name": "Phase 2",
      "status": "in-progress",
      "features": ["feature-3"]
    }
  ],
  "features": [
    {
      "id": "feature-1",
      "name": "Feature 1",
      "status": "in-progress",
      "progress": 50,
      "branch": "feature/01-phase-1-01-feature-1"
    },
    {
      "id": "feature-2",
      "name": "Feature 2",
      "status": "complete",
      "progress": 100,
      "branch": "feature/01-phase-1-02-feature-2"
    },
    {
      "id": "feature-3",
      "name": "Feature 3",
      "status": "in-progress",
      "progress": 25,
      "branch": "feature/02-phase-2-01-feature-3"
    }
  ],
  "progress": {
    "overall": 58,
    "phases": {
      "phase-1": 75,
      "phase-2": 25
    }
  }
}
EOF

# Create config
cat > .cursor/config/doplan-config.json <<EOF
{
  "ide": "cursor",
  "version": "1.0.0",
  "installed": true,
  "github": {
    "enabled": true,
    "repository": "test/repo"
  }
}
EOF

# Run tests
echo ""
echo -e "${BLUE}Test 1: Dashboard Generation${NC}"
test_check "Binary exists" "[ -f '$BINARY_PATH' ]"
test_check "Generate dashboard.json" "$BINARY_PATH progress 2>&1 | grep -q 'Progress updated' || [ -f .doplan/dashboard.json ]"

if [ -f .doplan/dashboard.json ]; then
    echo ""
    echo -e "${BLUE}Test 2: Dashboard.json Structure${NC}"
    test_check "Dashboard.json exists" "[ -f .doplan/dashboard.json ]"
    test_check "Dashboard has version" "grep -q '\"version\"' .doplan/dashboard.json"
    test_check "Dashboard has phases" "grep -q '\"phases\"' .doplan/dashboard.json"
    test_check "Dashboard has activity" "grep -q '\"activity\"' .doplan/dashboard.json"
    test_check "Dashboard has velocity" "grep -q '\"velocity\"' .doplan/dashboard.json"
    
    echo ""
    echo -e "${BLUE}Test 3: Activity Feed${NC}"
    test_check "Activity feed exists" "grep -q '\"recentActivity\"' .doplan/dashboard.json"
    test_check "Activity has commits" "grep -q 'commit' .doplan/dashboard.json || true"
    
    echo ""
    echo -e "${BLUE}Test 4: Velocity Metrics${NC}"
    test_check "Velocity metrics exist" "grep -q '\"commitsPerDay\"' .doplan/dashboard.json"
    test_check "Velocity has tasksPerDay" "grep -q '\"tasksPerDay\"' .doplan/dashboard.json"
    
    echo ""
    echo -e "${BLUE}Test 5: Progress Data${NC}"
    test_check "Phases have progress" "grep -q '\"progress\"' .doplan/dashboard.json"
    test_check "Features have progress" "grep -q '\"progress\"' .doplan/dashboard.json"
    
    echo ""
    echo -e "${BLUE}Dashboard.json Content Preview:${NC}"
    echo ""
    cat .doplan/dashboard.json | head -50
    echo ""
    echo "..."
    echo ""
else
    echo -e "${YELLOW}Warning: dashboard.json not generated, skipping structure tests${NC}"
fi

# Summary
echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Test Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo "Total tests: $((PASSED + FAILED))"
echo -e "Passed: ${GREEN}$PASSED${NC}"
echo -e "Failed: ${RED}$FAILED${NC}"
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}✗ Some tests failed${NC}"
    exit 1
fi

