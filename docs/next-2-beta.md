# **DoPlan CLI \- Official Release Plan**

This document outlines the features and implementation plan for the next two major beta releases, building on v0.0.17-beta.

## **📁 STEP 1: PROVIDE CONTEXT TO CURSOR**

Copy and paste this prompt to Cursor Composer with your folders added:

I'm implementing the new DoPlan CLI release plan, starting with v0.0.18-beta.

CONTEXT PROVIDED:  
✓ DoPlan CLI source code (Go \+ Cobra \+ Viper \+ Bubble Tea \+ Lipgloss)  
✓ The official \`doplan\_release\_plan.md\` (this file)  
✓ Test Project 1:  /Users/Dorgham/Documents/Work/Devleopment/Club Website
✓ Test Project 2: /Users/Dorgham/Documents/Work/Devleopment/mediabubble

OBJECTIVE:  
Implement the first release, \`v0.0.18-beta\`, by following its five phases (Architecture, Root Command, Dashboard, Documentation, and Integration).

Ready to start with \`v0.0.18-beta\` \- Phase 1: Architecture Setup?

## **🚀 Release: v0.0.17-beta (Baseline)**

**Change Notes:**

* This is the current version.  
* It establishes the initial DoPlan concept with a CLI-first approach.  
* **Known Issues:**  
  1. CONTEXT.md and README.md generation is incorrect.  
  2. TUI Dashboard (doplan \--tui) does not load or display project data.  
  3. Folder structure (01-phase) is not user-friendly.  
  4. No mechanism to adopt existing projects.  
  5. Workflow is tied to a CLI-first model, not a TUI or IDE-first model.

## **🚀 Release: v0.0.18-beta (The Foundation & Fixes Release)**

**Goal:** To fix all critical bugs from v0.0.17 and implement a new, stable, TUI-first/IDE-agnostic architecture. This release makes DoPlan stable, usable, and ready for advanced features.

**Change Notes:**

* **New TUI-First Entry:** doplan is now the only command, launching a smart TUI dashboard. All sub-commands are removed in favor of TUI/AI actions.  
* **New: Project Adoption:** doplan will now detect existing projects (without .doplan) and launch an "Adoption Wizard."  
* **New: Folder Naming Convention:** All project folders will now be created with the \#\#-slug-name format (e.g., 01-user-authentication/01-login-with-email). The CLI is now able to parse this structure.  
* **New: IDE-Agnostic Integration:** Adds a new TUI wizard (⚙️ Setup AI/IDE Integration) to auto-configure DoPlan for Cursor, VS Code (Copilot), Kiro, Qoder, etc. The AI logic is now stored generically in .doplan/ai/.  
* **New: Mandatory GitHub:** The core workflow now requires a GitHub repo, with a new TUI wizard for setup.  
* **Fix: Dashboard Data:** The TUI dashboard now correctly loads and displays all project data from .doplan/dashboard.json.  
* **Fix: Documentation:** CONTEXT.md is now correctly generated as project-specific (for AI), and README.md is now project-first (for humans).  
* **New: Foundational Polish:** Adds a robust new system for error handling, logging, animations, and graceful tool-installation prompts.

### **v0.0.18-beta \- Phase 1: Architecture Setup**

*(This was Phase 0 in the old plan)*

**Prompt for Cursor:**

PHASE 0: Architecture Setup

Create the foundational architecture for all new features:

1\. PROJECT STRUCTURE  
Update internal/ directory structure to support:  
\- internal/tui/ → Main TUI app, state management, navigation, components/, screens/  
\- internal/context/ → Smart context detection and project analysis  
\- internal/commands/ → Logic for new \*actions\* (run, undo, deploy, publish, create, security, fix, design).  
\- internal/wizard/ → New project and adoption wizards  
\- internal/deployment/ → Multi-platform deployment (Vercel, Netlify, Railway, Render, Coolify, custom)  
\- internal/publisher/ → Package publishing (npm, Homebrew, Scoop, Winget)  
\- internal/security/ → Security scanning (npm audit, trufflehog, git-secrets, gosec, dive)  
\- internal/fixer/ → Auto-fix system with AI integration  
\- internal/sops/ → Service setup procedures generator  
\- internal/rakd/ → Required API keys detector and validator  
\- internal/dpr/ → Design preferences questionnaire and generator  
\- internal/agents/ → AI agent file generator  
\- internal/workflow/ → New workflow guidance engine  
\- internal/integration/ → \*\*NEW: Logic for IDE & AI-specific setup (Cursor, VS Code, Kiro, etc.)\*\*

Create pkg/ directory for shared utilities:  
\- pkg/theme/ → Lipgloss color palette and reusable styles  
\- pkg/errors/ → Beautiful error types with fix suggestions  
\- pkg/logger/ → Beautiful logging with icons and colors  
\- pkg/animations/ → Spinners and smooth animations

2\. CONFIGURATION SYSTEM  
Update .doplan/config.yaml schema using Viper:  
\- project: name, type, version, \*\*ide\*\* (e.g., "cursor", "vscode", "kiro")  
\- github: repository (REQUIRED), enabled, autoBranch, autoPR  
\- design: hasPreferences, tokensPath  
\- security: lastScan, autoFix  
\- apis: configured\[\], required\[\]  
\- tui: theme, animations

3\. STATE MANAGEMENT  
Create .doplan/state.json for:  
\- Action history (for undo feature)  
\- Last 10 operations  
\- Timestamps and details  
\- Checkpoint references

4\. DASHBOARD JSON  
Create .doplan/dashboard.json (machine-readable) with:  
\- Project summary and progress  
\- Phase and feature details  
\- GitHub activity  
\- API keys status  
\- Velocity metrics  
\- Recent activity feed

5\. LIPGLOSS THEME  
Create comprehensive theme system in pkg/theme/:  
\- Color palette (Primary Blue, Secondary Purple, Success Green, Warning Amber, Error Red)  
\- Reusable styles (Header, Card, Button, Progress, Success, Error)  
\- Consistent styling across all TUI screens

Follow existing DoPlan patterns. Use Cobra for commands, Viper for config, Bubble Tea for TUI, Lipgloss for styling.

Confirm when architecture is ready.

### **v0.0.18-beta \- Phase 2: Smart Root Command & Context Detection**

*(This was Phase 1 in the old plan)*

**Prompt for Cursor:**

PHASE 1: Core Command & Root Behavior

Implement intelligent root command that's context-aware:

1\. CONTEXT DETECTION SYSTEM  
Create internal/context/detector.go:  
\- Detect 5 states: EmptyFolder, ExistingCodeNoDoPlan, DoPlanInstalled, InsideFeature, InsidePhase  
\- Check for .doplan folder existence  
\- Check for project files (package.json, go.mod, requirements.txt, etc.)  
\- \*\*Detect if inside phase/feature directories (by checking for the \`\#\#-slug-name\` pattern and matching against \`dashboard.json\`).\*\*  
\- Check for git repository

2\. SMART ROOT COMMAND BEHAVIOR  
Update cmd/doplan/main.go Execute() function:  
\- When user runs \`doplan\` with no arguments:  
  \* Empty folder → Launch full TUI new project wizard  
  \* Existing code but no .doplan → Show "Adopt Project?" wizard  
  \* .doplan exists → Open main TUI dashboard instantly  
  \* Inside feature folder → Show feature-specific view  
  \* Inside phase folder → Show phase-specific view  
\- Ensure NO OTHER sub-commands are registered. \`doplan\` is the only entry point.

3\. ADD DASHBOARD ALIAS (TUI ENTRY)  
Add aliases to the root \`doplan\` command: \[".", "dash", "d"\]  
So \`doplan .\` also opens the main TUI dashboard.

4\. NEW PROJECT WIZARD  
Create internal/wizard/new\_project.go with beautiful TUI flow:  
Screen 1: Welcome with DoPlan logo (lipgloss styled)  
Screen 2: Project name input  
Screen 3: Template gallery with preview (saas, mobile, ai-agent, landing, chrome-ext, electron, api, cli)  
Screen 4: GitHub repo setup (MANDATORY \- cannot proceed without)  
Screen 5: \*\*IDE & AI Selection (NEW): "Which AI/IDE will you use?" (Cursor, Kiro, Copilot, windsurf, qoder, Gemini, Claude, Other)\*\*  
Screen 6: Installation progress with spinner  
Screen 7: Success message \+ next steps  
\- \*\*After installation, automatically run the \`setup:ide\` logic (from Phase 5 of this release) based on the user's choice.\*\*  
Auto-open dashboard after completion.

5\. ADOPT EXISTING PROJECT WIZARD  
Create internal/wizard/adopt\_project.go:  
Screen 1: "Found existing project\!" message  
Screen 2: Analysis results (tech stack, files, existing docs)  
Screen 3: Options (Analyze & generate plan / Import existing docs / Start fresh)  
Screen 4: GitHub repo setup (MANDATORY)  
Screen 5: \*\*IDE & AI Selection (NEW): "Which AI/IDE will you use?" (Cursor, Kiro, Copilot, windsurf, qoder, Gemini, Claude, Other)\*\*  
Screen 6: Analysis progress with spinner  
Screen 7: Show generated plan preview  
Screen 8: Confirm adoption  
\- \*\*After adoption, automatically run the \`setup:ide\` logic (from Phase 5 of this release).\*\*

6\. PROJECT ANALYZER  
Create internal/context/analyzer.go:  
\- Detect tech stack from package.json, go.mod, etc.  
\- Map existing folder structure  
\- Find existing documentation files  
\- \*\*Identify potential phases/features from \`\#\#-slug-name\` folder names. If only names exist, suggest renaming to add number prefixes.\*\*  
\- Extract TODO comments as tasks  
\- Generate reverse plan from code structure

Use lipgloss extensively for beautiful output. Add loading spinners, progress bars, animations.

Test all scenarios: empty folder, Next.js app, Go API, Python project, existing DoPlan project.

Confirm when complete.

### **v0.0.18-beta \- Phase 3: Dashboard Supercharge (Bug Fix)**

*(This was Phase 7 in the old plan)*

**Prompt for Cursor:**

PHASE 7: Machine-Readable Dashboard

Fix dashboard data display and add real-time updates:

1\. DASHBOARD.JSON FORMAT  
Create .doplan/dashboard.json (machine-readable):

Structure:  
{  
  "version": "1.0",  
  "generated": "timestamp",  
  "project": {name, description, version, progress, status, dates},  
  "github": {repository, branch, commits, contributors, lastCommit},  
  "phases": \[  
    {  
      id, name, description, status, progress, dates,  
      features: \[  
        {id, name, status, progress, branch, pr, commits, lastActivity, tasks}  
      \],  
      stats: {totalFeatures, completed, inProgress, todo, totalTasks, completedTasks}  
    }  
  \],  
  "summary": {totalPhases, completed, inProgress, todo, totalFeatures, totalTasks, etc.},  
  "activity": {  
    last24Hours: {commits, tasksCompleted, filesChanged},  
    last7Days: {similar},  
    recentActivity: \[{type, message, timestamp}, ...\]  
  },  
  "apiKeys": {total, configured, pending, optional, completion},  
  "velocity": {tasksPerDay, commitsPerDay, estimatedCompletion, daysToLaunch}  
}

2\. DASHBOARD GENERATOR  
Update internal/dashboard/generator.go:  
\- Read all progress.json files  
\- Parse phase and feature states  
\- Fetch GitHub data (commits, PRs, branches)  
\- Calculate velocity metrics  
\- Generate activity feed  
\- Check API keys status from RAKD (it will be empty for now, that's OK)  
\- Output to dashboard.json  
\- Also generate dashboard.md (human-readable markdown)

3\. UPDATE TUI TO READ JSON  
Update internal/tui/screens/dashboard.go:  
\- Load dashboard.json instead of parsing markdown  
\- Instant loading (\<100ms)  
\- Display with beautiful lipgloss styling:  
  \* Header with project name and GitHub badge  
  \* Overall progress bar with percentage  
  \* Stats grid (phases, features, tasks, API keys)  
  \* Velocity section with sparkline  
  \* Current phase details with feature progress  
  \* Recent activity feed (last 5 items)  
\- Auto-refresh every 30 seconds  
\- Show last update time

4\. SPARKLINES  
Implement sparkline visualization:  
\- Show velocity trend over last 14 days  
\- Use characters:  ▂▃▄▅▆▇█  
\- Display in velocity section  
\- Color code (green=good, amber=slowing, red=stalled)

5\. REAL-TIME PROGRESS BARS  
Fix progress bars to show actual data:  
\- Read from dashboard.json  
\- Display with correct percentages  
\- Use filled/empty characters: █░  
\- Color code by status (green=complete, blue=progress, gray=todo)  
\- Show mini progress bars for features

6\. ACTIVITY FEED  
Generate from:  
\- Git commits (last 10\)  
\- Task completions (from progress updates)  
\- PR merges  
\- Feature status changes  
\- Phase completions  
Show with icons and time ago (2m, 1h, 1d)

7\. AUTO-UPDATE TRIGGERS  
Update dashboard.json when:  
\- \`/progress\` runs  
\- \`/github\` runs  
\- Feature status changes  
\- Tasks completed  
\- Any git operation  
\- Config changes

Keep dashboard.md for human reading in GitHub. Both should be in sync.

Test with real project data. Verify instant loading and accurate percentages.

Confirm when dashboard is supercharged.

### **v0.0.18-beta \- Phase 4: Project-First Documentation (Bug Fix)**

*(This was Phase 3 in the old plan)*

**Prompt for Cursor:**

PHASE 3: Fix Documentation Generation

Implement project-first, AI-ready documentation.

Part A: CONTEXT.md Improvements  
Fix CONTEXT.md to be project-specific instead of DoPlan-focused.

Current Issue: CONTEXT.md contains only DoPlan documentation links, not actual project context for the AI.

Required Changes:  
1\. CONTEXT.md Structure:  
   \# Project Context: \[Project Name\]

   \#\# Project Overview  
   \- Brief description  
   \- Target audience  
   \- Core features

   \#\# Technology Stack  
   \#\#\# Frontend  
   \- Framework: \[e.g., Next.js 14\]  
   \- UI Library: \[e.g., Tailwind CSS\]  
   \- State Management: \[e.g., Zustand\]  
     
   \#\#\# Backend  
   \- Runtime: \[e.g., Node.js\]  
   \- Framework: \[e.g., Express\]  
   \- Database: \[e.g., PostgreSQL\]  
     
   \#\#\# Services & APIs  
   \- Authentication: \[service name \+ link to SOP\]  
   \- Storage: \[service name \+ link to SOP\]  
   \- Payment: \[service name \+ link to SOP\]

   \#\# Project-Specific Documentation  
   \- \[Custom API Docs\](./contracts/api-spec.json)  
   \- \[Data Models\](./contracts/data-model.md)  
   \- \[Design System\](./design/DPR.md)

   \#\# Development Guidelines  
   \- Coding standards  
   \- File naming conventions  
   \- Component patterns  
   \- Testing approach

   \#\# DoPlan Resources  
   \<details\>  
   \<summary\>DoPlan CLI Documentation\</summary\>  
   \[Minimized DoPlan links here\]  
   \</details\>

2\. Auto-generate project-specific content from:  
   \- package.json analysis  
   \- Code structure scanning  
   \- User inputs from /Discuss  
   \- Detected services

3\. Update CONTEXT.md whenever:  
   \- Dependencies change  
   \- New services added  
   \- Project structure updated

Modify: internal/generator/context.go

Part B: README.md Improvements  
Restructure README.md to prioritize project information over DoPlan.

New README.md Structure:  
\# \[Project Name\]

\[Project tagline/description\]

\#\# 🚀 Quick Start  
\[Installation and setup for THIS project\]

\#\# 📋 Features  
\[Project features, not DoPlan features\]

\#\# 🛠️ Tech Stack  
\[Project technologies with links\]

\#\# 📁 Project Structure  
\[This project's folder organization, \*\*reflecting the \`\#\#-phase-name/\#\#-feature-name\` structure.\*\*\]

\#\# 🔑 Environment Variables  
\[Required .env variables \- link to RAKD.md\]

\#\# 📚 Documentation  
\- \[Product Requirements\](./doplan/PRD.md)  
\- \[API Specification\](./doplan/contracts/api-spec.json)  
\- \[Design Guidelines\](./doplan/design/DPR.md)  
\- \[Development Progress\](./doplan/dashboard.md)

\#\# 🤝 Contributing  
\[Project-specific guidelines\]

\---

\<details\>  
\<summary\>💼 Powered by DoPlan\</summary\>

This project uses DoPlan for workflow automation and project management.

\#\#\# DoPlan Commands  
\- \`/Discuss\` \- Refine ideas  
\- \`/Generate\` \- Create docs  
\- \`/Plan\` \- Structure phases  
\- \[View all commands\](./.doplan/ai/commands/)

\[Minimal DoPlan info\]

\</details\>

Implementation:  
\- Create README.md template with placeholders  
\- Auto-populate from project data  
\- Move DoPlan info to collapsible section  
\- Update during /Generate command

Modify: internal/generator/readme.go

Confirm when complete.

### **v0.0.18-beta \- Phase 5: GitHub & IDE Integration**

*(This was Phases 4 & 10 in the old plan)*

**Prompt for Cursor:**

PHASE 4 & 10: Mandatory GitHub & IDE Integration

Part A: GitHub Repository Requirement  
Make GitHub repository mandatory for all DoPlan features:

1\. REPOSITORY REQUIREMENT CHECK  
Create internal/github/validator.go:  
\- Function RequireGitHubRepo(cmd string) that checks config  
\- If github.repository is empty → return error  
\- If repo exists → validate access with GitHub CLI  
\- List of actions that require GitHub: discuss, plan, implement, feature

2\. BLOCK ACTIONS WITHOUT REPO  
Update TUI and AI command logic:  
\- Before running a protected action (like /Discuss or TUI \> Discuss Idea)  
\- Check for GitHub repo using RequireGitHubRepo()  
\- If not configured → show beautiful TUI error  
\- Offer to launch GitHub setup wizard now  
\- Cannot proceed without setting repo

3\. GITHUB SETUP WIZARD TUI  
Create internal/tui/screens/github\_setup.go:  
Screen 1: "GitHub Required" explanation with reasons why  
Screen 2: Options (Enter existing URL / Create new / Skip with warning)  
Screen 3a: If existing → input URL, validate, check access  
Screen 3b: If new → name, description, visibility, initialize options  
Screen 4: Success confirmation  
Save to .doplan/config.yaml under github.repository

4\. INTEGRATE WITH /DISCUSS  
Update /Discuss logic:  
\- First screen checks for GitHub repo  
\- If not set → show GitHub setup before allowing discussion  
\- Display warning badge if skipped  
\- Remind at feature creation

5\. DASHBOARD BADGE  
Add permanent GitHub badge to dashboard:  
\- Show repository name as clickable badge  
\- Display at top of dashboard  
\- Style with lipgloss (rounded border, primary color)  
\- Include commit count, last commit time

6\. VALIDATION IN INSTALL  
During \`doplan install\` (New Project Wizard):  
\- Ask for GitHub repo as part of setup  
\- Make it clear it's required for full functionality  
\- Offer to create new repo during installation  
\- Save immediately to config

Part B: IDE & AI Integration (IDE Agnosticism)  
Create the logic to make the DoPlan AI system work with any IDE.

1\. CREATE IDE INTEGRATION WIZARD  
Create internal/integration/wizard.go:  
\- This TUI wizard asks the user to select their primary AI/IDE from the list in \`config.yaml\` (Cursor, Kiro, Copilot, windsurf, qoder, etc.).  
\- It saves this choice to \`.doplan/config.yaml\` under the \`project.ide\` key.

2\. CREATE INTEGRATION LOGIC  
Create internal/integration/setup.go with a main \`SetupIDE(ideName string)\` function:  
\- This function reads the \`ideName\` and performs specific setup actions.

3\. CURSOR INTEGRATION  
\`SetupIDE("cursor")\`:  
\- Check if \`.cursor/\` directory exists.  
\- Automatically create (or symlink) the following:  
  \- \`.cursor/agents/\` \-\> from \`.doplan/ai/agents/\`  
  \- \`.cursor/rules/\` \-\> from \`.doplan/ai/rules/\`  
  \- \`.cursor/commands/\` \-\> from \`.doplan/ai/commands/\`  
\- This makes all the generic agents, rules, and commands instantly available in Cursor.

4\. VS CODE \+ COPILOT INTEGRATION  
\`SetupIDE("copilot")\`:  
\- Check if \`.vscode/\` directory exists.  
\- Generate \`.vscode/tasks.json\` with tasks to trigger all \`doplan\` TUI actions (e.g., "Run Dev Server", "Deploy Project").  
\- Generate \`.vscode/settings.json\` to recommend extensions (e.g., Copilot Chat).  
\- Generate \`.vscode/prompts/\` (or similar) with files that \`@include\` the agent definitions from \`.doplan/ai/agents/\`, making them available to Copilot.

5\. GENERIC/OTHER INTEGRATION  
\`SetupIDE("kiro")\`, \`SetupIDE("windsurf")\`, \`SetupIDE("qoder")\`, \`SetupIDE("other")\`:  
\- These tools may have unknown configuration formats.  
\- Create a \`.doplan/guides/\` directory.  
\- Generate a markdown file, e.g., \`.doplan/guides/kiro\_setup.md\`.  
\- This guide will instruct the user:  
  "To integrate DoPlan with Kiro, please configure Kiro to read its agent prompts from the \`.doplan/ai/agents/\` directory and its rules from \`.doplan/ai/rules/\`.  
  You can trigger DoPlan actions using the \`doplan\` TUI in your terminal."  
\- This provides a clear path for any AI tool to adopt the DoPlan system.

6\. UPDATE TUI MENU  
\- Create the main TUI app and navigation in \`internal/tui/app.go\`, \`internal/tui/navigation.go\`, etc.  
\- The main menu (launched by \`doplan\`) should be the central hub. For this release, it will be \*partially\* populated.  
\- Add the following menu items:  
  \- 1\. 📊 View Dashboard  
  \- 9\. 💬 Discuss Idea  
  \- 10\. 📝 Generate Documents  
  \- 11\. 🗺️  Create Plan  
  \- 12\. 📈 Update Progress  
  \- \*\*15. ⚙️  Setup AI/IDE Integration (Triggers this logic)\*\*  
\- (The other items will be added in the next release)

Confirm when the integration layer is complete.

### **v0.0.18-beta \- Phase 6: Foundational Polish**

*(This was Phase 9 in the old plan)*

**Prompt for Cursor:**

PHASE 9: Final Polish \- Make It Feel Like Magic

Add finishing touches to create the perfect developer experience:

1\. LIPGLOSS EVERYWHERE  
Audit all TUI screens and ensure consistent use of lipgloss theme:  
\- All headers use HeaderStyle  
\- All cards use CardStyle  
\- All buttons use ButtonStyle  
\- All progress bars use ProgressFilled/Empty  
\- All success messages use SuccessStyle  
\- All errors use ErrorStyle  
\- Consistent padding, margins, borders  
\- Beautiful color palette throughout

2\. GRACEFUL TOOL INSTALLATION  
Create pkg/tools/installer.go:  
\- Before using any external tool (gh, npm, docker, playwright-cli, etc.)  
\- Check if installed  
\- If not → show beautiful prompt asking to install  
\- Show installation progress with spinner  
\- Verify installation succeeded  
\- Continue with operation  
\- Handle declined installation gracefully

3\. PERFECT ERROR HANDLING  
Create pkg/errors/errors.go:  
\- DoPlanError type with Code, Message, Details, Fix  
\- Beautiful error display with lipgloss styling  
\- Always include "How to fix" section  
\- Display error code for support  
\- Common errors predefined (NoGitHubRepo, InvalidProjectStructure, etc.)  
\- Show errors in red with icon  
\- Suggest next steps

4\. BEAUTIFUL LOGGING  
Create pkg/logger/logger.go:  
\- Info (ℹ icon, primary color)  
\- Success (✓ icon, green)  
\- Warning (⚠ icon, amber)  
\- Error (✗ icon, red)  
\- Step (\[1/5\] format for multi-step operations)  
\- Write to .doplan/logs/doplan.log  
\- Rotate logs (keep last 10 files)

5\. SMOOTH ANIMATIONS  
Create pkg/animations/spinner.go:  
\- Spinner with frames: ⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏  
\- 80ms frame duration  
\- Show during any operation \>1 second  
\- Clear line when done

Confirm when all polish is applied.

## **🚀 Release: v0.0.19-beta (The Advanced Features Release)**

**Goal:** To build on the stable v0.0.18 foundation by adding all the new advanced TUI actions and the complete AI Agent workflow.

**Change Notes:**

* **New TUI Actions:** The main TUI menu is now fully populated with powerful new actions, including:  
  * ▶️ Run Dev Server  
  * ↩️ Undo Last Action  
  * 🚀 Deploy Project  
  * 🛡️ Run Security Scan  
  * 🩹 Auto-fix Issues  
  * 📦 Publish Package  
* **New: AI Agent System:** The full, IDE-agnostic agent system is now live.  
  * Includes @planner, @coder, @designer, @reviewer, @devops, and @tester.  
  * Agents follow strict rules defined in .doplan/ai/rules/workflow.mdc.  
* **New: Playwright Testing:** The @tester agent now automatically runs Playwright (MCP) tests and saves screenshots to .doplan/artifacts/.  
* **New: Design System (DPR):** A new TUI menu option (🎨 Apply Design / DPR) launches an interactive questionnaire to generate a complete DPR.md and design tokens.  
* **New: Secrets Management (RAKD/SOPS):** A new TUI menu option (🔑 Manage API Keys) launches a wizard to detect, validate, and manage all API keys and service guides.  
* **New: Workflow Guidance:** The TUI now provides a "Recommended Next Step" after each action to guide the user through the perfect workflow.

### **v0.0.19-beta \- Phase 1: Unified TUI & AI Commands**

*(This was Phase 2 in the old plan)*

**Prompt for Cursor:**

PHASE 2: Unified TUI & AI Command Definitions

Implement the advanced TUI actions and the AI command definitions.

1\. CREATE AI COMMAND DEFINITIONS  
Create files in \`.doplan/ai/commands/\` for all new actions:  
\- .doplan/ai/commands/run.md (Starts dev server)  
\- .doplan/ai/commands/undo.md (Reverts last DoPlan action)  
\- .doplan/ai/commands/deploy.md (Starts deployment wizard)  
\- .doplan/ai/commands/publish.md (Starts package publishing wizard)  
\- .doplan/ai/commands/create.md (Starts new project wizard)  
\- .doplan/ai/commands/security.md (Runs security scan)  
\- .doplan/ai/commands/fix.md (Runs AI-powered auto-fix)  
\- .doplan/ai/commands/design.md  
\- .doplan/ai/commands/keys.md

Each .md file should define the command's purpose, behavior, and IDE-agnostic trigger (e.g., "/run"). These are the \*source of truth\*.

2\. IMPLEMENT BACKEND ACTION LOGIC  
Create the Go functions that perform the actions in \`internal/commands/\` (as defined in Phase 0):  
\- internal/commands/run.go (Auto-detects and runs dev server)  
\- internal/commands/undo.go (Uses state.json for time-machine undo)  
\- internal/commands/deploy.go (Launches multi-platform deployment wizard)  
\- internal/commands/publish.go (Launches package publishing wizard)  
\- internal/commands/create.go (Launches template gallery)  
\- internal/commands/security.go (Runs comprehensive security scan)  
\- internal/commands/fix.go (Runs AI-powered auto-fix)

3\. IMPLEMENT UNIFIED TUI  
Update the main TUI app and navigation in \`internal/tui/app.go\`:  
The main menu (launched by \`doplan\`) should now be \*fully populated\*:  
\- 1\. 📊 View Dashboard  
\- 2\. ▶️  Run Dev Server (triggers \`run\` logic)  
\- 3\. ↩️  Undo Last Action (triggers \`undo\` logic)  
\- 4\. 🚀 Deploy Project (triggers \`deploy\` logic)  
\- 5\. 📦 Publish Package (triggers \`publish\` logic)  
\- 6\. ✨ Create New Project (triggers \`create\` logic)  
\- 7\. 🛡️ Run Security Scan (triggers \`security\` logic)  
\- 8\. 🩹 Auto-fix Issues (triggers \`fix\` logic)  
\- 9\. 💬 Discuss Idea  
\- 10\. 📝 Generate Documents  
\- 11\. 🗺️  Create Plan  
\- 12\. 📈 Update Progress  
\- 13\. 🔑 Manage API Keys (triggers Phase 3 of this release)  
\- 14\. 🎨 Apply Design (triggers Phase 2 of this release)  
\- 15\. ⚙️  Setup AI/IDE Integration

Use Bubble Tea \+ Lipgloss for a beautiful, navigable interface with keyboard shortcuts.

Test all new actions from \*both\* the AI-integrated IDE and the main TUI menu.

### **v0.0.19-beta \- Phase 2: Design System (DPR) Generation**

*(This was Phase 5 in the old plan)*

**Prompt for Cursor:**

PHASE 5: Design Preferences & Requirements (DPR) Document

Create comprehensive design system generation with interactive questionnaire:

1\. DESIGN COMMAND & TUI  
\- Connect the TUI Menu Item: "🎨 Apply Design / DPR" (from the previous phase) to launch the interactive TUI questionnaire.  
\- This wizard generates:  
  \- \`DPR.md\` document  
  \- \`design-tokens.json\`  
  \- \`.doplan/ai/rules/design\_rules.mdc\`

2\. INTERACTIVE QUESTIONNAIRE TUI  
Create internal/dpr/questionnaire.go with 20-30 questions:  
(Audience, Emotional Design, Style, Colors, Typography, Layout, Components, Animation, References)  
...  
Show progress bar throughout questionnaire.  
Allow back navigation.  
Save partial progress.

3\. GENERATE DPR.MD  
Create internal/dpr/generator.go:  
Generate comprehensive document with sections:  
(Executive Summary, Audience Analysis, Design Principles, Visual Identity, Layout, Component Library, Animation, Wireframes, Accessibility, etc.)  
...

4\. GENERATE DESIGN TOKENS JSON  
Create internal/dpr/tokens.go:  
Export to doplan/design/design-tokens.json:  
(colors, typography, spacing, borderRadius, shadows, breakpoints)  
...

5\. GENERATE AI DESIGN RULES  
Create internal/dpr/cursor\_rules.go:  
Generate \`.doplan/ai/rules/design\_rules.mdc\` with:  
\- Color usage rules (only from DPR)  
\- Typography rules (type scale, line heights)  
\- Spacing rules (use design tokens)  
\- Component guidelines  
\- Responsive rules  
\- Accessibility requirements  
\- Code style (Tailwind utilities)

Use beautiful TUI with lipgloss. Make questionnaire feel conversational and helpful.

Test with different project types. Confirm when design system generation is complete.

### **v0.0.19-beta \- Phase 3: Secrets & API Keys (RAKD/SOPS)**

*(This was Phase 6 in the old plan)*

**Prompt for Cursor:**

PHASE 6: SOPS & RAKD Systems

Implement comprehensive API key and service management:

1\. SOPS (SERVICE OPERATING PROCEDURES) FOLDER  
Create internal/sops/generator.go:  
Auto-generate service setup guides in .doplan/SOPS/:  
(authentication/, database/, payment/, storage/, email/, analytics/, ai/)  
...  
Each SOP document template must include:  
(Service Overview, When to Use, Setup Steps, API Key Creation, Env Vars, Code Examples, Testing, Common Issues, Resources)  
...

2\. AUTO-DETECT SERVICES  
Create internal/sops/detector.go:  
\- Scan package.json dependencies  
\- Detect from imports in code  
...

3\. RAKD (REQUIRED API KEYS DOCUMENT)  
Create internal/rakd/generator.go:  
Generate doplan/RAKD.md with structure:  
(Quick Status, Configured section, Pending section, Optional section, Validation Results, Quick Actions, etc.)  
...

4\. API KEY DETECTION  
Create internal/rakd/detector.go:  
\- Detect required services from dependencies  
\- Check .env file for configured keys  
...

5\. API KEY VALIDATION  
Create internal/rakd/validator.go:  
\- Check .env file exists  
\- Validate format of API keys  
\- Test connections (if safe to do)  
...

6\. KEYS COMMAND & TUI  
\- Connect the TUI Menu Item: "🔑 Manage API Keys" (from Phase 1 of this release) to launch a TUI screen for key management.  
\- TUI Screen features:  
  \- Show RAKD status (list, status, priority)  
  \- Action: Validate all keys  
  \- Action: Check for missing keys  
  \- Action: Sync .env.example  
  \- Action: Launch setup wizard for a service  
  \- Action: Test API connections

7\. DASHBOARD WIDGET  
Add to main dashboard:  
Display API Keys Status card:  
\- Progress bar of configuration  
\- Count of configured/pending/optional  
\- Highlight high-priority missing keys  
\- Hotkey to open keys management TUI

8\. AUTO-UPDATE SYSTEM  
Update RAKD.md automatically when:  
\- Dependencies change  
\- Keys added to .env  
...

Use beautiful TUI. Color-code by status (green=configured, amber=pending, red=high priority missing).

Test with various project types. Confirm when secrets management is complete.

### **v0.0.19-beta \- Phase 4: AI Agents System**

*(This was Phase 8 in the old plan)*

**Prompt for Cursor:**

PHASE 8: Specialized AI Agents & Workflow Rules

Create \`.doplan/ai/agents/\` and \`.doplan/ai/rules/\` folders with specialized agent configurations and workflow definitions:

1\. AGENTS & RULES STRUCTURE (IDE-AGNOSTIC)  
Generate during \`doplan install\` into \`.doplan/ai/\`:  
.doplan/ai/agents/  
├── README.md (usage guide for agents)  
├── planner.agent.md  
├── coder.agent.md  
├── designer.agent.md  
├── reviewer.agent.md  
├── tester.agent.md  
└── devops.agent.md  
.doplan/ai/rules/  
├── workflow.mdc (The "perfect workflow" sequence: Plan \-\> Design \-\> Code \-\> Test \-\> Review \-\> Deploy)  
├── communication.mdc (How agents must interact and hand off tasks)  
└── design\_rules.mdc (Generated by Phase 2 of this release)

2\. AGENTS README.MD  
Create comprehensive guide:  
\- List all available agents (including @tester)  
\- Explain role of each agent  
\- How to activate (@agent-name in an AI chat)  
...  
\- \*\*Crucially: Explains that all agents MUST follow the \`.doplan/ai/rules/workflow.mdc\` and \`.doplan/ai/rules/communication.mdc\` files.\*\*  
\- Examples of multi-agent conversations

3\. PLANNER AGENT  
Create \`.doplan/ai/agents/planner.agent.md\` with detailed sections:  
\- Role & Identity (senior project planner)  
...  
\- \*\*Workflow & Rules: You MUST read and obey \`.doplan/ai/rules/workflow.mdc\` and \`.doplan/ai/rules/communication.mdc\`. Your job is the FIRST step.\*\*  
\- \*\*Folder Structure: You MUST create all phase and feature folders using a numbered and slugified name, for example: \`doplan/01-user-authentication/01-login-with-email/\`. This provides both human readability and clear ordering.\*\*  
\- Commands & Workflows (/Plan, /Plan:Phase, /Plan:Reorder, /Plan:Dependencies)  
...

4\. CODER AGENT  
Create \`.doplan/ai/agents/coder.agent.md\`:  
\- Role: Implementation specialist  
...  
\- \*\*Workflow & Rules: You MUST follow \`workflow.mdc\`. You only begin work AFTER @planner and @designer are finished. You MUST tag @tester when your work is ready for review, as defined in \`communication.mdc\`.\*\*  
...

5\. DESIGNER AGENT  
Create \`.doplan/ai/agents/designer.agent.md\`:  
\- Role: UI/UX specialist  
...  
\- \*\*Workflow & Rules: You MUST follow \`workflow.mdc\` and \`design\_rules.mdc\`. Your work happens AFTER @planner. You must provide clear specs for @coder.\*\*  
...

6\. REVIEWER AGENT  
Create \`.doplan/ai/agents/reviewer.agent.md\`:  
\- Role: Quality assurance  
...  
\- \*\*Workflow & Rules: You MUST follow \`workflow.mdc\`. Your review happens ONLY AFTER @tester has successfully run all tests. You MUST follow \`communication.mdc\` for approving or rejecting work.\*\*  
...

7\. TESTER AGENT  
Create \`.doplan/ai/agents/tester.agent.md\`:  
\- Role: QA & Test Automation Specialist  
\- \*\*Workflow & Rules: You MUST follow \`workflow.mdc\`. Your work begins WHEN tagged by @coder. You MUST tag @reviewer with a test report (pass or fail) as defined in \`communication.mdc\`.\*\*  
\- Responsibilities:  
  \- Generates end-to-end test scenarios from feature acceptance criteria.  
  \- Writes and executes automated tests using Playwright (MCP framework).  
  \- \*\*Captures screenshots of completed features and saves them to \`.doplan/artifacts/screenshots/{phase-name}/{feature-name}.png\` (e.g., \`.../01-user-authentication/01-login-with-email.png\`) for visual verification and historical records.\*\*  
  \- Performs visual regression checks.  
  \- Reports bugs with detailed steps to reproduce, console logs, and screenshots.  
...

8\. DEVOPS AGENT  
Create \`.doplan/ai/agents/devops.agent.md\`:  
\- Role: Deployment and infrastructure specialist  
\- \*\*Workflow & Rules: You MUST follow \`workflow.mdc\`. Your work begins ONLY AFTER @reviewer has approved a feature or release. You MUST report deployment status back to the team.\*\*  
...

9\. AGENT & RULES GENERATOR  
Create internal/agents/generator.go:  
\- \*\*Generate all agent files\*\* (including tester.agent.md) into \`.doplan/ai/agents/\` during installation.  
\- \*\*Generate the core \`workflow.mdc\` and \`communication.mdc\` files\*\* into \`.doplan/ai/rules/\` with the standard DoPlan workflow logic.  
...

Test agents with any supported AI IDE. Verify they reference correct project files and follow the new workflow rules.

Confirm when all agents and rules are created and documented.

### **v0.0.19-beta \- Phase 5: Workflow Guidance Engine**

*(This was the final part of Phase 9 in the old plan)*

**Prompt for Cursor:**

PHASE 9 (Partial): Workflow Guidance Engine

Implement the "Recommended Next Step" feature to guide users.

1\. WORKFLOW GUIDANCE ENGINE  
Create internal/workflow/recommender.go:  
\- Implements a function \`GetNextStep(lastAction string) (string, string)\`.  
\- This function maps a completed action (e.g., "plan\_complete", "feature\_implemented", "project\_created") to a recommended next step.  
\- Example: \`GetNextStep("plan\_complete")\` returns ("Implement First Feature", "@coder /implement \<feature-name\> or use the TUI menu").  
\- Update the main TUI loop (\`internal/tui/app.go\`) to call this function after every successful action.  
\- Display the recommendation in a "Recommended Next Step" box using Lipgloss (SuccessStyle border).  
\- This guides the user through the "perfect workflow" defined in \`.doplan/ai/rules/workflow.mdc\`.

Confirm when this feature is complete.

## **📝 TESTING PLAN**

### **1\. Testing for v0.0.18-beta**

* **Clean Installation Test:**  
  * Run doplan in an empty folder.  
  * Verify the new project wizard appears.  
  * Select "Cursor" as the IDE.  
  * Complete the GitHub setup wizard.  
  * **Verify doplan/ folders are NOT created yet.**  
  * Verify the TUI dashboard loads and is mostly empty (but not broken).  
* **Adoption Test:**  
  * Run doplan in an existing project (without .doplan).  
  * Verify the "Adopt Project?" wizard appears.  
  * Complete the wizard.  
  * **Verify CONTEXT.md and README.md are correctly fixed/generated.**  
  * **Verify the TUI dashboard loads and now shows data parsed from the project.**  
* **IDE Integration Tests:**  
  * Run the "Setup AI/IDE Integration" TUI action.  
  * **Test "Cursor":** Verify .cursor/agents/ and .cursor/rules/ are created.  
  * **Test "Copilot":** Verify .vscode/tasks.json and .vscode/prompts/ are created.  
  * **Test "Kiro":** Verify .doplan/guides/kiro\_setup.md is created.  
* **Documentation Test:**  
  * Manually inspect CONTEXT.md and README.md to confirm they match the new, correct structure.

### **2\. Testing for v0.0.19-beta**

* **TUI Actions Test:**  
  * Go through every new item in the TUI menu (Run Dev Server, Undo, Deploy, Security Scan, Fix Issues, Publish).  
  * Verify each one launches the correct TUI wizard or action.  
* **DPR & RAKD Test:**  
  * Run 🎨 Apply Design / DPR. Complete the questionnaire. Verify DPR.md, design-tokens.json, and .doplan/ai/rules/design\_rules.mdc are created.  
  * Run 🔑 Manage API Keys. Verify the RAKD TUI appears, detects services, and generates RAKD.md and SOPS/.  
* **Agent & Workflow Test (in Cursor Project):**  
  * Run /Plan. **Verify the @planner agent creates folders with the \#\#-slug-name format (e.g., doplan/01-authentication/01-email-login/).**  
  * Run /implement. Verify @coder agent runs.  
  * Tag @tester. **Verify it runs Playwright and saves screenshots to .doplan/artifacts/screenshots/01-authentication/01-email-login.png.**  
  * Tag @reviewer. Verify it comments on the test results.  
* **Workflow Guidance Test:**  
  * After /Plan completes, verify the TUI shows a "Recommended Next Step" box prompting you to implement a feature.