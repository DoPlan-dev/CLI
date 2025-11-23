# /finished

## Trigger
Exact match: /finished

## Action
When user types /finished:

1. **Mark Current Task Complete**: Mark the active task as completed in TASKS.md
2. **Update State**: Add task ID to completed array in .plan/active_state.json
3. **Clear Active Task**: Set active_task to null
4. **Auto-Commit**: Automatically commit changes with conventional commit format
5. **Auto-Push**: Push to current branch (feature/bugfix/hotfix)
6. **Update Changelog**: If significant, add entry to CHANGELOG.md
7. **Response**: "Task marked complete! Changes committed and pushed. Type /build to start the next task, or /progress to see overall progress."

## Agent Involvement
- **Engineering Lead**: Validates task completion
- **QA Engineer**: May review completed task
- **Release Captain**: Manages version control and changelog updates

## Files Modified
- .plan/TASKS.md (task marked [x])
- .plan/active_state.json (completed array updated, active_task cleared)
- CHANGELOG.md (if significant changes)
- Git: Auto-commit and push

## GitHub Automation
- Creates conventional commit message based on changes
- Pushes to current feature/bugfix/hotfix branch
- Follows branching strategy from @library/01-core-workflow/github-workflow-automation.md
- Triggers CI workflow on push
