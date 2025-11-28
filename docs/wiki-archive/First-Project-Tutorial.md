# First Project Tutorial

Complete step-by-step tutorial for creating your first project with DoPlan CLI.

---

## Overview

This tutorial will guide you through creating your first project from start to finish. By the end, you'll have:
- A complete project structure
- Planning documents (PRD, Architecture, Design System)
- Implementation tasks
- Your first feature implemented

**Estimated Time**: 30-45 minutes

---

## Prerequisites

Before starting, ensure you have:
- **DoPlan CLI installed** (see [Installation](Installation))
- **An AI-powered IDE** (Cursor, Claude Code, Antigravity, Windsurf, Cline, or OpenCode)
- **Basic familiarity** with command-line tools

---

## Step 1: Create Your Project

### 1.1 Run DoPlan CLI

Open your terminal and run:

```bash
npx @doplan-dev/cli
```

### 1.2 Interactive Wizard

The wizard will ask:

1. **Project Name**: Enter a name (e.g., `my-first-app`)
   - Use lowercase letters, numbers, and hyphens
   - No spaces or special characters

2. **IDE Choice**: Select your IDE from the menu
   - Use arrow keys to navigate
   - Press Enter to select

### 1.3 Project Generation

DoPlan CLI will generate your project in under 5 seconds. You'll see:
- ✅ Project directory created
- ✅ Agents generated
- ✅ Commands generated
- ✅ Rules library extracted
- ✅ GitHub workflows created
- ✅ Boilerplate code generated

**Output**:
```
✅ Project generated successfully!

Open with: code ./my-first-app
Then type /tell to begin
```

---

## Step 2: Open Your Project

### 2.1 Navigate to Project

```bash
cd my-first-app
```

### 2.2 Open in IDE

```bash
code .  # or your IDE's command
```

**Alternative IDE Commands**:
- Cursor: `cursor .`
- Claude Code: `claude .`
- Antigravity: `antigravity .`
- Windsurf: `windsurf .`
- Cline: `cline .`
- OpenCode: `opencode .`

---

## Step 3: Understand the Project Structure

### 3.1 Explore the Directory

Your project has this structure:

```
my-first-app/
├── .cursor/
│   ├── agents/          # 18 AI agent personas
│   ├── commands/       # Command definitions
│   └── rules/          # Rules library (1000+ rules)
├── .plan/
│   ├── 00_System/      # Planning documents
│   ├── TASKS.md        # Implementation tasks
│   └── active_state.json
├── .github/
│   └── workflows/      # CI/CD workflows
├── src/                # Your source code
├── README.md           # Project documentation
└── STANDUP.md          # Daily standup notes

Need the full breakdown of every folder? See the dedicated [Project Structure](Project-Structure) reference.
```

### 3.2 Key Files

- **`.cursor/agents/`** - AI agent definitions
- **`.cursor/commands/`** - Command definitions
- **`.cursor/rules/library/`** - Rules library
- **`.plan/00_System/`** - Planning documents (will be created)
- **`.plan/TASKS.md`** - Implementation tasks (will be created)

---

## Step 4: Capture Your Idea (`/tell`)

### 4.1 Use the `/tell` Command

In your IDE, type:

```
/tell
```

### 4.2 Enter Your Project Idea

Describe your project. For example:

```
/tell Build a simple todo app with the following features:
- Add, edit, and delete todos
- Mark todos as complete
- Filter todos (all, active, completed)
- Dark mode support
- Local storage persistence
```

### 4.3 What Happens

1. Your idea is saved to `.plan/00_System/IDEA.md`
2. Project Orchestrator and Product Manager are activated
3. Project state is updated

**Check the file**: Open `.plan/00_System/IDEA.md` to see your idea saved.

---

## Step 5: Brainstorm (`/improve`)

### 5.1 Use the `/improve` Command

Type:

```
/improve
```

### 5.2 Review Brainstorm Results

The command activates all Level 1 managers who provide ideas:
- Product Manager - Product insights
- Engineering Lead - Technical suggestions
- Design Manager - Design ideas
- QA Manager - Quality considerations
- Release Manager - Release planning
- Documentation Lead - Documentation needs

**Check the file**: Open `.plan/00_System/BRAINSTORM.md` to see all ideas.

### 5.3 What to Look For

- Feature suggestions
- Technical approaches
- Design considerations
- Potential issues
- Best practices

---

## Step 6: Generate Planning Documents (`/write`)

### 6.1 Use the `/write` Command

Type:

```
/write
```

### 6.2 What Gets Generated

Three comprehensive documents:

1. **PRD.md** (Product Requirements Document)
   - Product overview
   - User stories
   - Features and requirements
   - Success metrics

2. **ARCHITECTURE.md** (Technical Architecture)
   - System architecture
   - Technology stack
   - Component design
   - Data models

3. **DESIGN_SYSTEM.md** (Design System)
   - Design principles
   - Color palette
   - Typography
   - Component library

### 6.3 Review the Documents

