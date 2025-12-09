package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
)

// MemoryCard represents the persistent user-agent relationship data
type MemoryCard struct {
	// User Identity
	UserName        string    `json:"user_name,omitempty"`
	FirstMet        time.Time `json:"first_met,omitempty"`
	ProjectsCount   int       `json:"projects_count,omitempty"`
	LastProjectDate time.Time `json:"last_project_date,omitempty"`
	LastInteraction time.Time `json:"last_interaction,omitempty"` // Last time user interacted with DoPlan

	// User Preferences & Personality
	Interest        string `json:"interest,omitempty"`         // "learning" or "develop"
	WorkStyle       string `json:"work_style,omitempty"`       // "fast" or "thoughtful"
	Personality     string `json:"personality,omitempty"`      // "thinker" or "copier"
	Dream           string `json:"dream,omitempty"`            // "change_world" or "build_others"
	Motivation      string `json:"motivation,omitempty"`       // "money" or "success"
	ExperienceLevel string `json:"experience_level,omitempty"` // "beginner", "intermediate", "advanced"

	// Communication Preferences
	CommunicationStyle string `json:"communication_style,omitempty"` // "brief", "detailed", "balanced"
	FeedbackFrequency  string `json:"feedback_frequency,omitempty"`  // "frequent", "moderate", "minimal"
	DetailLevel        string `json:"detail_level,omitempty"`        // "high", "medium", "low"
	EncouragementStyle string `json:"encouragement_style,omitempty"` // "enthusiastic", "supportive", "professional"
	ErrorHandlingPref  string `json:"error_handling_pref,omitempty"` // "gentle", "direct", "educational"

	// Learning & Preferences
	PreferredTechStack []string `json:"preferred_tech_stack,omitempty"`
	ProjectTypes       []string `json:"project_types,omitempty"`
	Interests          []string `json:"interests,omitempty"`            // Topics user is interested in
	LearningGoals      []string `json:"learning_goals,omitempty"`       // What user wants to learn
	PainPoints         []string `json:"pain_points,omitempty"`          // Common challenges user faces
	ResolvedPainPoints []string `json:"resolved_pain_points,omitempty"` // Pain points that have been resolved

	// Relationship Data
	ConversationHistory []ConversationEntry    `json:"conversation_history,omitempty"` // Structured conversation history
	MemorableMoments    []MemorableMoment      `json:"memorable_moments,omitempty"`    // Special moments to remember
	Achievements        []Achievement          `json:"achievements,omitempty"`         // User milestones and wins
	Preferences         map[string]interface{} `json:"preferences,omitempty"`          // Flexible preferences storage

	// Usage Patterns
	CommandUsage      map[string]int    `json:"command_usage,omitempty"`      // How often each command is used
	FavoriteCommands  []string          `json:"favorite_commands,omitempty"`  // Commands user prefers
	StruggledFeatures []string          `json:"struggled_features,omitempty"` // Features user had trouble with
	HelpfulFeatures   []string          `json:"helpful_features,omitempty"`   // Features user found helpful
	TimePreferences   map[string]string `json:"time_preferences,omitempty"`   // Preferred times for different activities

	// Challenge Tracking
	CompletedChallenges []string       `json:"completed_challenges,omitempty"` // List of completed challenge IDs
	ChallengeAttempts   map[string]int `json:"challenge_attempts,omitempty"`   // Number of attempts per challenge

	// Relationship Metrics
	ToneLevel         int     `json:"tone_level,omitempty"`         // 0-10, increases with usage, affects formality and warmth
	RelationshipLevel int     `json:"relationship_level,omitempty"` // 0-100, overall relationship strength
	TrustLevel        int     `json:"trust_level,omitempty"`        // 0-10, how much user trusts agent suggestions
	EngagementScore   float64 `json:"engagement_score,omitempty"`   // 0-1, how engaged user is
	Score             int     `json:"score,omitempty"`              // Total achievement score

	// Context Awareness
	CurrentProject     string    `json:"current_project,omitempty"`      // Active project name
	CurrentPhase       string    `json:"current_phase,omitempty"`        // Current workflow phase
	LastCommand        string    `json:"last_command,omitempty"`         // Last command executed
	LastCommandTime    time.Time `json:"last_command_time,omitempty"`    // When last command was run
	SessionCount       int       `json:"session_count,omitempty"`        // Total number of sessions
	AverageSessionTime float64   `json:"average_session_time,omitempty"` // Average session duration in minutes

	// Streak Tracking
	LastUsageDate time.Time `json:"last_usage_date,omitempty"` // Last date user used DoPlan (date only, no time)
	CurrentStreak int       `json:"current_streak,omitempty"`  // Current consecutive days streak
	LongestStreak int       `json:"longest_streak,omitempty"`  // Longest streak ever achieved
	UsageDates    []string  `json:"usage_dates,omitempty"`     // Dates when user used DoPlan (YYYY-MM-DD format)
}

