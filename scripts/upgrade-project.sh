#!/bin/bash
# upgrade-project.sh
# Migrates existing DoPlan projects to the new .do/ structure

set -e  # Exit on error

PROJECT_DIR=${1:-.}
BACKUP_DIR="backup-$(date +%Y%m%d-%H%M%S)"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_NAME=$(basename "$(cd "$PROJECT_DIR" && pwd)")

echo "🔄 Starting project upgrade for: $PROJECT_NAME"
echo "📍 Project directory: $(cd "$PROJECT_DIR" && pwd)"
echo ""

# Step 1: Backup
echo "📦 Creating backup..."
if [ ! -d "$BACKUP_DIR" ]; then
    mkdir -p "$BACKUP_DIR"
fi
cp -r "$PROJECT_DIR" "$BACKUP_DIR/" 2>/dev/null || {
    echo "⚠️  Warning: Could not create full backup, continuing with selective backup..."
    mkdir -p "$BACKUP_DIR/$PROJECT_NAME"
}

# Backup critical files
if [ -d "$PROJECT_DIR/.plan" ]; then
    echo "  → Backing up .plan/ directory..."
    cp -r "$PROJECT_DIR/.plan" "$BACKUP_DIR/$PROJECT_NAME/" 2>/dev/null || true
fi

if [ -d "$PROJECT_DIR/.cursor" ]; then
    echo "  → Backing up .cursor/ directory..."
    cp -r "$PROJECT_DIR/.cursor" "$BACKUP_DIR/$PROJECT_NAME/" 2>/dev/null || true
fi

echo "✅ Backup created: $BACKUP_DIR"
echo ""

# Step 2: Navigate to project directory
cd "$PROJECT_DIR"

# Step 3: Create new structure
echo "📁 Creating new directory structure..."
mkdir -p .do/core
mkdir -p .do/system/{history,content}
mkdir -p .do/plan
mkdir -p .cursor/{agents,commands,rules/library}
echo "✅ Directory structure created"
echo ""

# Step 4: Migrate existing files
echo "📋 Migrating files..."

# Migrate from .plan/ to .do/ if exists
if [ -d ".plan/00_System" ]; then
    echo "  → Migrating planning documents from .plan/00_System/ to .do/system/..."
    cp -r .plan/00_System/* .do/system/ 2>/dev/null || {
        echo "    ⚠️  Some files could not be copied, but continuing..."
    }
    echo "    ✅ Planning documents migrated"
fi

if [ -f ".plan/TASKS.md" ]; then
    echo "  → Migrating TASKS.md to .do/plan/..."
    cp .plan/TASKS.md .do/plan/ 2>/dev/null || true
    echo "    ✅ TASKS.md migrated"
fi

if [ -f ".plan/active_state.json" ]; then
    echo "  → Migrating active_state.json to .do/system/history/..."
    cp .plan/active_state.json .do/system/history/ 2>/dev/null || true
    echo "    ✅ active_state.json migrated"
fi

# Check for old active_state.json in root of .plan/
if [ -f ".plan/active_state.json" ]; then
    echo "  → Found active_state.json in .plan/, already migrated"
fi

# Migrate any existing history files
if [ -d ".plan/history" ]; then
    echo "  → Migrating history files..."
    cp -r .plan/history/* .do/system/history/ 2>/dev/null || true
    echo "    ✅ History files migrated"
fi

# Migrate phase folders (01-foundation, 02-core_features, etc.)
echo "  → Migrating phase folders..."
for phase_dir in .plan/0[1-9]-*; do
    if [ -d "$phase_dir" ]; then
        phase_name=$(basename "$phase_dir")
        echo "    → Migrating $phase_name/ to .do/plan/$phase_name/..."
        cp -r "$phase_dir" .do/plan/ 2>/dev/null || {
            echo "      ⚠️  Could not copy $phase_name, continuing..."
        }
    fi
done
# Count migrated phase folders
migrated_phases=$(find .do/plan -maxdepth 1 -type d -name "0[1-9]-*" 2>/dev/null | wc -l | tr -d ' ')
if [ "$migrated_phases" -gt 0 ]; then
    echo "    ✅ Migrated $migrated_phases phase folder(s)"
else
    echo "    ℹ️  No phase folders found to migrate"
fi

echo "✅ File migration completed"
echo ""

# Step 5: Backup customizations
echo "💾 Backing up customizations..."

if [ -d ".cursor/agents" ]; then
    if [ ! -d ".cursor/agents.backup" ]; then
        echo "  → Backing up custom agents..."
        cp -r .cursor/agents .cursor/agents.backup 2>/dev/null || true
        echo "    ✅ Agents backed up to .cursor/agents.backup"
    else
        echo "  ℹ️  Agents backup already exists"
    fi
fi

if [ -d ".cursor/commands" ]; then
    if [ ! -d ".cursor/commands.backup" ]; then
        echo "  → Backing up custom commands..."
        cp -r .cursor/commands .cursor/commands.backup 2>/dev/null || true
        echo "    ✅ Commands backed up to .cursor/commands.backup"
    else
        echo "  ℹ️  Commands backup already exists"
    fi
fi

if [ -d ".cursor/rules" ]; then
    if [ ! -d ".cursor/rules.backup" ]; then
        echo "  → Backing up custom rules..."
        cp -r .cursor/rules .cursor/rules.backup 2>/dev/null || true
        echo "    ✅ Rules backed up to .cursor/rules.backup"
    else
        echo "  ℹ️  Rules backup already exists"
    fi
fi

echo "✅ Customizations backed up"
echo ""

# Step 6: Move content to .do/core/ and create category symlinks in .cursor/
echo "🔗 Setting up correct structure (source in .do/core/, category symlinks in .cursor/)..."

# Remove old symlinks if they exist (old structure where entire folders were symlinked)
if [ -L ".cursor/agents" ]; then
    echo "  → Removing old agents symlink..."
    rm -f .cursor/agents
fi
if [ -L ".cursor/commands" ]; then
    echo "  → Removing old commands symlink..."
    rm -f .cursor/commands
fi

# Ensure .cursor/agents and .cursor/commands are directories
mkdir -p .cursor/agents .cursor/commands

# Move agents from .cursor/ to .do/core/ if they exist (and not already moved)
if [ -d ".cursor/agents" ] && [ ! -d ".do/core/agents" ] && [ ! -L ".cursor/agents" ] && [ -n "$(ls -A .cursor/agents 2>/dev/null)" ]; then
    echo "  → Moving agents from .cursor/ to .do/core/..."
    # Check if it's a directory with files or symlinks
    if [ "$(find .cursor/agents -maxdepth 1 -type f 2>/dev/null | wc -l)" -gt 0 ] || [ "$(find .cursor/agents -maxdepth 1 -type d ! -name . 2>/dev/null | wc -l)" -gt 0 ]; then
        mv .cursor/agents/* .do/core/agents/ 2>/dev/null || true
        echo "    ✅ Agents moved to .do/core/"
    fi
fi

# Ensure .do/core/agents exists
mkdir -p .do/core/agents

# Create symlinks for each agent category folder
if [ -d ".do/core/agents" ]; then
    echo "  → Creating symlinks for agent category folders..."
    for category_dir in .do/core/agents/*/; do
        if [ -d "$category_dir" ]; then
            category=$(basename "$category_dir")
            if [ ! -L ".cursor/agents/$category" ]; then
                ln -sfn "../../.do/core/agents/$category" ".cursor/agents/$category" 2>/dev/null && \
                echo "    ✅ .cursor/agents/$category → ../../.do/core/agents/$category"
            fi
        fi
    done
