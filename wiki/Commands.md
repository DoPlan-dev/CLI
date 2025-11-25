# Commands Reference

Complete reference for all DoPlan CLI commands. Commands are used within your IDE to interact with the AI agent system.

---

## Command Overview

DoPlan CLI uses a slash-command system (`/command`) that activates specific AI agents to help with your project. Commands are organized into three categories:

- **Core Commands** - Essential commands for the development workflow
- **Team Commands** - Commands for managing and viewing agents
- **Specialized Commands** - Domain-specific commands for advanced use cases

---

## Core Commands

### `/tell` - Capture Project Idea

**Description**: Capture your project idea and save it to the project plan.

**Usage**:
```bash
/tell
/tell <idea>
```

**Examples**:
```bash
/tell
/tell Build a todo app with authentication and dark mode
/tell Create a REST API for a blog platform
```

**What it does**:
1. Captures your project idea (inline or prompts for it)
2. Saves to `.plan/00_System/IDEA.md`
3. Activates Project Orchestrator and Product Manager
4. Updates project state

**Files Modified**:
- `.plan/00_System/IDEA.md`
- `.plan/active_state.json`

**Agents Involved**:
- Project Orchestrator
- Product Manager

**Next Steps**: After capturing your idea, use `/improve` to brainstorm.

---

### `/improve` - Team Brainstorm Session

**Description**: Get ideas and improvements from all Level 1 managers.

**Usage**:
```bash
/improve
```

**Examples**:
```bash
/improve
```

**What it does**:
1. Activates all Level 1 managers (Product, Engineering, Design, QA, Release, Documentation)
2. Conducts a brainstorm session
3. Saves all ideas to `.plan/00_System/BRAINSTORM.md`

**Files Read**:
- `.plan/00_System/IDEA.md`

**Files Modified**:
- `.plan/00_System/BRAINSTORM.md`

**Agents Involved**:
- Product Manager
- Engineering Lead
- Design & UX Manager
- QA & Reliability Manager
- Release & Growth Manager
- Documentation Lead

**Next Steps**: Review BRAINSTORM.md, then use `/write` to generate planning documents.

---

### `/write` - Generate Planning Documents

**Description**: Generate comprehensive planning documents (PRD, Architecture, Design System).

**Usage**:
```bash
/write
```

**Examples**:
```bash
/write
```

**What it does**:
1. Product Manager creates PRD.md
2. Engineering Lead and System Architect create ARCHITECTURE.md
3. Design Manager and UI/UX Designer create DESIGN_SYSTEM.md
4. Saves all documents to `.plan/00_System/`

**Files Read**:
- `.plan/00_System/IDEA.md`
- `.plan/00_System/BRAINSTORM.md`

**Files Modified**:
- `.plan/00_System/PRD.md`
- `.plan/00_System/ARCHITECTURE.md`
- `.plan/00_System/DESIGN_SYSTEM.md`
- `.plan/active_state.json`

**Agents Involved**:
- Product Manager
- Engineering Lead
- System Architect
- Design & UX Manager
- UI/UX Designer
- Project Orchestrator

**Next Steps**: Review the documents, use `/change` to edit, or `/good` to approve.

---

### `/change` - Edit Documents

**Description**: Edit any planning document with natural language instructions.

**Usage**:
```bash
/change <document> <change>
```

**Examples**:
```bash
/change prd Add dark mode support
/change architecture Use PostgreSQL instead of MySQL
/change design Add mobile-first responsive design
```

**What it does**:
1. Parses the document name and change description
2. Loads the specified document
3. Applies the requested changes
4. Saves the updated document

**Files Read**:
- `.plan/00_System/*.md`

**Files Modified**:
- `.plan/00_System/*.md` (specified document)

**Agents Involved**:
- Project Orchestrator

**Document Names**:
- `prd` - Product Requirements Document
- `architecture` - Technical Architecture
- `design` - Design System

---

### `/good` - Approve & Lock Plan

**Description**: Approve the planning documents and lock them for task generation.

**Usage**:
```bash
/good
```

**Examples**:
```bash
/good
```

