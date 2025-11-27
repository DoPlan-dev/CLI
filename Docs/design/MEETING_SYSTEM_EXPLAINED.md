# Meeting System Explained

## Current Implementation

### How It Works Now

The `/meeting` command uses a **hybrid approach** combining:

1. **Question Template Libraries** (`.do/core/brainstorm/{experience}/{project-type}/phase-*.md`)
   - Pre-written question templates organized by experience level and project type
   - Each question has multiple variations/phrasings
   - Created during project initialization
   - Located in `.do/core/brainstorm/`

2. **Dynamic Question Selection & Variation** (Runtime)
   - **Random Variation**: Each question has 3-5 different phrasings, randomly selected each meeting
   - **Question Order**: Questions within a phase can be shuffled (optional)
   - **Follow-up Depth**: Adaptive based on user answers
   - **Project Type Adaptation**: Questions tailored to detected project type
   - **Meeting Speed Filtering**: Phases and questions filtered by speed selection
   - **Content Needs**: Additional questions added based on content requirements

### Current Flow

```
User types /meeting
    ↓
1. Detect project type from IDEA.md
    ↓
2. Ask user experience level
    ↓
3. Filter speed options based on experience
    ↓
4. User selects speed
    ↓
5. Load question templates from .do/core/brainstorm/{experience}/{project-type}/phase-*.md
    ↓
6. **For each question, randomly select a variation**:
   - Each question has 3-5 different phrasings
   - Random selection ensures each meeting feels different
   - Same core question, different wording
    ↓
7. **Optional: Shuffle question order** (within each phase)
   - Keeps meetings fresh even with same questions
    ↓
8. **Adaptive follow-ups**:
   - Probe deeper based on user answers
   - Skip irrelevant questions based on context
   - Add project-specific follow-ups
    ↓
9. Conduct interview phase by phase
    ↓
10. Save to BRAINSTORM.md
```

## Current Limitations

### Problem: Static Questions

The current phase templates contain **static questions** that are asked the same way every time. For example:

**Current Phase 01 Question (Static):**
```
"What are the key success metrics?"
- How will we measure success?
- What KPIs matter most?
```

**Problems:**
- Same question, same wording every meeting
- No variation or freshness
- Users might feel like they're filling out a form
- Doesn't adapt to conversation flow

### Problem: Language Adaptation is Manual

The meeting command says:
> "Adapt question complexity based on experience level:
> - Beginner: Use simple language, avoid technical jargon
> - Intermediate: Mix of simple and technical terms
> - Advanced: Technical terminology, architecture discussions"

But this adaptation happens **at runtime** by the AI agent, not from pre-written question libraries.

## Proposed Solution: Experience-Level + Project-Type Question Libraries with Variation

### Structure

Create **question libraries organized by both experience level AND project type**, with **multiple variations** for each question:

```
.do/core/brainstorm/
├── beginner/
│   ├── website/          ← Website/Agency/Personal/Blog projects
│   ├── saas/             ← SaaS/Startup/Enterprise projects
│   ├── mobile/           ← Cross-platform mobile apps
│   ├── ios/              ← iOS/iPhone apps (Swift/Objective-C)
│   ├── android/          ← Android apps (Java/Kotlin)
│   ├── webapp/           ← Web applications/PWA/SPA
│   ├── desktop/          ← Cross-platform desktop apps
│   ├── windows/          ← Windows desktop apps (.NET/WinUI/WPF)
│   ├── macos/            ← macOS desktop apps (Swift/Objective-C)
│   ├── linux/            ← Linux desktop apps (GTK/Qt)
│   ├── cli/              ← Command-line interface tools
│   ├── library/          ← Code libraries/packages/SDKs
│   ├── framework/        ← Frameworks/plugins/extensions
│   ├── game/             ← Game development projects
│   ├── embedded/         ← Embedded systems/IoT
│   ├── api/              ← API/Backend services
│   ├── microservice/     ← Microservices architecture
│   ├── patch-windows/    ← Windows patches/updates
│   ├── patch-macos/      ← macOS patches/updates
│   ├── patch-linux/      ← Linux patches/updates
│   ├── patch-web/        ← Web app patches/updates
│   ├── devops/           ← DevOps/CI/CD tools
│   ├── data-science/     ← Data science/ML projects
│   ├── cloud/            ← Cloud-native applications
│   └── general/          ← Generic/Other projects (fallback)
├── intermediate/
│   ├── website/
│   ├── saas/
│   ├── mobile/
│   ├── webapp/
│   ├── desktop/
│   ├── cli/
│   ├── library/
│   ├── patch/
│   └── general/
└── advanced/
    ├── website/
    ├── saas/
    ├── mobile/
    ├── webapp/
    ├── desktop/
    ├── cli/
    ├── library/
    ├── patch/
    └── general/
```

### Why Both Experience Level AND Project Type?

**Different project types need different questions:**

