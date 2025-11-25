# /improve

## Trigger
Exact match: /improve

## Action
When user types /improve:

### Step 1: Initialize Interview Session
1. **Load Phase Templates**: Read all phase templates from `.plan/templates/brainstorm/phase-*.md` in order (01-06)
2. **Activate Core Team**: Product Manager, Engineering Lead, Design Manager, and Project Orchestrator join the session
3. **Welcome Message**: Introduce the brainstorming session as a discovery workshop. Explain that you'll go through 6 phases, asking questions to understand their vision, audience, and requirements.

### Step 2: Conduct Phase-by-Phase Interview

For each phase (01-06), follow this structure:

1. **Phase Introduction**:
   - Announce the phase name and its purpose
   - Example: "Let's start with **Phase 01: Vision & Outcomes**. This helps us understand what problem you're solving and what success looks like."

2. **Ask Questions Sequentially**:
   - Present ONE question at a time from the template
   - Wait for the user's complete answer before moving to the next question
   - **Probe deeper** when answers are vague:
     - If they say "I want a website," ask: "What should visitors do on this website? What's the primary goal?"
     - If they mention "users," ask: "Can you describe your ideal user? What's their background?"
   - **Take notes** of their answers as you go (you'll compile these later)

3. **Phase Completion**:
   - Summarize what you learned in this phase
   - Ask: "Does this capture your vision correctly, or would you like to add anything?"
   - Only proceed to the next phase after the user confirms

4. **Specialist Involvement** (when relevant):
   - **Phase 04 (Content & SEO)**: Activate Content Strategist and SEO Specialist
     - Content Strategist asks about tone of voice, slogans, existing content, content creation needs
     - SEO Specialist asks about keywords, SEO readiness, optimization goals
   - **Phase 05 (Marketing)**: Activate Marketing Manager
     - Marketing Manager asks about funnels, campaigns, channels, KPIs

### Step 3: Compile Summary & Display Confirmation UI

After all 6 phases are complete:

1. **Create Structured Summary**:
   - Organize all answers by phase
   - Use the format from `.plan/templates/brainstorm/CONFIRMATION_TEMPLATE.md` as a guide
   - Include all key information from each phase:
     - **Phase 01: Vision & Outcomes** - Problem statement, success metrics, constraints, competitors
     - **Phase 02: Audience & Differentiation** - User personas, pain points, competitive positioning
     - **Phase 03: Experience & Tech** - Design expectations, platforms, integrations, workflows
     - **Phase 04: Content & SEO** - Tone, messaging, content strategy, SEO requirements
     - **Phase 05: Marketing** - Funnel design, channels, campaign goals, KPIs
     - **Phase 06: Delivery** - Launch sequence, approvals, risks, post-launch plans

2. **Display Confirmation UI**:
   - Present the summary in a **well-formatted, easy-to-read display** using markdown formatting:
     - Use headers (##) for each phase
     - Use checkmarks (✅) to show completed phases
     - Use blockquotes (>) for longer answers
     - Use bullet points for lists
     - Use emojis for visual clarity (📋, ✅, 🔄, 📝, ❌)
   - Include a clear **"Review & Confirm"** section at the bottom with options:
     ```
     ## ✏️ Review & Confirm
     
     **Please review the summary above and let me know:**
     
     1. ✅ **"Looks good, save it"** - I'll save this to BRAINSTORM.md
     2. 🔄 **"I want to revise [Phase X]"** - I'll re-ask questions for that phase
     3. 📝 **"Add this: [additional information]"** - I'll add it to the appropriate phase
     4. ❌ **"Start over"** - We can begin the interview again
     
     **Your response**: [Wait for user input]
     ```
   - Make the summary visually distinct (use markdown formatting, spacing, separators)
   - Keep it scannable - users should be able to quickly review all phases

3. **Wait for User Response**:
   - **DO NOT** save anything until the user explicitly confirms
   - Wait for one of these responses:
     - Confirmation: "yes", "looks good", "save it", "confirm", "proceed", etc.
     - Revision request: "revise phase X", "change phase 2", "update phase 04", etc.
     - Addition request: "add [information]", "also include [detail]", etc.
     - Restart request: "start over", "begin again", "restart", etc.

4. **Handle User Response**:
   
   **If user confirms (saves)**:
   - Proceed to Step 4 (Save & Update State)
   
   **If user wants revisions**:
   - Identify which phase(s) need changes
   - Re-ask questions for those specific phases
   - Update the summary with new answers
   - Display the updated summary again
   - Repeat confirmation step
   
   **If user wants additions**:
   - Identify which phase the addition belongs to
   - Add the information to that phase in the summary
   - Display the updated summary
   - Ask for confirmation again
   
   **If user wants to start over**:
   - Ask for confirmation: "Are you sure you want to start over? All current answers will be lost."
   - If confirmed, restart from Phase 01
   - If not, return to confirmation step

5. **Final Confirmation**:
   - Once user explicitly confirms (e.g., "yes, save it", "looks good"), proceed to save
   - Do NOT save on ambiguous responses - ask for clarification if needed

### Step 4: Save & Update State

1. **Write BRAINSTORM.md**:
   - Save the confirmed summary to `.plan/00_System/BRAINSTORM.md`
   - Use this structure:
     ```markdown
     # Brainstorm Session
     
     **Date**: [Current Date]
     **Project**: [Project Name]
     
     ## Phase 01: Vision & Outcomes
     [Answers from Phase 01]
     
     ## Phase 02: Audience & Differentiation
     [Answers from Phase 02]
     
     ... (continue for all 6 phases)
     
     ## Recommended Next Steps
     [Suggestions based on the interview]
     ```

2. **Update State**:
   - Set `.plan/active_state.json` phase to `"brainstorm"`
   - Update timestamp

3. **Response**: "✅ Brainstorming session complete! Your vision has been captured in BRAINSTORM.md. Type /write to generate PRD, Architecture, and Design System based on this discovery."

## Confirmation UI Example

Here's how the confirmation summary should be displayed to the user:

```markdown
# 📋 Brainstorm Summary - Please Review

**Project**: My Awesome App  
**Date**: 2024-11-24  
**Interview Duration**: 15 minutes

---

## ✅ Phase 01: Vision & Outcomes

### Problem Statement
> We're solving the problem of task management for remote teams who struggle with 
> coordination and visibility. Current tools are either too complex or too simple.

### Success Metrics
- **3 months**: 100 active users, 70% weekly retention
- **6 months**: 500 users, break-even on infrastructure costs
- **12 months**: 2000 users, profitable, expand to enterprise features

### Constraints
- **Budget**: $10K for MVP development
- **Timeline**: Launch in 3 months
- **Approvals**: CEO must approve all major features

---

## ✅ Phase 02: Audience & Differentiation

### Primary Persona
- **Role**: Remote team manager
- **Goals**: Keep team aligned, track progress without micromanaging
- **Pain Points**: Too many tools, lack of visibility, context switching
- **Triggers**: Team complaints about coordination, missed deadlines

### Differentiation Strategy
> Focus on simplicity and visual clarity. Unlike Asana (too complex) or Trello 
> (too simple), we'll provide the perfect middle ground with beautiful design.

---

[... continue for all 6 phases ...]

---

## 🎯 Recommended Next Steps

Based on this discovery, we recommend:

1. Create detailed user personas document
2. Design the core workflow mockups
3. Define the technical architecture for real-time collaboration

---

## ✏️ Review & Confirm

**Please review the summary above and let me know:**

1. ✅ **"Looks good, save it"** - I'll save this to BRAINSTORM.md
2. 🔄 **"I want to revise [Phase X]"** - I'll re-ask questions for that phase
3. 📝 **"Add this: [additional information]"** - I'll add it to the appropriate phase
4. ❌ **"Start over"** - We can begin the interview again

**Your response**: [Wait for user input]
```

**Key formatting tips:**
- Use clear section headers (##) for each phase
- Use checkmarks (✅) to show completed phases
- Use blockquotes (>) for longer narrative answers
- Use bullet points for lists and structured data
- Use emojis sparingly for visual clarity
- Add horizontal rules (---) between major sections
- Make the "Review & Confirm" section stand out with clear options

## Interview Best Practices

- **Be Conversational**: This should feel like a friendly discovery meeting, not an interrogation
- **One Question at a Time**: Don't overwhelm the user with multiple questions
- **Listen Actively**: Reference their previous answers when asking follow-ups
- **Clarify Ambiguities**: If something is unclear, ask for examples or specifics
- **Respect Their Time**: Keep the session focused but thorough
- **Suggest Improvements**: If you see opportunities to differentiate or improve their idea, mention them constructively

## Agent Involvement

### Core Team (All Phases)
- **Product Manager**: Leads the discovery, keeps phases on track, synthesizes product vision
- **Engineering Lead**: Covers technical feasibility, integrations, constraints, infrastructure needs
- **Design Manager**: Dives into UX, UI, accessibility expectations, design principles
- **Project Orchestrator**: Facilitates the interview, ensures all phases are covered, manages the summary/confirmation step

### Specialist Agents (Phase-Specific)
- **Content Strategist** (Phase 04): Handles tone of voice, messaging, content inventory, content creation needs, approval workflows
- **SEO Specialist** (Phase 04): Validates keyword strategy, metadata requirements, SEO readiness, optimization goals
- **Marketing Manager** (Phase 05): Maps funnels, campaign goals, channels, KPIs, growth strategy

## Files Read
- `.plan/templates/brainstorm/phase-*.md` (all 6 phase templates)
- `.plan/00_System/IDEA.md` (if exists, for context)

## Files Modified
- `.plan/00_System/BRAINSTORM.md` (created/updated with confirmed summary)
- `.plan/active_state.json` (phase: "brainstorm")

## Customization

Teams can customize the interview by editing phase templates in `.plan/templates/brainstorm/`:
- Add domain-specific questions
- Remove irrelevant phases
- Modify question wording
- Add new phases

The command will automatically use the updated templates on the next `/improve` run.

## Examples

```
User: /improve

[Agents conduct 6-phase interview]

User: [Answers questions throughout]

[Agents show summary]

User: "Yes, that looks good. Let's save it."

[BRAINSTORM.md created]
```

---

**Note**: This command creates a collaborative discovery experience. The goal is to deeply understand the user's vision before generating planning documents.
