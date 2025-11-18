# DoPlan v0.0.18-beta Migration Strategy

This document outlines the migration strategy for existing DoPlan projects (v0.0.17-beta) to the new v0.0.18-beta architecture.

## 🎯 Migration Goals

1. **Zero data loss** - All project data must be preserved
2. **Backward compatibility** - Support old structure during transition
3. **Automatic migration** - Minimize manual steps for users
4. **Rollback capability** - Allow reverting if issues occur
5. **Clear communication** - Inform users of changes

## 📊 Current vs. New Structure

### Folder Structure Changes

**Old Structure (v0.0.17):**
```
project-root/
├── .cursor/
│   └── config/
│       ├── doplan-config.json
│       └── doplan-state.json
└── doplan/
    ├── 01-phase/
    │   ├── phase-plan.md
    │   ├── phase-progress.json
    │   ├── 01-Feature/
    │   │   ├── plan.md
    │   │   ├── design.md
    │   │   ├── tasks.md
    │   │   └── progress.json
    │   └── 02-Feature/
    └── 02-phase/
```

**New Structure (v0.0.18):**
```
project-root/
├── .doplan/                    [NEW - Replaces .cursor/config/]
│   ├── config.yaml            [NEW - Replaces doplan-config.json]
│   ├── state.json             [NEW - Replaces doplan-state.json]
│   ├── dashboard.json         [NEW - Machine-readable dashboard]
│   └── ai/                     [NEW - IDE-agnostic AI files]
│       ├── agents/
│       ├── rules/
│       └── commands/
└── doplan/
    ├── 01-user-authentication/  [NEW - Slug-based naming]
    │   ├── phase-plan.md
    │   ├── phase-progress.json
    │   ├── 01-login-with-email/  [NEW - Slug-based naming]
    │   │   ├── plan.md
    │   │   ├── design.md
    │   │   ├── tasks.md
    │   │   └── progress.json
    │   └── 02-password-reset/
    └── 02-user-profile/
```

### Configuration Changes

**Old Config (doplan-config.json):**
```json
{
  "ide": "cursor",
  "installed": true,
  "version": "1.0.0",
  "github": {
    "enabled": true,
    "autoBranch": true,
    "autoPR": true
  }
}
```

**New Config (config.yaml):**
```yaml
project:
  name: "Project Name"
  type: "web"
  version: "1.0.0"
  ide: "cursor"  # NEW

github:
  repository: "owner/repo"  # REQUIRED
  enabled: true
  autoBranch: true
  autoPR: true

design:
  hasPreferences: false
  tokensPath: "doplan/design/design-tokens.json"

security:
  lastScan: null
  autoFix: false

apis:
  configured: []
  required: []

tui:
  theme: "default"
  animations: true
```

## 🔄 Migration Process

### Phase 1: Detection & Backup

**Step 1.1: Detect Old Structure**
- Check for `.cursor/config/doplan-config.json`
- Check for `doplan/01-phase/` structure
- Return true if old structure found

**Step 1.2: Create Backup**
- Create timestamped backup directory
- Copy `.cursor/config/` to backup
- Copy `doplan/` to backup
- Return backup path

### Phase 2: Config Migration

**Step 2.1: Migrate Config File**
- Read `doplan-config.json`
- Convert to YAML format
- Add new required fields with defaults
- Write to `.doplan/config.yaml`
- Validate migrated config

**Migration Mapping:**
- `ide` → `project.ide`
- `version` → `project.version`
- `github.enabled` → `github.enabled`
- `github.autoBranch` → `github.autoBranch`
- `github.autoPR` → `github.autoPR`
- Add `github.repository` (prompt user if missing)
- Add `project.name` (extract from directory or prompt)
- Add `project.type` (detect from project files)

**Step 2.2: Migrate State File**
- Read `doplan-state.json`
- Convert to new `state.json` format
- Add action history structure
- Write to `.doplan/state.json`

### Phase 3: Folder Structure Migration

**Step 3.1: Detect Old Folder Names**
- Scan `doplan/` directory
- Find folders matching `01-phase`, `02-phase`, etc.
- Find subfolders matching `01-Feature`, `02-Feature`, etc.
- Return list of old folders

**Step 3.2: Generate New Folder Names**

**Option A: Automatic (Recommended)**
- Read `phase-plan.md` or feature `plan.md`
- Extract phase/feature name
- Convert to slug (lowercase, hyphens)
- Return `01-slug-name` format

**Option B: Manual (Fallback)**
- If `plan.md` doesn't exist or name extraction fails
- Prompt user for slug name
- Or use existing name converted to slug

**Step 3.3: Rename Folders**
- For each folder:
  1. Create new folder with new name
  2. Copy all files
  3. Update references in `progress.json` files
  4. Update references in `dashboard.md`
  5. Delete old folder (after verification)

