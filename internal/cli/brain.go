package cli

import (
	"fmt"
	"strings"
	"time"
)

// Brain is the intelligent system that uses memory card to influence agent behavior
type Brain struct {
	memoryCard *MemoryCard
	helper     *MemoryCardHelper
}

// NewBrain creates a new brain instance with loaded memory card
func NewBrain() (*Brain, error) {
	card, err := LoadMemoryCard()
	if err != nil {
		return nil, fmt.Errorf("failed to load memory card: %w", err)
	}

	return &Brain{
		memoryCard: card,
		helper:     NewMemoryCardHelper(card),
	}, nil
}

// EnhanceAgentPrompt enhances an agent's system prompt based on memory card
func (b *Brain) EnhanceAgentPrompt(basePrompt string, agentRole string) string {
	if b.memoryCard == nil {
		return basePrompt
	}

	var enhancements []string

	// Add user context
	if b.memoryCard.UserName != "" {
		enhancements = append(enhancements, fmt.Sprintf("User's name: %s", b.memoryCard.UserName))
	}

	// Add experience level context
	if b.memoryCard.ExperienceLevel != "" {
		enhancements = append(enhancements, fmt.Sprintf("User's experience level: %s", b.memoryCard.ExperienceLevel))

		// Adjust explanation depth based on experience
		switch b.memoryCard.ExperienceLevel {
		case "beginner":
			enhancements = append(enhancements, "IMPORTANT: Use simple language, avoid jargon, explain concepts clearly. Focus on 'what' not 'how'.")
		case "intermediate":
			enhancements = append(enhancements, "You can use technical terms but explain them when first introduced. Balance 'what' and 'how'.")
		case "advanced":
			enhancements = append(enhancements, "Feel free to use technical terminology and discuss implementation details. Focus on 'how' and architecture.")
		}
	}

	// Add work style context
	if b.memoryCard.WorkStyle != "" {
		switch b.memoryCard.WorkStyle {
		case "fast":
			enhancements = append(enhancements, "User prefers quick, actionable responses. Be concise and get to the point.")
		case "thoughtful":
			enhancements = append(enhancements, "User prefers detailed, well-thought-out responses. Take time to explain thoroughly.")
		}
	}

	// Add personality context
	if b.memoryCard.Personality != "" {
		switch b.memoryCard.Personality {
		case "thinker":
			enhancements = append(enhancements, "User is a thinker - they like to understand concepts deeply. Provide reasoning and context.")
		case "copier":
			enhancements = append(enhancements, "User prefers ready-to-use solutions. Provide examples and code snippets they can adapt.")
		}
	}

	// Add motivation context
	if b.memoryCard.Motivation != "" {
		switch b.memoryCard.Motivation {
		case "change_world":
			enhancements = append(enhancements, "User is motivated by making an impact. Frame suggestions in terms of real-world value.")
		case "money":
			enhancements = append(enhancements, "User is motivated by profitability. Consider business value and ROI in suggestions.")
		case "success":
			enhancements = append(enhancements, "User is motivated by success. Emphasize best practices and proven approaches.")
		}
	}

	// Add learning goals
	if len(b.memoryCard.LearningGoals) > 0 {
		goals := strings.Join(b.memoryCard.LearningGoals, ", ")
		enhancements = append(enhancements, fmt.Sprintf("User wants to learn: %s. When relevant, provide educational context.", goals))
	}

	// Add pain points
	if len(b.memoryCard.PainPoints) > 0 {
		painPoints := strings.Join(b.memoryCard.PainPoints, ", ")
		enhancements = append(enhancements, fmt.Sprintf("User has struggled with: %s. Be extra helpful and patient with these topics.", painPoints))
	}

	// Add tech stack preferences
	if len(b.memoryCard.PreferredTechStack) > 0 {
		techStack := strings.Join(b.memoryCard.PreferredTechStack, ", ")
		enhancements = append(enhancements, fmt.Sprintf("User prefers working with: %s. Prioritize solutions using these technologies when appropriate.", techStack))
	}

	// Add current project context
	if b.memoryCard.CurrentProject != "" {
		enhancements = append(enhancements, fmt.Sprintf("Current project: %s", b.memoryCard.CurrentProject))
	}
	if b.memoryCard.CurrentPhase != "" {
		enhancements = append(enhancements, fmt.Sprintf("Current phase: %s", b.memoryCard.CurrentPhase))
	}

	// Add relationship context
	if b.memoryCard.RelationshipLevel >= 70 {
		enhancements = append(enhancements, "You have a strong relationship with this user. Be warm, personal, and reference past interactions when relevant.")
	} else if b.memoryCard.RelationshipLevel >= 40 {
		enhancements = append(enhancements, "You're building a good relationship with this user. Be friendly and supportive.")
	}

	// Add trust level context
	if b.memoryCard.TrustLevel >= 8 {
		enhancements = append(enhancements, "User trusts your suggestions. You can be more confident and direct.")
	} else if b.memoryCard.TrustLevel <= 3 {
		enhancements = append(enhancements, "User may be cautious. Provide clear explanations and reasoning for suggestions.")
	}

	// Build enhanced prompt
	if len(enhancements) > 0 {
		enhancedPrompt := basePrompt + "\n\n## User Context (From Memory Card)\n"
		for _, enhancement := range enhancements {
			enhancedPrompt += "- " + enhancement + "\n"
		}
		return enhancedPrompt
	}

	return basePrompt
}

