# `/hello` Command Design - First-Time Welcome Experience

## 🎯 Purpose
A warm, personalized first-time welcome that introduces users to DoPlan, explains the system, and sets up their development journey.

---

## 📋 Command Flow (Proposed)

### Step 1: Warm Greeting
**Agent says:**
```
Hello! 👋 I'm DoPlan, your AI development partner. 
I'm here to guide you on how to turn your idea into a real product.

It's great to meet you! Let's get to know each other a bit.
```

### Step 2: Personal Introduction
**Ask user:**
1. **Name**: "What should I call you?" 
   - Store as: `user_name` (use throughout conversation)
   - Use friendly, casual tone

2. **Work Experience**: "What's your experience with development?"
   - Options:
     - "I'm completely new to development"
     - "I have some basic experience"
     - "I'm an intermediate developer"
     - "I'm an experienced developer"
   - Store as: `user_experience_level`
   - This will also be used for `/meeting` command

**Agent responds:**
```
Nice to meet you, [Name]! I'll be here to help you every step of the way.
```

---

### Step 3: Tutorial - System Overview
**Agent explains:**

#### 3.1 What DoPlan Is
```
DoPlan is an AI-powered project director that helps you:
- Plan your project from idea to launch
- Generate all necessary documents (PRD, Architecture, Design System)
- Create content (SEO-ready, legal pages, marketing materials)
- Guide you through development with step-by-step instructions
- Track your progress and celebrate milestones
```

#### 3.2 Sub-Agent System Explanation
```
I work with a team of 18 specialized AI agents, each an expert in their field:

👔 Leadership Team:
- Project Orchestrator (me!) - Coordinates everything
- Product Manager - Plans features and requirements

💻 Engineering Team:
- Engineering Lead - Technical decisions
- System Architect - System design
- Frontend/Backend Leads - Code architecture
- DevOps Engineer - Infrastructure
- Security Lead - Security best practices
- Performance Engineer - Optimization

🎨 Design Team:
- Design Manager - Overall design strategy
- UI/UX Designer - User experience

✅ Quality Team:
- QA Manager - Testing strategy
- QA Engineer - Quality assurance

📈 Growth Team:
- Release Manager - Launch planning
- Growth Coach - User growth

📝 Documentation Team:
- Documentation Lead - Documentation strategy
- Documentation Writer - Technical writing

Each agent is a specialist, ensuring you get expert guidance in every area.
```

**Agent Hierarchy (Text-Based Tree - Recommended Format):**
```
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
```

> **Note**: This text-based tree format is used in both chat interfaces and documentation for optimal readability. See `docs/AGENT_HIERARCHY_CHAT_PREVIEW.md` for alternative formats.

#### 3.3 Accuracy & Reliability
```
How accurate is DoPlan?

✅ **High Accuracy**: Each agent is trained on best practices and industry standards
✅ **Context-Aware**: Agents understand your specific project and adapt accordingly
✅ **Transparent**: All decisions and recommendations are explained
✅ **Learnable**: The system learns from your preferences and adjusts

Think of us as your expert development team, always available, always learning.
```

---

### Step 4: Development Support Option
**Agent asks:**
```
[Name], I can support you throughout your development journey in two ways:

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

Which would you prefer? (guided/independent)
```

**Store preference as:** `development_support_mode`

**If Guided Mode selected:**
```
Perfect! I'll be your development companion. After each feature you complete, 
I'll check in with you, celebrate your progress, and guide you to the next step.

You can always change this later if you want more independence.
```

**If Independent Mode selected:**
```
Got it! I'll be here when you need me. Just ask if you need guidance or have questions.
```

---

### Step 5: First-Time Tutorial Walkthrough
**Agent says:**
```
Now, let's do a quick walkthrough together. This is just for learning - 
we won't create anything real yet. Ready? Let's go! 🚀
```

**Tutorial Steps:**

1. **Show `/tell` command:**
   ```
   First, you'll use /tell to capture your project idea.
   Example: /tell "I want to build a todo app"
   This saves your idea and activates our planning system.
   ```

