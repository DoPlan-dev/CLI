# Post-Launch Monitoring Playbook

Use this checklist immediately after a release to stay on top of support signals, adoption, and success metrics.

---

## 1. Monitor GitHub Issues
- **Cadence**: Daily during the first week, then 2–3× per week.
- **Commands**:
  - `gh issue list --label bug --limit 50`
  - `/github info` → refreshes `Docs/history/github-meta.json` with repo stats.
- **Workflow**:
  1. Label new issues (`bug`, `feedback`, `question`).
  2. If the issue came from `/feedback`, link it in the issue body to keep provenance.
  3. Update `Docs/history/github-meta.json` via `go run scripts/githubmeta/main.go` so `/report` can surface the latest counts.

## 2. Monitor npm Downloads
- **Cadence**: Weekly snapshot recorded in `Docs/history/metrics.md`.
- **Command**:
  ```bash
  npm view @doplan-dev/cli downloads --json
  ```
- **Workflow**:
  1. Record the daily/weekly totals in `Docs/history/metrics.md`.
  2. Compare against launch targets defined in `.plan/00_System/PRD.md`.
  3. If downloads dip >20% week-over-week, create a `/feedback note` item so growth tasks capture it.

## 3. Gather User Feedback
- **Primary tool**: `/feedback` command.
  - `Docs/history/feedback.json` → structured data consumed by `/report`.
  - `Docs/history/feedback.md` → human-readable log for quick triage.
- **Process**:
  1. After each support interaction or community message, log it via `/feedback`.
  2. When feedback implies a new task, open an issue with `/github issue` to keep Git + Docs in sync.
  3. Highlight top 3 themes in `Docs/release/post-launch-monitoring.md` (this file) each week so leadership gets a snapshot.

## 4. Track Success Metrics
- **Source of truth**: `.plan/00_System/PRD.md` → “Success Metrics” section.
- **Dashboard**: `/report` already embeds the latest `/progress` snapshot plus KPI deltas.
- **Workflow**:
  1. Define metric targets in `PRD.md`.
  2. Update `Docs/history/metrics.md` with actuals (GitHub stars, npm installs, feedback count, resolved issues).
  3. Run `/report --preset exec` for stakeholder updates; attach the generated `SCAN_DIFF_*.md` to weekly check-ins.

## 5. Plan v1.1 Features
- Gather inputs from the previous sections (issues, downloads, feedback, metrics).
- Capture candidates in `Docs/features/05_v1_1_planning/README.md` (create if missing).
- Once priorities are confirmed, use `/tell` + `/plan` to regenerate the backlog for the next phase.

---

**Quick Command Reference**
```
go run scripts/githubmeta/main.go --project .
go run scripts/progress/main.go --root .
go run scripts/scanreport/main.go --preset exec
/feedback <type> "<title>" "<details>"
/github issue "<title>" "<body>"
```

Store every artifact under `Docs/` so post-launch monitoring stays automated and auditable.

