# Development Workflow Guide

Complete guide to the DoPlan CLI development workflow from idea to release.

---

## Overview

DoPlan CLI uses a structured development workflow that guides you from initial idea to production release. The workflow is organized into four main phases:

1. **Planning Phase** - Capture idea, brainstorm, and create planning documents
2. **Development Phase** - Generate tasks and implement features
3. **Review Phase** - Review code, test, and ensure quality
4. **Release Phase** - Prepare and deploy releases

---

## Complete Development Workflow

### Workflow Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                    DoPlan CLI Workflow                      │
└─────────────────────────────────────────────────────────────┘

1. IDEA PHASE
   ┌──────────────┐
   │  /tell        │ → Capture project idea
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /improve    │ → Brainstorm with team
   └──────┬───────┘
          │
2. PLANNING PHASE
   ┌──────▼───────┐
   │  /write      │ → Generate PRD, Architecture, Design
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /change     │ → Edit documents (optional)
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /good       │ → Approve and lock plan
   └──────┬───────┘
          │
3. TASKS PHASE
   ┌──────▼───────┐
   │  /plan       │ → Generate implementation tasks
   └──────┬───────┘
          │
4. DEVELOPMENT PHASE
   ┌──────▼───────┐
   │  /state      │ → Snapshot before coding
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /build      │ → Start coding task
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /progress   │ → Check progress (optional)
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /finished   │ → Complete task (auto-commit)
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  /state      │ → Snapshot after coding
   └──────┬───────┘
          │
   ┌──────▼───────┐
   │  Repeat      │ → Continue with next task
   └──────────────┘
          │
5. OPERATIONS PHASE
   ┌──────▼─────────────┐
   │  /report + /feedback│ → Share scans + capture input
   └──────┬──────────────┘
          │
6. RELEASE PHASE
   ┌──────▼───────┐
   │  /ship       │ → Release management (after /safe + /cheap)
   └──────────────┘
