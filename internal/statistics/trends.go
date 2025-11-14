package statistics

import (
	"math"
	"time"
)

// TrendCalculator calculates trends from historical data
type TrendCalculator struct{}

// NewTrendCalculator creates a new trend calculator
func NewTrendCalculator() *TrendCalculator {
	return &TrendCalculator{}
}

// CalculateTrends computes trends from historical data
func (tc *TrendCalculator) CalculateTrends(current *StatisticsMetrics, history []*HistoricalData) *Trends {
	if len(history) == 0 {
		return &Trends{
			VelocityTrend:   "stable",
			CompletionTrend: "stable",
			QualityTrend:    "stable",
		}
	}

	trends := &Trends{}

	// Calculate velocity trend
	if current.Velocity != nil && len(history) > 0 {
		trends.VelocityTrend, trends.VelocityChange = tc.calculateVelocityTrend(current.Velocity, history)
	}

	// Calculate completion trend
	if current.Completion != nil && len(history) > 0 {
		trends.CompletionTrend, trends.CompletionChange = tc.calculateCompletionTrend(current.Completion, history)
	}

	// Calculate quality trend
	if current.Quality != nil && len(history) > 0 {
		trends.QualityTrend = tc.calculateQualityTrend(current.Quality, history)
	}

	return trends
}

// calculateVelocityTrend determines if velocity is improving, declining, or stable
func (tc *TrendCalculator) calculateVelocityTrend(current *VelocityMetrics, history []*HistoricalData) (string, float64) {
	if len(history) < 2 {
		return "stable", 0.0
	}

	// Get previous velocity (from 7 days ago or latest if less than 7 days)
	var previous *VelocityMetrics
	cutoff := time.Now().Add(-7 * 24 * time.Hour)

	for i := len(history) - 1; i >= 0; i-- {
		if history[i].Timestamp.Before(cutoff) && history[i].Metrics.Velocity != nil {
			previous = history[i].Metrics.Velocity
			break
		}
	}

	// If no data from 7 days ago, use the oldest available
	if previous == nil && len(history) > 0 {
		for _, entry := range history {
			if entry.Metrics.Velocity != nil {
				previous = entry.Metrics.Velocity
				break
			}
		}
	}

	if previous == nil {
		return "stable", 0.0
	}

	// Calculate change in features per day
	currentRate := current.FeaturesPerDay
	previousRate := previous.FeaturesPerDay

	if previousRate == 0 {
		if currentRate > 0 {
			return "improving", 100.0
		}
		return "stable", 0.0
	}

	changePercent := ((currentRate - previousRate) / previousRate) * 100.0

	// Determine trend
	if changePercent > 10.0 {
		return "improving", changePercent
	} else if changePercent < -10.0 {
		return "declining", changePercent
	}
	return "stable", changePercent
}

// calculateCompletionTrend determines if completion is improving, declining, or stable
func (tc *TrendCalculator) calculateCompletionTrend(current *CompletionRates, history []*HistoricalData) (string, float64) {
	if len(history) < 2 {
		return "stable", 0.0
	}

	// Get previous completion (from 7 days ago or latest if less than 7 days)
	var previous *CompletionRates
	cutoff := time.Now().Add(-7 * 24 * time.Hour)

	for i := len(history) - 1; i >= 0; i-- {
		if history[i].Timestamp.Before(cutoff) && history[i].Metrics.Completion != nil {
			previous = history[i].Metrics.Completion
			break
		}
	}

	// If no data from 7 days ago, use the oldest available
	if previous == nil && len(history) > 0 {
		for _, entry := range history {
			if entry.Metrics.Completion != nil {
				previous = entry.Metrics.Completion
				break
			}
		}
	}

	if previous == nil {
		return "stable", 0.0
	}

	// Calculate change in overall completion
	currentRate := float64(current.Overall)
	previousRate := float64(previous.Overall)

	changePercent := currentRate - previousRate

	// Determine trend
	if changePercent > 5.0 {
		return "improving", changePercent
	} else if changePercent < -5.0 {
		return "declining", changePercent
	}
	return "stable", changePercent
}

