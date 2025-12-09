---
name: hey
category: core
trigger: "/hey [--project <path>]"
description: "Onboarding, tutorial, and welcome experience"
agentInvolvement:
  - Project Orchestrator
  - Product Manager
filesRead:
  - ".do/system/user_profile.json (if exists)"
  - ".do/system/memory_card.json (if exists)"
filesModified:
  - ".do/system/user_profile.json (created/updated)"
  - ".do/system/QUICK_REFERENCE.md"
  - "docs/references/QUICK_REFERENCE.md"
  - "docs/overview/AGENT_HIERARCHY.md"
requirements: |
  - Detect first-time vs returning user
  - Show agent hierarchy and quick reference
  - Offer interactive tutorial and test-drive mode
examples:
  - "/hey"
  - "/hey --project ./path"
---

When user types /hey:

1) Detect first-time vs returning from user_profile.json and memory_card.
2) For new users: warm greeting, ask name/experience, show system overview, agent hierarchy, tutorial option, create quick reference files.
3) For returning users: personalized greeting, quick command reference, option to re-run tutorial, show current progress summary.
4) Save/update user profile and references. Keep responses fast for new projects (lightweight path if no memory card yet).
