package cli

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Challenge represents a first-time task or milestone that earns high scores
type Challenge struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"` // "integration", "database", "deployment", "testing", "workflow"
	Points      int       `json:"points"`   // High points for challenges
	Rarity      string    `json:"rarity"`   // Usually "rare", "epic", or "legendary"
	Icon        string    `json:"icon"`     // Emoji icon
	CompletedAt time.Time `json:"completed_at,omitempty"`
	Project     string    `json:"project,omitempty"`
	Attempts    int       `json:"attempts,omitempty"` // Number of attempts before completion
}

// ChallengeDefinition defines a challenge that can be detected
type ChallengeDefinition struct {
	ID          string
	Title       string
	Description string
	Category    string
	Points      int
	Rarity      string
	Icon        string
	Condition   ChallengeCondition // Function to check if challenge is completed
	IsFirstTime bool               // True if this is a first-time challenge
}

// ChallengeCondition checks if challenge should be awarded
type ChallengeCondition func(mc *MemoryCard, context map[string]interface{}) bool

// ChallengeSystem manages challenges
type ChallengeSystem struct {
	memoryCard  *MemoryCard
	definitions []ChallengeDefinition
}

// NewChallengeSystem creates a new challenge system
func NewChallengeSystem() (*ChallengeSystem, error) {
	mc, err := LoadMemoryCard()
	if err != nil {
		return nil, err
	}

	system := &ChallengeSystem{
		memoryCard:  mc,
		definitions: GetAllChallengeDefinitions(),
	}

	return system, nil
}

// CheckAndAwardChallenges checks for completed challenges
func (cs *ChallengeSystem) CheckAndAwardChallenges(context map[string]interface{}) ([]Challenge, error) {
	var newlyCompleted []Challenge

	// Check each challenge definition
	for _, def := range cs.definitions {
		// Skip if already completed
		if cs.isAlreadyCompleted(def.ID) {
			continue
		}

		// Check condition
		if def.Condition(cs.memoryCard, context) {
			// Increment attempt count
			attempts := cs.memoryCard.ChallengeAttempts[def.ID] + 1
			cs.memoryCard.ChallengeAttempts[def.ID] = attempts

			challenge := Challenge{
				ID:          def.ID,
				Title:       def.Title,
				Description: def.Description,
				Category:    def.Category,
				Points:      def.Points,
				Rarity:      def.Rarity,
				Icon:        def.Icon,
				CompletedAt: time.Now(),
				Attempts:    attempts,
			}

			// Add project if in context
			if project, ok := context["project"].(string); ok {
				challenge.Project = project
			}

			newlyCompleted = append(newlyCompleted, challenge)
		}
	}

	// Award challenges
	if len(newlyCompleted) > 0 {
		return cs.awardChallenges(newlyCompleted, context)
	}

	return []Challenge{}, nil
}

// awardChallenges awards challenges and updates memory card
func (cs *ChallengeSystem) awardChallenges(challenges []Challenge, context map[string]interface{}) ([]Challenge, error) {
	// Add to completed challenges
	for _, challenge := range challenges {
		cs.memoryCard.CompletedChallenges = append(cs.memoryCard.CompletedChallenges, challenge.ID)

		// Increase score
		cs.memoryCard.Score += challenge.Points

		// Create memorable moment
		cs.memoryCard.AddMemorableMoment(
			"challenge_completed",
			challenge.Title,
			challenge.Description,
			"proud",
			fmt.Sprintf("Completed %s challenge", challenge.Category),
		)

		// Add achievement to achievements list
		achievement := Achievement{
			ID:          challenge.ID,
			Title:       challenge.Title,
			Description: challenge.Description,
			Category:    challenge.Category,
			EarnedAt:    challenge.CompletedAt,
			Project:     challenge.Project,
		}
		cs.memoryCard.Achievements = append(cs.memoryCard.Achievements, achievement)
	}

	// Save memory card
	if err := SaveMemoryCard(cs.memoryCard); err != nil {
		return challenges, fmt.Errorf("failed to save memory card: %w", err)
	}

	return challenges, nil
}

// isAlreadyCompleted checks if challenge is already completed
func (cs *ChallengeSystem) isAlreadyCompleted(challengeID string) bool {
	for _, completed := range cs.memoryCard.CompletedChallenges {
		if completed == challengeID {
			return true
		}
	}
	return false
}

// GetPendingChallenges returns challenges that haven't been completed yet
func (cs *ChallengeSystem) GetPendingChallenges() []ChallengeDefinition {
	var pending []ChallengeDefinition

	for _, def := range cs.definitions {
		if !cs.isAlreadyCompleted(def.ID) {
			pending = append(pending, def)
		}
	}

	return pending
}

// GetChallengeProgress returns progress on a specific challenge
func (cs *ChallengeSystem) GetChallengeProgress(challengeID string) map[string]interface{} {
	progress := make(map[string]interface{})

	progress["completed"] = cs.isAlreadyCompleted(challengeID)
	progress["attempts"] = cs.memoryCard.ChallengeAttempts[challengeID]

	return progress
}

// CelebrateChallenge displays celebration for completed challenge
func CelebrateChallenge(challenge Challenge, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  🎯  CHALLENGE COMPLETED!  🎯\n")
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  %s  %s  %s\n", challenge.Icon, challenge.Title, challenge.Icon)
	fmt.Fprintf(out, "  %s\n", challenge.Description)
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  💰 Points Earned: +%d\n", challenge.Points)
	fmt.Fprintf(out, "  ⭐ Rarity: %s\n", strings.Title(challenge.Rarity))
	if challenge.Attempts > 1 {
		fmt.Fprintf(out, "  💪 Completed in %d attempts - Great persistence!\n", challenge.Attempts)
	} else {
		fmt.Fprintf(out, "  ⚡ Completed on first try - Amazing!\n")
	}
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  🎉 This is a significant milestone! You're making great progress!\n")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")
}

// CelebrateMultipleChallenges displays celebration for multiple challenges
func CelebrateMultipleChallenges(challenges []Challenge, out io.Writer) {
	if len(challenges) == 0 {
		return
	}

	totalPoints := 0
	for _, challenge := range challenges {
		totalPoints += challenge.Points
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  🚀🚀🚀  INCREDIBLE! You completed %d challenges!  🚀🚀🚀\n", len(challenges))
	fmt.Fprintln(out, "")

	for i, challenge := range challenges {
		fmt.Fprintf(out, "  %d. %s %s\n", i+1, challenge.Icon, challenge.Title)
		fmt.Fprintf(out, "     %s\n", challenge.Description)
		fmt.Fprintf(out, "     +%d points", challenge.Points)
		if challenge.Rarity == "legendary" || challenge.Rarity == "epic" {
			fmt.Fprintf(out, " ⭐ %s", strings.Title(challenge.Rarity))
		}
		if challenge.Attempts == 1 {
			fmt.Fprintf(out, " ⚡ First try!")
		}
		fmt.Fprintln(out, "")
	}

	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  💰 Total Points: +%d\n", totalPoints)
	fmt.Fprintf(out, "  📊 New Score: %d\n", totalPoints) // Will be actual score
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "  🎊 You're on fire! Keep up the amazing work!\n")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")
}
