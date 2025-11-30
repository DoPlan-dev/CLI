# Project Upgrade Script

This script helps migrate existing DoPlan projects to the new `.do/` directory structure.

## What It Does

The `upgrade-project.sh` script:

1. ✅ Creates a full backup of your project
2. ✅ Creates the new `.do/` directory structure
3. ✅ Migrates planning documents from `.plan/` to `.do/system/`
4. ✅ Migrates tasks from `.plan/TASKS.md` to `.do/plan/TASKS.md`
5. ✅ Migrates state files to `.do/system/history/`
6. ✅ Backs up your custom agents, commands, and rules
7. ✅ Sets up symlinks between `.do/core/` and `.cursor/`

## Usage

### Basic Usage (Current Directory)

```bash
cd /path/to/your-project
./scripts/upgrade-project.sh
```

Or from the DoPlan CLI repository:

```bash
./scripts/upgrade-project.sh /path/to/your-project
```

### Specify Project Directory

```bash
./scripts/upgrade-project.sh ~/projects/my-old-project
```

## Prerequisites

- Bash shell
- Project must have existing `.plan/` or `.cursor/` directories
- Write permissions in the project directory

## What Gets Migrated

| Old Location | New Location |
|--------------|--------------|
| `.plan/00_System/*.md` | `.do/system/*.md` |
| `.plan/TASKS.md` | `.do/plan/TASKS.md` |
| `.plan/active_state.json` | `.do/system/history/active_state.json` |
| `.plan/history/*` | `.do/system/history/*` |
| `.cursor/agents/` | Backed up to `.cursor/agents.backup` |
| `.cursor/commands/` | Backed up to `.cursor/commands.backup` |
| `.cursor/rules/` | Backed up to `.cursor/rules.backup` |

## Backup Location

Backups are created in:
```
backup-YYYYMMDD-HHMMSS/your-project-name/
```

Example:
```
backup-20251129-031800/my-project/
```

## After Running the Script

### 1. Verify Migration

Check that files were migrated correctly:

```bash
ls -la .do/system/
ls -la .do/plan/
ls -la .do/system/history/
```

### 2. Regenerate Missing Files

If you're missing agents, commands, or rules from the new CLI version:

```bash
# Create a temporary new project
cd /tmp
npx @doplan-dev/cli
# Name it: temp-upgrade

# Copy missing files to your project
cp -r temp-upgrade/.do/core/* /path/to/your-project/.do/core/
cp -r temp-upgrade/.cursor/* /path/to/your-project/.cursor/

# Clean up
rm -rf temp-upgrade
```

### 3. Merge Customizations

Your custom agents, commands, and rules are backed up in:
- `.cursor/agents.backup/`
- `.cursor/commands.backup/`
- `.cursor/rules.backup/`

Merge your customizations with the new files:

```bash
# Example: Merge custom agent
diff .cursor/agents.backup/my_custom_agent.md .cursor/agents/my_custom_agent.md
# Then manually merge or copy your customizations
```

### 4. Clean Up (After Verification)

Once you've verified everything works:

```bash
# Remove old .plan/ directory
rm -rf .plan/

# Remove backup directories (keep main backup for safety)
rm -rf .cursor/*.backup

# Keep the main backup directory until you're 100% sure
```

## Troubleshooting

### Script Fails with Permission Error

```bash
# Make script executable
chmod +x scripts/upgrade-project.sh

# Or run with bash
bash scripts/upgrade-project.sh /path/to/project
```

### Backup Fails

The script will continue even if full backup fails. Check for selective backups in the backup directory.

### Files Not Found

If you get warnings about files not found, that's normal if your project structure differs. The script will continue and migrate what it finds.

### Symlinks Don't Work

On some systems, symlinks might not work. The script will fall back to creating directories instead.

## Safety Features

- ✅ Creates full backup before making changes
- ✅ Backs up customizations separately
- ✅ Non-destructive: original files are preserved
- ✅ Continues even if some files are missing
- ✅ Shows clear summary of what was done

## Example Output

```
🔄 Starting project upgrade for: my-project
📍 Project directory: /Users/you/projects/my-project

📦 Creating backup...
  → Backing up .plan/ directory...
  → Backing up .cursor/ directory...
✅ Backup created: backup-20251129-031800

📁 Creating new directory structure...
✅ Directory structure created

📋 Migrating files...
  → Migrating planning documents from .plan/00_System/ to .do/system/...
    ✅ Planning documents migrated
  → Migrating TASKS.md to .do/plan/...
    ✅ TASKS.md migrated
  → Migrating active_state.json to .do/system/history/...
    ✅ active_state.json migrated
✅ File migration completed

💾 Backing up customizations...
  → Backing up custom agents...
    ✅ Agents backed up to .cursor/agents.backup
✅ Customizations backed up

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
✅ Upgrade script completed!
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## Need Help?

- 📖 See full [Migration Guide](../../docs/wiki-archive/Migration-Guide.md)
- 🐛 Report issues: https://github.com/DoPlan-dev/CLI/issues
- 💬 Ask questions: https://github.com/DoPlan-dev/CLI/discussions

