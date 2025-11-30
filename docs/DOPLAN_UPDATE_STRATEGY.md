# DoPlan System Update Strategy

## Overview
This document describes the recommended approach for updating DoPlan CLI system files while preserving your project work.

## The Challenge
When DoPlan CLI updates, the system files (`.do/core/`, `.do/system/`, `.cursor/`) may change. We need a way to:
- Update DoPlan system files to get new features/bug fixes
- Preserve your project work (code, planning, tasks)
- Avoid conflicts between old and new system structures

## Solution: Granular Backup & Restore

### Backup Types

DoPlan supports 4 backup types for different scenarios:

#### 1. Project Backup (`--type project`)
**What it backs up:**
- All project files (src/, docs/, config files, package.json, etc.)
- **Excludes:** `.do/`, `.cursor/`, `node_modules/`, `.git/`

**Use case:** Updating DoPlan system while preserving your work

**Output:** `backup-project-YYYYMMDD-HHMMSS.doplan`

#### 2. Plan Backup (`--type plan`)
**What it backs up:**
- Only `.do/plan/` directory (tasks, phases, feature folders)

**Use case:** Preserving planning structure separately

**Output:** `backup-plan-YYYYMMDD-HHMMSS.doplan`

#### 3. Project-Plan Backup (`--type project-plan`)
**What it backs up:**
- Project files + `.do/plan/` directory
- **Excludes:** `.do/core/`, `.do/system/`, `.cursor/`

**Use case:** Updating DoPlan but keeping your project work + planning

**Output:** `backup-project-plan-YYYYMMDD-HHMMSS.doplan`

#### 4. Full Backup (`--type full`)
**What it backs up:**
- Everything (project + all `.do/` + `.cursor/`)

**Use case:** Complete disaster recovery

**Output:** `backup-full-YYYYMMDD-HHMMSS.doplan`

## Recommended Update Workflow

### Step 1: Create Project Backup
Before updating DoPlan, create a project backup:

```bash
/sys backup --type project --description "Before DoPlan update to v1.4.0"
```

This creates `backup-project-YYYYMMDD-HHMMSS.doplan` containing only your project files.

### Step 2: Install New DoPlan Version
Install or update DoPlan CLI to the new version:

```bash
npx @doplan-dev/cli@latest
```

Or if updating an existing project:
```bash
cd /path/to/new/doplan/installation
# New DoPlan system is installed with updated .do/core/, .do/system/, .cursor/
```

### Step 3: Restore Your Project
Restore your project files to the new DoPlan installation:

```bash
# In your new/updated project directory
/sys restore backup-project-YYYYMMDD-HHMMSS.doplan
```

This restores your project files while preserving the new DoPlan system files.

### Step 4: Verify & Continue
- Verify your project files are intact
- Check that DoPlan system is working
- Continue development with updated DoPlan features

## Alternative: Project-Plan Backup

If you also want to preserve your planning structure:

```bash
# Before update
/sys backup --type project-plan --description "Project + planning before update"

# After installing new DoPlan
/sys restore backup-project-plan-YYYYMMDD-HHMMSS.doplan
```

This restores both your project files AND your `.do/plan/` directory while getting the new `.do/core/` and `.do/system/` from the updated DoPlan.

## Benefits of This Approach

### ✅ Clean Separation
- Project work is separate from DoPlan system files
- Clear boundaries between user content and system files

### ✅ Safe Updates
- Never lose your project work during updates
- Can always rollback by restoring backup

### ✅ Flexibility
- Choose what to backup based on your needs
- Support different update scenarios

### ✅ Future-Proof
- Works even if DoPlan structure changes dramatically
- Compatible with major version updates

## Backup Manifest

Each backup includes a `BACKUP_MANIFEST.json` with:

```json
{
  "timestamp": "2025-11-30T13:00:00Z",
  "description": "Before DoPlan update to v1.4.0",
  "backup_type": "project",
  "included_paths": ["src/", "docs/", "package.json", ...],
  "excluded_paths": [".do/", ".cursor/", "node_modules/"],
  "file_count": 150,
  "compressed_size_mb": 2.5,
  "doplan_version": "1.3.0",
  "purpose": "Project backup for DoPlan system updates"
}
```

## Best Practices

1. **Always backup before major updates**
   - Use `--type project` or `--type project-plan`
   - Add descriptive messages with `--description`

2. **Keep multiple backup generations**
   - Don't delete old backups immediately
   - Keep at least 2-3 recent backups

3. **Test restore in dry-run mode first**
   ```bash
   /sys restore backup-project-YYYYMMDD-HHMMSS.doplan --dry-run
   ```

4. **Document your backup strategy**
   - Note which backup type you used
   - Keep track of what each backup contains

5. **Version control consideration**
   - Consider committing `.do/backup/` folder to git (excluding the actual backup files)
   - Or keep backups in a separate location

## Migration from Old Projects

If you have an old project using an older DoPlan structure:

1. **Create full backup first** (safety net):
   ```bash
   /sys backup --type full --description "Full backup before migration"
   ```

2. **Create project backup**:
   ```bash
   /sys backup --type project --description "Project files for migration"
   ```

3. **Set up new DoPlan project**:
   ```bash
   npx @doplan-dev/cli
   # Creates new project with latest DoPlan structure
   ```

4. **Restore your project files**:
   ```bash
   cd new-project-name
   /sys restore backup-project-YYYYMMDD-HHMMSS.doplan
   ```

5. **Restore planning if needed**:
   ```bash
   /sys restore backup-plan-YYYYMMDD-HHMMSS.doplan
   ```

## FAQ

**Q: Will this work if DoPlan structure changes significantly?**
A: Yes! Project backups exclude DoPlan system files, so your work is always safe regardless of system changes.

**Q: What if I want to keep my old `.do/system/` files?**
A: Use `--type project-plan` to preserve planning, but note that system files may conflict with new DoPlan versions. It's recommended to use fresh system files and only restore project content.

**Q: Can I automate this process?**
A: Yes, you can script the backup/restore process. The commands support flags for non-interactive use.

**Q: What about my custom agents/commands?**
A: Custom agents/commands in `.do/core/` will be preserved in `full` backups. For updates, consider:
- Exporting custom agents/commands separately
- Or using `full` backup if you have extensive customizations

## Summary

The granular backup system provides flexible, safe ways to update DoPlan while preserving your work. The recommended approach is:

1. Backup project files (`--type project`)
2. Update DoPlan
3. Restore project files

This ensures you always get the latest DoPlan features while keeping your work intact.

