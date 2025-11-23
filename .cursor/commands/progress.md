# /progress

## Trigger
Exact match: /progress

## Action
When user types /progress:

1. **Read TASKS.md**: Load all tasks
2. **Read active_state.json**: Get completed tasks and current phase
3. **Calculate Progress**: 
   - Total tasks
   - Completed tasks
   - In progress tasks
   - Percentage complete
4. **Display Progress**: Show formatted progress report:
   - Phase: [current phase]
   - Tasks: X/Y completed (Z%)
   - Current task: [active task]
   - Next up: [next task]
5. **Response**: Display progress summary

## Agent Involvement
- **Project Orchestrator**: Provides progress overview

## Files Read
- .plan/TASKS.md
- .plan/active_state.json