// ConversationEntry represents a structured conversation interaction
type ConversationEntry struct {
	Timestamp     time.Time `json:"timestamp,omitempty"`
	Command       string    `json:"command,omitempty"`
	UserInput     string    `json:"user_input,omitempty"`
	AgentResponse string    `json:"agent_response,omitempty"`
	Sentiment     string    `json:"sentiment,omitempty"` // "positive", "neutral", "negative", "frustrated", "excited"
	Insight       string    `json:"insight,omitempty"`   // Key insight learned from this interaction
	Duration      float64   `json:"duration,omitempty"`  // Duration in seconds
}

// MemorableMoment represents a special moment in the user-agent relationship
type MemorableMoment struct {
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Type        string    `json:"type,omitempty"` // "achievement", "breakthrough", "joke", "challenge_overcome", "first_time"
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Emotion     string    `json:"emotion,omitempty"` // "happy", "proud", "excited", "relieved", "grateful"
	Context     string    `json:"context,omitempty"` // What was happening when this occurred
}

// Achievement represents a user milestone or accomplishment
type Achievement struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	EarnedAt    time.Time `json:"earned_at,omitempty"`
	Category    string    `json:"category,omitempty"` // "project", "learning", "productivity", "milestone"
	Project     string    `json:"project,omitempty"`  // Associated project if applicable
}

// LoadMemoryCard loads the memory card from user's home directory
func LoadMemoryCard() (*MemoryCard, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	cardPath := filepath.Join(homeDir, ".doplan", "memory_card.json")

	// If doesn't exist, create new one
	if !utils.PathExists(cardPath) {
		card := &MemoryCard{
			FirstMet:          time.Now(),
			LastInteraction:   time.Now(),
			ToneLevel:         0,
			RelationshipLevel: 0,
			TrustLevel:        5, // Start with neutral trust
			EngagementScore:   0.0,
			SessionCount:      0,
			Score:             0, // Start with 0 score
			Preferences:       make(map[string]interface{}),
			CommandUsage:      make(map[string]int),
			TimePreferences:   make(map[string]string),
		}
		if err := SaveMemoryCard(card); err != nil {
			return nil, err
		}
		return card, nil
	}

	data, err := os.ReadFile(cardPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read memory card: %w", err)
	}

	var card MemoryCard
	if err := json.Unmarshal(data, &card); err != nil {
		return nil, fmt.Errorf("failed to parse memory card: %w", err)
	}

	// Initialize nil maps/slices
	if card.Preferences == nil {
		card.Preferences = make(map[string]interface{})
	}
	if card.CommandUsage == nil {
		card.CommandUsage = make(map[string]int)
	}
	if card.TimePreferences == nil {
		card.TimePreferences = make(map[string]string)
	}
	if card.ConversationHistory == nil {
		card.ConversationHistory = []ConversationEntry{}
	}
	if card.MemorableMoments == nil {
		card.MemorableMoments = []MemorableMoment{}
	}
	if card.Achievements == nil {
		card.Achievements = []Achievement{}
	}
	if card.PreferredTechStack == nil {
		card.PreferredTechStack = []string{}
	}
	if card.ProjectTypes == nil {
		card.ProjectTypes = []string{}
	}
	if card.Interests == nil {
		card.Interests = []string{}
	}
	if card.LearningGoals == nil {
		card.LearningGoals = []string{}
	}
	if card.PainPoints == nil {
		card.PainPoints = []string{}
	}
	if card.FavoriteCommands == nil {
		card.FavoriteCommands = []string{}
	}
	if card.StruggledFeatures == nil {
		card.StruggledFeatures = []string{}
	}
	if card.HelpfulFeatures == nil {
		card.HelpfulFeatures = []string{}
	}
	if card.CompletedChallenges == nil {
		card.CompletedChallenges = []string{}
	}
	if card.ChallengeAttempts == nil {
		card.ChallengeAttempts = make(map[string]int)
	}
	if card.ResolvedPainPoints == nil {
		card.ResolvedPainPoints = []string{}
	}
	if card.UsageDates == nil {
		card.UsageDates = []string{}
	}

	return &card, nil
}

