# /finished

## Trigger
Exact match: /finished

## Action
When user types /finished:

1. **Verify Active Branch**: 
   - Check that we're on a task branch (from `active_branch` in `.plan/active_state.json`)
   - If on main/master, warn and ask for confirmation
2. **Check Dependencies**: 
   - Run `go run scripts/taskcomplete/main.go --task [ID] --project . --check` to verify all dependencies are complete
   - If dependencies are missing, **block completion** and list missing dependencies
3. **Mark Current Task Complete**: 
   - Run `go run scripts/taskcomplete/main.go --task [ID] --project .` to mark task complete in TASKS.md
   - This updates the task status to "✅ Complete" and marks all checklist items as [x]
4. **Update State**: 
   - Add task ID to completed array in `.plan/active_state.json`
   - Clear `active_task` and `active_branch` (set to null)
4. **Snapshot State**: Run `go run scripts/statehistory/main.go snapshot --reason "finished [ID]" --label finished` so the new status is recorded in `.plan/history/`
5. **Auto-Commit**: Automatically commit changes with conventional commit format (e.g., `feat(task-5.2): complete branch automation`)
6. **Auto-Push**: 
   - Run `go run scripts/branch/main.go --action push --project .` to push the current branch
   - This pushes the task branch to origin
7. **Update Changelog**: If significant, add entry to CHANGELOG.md
8. **Optional PR Creation**: 
   - If `gh` CLI is available, suggest creating a PR with: `gh pr create --title "feat: [task description]" --body "Completes task [ID]"`
   - This is optional and can be done manually
9. **Response**: "Task marked complete! Changes committed and pushed to [branch_name]. Type /build to start the next task, or /progress to see overall progress."

## Agent Involvement
- **Engineering Lead**: Validates task completion
- **QA Engineer**: May review completed task
- **Release Captain**: Manages version control and changelog updates

## Files Modified
- .plan/TASKS.md (task marked [x])
- .plan/active_state.json (completed array updated, active_task and active_branch cleared)
- .plan/history/state-*.json (new snapshot)
- CHANGELOG.md (if significant changes)
- Git: Auto-commit and push to task branch

## GitHub Automation
- Creates conventional commit message based on changes
- Pushes to current feature/bugfix/hotfix branch
- Follows branching strategy from @library/01-core-workflow/github-workflow-automation.md
- Triggers CI workflow on push
