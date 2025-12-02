# Brain System

The Brain System is the intelligence layer that reads your Memory Card and uses it to personalize every interaction with DoPlan. It's what makes DoPlan truly yours.

---

## 🧠 What is the Brain System?

The Brain System processes your Memory Card data and:
- **Personalizes agent responses** - Adapts to your style
- **Adjusts tone of voice** - Changes based on relationship
- **Provides context** - Remembers your projects and phases
- **Offers guidance** - Matches your experience level
- **Encourages you** - Personalized motivation

---

## 🏗️ Architecture

### Components

```
Memory Card (JSON)
    ↓
Brain (brain.go)
    ↓
AgentBrain (agent_brain.go)
    ↓
Personalized Responses
```

### 1. Brain (`brain.go`)

Core intelligence that:
- Reads memory card data
- Processes user context
- Generates personalized instructions
- Adjusts tone and style
- Provides behavioral guidance

### 2. AgentBrain (`agent_brain.go`)

Agent-specific wrapper that:
- Applies brain insights to agents
- Processes agent prompts
- Personalizes responses
- Tracks interactions
- Formats output

---

## 🎯 How Brain Works

### Step 1: Read Memory Card

Brain reads your Memory Card:
```go
memoryCard, err := LoadMemoryCard()
```

### Step 2: Process Context

Brain processes:
- User identity and preferences
- Communication preferences
- Learning goals and interests
- Relationship metrics
- Current context

### Step 3: Generate Instructions

Brain generates personalized instructions:
```go
instructions := brain.GetAgentInstructions("Product Manager")
```

**Example instructions**:
```
User Context and Preferences:
The user's name is Sarah.
The user's experience level is intermediate. Adjust technical depth accordingly.
The user prefers a fast work style. Be concise or detailed as appropriate.
The user's preferred tech stack includes: React, Node.js, PostgreSQL.
The current relationship level is 75/100, trust level is 8/10, and tone level is 7/10.
Adjust your tone and formality accordingly.
The current project is 'my-todo-app'.
The current phase is 'building'.
```

### Step 4: Apply to Agents

Instructions are applied to agent prompts:
```go
enhancedPrompt := brain.ProcessAgentPrompt(basePrompt, "Product Manager")
```

### Step 5: Personalize Responses

Agent responses are personalized:
```go
personalizedResponse := brain.ProcessAgentResponse(rawResponse, "/do")
```

---

## 🎨 Personalization Features

### 1. Tone of Voice Adjustment

Brain adjusts tone based on relationship level:

**Low Relationship (0-20)**:
- Formal language
- No contractions
- Professional structure
- Example: "Hello. I'm DoPlan. Let's begin."

**Moderate Relationship (20-40)**:
- Professional with hints of warmth
- Minimal contractions
- Structured responses
- Example: "Hello! I'm DoPlan. Let's get started."

**High Relationship (40-70)**:
- Friendly but professional
- Some contractions
- Occasional emojis
- Example: "Hey! 👋 I'm DoPlan. Ready to build something great?"

**Very High Relationship (70-100)**:
- Very warm and personal
- Full contractions
- Personal references
- More emojis
- Example: "Hey Sarah! 👋 Great to see you again! Ready to build something amazing?"

---

### 2. Detail Level Adaptation

Brain adjusts detail based on experience:

**Beginner**:
```
"I'll create a Git branch for you. A branch is like a separate workspace 
where you can make changes without affecting the main code. This is a 
best practice in software development."
```

**Intermediate**:
```
"Creating branch task/2.1. This isolates your changes and makes code 
review easier."
```

**Advanced**:
```
"Creating branch task/2.1. Ready for development."
```

---

### 3. Communication Style Matching

Brain matches your communication preferences:

**Brief Style**:
```
"Task complete. Next: /dev"
```

**Detailed Style**:
```
"✅ Task 2.1 marked complete!
   ⏱️  Task duration: 2h 15m
   ✓ Changes committed
   ✓ Changes pushed to task/2.1
   
   💡 Next steps:
      • Type /dev to start the next task
      • Type /status to see overall progress"
```

