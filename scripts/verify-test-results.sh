#!/bin/bash
# Verification script for manual test results
# Run this after completing manual tests to verify file structures

set -e

echo "🔍 Verifying Manual Test Results"
echo "=================================="
echo ""

TEST_DIR="/tmp/doplan-manual-test"
ERRORS=0

check_file() {
    if [ -f "$1" ] || [ -d "$1" ]; then
        echo "  ✅ $1 exists"
        return 0
    else
        echo "  ❌ $1 MISSING"
        ((ERRORS++))
        return 1
    fi
}

check_symlink() {
    if [ -L "$1" ]; then
        TARGET=$(readlink "$1")
        echo "  ✅ $1 → $TARGET"
        return 0
    else
        echo "  ❌ $1 is not a symlink"
        ((ERRORS++))
        return 1
    fi
}

echo "📁 Test Suite 2: New Project Wizard"
echo "-----------------------------------"
if [ -d "$TEST_DIR/empty" ]; then
    cd "$TEST_DIR/empty"
    check_file ".doplan/config.yaml"
    check_file "doplan/"
    echo ""
fi

echo "📁 Test Suite 3: Project Adoption"
echo "-----------------------------------"
if [ -d "$TEST_DIR/existing" ]; then
    cd "$TEST_DIR/existing"
    check_file ".cursor/config/doplan-state.json"
    check_file ".doplan/config.yaml"
    echo ""
fi

echo "📁 Test Suite 4: Migration"
echo "-----------------------------------"
if [ -d "$TEST_DIR/old" ]; then
    cd "$TEST_DIR/old"
    check_file ".doplan/config.yaml"
    check_file ".doplan/backup/"
    echo ""
fi

echo "📁 Test Suite 6: IDE Integration"
echo "-----------------------------------"
if [ -d "$TEST_DIR/new" ]; then
    cd "$TEST_DIR/new"
    
    # Cursor integration
    if [ -d ".cursor" ]; then
        echo "Cursor Integration:"
        check_symlink ".cursor/agents" || check_file ".cursor/agents"
        check_symlink ".cursor/rules" || check_file ".cursor/rules"
        check_symlink ".cursor/commands" || check_file ".cursor/commands"
    fi
    
    # VS Code integration
    if [ -d ".vscode" ]; then
        echo "VS Code Integration:"
        check_file ".vscode/tasks.json"
        check_file ".vscode/settings.json"
        check_file ".vscode/prompts/"
    fi
    
    # Generic IDE
    check_file ".doplan/guides/generic_ide_setup.md"
    echo ""
fi

echo "📊 Summary"
echo "=================================="
if [ $ERRORS -eq 0 ]; then
    echo "✅ All file structures verified"
    exit 0
else
    echo "❌ Found $ERRORS issues"
    exit 1
fi

