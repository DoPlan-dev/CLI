# Personalization

Personalization is how DoPlan adapts to you, making every interaction feel like it's designed specifically for you. It's powered by the Memory Card and Brain System.

---

## 🎯 What is Personalization?

Personalization means DoPlan:
- **Remembers you** - Your name, preferences, goals
- **Adapts to you** - Communication style, detail level
- **Learns from you** - Tech preferences, pain points
- **Grows with you** - Relationship builds over time

---

## 🎨 Personalization Aspects

### 1. Tone of Voice

DoPlan's tone changes based on your relationship level:

**Starting (Level 0-20)**:
```
"Hello. I'm DoPlan. Let's begin by capturing your project idea."
```
- Formal language
- No contractions
- Professional structure

**Developing (Level 20-40)**:
```
"Hello! I'm DoPlan. Let's get started with your project idea."
```
- Professional with warmth
- Minimal contractions
- Friendly structure

**Building (Level 40-70)**:
```
"Hey! 👋 I'm DoPlan. Ready to capture your project idea?"
```
- Friendly but professional
- Some contractions
- Occasional emojis

**Strong (Level 70-100)**:
```
"Hey [Your Name]! 👋 Great to see you again! Ready to build something amazing?"
```
- Very warm and personal
- Full contractions
- Personal references
- More emojis

---

### 2. Detail Level

DoPlan adjusts detail based on your experience:

**Beginner**:
```
"I'll create a Git branch for you. A branch is like a separate workspace 
where you can make changes without affecting the main code. This is a 
best practice in software development that helps teams work together 
without conflicts."
```

**Intermediate**:
```
"Creating branch task/2.1. This isolates your changes and makes code 
review easier. You can work on this feature independently."
```

**Advanced**:
```
"Creating branch task/2.1. Ready for development."
```

---

### 3. Communication Style

DoPlan matches your preferred communication style:

**Brief**:
```
"Task complete. Next: /dev"
```

**Detailed**:
```
"✅ Task 2.1 marked complete!
   ⏱️  Task duration: 2h 15m
   ✓ Changes committed: feat(task-2.1): complete user authentication
   ✓ Changes pushed to task/2.1
   
   💡 Next steps:
      • Type /dev to start the next task
      • Type /status to see overall progress
      • Type /sys engagement to view achievements"
```

**Balanced**:
```
"✅ Task complete! (2h 15m)
   Committed and pushed.
   Next: /dev"
```

---

### 4. Encouragement Style

DoPlan provides encouragement that matches your style:

**Enthusiastic**:
```
"🎉 Amazing work! You're crushing it! Keep up the fantastic progress! 🚀"
```

**Supportive**:
```
"Great work! You're making excellent progress. Keep it up!"
```

**Professional**:
```
"Task completed successfully. Good progress on the project."
```

---

### 5. Error Handling

DoPlan handles errors based on your preference:

**Gentle**:
```
"I notice there was an issue. Don't worry, let's work through this together. 
Here's what happened and how we can fix it..."
```

**Direct**:
```
"Error: Branch not found. Run '/dev' first to create a branch."
```

**Educational**:
```
"Error: Branch not found. This happens when you try to complete a task 
without starting one. Let me explain: when you run '/dev', DoPlan creates 
a branch for your task. You need to run '/dev' before '/done'. 
Would you like to start a task now?"
```

---

## 🎯 Personalization Examples

### Example 1: Greeting

**First Time**:
```
"Hello! 👋 I'm DoPlan, your AI development partner. Let's get started!"
```

**After 10 Projects**:
```
"Hey Sarah! 👋 Great to see you again! Ready to build something amazing?"
```

**After 50 Projects**:
```
"Hey Sarah! 👋 It's great to see you again! I remember you're working on 
React projects. Ready to build something amazing together?"
```

---

### Example 2: Task Completion

**New User**:
```
"✅ Task complete! Good work!"
```

**Experienced User**:
```
"✅ Task 2.1 marked complete!
   ⏱️  Task duration: 2h 15m
   I remember you prefer fast workflows, so here's a quick summary:
   • Committed and pushed
   • Next: /dev
   
   You're doing great! I noticed you've completed 10 tasks this week. 
   That's impressive! 🎉"
```

---

### Example 3: Learning Support

**Beginner Learning TypeScript**:
```
"I see you're learning TypeScript! This task will help you practice. 
Let me explain what we're doing:
- We're creating a type-safe API endpoint
- TypeScript helps catch errors before runtime
- This is a great learning opportunity!

I'll provide extra guidance throughout this task."
```

**Advanced User**:
```
"Creating TypeScript API endpoint. Type-safe implementation."
```

---

## 🔄 How Personalization Works

### Step 1: Memory Card Read

Every command reads your Memory Card:
```go
memoryCard, err := LoadMemoryCard()
```

### Step 2: Brain Processing

Brain processes your data:
```go
instructions := brain.GetAgentInstructions("Product Manager")
```

### Step 3: Response Generation

Response is generated with personalization:
```go
personalizedResponse := brain.ProcessAgentResponse(rawResponse, "/do")
```

### Step 4: Display

Personalized response is displayed to you.

---

## 🎯 Personalization Factors

### Relationship Level

- **0-20**: Formal, professional
- **20-40**: Professional with warmth
- **40-70**: Friendly, helpful
- **70-100**: Very warm, personal

### Experience Level

- **Beginner**: Detailed explanations, educational
- **Intermediate**: Balanced, helpful
- **Advanced**: Concise, efficient

### Work Style

- **Fast**: Quick, efficient responses
- **Thoughtful**: Detailed, comprehensive responses

### Communication Preferences

- **Brief**: Short, to the point
- **Detailed**: Comprehensive, thorough
- **Balanced**: Moderate detail

---

## 💡 Personalization Benefits

### For You

1. **Feels Personal** - Like it's designed for you
2. **Appropriate Level** - Matches your experience
3. **Relevant Help** - Knows your pain points
4. **Encouraging** - Personalized motivation
5. **Efficient** - Matches your work style

### For DoPlan

1. **Better Experience** - More relevant responses
2. **User Satisfaction** - Personalized interaction
3. **Relationship Growth** - Builds connection
4. **Effective Help** - Targeted assistance
5. **Engagement** - Keeps you coming back

---

## 🚀 Building Personalization

### Use DoPlan Regularly

- **Builds relationship** - Increases relationship level
- **Learns preferences** - Tracks your style
- **Remembers context** - Knows your projects
- **Adapts tone** - Gets more personal

### Set Preferences

- **Experience level** - Set in `/hey`
- **Work style** - Detected from usage
- **Communication style** - Learned from interactions
- **Learning goals** - Set explicitly

### Complete Projects

- **Increases context** - More to remember
- **Builds relationship** - Stronger connection
- **Tracks patterns** - Learns your style
- **Records achievements** - Celebrates with you

---

## 💡 Personalization Tips

### For Beginners

- **Set experience level** - In `/hey` tutorial
- **Ask questions** - DoPlan learns from you
- **Use regularly** - Builds relationship
- **Set learning goals** - Educational focus

### For Intermediate

- **Complete projects** - Builds context
- **Explore features** - DoPlan learns preferences
- **Use all commands** - Comprehensive learning
- **Track progress** - See personalization grow

### For Advanced

- **Fast workflow** - DoPlan adapts
- **Power features** - Advanced personalization
- **Customization** - Tailor to your needs
- **Relationship max** - Best Friends Forever

---

**Next**: [Relationship Building](./05-Relationship-Building.md)

