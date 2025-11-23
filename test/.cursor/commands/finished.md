# /finished

## Trigger
/finished

## Examples
- /finished


## Action
When user types /finished:

1. **Mark Task Complete**: Update task status in TASKS.md
2. **Update State**: Remove active_task and add to completed in active_state.json
3. **Auto-commit**: Commit changes with conventional commit format
4. **Auto-push**: Push to current branch
5. **Update CHANGELOG**: Update CHANGELOG.md if significant changes
6. **Response**: "Task completed! Changes committed and pushed. Type /build to start next task."

## Agent Involvement
- **Project Orchestrator**
- **Release Captain**

## Files Read
- .plan/TASKS.md
- .plan/active_state.json

## Files Modified
- .plan/TASKS.md
- .plan/active_state.json
- CHANGELOG.md

## GitHub Automation
After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md
