# /github

## Trigger
/github <subcommand>

## Examples
- "/github info"
- "/github issue \"Fix cache\" \"Cache misses spike\""
- "/github ci → Generate CI workflow"
- "/github release → Release management"


## Action
When user types /github <subcommand>:

### `/github info`
Runs:
```
go run scripts/githubmeta/main.go --project . --sync-readme
```
- Detects primary remote + default branch
- Extracts success metrics from `.do/system/PRD.md`
- Updates the README KPI block between `<!-- KPIS:START -->` / `<!-- KPIS:END -->`
- Persists metadata to `docs/history/github-meta.json` for offline use

### `/github issue "Title" "Body"`
Outputs a ready-to-run `gh issue create` command with the detected repo slug, e.g.:
```
go run scripts/githubmeta/main.go --project . --issue-title "Fix cache" --issue-body "Details here"
```
Copy/paste the printed `gh issue create` command (or pipe it) to open the issue.

### `/github milestone "Name" [due-date]`
Prints a `gh api` command to create a milestone:
```
go run scripts/githubmeta/main.go --project . --milestone-title "MVP" --milestone-due 2025-01-15T00:00:00Z
```

### `/github ci [regenerate]`
Generates CI workflow for branch prefixes:
1. Reads `docs/history/branch-matrix.json` to understand what jobs/required checks each branch prefix needs (e.g., `task/`, `feature/`, `hotfix/`).
2. Runs the generator:
   ```bash
   go run scripts/branchci/main.go --matrix docs/history/branch-matrix.json --out .github/workflows
   ```
3. Emits `.github/workflows/task-branches.yml`, a workflow that:
   - Triggers on pushes to `task/*` (and can be expanded for other prefixes)
   - Spins up jobs per branch prefix (lint/test/build/etc.)
   - Adds a summary job so reviewers know which checks are required per branch
4. Output: "Workflow generated: .github/workflows/task-branches.yml"

### `/github release`
Release management:
1. **Release Planning**: Release Captain plans the release
2. **Version Management**: Manage version numbers and semantic versioning
3. **Release Notes**: Generate release notes
4. **Deployment Planning**: Plan deployment strategy
5. **Response**: "Release planned! Review release notes and deployment plan."

## Agent Involvement
- **Release & Growth Manager**
- **Release Captain**
- **DevOps Engineer**

## Files Read
- "`.git/` metadata"
- "`.do/system/PRD.md`"
- "docs/history/branch-matrix.json"
- ".do/plan/TASKS.md"
- "docs/history/CHANGELOG.md"

## Files Modified
- "`docs/history/github-meta.json`"
- "`README.md` KPI section when `--sync-readme` is used"
- ".github/workflows/task-branches.yml"
- "docs/history/CHANGELOG.md"
- ".do/system/RELEASE.md"

## Notes
|
## Customize
Edit `docs/history/branch-matrix.json` to add or tweak prefixes, jobs, and required checks. Re-run `/github ci` after editing to regenerate the workflow.
## Offline Safety
If git remote detection fails, the script logs a warning and keeps the last cached metadata (`docs/history/github-meta.json`). You can still update KPIs from PRD without network access.