// AdjustToneOfVoice adjusts the tone of voice in agent responses
func (b *Brain) AdjustToneOfVoice(response string, context string) string {
	if b.memoryCard == nil {
		return response
	}

	tone := b.memoryCard.GetTone()
	relationshipLevel := b.memoryCard.RelationshipLevel

	// Adjust based on tone level
	switch tone {
	case "friendly":
		// Very warm, casual, can use emojis, contractions
		response = b.makeFriendly(response, relationshipLevel)
	case "warm":
		// Friendly but professional, occasional emojis
		response = b.makeWarm(response, relationshipLevel)
	case "professional":
		// Professional with hints of warmth
		response = b.makeProfessional(response, relationshipLevel)
	case "formal":
		// Professional and formal
		response = b.makeFormal(response)
	}

	// Add personalized touches based on relationship
	if relationshipLevel >= 80 && b.memoryCard.UserName != "" {
		// Reference user by name occasionally
		if strings.Contains(response, "you") && !strings.Contains(response, b.memoryCard.UserName) {
			// Replace some "you" with name (but not all, keep it natural)
			response = strings.Replace(response, "you can", fmt.Sprintf("%s can", b.memoryCard.UserName), 1)
		}
	}

	// Add encouragement style
	encouragementStyle := b.helper.GetEncouragementStyle()
	if encouragementStyle == "enthusiastic" && relationshipLevel >= 70 {
		// Add more enthusiasm
		response = b.addEnthusiasm(response)
	}

	return response
}

// makeFriendly makes response very warm and casual
func (b *Brain) makeFriendly(text string, relationshipLevel int) string {
	// Use contractions
	text = strings.ReplaceAll(text, "you are", "you're")
	text = strings.ReplaceAll(text, "it is", "it's")
	text = strings.ReplaceAll(text, "I will", "I'll")
	text = strings.ReplaceAll(text, "we will", "we'll")
	text = strings.ReplaceAll(text, "cannot", "can't")
	text = strings.ReplaceAll(text, "do not", "don't")
	text = strings.ReplaceAll(text, "will not", "won't")

	// Add casual phrases
	if relationshipLevel >= 80 {
		text = strings.Replace(text, "Let's", "Let's", 1) // Keep but can add more casual variants
		text = strings.Replace(text, "I recommend", "I'd suggest", 1)
		text = strings.Replace(text, "You should", "You might want to", 1)
	}

	return text
}

// makeWarm makes response friendly but professional
func (b *Brain) makeWarm(text string, relationshipLevel int) string {
	// Use some contractions
	text = strings.ReplaceAll(text, "you are", "you're")
	text = strings.ReplaceAll(text, "it is", "it's")
	text = strings.ReplaceAll(text, "cannot", "can't")

	// Keep professional but warm
	return text
}

