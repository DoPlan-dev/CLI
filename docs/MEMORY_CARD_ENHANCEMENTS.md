# Memory Card System Enhancements

## Overview
The memory card system has been significantly enhanced to build stronger, more personalized relationships between users and the DoPlan agent. The system now tracks detailed user preferences, learning patterns, achievements, and memorable moments to create a truly adaptive and personal experience.

## Key Enhancements

### 1. **Enhanced User Profiling**

#### Communication Preferences
- **Communication Style**: `brief`, `detailed`, or `balanced`
- **Feedback Frequency**: `frequent`, `moderate`, or `minimal`
- **Detail Level**: `high`, `medium`, or `low`
- **Encouragement Style**: `enthusiastic`, `supportive`, or `professional`
- **Error Handling Preference**: `gentle`, `direct`, or `educational`

#### Learning & Interests
- **Learning Goals**: Tracks what the user wants to learn
- **Pain Points**: Records common challenges the user faces
- **Interests**: Topics the user is interested in
- **Preferred Tech Stack**: Technologies the user prefers working with

### 2. **Relationship Building**

#### Conversation History
- **Structured Entries**: Each conversation is now a structured `ConversationEntry` with:
  - Timestamp
  - Command used
  - User input
  - Agent response
  - Sentiment analysis (`positive`, `neutral`, `negative`, `frustrated`, `excited`)
  - Key insights learned
  - Duration

#### Memorable Moments
- Tracks special moments in the relationship:
  - **Types**: `achievement`, `breakthrough`, `joke`, `challenge_overcome`, `first_time`
  - **Emotions**: `happy`, `proud`, `excited`, `relieved`, `grateful`
  - **Context**: What was happening when the moment occurred
  - Last 50 memorable moments are preserved

#### Achievements
- Tracks user milestones and accomplishments:
  - **Categories**: `project`, `learning`, `productivity`, `milestone`
  - **Associated Project**: Links achievements to specific projects
  - **Earned Date**: When the achievement was earned
  - Last 100 achievements are preserved

### 3. **Usage Pattern Analysis**

#### Command Usage Tracking
- Tracks how often each command is used
- Automatically identifies **Favorite Commands** (top 5 most used)
- Helps agent understand user's workflow preferences

#### Feature Interaction
- **Struggled Features**: Features the user had trouble with
- **Helpful Features**: Features the user found helpful
- Helps agent provide better guidance and support

#### Time Preferences
- Tracks preferred times for different activities
- Enables time-aware greetings and suggestions

### 4. **Relationship Metrics**

#### Relationship Level (0-100)
Calculated based on:
- Tone level (up to 50 points)
- Interaction frequency (up to 20 points)
- Memorable moments (2 points each)
- Achievements (1 point each)

#### Trust Level (0-10)
- Starts at 5 (neutral)
- Increases with successful interactions
- Decreases with negative experiences

#### Engagement Score (0-1)
- Based on average time between sessions
- < 24 hours = 1.0 (very engaged)
- < 1 week = 0.7 (engaged)
- < 1 month = 0.4 (moderate)
- > 1 month = 0.2 (low)

#### Tone Level (0-10)
- Gradually increases with each interaction
- Affects formality and warmth of communication
- Higher levels = more casual and friendly

### 5. **Context Awareness**

#### Current State Tracking
- **Current Project**: Active project name
- **Current Phase**: Current workflow phase (`ideation`, `planning`, `development`, etc.)
- **Last Command**: Most recently executed command
- **Last Command Time**: When last command was run
- **Session Count**: Total number of sessions
- **Average Session Time**: Average session duration

### 6. **Personalized Communication**

#### Smart Greetings
- **Time-aware**: Different greetings for morning, afternoon, evening, night
- **Returning User**: Special messages for users returning after a break
- **Relationship-based**: More personal for high relationship levels
- **Context-aware**: Adapts to current situation

#### Personalized Encouragement
- References past achievements
- Adapts to user's motivation (`change_world`, `money`, `learning`)
- Relationship level affects warmth and personalization

#### Contextual Messages
- `first_project`: Special message for first-time users
- `returning_after_break`: Welcome back messages
- `milestone_reached`: Celebration messages
- `struggling`: Supportive messages during challenges
- `high_engagement`: Recognition for active users

#### Personalized Tips
- Based on pain points
- References learning goals
- Mentions favorite commands
- Contextual to current situation

