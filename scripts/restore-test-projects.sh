#!/bin/bash
# restore-test-projects.sh - Restore test projects from backup

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <timestamp>"
    echo "Available backups:"
    ls -1 /tmp/doplan-test-backups/ 2>/dev/null || echo "No backups found"
    exit 1
fi

TIMESTAMP="$1"
TEST_DIR="/tmp/doplan-test"
BACKUP_DIR="/tmp/doplan-test-backups/$TIMESTAMP"

if [ ! -d "$BACKUP_DIR" ]; then
    echo "❌ Backup not found: $BACKUP_DIR"
    exit 1
fi

echo "Restoring test projects from backup: $TIMESTAMP"

for project in empty existing old new; do
    SRC="$BACKUP_DIR/$project"
    DST="$TEST_DIR/$project"
    
    if [ -d "$SRC" ]; then
        echo "Restoring $project..."
        rm -rf "$DST"
        cp -r "$SRC" "$DST"
        echo "✓ $project restored"
    else
        echo "⚠️  $project not in backup, skipping"
    fi
done

echo ""
echo "✓ Test projects restored"

