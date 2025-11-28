# Command Overview

A comprehensive overview of all DoPlan commands, organized by category and purpose.

---

## 📋 Command Categories

### 🎯 Workflow Commands (Core)

These are the commands you'll use daily in your development workflow.

| Command | Purpose | When to Use |
|---------|---------|-------------|
| `/hey` | Onboarding & tutorial | First time, or when you need help |
| `/do` | Capture idea & discovery | Starting a new project or feature |
| `/plan` | Generate execution plan | After capturing your idea |
| `/dev` | Start development | When ready to code |
| `/done` | Complete task | When task is finished |
| `/status` | View progress | Anytime to check status |

### ⚙️ System Commands (Control)

These commands give you control over DoPlan's systems and settings.

| Command | Purpose | When to Use |
|---------|---------|-------------|
| `/sys` | System control panel | View system tree |
| `/sys engagement` | Engagement dashboard | View achievements & progress |
| `/sys role` | Role management | Manage permissions |
| `/sys security` | Security settings | Security tests & audits |
| `/sys control` | System control | Enable/disable features |

---

## 🎯 Workflow Commands Deep Dive

### `/hey` - Welcome & Onboarding

**Purpose**: Interactive onboarding, tutorial, and command introductions

**When to use**:
- First time using DoPlan
- Need a refresher on commands
- Want to see system overview
- Need reference materials

**What it does**:
- Welcome message (personalized if returning)
- Interactive tutorial
- System overview
- Agent hierarchy explanation
- Command walkthrough
- Test drive mode
- Creates reference materials

**Example**:
```bash
/hey
# → Interactive tutorial starts
```

---

### `/do` - Idea Capture & Discovery

**Purpose**: Capture your project idea through iterative conversation

**When to use**:
- Starting a new project
- Adding a new feature idea
- Fast-tracking with detailed prompt
- Getting AI-suggested ideas

**Subcommands**:
- `/do` - Full ideation workflow (default)
- `/do feature` - Add single feature idea
- `/do now` - Fast-track with prompt/PRD
- `/do i'm lucky` - Get AI-suggested ideas

**What it does**:
- Iterative idea capture
- Discovery meeting (automatic)
- Refinement phase
- Generates IDEA.md, BRAINSTORM.md, REFINEMENTS.md
- Updates memory card

**Example**:
```bash
/do
# → Interactive idea capture starts
# → Then discovery meeting
# → Then refinement
```

---

### `/plan` - Generate Execution Plan

**Purpose**: Create structured execution plan from your idea

**When to use**:
- After capturing your idea with `/do`
- When IDEA.md and BRAINSTORM.md exist

**What it does**:
- Reads IDEA.md and BRAINSTORM.md
- Generates TASKS.md with organized phases
- Creates phase directories (01-Foundation, 02-Core, etc.)
- Generates feature folders with templates
- Syncs documentation to docs/
- Integrates with engagement system

**Example**:
```bash
/plan
# → Generates complete execution plan
```

---

### `/dev` - Start Development

**Purpose**: Begin development workflow for a task

**When to use**:
- Ready to start coding
- After plan is generated
- Starting next task

**What it does**:
- Finds next available task (or specific task)
- Creates/checks out Git branch
- Syncs documentation
- Starts time tracking
- Updates active state
- Shows personalized message
- Integrates with engagement system

**Example**:
```bash
/dev              # Start next task
/dev --feature "auth"  # Start specific feature
```

---

### `/done` - Complete Task

**Purpose**: Mark task complete with auto-commit and push

**When to use**:
- Task is finished
- Ready to commit and push
- Moving to next task

**What it does**:
- Verifies active branch
- Checks dependencies
- Marks task complete in TASKS.md
- Updates state
- Creates state snapshot
- **Auto-commits** (conventional format)
- **Auto-pushes** to remote
- Updates changelog
- Suggests PR creation
- Checks achievements/challenges
- Displays task duration

**Example**:
```bash
/done
# → Task marked complete
# → Auto-committed and pushed
# → Achievements checked
```

---

### `/status` - View Progress

**Purpose**: Check project progress and status

**When to use**:
- Anytime to see progress
- After completing tasks
- Before starting new work

**What it does**:
- Reads TASKS.md
- Reads active_state.json
- Calculates progress
- Shows current phase
- Displays task statistics
- Shows state deltas
- Next task suggestion

