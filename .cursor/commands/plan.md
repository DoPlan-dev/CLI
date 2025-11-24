# /plan

## Trigger
Exact match: /plan

## Action
When user types /plan:

1. **Parse TASKS.md**: Read `.plan/TASKS.md` to extract phases and features
2. **Scaffold Phase Folders**: Create phase directories (e.g., `01-Foundation`, `02-Core_Features`) in `.plan/`
3. **Generate Feature Folders**: For each task, create feature folder with templates:
   - `design.md` - Design decisions and UI/UX considerations
   - `plan.md` - Implementation plan and objectives
   - `tasks.md` - Subtask checklist
   - `prompts.md` - AI prompts and interactions log
   - `github.md` - GitHub issues, PRs, and branch strategy
4. **Create Contracts Directory**: Add `_contracts/` folder in each phase for shared API/data schemas
5. **Update State**: Update `.plan/active_state.json` to reference the new hierarchy
6. **Response**: "Planning hierarchy scaffolded! Phase folders created in `.plan/`. Type /build to start implementing."

## Agent Involvement
- **Product Manager**: Validates phase structure
- **Engineering Lead**: Reviews feature organization
- **Project Orchestrator**: Coordinates scaffolding

## Files Read
- `.plan/TASKS.md`

## Files Modified
- `.plan/[Phase-Folders]/` (new phase directories)
- `.plan/[Phase-Folders]/[Feature-Folders]/` (new feature directories with templates)
- `.plan/[Phase-Folders]/_contracts/` (contracts directories)
- `.plan/active_state.json` (updated with hierarchy reference)

## Requirements
- TASKS.md must exist with phase and task definitions
- Phases should follow format: `## Phase N: Name`
- Tasks should follow format: `### N.N Task Title`

## Examples
- `/plan` → Scaffold planning hierarchy from TASKS.md

---
End Command ---

