package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Command represents a slash command in the AI agency system
type Command struct {
	// Basic Information
	Name        string // Command name (e.g., "tell", "build")
	Trigger     string // Trigger pattern (e.g., "/tell or /tell <idea>")
	Description string // Brief description
	Action      string // Detailed action description
	Category    string // Command category (e.g., "start", "plan", "develop")

	// Agent Involvement
	AgentInvolvement []string // List of agents involved

	// Files
	FilesRead     []string // Files read by this command
	FilesModified []string // Files modified by this command

	// Additional Information
	Examples         []string // Example usage
	GitHubAutomation string   // GitHub automation details (if applicable)
	Requirements     string   // Requirements section (if applicable)
	Notes            string   // Notes section (if applicable)
	Customize        string   // Customize section (if applicable)
	Options          string   // Options section (if applicable)
	OfflineSafety    string   // Offline Safety section (if applicable)
}

// GetAllCommands returns all core and squad commands
func GetAllCommands() []Command {
	return []Command{
		// Core Commands - Essential daily workflow
		{
			Name:        "hello",
			Category:    "core",
			Trigger:     "/hello [<subcommand>]",
			Description: "Welcome, tutorial, and command introductions",
			Action: `When user types /hello or /hello <subcommand>:

1. **If no subcommand provided** (or subcommand is "goplan"):
   - **Check if first time**: Read .do/system/user_profile.json. If it exists and "first_hello_completed" is true, show a welcome back message instead. Otherwise, proceed with first-time tutorial.

2. **Warm Greeting**: Display:
   "Hello! 👋 I'm DoPlan, your AI development partner. I'm here to guide you on how to turn your idea into a real product. It's great to meet you! Let's get to know each other a bit."

3. **Personal Introduction**:
   - Ask: "What should I call you?" → Store as user_name
   - Ask: "What's your experience with development?"
     * Option 1: "I'm completely new to development"
     * Option 2: "I have some basic experience"
     * Option 3: "I'm an intermediate developer"
     * Option 4: "I'm an experienced developer"
     → Store as user_experience_level
   - Respond: "Nice to meet you, [Name]! I'll be here to help you every step of the way."

4. **Tutorial - System Overview**:
   - Explain what DoPlan is (planning, document generation, content creation, development guidance, progress tracking)
   - Explain the sub-agent system (18 specialized agents)
   - Display the agent hierarchy using the text-based tree format from docs/AGENT_HIERARCHY_CHAT_PREVIEW.md (Option 2). The hierarchy should be displayed as:
     Project Orchestrator (CEO/Manager) 👔
     │
     ├── Product Manager 📋
     │
     ├── Engineering Lead 💻
     │   ├── System Architect 🏗️
     │   ├── Frontend Lead 🎨
     │   ├── Backend Lead ⚙️
     │   ├── DevOps Engineer 🚀
     │   ├── Security Lead 🔒
     │   └── Performance Engineer ⚡
     │
     ├── Design & UX Manager 🎨
     │   └── UI/UX Designer ✨
     │
     ├── QA & Reliability Manager ✅
     │   └── QA Engineer 🧪
     │
     ├── Release & Growth Manager 📈
     │   ├── Release Captain 🚢
     │   └── Growth Coach 📊
     │
     └── Documentation Lead 📝
         └── Documentation Writer ✍️
   - Explain accuracy & reliability (high accuracy, context-aware, transparent, learnable)

5. **Development Support Option**:
   - Ask: "[Name], I can support you throughout your development journey in two ways:
     Option 1: **Guided Mode** (Recommended for beginners)
     - I'll send you progress messages after each completed feature
     - Tell you exactly what you accomplished
     - Guide you on what to do next
     - Sometimes provide ready-to-copy prompts for the next step
     - Celebrate your milestones with you
     
     Option 2: **Independent Mode**
     - You work at your own pace
     - I'm here when you need me (just ask!)
     - Less frequent check-ins
     
     Which would you prefer? (guided/independent)"
   - Store as development_support_mode
   - Respond appropriately based on selection

6. **First-Time Tutorial Walkthrough**:
   - Say: "Now, let's do a quick walkthrough together. This is just for learning - we won't create anything real yet. Ready? Let's go! 🚀"
   - Explain each command in order:
     * /tell - Capture project idea
     * /meeting - Discovery meeting with adaptive speed options
     * /write - Generate PRD, Architecture, Design System (show example output preview)
     * /content - Generate SEO-ready content
     * /plan - Create execution plan and task list
     * /build - Start coding with step-by-step guidance
   - Ask: "Does this workflow make sense, [Name]? Would you like to try a 'test drive' of one command to see how it works? This is just practice - we won't save anything. (yes/no)"

7. **Interactive Test Drive** (if user says yes):
   - Offer to test drive: /tell, /meeting, /write, /content
   - Guide user through interactive practice of selected command(s)
   - Explain what would happen in real mode vs. practice mode
   - Ask if they want to try another command or finish

8. **Personalized Tips** (based on experience level):
   - **Beginner**: Focus on guided mode, step-by-step instructions, don't hesitate to ask questions
   - **Intermediate**: Balance of guidance and independence, leverage agent expertise
   - **Advanced**: Deep technical details, comprehensive guidance, trust agent recommendations

9. **Save Reference Materials**:
   - Create .do/system/user_profile.json with: user_name, user_experience_level, development_support_mode, first_hello_completed=true, first_hello_date, tutorial_completed=true, test_drive_completed, quick_reference_saved=true
   - Create .do/system/QUICK_REFERENCE.md with command cheat sheet
   - Create docs/references/QUICK_REFERENCE.md (same content)
   - Create docs/overview/AGENT_HIERARCHY.md with agent hierarchy diagram (Option 2 format)
   - Create docs/references/COMMAND_EXAMPLES.md with example outputs
   - Create docs/tutorials/TUTORIAL_NOTES.md with tutorial summary

10. **Encouragement & Ready to Start**:
    - Say: "Well done, [Name]! 🎉 You're now ready to start developing. Here's what you've learned:
      ✅ How DoPlan works
      ✅ Our agent system and expertise (with visual hierarchy)
      ✅ The command workflow
      ✅ Your development support mode
      
      All reference materials have been saved to:
      📁 .do/system/QUICK_REFERENCE.md - Your personal cheat sheet
      📁 docs/references/QUICK_REFERENCE.md - Project documentation
      📁 docs/overview/AGENT_HIERARCHY.md - Agent structure diagram
      📁 docs/references/COMMAND_EXAMPLES.md - Example outputs and usage
      📁 docs/tutorials/TUTORIAL_NOTES.md - Tutorial summary
      
      Ready to start? Type /tell to capture your project idea!"

2. **Subcommand: meeting** (or /hello meeting):
   - **Storytelling Approach**: Start with a scenario:
     "Imagine you're about to build your dream project. You have the idea, but you're not sure where to start. That's where /meeting comes in - it's like having a smart product manager interview you to understand exactly what you want to build."
   
   - **Interactive Discovery Journey**:
     * Step 1: "First, we'll figure out what kind of project you're building. Are you building a simple website? A complex SaaS? A mobile app? We adapt our questions based on your answer."
     * Step 2: "Then, we'll assess your experience level. Are you new to development? We'll keep it simple. Are you experienced? We'll dive deep into technical details."
     * Step 3: "Next, you'll choose your meeting speed:
       🚀 Quick Start (5-10 min) - Perfect for simple projects or when you're in a hurry
       ⚡ Standard (15-20 min) - Balanced depth for most projects
       📋 Comprehensive (30-45 min) - Detailed planning for complex projects
       🔍 Deep Dive (60+ min) - Complete exploration for enterprise solutions"
     * Step 4: "We'll verify your GitHub setup - because good version control is the foundation of professional development."
     * Step 5: "We'll ask about content needs - do you need landing pages? Legal pages? Blog posts? We'll help you plan it all."
   
   - **Visual Example**: Show a sample flow:
     "Here's what a meeting might look like:\n" +
     "Code block:\n" +
     "You: /meeting\n" +
     "DoPlan: What's your experience level?\n" +
     "You: Intermediate\n" +
     "DoPlan: Great! I'll show you Quick Start, Standard, and Comprehensive options.\n" +
     "        Which speed would you like? [Recommended: Standard]\n" +
     "You: Standard\n" +
     "DoPlan: Perfect! Let's start with your project vision...\n" +
     "[Interactive Q&A session]\n" +
     "DoPlan: Meeting complete! Your BRAINSTORM.md has been saved."
   
   - **Why It's Special**: "Unlike static forms, /meeting adapts in real-time. It asks follow-up questions, probes deeper when needed, and skips irrelevant sections. It's like having a conversation with an expert product manager who knows exactly what to ask."
   
   - **Call to Action**: "Ready to discover your project? Type /meeting and let's start the journey! 🎯"

3. **Subcommand: plan** (or /hello plan):
   - **Visual Tour Approach**: "Let me take you on a tour of how DoPlan organizes your project. Think of it like building a house - you need a solid foundation, then walls, then the roof. That's exactly how we structure your project!"
   
   - **Folder Structure Walkthrough** (with ASCII art):
     "Project structure:\n" +
     ".do/plan/\n" +
     "├── TASKS.md                    ← Master task list (your project roadmap)\n" +
     "│\n" +
     "├── 01-Foundation/              ← Phase 1: Core infrastructure\n" +
     "│   ├── feature-auth/\n" +
     "│   │   ├── design.md           ← How it looks\n" +
     "│   │   ├── plan.md             ← How it works\n" +
     "│   │   ├── tasks.md            ← What to build\n" +
     "│   │   ├── prompts.md          ← AI prompts for this feature\n" +
     "│   │   └── github.md           ← Git workflow for this feature\n" +
     "│   ├── feature-database/\n" +
     "│   └── _contracts/             ← Shared data schemas\n" +
     "│\n" +
     "├── 02-Core/                    ← Phase 2: Main features\n" +
     "│   └── ...\n" +
     "│\n" +
     "└── 03-Enhancement/              ← Phase 3: Polish & optimization\n" +
     "    └── ..."
   
   - **Planning Strategies Explained**:
     * "You have flexibility! Choose your planning style:
       📦 Full Planning: Generate everything at once (great for small projects)
       🔄 Phase-by-Phase: Plan incrementally (perfect for large projects)
       🎯 Selective: Update specific phases (when requirements change)"
   
   - **Real-World Example**: "Imagine you're building an e-commerce site:
     - Phase 1 (Foundation): User authentication, database setup
     - Phase 2 (Core): Product catalog, shopping cart
     - Phase 3 (Enhancement): Reviews, recommendations, analytics
     
     Each phase has its own tasks, and each feature has detailed documentation!"
   
   - **Smart Features**:
     * "Before planning, we check if your documents are approved - we want to make sure you're ready!"
     * "We generate templates automatically - no more starting from scratch!"
     * "Each feature folder is self-contained - everything you need is right there!"
   
   - **Call to Action**: "Ready to see your project organized? Type /plan and watch the magic happen! ✨"

4. **Subcommand: build** (or /hello build):
   - **Journey Metaphor**: "Think of /build as your development journey. Every task is a destination, and we'll guide you there step by step. Let me show you how smooth this journey can be!"
   
   - **Step-by-Step Adventure**:
     "Development flow:\n" +
     "Step 1: Task Selection\n" +
     "   You: /build\n" +
     "   DoPlan: Found next task: 2.1 - User Authentication\n" +
     "   OR\n" +
     "   You: /build 2.1\n" +
     "   DoPlan: Starting task 2.1 - User Authentication\n" +
     "\n" +
     "Step 2: Project Bootstrap (First Time Only)\n" +
     "   DoPlan: Setting up your project structure...\n" +
     "   [Creates Next.js boilerplate, installs dependencies]\n" +
     "   DoPlan: Project ready!\n" +
     "\n" +
     "Step 3: Git Safety Check\n" +
     "   DoPlan: Checking Git status...\n" +
     "   [Ensures clean working tree]\n" +
     "   DoPlan: Ready to code!\n" +
     "\n" +
     "Step 4: Branch Creation\n" +
     "   DoPlan: Creating branch: task/2.1\n" +
     "   [Automatically creates and checks out branch]\n" +
     "   DoPlan: Working on task/2.1\n" +
     "\n" +
     "Step 5: Agent Activation\n" +
     "   DoPlan: Activating agents: Frontend Lead, Backend Lead, Security Lead\n" +
     "   [Agents ready to help with the task]\n" +
     "\n" +
     "Step 6: Development\n" +
     "   [You code with AI assistance]\n" +
     "   [Agents provide suggestions and guidance]\n" +
     "\n" +
     "Step 7: Auto-Completion Detection\n" +
     "   DoPlan: Task appears complete! Summary:\n" +
     "            All requirements met\n" +
     "            Code implemented\n" +
     "            Tests passing\n" +
     "            \n" +
     "            Mark as done? (yes/no)\n" +
     "   You: yes\n" +
     "   DoPlan: Task marked complete! Committing and pushing...\n" +
     "   [Auto-commit with conventional format]\n" +
     "   [Auto-push to remote]\n" +
     "   DoPlan: Task 2.1 complete! Ready for next task."
   
   - **Why It's Powerful**:
     * "No more manual branch management - we handle it!"
     * "No more forgetting to commit - we do it automatically!"
     * "No more wondering what's next - we track everything!"
     * "No more losing progress - we snapshot your state!"
   
   - **State Management Magic**: "Every time you start or finish a task, we take a snapshot. Want to see what changed? Type /status. Need to rollback? We've got you covered!"
   
   - **Real Example**: "Here's what happens in practice:\n" +
     "Morning: /build → Start task 3.2\n" +
     "[Work on feature]\n" +
     "Afternoon: Agent detects completion\n" +
     "[Auto-commit, auto-push]\n" +
     "Evening: /build → Start next task automatically\n" +
     "It's that seamless!"
   
   - **Call to Action**: "Ready to start building? Type /build and let's turn your plan into code! 🚀"

5. **Subcommand: github** (or /hello github):
   - **Storytelling Approach**: "Let me tell you a story. Once upon a time, developers spent hours managing Git branches, writing commit messages, and setting up CI/CD. Then DoPlan came along and automated it all! 🎉"
   
   - **The Problem We Solve**:
     "❌ Before DoPlan:
     - Manual branch creation (easy to make mistakes)
     - Inconsistent commit messages (hard to track changes)
     - Forgetting to push (lost work)
     - Complex CI/CD setup (hours of configuration)
     - Manual release management (error-prone)
     
     ✅ With DoPlan:
     - Automatic branch creation (task/2.1, feature/auth, etc.)
     - Conventional commits (feat:, fix:, docs:, etc.)
     - Auto-push after completion (never lose work)
     - One-command CI/CD setup (/github ci)
     - Automated releases (/github release)"
   
   - **Why GitHub Matters** (with examples):
     "GitHub isn't just code storage - it's your project's heartbeat:
     💓 Version Control: Track every change, rollback when needed
     🤝 Collaboration: Work with teams seamlessly
     🔍 Code Review: Catch issues before production
     🚀 CI/CD: Automate testing and deployment
     📊 Insights: Track progress and KPIs
     🏷️ Releases: Professional version management"
   
   - **DoPlan's GitHub Automation** (Interactive Demo):
     "Scenario: You just finished a task\n" +
     "\n" +
     "Traditional Way:\n" +
     "1. git checkout -b feature/new-feature\n" +
     "2. [work on code]\n" +
     "3. git add .\n" +
     "4. git commit -m \"added new feature\"  ← Vague message\n" +
     "5. git push origin feature/new-feature\n" +
     "6. [Create PR manually]\n" +
     "7. [Set up CI/CD manually]\n" +
     "\n" +
     "DoPlan Way:\n" +
     "1. /build 2.1\n" +
     "2. [work on code]\n" +
     "3. Agent: Task complete! Mark as done?\n" +
     "4. You: yes\n" +
     "5. DoPlan automatically:\n" +
     "   Creates branch: task/2.1\n" +
     "   Commits: feat(task-2.1): implement user authentication\n" +
     "   Pushes to remote\n" +
     "   Updates CHANGELOG.md\n" +
     "   Suggests PR creation"
   
   - **GitHub Command Features**:
     "The /github command is your repository control center:
     
     📊 /github info
        - Syncs README with project KPIs
        - Updates repository metadata
        - Shows project health
     
     🐛 /github issue \"Title\" \"Description\"
        - Creates issues with proper formatting
        - Links to project documentation
     
     🎯 /github milestone \"Name\"
        - Manages project milestones
        - Tracks progress
     
     ⚙️ /github ci
        - Generates CI/CD workflows
        - Configures branch-based testing
        - One command setup!
     
     🚢 /github release
        - Plans releases
        - Generates release notes
        - Manages versioning"
   
   - **Branching Strategy** (Visual):
     "Branch structure:\n" +
     "main/master\n" +
     "├── task/1.1          ← Feature branches\n" +
     "├── task/1.2\n" +
     "├── feature/auth      ← Larger features\n" +
     "├── bugfix/login      ← Bug fixes\n" +
     "└── hotfix/security   ← Urgent fixes\n" +
     "\n" +
     "Each branch has its own CI/CD pipeline!"
   
   - **Best Practices We Enforce**:
     "We follow industry best practices:
     ✅ Conventional Commits (feat:, fix:, docs:, etc.)
     ✅ Branch naming conventions (task/, feature/, bugfix/, hotfix/)
     ✅ Clean working tree before branching
     ✅ Automatic state snapshots
     ✅ CI/CD for every branch type"
   
   - **Real Impact**: "Here's what this means for you:
     - ⏱️ Save 2-3 hours per week on Git management
     - 🎯 Never lose work (auto-push)
     - 📈 Professional commit history
     - 🚀 CI/CD in minutes, not hours
     - 🏆 Industry-standard workflow"
   
   - **Call to Action**: "Ready to automate your GitHub workflow? Type /github to see what's possible! 🌟"`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".do/system/user_profile.json",
				"docs/AGENT_HIERARCHY_CHAT_PREVIEW.md",
			},
			FilesModified: []string{
				".do/system/user_profile.json",
				".do/system/QUICK_REFERENCE.md",
				"docs/references/QUICK_REFERENCE.md",
				"docs/overview/AGENT_HIERARCHY.md",
				"docs/references/COMMAND_EXAMPLES.md",
				"docs/tutorials/TUTORIAL_NOTES.md",
			},
			Examples: []string{
				"/hello → Full tutorial (same as /hello goplan)",
				"/hello goplan → Full tutorial",
				"/hello meeting → Introduce /meeting command",
				"/hello plan → Introduce plan structure",
				"/hello build → Introduce build process",
				"/hello github → Introduce GitHub workflow",
			},
		},
		{
			Name:        "access",
			Category:    "tools",
			Trigger:     "/access [<scope>]",
			Description: "Beginner-friendly patch to fix .do/ and docs/ permissions",
			Action: `When user types /access or /access <scope>:

1. **Explain Purpose**:
   - "This patch makes sure DoPlan can write reference docs and state files. It only creates missing folders/files and fixes permissions."

2. **Run Helper**:
   - Execute npx --yes @doplan-dev/cli goplan access <scope> (defaults to all).
   - Scope options:
     * all - patch .do/system, .do/plan, and docs/
     * .do/system or system
     * .do/plan or plan
     * docs

3. **Report Results**:
   - Show which folders/files were created or updated.
   - If script exits non-zero, surface stderr and prompt the user to re-run with DEBUG=1 for verbose output.

4. **Next Steps**:
   - Encourage the user to run /hello again once the patch succeeds.
   - Mention that the patch is safe to re-run at any time.`,
			AgentInvolvement: []string{
				"Project Orchestrator",
				"Documentation Lead",
			},
			FilesModified: []string{
				".do/system/**",
				".do/plan/**",
				"docs/**",
			},
			Examples: []string{
				"/access",
				"/access docs",
				"/access .do/system",
			},
			Notes: "Perfect for beginners who see the permission warning during /hello. It never overwrites existing content—only creates missing folders/files and ensures they're writable.",
		},
		{
			Name:        "tell",
			Category:    "core",
			Trigger:     "/tell or /tell <idea>",
			Description: "Capture project idea",
			Action: `When user types /tell or /tell <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .do/system/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /meeting to start the discovery meeting with the team."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
				"Product Manager",
			},
			FilesRead: []string{},
			FilesModified: []string{
				".do/system/IDEA.md",
				".do/system/history/active_state.json",
			},
			Examples: []string{
				"/tell",
				"/tell Build a todo app",
			},
		},
		{
			Name:        "meeting",
			Category:    "core",
			Trigger:     "/meeting",
			Description: "Adaptive discovery meeting with speed options",
			Action: `When user types /meeting:

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
     * Generate progress bar: Use RenderProgressBar() function or format: [█████████████████] [percentage]% (30 characters total)
     * Display timing information including phase duration
     * Use DisplayPhaseProgress() function for consistent formatting
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
   
   Progress: [██████████████████████████████] 100%
   
   Phase Started : [start_time]
   Phase ended   : [end_time]
   Phase Duration: [duration]
   
   * Use DisplayPhaseProgress() function for consistent formatting
   * Show all phase durations and total meeting time
   
   Meeting Started: [meeting_start_time]
   Meeting Ended: [meeting_end_time]
   Total Duration: [total_duration]

16. **Update State**: Set .do/system/history/active_state.json phase to "brainstorm". Also save user experience level and meeting session data for future reference. Save meeting_session.json with complete timing data.

17. **Response**: "✅ Meeting complete! Summary saved to BRAINSTORM.md. Content structure created in .do/system/content/. Type /write to generate PRD, Architecture, and Design System. Type /content to start generating content."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"Design Manager",
				"Content Strategist",
				"SEO Specialist",
				"Marketing Manager",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".do/system/IDEA.md",
				".do/core/brainstorm/phase-*.md",
				".do/core/brainstorm/CONFIRMATION_TEMPLATE.md",
				".do/core/brainstorm/TEMPLATE_BRAINSTORM.md",
			},
			FilesModified: []string{
				".do/system/BRAINSTORM.md",
				".do/system/content/",
				".do/system/history/active_state.json",
			},
			Requirements: "- Phase templates should exist in .do/core/brainstorm/\n- Interview should be conversational, one question at a time\n- Questions should be dynamically generated based on project type, not pre-made\n- Project type detection should support: website/agency/personal, SaaS/startup, mobile app, web app, desktop app, CLI, library/package, patch/update, and other types\n- User experience level assessment is mandatory and affects available speed options\n- Speed options must be filtered based on experience level:\n  * Beginner/Non-Developer: Only Quick Start and Standard (recommend Quick Start)\n  * Intermediate: Quick Start, Standard, Comprehensive (recommend Standard)\n  * Advanced: All 4 options (recommend based on project type)\n- Question complexity should adapt to user experience level (simple for beginners, technical for advanced)\n- Content creation questions should adapt to project type\n- Summary must be displayed in formatted confirmation UI before saving\n- User must explicitly confirm before any files are written\n- Use CONFIRMATION_TEMPLATE.md format for displaying summary\n- Use TEMPLATE_BRAINSTORM.md format for final saved document\n- Adapt phases and questions based on selected speed, project type, and experience level\n- Always check GitHub repository and workflow status\n- Content folder structure (.do/system/content/) is created automatically\n- All content should be SEO-ready with keyword optimization\n- Save user experience level in active_state.json for future reference",
			Examples: []string{
				"/meeting",
			},
		},
		{
			Name:        "write",
			Category:    "core",
			Trigger:     "/write [<subcommand>] [<args>]",
			Description: "Generate documents & content, or edit existing documents",
			Action: `When user types /write or /write <subcommand>:

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

6. **Response**: "Documents/content generated! Review files in .do/system/ or .do/system/content/."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"System Architect",
				"Design & UX Manager",
				"UI/UX Designer",
				"Content Strategist",
				"SEO Specialist",
				"Documentation Writer",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".do/system/IDEA.md",
				".do/system/BRAINSTORM.md",
				".do/system/*.md",
			},
			FilesModified: []string{
				".do/system/PRD.md",
				".do/system/ARCHITECTURE.md",
				".do/system/DESIGN_SYSTEM.md",
				".do/system/*.md",
				".do/system/content/**",
				".do/system/history/active_state.json",
			},
			Examples: []string{
				"/write → Generate all planning docs (first time) or show menu",
				"/write plan → Show planning document options",
				"/write content → Show content type options",
				"/write change prd Add dark mode → Edit PRD",
				"/write edit architecture Use PostgreSQL → Edit ARCHITECTURE",
				"/write prd → Regenerate PRD only",
				"/write legal → Generate legal pages",
			},
		},
		{
			Name:        "plan",
			Category:    "core",
			Trigger:     "/plan [<subcommand>] [<args>]",
			Description: "Generate execution plan, scaffold phases, and manage tasks",
			Action: `When user types /plan or /plan <subcommand>:

1. **If no subcommand provided** (or subcommand is "everything"):
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
     * If no → "Please review and approve documents first. Type /write to regenerate if needed."
   - **If Approved or User Confirmed**:
     * Synthesize Execution Tasks: Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md to generate .do/plan/TASKS.md
     * Parse TASKS.md: Use the generated tasks to determine phases and features
     * Scaffold Phase Folders: Create phase directories (e.g., 01-Foundation) in .do/plan/
     * Generate Feature Folders: For each task, create feature folders with templates (design.md, plan.md, tasks.md, prompts.md, github.md)
     * Create Contracts Directory: Add _contracts/ folder in each phase for shared schemas
     * Update State: Update .do/system/history/active_state.json to reference the new hierarchy and set phase to "tasks"
     * Mark plan as generated in active_state.json
   - **Response**: "Execution plan generated! TASKS.md and phase folders created in .do/plan/. Type /build to start implementing."

2. **Subcommand: phases** (or /plan phases):
   - Plan project phase by phase:
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * Generate phase structure but don't create all tasks at once
     * Create phase folders (01-Foundation, 02-Core, etc.) without feature folders yet
     * Set up phase-by-phase planning mode in active_state.json
     * After user finishes a feature with /build, system will ask: "Ready to plan the next phase? Type /plan next to continue."
   - **Response**: "Phase-by-phase planning mode activated! Phase folders created. Complete features with /build, then type /plan next to plan the next phase."

3. **Subcommand: next** (or /plan next):
   - Planning next phase:
     * Read active_state.json to determine current phase
     * Find the next unplanned phase
     * Check if previous phase features are complete (if applicable)
     * Generate tasks for the next phase only
     * Create feature folders for this phase with templates
     * Update active_state.json with current phase
     * **Response**: "Next phase planned! Phase [X] tasks and feature folders created. Type /build to start implementing."

4. **Subcommand: phase {no} tasks** (or /plan phase <number> tasks):
   - Create tasks.md for a specific phase:
     * Parse phase number from command (e.g., /plan phase 1 tasks)
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * Generate tasks for this phase only
     * Create tasks.md file inside .do/plan/[Phase-Folder]/ (e.g., .do/plan/01-Foundation/tasks.md)
     * Include all feature tasks for this phase
     * **Response**: "Phase [X] tasks created! tasks.md saved in .do/plan/[Phase-Folder]/tasks.md"

5. **Subcommand: phases tasks** (or /plan phases tasks):
   - Create tasks.md for all phases:
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * For each phase folder in .do/plan/:
       - Generate tasks for that phase
       - Create tasks.md file inside that phase folder
       - Include all feature tasks for that phase
     * **Response**: "All phase tasks created! tasks.md files saved in each phase folder."

6. **Subcommand: all tasks** (or /plan all tasks):
   - **This is the ONLY way to create the main TASKS.md file**:
     * Read .do/system/PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
     * Generate ALL tasks for the entire project
     * Create .do/plan/TASKS.md with all project tasks organized by phase
     * This is the master tasks file that contains everything
     * **Response**: "Main TASKS.md created! All project tasks saved in .do/plan/TASKS.md"`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".do/system/PRD.md",
				".do/system/ARCHITECTURE.md",
				".do/system/DESIGN_SYSTEM.md",
				".do/system/history/active_state.json",
				".do/plan/TASKS.md",
			},
			FilesModified: []string{
				".do/plan/TASKS.md",
				".do/plan/[Phase-Folders]/",
				".do/plan/[Phase-Folders]/[Feature-Folders]/",
				".do/plan/[Phase-Folders]/_contracts/",
				".do/system/history/active_state.json",
			},
			Requirements: "- Task generation templates live in `.do/core`\n- If documents are not approved, user must confirm before proceeding\n- /plan all tasks is the only way to create the main .do/plan/TASKS.md file",
			Examples: []string{
				"/plan → Generate execution plan and tasks (same as /plan everything)",
				"/plan everything → Generate full execution plan",
				"/plan phases → Plan project phase by phase",
				"/plan next → Planning next phase",
				"/plan phase 1 tasks → Create tasks.md for phase 1",
				"/plan phases tasks → Create tasks.md for all phases",
				"/plan all tasks → Create main TASKS.md with all project tasks",
			},
		},
		{
			Name:        "build",
			Category:    "core",
			Trigger:     "/build [<task_id>]",
			Description: "Start coding next task",
			Action:      "When user types /build or /build <task_id>:\n\n1. **Determine Task**:\n   - If task_id provided, load that task\n   - Otherwise, find next uncompleted task from TASKS.md\n2. **Bootstrap Boilerplate (first run only)**:\n   - If the project is still plan-only (no package.json / src/), prompt the user to scaffold code with their preferred stack tool (e.g., `npx create-next-app`, `go mod init`, etc.)\n   - DoPlan no longer ships the legacy `scripts/boilerplate` helper, so projects must bring or generate their own starter code\n   - Skip once code already exists for this project/stack\n3. **Check Git Status**:\n   - Verify working tree is clean (no uncommitted changes)\n   - If dirty, warn user and block until clean\n4. **Create/Checkout Task Branch**:\n   - Create or checkout branch `task/[ID]` manually (e.g., `git checkout -b task/5.2`)\n   - Store branch name in `active_branch` field of `.do/system/history/active_state.json`\n5. **Load Task Context**: Read task details, dependencies, and related code\n6. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)\n7. **Start Implementation**: Begin coding the task with full context\n8. **Update State**: Set `active_task` and `active_branch` in `.do/system/history/active_state.json`\n9. **Snapshot State**: Immediately log the new state with `go run scripts/statehistory/main.go snapshot --reason \"build [ID]\" --label build`\n10. **After Task Implementation** (when agent detects completion):\n   - Agent analyzes code changes, tests, requirements\n   - Agent checks if task criteria are met\n   - If complete:\n     - Agent asks: \"Task [ID] appears complete. Summary:\n       ✅ All requirements met\n       ✅ Code implemented\n       ✅ Tests passing (if applicable)\n       \n       Mark as done? (yes/no)\"\n     - If user says yes:\n       - Mark task complete in TASKS.md\n       - Update active_state.json\n       - Auto-commit and push\n       - Response: \"Task marked complete! Changes committed and pushed.\"\n     - If user says no:\n       - Continue working on task\n11. **Response**: \"Building task [ID]: [Description] on branch [branch_name]. Focus on this task only.\"",
			AgentInvolvement: []string{
				"Engineering Lead",
				"Relevant Team Leads",
				"Project Orchestrator",
				"QA Engineer",
			},
			FilesRead: []string{
				".do/plan/TASKS.md",
				".do/system/history/active_state.json",
			},
			FilesModified: []string{
				".do/system/history/active_state.json (active_task and active_branch updated)",
				".do/system/history/state-*.json (automatic snapshot for audit/rollback)",
				".do/plan/TASKS.md (task marked complete if auto-detected)",
				"Git: New branch created/checked out (task/[ID])",
				"src/** (code files created/modified)",
			},
			Examples: []string{
				"/build → Start next uncompleted task",
				"/build 1.2 → Start specific task 1.2",
				"/build 3 → Start task 3",
			},
			GitHubAutomation: `After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
   - Update docs/history/CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md`,
		},
		{
			Name:        "status",
			Category:    "core",
			Trigger:     "/status [<subcommand>]",
			Description: "Show project progress and generate reports",
			Action:      "When user types /status or /status <subcommand>:\n\n1. **If no subcommand provided** (default: show progress):\n   - Read TASKS.md: Load all tasks\n   - Read active_state.json: Get completed tasks and current phase\n   - Run Progress Tool: Execute `go run scripts/progress/main.go --root <project>`\n     This parses `.do/plan/TASKS.md`, `.do/system/history/active_state.json`, and `.do/system/history/` to compute stats and state deltas.\n   - Calculate Progress:\n     * Total tasks\n     * Completed tasks\n     * In progress tasks\n     * Percentage complete\n   - Display Progress: Show formatted progress report:\n     * Phase: [current phase]\n     * Tasks: X/Y completed (Z%)\n     * Current task: [active task]\n     * Next up: [next task]\n     * State Delta: summarize what changed between the last two snapshots (phase/task/branch/completed)\n   - Response: Display progress summary with the state delta footer\n\n2. **Subcommand: report** (or /status report):\n   - Select Project:\n     * Default: current workspace (.)\n     * Optional: `/status report test/qr-generator/test-no01`\n   - Generate Metadata:\n     * Runs `go run scripts/scanreport/main.go --project <path>`\n     * Parses `.do/reports/SCAN_REPORT_*.md`\n     * Creates/updates matching JSON files with structured metadata (scan date, project, executive summary, findings, next actions, summary hash)\n   - Compute Diff:\n     * When >=2 reports exist, compares the newest vs previous\n     * Builds `SCAN_DIFF_<date>.md` highlighting added/removed bullets in Executive Summary, Findings & Risks, Recommended Next Actions, **and** the latest `.do/history` state changes (phase/task/branch/completed deltas)\n     * Appends preset-specific sections: progress snapshot (from `scripts/progress`), ASCII visuals, and a dependency audit when manifests are detected\n   - Output:\n     * Terminal summary showing metadata count + diff file path\n     * Diff markdown stored alongside the reports for sharing\n\n3. **Subcommand: full** (or /status full):\n   - Show both progress and report in one comprehensive view",
			AgentInvolvement: []string{
				"Project Orchestrator",
				"QA Engineer",
				"Documentation Lead",
			},
			FilesRead: []string{
				".do/plan/TASKS.md",
				".do/system/history/active_state.json",
				".do/system/history/state-*.json",
				"<project>/.do/reports/SCAN_REPORT_*.md",
			},
			FilesModified: []string{
				"<project>/.do/reports/SCAN_REPORT_*.json",
				"<project>/.do/reports/SCAN_DIFF_<date>.md",
			},
			Options:      "- `--preset standard` *(default)* – complete report\n- `--preset exec` – condensed executive view + visuals\n- `--preset detailed` – expanded sections with dependency audit\n- `.do/reports/config.json` (optional) can set preset and sections",
			Requirements: "- Go 1.21+\n- Reports must follow `SCAN_REPORT_YYYY-MM-DD.md` naming",
			Examples: []string{
				"/status → Show current progress",
				"/status report → Generate scan report metadata and diffs",
				"/status full → Show progress and report together",
			},
		},
		{
			Name:        "feedback",
			Category:    "tools",
			Trigger:     "/feedback <type> \"Title\" \"Details\" [--github <url>] [--author <name>]",
			Description: "Log feedback (bug, feature, question, note)",
			Action:      "When you run `/feedback ...`:\n\n1. **Parse arguments**\n   - `type`: bug | feature | question | note (defaults to `note`)\n   - `title`: short summary (required)\n   - `details`: multiline description (optional)\n   - `--author`: person filing feedback (defaults to `anonymous`)\n   - `--github`: optional issue URL if mirrored upstream\n2. **Log entry** via `go run scripts/feedback/main.go ...`\n   - Appends markdown to `docs/history/feedback.md`\n   - Updates JSON log `docs/history/feedback.json` for automation\n3. **Surface in workflow**\n   - `/report` command ingests latest feedback when generating scan metadata/diffs\n   - Future scans can summarize outstanding feedback items\n4. **Response**\n   - \"Feedback logged (type=bug) → docs/history/feedback.md\"",
			AgentInvolvement: []string{
				"Product Manager",
				"QA Engineer",
				"Documentation Lead",
			},
			FilesRead: []string{
				"docs/history/feedback.md (created if missing)",
				"docs/history/feedback.json",
			},
			FilesModified: []string{
				"docs/history/feedback.md",
				"docs/history/feedback.json",
			},
			Examples: []string{
				"/feedback bug \"QR download fails\" \"API returns 500 when Accept header missing\"",
				"/feedback feature \"Add dark mode\" \"Marketing wants dark hero section\" --author PM",
				"/feedback question \"Rate limit\" \"What are the prod limits?\" --github https://github.com/org/repo/issues/123",
			},
			Notes: "- Requires Go 1.21+. Command shells run: `go run scripts/feedback/main.go --type <type> --title \"...\" --details \"...\" --author \"...\" --github <url>`\n- Works in any generated project (paths relative to project root).\n- Add new feedback types by passing a custom string (stored as lowercase).",
		},
		{
			Name:        "state",
			Category:    "tools",
			Trigger:     "/state <subcommand>",
			Description: "Manage project state history",
			Action:      "The `/state` helper wraps `go run scripts/statehistory/main.go` so you can manage `.do/active_state.json` history safely.\n\n### snapshot\n1. Writes the current `.do/active_state.json` into `.do/history/state-<timestamp>.json`\n2. Accepts optional flags:\n   - `--reason` → stored in the snapshot metadata\n   - `--label` → appended to the file name (e.g., build, finished)\n3. Output: `Snapshot saved: .do/history/state-20251124T120000Z-build.json`\n\n### list\n1. Lists recent entries (default: last 10)\n2. `--json` emits machine-readable summaries for scripts/CI\n\n### diff\n1. Compares two snapshots (default: latest vs previous)\n2. Shows Markdown summary (phase/task/branch/completed deltas) or JSON if `--json`\n3. Used by `/progress` and `/report` to surface state deltas\n\n### restore\n1. Requires `--file <id>` and `--yes` confirmation for guardrails\n2. Restores `.do/active_state.json` from the selected snapshot\n3. Optionally captures a new snapshot (`--snapshot=false` to skip) so rollbacks themselves are logged\n4. Respond with confirmation + reminder to rerun `/progress`",
			AgentInvolvement: []string{
				"Project Orchestrator",
				"QA Engineer",
			},
			FilesRead: []string{
				"`.do/active_state.json`",
				"`.do/history/state-*.json`",
			},
			FilesModified: []string{
				"`.do/history/state-*.json` (new entries)",
				"`.do/active_state.json` (when restoring)",
			},
			Examples: []string{
				"/state snapshot --reason \"after /build 5.8\"",
				"/state list --limit 5",
				"/state diff --json",
				"/state restore --file state-20251124T120000Z.json --yes",
			},
			Notes: "- State history is now required before/after `/build` and `/finished`\n- Restores should be rare; always snapshot first so you can undo mistakes",
		},
		{
			Name:        "github",
			Category:    "tools",
			Trigger:     "/github <subcommand>",
			Description: "GitHub operations: metadata, CI, releases, issues, PRs",
			Action:      "When user types /github <subcommand>:\n\n### `/github info`\nRuns:\n```\ngo run scripts/githubmeta/main.go --project . --sync-readme\n```\n- Detects primary remote + default branch\n- Extracts success metrics from `.do/system/PRD.md`\n- Updates the README KPI block between `<!-- KPIS:START -->` / `<!-- KPIS:END -->`\n- Persists metadata to `docs/history/github-meta.json` for offline use\n\n### `/github issue \"Title\" \"Body\"`\nOutputs a ready-to-run `gh issue create` command with the detected repo slug, e.g.:\n```\ngo run scripts/githubmeta/main.go --project . --issue-title \"Fix cache\" --issue-body \"Details here\"\n```\nCopy/paste the printed `gh issue create` command (or pipe it) to open the issue.\n\n### `/github milestone \"Name\" [due-date]`\nPrints a `gh api` command to create a milestone:\n```\ngo run scripts/githubmeta/main.go --project . --milestone-title \"MVP\" --milestone-due 2025-01-15T00:00:00Z\n```\n\n### `/github ci [regenerate]`\nGenerates CI workflow for branch prefixes:\n1. Reads `docs/history/branch-matrix.json` to understand what jobs/required checks each branch prefix needs (e.g., `task/`, `feature/`, `hotfix/`).\n2. Runs the generator:\n   ```bash\n   go run scripts/branchci/main.go --matrix docs/history/branch-matrix.json --out .github/workflows\n   ```\n3. Emits `.github/workflows/task-branches.yml`, a workflow that:\n   - Triggers on pushes to `task/*` (and can be expanded for other prefixes)\n   - Spins up jobs per branch prefix (lint/test/build/etc.)\n   - Adds a summary job so reviewers know which checks are required per branch\n4. Output: \"Workflow generated: .github/workflows/task-branches.yml\"\n\n### `/github release`\nRelease management:\n1. **Release Planning**: Release Captain plans the release\n2. **Version Management**: Manage version numbers and semantic versioning\n3. **Release Notes**: Generate release notes\n4. **Deployment Planning**: Plan deployment strategy\n5. **Response**: \"Release planned! Review release notes and deployment plan.\"",
			AgentInvolvement: []string{
				"Release & Growth Manager",
				"Release Captain",
				"DevOps Engineer",
			},
			FilesRead: []string{
				"`.git/` metadata",
				"`.do/system/PRD.md`",
				"docs/history/branch-matrix.json",
				".do/plan/TASKS.md",
				"docs/history/CHANGELOG.md",
			},
			FilesModified: []string{
				"`docs/history/github-meta.json`",
				"`README.md` KPI section when `--sync-readme` is used",
				".github/workflows/task-branches.yml",
				"docs/history/CHANGELOG.md",
				".do/system/RELEASE.md",
			},
			Customize:     "Edit `docs/history/branch-matrix.json` to add or tweak prefixes, jobs, and required checks. Re-run `/github ci` after editing to regenerate the workflow.",
			Notes:         "- Generated workflow expects Go 1.21 and the standard lint/test/build jobs. Adapt `scripts/branchci/main.go` if your stack differs.\n- Use `/github ci` whenever you add a new branch naming convention or need different CI steps per branch type.",
			OfflineSafety: "- If git remote detection fails, the script logs a warning and keeps the last cached metadata (`docs/history/github-meta.json`). You can still update KPIs from PRD without network access.",
			Examples: []string{
				"/github info",
				"/github issue \"Fix cache\" \"Cache misses spike\"",
				"/github ci → Generate CI workflow",
				"/github release → Release management",
			},
		},
		{
			Name:        "security",
			Category:    "tools",
			Trigger:     "/security [<subcommand>]",
			Description: "Security review and audit",
			Action: `When user types /security or /security <subcommand>:

1. **If no subcommand provided**: Show menu:
   - "Select security operation:"
   - "1. Review - Security review"
   - "2. Audit - Security audit"
   - "3. Both - Review and audit"
   - Wait for user selection

2. **Subcommand: review** (or user selects option 1):
   - Security Review: Security Lead conducts security review
   - Vulnerability Assessment: Identify and document security vulnerabilities
   - Generate Report: Create security review report
   - Response: "Security review complete! Review security findings in SECURITY.md"

3. **Subcommand: audit** (or user selects option 2):
   - Security Audit: Security Lead conducts comprehensive security audit
   - Vulnerability Scanning: Scan for security vulnerabilities
   - Compliance Check: Verify compliance with security standards
   - Generate Audit Report: Create security audit report
   - Response: "Security audit complete! Review audit findings in SECURITY_AUDIT.md"

4. **Subcommand: both** (or user selects option 3):
   - Run both review and audit in sequence
   - Generate comprehensive security report
   - Response: "Security review and audit complete! Review findings in SECURITY.md and SECURITY_AUDIT.md"`,
			AgentInvolvement: []string{
				"Security Lead",
			},
			FilesRead: []string{
				"src/**",
				".do/system/ARCHITECTURE.md",
				".do/system/SECURITY.md",
			},
			FilesModified: []string{
				".do/system/SECURITY.md",
				".do/system/SECURITY_AUDIT.md",
			},
			Examples: []string{
				"/security → Show security menu",
				"/security review → Security review",
				"/security audit → Security audit",
				"/security both → Review and audit",
			},
		},
		{
			Name:        "permissions",
			Category:    "tools",
			Trigger:     "/permissions",
			Description: "Design RBAC system (role-based access control)",
			Action: `When user types /permissions:

1. **Design RBAC**: Security Lead and Backend Lead design role-based access control
2. **Define Roles**: Create role definitions and permissions
3. **Generate Documentation**: Document RBAC system
4. **Response**: "RBAC system designed! Review role definitions in RBAC.md"`,
			AgentInvolvement: []string{
				"Security Lead",
				"Backend Lead",
			},
			FilesRead: []string{
				".do/system/PRD.md",
				".do/system/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".do/system/RBAC.md",
			},
			Examples: []string{
				"/permissions",
			},
		},
		{
			Name:        "optimize",
			Category:    "optimize",
			Trigger:     "/optimize [<subcommand>]",
			Description: "Project optimization hub",
			Action: `When user types /optimize or /optimize <subcommand>:

1. **If no subcommand provided**: Show menu:
   - "Select optimization type:"
   - "1. Design - UI/UX improvements"
   - "2. Finance - Cost optimization"
   - "3. Performance - Performance optimization"
   - "4. All - Run all optimizations"
   - Wait for user selection

2. **Subcommand: design** (or user selects option 1):
   - UI/UX Review: Design Manager and UI/UX Designer review interface
   - Improvement Suggestions: Provide UI/UX improvement recommendations
   - Update Design System: Update DESIGN_SYSTEM.md with improvements
   - Response: "UI/UX improvements suggested! Review design updates in DESIGN_SYSTEM.md"

3. **Subcommand: finance** (or user selects option 2):
   - Cost Analysis: Analyze current infrastructure and service costs
   - Optimization Recommendations: Provide cost optimization suggestions
   - Generate Cost Plan: Create cost optimization plan
   - Response: "Cost analysis complete! Review optimization recommendations in COST_OPTIMIZATION.md"

4. **Subcommand: performance** (or user selects option 3):
   - Performance Analysis: Performance Engineer analyzes application performance
   - Identify Bottlenecks: Detect performance bottlenecks and issues
   - Optimization Suggestions: Provide performance optimization recommendations
   - Generate Performance Report: Create performance optimization report
   - Response: "Performance analysis complete! Review optimization recommendations in PERFORMANCE_OPTIMIZATION.md"

5. **Subcommand: all** (or user selects option 4):
   - Run all three optimizations in sequence (design, finance, performance)
   - Generate comprehensive optimization report
   - Response: "All optimizations complete! Review reports in DESIGN_SYSTEM.md, COST_OPTIMIZATION.md, and PERFORMANCE_OPTIMIZATION.md"`,
			AgentInvolvement: []string{
				"Design & UX Manager",
				"UI/UX Designer",
				"DevOps Engineer",
				"Performance Engineer",
			},
			FilesRead: []string{
				".do/system/DESIGN_SYSTEM.md",
				".do/system/ARCHITECTURE.md",
				"src/**",
			},
			FilesModified: []string{
				".do/system/DESIGN_SYSTEM.md",
				".do/system/COST_OPTIMIZATION.md",
				".do/system/PERFORMANCE_OPTIMIZATION.md",
			},
			Examples: []string{
				"/optimize → Show optimization menu",
				"/optimize design → UI/UX improvements",
				"/optimize finance → Cost optimization",
				"/optimize performance → Performance optimization",
				"/optimize all → Run all optimizations",
			},
		},
	}
}

