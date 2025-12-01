# /do

## Trigger
/do [<subcommand>]

## Examples
- "/do → Full ideation workflow"
- "/do feature → Add single feature idea"
- "/do now → Fast-track with detailed prompt/PRD"
- "/do i'm lucky → Get AI-suggested ideas"
- "/do meeting → Conduct discovery meeting"
- "/do refine → Refine project plan"


## Action
When user types /do or /do <subcommand>:

**Phase 1: Ideation (if no IDEA.md exists or subcommand is "feature" or "now" or "i'm lucky")**

1. **If no subcommand provided and IDEA.md doesn't exist**:
   - **Capture the idea**: Prompt user for their project idea
   - **Save to IDEA.md**: Write the idea to .do/system/IDEA.md
   - **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents
   - **Auto-proceed to meeting**: Automatically continue to Phase 2 (Meeting)

2. **Subcommand: feature** (or /do feature):
   - Capture a single feature idea
   - Append to existing IDEA.md or create new one
   - Response: "Feature idea captured! Continue with /do to conduct the meeting."

3. **Subcommand: now** (or /do now):
   - Fast-track mode: Ask for detailed prompt/PRD directly
   - Skip full meeting, generate documents immediately
   - Response: "Fast-track complete! Documents generated."

4. **Subcommand: i'm lucky** (or /do i'm lucky):
   - Get AI-suggested project ideas based on current trends
   - User selects from suggestions
   - Save selected idea to IDEA.md
   - Auto-proceed to meeting

**Phase 2: Meeting (automatic after ideation, or if subcommand is "meeting")**

0. **Initialize Meeting Session**:
   - Record meeting start time: Get current timestamp and store in meeting session (e.g., "2025-01-15 14:30:00")
   - Initialize progress tracking: Set current phase = 0, total phases = [based on speed selection]
   - Store meeting metadata: Start time, selected speed, project type, user experience level
   - Create meeting session file: .do/system/meeting_session.json with:
     {
       "start_time": "2025-01-15T14:30:00Z",
       "speed": "standard",
       "project_type": "saas",
       "user_experience": "intermediate",
       "total_phases": 4,
       "current_phase": 0,
       "phases": []
     }

2. **Determine Project Type**: Read .do/system/IDEA.md to understand the project type. Detect one of:
   - **Website**: Company website, agency portfolio, personal website, blog
   - **SaaS**: Software-as-a-Service, startup product, enterprise solution
   - **Mobile App**: Cross-platform mobile app (React Native, Flutter)
   - **iOS App**: iPhone/iPad native app (Swift, Objective-C)
   - **Android App**: Android native app (Java, Kotlin)
   - **Web App**: Progressive Web App (PWA), single-page app (SPA)
   - **Desktop App**: Cross-platform desktop app (Electron, Tauri)
   - **Windows App**: Windows native app (.NET, WPF, WinUI)
   - **macOS App**: macOS native app (Swift, SwiftUI)
   - **Linux App**: Linux native app (GTK, Qt)
   - **CLI Tool**: Command-line interface tool, terminal application
   - **Library/Package**: Code library, npm package, Python package, SDK
   - **Framework**: Development framework, plugin, extension
   - **API**: REST API, GraphQL API, backend service
   - **Microservice**: Microservices architecture, distributed system
   - **Game**: Video game, game engine, game tool
   - **Embedded/IoT**: Embedded system, IoT device, firmware
   - **Data Science/ML**: Machine learning, data analysis, AI project
   - **Cloud**: Cloud-native app, serverless function, cloud service
   - **DevOps**: DevOps tool, CI/CD pipeline, infrastructure as code
   - **Patch (Windows)**: Windows app patch, update, hotfix
   - **Patch (macOS)**: macOS app patch, update, hotfix
   - **Patch (Linux)**: Linux app patch, update, hotfix
   - **Patch (Web)**: Web app patch, update, hotfix
   - **Other**: Any other project type (uses general templates)

2. **Assess User Experience Level**: Ask user: "What's your experience level with development?"
   - **Option 1: Beginner/Non-Developer** - "I'm new to development or don't code"
   - **Option 2: Intermediate** - "I have some development experience"
   - **Option 3: Advanced** - "I'm an experienced developer"
   
   Based on experience level, determine available speed options:
   - **Beginner/Non-Developer**: Show only Quick Start and Standard. Recommend Quick Start.
   - **Intermediate**: Show Quick Start, Standard, and Comprehensive. Recommend Standard.
   - **Advanced**: Show all 4 options (Quick Start, Standard, Comprehensive, Deep Dive). Recommend based on project type:
     * Website/Agency/Personal/Mobile App/Web App/Desktop App/CLI/Library/Patch: Recommend Standard or Comprehensive
     * SaaS/Startup: Recommend Comprehensive or Deep Dive

