# Dashboard Changes Reference (Comprehensive)

Use this checklist to reapply every dashboard fix since opening the project in Cursor. Paths are relative to `dashboard/`.

## Shared Header (all pages)
- Brand: `DoPlan` stays blue `#206bc4`; `.dev` is black.
- Nav items: icons blue `#206bc4`; text black `#111`; active link uses a light blue background.
- Right-side links (desktop only): Docs → `https://doplan.dev/docs`; GitHub → `https://github.com/doplan/cli` with blue icons, black text.
- Applied to: `index.html`, `settings.html`, `plan.html`, `meetings.html`, `achievements.html`.
- Mobile menu: unchanged; new links are hidden on small screens (`d-none d-md-flex`).

## Home (`index.html`)
- Removed “View Reports”; only “Go to Plan”.
- Enlarged hero artwork to 384×288.
- Project Info card (beside welcome):
  - Shows project name, type badge, GitHub repo link, branch count, current branch badge, tech stack badges (with icons), lines of code, total size.
  - Data source: `data/project.json` with added fields:
    - `github.repo`, `github.branches`, `github.current_branch`
    - `tech_stack` (array)
    - `stats.lines_of_code`, `stats.total_size`
  - Tech stack icon map: js, ts, python, go, react, vue, node, html, css, git, github.
  - Fallbacks: display “Not available” in gray when missing.
- Next Recommendation block:
  - Phase/plan/tasks-driven commands: `/plan`, `/build`, `/status report`.
  - Uses progress (completed/total) and phase; also set in the fallback loader.
- Progress cleanup: removed hardcoded Sprint `+12%` and Pending `-2%`; now fully dynamic from `project.json`.

## Plan (`plan.html`)
- Columns: only To Do and Done (In Progress removed).
- Task cards: compact, expandable; phase badge + time on top; feature name; task title with chevron on the right; description collapsible.
- Phase-colored left border; distinct colors per phase.
- Removed gray task ID badge.

## Agents (`settings.html`)
- Agent files served from `dashboard/data/agents/` (copied, not symlinked; Python `http.server` blocks `..` and symlinks).
- Loader tries paths in order: `data/agents/{category}/{agent}.md`, then `.do/core/agents/**`, then `.cursor/agents/**`.
- Regex: name from first `# Heading`; role from `## Role` section.
- If none load, show a warning row. Added console logs for troubleshooting.
- Full 18 agents included (leadership, product, engineering: lead/system/frontend/backend/devops/security/performance, design: manager/uiux, documentation: lead/writer, quality: qa manager/engineer, release: manager/captain/growth).

## Tasks & Data Access
- `plan.html` fetches `data/TASKS.md` (copy of `.do/plan/TASKS.md` because `..` is blocked).
- `index.html` uses `data/project.json` for progress/cards.
- `project.json` now includes GitHub, tech_stack, and stats sections.

## Why copy instead of symlink
- `python3 -m http.server` blocks `..` paths and often refuses to serve symlinks. Copying into `dashboard/data/` ensures fetch works.

## Files touched
- HTML: `index.html`, `settings.html`, `plan.html`, `meetings.html`, `achievements.html`
- Data: `data/project.json`, `data/TASKS.md`, `data/agents/**` (copied from `.do/core/agents/**`)

## Regeneration Steps (manual)
1) Copy data:
   - Tasks: `cp .do/plan/TASKS.md dashboard/data/TASKS.md`
   - Agents:
     ```
     rm -rf dashboard/data/agents && mkdir -p dashboard/data/agents
     for d in .do/core/agents/*/; do b=$(basename "$d"); mkdir -p dashboard/data/agents/$b; cp "$d"*.md dashboard/data/agents/$b/; done
     ```
   - Ensure `dashboard/data/project.json` includes:
     ```
     "github": { "repo": "", "branches": 0, "current_branch": "main" },
     "tech_stack": [],
     "stats": { "lines_of_code": 0, "total_size": "-" }
     ```
2) Reapply header styling/links across all pages.
3) Serve locally and verify:
   - `python3 -m http.server 8000 --directory dashboard`
   - Check progress on home and tasks on plan load correctly.

## Quick Troubleshooting
- Agents not loading: confirm files exist in `dashboard/data/agents/**`; check browser console logs.
- Tasks not loading: ensure `data/TASKS.md` exists (copied).
- Progress wrong: confirm `data/project.json` has correct `progress` fields and added GitHub/tech_stack/stats sections.