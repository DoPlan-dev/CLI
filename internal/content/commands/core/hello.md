---
name: hello
category: core
trigger: "/hello [<subcommand>]"
description: "Welcome, tutorial, and command introductions"
agentInvolvement:
  - Project Orchestrator
  - Product Manager
examples:
  - "/hello"
  - "/hello meeting"
  - "/hello plan"
filesRead:
  - ".do/system/user_profile.json"
filesModified:
  - ".do/system/user_profile.json"
  - ".do/system/QUICK_REFERENCE.md"
  - "docs/references/QUICK_REFERENCE.md"
  - "docs/overview/AGENT_HIERARCHY.md"
  - "docs/references/COMMAND_EXAMPLES.md"
  - "docs/tutorials/TUTORIAL_NOTES.md"
---

When user types /hello or /hello <subcommand>:

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
   - Display the agent hierarchy
   - Explain accuracy & reliability (high accuracy, context-aware, transparent, learnable)

5. **Development Support Option**:
   - Ask user to choose between Guided Mode or Independent Mode
   - Store as development_support_mode

6. **First-Time Tutorial Walkthrough**:
   - Explain each command: /tell, /meeting, /write, /content, /plan, /build
   - Offer interactive test drive

7. **Save Reference Materials**:
   - Create user profile and reference documents
   - Save to .do/system/ and the organized docs subfolders (`docs/overview`, `docs/references`, `docs/tutorials`)

**Subcommands**: meeting, plan, build, github (see full action in source)