2. **Show `/meeting` command:**
   ```
   Next, use /meeting to have a discovery session with our team.
   We'll ask you questions about your project, adapt to your experience level,
   and create a comprehensive plan. This is where we get to know your project!
   ```

3. **Show `/write` command with Example Output:**
   ```
   After the meeting, use /write to generate your project documents:
   - PRD (Product Requirements Document)
   - Architecture document
   - Design System
   These are your project's foundation.
   ```

   **Example Output Preview:**
   ```
   📄 PRD.md will include:
   - Executive Summary
   - User Personas
   - Feature Requirements
   - Success Metrics
   - Timeline

   🏗️ ARCHITECTURE.md will include:
   - System Overview
   - Technology Stack
   - Database Design
   - API Structure
   - Deployment Strategy

   🎨 DESIGN_SYSTEM.md will include:
   - Color Palette
   - Typography
   - Component Library
   - UI Patterns
   - Brand Guidelines
   ```

   **Show actual example snippet:**
   ```
   Here's a preview of what PRD.md looks like:

   # Product Requirements Document
   
   ## Executive Summary
   [Your project summary based on meeting]
   
   ## User Personas
   - Primary User: [Based on your input]
   - Secondary User: [Based on your input]
   
   ## Core Features
   1. Feature 1: [Description]
   2. Feature 2: [Description]
   
   ...and much more!
   ```

4. **Show `/content` command:**
   ```
   Use /content to generate all the content you need:
   - Website pages
   - Legal documents
   - Marketing materials
   - SEO-optimized copy
   All ready to use!
   ```

5. **Show `/plan` command:**
   ```
   Use /plan to create your execution plan and task list.
   This breaks down your project into manageable steps.
   ```

6. **Show `/build` command:**
   ```
   Finally, use /build to start coding!
   I'll guide you through each task, one step at a time.
   ```

**Agent asks:**
```
Does this workflow make sense, [Name]? 

Would you like to try a "test drive" of one command to see how it works?
This is just practice - we won't save anything. (yes/no)
```

**If user says yes to test drive:**
```
Great! Let's try an interactive test drive, [Name]!

I'll guide you through trying a command yourself. This is safe - 
nothing will be saved. Ready? Let's go! 🚀

---

**Interactive Test Drive: /tell Command**

Let's practice using /tell. I'll ask you a question, and you can respond!

Me: "What's your project idea? (This is just practice - type anything!)"
[Wait for user input]

[User types something, e.g., "I want to build a todo app"]

Me: "Perfect! Here's what would happen in real mode:

✅ Your idea: '[user's input]'
✅ Saved to IDEA.md
✅ Project Orchestrator activated
✅ Ready for next step: /meeting

In real mode, this would be saved to .do/system/IDEA.md
But since this is practice, nothing was saved!

---

Would you like to try another command? (yes/no)
- /meeting (see how the discovery meeting works)
- /write (see what documents get generated)
- /content (see content generation)
- Or type 'done' to finish test drive
```

**If user wants to try another command:**

**Option: /meeting test drive:**
```
Great! Let's try /meeting in test mode.

I'll show you how the meeting works. First, I need to know:
"What's your experience level?" (beginner/intermediate/advanced)

[User responds]

"Perfect! Based on that, I would show you these speed options:
[Show filtered options based on their level]

Which speed would you like to try? (1/2/3)

[User selects]

"Great choice! Now, let me ask you a few sample questions:

[Ask 2-3 sample questions from Phase 01, adapted to their level]

[User answers]

"Excellent! In real mode, I would:
✅ Ask more questions based on your answers
✅ Adapt questions to your project type
✅ Create a comprehensive plan
✅ Save everything to BRAINSTORM.md

This gives you a feel for how the meeting works!
```

**Option: /write test drive:**
```
Let's see what /write generates!

Based on your test idea, here's what I would create:

📄 PRD.md Preview:
[Show actual example structure]

🏗️ ARCHITECTURE.md Preview:
[Show actual example structure]

🎨 DESIGN_SYSTEM.md Preview:
[Show actual example structure]

In real mode, these would be full documents saved to .do/system/
But you can see the structure and quality!

Want to try another command or finish? (yes/done)
```

