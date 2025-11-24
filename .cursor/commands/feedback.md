# /feedback

## Trigger
`/feedback <type> "Title" "Details" [--github <url>] [--author <name>]`

### Examples
- `/feedback bug "QR download fails" "API returns 500 when Accept header missing"`
- `/feedback feature "Add dark mode" "Marketing wants dark hero section" --author PM`
- `/feedback question "Rate limit" "What are the prod limits?" --github https://github.com/org/repo/issues/123`

## Action
When you run `/feedback ...`:

1. **Parse arguments**
   - `type`: bug | feature | question | note (defaults to `note`)
   - `title`: short summary (required)
   - `details`: multiline description (optional)
   - `--author`: person filing feedback (defaults to `anonymous`)
   - `--github`: optional issue URL if mirrored upstream
2. **Log entry** via `go run scripts/feedback/main.go ...`
   - Appends markdown to `Docs/history/feedback.md`
   - Updates JSON log `Docs/history/feedback.json` for automation
3. **Surface in workflow**
   - `/report` command ingests latest feedback when generating scan metadata/diffs
   - Future scans can summarize outstanding feedback items
4. **Response**
   - “Feedback logged (type=bug) → Docs/history/feedback.md”

## Agent Involvement
- **Product Manager**: prioritizes feature/bug requests
- **QA Engineer**: triages bug feedback
- **Documentation Lead**: ensures items appear in scan reports

## Files Read
- Docs/history/feedback.md (created if missing)
- Docs/history/feedback.json

## Files Modified
- Docs/history/feedback.md
- Docs/history/feedback.json

## Notes
- Requires Go 1.21+. Command shells run: `go run scripts/feedback/main.go --type <type> --title "..." --details "..." --author "..." --github <url>`
- Works in any generated project (paths relative to project root).
- Add new feedback types by passing a custom string (stored as lowercase).
