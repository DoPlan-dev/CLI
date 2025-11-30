package cli

import (
	"time"
)

// GetAllAchievementDefinitions returns all achievement definitions
func GetAllAchievementDefinitions() []AchievementDefinition {
	definitions := []AchievementDefinition{}

	// Score Milestones (10 achievements)
	definitions = append(definitions, getScoreMilestoneAchievements()...)

	// Project Achievements (30+ achievements)
	definitions = append(definitions, getProjectAchievements()...)

	// Command Usage Achievements (20+ achievements)
	definitions = append(definitions, getCommandUsageAchievements()...)

	// Learning Achievements (20+ achievements)
	definitions = append(definitions, getLearningAchievements()...)

	// Productivity Achievements (20+ achievements)
	definitions = append(definitions, getProductivityAchievements()...)

	// Streak Achievements (15+ achievements)
	definitions = append(definitions, getStreakAchievements()...)

	// Relationship Achievements (15+ achievements)
	definitions = append(definitions, getRelationshipAchievements()...)

	// Milestone Achievements (20+ achievements)
	definitions = append(definitions, getMilestoneAchievements()...)

	// Special Achievements (30+ achievements)
	definitions = append(definitions, getSpecialAchievements()...)

	return definitions
}

// Score Milestone Achievements
func getScoreMilestoneAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "score_100",
			Title:       "Getting Started",
			Description: "Reach 100 points",
			Category:    "score",
			Type:        "achievement",
			Points:      10,
			Rarity:      "common",
			Icon:        "🎯",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 100 && score < 250
				}
				return mc.Score >= 100 && mc.Score < 250
			},
		},
		{
			ID:          "score_250",
			Title:       "On the Rise",
			Description: "Reach 250 points",
			Category:    "score",
			Type:        "achievement",
			Points:      25,
			Rarity:      "common",
			Icon:        "📈",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 250 && score < 500
				}
				return mc.Score >= 250 && mc.Score < 500
			},
		},
		{
			ID:          "score_500",
			Title:       "Halfway Hero",
			Description: "Reach 500 points",
			Category:    "score",
			Type:        "achievement",
			Points:      50,
			Rarity:      "uncommon",
			Icon:        "⭐",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 500 && score < 1000
				}
				return mc.Score >= 500 && mc.Score < 1000
			},
		},
		{
			ID:          "score_1000",
			Title:       "Thousand Club",
			Description: "Reach 1,000 points",
			Category:    "score",
			Type:        "trophy",
			Points:      100,
			Rarity:      "rare",
			Icon:        "🏆",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 1000 && score < 2500
				}
				return mc.Score >= 1000 && mc.Score < 2500
			},
		},
		{
			ID:          "score_2500",
			Title:       "Elite Developer",
			Description: "Reach 2,500 points",
			Category:    "score",
			Type:        "trophy",
			Points:      250,
			Rarity:      "rare",
			Icon:        "💎",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 2500 && score < 5000
				}
				return mc.Score >= 2500 && mc.Score < 5000
			},
		},
		{
			ID:          "score_5000",
			Title:       "Master Builder",
			Description: "Reach 5,000 points",
			Category:    "score",
			Type:        "trophy",
			Points:      500,
			Rarity:      "epic",
			Icon:        "👑",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 5000 && score < 10000
				}
				return mc.Score >= 5000 && mc.Score < 10000
			},
		},
		{
			ID:          "score_10000",
			Title:       "Legendary Coder",
			Description: "Reach 10,000 points",
			Category:    "score",
			Type:        "trophy",
			Points:      1000,
			Rarity:      "epic",
			Icon:        "🌟",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 10000 && score < 25000
				}
				return mc.Score >= 10000 && mc.Score < 25000
			},
		},
		{
			ID:          "score_25000",
			Title:       "Unstoppable Force",
			Description: "Reach 25,000 points",
			Category:    "score",
			Type:        "trophy",
			Points:      2500,
			Rarity:      "legendary",
			Icon:        "🚀",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 25000 && score < 50000
				}
				return mc.Score >= 25000 && mc.Score < 50000
			},
		},
		{
			ID:          "score_50000",
			Title:       "God Mode",
			Description: "Reach 50,000 points",
			Category:    "score",
			Type:        "trophy",
			Points:      5000,
			Rarity:      "legendary",
			Icon:        "⚡",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["score"].(int); ok {
					return score >= 50000
				}
				return mc.Score >= 50000
			},
		},
	}
}