// SaveMemoryCard saves the memory card to user's home directory
func SaveMemoryCard(card *MemoryCard) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	cardDir := filepath.Join(homeDir, ".doplan")
	if err := utils.CreateDirectory(cardDir); err != nil {
		return fmt.Errorf("failed to create .doplan directory: %w", err)
	}

	cardPath := filepath.Join(cardDir, "memory_card.json")

	// Update interaction tracking
	card.LastInteraction = time.Now()

	// Update streak tracking
	card.UpdateStreak()

	// Only increment project count if this is a new project
	// (This should be called explicitly when starting a new project, not on every save)

	// Update relationship metrics
	card.UpdateRelationshipMetrics()

	data, err := json.MarshalIndent(card, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal memory card: %w", err)
	}

	// Invalidate cache after saving
	err = utils.WriteFile(cardPath, data)
	if err == nil {
		invalidateMemoryCardCache()
	}
	return err
}

// UpdateFromConversation updates memory card based on user responses
func (m *MemoryCard) UpdateFromConversation(insights map[string]string) {
	// Update tone level (increases with each interaction, but more slowly)
	if m.ToneLevel < 10 && len(insights) > 0 {
		// Increment by 0.1 per interaction, but only show integer part
		m.ToneLevel = int(float64(m.ToneLevel) + 0.1)
		if m.ToneLevel > 10 {
			m.ToneLevel = 10
		}
	}

	// Store structured conversation entry
	entry := ConversationEntry{
		Timestamp: time.Now(),
		Insight:   fmt.Sprintf("%v", insights),
	}
	if cmd, ok := insights["command"]; ok {
		entry.Command = cmd
		m.RecordCommandUsage(cmd)
	}
	if input, ok := insights["user_input"]; ok {
		entry.UserInput = input
	}
	if sentiment, ok := insights["sentiment"]; ok {
		entry.Sentiment = sentiment
	} else {
		entry.Sentiment = "neutral" // Default sentiment
	}

	m.ConversationHistory = append(m.ConversationHistory, entry)
	// Keep only last 100 entries
	if len(m.ConversationHistory) > 100 {
		m.ConversationHistory = m.ConversationHistory[len(m.ConversationHistory)-100:]
	}

	// Update relationship metrics
	m.UpdateRelationshipMetrics()
}

// RecordCommandUsage tracks which commands the user uses
func (m *MemoryCard) RecordCommandUsage(command string) {
	if m.CommandUsage == nil {
		m.CommandUsage = make(map[string]int)
	}
	m.CommandUsage[command]++

	// Update favorite commands (top 5 most used)
	m.updateFavoriteCommands()
}

// updateFavoriteCommands updates the list of favorite commands based on usage
func (m *MemoryCard) updateFavoriteCommands() {
	type cmdUsage struct {
		cmd   string
		count int
	}

	var sorted []cmdUsage
	for cmd, count := range m.CommandUsage {
		sorted = append(sorted, cmdUsage{cmd, count})
	}

	// Simple sort (bubble sort for small lists)
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i].count < sorted[j].count {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	// Get top 5
	m.FavoriteCommands = []string{}
	for i := 0; i < len(sorted) && i < 5; i++ {
		m.FavoriteCommands = append(m.FavoriteCommands, sorted[i].cmd)
	}
}

