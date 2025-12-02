# Complete Workflow

The DoPlan workflow takes you from idea to production in a simple, engaging, and automated way. This guide covers the complete end-to-end process.

---

## 🎯 The Complete Journey

```
1. Onboarding (/hey)
   ↓
2. Ideation (/do)
   ↓
3. Planning (/plan)
   ↓
4. Development Loop (/dev → /done)
   ↓
5. Progress Tracking (/status)
   ↓
6. Engagement (/sys engagement)
```

---

## Phase 1: Onboarding (`/hey`)

### Purpose

Get familiar with DoPlan, learn the system, and set up your preferences.

### Steps

1. **Run `/hey`**
   ```bash
   /hey
   ```

2. **Interactive Tutorial**
   - Welcome and introduction
   - System overview
   - Agent hierarchy explanation
   - Command walkthrough
   - Test drive mode (optional)

3. **Personalization**
   - Name collection
   - Experience level
   - Development support mode
   - Preferences

4. **Reference Materials**
   - Quick reference created
   - Agent hierarchy saved
   - Command examples saved

### Output

- `.do/system/user_profile.json` - Your profile
- `docs/references/QUICK_REFERENCE.md` - Cheat sheet
- Memory Card initialized
- Relationship started

### Time: 5-10 minutes (first time)

---

## Phase 2: Ideation (`/do`)

### Purpose

Capture your project idea through iterative conversation, then conduct discovery and refinement.

### Steps

1. **Run `/do`**
   ```bash
   /do
   ```

2. **Ideation Phase** (Iterative Conversation)
   ```
   You: I want to build a todo app
   
   DoPlan: ✨ Great start! Tell me more...
   
   You: It should have categories
   
   DoPlan: 🙏 Thank you! That will improve results. Tell me more...
   
   You: And sync across devices
   
   DoPlan: 🙏 Thank you! Anything else?
   
   You: done
   ```

3. **Meeting Phase** (Automatic)
   - Discovery meeting starts automatically
   - Adaptive questions based on project type
   - Experience level consideration
   - Generates BRAINSTORM.md

4. **Refinement Phase** (Automatic)
   - Enhances idea with suggestions
   - Generates REFINEMENTS.md
   - Updates IDEA.md

### Output

- `.do/system/IDEA.md` - Your complete idea
- `.do/system/BRAINSTORM.md` - Discovery results
- `.do/system/REFINEMENTS.md` - Refinements
- Memory Card updated

### Time: 15-30 minutes (depending on complexity)

---

## Phase 3: Planning (`/plan`)

### Purpose

Generate structured execution plan from your idea documents.

### Steps

1. **Run `/plan`**
   ```bash
   /plan
   ```

2. **Plan Generation**
   - Reads IDEA.md and BRAINSTORM.md
   - Generates TASKS.md
   - Creates phase structure
   - Generates feature folders

3. **Documentation Sync**
   - Syncs to docs/ directory
   - Creates feature documentation
   - Updates project structure

### Output

- `.do/plan/TASKS.md` - Master task list
- `.do/plan/01-Foundation/` - Phase 1
- `.do/plan/02-Core/` - Phase 2
- `.do/plan/03-Enhancement/` - Phase 3
- Feature folders with templates

### Time: 1-2 minutes

---

## Phase 4: Development Loop (`/dev` → `/done`)

### Purpose

Build your project task by task with automatic tracking and Git automation.

### The Loop

```
/dev    → Start task
  ↓
[Code]  → Develop feature
  ↓
/done   → Complete task (auto-commit/push)
  ↓
/dev    → Next task
  ↓
[Code]  → Develop feature
  ↓
/done   → Complete task
  ↓
...repeat...
```

### Starting a Task (`/dev`)

1. **Run `/dev`**
   ```bash
   /dev
   # Or specific task
   /dev --feature "auth"
   ```

2. **What Happens**
   - Finds next available task
   - Displays task information
   - Creates/checks out Git branch
   - Syncs documentation
   - Starts time tracking
   - Updates active state
   - Shows personalized message

3. **Development**
   - Review feature documentation
   - Code your feature
   - Test your code
   - Use your IDE

### Completing a Task (`/done`)

1. **Run `/done`**
   ```bash
   /done
   ```

