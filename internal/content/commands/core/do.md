---
name: do
category: core
trigger: "/do [feature|now|i'm lucky|--prompt <text>|--prd <path>]"
description: "Capture idea, discovery meeting, and refinement"
agentInvolvement:
  - Product Manager
  - Engineering Lead
  - Design Manager
  - Project Orchestrator
filesRead:
  - ".do/system/IDEA.md (if exists)"
  - ".do/system/BRAINSTORM.md (if exists)"
filesModified:
  - ".do/system/IDEA.md"
  - ".do/system/BRAINSTORM.md"
  - ".do/system/REFINEMENTS.md"
  - ".do/system/history/active_state.json"
requirements: |
  - Interactive idea capture with optional fast-track (prompt/PRD)
  - Discovery meeting with speed selection and experience-aware questions
  - Refinement phase and confirmation before saving
examples:
  - "/do"
  - "/do feature"
  - "/do now --prompt \"Build a todo app\""
  - "/do now --prd ./requirements.md"
  - "/do i'm lucky"
---

When user types /do:

1) Start idea capture (interactive or fast-track via prompt/PRD).
2) Run discovery meeting with experience-aware speed options; adapt questions by project type.
3) Build summary and ask for explicit confirmation before saving.
4) Save IDEA.md, BRAINSTORM.md, and REFINEMENTS.md; update active_state.json phase.
