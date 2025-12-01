# DoPlan CLI - Complete Mind Map

## 🧠 Core Philosophy
**DoPlan transforms development from a chore into an engaging, educational, and fun experience.**

---

## 📊 SYSTEM ARCHITECTURE

### 1. COMMAND SYSTEM
```
Core Commands (Workflow)
├── /hey          → Onboarding, tutorial, welcome
├── /do           → Idea capture & discovery
│   ├── /do app idea prompt
│   ├── /do feature
│   ├── /do now
│   └── /do i'm lucky
├── /plan         → Generate execution plan
├── /dev          → Start development workflow (auto-detects completion)
└── /sys          → System control panel
    ├── /sys status → View progress
    ├── /sys engagement → Engagement dashboard
    ├── /sys role → Role management
    ├── /sys security → Security settings
    └── /sys control → System control

System Commands (Control)
└── /sys          → System control panel
    ├── /sys engagement  → Engagement dashboard
    ├── /sys role        → Role management
    ├── /sys security    → Security settings
    └── /sys control     → System control (kill switch)
```

### 2. ENGAGEMENT SYSTEM (The Fun Factor)
```
Engagement Orchestrator
├── Brain System
│   ├── Memory Card Integration
│   ├── Personalized Prompts
│   ├── Tone of Voice Adjustment
│   └── Context-Aware Responses
│
├── Achievement System
│   ├── 200+ Achievements
│   ├── Score System (0-100,000+)
│   ├── Categories:
│   │   ├── Score Milestones
│   │   ├── Project Achievements
│   │   ├── Command Usage
│   │   ├── Learning Goals
│   │   ├── Productivity
│   │   ├── Streaks
│   │   ├── Relationship
│   │   └── Special Events
│   └── Celebration System
│
├── Challenge System
│   ├── 30+ High-Scoring Challenges
│   ├── 300-2000 Points Each
│   ├── Categories:
│   │   ├── Integration (API, Webhooks)
│   │   ├── Database (Migrations, Merges)
│   │   ├── Deployment (Staging, Production, CI/CD)
│   │   ├── Testing (Coverage, E2E, Performance)
│   │   ├── Workflow (GitHub, Commits, PRs)
│   │   ├── Release (Public Launch, v1.0)
│   │   ├── Performance (Optimization)
│   │   └── Security (Audits, Vulnerabilities)
│   └── First-Time Detection
│
└── Dopamine Timing System
    ├── Strategic Reward Delays
    ├── Anticipation Building
    ├── Multi-Reward Bursts
    └── Psychological Optimization
```

### 3. MEMORY CARD SYSTEM (The Relationship)
```
Memory Card (~/.doplan/memory_card.json)
├── User Identity
│   ├── Name, First Met, Projects Count
│   └── Last Interaction
│
├── Preferences & Personality
│   ├── Interest (learning/develop)
│   ├── Work Style (fast/thoughtful)
│   ├── Personality (thinker/copier)
│   ├── Dream (change_world/build_others)
│   ├── Motivation (money/success)
│   └── Experience Level
│
├── Communication Preferences
│   ├── Style (brief/detailed/balanced)
│   ├── Feedback Frequency
│   ├── Detail Level
│   ├── Encouragement Style
│   └── Error Handling Preference
│
├── Learning & Preferences
│   ├── Preferred Tech Stack
│   ├── Project Types
│   ├── Interests
│   ├── Learning Goals
│   └── Pain Points
│
├── Relationship Data
│   ├── Conversation History
│   ├── Memorable Moments
│   ├── Achievements
│   └── Preferences Map
│
├── Usage Patterns
│   ├── Command Usage Stats
│   ├── Favorite Commands
│   ├── Struggled Features
│   ├── Helpful Features
│   └── Time Preferences
│
├── Challenge Tracking
│   ├── Completed Challenges
│   └── Challenge Attempts
│
├── Relationship Metrics
│   ├── Tone Level (0-10)
│   ├── Relationship Level (0-100)
│   ├── Trust Level (0-10)
│   ├── Engagement Score (0-1)
│   └── Score (Total Points)
│
└── Context Awareness
    ├── Current Project
    ├── Current Phase
    ├── Last Command
    ├── Session Count
    └── Average Session Time
```

