# Phase-by-Phase Guide

A detailed breakdown of each phase in the DoPlan workflow, from onboarding to production.

---

## Phase 1: Onboarding (`/hey`)

### Purpose

Get familiar with DoPlan, learn the system, understand commands, and set up your preferences.

### Duration

- **First time**: 5-10 minutes
- **Returning**: 30 seconds (personalized greeting)

### Steps

1. **Welcome & Introduction**
   - Personalized greeting (if returning user)
   - System overview
   - What DoPlan is and does

2. **Personal Information** (First time only)
   - Name collection
   - Experience level assessment
   - Development support mode selection

3. **System Overview**
   - Agent hierarchy explanation
   - Command structure
   - Workflow overview

4. **Interactive Tutorial**
   - Command walkthrough
   - Test drive mode (optional)
   - Practice commands

5. **Reference Materials**
   - Quick reference created
   - Agent hierarchy saved
   - Command examples saved

### Output Files

- `.do/system/user_profile.json` - Your profile
- `docs/references/QUICK_REFERENCE.md` - Cheat sheet
- `docs/overview/AGENT_HIERARCHY.md` - Agent structure
- `docs/references/COMMAND_EXAMPLES.md` - Examples
- Memory Card initialized

### Engagement

- Onboarding completion tracked
- Memory Card updated
- Relationship started
- First achievements possible

---

## Phase 2: Ideation (`/do`)

### Purpose

Capture your project idea through iterative conversation, then conduct discovery and refinement.

### Duration

- **Simple idea**: 10-15 minutes
- **Complex idea**: 20-30 minutes
- **Fast-track**: 1-2 minutes

### Sub-phases

#### 2.1 Ideation Phase (Iterative Conversation)

**What happens**:
- Initial idea capture
- Iterative conversation loop
- Encouragement for more details
- Continues until "done"

**Example**:
```
You: I want to build a todo app

DoPlan: ✨ Great start! Tell me more...

You: It should have categories

DoPlan: 🙏 Thank you! That will improve results. Tell me more...

You: And sync across devices

DoPlan: 🙏 Thank you! Anything else?

You: done
```

**Output**: All input combined into IDEA.md

#### 2.2 Meeting Phase (Automatic)

**What happens**:
- Discovery meeting starts automatically
- Adaptive questions based on:
  - Project type
  - Experience level
  - Idea complexity
- Generates BRAINSTORM.md

**Questions might include**:
- Target users
- Key features
- Success metrics
- Technical requirements
- Timeline considerations

**Output**: BRAINSTORM.md with discovery results

#### 2.3 Refinement Phase (Automatic)

**What happens**:
- Enhances idea with suggestions
- Adds improvements
- Generates REFINEMENTS.md
- Updates IDEA.md

**Output**: REFINEMENTS.md with enhancements

### Output Files

- `.do/system/IDEA.md` - Your complete idea
- `.do/system/BRAINSTORM.md` - Discovery results
- `.do/system/REFINEMENTS.md` - Refinements

### Engagement

- Conversation history tracked
- Memory Card updated
- Achievements possible
- Relationship building

### Fast-Track Options

- `/do now --prompt "..."` - Skip to planning
- `/do feature` - Single feature idea
- `/do i'm lucky` - AI-suggested ideas

---

## Phase 3: Planning (`/plan`)

### Purpose

Generate structured execution plan from your idea documents.

### Duration

- **Small project**: 1-2 minutes
- **Medium project**: 2-3 minutes
- **Large project**: 3-5 minutes

### Steps

1. **Read Documents**
   - Loads IDEA.md
   - Loads BRAINSTORM.md
   - Processes content

2. **Generate TASKS.md**
   - Creates master task list
   - Organizes by phases
   - Assigns task IDs
   - Sets initial status

3. **Create Phase Structure**
   - Phase directories (01-Foundation, 02-Core, etc.)
   - Feature folders
   - Template files

4. **Sync Documentation**
   - Syncs to docs/ directory
   - Creates feature documentation
   - Updates project structure

### Output Structure

```
.do/plan/
├── TASKS.md                    # Master task list
├── 01-Foundation/
│   ├── feature-auth/
│   │   ├── design.md
│   │   ├── plan.md
│   │   ├── tasks.md
│   │   ├── prompts.md
│   │   └── github.md
│   └── feature-database/
├── 02-Core/
│   └── ...
└── 03-Enhancement/
    └── ...
```

### TASKS.md Format

```markdown
# Implementation Tasks

## Phase 1: Foundation

### 1.1 Set up project structure
**Description**: Initialize project with basic configuration
**Status**: ⏳ Pending
**Tasks**:
- [ ] Create project structure
- [ ] Set up configuration files
- [ ] Initialize Git repository

### 1.2 Configure database
**Description**: Set up database connection
**Status**: ⏳ Pending
...
```

### Engagement

- Planning completion tracked
- "Planner" achievement milestones
- Memory Card updated
- Challenges possible (first plan)

---