- **Website**: "What pages do you need?" (landing, about, services, contact)
- **SaaS**: "What's your pricing model?" (freemium, subscription, usage-based)
- **Mobile App**: "Which platforms?" (iOS, Android, both)
- **CLI Tool**: "What commands will it have?" (command structure, help system)
- **Library**: "What's the API design?" (function signatures, versioning)

**Same project type, different experience levels need different language:**

- **Beginner + Website**: "What pages do you want on your website?"
- **Advanced + Website**: "What's the information architecture and page hierarchy?"

### Example: Phase 01 - Vision & Outcomes

#### Beginner + Website Version
```markdown
# Phase 01: Vision & Outcomes
**Led by: Product Manager**

## Purpose
Let's figure out what you want to build and why it matters.

## Key Questions

1. **What problem are you trying to solve?**
   - What's bothering you or your users?
   - Why does this problem need to be fixed?
   - Can you give me an example?

2. **What do you want to build?**
   - Describe your idea in simple terms
   - What will people see when they use it?
   - What's the main thing it does?

3. **How will you know it's working?**
   - What would make you happy with the result?
   - How many people should use it?
   - What's a good sign that people like it?

4. **What absolutely must work?**
   - What features can't be missing?
   - What would make this project fail if it's not there?
   - What's the most important thing?

5. **What makes your idea special?**
   - How is it different from other similar things?
   - Why would someone choose yours instead of something else?
```

#### Intermediate + Website Version
```markdown
# Phase 01: Vision & Outcomes
**Led by: Product Manager**

## Purpose
Define the core vision, success metrics, and desired outcomes for the project.

## Key Questions

1. **What is the primary problem we're solving?**
   - What pain point does this address?
   - Why is this problem important?
   - What's the current workaround?

2. **What is the vision for this product?**
   - What does success look like in 6 months? 1 year?
   - What impact do we want to make?
   - What's the long-term goal?

3. **What are the key success metrics?**
   - How will we measure success?
   - What KPIs matter most? (e.g., user count, engagement, revenue)
   - What are the target numbers?

4. **What are the must-have outcomes?**
   - What absolutely must be achieved?
   - What would make this project a failure if missing?
   - What are the non-negotiables?

5. **What is the unique value proposition?**
   - What makes this different from existing solutions?
   - Why would users choose this over alternatives?
   - What's our competitive advantage?
```

#### Advanced + Website Version
```markdown
# Phase 01: Vision & Outcomes
**Led by: Product Manager**

## Purpose
Define the core vision, success metrics, and desired outcomes for the project with strategic depth.

## Key Questions

1. **What is the primary problem we're solving?**
   - What specific pain point or market gap does this address?
   - What's the root cause analysis?
   - What are the current alternatives and their limitations?
   - What's the market opportunity size?

2. **What is the strategic vision for this product?**
   - What does success look like at 3, 6, 12, and 24 months?
   - What's the strategic positioning in the market?
   - What's the long-term vision and potential pivots?
   - How does this align with business objectives?

3. **What are the key success metrics and KPIs?**
   - What are the North Star metrics?
   - How will we measure product-market fit?
   - What are the leading vs. lagging indicators?
   - What's the target conversion funnel?
   - What are the technical performance benchmarks?

4. **What are the must-have outcomes and success criteria?**
   - What are the critical success factors (CSFs)?
   - What would constitute a failure condition?
   - What are the minimum viable outcomes (MVOs)?
   - What are the go/no-go criteria?

5. **What is the unique value proposition and competitive differentiation?**
   - What's the unique selling proposition (USP)?
   - How do we differentiate from direct and indirect competitors?
   - What's our competitive moat?
   - What's the value proposition canvas?
```

### How It Works

1. **User selects experience level** → System knows which library to use
2. **Load appropriate phase templates** → Beginner/intermediate/advanced versions
3. **Further adapt based on:**
   - Project type (add project-specific questions)
   - Meeting speed (skip questions for Quick Start)
   - Content needs (add content-related questions)

### Benefits

✅ **Pre-written, tested questions** - No reliance on AI to adapt language  
✅ **Consistent core questions** - Same topics covered, different phrasings  
✅ **Fresh every time** - Random variation selection makes each meeting unique  
✅ **Clear language differences** - Beginners get simple explanations  
✅ **Professional depth** - Advanced users get technical discussions  
✅ **Maintainable** - Easy to update question libraries  
✅ **Extensible** - Can add project-type-specific variants  
✅ **Conversational** - Multiple phrasings feel natural, not robotic  
✅ **Adaptive follow-ups** - Can use different variation if answer is vague

## Implementation Plan

### Step 1: Create Question Library Structure
- Generate three versions of each phase template
- Organize by experience level in `.do/core/brainstorm/{level}/`

### Step 2: Update Meeting Command
- Modify `/meeting` to load from appropriate library based on user experience
- Keep dynamic adaptation for project type and speed