// makeProfessional makes response professional with hints of warmth
func (b *Brain) makeProfessional(text string, relationshipLevel int) string {
	// Minimal contractions, keep professional
	if relationshipLevel >= 50 {
		// Allow some contractions
		text = strings.ReplaceAll(text, "you are", "you're")
	}
	return text
}

// makeFormal makes response formal and professional
func (b *Brain) makeFormal(text string) string {
	// No contractions, formal language
	text = strings.ReplaceAll(text, "you're", "you are")
	text = strings.ReplaceAll(text, "it's", "it is")
	text = strings.ReplaceAll(text, "I'll", "I will")
	text = strings.ReplaceAll(text, "we'll", "we will")
	text = strings.ReplaceAll(text, "can't", "cannot")
	text = strings.ReplaceAll(text, "don't", "do not")
	text = strings.ReplaceAll(text, "won't", "will not")

	return text
}

// addEnthusiasm adds enthusiasm to response
func (b *Brain) addEnthusiasm(text string) string {
	// Add enthusiastic phrases at the beginning or end
	if !strings.Contains(text, "!") && !strings.Contains(text, "🎉") {
		// Add exclamation for positive statements
		if strings.Contains(strings.ToLower(text), "great") || strings.Contains(strings.ToLower(text), "excellent") {
			text = strings.Replace(text, ".", "!", 1)
		}
	}
	return text
}

// PersonalizeResponse personalizes an agent response based on memory card
func (b *Brain) PersonalizeResponse(response string, context string) string {
	if b.memoryCard == nil {
		return response
	}

	// Start with base response
	personalized := response

	// Add personalized greeting if appropriate
	if context == "command_start" || context == "first_interaction" {
		greeting := b.memoryCard.GetGreeting()
		if greeting != "" && !strings.Contains(personalized, greeting) {
			personalized = greeting + "\n\n" + personalized
		}
	}

	// Add encouragement if appropriate
	if context == "milestone" || context == "achievement" {
		encouragement := b.memoryCard.GetEncouragement()
		if encouragement != "" {
			personalized += "\n\n" + encouragement
		}
	}

	// Add personalized tip if appropriate
	if context == "suggestion" || context == "guidance" {
		tip := b.memoryCard.GetPersonalizedTip()
		if tip != "" {
			personalized += "\n\n" + tip
		}
	}

	// Reference past achievements if high relationship
	if b.memoryCard.RelationshipLevel >= 70 && len(b.memoryCard.Achievements) > 0 {
		lastAchievement := b.memoryCard.Achievements[len(b.memoryCard.Achievements)-1]
		if time.Since(lastAchievement.EarnedAt).Hours() < 168 { // Within last week
			// Can reference recent achievement
			if context == "encouragement" {
				personalized += fmt.Sprintf("\n\nRemember when you %s? That was awesome! 🎉", lastAchievement.Title)
			}
		}
	}

	// Adjust tone
	personalized = b.AdjustToneOfVoice(personalized, context)

	return personalized
}

