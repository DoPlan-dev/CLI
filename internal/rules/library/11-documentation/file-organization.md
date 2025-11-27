# File Organization Rules

## When to use this rule
- Creating any new markdown (.md) files
- Organizing project documentation
- Adding new documentation categories
- Reviewing file structure

## Core Principle
**All markdown files must be organized in the `docs/` folder, never in the project root.**

## Required Structure

### 1. Root Directory
The project root must remain clean. Only these files are allowed at root level:
- `README.md` - Main project documentation
- Configuration files (package.json, tsconfig.json, etc.)
- Source code files

### 2. Documentation Folder Structure
All markdown files must be placed in `docs/` with proper categorization:

```
docs/
├── README.md              # Documentation index
├── CHANGELOG.md           # Project changelog
├── CLAUDE.md              # IDE configuration docs
├── foundation/            # Global system references
│   ├── architecture.md
│   ├── design-system.md
│   └── roadmap.md
├── features/              # Feature-specific documentation
│   ├── 01_Auth_Feature/
│   │   ├── spec.md
│   │   ├── tasks.md
│   │   └── decisions.md
│   └── 02_Payment_Feature/
│       └── ...
├── release/               # Release documentation
│   ├── v1.0.0-release-notes.md
│   └── retrospectives/
├── history/               # Historical records
│   ├── prompt-logs/
│   └── git-summaries/
└── [category]/            # Additional categories as needed
```

## Folder Naming Patterns

### Standard Categories
Use these predefined categories when possible:
- `foundation/` - System-wide documentation
- `features/` - Feature-specific docs
- `release/` - Release-related documentation
- `history/` - Historical records and logs

### Feature Folder Naming
Feature folders must follow this pattern:
```
docs/features/[ID]_[Feature_Name]/
```

Examples:
- `docs/features/01_Auth_Feature/`
- `docs/features/02_Payment_Integration/`
- `docs/features/03_User_Dashboard/`

**Rules:**
- Use numeric prefix (01, 02, 03...) for ordering
- Use PascalCase or Title Case for feature names
- Separate words with underscores or hyphens
- Match the task ID from `.do/plan/TASKS.md` when applicable

### Custom Category Naming
When creating new categories, follow this pattern:
- Use lowercase letters
- Use underscores for multi-word names
- Be descriptive and specific
- Examples: `api-reference/`, `user-guides/`, `troubleshooting/`

## File Naming Conventions

### Standard File Names
- `README.md` - Directory index/overview
- `spec.md` - Feature specification
- `tasks.md` - Task checklist
- `decisions.md` - Architecture decisions
- `notes.md` - General notes
- `guide.md` - How-to guide

### Naming Pattern
```
[type]-[descriptive-name].md
```

Examples:
- `api-authentication.md`
- `database-schema.md`
- `deployment-guide.md`
- `troubleshooting-common-issues.md`

## Enforcement Rules

### 1. New File Creation
When creating a new markdown file:
1. **Always** place it in `docs/` folder
2. Choose the appropriate category subfolder
3. If category doesn't exist, create it following naming patterns
4. Update `docs/README.md` to document the new file/category

### 2. File Categorization
- **Foundation docs** → `docs/foundation/`
  - Architecture, design systems, roadmaps
  - Global references and guides
  
- **Feature docs** → `docs/features/[ID]_[Feature_Name]/`
  - Feature specifications
  - Feature-specific tasks
  - Feature decisions
  
- **Release docs** → `docs/release/`
  - Release notes
  - Launch plans
  - Retrospectives
  
- **Historical docs** → `docs/history/`
  - Prompt logs
  - Git summaries
  - Audit trails

### 3. Category Creation
When creating a new category:
1. Create folder in `docs/` following naming patterns
2. Add `README.md` in the new category explaining its purpose
3. Update `docs/README.md` to list the new category
4. Document the category's purpose and usage

## Examples

### ✅ Correct
```
docs/features/01_User_Authentication/spec.md
docs/foundation/architecture.md
docs/release/v1.0.0-notes.md
docs/history/prompt-logs/2024-01-15.md
```

### ❌ Incorrect
```
User_Authentication.md                    # Root level
docs/User_Authentication.md               # Missing category
docs/features/user-auth.md                # Wrong naming pattern
docs/01_User_Authentication/spec.md        # Missing category prefix
```

## Checklist
Before creating any markdown file:
- [ ] File will be placed in `docs/` folder
- [ ] Appropriate category subfolder is selected/created
- [ ] Folder name follows naming pattern
- [ ] File name is descriptive and follows conventions
- [ ] `docs/README.md` will be updated if adding new category
- [ ] File is not placed in project root

## Integration with Commands
When using DoPlan commands:
- `/write` - Generates docs in appropriate `docs/` subfolders
- `/plan` - Mirrors `.do/` content to `docs/foundation/`
- `/build` - Creates feature docs in `docs/features/` and (on first run) prompts the user to scaffold code with their preferred tool (the legacy `scripts/boilerplate` helper has been removed)
- `/finished` - Moves completed feature docs to appropriate locations

## Maintenance
- Regularly review `docs/` structure for consistency
- Merge duplicate categories when found
- Archive old documentation to `docs/history/`
- Keep `docs/README.md` up-to-date with current structure

