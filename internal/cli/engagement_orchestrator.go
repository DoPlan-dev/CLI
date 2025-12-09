package cli

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/generator"
)

// EngagementOrchestrator coordinates all engagement systems
type EngagementOrchestrator struct {
	brain          *Brain
	achievementSys *AchievementSystem
	challengeSys   *ChallengeSystem
	dopamineTiming *DopamineTiming
	memoryCard     *MemoryCard
}

// NewEngagementOrchestrator creates a new orchestrator that coordinates all systems
func NewEngagementOrchestrator() (*EngagementOrchestrator, error) {
	// Load memory card
	mc, err := LoadMemoryCard()
	if err != nil {
		return nil, fmt.Errorf("failed to load memory card: %w", err)
	}

	// Initialize brain
	brain, err := NewBrain()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize brain: %w", err)
	}

	// Initialize achievement system
	achievementSys, err := NewAchievementSystem()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize achievement system: %w", err)
	}

	// Initialize challenge system
	challengeSys, err := NewChallengeSystem()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize challenge system: %w", err)
	}

	// Initialize dopamine timing
	dopamineTiming, err := NewDopamineTiming()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dopamine timing: %w", err)
	}

	return &EngagementOrchestrator{
		brain:          brain,
		achievementSys: achievementSys,
		challengeSys:   challengeSys,
		dopamineTiming: dopamineTiming,
		memoryCard:     mc,
	}, nil
}

// ProcessCommandWithEngagement processes a command with full engagement system
func (eo *EngagementOrchestrator) ProcessCommandWithEngagement(
	command string,
	context map[string]interface{},
	out io.Writer,
) error {
	// 0. Update streak tracking (must happen before achievement checks)
	if eo.memoryCard != nil {
		eo.memoryCard.UpdateStreak()
	}

	// 1. Check for pending rewards first (dopamine timing)
	if eo.dopamineTiming != nil {
		rewards, err := eo.dopamineTiming.CheckAndReleaseRewards(out)
		if err != nil {
			// Log but don't fail
			fmt.Fprintf(out, "⚠️  Warning: Could not check pending rewards: %v\n", err)
		}
		if len(rewards) > 0 {
			// Rewards were released, update systems
			eo.refreshSystems()
		}
	}

	// 2. Show anticipation message if appropriate
	if eo.dopamineTiming != nil && eo.dopamineTiming.ShouldShowAnticipation() {
		msg := eo.dopamineTiming.GetAnticipationMessage()
		if msg != "" {
			fmt.Fprintln(out, msg)
			fmt.Fprintln(out, "")
		}
	}

	// 3. Check for new achievements
	if eo.achievementSys != nil {
		// Add command context
		context["command"] = command
		context["score"] = eo.memoryCard.Score
		context["achievements"] = len(eo.memoryCard.Achievements)

		earned, err := eo.achievementSys.CheckAndAwardAchievements(context)
		if err != nil {
			fmt.Fprintf(out, "⚠️  Warning: Achievement check failed: %v\n", err)
		} else if len(earned) > 0 {
			// Convert to notification format for IDE display
			notifications := make([]generator.AchievementNotification, len(earned))
			for i, ach := range earned {
				notifications[i] = generator.AchievementNotification{
					ID:          ach.ID,
					Title:       ach.Title,
					Description: ach.Description,
					Points:      ach.Points,
					Rarity:      ach.Rarity,
					Icon:        ach.Icon,
				}
			}

			// Display achievements with IDE-formatted notifications
			ideNotification := generator.FormatMultipleAchievements(notifications)
			if ideNotification != "" {
				fmt.Fprint(out, ideNotification)
			}

			// Also display traditional celebration
			CelebrateAchievements(earned, out)

			// Also schedule for dopamine timing (for delayed rewards if needed)
			if eo.dopamineTiming != nil {
				for _, achievement := range earned {
					eo.dopamineTiming.ScheduleReward(
						"achievement",
						achievement.ID,
						achievement.Title,
						achievement.Description,
						achievement.Points,
						achievement.Rarity,
						achievement.Icon,
						achievement.Project,
					)
				}
			}

			// Refresh systems after achievements awarded
			eo.refreshSystems()
		}

		// 4. Check for pain point resolution
		if eo.memoryCard != nil && len(eo.memoryCard.PainPoints) > 0 {
			eo.detectPainPointResolution(context)
			// Save memory card after pain point resolution (may have modified ResolvedPainPoints)
			if eo.memoryCard != nil {
				if err := SaveMemoryCard(eo.memoryCard); err != nil {
					fmt.Fprintf(out, "⚠️  Warning: Failed to save memory card after pain point resolution: %v\n", err)
				}
			}
		}
	}

	// 5. Check for new challenges
	if eo.challengeSys != nil {
		completed, err := eo.challengeSys.CheckAndAwardChallenges(context)
		if err != nil {
			fmt.Fprintf(out, "⚠️  Warning: Challenge check failed: %v\n", err)
		} else if len(completed) > 0 {
			// Schedule challenges for dopamine timing
			if eo.dopamineTiming != nil {
				for _, challenge := range completed {
					eo.dopamineTiming.ScheduleReward(
						"challenge",
						challenge.ID,
						challenge.Title,
						challenge.Description,
						challenge.Points,
						challenge.Rarity,
						challenge.Icon,
						challenge.Project,
					)
				}
			}
		}
	}

	// 6. Check if rewards should be released now (immediate for first-time, delayed for others)
	if eo.dopamineTiming != nil {
		rewards, err := eo.dopamineTiming.CheckAndReleaseRewards(out)
		if err != nil {
			fmt.Fprintf(out, "⚠️  Warning: Reward release failed: %v\n", err)
		}
		if len(rewards) > 0 {
			// Refresh systems after rewards released
			eo.refreshSystems()
		}
	}

	return nil
}

