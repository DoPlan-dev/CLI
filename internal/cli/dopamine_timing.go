package cli

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// DopamineTiming manages reward scheduling for optimal dopamine release
type DopamineTiming struct {
	memoryCard  *MemoryCard
	rewardQueue []PendingReward
}

// PendingReward represents a reward waiting to be released
type PendingReward struct {
	Type         string // "achievement", "challenge", "milestone"
	ID           string
	Title        string
	Description  string
	Points       int
	Rarity       string
	Icon         string
	EarnedAt     time.Time
	ScheduledFor time.Time
	Project      string
}

// NewDopamineTiming creates a new dopamine timing system
func NewDopamineTiming() (*DopamineTiming, error) {
	mc, err := LoadMemoryCard()
	if err != nil {
		return nil, err
	}

	dt := &DopamineTiming{
		memoryCard:  mc,
		rewardQueue: []PendingReward{},
	}

	// Load pending rewards from memory card
	dt.loadPendingRewards()

	return dt, nil
}

// ScheduleReward schedules a reward for optimal dopamine release
func (dt *DopamineTiming) ScheduleReward(rewardType, id, title, description string, points int, rarity, icon, project string) {
	now := time.Now()

	// Calculate optimal release time
	releaseTime := dt.calculateOptimalReleaseTime(now)

	reward := PendingReward{
		Type:         rewardType,
		ID:           id,
		Title:        title,
		Description:  description,
		Points:       points,
		Rarity:       rarity,
		Icon:         icon,
		EarnedAt:     now,
		ScheduledFor: releaseTime,
		Project:      project,
	}

	dt.rewardQueue = append(dt.rewardQueue, reward)
	dt.savePendingRewards()
}

// CheckAndReleaseRewards checks if any rewards should be released
func (dt *DopamineTiming) CheckAndReleaseRewards(out io.Writer) ([]PendingReward, error) {
	now := time.Now()
	var readyToRelease []PendingReward
	var remainingQueue []PendingReward

	// Check each pending reward
	for _, reward := range dt.rewardQueue {
		if now.After(reward.ScheduledFor) || now.Equal(reward.ScheduledFor) {
			readyToRelease = append(readyToRelease, reward)
		} else {
			remainingQueue = append(remainingQueue, reward)
		}
	}

	// Update queue
	dt.rewardQueue = remainingQueue
	dt.savePendingRewards()

	// Release rewards if any are ready
	if len(readyToRelease) > 0 {
		dt.releaseRewards(readyToRelease, out)
	}

	return readyToRelease, nil
}

// calculateOptimalReleaseTime calculates when to release reward for maximum dopamine
func (dt *DopamineTiming) calculateOptimalReleaseTime(earnedAt time.Time) time.Time {
	// Get time since last reward
	timeSinceLastReward := dt.getTimeSinceLastReward()

	// Get user engagement level
	engagementLevel := dt.getEngagementLevel()

	// Calculate delay based on engagement and time since last reward
	var delay time.Duration

	// Strategy: Variable Interval Reinforcement
	if timeSinceLastReward < 5*time.Minute {
		// Very recent reward - delay longer to build anticipation
		delay = 30 * time.Minute
	} else if timeSinceLastReward < 15*time.Minute {
		// Recent reward - medium delay
		delay = 15 * time.Minute
	} else if timeSinceLastReward < 30*time.Minute {
		// Some time passed - shorter delay
		delay = 5 * time.Minute
	} else if timeSinceLastReward < 60*time.Minute {
		// Good gap - immediate or short delay
		delay = 2 * time.Minute
	} else if timeSinceLastReward < 120*time.Minute {
		// Long gap - user might be feeling bad, release soon but with buildup
		delay = 1 * time.Minute
	} else {
		// Very long gap - user definitely feeling bad, release immediately with celebration
		delay = 0
	}

	// Adjust based on engagement level
	if engagementLevel == "high" {
		// Highly engaged users can handle longer delays
		delay = time.Duration(float64(delay) * 1.5)
	} else if engagementLevel == "low" {
		// Low engagement - release faster to re-engage
		delay = time.Duration(float64(delay) * 0.5)
	}

	// Adjust based on reward rarity
	// Epic/Legendary rewards can wait longer (bigger payoff)
	// Common rewards should come faster

	return earnedAt.Add(delay)
}