// Project Achievements
func getProjectAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "first_project",
			Title:       "First Steps",
			Description: "Complete your first project",
			Category:    "project",
			Type:        "achievement",
			Points:      50,
			Rarity:      "common",
			Icon:        "🎉",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.ProjectsCount == 1
			},
		},
		{
			ID:          "second_mvp",
			Title:       "Double Trouble",
			Description: "Complete MVP for the second time",
			Category:    "project",
			Type:        "achievement",
			Points:      75,
			Rarity:      "uncommon",
			Icon:        "🎊",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if phase, ok := context["phase"].(string); ok && phase == "completed" {
					return mc.ProjectsCount == 2
				}
				return false
			},
		},
		{
			ID:          "five_projects",
			Title:       "Serial Builder",
			Description: "Complete 5 projects",
			Category:    "project",
			Type:        "trophy",
			Points:      200,
			Rarity:      "rare",
			Icon:        "🏗️",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.ProjectsCount == 5
			},
		},
		{
			ID:          "ten_projects",
			Title:       "Decade Developer",
			Description: "Complete 10 projects",
			Category:    "project",
			Type:        "trophy",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🎖️",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.ProjectsCount == 10
			},
		},
		{
			ID:          "twenty_five_projects",
			Title:       "Project Master",
			Description: "Complete 25 projects",
			Category:    "project",
			Type:        "trophy",
			Points:      1500,
			Rarity:      "legendary",
			Icon:        "👑",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.ProjectsCount == 25
			},
		},
		{
			ID:          "fifty_projects",
			Title:       "Unstoppable Creator",
			Description: "Complete 50 projects",
			Category:    "project",
			Type:        "trophy",
			Points:      5000,
			Rarity:      "legendary",
			Icon:        "⚡",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.ProjectsCount == 50
			},
		},
		{
			ID:          "project_types_5",
			Title:       "Versatile Builder",
			Description: "Work on 5 different project types",
			Category:    "project",
			Type:        "achievement",
			Points:      100,
			Rarity:      "uncommon",
			Icon:        "🎨",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.ProjectTypes) >= 5
			},
		},
		{
			ID:          "project_types_10",
			Title:       "Jack of All Trades",
			Description: "Work on 10 different project types",
			Category:    "project",
			Type:        "trophy",
			Points:      300,
			Rarity:      "rare",
			Icon:        "🌟",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.ProjectTypes) >= 10
			},
		},
		// Add more project achievements...
	}
}

// Command Usage Achievements
func getCommandUsageAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "command_hey_10",
			Title:       "Hello There!",
			Description: "Use /hey command 10 times",
			Category:    "productivity",
			Type:        "achievement",
			Points:      25,
			Rarity:      "common",
			Icon:        "👋",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.CommandUsage["/hey"] == 10
			},
		},
		{
			ID:          "command_do_50",
			Title:       "Do It Again",
			Description: "Use /do command 50 times",
			Category:    "productivity",
			Type:        "achievement",
			Points:      75,
			Rarity:      "uncommon",
			Icon:        "🚀",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.CommandUsage["/do"] == 50
			},
		},
		{
			ID:          "command_plan_25",
			Title:       "Planner",
			Description: "Use /plan command 25 times",
			Category:    "productivity",
			Type:        "achievement",
			Points:      50,
			Rarity:      "common",
			Icon:        "📋",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.CommandUsage["/plan"] == 25
			},
		},
		{
			ID:          "command_dev_100",
			Title:       "Code Machine",
			Description: "Use /dev command 100 times",
			Category:    "productivity",
			Type:        "trophy",
			Points:      200,
			Rarity:      "rare",
			Icon:        "💻",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.CommandUsage["/dev"] == 100
			},
		},
		{
			ID:          "all_commands",
			Title:       "Command Master",
			Description: "Use all core commands at least once",
			Category:    "productivity",
			Type:        "achievement",
			Points:      100,
			Rarity:      "uncommon",
			Icon:        "🎯",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				coreCommands := []string{"/hey", "/do", "/plan", "/dev"}
				for _, cmd := range coreCommands {
					if mc.CommandUsage[cmd] == 0 {
						return false
					}
				}
				return true
			},
		},
		// Add more command achievements...
	}
}

