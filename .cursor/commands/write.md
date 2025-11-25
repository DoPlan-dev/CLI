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
4. **Generate API Contract**: API Owner captures machine-readable contract in `.plan/contracts/`
   - Endpoint catalog (paths, methods, descriptions)
   - Request/response schemas with required headers
   - Authentication flows and security requirements
   - Status codes, error shapes, and constraints
5. **Save All Files**: Write core docs to `.plan/00_System/` and contracts to `.plan/contracts/`
6. **Response**: "Documents generated! Review PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md, and the API contracts in .plan/contracts. Type /change to edit any document, or /good to approve."

## Agent Involvement
- **Product Manager**: Creates PRD.md
- **Engineering Lead**: Creates ARCHITECTURE.md (with System Architect)
- **Design Manager**: Creates DESIGN_SYSTEM.md (with UI/UX Designer)
- **API Owner**: Authors API contracts within `.plan/contracts/`
- **Project Orchestrator**: Coordinates the document generation

## Files Created
- .plan/00_System/PRD.md
- .plan/00_System/ARCHITECTURE.md
- .plan/00_System/DESIGN_SYSTEM.md
- .plan/contracts/API_CONTRACT.md (or service-specific contracts)
- .plan/active_state.json (phase: "writing")
