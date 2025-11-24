# Docs Folder Structure Rules

## Canonical Layout
- Every handcrafted document belongs under `Docs/`
  - `foundation/` → global references (IDEA, PRD, Architecture, Design System, brainstorm logs)
  - `features/` → feature/phase specific specs, tasks, decisions (`Docs/features/<Feature_Name>/`)
  - `release/` → launch plans, retrospectives, comms
  - `history/` → prompt transcripts, changelog excerpts, audit trails
- Root stays clean: only `README.md`, `CHANGELOG.md`, generated code/config may live outside `Docs/`.

## Generation Requirements
1. `/plan` must mirror `.plan/00_System/*.md` into `Docs/foundation/`.
2. `/plan` must copy `TASKS.md` into `Docs/features/_plan/TASKS.md` so feature teams see the checklist without digging into `.plan/`.
3. Documentation generators must create the base `Docs/` hierarchy (with `.gitkeep` placeholders) for every new project.

## Checklist
- [ ] When adding a new doc, confirm it lives under the correct `Docs/` category.
- [ ] Feature folders must include the task/feature identifier (e.g., `Docs/features/01_Auth_Feature/`).
- [ ] If a new category is required, update `Docs/README.md` so all teammates + agents share the same map.
- [ ] Scans and `/report` should read from `Docs/` when referencing long-form docs to keep root uncluttered.
# Documentation Folder & Clean Root Rules

## When to use this rule
- Creating or updating project documentation
- Adding new auto-generated docs via commands like `/plan`, `/write`, `/build`, `/finished`
- Reviewing PRs that introduce new markdown files

## Required structure
1. **All documents live under `Docs/`.** Never drop standalone `.md` files into the repository root.
2. **Use categories:**
   - `Docs/foundation/` – Global guides (`the-guide.md`, `roadmap.md`, etc.)
   - `Docs/features/<Feature_Name>/` – Feature-specific specs (`spec.md`, `tasks.md`, `history.md`)
   - `Docs/release/` – Launch assets, retro notes, announcements
   - `Docs/history/` – Prompt logs, branch histories, retrospectives
3. **Feature folders mirror `/plan` output.** If `/plan` creates `01_Auth_Feature/`, the corresponding docs reside in `Docs/features/01_Auth_Feature/`.

## Root cleanliness
- The repo root should only contain source code, configuration, `README.md`, and `CHANGELOG.md`.
- Any new supporting write-up must be placed under an appropriate `Docs/` subdirectory and referenced from `Docs/README.md`.

## Enforcement

### Automated Checks
- **Lint Script**: `scripts/check-docs-organization.sh` validates documentation organization
  - Run before committing: `./scripts/check-docs-organization.sh`
  - Checks for root-level `.md` files (blocks if found)
  - Validates `Docs/` structure exists with required subdirectories
- **CI/CD Integration**: Automated checks run in GitHub Actions on every PR
  - Job: `docs-check` in `.github/workflows/ci.yml`
  - Blocks PRs that violate clean root policy

### Manual Enforcement
- Update `Docs/README.md` whenever you add a document so future contributors know where to find it.
- Add `.gitkeep` files to empty directories to keep the category structure intact.
- PR reviewers should block merges that introduce root-level `.md` files or documents outside `Docs/`.

