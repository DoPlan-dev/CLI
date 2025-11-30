---
name: sys
category: core
trigger: "/sys [<subcommand>]"
description: "System commands for project management and configuration"
agentInvolvement:
  - Project Orchestrator
  - System Administrator
filesRead:
  - ".do/system/history/active_state.json"
  - ".do/system/user_profile.json"
filesModified:
  - ".do/system/** (various system files)"
examples:
  - "/sys status → Show system status"
  - "/sys config → Configure system settings"
  - "/sys engagement → Show engagement metrics"
  - "/sys backup → Create compressed backup (.doplan file)"
  - "/sys restore → Restore from backup file"
  - "/sys memory → Export or restore memory card"
---

# /sys

## Trigger
Exact match: /sys or /sys <subcommand>

## Description
System commands for project management, configuration, and system-level operations.

## Available Subcommands

### /sys status
Show current system status, active tasks, and project state.

### /sys config
Configure system settings and preferences.

### /sys engagement
Show user engagement metrics and interaction statistics.

### /sys reset
Reset system state or configuration (with confirmation).

### /sys backup
Create a compressed backup with flexible options for different backup scenarios. This command provides granular control over what gets backed up, supporting various update and migration strategies.

**Usage**: `/sys backup [--type <type>] [--description "backup description"]`

**Backup Types** (interactive selection if not specified):

1. **`--type project`** - Project files only (recommended for DoPlan updates)
   - Backs up all project files (src/, docs/, config files, etc.)
   - **Excludes**: `.do/` and `.cursor/` directories
   - Use when: Updating DoPlan system, migrating to new DoPlan version
   - Output: `backup-project-YYYYMMDD-HHMMSS.doplan`

2. **`--type plan`** - Planning folder only
   - Backs up only `.do/plan/` directory (tasks, phases, feature folders)
   - Use when: You want to preserve planning structure separately
   - Output: `backup-plan-YYYYMMDD-HHMMSS.doplan`

3. **`--type project-plan`** - Project files + planning folder
   - Backs up project files + `.do/plan/` directory
   - **Excludes**: `.do/core/`, `.do/system/`, `.cursor/`
   - Use when: Updating DoPlan but want to keep your planning
   - Output: `backup-project-plan-YYYYMMDD-HHMMSS.doplan`

4. **`--type full`** - Everything (default)
   - Backs up complete project including all `.do/` and `.cursor/` directories
   - Use when: Complete snapshot for disaster recovery
   - Output: `backup-full-YYYYMMDD-HHMMSS.doplan`

**Options**:
- `--type`: Specify backup type (project, plan, project-plan, full)
- `--description`: Add a description to the backup

**Examples**:
- `/sys backup --type project --description "Before DoPlan update"`
- `/sys backup --type plan`
- `/sys backup --type project-plan`
- `/sys backup` (interactive selection)

**Output**: Creates timestamped `.doplan` compressed archive in `.do/backup/`

**Format**: The `.doplan` file is a compressed archive (tar.gz) containing:
- Selected files/folders based on backup type
- `BACKUP_MANIFEST.json` with backup metadata (type, contents, purpose)

**Update Strategy**: Use `--type project` when updating DoPlan system, then restore project files to new DoPlan installation.

### /sys restore
Restore a project from a compressed backup. This command:
- Lists available `.doplan` backup files in `.do/backup/`
- Shows backup type and manifest information for each backup
- Allows you to select a backup to restore
- Intelligently restores based on backup type
- Creates a new backup before restoring (safety measure)

**Usage**: `/sys restore [backup-file.doplan] [--dry-run]`

**Options**:
- `--dry-run`: Preview what would be restored without actually restoring
- `backup-file.doplan`: Specify a specific backup file to restore (optional, without path - just filename)

**Examples**:
- `/sys restore backup-project-20251130-130000.doplan --dry-run`
- `/sys restore` (interactive selection)

**Restore Behavior by Backup Type**:
- **project**: Restores project files only (safe to use after installing new DoPlan system)
- **plan**: Restores only `.do/plan/` directory
- **project-plan**: Restores project files + `.do/plan/` directory
- **full**: Restores everything including `.do/` and `.cursor/` directories

**Format**: Extracts the `.doplan` compressed archive and restores files to their original locations based on backup type

### /sys memory
Manage memory card data for transferring user preferences and engagement data between projects. This command:
- Exports your current memory card (preferences, achievements, engagement data)
- Restores memory card from your last project
- Allows you to carry forward your DoPlan experience to new projects

**Usage**: `/sys memory [export|restore]`

**Options**:
- `export`: Export current memory card to `.do/backup/` for use in next project
- `restore`: Restore memory card from the last exported backup

**Example**: `/sys memory export` or `/sys memory restore`

**What's in Memory Card**:
- User preferences and work style
- Command usage statistics
- Achievements and challenges completed
- Engagement metrics and relationship level
- Preferred tech stack
- Memorable moments and interactions
- Project history and context

**Note**: Memory card is stored globally at `~/.doplan/memory_card.json`. Export/restore allows you to save project-specific versions.

## Action
When user types /sys or /sys <subcommand>:

The system command provides access to project management and configuration functions. Each subcommand performs specific system-level operations.

### Backup Action Details
When user types `/sys backup`:

1. **Select Backup Type**:
   - If `--type` flag provided: Use specified type
   - Otherwise: Display interactive menu:
     ```
     What would you like to backup?
     1. Project files only (for DoPlan updates) - excludes .do/ and .cursor/
     2. Planning folder only (.do/plan/)
     3. Project files + planning folder
     4. Everything (complete backup)
     Select option [1-4]:
     ```