**What it does**:
1. Validates that PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md exist
2. Sets `locked: true` in active_state.json
3. Updates phase to "approved"

**Files Read**:
- `.plan/00_System/PRD.md`
- `.plan/00_System/ARCHITECTURE.md`
- `.plan/00_System/DESIGN_SYSTEM.md`
- `.plan/active_state.json`

**Files Modified**:
- `.plan/active_state.json`

**Agents Involved**:
- Project Orchestrator

**Next Steps**: Use `/plan` to generate implementation tasks.

---

### `/plan` - Generate Implementation Tasks

**Description**: Generate organized implementation tasks from the approved plan.

**Usage**:
```bash
/plan
```

**Examples**:
```bash
/plan
```

**What it does**:
1. Reads the approved plan (PRD, Architecture, Design System)
2. Generates implementation tasks organized by phases
3. Creates `.plan/TASKS.md`
4. Updates phase to "tasks"

**Files Read**:
- `.plan/00_System/PRD.md`
- `.plan/00_System/ARCHITECTURE.md`
- `.plan/00_System/DESIGN_SYSTEM.md`
- `.plan/active_state.json`

**Files Modified**:
- `.plan/TASKS.md`
- `.plan/active_state.json`

**Agents Involved**:
- Project Orchestrator
- Engineering Lead
- Product Manager

**Next Steps**: Review TASKS.md, then use `/build` to start coding.

---

### `/build` - Start Coding

**Description**: Start implementing the next task or a specific task.

**Usage**:
```bash
/build
/build <task_id>
```

**Examples**:
```bash
/build
/build 1.2
/build 3
```

**What it does**:
1. Determines the task (next uncompleted or specific task_id)
2. Loads task context, dependencies, and related code
3. Activates relevant agents (Frontend Lead, Backend Lead, etc.)
4. Begins implementation
5. Updates active_task in active_state.json

**Files Read**:
- `.plan/TASKS.md`
- `.plan/active_state.json`

**Files Modified**:
- `.plan/active_state.json`
- Source code files

**Agents Involved**:
- Engineering Lead
- Project Orchestrator
- Relevant team leads (based on task)

**GitHub Automation**: After task completion, changes are auto-committed and pushed.

**Next Steps**: Complete the task, then use `/finished` to mark it complete.

---

### `/progress` - Track Progress

**Description**: Show current project progress and status.

**Usage**:
```bash
/progress
```

**Examples**:
```bash
/progress
```

**What it does**:
1. Reads TASKS.md and active_state.json
2. Calculates progress metrics:
   - Total tasks
   - Completed tasks
   - In progress tasks
   - Percentage complete
3. Displays formatted progress report

**Files Read**:
- `.plan/TASKS.md`
- `.plan/active_state.json`

**Agents Involved**:
- Project Orchestrator

**Output Example**:
```
Phase: building
Tasks: 15/30 completed (50%)
Current task: 1.2 - Create API endpoints
Next up: 1.3 - Implement authentication
```

---

### `/state` - Manage Project State History

**Description**: Snapshot, diff, and restore `.plan/active_state.json` safely.

**Usage**:
```bash
/state snapshot --reason "before build 2.1"
/state list --limit 5
/state diff --json
/state restore --file state-20251124T120000Z.json --yes
```

**What it does**:
1. Wraps `go run scripts/statehistory/main.go`.
2. Creates timestamped entries in `.plan/history/state-*.json`.
3. Produces human-readable or JSON diffs between snapshots.
4. Restores state with guardrails (confirmation + optional auto-snapshot).

**Best Practice**: Snapshot before `/build` and right after `/finished` so `/progress` and `/report` can highlight exact deltas.

---

### `/finished` - Complete Task

**Description**: Mark the current task as complete and auto-commit changes.

**Usage**:
```bash
/finished
```

**Examples**:
```bash
/finished
```

**What it does**:
1. Marks task complete in TASKS.md
2. Updates active_state.json (removes active_task, adds to completed)
3. Auto-commits changes with conventional commit format
4. Auto-pushes to current branch
5. Updates CHANGELOG.md if significant changes