// AddMemorableMoment records a special moment in the relationship
func (m *MemoryCard) AddMemorableMoment(momentType, title, description, emotion, context string) {
	moment := MemorableMoment{
		Timestamp:   time.Now(),
		Type:        momentType,
		Title:       title,
		Description: description,
		Emotion:     emotion,
		Context:     context,
	}
	m.MemorableMoments = append(m.MemorableMoments, moment)

	// Keep only last 50 memorable moments
	if len(m.MemorableMoments) > 50 {
		m.MemorableMoments = m.MemorableMoments[len(m.MemorableMoments)-50:]
	}

	// Increase relationship level
	if m.RelationshipLevel < 100 {
		m.RelationshipLevel += 2
	}
}

// AddAchievement records a user achievement
func (m *MemoryCard) AddAchievement(id, title, description, category, project string) {
	achievement := Achievement{
		ID:          id,
		Title:       title,
		Description: description,
		EarnedAt:    time.Now(),
		Category:    category,
		Project:     project,
	}
	m.Achievements = append(m.Achievements, achievement)

	// Keep only last 100 achievements
	if len(m.Achievements) > 100 {
		m.Achievements = m.Achievements[len(m.Achievements)-100:]
	}

	// Create memorable moment for significant achievements
	if category == "milestone" || category == "project" {
		m.AddMemorableMoment("achievement", title, description, "proud", fmt.Sprintf("Completed %s in project %s", title, project))
	}
}

// UpdateRelationshipMetrics recalculates relationship metrics
func (m *MemoryCard) UpdateRelationshipMetrics() {
	// Relationship level based on multiple factors
	baseLevel := m.ToneLevel * 5 // Tone contributes up to 50 points

	// Add points for interaction frequency
	daysSinceFirst := time.Since(m.FirstMet).Hours() / 24
	if daysSinceFirst > 0 {
		interactionFrequency := float64(m.SessionCount) / daysSinceFirst
		if interactionFrequency > 0.5 { // More than once every 2 days
			baseLevel += 20
		} else if interactionFrequency > 0.1 { // More than once every 10 days
			baseLevel += 10
		}
	}

	// Add points for memorable moments
	baseLevel += len(m.MemorableMoments) * 2

	// Add points for achievements
	baseLevel += len(m.Achievements)

	// Cap at 100
	if baseLevel > 100 {
		baseLevel = 100
	}
	m.RelationshipLevel = baseLevel

	// Update engagement score (0-1)
	if m.SessionCount > 0 {
		avgTimeBetweenSessions := time.Since(m.FirstMet).Hours() / float64(m.SessionCount)
		if avgTimeBetweenSessions < 24 { // Less than a day between sessions
			m.EngagementScore = 1.0
		} else if avgTimeBetweenSessions < 168 { // Less than a week
			m.EngagementScore = 0.7
		} else if avgTimeBetweenSessions < 720 { // Less than a month
			m.EngagementScore = 0.4
		} else {
			m.EngagementScore = 0.2
		}
	}
}

// RecordSession tracks a new session
func (m *MemoryCard) RecordSession(durationMinutes float64) {
	m.SessionCount++
	if m.AverageSessionTime == 0 {
		m.AverageSessionTime = durationMinutes
	} else {
		// Running average
		m.AverageSessionTime = (m.AverageSessionTime*float64(m.SessionCount-1) + durationMinutes) / float64(m.SessionCount)
	}
	m.LastInteraction = time.Now()
	m.UpdateStreak()
	m.UpdateRelationshipMetrics()
}

// UpdateStreak updates the user's usage streak
func (m *MemoryCard) UpdateStreak() {
	if m.UsageDates == nil {
		m.UsageDates = []string{}
	}

	now := time.Now()
	today := now.Format("2006-01-02")

	// Check if we already recorded today
	alreadyRecorded := false
	for _, date := range m.UsageDates {
		if date == today {
			alreadyRecorded = true
			break
		}
	}

	// If not recorded today, add it
	if !alreadyRecorded {
		m.UsageDates = append(m.UsageDates, today)
		// Keep only last 60 days to avoid unbounded growth
		if len(m.UsageDates) > 60 {
			m.UsageDates = m.UsageDates[len(m.UsageDates)-60:]
		}
	}

	// Calculate current streak
	m.CurrentStreak = m.calculateCurrentStreak()

	// Update longest streak if current is longer
	if m.CurrentStreak > m.LongestStreak {
		m.LongestStreak = m.CurrentStreak
	}

	// Update last usage date
	m.LastUsageDate = now
}

