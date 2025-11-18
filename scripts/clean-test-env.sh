#!/bin/bash
# clean-test-env.sh - Clean up test environment

set -e

TEST_DIR="/tmp/doplan-test"
BACKUP_DIR="/tmp/doplan-test-backups"

read -p "This will delete all test projects. Are you sure? (y/N) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Cancelled"
    exit 0
fi

echo "Cleaning test environment..."

if [ -d "$TEST_DIR" ]; then
    rm -rf "$TEST_DIR"
    echo "✓ Test directory removed"
fi

if [ -d "$BACKUP_DIR" ]; then
    read -p "Also delete backups? (y/N) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -rf "$BACKUP_DIR"
        echo "✓ Backups removed"
    fi
fi

echo "✓ Cleanup complete"

