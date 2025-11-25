package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Command represents a slash command in the AI agency system
type Command struct {
	// Basic Information
	Name        string // Command name (e.g., "tell", "build")
	Trigger     string // Trigger pattern (e.g., "/tell or /tell <idea>")
	Description string // Brief description
	Action      string // Detailed action description

	// Agent Involvement
	AgentInvolvement []string // List of agents involved

	// Files
	FilesRead     []string // Files read by this command
	FilesModified []string // Files modified by this command

	// Additional Information
	Examples         []string // Example usage
	GitHubAutomation string   // GitHub automation details (if applicable)
	Requirements     string   // Requirements section (if applicable)
	Notes            string   // Notes section (if applicable)
	Customize        string   // Customize section (if applicable)
	Options          string   // Options section (if applicable)
	OfflineSafety    string   // Offline Safety section (if applicable)
}

// GetAllCommands returns all core and squad commands
func GetAllCommands() []Command {
	return []Command{
		// Core Commands (11)
		{
			Name:        "tell",
			Trigger:     "/tell or /tell <idea>",
			Description: "Capture project idea",
			Action: `When user types /tell or /tell <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .plan/00_System/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /improve to brainstorm with the team."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
				"Product Manager",
			},
			FilesRead: []string{},
			FilesModified: []string{
				".plan/00_System/IDEA.md",
				".plan/active_state.json",
			},
			Examples: []string{
				"/tell",
				"/tell Build a todo app",
			},
		},
		{
			Name:        "improve",
			Trigger:     "/improve",
			Description: "Phase-based discovery interview & brainstorming",
			Action: `When user types /improve:

1. **Load Phase Templates**: Read interview phases from .plan/templates/brainstorm/phase-*.md (01-06) to enable customization without code changes
2. **Conduct Phase-by-Phase Interview**:
   - Phase 01: Vision & Outcomes (Product Manager leads)
   - Phase 02: Audience & Differentiation (Product Manager + Design Manager)
   - Phase 03: Experience, UI/UX & Tech (Design Manager + Engineering Lead)
   - Phase 04: Content & SEO (Content Strategist + SEO Specialist join)
   - Phase 05: Marketing & Growth (Marketing Manager joins)
   - Phase 06: Delivery, Ops & Risks (Engineering Lead + Project Orchestrator)
   - For each phase: Ask questions one at a time, probe deeper when answers are vague, wait for user confirmation before moving to next phase
3. **Compile Summary**: Organize all answers by phase into a structured summary using format from .plan/templates/brainstorm/CONFIRMATION_TEMPLATE.md
4. **Display Confirmation UI**: 
   - Present the summary in a well-formatted markdown display with clear sections, checkmarks (✅), blockquotes for longer answers, and visual separators
   - Include a "Review & Confirm" section with 4 clear options: (1) Save it, (2) Revise a phase, (3) Add information, (4) Start over
   - Wait for explicit user confirmation - DO NOT save until user explicitly confirms
5. **Handle User Response**:
   - If confirmed: Proceed to save
   - If revision requested: Re-ask questions for specified phase(s), update summary, show again
   - If addition requested: Add information to appropriate phase, show updated summary
   - If restart requested: Confirm intent, then restart from Phase 01 if confirmed
6. **Save to BRAINSTORM.md**: Once explicitly confirmed, write the approved summary (organized by phase) to .plan/00_System/BRAINSTORM.md using structure from .plan/templates/brainstorm/TEMPLATE_BRAINSTORM.md
7. **Update State**: Set .plan/active_state.json phase to "brainstorm"
8. **Response**: "✅ Brainstorming session complete! Summary saved to BRAINSTORM.md. Type /write to generate PRD, Architecture, and Design System."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"Design Manager",
				"Content Strategist",
				"SEO Specialist",
				"Marketing Manager",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/IDEA.md",
				".plan/templates/brainstorm/phase-*.md",
				".plan/templates/brainstorm/CONFIRMATION_TEMPLATE.md",
				".plan/templates/brainstorm/TEMPLATE_BRAINSTORM.md",
			},
			FilesModified: []string{
				".plan/00_System/BRAINSTORM.md",
				".plan/active_state.json",
			},
			Requirements: "- Phase templates should exist in .plan/templates/brainstorm/\n- Interview should be conversational, one question at a time\n- Summary must be displayed in formatted confirmation UI before saving\n- User must explicitly confirm before any files are written\n- Use CONFIRMATION_TEMPLATE.md format for displaying summary\n- Use TEMPLATE_BRAINSTORM.md format for final saved document",
			Examples: []string{
				"/improve",
			},
		},
		{
			Name:        "team",
			Trigger:     "/team",
			Description: "Show active agents and hierarchy",
			Action: `When user types /team:

1. **Load Agent Definitions**: Read all agent files from .cursor/agents/
2. **Display Hierarchy**: Show the hierarchical structure of all agents
3. **Show Roles**: Display each agent's role and responsibilities
4. **Response**: Display formatted agent hierarchy and roles`,
			AgentInvolvement: []string{},
			FilesRead: []string{
				".cursor/agents/*.md",
			},
			FilesModified: []string{},
			Examples: []string{
				"/team",
			},
		},
		{
			Name:        "write",
			Trigger:     "/write",
			Description: "Generate PRD + ARCHITECTURE + DESIGN_SYSTEM",
			Action: `When user types /write:

1. **Generate PRD.md**: Product Manager creates comprehensive Product Requirements Document
2. **Generate ARCHITECTURE.md**: Engineering Lead and System Architect create technical architecture
3. **Generate DESIGN_SYSTEM.md**: Design Manager and UI/UX Designer create design system
4. **Save All Files**: Write to .plan/00_System/
5. **Response**: "Documents generated! Review PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md. Type /change to edit any document, or /good to approve."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"System Architect",
				"Design & UX Manager",
				"UI/UX Designer",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/IDEA.md",
				".plan/00_System/BRAINSTORM.md",
			},
			FilesModified: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/DESIGN_SYSTEM.md",
				".plan/active_state.json",
			},
			Examples: []string{
				"/write",
			},
		},
		{
			Name:        "change",
			Trigger:     "/change <document> <change>",
			Description: "Edit any document",
			Action: `When user types /change <document> <change>:

1. **Parse Command**: Extract document name and change description
2. **Load Document**: Read the specified document from .plan/00_System/
3. **Apply Changes**: Update the document with the requested changes
4. **Save Document**: Write updated document back to file
5. **Response**: "Document updated! Changes saved to [document].md"`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/*.md",
			},
			FilesModified: []string{
				".plan/00_System/*.md",
			},
			Examples: []string{
				"/change prd Add dark mode",
				"/change architecture Use PostgreSQL instead of MySQL",
			},
		},
		{
			Name:        "good",
			Trigger:     "/good",
			Description: "Approve & lock plan",
			Action: `When user types /good:

1. **Validate Documents**: Ensure PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md exist
2. **Lock Plan**: Set locked: true in .plan/active_state.json
3. **Update Phase**: Set phase to "approved" in active_state.json
4. **Response**: "Plan approved and locked! Type /plan to generate the execution plan and tasks."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/DESIGN_SYSTEM.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/active_state.json",
				".plan/history/state-*.json",
			},
			Examples: []string{
				"/good",
			},
		},
		{
			Name:        "plan",
			Trigger:     "/plan",
			Description: "Generate execution plan & scaffold phases",
			Action: `When user types /plan:

1. **Synthesize Execution Tasks**: Read .plan/00_System/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md to generate .plan/TASKS.md
2. **Parse TASKS.md**: Use the generated tasks to determine phases and features
3. **Scaffold Phase Folders**: Create phase directories (e.g., 01-Foundation) in .plan/
4. **Generate Feature Folders**: For each task, create feature folders with templates (design.md, plan.md, tasks.md, prompts.md, github.md)
5. **Create Contracts Directory**: Add _contracts/ folder in each phase for shared schemas
6. **Update State**: Update .plan/active_state.json to reference the new hierarchy and set phase to "tasks"
7. **Response**: "Execution plan generated! TASKS.md and phase folders created in .plan/. Type /build to start implementing."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/DESIGN_SYSTEM.md",
				".plan/TASKS.md",
			},
			FilesModified: []string{
				".plan/TASKS.md",
				".plan/[Phase-Folders]/",
				".plan/[Phase-Folders]/[Feature-Folders]/",
				".plan/[Phase-Folders]/_contracts/",
				".plan/active_state.json",
			},
			Requirements: "- PRD, ARCHITECTURE, and DESIGN_SYSTEM must be approved via `/good`\n- Task generation templates live in `.plan/templates`",
			Examples: []string{
				"/plan → Generate execution plan and tasks",
			},
		},
		{
			Name:        "load",
			Trigger:     "/load <path>",
			Description: "Inject context into AI agents",
			Action: `When user types /load <path>:

1. **Parse Path**: Extract file or directory path
2. **Load Content**: Read the specified file or files from directory
3. **Inject Context**: Add content to agent context for current session
4. **Response**: "Context loaded! [path] is now available to agents."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".cursor/rules/library/**",
				".plan/**",
			},
			FilesModified: []string{},
			Examples: []string{
				"/load @library/04-frameworks/frontend/nextjs.md",
				"/load .plan/00_System/PRD.md",
			},
		},
		{
			Name:        "build",
			Trigger:     "/build [<task_id>]",
			Description: "Start coding next task",
			Action:      "When user types /build or /build <task_id>:\n\n1. **Determine Task**:\n   - If task_id provided, load that task\n   - Otherwise, find next uncompleted task from TASKS.md\n2. **Bootstrap Boilerplate (first run only)**:\n   - If the project is still plan-only (no package.json / src/), run `go run scripts/boilerplate/main.go --project .`\n   - This materializes the default stack (Next.js today) so the task has code to modify\n   - Skip if code already exists for this project/stack\n3. **Check Git Status**:\n   - Verify working tree is clean (no uncommitted changes)\n   - If dirty, warn user and block until clean\n4. **Create/Checkout Task Branch**:\n   - Run `go run scripts/branch/main.go --action create --task [ID] --project .`\n   - This creates/checks out branch `task/[ID]` (e.g., `task/5.2`)\n   - Store branch name in `active_branch` field of `.plan/active_state.json`\n5. **Load Task Context**: Read task details, dependencies, and related code\n6. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)\n7. **Start Implementation**: Begin coding the task with full context\n8. **Update State**: Set `active_task` and `active_branch` in `.plan/active_state.json`\n9. **Snapshot State**: Immediately log the new state with `go run scripts/statehistory/main.go snapshot --reason \"build [ID]\" --label build`\n10. **Response**: \"Building task [ID]: [Description] on branch [branch_name]. Focus on this task only.\"",
			AgentInvolvement: []string{
				"Engineering Lead",
				"Relevant Team Leads",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/active_state.json (active_task and active_branch updated)",
				".plan/history/state-*.json (automatic snapshot for audit/rollback)",
				"Git: New branch created/checked out (task/[ID])",
				"src/** (code files created/modified)",
			},
			Examples: []string{
				"/build → Start next uncompleted task",
				"/build 1.2 → Start specific task 1.2",
				"/build 3 → Start task 3",
			},
			GitHubAutomation: `After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update docs/CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md`,
		},
		{
			Name:        "progress",
			Trigger:     "/progress",
			Description: "Show current progress",
			Action:      "When user types /progress:\n\n1. **Read TASKS.md**: Load all tasks\n2. **Read active_state.json**: Get completed tasks and current phase\n3. **Run Progress Tool**: Execute  \n   ```\n   go run scripts/progress/main.go --root <project>\n   ```  \n   This parses `.plan/TASKS.md`, `.plan/active_state.json`, and `.plan/history/` to compute stats and state deltas.\n4. **Calculate Progress**: \n   - Total tasks\n   - Completed tasks\n   - In progress tasks\n   - Percentage complete\n5. **Display Progress**: Show formatted progress report:\n   - Phase: [current phase]\n   - Tasks: X/Y completed (Z%)\n   - Current task: [active task]\n   - Next up: [next task]\n   - State Delta: summarize what changed between the last two snapshots (phase/task/branch/completed)\n6. **Response**: Display progress summary with the state delta footer",
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
				".plan/history/state-*.json",
			},
			FilesModified: []string{},
			Examples: []string{
				"/progress",
			},
		},
		{
			Name:        "finished",
			Trigger:     "/finished [<task_id>]",
			Description: "Mark current task done",
			Action:      "When user types /finished:\n\n1. **Verify Active Branch**: \n   - Check that we're on a task branch (from `active_branch` in `.plan/active_state.json`)\n   - If on main/master, warn and ask for confirmation\n2. **Check Dependencies**: \n   - Run `go run scripts/taskcomplete/main.go --task [ID] --project . --check` to verify all dependencies are complete\n   - If dependencies are missing, **block completion** and list missing dependencies\n3. **Mark Current Task Complete**: \n   - Run `go run scripts/taskcomplete/main.go --task [ID] --project .` to mark task complete in TASKS.md\n   - This updates the task status to \"✅ Complete\" and marks all checklist items as [x]\n4. **Update State**: \n   - Add task ID to completed array in `.plan/active_state.json`\n   - Clear `active_task` and `active_branch` (set to null)\n4. **Snapshot State**: Run `go run scripts/statehistory/main.go snapshot --reason \"finished [ID]\" --label finished` so the new status is recorded in `.plan/history/`\n5. **Auto-Commit**: Automatically commit changes with conventional commit format (e.g., `feat(task-5.2): complete branch automation`)\n6. **Auto-Push**: \n   - Run `go run scripts/branch/main.go --action push --project .` to push the current branch\n   - This pushes the task branch to origin\n7. **Update Changelog**: If significant, add entry to CHANGELOG.md\n8. **Optional PR Creation**: \n   - If `gh` CLI is available, suggest creating a PR with: `gh pr create --title \"feat: [task description]\" --body \"Completes task [ID]\"`\n   - This is optional and can be done manually\n9. **Response**: \"Task marked complete! Changes committed and pushed to [branch_name]. Type /build to start the next task, or /progress to see overall progress.\"",
			AgentInvolvement: []string{
				"Engineering Lead",
				"QA Engineer",
				"Release Captain",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/TASKS.md (task marked [x])",
				".plan/active_state.json (completed array updated, active_task and active_branch cleared)",
				".plan/history/state-*.json (new snapshot)",
				"docs/CHANGELOG.md (if significant changes)",
				"Git: Auto-commit and push to task branch",
			},
			Examples: []string{
				"/finished",
			},
			GitHubAutomation: `- Creates conventional commit message based on changes
- Pushes to current feature/bugfix/hotfix branch
- Follows branching strategy from @library/01-core-workflow/github-workflow-automation.md
- Triggers CI workflow on push`,
		},
		// Workflow Commands
		{
			Name:        "branchci",
			Trigger:     "/branchci [regenerate]",
			Description: "Generate CI workflow for branch prefixes",
			Action:      "When you run `/branchci`:\n\n1. Reads `docs/history/branch-matrix.json` to understand what jobs/required checks each branch prefix needs (e.g., `task/`, `feature/`, `hotfix/`).\n2. Runs the generator:\n   ```bash\n   go run scripts/branchci/main.go --matrix docs/history/branch-matrix.json --out .github/workflows\n   ```\n3. Emits `.github/workflows/task-branches.yml`, a workflow that:\n   - Triggers on pushes to `task/*` (and can be expanded for other prefixes)\n   - Spins up jobs per branch prefix (lint/test/build/etc.)\n   - Adds a summary job so reviewers know which checks are required per branch\n4. Output: \"Workflow generated: .github/workflows/task-branches.yml\"",
			AgentInvolvement: []string{
				"DevOps Engineer",
			},
			FilesRead: []string{
				"docs/history/branch-matrix.json",
			},
			FilesModified: []string{
				".github/workflows/task-branches.yml",
			},
			Customize: "Edit `docs/history/branch-matrix.json` to add or tweak prefixes, jobs, and required checks. Re-run `/branchci` after editing to regenerate the workflow.",
			Notes:     "- Generated workflow expects Go 1.21 and the standard lint/test/build jobs. Adapt `scripts/branchci/main.go` if your stack differs.\n- Use this command whenever you add a new branch naming convention or need different CI steps per branch type.",
			Examples: []string{
				"/branchci",
				"/branchci regenerate",
			},
		},
		{
			Name:        "report",
			Trigger:     "/report or /report <project_path>",
			Description: "Generate scan report metadata and diffs",
			Action:      "When you type /report:\n\n1. **Select Project**:\n   - Default: current workspace (.)\n   - Optional: `/report test/qr-generator/test-no01`\n2. **Generate Metadata**:\n   - Runs `go run scripts/scanreport/main.go --project <path>`\n   - Parses `.plan/reports/SCAN_REPORT_*.md`\n   - Creates/updates matching JSON files with structured metadata (scan date, project, executive summary, findings, next actions, summary hash)\n3. **Compute Diff**:\n   - When >=2 reports exist, compares the newest vs previous\n   - Builds `SCAN_DIFF_<date>.md` highlighting added/removed bullets in Executive Summary, Findings & Risks, Recommended Next Actions, **and** the latest `.plan/history` state changes (phase/task/branch/completed deltas)\n   - Appends preset-specific sections: progress snapshot (from `scripts/progress`), ASCII visuals, and a dependency audit when manifests are detected\n4. **Output**:\n   - Terminal summary showing metadata count + diff file path\n   - Diff markdown stored alongside the reports for sharing",
			AgentInvolvement: []string{
				"QA Engineer",
				"Documentation Lead",
			},
			FilesRead: []string{
				"<project>/.plan/reports/SCAN_REPORT_*.md",
				"<project>/.plan/history/state-*.json (for the state delta section)",
			},
			FilesModified: []string{
				"<project>/.plan/reports/SCAN_REPORT_*.json",
				"<project>/.plan/reports/SCAN_DIFF_<date>.md",
			},
			Options:      "- `--preset standard` *(default)* – complete report\n- `--preset exec` – condensed executive view + visuals\n- `--preset detailed` – expanded sections with dependency audit\n- `.plan/reports/config.json` (optional) can set:\n  ```json\n  {\n    \"preset\": \"exec\",\n    \"sections\": [\"executive\", \"progress\", \"visuals\", \"state\", \"feedback\"]\n  }\n  ```\n  CLI flags override config; custom `sections` let teams reorder or omit report blocks.",
			Requirements: "- Go 1.21+\n- Reports must follow `SCAN_REPORT_YYYY-MM-DD.md` naming",
			Examples: []string{
				"/report → run against current repo",
				"/report test/qr-generator/test-no01 → run inside test fixture",
			},
		},
		{
			Name:        "feedback",
			Trigger:     "/feedback <type> \"Title\" \"Details\" [--github <url>] [--author <name>]",
			Description: "Log feedback (bug, feature, question, note)",
			Action:      "When you run `/feedback ...`:\n\n1. **Parse arguments**\n   - `type`: bug | feature | question | note (defaults to `note`)\n   - `title`: short summary (required)\n   - `details`: multiline description (optional)\n   - `--author`: person filing feedback (defaults to `anonymous`)\n   - `--github`: optional issue URL if mirrored upstream\n2. **Log entry** via `go run scripts/feedback/main.go ...`\n   - Appends markdown to `docs/history/feedback.md`\n   - Updates JSON log `docs/history/feedback.json` for automation\n3. **Surface in workflow**\n   - `/report` command ingests latest feedback when generating scan metadata/diffs\n   - Future scans can summarize outstanding feedback items\n4. **Response**\n   - \"Feedback logged (type=bug) → docs/history/feedback.md\"",
			AgentInvolvement: []string{
				"Product Manager",
				"QA Engineer",
				"Documentation Lead",
			},
			FilesRead: []string{
				"docs/history/feedback.md (created if missing)",
				"docs/history/feedback.json",
			},
			FilesModified: []string{
				"docs/history/feedback.md",
				"docs/history/feedback.json",
			},
			Examples: []string{
				"/feedback bug \"QR download fails\" \"API returns 500 when Accept header missing\"",
				"/feedback feature \"Add dark mode\" \"Marketing wants dark hero section\" --author PM",
				"/feedback question \"Rate limit\" \"What are the prod limits?\" --github https://github.com/org/repo/issues/123",
			},
			Notes: "- Requires Go 1.21+. Command shells run: `go run scripts/feedback/main.go --type <type> --title \"...\" --details \"...\" --author \"...\" --github <url>`\n- Works in any generated project (paths relative to project root).\n- Add new feedback types by passing a custom string (stored as lowercase).",
		},
		{
			Name:        "state",
			Trigger:     "/state <subcommand>",
			Description: "Manage project state history",
			Action:      "The `/state` helper wraps `go run scripts/statehistory/main.go` so you can manage `.plan/active_state.json` history safely.\n\n### snapshot\n1. Writes the current `.plan/active_state.json` into `.plan/history/state-<timestamp>.json`\n2. Accepts optional flags:\n   - `--reason` → stored in the snapshot metadata\n   - `--label` → appended to the file name (e.g., build, finished)\n3. Output: `Snapshot saved: .plan/history/state-20251124T120000Z-build.json`\n\n### list\n1. Lists recent entries (default: last 10)\n2. `--json` emits machine-readable summaries for scripts/CI\n\n### diff\n1. Compares two snapshots (default: latest vs previous)\n2. Shows Markdown summary (phase/task/branch/completed deltas) or JSON if `--json`\n3. Used by `/progress` and `/report` to surface state deltas\n\n### restore\n1. Requires `--file <id>` and `--yes` confirmation for guardrails\n2. Restores `.plan/active_state.json` from the selected snapshot\n3. Optionally captures a new snapshot (`--snapshot=false` to skip) so rollbacks themselves are logged\n4. Respond with confirmation + reminder to rerun `/progress`",
			AgentInvolvement: []string{
				"Project Orchestrator",
				"QA Engineer",
			},
			FilesRead: []string{
				"`.plan/active_state.json`",
				"`.plan/history/state-*.json`",
			},
			FilesModified: []string{
				"`.plan/history/state-*.json` (new entries)",
				"`.plan/active_state.json` (when restoring)",
			},
			Examples: []string{
				"/state snapshot --reason \"after /build 5.8\"",
				"/state list --limit 5",
				"/state diff --json",
				"/state restore --file state-20251124T120000Z.json --yes",
			},
			Notes: "- State history is now required before/after `/build` and `/finished`\n- Restores should be rare; always snapshot first so you can undo mistakes",
		},
		{
			Name:        "github",
			Trigger:     "/github <subcommand>",
			Description: "GitHub metadata and issue management",
			Action:      "### `/github info`\nRuns:\n```\ngo run scripts/githubmeta/main.go --project . --sync-readme\n```\n- Detects primary remote + default branch\n- Extracts success metrics from `.plan/00_System/PRD.md`\n- Updates the README KPI block between `<!-- KPIS:START -->` / `<!-- KPIS:END -->`\n- Persists metadata to `docs/history/github-meta.json` for offline use\n\n### `/github issue \"Title\" \"Body\"`\nOutputs a ready-to-run `gh issue create` command with the detected repo slug, e.g.:\n```\ngo run scripts/githubmeta/main.go --project . --issue-title \"Fix cache\" --issue-body \"Details here\"\n```\nCopy/paste the printed `gh issue create` command (or pipe it) to open the issue.\n\n### `/github milestone \"Name\" [due-date]`\nPrints a `gh api` command to create a milestone:\n```\ngo run scripts/githubmeta/main.go --project . --milestone-title \"MVP\" --milestone-due 2025-01-15T00:00:00Z\n```",
			AgentInvolvement: []string{
				"Release & Growth Manager",
			},
			FilesRead: []string{
				"`.git/` metadata",
				"`.plan/00_System/PRD.md`",
			},
			FilesModified: []string{
				"`docs/history/github-meta.json`",
				"`README.md` KPI section when `--sync-readme` is used",
			},
			OfflineSafety: "- If git remote detection fails, the script logs a warning and keeps the last cached metadata (`docs/history/github-meta.json`). You can still update KPIs from PRD without network access.",
			Examples: []string{
				"/github info",
				"/github issue \"Fix cache\" \"Cache misses spike\"",
			},
		},
		// Squad Commands
		{
			Name:        "secure",
			Trigger:     "/secure",
			Description: "Security review",
			Action: `When user types /secure:

1. **Security Review**: Security Lead conducts security review
2. **Vulnerability Assessment**: Identify and document security vulnerabilities
3. **Generate Report**: Create security review report
4. **Response**: "Security review complete! Review security findings."`,
			AgentInvolvement: []string{
				"Security Lead",
			},
			FilesRead: []string{
				"src/**",
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/SECURITY.md",
			},
			Examples: []string{
				"/secure",
			},
		},
		{
			Name:        "roles",
			Trigger:     "/roles",
			Description: "Design RBAC system",
			Action: `When user types /roles:

1. **Design RBAC**: Security Lead and Backend Lead design role-based access control
2. **Define Roles**: Create role definitions and permissions
3. **Generate Documentation**: Document RBAC system
4. **Response**: "RBAC system designed! Review role definitions."`,
			AgentInvolvement: []string{
				"Security Lead",
				"Backend Lead",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/RBAC.md",
			},
			Examples: []string{
				"/roles",
			},
		},
		{
			Name:        "money",
			Trigger:     "/money",
			Description: "Billing & payment setup",
			Action: `When user types /money:

1. **Payment System Design**: Design billing and payment system
2. **Integration Planning**: Plan payment gateway integration
3. **Generate Documentation**: Document payment system architecture
4. **Response**: "Payment system designed! Review billing architecture."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Backend Lead",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/BILLING.md",
			},
			Examples: []string{
				"/money",
			},
		},
		{
			Name:        "pretty",
			Trigger:     "/pretty",
			Description: "UI/UX improvements",
			Action: `When user types /pretty:

1. **UI/UX Review**: Design Manager and UI/UX Designer review interface
2. **Improvement Suggestions**: Provide UI/UX improvement recommendations
3. **Update Design System**: Update DESIGN_SYSTEM.md with improvements
4. **Response**: "UI/UX improvements suggested! Review design updates."`,
			AgentInvolvement: []string{
				"Design & UX Manager",
				"UI/UX Designer",
			},
			FilesRead: []string{
				".plan/00_System/DESIGN_SYSTEM.md",
				"src/**",
			},
			FilesModified: []string{
				".plan/00_System/DESIGN_SYSTEM.md",
			},
			Examples: []string{
				"/pretty",
			},
		},
		{
			Name:        "seo",
			Trigger:     "/seo",
			Description: "SEO optimization",
			Action: `When user types /seo:

1. **SEO Analysis**: Analyze current SEO implementation
2. **Optimization Recommendations**: Provide SEO optimization suggestions
3. **Generate SEO Plan**: Create SEO optimization plan
4. **Response**: "SEO analysis complete! Review SEO recommendations."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Frontend Lead",
			},
			FilesRead: []string{
				"src/**",
				".plan/00_System/PRD.md",
			},
			FilesModified: []string{
				".plan/00_System/SEO.md",
			},
			Examples: []string{
				"/seo",
			},
		},
		{
			Name:        "ship",
			Trigger:     "/ship",
			Description: "Release management",
			Action: `When user types /ship:

1. **Release Planning**: Release Captain plans the release
2. **Version Management**: Manage version numbers and semantic versioning
3. **Release Notes**: Generate release notes
4. **Deployment Planning**: Plan deployment strategy
5. **Response**: "Release planned! Review release notes and deployment plan."`,
			AgentInvolvement: []string{
				"Release Captain",
				"Release & Growth Manager",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				"docs/CHANGELOG.md",
			},
			FilesModified: []string{
				"docs/CHANGELOG.md",
				".plan/00_System/RELEASE.md",
			},
			Examples: []string{
				"/ship",
			},
		},
		{
			Name:        "safe",
			Trigger:     "/safe",
			Description: "Security audit",
			Action: `When user types /safe:

1. **Security Audit**: Security Lead conducts comprehensive security audit
2. **Vulnerability Scanning**: Scan for security vulnerabilities
3. **Compliance Check**: Verify compliance with security standards
4. **Generate Audit Report**: Create security audit report
5. **Response**: "Security audit complete! Review audit findings."`,
			AgentInvolvement: []string{
				"Security Lead",
			},
			FilesRead: []string{
				"src/**",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/SECURITY.md",
			},
			FilesModified: []string{
				".plan/00_System/SECURITY_AUDIT.md",
			},
			Examples: []string{
				"/safe",
			},
		},
		{
			Name:        "cheap",
			Trigger:     "/cheap",
			Description: "Cost optimization",
			Action: `When user types /cheap:

1. **Cost Analysis**: Analyze current infrastructure and service costs
2. **Optimization Recommendations**: Provide cost optimization suggestions
3. **Generate Cost Plan**: Create cost optimization plan
4. **Response**: "Cost analysis complete! Review optimization recommendations."`,
			AgentInvolvement: []string{
				"DevOps Engineer",
				"Performance Engineer",
			},
			FilesRead: []string{
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/COST_OPTIMIZATION.md",
			},
			Examples: []string{
				"/cheap",
			},
		},
	}
}

