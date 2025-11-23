# /improve

## Trigger
Exact match: /improve

## Action
When user types /improve:

1. **Team Brainstorm**: Activate Product Manager, Engineering Lead, and Design Manager for interactive brainstorming session.
2. **Ask Questions**: The team asks clarifying questions about:
   - Target users
   - Core features
   - Technical constraints
   - Design preferences
   - Timeline expectations
3. **Save to BRAINSTORM.md**: Document all brainstorming insights to .plan/00_System/BRAINSTORM.md
4. **Response**: "Brainstorming session complete! Type /write to generate PRD, Architecture, and Design System."

## Agent Involvement
- **Product Manager**: Leads the brainstorming, asks product questions
- **Engineering Lead**: Asks technical questions
- **Design Manager**: Asks design and UX questions
- **Project Orchestrator**: Coordinates the session

## Files Modified
- .plan/00_System/BRAINSTORM.md (created/updated)
- .plan/active_state.json (phase: "brainstorm")
