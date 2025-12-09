# /write

## Trigger
/write [<subcommand>] [<args>]

## Examples
- "/write → Generate all planning docs (first time) or show menu"
- "/write plan → Show planning document options"
- "/write content → Show content type options"
- "/write change prd Add dark mode → Edit PRD"
- "/write edit architecture Use PostgreSQL → Edit ARCHITECTURE"
- "/write prd → Regenerate PRD only"
- "/write legal → Generate legal pages"


## Action
When user types /write or /write <subcommand>:

1. **If no subcommand provided**: 
   - Check if PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md exist
   - If they don't exist: Generate all three planning documents (PRD, ARCHITECTURE, DESIGN_SYSTEM)
   - If they exist: Show interactive menu:
     - "What would you like to generate or edit?"
     - "1. Planning Documents (PRD, Architecture, Design System)"
     - "2. Content (app pages, legal, blog, social, marketing, email, docs, SEO)"
     - "3. Edit Document (change existing document)"
     - Wait for user selection

2. **Subcommand: plan** (or user selects option 1):
   - Show planning document options:
     - "1. PRD only"
     - "2. ARCHITECTURE only"
     - "3. DESIGN_SYSTEM only"
     - "4. All planning documents"
   - Generate selected document(s)

3. **Subcommand: content** (or user selects option 2):
   - Load content requirements from BRAINSTORM.md
   - Show content type options based on meeting requirements:
     - App pages, Legal pages, Blog posts, Social media, Marketing, Email templates, Documentation, SEO content
   - Generate selected content type(s)

4. **Subcommand: change <document> <change>** (or user selects option 3):
   - Parse document name and change description
   - Load the specified document from .do/system/
   - Apply changes to the document
   - Save updated document back to file
   - Response: "Document updated! Changes saved to [document].md"
   - Alternative: /write edit <document> <change> (alias for change)

5. **Other subcommands**:
   - /write prd → Regenerate PRD only
   - /write architecture → Regenerate ARCHITECTURE only
   - /write design → Regenerate DESIGN_SYSTEM only
   - /write app-pages → Generate app pages content
   - /write legal → Generate legal pages
   - /write blog → Generate blog posts
   - /write social → Generate social media content
   - /write marketing → Generate marketing content
   - /write email → Generate email templates
   - /write docs → Generate documentation pages
   - /write seo → Generate SEO content
   - /write all → Generate everything

6. **Response**: "Documents/content generated! Review files in .do/system/ or .do/system/content/."

## Agent Involvement
- **Product Manager**
- **Engineering Lead**
- **System Architect**
- **Design & UX Manager**
- **UI/UX Designer**
- **Content Strategist**
- **SEO Specialist**
- **Documentation Writer**
- **Project Orchestrator**

## Files Read
- ".do/system/IDEA.md"
- ".do/system/BRAINSTORM.md"
- ".do/system/*.md"

## Files Modified
- ".do/system/PRD.md"
- ".do/system/ARCHITECTURE.md"
- ".do/system/DESIGN_SYSTEM.md"
- ".do/system/*.md"
- ".do/system/content/**"
- ".do/system/history/active_state.json"


