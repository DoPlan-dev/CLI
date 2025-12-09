---
name: plan
category: core
trigger: "/plan [--project <path>]"
description: "Generate execution plan from captured idea"
agentInvolvement:
  - Project Orchestrator
  - Engineering Lead
  - Product Manager
filesRead:
  - ".do/system/IDEA.md"
  - ".do/system/BRAINSTORM.md"
filesModified:
  - ".do/plan/TASKS.md"
  - ".do/system/history/active_state.json (phase/progress)"
requirements: |
  - Reads IDEA and BRAINSTORM to produce phased TASKS.md
  - Creates phase directories and feature folders as needed
  - Integrates with engagement/memory systems
examples:
  - "/plan"
  - "/plan --project ./my-project"
---

When user types /plan:

1) Load IDEA.md and BRAINSTORM.md context.
2) Generate TASKS.md with phases, tasks, and checklists; scaffold phase/feature folders if required.
3) Update active_state.json to reflect planning progress.