### Step 3: Project Type Variants (Already in Structure)
- Project-type-specific folders are already part of the structure
- Each experience level has all project type variants
- Questions are tailored to both experience level AND project type

### Step 4: Meeting Speed Filtering
- Load questions from `{experience}/{project-type}/phase-*.md`
- Filter phases based on speed:
  - Quick Start: Phases 01, 03 only
  - Standard: Phases 01, 02, 03, 06
  - Comprehensive: All phases, condensed questions
  - Deep Dive: All phases, full questions
- Add content-related questions dynamically if user needs content

## Example: Complete Question Flow

### Scenario 1: Beginner building a website

1. **System detects**: Project type = Website (from IDEA.md)
2. **User selects**: Beginner experience, Quick Start speed
3. **System loads**: `.do/core/brainstorm/beginner/website/phase-01.md`
   - Questions are already website-specific AND beginner-friendly
   - No need to add project-specific questions (already included)
4. **System filters**: Only phases 01 and 03 (Quick Start = 2 phases)
5. **System uses**: Simple language, website-focused questions

**Example questions loaded:**
- "What pages do you want on your website?" (beginner + website)
- "What should people see on the home page?" (beginner + website)
- "Do you need a contact form?" (beginner + website)

### Scenario 2: Advanced building a SaaS

1. **System detects**: Project type = SaaS (from IDEA.md)
2. **User selects**: Advanced experience, Deep Dive speed
3. **System loads**: `.do/core/brainstorm/advanced/saas/phase-01.md`
   - Questions are already SaaS-specific AND advanced-level
   - Technical terminology, deep strategic questions
4. **System includes**: All 6 phases with full questions
5. **System uses**: Technical terminology, SaaS-specific concepts

**Example questions loaded:**
- "What's the monetization strategy?" (advanced + saas)
- "What's the target LTV:CAC ratio?" (advanced + saas)
- "What's the product-market fit hypothesis?" (advanced + saas)

### Scenario 3: Intermediate building a mobile app

1. **System detects**: Project type = Mobile App (from IDEA.md)
2. **User selects**: Intermediate experience, Standard speed
3. **System loads**: `.do/core/brainstorm/intermediate/mobile/phase-*.md`
   - Questions are mobile-specific with balanced language
4. **System includes**: Phases 01, 02, 03, 06 (Standard = 4 phases)
5. **System uses**: Mix of simple and technical terms, mobile-focused

**Example questions loaded:**
- "Which platforms will you support?" (intermediate + mobile)
- "What are the key user flows on mobile?" (intermediate + mobile)
- "What are the performance requirements for mobile?" (intermediate + mobile)

## Current vs. Proposed

| Aspect | Current | Proposed |
|--------|---------|----------|
| **Question Source** | Generic templates + AI adaptation | Pre-written libraries by experience level |
| **Language Adaptation** | AI adapts at runtime | Pre-written versions for each level |
| **Consistency** | Varies by AI interpretation | Consistent for same experience level |
| **Maintainability** | Hard to update | Easy to update question libraries |
| **Project Type** | Dynamic adaptation | Pre-written project-type-specific libraries |
| **Speed Filtering** | Dynamic | Dynamic (same as now) |

## Next Steps

1. **Create question library structure** in `generateBrainstormTemplates()`
   - Create folders: `beginner/{project-types}/`, `intermediate/{project-types}/`, `advanced/{project-types}/`
   
2. **Generate question libraries with variations** for each combination:
   - 3 experience levels × 25+ project types = 75+ sets of phase templates
   - Each set has 6 phases = 450+ phase template files total
   - **Each question has 3-5 variations** stored in the template
   - Or use a hybrid: base templates + project-type overlays (more maintainable)

3. **Update meeting command** to:
   - Detect project type from IDEA.md
   - Load from `{experience}/{project-type}/phase-*.md`
   - **Randomly select question variations** for each question
   - **Track selected variations** in meeting session to avoid repetition
   - **Optional: Shuffle question order** within each phase
   - Filter phases based on meeting speed

4. **Variation Selection Logic**:
   - For each question, randomly pick one variation
   - Store selection in `meeting_session.json` to track what was asked
   - If user gives vague answer, use a different variation as follow-up
   - On repeat meetings, prefer variations not used recently

5. **Alternative: Hybrid Approach** (More Maintainable):
   - Base questions with variations in `{experience}/general/phase-*.md`
   - Project-type-specific additions in `{experience}/{project-type}/phase-*.md`
   - System merges base + project-specific questions at runtime
   - Reduces from 450+ files to ~90 files (base + overlays)

6. **Test variation system**:
   - Run same meeting twice, verify different phrasings
   - Test with Beginner + Website
   - Test with Advanced + SaaS
   - Ensure questions feel natural and varied

---

**Generated by: DoPlan CLI v1.2.0**

**Sub-Agent: Documentation Lead**

**Date: 2025-01-15**

