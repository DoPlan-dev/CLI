package cli

import (
	"fmt"
	"io"
	"strings"
)

// AchievementIntegration provides easy integration of achievements into commands
type AchievementIntegration struct {
	system *AchievementSystem
}

// NewAchievementIntegration creates a new achievement integration helper
func NewAchievementIntegration() (*AchievementIntegration, error) {
	system, err := NewAchievementSystem()
	if err != nil {
		return nil, err
	}

	return &AchievementIntegration{
		system: system,
	}, nil
}

// CheckAndCelebrate checks for achievements and celebrates them
func (ai *AchievementIntegration) CheckAndCelebrate(context map[string]interface{}, out io.Writer) error {
	if ai.system == nil {
		return nil // Graceful degradation
	}

	// Check for achievements
	earned, err := ai.system.CheckAndAwardAchievements(context)
	if err != nil {
		return fmt.Errorf("failed to check achievements: %w", err)
	}

	// Celebrate if any earned
	if len(earned) > 0 {
		CelebrateAchievements(earned, out)
	}

	return nil
}

// CheckOnCommandCompletion checks achievements when a command completes
func (ai *AchievementIntegration) CheckOnCommandCompletion(command string, success bool, context map[string]interface{}, out io.Writer) error {
	if !success {
		return nil // Don't check achievements on failure
	}

	// Add command-specific context
	context["command"] = command
	context["success"] = true

	// Check and celebrate
	return ai.CheckAndCelebrate(context, out)
}

// CheckOnProjectMilestone checks achievements when project reaches a milestone
func (ai *AchievementIntegration) CheckOnProjectMilestone(projectName string, milestone string, context map[string]interface{}, out io.Writer) error {
	context["project"] = projectName
	context["milestone"] = milestone
	context["phase"] = milestone

	return ai.CheckAndCelebrate(context, out)
}

// CheckOnScoreIncrease checks achievements when score increases
func (ai *AchievementIntegration) CheckOnScoreIncrease(newScore int, context map[string]interface{}, out io.Writer) error {
	context["score"] = newScore
	context["achievements"] = len(ai.system.memoryCard.Achievements)

	return ai.CheckAndCelebrate(context, out)
}

// GetScore returns current score
func (ai *AchievementIntegration) GetScore() int {
	if ai.system == nil {
		return 0
	}
	return ai.system.GetScore()
}

// GetAchievementSummary returns formatted achievement summary
func (ai *AchievementIntegration) GetAchievementSummary() string {
	if ai.system == nil {
		return ""
	}
	return GetAchievementSummary(ai.system.memoryCard)
}

// GetNextAchievements returns hints about next achievements
func (ai *AchievementIntegration) GetNextAchievements(limit int) []string {
	if ai.system == nil {
		return []string{}
	}
	return GetNextAchievements(ai.system.memoryCard, limit)
}

// DisplayAchievementProgress displays progress towards next achievements
func (ai *AchievementIntegration) DisplayAchievementProgress(out io.Writer) {
	if ai.system == nil {
		return
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🎯 Achievement Progress")
	fmt.Fprintln(out, strings.Repeat("-", 40))

	score := ai.system.GetScore()
	fmt.Fprintf(out, "Current Score: %d points\n", score)
	fmt.Fprintf(out, "Total Achievements: %d\n", len(ai.system.memoryCard.Achievements))

	nextHints := ai.GetNextAchievements(3)
	if len(nextHints) > 0 {
		fmt.Fprintln(out, "\nNext Achievements:")
		for _, hint := range nextHints {
			fmt.Fprintf(out, "  %s\n", hint)
		}
	}

	fmt.Fprintln(out, "")
}
