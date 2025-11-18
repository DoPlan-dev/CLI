#!/bin/bash
# test-ide-integration.sh - Test IDE integration setup

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEST_DIR="/tmp/doplan-test"
NEW_DIR="$TEST_DIR/new"

echo "Testing IDE integration..."

if [ ! -d "$NEW_DIR" ]; then
    echo "❌ New DoPlan structure not found. Run setup-test-env.sh first."
    exit 1
fi

cd "$NEW_DIR"

# Test Cursor integration
echo "Testing Cursor integration..."
if [ -d ".cursor" ]; then
    echo "✓ .cursor directory exists"
    if [ -L ".cursor/agents" ] || [ -d ".cursor/agents" ]; then
        echo "✓ .cursor/agents exists"
    else
        echo "❌ .cursor/agents missing"
    fi
else
    echo "⚠️  .cursor directory not found (will be created on setup)"
fi

# Test VS Code integration
echo "Testing VS Code integration..."
if [ -d ".vscode" ]; then
    echo "✓ .vscode directory exists"
    if [ -f ".vscode/tasks.json" ]; then
        echo "✓ .vscode/tasks.json exists"
    fi
else
    echo "⚠️  .vscode directory not found (will be created on setup)"
fi

echo "✓ IDE integration test complete"

