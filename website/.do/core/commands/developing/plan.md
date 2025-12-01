# /plan

## Trigger
/plan [<subcommand>] [<args>]

## Examples
- "/plan → Generate all planning docs (first time) or show menu"
- "/plan docs → Generate PRD, Architecture, Design System"
- "/plan content → Generate content"
- "/plan everything → Generate full execution plan"
- "/plan phases → Plan project phase by phase"
- "/plan next → Planning next phase"
- "/plan phase 1 tasks → Create tasks.md for phase 1"
- "/plan phases tasks → Create tasks.md for all phases"
- "/plan all tasks → Create main TASKS.md with all project tasks"


## Action
When user types /plan or /plan <subcommand>:

**Document Generation (write functionality)**

1. **If no subcommand provided**: 
   - Check if PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md exist
   - If they don't exist: Generate all three planning documents (PRD, ARCHITECTURE, DESIGN_SYSTEM)
   - If they exist: Show interactive menu:
     - "What would you like to generate or edit?"
     - "1. Planning Documents (PRD, Architecture, Design System)"
     - "2. Content (app pages, legal, blog, social, marketing, email, docs, SEO)"
     - "3. Execution Plan & Tasks"
     - "4. Edit Document (change existing document)"
     - Wait for user selection

2. **Subcommand: docs** (or user selects option 1):
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

4. **Subcommand: change <document> <change>** (or user selects option 4):
   - Parse document name and change description
   - Load the specified document from .do/system/
   - Apply changes to the document
   - Save updated document back to file
   - Response: "Document updated! Changes saved to [document].md"
   - Alternative: /plan edit <document> <change> (alias for change)

5. **Document subcommands**:
   - /plan prd → Regenerate PRD only
   - /plan architecture → Regenerate ARCHITECTURE only
   - /plan design → Regenerate DESIGN_SYSTEM only
   - /plan app-pages → Generate app pages content
   - /plan legal → Generate legal pages
   - /plan blog → Generate blog posts
   - /plan social → Generate social media content
   - /plan marketing → Generate marketing content
   - /plan email → Generate email templates
   - /plan docs → Generate documentation pages
   - /plan seo → Generate SEO content
   - /plan all → Generate everything

**Execution Plan Generation (plan functionality)**

6. **Subcommand: everything** (or user selects option 3):
   - **Check Documents**: Verify PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md exist in .do/system/
   - **Check Approval Status**: Read .do/system/history/active_state.json for "approved" status
   - **If Not Approved**:
     * Show warning: "⚠️ Planning documents have not been approved yet."
     * Show documents status:
       - PRD.md: [exists/doesn't exist]
       - ARCHITECTURE.md: [exists/doesn't exist]
       - DESIGN_SYSTEM.md: [exists/doesn't exist]
     * Ask: "Do you want to proceed with generating the execution plan anyway? (yes/no)"
     * If yes → Proceed to next step
     * If no → "Please review and approve documents first. Type /plan docs to regenerate if needed."
   - **If Approved or User Confirmed**:
     * Synthesize Execution Tasks: Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md to generate .do/plan/TASKS.md
     * Parse TASKS.md: Use the generated tasks to determine phases and features
     * Scaffold Phase Folders: Create phase directories (e.g., 01-Foundation) in .do/plan/
     * Generate Feature Folders: For each task, create feature folders with templates (design.md, plan.md, tasks.md, prompts.md, github.md)
     * Create Contracts Directory: Add _contracts/ folder in each phase for shared schemas
     * Update State: Update .do/system/history/active_state.json to reference the new hierarchy and set phase to "tasks"
     * Mark plan as generated in active_state.json
   - **Response**: "Execution plan generated! TASKS.md and phase folders created in .do/plan/. Type /dev to start implementing."

7. **Subcommand: phases** (or /plan phases):
   - Plan project phase by phase:
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * Generate phase structure but don't create all tasks at once
     * Create phase folders (01-Foundation, 02-Core, etc.) without feature folders yet
     * Set up phase-by-phase planning mode in active_state.json
     * After user finishes a feature with /dev, system will ask: "Ready to plan the next phase? Type /plan next to continue."
   - **Response**: "Phase-by-phase planning mode activated! Phase folders created. Complete features with /dev, then type /plan next to plan the next phase."

8. **Subcommand: next** (or /plan next):
   - Planning next phase:
     * Read active_state.json to determine current phase
     * Find the next unplanned phase
     * Check if previous phase features are complete (if applicable)
     * Generate tasks for the next phase only
     * Create feature folders for this phase with templates
     * Update active_state.json with current phase
     * **Response**: "Next phase planned! Phase [X] tasks and feature folders created. Type /dev to start implementing."

9. **Subcommand: phase {no} tasks** (or /plan phase <number> tasks):
   - Create tasks.md for a specific phase:
     * Parse phase number from command (e.g., /plan phase 1 tasks)
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * Generate tasks for this phase only
     * Create tasks.md file inside .do/plan/[Phase-Folder]/ (e.g., .do/plan/01-Foundation/tasks.md)
     * Include all feature tasks for this phase
     * **Response**: "Phase [X] tasks created! tasks.md saved in .do/plan/[Phase-Folder]/tasks.md"

10. **Subcommand: phases tasks** (or /plan phases tasks):
   - Create tasks.md for all phases:
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * For each phase folder in .do/plan/:
       - Generate tasks for that phase
       - Create tasks.md file inside that phase folder
       - Include all feature tasks for that phase
     * **Response**: "All phase tasks created! tasks.md files saved in each phase folder."

11. **Subcommand: all tasks** (or /plan all tasks):
   - **This is the ONLY way to create the main TASKS.md file**:
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * Generate ALL tasks for the entire project
     * Create .do/plan/TASKS.md with all project tasks organized by phase
     * This is the master tasks file that contains everything
     * **Response**: "Main TASKS.md created! All project tasks saved in .do/plan/TASKS.md"

## Agent Involvement
- **Product Manager**
- **Engineering Lead**
- **Project Orchestrator**

## Files Read
- ".do/system/PRD.md"
- ".do/system/ARCHITECTURE.md"
- ".do/system/DESIGN_SYSTEM.md"
- ".do/system/history/active_state.json"
- ".do/plan/TASKS.md"

## Files Modified
- ".do/plan/TASKS.md"
- ".do/plan/[Phase-Folders]/"
- ".do/plan/[Phase-Folders]/[Feature-Folders]/"
- ".do/plan/[Phase-Folders]/_contracts/"
- ".do/system/history/active_state.json"

## Requirements
|

