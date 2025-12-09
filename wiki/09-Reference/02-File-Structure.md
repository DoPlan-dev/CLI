# File Structure

Understanding the `.do/` directory structure and project organization.

---

## 📁 Project Structure

```
your-project/
├── .do/
│   ├── system/
│   │   ├── IDEA.md
│   │   ├── BRAINSTORM.md
│   │   ├── REFINEMENTS.md
│   │   ├── PRD.md
│   │   ├── ARCHITECTURE.md
│   │   ├── DESIGN_SYSTEM.md
│   │   ├── time-tracker.jsonl
│   │   └── history/
│   │       ├── active_state.json
│   │       └── state-*.json
│   │
│   ├── plan/
│   │   ├── TASKS.md
│   │   ├── 01-Foundation/
│   │   │   ├── feature-auth/
│   │   │   │   ├── design.md
│   │   │   │   ├── plan.md
│   │   │   │   ├── tasks.md
│   │   │   │   ├── prompts.md
│   │   │   │   └── github.md
│   │   │   └── feature-database/
│   │   ├── 02-Core/
│   │   └── 03-Enhancement/
│   │
│   └── core/
│       └── (templates)
│
└── docs/
    ├── foundation/
    ├── features/
    └── reference/
```

---

## 📂 `.do/system/` - System Files

### Core Documents

- **IDEA.md** - Your project idea
- **BRAINSTORM.md** - Discovery meeting results
- **REFINEMENTS.md** - Refinements and enhancements
- **PRD.md** - Product Requirements Document
- **ARCHITECTURE.md** - System architecture
- **DESIGN_SYSTEM.md** - Design system

### Tracking Files

- **time-tracker.jsonl** - Time tracking data (JSONL format)

### History Directory

- **active_state.json** - Current project state
- **state-*.json** - State snapshots

---

## 📂 `.do/plan/` - Execution Plan

### Master Task List

- **TASKS.md** - Complete task list organized by phases

### Phase Directories

- **01-Foundation/** - Phase 1 tasks
- **02-Core/** - Phase 2 tasks
- **03-Enhancement/** - Phase 3 tasks

### Feature Folders

Each feature has:
- **design.md** - Design specifications
- **plan.md** - Implementation plan
- **tasks.md** - Feature tasks
- **prompts.md** - AI prompts
- **github.md** - GitHub workflow

---

## 📂 `.do/core/` - Templates

Template files for:
- Feature structure
- Documentation
- Code templates

---

## 📂 `docs/` - Documentation

### Foundation

- **IDEA.md** - Project idea (synced)
- **BRAINSTORM.md** - Discovery results (synced)
- **PRD.md** - Requirements (synced)
- **ARCHITECTURE.md** - Architecture (synced)
- **DESIGN_SYSTEM.md** - Design system (synced)

### Features

- **Feature documentation** - Synced from plan
- **Task documentation** - Implementation details

### Reference

- **Quick reference** - Command cheat sheet
- **Examples** - Usage examples
- **Guides** - How-to guides

---

## 🎯 File Purposes

### IDEA.md

**Purpose**: Your complete project idea
**Created**: By `/do` command
**Updated**: During ideation and refinement

### BRAINSTORM.md

**Purpose**: Discovery meeting results
**Created**: By `/do` meeting phase
**Content**: Questions, answers, insights

### REFINEMENTS.md

**Purpose**: Refinements and enhancements
**Created**: By `/do` refinement phase
**Content**: Improvements, suggestions

### TASKS.md

**Purpose**: Master task list
**Created**: By `/plan` command
**Updated**: When `/dev` auto-completes tasks

### active_state.json

**Purpose**: Current project state
**Created**: Automatically
**Updated**: On every state change

### time-tracker.jsonl

**Purpose**: Time tracking data
**Created**: Automatically
**Updated**: On every command execution

---

## 💡 File Management

### Reading Files

```bash
# Read IDEA.md
cat .do/system/IDEA.md

# Read TASKS.md
cat .do/plan/TASKS.md

# Read active state
cat .do/system/history/active_state.json
```

### Editing Files

**Manual editing**:
- Edit IDEA.md, BRAINSTORM.md directly
- Update TASKS.md manually (not recommended)
- Modify state files (advanced)

**Recommended**:
- Use DoPlan commands
- Let DoPlan manage files
- Manual edits only when needed

---

**Next**: [State Management](./03-State-Management.md)

