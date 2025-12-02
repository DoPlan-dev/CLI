# Memory Card Schema

Complete schema for the Memory Card JSON structure.

---

## 📁 Location

```
~/.doplan/memory_card.json
```

---

## 📊 Complete Schema

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

## 📝 Type Definitions

### ConversationEntry

```json
{
  "timestamp": "timestamp",
  "command": "string",
  "user_input": "string",
  "agent_response": "string",
  "sentiment": "positive|neutral|negative|frustrated|excited",
  "insight": "string",
  "duration": 0.0
}
```

### MemorableMoment

```json
{
  "timestamp": "timestamp",
  "type": "achievement|breakthrough|joke|challenge_overcome|first_time",
  "title": "string",
  "description": "string",
  "emotion": "happy|proud|excited|relieved|grateful",
  "context": "string"
}
```

### Achievement

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "earned_at": "timestamp",
  "category": "project|learning|productivity|milestone|score|streak|social",
  "project": "string",
  "points": 0,
  "rarity": "common|uncommon|rare|epic|legendary"
}
```

---

## 💡 Field Descriptions

### Relationship Metrics

- **tone_level** (0-10) - Affects formality
- **relationship_level** (0-100) - Overall relationship strength
- **trust_level** (0-10) - Trust in suggestions
- **engagement_score** (0-1) - Engagement level
- **score** (0-100000+) - Total achievement points

### Context Awareness

- **current_project** - Active project name
- **current_phase** - Current workflow phase
- **last_command** - Last command executed
- **last_command_time** - When last command was run
- **session_count** - Total sessions
- **average_session_time** - Average session duration (minutes)

---

**Wiki Complete!** 🎉