**If user says done:**
```
Perfect! You've got a good feel for how commands work, [Name]!

Remember: In real mode, everything gets saved and you can build on it.
The test drive was just to show you how it works.

Ready to continue with the tutorial?
```

**If user says no:**
```
No problem! You can always try commands when you're ready.
```

**Agent asks:**
```
Any questions before we start?
```

---

### Step 6: Personalized Tips Based on Experience Level

**For Beginners:**
```
[Name], as a beginner, here are some tips to help you succeed:

💡 Start Simple: Don't try to build everything at once. Start with core features.
💡 Ask Questions: If something is unclear, ask me! I'm here to help.
💡 Take Your Time: There's no rush. Learn at your own pace.
💡 Use Guided Mode: I'll check in after each step to keep you on track.
💡 Celebrate Small Wins: Every feature you complete is progress!

Remember: Every expert was once a beginner. You've got this! 💪
```

**For Intermediate:**
```
[Name], with your experience, you'll move quickly! Here are some tips:

💡 Leverage Our Agents: Each agent is an expert - trust their recommendations.
💡 Customize as Needed: Feel free to adjust generated documents to your style.
💡 Focus on Architecture: Pay attention to the Architecture document - it's crucial.
💡 Use Independent Mode: You might prefer working at your own pace.
💡 Iterate: Don't worry about perfection - you can always refine later.

You have the skills - I'll help you apply them efficiently! 🚀
```

**For Advanced:**
```
[Name], as an experienced developer, you'll appreciate:

💡 Deep Technical Details: Our agents provide comprehensive technical guidance.
💡 Best Practices: All recommendations follow industry best practices.
💡 Customization: Full control to modify and extend generated content.
💡 Efficiency: Skip basic explanations, dive into advanced topics.
💡 Architecture Focus: Detailed system design and scalability considerations.

You know what you're doing - I'll help you do it faster and better! ⚡
```

### Step 7: Save Reference Materials to docs/

**Agent says:**
```
Before we finish, let me save some helpful reference materials to your docs/ folder
so you can access them anytime!
```

**What gets saved:**

1. **Quick Reference Card** → `docs/references/QUICK_REFERENCE.md`
   - Complete command cheat sheet
   - Organized by category
   - Easy to scan format

2. **Agent Hierarchy Diagram** → `docs/overview/AGENT_HIERARCHY.md`
   - Visual hierarchy diagram
   - Agent descriptions
   - Team structure explanation

3. **Command Examples** → `docs/references/COMMAND_EXAMPLES.md`
   - Example outputs for each command
   - Usage examples
   - What to expect from each command

4. **Tutorial Notes** → `docs/tutorials/TUTORIAL_NOTES.md`
   - Personalized notes based on your experience level
   - Tips specific to your skill level
   - Quick reminders

**Agent says:**
```
✅ All reference materials saved to docs/ folder!
You can find them anytime in your project's docs/ directory.
```

### Step 8: Quick Reference Card Display
**Agent says:**
```
Before we finish, let me give you a quick reference card you can use anytime!
```

