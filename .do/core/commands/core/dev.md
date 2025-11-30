---
name: dev
category: core
trigger: "/dev or /dev <task_id>"
description: "Start development on a task"
agentInvolvement:
  - Engineering Lead
  - Relevant Team Leads
  - Project Orchestrator
filesRead:
  - ".do/plan/TASKS.md"
  - ".do/system/history/active_state.json"
filesModified:
  - ".do/system/history/active_state.json (active_task and active_branch updated)"
  - ".do/system/history/state-*.json (automatic snapshot for audit/rollback)"
  - "Git: New branch created/checked out (task/[ID])"
  - "src/** (code files created/modified)"
examples:
  - "/dev → Start next uncompleted task"
  - "/dev 1.2 → Start specific task 1.2"
  - "/dev 3 → Start task 3"
---

# /dev

## Trigger
Exact match: /dev or /dev <task_id>

Examples:
- /dev → Start next uncompleted task
- /dev 1.2 → Start specific task 1.2
- /dev 3 → Start task 3

## Action
When user types /dev or /dev <task_id>:

1. **Determine Task**: 
   - If task_id provided, load that task
   - Otherwise, find next uncompleted task from TASKS.md
2. **Bootstrap Boilerplate (first run only)**:
   - If source files do not exist yet, run `go run scripts/boilerplate/main.go --project .`
   - This materializes the appropriate stack (Next.js by default) so the task has code to modify
   - Skip if the project already contains code for this task/stack
3. **Check Git Status**: 
   - Verify working tree is clean (no uncommitted changes)
   - If dirty, warn user and block until clean
4. **Create/Checkout Task Branch**: 
   - Run `go run scripts/branch/main.go --action create --task [ID] --project .`
   - This creates/checks out branch `task/[ID]` (e.g., `task/5.2`)
   - Store branch name in `active_branch` field of `.do/system/history/active_state.json`
5. **Load Task Context**: Read task details, dependencies, and related code
6. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)
7. **Start Implementation**: Begin coding the task with full context
8. **Update State**: Set `active_task` and `active_branch` in `.do/system/history/active_state.json`
9. **Snapshot State**: Immediately log the new state with `go run scripts/statehistory/main.go snapshot --reason "dev [ID]" --label dev`
10. **Response**: "Building task [ID]: [Description] on branch [branch_name]. Focus on this task only."

## Agent Involvement
- **Engineering Lead**: Coordinates task execution
- **Relevant Team Leads**: Frontend Lead, Backend Lead, etc. based on task
- **Project Orchestrator**: Monitors progress

## Files Read
- .do/plan/TASKS.md
- .do/system/history/active_state.json

## Files Modified
- .do/system/history/active_state.json (active_task and active_branch updated)
- .do/system/history/state-*.json (automatic snapshot for audit/rollback)
- Git: New branch created/checked out (task/[ID])
- src/** (code files created/modified)

## GitHub Automation
After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md

