# Documentation Index

This document provides a comprehensive index of all available documentation in this project for use with Cursor's @docs feature.

<!-- Generated: 2025-01-27 -->

## Quick Reference

To use documentation in Cursor:
- Type `@docs` to see available documentation sources
- Type `@filename` to reference specific files
- Type `@folder` to include entire directories

## Documentation Structure

### Core Documentation

#### `docs/README.md`
**Purpose**: Documentation directory structure and organization rules
**Content**: 
- Directory structure explanation
- Folder purposes and contracts
- Documentation organization rules

#### `docs/overview/README.md`
**Purpose**: Project overview and getting started guide
**Content**:
- Project quick start
- Command reference
- Project structure
- Workflow guide
- Agent hierarchy overview

#### `docs/history/CHANGELOG.md`
**Purpose**: Project changelog and version history
**Content**:
- Version history
- Change categories (Added, Changed, Deprecated, Removed, Fixed, Security)
- Semantic versioning format

### Directory Structure

```
docs/
├── README.md              # Documentation index and structure
├── DOCS_INDEX.md          # This file - comprehensive docs index
├── overview/              # Project overview and guides
│   └── README.md          # Main project documentation
├── references/            # Command cheat sheets and examples
├── tutorials/             # Guided walkthroughs and training
├── history/               # Changelog and historical records
│   └── CHANGELOG.md       # Version history
├── ops/                   # Operations and deployment docs
└── research/              # Research notes and experiments
```

## Recommended @Docs Configuration

For optimal Cursor @docs integration, add these sources in Settings → Features → Docs:

1. **`docs/`** - Entire documentation directory
2. **`docs/overview/`** - Project overview and guides
3. **`docs/history/`** - Changelog and historical records
4. **`.cursorrules`** - Project configuration and rules (if not already included)

## Documentation Categories

### Overview Documentation
- **Location**: `docs/overview/`
- **Purpose**: Project overview, mission, agent hierarchy
- **Key Files**:
  - `README.md` - Main project documentation

### Reference Documentation
- **Location**: `docs/references/`
- **Purpose**: Command cheat sheets and reusable examples
- **Status**: Directory exists, ready for content

### Tutorial Documentation
- **Location**: `docs/tutorials/`
- **Purpose**: Guided walkthroughs and training notes
- **Status**: Directory exists, ready for content

### History Documentation
- **Location**: `docs/history/`
- **Purpose**: Changelog, meeting notes, audit logs
- **Key Files**:
  - `CHANGELOG.md` - Project changelog

### Operations Documentation
- **Location**: `docs/ops/`
- **Purpose**: Release checklists, security reviews, CI/CD docs
- **Status**: Directory exists, ready for content

### Research Documentation
- **Location**: `docs/research/`
- **Purpose**: Brainstorms, interviews, experiment logs
- **Status**: Directory exists, ready for content

## Using Documentation in Cursor

### Basic Usage
1. **Reference entire directory**: `@docs` or `@docs/overview`
2. **Reference specific file**: `@docs/overview/README.md`
3. **Reference folder**: `@docs/history`

### Common Workflows

#### Getting Started
```
@docs/overview/README.md How do I get started with this project?
```

#### Understanding Project Structure
```
@docs What is the project structure and where should I add new features?
```

#### Checking History
```
@docs/history/CHANGELOG.md What changes were made in recent versions?
```

#### Understanding Commands
```
@docs/overview/README.md What commands are available and how do I use them?
```

## Documentation Best Practices

1. **Keep docs current**: Update documentation as code changes
2. **LLM-optimized format**: Structure docs for both humans and AI
   - Include concrete file references
   - Use clear section headings
   - Provide examples with file paths
3. **Timestamp tracking**: Include last-update timestamps in docs
4. **Link, don't duplicate**: Reference existing docs rather than copying content

## Maintenance

To check if documentation is up to date:
- Ask Cursor: "Are any docs outdated based on recent changes?"
- Review timestamps in documentation files
- Update this index when new documentation is added

## Adding New Documentation

When adding new documentation:

1. **Place in appropriate category**: Use the existing directory structure
2. **Update this index**: Add new files to the relevant section
3. **Follow naming conventions**: Use clear, descriptive filenames
4. **Include timestamps**: Add generation timestamps to new docs
5. **Link from README**: Update `docs/README.md` if adding new categories

## Related Resources

- **Project Rules**: `.cursorrules` - Project-wide configuration
- **Agent Definitions**: `.cursor/agents/` - AI agent personalities
- **Command Definitions**: `.cursor/commands/` - Available commands
- **Rules Library**: `.cursor/rules/` - Stack-specific rules

---

**Last Updated**: 2025-01-27
**Maintained By**: Documentation Lead Agent