// GetCommandByName returns a command by name, or nil if not found
func GetCommandByName(name string) *Command {
	commands := GetAllCommands()
	for i := range commands {
		if commands[i].Name == name {
			return &commands[i]
		}
	}
	return nil
}

// GetCoreCommands returns only the core commands (essential daily workflow)
func GetCoreCommands() []Command {
	allCommands := GetAllCommands()
	coreNames := map[string]bool{
		"hello":   true,
		"tell":    true,
		"meeting": true,
		"write":   true,
		"plan":    true,
		"build":   true,
		"status":  true,
	}
	var coreCommands []Command
	for _, cmd := range allCommands {
		if coreNames[cmd.Name] {
			coreCommands = append(coreCommands, cmd)
		}
	}
	return coreCommands
}

// GetSquadCommands returns only the tools and optimize commands
func GetSquadCommands() []Command {
	allCommands := GetAllCommands()
	coreNames := map[string]bool{
		"hello":   true,
		"tell":    true,
		"meeting": true,
		"write":   true,
		"plan":    true,
		"build":   true,
		"status":  true,
	}
	var squadCommands []Command
	for _, cmd := range allCommands {
		if !coreNames[cmd.Name] {
			squadCommands = append(squadCommands, cmd)
		}
	}
	return squadCommands
}

