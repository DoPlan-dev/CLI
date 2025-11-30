package cli

import (
	"fmt"
	"time"
)

// MemoryCardHelper provides convenient methods for working with memory cards
type MemoryCardHelper struct {
	card *MemoryCard
}

// NewMemoryCardHelper creates a new helper instance
func NewMemoryCardHelper(card *MemoryCard) *MemoryCardHelper {
	return &MemoryCardHelper{card: card}
}

// TrackCommandExecution records a command execution with full context
func (h *MemoryCardHelper) TrackCommandExecution(command string, userInput string, agentResponse string, durationSeconds float64, sentiment string) {
	if h.card == nil {
		return
	}

	insights := map[string]string{
		"command":        command,
		"user_input":     userInput,
		"agent_response": agentResponse,
		"sentiment":      sentiment,
	}

	h.card.UpdateFromConversation(insights)
	h.card.RecordCommandUsage(command)
	h.card.LastCommand = command
	h.card.LastCommandTime = time.Now()

	// Update session tracking
	if durationSeconds > 0 {
		h.card.RecordSession(durationSeconds / 60.0) // Convert to minutes
	}
}

// RecordSuccess records a successful interaction
func (h *MemoryCardHelper) RecordSuccess(context string, details string) {
	if h.card == nil {
		return
	}

	h.card.AddMemorableMoment("achievement", "Success", details, "happy", context)

	// Increase trust level
	if h.card.TrustLevel < 10 {
		h.card.TrustLevel++
	}
}

// RecordStruggle records when user faces challenges
func (h *MemoryCardHelper) RecordStruggle(feature string, issue string) {
	if h.card == nil {
		return
	}

	// Add to struggled features if not already there
	found := false
	for _, f := range h.card.StruggledFeatures {
		if f == feature {
			found = true
			break
		}
	}
	if !found {
		h.card.StruggledFeatures = append(h.card.StruggledFeatures, feature)
	}

	// Add pain point
	h.card.PainPoints = append(h.card.PainPoints, issue)
	if len(h.card.PainPoints) > 20 {
		h.card.PainPoints = h.card.PainPoints[len(h.card.PainPoints)-20:]
	}

	// Add memorable moment for overcoming challenge
	h.card.AddMemorableMoment("challenge_overcome", fmt.Sprintf("Overcame %s", feature), issue, "relieved", "User faced challenge but persisted")
}

// RecordHelpfulFeature records when user finds a feature helpful
func (h *MemoryCardHelper) RecordHelpfulFeature(feature string) {
	if h.card == nil {
		return
	}

	found := false
	for _, f := range h.card.HelpfulFeatures {
		if f == feature {
			found = true
			break
		}
	}
	if !found {
		h.card.HelpfulFeatures = append(h.card.HelpfulFeatures, feature)
		if len(h.card.HelpfulFeatures) > 20 {
			h.card.HelpfulFeatures = h.card.HelpfulFeatures[len(h.card.HelpfulFeatures)-20:]
		}
	}
}

// RecordLearningGoal records what user wants to learn
func (h *MemoryCardHelper) RecordLearningGoal(goal string) {
	if h.card == nil {
		return
	}

	found := false
	for _, g := range h.card.LearningGoals {
		if g == goal {
			found = true
			break
		}
	}
	if !found {
		h.card.LearningGoals = append(h.card.LearningGoals, goal)
		if len(h.card.LearningGoals) > 10 {
			h.card.LearningGoals = h.card.LearningGoals[len(h.card.LearningGoals)-10:]
		}
	}
}

