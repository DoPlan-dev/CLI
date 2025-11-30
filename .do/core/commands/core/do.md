---
name: do
category: core
trigger: "/do or /do <idea>"
description: "Capture project idea or perform action"
agentInvolvement:
  - Project Orchestrator
  - Product Manager
filesRead: []
filesModified:
  - ".do/system/IDEA.md"
  - ".do/system/history/active_state.json"
examples:
  - "/do"
  - "/do Build a todo app"
---

When user types /do or /do <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .do/system/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /meeting to start the discovery meeting with the team."

## Agent Involvement
- **Project Orchestrator**: Analyzes the idea and determines project scope
- **Product Manager**: Begins thinking about requirements

## Files Modified
- .do/system/IDEA.md (created/updated)
- .do/system/history/active_state.json (phase: "idea")

