# Docs Directory

DoPlan keeps every long-form artifact inside `docs/` so the repository root stays clean. Use these canonical subfolders:

| Directory | Purpose |
| --- | --- |
| `docs/overview/` | Project overview, mission, agent hierarchy |
| `docs/references/` | Command cheat sheets and reusable examples |
| `docs/tutorials/` | Guided walkthroughs and training notes |
| `docs/history/` | Changelog, meeting notes, audit logs |
| `docs/ops/` | Release checklists, security reviews, CI/CD docs |
| `docs/research/` | Brainstorms, interviews, experiment logs |

## Rules

1. **All docs live under `docs/`.** Keep the repository root clean.
2. **Folder = contract.** When you add a new category, document it here so agents know where to write.
3. **History lives in `docs/history/`.** Capture retrospectives, decision logs, and state snapshots there.
4. **Link, don't duplicate.** Reference existing docs rather than copying content across folders.

Need another category? Create it under `docs/` and update this README so future scans stay consistent.
