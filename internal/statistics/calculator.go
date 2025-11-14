package statistics

import (
	"sort"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Calculator computes metrics from collected data
type Calculator struct {
	projectStartDate time.Time
}

// NewCalculator creates a new statistics calculator
func NewCalculator(projectStartDate time.Time) *Calculator {
	return &Calculator{
		projectStartDate: projectStartDate,
	}
}

// Calculate computes all metrics from statistics data
func (c *Calculator) Calculate(data *StatisticsData, state *models.State, githubData *models.GitHubData) *StatisticsMetrics {
	metrics := &StatisticsMetrics{
		CalculatedAt: time.Now(),
	}

	metrics.Velocity = c.CalculateVelocity(data, state, githubData)
	metrics.Completion = c.CalculateCompletionRates(data, state)
	metrics.Time = c.CalculateTimeMetrics(data, state)
	metrics.Quality = c.CalculateQualityMetrics(data, githubData)
	metrics.Testing = c.CalculateTestingMetrics(data.Testing)

	return metrics
}

// CalculateVelocity computes velocity metrics
func (c *Calculator) CalculateVelocity(data *StatisticsData, state *models.State, githubData *models.GitHubData) *VelocityMetrics {
	daysSinceStart := c.daysSinceStart()
	if daysSinceStart == 0 {
		daysSinceStart = 1 // Avoid division by zero
	}

	weeksSinceStart := float64(daysSinceStart) / 7.0
	if weeksSinceStart == 0 {
		weeksSinceStart = 1
	}

	metrics := &VelocityMetrics{}

	// Features per day/week
	if data.State != nil {
		completedFeatures := float64(data.State.CompletedFeatures)
		metrics.FeaturesPerDay = completedFeatures / float64(daysSinceStart)
		metrics.FeaturesPerWeek = completedFeatures / weeksSinceStart
	}

	// Commits per day/week
	if data.GitHub != nil {
		totalCommits := float64(data.GitHub.TotalCommits)
		metrics.CommitsPerDay = totalCommits / float64(daysSinceStart)
		metrics.CommitsPerWeek = totalCommits / weeksSinceStart
	}

	// Tasks per day
	if data.Tasks != nil {
		completedTasks := float64(data.Tasks.CompletedTasks)
		metrics.TasksPerDay = completedTasks / float64(daysSinceStart)
	}

	// PRs per week
	if data.GitHub != nil {
		mergedPRs := float64(data.GitHub.MergedPRs)
		metrics.PRsPerWeek = mergedPRs / weeksSinceStart
	}

	return metrics
}

// CalculateCompletionRates computes completion percentages
func (c *Calculator) CalculateCompletionRates(data *StatisticsData, state *models.State) *CompletionRates {
	rates := &CompletionRates{
		Phases:   make(map[string]int),
		Features: make(map[string]int),
	}

	// Overall completion
	if data.State != nil && data.State.TotalFeatures > 0 {
		rates.Overall = (data.State.CompletedFeatures * 100) / data.State.TotalFeatures
	}

	// Phase completion
	for _, phase := range state.Phases {
		phaseFeatures := 0
		completedPhaseFeatures := 0

		for _, feature := range state.Features {
			if feature.Phase == phase.ID {
				phaseFeatures++
				if feature.Status == "complete" {
					completedPhaseFeatures++
				}
			}
		}

		if phaseFeatures > 0 {
			rates.Phases[phase.ID] = (completedPhaseFeatures * 100) / phaseFeatures
		}
	}

	// Feature completion (individual)
	for _, feature := range state.Features {
		rates.Features[feature.ID] = feature.Progress
	}

	// Task completion
	if data.Tasks != nil {
		rates.Tasks = data.Tasks.CompletionRate
	}

	return rates
}

// CalculateTimeMetrics computes time-related statistics
func (c *Calculator) CalculateTimeMetrics(data *StatisticsData, state *models.State) *TimeMetrics {
	metrics := &TimeMetrics{
		ProjectStartDate: c.projectStartDate,
		DaysSinceStart:   c.daysSinceStart(),
	}

	// Calculate average feature time
	completedFeatures := 0
	totalFeatureDays := 0.0

	for _, feature := range state.Features {
		if feature.Status == "complete" && feature.StartDate != "" && feature.TargetDate != "" {
			start, err1 := time.Parse("2006-01-02", feature.StartDate)
			target, err2 := time.Parse("2006-01-02", feature.TargetDate)

			if err1 == nil && err2 == nil {
				duration := target.Sub(start).Hours() / 24.0
				if duration > 0 {
					totalFeatureDays += duration
					completedFeatures++
				}
			}
		}
	}

	if completedFeatures > 0 {
		metrics.AvgFeatureTime = totalFeatureDays / float64(completedFeatures)
	}

	// Calculate average phase time
	completedPhases := 0
	totalPhaseDays := 0.0

	for _, phase := range state.Phases {
		if phase.Status == "complete" && phase.StartDate != "" && phase.TargetDate != "" {
			start, err1 := time.Parse("2006-01-02", phase.StartDate)
			target, err2 := time.Parse("2006-01-02", phase.TargetDate)

			if err1 == nil && err2 == nil {
				duration := target.Sub(start).Hours() / 24.0
				if duration > 0 {
					totalPhaseDays += duration
					completedPhases++
				}
			}
		}
	}

	if completedPhases > 0 {
		metrics.AvgPhaseTime = totalPhaseDays / float64(completedPhases)
	}

	// Estimate completion date
	if metrics.AvgFeatureTime > 0 && data.State != nil {
		remainingFeatures := data.State.TotalFeatures - data.State.CompletedFeatures
		if remainingFeatures > 0 {
			daysRemaining := int(metrics.AvgFeatureTime * float64(remainingFeatures))
			metrics.EstimatedCompletion = time.Now().AddDate(0, 0, daysRemaining)
		}
	}

	return metrics
}

// CalculateQualityMetrics computes quality-related statistics
func (c *Calculator) CalculateQualityMetrics(data *StatisticsData, githubData *models.GitHubData) *QualityMetrics {
	metrics := &QualityMetrics{}

	// PR merge rate
	if data.GitHub != nil && data.GitHub.TotalPRs > 0 {
		metrics.PRMergeRate = (float64(data.GitHub.MergedPRs) / float64(data.GitHub.TotalPRs)) * 100
	}

	// Checkpoint frequency
	daysSinceStart := c.daysSinceStart()
	if daysSinceStart > 0 && data.Checkpoints != nil {
		weeksSinceStart := float64(daysSinceStart) / 7.0
		if weeksSinceStart > 0 {
			metrics.CheckpointFrequency = float64(data.Checkpoints.TotalCheckpoints) / weeksSinceStart
		}
	}

	// Average branch lifetime (simplified - would need more data)
	if data.GitHub != nil && data.GitHub.MergedPRs > 0 {
		// Estimate based on merged PRs
		metrics.AvgBranchLifetime = 3.5 // Default estimate, would need actual branch data
	}

	return metrics
}

// CalculateTestingMetrics computes coverage metrics from testing data
func (c *Calculator) CalculateTestingMetrics(testing *TestingStats) *TestingMetrics {
	if testing == nil || testing.TotalStatements == 0 {
		return nil
	}

	metrics := &TestingMetrics{
		OverallCoverage: percentage(testing.CoveredStatements, testing.TotalStatements),
	}

	for _, pkgStats := range testing.PackageStats {
		if pkgStats.Statements == 0 {
			continue
		}
		metrics.Packages = append(metrics.Packages, PackageCoverageMetric{
			Name:     pkgStats.Name,
			Coverage: percentage(pkgStats.CoveredStatements, pkgStats.Statements),
		})
	}

	sort.Slice(metrics.Packages, func(i, j int) bool {
		return metrics.Packages[i].Name < metrics.Packages[j].Name
	})

	return metrics
}

func percentage(covered, total int) float64 {
	if total == 0 {
		return 0
	}
	return (float64(covered) / float64(total)) * 100
}

// daysSinceStart calculates days since project start
func (c *Calculator) daysSinceStart() int {
	if c.projectStartDate.IsZero() {
		return 0
	}
	return int(time.Since(c.projectStartDate).Hours() / 24)
}
