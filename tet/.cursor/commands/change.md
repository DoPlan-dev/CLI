# /change

## Trigger
/change <document> <change>

## Examples
- /change prd Add dark mode
- /change architecture Use PostgreSQL instead of MySQL


## Action
When user types /change <document> <change>:

1. **Parse Command**: Extract document name and change description
2. **Load Document**: Read the specified document from .plan/00_System/
3. **Apply Changes**: Update the document with the requested changes
4. **Save Document**: Write updated document back to file
5. **Response**: "Document updated! Changes saved to [document].md"

## Agent Involvement
- **Project Orchestrator**

## Files Read
- .plan/00_System/*.md

## Files Modified
- .plan/00_System/*.md

