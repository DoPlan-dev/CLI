# Common Patterns

Real-world usage patterns and examples for different scenarios. Learn how others use DoPlan and adapt patterns to your needs.

---

## 🎯 Pattern 1: First Project (Beginner)

### Scenario

New user creating their first project with DoPlan.

### Workflow

```bash
# 1. Onboarding
/hey
# → Complete tutorial
# → Set preferences
# → Learn commands

# 2. Capture idea
/do
# → Iterative conversation
# → "I want to build a todo app"
# → Add details
# → Discovery meeting
# → Refinement

# 3. Generate plan
/plan
# → Review TASKS.md
# → Understand structure

# 4. Start development
/dev
# → Task 1.1 started
# → Code feature
# → Auto-completes when finished (commit/push/achievements)
# → Achievement: "First Steps" (50 points)

# 5. Continue
/dev
/dev
# ... repeat with auto-completion ...

# 6. View engagement
/sys engagement
# → See achievements earned
```

### Key Points

- Complete onboarding thoroughly
- Use iterative conversation
- Review plan before coding
- Let `/dev` auto-complete tasks
- Watch engagement for progress cues

### Expected Outcomes

- First project completed
- Multiple achievements earned
- Relationship level: 10-20
- Score: 100-200 points

---

## 🎯 Pattern 2: Fast Development (Experienced)

### Scenario

Experienced developer with detailed requirements, wants to start coding quickly.

### Workflow

```bash
# 1. Fast-track ideation
/do now --prompt "Build a REST API with Node.js, Express, PostgreSQL, 
authentication, and CRUD operations for a blog system"

# 2. Generate plan
/plan
# → Quick review of TASKS.md

# 3. Rapid development loop
/dev
# → Code (auto-commit/push on completion)

/dev
# → Code (auto-commit/push on completion)
# ... fast iteration ...
```

### Key Points

- Use `/do now` for fast-tracking
- Skip meeting/refinement phases
- Rapid development loop with auto-completion

### Expected Outcomes

- Fast project setup
- Quick development start
- Efficient workflow
- Time saved on ideation

---

## 🎯 Pattern 3: Learning Project (Student)

### Scenario

Student learning new technologies, wants educational support.

### Workflow

```bash
# 1. Onboarding with learning focus
/hey
# → Set experience: "beginner"
# → Set interest: "learning"
# → Set learning goals: ["TypeScript", "Docker"]

# 2. Capture idea
/do
# → "I want to learn TypeScript by building a CLI tool"
# → Detailed conversation
# → Educational focus

# 3. Generate plan
/plan
# → Review plan
# → Understand structure

# 4. Development with learning
/dev
# → DoPlan provides explanations
# → Educational guidance
# → Learning support
# → Auto-completes and awards learning achievements

# 5. Track learning
/sys engagement
# → See learning achievements
```

### Key Points

- Set learning goals
- Use beginner mode
- Request explanations
- Track learning progress
- Earn learning achievements

### Expected Outcomes

- Educational experience
- Learning goals tracked
- Tech exploration achievements
- Skill growth

---

## 🎯 Pattern 4: Feature Addition (Existing Project)

### Scenario

Adding a new feature to an existing project.

### Workflow

```bash
# 1. Add feature idea
/do feature
# → "Add user profile editing"
# → Quick addition

# 2. Update plan (if needed)
/plan
# → Or manually update TASKS.md

# 3. Develop feature
/dev --feature "profile"
# → Specific feature
# → Code (auto-completes when finished)
```

### Key Points

- Use `/do feature` for single features
- Target specific features
- Quick addition
- Minimal workflow

### Expected Outcomes

- Feature added quickly
- Minimal overhead
- Efficient addition

---

## 🎯 Pattern 5: Challenge Completion (Advanced)

### Scenario

Advanced user focusing on completing high-scoring challenges.

### Workflow

```bash
# 1. Regular workflow
/do
/plan
/dev
# → Auto-completes tasks while you build the project

# 2. Complete first API
# → Code API integration
/dev
# → Auto-completes and checks challenges:
#   "API Integration Master" (500 points)
#   "Integration Tested" (750 points)

# 3. Deploy to staging
# → Deploy project
/dev
# → Auto-completes and checks "First Launch!" (1,000 points)

# 4. Achieve test coverage
# → Write tests
# → Achieve 80% coverage
/dev
# → Auto-completes and checks "Test Coverage Champion" (700 points)

# 5. View engagement
/sys engagement
# → See all challenges completed
# → High score achieved
```

### Key Points

- Focus on first-time tasks
- Complete high-scoring challenges
- Deploy to staging/production
- Achieve quality metrics
- Track engagement

### Expected Outcomes

- Multiple challenges completed
- High scores (3,000+ points)
- Epic/legendary achievements
- Strong relationship

---

## 🎯 Pattern 6: Relationship Building (Long-Term)

### Scenario

User building long-term relationship with DoPlan across multiple projects.

### Workflow

