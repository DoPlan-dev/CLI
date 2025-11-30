# /plan

## Trigger
Exact match: /plan

## Action
When user types /plan:

1. **Synthesize Execution Tasks**: Read `.plan/00_System/PRD.md`, `ARCHITECTURE.md`, and `DESIGN_SYSTEM.md` to generate a detailed `.plan/TASKS.md`
2. **Parse TASKS.md**: Use the generated tasks to extract phases and features
3. **Scaffold Phase Folders**: Create phase directories (e.g., `01-Foundation`, `02-Core_Features`) in `.plan/`
4. **Generate Feature Folders**: For each task, create feature folder with templates:
   - `design.md` - Design decisions and UI/UX considerations
   - `plan.md` - Implementation plan and objectives
   - `tasks.md` - Subtask checklist
   - `prompts.md` - AI prompts and interactions log
   - `github.md` - GitHub issues, PRs, and branch strategy
5. **Create Contracts Directory**: Add `_contracts/` folder in each phase for shared API/data schemas
6. **Update State**: Update `.plan/active_state.json` to reference the new hierarchy and set phase to "tasks"
7. **Response**: "Execution plan generated! TASKS.md and phase folders created in `.plan/`. Type /build to start implementing."

## Agent Involvement
- **Product Manager**: Validates phase structure
- **Engineering Lead**: Reviews feature organization
- **Project Orchestrator**: Coordinates scaffolding

## Files Read
- `.plan/00_System/PRD.md`
- `.plan/00_System/ARCHITECTURE.md`
- `.plan/00_System/DESIGN_SYSTEM.md`
- `.plan/TASKS.md`

## Files Modified
- `.plan/[Phase-Folders]/` (new phase directories)
- `.plan/[Phase-Folders]/[Feature-Folders]/` (new feature directories with templates)
- `.plan/[Phase-Folders]/_contracts/` (contracts directories)
- `.plan/active_state.json` (updated with hierarchy reference)

## Requirements
- PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md are approved via `/good`
- Task generation templates in `.plan/templates` can be customized per team

## Examples
- `/plan` → Scaffold planning hierarchy from TASKS.md

---
End Command ---

