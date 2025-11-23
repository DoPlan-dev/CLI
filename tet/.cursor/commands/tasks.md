# /tasks

## Trigger
/tasks

## Examples
- /tasks


## Action
When user types /tasks:

1. **Read Approved Plan**: Load PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
2. **Generate Tasks**: Create implementation tasks organized by phases
3. **Create TASKS.md**: Write tasks to .plan/TASKS.md
4. **Update State**: Set phase to "tasks" in active_state.json
5. **Response**: "Tasks generated! Review .plan/TASKS.md. Type /build to start coding."

## Agent Involvement
- **Project Orchestrator**
- **Engineering Lead**
- **Product Manager**

## Files Read
- .plan/00_System/PRD.md
- .plan/00_System/ARCHITECTURE.md
- .plan/00_System/DESIGN_SYSTEM.md
- .plan/active_state.json

## Files Modified
- .plan/TASKS.md
- .plan/active_state.json

