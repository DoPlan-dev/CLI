# /tasks

## Trigger
Exact match: /tasks

## Action
When user types /tasks:

1. **Read Approved Documents**: Engineering Lead reads PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
2. **Break Down Tasks**: Convert approved plan into granular, actionable tasks
3. **Generate TASKS.md**: Create comprehensive task list with:
   - Task IDs (1.1, 1.2, 2.1, etc.)
   - Task descriptions
   - Dependencies
   - Estimated effort
   - Assigned agents/teams
4. **Save to TASKS.md**: Write to .plan/TASKS.md
5. **Response**: "Tasks generated! Review TASKS.md. Type /load to inject context, then /build to start coding."

## Agent Involvement
- **Engineering Lead**: Creates the task breakdown
- **System Architect**: Provides technical task details
- **Product Manager**: Validates task alignment with PRD
- **Project Orchestrator**: Reviews and approves task list

## Files Created
- .plan/TASKS.md
- .plan/active_state.json (phase: "tasks")