3. **Present Speed Options** (Filtered based on experience level): Display available meeting speed options with clear descriptions:
   - **Quick Start** (Very Fast): ~5-10 minutes | Essential questions only | Best for: Simple projects, quick prototypes, beginners
   - **Standard** (Fast): ~15-20 minutes | Core phases with key questions | Best for: Most projects, MVPs, intermediate users
   - **Comprehensive** (Medium): ~30-45 minutes | All phases with detailed questions | Best for: Complex projects, established businesses, advanced users
   - **Deep Dive** (Long): ~60+ minutes | Full 6-phase interview with extensive probing | Best for: SaaS products, startups, enterprise solutions, advanced users

4. **User Selects Speed**: Wait for user to choose from available options. Show recommendation clearly (e.g., "⭐ Recommended: Quick Start").

5. **GitHub Repository Check**: 
   - Ask user: "Do you have a GitHub repository for this project? (yes/no)"
   - If yes: Ask for repository URL and verify it exists
   - Check if automated workflows (.github/workflows/) are set up correctly
   - Verify that automated committing and pushing workflows are configured
   - If missing or incorrect: Offer to set up GitHub Actions workflows for automated commits and pushes
   - Document GitHub repo info in the meeting summary

6. **Content Creation Needs Assessment** (Dynamic based on project type):
   - **For All Projects**: Ask "Do you need content creation for your project? (yes/no)"
   - If yes, present content types relevant to project type:
     - **Website/Agency/Personal**: 
       - App pages (landing, about, services, contact)
       - Legal pages (Privacy Policy, Terms of Service)
       - Social media content
       - Blog posts (if applicable)
       - SEO content
     - **SaaS/Startup**:
       - App pages (landing, features, pricing, about)
       - Legal pages (Privacy Policy, Terms of Service, User Agreement)
       - Social media content
       - Blog posts (content marketing)
       - Documentation (user guides, API docs)
       - Marketing content (ad copy, campaigns)
       - Email templates (welcome, newsletters)
       - SEO content (meta descriptions, keywords)
   - For each selected content type:
     - Ask: "How would you like to handle [content type]?"
       - Option 1: "Let LLM create 100% automatically"
       - Option 2: "I'll provide keywords, LLM creates content"
       - Option 3: "I'll provide initial draft, LLM refines and optimizes"
     - If keywords requested: Ask for relevant keywords, target audience, tone, and any specific requirements
     - Document content creation preferences in meeting summary
   - Note: Content will be created in .do/system/content/ organized by type, all SEO-ready

7. **Adaptive Phase Selection Based on Speed**:
   - **Quick Start**: Phases 01, 03 only (Vision & Outcomes, Experience & Tech) - Skip marketing, SEO, detailed ops
   - **Standard**: Phases 01, 02, 03, 06 (Vision, Audience, Experience, Delivery) - Skip detailed content/SEO and marketing
   - **Comprehensive**: Phases 01-06 but with condensed questions per phase
   - **Deep Dive**: All 6 phases with full detailed questioning

8. **Project Type Adaptation** (Dynamic based on detected type):
   - **Website/Agency/Personal**: Focus on design, content, basic functionality. Skip complex business models, growth strategies, scalability concerns.
   - **SaaS/Startup**: Emphasize business model, user acquisition, scalability, monetization, competitive analysis, technical architecture, growth strategies.
   - **Mobile App**: Focus on platform-specific requirements (iOS/Android), app store optimization, mobile UX patterns, device capabilities, offline functionality.
   - **Web App**: Emphasize browser compatibility, responsive design, performance optimization, PWA features, accessibility.
   - **Desktop App**: Focus on platform integration, system requirements, installation, updates, native features.
   - **CLI Tool**: Emphasize command structure, help system, configuration, output formatting, scripting integration.
   - **Library/Package**: Focus on API design, documentation, versioning, dependencies, distribution, backward compatibility.
   - **Patch/Update**: Emphasize change scope, testing requirements, backward compatibility, migration path, rollback strategy.
   - **Other**: Adapt questions based on project characteristics detected from IDEA.md.

9. **Load Phase Templates**: Read interview phases from .do/core/brainstorm/phase-*.md (01-06) based on selected speed. Questions should be dynamically generated based on:
   - Project type (website/agency/personal, SaaS/startup, mobile app, web app, desktop app, CLI, library, patch, etc.)
   - User experience level (beginner, intermediate, advanced)
   - Selected content creation needs
   - User's keyword preferences (if provided)
   - Meeting speed selected
   
   **Adapt question complexity based on experience level:**
   - **Beginner**: Use simple language, avoid technical jargon, explain concepts, focus on "what" not "how"
   - **Intermediate**: Mix of simple and technical terms, can discuss implementation details
   - **Advanced**: Technical terminology, architecture discussions, implementation strategies

