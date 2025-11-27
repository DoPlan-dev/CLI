# Question Variation System

## Overview

The meeting system uses **question variation** to ensure each meeting feels fresh and conversational, even when covering the same topics. Every question has multiple phrasings that are randomly selected each time.

## Why Variation?

### Problem with Static Questions
- Same question, same wording every meeting → Feels robotic
- Users might feel like they're filling out a form
- No conversational flow or natural feel
- Repetitive experience discourages repeat meetings

### Solution: Question Variations
- **Same core question, different phrasings** → Feels natural
- **Random selection** → Each meeting is unique
- **Adaptive follow-ups** → Can use different variation if answer is vague
- **Conversational flow** → Feels like talking to a person, not a form

## How It Works

### 1. Question Template Structure

Each question in the phase templates has multiple variations:

```markdown
### Question 1: Problem Identification

**Variations:**
1. "What problem are you trying to solve?"
2. "What's the main issue you want to address?"
3. "What challenge is your project solving?"
4. "What problem does this project tackle?"
5. "What's the pain point you're addressing?"

**Follow-ups:**
- What pain point does this address?
- Why is this problem important?
- Can you give me an example?
- How does this affect your users?
```

### 2. Random Selection Process

When the meeting runs:

1. **Load question template** from `{experience}/{project-type}/phase-*.md`
2. **For each question:**
   - Randomly select one variation (e.g., variation #3)
   - Store selection in `meeting_session.json`
   - Ask the selected variation
3. **If user gives vague answer:**
   - Use a different variation as follow-up
   - Probe deeper with alternative phrasing

### 3. Variation Tracking

The system tracks which variations were used:

```json
{
  "meeting_session": {
    "start_time": "2025-01-15T14:30:00Z",
    "questions_asked": {
      "phase-01": {
        "question-1": {
          "variation_used": 3,
          "phrasing": "What challenge is your project solving?",
          "timestamp": "2025-01-15T14:32:15Z"
        },
        "question-2": {
          "variation_used": 1,
          "phrasing": "What do you want to build?",
          "timestamp": "2025-01-15T14:33:45Z"
        }
      }
    }
  }
}
```

### 4. Avoiding Repetition

On repeat meetings:
- System checks previous meeting sessions
- Prefers variations not used recently
- Ensures different phrasings even for same user

## Example: Same Question, Different Meetings

### Meeting 1 (First Time)
```
DoPlan: "What problem are you trying to solve?"
User: "People can't find good restaurants nearby"
```

### Meeting 2 (Repeat Meeting)
```
DoPlan: "What's the main issue you want to address?"
User: "People can't find good restaurants nearby"
```

### Meeting 3 (Another Repeat)
```
DoPlan: "What challenge is your project solving?"
User: "People can't find good restaurants nearby"
```

**Same core question, different phrasings, fresh experience each time!**

## Question Order Variation (Optional)

In addition to phrasing variation, questions within each phase can be shuffled:

### Standard Order
1. Problem identification
2. Project vision
3. Success metrics
4. Must-have features
5. Differentiation

### Shuffled Order (Example)
1. Project vision
2. Must-have features
3. Problem identification
4. Differentiation
5. Success metrics

This adds another layer of variation while maintaining logical flow.

## Variation Guidelines

### Creating Variations

1. **Keep core meaning** - All variations should ask the same thing
2. **Vary sentence structure** - Different ways to phrase the same question
3. **Match experience level** - Beginner variations are simpler, advanced are more technical
4. **Match project type** - Include project-specific context when relevant
5. **Natural language** - Sound conversational, not robotic

### Good Variations
```
✅ "What problem are you trying to solve?"
✅ "What's the main issue you want to address?"
✅ "What challenge is your project solving?"
✅ "What problem does this project tackle?"
```

### Bad Variations (Too Different)
```
❌ "What problem are you trying to solve?"
❌ "What's your favorite color?"  ← Completely different question
❌ "How old are you?"  ← Not related
```

## Implementation Details

### Template Format

```markdown
### Question 1: [Question Topic]

**Variations:**
1. "[Variation 1]"
2. "[Variation 2]"
3. "[Variation 3]"
4. "[Variation 4]"
5. "[Variation 5]"

**Follow-ups:**
- [Follow-up question 1]
- [Follow-up question 2]
- [Follow-up question 3]

**Context Notes:**
- When to use this question
- What to probe for
- Project-type-specific considerations
```

### Selection Algorithm

```go
// Pseudocode
func selectVariation(question Question, previousMeetings []MeetingSession) int {
    // Get available variations
    variations := question.Variations
    
    // Filter out recently used variations
    recentVariations := getRecentVariations(previousMeetings, question.ID)
    availableVariations := filterOut(recentVariations, variations)
    
    // If all variations used recently, use all
    if len(availableVariations) == 0 {
        availableVariations = variations
    }
    
    // Randomly select from available
    return randomSelect(availableVariations)
}
```

## Benefits

✅ **Fresh Experience** - Each meeting feels different  
✅ **Natural Conversation** - Multiple phrasings feel human, not robotic  
✅ **Adaptive Follow-ups** - Can rephrase if answer is vague  
✅ **Maintains Consistency** - Same topics covered, different wording  
✅ **Encourages Repeat Use** - Users won't feel like filling the same form  
✅ **Better Engagement** - Varied questions keep users interested

## Future Enhancements

1. **Context-Aware Variations** - Select variations based on user's previous answers
2. **Tone Adaptation** - Match variation tone to user's communication style
3. **Learning from Feedback** - Track which variations get better responses
4. **A/B Testing** - Test which variations work best for different user types
5. **Dynamic Generation** - AI-generated variations based on base templates (hybrid approach)

---

**Generated by: DoPlan CLI v1.2.0**

**Sub-Agent: Documentation Lead**

**Date: 2025-01-15**

