# API Reference

Complete API reference for DoPlan CLI and generated projects.

---

## Overview

DoPlan CLI provides several APIs:
- **Command Line API** - CLI commands and flags
- **Generated Files API** - Structure and format of generated files
- **Agent System API** - Agent definitions and interactions
- **Rules System API** - Rules library structure
- **Extension API** - Customization points

---

## Command Line API

### Command: `doplan`

**Description**: Main DoPlan CLI command.

**Usage**:
```bash
doplan [flags]
```

**Flags**:
- `--version, -v` - Print version information and exit
- `--help, -h` - Show help information

**Examples**:
```bash
doplan              # Run interactive wizard
doplan --version    # Show version
doplan --help       # Show help
```

**Exit Codes**:
- `0` - Success
- `1` - Error

---

### Command: `npx @doplan-dev/cli`

**Description**: Zero-install wrapper command.

**Usage**:
```bash
npx @doplan-dev/cli
```

**Behavior**:
1. Downloads binary for your platform
2. Runs DoPlan CLI
3. No installation required

---

## Generated Files API

### Project Structure

**Format**: Directory tree

**Structure**:
```
project-name/
├── .cursor/
│   ├── agents/
│   ├── commands/
│   └── rules/
├── .plan/
│   ├── 00_System/
│   ├── TASKS.md
│   └── active_state.json
├── .github/
│   └── workflows/
├── src/
└── [root files]
```

---

### active_state.json

**Format**: JSON

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
- `phase` (string, required) - Current workflow phase
- `active_task` (string|null, optional) - Active task ID
- `completed` (array, required) - Completed task IDs
- `locked` (boolean, required) - Plan lock status

**Example**:
```json
{
  "phase": "building",
  "active_task": "1.2",
  "completed": ["1.1"],
  "locked": true
}
```

---

### Agent Files

**Format**: Markdown

**Structure**:
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

**Location**: `.cursor/agents/{name}.md`

**Required Fields**:
- `# Agent Name` - Agent title
- `## Role` - Role description
- `## System Prompt` - Agent persona

---

### Command Files

**Format**: Markdown

**Structure**:
```markdown
# Command Name

## Trigger
Command trigger pattern

## Action
What happens when command is used

## Agent Involvement
- Agent 1
- Agent 2

## Files Read
- file1.md
- file2.md

## Files Modified
- file1.md
- file2.md
```

**Location**: `.cursor/commands/{name}.md`

**Required Fields**:
- `# Command Name` - Command title
- `## Trigger` - Trigger pattern
- `## Action` - Action description

---

### Planning Documents

**Format**: Markdown

**Files**:
- `IDEA.md` - Project idea
- `BRAINSTORM.md` - Brainstorming results
- `PRD.md` - Product Requirements Document
- `ARCHITECTURE.md` - Technical Architecture
- `DESIGN_SYSTEM.md` - Design System

**Location**: `.plan/00_System/`

**Structure**: Free-form markdown (no strict format)

---

### TASKS.md

**Format**: Markdown

**Structure**:
```markdown
# Tasks

## Phase 1: Description
**Goal**: Phase goal

### Task ID
**ID**: X.Y
**Description**: Task description
**Status**: Pending|In Progress|Completed
```

**Location**: `.plan/TASKS.md`

---

## Agent System API

### Agent Definition

**Format**: Markdown file

**Required Sections**:
- `# Agent Name` - Agent title
- `## Role` - Role description
- `## System Prompt` - Agent persona

**Optional Sections**:
- `## Responsibilities` - List of responsibilities
- `## Current Project Context` - Project-specific context

---

### Agent Activation

**Trigger**: Commands activate agents based on `Agent Involvement` section.

**Process**:
1. Command file specifies agents
2. Agents are loaded from `.cursor/agents/`
3. Agents receive context
4. Agents execute their role

---

## Rules System API

### Rule File Format

**Format**: Markdown

