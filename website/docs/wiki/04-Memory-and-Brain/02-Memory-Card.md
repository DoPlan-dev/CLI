# Memory Card

The Memory Card is your personal relationship file that DoPlan uses to remember you, your preferences, and your journey. It's stored in your home directory and works across all your projects.

---

## 📁 What is the Memory Card?

The Memory Card is a JSON file stored at:
```
~/.doplan/memory_card.json
```

It's like DoPlan's memory of you - it remembers everything about your relationship, preferences, and journey.

---

## 🎯 Key Feature: It Moves With You

The Memory Card is stored in your **home directory**, not in project folders. This means:
- ✅ **Works across all projects** - Same memory everywhere
- ✅ **Builds relationship over time** - Grows with usage
- ✅ **Remembers you everywhere** - Consistent experience
- ✅ **Personal to you** - Your unique relationship

---

## 📊 What's Stored in the Memory Card?

### 1. User Identity

```json
{
  "user_name": "Your Name",
  "first_met": "2025-01-15T10:00:00Z",
  "projects_count": 5,
  "last_project_date": "2025-01-20T14:30:00Z",
  "last_interaction": "2025-01-20T15:00:00Z"
}
```

**Tracks**:
- Your name
- When you first met DoPlan
- Number of projects
- Last project date
- Last interaction time

---

### 2. User Preferences & Personality

```json
{
  "interest": "learning",  // or "develop"
  "work_style": "fast",     // or "thoughtful"
  "personality": "thinker", // or "copier"
  "dream": "change_world",  // or "build_others"
  "motivation": "money",    // or "success"
  "experience_level": "intermediate"  // "beginner", "intermediate", "advanced"
}
```

**Tracks**:
- Your primary interest
- Your work style preference
- Your personality type
- Your long-term dream
- Your motivation
- Your experience level

---

### 3. Communication Preferences

```json
{
  "communication_style": "brief",      // "brief", "detailed", "balanced"
  "feedback_frequency": "frequent",     // "frequent", "moderate", "minimal"
  "detail_level": "high",               // "high", "medium", "low"
  "encouragement_style": "enthusiastic", // "enthusiastic", "supportive", "professional"
  "error_handling_pref": "gentle"      // "gentle", "direct", "educational"
}
```

**Tracks**:
- How you prefer to communicate
- How often you want feedback
- How much detail you want
- What encouragement style works
- How you prefer error handling

---

### 4. Learning & Preferences

```json
{
  "preferred_tech_stack": ["React", "Node.js", "PostgreSQL"],
  "project_types": ["Web App", "API", "CLI"],
  "interests": ["Frontend", "Backend", "DevOps"],
  "learning_goals": ["TypeScript", "Docker", "Kubernetes"],
  "pain_points": ["Database migrations", "Deployment"]
}
```

**Tracks**:
- Technologies you prefer
- Types of projects you work on
- Topics you're interested in
- What you want to learn
- Where you struggle

---

### 5. Relationship Data

```json
{
  "conversation_history": [
    {
      "timestamp": "2025-01-15T10:00:00Z",
      "command": "/do",
      "user_input": "I want to build a todo app",
      "agent_response": "Great idea! Tell me more...",
      "sentiment": "positive",
      "insight": "User prefers web apps",
      "duration": 120.5
    }
  ],
  "memorable_moments": [
    {
      "timestamp": "2025-01-15T10:00:00Z",
      "type": "achievement",
      "title": "First Steps",
      "description": "Completed first task",
      "emotion": "excited",
      "context": "First project"
    }
  ],
  "achievements": [
    {
      "id": "first_steps",
      "title": "First Steps",
      "description": "Complete your first task",
      "earned_at": "2025-01-15T10:00:00Z",
      "category": "project",
      "points": 50,
      "rarity": "common"
    }
  ]
}
```

**Tracks**:
- Conversation history (last 50 entries)
- Memorable moments (last 20)
- All achievements earned
- Flexible preferences storage

---

### 6. Usage Patterns

```json
{
  "command_usage": {
    "/hey": 1,
    "/do": 5,
    "/plan": 3,
    "/dev": 25,
    "/done": 20
  },
  "favorite_commands": ["/dev", "/done"],
  "struggled_features": ["Database migrations"],
  "helpful_features": ["Auto-commit", "Time tracking"],
  "time_preferences": {
    "development": "morning",
    "planning": "afternoon"
  }
}
```

**Tracks**:
- How often each command is used
- Your favorite commands
- Features you struggled with
- Features you found helpful
- Preferred times for activities

---

### 7. Challenge Tracking

```json
{
  "completed_challenges": [
    "api_integration_master",
    "first_launch"
  ],
  "challenge_attempts": {
    "database_merger": 2,
    "production_deploy": 1
  }
}
```

**Tracks**:
- Completed challenges
- Challenge attempt counts
- Learning from failures

---

### 8. Relationship Metrics

```json
{
  "tone_level": 7,           // 0-10, affects formality
  "relationship_level": 75,   // 0-100, overall relationship strength
  "trust_level": 8,          // 0-10, how much you trust suggestions
  "engagement_score": 0.85,   // 0-1, how engaged you are
  "score": 1250               // Total achievement score
}
```