```

---

## Phase 1: Planning Phase

The planning phase transforms your idea into a comprehensive plan.

### Step 1: Capture Your Idea (`/tell`)

**Command**: `/tell` or `/tell <idea>`

**What happens**:
1. Your project idea is captured
2. Saved to `.plan/00_System/IDEA.md`
3. Project Orchestrator and Product Manager are activated
4. Project state updated to "idea" phase

**Example**:
```bash
/tell Build a modern todo app with authentication, dark mode, and offline support
```

**Output**: Idea saved to `.plan/00_System/IDEA.md`

**Next Step**: Use `/improve` to brainstorm

---

### Step 2: Brainstorm (`/improve`)

**Command**: `/improve`

**What happens**:
1. All Level 1 managers are activated:
   - Product Manager
   - Engineering Lead
   - Design & UX Manager
   - QA & Reliability Manager
   - Release & Growth Manager
   - Documentation Lead
2. Each manager provides ideas and improvements
3. All ideas saved to `.plan/00_System/BRAINSTORM.md`

**Example**:
```bash
/improve
```

**Output**: Brainstorm ideas saved to `.plan/00_System/BRAINSTORM.md`

**Next Step**: Use `/write` to generate planning documents

---

### Step 3: Generate Planning Documents (`/write`)

**Command**: `/write`

**What happens**:
1. **Product Manager** creates `PRD.md` (Product Requirements Document)
2. **Engineering Lead** and **System Architect** create `ARCHITECTURE.md`
3. **Design Manager** and **UI/UX Designer** create `DESIGN_SYSTEM.md`
4. All documents saved to `.plan/00_System/`
5. Project state updated to "writing" phase

**Example**:
```bash
/write
```

**Output**: Three planning documents generated:
- `.plan/00_System/PRD.md`
- `.plan/00_System/ARCHITECTURE.md`
- `.plan/00_System/DESIGN_SYSTEM.md`

**Next Step**: Review documents, use `/change` to edit, or `/good` to approve

---

### Step 4: Edit Documents (Optional) (`/change`)

**Command**: `/change <document> <change>`

**What happens**:
1. Document is loaded
2. Changes are applied
3. Updated document is saved

**Examples**:
```bash
/change prd Add dark mode support
/change architecture Use PostgreSQL instead of MySQL
/change design Add mobile-first responsive design
```

**Next Step**: Use `/good` to approve when ready

---

### Step 5: Approve Plan (`/good`)

**Command**: `/good`

**What happens**:
1. Validates that all planning documents exist
2. Sets `locked: true` in `.plan/active_state.json`
3. Updates phase to "approved"
4. Plan is now locked and ready for task generation

**Example**:
```bash
/good
```

**Output**: Plan approved and locked

**Next Step**: Use `/plan` to generate implementation tasks

---

## Phase 2: Development Phase

The development phase transforms your approved plan into working code.

### Step 1: Generate Tasks (`/plan`)

**Command**: `/plan`

**What happens**:
1. Reads approved plan (PRD, Architecture, Design System)
2. **Project Orchestrator**, **Engineering Lead**, and **Product Manager** generate tasks
3. Tasks organized by phases
4. Tasks saved to `.plan/TASKS.md`
5. Project state updated to "tasks" phase

**Example**:
```bash
/plan
```

**Output**: Implementation tasks generated in `.plan/TASKS.md`

**Next Step**: Snapshot state before you begin coding.

---

### Step 2: Snapshot State (`/state snapshot`)

**Command**: `/state snapshot --reason "before build <task>" --label build`

**What happens**:
1. Captures the current `.plan/active_state.json`.
2. Stores a timestamped entry in `.plan/history/`.
3. Tags the snapshot with your reason/label for quick recall.
4. Forms the baseline that `/report` and `/progress` will diff against.

**Example**:
```bash
/state snapshot --reason "before build 2.1" --label pre-build
```

**Output**: `Snapshot saved: .plan/history/state-20251125T120000Z-pre-build.json`

**Next Step**: Start coding with `/build`.

---

### Step 3: Start Coding (`/build`)

**Command**: `/build` or `/build <task_id>`

**What happens**:
1. Determines task (next uncompleted or specific task_id)
2. Loads task context, dependencies, and related code
3. Activates relevant agents (Frontend Lead, Backend Lead, etc.)
4. Begins implementation
5. Updates `active_task` in `.plan/active_state.json`

**Examples**:
```bash
/build          # Start next uncompleted task
/build 1.2      # Start specific task 1.2
/build 3        # Start task 3
```

**Output**: Task implementation begins

**Next Step**: Track progress or complete the task, then use `/finished`.

---

### Step 4: Track Progress (Optional) (`/progress`)

**Command**: `/progress`

**What happens**:
1. Reads TASKS.md and active_state.json
2. Calculates progress metrics
3. Displays formatted progress report

**Example**:
```bash
/progress
```

**Output**:
```
Phase: building
Tasks: 15/30 completed (50%)
Current task: 1.2 - Create API endpoints
Next up: 1.3 - Implement authentication
```

**Next Step**: Continue with `/build` or `/finished`.

---

### Step 5: Complete Task (`/finished`)

**Command**: `/finished`

**What happens**:
1. Marks task complete in TASKS.md
2. Updates active_state.json (removes active_task, adds to completed)
3. **Auto-commits** changes with conventional commit format
4. **Auto-pushes** to current branch
5. Updates CHANGELOG.md if significant changes
6. Triggers CI workflow

**Example**:
```bash
/finished
```

**Output**: Task completed, changes committed and pushed

**Next Step**: Snapshot state again (`/state snapshot --reason "after build <task>"`), then loop back to `/build` or move toward `/report`/`/ship`.

---

## Phase 3: Review Phase

The review phase ensures code quality before release.

### Code Review

**Process**:
1. Review auto-committed changes
2. Check CI workflow status
3. Review code quality
4. Test functionality

**Agents Involved**:
- QA Engineer
- Security Lead
- Performance Engineer

**Commands**:
- `/safe` - Security audit
- `/progress` - Check completion status

---

### Testing

**Process**:
1. Run automated tests
2. Manual testing
3. Integration testing
4. Performance testing

**Agents Involved**:
- QA Engineer
- QA & Reliability Manager

---

### Reporting & Stakeholder Updates

**Purpose**: Keep leadership, QA, and product stakeholders informed with consistent artifacts.

**Commands**:
- `/report [path] [--preset exec|detailed]`
  - Runs `scripts/scanreport` to refresh `.plan/reports/SCAN_REPORT_<date>.md` + JSON.
  - Builds `SCAN_DIFF_<date>.md` highlighting state changes, `/progress` stats, visuals, and outstanding feedback.
- `/feedback <type> "Title" "Details" [--author ...] [--github url]`
  - Logs structured entries to `Docs/history/feedback.md` and `.json`.
  - Feeds subsequent `/report` runs so regressions, bugs, and feature ideas stay visible.

**Best Practice**:
1. Run `/report` after every major `/finished` milestone or before demos.
2. Capture stakeholder reactions via `/feedback` immediately so they are traceable.

---

## Phase 4: Release Phase

The release phase prepares and deploys your project.

### Release Management (`/ship`)

**Command**: `/ship`

**What happens**:
1. **Release Captain** coordinates release process
2. Manages versioning
3. Handles deployment workflows
4. Updates release notes

**Example**:
```bash
/ship
```

**Agents Involved**:
- Release Captain
- Release & Growth Manager

---

### Operational Automation

Use integration helpers to keep automation and GitHub metadata aligned with the workflow.

#### `/branchci`
- Reads `Docs/history/branch-matrix.json`.
- Regenerates `.github/workflows/task-branches.yml` so task/, feature/, hotfix/ branches run the right checks.
- Run it whenever you adjust branch naming or CI requirements.

#### `/github info|issue|milestone`
- `info` updates the README KPI block and refreshes `Docs/history/github-meta.json`.
- `issue` and `milestone` print ready-to-run `gh` commands that inherit the detected repo slug.
- Keeps documentation KPIs and GitHub tracking in lockstep with the CLI workflow.

---

## Workflow Best Practices

### 1. Follow the Workflow Order

Always follow the workflow in order:
1. `/tell` → Capture idea
2. `/improve` → Brainstorm
3. `/write` + `/change` → Generate/refine plans
4. `/good` → Approve
5. `/plan` → Generate tasks
6. `/state snapshot` → Capture baseline
7. `/build` → Start coding
8. `/progress` → Check status
9. `/finished` → Complete task (auto-commit/push)
10. `/state snapshot` → Capture delta
11. `/report` + `/feedback` → Inform stakeholders
12. `/safe` + `/cheap` + `/ship` → Release

### 2. Review Before Approving

Always review planning documents before using `/good`:
- Check PRD for completeness
- Verify Architecture is sound
- Ensure Design System is comprehensive

### 3. Use `/change` Liberally

Don't hesitate to use `/change` to refine documents:
- Add missing features
- Fix technical decisions
- Update design requirements

### 4. Track Progress Regularly

Use `/progress` and `/state diff` regularly to:
- Monitor completion status
- Identify blockers
- Capture precise before/after deltas
- Plan next steps

### 5. Complete Tasks Incrementally

Work on one task at a time:
- Focus on current task
- Complete before moving on
- Use `/finished` to track completion

### 6. Leverage Agents & Automation

Use specialized commands when needed:
- `/report` - Executive-ready scan reports
- `/feedback` - Structured stakeholder input
- `/safe` - Security review
- `/cheap` - Cost optimization
- `/ship` - Release management
- `/branchci` - CI guardrails per branch type
- `/github info|issue|milestone` - Keep KPIs + GitHub tracking synced

---

## Common Workflow Patterns

### Pattern 1: Rapid Prototyping

For quick prototypes:
1. `/tell` - Capture idea
2. `/write` - Generate plans
3. `/good` - Approve quickly
4. `/plan` - Generate tasks
5. `/build` - Start coding immediately

### Pattern 2: Thorough Planning

For complex projects:
1. `/tell` - Capture idea
2. `/improve` - Extensive brainstorming
3. `/write` - Generate plans
4. `/change` - Multiple iterations
5. `/good` - Approve when perfect
6. `/plan` - Generate detailed tasks
7. `/build` - Implement carefully

### Pattern 3: Iterative Development

For iterative projects:
1. Complete planning phase
2. Generate initial tasks
3. Build and complete tasks
4. Use `/change` to update plans
5. Regenerate tasks with `/plan`
6. Continue building

### Pattern 4: Feature Development

For adding features:
1. `/change prd Add new feature`
2. `/change architecture Update architecture for feature`
3. `/good` - Approve changes
4. `/plan` - Generate new tasks
5. `/build` - Implement feature

---

## Workflow Examples

### Example 1: Simple Todo App

```bash
# 1. Capture idea
/tell Build a simple todo app

