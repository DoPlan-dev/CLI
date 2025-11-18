#!/bin/bash
# backup-test-projects.sh - Create backups of test projects

set -e

TEST_DIR="/tmp/doplan-test"
BACKUP_DIR="/tmp/doplan-test-backups"
TIMESTAMP=$(date +%Y%m%d-%H%M%S)

echo "Backing up test projects..."

mkdir -p "$BACKUP_DIR/$TIMESTAMP"

for project in empty existing old new; do
    SRC="$TEST_DIR/$project"
    DST="$BACKUP_DIR/$TIMESTAMP/$project"
    
    if [ -d "$SRC" ]; then
        echo "Backing up $project..."
        cp -r "$SRC" "$DST"
        echo "✓ $project backed up"
    else
        echo "⚠️  $project not found, skipping"
    fi
done

echo ""
echo "✓ Backups created in: $BACKUP_DIR/$TIMESTAMP"
echo ""
echo "To restore:"
echo "  ./scripts/restore-test-projects.sh $TIMESTAMP"

