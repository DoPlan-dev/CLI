# /sys

## Trigger
/sys [<subcommand>]

## Examples
- "/sys → Show system control panel"
- "/sys status → Show project progress and generate reports"
- "/sys performance → View performance metrics"
- "/sys backup → Create project backup"
- "/sys restore → Restore from backup"
- "/sys memory → Export/import memory card"
- "/sys engagement → View engagement dashboard"
- "/sys optimize → Project optimization hub"
- "/sys state → Manage project state history"
- "/sys feedback → Log feedback"
- "/sys github → GitHub operations"
- "/sys security → Security review and audit"
- "/sys permissions → Design RBAC system"
- "/sys access → Fix permissions"


## Action
When user types /sys or /sys <subcommand>:

1. **If no subcommand provided**: Show menu:
   - "System Control Panel:"
   - "1. status - Show project progress and generate reports"
   - "2. engagement - View engagement dashboard"
   - "3. performance - View performance metrics"
   - "4. optimize - Project optimization hub"
   - "5. backup - Create project backups"
   - "6. restore - Restore from backup"
   - "7. memory - Export/import memory card"
   - "8. state - Manage project state history"
   - "9. feedback - Log feedback"
   - "10. github - GitHub operations"
   - "11. security - Security review and audit"
   - "12. permissions - Design RBAC system"
   - "13. access - Fix permissions"
   - Wait for user selection

2. **Subcommand: status** (or /sys status):
   - Show project progress and generate reports
   - Default: Show current progress (tasks completed, percentage, current task, next task, state delta)
   - Subcommands:
     * /sys status report → Generate scan report metadata and diffs
     * /sys status full → Show both progress and report in one comprehensive view
   - Response: "Progress displayed!" or "Report generated!"

3. **Subcommand: engagement** (or /sys engagement):
   - View engagement dashboard and statistics
   - Show achievements, challenges, score tracking
   - Response: "Engagement dashboard displayed!"

4. **Subcommand: performance** (or /sys performance):
   - View performance metrics and cache statistics
   - Show command execution metrics, cache hit rates
   - Response: "Performance metrics displayed!"

5. **Subcommand: optimize** (or /sys optimize):
   - Project optimization hub
   - If no subcommand: Show menu:
     * "1. Design - UI/UX improvements"
     * "2. Finance - Cost optimization"
     * "3. Performance - Performance optimization"
     * "4. All - Run all optimizations"
   - Subcommands:
     * /sys optimize design → UI/UX improvements
     * /sys optimize finance → Cost optimization
     * /sys optimize performance → Performance optimization
     * /sys optimize all → Run all optimizations
   - Response: "Optimization complete! Review recommendations in [report].md"

6. **Subcommand: backup** (or /sys backup):
   - Create compressed project backups
   - Multiple backup types: project, plan, project-plan, full
   - Response: "Backup created!"

7. **Subcommand: restore** (or /sys restore):
   - Restore project from backup
   - Dry-run, safety backups, version compatibility checks
   - Response: "Project restored from backup!"

8. **Subcommand: memory** (or /sys memory):
   - Export/import memory card
   - Manage personalization data
   - Response: "Memory card exported/imported!"

9. **Subcommand: state** (or /sys state):
   - Manage project state history
   - Wraps `go run scripts/statehistory/main.go` for safe state management
   - Subcommands:
     * /sys state snapshot [--reason "text"] [--label "label"] → Save current state snapshot
     * /sys state list [--limit N] [--json] → List recent state snapshots
     * /sys state diff [--json] → Compare snapshots (default: latest vs previous)
     * /sys state restore --file <id> --yes → Restore state from snapshot
   - Response: "State operation completed!"

10. **Subcommand: feedback** (or /sys feedback):
   - Log feedback (bug, feature, question, note)
   - Usage: /sys feedback <type> "Title" "Details" [--github <url>] [--author <name>]
   - Types: bug | feature | question | note (defaults to note)
   - Logs to `docs/history/feedback.md` and `docs/history/feedback.json`
   - Response: "Feedback logged (type=[type]) → docs/history/feedback.md"

11. **Subcommand: github** (or /sys github):
   - GitHub operations
   - Subcommands:
     * /sys github info → Sync README with project KPIs, update repository metadata
     * /sys github issue "Title" "Body" → Generate gh issue create command
     * /sys github milestone "Name" [due-date] → Generate gh api command for milestone
     * /sys github ci [regenerate] → Generate CI/CD workflows for branch prefixes
     * /sys github release → Release management (planning, versioning, release notes)
   - Response: "GitHub operation completed!"

12. **Subcommand: security** (or /sys security):
   - Security review and audit
   - If no subcommand: Show menu (review, audit, both)
   - Subcommands:
     * /sys security review → Security review, vulnerability assessment, generate report
     * /sys security audit → Comprehensive security audit, vulnerability scanning, compliance check
     * /sys security both → Run both review and audit
   - Response: "Security review/audit complete! Review findings in SECURITY.md"

13. **Subcommand: permissions** (or /sys permissions):
   - Design RBAC system (role-based access control)
   - Security Lead and Backend Lead design role-based access control
   - Define roles and permissions
   - Generate documentation
   - Response: "RBAC system designed! Review role definitions in RBAC.md"

14. **Subcommand: access** (or /sys access):
   - Fix permissions for .do/ and docs/ directories
   - Execute `npx --yes @doplan-dev/cli goplan access <scope>` (defaults to all)
   - Scope options: all, .do/system, .do/plan, docs
   - Creates missing folders/files and fixes permissions
   - Response: "Permissions fixed! You can now run /hey again."

## Agent Involvement
- **Project Orchestrator**
- **Product Manager**
- **QA Engineer**
- **Documentation Lead**
- **Release & Growth Manager**
- **Release Captain**
- **DevOps Engineer**
- **Security Lead**
- **Backend Lead**
- **Design & UX Manager**
- **UI/UX Designer**
- **Performance Engineer**

## Files Read
- ".do/plan/TASKS.md"
- ".do/system/history/active_state.json"
- ".do/system/history/state-*.json"
- ".do/reports/SCAN_REPORT_*.md"
- "docs/history/feedback.md"
- "docs/history/feedback.json"
- ".git/"
- ".do/system/PRD.md"
- "docs/history/branch-matrix.json"
- "docs/history/CHANGELOG.md"
- "src/**"
- ".do/system/ARCHITECTURE.md"
- ".do/system/DESIGN_SYSTEM.md"
- ".do/system/SECURITY.md"

## Files Modified
- ".do/reports/SCAN_REPORT_*.json"
- ".do/reports/SCAN_DIFF_<date>.md"
- ".do/history/state-*.json"
- ".do/active_state.json"
- "docs/history/feedback.md"
- "docs/history/feedback.json"
- "docs/history/github-meta.json"
- "README.md"
- ".github/workflows/task-branches.yml"
- ".do/system/RELEASE.md"
- ".do/system/SECURITY.md"
- ".do/system/SECURITY_AUDIT.md"
- ".do/system/RBAC.md"
- ".do/system/DESIGN_SYSTEM.md"
- ".do/system/COST_OPTIMIZATION.md"
- ".do/system/PERFORMANCE_OPTIMIZATION.md"
- ".do/system/**"
- ".do/plan/**"
- "docs/**"