**Structure**:
```markdown
# Rule Title

## Overview
Rule description

## Guidelines
- Guideline 1
- Guideline 2

## Examples
### Good
```code
// Good example
```

### Bad
```code
// Bad example
```
```

**Location**: `.cursor/rules/library/{category}/{name}.md`

---

### Rule Categories

**Structure**: 15 numbered categories:
- `01-core-workflow/`
- `02-ai-agents/`
- `03-languages/`
- `04-frameworks/`
- `05-ui-libraries/`
- `06-cloud-infrastructure/`
- `07-databases/`
- `08-testing/`
- `09-devops-ci-cd/`
- `10-code-quality/`
- `11-documentation/`
- `12-security/`
- `13-development-practices/`
- `14-mcp-tools/`
- `15-project-specific/`

---

### Loading Rules

**Command**: `/load @library/{path}`

**Format**: `@library/{category}/{file}.md`

**Examples**:
```bash
/load @library/04-frameworks/nextjs.md
/load @library/07-databases/postgresql.md
```

---

## Extension API

### Custom Agents

**Location**: `.cursor/agents/`

**Process**:
1. Create new `.md` file
2. Follow agent file format
3. Reference in commands if needed

**Example**:
```markdown
# Custom Agent

## Role
Custom role description

## System Prompt
Custom agent persona
```

---

### Custom Commands

**Location**: `.cursor/commands/`

**Process**:
1. Create new `.md` file
2. Follow command file format
3. Define trigger and action

**Example**:
```markdown
# Custom Command

## Trigger
/custom

## Action
Custom action description

## Agent Involvement
- Custom Agent
```

---

### Custom Rules

**Location**: `.cursor/rules/library/`

**Process**:
1. Create rule file in appropriate category
2. Follow rule file format
3. Load with `/load` command

**Example**:
```markdown
# Custom Rule

## Overview
Custom rule description

## Guidelines
- Guideline 1
```

---

## Request/Response Formats

### Command Execution

**Request**: User types command in IDE

**Response**: Command executes, files modified

**No explicit API** - Commands work through IDE integration

---

### State Updates

**Request**: Command execution

**Response**: `active_state.json` updated

**Format**: JSON

---

## Error Codes

### CLI Errors

**Exit Code 1**: General error
- Invalid arguments
- File system errors
- Generation failures

**Error Messages**: Printed to stderr

---

### Project Errors

**No explicit error codes** - Errors handled through:
- Command responses
- File validation
- State validation

---

## Versioning

### DoPlan CLI Version

**Format**: Semantic versioning (MAJOR.MINOR.PATCH)

**Example**: `1.0.4`

**Check Version**:
```bash
doplan --version
```

---

### API Versioning

**Current Version**: 1.0

**Breaking Changes**: Documented in [Migration Guide](Migration-Guide)

**Backward Compatibility**: Maintained within major versions

---

## Examples

### Example 1: Reading State

```bash
# Read active_state.json
cat .plan/active_state.json
```

**Response**:
```json
{
  "phase": "building",
  "active_task": "1.2",
  "completed": ["1.1"],
  "locked": true
}
```

---

### Example 2: Custom Agent

Create `.cursor/agents/custom_agent.md`:
```markdown
# Custom Agent

## Role
Custom role

## System Prompt
Custom agent persona
```

---

### Example 3: Custom Command

Create `.cursor/commands/custom.md`:
```markdown
# Custom Command

## Trigger
/custom

## Action
Custom action

## Agent Involvement
- Custom Agent
```

---

### Example 4: Custom Rule

Create `.cursor/rules/library/15-project-specific/custom.md`:
```markdown
# Custom Rule

## Overview
Custom rule

## Guidelines
- Guideline 1
```

Load with:
```bash
/load @library/15-project-specific/custom.md
```

---

## Related Pages

- [Commands Reference](Commands) - Command details
- [Project Structure](Project-Structure) - File structure
- [Configuration Reference](Configuration) - Configuration API
- [Customization Guide](Customization) - Extension examples
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

