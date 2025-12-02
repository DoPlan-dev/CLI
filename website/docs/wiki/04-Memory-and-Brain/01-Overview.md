# Memory and Brain Overview

The Memory and Brain System work together to make DoPlan **truly personal**. They learn about you, remember your preferences, and adapt every interaction to your style.

---

## 🧠 What is the Memory Card?

The **Memory Card** is your personal relationship file stored at `~/.doplan/memory_card.json`. It's like DoPlan's memory of you - it remembers everything about your relationship.

### What It Stores

- **Your identity** - Name, experience level, preferences
- **Your personality** - Work style, interests, dreams
- **Your communication** - Preferred style, detail level
- **Your learning** - Goals, tech stack, pain points
- **Your relationship** - Relationship level, trust, engagement
- **Your history** - Conversations, achievements, memorable moments
- **Your usage** - Command patterns, favorite features

### Key Point: It Moves With You

The Memory Card is stored in your **home directory**, not in project folders. This means:
- ✅ Works across all projects
- ✅ Builds relationship over time
- ✅ Remembers you everywhere
- ✅ Grows with you

---

## 🧠 What is the Brain System?

The **Brain System** reads your Memory Card and uses it to:
- **Personalize agent responses** - Adapts to your style
- **Adjust tone of voice** - Changes based on relationship
- **Provide context** - Remembers your projects and phases
- **Offer guidance** - Matches your experience level
- **Encourage you** - Personalized motivation

### How It Works

```
Memory Card (JSON)
    ↓
Brain (Processes Data)
    ↓
Agent Brain (Applies to Agents)
    ↓
Personalized Responses
```

---

## 🎯 How They Work Together

### Example: First Time User

**Memory Card**: Empty (new user)
**Brain**: No personalization data
**Response**: "Hello! 👋 I'm DoPlan. Let's get started!"

### Example: After 10 Projects

**Memory Card**: 
- Name: "Sarah"
- Experience: "intermediate"
- Relationship Level: 75
- Favorite Commands: ["/dev", "/done"]

**Brain**: Processes all this data
**Response**: "Hey Sarah! 👋 Great to see you again! I know you love using /dev and /done. Ready to build something amazing?"

---

## 💡 Key Benefits

### For You

1. **Personal Connection** - DoPlan knows you
2. **Adaptive Experience** - Matches your style
3. **Learning Support** - Understands your goals
4. **Relationship Growth** - Gets better over time
5. **Context Awareness** - Remembers your projects

### For DoPlan

1. **Better Responses** - Personalized to you
2. **Appropriate Tone** - Matches relationship
3. **Relevant Guidance** - Based on experience
4. **Effective Encouragement** - Personalized motivation
5. **Contextual Help** - Remembers your pain points

---

## 🎯 Memory Card Location

Your Memory Card is stored at:
```
~/.doplan/memory_card.json
```

**Why home directory?**
- Works across all projects
- Builds relationship over time
- Persistent across sessions
- Personal to you

---

## 🧠 Brain System Components

### 1. Brain (`brain.go`)

Core intelligence that:
- Reads memory card data
- Processes user context
- Generates personalized instructions
- Adjusts tone and style

### 2. AgentBrain (`agent_brain.go`)

Agent-specific wrapper that:
- Applies brain insights to agents
- Processes agent prompts
- Personalizes responses
- Tracks interactions

### 3. Integration

Used throughout DoPlan:
- Command responses
- Agent interactions
- Achievement messages
- Error handling
- Encouragement

---

## 📊 Relationship Metrics

### Relationship Level (0-100)

Tracks overall relationship strength:
- **0-20**: Starting - Formal communication
- **20-40**: Developing - Professional with warmth
- **40-70**: Building - Friendly and helpful
- **70-100**: Strong! - Very warm and personal

**How it increases**:
- Regular usage
- Project completion
- Command usage
- Time spent together

### Trust Level (0-10)

How much you trust DoPlan's suggestions:
- **0-3**: Low trust - More cautious
- **4-6**: Moderate trust - Balanced
- **7-9**: High trust - Confident suggestions
- **10**: Complete trust - Maximum confidence

**How it increases**:
- Successful commands
- Positive outcomes
- Reliable suggestions

### Tone Level (0-10)

Affects communication formality:
- **0-2**: Formal - Professional language
- **3-5**: Professional - Balanced
- **6-8**: Warm - Friendly
- **9-10**: Friendly - Very casual

**How it increases**:
- Frequent interaction
- Regular usage
- Relationship building

---

## 🎯 Personalization Examples

### Tone of Voice

**Low Relationship (Level 10)**:
```
"Hello. I'm DoPlan. Let's begin by capturing your project idea."
```

**High Relationship (Level 80)**:
```
"Hey [Name]! 👋 Great to see you again! Ready to build something amazing? 
I remember you're working on React projects. Let's capture your next idea!"
```

### Detail Level

**Beginner**:
```
"I'll create a Git branch for you. A branch is like a separate workspace 
where you can make changes without affecting the main code. This is a 
best practice in software development."
```

**Advanced**:
```
"Creating branch task/2.1. Ready for development."
```

### Encouragement

**First Project**:
```
"Great start! You're on your way to building something amazing!"
```

**After 10 Projects**:
```
"I remember you recently earned 'Serial Builder'. You're doing fantastic! 
Keep up the great work! 🎉"
```

---

## 🚀 Getting Started

### Your First Interaction

When you first use DoPlan:
1. Memory Card is created automatically
2. Basic information collected
3. Relationship starts at 0
4. Brain begins learning

### Building Relationship

As you use DoPlan:
1. Memory Card updates automatically
2. Relationship level increases
3. Brain learns your preferences
4. Responses become more personal

### Viewing Your Memory Card

Your Memory Card is stored at:
```
~/.doplan/memory_card.json
```

You can view it (if you have access):
```bash
cat ~/.doplan/memory_card.json
```

---

## 💡 Key Concepts

### Memory Card
- **Location**: `~/.doplan/memory_card.json`
- **Scope**: Global (all projects)
- **Updates**: Automatic
- **Purpose**: Remember you

### Brain System
- **Purpose**: Personalize responses
- **Input**: Memory Card data
- **Output**: Personalized instructions
- **Integration**: All commands and agents

### Relationship
- **Metric**: 0-100 relationship level
- **Growth**: Over time with usage
- **Impact**: Affects tone and style
- **Benefit**: More personal experience

---

## 📚 Next Steps

1. **[Memory Card](./02-Memory-Card.md)** - Complete Memory Card guide
2. **[Brain System](./03-Brain-System.md)** - How Brain works
3. **[Personalization](./04-Personalization.md)** - How DoPlan adapts
4. **[Relationship Building](./05-Relationship-Building.md)** - Growing with DoPlan

---

**Ready to learn more?** → [Memory Card](./02-Memory-Card.md)