fi

# Move commands from .cursor/ to .do/core/ if they exist
if [ -d ".cursor/commands" ] && [ ! -d ".do/core/commands" ] && [ ! -L ".cursor/commands" ] && [ -n "$(ls -A .cursor/commands 2>/dev/null)" ]; then
    echo "  → Moving commands from .cursor/ to .do/core/..."
    if [ "$(find .cursor/commands -maxdepth 1 -type f 2>/dev/null | wc -l)" -gt 0 ] || [ "$(find .cursor/commands -maxdepth 1 -type d ! -name . 2>/dev/null | wc -l)" -gt 0 ]; then
        mv .cursor/commands/* .do/core/commands/ 2>/dev/null || true
        echo "    ✅ Commands moved to .do/core/"
    fi
fi

# Ensure .do/core/commands exists
mkdir -p .do/core/commands

# Create symlinks for each command category folder
if [ -d ".do/core/commands" ]; then
    echo "  → Creating symlinks for command category folders..."
    for category_dir in .do/core/commands/*/; do
        if [ -d "$category_dir" ]; then
            category=$(basename "$category_dir")
            if [ ! -L ".cursor/commands/$category" ]; then
                ln -sfn "../../.do/core/commands/$category" ".cursor/commands/$category" 2>/dev/null && \
                echo "    ✅ .cursor/commands/$category → ../../.do/core/commands/$category"
            fi
        fi
    done
fi

# Move rules/library from .cursor/rules/ to .do/core/ if it exists
if [ -d ".cursor/rules/library" ] && [ ! -d ".do/core/library" ] && [ ! -L ".cursor/rules/library" ]; then
    echo "  → Moving rules library from .cursor/rules/ to .do/core/..."
    mv .cursor/rules/library .do/core/library 2>/dev/null || true
    echo "    ✅ Rules library moved to .do/core/"
fi

# Create symlink in .cursor/rules/ pointing to .do/core/library
if [ -d ".do/core/library" ] && [ ! -L ".cursor/rules/library" ]; then
    echo "  → Creating symlink: .cursor/rules/library → ../../.do/core/library"
    ln -sfn "../../.do/core/library" .cursor/rules/library 2>/dev/null || true
fi

echo "✅ Structure corrected - source in .do/core/, category symlinks in .cursor/"
echo ""

# Step 7: Summary
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ Upgrade script completed!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "📋 Summary:"
echo "  • Backup created: $BACKUP_DIR"
echo "  • New structure: .do/ directory created"
echo "  • Files migrated: Planning docs, tasks, state"
echo "  • Customizations: Backed up with .backup suffix"
echo ""
echo "📝 Next steps:"
echo "  1. Verify the migrated files in .do/ directory"
echo "  2. Regenerate missing agents/commands from new CLI version:"
echo "     npx @doplan-dev/cli"
echo "     (Create a temp project and copy missing files)"
echo "  3. Merge your customizations from .backup directories"
echo "  4. Test your project to ensure everything works"
echo "  5. Remove old .plan/ directory if migration successful:"
echo "     rm -rf .plan/"
echo ""
echo "⚠️  Important:"
echo "  • Keep the backup directory until you verify everything works"
echo "  • Review differences between old and new structure"
echo "  • Test all commands and workflows before removing backups"
echo ""
echo "📖 For detailed migration guide, see:"
echo "   docs/wiki-archive/Migration-Guide.md"
echo ""

