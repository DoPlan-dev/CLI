package cli

import (
	"fmt"
	"strings"
)

// AgentBrain wraps agent interactions with memory card intelligence
type AgentBrain struct {
	brain *Brain
}

// NewAgentBrain creates a new agent brain instance
func NewAgentBrain() (*AgentBrain, error) {
	brain, err := NewBrain()
	if err != nil {
		return nil, err
	}

	return &AgentBrain{
		brain: brain,
	}, nil
}

// ProcessAgentPrompt processes an agent's system prompt with memory card enhancements
func (ab *AgentBrain) ProcessAgentPrompt(basePrompt string, agentRole string) string {
	if ab.brain == nil {
		return basePrompt
	}

	// Enhance prompt with user context
	enhancedPrompt := ab.brain.EnhanceAgentPrompt(basePrompt, agentRole)

	// Add agent-specific instructions
	instructions := ab.brain.GetAgentInstructions(agentRole)
	if instructions != "" {
		enhancedPrompt += "\n\n## Behavioral Instructions\n" + instructions
	}

	return enhancedPrompt
}

// ProcessAgentResponse processes an agent's response with personalization and tone adjustment
func (ab *AgentBrain) ProcessAgentResponse(response string, context string) string {
	if ab.brain == nil {
		return response
	}

	// Personalize response
	personalized := ab.brain.PersonalizeResponse(response, context)

	// Adjust tone of voice
	personalized = ab.brain.AdjustToneOfVoice(personalized, context)

	return personalized
}

// GetAgentSystemContext returns system context for agent to include in responses
func (ab *AgentBrain) GetAgentSystemContext() string {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return ""
	}

	var context []string

	// User identity
	if ab.brain.memoryCard.UserName != "" {
		context = append(context, fmt.Sprintf("User: %s", ab.brain.memoryCard.UserName))
	}

	// Current state
	if ab.brain.memoryCard.CurrentProject != "" {
		context = append(context, fmt.Sprintf("Current Project: %s", ab.brain.memoryCard.CurrentProject))
	}
	if ab.brain.memoryCard.CurrentPhase != "" {
		context = append(context, fmt.Sprintf("Current Phase: %s", ab.brain.memoryCard.CurrentPhase))
	}

	// Relationship
	if ab.brain.memoryCard.RelationshipLevel >= 70 {
		context = append(context, "Relationship: Strong - You have a close working relationship with this user.")
	} else if ab.brain.memoryCard.RelationshipLevel >= 40 {
		context = append(context, "Relationship: Good - You're building a solid relationship with this user.")
	}

	// Recent achievements
	if len(ab.brain.memoryCard.Achievements) > 0 {
		lastAchievement := ab.brain.memoryCard.Achievements[len(ab.brain.memoryCard.Achievements)-1]
		context = append(context, fmt.Sprintf("Recent Achievement: %s", lastAchievement.Title))
	}

	if len(context) > 0 {
		return strings.Join(context, "\n")
	}

	return ""
}

// ShouldUseTechnicalLanguage determines if technical language is appropriate
func (ab *AgentBrain) ShouldUseTechnicalLanguage() bool {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return false
	}

	return ab.brain.memoryCard.ExperienceLevel == "advanced" ||
		ab.brain.memoryCard.ExperienceLevel == "intermediate"
}

// ShouldProvideExamples determines if examples should be provided
func (ab *AgentBrain) ShouldProvideExamples() bool {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return true // Default to yes
	}

	// Always provide examples for beginners
	if ab.brain.memoryCard.ExperienceLevel == "beginner" {
		return true
	}

	// Provide examples if user is a "copier" personality
	if ab.brain.memoryCard.Personality == "copier" {
		return true
	}

	// Provide examples if detail level is high
	if ab.brain.memoryCard.DetailLevel == "high" {
		return true
	}

	return true // Default to yes
}

// GetPreferredTechStack returns user's preferred tech stack for suggestions
func (ab *AgentBrain) GetPreferredTechStack() []string {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return []string{}
	}

	return ab.brain.memoryCard.PreferredTechStack
}

// GetLearningGoals returns user's learning goals
func (ab *AgentBrain) GetLearningGoals() []string {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return []string{}
	}

	return ab.brain.memoryCard.LearningGoals
}

// GetPainPoints returns user's pain points
func (ab *AgentBrain) GetPainPoints() []string {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return []string{}
	}

	return ab.brain.memoryCard.PainPoints
}

// FormatResponseForUser formats a response according to user preferences
func (ab *AgentBrain) FormatResponseForUser(response string, responseType string) string {
	if ab.brain == nil {
		return response
	}

	// Get response length preference
	length := ab.brain.GetResponseLength()

	// Adjust length if needed
	if length == "short" && len(response) > 500 {
		// Truncate to first paragraph or 500 chars
		if idx := strings.Index(response, "\n\n"); idx > 0 && idx < 500 {
			response = response[:idx]
		} else if len(response) > 500 {
			response = response[:500] + "..."
		}
	} else if length == "long" && len(response) < 200 {
		// Expand with more detail if user prefers detailed
		// This would typically be done by the agent, but we can add context
		response += "\n\nWould you like more details on any specific aspect?"
	}

	// Add encouragement if appropriate
	encouragementLevel := ab.brain.GetEncouragementLevel()
	if encouragementLevel == "high" && responseType == "completion" {
		encouragement := ab.brain.memoryCard.GetEncouragement()
		if encouragement != "" {
			response += "\n\n" + encouragement
		}
	}

	return response
}

// GetPersonalizedCommandSuggestion suggests a command based on user patterns
func (ab *AgentBrain) GetPersonalizedCommandSuggestion(context string) string {
	if ab.brain == nil || ab.brain.memoryCard == nil {
		return ""
	}

	// Suggest favorite commands when appropriate
	if len(ab.brain.memoryCard.FavoriteCommands) > 0 {
		favCmd := ab.brain.memoryCard.FavoriteCommands[0]

		switch context {
		case "after_plan":
			if favCmd == "/dev" {
				return fmt.Sprintf("Since you often use /%s, you might want to start development now!", favCmd)
			}
		case "after_idea":
			if favCmd == "/plan" {
				return fmt.Sprintf("You usually like to /%s next. Ready to create your execution plan?", favCmd)
			}
		}
	}

	return ""
}

// TrackInteraction tracks an agent interaction for learning
func (ab *AgentBrain) TrackInteraction(command string, userInput string, agentResponse string, durationSeconds float64, sentiment string) {
	if ab.brain == nil || ab.brain.helper == nil {
		return
	}

	ab.brain.helper.TrackCommandExecution(command, userInput, agentResponse, durationSeconds, sentiment)

	// Save memory card after tracking
	if ab.brain.memoryCard != nil {
		SaveMemoryCard(ab.brain.memoryCard)
	}
}
