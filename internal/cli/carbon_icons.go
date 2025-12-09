package cli

// CarbonIcon represents a Carbon Design System icon name
// These are the icon names from @carbon/icons-react
// Reference: https://carbondesignsystem.com/elements/icons/library/

// GetCarbonIcon returns the Carbon icon name for an achievement
// First parameter can be achievement ID or category
func GetCarbonIcon(achievementIDOrCategory, achievementType, rarity string) string {
	// Map by achievement ID first (most specific)
	iconMap := map[string]string{
		// Score milestones
		"score_100":   "Target",
		"score_250":   "TrendingUp",
		"score_500":   "Star",
		"score_1000":  "Trophy",
		"score_2500":  "Diamond",
		"score_5000":  "Crown",
		"score_10000": "Sparkles",
		"score_25000": "Rocket",
		"score_50000": "Lightning",

		// Project achievements
		"first_project":        "Celebration",
		"second_mvp":           "Confetti",
		"five_projects":        "Building",
		"ten_projects":         "Medal",
		"twenty_five_projects": "Crown",
		"fifty_projects":       "Lightning",
		"project_types_5":      "Palette",
		"project_types_10":     "Sparkles",

		// Command usage
		"command_hey_10":  "Wave",
		"command_do_50":   "Rocket",
		"command_plan_25": "Document",
		"command_dev_100": "Code",
		"all_commands":    "Target",

		// Learning
		"learning_goal_1": "Book",
		"learning_goal_5": "Education",
		"tech_stack_3":    "Tools",
		"tech_stack_10":   "Settings",

		// Productivity
		"session_10":  "Time",
		"session_50":  "Fire",
		"session_100": "Number_100",
		"session_500": "Lightning",

		// Streak
		"streak_3":  "Fire",
		"streak_7":  "Strength",
		"streak_30": "Crown",

		// Relationship
		"relationship_40":      "Handshake",
		"relationship_70":      "Diamond",
		"relationship_100":     "Heart",
		"trust_10":             "Handshake",
		"memorable_moments_10": "Camera",
		"memorable_moments_50": "Book",

		// Milestones
		"achievement_10":  "Medal",
		"achievement_25":  "Award",
		"achievement_50":  "Trophy",
		"achievement_100": "Crown",

		// Special
		"first_met_anniversary": "Cake",
		"night_owl":             "Moon",
		"early_bird":            "Sun",
		"overcome_pain_point":   "Strength",
	}

	// Try to get specific icon by achievement ID
	if icon, ok := iconMap[achievementIDOrCategory]; ok {
		return icon
	}

	// Fallback to category-based icons (if achievementIDOrCategory is actually a category)
	categoryIcons := map[string]string{
		"score":        "Target",
		"project":      "Building",
		"command":      "Command",
		"learning":     "Book",
		"productivity": "Time",
		"streak":       "Fire",
		"relationship": "User",
		"milestone":    "Medal",
		"special":      "Sparkles",
	}

	if icon, ok := categoryIcons[achievementIDOrCategory]; ok {
		return icon
	}

	// Default icon based on rarity
	rarityIcons := map[string]string{
		"common":    "Circle",
		"uncommon":  "Star",
		"rare":      "Diamond",
		"epic":      "Trophy",
		"legendary": "Crown",
	}

	if icon, ok := rarityIcons[rarity]; ok {
		return icon
	}

	// Ultimate fallback
	return "Checkmark"
}