**Display Quick Reference Card:**
```
╔══════════════════════════════════════════════════════════════╗
║           DoPlan Command Cheat Sheet                         ║
╠══════════════════════════════════════════════════════════════╣
║                                                               ║
║  🚀 STARTING YOUR PROJECT                                     ║
║  ─────────────────────────────────────────────────────────── ║
║  /hello      → Welcome & tutorial (first time only)          ║
║  /tell       → Capture your project idea                      ║
║  /meeting    → Discovery meeting with adaptive speed          ║
║  /team       → Show all agents and hierarchy                  ║
║                                                               ║
║  📋 PLANNING                                                  ║
║  ─────────────────────────────────────────────────────────── ║
║  /write      → Generate PRD, Architecture, Design System     ║
║  /content    → Generate SEO-ready content                     ║
║  /change     → Edit any document                              ║
║  /good       → Approve & lock plan                            ║
║  /plan       → Generate execution plan & tasks                ║
║                                                               ║
║  💻 DEVELOPMENT                                               ║
║  ─────────────────────────────────────────────────────────── ║
║  /load       → Inject context into AI agents                 ║
║  /build      → Start coding next task                         ║
║  /progress   → Show current progress                          ║
║  /finished   → Mark current task done                         ║
║  /state      → Manage project state                           ║
║                                                               ║
║  📤 PUBLISHING                                                ║
║  ─────────────────────────────────────────────────────────── ║
║  /ship       → Release management                             ║
║  /github     → GitHub operations                              ║
║                                                               ║
║  🛠️ MANAGEMENT                                               ║
║  ─────────────────────────────────────────────────────────── ║
║  /branchci   → CI/CD workflows                                ║
║  /report     → Generate project report                        ║
║  /feedback   → Log feedback                                   ║
║                                                               ║
║  ✨ QUALITY & BUSINESS                                        ║
║  ─────────────────────────────────────────────────────────── ║
║  /secure     → Security review                                ║
║  /safe       → Security audit                                 ║
║  /pretty     → UI/UX improvements                             ║
║  /seo        → SEO optimization                               ║
║  /roles      → RBAC design                                    ║
║  /money      → Billing setup                                  ║
║  /cheap      → Cost optimization                              ║
║                                                               ║
║  💡 TIP: Type /help [command] for detailed help               ║
╚══════════════════════════════════════════════════════════════╝
```

**Save reference card to:** `.do/system/QUICK_REFERENCE.md`

**Agent says:**
```
I've saved this cheat sheet in two places for easy access:
✅ .do/system/QUICK_REFERENCE.md (in your project)
✅ docs/references/QUICK_REFERENCE.md (in documentation folder)

You can reference it anytime, or just type /help to see all commands!
```

**Also save to:** `docs/references/QUICK_REFERENCE.md` (project root docs folder)
- Accessible from project root
- Part of project documentation
- Can be version controlled
- Easy to share with team members

### Step 9: Encouragement & Ready to Start
**Agent says:**
```
Excellent! You're all set, [Name]! 🎉

You now understand:
✅ How DoPlan works
✅ Our agent system and expertise (with visual hierarchy)
✅ The command workflow (with examples)
✅ How I'll support you (if you chose guided mode)
✅ Have reference materials saved in docs/ folder
✅ Received personalized tips for your experience level
✅ Tried commands in interactive test drive

**Reference Materials Available:**
📁 docs/references/QUICK_REFERENCE.md - Command cheat sheet
📁 docs/overview/AGENT_HIERARCHY.md - Agent structure diagram
📁 docs/references/COMMAND_EXAMPLES.md - Example outputs
📁 docs/tutorials/TUTORIAL_NOTES.md - Your personalized notes

You're ready to start developing! 

Here's what to do next:
1. Type /tell to share your project idea
2. Then type /meeting to start planning
3. I'll guide you through the rest!

Remember: I'm here to help. If you get stuck or have questions, just ask!

Let's turn your idea into reality! 🚀

Ready when you are, [Name]!
```

---

## 💾 Data Storage

**Save to:** `.do/system/user_profile.json`

```json
{
  "user_name": "John",
  "user_experience_level": "intermediate",
  "development_support_mode": "guided",
  "first_hello_completed": true,
  "first_hello_date": "2025-01-XX",
  "tutorial_completed": true,
  "test_drive_completed": false,
  "quick_reference_saved": true
}
```

**Also create:** 
1. `.do/system/QUICK_REFERENCE.md` (project-specific)
   - Contains the cheat sheet shown in tutorial
   - Can be updated as new commands are added
   - Easy to reference anytime

2. `docs/references/QUICK_REFERENCE.md` (project documentation)
   - Same content, saved in docs folder
   - Part of project documentation
   - Version controlled
   - Shareable with team

**Additional Reference Materials Saved to docs/:**
- `docs/overview/AGENT_HIERARCHY.md` - Visual agent hierarchy diagram
- `docs/references/COMMAND_EXAMPLES.md` - Example outputs and usage
- `docs/tutorials/TUTORIAL_NOTES.md` - Personalized tutorial notes based on user's experience level