## Phase 4: Development (`/dev` auto-completes)

### Purpose

Build your project task by task with automatic tracking and Git automation.

### Duration

- **Per task**: Minutes to hours (varies)
- **Task overhead**: ~1 minute (start + complete)

### The Development Loop

```
/dev    → Start task
  ↓
[Code]  → Develop feature
  ↓
/dev    → Auto-completes when finished (commit/push/achievements)
  ↓
/dev    → Next task
  ↓
...repeat...
```

### Starting a Task (`/dev`)

**What happens**:
1. Finds next available task (or specific task)
2. Displays task information
3. Checks Git status
4. Creates/checks out branch (`task/[ID]`)
5. Syncs documentation
6. Starts time tracking
7. Updates active state
8. Shows personalized message

**Example output**:
```
🚀 Starting development workflow

📋 Task: 2.1 - User Authentication
   Description: Implement user login and registration
   Phase: Core Features
   Task ID: 2.1

✓ Git repository detected
🌿 Branch: task/2.1
   ✓ Branch created/checked out

📚 Syncing documentation...
   ✓ Feature documentation synced

✅ Development environment ready!
💡 Working on: User Authentication
   • Feature documentation synced
   • Branch created/checked out

📝 Next steps:
   • Review feature documentation in .do/plan/
   • Start coding with your IDE
   • Keep coding—completion runs automatically when finished
```

**Time tracking starts**: `task_started_at` stored in active_state.json

### Development

**What you do**:
- Review feature documentation
- Code your feature
- Test your code
- Use your IDE

**Time**: Varies (minutes to hours)

### Auto-Completion (handled by `/dev`)

**What happens automatically**:
1. Verifies active branch and dependencies
2. Marks task complete in TASKS.md
3. Updates state (adds to completed, clears active)
4. Creates state snapshot
5. **Auto-commits** (conventional format) and **auto-pushes**
6. Updates changelog (if significant)
7. Suggests PR creation
8. Checks achievements/challenges
9. Displays task duration and next steps

**Time tracking ends**: Duration calculated and displayed

### Engagement

- Task completion tracked
- "Code Machine" achievement milestones
- Challenges detected (first-time tasks)
- Memory Card updated
- Personalized encouragement

---

## Phase 5: Engagement (`/sys engagement`)

### Purpose

View your achievements, challenges, and engagement metrics.

### Duration

- **Instant** - No processing time

### What It Shows

1. **Total Score** - Your current points
2. **Achievements** - Count of achievements earned
3. **Challenges** - Count of challenges completed
4. **Relationship Level** - 0-100 relationship strength
5. **Engagement Score** - 0-100% engagement
6. **Last Reward** - Time since last achievement/challenge
7. **Pending Rewards** - Rewards scheduled for release
8. **Next Milestones** - Hints for upcoming achievements

### Example Output

```
📊 DoPlan Engagement Dashboard
============================================================

  💰 Total Score: 1,250 points
  🏆 Achievements: 8
  🎯 Challenges: 2

  🤝 Relationship Level: 45/100 💪 Developing
  📈 Engagement: 80% 👍 Good!

  ⏰ Last Reward: 2 hours ago
  ⏳ Pending Rewards: 1 (coming soon!)

  🎯 Next Milestones:
     🎯 'On the Rise': Reach 250 points
     📋 'Planner': Use /plan command 25 times
     🚀 'Do It Again': Use /do command 50 times
```

### When to Use

- Check your progress
- See achievements earned
- View relationship level
- Get motivated
- Check next milestones

---

## 🔄 Phase Transitions

### Onboarding → Ideation

**Trigger**: Run `/do` after `/hey`
**State change**: "onboarding" → "ideation"
**What happens**: Memory Card updated, relationship building begins

### Ideation → Planning

**Trigger**: Run `/plan` after `/do`
**State change**: "ideation" → "planning"
**What happens**: Plan generated, phase structure created

### Planning → Development

**Trigger**: Run `/dev` after `/plan`
**State change**: "planning" → "building"
**What happens**: Task started, branch created, time tracking begins

### Development → Completion

**Trigger**: `/dev` detects task is finished
**State change**: Task completed, state updated
**What happens**: Task marked complete, committed, pushed, achievements checked

### Development Loop

**Trigger**: Run `/dev` again to start the next task
**State change**: Next task started
**What happens**: New task, new branch, time tracking continues

---

## 💡 Phase Tips

### Onboarding
- Complete tutorial first time
- Review reference materials
- Set preferences accurately
- Ask questions if needed

### Ideation
- Use iterative conversation
- Provide complete details
- Don't rush the process
- Use fast-track if you have detailed prompt

### Planning
- Review plan before development
- Understand phase structure
- Check task dependencies
- Verify documentation

### Development
- Review feature docs before coding
- Let `/dev` auto-complete when finished
- View engagement for motivation

### Engagement
- View dashboard regularly
- Track achievements
- Monitor relationship level
- Stay motivated

---

**Next**: [Best Practices](./03-Best-Practices.md)

