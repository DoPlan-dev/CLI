# Workflow Commands

Detailed documentation for each workflow command: `/hey`, `/do`, `/plan`, and `/dev`.

---

## `/hey` - Welcome & Onboarding

### Overview

The `/hey` command provides an interactive onboarding experience for new users and a personalized welcome for returning users.

### When to Use

- **First time**: Complete tutorial and learn DoPlan
- **Returning**: Get personalized greeting and quick reference
- **Need help**: Review commands and system overview
- **Want tutorial**: Re-run the interactive tutorial

### What It Does

1. **Checks if first time** - Reads `.do/system/user_profile.json`
2. **Personalized greeting** - Based on memory card relationship
3. **System overview** - Explains what DoPlan is
4. **Agent hierarchy** - Shows the 18-agent system
5. **Command walkthrough** - Explains each command
6. **Test drive mode** - Practice commands without saving
7. **Reference materials** - Creates quick reference files
8. **Personalized tips** - Based on experience level

### First-Time Experience

```
/hey
```

**What happens**:
1. Warm greeting: "Hello! 👋 I'm DoPlan, your AI development partner..."
2. Personal introduction questions:
   - "What should I call you?"
   - "What's your experience with development?"
3. System overview explanation
4. Agent hierarchy display (text-based tree)
5. Development support mode selection (Guided/Independent)
6. Interactive tutorial walkthrough
7. Test drive option (practice mode)
8. Personalized tips based on experience
9. Saves reference materials

### Returning User Experience

```
/hey
```

**What happens**:
- Personalized greeting based on relationship level
- Quick command reference
- Option to re-run tutorial
- Current progress summary

**Example greetings**:
- **New relationship**: "Hello! 👋 I'm DoPlan. Let's get started!"
- **Growing relationship**: "Hey [Name]! 👋 Good to have you back!"
- **Strong relationship**: "Hey [Name]! 👋 It's great to see you again. Ready to build something amazing?"

### Files Created

- `.do/system/user_profile.json` - User profile
- `.do/system/QUICK_REFERENCE.md` - Personal cheat sheet
- `docs/references/QUICK_REFERENCE.md` - Project documentation
- `docs/overview/AGENT_HIERARCHY.md` - Agent structure
- `docs/references/COMMAND_EXAMPLES.md` - Example outputs
- `docs/tutorials/TUTORIAL_NOTES.md` - Tutorial summary

### Engagement Integration

- Tracks onboarding completion
- Updates memory card with preferences
- Awards achievements for tutorial completion
- Builds relationship level
- Records memorable moments

### Performance Optimization

**Fast Path for New Projects**

The `/hey` command is optimized for instant response on new projects:

- **New Projects**: Skips heavy engagement system initialization for immediate response
  - No memory card exists → Fast path activated
  - No achievement/challenge checks needed
  - No reward processing required
  - Lightweight tracking only

- **Existing Projects**: Full engagement system with personalized experience
  - Memory card detected → Full initialization
  - Achievement and challenge checks
  - Personalized greetings and rewards
  - Complete engagement tracking

**Why This Matters**

For new projects, everything is already saved in files, so `/hey` doesn't need to:
- Load memory card multiple times (Brain, Achievements, Challenges, DopamineTiming)
- Process engagement checks for users who haven't earned anything yet
- Initialize systems that aren't needed for first-time users

**Result**: `/hey` responds instantly on new projects while maintaining full functionality for existing projects.

---

## `/do` - Idea Capture & Discovery

### Overview

The `/do` command captures your project idea through iterative conversation, then conducts a discovery meeting and refinement phase.

### When to Use

- Starting a new project
- Adding a new feature idea
- Fast-tracking with detailed information
- Getting AI-suggested ideas

### Subcommands

#### `/do` (Default - Full Workflow)

Complete ideation workflow with three phases:

1. **Ideation Phase** - Iterative idea capture
2. **Meeting Phase** - Discovery meeting
3. **Refinement Phase** - Enhancements and suggestions

#### `/do feature` - Single Feature Idea

Quickly add a single feature idea to your project.

```bash
/do feature
```

**What happens**:
- Prompts for feature idea
- Appends to IDEA.md
- Skips meeting and refinement
- Updates state to "planning"

#### `/do now` - Fast-Track

Fast-track workflow with detailed prompt or PRD file.

```bash
/do now --prompt "Build a todo app with React and Node.js"
/do now --prd ./path/to/PRD.md
```

**What happens**:
- Combines prompt/PRD into IDEA.md
- Skips meeting and refinement phases
- Updates state to "planning"
- Ready for `/plan` immediately

#### `/do i'm lucky` - AI-Suggested Ideas

Get AI-suggested project ideas and choose one.

```bash
/do i'm lucky
```

**What happens**:
- AI suggests 2 unique ideas
- You choose one or ask for "more"
- AI learns from rejections (never repeats)
- Saves chosen idea to IDEA.md
- Skips meeting and refinement
- Updates state to "planning"

### Ideation Phase (Iterative Conversation)

The ideation phase uses **iterative conversation** to capture complete ideas:

```
You: I want to build a todo app

DoPlan: ✨ Great start! You can tell me more about your idea.
        Often, users remember more details as they talk about it,
        and new improvements come to mind.
        
        Tell me more (or type 'done' when finished):

You: It should have categories

DoPlan: 🙏 Thank you for those details! It's great you didn't forget
        about that - it will definitely improve the results.
        
        You can tell me more, or if you have notes, additional features,
        or any other thoughts, share them now (or type 'done' to finish):

You: And it should sync across devices

DoPlan: 🙏 Thank you for those details!
        
        Tell me more (or type 'done' to finish):

You: done
```

