# DoPlan v0.0.18-beta Utility Scripts

This directory contains utility scripts for development, testing, and maintenance of DoPlan v0.0.18-beta.

## 📋 Scripts Overview

### Test Environment Scripts

#### `setup-test-env.sh`
Sets up a complete test environment with various project states.

**Usage:**
```bash
./scripts/setup-test-env.sh
```

**Creates:**
- `/tmp/doplan-test/empty` - Empty folder
- `/tmp/doplan-test/existing` - Existing code without DoPlan
- `/tmp/doplan-test/old` - Old DoPlan structure (v0.0.17)
- `/tmp/doplan-test/new` - New DoPlan structure (v0.0.18)

#### `clean-test-env.sh`
Cleans up the test environment.

**Usage:**
```bash
./scripts/clean-test-env.sh
```

#### `backup-test-projects.sh`
Creates backups of test projects.

**Usage:**
```bash
./scripts/backup-test-projects.sh
```

**Backups stored in:** `/tmp/doplan-test-backups/TIMESTAMP/`

#### `restore-test-projects.sh`
Restores test projects from backup.

**Usage:**
```bash
./scripts/restore-test-projects.sh <timestamp>
```

**Example:**
```bash
./scripts/restore-test-projects.sh 20240115-103000
```

### Testing Scripts

#### `run-all-tests.sh`
Runs all test suites (unit + integration).

**Usage:**
```bash
./scripts/run-all-tests.sh
```

**Runs:**
- Unit tests for all packages
- Integration tests
- Migration tests
- IDE integration tests

#### `test-migration.sh`
Tests migration functionality.

**Usage:**
```bash
./scripts/test-migration.sh
```

**Requires:** Old DoPlan structure (created by `setup-test-env.sh`)

#### `test-ide-integration.sh`
Tests IDE integration setup.

**Usage:**
```bash
./scripts/test-ide-integration.sh
```

**Checks:**
- Cursor integration
- VS Code integration
- Directory structure
- Symlinks/files

### Validation Scripts

#### `validate-migration.sh`
Validates a completed migration.

**Usage:**
```bash
./scripts/validate-migration.sh [project-root]
```

**Checks:**
- New config exists and is valid
- Old config removed/backed up
- Folder structure migrated
- .doplan structure exists

#### `check-dependencies.sh`
Checks required dependencies.

**Usage:**
```bash
./scripts/check-dependencies.sh
```

**Checks:**
- Go (required, 1.24+)
- Git (required)
- GitHub CLI (optional)
- jq (optional)
- yq (optional)

## 🔧 Development Workflow

### Initial Setup
```bash
# 1. Check dependencies
./scripts/check-dependencies.sh

# 2. Set up test environment
./scripts/setup-test-env.sh

# 3. Create backup before making changes
./scripts/backup-test-projects.sh
```

### During Development
```bash
# Run tests frequently
./scripts/run-all-tests.sh

# Test specific features
./scripts/test-migration.sh
./scripts/test-ide-integration.sh
```

### Before Committing
```bash
# Run all tests
./scripts/run-all-tests.sh

# Validate any migrations
./scripts/validate-migration.sh
```

## 📝 Script Details

### Test Environment Structure

After running `setup-test-env.sh`, you'll have:

```
/tmp/doplan-test/
├── empty/          # Empty folder (for new project wizard)
├── existing/       # Existing code (for adoption wizard)
├── old/            # Old DoPlan structure (for migration)
└── new/            # New DoPlan structure (for testing)
```

### Backup Structure

Backups are stored with timestamps:

```
/tmp/doplan-test-backups/
└── 20240115-103000/
    ├── empty/
    ├── existing/
    ├── old/
    └── new/
```

## 🐛 Troubleshooting

### Script fails with "permission denied"
```bash
chmod +x scripts/*.sh
```

### Test environment not found
```bash
./scripts/setup-test-env.sh
```

### Backup not found
```bash
# List available backups
ls -1 /tmp/doplan-test-backups/
```

## 📚 Related Documentation

- [Testing Scenarios](../docs/development/TESTING_SCENARIOS.md)
- [Troubleshooting Guide](../docs/development/TROUBLESHOOTING_GUIDE.md)
- [Migration Strategy](../docs/development/MIGRATION_STRATEGY.md)