// GetCommandByName returns a command by name, or nil if not found
func GetCommandByName(name string) *Command {
	commands := GetAllCommands()
	for i := range commands {
		if commands[i].Name == name {
			return &commands[i]
		}
	}
	return nil
}

// GetCoreCommands returns only the 11 core commands
func GetCoreCommands() []Command {
	allCommands := GetAllCommands()
	return allCommands[:11] // First 11 are core commands
}

// GetSquadCommands returns only the squad-specific commands
func GetSquadCommands() []Command {
	allCommands := GetAllCommands()
	return allCommands[11:] // Remaining are squad commands
}

// commandTemplate is the markdown template for command files
const commandTemplate = `# /{{.Name}}

## Trigger
{{.Trigger}}{{if .Examples}}

## Examples
{{range .Examples}}- {{.}}
{{end}}{{end}}

## Action
{{.Action}}

## Agent Involvement
{{range .AgentInvolvement}}- **{{.}}**
{{end}}{{if .FilesRead}}
## Files Read
{{range .FilesRead}}- {{.}}
{{end}}{{end}}{{if .FilesModified}}
## Files Modified
{{range .FilesModified}}- {{.}}
{{end}}{{end}}{{if .GitHubAutomation}}
## GitHub Automation
{{.GitHubAutomation}}{{end}}{{if .Requirements}}
## Requirements
{{.Requirements}}{{end}}{{if .Notes}}
## Notes
{{.Notes}}{{end}}{{if .Customize}}
## Customize
{{.Customize}}{{end}}{{if .Options}}
## Options
{{.Options}}{{end}}{{if .OfflineSafety}}
## Offline Safety
{{.OfflineSafety}}{{end}}
`