**Example Migration:**
```
Old: doplan/01-phase/01-Feature/
New: doplan/01-user-authentication/01-login-with-email/

Process:
1. Read 01-phase/phase-plan.md → Extract "User Authentication"
2. Read 01-phase/01-Feature/plan.md → Extract "Login with Email"
3. Generate slugs: "user-authentication", "login-with-email"
4. Create new folders
5. Copy all files
6. Update progress.json references
7. Update dashboard.md
8. Verify migration
9. Delete old folders
```

### Phase 4: Dashboard Migration

**Step 4.1: Generate dashboard.json**
- Read all `progress.json` files
- Parse phase and feature data
- Generate `dashboard.json` with new structure
- Keep `dashboard.md` for backward compatibility

### Phase 5: IDE Integration Migration

**Step 5.1: Migrate AI Files**
- If `.cursor/` exists:
  - Copy `.cursor/agents/` → `.doplan/ai/agents/`
  - Copy `.cursor/rules/` → `.doplan/ai/rules/`
  - Copy `.cursor/commands/` → `.doplan/ai/commands/`
- Create symlinks back to `.cursor/` (for Cursor compatibility)

## 🛡️ Safety Measures

### 1. Pre-Migration Validation
- Check disk space
- Verify write permissions
- Check for active git operations
- Verify backup creation succeeded

### 2. Migration Rollback
- Restore `.cursor/config/` from backup
- Restore `doplan/` structure from backup
- Remove `.doplan/` directory
- Verify restoration

### 3. Migration Verification
- Verify `.doplan/config.yaml` exists and is valid
- Verify `.doplan/state.json` exists
- Verify folder structure matches new format
- Verify all `progress.json` files updated
- Verify `dashboard.json` generated
- Return list of issues if any

## 📋 Migration Wizard Flow

See `TUI_WIZARD_FLOW.md` for detailed TUI wizard screens.

## 🔧 Manual Migration (Fallback)

If automatic migration fails, provide manual steps:

### Step 1: Backup
```bash
# Create backup
cp -r .cursor .cursor.backup
cp -r doplan doplan.backup
```

### Step 2: Create .doplan directory
```bash
mkdir -p .doplan/ai/{agents,rules,commands}
```

### Step 3: Migrate config
```bash
# Convert config (manual or with script)
doplan migrate-config
```

### Step 4: Rename folders
```bash
# Rename phases
cd doplan
mv 01-phase 01-user-authentication
mv 02-phase 02-user-profile

# Rename features
cd 01-user-authentication
mv 01-Feature 01-login-with-email
```

### Step 5: Update references
```bash
# Update progress.json files
# Update dashboard.md
doplan migrate-references
```

## 📝 Migration Checklist for Users

- [ ] **Pre-migration**
  - [ ] Commit all changes to git
  - [ ] Create manual backup (optional but recommended)
  - [ ] Close any active DoPlan processes

- [ ] **During migration**
  - [ ] Run `doplan` (will auto-detect and prompt)
  - [ ] Review folder rename suggestions
  - [ ] Provide GitHub repository URL
  - [ ] Select IDE integration

- [ ] **Post-migration**
  - [ ] Verify dashboard loads correctly
  - [ ] Check folder structure
  - [ ] Test IDE integration
  - [ ] Verify progress tracking works
  - [ ] Report any issues

## 🚨 Troubleshooting

### Issue: Migration fails mid-process
**Solution:** Use rollback command
```bash
doplan migrate-rollback .doplan/backup/2024-01-15-10-30/
```

### Issue: Folder names not auto-detected
**Solution:** Manual rename with migration wizard
- Wizard will prompt for each folder name
- Or skip and rename manually later

### Issue: Config migration errors
**Solution:** Manual config creation
```bash
# Create .doplan/config.yaml manually
# Use template from documentation
doplan migrate-config --manual
```

### Issue: IDE integration not working
**Solution:** Re-run IDE setup
```bash
doplan  # Opens TUI
# Select: ⚙️ Setup AI/IDE Integration
```

## 📊 Migration Status Tracking

Track migration in `.doplan/migration.log`:
```json
{
  "version": "0.0.18-beta",
  "migratedAt": "2024-01-15T10:30:00Z",
  "backupPath": ".doplan/backup/2024-01-15-10-30/",
  "steps": {
    "config": "completed",
    "state": "completed",
    "folders": "completed",
    "dashboard": "completed",
    "ide": "completed"
  },
  "issues": []
}
```

## ✅ Success Criteria

Migration is successful when:
1. ✅ `.doplan/config.yaml` exists and is valid
2. ✅ `.doplan/state.json` exists
3. ✅ All folders use new naming (`##-slug-name`)
4. ✅ `dashboard.json` is generated and accurate
5. ✅ IDE integration files are in place
6. ✅ Dashboard TUI loads correctly
7. ✅ All progress data is preserved
8. ✅ No data loss occurred

