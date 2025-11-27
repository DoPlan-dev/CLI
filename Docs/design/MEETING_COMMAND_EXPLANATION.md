# How `/meeting` Command Works

**Command**: `/meeting`  
**Category**: Start (Project Initiation)  
**Purpose**: Adaptive discovery meeting with speed options and dynamic content planning

---

## 🎯 Overview

The `/meeting` command conducts an intelligent, adaptive discovery interview that adapts to your project type, time constraints, and content needs. Unlike static questionnaires, it dynamically generates questions based on your project and provides multiple speed options.

---

## 📋 Step-by-Step Process

### Step 1: Project Type Detection
**What happens:**
- Reads `.do/system/IDEA.md` to understand your project
- Analyzes keywords and context to determine project type:
  - **Website/Agency/Personal**: Simple websites, agency portfolios, personal projects
  - **SaaS/Startup**: Software-as-a-Service products, startups, enterprise solutions

**Why it matters:**
- Determines which questions to ask
- Suggests appropriate meeting speed
- Adapts content creation options

---

### Step 2: Speed Options Presentation
**What you see:**
Four meeting speed options with clear descriptions:

1. **Quick Start** (Very Fast)
   - Duration: ~5-10 minutes
   - Phases: 01, 03 only (Vision & Outcomes, Experience & Tech)
   - Best for: Simple websites, personal projects, quick prototypes
   - Skips: Marketing, SEO, detailed operations

2. **Standard** (Fast)
   - Duration: ~15-20 minutes
   - Phases: 01, 02, 03, 06 (Vision, Audience, Experience, Delivery)
   - Best for: Company websites, agency projects, MVPs
   - Skips: Detailed content/SEO and marketing

3. **Comprehensive** (Medium)
   - Duration: ~30-45 minutes
   - Phases: All 6 phases with condensed questions
   - Best for: Complex projects, established businesses
   - Includes: All phases but streamlined

4. **Deep Dive** (Long)
   - Duration: ~60+ minutes
   - Phases: All 6 phases with full detailed questioning
   - Best for: SaaS products, startups, enterprise solutions
   - Includes: Extensive probing and detailed questions

**Smart Suggestions:**
- If project type is **Website/Agency/Personal**: Suggests Quick Start or Standard
- If project type is **SaaS/Startup**: Suggests Comprehensive or Deep Dive

---

### Step 3: User Speed Selection
**What happens:**
- You choose speed option (1, 2, 3, or 4)
- System adapts interview depth based on your choice
- Meeting structure adjusts accordingly

---

### Step 4: GitHub Repository Check
**What happens:**
1. Asks: "Do you have a GitHub repository for this project? (yes/no)"
2. If **yes**:
   - Asks for repository URL
   - Verifies repository exists
   - Checks automated workflows (`.github/workflows/`)
   - Verifies automated committing and pushing workflows
   - If missing/incorrect: Offers to set up GitHub Actions workflows
3. Documents GitHub repo info in meeting summary

**Why it matters:**
- Ensures automation is set up correctly
- Enables automated workflows from the start
- Prevents manual work later

---

### Step 5: Content Creation Needs Assessment
**This is where it gets dynamic!**

**Initial Question:**
"Do you need content creation for your project? (yes/no)"

**If Yes - Content Types by Project Type:**

#### For Website/Agency/Personal Projects:
- ✅ **App Pages**: Landing, about, services, contact pages
- ✅ **Legal Pages**: Privacy Policy, Terms of Service
- ✅ **Social Media**: Social media posts and content
- ✅ **Blog Posts**: If applicable
- ✅ **SEO Content**: Meta descriptions, keywords

#### For SaaS/Startup Projects:
- ✅ **App Pages**: Landing, features, pricing, about pages
- ✅ **Legal Pages**: Privacy Policy, Terms of Service, User Agreement
- ✅ **Social Media**: Social media posts and content
- ✅ **Blog Posts**: Content marketing articles
- ✅ **Documentation**: User guides, API documentation
- ✅ **Marketing Content**: Ad copy, campaigns
- ✅ **Email Templates**: Welcome emails, newsletters
- ✅ **SEO Content**: Meta descriptions, keywords

**For Each Selected Content Type:**
Asks: "How would you like to handle [content type]?"

**Three Options:**
1. **Full LLM Generation** (100% Automatic)
   - AI creates all content automatically
   - SEO-optimized by default
   - No input required from you