# 2. Brainstorm
/improve

# 3. Generate plans
/write

# 4. Approve
/good

# 5. Generate tasks
/plan

# 6. Build
/build
# ... implement task ...
/finished

# 7. Continue
/build
/finished
# ... repeat ...
```

### Example 2: Complex Web Application

```bash
# 1. Capture detailed idea
/tell Build a social media platform with authentication, posts, comments, and real-time notifications

# 2. Extensive brainstorming
/improve

# 3. Generate comprehensive plans
/write

# 4. Refine documents
/change prd Add user profiles
/change architecture Use microservices
/change design Add dark mode

# 5. Approve when ready
/good

# 6. Generate tasks
/plan

# 7. Build incrementally
/build
# ... work on task ...
/finished

# 8. Track progress
/progress

# 9. Continue building
/build
/finished
# ... repeat for all tasks ...
```

---

## State Management

The workflow uses `active_state.json` to track progress:

```json
{
  "phase": "idea|brainstorm|writing|approved|tasks|building",
  "active_task": null | "task_id",
  "completed": ["task_id", ...],
  "locked": false | true
}
```

**Phases**:
- `idea` - Idea captured
- `brainstorm` - Brainstorming complete
- `writing` - Planning documents generated
- `approved` - Plan approved and locked
- `tasks` - Tasks generated
- `building` - Actively building

---

## Related Pages

- [Commands Reference](Commands) - All available commands
- [Agents Documentation](Agents) - Understanding AI agents
- [Quick Start](Quick-Start) - Getting started guide
- [First Project Tutorial](First-Project-Tutorial) - Detailed tutorial
- [Best Practices](Best-Practices) - Development best practices
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