// RecordProjectStart records when user starts a new project
func (h *MemoryCardHelper) RecordProjectStart(projectName string, projectType string) {
	if h.card == nil {
		return
	}

	h.card.ProjectsCount++
	h.card.LastProjectDate = time.Now()
	h.card.CurrentProject = projectName
	h.card.CurrentPhase = "ideation"

	// Add project type if new
	found := false
	for _, pt := range h.card.ProjectTypes {
		if pt == projectType {
			found = true
			break
		}
	}
	if !found {
		h.card.ProjectTypes = append(h.card.ProjectTypes, projectType)
	}

	// Create memorable moment for first project
	if h.card.ProjectsCount == 1 {
		h.card.AddMemorableMoment("first_time", "First Project", fmt.Sprintf("Started first project: %s", projectName), "excited", "Beginning of journey")
	}

	// Add achievement
	h.card.AddAchievement(
		fmt.Sprintf("project_%d", h.card.ProjectsCount),
		fmt.Sprintf("Started Project #%d", h.card.ProjectsCount),
		fmt.Sprintf("Began working on %s", projectName),
		"project",
		projectName,
	)
}

// RecordPhaseTransition records when user moves to a new phase
func (h *MemoryCardHelper) RecordPhaseTransition(fromPhase string, toPhase string) {
	if h.card == nil {
		return
	}

	h.card.CurrentPhase = toPhase

	// Create memorable moment for significant transitions
	significantTransitions := map[string]string{
		"ideation":    "planning",
		"planning":    "development",
		"development": "testing",
		"testing":     "deployment",
	}

	if significantTransitions[fromPhase] == toPhase {
		h.card.AddMemorableMoment(
			"breakthrough",
			fmt.Sprintf("Moved to %s", toPhase),
			fmt.Sprintf("Progressed from %s to %s", fromPhase, toPhase),
			"proud",
			fmt.Sprintf("Project: %s", h.card.CurrentProject),
		)
	}
}

// RecordMilestone records a project milestone
func (h *MemoryCardHelper) RecordMilestone(milestone string, description string) {
	if h.card == nil {
		return
	}

	h.card.AddAchievement(
		fmt.Sprintf("milestone_%d", len(h.card.Achievements)+1),
		milestone,
		description,
		"milestone",
		h.card.CurrentProject,
	)

	h.card.AddMemorableMoment("achievement", milestone, description, "proud", fmt.Sprintf("Project: %s", h.card.CurrentProject))
}

// GetPersonalizedResponse returns a response tailored to the user
func (h *MemoryCardHelper) GetPersonalizedResponse(context string) string {
	if h.card == nil {
		return ""
	}

	// Check for contextual message first
	if msg := h.card.GetContextualMessage(context); msg != "" {
		return msg
	}

	// Return personalized tip
	return h.card.GetPersonalizedTip()
}

// ShouldUseDetailedExplanation returns whether to provide detailed explanations
func (h *MemoryCardHelper) ShouldUseDetailedExplanation() bool {
	if h.card == nil {
		return false
	}

	if h.card.DetailLevel == "high" {
		return true
	}
	if h.card.ExperienceLevel == "beginner" {
		return true
	}
	if h.card.DetailLevel == "low" {
		return false
	}
	// Default: medium detail
	return h.card.ExperienceLevel == "intermediate"
}

// ShouldProvideFrequentFeedback returns whether to provide frequent feedback
func (h *MemoryCardHelper) ShouldProvideFrequentFeedback() bool {
	if h.card == nil {
		return false
	}

	return h.card.FeedbackFrequency == "frequent" || h.card.WorkStyle == "fast"
}

// GetEncouragementStyle returns the appropriate encouragement style
func (h *MemoryCardHelper) GetEncouragementStyle() string {
	if h.card == nil {
		return "professional"
	}

	if h.card.EncouragementStyle != "" {
		return h.card.EncouragementStyle
	}

	// Default based on relationship level
	if h.card.RelationshipLevel >= 70 {
		return "enthusiastic"
	} else if h.card.RelationshipLevel >= 40 {
		return "supportive"
	}
	return "professional"
}

// GetErrorHandlingStyle returns how to handle errors
func (h *MemoryCardHelper) GetErrorHandlingStyle() string {
	if h.card == nil {
		return "direct"
	}

	if h.card.ErrorHandlingPref != "" {
		return h.card.ErrorHandlingPref
	}

	// Default based on experience
	if h.card.ExperienceLevel == "beginner" {
		return "educational"
	} else if h.card.ExperienceLevel == "advanced" {
		return "direct"
	}
	return "gentle"
}