10. **Conduct Phase-by-Phase Interview**:
   - **For each phase, before starting:**
     * Record phase start time (store in meeting_session.json)
   
   - Phase 01: Vision & Outcomes (Product Manager leads)
     * Phase start time recorded
     * Questions asked one at a time
     * When phase complete: Record phase end time, calculate duration, display progress summary
   
   - Phase 02: Audience & Differentiation (Product Manager + Design Manager) - Skip if Quick Start
     * Phase start time recorded
     * Questions asked one at a time
     * When phase complete: Record phase end time, calculate duration, display progress summary
   
   - Phase 03: Experience, UI/UX & Tech (Design Manager + Engineering Lead)
     * Phase start time recorded
     * Questions asked one at a time
     * When phase complete: Record phase end time, calculate duration, display progress summary
   
   - Phase 04: Content & SEO (Content Strategist + SEO Specialist) - Skip if Quick Start or Standard
     * Phase start time recorded
     * Questions asked one at a time
     * When phase complete: Record phase end time, calculate duration, display progress summary
   
   - Phase 05: Marketing & Growth (Marketing Manager) - Skip if Quick Start or Standard, emphasize for SaaS/Startup
     * Phase start time recorded
     * Questions asked one at a time
     * When phase complete: Record phase end time, calculate duration, display progress summary
   
   - Phase 06: Delivery, Ops & Risks (Engineering Lead + Project Orchestrator)
     * Phase start time recorded
     * Questions asked one at a time
     * When phase complete: Record phase end time, calculate duration, display progress summary
   
   - **After each phase completion, display simple progress summary:**
     Format:
     Phase [Number] - [Phase Name]
     
     Progress: [█████████████████] [percentage]%
     
     Phase Started : [start_time]
     Phase ended   : [end_time]
     
     * Calculate progress: (completed_phases / total_phases) * 100
     * Generate progress bar: Use █ for completed, ░ for remaining (20 characters total)
     * Display timing information
     * Wait for user confirmation before moving to next phase
   
   - **General phase behavior:**
     * Ask questions one at a time
     * Probe deeper when answers are vague
     * Adapt question depth based on selected speed

11. **Compile Summary**: Organize all answers by phase into a structured summary using format from .do/core/brainstorm/CONFIRMATION_TEMPLATE.md. Include:
   - Project type detected
   - User experience level
   - Meeting speed selected
   - GitHub repository information if provided
   - Content creation needs and preferences
   - Keywords provided (if any)
   - Content types selected

12. **Display Confirmation UI**: 
   - Present the summary in a well-formatted markdown display with clear sections, checkmarks (✅), blockquotes for longer answers, and visual separators
   - Include a "Review & Confirm" section with 4 clear options: (1) Save it, (2) Revise a phase, (3) Add information, (4) Start over
   - Wait for explicit user confirmation - DO NOT save until user explicitly confirms

13. **Handle User Response**:
   - If confirmed: Proceed to save
   - If revision requested: Re-ask questions for specified phase(s), update summary, show again
   - If addition requested: Add information to appropriate phase, show updated summary
   - If restart requested: Confirm intent, then restart from speed selection if confirmed

14. **Save to BRAINSTORM.md**: Once explicitly confirmed, write the approved summary (organized by phase) to .do/system/BRAINSTORM.md using structure from .do/core/brainstorm/TEMPLATE_BRAINSTORM.md. Include:
   - Project type detected
   - User experience level
   - Meeting speed selected
   - GitHub repository information
   - Content creation needs and preferences
   - Keywords and content requirements
   - Meeting timing summary (start time, end time, total duration, phase durations)

15. **Display Final Meeting Summary** (after last phase):
   Format:
   Phase [Last Phase Number] - [Phase Name]
   
   Progress: [█████████████████] 100%
   
   Phase Started : [start_time]
   Phase ended   : [end_time]
   
   Meeting Started: [meeting_start_time]
   Meeting Ended: [meeting_end_time]
   Total Duration: [total_duration]

16. **Update State**: Set .do/system/history/active_state.json phase to "brainstorm". Also save user experience level and meeting session data for future reference. Save meeting_session.json with complete timing data.

17. **Response**: "✅ Meeting complete! Summary saved to BRAINSTORM.md. Content structure created in .do/system/content/. Type /plan to generate PRD, Architecture, Design System, and execution plan."

**Phase 3: Refining (optional, if subcommand is "refine")**

1. **Review BRAINSTORM.md**: Load the meeting summary
2. **Suggest Refinements**: AI analyzes and suggests improvements
3. **User Reviews**: Present suggestions for user approval
4. **Update Documents**: Apply approved refinements to IDEA.md and BRAINSTORM.md
5. **Response**: "Refinements applied! Your project plan has been updated."

## Agent Involvement
- **Product Manager**
- **Engineering Lead**
- **Design Manager**
- **Content Strategist**
- **SEO Specialist**
- **Marketing Manager**
- **Project Orchestrator**

## Files Read
- ".do/system/IDEA.md"
- ".do/core/brainstorm/phase-*.md"
- ".do/core/brainstorm/CONFIRMATION_TEMPLATE.md"
- ".do/core/brainstorm/TEMPLATE_BRAINSTORM.md"

## Files Modified
- ".do/system/BRAINSTORM.md"
- ".do/system/content/"
- ".do/system/history/active_state.json"

## Requirements
|

