# Quick Tour - See DoPlan's Power

Take a 5-minute tour of DoPlan's most powerful features. See how DoPlan transforms development into an engaging, educational, and fun experience.

---

## 🎮 Feature 1: Engagement System

DoPlan makes development **fun** with achievements, challenges, and rewards.

### Try It

```bash
# Check your engagement
/sys engagement
```

**What you'll see**:
- Your total score (points)
- Achievements earned
- Challenges completed
- Relationship level
- Next milestones

### Earn Achievements

Achievements are earned automatically as you use DoPlan:

- **First Steps** - Complete your first task (50 points)
- **On the Rise** - Reach 250 points (25 points)
- **Planner** - Use /plan 25 times (50 points)
- **Code Machine** - Use /dev 100 times (200 points) 💻

**200+ achievements** waiting to be earned!

### Complete Challenges

Challenges reward high scores for first-time tasks:

- **API Integration Master** - First API integration (500 points) 🔌
- **First Launch** - First deployment (1000 points) 🚀
- **Public Launch** - First public release (2000 points) 🌐

**30+ challenges** with 300-2000 points each!

---

## 🧠 Feature 2: Memory Card & Personalization

DoPlan learns about you and adapts to your style.

### Your Memory Card

Stored at `~/.doplan/memory_card.json`, it remembers:
- Your preferences and personality
- Learning goals and interests
- Achievements and challenges
- Relationship metrics
- Usage patterns

### See Your Memory Card

```bash
# View your memory card (if you have access)
cat ~/.doplan/memory_card.json
```

**What's stored**:
```json
{
  "user_name": "Your Name",
  "experience_level": "intermediate",
  "work_style": "fast",
  "score": 1250,
  "achievements": [...],
  "relationship_level": 45,
  "learning_goals": ["React", "TypeScript"],
  ...
}
```

### Personalization in Action

**First time**:
```
DoPlan: Hello! 👋 I'm DoPlan. Let's get started!
```

**After 10 projects**:
```
DoPlan: Hey [Your Name]! 👋 Great to see you again!
        I remember you're working on React projects.
        Ready to build something amazing?
```

---

## ⏱️ Feature 3: Automatic Time Tracking

DoPlan tracks **everything** automatically - no manual logging needed.

### How It Works

Time is tracked for:
- Every command execution
- Each phase (ideation, planning, development)
- Complete task duration (from /dev to /done)
- Session duration

### View Time Data

Time data is stored in `.do/system/time-tracker.jsonl`:

```bash
# View your time tracking (if available)
cat .do/system/time-tracker.jsonl | tail -5
```

**Format** (JSONL - one entry per line):
```json
{
  "project": "/path/to/project",
  "command": "/dev",
  "phase": "development",
  "started_at": "2025-01-15T10:00:00Z",
  "ended_at": "2025-01-15T10:15:30Z",
  "duration_ms": 930000,
  "status": "completed"
}
```

### See Duration After Tasks

When you complete a task with `/done`:
```
✅ Task 2.1 marked complete!
   ⏱️  Task duration: 2h 15m
```

**Perfect for**:
- Time analysis
- Productivity insights
- Billing (if needed)
- Future dashboards

---

## 🔄 Feature 4: State Management

DoPlan maintains complete state history with rollback capability.

### Active State

Current state in `.do/system/history/active_state.json`:
```json
{
  "phase": "building",
  "active_task": "2.1",
  "active_branch": "task/2.1",
  "completed": ["1.1", "1.2"],
  "task_started_at": "2025-01-15T10:00:00Z"
}
```

### State Snapshots

Automatic snapshots in `.do/system/history/`:
- `state-20250115T100000Z-build.json` - Before /dev
- `state-20250115T101530Z-done.json` - After /done

**Benefits**:
- Full history
- Diff capability
- Rollback support
- Audit trail

---

## 🚀 Feature 5: Git Automation

DoPlan automates Git operations so you can focus on coding.

### Auto-Commit

When you run `/done`:
```
✓ Changes committed
  Commit: feat(task-2.1): complete user authentication
```

**Conventional commit format** automatically:
- `feat(task-X.X):` for features
- `fix(task-X.X):` for bug fixes
- `docs(task-X.X):` for documentation

### Auto-Push

```
✓ Changes pushed to task/2.1
```

**No manual commands needed!**

### PR Suggestions

```
💡 Suggestion: Create a pull request?
   Run: gh pr create --title "feat: User authentication" --body "Completes task 2.1"
```

---