**Files Read**:
- `.plan/TASKS.md`
- `.plan/active_state.json`

**Files Modified**:
- `.plan/TASKS.md`
- `.plan/active_state.json`
- CHANGELOG.md (if applicable)

**Agents Involved**:
- Project Orchestrator
- Release Captain

**GitHub Automation**:
- Auto-commit with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update CHANGELOG.md
- Follow branching strategy

**Next Steps**: Use `/build` to start the next task.

---

## Team Commands

### `/team` - Show Agents

**Description**: Display the hierarchical structure of all AI agents and their roles.

**Usage**:
```bash
/team
```

**Examples**:
```bash
/team
```

**What it does**:
1. Loads all agent definitions from `.cursor/agents/`
2. Displays the hierarchical structure
3. Shows each agent's role and responsibilities

**Files Read**:
- `.cursor/agents/*.md`

**Agents Involved**: None (display only)

**Output**: Shows the complete agent hierarchy with roles.

---

### `/load` - Inject Context

**Description**: Load rules or files into agent context for the current session.

**Usage**:
```bash
/load <path>
```

**Examples**:
```bash
/load @library/04-frameworks/nextjs.md
/load .plan/00_System/PRD.md
/load @library/07-databases/postgresql.md
```

**What it does**:
1. Parses the file or directory path
2. Loads the content
3. Injects it into agent context for the current session

**Files Read**:
- `.cursor/rules/library/**`
- `.plan/**`

**Agents Involved**:
- Project Orchestrator

**Use Cases**:
- Load framework-specific rules
- Load database rules
- Load planning documents for reference
- Load custom rules

---

## Operations & Reporting Commands

### `/feedback` - Log Product/Bug Feedback

**Description**: Structured feedback logger that keeps `Docs/history` in sync.

**Usage**:
```bash
/feedback bug "QR download fails" "API returns 500 when Accept header missing" --author QA
/feedback feature "Add dark mode" "Marketing wants a themed hero" --github https://github.com/org/repo/issues/123
```

**What it does**:
1. Parses `type`, `title`, `details`, `--author`, and `--github`.
2. Runs `go run scripts/feedback/main.go`.
3. Appends Markdown entries to `Docs/history/feedback.md`.
4. Updates machine-readable `Docs/history/feedback.json`.
5. Surfaces entries automatically inside `/report`.

---

### `/report` - Generate SCAN Reports & Diffs

**Description**: Produce executive-ready progress reports and change diffs.

**Usage**:
```bash
/report
/report ./test/qr-generator/test-no01 --preset detailed
```

**What it does**:
1. Runs `go run scripts/scanreport/main.go --project <path>`.
2. Builds `.plan/reports/SCAN_REPORT_<date>.md` + JSON metadata.
3. Creates `SCAN_DIFF_<date>.md` by diffing latest vs previous report.
4. Embeds `/progress` snapshot, `.plan/history` deltas, and recent `/feedback`.
5. Supports presets: `standard`, `exec`, `detailed`, or custom `.plan/reports/config.json`.

---

### `/branchci` - Regenerate Branch-Aware Workflows

**Description**: Keeps `.github/workflows/task-branches.yml` aligned with branch policies.

**Usage**:
```bash
/branchci
/branchci regenerate
```

**What it does**:
1. Reads `Docs/history/branch-matrix.json` for branch prefixes + required jobs.
2. Runs `go run scripts/branchci/main.go --matrix Docs/history/branch-matrix.json --out .github/workflows`.
3. Outputs workflows that enforce lint/test/build suites per prefix.

---

### `/github` - Sync KPIs, Issues, and Milestones

**Description**: Bridges GitHub metadata with README + Docs/history cache.

**Usage**:
```bash
/github info
/github issue "Fix cache invalidation" "Details here"
/github milestone "v1 GA" 2025-01-15
```

**What it does**:
- `info`: Updates the README KPI block (between `<!-- KPIS:START -->` markers) and refreshes `Docs/history/github-meta.json`.
- `issue`: Prints a ready-to-run `gh issue create` command with repo slug, title, and body.
- `milestone`: Prints a `gh api` command with milestone name and due date.