// calculateCurrentStreak calculates the current consecutive day streak
func (m *MemoryCard) calculateCurrentStreak() int {
	if len(m.UsageDates) == 0 {
		return 0
	}

	// Sort dates (they should already be in order, but ensure it)
	dates := make([]string, len(m.UsageDates))
	copy(dates, m.UsageDates)

	// Get today's date
	today := time.Now()

	// Count backwards from today
	streak := 0
	for i := 0; i < 60; i++ { // Check up to 60 days back
		checkDate := today.AddDate(0, 0, -i)
		dateStr := checkDate.Format("2006-01-02")

		// Check if this date exists in usage dates
		found := false
		for _, d := range dates {
			if d == dateStr {
				found = true
				break
			}
		}

		if found {
			streak++
		} else {
			// If we're checking today (i=0) and it's not found, assume today counts
			// (user is currently using DoPlan, so today should be counted)
			// Otherwise, break the streak
			if i == 0 {
				streak++
				continue
			}
			break
		}
	}

	return streak
}

// ResolvePainPoint marks a pain point as resolved
func (m *MemoryCard) ResolvePainPoint(painPoint string) {
	if m.ResolvedPainPoints == nil {
		m.ResolvedPainPoints = []string{}
	}

	// Check if already resolved
	for _, resolved := range m.ResolvedPainPoints {
		if resolved == painPoint {
			return // Already resolved
		}
	}

	// Add to resolved list
	m.ResolvedPainPoints = append(m.ResolvedPainPoints, painPoint)

	// Remove from active pain points if present
	if m.PainPoints != nil {
		newPainPoints := []string{}
		for _, pp := range m.PainPoints {
			if pp != painPoint {
				newPainPoints = append(newPainPoints, pp)
			}
		}
		m.PainPoints = newPainPoints
	}
}

// HasResolvedPainPoint checks if a pain point has been resolved
func (m *MemoryCard) HasResolvedPainPoint(painPoint string) bool {
	if m.ResolvedPainPoints == nil {
		return false
	}

	for _, resolved := range m.ResolvedPainPoints {
		if resolved == painPoint {
			return true
		}
	}

	return false
}

// GetGreeting returns a personalized greeting based on memory card
func (m *MemoryCard) GetGreeting() string {
	now := time.Now()
	hour := now.Hour()
	timeOfDay := "day"
	if hour < 6 {
		timeOfDay = "night"
	} else if hour < 12 {
		timeOfDay = "morning"
	} else if hour < 18 {
		timeOfDay = "afternoon"
	} else {
		timeOfDay = "evening"
	}

	// Check if returning after a long absence
	daysSinceLastInteraction := time.Since(m.LastInteraction).Hours() / 24
	isReturning := daysSinceLastInteraction > 7

	if m.UserName != "" {
		// High relationship level (80+)
		if m.RelationshipLevel >= 80 {
			if isReturning {
				return fmt.Sprintf("Hey %s! 👋 Long time no see! I've missed our coding sessions. Ready to pick up where we left off?", m.UserName)
			}
			if timeOfDay == "morning" {
				return fmt.Sprintf("Good morning, %s! ☀️ Great to see you again. Let's make today productive!", m.UserName)
			} else if timeOfDay == "evening" {
				return fmt.Sprintf("Hey %s! 🌙 Working late? I'm here to help. What are we building tonight?", m.UserName)
			}
			return fmt.Sprintf("Hey %s! 👋 Back for more? Let's build something awesome together!", m.UserName)
		}

		// Medium relationship level (40-79)
		if m.RelationshipLevel >= 40 {
			if isReturning {
				return fmt.Sprintf("Welcome back, %s! 👋 Good to see you again. Ready to continue building?", m.UserName)
			}
			return fmt.Sprintf("Hey %s! 👋 Good to have you back. Let's create something great together!", m.UserName)
		}

		// Low relationship level (0-39)
		if m.ProjectsCount > 1 {
			return fmt.Sprintf("Hello %s! 👋 Welcome back. Ready to work on another project?", m.UserName)
		}
		return fmt.Sprintf("Hello %s! 👋 Let's turn your idea into reality!", m.UserName)
	}

	// First time or no name
	if m.RelationshipLevel == 0 {
		return "Hello! 👋 I'm DoPlan, your AI development partner. Let's get started!"
	}
	return "Hello! 👋 Welcome back to DoPlan. Ready to build something amazing?"
}