// RenderCommandMarkdown renders a command to markdown format
func RenderCommandMarkdown(cmd *Command) (string, error) {
	if cmd == nil {
		return "", fmt.Errorf("command cannot be nil")
	}

	tmpl, err := template.New("command").Parse(commandTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cmd); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// CommandsGenerator generates the command markdown files.
type CommandsGenerator struct{}

// Name returns the name of the generator.
func (g *CommandsGenerator) Name() string {
	return "Commands"
}

// Generate creates the .cursor/commands directory and all command markdown files.
func (g *CommandsGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	commandsDir := filepath.Join(projectPath, ".cursor", "commands")

	// Create .cursor/commands directory
	if err := utils.CreateDirectory(commandsDir); err != nil {
		return fmt.Errorf("failed to create .cursor/commands directory: %w", err)
	}

	// Get all commands
	commands := GetAllCommands()

	// Generate each command file
	for _, cmd := range commands {
		// Render command markdown
		markdown, err := RenderCommandMarkdown(&cmd)
		if err != nil {
			return fmt.Errorf("failed to render command %s: %w", cmd.Name, err)
		}

		// Write command file
		commandPath := filepath.Join(commandsDir, cmd.Name+".md")
		if err := utils.WriteFile(commandPath, []byte(markdown)); err != nil {
			return fmt.Errorf("failed to write command file %s: %w", cmd.Name+".md", err)
		}
	}

	return nil
}

// GenerateCommands is a convenience function that creates a CommandsGenerator and generates commands
func GenerateCommands(request *models.ProjectRequest, projectPath string) error {
	generator := &CommandsGenerator{}
	return generator.Generate(request, projectPath)
}
