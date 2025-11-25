# /good

## Trigger
Exact match: /good

## Action
When user types /good:

1. **Approve Plan**: Mark the current plan (PRD, Architecture, Design System) as approved and locked.
2. **Update State**: Set locked: true in .plan/active_state.json
3. **Lock Documents**: Mark PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md as approved (no more structural changes allowed).
4. **Response**: "Plan approved and locked! Type /plan to generate the execution plan and task hierarchy."

## Agent Involvement
- **Project Orchestrator**: Approves and locks the plan

## Files Modified
- .plan/active_state.json (locked: true, phase: "approved")
