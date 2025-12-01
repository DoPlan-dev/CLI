# Command Definitions

This directory contains markdown definitions for all slash commands in the DoPlan CLI system.

## Structure

Commands are organized by category:
- `onboarding/` - Welcome and project setup commands
- `developing/` - Development workflow commands
- `system/` - System control and management commands

## Commands

### Onboarding
- `hey.md` - Welcome, tutorial, and command introductions
- `do.md` - Capture project idea, conduct meeting, and refine

### Developing
- `plan.md` - Generate documents, content, execution plan, scaffold phases, and manage tasks
- `dev.md` - Start development workflow for a feature

### System
- `sys.md` - System control panel (engagement, performance, backup, etc.)

## Markdown Format

Each command markdown file uses YAML frontmatter for metadata, followed by the full action description:

**Frontmatter (YAML)**:
- `name`: Command name (e.g., "hey", "do", "plan", "dev", "sys")
- `category`: Command category (onboarding, developing, system)
- `trigger`: Trigger pattern (e.g., "/hey [<subcommand>]")
- `description`: Brief description
- `agentInvolvement`: Array of agent names involved
- `examples`: Array of example usage
- `filesRead`: Array of files this command reads
- `filesModified`: Array of files this command modifies
- Optional: `githubAutomation`, `requirements`, `notes`, `customize`, `options`, `offlineSafety`

**Markdown Content**:
- Full action description with detailed instructions

## Usage

These markdown files are loaded by the generator and converted to command definitions in the `.do/core/commands/` directory (and other IDE-specific locations).