### 4. TIME TRACKING SYSTEM
```
Time Tracker (.do/system/time-tracker.jsonl)
├── Automatic Tracking
│   ├── Command Execution
│   ├── Phase Duration
│   ├── Task Duration (from /dev start to completion detection)
│   └── Session Duration
│
├── Metadata
│   ├── Project Path
│   ├── Command Name
│   ├── Phase
│   ├── Args
│   ├── Work Style
│   └── Experience Level
│
└── Analytics Ready
    ├── JSONL Format
    ├── Timestamped
    └── Status Tracking
```

### 5. STATE MANAGEMENT SYSTEM
```
State System
├── Active State (.do/system/history/active_state.json)
│   ├── Phase
│   ├── Active Task
│   ├── Active Branch
│   ├── Completed Tasks
│   ├── Task Started At
│   └── Locked Status
│
└── State History (.do/system/history/state-*.json)
    ├── Automatic Snapshots
    ├── Before/After /dev
    ├── On task completion (auto-detected)
    ├── Diff Capability
    └── Rollback Support
```

### 6. PROJECT STRUCTURE
```
Project Root
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

## 🎯 WORKFLOW PHASES

### Phase 1: ONBOARDING (/hey)
- First-time welcome
- Interactive tutorial
- System overview
- Agent hierarchy explanation
- Command walkthrough
- Test drive mode
- Personalized tips
- Reference materials creation

### Phase 2: IDEATION (/do)
- Iterative idea capture
- Multiple conversation rounds
- Encouragement and guidance
- IDEA.md generation
- Meeting phase (discovery)
- Refinement phase
- BRAINSTORM.md & REFINEMENTS.md

### Phase 3: PLANNING (/plan)
- Read IDEA.md & BRAINSTORM.md
- Generate TASKS.md
- Create phase structure
- Feature folders with templates
- Documentation sync
- Engagement integration

### Phase 4: DEVELOPMENT (/dev)
- Task selection (next or specific)
- Git branch creation
- Documentation sync
- Time tracking start
- Engagement processing
- Personalized messages
- Memory card updates

### Phase 5: COMPLETION (Auto-Detected by /dev)
- Task completion detection
- Completion verification
- Dependency checking
- TASKS.md update
- State update
- State snapshot
- Auto-commit (conventional format)
- Auto-push
- Changelog update
- PR suggestion
- Achievement/challenge checking
- Duration display

### Phase 6: MONITORING (/sys status)
- Progress tracking
- State deltas
- Task statistics
- Phase overview

---

## 🚀 KEY FEATURES

### 1. PERSONALIZATION
- Brain-powered agent responses
- Memory card-driven adaptation
- Relationship-based tone
- Context-aware suggestions
- Learning goal integration

### 2. GAMIFICATION
- 200+ achievements
- 30+ high-scoring challenges
- Score system (0-100,000+)
- Dopamine timing optimization
- Celebration system
- Multi-achievement bursts

### 3. LEARNING SUPPORT
- Beginner-friendly guidance
- Educational explanations
- Pain point assistance
- Learning goal tracking
- Tech stack preferences
- Experience level adaptation

### 4. AUTOMATION
- Auto-commit (conventional format)
- Auto-push to remote
- State snapshots
- Time tracking
- Changelog updates
- PR suggestions

### 5. TRACKING & ANALYTICS
- Time tracking (JSONL)
- Command usage stats
- Session tracking
- Progress metrics
- Engagement dashboard

### 6. SAFETY & CONTROL
- State history (rollback)
- Dependency checking
- Branch verification
- System control panel
- Security settings
- Role management

---

## 🎓 EDUCATIONAL VALUE

### For Beginners
- Step-by-step guidance
- Explanations at every step
- Test drive mode
- Learning goal support
- Pain point assistance
- Encouragement system

### For Professionals
- Fast workflow
- Detailed technical info
- Advanced features
- Customization options
- Power user features
- Professional automation

---

## 💡 UNIQUE SELLING POINTS

1. **Engagement-First Design**: Makes development fun and exciting
2. **Learning-Focused**: Built for both learning and development
3. **Relationship Building**: Memory card creates personal connection
4. **Gamification**: Achievements and challenges motivate users
5. **Personalization**: Brain system adapts to each user
6. **Time Tracking**: Automatic, comprehensive time analytics
7. **State Management**: Full history and rollback capability
8. **Automation**: Git, commits, pushes, PRs all automated
9. **Beginner-Friendly**: Onboarding and guidance for all levels
10. **Professional-Grade**: Production-ready from day one

---

## 🔄 USER JOURNEY

```
New User
  ↓
