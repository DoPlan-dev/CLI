# /build

## Trigger
/build or /build <task_id>

## Examples
- /build
- /build 1.2
- /build 3


## Action
When user types /build or /build <task_id>:

1. **Determine Task**: 
   - If task_id provided, load that task
   - Otherwise, find next uncompleted task from TASKS.md
2. **Load Task Context**: Read task details, dependencies, and related code
3. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)
4. **Prepare Git Branch**:
   - Ensure working tree is clean (`git status -sb`)
   - Determine branch name: `branch = task/<TASK-ID>-<slug>` (lowercase slug from task title, max 4 words). Example: `task/TASK-053-error-messages`
   - Switch to base branch (main/develop), pull latest
   - If `branch` exists → `git switch branch`; otherwise → `git switch -c branch`
   - Record `branch` in `.plan/active_state.json` under `active_branch`
5. **Start Implementation**: Begin coding the task with full context on the new branch
6. **Update State**: Set `active_task` (and `active_branch`) in `.plan/active_state.json`
7. **Snapshot State**: `go run scripts/statehistory/main.go snapshot --reason "build [ID]" --label build` so the new branch + task pairing is recorded
8. **Response**: "Building task [ID] on branch [branch]. Focus on this task only."

## Agent Involvement
- **Engineering Lead**
- **Project Orchestrator**

## Files Read
- .plan/TASKS.md
- .plan/active_state.json

## Files Modified
- .plan/active_state.json (active_task + active_branch)
- .plan/history/state-*.json (automatic snapshot for rollback)

## GitHub Automation
After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push the `task/<TASK-ID>-<slug>` branch
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md
