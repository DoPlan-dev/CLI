# /tell

## Trigger
Exact match: /tell or /tell <idea>

## Action
When user types /tell or /tell <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .plan/00_System/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /improve to brainstorm with the team."

## Agent Involvement
- **Project Orchestrator**: Analyzes the idea and determines project scope
- **Product Manager**: Begins thinking about requirements

## Files Modified
- .plan/00_System/IDEA.md (created/updated)
- .plan/active_state.json (phase: "idea")
