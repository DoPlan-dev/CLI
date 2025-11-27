---
name: build
category: core
trigger: "/build [<task_id>]"
description: "Start coding next task"
agentInvolvement:
  - Engineering Lead
  - Relevant Team Leads
  - Project Orchestrator
  - QA Engineer
filesRead:
  - ".do/plan/TASKS.md"
  - ".do/system/history/active_state.json"
filesModified:
  - ".do/system/history/active_state.json (active_task and active_branch updated)"
  - ".do/system/history/state-*.json (automatic snapshot for audit/rollback)"
  - ".do/plan/TASKS.md (task marked complete if auto-detected)"
  - "Git: New branch created/checked out (task/[ID])"
  - "src/** (code files created/modified)"
githubAutomation: |
  After task completion, the system will:
  - Auto-commit changes with conventional commit format
  - Auto-push to current branch (feature/bugfix/hotfix)
  - Update docs/history/CHANGELOG.md if significant changes
  - Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md
examples:
  - "/build → Start next uncompleted task"
  - "/build 1.2 → Start specific task 1.2"
  - "/build 3 → Start task 3"
---

When user types /build or /build <task_id>:

1. **Determine Task**:
   - If task_id provided, load that task
   - Otherwise, find next uncompleted task from TASKS.md

2. **Bootstrap Boilerplate (first run only)g 
   - If the project is still plan-only (no package.json / src/), prompt the user to scaffold code with their preferred stack tool (e.g., `npx create-next-app`, `go mod init`, etc.)
   - DoPlan no longer ships the legacy `scripts/boilerplate` helper, so projects are expected to bring or generate their own starter code
   - Skip once code already exists for this project/stack

3. **Check Git Status**:
   - Verify working tree is clean (no uncommitted changes)
   - If dirty, warn user and block until clean

4. **Create/Checkout Task Branch**:
   - Create or checkout branch `task/[ID]` manually (e.g., `git checkout -b task/5.2`)
   - Store the selected branch name in `active_branch` inside `.do/system/history/active_state.json`

5. **Load Task Context**: Read task details, dependencies, and related code

6. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)

7. **Start Implementation**: Begin coding the task with full context

8. **Update State**: Set `active_task` and `active_branch` in `.do/system/history/active_state.json`

9. **Snapshot State**: Immediately log the new state with `go run scripts/statehistory/main.go snapshot --reason "build [ID]" --label build`

10. **After Task Implementation** (when agent detects completion):
    - Agent analyzes code changes, tests, requirements
    - Agent checks if task criteria are met
    - If complete:
      - Agent asks: "Task [ID] appears complete. Summary:
        ✅ All requirements met
        ✅ Code implemented
        ✅ Tests passing (if applicable)
        
        Mark as done? (yes/no)"
      - If user says yes:
        - Mark task complete in TASKS.md
        - Update active_state.json
        - Auto-commit and push
        - Response: "Task marked complete! Changes committed and pushed."
      - If user says no:
        - Continue working on task

11. **Response**: "Building task [ID]: [Description] on branch [branch_name]. Focus on this task only."