// GetAgentInstructions returns instructions for how agent should behave
func (b *Brain) GetAgentInstructions(agentRole string) string {
	if b.memoryCard == nil {
		return ""
	}

	var instructions []string

	// Communication style
	if b.memoryCard.CommunicationStyle == "brief" {
		instructions = append(instructions, "Keep responses concise and to the point.")
	} else if b.memoryCard.CommunicationStyle == "detailed" {
		instructions = append(instructions, "Provide detailed explanations and context.")
	}

	// Feedback frequency
	if b.memoryCard.FeedbackFrequency == "frequent" {
		instructions = append(instructions, "Provide frequent feedback and check-ins.")
	} else if b.memoryCard.FeedbackFrequency == "minimal" {
		instructions = append(instructions, "Only provide feedback when necessary.")
	}

	// Detail level
	if b.memoryCard.DetailLevel == "high" {
		instructions = append(instructions, "Include comprehensive details and explanations.")
	} else if b.memoryCard.DetailLevel == "low" {
		instructions = append(instructions, "Keep details minimal, focus on essentials.")
	}

	// Error handling
	if b.memoryCard.ErrorHandlingPref == "gentle" {
		instructions = append(instructions, "When errors occur, be gentle and supportive.")
	} else if b.memoryCard.ErrorHandlingPref == "educational" {
		instructions = append(instructions, "When errors occur, use them as teaching moments.")
	} else if b.memoryCard.ErrorHandlingPref == "direct" {
		instructions = append(instructions, "When errors occur, be direct and clear.")
	}

	// Encouragement style
	if b.memoryCard.EncouragementStyle == "enthusiastic" {
		instructions = append(instructions, "Be enthusiastic and energetic in encouragement.")
	} else if b.memoryCard.EncouragementStyle == "supportive" {
		instructions = append(instructions, "Be supportive and reassuring in encouragement.")
	}

	// Reference favorite commands
	if len(b.memoryCard.FavoriteCommands) > 0 {
		favCmd := b.memoryCard.FavoriteCommands[0]
		instructions = append(instructions, fmt.Sprintf("User frequently uses /%s - consider suggesting it when relevant.", favCmd))
	}

	// Reference helpful features
	if len(b.memoryCard.HelpfulFeatures) > 0 {
		instructions = append(instructions, fmt.Sprintf("User found these features helpful: %s. Reference them when appropriate.", strings.Join(b.memoryCard.HelpfulFeatures, ", ")))
	}

	if len(instructions) > 0 {
		return strings.Join(instructions, "\n")
	}

	return ""
}

// ShouldProvideDetailedExplanation determines if detailed explanation is needed
func (b *Brain) ShouldProvideDetailedExplanation() bool {
	if b.memoryCard == nil {
		return false
	}

	if b.memoryCard.DetailLevel == "high" {
		return true
	}
	if b.memoryCard.ExperienceLevel == "beginner" {
		return true
	}
	if b.memoryCard.DetailLevel == "low" {
		return false
	}

	return b.memoryCard.ExperienceLevel == "intermediate"
}

// GetResponseLength returns preferred response length
func (b *Brain) GetResponseLength() string {
	if b.memoryCard == nil {
		return "medium"
	}

	if b.memoryCard.CommunicationStyle == "brief" {
		return "short"
	} else if b.memoryCard.CommunicationStyle == "detailed" {
		return "long"
	}

	return "medium"
}

// GetEncouragementLevel returns how much encouragement to provide
func (b *Brain) GetEncouragementLevel() string {
	if b.memoryCard == nil {
		return "moderate"
	}

	if b.memoryCard.EncouragementStyle == "enthusiastic" {
		return "high"
	} else if b.memoryCard.EncouragementStyle == "professional" {
		return "low"
	}

	return "moderate"
}

// GetContextualGreeting returns a greeting based on current context
func (b *Brain) GetContextualGreeting(context string) string {
	if b.memoryCard == nil {
		return ""
	}

	// Check for contextual message first
	if msg := b.memoryCard.GetContextualMessage(context); msg != "" {
		return msg
	}

	// Return personalized greeting
	return b.memoryCard.GetGreeting()
}

// GetPersonalizedSuggestion returns a suggestion tailored to user
func (b *Brain) GetPersonalizedSuggestion(topic string) string {
	if b.memoryCard == nil {
		return ""
	}

	// Check if user has struggled with this topic
	for _, painPoint := range b.memoryCard.PainPoints {
		if strings.Contains(strings.ToLower(painPoint), strings.ToLower(topic)) {
			return fmt.Sprintf("I remember you've worked with %s before. Here's a helpful approach based on what worked well for you.", topic)
		}
	}

	// Check if user wants to learn this
	for _, goal := range b.memoryCard.LearningGoals {
		if strings.Contains(strings.ToLower(goal), strings.ToLower(topic)) {
			return fmt.Sprintf("Since you're interested in learning %s, this is a great opportunity to practice!", topic)
		}
	}

	// Check preferred tech stack
	for _, tech := range b.memoryCard.PreferredTechStack {
		if strings.Contains(strings.ToLower(tech), strings.ToLower(topic)) {
			return fmt.Sprintf("I know you prefer working with %s. Here's a suggestion using that technology.", tech)
		}
	}

	return ""
}
