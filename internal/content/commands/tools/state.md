---
name: state
category: tools
trigger: "/state <subcommand>"
description: "Manage project state history"
agentInvolvement:
  - Project Orchestrator
  - QA Engineer
filesRead:
  - "`.do/active_state.json`"
  - "`.do/history/state-*.json`"
filesModified:
  - "`.do/history/state-*.json` (new entries)"
  - "`.do/active_state.json` (when restoring)"
notes: |
  - State history is now required before/after `/build` and `/finished`
  - Restores should be rare; always snapshot first so you can undo mistakes
examples:
  - "/state snapshot --reason \"after /build 5.8\""
  - "/state list --limit 5"
  - "/state diff --json"
  - "/state restore --file state-20251124T120000Z.json --yes"
---

The `/state` helper wraps `go run scripts/statehistory/main.go` so you can manage `.do/active_state.json` history safely.

### snapshot
1. Writes the current `.do/active_state.json` into `.do/history/state-<timestamp>.json`
2. Accepts optional flags:
   - `--reason` → stored in the snapshot metadata
   - `--label` → appended to the file name (e.g., build, finished)
3. Output: `Snapshot saved: .do/history/state-20251124T120000Z-build.json`

### list
1. Lists recent entries (default: last 10)
2. `--json` emits machine-readable summaries for scripts/CI

### diff
1. Compares two snapshots (default: latest vs previous)
2. Shows Markdown summary (phase/task/branch/completed deltas) or JSON if `--json`
3. Used by `/progress` and `/report` to surface state deltas

### restore
1. Requires `--file <id>` and `--yes` confirmation for guardrails
2. Restores `.do/active_state.json` from the selected snapshot
3. Optionally captures a new snapshot (`--snapshot=false` to skip) so rollbacks themselves are logged
4. Respond with confirmation + reminder to rerun `/progress`