## 🎓 Feature 6: Learning Support

DoPlan is built for learning, not just development.

### Learning Goals

Set learning goals and DoPlan helps you achieve them:

```bash
# When you use a new technology
DoPlan: I see you're using TypeScript for the first time!
        This is great for your learning goal "Learn TypeScript"!
        
        [Provides explanation and resources]
        
        🎯 Achievement: Tech Explorer (50 points)
```

### Educational Explanations

DoPlan explains **why**, not just **how**:

```
DoPlan: I'm creating a feature branch because:
        1. It isolates your changes
        2. Makes code review easier
        3. Allows parallel development
        4. Follows Git best practices
```

### Pain Point Assistance

DoPlan remembers where you struggle and provides extra help:

```
DoPlan: I noticed you've had trouble with database migrations before.
        Let me provide extra guidance for this task.
        
        [Detailed migration help]
```

---

## 🎯 Feature 7: Iterative Idea Capture

DoPlan uses **iterative conversation** to capture complete ideas.

### How It Works

Instead of a single input, DoPlan encourages multiple rounds:

```
You: I want to build a todo app

DoPlan: ✨ Great start! Tell me more...

You: It should have categories

DoPlan: 🙏 Thank you! That will improve results. Tell me more...

You: And sync across devices

DoPlan: 🙏 Thank you! Anything else?

You: done
```

**Result**: Complete idea captured with all details!

---

## 🧩 Feature 8: Subcommands & Flexibility

DoPlan commands have powerful subcommands for different needs.

### /do Subcommands

```bash
/do                    # Full ideation workflow
/do feature            # Add a single feature idea
/do now                # Fast-track with detailed prompt
/do i'm lucky          # Get AI-suggested ideas
```

### /sys Subcommands

```bash
/sys                   # Show system tree
/sys engagement        # Engagement dashboard
/sys role              # Role management
/sys security          # Security settings
/sys control           # System control (kill switch)
```

---

## 📊 Feature 9: Progress Tracking

Always know where you are and what's next.

### /status Command

```bash
/status
```

**Shows**:
- Current phase
- Tasks completed/total
- Percentage complete
- Current task
- Next task
- State deltas

### Visual Progress

```
📊 Project Progress

Phase: Foundation
Tasks: 3/4 completed (75%)

Current task: 1.4 - Set up testing
Next up: 2.1 - User authentication

State Delta:
  • Task 1.3 completed
  • Phase: Foundation (no change)
```

---

## 🎨 Feature 10: Project Structure

DoPlan creates a **complete, production-ready** structure.

### Generated Structure

```
your-project/
├── .do/
│   ├── system/          # System files
│   │   ├── IDEA.md
│   │   ├── BRAINSTORM.md
│   │   ├── PRD.md
│   │   ├── ARCHITECTURE.md
│   │   └── history/     # State snapshots
│   │
│   ├── plan/            # Execution plan
│   │   ├── TASKS.md
│   │   ├── 01-Foundation/
│   │   ├── 02-Core/
│   │   └── 03-Enhancement/
│   │
│   └── core/            # Templates
│
└── docs/                # Documentation
    ├── foundation/
    ├── features/
    └── reference/
```

**Everything organized and ready!**

---

## 🎉 What Makes DoPlan Special

After this tour, you've seen:

1. ✅ **Engagement System** - Makes development fun
2. ✅ **Personalization** - Adapts to you
3. ✅ **Time Tracking** - Automatic analytics
4. ✅ **State Management** - Full history
5. ✅ **Git Automation** - No manual work
6. ✅ **Learning Support** - Educational focus
7. ✅ **Iterative Capture** - Complete ideas
8. ✅ **Flexible Commands** - Power when needed
9. ✅ **Progress Tracking** - Always know status
10. ✅ **Complete Structure** - Production-ready

---

## 🚀 Next Steps

Now that you've seen DoPlan's power:

1. **[Learn the Commands](../02-Commands/)** - Master every command
2. **[Understand Engagement](../03-Engagement-System/)** - Deep dive into achievements
3. **[Explore Memory Card](../04-Memory-Card-Brain/)** - Understand personalization
4. **[Master the Workflow](../05-Workflow/)** - Optimize your process

---

## 💡 Pro Tips

- **Check `/sys engagement`** regularly to see your progress
- **Use `/status`** to track where you are
- **Complete challenges** for high scores
- **Set learning goals** to track your growth
- **Explore subcommands** for power features

---

**Ready to master DoPlan?** → [Commands Guide](../02-Commands/)

