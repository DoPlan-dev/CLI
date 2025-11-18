package dashboard

import (
	"fmt"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// ActivityData represents activity data (to avoid import cycle)
type ActivityData struct {
	Last24Hours    ActivityPeriodData
	Last7Days      ActivityPeriodData
	RecentActivity []ActivityItemData
}

// ActivityPeriodData represents activity for a time period
type ActivityPeriodData struct {
	Commits        int
	TasksCompleted int
	FilesChanged   int
}

// ActivityItemData represents a single activity item
type ActivityItemData struct {
	Type      string
	Message   string
	Timestamp string
}

// ActivityGenerator generates activity feed from various sources
type ActivityGenerator struct {
	state        *models.State
	githubData   *models.GitHubData
	progressData map[string]*ProgressData
}

// NewActivityGenerator creates a new activity generator
func NewActivityGenerator(state *models.State, githubData *models.GitHubData, progressData map[string]*ProgressData) *ActivityGenerator {
	return &ActivityGenerator{
		state:        state,
		githubData:   githubData,
		progressData: progressData,
	}
}

// GenerateActivityFeed generates a comprehensive activity feed
func (a *ActivityGenerator) GenerateActivityFeed() ActivityData {
	now := time.Now()
	last24Hours := now.Add(-24 * time.Hour)
	last7Days := now.Add(-7 * 24 * time.Hour)

	activity := ActivityData{
		Last24Hours:    ActivityPeriodData{},
		Last7Days:      ActivityPeriodData{},
		RecentActivity: []ActivityItemData{},
	}

	// Collect activities from different sources
	var allActivities []ActivityItemData

	// Add commits
	if a.githubData != nil {
		commits24h := 0
		commits7d := 0
		for _, commit := range a.githubData.Commits {
			commitTime, err := ParseTime(commit.Date)
			if err != nil {
				continue
			}

			if commitTime.After(last24Hours) {
				commits24h++
			}
			if commitTime.After(last7Days) {
				commits7d++
			}

			// Add to recent activity (last 10)
			if len(allActivities) < 10 {
				allActivities = append(allActivities, ActivityItemData{
					Type:      "commit",
					Message:   fmt.Sprintf("📝 commit: %s", commit.Message),
					Timestamp: commit.Date,
				})
			}
		}
		activity.Last24Hours.Commits = commits24h
		activity.Last7Days.Commits = commits7d
	}

	// Add task completions from progress data
	tasks24h := 0
	tasks7d := 0
	for _, progress := range a.progressData {
		if progress.LastUpdated.IsZero() {
			continue
		}

		if progress.LastUpdated.After(last24Hours) {
			tasks24h++
		}
		if progress.LastUpdated.After(last7Days) {
			tasks7d++
		}

		// Add feature status changes
		if progress.Status != "" && progress.LastUpdated.After(last7Days) {
			statusIcon := "✨"
			if progress.Status == "complete" {
				statusIcon = "✅"
			} else if progress.Status == "in-progress" {
				statusIcon = "🚧"
			}

			featureName := progress.FeatureName
			if featureName == "" {
				featureName = progress.FeatureID
			}

			allActivities = append(allActivities, ActivityItemData{
				Type:      "feature",
				Message:   fmt.Sprintf("%s feature: %s %s", statusIcon, featureName, progress.Status),
				Timestamp: progress.LastUpdated.Format(time.RFC3339),
			})
		}
	}
	activity.Last24Hours.TasksCompleted = tasks24h
	activity.Last7Days.TasksCompleted = tasks7d

	// Add PR merges
	if a.githubData != nil {
		for _, pr := range a.githubData.PRs {
			if pr.Status == "merged" || pr.Status == "closed" {
				// Try to parse PR timestamp (would need to be added to GitHubData)
				allActivities = append(allActivities, ActivityItemData{
					Type:      "pr",
					Message:   fmt.Sprintf("🔀 PR #%d: %s merged", pr.Number, pr.Title),
					Timestamp: time.Now().Format(time.RFC3339), // TODO: Get actual merge time
				})
			}
		}
	}

	// Add phase completions from state
	if a.state != nil {
		for _, phase := range a.state.Phases {
			if phase.Status == "complete" {
				allActivities = append(allActivities, ActivityItemData{
					Type:      "phase",
					Message:   fmt.Sprintf("🎯 phase: %s completed", phase.Name),
					Timestamp: time.Now().Format(time.RFC3339), // TODO: Get actual completion time
				})
			}
		}
	}

	// Sort activities by timestamp (newest first)
	activity.RecentActivity = sortActivitiesByTime(allActivities)

	// Limit to last 10
	if len(activity.RecentActivity) > 10 {
		activity.RecentActivity = activity.RecentActivity[:10]
	}

	return activity
}

// FormatTimeAgo formats a timestamp as relative time
func FormatTimeAgo(timestamp string) string {
	t, err := ParseTime(timestamp)
	if err != nil {
		return timestamp
	}

	now := time.Now()
	diff := now.Sub(t)

	if diff < time.Minute {
		return "just now"
	}

	if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf("%dm ago", minutes)
	}

	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		return fmt.Sprintf("%dh ago", hours)
	}

	days := int(diff.Hours() / 24)
	return fmt.Sprintf("%dd ago", days)
}

// ParseTime parses various time formats (public for testing)
func ParseTime(timeStr string) (time.Time, error) {
	// Try RFC3339 first
	if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
		return t, nil
	}

	// Try ISO 8601
	if t, err := time.Parse(time.RFC3339Nano, timeStr); err == nil {
		return t, nil
	}

	// Try common formats
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", timeStr)
}

// sortActivitiesByTime sorts activities by timestamp (newest first)
func sortActivitiesByTime(activities []ActivityItemData) []ActivityItemData {
	// Simple bubble sort (can be optimized if needed)
	for i := 0; i < len(activities); i++ {
		for j := i + 1; j < len(activities); j++ {
			timeI, errI := ParseTime(activities[i].Timestamp)
			timeJ, errJ := ParseTime(activities[j].Timestamp)

			if errI != nil || errJ != nil {
				continue
			}

			if timeJ.After(timeI) {
				activities[i], activities[j] = activities[j], activities[i]
			}
		}
	}

	return activities
}

// CalculateActivityPeriods calculates activity for specific time periods
func CalculateActivityPeriods(activities []ActivityItemData, since time.Time) ActivityPeriodData {
	period := ActivityPeriodData{}

	for _, activity := range activities {
		activityTime, err := ParseTime(activity.Timestamp)
		if err != nil {
			continue
		}

		if activityTime.After(since) {
			switch activity.Type {
			case "commit":
				period.Commits++
			case "task", "feature":
				period.TasksCompleted++
			}
			// FilesChanged would need to be tracked separately
		}
	}

	return period
}