**Example**:
```bash
/status
# → Shows progress dashboard
```

---

## ⚙️ System Commands Deep Dive

### `/sys` - System Control Panel

**Purpose**: Access system settings and controls

**When to use**:
- View system tree
- Access system features
- Manage settings

**Subcommands**:
- `/sys engagement` - Engagement dashboard
- `/sys role` - Role management
- `/sys security` - Security settings
- `/sys control` - System control

**Example**:
```bash
/sys
# → Shows system tree with suggestions
```

---

### `/sys engagement` - Engagement Dashboard

**Purpose**: View comprehensive engagement statistics

**When to use**:
- Check your score and achievements
- See relationship level
- View pending rewards
- See next milestones

**What it shows**:
- Total score
- Achievements count
- Challenges completed
- Relationship level
- Engagement score
- Last reward time
- Pending rewards
- Next milestones

**Example**:
```bash
/sys engagement
# → Shows full engagement dashboard
```

---

### `/sys role` - Role Management

**Purpose**: Manage roles and permissions

**When to use**:
- View role hierarchy
- Assign roles
- Check permissions

**Subcommands**:
- `/sys role tree` - Show role hierarchy
- `/sys role list` - List all roles
- `/sys role show <role>` - Show role details
- `/sys role assign <role>` - Assign role

**Example**:
```bash
/sys role tree
# → Shows role hierarchy
```

---

### `/sys security` - Security Settings

**Purpose**: Security tests and audits

**When to use**:
- Run security tests
- Pre-release security check
- Security audit

**Subcommands**:
- `/sys security status` - Show security status
- `/sys security test` - Run standard tests
- `/sys security release test` - Pre-release tests
- `/sys security audit` - Comprehensive audit

**Example**:
```bash
/sys security test
# → Runs security tests
```

---

### `/sys control` - System Control

**Purpose**: Enable/disable system features

**When to use**:
- Turn agents on/off
- Enable/disable roles
- Global kill switch (with confirmation)

**Subcommands**:
- `/sys control system on|off` - Global kill switch
- `/sys control agents on|off` - Enable/disable agents
- `/sys control roles on|off` - Enable/disable roles

**Example**:
```bash
/sys control system off
# → Requires strong confirmation
# → Disables entire DoPlan system
```

---

## 🎯 Command Flow Examples

### Complete Workflow

```bash
# 1. Onboarding (first time)
/hey

# 2. Capture idea
/do

# 3. Generate plan
/plan

# 4. Start development
/dev

# 5. Complete task
/done

# 6. Check progress
/status

# 7. View engagement
/sys engagement

# 8. Continue development
/dev
/done
# ... repeat ...
```

### Fast-Track Workflow

```bash
# 1. Fast-track with detailed prompt
/do now --prompt "Build a todo app with React and Node.js"

# 2. Generate plan
/plan

# 3. Start development
/dev
/done
# ... repeat ...
```

### Lucky Mode (Inspiration)

```bash
# 1. Get AI-suggested ideas
/do i'm lucky

# 2. Choose from suggestions
# → AI suggests 2 ideas
# → You choose one or ask for more
# → AI learns from rejections

# 3. Generate plan
/plan

# 4. Start development
/dev
/done
```

---

## 💡 Command Tips

### Efficiency Tips
- Use `/do now` when you have a detailed prompt
- Use `/do feature` for single feature ideas
- Use `/do i'm lucky` for inspiration
- Check `/status` regularly
- View `/sys engagement` to stay motivated

### Best Practices
- Always use `/done` when tasks are complete
- Check `/status` before starting new work
- View `/sys engagement` to see achievements
- Use `/hey` if you forget commands
- Use `/sys` to explore system features

### Power User Tips
- Use `/sys control` to customize system
- Use `/sys role` to manage permissions
- Use `/sys security` before releases
- Check state history for rollback
- Use time tracker data for analytics

---

## 📚 Next Steps

Now that you understand the commands:

1. **[Workflow Commands](./02-Workflow-Commands.md)** - Deep dive into each command
2. **[System Commands](./03-System-Commands.md)** - Master system control
3. **[Command Reference](./04-Command-Reference.md)** - Complete reference

---

**Ready to master the commands?** → [Workflow Commands](./02-Workflow-Commands.md)