// commandTemplate is the markdown template for command files
const commandTemplate = `# /{{.Name}}

## Trigger
{{.Trigger}}{{if .Examples}}

## Examples
{{range .Examples}}- {{.}}
{{end}}{{end}}

## Action
{{.Action}}

## Agent Involvement
{{range .AgentInvolvement}}- **{{.}}**
{{end}}{{if .FilesRead}}
## Files Read
{{range .FilesRead}}- {{.}}
{{end}}{{end}}{{if .FilesModified}}
## Files Modified
{{range .FilesModified}}- {{.}}
{{end}}{{end}}{{if .GitHubAutomation}}
## GitHub Automation
{{.GitHubAutomation}}{{end}}{{if .Requirements}}
## Requirements
{{.Requirements}}{{end}}{{if .Notes}}
## Notes
{{.Notes}}{{end}}{{if .Customize}}
## Customize
{{.Customize}}{{end}}{{if .Options}}
## Options
{{.Options}}{{end}}{{if .OfflineSafety}}
## Offline Safety
{{.OfflineSafety}}{{end}}
`

// RenderCommandMarkdown renders a command to markdown format
// Uses file-based template if available, falls back to hardcoded template
func RenderCommandMarkdown(cmd *Command) (string, error) {
	if cmd == nil {
		return "", fmt.Errorf("command cannot be nil")
	}

	// Try file-based template first
	rendered, err := RenderCommandMarkdownFileBased(cmd)
	if err == nil {
		return rendered, nil
	}

	// Fallback to hardcoded template
	tmpl, err := template.New("command").Parse(commandTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cmd); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// CommandsGenerator generates the command markdown files.
type CommandsGenerator struct{}

// Name returns the name of the generator.
func (g *CommandsGenerator) Name() string {
	return "Commands"
}

// Generate creates IDE-specific commands/ directories and all command markdown files.
func (g *CommandsGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// Ensure IDEs list is populated (for backward compatibility)
	if len(request.IDEs) == 0 && request.IDE != "" {
		request.IDEs = []string{request.IDE}
	}

	// Central location for commands: .do/core/commands/
	centralCommandsDir := filepath.Join(projectPath, ".do", "core", "commands")

	// Create central commands directory
	if err := utils.CreateDirectory(centralCommandsDir); err != nil {
		return fmt.Errorf("failed to create central commands directory: %w", err)
	}

	// Get all commands
	commands := GetAllCommands()

	// Get all commands (try file-based first, fallback to hardcoded)
	allCommands, err := GetAllCommandsFileBased()
	if err != nil {
		// Log the error but continue with fallback
		// The GetAllCommandsFileBased already falls back to GetAllCommands() internally
		allCommands = GetAllCommands()
	}

	// Always regenerate commands to ensure all IDEs get the latest updates
	// Group commands by category
	commandsByCategory := make(map[string][]Command)
	for _, cmd := range allCommands {
		category := cmd.Category
		if category == "" {
			category = "other" // Default category for commands without category
		}
		commandsByCategory[category] = append(commandsByCategory[category], cmd)
	}

	// Generate commands organized by category (always regenerate to ensure updates)
	for category, categoryCommands := range commandsByCategory {
		categoryDir := filepath.Join(centralCommandsDir, category)
		if err := utils.CreateDirectory(categoryDir); err != nil {
			return fmt.Errorf("failed to create category directory %s: %w", category, err)
		}

		// Generate each command file in its category folder
		for _, cmd := range categoryCommands {
			// Render command markdown
			markdown, err := RenderCommandMarkdown(&cmd)
			if err != nil {
				return fmt.Errorf("failed to render command %s: %w", cmd.Name, err)
			}

			// Write command file in category folder (overwrite existing to ensure updates)
			commandPath := filepath.Join(categoryDir, cmd.Name+".md")
			if err := utils.WriteFile(commandPath, []byte(markdown)); err != nil {
				return fmt.Errorf("failed to write command file %s: %w", cmd.Name+".md", err)
			}
		}
	}

	// Create symlinks in each IDE's commands directory pointing to central location
	for _, ide := range request.IDEs {
		ideCommandsDir, err := getIDECommandsDir(projectPath, ide)
		if err != nil {
			return fmt.Errorf("failed to get commands directory for %s: %w", ide, err)
		}

		// Ensure parent directory exists
		parentDir := filepath.Dir(ideCommandsDir)
		if err := utils.CreateDirectory(parentDir); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", ide, err)
		}

		// Ensure IDE commands directory exists
		if err := utils.CreateDirectory(ideCommandsDir); err != nil {
			return fmt.Errorf("failed to create commands directory for %s: %w", ide, err)
		}

		// Create symlinks for each command category folder
		if err := createCommandCategorySymlinks(ideCommandsDir, centralCommandsDir); err != nil {
			// Fallback: copy commands if symlink creation fails
			if err := copyCommands(ideCommandsDir, centralCommandsDir, commands); err != nil {
				return fmt.Errorf("failed to create commands for %s (symlink and copy both failed): %w", ide, err)
			}
		}
	}

	return nil
}

