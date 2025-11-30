# Brain System - Memory Card Integration

## Overview

The Brain System is an intelligent layer that reads the memory card and uses it to influence agent behavior, responses, and tone of voice. It acts as the "brain" that makes agents truly personalized and adaptive.

## Architecture

```
Memory Card (JSON)
    ↓
Brain (brain.go)
    ↓
Agent Brain (agent_brain.go)
    ↓
Agent Responses (Personalized & Tone-Adjusted)
```

## Components

### 1. Brain (`brain.go`)
Core intelligence that:
- Reads and processes memory card data
- Enhances agent prompts with user context
- Adjusts tone of voice
- Personalizes responses
- Provides behavioral instructions

### 2. AgentBrain (`agent_brain.go`)
Agent-specific wrapper that:
- Processes agent prompts
- Processes agent responses
- Provides system context
- Tracks interactions
- Formats responses for user preferences

## Key Features

### Prompt Enhancement
Agents receive enhanced prompts that include:
- User experience level (affects explanation depth)
- Work style (affects response length)
- Personality (affects approach - thinker vs copier)
- Motivation (affects framing)
- Learning goals (educational context)
- Pain points (extra help areas)
- Tech stack preferences (solution prioritization)
- Current project/phase context
- Relationship level (warmth and personalization)
- Trust level (confidence in suggestions)

### Tone of Voice Adjustment
Automatically adjusts based on relationship level:

**Formal (Level 0-1)**
- No contractions
- Professional language
- Formal structure

**Professional (Level 2-4)**
- Minimal contractions
- Professional with hints of warmth
- Structured responses

**Warm (Level 5-7)**
- Some contractions
- Friendly but professional
- Occasional emojis

**Friendly (Level 8-10)**
- Full contractions
- Casual phrases
- Personal references
- More emojis
- References past interactions

### Response Personalization
- Adds personalized greetings
- Includes encouragement when appropriate
- References past achievements
- Provides contextual tips
- Adjusts length based on preferences
- Uses preferred communication style

## Usage Examples

### Basic Usage

```go
// Create brain instance
brain, err := NewBrain()
if err != nil {
    // Handle error
}

// Enhance agent prompt
basePrompt := "You are a helpful assistant..."
enhancedPrompt := brain.EnhanceAgentPrompt(basePrompt, "Product Manager")

// Personalize response
agentResponse := "Here's your plan..."
personalizedResponse := brain.PersonalizeResponse(agentResponse, "planning")

// Adjust tone
finalResponse := brain.AdjustToneOfVoice(personalizedResponse, "planning")
```

### Agent Brain Usage

```go
// Create agent brain
agentBrain, err := NewAgentBrain()
if err != nil {
    // Handle error
}

// Process agent prompt
enhancedPrompt := agentBrain.ProcessAgentPrompt(basePrompt, "Engineering Lead")

// Process agent response
personalizedResponse := agentBrain.ProcessAgentResponse(agentResponse, "development")

// Track interaction
agentBrain.TrackInteraction("/dev", "user input", "agent response", 45.2, "positive")
```

### Integration in Commands

```go
func runCommand(cmd *cobra.Command, args []string) error {
    // Initialize brain
    agentBrain, err := NewAgentBrain()
    if err != nil {
        return err
    }

    // Get enhanced prompt for agent
    basePrompt := getAgentPrompt("Product Manager")
    enhancedPrompt := agentBrain.ProcessAgentPrompt(basePrompt, "Product Manager")

    // Use enhanced prompt with agent
    agentResponse := callAgent(enhancedPrompt, userInput)

    // Process and personalize response
    personalizedResponse := agentBrain.ProcessAgentResponse(agentResponse, "command_context")

    // Track interaction
    agentBrain.TrackInteraction("/plan", userInput, personalizedResponse, duration, "positive")

    // Output personalized response
    fmt.Fprintln(cmd.OutOrStdout(), personalizedResponse)

    return nil
}
```

## Tone Adjustment Examples

### Formal (New User)
**Before:** "You can use this approach to solve the problem."
**After:** "You can use this approach to solve the problem."

### Professional (Building Relationship)
**Before:** "You can use this approach to solve the problem."
**After:** "You can use this approach to solve the problem."

### Warm (Good Relationship)
**Before:** "You can use this approach to solve the problem."
**After:** "You can use this approach to solve the problem."

### Friendly (Strong Relationship)
**Before:** "You can use this approach to solve the problem."
**After:** "You can use this approach to solve the problem. I remember you've worked on similar things before - this should feel familiar!"

## Prompt Enhancement Examples

### Beginner User
```
## User Context (From Memory Card)
- User's experience level: beginner
- IMPORTANT: Use simple language, avoid jargon, explain concepts clearly. Focus on 'what' not 'how'.
- User prefers working with: React, Node.js
- Current project: My First App
- Current phase: planning
```

### Advanced User
```
## User Context (From Memory Card)
- User's experience level: advanced
- Feel free to use technical terminology and discuss implementation details. Focus on 'how' and architecture.
- User prefers working with: Go, TypeScript, PostgreSQL
- Current project: Enterprise SaaS Platform
- Current phase: development
- You have a strong relationship with this user. Be warm, personal, and reference past interactions when relevant.
```

## Behavioral Instructions

The brain provides instructions like:

```
Keep responses concise and to the point.
Provide frequent feedback and check-ins.
When errors occur, be gentle and supportive.
Be enthusiastic and energetic in encouragement.
User frequently uses /dev - consider suggesting it when relevant.
User found these features helpful: time tracking, memory card. Reference them when appropriate.
```

## Response Formatting

Based on user preferences:

**Brief Communication Style:**
- Truncates long responses
- Gets to the point quickly
- Minimal explanations

**Detailed Communication Style:**
- Expands short responses
- Provides comprehensive context
- Includes examples

**Balanced Communication Style:**
- Medium length responses
- Key points with context
- Selective examples

## Integration Points

### Commands
- `/hey` - Uses brain for personalized onboarding
- `/do` - Uses brain for idea capture and refinement
- `/plan` - Uses brain for planning suggestions
- `/dev` - Uses brain for development guidance

### Agent Interactions
- All agent prompts are enhanced
- All agent responses are personalized
- All interactions are tracked
- Tone is automatically adjusted

## Benefits

### For Users
1. **Personalized Experience**: Every interaction feels tailored
2. **Consistent Tone**: Agent maintains appropriate relationship level
3. **Relevant Suggestions**: Based on past behavior and preferences
4. **Learning Support**: References learning goals and pain points
5. **Relationship Building**: Agent becomes more personal over time

### For Agents
1. **Better Context**: Understands user deeply
2. **Appropriate Communication**: Matches user's preferences
3. **Trust Building**: Adapts based on trust level
4. **Efficiency**: Knows when to be brief or detailed
5. **Personalization**: Can reference past interactions

## Technical Details

### Memory Card Reading
- Loaded once per command execution
- Cached in Brain instance
- Updated after each interaction
- Persisted to disk automatically

### Performance
- Minimal overhead (< 10ms per interaction)
- Efficient string processing
- Cached memory card data
- Lazy evaluation where possible

### Error Handling
- Gracefully handles missing memory card
- Falls back to defaults if data unavailable
- Continues working even if brain fails to load
- Logs errors without breaking flow

## Future Enhancements

Potential improvements:
- Real-time sentiment analysis
- Predictive suggestions
- Multi-project context switching
- Team collaboration features
- Advanced pattern recognition
- Machine learning integration

