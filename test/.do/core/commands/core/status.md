# /status

## Trigger
/status [<subcommand>]

## Examples
- "/status → Show current progress"
- "/status report → Generate scan report metadata and diffs"
- "/status full → Show progress and report together"


## Action
When user types /status or /status <subcommand>:

1. **If no subcommand provided** (default: show progress):
   - Read TASKS.md: Load all tasks
   - Read active_state.json: Get completed tasks and current phase
   - Run Progress Tool: Execute `go run scripts/progress/main.go --root <project>`
     This parses `.do/plan/TASKS.md`, `.do/system/history/active_state.json`, and `.do/system/history/` to compute stats and state deltas.
   - Calculate Progress:
     * Total tasks
     * Completed tasks
     * In progress tasks
     * Percentage complete
   - Display Progress: Show formatted progress report:
     * Phase: [current phase]
     * Tasks: X/Y completed (Z%)
     * Current task: [active task]
     * Next up: [next task]
     * State Delta: summarize what changed between the last two snapshots (phase/task/branch/completed)
   - Response: Display progress summary with the state delta footer

2. **Subcommand: report** (or /status report):
   - Select Project:
     * Default: current workspace (.)
     * Optional: `/status report test/qr-generator/test-no01`
   - Generate Metadata:
     * Runs `go run scripts/scanreport/main.go --project <path>`
     * Parses `.do/reports/SCAN_REPORT_*.md`
     * Creates/updates matching JSON files with structured metadata (scan date, project, executive summary, findings, next actions, summary hash)
   - Compute Diff:
     * When >=2 reports exist, compares the newest vs previous
     * Builds `SCAN_DIFF_<date>.md` highlighting added/removed bullets in Executive Summary, Findings & Risks, Recommended Next Actions, **and** the latest `.do/history` state changes (phase/task/branch/completed deltas)
     * Appends preset-specific sections: progress snapshot (from `scripts/progress`), ASCII visuals, and a dependency audit when manifests are detected
   - Output:
     * Terminal summary showing metadata count + diff file path
     * Diff markdown stored alongside the reports for sharing

3. **Subcommand: full** (or /status full):
   - Show both progress and report in one comprehensive view

## Agent Involvement
- **Project Orchestrator**
- **QA Engineer**
- **Documentation Lead**

## Files Read
- ".do/plan/TASKS.md"
- ".do/system/history/active_state.json"
- ".do/system/history/state-*.json"
- "<project>/.do/reports/SCAN_REPORT_*.md"

## Files Modified
- "<project>/.do/reports/SCAN_REPORT_*.json"
- "<project>/.do/reports/SCAN_DIFF_<date>.md"

## Requirements
|
## Options
|

