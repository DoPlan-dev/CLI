# IDE Command Formats Supported by DoPlan

DoPlan supports multiple IDEs, each with their own command format. This document explains how commands are installed for each IDE.

## Supported IDEs

1. **Cursor** - Markdown files (`.md`)
2. **Gemini CLI** - TOML files (`.toml`)
3. **Claude Code** - Markdown files with YAML frontmatter (`.md`)
4. **Codex CLI** - Markdown files with YAML frontmatter (`.md`)
5. **OpenCode** - Markdown files with YAML frontmatter (`.md`)
6. **Qwen Code** - TOML files (`.toml`) with `[command]` section

## Cursor Commands

**Format:** Markdown files (`.md`)  
**Location:** `.cursor/commands/`  
**Documentation:** [Cursor Commands](https://cursor.com/docs/agent/chat/commands)

### Structure
```markdown
# Command Name

## Overview
Description of what the command does

## Workflow
Step-by-step instructions
```

### Usage
- Type `/` in Cursor chat
- Commands appear as `/discuss`, `/plan`, etc.
- File `discuss.md` becomes command `/discuss`

### Example
File: `.cursor/commands/discuss.md`
```markdown
# Discuss

## Overview
Start idea discussion and refinement workflow...

## Workflow
1. Ask comprehensive questions...
```

## Gemini CLI Commands

**Format:** TOML files (`.toml`)  
**Location:** `.gemini/commands/` (project) or `~/.gemini/commands/` (global)  
**Documentation:** [Gemini CLI Custom Commands](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/custom-commands.md)

### Structure
```toml
description = "Brief description of the command"

prompt = """
The full prompt that will be sent to the model.
Supports {{args}} for argument injection.
Supports !{...} for shell commands.
Supports @{...} for file injection.
"""
```

### Usage
- Type `/` in Gemini CLI
- Commands appear as `/discuss`, `/plan`, etc.
- File `discuss.toml` becomes command `/discuss`
- Use `/discuss "my idea"` to pass arguments

### Example
File: `.gemini/commands/discuss.toml`
```toml
description = "Start idea discussion and refinement workflow"

prompt = """
Start the DoPlan idea discussion workflow...
"""
```

## Claude Code Commands

**Format:** Markdown files with optional YAML frontmatter (`.md`)  
**Location:** `.claude/commands/` (project) or `~/.claude/commands/` (global)  
**Documentation:** [Claude Code Slash Commands](https://code.claude.com/docs/en/slash-commands)

### Structure
```markdown
---
description: Brief description shown in popup
argument-hint: [KEY="<value>"] (optional)
allowed-tools: Bash(git add:*), Read, Grep (optional)
model: claude-sonnet-4-5-20250929 (optional)
disable-model-invocation: false (optional)
---

The prompt content that defines the command behavior.

Supports:
- $ARGUMENTS - All arguments
- $1, $2, etc. - Positional arguments
- !`command` - Bash command execution
- @path/to/file - File references
```

### Usage
- Type `/` in Claude Code
- Commands appear as `/discuss`, `/plan`, etc.
- File `discuss.md` becomes command `/discuss`
- Use `/discuss "my idea"` to pass arguments
- Supports `$ARGUMENTS`, `$1`, `$2`, etc. for arguments
- Use `!` prefix for bash commands: `!`git status`
- Use `@` prefix for file references: `@doplan/PRD.md`

### Example
File: `.claude/commands/discuss.md`
```markdown
---
description: Start idea discussion and refinement workflow
---

Start the DoPlan idea discussion workflow...
```

## Codex CLI Commands

**Format:** Markdown files with YAML frontmatter (`.md`)  
**Location:** `.codex/prompts/` (project) or `~/.codex/prompts/` (global)  
**Documentation:** [Codex Slash Commands](https://developers.openai.com/codex/guides/slash-commands)

### Structure
```markdown
---
description: Brief description shown in popup
argument-hint: [KEY="<value>"] [KEY2=<value>]
---

The prompt content that will be executed.

Supports placeholders:
- $1-$9: Positional arguments
- $ARGUMENTS: All arguments
- $KEY: Named arguments (uppercase)
- $$: Literal dollar sign
```

### Usage
- Type `/` in Codex CLI
- Commands appear as `/prompts:discuss`, `/prompts:plan`, etc.
- File `discuss.md` becomes command `/prompts:discuss`
- Use `/prompts:discuss TOPIC="my idea"` to pass arguments

### Example
File: `.codex/prompts/discuss.md`
```markdown
---
description: Start idea discussion and refinement workflow
argument-hint: [TOPIC="<idea description>"]
---

Start the DoPlan idea discussion workflow...

$ARGUMENTS
```

## OpenCode Commands

**Format:** Markdown files with YAML frontmatter (`.md`)  
**Location:** `.opencode/command/` (project) or `~/.config/opencode/command/` (global)  
**Documentation:** [OpenCode Commands](https://opencode.ai/docs/commands/)

### Structure
```markdown
---
description: Brief description shown in popup
agent: build (optional)
model: anthropic/claude-3-5-sonnet-20241022 (optional)
subtask: false (optional)
---

The prompt content that defines the command behavior.

Supports:
- $ARGUMENTS - All arguments
- $1, $2, etc. - Positional arguments
- !`command` - Bash command execution
- @path/to/file - File references
```

### Usage
- Type `/` in OpenCode TUI
- Commands appear as `/discuss`, `/plan`, etc.
- File `discuss.md` becomes command `/discuss`
- Use `/discuss "my idea"` to pass arguments
- Supports `$ARGUMENTS`, `$1`, `$2`, etc. for arguments
- Use `!` prefix for bash commands: `!`git status`
- Use `@` prefix for file references: `@doplan/PRD.md`

### Example
File: `.opencode/command/discuss.md`
```markdown
---
description: Start idea discussion and refinement workflow
---

Start the DoPlan idea discussion workflow...

## Context
- Current project state: @.cursor/config/doplan-state.json
- Existing idea notes: @doplan/idea-notes.md

$ARGUMENTS
```

## Qwen Code Commands

**Format:** TOML files (`.toml`) with `[command]` section  
**Location:** `.qwen/commands/` (project) or `~/.qwen/commands/` (global)

### Structure
```toml
[command]
name = "command-name"
description = "Brief description of what the command does"
prompt = """
The prompt content that defines the command behavior.
Multi-line strings are supported.
"""
```

### Usage
- Type `/` in Qwen Code
- Commands appear as `/discuss`, `/plan`, etc.
- File `discuss.toml` becomes command `/discuss`
- Supports namespacing: subdirectories become `/namespace:command`
- Example: `.qwen/commands/my_category/my_command.toml` becomes `/my_category:my_command`

### Example
File: `.qwen/commands/discuss.toml`
```toml
[command]
name = "discuss"
description = "Start idea discussion and refinement workflow"
prompt = """Start the DoPlan idea discussion workflow...

## Workflow
1. Ask comprehensive questions about the idea
2. Suggest improvements and enhancements
...
"""
```

## Command Installation

When you run `doplan install`, DoPlan:

1. **Detects your IDE choice** from the installation menu
2. **Creates appropriate directory**:
   - Cursor: `.cursor/commands/`
   - Gemini: `.gemini/commands/`
   - Claude Code: `.claude/commands/`
   - Codex: `.codex/prompts/`
   - OpenCode: `.opencode/command/`
   - Qwen Code: `.qwen/commands/`
3. **Generates commands** in the correct format for your IDE
4. **Creates 8 commands**:
   - `discuss` - Start idea discussion
   - `refine` - Refine idea
   - `generate` - Generate PRD and contracts
   - `plan` - Create phase structure
   - `dashboard` - Show dashboard
   - `implement` - Start feature implementation
   - `next` - Get next action recommendation
   - `progress` - Update progress tracking

## Switching IDEs

If you want to use a different IDE:

1. Run `doplan install` again
2. Select the new IDE
3. DoPlan will create commands in the new format
4. Old commands remain but won't interfere

## Command Features by IDE

| Feature | Cursor | Gemini | Claude Code | Codex | OpenCode | Qwen Code |
|---------|--------|--------|-------------|-------|----------|-----------|
| Format | Markdown | TOML | Markdown + YAML | Markdown + YAML | Markdown + YAML | TOML |
| Arguments | Via prompt | `{{args}}` | `$ARGUMENTS`, `$1-$9` | `$ARGUMENTS`, `$KEY` | `$ARGUMENTS`, `$1-$9` | Via prompt |
| Shell Commands | No | `!{...}` | `!`command`` | Via Codex tools | `!`command`` | No |
| File Injection | No | `@{...}` | `@path/to/file` | Via `/mention` | `@path/to/file` | No |
| Description | In markdown | `description` field | YAML frontmatter | YAML frontmatter | YAML frontmatter | `description` in TOML |
| Argument Hints | No | No | `argument-hint` field | `argument-hint` field | No | No |
| Tools Config | No | No | `allowed-tools` in YAML | No | No | No |
| Model Config | No | No | `model` in YAML | `model` in YAML | `model` in YAML | No |
| Agent Config | No | No | No | No | `agent` in YAML | No |
| Subtask Config | No | No | No | No | `subtask` in YAML | No |
| Namespacing | No | Yes (subdirs) | Yes (subdirs) | No | No | Yes (subdirs) |
| SlashCommand Tool | No | No | Yes (programmatic) | No | No | No |

## References

- [Cursor Commands Documentation](https://cursor.com/docs/agent/chat/commands)
- [Gemini CLI Custom Commands](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/custom-commands.md)
- [Claude Code Slash Commands](https://code.claude.com/docs/en/slash-commands)
- [Codex Slash Commands](https://developers.openai.com/codex/guides/slash-commands)
- [OpenCode Commands](https://opencode.ai/docs/commands/)