2. **What Happens**
   - Verifies active branch
   - Checks dependencies
   - Marks task complete in TASKS.md
   - Updates state
   - Creates state snapshot
   - **Auto-commits** (conventional format)
   - **Auto-pushes** to remote
   - Checks achievements/challenges
   - Displays task duration

3. **Output**
   ```
   ✅ Task 2.1 marked complete!
      ⏱️  Task duration: 2h 15m
      ✓ Changes committed
      ✓ Changes pushed to task/2.1
   
   💡 Next steps:
      • Type /dev to start the next task
      • Type /status to see overall progress
   ```

### Time Per Task: Varies (minutes to hours)

---

## Phase 5: Progress Tracking (`/status`)

### Purpose

Monitor your progress and see where you are in the project.

### Steps

1. **Run `/status`**
   ```bash
   /status
   ```

2. **What You See**
   - Current phase
   - Tasks completed/total
   - Percentage complete
   - Current task (if any)
   - Next task
   - State deltas

### Output

```
📊 Project Progress

Phase: Foundation
Tasks: 3/4 completed (75%)

Current task: 1.4 - Set up testing
Next up: 2.1 - User authentication

State Delta (since last snapshot):
  • Task 1.3 completed
  • Phase: Foundation (no change)
```

### Time: Instant

---

## Phase 6: Engagement (`/sys engagement`)

### Purpose

View your achievements, challenges, and engagement metrics.

### Steps

1. **Run `/sys engagement`**
   ```bash
   /sys engagement
   ```

2. **What You See**
   - Total score
   - Achievements earned
   - Challenges completed
   - Relationship level
   - Engagement score
   - Next milestones

### Time: Instant

---

## 🔄 Complete Example

### Day 1: Setup

```bash
# 1. Onboarding
/hey
# → Tutorial, learn system

# 2. Capture idea
/do
# → Iterative conversation
# → Discovery meeting
# → Refinement

# 3. Generate plan
/plan
# → TASKS.md created
# → Phase structure created
```

### Day 2-7: Development

```bash
# Development loop
/dev
# → Task 1.1 started
# → Code feature
/done
# → Task 1.1 complete (2h 15m)
# → Auto-committed and pushed

/dev
# → Task 1.2 started
# → Code feature
/done
# → Task 1.2 complete (1h 30m)

# ... continue ...
```

### Regular Checks

```bash
# Check progress
/status
# → See completion percentage

# View engagement
/sys engagement
# → See achievements and score
```

---

## ⏱️ Time Breakdown

### First Time Setup
- **Onboarding**: 5-10 minutes
- **Ideation**: 15-30 minutes
- **Planning**: 1-2 minutes
- **Total**: ~30-45 minutes

### Development Per Task
- **Task start**: 10-30 seconds
- **Development**: Varies (minutes to hours)
- **Task completion**: 10-30 seconds
- **Total overhead**: ~1 minute per task

### Regular Checks
- **Status**: Instant
- **Engagement**: Instant

---

## 🎯 Workflow Benefits

### Automation
- ✅ Auto-commit (conventional format)
- ✅ Auto-push to remote
- ✅ State snapshots
- ✅ Time tracking
- ✅ Documentation sync

### Engagement
- ✅ Achievement checking
- ✅ Challenge detection
- ✅ Score tracking
- ✅ Relationship building
- ✅ Personalized messages

### Safety
- ✅ State history
- ✅ Rollback capability
- ✅ Dependency checking
- ✅ Branch verification

---

## 💡 Workflow Tips

### Efficiency
- Use `/do now` for fast-tracking
- Use `/do feature` for single features
- Check `/status` regularly
- Always use `/done` when complete

### Best Practices
- Complete `/hey` tutorial first time
- Use iterative conversation in `/do`
- Review plan before `/dev`
- Check progress regularly
- View engagement for motivation

### Power User
- Use `/do i'm lucky` for inspiration
- Track time with automatic tracking
- Use state history for rollback
- Leverage engagement system
- Complete challenges for high scores

---

## 🚀 Next Steps

1. **[Phase-by-Phase Guide](./02-Phase-by-Phase.md)** - Deep dive into each phase
2. **[Best Practices](./03-Best-Practices.md)** - Optimize your workflow
3. **[Common Patterns](./04-Common-Patterns.md)** - Real-world examples

---

**Ready to optimize?** → [Phase-by-Phase Guide](./02-Phase-by-Phase.md)

