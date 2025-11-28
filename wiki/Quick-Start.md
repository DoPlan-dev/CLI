# Quick Start Guide

Get up and running with DoPlan CLI in 5 minutes! This guide will walk you through creating your first project.

---

## Prerequisites

Before you begin, ensure you have:
- **Node.js** >= 14.0.0 installed (check with `node --version`)
- An AI-powered IDE installed (Cursor, Claude Code, Antigravity, Windsurf, Cline, or OpenCode)

If you haven't installed DoPlan CLI yet, see the [Installation Guide](Installation).

---

## Your First Project (5-Minute Tutorial)

### Step 1: Run DoPlan CLI

Open your terminal and run:

```bash
npx @doplan-dev/cli
```

Or if you have DoPlan CLI installed:

```bash
doplan
```

### Step 2: Interactive Wizard

DoPlan CLI will launch an interactive wizard that asks:

1. **Project Name**: Enter a name for your project (e.g., `my-awesome-app`)
   - Use lowercase letters, numbers, and hyphens
   - No spaces or special characters

2. **IDE Choice**: Select your preferred IDE from the list:
   - Cursor
   - Claude Code
   - Antigravity
   - Windsurf
   - Cline
   - OpenCode

### Step 3: Project Generation

DoPlan CLI will generate your complete project structure in under 5 seconds. You'll see progress indicators as it creates:

- Project directory structure
- 18 AI agent definitions
- Command definitions
- Rules library (1000+ rules)
- GitHub workflows
- Configuration files
- Boilerplate code

### Step 4: Open Your Project

Once generation is complete, DoPlan CLI will display:

```
✅ Project generated successfully!

Open with: code ./my-awesome-app
Then type /tell to begin
```

Open your project in your IDE:

```bash
cd my-awesome-app
code .  # or your IDE's command
```

---

## What Gets Generated

Your new project includes:

### 📁 Project Structure

```
my-awesome-app/
├── .cursor/
│   ├── agents/              # 18 AI agent personas
│   ├── commands/            # Command definitions
│   └── rules/               # Rules library
│       └── library/         # 15 categories, 1000+ rules
├── .plan/
│   ├── 00_System/           # Planning documents
│   ├── TASKS.md             # Implementation tasks
│   └── active_state.json    # Project state
├── .github/
│   └── workflows/           # CI/CD workflows
├── src/                     # Your source code
├── README.md                # Project documentation
└── STANDUP.md               # Daily standup notes
```

Want to see how each directory evolves during the workflow? Jump to the full [Project Structure](Project-Structure) reference.
### 🤖 AI Agents

18 specialized agents ready to help:
- Project Orchestrator
- Product Manager
- Engineering Lead
- System Architect
- Frontend Lead
- Backend Lead
- DevOps Engineer
- Security Lead
- Performance Engineer
- Design & UX Manager
- UI/UX Designer
- QA & Reliability Manager
- QA Engineer
- Release & Growth Manager
- Release Captain
- Growth Coach
- Documentation Lead
- Documentation Writer

### 📝 Commands

Core commands available:
- `/tell` - Capture your project idea
- `/improve` - Brainstorm with the team
- `/write` - Generate planning documents
- `/build` - Start coding
- And more! See [Commands Reference](Commands)

### 📚 Rules Library

1000+ embedded rules covering:
- Core workflow
- Languages (Go, TypeScript, Python, etc.)
- Frameworks (Next.js, React, Express, etc.)
- Databases
- Testing
- CI/CD
- Security
- And more!

---

## Next Steps

### 1. Capture Your Idea

In your IDE, type:

```
/tell
```

Then describe your project idea. For example:

```
/tell Build a todo app with authentication, dark mode, and offline support
```

This saves your idea to `.plan/00_System/IDEA.md` and activates the Project Orchestrator.

### 2. Brainstorm

Get ideas from your AI team:

```
/improve
```

This activates all Level 1 managers (Product Manager, Engineering Lead, Design Manager, etc.) to brainstorm improvements and ideas.

### 3. Generate Planning Documents

