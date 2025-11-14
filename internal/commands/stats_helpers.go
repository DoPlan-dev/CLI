package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/statistics"
)

// parseTimeInput parses time input in various formats
// Supports: "7d", "2w", "1m", "2025-01-01", "2025-01-01 15:04:05", RFC3339
func parseTimeInput(input string) (time.Time, error) {
	// Try duration first (e.g., "7d", "2w", "1m")
	if d, err := parseDuration(input); err == nil {
		return time.Now().Add(-d), nil
	}

	// Try date formats
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, input); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid time format: %s (supported: 7d, 2w, 1m, 2025-01-01, RFC3339)", input)
}

// parseDuration parses duration strings like "7d", "2w", "1m"
func parseDuration(input string) (time.Duration, error) {
	input = strings.TrimSpace(input)
	if len(input) < 2 {
		return 0, fmt.Errorf("invalid duration format")
	}

	unit := input[len(input)-1:]
	valueStr := input[:len(input)-1]

	var value int
	if _, err := fmt.Sscanf(valueStr, "%d", &value); err != nil {
		return 0, err
	}

	switch unit {
	case "d":
		return time.Duration(value) * 24 * time.Hour, nil
	case "w":
		return time.Duration(value) * 7 * 24 * time.Hour, nil
	case "m":
		return time.Duration(value) * 30 * 24 * time.Hour, nil
	case "h":
		return time.Duration(value) * time.Hour, nil
	default:
		return 0, fmt.Errorf("unsupported duration unit: %s (use: d, w, m, h)", unit)
	}
}

// parseRangeInput parses a date range in format "start:end"
func parseRangeInput(input string) (time.Time, time.Time, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, fmt.Errorf("range must be in format 'start:end'")
	}

	start, err := parseTimeInput(strings.TrimSpace(parts[0]))
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start time: %w", err)
	}

	end, err := parseTimeInput(strings.TrimSpace(parts[1]))
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end time: %w", err)
	}

	if start.After(end) {
		return time.Time{}, time.Time{}, fmt.Errorf("start time must be before end time")
	}

	return start, end, nil
}

// filterMetrics filters metrics based on the filter string
// filter can be: "velocity", "completion", "time", "quality", or comma-separated
func filterMetrics(metrics *statistics.StatisticsMetrics, filter string) *statistics.StatisticsMetrics {
	if filter == "all" {
		return metrics
	}

	filtered := &statistics.StatisticsMetrics{
		CalculatedAt: metrics.CalculatedAt,
	}

	filters := strings.Split(filter, ",")
	for _, f := range filters {
		f = strings.TrimSpace(strings.ToLower(f))
		switch f {
		case "velocity":
			filtered.Velocity = metrics.Velocity
		case "completion":
			filtered.Completion = metrics.Completion
		case "time":
			filtered.Time = metrics.Time
		case "quality":
			filtered.Quality = metrics.Quality
		}
	}

	return filtered
}

// aggregateHistoricalMetrics aggregates metrics from historical data
func aggregateHistoricalMetrics(history []*statistics.HistoricalData) *statistics.StatisticsMetrics {
	if len(history) == 0 {
		return nil
	}

	// Use the most recent entry
	latest := history[len(history)-1]
	return latest.Metrics
}
