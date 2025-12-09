---
name: sys
category: core
trigger: "/sys [engagement|role|security|control]"
description: "System control panel and subcommands"
agentInvolvement:
  - Project Orchestrator
  - Security Lead
  - Engineering Lead
filesRead:
  - ".do/system/memory_card.json (achievements/streaks)"
  - ".do/system/history/active_state.json"
filesModified:
  - ".do/system/memory_card.json (when engagement actions occur)"
  - ".do/system/history/active_state.json (settings/flags)"
requirements: |
  - Provide overview and route to subcommands
  - Subcommands: engagement, role, security, control
examples:
  - "/sys"
  - "/sys engagement"
  - "/sys role"
  - "/sys security"
  - "/sys control"
---

Subcommands:
- **/sys engagement**: Show achievements, XP, streaks, challenges.
- **/sys role**: Manage roles/permissions.
- **/sys security**: Security settings and checks.
- **/sys control**: Toggle system features on/off.
