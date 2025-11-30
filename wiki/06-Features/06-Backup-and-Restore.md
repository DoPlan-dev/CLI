# Backup and Restore

DoPlan provides comprehensive backup and restore capabilities to protect your work and enable safe project migrations.

---

## 🎯 Overview

DoPlan's backup system allows you to:
- **Create backups** with flexible options (project files, planning data, or everything)
- **Restore projects** from compressed backups
- **Export/import memory cards** for transferring preferences between projects
- **Migrate safely** when updating DoPlan versions
- **Protect your work** with automatic safety backups before restore

---

## 📦 Backup Types

### 1. Project Backup (`project`)

**What it backs up:**
- All project files (source code, configs, etc.)
- Excludes `.do/` and `.cursor/` directories

**When to use:**
- Before DoPlan updates
- When you want to preserve only your work
- For sharing project code without DoPlan structure

**Example:**
```bash
/sys backup --type project --description "Before DoPlan update"
```

### 2. Plan Backup (`plan`)

**What it backs up:**
- Only `.do/plan/` directory
- Task lists, feature documentation, planning files

**When to use:**
- Before major planning changes
- To preserve task structure
- For planning data recovery

**Example:**
```bash
/sys backup --type plan --description "Before restructuring tasks"
```

### 3. Project + Plan Backup (`project-plan`)

**What it backs up:**
- Project files (source code, configs)
- `.do/plan/` directory
- Excludes `.do/core/`, `.do/system/`, `.cursor/`

**When to use:**
- Before DoPlan migrations
- When you want work + planning but not system files
- For project transfers

**Example:**
```bash
/sys backup --type project-plan --description "Migration backup"
```

### 4. Full Backup (`full`)

**What it backs up:**
- Everything in the project
- All `.do/` directories
- All `.cursor/` directories
- Complete project state

**When to use:**
- Complete project backup
- Before major changes
- For disaster recovery
- When you want everything preserved

**Example:**
```bash
/sys backup --type full --description "Complete backup before refactoring"
```

---

## 💾 Creating Backups

### Interactive Mode

```bash
/sys backup
```

**What happens:**
1. Prompts you to select backup type
2. Optionally asks for description
3. Creates compressed backup in `.do/backup/`
4. Shows backup information

**Output:**
```
✅ Backup created: backup-20250115-143022-project.doplan
   Location: /path/to/project/.do/backup/backup-20250115-143022-project.doplan
   Type: project
   Size: 2.45 MB
   Description: Before DoPlan update
```

### Command Line Mode

```bash
# With type and description
/sys backup --type project --description "Before update"

# Just type
/sys backup --type full

# Interactive selection
/sys backup
```

### Backup File Format

**Location:** `.do/backup/`

**Naming:** `backup-[timestamp]-[type].doplan`

**Format:** Compressed tar.gz archive with manifest

**Contents:**
- `BACKUP_MANIFEST.json` - Backup metadata
- All selected files and directories

---

## 🔄 Restoring Backups

### Interactive Restore

```bash
/sys restore
```

**What happens:**
1. Lists available backups
2. Shows backup information (type, size, date, description)
3. Verifies backup integrity
4. Checks version compatibility
5. Creates safety backup of current state
6. Restores files
7. Shows restore results

### Restore by Filename

```bash
/sys restore backup-20250115-143022-project.doplan
```

Or just the base name:
```bash
/sys restore backup-20250115-143022-project
```

### Dry Run Mode

Preview what would be restored without actually restoring:

```bash
/sys restore --dry-run backup-20250115-143022-project.doplan
```

**Shows:**
- Files that would be restored
- Files that would be skipped
- Conflicts that would occur

---

## 🔍 Backup Information

### Backup Manifest

Each backup contains a manifest with:

```json
{
  "backup_type": "project",
  "timestamp": "2025-01-15T14:30:22Z",
  "doplan_version": "1.0.0",
  "description": "Before DoPlan update",
  "file_count": 127,
  "total_size": 2560000
}
```

### Viewing Backup Info

When you restore, DoPlan shows:
- Backup file name
- Backup type
- Creation timestamp
- Description (if provided)
- DoPlan version
- File count

---

## ✅ Backup Verification

### Integrity Check

Before restore, DoPlan automatically:
1. **Verifies backup integrity** - Ensures backup file is valid
2. **Checks version compatibility** - Warns if versions differ significantly
3. **Validates manifest** - Ensures backup metadata is correct

### Version Compatibility

DoPlan checks if backup version is compatible:
- **Compatible**: Same or compatible version → Proceeds
- **Warning**: Different version → Shows warning, asks for confirmation
- **Incompatible**: Major version difference → Warns strongly

---

## 🛡️ Safety Features

### Automatic Safety Backup

Before restoring, DoPlan automatically creates a safety backup of your current state:

```
💾 Creating safety backup of current state...
✅ Safety backup created: backup-20250115-150000-full.doplan
```

**What's backed up:**
- Complete project state (full backup)
- Everything in case you need to rollback

**When created:**
- Automatically before every restore operation
- Not created in dry-run mode

### Conflict Handling

When restoring, DoPlan handles conflicts:

- **Existing files**: Detected and reported
- **Merge strategy**: Can overwrite, merge, or skip
- **Warnings**: All conflicts shown in restore results

---

## 📊 Restore Results

After restore, DoPlan shows:

```
✅ Restore Complete!
   Files restored: 127
   Files skipped: 3
   Files merged: 2
   Conflicts: 0

⚠️  Warnings:
   • Version mismatch detected
   • Some files were skipped due to conflicts
```

