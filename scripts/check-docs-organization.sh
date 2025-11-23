#!/bin/bash
# Script to check documentation organization
# Ensures root stays clean and docs are organized

set -e

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ERRORS=0

echo "🔍 Checking documentation organization..."

# Check for .md files in root (except README.md and CHANGELOG.md)
ROOT_MD_FILES=$(find "$ROOT_DIR" -maxdepth 1 -name "*.md" ! -name "README.md" ! -name "CHANGELOG.md" -type f)

if [ -n "$ROOT_MD_FILES" ]; then
    echo "❌ Error: Found .md files in root directory (should be in docs/):"
    echo "$ROOT_MD_FILES"
    ERRORS=$((ERRORS + 1))
fi

# Check that docs/README.md exists
if [ ! -f "$ROOT_DIR/docs/README.md" ]; then
    echo "⚠️  Warning: docs/README.md not found"
fi

# Check for build artifacts in root
if [ -f "$ROOT_DIR/doplan" ] && ! git check-ignore -q "$ROOT_DIR/doplan"; then
    echo "⚠️  Warning: doplan binary found in root (should be in .gitignore)"
fi

if [ -f "$ROOT_DIR/coverage.out" ] && ! git check-ignore -q "$ROOT_DIR/coverage.out"; then
    echo "⚠️  Warning: coverage.out found in root (should be in .gitignore)"
fi

if [ $ERRORS -eq 0 ]; then
    echo "✅ Documentation organization looks good!"
    exit 0
else
    echo "❌ Found $ERRORS error(s). Please fix before committing."
    exit 1
fi

