package dashboard

import (
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/charmbracelet/lipgloss"
)

// Sparkline characters for visualization
const (
	SparklineMin    = "▂"
	SparklineLow    = "▃"
	SparklineMid    = "▄"
	SparklineHigh   = "▅"
	SparklineHigher = "▆"
	SparklineMax    = "▇"
	SparklineFull   = "█"
)

// GenerateSparkline generates a sparkline visualization from data
func GenerateSparkline(data []float64, width int) string {
	if len(data) == 0 {
		return ""
	}

	if width <= 0 {
		width = len(data)
	}

	// Normalize data to fit within character range
	min, max := findMinMax(data)
	rangeVal := max - min
	if rangeVal == 0 {
		// All values are the same
		return strings.Repeat(SparklineMid, width)
	}

	// Generate sparkline
	var result string
	step := float64(len(data)) / float64(width)

	for i := 0; i < width; i++ {
		startIdx := int(float64(i) * step)
		endIdx := int(float64(i+1) * step)
		if endIdx > len(data) {
			endIdx = len(data)
		}
		if startIdx >= len(data) {
			startIdx = len(data) - 1
		}

		// Average values in this range
		var sum float64
		count := 0
		for j := startIdx; j < endIdx && j < len(data); j++ {
			sum += data[j]
			count++
		}

		var avg float64
		if count > 0 {
			avg = sum / float64(count)
		} else {
			avg = data[startIdx]
		}

		// Normalize to 0-7 range for character selection
		normalized := (avg - min) / rangeVal
		char := getSparklineChar(normalized)
		result += char
	}

	return result
}

// GenerateSparklineWithTrend generates a sparkline with trend indicator
func GenerateSparklineWithTrend(data []float64, width int) (string, string, string) {
	sparkline := GenerateSparkline(data, width)

	if len(data) < 2 {
		return sparkline, "→", "stable"
	}

	// Calculate trend
	firstHalf := average(data[:len(data)/2])
	secondHalf := average(data[len(data)/2:])

	change := ((secondHalf - firstHalf) / firstHalf) * 100

	var trendIcon string
	var trendText string

	if change > 5 {
		trendIcon = "↑"
		trendText = "increasing"
	} else if change < -5 {
		trendIcon = "↓"
		trendText = "decreasing"
	} else {
		trendIcon = "→"
		trendText = "stable"
	}

	return sparkline, trendIcon, trendText
}

// ColorCodeSparkline color codes a sparkline based on trend
func ColorCodeSparkline(sparkline, trendText string) string {
	style := lipgloss.NewStyle()

	switch trendText {
	case "increasing":
		style = style.Foreground(lipgloss.Color("#10b981")) // Green
	case "decreasing":
		style = style.Foreground(lipgloss.Color("#ef4444")) // Red
	default:
		style = style.Foreground(lipgloss.Color("#f59e0b")) // Amber
	}

	return style.Render(sparkline)
}

// getSparklineChar returns the appropriate sparkline character for normalized value (0-1)
func getSparklineChar(normalized float64) string {
	// Clamp to 0-1
	if normalized < 0 {
		normalized = 0
	}
	if normalized > 1 {
		normalized = 1
	}

	// Map to 7 levels
	level := int(normalized * 7)

	switch level {
	case 0:
		return SparklineMin
	case 1:
		return SparklineLow
	case 2:
		return SparklineMid
	case 3:
		return SparklineHigh
	case 4:
		return SparklineHigher
	case 5:
		return SparklineMax
	case 6, 7:
		return SparklineFull
	default:
		return SparklineMid
	}
}

// findMinMax finds min and max values in a slice
func findMinMax(data []float64) (float64, float64) {
	if len(data) == 0 {
		return 0, 0
	}

	min := data[0]
	max := data[0]

	for _, val := range data {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return min, max
}

// average calculates the average of a float slice
func average(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	var sum float64
	for _, val := range data {
		sum += val
	}

	return sum / float64(len(data))
}

// GenerateVelocityHistory generates velocity data for the last N days
func GenerateVelocityHistory(dailyData []DailyVelocity, days int) []float64 {
	if len(dailyData) == 0 {
		return make([]float64, days)
	}

	result := make([]float64, days)

	// Fill with zeros initially
	for i := range result {
		result[i] = 0
	}

	// Fill with actual data (most recent first)
	for i, data := range dailyData {
		if i >= days {
			break
		}
		daysAgo := days - 1 - i
		if daysAgo >= 0 && daysAgo < days {
			result[daysAgo] = data.Value
		}
	}

	return result
}

// DailyVelocity represents velocity for a single day
type DailyVelocity struct {
	Date  time.Time
	Value float64
}

// CalculateVelocityHistory calculates velocity history from commits and tasks
func CalculateVelocityHistory(commits []models.Commit, tasks []TaskProgress, days int) []DailyVelocity {
	history := make([]DailyVelocity, days)
	now := time.Now()

	// Initialize with zeros
	for i := 0; i < days; i++ {
		history[i] = DailyVelocity{
			Date:  now.AddDate(0, 0, -i),
			Value: 0,
		}
	}

	// Count commits per day
	commitsPerDay := make(map[string]int)
	for _, commit := range commits {
		commitTime, err := ParseTime(commit.Date)
		if err != nil {
			continue
		}
		dayKey := commitTime.Format("2006-01-02")
		commitsPerDay[dayKey]++
	}

	// Count tasks per day
	tasksPerDay := make(map[string]int)
	for _, task := range tasks {
		if task.Completed && !task.CompletedAt.IsZero() {
			dayKey := task.CompletedAt.Format("2006-01-02")
			tasksPerDay[dayKey]++
		}
	}

	// Fill history
	for i := 0; i < days; i++ {
		date := now.AddDate(0, 0, -i)
		dayKey := date.Format("2006-01-02")

		commits := float64(commitsPerDay[dayKey])
		tasks := float64(tasksPerDay[dayKey])

		// Combine commits and tasks (weighted)
		history[i].Value = commits + (tasks * 0.5)
		history[i].Date = date
	}

	return history
}