2. **Keyword-Guided** (You provide keywords)
   - You provide: Keywords, target audience, tone, specific requirements
   - AI creates content around your keywords
   - SEO-optimized with your keywords

3. **Hybrid** (You provide draft, AI refines)
   - You provide initial draft
   - AI refines and optimizes
   - Adds SEO elements, improves structure

**What Gets Documented:**
- Content types selected
- Creation preferences for each type
- Keywords provided (if any)
- Target audience and tone
- Specific requirements

---

### Step 6: Adaptive Phase Selection
**Based on selected speed:**

| Speed | Phases Included | What's Skipped |
|-------|----------------|----------------|
| **Quick Start** | 01, 03 | Marketing, SEO, detailed ops |
| **Standard** | 01, 02, 03, 06 | Detailed content/SEO, marketing |
| **Comprehensive** | 01-06 (condensed) | None, but streamlined |
| **Deep Dive** | 01-06 (full) | Nothing - everything included |

**Phase Breakdown:**
- **Phase 01**: Vision & Outcomes (Product Manager leads)
- **Phase 02**: Audience & Differentiation (Product Manager + Design Manager)
- **Phase 03**: Experience, UI/UX & Tech (Design Manager + Engineering Lead)
- **Phase 04**: Content & SEO (Content Strategist + SEO Specialist)
- **Phase 05**: Marketing & Growth (Marketing Manager)
- **Phase 06**: Delivery, Ops & Risks (Engineering Lead + Project Orchestrator)

---

### Step 7: Project Type Adaptation
**Questions adapt based on project type:**

#### Website/Agency/Personal:
- Focus: Design, content, basic functionality
- Skips: Complex business models, growth strategies, scalability concerns
- Emphasizes: Visual design, user experience, content needs

#### SaaS/Startup:
- Focus: Business model, scalability, monetization
- Emphasizes: User acquisition, competitive analysis, technical architecture, growth strategies
- Includes: Market analysis, pricing models, scaling concerns

---

### Step 8: Dynamic Question Generation
**Key Feature: Questions are NOT pre-made!**

**How it works:**
1. Loads phase templates from `.do/core/brainstorm/phase-*.md`
2. **Dynamically generates questions** based on:
   - Project type (website vs SaaS)
   - Selected content creation needs
   - User's keyword preferences (if provided)
   - Meeting speed selected
   - Previous answers

**Why dynamic?**
- Each project is unique
- Questions should be relevant to YOUR project
- Avoids asking irrelevant questions
- More efficient and focused

**Example:**
- If you selected "app-pages" content: Questions focus on page structure, content hierarchy
- If you selected "legal" content: Questions about compliance, jurisdiction
- If you're building SaaS: Questions about user onboarding, subscription models

---

### Step 9: Phase-by-Phase Interview
**How it works:**
- **One question at a time**: Not overwhelming, conversational
- **Probes deeper**: If answers are vague, asks follow-up questions
- **Waits for confirmation**: Won't move to next phase until you confirm
- **Adapts depth**: Based on selected speed (quick vs detailed)

**Interview Style:**
- Conversational, not robotic
- Context-aware follow-ups
- Respects your time constraints
- Focuses on what matters for your project

---

### Step 10: Compile Summary
**What gets compiled:**
- All answers organized by phase
- GitHub repository information (if provided)
- Content creation needs and preferences
- Keywords provided (if any)
- Content types selected
- Meeting speed selected

**Format:**
- Uses `.do/core/brainstorm/CONFIRMATION_TEMPLATE.md` format
- Well-structured, easy to review
- Clear sections and visual separators

---

### Step 11: Display Confirmation UI
**What you see:**
- Well-formatted markdown display
- Clear sections with checkmarks (✅)
- Blockquotes for longer answers
- Visual separators
- **Review & Confirm** section with 4 options:
  1. **Save it** - Proceed to save
  2. **Revise a phase** - Re-ask questions for specific phase(s)
  3. **Add information** - Add more details to appropriate phase
  4. **Start over** - Restart from speed selection

**Important:** 
- **DO NOT save until you explicitly confirm**
- You have full control
- Can revise any phase before saving

---

### Step 12: Handle User Response
**Based on your choice:**

1. **If confirmed (Save it)**:
   - Proceeds to save everything
   - Creates content folder structure
   - Saves meeting summary

