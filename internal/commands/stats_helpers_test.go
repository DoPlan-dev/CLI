package commands

import (
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/internal/statistics"
	"github.com/stretchr/testify/assert"
)

func TestParseTimeInput_Duration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		wantErr  bool
	}{
		{"7d", 7 * 24 * time.Hour, false},
		{"2w", 2 * 7 * 24 * time.Hour, false},
		{"1m", 30 * 24 * time.Hour, false},
		{"12h", 12 * time.Hour, false},
		{"invalid", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseTimeInput(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// Allow 1 hour difference for timing
				assert.WithinDuration(t, time.Now().Add(-tt.expected), result, time.Hour)
			}
		})
	}
}

func TestParseTimeInput_Date(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"2025-01-01", false},
		{"2025-01-01 15:04:05", false},
		{"2025-01-01T15:04:05Z", false},
		{"invalid-date", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseTimeInput(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.False(t, result.IsZero())
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		wantErr  bool
	}{
		{"7d", 7 * 24 * time.Hour, false},
		{"2w", 2 * 7 * 24 * time.Hour, false},
		{"1m", 30 * 24 * time.Hour, false},
		{"12h", 12 * time.Hour, false},
		{"5d", 5 * 24 * time.Hour, false},
		{"invalid", 0, true},
		{"x", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseDuration(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestParseRangeInput(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"2025-01-01:2025-01-15", false},
		{"7d:1d", false}, // 7 days ago to 1 day ago is valid (start before end)
		{"1d:7d", true},  // 1 day ago to 7 days ago should error (start after end)
		{"invalid", true},
		{"2025-01-01", true}, // missing colon
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			start, end, err := parseRangeInput(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.False(t, start.IsZero())
				assert.False(t, end.IsZero())
				assert.True(t, start.Before(end) || start.Equal(end))
			}
		})
	}
}

func TestFilterMetrics(t *testing.T) {
	metrics := &statistics.StatisticsMetrics{
		Velocity: &statistics.VelocityMetrics{
			FeaturesPerDay: 0.5,
		},
		Completion: &statistics.CompletionRates{
			Overall: 75,
		},
		Time: &statistics.TimeMetrics{
			DaysSinceStart: 30,
		},
		Quality: &statistics.QualityMetrics{
			PRMergeRate: 80.0,
		},
	}

	tests := []struct {
		filter  string
		hasVel  bool
		hasComp bool
		hasTime bool
		hasQual bool
	}{
		{"all", true, true, true, true},
		{"velocity", true, false, false, false},
		{"completion", false, true, false, false},
		{"time", false, false, true, false},
		{"quality", false, false, false, true},
		{"velocity,completion", true, true, false, false},
		{"velocity,time,quality", true, false, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.filter, func(t *testing.T) {
			filtered := filterMetrics(metrics, tt.filter)
			assert.NotNil(t, filtered)
			assert.Equal(t, tt.hasVel, filtered.Velocity != nil)
			assert.Equal(t, tt.hasComp, filtered.Completion != nil)
			assert.Equal(t, tt.hasTime, filtered.Time != nil)
			assert.Equal(t, tt.hasQual, filtered.Quality != nil)
		})
	}
}

func TestAggregateHistoricalMetrics(t *testing.T) {
	now := time.Now()
	history := []*statistics.HistoricalData{
		{
			Timestamp: now.Add(-2 * time.Hour),
			Metrics: &statistics.StatisticsMetrics{
				Velocity: &statistics.VelocityMetrics{FeaturesPerDay: 0.3},
			},
		},
		{
			Timestamp: now.Add(-1 * time.Hour),
			Metrics: &statistics.StatisticsMetrics{
				Velocity: &statistics.VelocityMetrics{FeaturesPerDay: 0.5},
			},
		},
	}

	result := aggregateHistoricalMetrics(history)
	assert.NotNil(t, result)
	assert.Equal(t, 0.5, result.Velocity.FeaturesPerDay) // Should use latest
}

func TestAggregateHistoricalMetrics_Empty(t *testing.T) {
	result := aggregateHistoricalMetrics([]*statistics.HistoricalData{})
	assert.Nil(t, result)
}
