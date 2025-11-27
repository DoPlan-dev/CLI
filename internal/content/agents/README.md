# Agent Definitions

This directory contains JSON definitions for all AI agents in the DoPlan hierarchical agency system.

## Structure

Agents are organized by category:
- `leadership/` - Top-level management roles
- `engineering/` - Technical leadership and engineering roles
- `product/` - Product management roles
- `design/` - Design and UX roles
- `quality/` - QA and testing roles
- `release/` - Release and growth roles
- `documentation/` - Documentation roles

## Markdown Format

Each agent markdown file uses YAML frontmatter for metadata, followed by the full markdown content:

**Frontmatter (YAML)**:
- `name`: Agent display name
- `role`: Brief role description
- `reportsTo`: Manager in hierarchy (empty for top level)
- `manages`: Array of subordinate agent names (empty array if none)
- `responsibilities`: Array of key responsibilities
- `category`: Category folder name

**Markdown Content**:
- Full markdown document with sections: Role, System Prompt, Responsibilities, Reports To, Manages
- The System Prompt is extracted from the "## System Prompt" section

## Usage

These JSON files are loaded by the generator and converted to markdown files in the `.cursor/agents/` directory (and other IDE-specific locations).