Open and review each document:
- `.plan/00_System/PRD.md`
- `.plan/00_System/ARCHITECTURE.md`
- `.plan/00_System/DESIGN_SYSTEM.md`

**Take your time**: Review carefully before approving.

---

## Step 7: Edit Documents (Optional) (`/change`)

### 7.1 Make Changes if Needed

If you want to modify any document:

```
/change prd Add user authentication feature
/change architecture Use PostgreSQL instead of SQLite
/change design Add mobile-first responsive design
```

### 7.2 Review Changes

Check the updated documents to ensure changes are correct.

---

## Step 8: Approve the Plan (`/good`)

### 8.1 Use the `/good` Command

When you're satisfied with the planning documents:

```
/good
```

### 8.2 What Happens

1. All planning documents are validated
2. Plan is locked (`locked: true` in `active_state.json`)
3. Phase is updated to "approved"
4. Plan is ready for task generation

**Important**: Once approved, the plan is locked. You can still use `/change` to modify, but you'll need to approve again.

---

## Step 9: Generate Tasks (`/plan`)

### 9.1 Use the `/plan` Command

Type:

```
/plan
```

### 9.2 What Gets Generated

Implementation tasks are generated in `.plan/TASKS.md`:
- Tasks organized by phases
- Dependencies identified
- Effort estimates
- Acceptance criteria

### 9.3 Review Tasks

Open `.plan/TASKS.md` and review:
- Task organization
- Dependencies
- Effort estimates
- Completeness

---

## Step 10: Build Your First Feature (`/build`)

### 10.1 Use the `/build` Command

Type:

```
/build
```

This starts the first uncompleted task.

### 10.2 What Happens

1. Task context is loaded
2. Relevant agents are activated
3. Implementation begins
4. Active task is set in `active_state.json`

### 10.3 Work on the Task

Follow the agent's guidance to implement the task:
- Read the task requirements
- Review related code
- Implement the feature
- Test your changes

---

## Step 11: Complete Your First Task (`/done`)

### 11.1 Use the `/finished` Command

When you've completed the task:

```
/finished
```

### 11.2 What Happens

1. Task is marked complete in `TASKS.md`
2. Changes are auto-committed
3. Changes are auto-pushed to your branch
4. CHANGELOG.md is updated (if significant)
5. Progress is updated

### 11.3 Verify Completion

Check:
- Task marked complete in `TASKS.md`
- Git commit created
- Changes pushed to branch
- Progress updated

---

## Step 12: Continue Building

### 12.1 Build Next Task

```
/build
```

### 12.2 Track Progress

```
/progress
```

### 12.3 Repeat

Continue building and completing tasks:
1. `/build` - Start task
2. Implement feature
3. `/finished` - Complete task
4. Repeat

---

## Common Pitfalls and Solutions

### Pitfall 1: Skipping Planning

**Problem**: Jumping straight to coding without planning.

**Solution**: Always follow the workflow:
1. `/tell` - Capture idea
2. `/improve` - Brainstorm
3. `/write` - Generate plans
4. `/good` - Approve
5. `/plan` - Generate tasks
6. `/build` - Start coding

### Pitfall 2: Not Reviewing Documents

**Problem**: Approving plans without reviewing.

**Solution**: Always review:
- PRD.md for completeness
- ARCHITECTURE.md for soundness
- DESIGN_SYSTEM.md for clarity

### Pitfall 3: Not Using `/change`

**Problem**: Accepting generated documents as-is.

**Solution**: Use `/change` to refine:
- Add missing features
- Fix technical decisions
- Update design requirements

### Pitfall 4: Building Without Approval

**Problem**: Trying to build before approving plan.

**Solution**: Always approve first:
1. Review documents
2. Use `/change` if needed
3. Use `/good` to approve
4. Then use `/plan` and `/build`

### Pitfall 5: Not Tracking Progress

**Problem**: Losing track of what's done.

**Solution**: Use `/progress` regularly:
- Check completion status
- Identify blockers
- Plan next steps

---

## Next Steps

Now that you've completed your first project:

1. **[Workflow Guide](Workflow)** - Understand the complete workflow
2. **[Commands Reference](Commands)** - Learn all available commands
3. **[Agents Documentation](Agents)** - Understand the AI agents
4. **[Examples](Examples)** - See example projects and use cases
5. **[Development Guide](Development)** - Advanced patterns and techniques

---

## Summary

You've learned:
- ✅ How to create a project with DoPlan CLI
- ✅ How to capture and refine your idea
- ✅ How to generate planning documents
- ✅ How to approve and lock your plan
- ✅ How to generate implementation tasks
- ✅ How to build and complete features
- ✅ How to track progress

**Congratulations!** You're now ready to build amazing projects with DoPlan CLI.

---

## Related Pages

- [Quick Start](Quick-Start) - Quick 5-minute guide
- [Workflow Guide](Workflow) - Complete workflow
- [Commands Reference](Commands) - All commands
- [FAQ](FAQ) - Common questions
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

