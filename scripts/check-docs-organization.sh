#!/bin/bash
# Script to check documentation organization
# Ensures root stays clean and docs are organized according to rules
# See Docs/README.md for full guidelines
# Rule: .cursor/rules/library/11-documentation/docs-folder-structure.md

set -e

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ERRORS=0
WARNINGS=0

echo "🔍 Checking documentation organization..."
echo ""

# Check for .md files in root (except README.md and CHANGELOG.md)
ROOT_MD_FILES=$(find "$ROOT_DIR" -maxdepth 1 -name "*.md" ! -name "README.md" ! -name "CHANGELOG.md" -type f)

if [ -n "$ROOT_MD_FILES" ]; then
    echo "❌ Error: Found .md files in root directory (should be in Docs/):"
    echo "$ROOT_MD_FILES" | sed 's/^/   /'
    echo ""
    echo "   → Move to appropriate Docs/ subdirectory:"
    echo "     - Foundation docs → Docs/foundation/"
    echo "     - Feature docs → Docs/features/<Feature_Name>/"
    echo "     - Release docs → Docs/release/"
    echo "     - History docs → Docs/history/"
    echo ""
    echo "   See: .cursor/rules/library/11-documentation/docs-folder-structure.md"
    echo ""
    ERRORS=$((ERRORS + 1))
fi

# Check that Docs/ directory exists (capital D)
if [ ! -d "$ROOT_DIR/Docs" ]; then
    echo "⚠️  Warning: Docs/ directory not found (should be created by generator)"
    WARNINGS=$((WARNINGS + 1))
else
    # Check that Docs/README.md exists
    if [ ! -f "$ROOT_DIR/Docs/README.md" ]; then
        echo "⚠️  Warning: Docs/README.md not found"
        WARNINGS=$((WARNINGS + 1))
    fi
    
    # Check for canonical subdirectories
    for subdir in "foundation" "features" "release" "history"; do
        if [ ! -d "$ROOT_DIR/Docs/$subdir" ]; then
            echo "⚠️  Warning: Docs/$subdir/ directory not found"
            WARNINGS=$((WARNINGS + 1))
        fi
    done
fi

# Check for lowercase docs/ directory (should be Docs/)
if [ -d "$ROOT_DIR/docs" ] && [ ! -d "$ROOT_DIR/Docs" ]; then
    echo "⚠️  Warning: Found lowercase 'docs/' directory. Should be 'Docs/' (capital D)"
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
    echo "📖 Remember: All docs must live under Docs/ (capital D)"
    echo "   See: .cursor/rules/library/11-documentation/docs-folder-structure.md"
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo "⚠️  Found $WARNINGS warning(s) (non-blocking)"
    echo ""
    echo "📖 See Docs/README.md and .cursor/rules/library/11-documentation/docs-folder-structure.md"
    exit 0
else
    echo "❌ Found $ERRORS error(s) and $WARNINGS warning(s)."
    echo ""
    echo "📖 Documentation organization rules:"
    echo "   - .cursor/rules/library/11-documentation/docs-folder-structure.md"
    echo "   - Docs/README.md"
    echo ""
    echo "   Root must stay clean: only README.md and CHANGELOG.md allowed in root"
    exit 1
fi