// getTimeSinceLastReward returns time since last achievement/challenge
func (dt *DopamineTiming) getTimeSinceLastReward() time.Duration {
	if dt.memoryCard == nil {
		return 24 * time.Hour // Default to long time
	}

	// Check last achievement
	var lastRewardTime time.Time
	if len(dt.memoryCard.Achievements) > 0 {
		lastAchievement := dt.memoryCard.Achievements[len(dt.memoryCard.Achievements)-1]
		lastRewardTime = lastAchievement.EarnedAt
	}

	// Check last interaction
	if dt.memoryCard.LastInteraction.After(lastRewardTime) {
		// Use last interaction if more recent
		return time.Since(dt.memoryCard.LastInteraction)
	}

	if lastRewardTime.IsZero() {
		return 24 * time.Hour
	}

	return time.Since(lastRewardTime)
}

// getEngagementLevel returns user's engagement level
func (dt *DopamineTiming) getEngagementLevel() string {
	if dt.memoryCard == nil {
		return "medium"
	}

	score := dt.memoryCard.EngagementScore
	if score >= 0.8 {
		return "high"
	} else if score >= 0.4 {
		return "medium"
	}
	return "low"
}

// releaseRewards releases scheduled rewards with celebration
func (dt *DopamineTiming) releaseRewards(rewards []PendingReward, out io.Writer) {
	if len(rewards) == 0 {
		return
	}

	// Calculate total points
	totalPoints := 0
	for _, reward := range rewards {
		totalPoints += reward.Points
	}

	// Determine celebration level based on time since last reward
	timeSinceLast := dt.getTimeSinceLastReward()
	celebrationLevel := dt.determineCelebrationLevel(timeSinceLast)

	// Display celebration
	dt.displayCelebration(rewards, totalPoints, celebrationLevel, out)

	// Update memory card with rewards
	dt.applyRewardsToMemoryCard(rewards)
}

// determineCelebrationLevel determines how excited to be
func (dt *DopamineTiming) determineCelebrationLevel(timeSinceLast time.Duration) string {
	if timeSinceLast >= 120*time.Minute {
		return "extreme" // User was waiting, make it special!
	} else if timeSinceLast >= 60*time.Minute {
		return "high"
	} else if timeSinceLast >= 30*time.Minute {
		return "medium"
	}
	return "normal"
}

// displayCelebration displays the reward celebration
func (dt *DopamineTiming) displayCelebration(rewards []PendingReward, totalPoints int, level string, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")

	// Build anticipation message based on time since last reward
	timeSinceLast := dt.getTimeSinceLastReward()
	if timeSinceLast >= 120*time.Minute {
		fmt.Fprintf(out, "  🎊🎊🎊  WOW! You've been working hard! Here's your reward!  🎊🎊🎊\n")
		fmt.Fprintf(out, "  (You've been building for over %d minutes - amazing dedication!)\n", int(timeSinceLast.Minutes()))
	} else if timeSinceLast >= 60*time.Minute {
		fmt.Fprintf(out, "  🎉🎉  Great work! Your achievements are here!  🎉🎉\n")
	} else {
		fmt.Fprintf(out, "  🎯  Achievement Unlocked!  🎯\n")
	}

	fmt.Fprintln(out, "")

	if len(rewards) == 1 {
		reward := rewards[0]
		fmt.Fprintf(out, "  %s  %s  %s\n", reward.Icon, reward.Title, reward.Icon)
		fmt.Fprintf(out, "  %s\n", reward.Description)
		fmt.Fprintf(out, "  💰 +%d points\n", reward.Points)
		fmt.Fprintf(out, "  ⭐ %s\n", strings.Title(reward.Rarity))
	} else {
		fmt.Fprintf(out, "  🚀🚀🚀  INCREDIBLE! You earned %d rewards!  🚀🚀🚀\n\n", len(rewards))

		for i, reward := range rewards {
			fmt.Fprintf(out, "  %d. %s %s\n", i+1, reward.Icon, reward.Title)
			fmt.Fprintf(out, "     %s (+%d points", reward.Description, reward.Points)
			if reward.Rarity == "legendary" || reward.Rarity == "epic" {
				fmt.Fprintf(out, " ⭐ %s", strings.Title(reward.Rarity))
			}
			fmt.Fprintf(out, ")\n\n")
		}

		fmt.Fprintf(out, "  💰 Total Points: +%d\n", totalPoints)
	}

	// Add encouragement based on level
	if level == "extreme" {
		fmt.Fprintf(out, "\n  🌟 You've been so patient and persistent! This reward is well-deserved! 🌟\n")
	} else if level == "high" {
		fmt.Fprintf(out, "\n  💪 Your dedication is paying off! Keep up the amazing work! 💪\n")
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, strings.Repeat("=", 70))
	fmt.Fprintln(out, "")
}

