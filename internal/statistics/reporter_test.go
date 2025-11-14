package statistics

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewReporter(t *testing.T) {
	reporter := NewReporter()
	assert.NotNil(t, reporter)
}

func TestReporter_ReportCLI(t *testing.T) {
	t.Setenv("DOPLAN_NO_ANIMATION", "1")
	reporter := NewReporter()
	metrics := &StatisticsMetrics{
		Velocity: &VelocityMetrics{
			FeaturesPerDay:  1.5,
			FeaturesPerWeek: 10.5,
			CommitsPerDay:   5.0,
			CommitsPerWeek:  35.0,
			TasksPerDay:     10.0,
			PRsPerWeek:      2.0,
		},
		Completion: &CompletionRates{
			Overall: 75,
			Tasks:   80,
			Phases:  map[string]int{"phase1": 50, "phase2": 100},
		},
		Time: &TimeMetrics{
			DaysSinceStart:      30,
			AvgFeatureTime:      5.5,
			AvgPhaseTime:        15.0,
			EstimatedCompletion: time.Now().Add(30 * 24 * time.Hour),
		},
		Quality: &QualityMetrics{
			PRMergeRate:         85.5,
			CheckpointFrequency: 2.5,
			AvgBranchLifetime:   3.0,
		},
		Trends: &Trends{
			VelocityTrend:    "improving",
			CompletionTrend:  "stable",
			QualityTrend:     "improving",
			VelocityChange:   10.5,
			CompletionChange: 5.0,
		},
		Testing: &TestingMetrics{
			OverallCoverage: 82.3,
			Packages: []PackageCoverageMetric{
				{Name: "github.com/example/pkg", Coverage: 80},
			},
		},
	}

	err := reporter.ReportCLI(metrics)
	assert.NoError(t, err)
}

func TestReporter_ReportCLI_NilMetrics(t *testing.T) {
	t.Setenv("DOPLAN_NO_ANIMATION", "1")
	reporter := NewReporter()
	metrics := &StatisticsMetrics{}

	err := reporter.ReportCLI(metrics)
	assert.NoError(t, err)
}

func TestReporter_ReportJSON(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	reporter := NewReporter()
	metrics := &StatisticsMetrics{
		Velocity: &VelocityMetrics{
			FeaturesPerDay: 1.5,
		},
	}

	outputPath := filepath.Join(projectRoot, "stats.json")
	err := reporter.ReportJSON(metrics, outputPath)
	require.NoError(t, err)

	// Verify file exists
	assert.FileExists(t, outputPath)

	// Verify content
	data, err := os.ReadFile(outputPath)
	require.NoError(t, err)
	assert.Contains(t, string(data), "featuresPerDay")
}

func TestReporter_ReportJSON_Stdout(t *testing.T) {
	reporter := NewReporter()
	metrics := &StatisticsMetrics{
		Velocity: &VelocityMetrics{
			FeaturesPerDay: 1.5,
		},
	}

	err := reporter.ReportJSON(metrics, "")
	assert.NoError(t, err)
}

func TestReporter_ReportMarkdown(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	reporter := NewReporter()
	metrics := &StatisticsMetrics{
		Velocity: &VelocityMetrics{
			FeaturesPerDay:  1.5,
			FeaturesPerWeek: 10.5,
			CommitsPerDay:   5.0,
			TasksPerDay:     10.0,
			PRsPerWeek:      2.0,
		},
		Completion: &CompletionRates{
			Overall: 75,
			Tasks:   80,
			Phases:  map[string]int{"phase1": 50},
		},
		Time: &TimeMetrics{
			DaysSinceStart: 30,
			AvgFeatureTime: 5.5,
		},
		Quality: &QualityMetrics{
			PRMergeRate: 85.5,
		},
		Testing: &TestingMetrics{
			OverallCoverage: 90.5,
			Packages: []PackageCoverageMetric{
				{Name: "pkg/a", Coverage: 95},
			},
		},
	}

	outputPath := filepath.Join(projectRoot, "stats.md")
	err := reporter.ReportMarkdown(metrics, outputPath)
	require.NoError(t, err)

	assert.FileExists(t, outputPath)
	data, err := os.ReadFile(outputPath)
	require.NoError(t, err)
	assert.Contains(t, string(data), "# DoPlan Statistics")
	assert.Contains(t, string(data), "Features/day")
	assert.Contains(t, string(data), "Testing Metrics")
}

func TestReporter_ReportHTML(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	reporter := NewReporter()
	metrics := &StatisticsMetrics{
		Velocity: &VelocityMetrics{
			FeaturesPerDay: 1.5,
		},
		Completion: &CompletionRates{
			Overall: 75,
		},
		Time: &TimeMetrics{
			DaysSinceStart: 30,
		},
		Quality: &QualityMetrics{
			PRMergeRate: 85.5,
		},
		Testing: &TestingMetrics{
			OverallCoverage: 70.0,
			Packages: []PackageCoverageMetric{
				{Name: "pkg/html", Coverage: 65},
			},
		},
	}

	outputPath := filepath.Join(projectRoot, "stats.html")
	err := reporter.ReportHTML(metrics, outputPath)
	require.NoError(t, err)

	assert.FileExists(t, outputPath)
	data, err := os.ReadFile(outputPath)
	require.NoError(t, err)
	assert.Contains(t, string(data), "<!DOCTYPE html>")
	assert.Contains(t, string(data), "DoPlan Statistics")
	assert.Contains(t, string(data), "Testing Metrics")
}
