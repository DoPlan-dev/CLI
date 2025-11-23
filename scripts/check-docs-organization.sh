#!/bin/bash
# Script to check documentation organization
# Ensures root stays clean and docs are organized according to rules
# See docs/CONTRIBUTING.md for full guidelines

set -e

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ERRORS=0
WARNINGS=0

echo "🔍 Checking documentation organization..."
echo ""

# Check for .md files in root (except README.md and CHANGELOG.md)
ROOT_MD_FILES=$(find "$ROOT_DIR" -maxdepth 1 -name "*.md" ! -name "README.md" ! -name "CHANGELOG.md" -type f)

if [ -n "$ROOT_MD_FILES" ]; then
    echo "❌ Error: Found .md files in root directory (should be in docs/):"
    echo "$ROOT_MD_FILES" | sed 's/^/   /'
    echo ""
    echo "   → Move to appropriate docs/ subdirectory:"
    echo "     - Development docs → docs/development/"
    echo "     - Release docs → docs/release/"
    echo "     - Security docs → docs/security/"
    echo "     - General docs → docs/"
    echo ""
    ERRORS=$((ERRORS + 1))
fi

# Check that docs/README.md exists
if [ ! -f "$ROOT_DIR/docs/README.md" ]; then
    echo "⚠️  Warning: docs/README.md not found"
    WARNINGS=$((WARNINGS + 1))
fi

# Check that docs/CONTRIBUTING.md exists
if [ ! -f "$ROOT_DIR/docs/CONTRIBUTING.md" ]; then
    echo "⚠️  Warning: docs/CONTRIBUTING.md not found"
    WARNINGS=$((WARNINGS + 1))
fi

# Check for build artifacts in root
if [ -f "$ROOT_DIR/doplan" ] && ! git check-ignore -q "$ROOT_DIR/doplan"; then
    echo "⚠️  Warning: doplan binary found in root (should be in .gitignore)"
    WARNINGS=$((WARNINGS + 1))
fi

if [ -f "$ROOT_DIR/coverage.out" ] && ! git check-ignore -q "$ROOT_DIR/coverage.out"; then
    echo "⚠️  Warning: coverage.out found in root (should be in .gitignore)"
    WARNINGS=$((WARNINGS + 1))
fi

# Summary
echo ""
if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    echo "✅ Documentation organization looks good!"
    echo ""
    echo "📖 Remember: See docs/CONTRIBUTING.md for rules on where to create new docs"
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo "⚠️  Found $WARNINGS warning(s) (non-blocking)"
    echo ""
    echo "📖 See docs/CONTRIBUTING.md for documentation organization rules"
    exit 0
else
    echo "❌ Found $ERRORS error(s) and $WARNINGS warning(s)."
    echo ""
    echo "📖 See docs/CONTRIBUTING.md for documentation organization rules"
    exit 1
fi