// GetTone returns the appropriate tone based on relationship level
func (m *MemoryCard) GetTone() string {
	if m.ToneLevel >= 8 {
		return "friendly" // Very warm, casual
	} else if m.ToneLevel >= 5 {
		return "warm" // Friendly but professional
	} else if m.ToneLevel >= 2 {
		return "professional" // Professional with hints of warmth
	}
	return "formal" // Professional and formal
}

// GetEncouragement returns personalized encouragement based on user profile
func (m *MemoryCard) GetEncouragement() string {
	// High relationship level - more personal
	if m.RelationshipLevel >= 70 {
		if len(m.Achievements) > 0 {
			lastAchievement := m.Achievements[len(m.Achievements)-1]
			return fmt.Sprintf("Remember when you %s? That was awesome! 🎉 Let's create another win like that!", lastAchievement.Title)
		}
		if m.Motivation == "change_world" {
			return "Your vision to change the world inspires me! 🌍 Every project you build gets us closer to that goal."
		}
		return "I believe in you! We've accomplished so much together. Let's keep the momentum going! 💪"
	}

	// Medium relationship level
	if m.RelationshipLevel >= 40 {
		if m.Motivation == "change_world" {
			return "I love your ambition to change the world! 🌍 Let's build something that makes a real impact."
		} else if m.Motivation == "money" {
			return "Let's build something profitable! 💰 I'll help you create value that users will pay for."
		} else if m.Interest == "learning" {
			return "Great choice to focus on learning! 📚 This project will teach you so much."
		}
		return "Let's build something amazing together! 🚀"
	}

	// Low relationship level - more general
	if m.Motivation == "change_world" {
		return "Building something to change the world is ambitious! 🌍 I'm here to help you make it happen."
	} else if m.Motivation == "money" {
		return "Let's create something valuable! 💰 I'll guide you through building a profitable product."
	} else if m.Interest == "learning" {
		return "Focusing on learning is smart! 📚 This project will be a great learning experience."
	}
	return "Let's build something amazing together! 🚀"
}

// GetPersonalizedTip returns a tip based on user's patterns and preferences
func (m *MemoryCard) GetPersonalizedTip() string {
	// Check for common pain points
	if len(m.PainPoints) > 0 {
		lastPainPoint := m.PainPoints[len(m.PainPoints)-1]
		return fmt.Sprintf("💡 Tip: I noticed you've struggled with %s before. Want me to help you approach it differently this time?", lastPainPoint)
	}

	// Check for learning goals
	if len(m.LearningGoals) > 0 {
		goal := m.LearningGoals[len(m.LearningGoals)-1]
		return fmt.Sprintf("💡 Tip: Since you're interested in learning %s, this project is a great opportunity to practice!", goal)
	}

	// Check favorite commands
	if len(m.FavoriteCommands) > 0 {
		favCmd := m.FavoriteCommands[0]
		return fmt.Sprintf("💡 Tip: I see you love using /%s! It's one of my favorites too. Want to explore it more?", favCmd)
	}

	return "💡 Tip: Take it one step at a time. Every great project starts with a single command!"
}

// GetContextualMessage returns a message based on current context
func (m *MemoryCard) GetContextualMessage(context string) string {
	switch context {
	case "first_project":
		return "🎉 Exciting! Your first project with DoPlan. I'm here to guide you every step of the way!"
	case "returning_after_break":
		days := int(time.Since(m.LastInteraction).Hours() / 24)
		return fmt.Sprintf("👋 Welcome back! It's been %d days. I'm excited to see what you're building next!", days)
	case "milestone_reached":
		return "🎊 Congratulations on reaching this milestone! You're making great progress!"
	case "struggling":
		return "💪 I see you're facing some challenges. That's okay - every developer goes through this. Let's work through it together!"
	case "high_engagement":
		return "🔥 You're on fire! I love seeing this level of engagement. Keep it up!"
	default:
		return ""
	}
}