// createCommandCategorySymlinks creates symlinks for each category folder in commands
func createCommandCategorySymlinks(ideCommandsDir, centralCommandsDir string) error {
	// Read all folders in the central commands directory
	entries, err := os.ReadDir(centralCommandsDir)
	if err != nil {
		return fmt.Errorf("failed to read central commands directory: %w", err)
	}

	var firstError error
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		folderName := entry.Name()
		centralFolderPath := filepath.Join(centralCommandsDir, folderName)
		ideFolderPath := filepath.Join(ideCommandsDir, folderName)

		// Remove existing symlink/directory if it exists
		if utils.PathExists(ideFolderPath) {
			// Check if it's already a symlink pointing to the right place
			if utils.IsSymlink(ideFolderPath) {
				target, err := os.Readlink(ideFolderPath)
				if err == nil {
					// Resolve to absolute path for comparison
					absTarget, _ := filepath.Abs(target)
					absCentral, _ := filepath.Abs(centralFolderPath)
					if absTarget == absCentral || filepath.Clean(absTarget) == filepath.Clean(absCentral) {
						// Already correctly linked, skip
						continue
					}
				}
			}
			// Remove existing directory/link
			if err := os.RemoveAll(ideFolderPath); err != nil {
				if firstError == nil {
					firstError = fmt.Errorf("failed to remove existing folder %s: %w", folderName, err)
				}
				continue
			}
		}

		// Create symlink for this folder
		if err := utils.CreateSymlink(ideFolderPath, centralFolderPath); err != nil {
			if firstError == nil {
				firstError = fmt.Errorf("failed to create symlink for folder %s: %w", folderName, err)
			}
			// Continue trying other folders
			continue
		}
	}

	return firstError
}