2. **If revision requested**:
   - Re-asks questions for specified phase(s)
   - Updates summary
   - Shows updated summary again

3. **If addition requested**:
   - Adds information to appropriate phase
   - Shows updated summary

4. **If restart requested**:
   - Confirms intent
   - Restarts from speed selection if confirmed

---

### Step 13: Save to BRAINSTORM.md
**What gets saved:**
- Approved summary (organized by phase)
- Structure from `.do/core/brainstorm/TEMPLATE_BRAINSTORM.md`
- Meeting speed selected
- GitHub repository information
- Content creation needs and preferences
- Keywords and content requirements

**Location:** `.do/system/BRAINSTORM.md`

---

### Step 14: Update State
**What happens:**
- Sets `.do/system/history/active_state.json` phase to "brainstorm"
- Tracks meeting completion
- Enables next steps in workflow

---

### Step 15: Content Structure Creation
**What gets created:**
- `.do/system/content/` directory structure
- Organized folders by content type:
  - `app-pages/`
  - `legal/`
  - `social-media/`
  - `blog/`
  - `documentation/`
  - `marketing/`
  - `email/`
  - `seo/`
- Each folder includes README explaining purpose

**Ready for:**
- Content generation via `/content` command
- SEO-optimized content creation
- Organized content management

---

### Step 16: Final Response
**What you see:**
```
✅ Meeting complete! Summary saved to BRAINSTORM.md. 
Content structure created in .do/system/content/. 
Type /write to generate PRD, Architecture, and Design System. 
Type /content to start generating content.
```

---

## 🎨 Key Features

### 1. **Adaptive Intelligence**
- Questions adapt to your project type
- Not a one-size-fits-all questionnaire
- Dynamic question generation

### 2. **Time Flexibility**
- Choose your own pace
- Quick for simple projects
- Comprehensive for complex projects

### 3. **Content Planning**
- Integrated content creation planning
- SEO-ready by default
- Flexible generation options

### 4. **User Control**
- Confirm before saving
- Revise any phase
- Add information anytime
- Restart if needed

### 5. **Smart Defaults**
- Suggests appropriate speed
- Adapts to project type
- Focuses on what matters

---

## 🔄 Workflow Integration

**Before `/meeting`:**
- Type `/tell` to capture your idea

**After `/meeting`:**
- Type `/write` to generate PRD, Architecture, Design System
- Type `/content` to generate content based on meeting requirements
- Type `/good` to approve plan
- Type `/plan` to generate execution plan

---

## 💡 Best Practices

1. **Choose the right speed:**
   - Don't rush complex projects
   - Don't overthink simple projects

2. **Be specific about content needs:**
   - Think about what content you'll need
   - Consider SEO from the start

3. **Provide keywords if you have them:**
   - Helps AI create better content
   - Improves SEO optimization

4. **Review the summary carefully:**
   - Make sure everything is correct
   - Revise if needed before saving

5. **Use the confirmation step:**
   - Don't skip the review
   - Add missing information

---

## 🚀 What Makes It Special

1. **Not Pre-Made Questions**: Questions are dynamically generated based on YOUR project
2. **Adaptive**: Adjusts to project type, speed, and content needs
3. **Comprehensive**: Covers everything from vision to content to operations
4. **Flexible**: Multiple speed options and revision capabilities
5. **SEO-Ready**: Content planning includes SEO from the start
6. **User-Friendly**: Conversational, one question at a time, confirmation before saving

---

## 📊 Example Flow

**Scenario: Building a SaaS Startup**

1. **Project Type Detected**: SaaS/Startup
2. **Speed Suggested**: Comprehensive or Deep Dive
3. **You Choose**: Comprehensive
4. **GitHub Check**: Repository verified, workflows checked
5. **Content Needs**: 
   - Selected: App pages, Legal, Blog, Documentation, Marketing, Email, SEO
   - Mode: Keyword-guided for most, full LLM for legal
6. **Interview**: All 6 phases with condensed questions
7. **Questions Focus**: Business model, scalability, user acquisition, technical architecture
8. **Summary**: Comprehensive summary with all details
9. **Review**: You review and confirm
10. **Save**: Everything saved, content structure created
11. **Next**: Ready for `/write` and `/content` commands

---

**The `/meeting` command is your intelligent discovery partner that adapts to your needs, saves you time, and ensures nothing important is missed!**

