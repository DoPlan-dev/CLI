# /finished

## Trigger
/finished

## Examples
- /finished


## Action
When user types /finished:

1. **Verify Dependencies**: Parse the active task block in `.plan/TASKS.md` and ensure every ID listed under `**Dependencies**` is already marked `[x]`. If any dependency remains `[ ]`, stop immediately and respond: `"Cannot finish TASK-XXX. Dependency TASK-YYY is still open."`
2. **Mark Task Complete**: Within `.plan/TASKS.md`, locate the checklist items for the active task and flip each `- [ ]` to `- [x]`. (If the task has subtasks, only mark the ones truly finished.)
3. **Update State**: Remove `active_task`, clear `active_branch`, and add the task ID to `completed` in `.plan/active_state.json`
4. **Snapshot State**: `go run scripts/statehistory/main.go snapshot --reason "finished [ID]" --label finished` so rollback + scan reports can reference the change
5. **Run Checks**: Execute lint/tests per task requirements before committing
6. **Auto-commit**: Commit changes on the recorded `task/<TASK-ID>-<slug>` branch with conventional commit format
7. **Auto-push**: Push that branch (`git push -u origin branch`)
8. **Optional PR**: Open a draft/PR if workflow requires, referencing the task ID
9. **Update CHANGELOG**: Update CHANGELOG.md if significant changes
10. **Response**: "Task completed on [branch]! Changes committed and pushed. Type /build to start next task."

## Agent Involvement
- **Project Orchestrator**
- **Release Captain**

## Files Read
- .plan/TASKS.md
- .plan/active_state.json

## Files Modified
- .plan/TASKS.md
- .plan/active_state.json (active_task + active_branch)
- .plan/history/state-*.json (new snapshot)
- CHANGELOG.md

## GitHub Automation
After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push the task branch (`task/<TASK-ID>-<slug>`)
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md