// applyRewardsToMemoryCard applies rewards to memory card
func (dt *DopamineTiming) applyRewardsToMemoryCard(rewards []PendingReward) {
	if dt.memoryCard == nil {
		return
	}

	for _, reward := range rewards {
		// Add to achievements
		achievement := Achievement{
			ID:          reward.ID,
			Title:       reward.Title,
			Description: reward.Description,
			Category:    reward.Type,
			EarnedAt:    reward.EarnedAt,
			Project:     reward.Project,
		}
		dt.memoryCard.Achievements = append(dt.memoryCard.Achievements, achievement)

		// Increase score
		dt.memoryCard.Score += reward.Points

		// Create memorable moment
		dt.memoryCard.AddMemorableMoment(
			"achievement",
			reward.Title,
			reward.Description,
			"proud",
			fmt.Sprintf("Earned %s reward", reward.Rarity),
		)
	}

	// Save memory card
	SaveMemoryCard(dt.memoryCard)
}

// loadPendingRewards loads pending rewards from memory card
func (dt *DopamineTiming) loadPendingRewards() {
	// In a real implementation, this would load from memory card or separate file
	// For now, we'll check memory card preferences
	if dt.memoryCard != nil && dt.memoryCard.Preferences != nil {
		if pending, ok := dt.memoryCard.Preferences["pending_rewards"].([]interface{}); ok {
			// Convert back to PendingReward (simplified for now)
			_ = pending
		}
	}
}

// savePendingRewards saves pending rewards to memory card
func (dt *DopamineTiming) savePendingRewards() {
	if dt.memoryCard == nil {
		return
	}

	if dt.memoryCard.Preferences == nil {
		dt.memoryCard.Preferences = make(map[string]interface{})
	}

	// Convert to serializable format
	pending := make([]map[string]interface{}, 0, len(dt.rewardQueue))
	for _, reward := range dt.rewardQueue {
		pending = append(pending, map[string]interface{}{
			"type":          reward.Type,
			"id":            reward.ID,
			"title":         reward.Title,
			"description":   reward.Description,
			"points":        reward.Points,
			"rarity":        reward.Rarity,
			"icon":          reward.Icon,
			"earned_at":     reward.EarnedAt,
			"scheduled_for": reward.ScheduledFor,
			"project":       reward.Project,
		})
	}

	dt.memoryCard.Preferences["pending_rewards"] = pending
	SaveMemoryCard(dt.memoryCard)
}

// GetAnticipationMessage returns a message to build anticipation
func (dt *DopamineTiming) GetAnticipationMessage() string {
	timeSinceLast := dt.getTimeSinceLastReward()

	if timeSinceLast >= 120*time.Minute {
		return "💪 You've been working hard! Something exciting is coming your way..."
	} else if timeSinceLast >= 60*time.Minute {
		return "🎯 Keep going! You're building towards something great..."
	} else if timeSinceLast >= 30*time.Minute {
		return "✨ Great progress! Your efforts are being recognized..."
	}

	return ""
}

// ShouldShowAnticipation returns whether to show anticipation message
func (dt *DopamineTiming) ShouldShowAnticipation() bool {
	timeSinceLast := dt.getTimeSinceLastReward()
	return timeSinceLast >= 30*time.Minute && len(dt.rewardQueue) > 0
}