// calculateQualityTrend determines if quality is improving, declining, or stable
func (tc *TrendCalculator) calculateQualityTrend(current *QualityMetrics, history []*HistoricalData) string {
	if len(history) < 2 {
		return "stable"
	}

	// Get previous quality (from 7 days ago or latest if less than 7 days)
	var previous *QualityMetrics
	cutoff := time.Now().Add(-7 * 24 * time.Hour)

	for i := len(history) - 1; i >= 0; i-- {
		if history[i].Timestamp.Before(cutoff) && history[i].Metrics.Quality != nil {
			previous = history[i].Metrics.Quality
			break
		}
	}

	// If no data from 7 days ago, use the oldest available
	if previous == nil && len(history) > 0 {
		for _, entry := range history {
			if entry.Metrics.Quality != nil {
				previous = entry.Metrics.Quality
				break
			}
		}
	}

	if previous == nil {
		return "stable"
	}

	// Compare PR merge rate
	currentRate := current.PRMergeRate
	previousRate := previous.PRMergeRate

	change := currentRate - previousRate

	// Determine trend
	if change > 5.0 {
		return "improving"
	} else if change < -5.0 {
		return "declining"
	}
	return "stable"
}

// CalculateAverageVelocity calculates average velocity over a time period
func (tc *TrendCalculator) CalculateAverageVelocity(history []*HistoricalData, days int) *VelocityMetrics {
	if len(history) == 0 {
		return &VelocityMetrics{}
	}

	cutoff := time.Now().AddDate(0, 0, -days)
	filtered := []*HistoricalData{}

	for _, entry := range history {
		if entry.Timestamp.After(cutoff) && entry.Metrics.Velocity != nil {
			filtered = append(filtered, entry)
		}
	}

	if len(filtered) == 0 {
		return &VelocityMetrics{}
	}

	// Calculate averages
	avg := &VelocityMetrics{}
	count := 0

	for _, entry := range filtered {
		if entry.Metrics.Velocity != nil {
			avg.FeaturesPerDay += entry.Metrics.Velocity.FeaturesPerDay
			avg.FeaturesPerWeek += entry.Metrics.Velocity.FeaturesPerWeek
			avg.CommitsPerDay += entry.Metrics.Velocity.CommitsPerDay
			avg.CommitsPerWeek += entry.Metrics.Velocity.CommitsPerWeek
			avg.TasksPerDay += entry.Metrics.Velocity.TasksPerDay
			avg.PRsPerWeek += entry.Metrics.Velocity.PRsPerWeek
			count++
		}
	}

	if count > 0 {
		avg.FeaturesPerDay /= float64(count)
		avg.FeaturesPerWeek /= float64(count)
		avg.CommitsPerDay /= float64(count)
		avg.CommitsPerWeek /= float64(count)
		avg.TasksPerDay /= float64(count)
		avg.PRsPerWeek /= float64(count)
	}

	return avg
}

// CalculateProjection projects future completion based on current velocity
func (tc *TrendCalculator) CalculateProjection(current *StatisticsMetrics, history []*HistoricalData) time.Time {
	if current.Velocity == nil || current.Completion == nil {
		return time.Time{}
	}

	// Use average velocity from last 7 days
	avgVelocity := tc.CalculateAverageVelocity(history, 7)
	if avgVelocity.FeaturesPerDay == 0 {
		avgVelocity = current.Velocity
	}

	// Calculate remaining work
	remainingPercent := 100.0 - float64(current.Completion.Overall)
	if remainingPercent <= 0 {
		return time.Now()
	}

	// Estimate days to completion
	// Assuming we need to complete remaining percentage
	// This is a simplified calculation
	daysRemaining := remainingPercent / (avgVelocity.FeaturesPerDay * 10.0) // Rough estimate

	if math.IsInf(daysRemaining, 0) || math.IsNaN(daysRemaining) || daysRemaining < 0 {
		return time.Time{}
	}

	return time.Now().AddDate(0, 0, int(daysRemaining))
}

