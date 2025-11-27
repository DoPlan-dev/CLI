# Docs Directory (Generated Projects)

This folder mirrors the structure we expect every DoPlan-generated project to ship with. Treat it as the canonical reference for how documentation must be organized.

## 📁 Categories

| Directory | Purpose | Examples |
| --- | --- | --- |
| `docs/foundation/` | Global system references shared across the whole product | `foundation/the-guide.md`, `foundation/roadmap.md` |
| `docs/features/` | Feature-scoped specs (one folder per feature/phase) | `features/01_foundation/README.md`, `features/_plan/TASKS.md` |
| `docs/release/` | Launch-readiness material | `release/launch-checklist.md`, `release/retro.md` |
| `docs/history/` | Prompt logs, retro notes, Git logs | `history/prompts.md`, `history/changelogs.md` |

> Need another category? Create a subfolder under `docs/` and document it here so every project stays consistent.

## 🔨 Generating Source Code

Projects now start with planning + documentation only. When you're ready to build, scaffold the codebase with your preferred tool (e.g., `npx create-next-app`, `pnpm create`, `go mod init`). The legacy `scripts/boilerplate` helper has been removed, so `/build` expects you to bring or generate the starter code manually.

## 🧭 Rules

1. **Root stays clean.** Only `README.md` and generated code/config live at the repo root—every other document belongs somewhere under `docs/`.
   - **Enforcement**: `scripts/check-docs-organization.sh` validates this policy
   - **CI Integration**: This check runs in CI/CD to block non-compliant PRs
   - **Rule Reference**: `.cursor/rules/library/11-documentation/docs-folder-structure.md`
2. **One document = one place.** Don't duplicate docs in multiple categories; link instead.
3. **Feature folders mirror phase/task IDs.** `/plan` copies `.plan/00_System/*.md` into `docs/foundation/`, `TASKS.md` into `docs/features/_plan/`, and creates `docs/features/<Phase_Title>/README.md` for each phase.
4. **History never leaves `docs/history/`.** Prompt transcripts, Git timelines, or retros all live there.

## 🚫 Clean Root Policy

**All documentation must live under `docs/`.** The repository root must remain clean with only:
- `README.md` - Project overview
- Generated code and configuration files

### Enforcement

- **Lint Script**: Run `./scripts/check-docs-organization.sh` before committing
- **CI/CD**: Automated checks run on every PR
- **PR Review**: PRs introducing root-level `.md` files will be blocked

### Where to Put New Docs

- **Foundation docs** → `docs/foundation/`
- **Feature specs** → `docs/features/<Feature_Name>/`
- **Release plans** → `docs/release/`
- **History/audit logs** → `docs/history/`

See `.cursor/rules/library/11-documentation/docs-folder-structure.md` for complete guidelines.

## 🎨 Branding assets

- `docs/ascii.md` holds the canonical ASCII wordmark. Reference or embed it from there anytime a doc (e.g., READMEs, TUI specs) needs the project banner so we only maintain it in one place.
- `docs/tui/header.md` mirrors the same art for the interactive TUI docs; if the banner changes update `docs/ascii.md` first, then sync this file.

## 🔗 Reference templates

- [`docs/foundation/the-guide.md`](foundation/the-guide.md) mirrors the canonical test project's end-to-end workflow (`test/qr-generator/test-no01/docs/the-guide.md`). Use it as the authoritative example when describing commands or process in new projects.
