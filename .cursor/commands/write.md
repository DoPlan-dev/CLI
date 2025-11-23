# /write

## Trigger
Exact match: /write

## Action
When user types /write:

1. **Generate PRD.md**: Product Manager creates comprehensive Product Requirements Document
   - User personas
   - Feature list
   - User stories
   - Success metrics
   - Timeline
2. **Generate ARCHITECTURE.md**: Engineering Lead and System Architect create technical architecture
   - System design
   - Technology stack
   - Database schema
   - API design
   - Infrastructure
3. **Generate DESIGN_SYSTEM.md**: Design Manager and UI/UX Designer create design system
   - Design principles
   - Color palette
   - Typography
   - Component library
   - UI patterns
4. **Save All Files**: Write to .plan/00_System/
5. **Response**: "Documents generated! Review PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md. Type /change to edit any document, or /good to approve."

## Agent Involvement
- **Product Manager**: Creates PRD.md
- **Engineering Lead**: Creates ARCHITECTURE.md (with System Architect)
- **Design Manager**: Creates DESIGN_SYSTEM.md (with UI/UX Designer)
- **Project Orchestrator**: Coordinates the document generation

## Files Created
- .plan/00_System/PRD.md
- .plan/00_System/ARCHITECTURE.md
- .plan/00_System/DESIGN_SYSTEM.md
- .plan/active_state.json (phase: "writing")
