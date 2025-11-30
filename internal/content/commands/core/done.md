---
name: done
category: core
trigger: "/done"
description: "Mark current task as complete and auto-commit/push"
agentInvolvement:
  - Engineering Lead
  - QA Engineer
  - Release Captain
filesRead:
  - ".do/plan/TASKS.md"
  - ".do/system/history/active_state.json"
filesModified:
  - ".do/plan/TASKS.md (task marked [x] and status updated)"
  - ".do/system/history/active_state.json (completed array updated, active_task and active_branch cleared)"
  - ".do/system/history/state-*.json (new snapshot)"
  - "CHANGELOG.md (if significant changes)"
  - "Git: Auto-commit and push to task branch"
githubAutomation: |
  After task completion, the system will:
  - Auto-commit changes with conventional commit format (e.g., `feat(task-5.2): complete branch automation`)
  - Auto-push to current branch using scripts/branch/main.go
  - Update CHANGELOG.md if significant changes
  - Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md
  - Trigger CI workflow on push
  - Optional: Suggest PR creation with `gh pr create` if gh CLI is available
examples:
  - "/done → Mark current task complete"
---
# /done

## Trigger
Exact match: /done

## Action
When user types /done:

1. **Verify Active Branch**: 
   - Check that we're on a task branch (from `active_branch` in `.do/system/history/active_state.json`)
   - If on main/master, warn and ask for confirmation
2. **Check Dependencies**: 
   - Run `go run scripts/taskcomplete/main.go --task [ID] --project . --check` to verify all dependencies are complete
   - If dependencies are missing, **block completion** and list missing dependencies
3. **Mark Current Task Complete**: 
   - Run `go run scripts/taskcomplete/main.go --task [ID] --project .` to mark task complete in TASKS.md
   - This updates the task status to "✅ Complete" and marks all checklist items as [x]
4. **Update State**: 
   - Add task ID to completed array in `.do/system/history/active_state.json`
   - Clear `active_task` and `active_branch` (set to null)
5. **Snapshot State**: Run `go run scripts/statehistory/main.go snapshot --reason "done [ID]" --label done` so the new status is recorded in `.do/system/history/`
6. **Auto-Commit**: Automatically commit changes with conventional commit format (e.g., `feat(task-5.2): complete branch automation`)
7. **Auto-Push**: 
   - Run `go run scripts/branch/main.go --action push --project .` to push the current branch
   - This pushes the task branch to origin
8. **Update Changelog**: If significant, add entry to CHANGELOG.md
9. **Optional PR Creation**: 
   - If `gh` CLI is available, suggest creating a PR with: `gh pr create --title "feat: [task description]" --body "Completes task [ID]"`
   - This is optional and can be done manually
10. **Response**: "Task marked complete! Changes committed and pushed to [branch_name]. Type /dev to start the next task, or /status to see overall progress."

## Agent Involvement
- **Engineering Lead**: Validates task completion
- **QA Engineer**: May review completed task
- **Release Captain**: Manages version control and changelog updates

## Files Modified
- .do/plan/TASKS.md (task marked [x])
- .do/system/history/active_state.json (completed array updated, active_task and active_branch cleared)
- .do/system/history/state-*.json (new snapshot)
- CHANGELOG.md (if significant changes)
- Git: Auto-commit and push to task branch

## GitHub Automation
- Creates conventional commit message based on changes
- Pushes to current feature/bugfix/hotfix branch
- Follows branching strategy from @library/01-core-workflow/github-workflow-automation.md
- Triggers CI workflow on push

## Status
✅ **IMPLEMENTED**: Command fully implemented and ready for use.