```bash
# Project 1
/hey
/do
/plan
/dev (auto-completes) (repeat)
# → Relationship level: 20

# Project 2
/do
/plan
/dev (auto-completes) (repeat)
# → Relationship level: 40
# → Achievement: "Building Connection"

# Project 3
/do
/plan
/dev (auto-completes) (repeat)
# → Relationship level: 60

# Project 5
/do
/plan
/dev (auto-completes) (repeat)
# → Relationship level: 75
# → Achievement: "True Partner"
# → Achievement: "Serial Builder"

# Project 10
/do
/plan
/dev (auto-completes) (repeat)
# → Relationship level: 100
# → Achievement: "Best Friends Forever"
# → Achievement: "Decade Developer"
```

### Key Points

- Use DoPlan regularly
- Complete multiple projects
- Build relationship over time
- Earn relationship achievements
- Maximum personalization

### Expected Outcomes

- Strong relationship (100/100)
- High trust level (10/10)
- Warm communication
- Maximum personalization
- Best Friends Forever achievement

---

## 🎯 Pattern 7: Team Collaboration

### Scenario

Team using DoPlan with shared workflow.

### Workflow

```bash
# Each team member
/hey
# → Set individual preferences

# Shared project
/do
# → Team input
/plan
# → Shared plan

# Individual development
/dev
# → Work on assigned tasks
# → Auto-commit/push and PR suggestion handled automatically
```

### Key Points

- Individual Memory Cards
- Shared project structure
- Individual development
- Coordinated progress
- PR workflow

### Expected Outcomes

- Coordinated development
- Individual personalization
- Shared progress tracking
- Team efficiency

---

## 🎯 Pattern 8: Inspiration Mode

### Scenario

User wants AI-suggested project ideas.

### Workflow

```bash
# 1. Get suggestions
/do i'm lucky
# → AI suggests 2 ideas
# → "Build a habit tracker"
# → "Create a recipe manager"

# 2. Choose or ask for more
# → Choose "habit tracker"
# → Or type "more" for more suggestions

# 3. Generate plan
/plan
# → Plan for chosen idea

# 4. Develop
/dev (auto-completes) (repeat)
```

### Key Points

- Use `/do i'm lucky` for inspiration
- Choose from suggestions
- Ask for more if needed
- AI learns from rejections

### Expected Outcomes

- Creative ideas
- Learning from AI
- Inspiration for projects
- Unique project ideas

---

## 🎯 Pattern 9: Production Release

### Scenario

Completing project and making public release.

### Workflow

```bash
# 1. Complete development
/dev (auto-completes) (repeat)
# → All tasks complete

# 2. Final testing
# → Run tests
# → Achieve coverage
/dev
# → Auto-completes and checks "Test Coverage Champion" (700 points)

# 3. Deploy to production
# → Deploy
/dev
# → Auto-completes and checks "Production Ready" (1,500 points)

# 4. Public release
# → Release publicly
/dev
# → Auto-completes and checks:
#   "Public Launch!" (2,000 points)
#   "Version 1.0" (1,500 points)
# → Total: 3,500 points!

# 5. Celebrate
/sys engagement
# → See all achievements
# → High score!
```

### Key Points

- Complete all tasks
- Achieve quality metrics
- Deploy to production
- Make public release
- Celebrate achievements

### Expected Outcomes

- Project completed
- Multiple legendary challenges
- High scores (5,000+ points)
- Production-ready project

---

## 🎯 Pattern 10: Daily Development

### Scenario

Daily development routine with DoPlan.

### Workflow

```bash
# Morning
/dev
# → Start task and code
# → Auto-completes when finished

# Continue
# → Pick up next task with /dev as needed
# → Auto-completes each task

# End of day
/sys engagement
# → See daily progress
# → Check achievements
# → Stay motivated
```

### Key Points

- Start with `/dev`, let it auto-complete
- Regular development loop
- End with engagement view
- Daily consistency

### Expected Outcomes

- Consistent progress
- Daily streaks
- Regular achievements
- Relationship growth
- Productivity tracking

---

## 💡 Pattern Selection Guide

### Choose Pattern Based On:

**Experience Level**:
- Beginner → Pattern 1 (First Project)
- Intermediate → Pattern 2 (Fast Development)
- Advanced → Pattern 5 (Challenge Completion)

**Goal**:
- Learning → Pattern 3 (Learning Project)
- Speed → Pattern 2 (Fast Development)
- Quality → Pattern 5 (Challenge Completion)
- Relationship → Pattern 6 (Relationship Building)

**Project Type**:
- New Project → Pattern 1 or 2
- Feature Addition → Pattern 4
- Production Release → Pattern 9
- Daily Work → Pattern 10

**Inspiration**:
- Need Ideas → Pattern 8 (Inspiration Mode)

---

## 🎯 Combining Patterns

You can combine patterns:

**Example**: Learning + Challenge Completion
```bash
# Set learning goals
/hey
# → Learning focus

# Regular development
/dev (auto-completes)

# Complete challenges
# → First API (learning + challenge)
# → Deploy (challenge)
# → Tests (learning + challenge)
```

**Example**: Fast Development + Relationship Building
```bash
# Fast-track multiple projects
/do now --prompt "..."
/plan
/dev (auto-completes)

# Build relationship over time
# → Multiple projects
# → Relationship grows
# → Better personalization
```

---

**Next**: [Features Guide](../06-Features/)