2. **Create Backup Archive Based on Type**:
   
   **Type: project**:
   - Generate filename: `backup-project-YYYYMMDD-HHMMSS.doplan`
   - Include: All project files (src/, docs/, config files, package.json, etc.)
   - Exclude: `.do/`, `.cursor/`, `node_modules/`, `.git/`
   - Purpose: Clean project backup for DoPlan system updates
   
   **Type: plan**:
   - Generate filename: `backup-plan-YYYYMMDD-HHMMSS.doplan`
   - Include: Only `.do/plan/` directory (tasks, phases, features)
   - Exclude: Everything else
   - Purpose: Preserve planning structure separately
   
   **Type: project-plan**:
   - Generate filename: `backup-project-plan-YYYYMMDD-HHMMSS.doplan`
   - Include: Project files + `.do/plan/` directory
   - Exclude: `.do/core/`, `.do/system/`, `.do/backup/`, `.cursor/`
   - Purpose: Project work + planning, ready for DoPlan update
   
   **Type: full**:
   - Generate filename: `backup-full-YYYYMMDD-HHMMSS.doplan`
   - Include: Everything (project files + all `.do/` + `.cursor/`)
   - Exclude: `node_modules/`, `.git/` (unless specified)
   - Purpose: Complete disaster recovery backup

3. **Create Backup Manifest**:
   - Create `BACKUP_MANIFEST.json` with metadata:
     * timestamp (backup creation time)
     * description (user-provided description)
     * backup type (project, plan, project-plan, full)
     * included paths (list of what's backed up)
     * excluded paths (list of what's excluded)
     * file count
     * compressed size
     * version (DoPlan CLI version)
     * purpose (recommended use case)
   - Add manifest to archive

4. **Compress Archive**:
   - Compress selected files into `.doplan` format (tar.gz)
   - Save to `.do/backup/backup-[type]-[timestamp].doplan`

5. **Response**: "✅ Backup created: backup-[type]-[timestamp].doplan in .do/backup/ (size: X MB, type: [type])"

### Restore Action Details
When user types `/sys restore`:

1. **List Available Backups**:
   - Scan `.do/backup/` for `.doplan` backup files
   - Extract manifest from each backup to display metadata
   - Display list with:
     * Backup filename
     * Backup type (project, plan, project-plan, full)
     * Timestamp
     * Description
     * Size
     * Purpose/use case

2. **Select Backup**:
   - If backup filename provided: use that backup file
   - Otherwise: prompt user to select from list with backup types clearly shown

3. **Safety Check**:
   - Create current-state backup before restoring (prevent data loss)
   - Read backup manifest to understand what will be restored
   - Show summary: "This backup contains: [type description]"
   - Request confirmation

4. **Dry Run Mode** (if `--dry-run` flag):
   - Extract manifest from backup archive
   - Show what would be restored based on backup type
   - List files that would be overwritten
   - Display restore strategy
   - Do not actually extract or restore

5. **Restore Process** (based on backup type):

   **Project Backup**:
   - Extract project files only
   - Restore to current directory
   - Skip `.do/` and `.cursor/` (preserve current DoPlan system)
   - Perfect for: Restoring work after DoPlan update

   **Plan Backup**:
   - Extract only `.do/plan/` directory
   - Restore planning structure and tasks
   - Preserve existing `.do/core/` and `.do/system/`

   **Project-Plan Backup**:
   - Extract project files + `.do/plan/`
   - Restore to appropriate locations
   - Preserve `.do/core/` and `.do/system/` (DoPlan system files)

   **Full Backup**:
   - Extract everything including all `.do/` directories
   - Full restore of project + DoPlan system
   - Use for complete restoration

6. **Response**: "✅ Restore complete! Files restored from backup-[type]-[timestamp].doplan"

### Memory Action Details
When user types `/sys memory`:

1. **Interactive Menu**:
   - Display options:
     * Option 1: Export memory to use in next project
     * Option 2: Restore memory from last project
   - Wait for user selection

2. **Export Memory** (if user selects export):
   - Read memory card from `~/.doplan/memory_card.json`
   - Create backup directory if needed: `.do/backup/`
   - Export memory card to: `.do/backup/memory_card_export-YYYYMMDD-HHMMSS.json`
   - Create export manifest with metadata:
     * Export timestamp
     * Project name (current project)
     * Memory card version
     * Data summary (achievements count, relationship level, etc.)
   - **Response**: "✅ Memory exported! Saved to .do/backup/memory_card_export-[timestamp].json. Use this in your next project with /sys memory restore"

3. **Restore Memory** (if user selects restore):
   - Scan `.do/backup/` for exported memory card files
   - List available exports with timestamps and project info
   - If multiple found: prompt user to select which to restore
   - If single found: show info and ask for confirmation
   - Create backup of current memory card before restoring (safety)
   - Restore memory card data to `~/.doplan/memory_card.json`
   - Update global memory card with restored data
   - **Response**: "✅ Memory restored! Your preferences, achievements, and engagement data have been loaded from [export-file]"

4. **Memory Card Data**:
   - User preferences (work style, experience level)
   - Command usage statistics
   - Achievements and challenges
   - Engagement metrics and relationship level
   - Preferred tech stack
   - Memorable moments
   - Project context and history

## Agent Involvement
- **Project Orchestrator**: Coordinates system operations
- **System Administrator**: Manages configuration and state

## Files Modified
- Various system files in .do/system/ depending on subcommand

## Status
✅ **ACTIVE**: System command available for use.

