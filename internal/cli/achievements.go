package cli

import (
	"fmt"
	"time"
)

// AchievementExtended extends Achievement with scoring fields
type AchievementExtended struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"` // "project", "learning", "productivity", "milestone", "score", "streak", "social"
	Type        string    `json:"type"`     // "achievement", "trophy", "badge"
	Points      int       `json:"points"`   // Points awarded for earning this
	Rarity      string    `json:"rarity"`   // "common", "uncommon", "rare", "epic", "legendary"
	Icon        string    `json:"icon"`     // Emoji or icon identifier
	EarnedAt    time.Time `json:"earned_at,omitempty"`
	Project     string    `json:"project,omitempty"` // Associated project if applicable
}

// ToAchievement converts AchievementExtended to Achievement for memory card
func (ae AchievementExtended) ToAchievement() Achievement {
	return Achievement{
		ID:          ae.ID,
		Title:       ae.Title,
		Description: ae.Description,
		Category:    ae.Category,
		EarnedAt:    ae.EarnedAt,
		Project:     ae.Project,
	}
}

// AchievementDefinition defines how to detect and award an achievement
type AchievementDefinition struct {
	ID          string
	Title       string
	Description string
	Category    string
	Type        string
	Points      int
	Rarity      string
	Icon        string
	Condition   AchievementCondition // Function to check if achievement is earned
}

// AchievementCondition is a function that checks if achievement should be awarded
type AchievementCondition func(mc *MemoryCard, context map[string]interface{}) bool

// AchievementSystem manages achievements and scoring
type AchievementSystem struct {
	memoryCard   *MemoryCard
	score        int
	achievements []Achievement
	definitions  []AchievementDefinition
}

// NewAchievementSystem creates a new achievement system
func NewAchievementSystem() (*AchievementSystem, error) {
	mc, err := LoadMemoryCard()
	if err != nil {
		return nil, err
	}

	system := &AchievementSystem{
		memoryCard:   mc,
		score:        mc.Score,
		achievements: mc.Achievements,
		definitions:  GetAllAchievementDefinitions(),
	}

	return system, nil
}

// CheckAndAwardAchievements checks for new achievements and awards them
func (as *AchievementSystem) CheckAndAwardAchievements(context map[string]interface{}) ([]AchievementExtended, error) {
	var newlyEarned []AchievementExtended

	// Check each achievement definition
	for _, def := range as.definitions {
		// Skip if already earned
		if as.isAlreadyEarned(def.ID) {
			continue
		}

		// Check condition
		if def.Condition(as.memoryCard, context) {
			achievement := AchievementExtended{
				ID:          def.ID,
				Title:       def.Title,
				Description: def.Description,
				Category:    def.Category,
				Type:        def.Type,
				Points:      def.Points,
				Rarity:      def.Rarity,
				Icon:        def.Icon,
				EarnedAt:    time.Now(),
			}

			// Add project if in context
			if project, ok := context["project"].(string); ok {
				achievement.Project = project
			}

			newlyEarned = append(newlyEarned, achievement)
		}
	}

	// Award achievements (this increases score, which may trigger more achievements)
	if len(newlyEarned) > 0 {
		return as.awardAchievements(newlyEarned, context)
	}

	return []AchievementExtended{}, nil
}

// awardAchievements awards achievements and checks for cascading achievements
func (as *AchievementSystem) awardAchievements(achievements []AchievementExtended, context map[string]interface{}) ([]AchievementExtended, error) {
	var allEarned []AchievementExtended

	// Award each achievement
	for _, achievement := range achievements {
		// Add to memory card (convert to base Achievement)
		as.memoryCard.Achievements = append(as.memoryCard.Achievements, achievement.ToAchievement())

		// Increase score
		as.score += achievement.Points
		as.memoryCard.Score = as.score

		// Add to earned list
		allEarned = append(allEarned, achievement)

		// Create memorable moment
		as.memoryCard.AddMemorableMoment(
			"achievement",
			achievement.Title,
			achievement.Description,
			"proud",
			fmt.Sprintf("Earned %s achievement", achievement.Rarity),
		)
	}

	// Save memory card
	if err := SaveMemoryCard(as.memoryCard); err != nil {
		return allEarned, fmt.Errorf("failed to save memory card: %w", err)
	}

	// Check for cascading achievements (score milestones, etc.)
	cascadingContext := map[string]interface{}{
		"score":        as.score,
		"achievements": len(as.memoryCard.Achievements),
	}
	for k, v := range context {
		cascadingContext[k] = v
	}

	cascading, err := as.CheckAndAwardAchievements(cascadingContext)
	if err != nil {
		return allEarned, err
	}

	allEarned = append(allEarned, cascading...)

	return allEarned, nil
}

// isAlreadyEarned checks if achievement is already earned
func (as *AchievementSystem) isAlreadyEarned(achievementID string) bool {
	for _, achievement := range as.memoryCard.Achievements {
		if achievement.ID == achievementID {
			return true
		}
	}
	return false
}

// GetScore returns current score
func (as *AchievementSystem) GetScore() int {
	return as.score
}

// GetAchievementsByCategory returns achievements grouped by category
func (as *AchievementSystem) GetAchievementsByCategory() map[string][]Achievement {
	categories := make(map[string][]Achievement)
	for _, achievement := range as.memoryCard.Achievements {
		categories[achievement.Category] = append(categories[achievement.Category], achievement)
	}
	return categories
}

// GetRecentAchievements returns recently earned achievements
func (as *AchievementSystem) GetRecentAchievements(limit int) []Achievement {
	if limit <= 0 {
		limit = 10
	}

	achievements := as.memoryCard.Achievements
	if len(achievements) == 0 {
		return []Achievement{}
	}

	// Sort by earned date (most recent first)
	recent := make([]Achievement, 0, limit)
	for i := len(achievements) - 1; i >= 0 && len(recent) < limit; i-- {
		recent = append(recent, achievements[i])
	}

	return recent
}

// GetProgress returns progress towards next achievement
func (as *AchievementSystem) GetProgress() map[string]interface{} {
	progress := make(map[string]interface{})

	progress["score"] = as.score
	progress["total_achievements"] = len(as.memoryCard.Achievements)
	progress["next_score_milestone"] = as.getNextScoreMilestone()
	progress["categories"] = as.GetAchievementsByCategory()

	return progress
}

// getNextScoreMilestone returns the next score milestone
func (as *AchievementSystem) getNextScoreMilestone() int {
	milestones := []int{100, 250, 500, 1000, 2500, 5000, 10000, 25000, 50000, 100000}
	for _, milestone := range milestones {
		if as.score < milestone {
			return milestone
		}
	}
	return 0 // All milestones reached
}
