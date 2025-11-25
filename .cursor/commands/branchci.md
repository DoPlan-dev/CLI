# /branchci

## Trigger
`/branchci` or `/branchci regenerate`

## Action
When you run `/branchci`:

1. Reads `Docs/history/branch-matrix.json` to understand what jobs/required checks each branch prefix needs (e.g., `task/`, `feature/`, `hotfix/`).
2. Runs the generator:
   ```bash
   go run scripts/branchci/main.go --matrix Docs/history/branch-matrix.json --out .github/workflows
   ```
3. Emits `.github/workflows/task-branches.yml`, a workflow that:
   - Triggers on pushes to `task/*` (and can be expanded for other prefixes)
   - Spins up jobs per branch prefix (lint/test/build/etc.)
   - Adds a summary job so reviewers know which checks are required per branch
4. Output: “Workflow generated: .github/workflows/task-branches.yml”

## Customize
Edit `Docs/history/branch-matrix.json` to add or tweak prefixes, jobs, and required checks. Re-run `/branchci` after editing to regenerate the workflow.

## Files Read
- `Docs/history/branch-matrix.json`

## Files Modified
- `.github/workflows/task-branches.yml`

## Notes
- Generated workflow expects Go 1.21 and the standard lint/test/build jobs. Adapt `scripts/branchci/main.go` if your stack differs.
- Use this command whenever you add a new branch naming convention or need different CI steps per branch type.
