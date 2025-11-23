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
4. **Start Implementation**: Begin coding the task with full context
5. **Update State**: Set active_task in .plan/active_state.json
6. **Response**: "Building task [ID]: [Description]. Focus on this task only."

## Agent Involvement
- **Engineering Lead**
- **Project Orchestrator**

## Files Read
- .plan/TASKS.md
- .plan/active_state.json

## Files Modified
- .plan/active_state.json

## GitHub Automation
After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md
