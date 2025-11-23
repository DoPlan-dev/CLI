# /load

## Trigger
Exact match: /load

## Action
When user types /load:

1. **Read Current Context**: Read PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md, and TASKS.md
2. **Inject into Agents**: Update all agent prompts with current project context
3. **Update Agent Files**: Modify .cursor/agents/*.md files to include:
   - Current project goals
   - Technical constraints
   - Design requirements
   - Active tasks
4. **Response**: "Context loaded! All agents now have full project context. Type /build to start coding."

## Agent Involvement
- **All Agents**: Receive updated context
- **Project Orchestrator**: Coordinates context injection

## Files Modified
- .cursor/agents/*.md (all agent files updated with context)