// Learning Achievements
func getLearningAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "learning_goal_1",
			Title:       "Student",
			Description: "Set your first learning goal",
			Category:    "learning",
			Type:        "achievement",
			Points:      25,
			Rarity:      "common",
			Icon:        "📚",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.LearningGoals) == 1
			},
		},
		{
			ID:          "learning_goal_5",
			Title:       "Knowledge Seeker",
			Description: "Set 5 learning goals",
			Category:    "learning",
			Type:        "achievement",
			Points:      75,
			Rarity:      "uncommon",
			Icon:        "🎓",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.LearningGoals) == 5
			},
		},
		{
			ID:          "tech_stack_3",
			Title:       "Tech Explorer",
			Description: "Work with 3 different technologies",
			Category:    "learning",
			Type:        "achievement",
			Points:      50,
			Rarity:      "common",
			Icon:        "🔧",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.PreferredTechStack) >= 3
			},
		},
		{
			ID:          "tech_stack_10",
			Title:       "Tech Master",
			Description: "Work with 10 different technologies",
			Category:    "learning",
			Type:        "trophy",
			Points:      300,
			Rarity:      "rare",
			Icon:        "⚙️",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.PreferredTechStack) >= 10
			},
		},
		// Add more learning achievements...
	}
}

// Productivity Achievements
func getProductivityAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "session_10",
			Title:       "Getting Into It",
			Description: "Complete 10 sessions",
			Category:    "productivity",
			Type:        "achievement",
			Points:      50,
			Rarity:      "common",
			Icon:        "⏱️",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.SessionCount == 10
			},
		},
		{
			ID:          "session_50",
			Title:       "Dedicated",
			Description: "Complete 50 sessions",
			Category:    "productivity",
			Type:        "achievement",
			Points:      150,
			Rarity:      "uncommon",
			Icon:        "🔥",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.SessionCount == 50
			},
		},
		{
			ID:          "session_100",
			Title:       "Centurion",
			Description: "Complete 100 sessions",
			Category:    "productivity",
			Type:        "trophy",
			Points:      400,
			Rarity:      "rare",
			Icon:        "💯",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.SessionCount == 100
			},
		},
		{
			ID:          "session_500",
			Title:       "Power User",
			Description: "Complete 500 sessions",
			Category:    "productivity",
			Type:        "trophy",
			Points:      2000,
			Rarity:      "epic",
			Icon:        "⚡",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.SessionCount == 500
			},
		},
		// Add more productivity achievements...
	}
}

// Streak Achievements
func getStreakAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "streak_3",
			Title:       "Three Day Streak",
			Description: "Use DoPlan 3 days in a row",
			Category:    "streak",
			Type:        "achievement",
			Points:      30,
			Rarity:      "common",
			Icon:        "🔥",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				// Check if used 3 days in a row (simplified check)
				return false // Would need streak tracking
			},
		},
		{
			ID:          "streak_7",
			Title:       "Week Warrior",
			Description: "Use DoPlan 7 days in a row",
			Category:    "streak",
			Type:        "achievement",
			Points:      100,
			Rarity:      "uncommon",
			Icon:        "💪",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return false // Would need streak tracking
			},
		},
		{
			ID:          "streak_30",
			Title:       "Monthly Master",
			Description: "Use DoPlan 30 days in a row",
			Category:    "streak",
			Type:        "trophy",
			Points:      500,
			Rarity:      "epic",
			Icon:        "👑",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return false // Would need streak tracking
			},
		},
		// Add more streak achievements...
	}
}

