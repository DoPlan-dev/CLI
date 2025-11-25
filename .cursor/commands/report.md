# /report

## Trigger
/report or /report <project_path>

## Action
When you type /report:

1. **Select Project**:
   - Default: current workspace (.)
   - Optional: `/report test/qr-generator/test-no01`
2. **Generate Metadata**:
   - Runs `go run scripts/scanreport/main.go --project <path>`
   - Parses `.plan/reports/SCAN_REPORT_*.md`
   - Creates/updates matching JSON files with structured metadata (scan date, project, executive summary, findings, next actions, summary hash)
3. **Compute Diff**:
   - When >=2 reports exist, compares the newest vs previous
   - Builds `SCAN_DIFF_<date>.md` highlighting added/removed bullets in Executive Summary, Findings & Risks, Recommended Next Actions, **and** the latest `.plan/history` state changes (phase/task/branch/completed deltas)
   - Appends preset-specific sections: progress snapshot (from `scripts/progress`), ASCII visuals, and a dependency audit when manifests are detected
4. **Output**:
   - Terminal summary showing metadata count + diff file path
   - Diff markdown stored alongside the reports for sharing

## Options
- `--preset standard` *(default)* – complete report
- `--preset exec` – condensed executive view + visuals
- `--preset detailed` – expanded sections with dependency audit
- `.plan/reports/config.json` (optional) can set:
  ```json
  {
    "preset": "exec",
    "sections": ["executive", "progress", "visuals", "state", "feedback"]
  }
  ```
  CLI flags override config; custom `sections` let teams reorder or omit report blocks.

## Agent Involvement
- **QA Engineer**: Validates scan data
- **Documentation Lead**: Reviews diff output

## Files Read
- `<project>/.plan/reports/SCAN_REPORT_*.md`
- `<project>/.plan/history/state-*.json` (for the state delta section)

## Files Modified
- `<project>/.plan/reports/SCAN_REPORT_*.json`
- `<project>/.plan/reports/SCAN_DIFF_<date>.md`

## Requirements
- Go 1.21+
- Reports must follow `SCAN_REPORT_YYYY-MM-DD.md` naming

## Examples
- `/report` → run against current repo
- `/report test/qr-generator/test-no01` → run inside test fixture