### Result Information

- **Files restored**: Number of files successfully restored
- **Files skipped**: Files that were skipped (conflicts, etc.)
- **Files merged**: Files that were merged intelligently
- **Conflicts**: Number of file conflicts encountered
- **Warnings**: Any warnings or issues during restore

---

## 🧠 Memory Card Backup

### Export Memory Card

Export your memory card (user preferences, engagement data) for use in other projects:

```bash
/sys memory export
```

**What's exported:**
- User preferences and work style
- Command usage statistics
- Achievements and challenges
- Engagement metrics
- Project history and context

**Location:** `.do/backup/memory_card_export-[timestamp].json`

### Restore Memory Card

Restore a memory card export in your current project:

```bash
/sys memory restore
```

**What happens:**
1. Lists available memory card exports
2. Backs up current memory card (if exists)
3. Restores selected memory card
4. Confirms restoration

**Use cases:**
- Transfer preferences to new project
- Restore engagement data
- Share user profile between projects

---

## 🔄 Migration Assistant

### Guided Migration

When updating DoPlan, use the migration assistant:

```bash
/sys migrate
```

**What it does:**
1. Detects old DoPlan structure
2. Suggests appropriate backup type
3. Guides you through update process
4. Helps restore after update

### Migration Steps

The assistant provides:
- **Detection**: Identifies old structure indicators
- **Recommendation**: Suggests best backup type
- **Guidance**: Step-by-step instructions
- **Safety**: Ensures your work is protected

---

## 📁 Backup Storage

### Location

All backups stored in:
```
.do/backup/
```

### File Management

- **Automatic organization**: All backups in one directory
- **Easy listing**: Use `/sys restore` to see all backups
- **Manual access**: Can access files directly if needed

### Backup Retention

- **No automatic deletion**: All backups kept
- **Manual cleanup**: Remove old backups manually if needed
- **Storage management**: Monitor `.do/backup/` directory size

---

## 🎯 Use Cases

### 1. Before DoPlan Update

```bash
# Create backup
/sys backup --type project-plan --description "Before DoPlan update"

# Update DoPlan
# ... update process ...

# Restore if needed
/sys restore backup-20250115-143022-project-plan.doplan
```

### 2. Before Major Changes

```bash
# Full backup before refactoring
/sys backup --type full --description "Before major refactoring"

# Make changes
# ... work ...

# Restore if something goes wrong
/sys restore backup-20250115-143022-full.doplan
```

### 3. Project Transfer

```bash
# Create full backup
/sys backup --type full --description "Project transfer"

# Copy backup file to new location
# Restore in new location
/sys restore backup-20250115-143022-full.doplan
```

### 4. Planning Recovery

```bash
# Backup planning data
/sys backup --type plan --description "Before task restructuring"

# Restore if needed
/sys restore backup-20250115-143022-plan.doplan
```

### 5. Memory Card Transfer

```bash
# In old project
/sys memory export

# In new project
/sys memory restore
```

---

## 💡 Best Practices

### Regular Backups

- **Before updates**: Always backup before DoPlan updates
- **Before major changes**: Backup before refactoring or restructuring
- **Periodic backups**: Create periodic full backups
- **Planning backups**: Backup planning data before major task changes

### Backup Types

- **Project updates**: Use `project` or `project-plan`
- **Planning changes**: Use `plan`
- **Major changes**: Use `full`
- **Migrations**: Use `project-plan` or `full`

### Restore Safety

- **Dry run first**: Use `--dry-run` to preview
- **Safety backups**: Automatic safety backups protect you
- **Version checks**: Pay attention to version compatibility warnings
- **Review results**: Check restore results for warnings

### Storage Management

- **Monitor size**: Keep an eye on `.do/backup/` directory size
- **Cleanup old backups**: Remove old backups when no longer needed
- **Keep important backups**: Don't delete backups you might need

---

## 🔧 Advanced Features

### Backup Manifest

Each backup includes a manifest with:
- Backup type
- Timestamp
- DoPlan version
- Description
- File count
- Total size

### Version Compatibility

DoPlan checks version compatibility:
- **Same version**: Fully compatible
- **Compatible versions**: Works with warnings
- **Incompatible versions**: Strong warnings, may not work

### Merge Strategies

When restoring, you can choose merge strategies:
- **Overwrite**: Replace existing files (default)
- **Merge**: Intelligently merge files (for `.do/plan/`)
- **Skip**: Skip existing files

---

## 🚨 Troubleshooting

### Backup Creation Fails

**Issue**: Cannot create backup

**Solutions:**
- Check disk space
- Verify write permissions
- Check `.do/backup/` directory exists
- Try different backup type

### Restore Fails

**Issue**: Restore operation fails

**Solutions:**
- Verify backup file integrity
- Check version compatibility
- Ensure sufficient disk space
- Try dry-run first
- Check file permissions

### Version Warnings

**Issue**: Version compatibility warnings

**Solutions:**
- Review warning message
- Check DoPlan version compatibility
- Consider updating DoPlan
- Proceed with caution if incompatible

### Missing Backups

**Issue**: Backups not found

**Solutions:**
- Check `.do/backup/` directory exists
- Verify backup file names
- Check file permissions
- List backups with `/sys restore`

---

## 📚 Related Topics

- [State Management](./02-State-Management.md) - Project state tracking
- [System Commands](../02-Commands/03-System-Commands.md) - All system commands
- [Troubleshooting](../08-Advanced/04-Troubleshooting.md) - Common issues

---

**Next**: [Learning Support](./04-Learning-Support.md)