// Relationship Achievements
func getRelationshipAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "relationship_40",
			Title:       "Building Connection",
			Description: "Reach relationship level 40",
			Category:    "relationship",
			Type:        "achievement",
			Points:      50,
			Rarity:      "common",
			Icon:        "🤝",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.RelationshipLevel >= 40 && mc.RelationshipLevel < 70
			},
		},
		{
			ID:          "relationship_70",
			Title:       "True Partner",
			Description: "Reach relationship level 70",
			Category:    "relationship",
			Type:        "trophy",
			Points:      200,
			Rarity:      "rare",
			Icon:        "💎",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.RelationshipLevel >= 70 && mc.RelationshipLevel < 100
			},
		},
		{
			ID:          "relationship_100",
			Title:       "Best Friends Forever",
			Description: "Reach maximum relationship level",
			Category:    "relationship",
			Type:        "trophy",
			Points:      1000,
			Rarity:      "legendary",
			Icon:        "❤️",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.RelationshipLevel >= 100
			},
		},
		{
			ID:          "trust_10",
			Title:       "Complete Trust",
			Description: "Reach maximum trust level",
			Category:    "relationship",
			Type:        "trophy",
			Points:      300,
			Rarity:      "epic",
			Icon:        "🤝",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return mc.TrustLevel >= 10
			},
		},
		{
			ID:          "memorable_moments_10",
			Title:       "Memory Keeper",
			Description: "Create 10 memorable moments",
			Category:    "relationship",
			Type:        "achievement",
			Points:      75,
			Rarity:      "uncommon",
			Icon:        "📸",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.MemorableMoments) >= 10
			},
		},
		{
			ID:          "memorable_moments_50",
			Title:       "Storyteller",
			Description: "Create 50 memorable moments",
			Category:    "relationship",
			Type:        "trophy",
			Points:      500,
			Rarity:      "epic",
			Icon:        "📖",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				return len(mc.MemorableMoments) >= 50
			},
		},
		// Add more relationship achievements...
	}
}

// Milestone Achievements
func getMilestoneAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "achievement_10",
			Title:       "Achievement Hunter",
			Description: "Earn 10 achievements",
			Category:    "milestone",
			Type:        "achievement",
			Points:      50,
			Rarity:      "common",
			Icon:        "🎖️",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if count, ok := context["achievements"].(int); ok {
					return count == 10
				}
				return len(mc.Achievements) == 10
			},
		},
		{
			ID:          "achievement_25",
			Title:       "Collector",
			Description: "Earn 25 achievements",
			Category:    "milestone",
			Type:        "achievement",
			Points:      150,
			Rarity:      "uncommon",
			Icon:        "🏅",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if count, ok := context["achievements"].(int); ok {
					return count == 25
				}
				return len(mc.Achievements) == 25
			},
		},
		{
			ID:          "achievement_50",
			Title:       "Trophy Hunter",
			Description: "Earn 50 achievements",
			Category:    "milestone",
			Type:        "trophy",
			Points:      400,
			Rarity:      "rare",
			Icon:        "🏆",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if count, ok := context["achievements"].(int); ok {
					return count == 50
				}
				return len(mc.Achievements) == 50
			},
		},
		{
			ID:          "achievement_100",
			Title:       "Legendary Collector",
			Description: "Earn 100 achievements",
			Category:    "milestone",
			Type:        "trophy",
			Points:      1000,
			Rarity:      "legendary",
			Icon:        "👑",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if count, ok := context["achievements"].(int); ok {
					return count == 100
				}
				return len(mc.Achievements) == 100
			},
		},
		// Add more milestone achievements...
	}
}

// Special Achievements
func getSpecialAchievements() []AchievementDefinition {
	return []AchievementDefinition{
		{
			ID:          "first_met_anniversary",
			Title:       "Anniversary",
			Description: "Celebrate your first year with DoPlan",
			Category:    "special",
			Type:        "trophy",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🎂",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				daysSinceFirst := time.Since(mc.FirstMet).Hours() / 24
				return daysSinceFirst >= 365 && daysSinceFirst < 730
			},
		},
		{
			ID:          "night_owl",
			Title:       "Night Owl",
			Description: "Use DoPlan after midnight",
			Category:    "special",
			Type:        "achievement",
			Points:      25,
			Rarity:      "common",
			Icon:        "🦉",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				hour := time.Now().Hour()
				return hour >= 0 && hour < 6
			},
		},
		{
			ID:          "early_bird",
			Title:       "Early Bird",
			Description: "Use DoPlan before 6 AM",
			Category:    "special",
			Type:        "achievement",
			Points:      25,
			Rarity:      "common",
			Icon:        "🐦",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				hour := time.Now().Hour()
				return hour >= 4 && hour < 6
			},
		},
		{
			ID:          "overcome_pain_point",
			Title:       "Problem Solver",
			Description: "Overcome a pain point you've struggled with",
			Category:    "special",
			Type:        "achievement",
			Points:      100,
			Rarity:      "uncommon",
			Icon:        "💪",
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				// Would need to track when pain points are resolved
				return false
			},
		},
		// Add many more special achievements...
	}
}
