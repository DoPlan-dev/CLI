#!/bin/bash
# validate-migration.sh - Validate a completed migration

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="${1:-$(pwd)}"

echo "Validating migration in: $PROJECT_ROOT"
echo ""

cd "$PROJECT_ROOT"

ERRORS=0

# Check new config exists
if [ -f ".doplan/config.yaml" ]; then
    echo "✓ New config exists"
    
    # Validate YAML syntax
    if command -v yq &> /dev/null; then
        yq eval '.' .doplan/config.yaml > /dev/null 2>&1 && echo "✓ Config YAML is valid" || {
            echo "❌ Config YAML is invalid"
            ERRORS=$((ERRORS + 1))
        }
    fi
else
    echo "❌ New config not found"
    ERRORS=$((ERRORS + 1))
fi

# Check old config is gone or backed up
if [ -f ".cursor/config/doplan-config.json" ]; then
    echo "⚠️  Old config still exists (should be backed up)"
else
    echo "✓ Old config removed or backed up"
fi

# Check new folder structure
if [ -d "doplan" ]; then
    echo "✓ doplan directory exists"
    
    # Check for old-style folders
    OLD_FOLDERS=$(find doplan -maxdepth 1 -type d -name "*-phase" 2>/dev/null | wc -l)
    if [ "$OLD_FOLDERS" -gt 0 ]; then
        echo "❌ Old-style phase folders found: $OLD_FOLDERS"
        ERRORS=$((ERRORS + 1))
    else
        echo "✓ No old-style phase folders"
    fi
    
    # Check for new-style folders
    NEW_FOLDERS=$(find doplan -maxdepth 1 -type d -regex ".*/[0-9]+-.*" 2>/dev/null | wc -l)
    if [ "$NEW_FOLDERS" -gt 0 ]; then
        echo "✓ New-style folders found: $NEW_FOLDERS"
    fi
else
    echo "❌ doplan directory not found"
    ERRORS=$((ERRORS + 1))
fi

# Check .doplan structure
if [ -d ".doplan" ]; then
    echo "✓ .doplan directory exists"
    
    [ -d ".doplan/ai" ] && echo "✓ .doplan/ai exists" || echo "⚠️  .doplan/ai missing"
    [ -f ".doplan/dashboard.json" ] && echo "✓ dashboard.json exists" || echo "⚠️  dashboard.json missing"
else
    echo "❌ .doplan directory not found"
    ERRORS=$((ERRORS + 1))
fi

echo ""
if [ $ERRORS -eq 0 ]; then
    echo "✓ Migration validation passed!"
    exit 0
else
    echo "❌ Migration validation failed with $ERRORS error(s)"
    exit 1
fi