// copyCommands copies command category folders from central location to IDE directory (fallback)
func copyCommands(ideCommandsDir, centralCommandsDir string, commands []Command) error {
	// Read all folders in the central commands directory
	entries, err := os.ReadDir(centralCommandsDir)
	if err != nil {
		return fmt.Errorf("failed to read central commands directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		folderName := entry.Name()
		centralFolderPath := filepath.Join(centralCommandsDir, folderName)
		ideFolderPath := filepath.Join(ideCommandsDir, folderName)

		// Copy the entire folder recursively
		if err := utils.CopyDirectory(centralFolderPath, ideFolderPath); err != nil {
			return fmt.Errorf("failed to copy folder %s: %w", folderName, err)
		}
	}

	return nil
}

// getIDECommandsDir returns the commands directory path for the given IDE
func getIDECommandsDir(projectPath, ide string) (string, error) {
	switch ide {
	case "Cursor":
		return filepath.Join(projectPath, ".cursor", "commands"), nil
	case "Claude Code":
		return filepath.Join(projectPath, ".claude", "commands"), nil
	case "Antigravity":
		return filepath.Join(projectPath, ".antigravity", "commands"), nil
	case "Windsurf":
		return filepath.Join(projectPath, ".windsurf", "commands"), nil
	case "Cline":
		return filepath.Join(projectPath, ".cline", "commands"), nil
	case "OpenCode":
		return filepath.Join(projectPath, ".opencode", "commands"), nil
	default:
		return "", fmt.Errorf("unsupported IDE: %s", ide)
	}
}

// GenerateCommands is a convenience function that creates a CommandsGenerator and generates commands
func GenerateCommands(request *models.ProjectRequest, projectPath string) error {
	generator := &CommandsGenerator{}
	return generator.Generate(request, projectPath)
}