### 7. **Helper Functions**

The `MemoryCardHelper` provides convenient methods:

- **TrackCommandExecution**: Records full command context with sentiment
- **RecordSuccess**: Records successful interactions
- **RecordStruggle**: Tracks challenges and pain points
- **RecordHelpfulFeature**: Notes features user finds helpful
- **RecordLearningGoal**: Tracks learning objectives
- **RecordProjectStart**: Records new project initiation
- **RecordPhaseTransition**: Tracks workflow progress
- **RecordMilestone**: Records project milestones
- **GetPersonalizedResponse**: Returns context-appropriate messages
- **ShouldUseDetailedExplanation**: Determines explanation depth
- **ShouldProvideFrequentFeedback**: Determines feedback frequency
- **GetEncouragementStyle**: Returns appropriate encouragement style
- **GetErrorHandlingStyle**: Returns appropriate error handling approach

## Integration Points

### Commands Using Memory Card

1. **`/hey`** - Onboarding and tutorial
   - Records first meeting
   - Captures user preferences
   - Creates memorable moment for first interaction

2. **`/do`** - Idea capture and refinement
   - Tracks ideation patterns
   - Records project start
   - Learns from user's idea style

3. **`/plan`** - Planning phase
   - Records planning preferences
   - Tracks planning style (quick/detailed)
   - Learns project types

4. **`/dev`** - Development phase
   - Tracks tech stack preferences
   - Records feature work
   - Learns development patterns

### Time Tracking Integration

All commands now record:
- Command execution time
- Phase duration
- Session duration
- Time between interactions

## Benefits

### For Users
1. **Personalized Experience**: Agent remembers preferences and adapts communication
2. **Relationship Building**: Agent becomes more friendly and personal over time
3. **Achievement Recognition**: Milestones and accomplishments are celebrated
4. **Learning Support**: Agent remembers learning goals and provides relevant tips
5. **Context Awareness**: Agent understands current project state and phase

### For Agent
1. **Better Guidance**: Understands user's experience level and preferences
2. **Adaptive Communication**: Adjusts tone, detail level, and feedback frequency
3. **Pattern Recognition**: Learns from user behavior to provide better suggestions
4. **Relationship Metrics**: Tracks relationship strength to personalize interactions
5. **Context Understanding**: Knows current project, phase, and recent activity

## Future Enhancements

Potential future improvements:
- Sentiment analysis from user input
- Predictive suggestions based on patterns
- Collaborative filtering (learn from similar users)
- Mood detection and adaptation
- Long-term goal tracking
- Habit formation support
- Progress visualization
- Relationship milestones and celebrations

## Technical Details

### File Location
- **Path**: `~/.doplan/memory_card.json`
- **Format**: JSON with indentation
- **Size**: Typically < 50KB even with extensive history

### Data Retention
- **Conversation History**: Last 100 entries
- **Memorable Moments**: Last 50 moments
- **Achievements**: Last 100 achievements
- **Pain Points**: Last 20 pain points
- **Learning Goals**: Last 10 goals

### Privacy & Security
- Stored locally in user's home directory
- Never transmitted to external servers
- User has full control over data
- Can be deleted or reset at any time

## Usage Examples

### Basic Usage
```go
// Load memory card
card, err := LoadMemoryCard()
if err != nil {
    // Handle error
}

// Get personalized greeting
greeting := card.GetGreeting()

// Track command execution
helper := NewMemoryCardHelper(card)
helper.TrackCommandExecution("/plan", "user input", "agent response", 45.2, "positive")

// Record achievement
helper.RecordMilestone("First Feature Complete", "Successfully implemented user authentication")

// Save updates
SaveMemoryCard(card)
```

### Advanced Usage
```go
// Record project start
helper.RecordProjectStart("My Awesome App", "SaaS")

// Track phase transition
helper.RecordPhaseTransition("planning", "development")

// Record learning goal
helper.RecordLearningGoal("Learn React hooks")

// Get personalized response
response := helper.GetPersonalizedResponse("milestone_reached")

// Check if detailed explanation needed
if helper.ShouldUseDetailedExplanation() {
    // Provide detailed explanation
}
```

## Conclusion

The enhanced memory card system transforms DoPlan from a tool into a true development partner. By building a deep understanding of each user, the agent can provide increasingly personalized, helpful, and supportive guidance throughout the development journey.