**Balanced Style**:
```
"✅ Task complete! (2h 15m)
   Next: /dev"
```

---

### 4. Encouragement Personalization

Brain provides personalized encouragement:

**Based on Motivation**:
- "change_world": "I love your ambition to change the world! 🌍"
- "money": "Let's build something profitable! 💰"

**Based on Interest**:
- "learning": "Great choice to focus on learning! 📚"

**Based on Achievements**:
- "Keep up the great work! I remember you recently earned 'Serial Builder'. You're doing fantastic! 🎉"

---

### 5. Context Awareness

Brain remembers:
- **Current project** - Which project you're working on
- **Current phase** - Where you are in workflow
- **Last command** - What you just did
- **Learning goals** - What you want to learn
- **Pain points** - Where you struggle

**Example**:
```
"I see you're working on 'my-todo-app' in the building phase.
I remember you're learning TypeScript. This task will help with that!
Also, I noticed you've had trouble with database migrations before.
Let me provide extra guidance for this task."
```

---

## 🎯 Brain Integration

### In Commands

Every command uses Brain:
- `/hey` - Personalized greeting
- `/do` - Adaptive conversation
- `/plan` - Context-aware planning
- `/dev` - Personalized messages
- `/done` - Encouragement

### In Agents

All agents receive:
- Enhanced prompts with user context
- Personalized instructions
- Tone-adjusted responses
- Context-aware suggestions

### In Engagement

Engagement system uses Brain for:
- Personalized achievement messages
- Context-aware challenges
- Relationship-based celebrations
- Encouragement matching

---

## 💡 Brain Examples

### Example 1: First Time User

**Memory Card**: Empty (new user)
**Brain Processing**: No personalization data
**Response**: "Hello! 👋 I'm DoPlan. Let's get started!"

### Example 2: After 10 Projects

**Memory Card**:
- Name: "Sarah"
- Experience: "intermediate"
- Relationship Level: 75
- Favorite Commands: ["/dev", "/done"]
- Current Project: "my-todo-app"

**Brain Processing**: Full personalization
**Response**: "Hey Sarah! 👋 Great to see you again! I know you love using /dev and /done. Ready to build something amazing on 'my-todo-app'?"

### Example 3: Learning Focus

**Memory Card**:
- Interest: "learning"
- Learning Goals: ["TypeScript", "Docker"]
- Experience: "beginner"

**Brain Processing**: Educational focus
**Response**: "I see you're learning TypeScript! This task will help you practice. Let me explain what we're doing and why..."

---

## 🔄 Brain Update Cycle

### Real-Time Updates

Brain reads Memory Card:
- **On every command** - Fresh data
- **After interactions** - Updated context
- **With state changes** - Current phase/project

### Processing Flow

1. Command executed
2. Brain loads Memory Card
3. Processes context
4. Generates instructions
5. Applies to agents
6. Personalizes responses
7. Updates Memory Card (if needed)

---

## 🎯 Brain Benefits

### For You

1. **Personal Experience** - Every interaction is yours
2. **Appropriate Guidance** - Matches your level
3. **Relevant Help** - Knows your pain points
4. **Encouraging** - Personalized motivation
5. **Context-Aware** - Remembers your projects

### For DoPlan

1. **Better Responses** - More relevant
2. **Appropriate Tone** - Matches relationship
3. **Effective Help** - Targeted assistance
4. **User Satisfaction** - Personalized experience
5. **Relationship Growth** - Builds connection

---

## 💡 Brain Tips

### Building Your Brain

- **Use DoPlan regularly** - Builds relationship
- **Complete projects** - Increases context
- **Set learning goals** - Educational focus
- **Have conversations** - Builds history
- **Earn achievements** - Tracks progress

### Understanding Brain

- **Relationship level** - Affects tone
- **Trust level** - Affects confidence
- **Experience level** - Affects detail
- **Communication style** - Affects format
- **Learning goals** - Affects guidance

---

**Next**: [Personalization](./04-Personalization.md)

