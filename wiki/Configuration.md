# Configuration Reference

Complete reference for all configuration files and settings in DoPlan CLI projects.

---

## Overview

DoPlan CLI projects use a minimal configuration approach. Most configuration is embedded in:
- Agent definitions (`.cursor/agents/`)
- Command definitions (`.cursor/commands/`)
- Rules library (`.cursor/rules/library/`)
- Project state (`.plan/active_state.json`)

---

## Configuration Files

### `.plan/active_state.json`

**Purpose**: Tracks current project state and progress.

**Location**: `.plan/active_state.json`

**Structure**:
```json
{
  "phase": "idea|brainstorm|writing|approved|tasks|building",
  "active_task": null | "task_id",
  "completed": ["task_id", ...],
  "locked": false | true
}
```

**Fields**:
- `phase` - Current workflow phase
- `active_task` - Currently active task ID (null if none)
- `completed` - Array of completed task IDs
- `locked` - Whether plan is locked

**Default Values**:
```json
{
  "phase": "idea",
  "active_task": null,
  "completed": [],
  "locked": false
}
```

**Updated By**:
- `/tell` - Sets phase to "idea"
- `/improve` - Sets phase to "brainstorm"
- `/write` - Sets phase to "writing"
- `/good` - Sets phase to "approved", locked to true
- `/tasks` - Sets phase to "tasks"
- `/build` - Sets active_task
- `/finished` - Updates completed, clears active_task

---

## Environment Variables

### DoPlan CLI Environment Variables

DoPlan CLI itself doesn't require environment variables. However, you can set:

#### `DOPLAN_DEBUG`

**Purpose**: Enable debug mode.

**Usage**:
```bash
export DOPLAN_DEBUG=1
doplan
```

**Effect**: Enables verbose logging and debug output.

---

## CLI Flags

### DoPlan CLI Flags

#### `--version` or `-v`

**Purpose**: Show version information.

**Usage**:
```bash
doplan --version
```

**Output**:
```
doplan version 1.0.4
```

---

#### `--help` or `-h`

**Purpose**: Show help information.

**Usage**:
```bash
doplan --help
```

---

## Project Settings

### Project Name

**Setting**: Project name

**Location**: Project directory name

**Format**: 
- Lowercase letters, numbers, and hyphens
- No spaces or special characters

**Example**: `my-awesome-app`

---

### IDE Choice

**Setting**: Preferred IDE

**Location**: IDE-specific config files

**Options**:
- Cursor
- Claude Code
- Antigravity
- Windsurf
- Cline
- OpenCode

**Effect**: Generates IDE-specific configuration files.

---

## Agent Configuration

### Agent Definitions

**Location**: `.cursor/agents/*.md`

**Format**: Markdown files with:
- Role description
- System prompt
- Responsibilities
- Current project context

**Customization**: Edit agent files directly.

**Example Structure**:
```markdown
# Agent Name

## Role
Brief role description

## System Prompt
Detailed persona and responsibilities

## Responsibilities
- Responsibility 1
- Responsibility 2
```

---

## Command Configuration

### Command Definitions

**Location**: `.cursor/commands/*.md`

**Format**: Markdown files with:
- Trigger pattern
- Action description
- Agent involvement
- Files read/modified

**Customization**: Edit command files directly.

**Example Structure**:
```markdown
# Command Name

## Trigger
Command trigger pattern

## Action
What happens when command is used

## Agent Involvement
- Agent 1
- Agent 2
```

---

## Rules Configuration

### Rules Library

**Location**: `.cursor/rules/library/`

**Structure**: 15 categories of rules

**Customization**:
- Edit existing rules
- Add new rules
- Create custom categories

**Loading Rules**: Use `/load` command:
```bash
/load @library/04-frameworks/nextjs.md
```

---

## GitHub Workflow Configuration

### CI Workflow

**Location**: `.github/workflows/ci.yml`

**Configuration**: YAML file defining CI pipeline

**Default**: Runs on all branches

**Customization**: Edit YAML file

---

### Release Workflow

**Location**: `.github/workflows/release.yml`

**Configuration**: YAML file defining release process

**Default**: Triggers on version tags

**Customization**: Edit YAML file

---

### Changelog Workflow

**Location**: `.github/workflows/changelog.yml`

**Configuration**: YAML file defining changelog updates

**Default**: Auto-updates on CHANGELOG.md changes

**Customization**: Edit YAML file

---

## IDE Configuration

### Cursor Configuration

**File**: `.cursorrules`

**Purpose**: Cursor IDE rules and configuration

**Format**: Markdown

**Customization**: Edit directly

---

### Claude Code Configuration

**File**: `CLAUDE.md`

**Purpose**: Claude Code IDE configuration

**Format**: Markdown

**Customization**: Edit directly

---

### Other IDEs

Similar configuration files for:
- Antigravity: `.antigravity/`
- Windsurf: `.windsurf/`
- Cline: `.cline/`
- OpenCode: `.opencode/`

---

## Configuration Validation

### State Validation

The `active_state.json` file is validated when:
- Commands are executed
- State transitions occur
- Tasks are generated

**Validation Rules**:
- Phase must be valid
- Active task must exist in TASKS.md (if set)
- Completed tasks must exist in TASKS.md

---

## Configuration Examples

### Example 1: Initial State

```json
{
  "phase": "idea",
  "active_task": null,
  "completed": [],
  "locked": false
}
```

### Example 2: After Planning

```json
{
  "phase": "approved",
  "active_task": null,
  "completed": [],
  "locked": true
}
```

### Example 3: During Development

```json
{
  "phase": "building",
  "active_task": "1.2",
  "completed": ["1.1"],
  "locked": true
}
```

### Example 4: After Tasks Complete

```json
{
  "phase": "building",
  "active_task": null,
  "completed": ["1.1", "1.2", "1.3", "2.1"],
  "locked": true
}
```

---

## Default Values

### Project State Defaults

- **Phase**: `"idea"`
- **Active Task**: `null`
- **Completed**: `[]`
- **Locked**: `false`

### Agent Defaults

- 18 agents generated by default
- Each agent has default persona
- Customizable via agent files

### Command Defaults

- 14 commands generated by default
- Each command has default behavior
- Customizable via command files

### Rules Defaults

- 15 categories of rules
- 1000+ rules embedded
- All extracted to `.cursor/rules/library/`

---

## Configuration Best Practices

### 1. Version Control

Commit all configuration files:
- `.plan/active_state.json`
- `.cursor/agents/*.md`
- `.cursor/commands/*.md`
- `.cursor/rules/library/**`
- `.github/workflows/*.yml`

### 2. Don't Commit

Don't commit:
- IDE-specific local settings
- Personal customizations (unless shared)
- Temporary files

### 3. Document Customizations

Document any customizations:
- Update README.md
- Add comments in config files
- Document in project docs

### 4. Test Changes

Test configuration changes:
- Verify state transitions
- Test commands
- Check agent behavior

---

## Related Pages

- [Project Structure](Project-Structure) - Project layout
- [Workflow Guide](Workflow) - How configuration is used
- [Customization Guide](Customization) - Customizing configuration
- [API Reference](API-Reference) - Configuration API
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

