# Docs Directory (Generated Projects)

This folder mirrors the structure we expect every DoPlan-generated project to ship with. Treat it as the canonical reference for how documentation must be organized.

## 📁 Categories

| Directory | Purpose | Examples |
| --- | --- | --- |
| `Docs/foundation/` | Global system references shared across the whole product | `foundation/the-guide.md`, `foundation/roadmap.md` |
| `Docs/features/` | Feature-scoped specs (one folder per feature/phase) | `features/01_foundation/README.md`, `features/_plan/TASKS.md` |
| `Docs/release/` | Launch-readiness material | `release/launch-checklist.md`, `release/retro.md` |
| `Docs/history/` | Prompt logs, retro notes, Git logs | `history/prompts.md`, `history/changelogs.md` |

> Need another category? Create a subfolder under `Docs/` and document it here so every project stays consistent.

## 🧭 Rules

1. **Root stays clean.** Only `README.md`, `CHANGELOG.md`, and generated code/config live at the repo root—every other document belongs somewhere under `Docs/`.
2. **One document = one place.** Don’t duplicate docs in multiple categories; link instead.
3. **Feature folders mirror phase/task IDs.** `/plan` copies `.plan/00_System/*.md` into `Docs/foundation/`, `TASKS.md` into `Docs/features/_plan/`, and creates `Docs/features/<Phase_Title>/README.md` for each phase.
4. **History never leaves `Docs/history/`.** Prompt transcripts, Git timelines, or retros all live there.

## 🎨 Branding assets

- `docs/ascii.md` holds the canonical ASCII wordmark. Reference or embed it from there anytime a doc (e.g., READMEs, TUI specs) needs the project banner so we only maintain it in one place.
- `docs/tui/header.md` mirrors the same art for the interactive TUI docs; if the banner changes update `docs/ascii.md` first, then sync this file.

## 🔗 Reference templates

- [`Docs/foundation/the-guide.md`](foundation/the-guide.md) mirrors the canonical test project's end-to-end workflow (`test/qr-generator/test-no01/Docs/the-guide.md`). Use it as the authoritative example when describing commands or process in new projects.