**Benefits**:
- Captures complete ideas
- Encourages thinking through details
- Improves plan quality
- Updates memory card with conversation

### Meeting Phase (Automatic)

After ideation, DoPlan automatically starts a discovery meeting:

```
DoPlan: Let's have a discovery meeting to understand your project better!
        This will help us create a better plan.
        
        [Interactive Q&A session]
```

**What happens**:
- Adaptive questions based on project type
- Experience level consideration
- Generates BRAINSTORM.md
- Updates memory card

### Refinement Phase

After meeting, refinement phase:

```
DoPlan: Great! Now let's refine and enhance your idea with suggestions.
        
        [Refinement process]
```

**What happens**:
- Enhances idea with suggestions
- Generates REFINEMENTS.md
- Updates IDEA.md with improvements
- Updates memory card

### Files Created

- `.do/system/IDEA.md` - Your project idea
- `.do/system/BRAINSTORM.md` - Discovery meeting results
- `.do/system/REFINEMENTS.md` - Refinements and enhancements

### Engagement Integration

- Tracks ideation rounds
- Updates conversation history
- Awards achievements for idea capture
- Learns user preferences
- Builds relationship

---

## `/plan` - Generate Execution Plan

### Overview

The `/plan` command reads your IDEA.md and BRAINSTORM.md, then generates a structured execution plan with phases and tasks.

### When to Use

- After capturing your idea with `/do`
- When IDEA.md and BRAINSTORM.md exist
- Ready to create execution plan

### What It Does

1. **Reads documents** - IDEA.md and BRAINSTORM.md
2. **Generates TASKS.md** - Organized by phases
3. **Creates phase structure** - 01-Foundation, 02-Core, etc.
4. **Generates feature folders** - With templates for each task
5. **Syncs documentation** - To docs/ directory
6. **Updates state** - Sets phase to "planning"
7. **Integrates engagement** - Checks achievements

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

### Engagement Integration

- Awards "Planner" achievement milestones
- Tracks planning completion
- Updates memory card
- Checks challenges (first plan, etc.)

---

## `/dev` - Start Development & Auto-Complete

### Overview

The `/dev` command initiates development for a specific feature or the next available task and now automatically performs completion actions (former `/done` behavior) once the AI agent determines the work is finished.

### When to Use

- Ready to start coding
- After plan is generated
- Starting next task
- Working on specific feature

### Usage

```bash
/dev                    # Start next available task
/dev --feature "auth"  # Start specific feature by name
/dev --feature "1.2"   # Start specific task by ID
```

### What It Does

1. **Finds task** - Next available or specific task
2. **Displays task info** - Title, description, phase, ID
3. **Checks Git status** - Verifies clean working tree
4. **Creates/checks out branch** - `task/[ID]` format
5. **Syncs documentation** - Feature docs to project
6. **Starts time tracking** - Automatic tracking begins
7. **Updates active state** - Sets active_task and active_branch
8. **Shows personalized message** - Based on memory card
9. **Monitors completion** - AI agent decides when work is done
10. **Auto-completes task** - Commits, pushes, records duration, and updates achievements
11. **Integrates engagement** - Checks achievements/challenges

### Example Output

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
⚡ Quick setup complete! Let's build fast!

💡 Working on: User Authentication
   • Feature documentation synced
   • Branch created/checked out

📝 Next steps:
   • Review feature documentation in .do/plan/
   • Start coding with your IDE
   • Keep working—completion runs automatically when finished
```

### Time Tracking

Time tracking starts automatically:
- **Task start time** stored in active_state.json
- **Command execution** tracked in time-tracker.jsonl
- **Full task duration** calculated when `/dev` auto-completes the task

### Auto-Completion Actions (former `/done`, now silent)

When the AI agent determines the task is finished, `/dev` performs the full completion workflow automatically:

- Marks the active task complete
- Stops time tracking and records duration
- Creates a conventional commit and pushes the task branch
- Runs achievement and challenge checks
- Suggests next steps (e.g., PR creation) without requiring another command

### Auto-Commit Format

Conventional commit format automatically:
```
feat(task-2.1): complete user authentication
fix(task-3.2): resolve login bug
docs(task-1.1): update API documentation
```

### Engagement Integration

- Awards "Code Machine" achievement milestones
- Tracks development patterns
- Updates memory card with tech preferences
- Checks development challenges
- Personalized encouragement

## 💡 Command Tips

### Efficiency
- Use `/do now` for fast-tracking
- Use `/do feature` for single features
- Use `/do i'm lucky` for inspiration
- Let `/dev` auto-complete tasks when finished

### Best Practices
- Complete `/hey` tutorial first time
- Use iterative conversation in `/do`
- Review plan before `/dev`
- View `/sys engagement` for motivation

### Power User
- Use `/do now` with detailed prompts
- Use `/do i'm lucky` to explore ideas
- Track time with automatic tracking
- Use state history for rollback
- Leverage engagement system

---

## 🔄 Complete Workflow Example

```bash
# 1. Onboarding
/hey

# 2. Capture idea (iterative)
/do
# → Iterative conversation
# → Discovery meeting
# → Refinement

# 3. Generate plan
/plan

# 4. Start development
/dev
# → Code your task
# → Auto-completes when finished (commit/push/achievements)

# 5. Continue loop
/dev
# → Auto-completes when finished
# ... repeat ...

# 6. View engagement
/sys engagement
```

---

**Next**: [System Commands](./03-System-Commands.md)

