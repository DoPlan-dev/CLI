---
name: dev
category: core
trigger: "/dev [<task_id>]"
description: "Start development and auto-complete tasks"
agentInvolvement:
  - Engineering Lead
  - Relevant Team Leads
  - QA Engineer
filesRead:
  - ".do/plan/TASKS.md"
  - ".do/system/history/active_state.json"
filesModified:
  - ".do/system/history/active_state.json (active_task, active_branch, progress)"
  - ".do/plan/TASKS.md (mark task complete when confirmed)"
requirements: |
  - Select next or specified task; prepare branch
  - Track active_task/branch and progress
  - Offer auto-complete when task criteria met
examples:
  - "/dev"
  - "/dev 2.1"
---

When user types /dev or /dev <task_id>:

1) Pick the specified task or next uncompleted task from TASKS.md.
2) Prompt/prepare branch (task/[ID]) and update active_state.json with active_task and active_branch.
3) Guide implementation; track progress.
4) When work is done, confirm completion, mark task in TASKS.md, update active_state.json, and auto-commit/push if enabled.