// refreshSystems refreshes all systems after memory card updates
func (eo *EngagementOrchestrator) refreshSystems() {
	// Reload memory card
	mc, err := LoadMemoryCard()
	if err == nil {
		eo.memoryCard = mc
		// Update brain
		if eo.brain != nil {
			eo.brain.memoryCard = mc
			eo.brain.helper = NewMemoryCardHelper(mc)
		}
		// Update achievement system
		if eo.achievementSys != nil {
			eo.achievementSys.memoryCard = mc
			eo.achievementSys.score = mc.Score
			eo.achievementSys.achievements = mc.Achievements
		}
		// Update challenge system
		if eo.challengeSys != nil {
			eo.challengeSys.memoryCard = mc
		}
		// Update dopamine timing
		if eo.dopamineTiming != nil {
			eo.dopamineTiming.memoryCard = mc
		}
	}
}

// GetPersonalizedGreeting returns a greeting enhanced by all systems
func (eo *EngagementOrchestrator) GetPersonalizedGreeting(context string) string {
	if eo.brain == nil {
		return "Hello! 👋"
	}
	return eo.brain.GetContextualGreeting(context)
}

// GetEngagementSummary returns a summary of user's engagement
func (eo *EngagementOrchestrator) GetEngagementSummary() string {
	if eo.memoryCard == nil {
		return ""
	}

	var summary strings.Builder
	summary.WriteString("📊 Engagement Summary\n")
	summary.WriteString(fmt.Sprintf("   Score: %d points\n", eo.memoryCard.Score))
	summary.WriteString(fmt.Sprintf("   Achievements: %d\n", len(eo.memoryCard.Achievements)))
	summary.WriteString(fmt.Sprintf("   Challenges: %d\n", len(eo.memoryCard.CompletedChallenges)))
	summary.WriteString(fmt.Sprintf("   Relationship Level: %d/100\n", eo.memoryCard.RelationshipLevel))
	summary.WriteString(fmt.Sprintf("   Engagement Score: %.1f%%\n", eo.memoryCard.EngagementScore*100))

	// Time since last reward
	if eo.dopamineTiming != nil {
		timeSince := eo.dopamineTiming.getTimeSinceLastReward()
		if timeSince < 60*time.Minute {
			summary.WriteString(fmt.Sprintf("   Last Reward: %d minutes ago\n", int(timeSince.Minutes())))
		} else {
			summary.WriteString(fmt.Sprintf("   Last Reward: %d hours ago\n", int(timeSince.Hours())))
		}

		// Pending rewards
		if len(eo.dopamineTiming.rewardQueue) > 0 {
			summary.WriteString(fmt.Sprintf("   Pending Rewards: %d (coming soon!)\n", len(eo.dopamineTiming.rewardQueue)))
		}
	}

	return summary.String()
}

// TrackInteraction tracks a user interaction and updates all systems
func (eo *EngagementOrchestrator) TrackInteraction(
	command string,
	userInput string,
	agentResponse string,
	durationSeconds float64,
	sentiment string,
) {
	if eo.memoryCard == nil {
		return
	}

	// Update memory card helper
	helper := NewMemoryCardHelper(eo.memoryCard)
	helper.TrackCommandExecution(command, userInput, agentResponse, durationSeconds, sentiment)

	// Update last command
	eo.memoryCard.LastCommand = command
	eo.memoryCard.LastCommandTime = time.Now()

	// Save memory card
	SaveMemoryCard(eo.memoryCard)

	// Refresh systems
	eo.refreshSystems()
}

