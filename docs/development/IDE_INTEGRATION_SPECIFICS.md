# IDE Integration Specifics - Step-by-Step Setup Guide

This document provides detailed, step-by-step instructions for integrating DoPlan with each supported IDE and AI tool.

## 🎯 Integration Architecture

### Core Concept
DoPlan stores all AI-related files in `.doplan/ai/` (IDE-agnostic):
```
.doplan/
└── ai/
    ├── agents/          # AI agent definitions
    ├── rules/           # Workflow and design rules
    └── commands/        # Command definitions
```

Each IDE integration creates appropriate symlinks or configuration files to point to these directories.

## 🖱️ Cursor Integration

### Step-by-Step Setup

#### Step 1: Detect Cursor Installation
```go
func DetectCursor(projectRoot string) bool {
    cursorDir := filepath.Join(projectRoot, ".cursor")
    if _, err := os.Stat(cursorDir); err == nil {
        return true
    }
    return false
}
```

#### Step 2: Create Directory Structure
Create `.cursor/` subdirectories if they don't exist:
- `.cursor/agents/`
- `.cursor/rules/`
- `.cursor/commands/`

#### Step 3: Create Symlinks
Create symlinks from `.cursor/` to `.doplan/ai/`:
- `.cursor/agents/` → `.doplan/ai/agents/`
- `.cursor/rules/` → `.doplan/ai/rules/`
- `.cursor/commands/` → `.doplan/ai/commands/`

**Note**: On Windows, use file copy instead of symlinks.

#### Step 4: Generate Cursor-Specific Files

**File: `.cursor/rules/doplan.mdc`**
```markdown
# DoPlan Workflow Rules

You are working on a project managed by DoPlan CLI.

## Project Structure
- All phases are in `doplan/##-phase-name/`
- All features are in `doplan/##-phase-name/##-feature-name/`
- Documentation is in `doplan/` root

## Workflow
1. Always read the phase plan before implementing features
2. Follow the feature design specifications
3. Update progress.json after completing tasks
4. Use the DoPlan commands for project management
```

#### Step 5: Verify Integration
Check that all symlinks/files exist and are accessible.

## 💻 VS Code + GitHub Copilot Integration

### Step-by-Step Setup

#### Step 1: Detect VS Code
Check for `.vscode/` directory.

#### Step 2: Generate tasks.json
Create `.vscode/tasks.json` with DoPlan tasks:
- DoPlan: Dashboard
- DoPlan: Discuss Idea
- DoPlan: Generate Docs

#### Step 3: Generate settings.json
Create `.vscode/settings.json` with:
- GitHub Copilot settings
- File associations
- Extension recommendations

#### Step 4: Create Prompts Directory
Copy agent files to `.vscode/prompts/` with Copilot-specific headers.

#### Step 5: Generate Copilot Chat Instructions
Create `.vscode/prompts/doplan-context.md` with project context.

## 🎨 Kiro Integration

### Step-by-Step Setup

#### Step 1: Detect Kiro
Check for `.kiro/` or `kiro.config.json`.

#### Step 2: Create Guide
Create `.doplan/guides/kiro_setup.md` with manual setup instructions.

#### Step 3: Copy Files (if .kiro exists)
Copy files from `.doplan/ai/` to `.kiro/` directories.

## 🌊 Windsurf Integration

Similar to Cursor - create symlinks from `.windsurf/` to `.doplan/ai/`.

## 🔧 Qoder Integration

Create `.qoder/doplan.json` config file pointing to `.doplan/ai/` directories.

## 🤖 CLI Tools Integration (Gemini, Claude, etc.)

### Step-by-Step Setup

#### Step 1: Create CLI Wrapper Scripts
Create wrapper scripts in `.doplan/scripts/`:
- `gemini-doplan.sh`
- `claude-doplan.sh`

#### Step 2: Create Usage Guide
Create `.doplan/guides/cli_integration.md` with usage instructions.

## ⚙️ Generic/Other IDE Integration

### Step-by-Step Setup

#### Step 1: Create Generic Guide
Create `.doplan/guides/generic_ide_setup.md` with comprehensive instructions.

## 🔄 Integration Maintenance

### Update Integration
```go
func UpdateIDEIntegration(projectRoot, ide string) error {
    switch ide {
    case "cursor":
        return SetupCursor(projectRoot)
    case "vscode", "copilot":
        return SetupVSCode(projectRoot)
    // ... more cases
    }
}
```

### Verify Integration
Check that all required files/directories exist for the selected IDE.

## 📋 Integration Checklist

For each IDE:
- [ ] Detect IDE installation
- [ ] Create necessary directories
- [ ] Create symlinks or copy files
- [ ] Generate IDE-specific configuration
- [ ] Create usage guide
- [ ] Verify integration works
- [ ] Test agent access
- [ ] Test command execution
- [ ] Test rule application

## 🐛 Troubleshooting Common Issues

### Issue: Symlinks not working (Windows)
**Solution**: Use file copy instead of symlinks

### Issue: IDE not detecting files
**Solution**: 
1. Verify file paths are correct
2. Check file permissions
3. Restart IDE
4. Check IDE logs for errors

### Issue: Agents not loading
**Solution**:
1. Verify agent file format matches IDE requirements
2. Check file encoding (UTF-8)
3. Verify file extensions (.md, .mdc, etc.)

