package cli

import (
	"fmt"
	"io"
	"strings"
)

// CelebrateAchievements displays a celebration for earned achievements
func CelebrateAchievements(achievements []AchievementExtended, out io.Writer) {
	if len(achievements) == 0 {
		return
	}

	// Calculate total points
	totalPoints := 0
	for _, achievement := range achievements {
		totalPoints += achievement.Points
	}

	// Determine celebration level
	celebrationLevel := determineCelebrationLevel(achievements)

	// Display celebration
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 60))
	fmt.Fprintln(out, "")

	if len(achievements) == 1 {
		displaySingleAchievement(achievements[0], totalPoints, out)
	} else {
		displayMultipleAchievements(achievements, totalPoints, celebrationLevel, out)
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 60))
	fmt.Fprintln(out, "")
}

// displaySingleAchievement displays a single achievement celebration
func displaySingleAchievement(achievement AchievementExtended, points int, out io.Writer) {
	fmt.Fprintf(out, "  %s  %s  %s\n", achievement.Icon, achievement.Title, achievement.Icon)
	fmt.Fprintf(out, "  %s\n", achievement.Description)
	fmt.Fprintf(out, "  +%d points\n", points)
	fmt.Fprintf(out, "  Rarity: %s\n", strings.Title(achievement.Rarity))
}

// displayMultipleAchievements displays multiple achievements (dopamine release!)
func displayMultipleAchievements(achievements []AchievementExtended, totalPoints int, level string, out io.Writer) {
	// Exciting header based on number of achievements
	if len(achievements) >= 3 {
		fmt.Fprintf(out, "  🎉🎉🎉  AMAZING! You earned %d achievements!  🎉🎉🎉\n", len(achievements))
	} else {
		fmt.Fprintf(out, "  🎊  Congratulations! You earned %d achievements!  🎊\n", len(achievements))
	}

	fmt.Fprintln(out, "")

	// Display each achievement
	for i, achievement := range achievements {
		fmt.Fprintf(out, "  %d. %s %s\n", i+1, achievement.Icon, achievement.Title)
		fmt.Fprintf(out, "     %s (+%d points)\n", achievement.Description, achievement.Points)
		if achievement.Rarity == "legendary" || achievement.Rarity == "epic" {
			fmt.Fprintf(out, "     ⭐ %s\n", strings.Title(achievement.Rarity))
		}
		fmt.Fprintln(out, "")
	}

	// Total points
	fmt.Fprintf(out, "  💰 Total Points: +%d\n", totalPoints)
	fmt.Fprintf(out, "  📊 New Score: %d\n", totalPoints) // Will be updated with actual score
}

// determineCelebrationLevel determines how excited to be
func determineCelebrationLevel(achievements []AchievementExtended) string {
	hasLegendary := false
	hasEpic := false
	hasRare := false

	for _, achievement := range achievements {
		switch achievement.Rarity {
		case "legendary":
			hasLegendary = true
		case "epic":
			hasEpic = true
		case "rare":
			hasRare = true
		}
	}

	if hasLegendary {
		return "legendary"
	}
	if hasEpic {
		return "epic"
	}
	if hasRare {
		return "rare"
	}
	if len(achievements) >= 3 {
		return "multiple"
	}
	return "normal"
}

// GetAchievementSummary returns a summary of user's achievements
func GetAchievementSummary(mc *MemoryCard) string {
	if mc == nil {
		return ""
	}

	var summary strings.Builder

	summary.WriteString("🏆 Achievement Summary\n")
	summary.WriteString(fmt.Sprintf("   Score: %d points\n", mc.Score))
	summary.WriteString(fmt.Sprintf("   Total Achievements: %d\n", len(mc.Achievements)))

	// Group by category
	categories := make(map[string]int)
	for _, achievement := range mc.Achievements {
		categories[achievement.Category]++
	}

	if len(categories) > 0 {
		summary.WriteString("\n   By Category:\n")
		for category, count := range categories {
			summary.WriteString(fmt.Sprintf("   - %s: %d\n", strings.Title(category), count))
		}
	}

	// Recent achievements
	if len(mc.Achievements) > 0 {
		summary.WriteString("\n   Recent Achievements:\n")
		recent := mc.Achievements
		if len(recent) > 5 {
			recent = recent[len(recent)-5:]
		}
		for _, achievement := range recent {
			summary.WriteString(fmt.Sprintf("   - %s\n", achievement.Title))
		}
	}

	return summary.String()
}

// CheckAchievementsOnCommand checks for achievements after a command execution
func CheckAchievementsOnCommand(command string, context map[string]interface{}) ([]AchievementExtended, error) {
	// Initialize achievement system
	as, err := NewAchievementSystem()
	if err != nil {
		return []AchievementExtended{}, nil // Graceful degradation
	}

	// Add command context
	context["command"] = command
	context["score"] = as.GetScore()
	context["achievements"] = len(as.memoryCard.Achievements)

	// Check for achievements
	earned, err := as.CheckAndAwardAchievements(context)
	if err != nil {
		return []AchievementExtended{}, err
	}

	return earned, nil
}

// GetNextAchievements returns hints about next achievable achievements
func GetNextAchievements(mc *MemoryCard, limit int) []string {
	if mc == nil || limit <= 0 {
		return []string{}
	}

	var hints []string

	// Score milestones
	nextScoreMilestone := getNextScoreMilestone(mc.Score)
	if nextScoreMilestone > 0 {
		pointsNeeded := nextScoreMilestone - mc.Score
		hints = append(hints, fmt.Sprintf("🎯 %d more points to reach %d (Score Milestone)", pointsNeeded, nextScoreMilestone))
	}

	// Project milestones
	if mc.ProjectsCount < 5 {
		hints = append(hints, fmt.Sprintf("🏗️  Complete %d more projects to reach 5 projects", 5-mc.ProjectsCount))
	} else if mc.ProjectsCount < 10 {
		hints = append(hints, fmt.Sprintf("🏗️  Complete %d more projects to reach 10 projects", 10-mc.ProjectsCount))
	}

	// Achievement count milestones
	achievementCount := len(mc.Achievements)
	if achievementCount < 10 {
		hints = append(hints, fmt.Sprintf("🎖️  Earn %d more achievements to reach 10", 10-achievementCount))
	} else if achievementCount < 25 {
		hints = append(hints, fmt.Sprintf("🎖️  Earn %d more achievements to reach 25", 25-achievementCount))
	}

	// Relationship milestones
	if mc.RelationshipLevel < 40 {
		hints = append(hints, fmt.Sprintf("🤝 Reach relationship level 40 (currently %d)", mc.RelationshipLevel))
	} else if mc.RelationshipLevel < 70 {
		hints = append(hints, fmt.Sprintf("💎 Reach relationship level 70 (currently %d)", mc.RelationshipLevel))
	}

	// Limit results
	if len(hints) > limit {
		hints = hints[:limit]
	}

	return hints
}

// getNextScoreMilestone returns the next score milestone
func getNextScoreMilestone(currentScore int) int {
	milestones := []int{100, 250, 500, 1000, 2500, 5000, 10000, 25000, 50000, 100000}
	for _, milestone := range milestones {
		if currentScore < milestone {
			return milestone
		}
	}
	return 0
}