// GetNextMilestones returns hints about upcoming milestones
func (eo *EngagementOrchestrator) GetNextMilestones(limit int) []string {
	var hints []string

	if eo.achievementSys != nil {
		// Score milestones
		nextScore := eo.getNextScoreMilestone()
		if nextScore > 0 {
			pointsNeeded := nextScore - eo.memoryCard.Score
			hints = append(hints, fmt.Sprintf("🎯 %d more points to reach %d (Score Milestone)", pointsNeeded, nextScore))
		}
	}

	// Project milestones
	if eo.memoryCard.ProjectsCount < 5 {
		hints = append(hints, fmt.Sprintf("🏗️  Complete %d more projects to reach 5 projects", 5-eo.memoryCard.ProjectsCount))
	}

	// Achievement count
	achievementCount := len(eo.memoryCard.Achievements)
	if achievementCount < 10 {
		hints = append(hints, fmt.Sprintf("🎖️  Earn %d more achievements to reach 10", 10-achievementCount))
	}

	// Relationship milestones
	if eo.memoryCard.RelationshipLevel < 40 {
		hints = append(hints, fmt.Sprintf("🤝 Reach relationship level 40 (currently %d)", eo.memoryCard.RelationshipLevel))
	}

	// Limit results
	if len(hints) > limit {
		hints = hints[:limit]
	}

	return hints
}

// DisplayEngagementDashboard displays a comprehensive engagement dashboard
func (eo *EngagementOrchestrator) DisplayEngagementDashboard(out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "  📊 DoPlan Engagement Dashboard")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")

	// Score and achievements
	fmt.Fprintf(out, "  💰 Total Score: %d points\n", eo.memoryCard.Score)
	fmt.Fprintf(out, "  🏆 Achievements: %d\n", len(eo.memoryCard.Achievements))
	fmt.Fprintf(out, "  🎯 Challenges: %d\n", len(eo.memoryCard.CompletedChallenges))
	fmt.Fprintln(out, "")

	// Note: Feature and phase time tracking available via /sys performance or time-tracker.jsonl
	fmt.Fprintln(out, "  💡 Tip: Check .do/system/time-tracker.jsonl for detailed time analytics")
	fmt.Fprintln(out, "")

	// Relationship
	fmt.Fprintf(out, "  🤝 Relationship Level: %d/100", eo.memoryCard.RelationshipLevel)
	if eo.memoryCard.RelationshipLevel >= 70 {
		fmt.Fprintf(out, " ⭐ Strong!")
	} else if eo.memoryCard.RelationshipLevel >= 40 {
		fmt.Fprintf(out, " 💪 Building!")
	}
	fmt.Fprintln(out, "")

	// Engagement
	fmt.Fprintf(out, "  📈 Engagement: %.0f%%", eo.memoryCard.EngagementScore*100)
	if eo.memoryCard.EngagementScore >= 0.8 {
		fmt.Fprintf(out, " 🔥 Very High!")
	} else if eo.memoryCard.EngagementScore >= 0.4 {
		fmt.Fprintf(out, " 👍 Good!")
	}
	fmt.Fprintln(out, "")

	// Time since last reward
	if eo.dopamineTiming != nil {
		timeSince := eo.dopamineTiming.getTimeSinceLastReward()
		fmt.Fprintf(out, "  ⏰ Last Reward: ")
		if timeSince < 5*time.Minute {
			fmt.Fprintf(out, "Just now! 🎉\n")
		} else if timeSince < 60*time.Minute {
			fmt.Fprintf(out, "%d minutes ago\n", int(timeSince.Minutes()))
		} else {
			fmt.Fprintf(out, "%d hours ago\n", int(timeSince.Hours()))
		}

		// Pending rewards
		if len(eo.dopamineTiming.rewardQueue) > 0 {
			fmt.Fprintf(out, "  ⏳ Pending Rewards: %d (coming soon!)\n", len(eo.dopamineTiming.rewardQueue))
		}
	}

	fmt.Fprintln(out, "")

	// Next milestones
	nextMilestones := eo.GetNextMilestones(3)
	if len(nextMilestones) > 0 {
		fmt.Fprintln(out, "  🎯 Next Milestones:")
		for _, milestone := range nextMilestones {
			fmt.Fprintf(out, "     %s\n", milestone)
		}
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")
}

// getNextScoreMilestone returns the next score milestone
func (eo *EngagementOrchestrator) getNextScoreMilestone() int {
	currentScore := eo.memoryCard.Score
	milestones := []int{100, 250, 500, 1000, 2500, 5000, 10000}

	for _, milestone := range milestones {
		if currentScore < milestone {
			return milestone
		}
	}
	return 0 // No more milestones
}

// detectPainPointResolution detects when a user has overcome a pain point
func (eo *EngagementOrchestrator) detectPainPointResolution(context map[string]interface{}) {
	if eo.memoryCard == nil {
		return
	}

	// Check if user successfully completed a task/feature they struggled with
	if success, ok := context["success"].(bool); ok && success {
		// If a feature was completed successfully and it was in struggled features, mark as resolved
		if featureName, ok := context["feature_name"].(string); ok && featureName != "" {
			for _, struggledFeature := range eo.memoryCard.StruggledFeatures {
				if struggledFeature == featureName {
					// User successfully completed a feature they struggled with
					// Treat as resolved pain point for that feature
					eo.memoryCard.ResolvePainPoint(struggledFeature)
					break
				}
			}
		}

		// Check if user explicitly indicated they resolved a pain point
		if resolved, ok := context["pain_point_resolved"].(string); ok && resolved != "" {
			eo.memoryCard.ResolvePainPoint(resolved)
		}
	}
}
