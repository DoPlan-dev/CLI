# Migration Guide

Complete guide for upgrading DoPlan CLI and migrating projects.

---

## Upgrading DoPlan CLI

### Using npx (Recommended)

**No upgrade needed!** `npx` always uses the latest version:

```bash
npx @doplan-dev/cli
```

---

### Using Homebrew

```bash
brew update
brew upgrade doplan
```

---

### Using Scoop

```bash
scoop update doplan
```

---

### Direct Binary

1. Download latest release from [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Replace existing binary
3. Verify:
   ```bash
   doplan --version
   ```

---

### Building from Source

```bash
git pull origin main
go build -o doplan ./cmd/doplan
```

---

## Migrating Projects

### Project Structure Changes

**Version 1.0 to 1.1**: No breaking changes

**Version 1.1 to 1.2**: 
- New agent added
- Update agent files if customized

**Migration Steps**:
1. Pull latest DoPlan CLI
2. Regenerate project (optional)
3. Merge customizations

---

### Agent Updates

**When**: New agents added or updated

**Migration**:
1. Check new agent files
2. Merge customizations
3. Update references if needed

---

### Command Updates

**When**: Commands updated or new commands added

**Migration**:
1. Check command files
2. Merge customizations
3. Update usage if needed

---

### Rules Updates

**When**: Rules library updated

**Migration**:
1. Check updated rules
2. Merge custom rules
3. Update references if needed

---

## Breaking Changes

### Version 1.0

**No breaking changes** - Initial release

---

### Future Versions

Breaking changes will be documented here and in release notes.

---

## Compatibility Information

### DoPlan CLI Versions

| Version | Release Date | Status |
|---------|--------------|--------|
| 1.0.0   | 2025-01-XX   | Current |
| 1.0.1   | TBD          | Future |
| 1.1.0   | TBD          | Future |

---

### Project Compatibility

**Version 1.0 Projects**: Compatible with all 1.0.x versions

**Future Versions**: Compatibility will be maintained within major versions

---

## Upgrade Step-by-Step Guides

### Guide 1: Quick Upgrade (npx)

**Time**: < 1 minute

**Steps**:
1. No action needed
2. `npx` uses latest version automatically

---

### Guide 2: Homebrew Upgrade

**Time**: 2-3 minutes

**Steps**:
1. Update Homebrew:
   ```bash
   brew update
   ```
2. Upgrade DoPlan CLI:
   ```bash
   brew upgrade doplan
   ```
3. Verify:
   ```bash
   doplan --version
   ```

---

### Guide 3: Manual Upgrade

**Time**: 5 minutes

**Steps**:
1. Download latest release
2. Backup current binary (optional)
3. Replace binary
4. Verify:
   ```bash
   doplan --version
   ```

---

## Migration Scripts

### Script 1: Backup Project

```bash
#!/bin/bash
# Backup project before migration

PROJECT_NAME=$1
BACKUP_DIR="backup-$(date +%Y%m%d)"

mkdir -p $BACKUP_DIR
cp -r $PROJECT_NAME $BACKUP_DIR/
echo "Backup created: $BACKUP_DIR"
```

---

### Script 2: Update Agents

```bash
#!/bin/bash
# Update agent files

# Backup custom agents
cp -r .cursor/agents .cursor/agents.backup

# Update agents (manual merge required)
echo "Update agent files manually"
```

---

## Rollback Procedures

### Rollback DoPlan CLI

**Homebrew**:
```bash
brew uninstall doplan
brew install doplan@1.0.0
```

**Direct Binary**:
1. Download previous version
2. Replace binary
3. Verify version

---

### Rollback Project

**If using Git**:
```bash
git checkout <previous-commit>
```

**If using backup**:
```bash
cp -r backup-YYYYMMDD/project-name .
```

---

## Version Compatibility Matrix

| DoPlan CLI | Project v1.0 | Project v1.1 | Project v1.2 |
|------------|--------------|--------------|--------------|
| 1.0.x      | ✅           | ✅           | ✅           |
| 1.1.x      | ✅           | ✅           | ✅           |
| 1.2.x      | ⚠️           | ✅           | ✅           |

**Legend**:
- ✅ Fully compatible
- ⚠️ Compatible with migration
- ❌ Not compatible

---

## Troubleshooting Upgrades

### Issue: Version Mismatch

**Symptoms**: Project generated with different version

**Solution**:
1. Check DoPlan CLI version:
   ```bash
   doplan --version
   ```
2. Upgrade if needed
3. Regenerate project if necessary

---

### Issue: Missing Features

**Symptoms**: Features not available after upgrade

**Solution**:
1. Check release notes
2. Verify version
3. Update project structure if needed

---

### Issue: Customizations Lost

**Symptoms**: Custom agents/commands/rules missing

**Solution**:
1. Restore from backup
2. Merge with new version
3. Update references

---

## Best Practices

### 1. Backup Before Upgrade

Always backup:
- Custom agents
- Custom commands
- Custom rules
- Project state

---

### 2. Test After Upgrade

Test:
- Project generation
- Commands
- Agents
- Rules

---

### 3. Read Release Notes

Before upgrading:
- Read release notes
- Check breaking changes
- Review migration guide

---

### 4. Gradual Migration

For major versions:
- Test in development
- Migrate one project at a time
- Verify before full migration

---

## Related Pages

- [Installation](Installation) - Installation guide
- [Release Notes](Release-Notes) - Version history
- [Troubleshooting](Troubleshooting) - Common issues
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