Create comprehensive planning documents:

```
/write
```

This generates:
- **PRD.md** - Product Requirements Document
- **ARCHITECTURE.md** - Technical Architecture
- **DESIGN_SYSTEM.md** - Design System

### 4. Review and Approve

Review the generated documents in `.plan/00_System/`. When ready, approve:

```
/good
```

This locks the plan and prepares it for task generation.

### 5. Generate Tasks

Create implementation tasks from your approved plan:

```
/plan
```

This generates `TASKS.md` with organized, actionable tasks.

### 6. Start Building

Begin coding your first task:

```
/build
```

This starts the next uncompleted task. Or build a specific task:

```
/build 1.2
```

### 7. Track Progress

See how you're doing:

```
/progress
```

### 8. Complete Tasks

When you finish a task:

```
/done
```

This marks the task complete, auto-commits your changes, and updates progress.

---

## Common First-Time Questions

### Q: Do I need to install anything else?

**A**: No! DoPlan CLI generates everything you need. Just open your project in your IDE and start using the commands.

### Q: How do I know which command to use?

**A**: Start with `/tell` to capture your idea, then follow the workflow:
1. `/tell` → Capture idea
2. `/improve` → Brainstorm
3. `/write` → Generate plans
4. `/good` → Approve
5. `/plan` → Generate tasks
6. `/build` → Start coding

See the [Commands Reference](Commands) for all available commands.

### Q: Can I modify the generated files?

**A**: Absolutely! All files are yours to modify. The AI agents will work with your changes.

### Q: What if I make a mistake?

**A**: You can always:
- Edit documents directly
- Use `/change` to modify documents
- Regenerate tasks with `/plan`
- Start over by deleting and regenerating

### Q: How do the AI agents work?

**A**: Each agent has a specific role and expertise. When you use commands, the relevant agents are activated. See the [Agents Documentation](Agents) for details.

### Q: Can I customize the agents or commands?

**A**: Yes! All agent definitions are in `.cursor/agents/` and commands are in `.cursor/commands/`. You can modify them to fit your needs. See the [Configuration Guide](Configuration).

### Q: What if I need help?

**A**: Check out:
- [FAQ](FAQ) - Common questions
- [Troubleshooting](Troubleshooting) - Common issues
- [Commands Reference](Commands) - All commands
- [Workflow Guide](Workflow) - Complete workflow

---

## Example: Creating a Todo App

Here's a complete example workflow:

```bash
# 1. Generate project
npx @doplan-dev/cli
# Enter: todo-app
# Select: Cursor

# 2. Open in IDE
cd todo-app
code .

# 3. In IDE, capture idea
/tell Build a modern todo app with authentication, dark mode, and offline support

# 4. Brainstorm
/improve

# 5. Generate plans
/write

# 6. Review and approve
/good

# 7. Generate tasks
/plan

# 8. Start building
/build

# 9. Track progress
/progress

# 10. Complete tasks
/done
```

---

## Tips for Success

1. **Start Simple**: Begin with a clear, focused idea
2. **Review Plans**: Always review generated documents before approving
3. **Use Commands**: Leverage the command system - it's designed to help
4. **Check Progress**: Use `/progress` regularly to stay on track
5. **Ask Questions**: Use `/improve` to brainstorm when stuck
6. **Customize**: Modify agents and rules to fit your needs

---

## What's Next?

Now that you've created your first project:

1. **[First Project Tutorial](First-Project-Tutorial)** - Detailed step-by-step guide
2. **[Commands Reference](Commands)** - Learn all available commands
3. **[Workflow Guide](Workflow)** - Understand the complete workflow
4. **[Agents Documentation](Agents)** - Learn about the AI agents
5. **[Workflow Guide](Workflow)** - Complete workflow and best practices

---

## Related Pages

- [Installation](Installation) - Install DoPlan CLI
- [Commands Reference](Commands) - All available commands
- [Workflow Guide](Workflow) - Complete development workflow
- [FAQ](FAQ) - Frequently asked questions
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

