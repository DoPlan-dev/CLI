# /dev

## Trigger
/dev [<task_id>]

## Examples
- "/dev → Start next task"
- "/dev --feature \"auth\" → Start specific feature"
- "/dev 1.2 → Start specific task 1.2"


## Action
When user types /dev or /dev <task_id>:

1. **Determine Task**:
   - If task_id provided, load that task
   - Otherwise, find next uncompleted task from TASKS.md

2. **Bootstrap Boilerplate (first run only)**:
   - If the project is still plan-only (no package.json / src/), prompt the user to scaffold code with their preferred stack tool (e.g., `npx create-next-app`, `go mod init`, etc.)
   - DoPlan no longer ships the legacy `scripts/boilerplate` helper, so projects must bring or generate their own starter code
   - Skip once code already exists for this project/stack

3. **Check Git Status**:
   - Verify working tree is clean (no uncommitted changes)
   - If dirty, warn user and block until clean

4. **Create/Checkout Task Branch**:
   - Create or checkout branch `task/[ID]` manually (e.g., `git checkout -b task/5.2`)
   - Store branch name in `active_branch` field of `.do/system/history/active_state.json`

5. **Load Task Context**: Read task details, dependencies, and related code

6. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)

7. **Start Implementation**: Begin coding the task with full context

8. **Update State**: Set `active_task` and `active_branch` in `.do/system/history/active_state.json`

9. **Snapshot State**: Immediately log the new state with `go run scripts/statehistory/main.go snapshot --reason "dev [ID]" --label dev`

10. **Automatic Completion Detection** (System continuously monitors):
   - Agent continuously analyzes code changes, tests, requirements
   - Agent checks if task criteria are met
   - **When System Detects Completion** (automatic, no user prompt needed):
     - **Verify Active Branch**: Check that we're on a task branch (from `active_branch` in `.do/system/history/active_state.json`)
     - **Check Dependencies**: Run `go run scripts/taskcomplete/main.go --task [ID] --project . --check` to verify all dependencies are complete
       - If dependencies are missing, **block completion** and list missing dependencies
     - **Mark Task Complete**: Run `go run scripts/taskcomplete/main.go --task [ID] --project .` to mark task complete in TASKS.md
       - Updates task status to "✅ Complete" and marks all checklist items as [x]
     - **Update State**: 
       - Add task ID to completed array in `.do/system/history/active_state.json`
       - Clear `active_task` and `active_branch` (set to null)
     - **Snapshot State**: Run `go run scripts/statehistory/main.go snapshot --reason "dev [ID] complete" --label dev` to record completion
     - **Auto-Commit**: Automatically commit changes with conventional commit format (e.g., `feat(task-5.2): complete [description]`)
     - **Auto-Push**: Run `go run scripts/branch/main.go --action push --project .` to push the task branch to origin
     - **Update Changelog**: If significant, add entry to CHANGELOG.md
     - **Optional PR Suggestion**: If `gh` CLI is available, suggest creating a PR
     - **Response**: "✅ Task [ID] complete! Changes committed and pushed to [branch_name]. Ready for next task. Type /dev to continue."
   - **If Not Complete**: Continue working on task, agent provides guidance

11. **Response**: "Building task [ID]: [Description] on branch [branch_name]. Focus on this task only. System will automatically detect when complete."

## Agent Involvement
- **Engineering Lead**
- **Relevant Team Leads**
- **Project Orchestrator**
- **QA Engineer**

## Files Read
- ".do/plan/TASKS.md"
- ".do/system/history/active_state.json"

## Files Modified
- ".do/system/history/active_state.json (active_task and active_branch updated, then cleared on completion)"
- ".do/system/history/state-*.json (automatic snapshot for audit/rollback)"
- ".do/plan/TASKS.md (task marked complete when system detects completion)"
- "CHANGELOG.md (if significant changes on completion)"
- "Git: New branch created/checked out (task/[ID]), then auto-commit and push on completion"
- "src/** (code files created/modified)"

## GitHub Automation
|

