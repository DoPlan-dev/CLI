# Your First Project

Create your first project with DoPlan in just 5 minutes! This guide will walk you through the entire process step by step.

---

## 🎯 What You'll Create

In this tutorial, you'll create a simple project and see DoPlan's key features:
- Project structure generation
- Idea capture
- Planning
- Development workflow
- Task completion

---

## 🚀 Step 1: Start DoPlan

Open your terminal and run:

```bash
# If installed globally
doplan

# Or with npx
npx @doplan-dev/cli
```

You'll see the interactive wizard!

---

## 📝 Step 2: Project Setup

The wizard will ask you:

1. **Project Name**: Choose a name (e.g., "my-first-app")
2. **Project Type**: Select your type (Web App, API, CLI, etc.)
3. **Tech Stack**: Choose your technologies
4. **IDE**: Select your preferred IDE (Cursor, VS Code, etc.)

**Example**:
```
Project Name: my-first-app
Project Type: Web Application
Tech Stack: React, Node.js, PostgreSQL
IDE: Cursor
```

Press Enter after each selection. The wizard creates your project structure automatically!

---

## 💡 Step 3: Capture Your Idea

After project creation, you'll be in your project directory. Start with the `/hey` command for onboarding:

```bash
/hey
```

This will:
- Welcome you to DoPlan
- Show you the system overview
- Explain the agent hierarchy
- Walk you through commands
- Create reference materials

**First time?** You'll get a full tutorial. **Returning user?** You'll get a personalized greeting!

---

## 🎨 Step 4: Capture Your Idea

Now let's capture your project idea:

```bash
/do
```

You'll be prompted to describe your idea. **Here's the magic**: DoPlan uses **iterative conversation**!

**Example conversation**:
```
You: I want to build a todo app

DoPlan: ✨ Great start! You can tell me more about your idea.
        Often, users remember more details as they talk about it.
        
        Tell me more (or type 'done' when finished):

You: It should have categories and due dates

DoPlan: 🙏 Thank you for those details! It will definitely improve the results.
        
        You can tell me more, or if you have notes, additional features,
        or any other thoughts, share them now (or type 'done' to finish):

You: And it should sync across devices

DoPlan: 🙏 Thank you for those details!
        
        Tell me more (or type 'done' to finish):

You: done
```

**What happened?**
- All your input was captured
- Saved to `.do/system/IDEA.md`
- DoPlan learned about your preferences
- Your memory card was updated

---

## 🎯 Step 5: Discovery Meeting

After capturing your idea, DoPlan automatically starts a discovery meeting:

```
DoPlan: Let's have a discovery meeting to understand your project better!
        This will help us create a better plan.
        
        [Interactive Q&A session]
```

The meeting:
- Asks relevant questions
- Adapts to your experience level
- Generates `BRAINSTORM.md`
- Creates `REFINEMENTS.md`

---

## 📋 Step 6: Generate Your Plan

Now create your execution plan:

```bash
/plan
```

This will:
- Read your IDEA.md and BRAINSTORM.md
- Generate TASKS.md with organized phases
- Create phase directories (01-Foundation, 02-Core, etc.)
- Generate feature folders with templates
- Sync documentation

**You'll see**:
```
🚀 Generating execution plan...

📋 Reading project documents...
✅ IDEA.md loaded
✅ BRAINSTORM.md loaded

📝 Generating TASKS.md...
✅ Created 12 tasks across 3 phases

📁 Creating phase structure...
✅ Phase 01-Foundation created
✅ Phase 02-Core created
✅ Phase 03-Enhancement created

📚 Syncing documentation...
✅ Documentation synced to docs/

✅ Plan generated successfully!
```

**Check your plan**:
```bash
cat .do/plan/TASKS.md
```

---

## 💻 Step 7: Start Development

Now let's start building! DoPlan will find the next available task:

```bash
/dev
```

**What happens**:
1. Finds next uncompleted task from TASKS.md
2. Creates/checks out Git branch: `task/1.1`
3. Syncs documentation for that feature
4. Starts time tracking
5. Updates active state
6. Shows personalized message

**You'll see**:
```
🚀 Starting development workflow

📋 Task: 1.1 - Set up project structure
   Description: Initialize project with basic configuration
   Phase: Foundation
   Task ID: 1.1

🌿 Branch: task/1.1
   ✓ Branch created/checked out

📚 Syncing documentation...
   ✓ Feature documentation synced

✅ Development environment ready!
💡 Working on: Set up project structure
   • Feature documentation synced
   • Branch created/checked out

📝 Next steps:
   • Review feature documentation in .do/plan/
   • Start coding with your IDE
   • Keep coding—completion runs automatically when finished
```

**Now you can code!** Open your IDE and start building.

---

## ✅ Step 8: Completion Happens Automatically

When the AI agent decides the task is done, `/dev` auto-completes for you:

- Marks the task complete and updates TASKS.md/state
- Snapshots state history
- Creates a conventional commit and pushes
- Checks achievements/challenges and shows duration
- Suggests next steps (e.g., PR creation)

**Next steps after completion**:
- Run `/dev` again to start the next task
- Use `/sys engagement` to review achievements

---

## 🎮 Step 9: View Your Engagement

See your achievements and progress:

```bash
/sys engagement
```

**You'll see**:
```
📊 DoPlan Engagement Dashboard
============================================================

  💰 Total Score: 50 points
  🏆 Achievements: 1
  🎯 Challenges: 0

  🤝 Relationship Level: 5/100 💪 Developing
  📈 Engagement: 80% 👍 Good!

  🎯 Next Milestones:
     🎯 'On the Rise': Reach 250 points
     📋 'Planner': Use /plan command 25 times
     🚀 'Do It Again': Use /do command 50 times
```

---

## 🔄 Continue the Loop

Now you can continue developing:

```bash
/dev    # Start next task
# ... code ...
/dev    # Auto-completes when finished
# ... repeat as you pick up new tasks ...
```

Each cycle:
- ✅ Tracks time automatically
- ✅ Updates progress
- ✅ Checks achievements
- ✅ Commits and pushes
- ✅ Shows duration

---

## 🎓 What You Learned

In this tutorial, you:
- ✅ Created a project with DoPlan
- ✅ Captured your idea iteratively
- ✅ Generated an execution plan
- ✅ Started development workflow
- ✅ Completed a task automatically (commit/push included)
- ✅ Viewed engagement dashboard

---

## 🚀 Next Steps

Now that you've created your first project:

1. **[Take a Quick Tour](./04-Quick-Tour.md)** - Explore more features
2. **[Learn the Commands](../02-Commands/)** - Master all commands
3. **[Understand the Workflow](../05-Workflow/)** - Deep dive into the process
4. **[Explore Engagement System](../03-Engagement-System/)** - Learn about achievements

---

## 💡 Tips

- **Use `/hey`** anytime to review commands and get help
- **View `/sys engagement`** to see your achievements
- **Use `/dev`** to start tasks and let it auto-complete them
- **Explore achievements** - they make development fun!

---

**Congratulations!** You've created your first project with DoPlan! 🎉

**Ready to learn more?** → [Quick Tour](./04-Quick-Tour.md)