/hey (Onboarding)
  ├── Tutorial
  ├── System Overview
  ├── Test Drive
  └── Reference Materials
  ↓
/do (Ideation)
  ├── Iterative Idea Capture
  ├── Discovery Meeting
  └── Refinement
  ↓
/plan (Planning)
  ├── TASKS.md Generation
  ├── Phase Structure
  └── Feature Templates
  ↓
/dev (Development Loop)
  ├── Task Selection
  ├── Branch Creation
  ├── Development
  └── Auto-Completion Detection
      ├── Auto-Commit
      ├── Auto-Push
      ├── Achievements
      └── Next Task
  ↓
Repeat /dev (auto-detects completion)
  ↓
/sys status (Progress)
  └── Dashboard View
  ↓
/sys engagement (Engagement)
  └── Full Dashboard
```

---

## 🎯 TARGET AUDIENCES

### Beginners
- Learning-focused features
- Step-by-step guidance
- Educational explanations
- Test drive mode
- Encouragement system

### Intermediate
- Balanced guidance
- Learning goals
- Tech exploration
- Achievement hunting

### Advanced
- Fast workflow
- Power features
- Customization
- Professional automation
- Challenge completion

---

## 📈 METRICS & TRACKING

### User Metrics
- Score (0-100,000+)
- Achievements (200+)
- Challenges (30+)
- Relationship Level (0-100)
- Trust Level (0-10)
- Engagement Score (0-1)

### Project Metrics
- Tasks Completed
- Time Spent
- Commands Used
- Phases Completed
- Features Developed

### Learning Metrics
- Learning Goals Set
- Tech Stack Explored
- Pain Points Overcome
- Sessions Completed

---

## 🛠️ TECHNICAL ARCHITECTURE

### Languages & Frameworks
- Go (CLI backend)
- Cobra (CLI framework)
- Bubbletea (TUI)
- JSON/JSONL (Data storage)
- Markdown (Documentation)

### Data Storage
- Memory Card (JSON, ~/.doplan/)
- Time Tracker (JSONL, .do/system/)
- State History (JSON, .do/system/history/)
- Project Files (Markdown, .do/)

### Integration Points
- Git (auto-commit/push)
- GitHub CLI (PR suggestions)
- IDE Support (6 IDEs)
- CI/CD (workflow automation)

---

## 🎨 DESIGN PRINCIPLES

1. **User-Centric**: Everything adapts to the user
2. **Engagement-First**: Fun and motivating
3. **Learning-Focused**: Educational at every step
4. **Automation-Heavy**: Reduce manual work
5. **Transparent**: All logic visible in markdown
6. **Safe**: State history and rollback
7. **Personal**: Memory card builds relationship
8. **Flexible**: Works for all skill levels
9. **Professional**: Production-ready output
10. **Comprehensive**: Full workflow coverage

---

## 🔮 FUTURE POTENTIAL

### Analytics Dashboard
- Visual progress charts
- Time analysis
- Achievement gallery
- Challenge progress
- Learning path visualization

### IDE Extension
- Real-time integration
- In-editor achievements
- Live progress tracking
- Command palette integration

### Community Features
- Achievement sharing
- Leaderboards
- Challenge competitions
- Learning groups

### Advanced AI
- Predictive suggestions
- Proactive help
- Pattern recognition
- Optimization recommendations

---

## 📚 DOCUMENTATION STRUCTURE

```
Wiki/
├── 01-Getting-Started/
├── 02-Commands/
├── 03-Engagement-System/
├── 04-Memory-Card/
├── 05-Workflow/
├── 06-Features/
├── 07-Advanced/
└── 08-Reference/
```

---

**This mind map represents the complete DoPlan ecosystem - a revolutionary tool that makes development engaging, educational, and fun for everyone.**

