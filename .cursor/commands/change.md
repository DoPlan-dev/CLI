# /change

## Trigger
Exact match: /change <doc> <instructions>

Examples:
- /change prd Add dark mode requirement
- /change architecture Use PostgreSQL instead of MongoDB
- /change design Use blue instead of green

## Action
When user types /change <doc> <instructions>:

1. **Parse Command**: Extract document name and change instructions
2. **Identify Document**: Map to correct file:
   - prd → .plan/00_System/PRD.md
   - architecture or arch → .plan/00_System/ARCHITECTURE.md
   - design or design_system → .plan/00_System/DESIGN_SYSTEM.md
3. **Activate Relevant Agent**: 
   - PRD changes → Product Manager
   - Architecture changes → Engineering Lead
   - Design changes → Design Manager
4. **Apply Changes**: Agent reads current document, applies changes, saves updated version
5. **Response**: "Changes applied to [document]. Review the updated file."

## Agent Involvement
- **Product Manager**: Handles PRD.md changes
- **Engineering Lead**: Handles ARCHITECTURE.md changes
- **Design Manager**: Handles DESIGN_SYSTEM.md changes

## Files Modified
- .plan/00_System/[DOCUMENT].md (updated)