**Also update:** `.do/system/history/active_state.json`
```json
{
  "phase": "onboarding",
  "user_profile_setup": true,
  "tutorial_completed": true
}
```

---

## 🔄 Subsequent Uses

**If user runs `/hello` again:**
```
Hey [Name]! 👋 Welcome back!

You've already completed the tutorial. 
Is there something specific you'd like help with?

- Type /tell to start a new project
- Type /help to see all available commands
- Or just tell me what you need!
```

---

## 🎨 Design Principles

1. **Warm & Personal**: Use user's name throughout, friendly tone
2. **Educational**: Explain system clearly, don't assume knowledge
3. **Encouraging**: Celebrate readiness, build confidence
4. **Non-Intrusive**: Ask permission, respect user choice
5. **Adaptive**: Adjust language based on experience level
6. **Supportive**: Offer help, make user feel supported

---

## 🤔 Questions to Consider

1. **Should tutorial be skippable?**
   - Current: Required on first run
   - Alternative: Offer to skip, but recommend completing

3. **Should we save conversation history?**
   - Store first conversation for personalization
   - Reference in future interactions

4. **What if user is returning (not first time)?**
   - Check if `first_hello_completed` exists
   - Show different welcome message
   - Offer to redo tutorial if needed

5. **Should guided mode be changeable?**
   - Yes, allow user to toggle later
   - Store preference, respect it in other commands

6. **How detailed should tutorial be?**
   - Current: Overview of each command
   - Alternative: More detailed with examples

7. **Should we show examples during tutorial?**
   - Yes, show example commands
   - Maybe even simulate one command

---

## 📚 Reference Materials Generated

### Files Created in docs/ Folder:

1. **docs/references/QUICK_REFERENCE.md**
   - Complete command cheat sheet
   - Organized by workflow category
   - Quick lookup format

2. **docs/overview/AGENT_HIERARCHY.md**
   - Visual ASCII diagram of agent hierarchy
   - Text-based tree structure
   - Agent descriptions and roles
   - Team organization explanation

3. **docs/references/COMMAND_EXAMPLES.md**
   - Example outputs for each major command
   - What /write generates (PRD, Architecture, Design System)
   - What /content generates
   - What /meeting creates
   - Usage examples with sample inputs/outputs

4. **docs/tutorials/TUTORIAL_NOTES.md**
   - Personalized notes based on user's experience level
   - Tips specific to beginner/intermediate/advanced
   - Quick reminders
   - Links to relevant documentation

### Files Created in .do/system/:

1. **.do/system/QUICK_REFERENCE.md**
   - Same as docs version
   - Project-specific location
   - Quick access during development

---

## 🚀 Future Enhancements

1. **Progress Tracking**: Show user's progress through tutorial with progress bar
2. **Welcome Back Messages**: Different messages for returning users with personalized stats
3. **Achievement System**: Celebrate completing tutorial with badges
4. **Interactive Agent Explorer**: Click on agents in hierarchy to learn more
5. **Command Examples Library**: Show real examples from other projects (anonymized)
6. **Video Tutorial Option**: Link to video walkthrough for visual learners
7. **Tutorial Replay**: Option to replay tutorial anytime
8. **Personalized Onboarding Path**: Different tutorial paths based on project type
9. **Interactive Command Builder**: Let users build custom commands in test mode
10. **Reference Material Updates**: Auto-update reference docs when new commands are added

---

## 📝 Command Definition Structure

```go
{
    Name:        "hello",
    Category:    "start",
    Trigger:     "/hello",
    Description: "First-time welcome and tutorial",
    Action: `[Detailed action steps as above]`,
    AgentInvolvement: []string{
        "Project Orchestrator",
    },
    FilesRead: []string{
        ".do/system/user_profile.json", // Check if already completed
    },
    FilesModified: []string{
        ".do/system/user_profile.json",
        ".do/system/history/active_state.json",
    },
}
```

---

**What do you think? Should we adjust anything before implementation?**

