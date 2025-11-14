package statistics

import (
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestNewCalculator(t *testing.T) {
	startDate := time.Now().Add(-30 * 24 * time.Hour)
	calculator := NewCalculator(startDate)

	assert.NotNil(t, calculator)
	assert.Equal(t, startDate, calculator.projectStartDate)
}

func TestCalculateVelocity(t *testing.T) {
	startDate := time.Now().Add(-30 * 24 * time.Hour)
	calculator := NewCalculator(startDate)

	data := &StatisticsData{
		State: &StateData{
			CompletedFeatures: 15,
		},
		GitHub: &GitHubStats{
			TotalCommits: 90,
			MergedPRs:    6,
		},
		Tasks: &TaskStats{
			CompletedTasks: 60,
		},
	}

	state := &models.State{}
	githubData := &models.GitHubData{}

	velocity := calculator.CalculateVelocity(data, state, githubData)

	assert.NotNil(t, velocity)
	assert.InDelta(t, 0.5, velocity.FeaturesPerDay, 0.1) // 15/30
	assert.InDelta(t, 3.0, velocity.CommitsPerDay, 0.1)  // 90/30
	assert.InDelta(t, 2.0, velocity.TasksPerDay, 0.1)    // 60/30
	assert.InDelta(t, 1.4, velocity.PRsPerWeek, 0.1)     // 6/4.3 weeks
}

func TestCalculateCompletionRates(t *testing.T) {
	startDate := time.Now()
	calculator := NewCalculator(startDate)

	data := &StatisticsData{
		State: &StateData{
			TotalFeatures:     10,
			CompletedFeatures: 6,
		},
		Tasks: &TaskStats{
			TotalTasks:     20,
			CompletedTasks: 15,
			CompletionRate: 75,
		},
	}

	state := &models.State{
		Phases: []models.Phase{
			{ID: "phase-1", Name: "Phase 1"},
			{ID: "phase-2", Name: "Phase 2"},
		},
		Features: []models.Feature{
			{ID: "f1", Phase: "phase-1", Status: "complete", Progress: 100},
			{ID: "f2", Phase: "phase-1", Status: "complete", Progress: 100},
			{ID: "f3", Phase: "phase-1", Status: "in-progress", Progress: 50},
			{ID: "f4", Phase: "phase-2", Status: "complete", Progress: 100},
			{ID: "f5", Phase: "phase-2", Status: "pending", Progress: 0},
		},
	}

	rates := calculator.CalculateCompletionRates(data, state)

	assert.NotNil(t, rates)
	assert.Equal(t, 60, rates.Overall) // 6/10 * 100
	assert.Equal(t, 75, rates.Tasks)
	assert.Equal(t, 66, rates.Phases["phase-1"]) // 2/3 * 100
	assert.Equal(t, 50, rates.Phases["phase-2"]) // 1/2 * 100
	assert.Equal(t, 100, rates.Features["f1"])
	assert.Equal(t, 50, rates.Features["f3"])
}

func TestCalculateTimeMetrics(t *testing.T) {
	startDate := time.Now().Add(-30 * 24 * time.Hour)
	calculator := NewCalculator(startDate)

	data := &StatisticsData{
		State: &StateData{
			TotalFeatures:     10,
			CompletedFeatures: 5,
		},
	}

	state := &models.State{
		Features: []models.Feature{
			{
				ID:         "f1",
				Status:     "complete",
				StartDate:  time.Now().Add(-20 * 24 * time.Hour).Format("2006-01-02"),
				TargetDate: time.Now().Add(-10 * 24 * time.Hour).Format("2006-01-02"),
			},
			{
				ID:         "f2",
				Status:     "complete",
				StartDate:  time.Now().Add(-15 * 24 * time.Hour).Format("2006-01-02"),
				TargetDate: time.Now().Add(-5 * 24 * time.Hour).Format("2006-01-02"),
			},
		},
	}

	metrics := calculator.CalculateTimeMetrics(data, state)

	assert.NotNil(t, metrics)
	assert.Equal(t, startDate, metrics.ProjectStartDate)
	assert.Equal(t, 30, metrics.DaysSinceStart)
	assert.InDelta(t, 10.0, metrics.AvgFeatureTime, 1.0) // Average of 10 and 10 days
}

func TestCalculateQualityMetrics(t *testing.T) {
	startDate := time.Now().Add(-30 * 24 * time.Hour)
	calculator := NewCalculator(startDate)

	data := &StatisticsData{
		GitHub: &GitHubStats{
			TotalPRs:  10,
			MergedPRs: 8,
		},
		Checkpoints: &CheckpointStats{
			TotalCheckpoints: 12,
		},
	}

	githubData := &models.GitHubData{}

	metrics := calculator.CalculateQualityMetrics(data, githubData)

	assert.NotNil(t, metrics)
	assert.Equal(t, 80.0, metrics.PRMergeRate)               // 8/10 * 100
	assert.InDelta(t, 2.8, metrics.CheckpointFrequency, 0.1) // 12/4.3 weeks
}

func TestDaysSinceStart(t *testing.T) {
	startDate := time.Now().Add(-15 * 24 * time.Hour)
	calculator := NewCalculator(startDate)

	days := calculator.daysSinceStart()
	assert.Equal(t, 15, days)
}

func TestDaysSinceStart_Zero(t *testing.T) {
	calculator := NewCalculator(time.Time{})
	days := calculator.daysSinceStart()
	assert.Equal(t, 0, days)
}