**Tracks**:
- Tone level (formality)
- Relationship strength
- Trust in DoPlan
- Engagement level
- Total score

---

### 9. Context Awareness

```json
{
  "current_project": "my-todo-app",
  "current_phase": "building",
  "last_command": "/dev",
  "last_command_time": "2025-01-20T15:00:00Z",
  "session_count": 15,
  "average_session_time": 45.5  // minutes
}
```

**Tracks**:
- Current project name
- Current workflow phase
- Last command executed
- When last command was run
- Total sessions
- Average session duration

---

## 🔄 How the Memory Card Updates

### Automatic Updates

The Memory Card updates automatically when you:
- Use commands
- Complete tasks
- Earn achievements
- Complete challenges
- Have conversations
- Set preferences

### Manual Updates

You can also update it through:
- `/hey` command (first-time setup)
- Conversation interactions
- Preference settings
- Learning goal setting

---

## 📊 Memory Card Structure

### Complete Schema

```json
{
  // User Identity
  "user_name": "string",
  "first_met": "timestamp",
  "projects_count": 0,
  "last_project_date": "timestamp",
  "last_interaction": "timestamp",
  
  // Preferences & Personality
  "interest": "learning|develop",
  "work_style": "fast|thoughtful",
  "personality": "thinker|copier",
  "dream": "change_world|build_others",
  "motivation": "money|success",
  "experience_level": "beginner|intermediate|advanced",
  
  // Communication Preferences
  "communication_style": "brief|detailed|balanced",
  "feedback_frequency": "frequent|moderate|minimal",
  "detail_level": "high|medium|low",
  "encouragement_style": "enthusiastic|supportive|professional",
  "error_handling_pref": "gentle|direct|educational",
  
  // Learning & Preferences
  "preferred_tech_stack": ["string"],
  "project_types": ["string"],
  "interests": ["string"],
  "learning_goals": ["string"],
  "pain_points": ["string"],
  
  // Relationship Data
  "conversation_history": [ConversationEntry],
  "memorable_moments": [MemorableMoment],
  "achievements": [Achievement],
  "preferences": {},
  
  // Usage Patterns
  "command_usage": {},
  "favorite_commands": ["string"],
  "struggled_features": ["string"],
  "helpful_features": ["string"],
  "time_preferences": {},
  
  // Challenge Tracking
  "completed_challenges": ["string"],
  "challenge_attempts": {},
  
  // Relationship Metrics
  "tone_level": 0,
  "relationship_level": 0,
  "trust_level": 5,
  "engagement_score": 0.0,
  "score": 0,
  
  // Context Awareness
  "current_project": "string",
  "current_phase": "string",
  "last_command": "string",
  "last_command_time": "timestamp",
  "session_count": 0,
  "average_session_time": 0.0
}
```

---

## 🎯 How Memory Card is Used

### By Commands

Every command:
- Reads memory card for personalization
- Updates memory card with usage
- Tracks interactions
- Updates metrics

### By Brain System

The Brain System:
- Reads memory card data
- Generates personalized instructions
- Adjusts tone and style
- Provides context

### By Engagement System

The Engagement System:
- Tracks achievements
- Updates score
- Records challenges
- Builds relationship

---

## 💡 Memory Card Benefits

### For You

1. **Personal Experience** - DoPlan knows you
2. **Consistent Across Projects** - Same memory everywhere
3. **Relationship Growth** - Builds over time
4. **Learning Support** - Tracks your goals
5. **Progress Tracking** - See your journey

### For DoPlan

1. **Better Responses** - Personalized to you
2. **Appropriate Tone** - Matches relationship
3. **Relevant Guidance** - Based on experience
4. **Effective Help** - Knows your pain points
5. **Context Awareness** - Remembers your projects

---

## 🔒 Privacy & Security

### Location

Stored in your home directory:
```
~/.doplan/memory_card.json
```

### Access

- Only DoPlan reads/writes it
- You can view it (if you have access)
- Not shared with anyone
- Local to your machine

### Data

- All data is local
- No cloud sync (currently)
- You control the data
- Can be deleted/reset

---

## 🚀 Viewing Your Memory Card

### View File

```bash
# View your memory card
cat ~/.doplan/memory_card.json
```

### Format with jq (if available)

```bash
# Pretty print
cat ~/.doplan/memory_card.json | jq .

# View specific fields
cat ~/.doplan/memory_card.json | jq .relationship_level
cat ~/.doplan/memory_card.json | jq .achievements
```

---

## 💡 Memory Card Tips

### Building Your Memory Card

- **Use DoPlan regularly** - Builds relationship
- **Complete projects** - Increases project count
- **Earn achievements** - Tracks progress
- **Set learning goals** - Tracks education
- **Have conversations** - Builds history

### Understanding Your Data

- **Relationship level** - How strong your connection is
- **Trust level** - How much you trust DoPlan
- **Score** - Your total achievement points
- **Usage patterns** - How you use DoPlan
- **Learning goals** - What you want to learn

---

**Next**: [Brain System](./03-Brain-System.md)

