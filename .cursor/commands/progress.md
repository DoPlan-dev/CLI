# /progress

## Trigger
Exact match: /progress

## Action
When user types /progress:

1. **Read TASKS.md**: Load all tasks
2. **Read active_state.json**: Get completed tasks and current phase
3. **Run Progress Tool**: Execute  
   ```
   go run scripts/progress/main.go --root <project>
   ```  
   This parses `.plan/TASKS.md`, `.plan/active_state.json`, and `.plan/history/` to compute stats and state deltas.
4. **Calculate Progress**: 
   - Total tasks
   - Completed tasks
   - In progress tasks
   - Percentage complete
5. **Display Progress**: Show formatted progress report:
   - Phase: [current phase]
   - Tasks: X/Y completed (Z%)
   - Current task: [active task]
   - Next up: [next task]
   - State Delta: summarize what changed between the last two snapshots (phase/task/branch/completed)
6. **Response**: Display progress summary with the state delta footer

## Agent Involvement
- **Project Orchestrator**: Provides progress overview

## Files Read
- .plan/TASKS.md
- .plan/active_state.json
- .plan/history/state-*.json
