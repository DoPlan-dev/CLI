## Beginner Experience Roadmap (Target: GoPlan CLI 1.2.0)

Captured from internal discussion on improving onboarding for non-experts and AI-first beginners.

### 1. Guided Onboarding Mode
- Beginner preset toggle in Bubbletea wizard or `--beginner` flag.
- Preselect defaults (Cursor, safe project name) and show contextual tips in wizard footer.
- Auto-generate `Docs/BEGINNER_CHECKLIST.md` with step-by-step “Next run /tell → /improve → /write → /plan”.
- Post-generation prompt to open the checklist in the IDE; optional `/tutorial start` auto-run.

### 2. Interactive Tutorials
- New `/tutorial` command family with segments (`start`, `step <n>`, `docs`) referencing shared JSON/script data.
- Embed 30-minute playbook: idea capture, document review, planning.
- Provide copy/paste prompt snippets for AI IDE chats plus Markdown quickstart in `Docs/quickstart.md`.
- Align tutorial content with beginner preset so messaging stays consistent.

### 3. Errors-as-Guidance
- Central error help registry mapping known errors to fixes + wiki links.
- CLI commands funnel errors through helper that prints “Summary → Fix command → Docs reference”.
- Wizard `stateError` uses the same suggestions; highlight recovery keys (`r` retry, `b` back).
- Optional logging to `.plan/history/errors.log` to track common failure reasons.

### Next Release Plan
- Aim for GoPlan CLI **1.2.0** once at least the beginner preset + tutorial system ship together.
- Use 1.1.x patch releases for any bugs discovered post-1.1.0.