---

## Specialized Commands

### `/ship` - Release Management

**Description**: Manage releases, versioning, and deployment.

**Usage**:
```bash
/ship
```

**Examples**:
```bash
/ship
```

**What it does**:
- Coordinates release process
- Manages versioning
- Handles deployment workflows

**Agents Involved**:
- Release Captain
- Release & Growth Manager

---

### `/safe` - Security Audit

**Description**: Perform security audit and review.

**Usage**:
```bash
/safe
```

**Examples**:
```bash
/safe
```

**What it does**:
- Reviews code for security issues
- Checks for vulnerabilities
- Provides security recommendations

**Agents Involved**:
- Security Lead

---

### `/cheap` - Cost Optimization

**Description**: Optimize costs and resource usage.

**Usage**:
```bash
/cheap
```

**Examples**:
```bash
/cheap
```

**What it does**:
- Analyzes resource usage
- Identifies cost optimization opportunities
- Provides recommendations

**Agents Involved**:
- DevOps Engineer
- Performance Engineer

---

## Command Tips & Tricks

### 1. Workflow Order

Follow this order for best results:
1. `/tell` - Capture idea
2. `/improve` - Brainstorm
3. `/write` - Generate plans (`/change` as needed)
4. `/good` - Approve and lock
5. `/plan` - Generate tasks
6. `/state snapshot --reason "pre-build"` - Record baseline
7. `/build` - Start coding
8. `/progress` - Check status
9. `/finished` - Complete task (auto-commit + push)
10. `/state snapshot --reason "post-finish"` - Capture delta
11. `/report` - Share scan metadata/diffs
12. `/safe` + `/cheap` + `/ship` - Prepare release

### 2. Using `/change` Effectively

Be specific with your change requests:
- ✅ Good: `/change prd Add user authentication with OAuth2`
- ❌ Bad: `/change prd fix it`

### 3. Building Specific Tasks

You can build any task by ID:
```bash
/build 1.2    # Build task 1.2
/build 3.5     # Build task 3.5
```

### 4. Loading Context

Load relevant rules before building:
```bash
/load @library/04-frameworks/nextjs.md
/build
```

### 5. Tracking Progress

Check progress regularly:
```bash
/progress
```

### 6. Combining Commands

You can use multiple commands in sequence:
```bash
/tell Build a blog
/improve
/write
/good
/plan
/build
```

---

## Command Syntax

### General Syntax

```
/command [arguments]
```

### Arguments

- **Required arguments**: Must be provided
- **Optional arguments**: Can be omitted
- **Inline arguments**: Can be provided inline or prompted

### Examples

```bash
# No arguments
/tell
/improve
/write

# With inline arguments
/tell Build a todo app
/change prd Add dark mode

# With task ID
/build 1.2
```

---

## Command Reference Table

| Command | Category | Arguments | Auto-commit |
|---------|----------|-----------|-------------|
| `/tell` | Core | Optional: idea | No |
| `/improve` | Core | None | No |
| `/write` | Core | None | No |
| `/change` | Core | Required: doc, change | No |
| `/good` | Core | None | No |
| `/plan` | Core | None | No |
| `/build` | Core | Optional: task_id | No |
| `/progress` | Core | None | No |
| `/state` | Core | Subcommand: snapshot/list/diff/restore | No |
| `/finished` | Core | None | Yes |
| `/team` | Team | None | No |
| `/load` | Team | Required: path/context | No |
| `/feedback` | Operations | Required: type, title; optional flags | No |
| `/report` | Operations | Optional: path, preset | No |
| `/ship` | Specialized | None | No |
| `/safe` | Specialized | None | No |
| `/cheap` | Specialized | None | No |
| `/branchci` | Integrations | Optional: regenerate flag | No |
| `/github` | Integrations | Subcommand: info/issue/milestone | No |

---

## Related Pages

- [Workflow Guide](Workflow) - Complete development workflow
- [Agents Documentation](Agents) - Understanding AI agents
- [Quick Start](Quick-Start) - Getting started guide
- [FAQ](FAQ) - Frequently asked questions
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

