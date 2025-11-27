---
name: tell
category: core
trigger: "/tell or /tell <idea>"
description: "Capture project idea"
agentInvolvement:
  - Project Orchestrator
  - Product Manager
filesRead: []
filesModified:
  - ".do/system/IDEA.md"
  - ".do/system/history/active_state.json"
examples:
  - "/tell"
  - "/tell Build a todo app"
---

When user types /tell or /tell <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .do/system/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /meeting to start the discovery meeting with the team."

