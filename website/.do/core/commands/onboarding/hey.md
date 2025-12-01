# /hey

## Trigger
/hey [<subcommand>]

## Examples
- "/hey → Full tutorial (same as /hey goplan)"
- "/hey goplan → Full tutorial"
- "/hey meeting → Introduce /do command"
- "/hey plan → Introduce plan structure"
- "/hey build → Introduce /dev process"
- "/hey github → Introduce GitHub workflow"


## Action
When user types /hey or /hey <subcommand>:

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
     * /do - Capture project idea, conduct meeting, and refine
     * /plan - Generate documents and create execution plan
     * /dev - Start coding with step-by-step guidance
     * /sys - System control panel (engagement, performance, backup, etc.)
   - Ask: "Does this workflow make sense, [Name]? Would you like to try a 'test drive' of one command to see how it works? This is just practice - we won't save anything. (yes/no)"

7. **Interactive Test Drive** (if user says yes):
   - Offer to test drive: /do, /plan
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
      
      Ready to start? Type /do to capture your project idea!"

2. **Subcommand: meeting** (or /hey meeting):
   - **Storytelling Approach**: Start with a scenario:
     "Imagine you're about to build your dream project. You have the idea, but you're not sure where to start. That's where /do comes in - it's like having a smart product manager interview you to understand exactly what you want to build."
   
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
     "You: /do\n" +
     "DoPlan: What's your experience level?\n" +
     "You: Intermediate\n" +
     "DoPlan: Great! I'll show you Quick Start, Standard, and Comprehensive options.\n" +
     "        Which speed would you like? [Recommended: Standard]\n" +
     "You: Standard\n" +
     "DoPlan: Perfect! Let's start with your project vision...\n" +
     "[Interactive Q&A session]\n" +
     "DoPlan: Meeting complete! Your BRAINSTORM.md has been saved."
   
   - **Why It's Special**: "Unlike static forms, /do adapts in real-time. It asks follow-up questions, probes deeper when needed, and skips irrelevant sections. It's like having a conversation with an expert product manager who knows exactly what to ask."
   
   - **Call to Action**: "Ready to discover your project? Type /do and let's start the journey! 🎯"

3. **Subcommand: plan** (or /hey plan):
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

4. **Subcommand: build** (or /hey build):
   - **Journey Metaphor**: "Think of /dev as your development journey. Every task is a destination, and we'll guide you there step by step. Let me show you how smooth this journey can be!"
   
   - **Step-by-Step Adventure**:
     "Development flow:\n" +
     "Step 1: Task Selection\n" +
     "   You: /dev\n" +
     "   DoPlan: Found next task: 2.1 - User Authentication\n" +
     "   OR\n" +
     "   You: /dev 2.1\n" +
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
     "            Automatically marking as done...\n" +
     "   DoPlan: Task marked complete! Committing and pushing...\n" +
     "   [Auto-commit with conventional format]\n" +
     "   [Auto-push to remote]\n" +
     "   DoPlan: Task 2.1 complete! Ready for next task."
   
   - **Why It's Powerful**:
     * "No more manual branch management - we handle it!"
     * "No more forgetting to commit - we do it automatically!"
     * "No more wondering what's next - we track everything!"
     * "No more losing progress - we snapshot your state!"
   
   - **State Management Magic**: "Every time you start or finish a task, we take a snapshot. Want to see what changed? Type /sys status. Need to rollback? We've got you covered!"
   
   - **Real Example**: "Here's what happens in practice:\n" +
     "Morning: /dev → Start task 3.2\n" +
     "[Work on feature]\n" +
     "Afternoon: Agent detects completion\n" +
     "[Auto-commit, auto-push]\n" +
     "Evening: /dev → Start next task automatically\n" +
     "It's that seamless!"
   
   - **Call to Action**: "Ready to start building? Type /dev and let's turn your plan into code! 🚀"

5. **Subcommand: github** (or /hey github):
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
     - One-command CI/CD setup (/sys github ci)
     - Automated releases (/sys github release)"
   
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
     "1. /dev 2.1\n" +
     "2. [work on code]\n" +
     "3. Agent: Task complete! Automatically marking as done...\n" +
     "4. DoPlan automatically:\n" +
     "   Creates branch: task/2.1\n" +
     "   Commits: feat(task-2.1): implement user authentication\n" +
     "   Pushes to remote\n" +
     "   Updates CHANGELOG.md\n" +
     "   Suggests PR creation"
   
   - **GitHub Command Features**:
     "The /sys github command is your repository control center:
     
     📊 /sys github info
        - Syncs README with project KPIs
        - Updates repository metadata
        - Shows project health
     
     🐛 /sys github issue \"Title\" \"Description\"
        - Creates issues with proper formatting
        - Links to project documentation
     
     🎯 /sys github milestone \"Name\"
        - Manages project milestones
        - Tracks progress
     
     ⚙️ /sys github ci
        - Generates CI/CD workflows
        - Configures branch-based testing
        - One command setup!
     
     🚢 /sys github release
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
   
   - **Call to Action**: "Ready to automate your GitHub workflow? Type /sys github to see what's possible! 🌟"

## Agent Involvement
- **Project Orchestrator**

## Files Read
- ".do/system/user_profile.json"
- "docs/AGENT_HIERARCHY_CHAT_PREVIEW.md"

## Files Modified
- ".do/system/user_profile.json"
- ".do/system/QUICK_REFERENCE.md"
- "docs/references/QUICK_REFERENCE.md"
- "docs/overview/AGENT_HIERARCHY.md"
- "docs/references/COMMAND_EXAMPLES.md"
- "docs/tutorials/TUTORIAL_NOTES.md"


