#!/bin/bash
# test-migration.sh - Test migration functionality

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEST_DIR="/tmp/doplan-test"
OLD_DIR="$TEST_DIR/old"

echo "Testing migration functionality..."

if [ ! -d "$OLD_DIR" ]; then
    echo "❌ Old DoPlan structure not found. Run setup-test-env.sh first."
    exit 1
fi

cd "$OLD_DIR"

echo "Testing migration from old structure..."
echo "Current structure:"
ls -la .cursor/config/ 2>/dev/null || echo "No .cursor/config"
ls -la doplan/ 2>/dev/null || echo "No doplan/"

# Run migration (when implemented)
# doplan migrate

echo "✓ Migration test complete"

