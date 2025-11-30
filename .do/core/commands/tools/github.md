# /github

## Trigger
`/github info` or `/github issue ...`

## Actions

### `/github info`
Runs:
```
go run scripts/githubmeta/main.go --project . --sync-readme
```
- Detects primary remote + default branch
- Extracts success metrics from `.plan/00_System/PRD.md`
- Updates the README KPI block between `<!-- KPIS:START -->` / `<!-- KPIS:END -->`
- Persists metadata to `Docs/history/github-meta.json` for offline use

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

## Files Read
- `.git/` metadata
- `.plan/00_System/PRD.md`

## Files Modified
- `Docs/history/github-meta.json`
- `README.md` KPI section when `--sync-readme` is used

## Offline Safety
- If git remote detection fails, the script logs a warning and keeps the last cached metadata (`Docs/history/github-meta.json`). You can still update KPIs from PRD without network access.
