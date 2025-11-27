# Command Definitions

This directory contains markdown definitions for all slash commands in the DoPlan CLI system.

## Structure

Commands are organized by category:
- `core/` - Essential daily workflow commands
- `tools/` - Utility and helper commands
- `optimize/` - Optimization commands

## Markdown Format

Each command markdown file uses YAML frontmatter for metadata, followed by the full action description:

**Frontmatter (YAML)**:
- `name`: Command name (e.g., "hello", "tell")
- `category`: Command category (core, tools, optimize)
- `trigger`: Trigger pattern (e.g., "/hello [<subcommand>]")
- `description`: Brief description
- `agentInvolvement`: Array of agent names involved
- `examples`: Array of example usage
- `filesRead`: Array of files this command reads
- `filesModified`: Array of files this command modifies
- Optional: `githubAutomation`, `requirements`, `notes`, `customize`, `options`, `offlineSafety`

**Markdown Content**:
- Full action description with detailed instructions

## Usage

These markdown files are loaded by the generator and converted to command definitions in the `.cursor/commands/` directory (and other IDE-specific locations).

