# /good

## Trigger
/good

## Examples
- /good


## Action
When user types /good:

1. **Validate Documents**: Ensure PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md exist
2. **Lock Plan**: Set locked: true in .plan/active_state.json
3. **Update Phase**: Set phase to "approved" in active_state.json
4. **Response**: "Plan approved and locked! Type /tasks to generate implementation tasks."

## Agent Involvement
- **Project Orchestrator**

## Files Read
- .plan/00_System/PRD.md
- .plan/00_System/ARCHITECTURE.md
- .plan/00_System/DESIGN_SYSTEM.md
- .plan/active_state.json

## Files Modified
- .plan/active_state.json

