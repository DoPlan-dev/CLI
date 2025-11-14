package statistics

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/fatih/color"
)

// Reporter formats and outputs statistics
type Reporter struct{}

// NewReporter creates a new statistics reporter
func NewReporter() *Reporter {
	return &Reporter{}
}

// ReportCLI outputs statistics in table format
func (r *Reporter) ReportCLI(metrics *StatisticsMetrics) error {
	animations := utils.AnimationsEnabled()

	fmt.Print(color.CyanString("\n📊 DoPlan Statistics\n"))
	fmt.Print(color.CyanString("===================\n\n"))

	r.printVelocityCLI(metrics.Velocity)
	r.printCompletionCLI(metrics.Completion, animations)
	r.printTimeCLI(metrics.Time)
	r.printQualityCLI(metrics.Quality)
	r.printTestingCLI(metrics.Testing, animations)
	r.printTrendsCLI(metrics.Trends)

	return nil
}

// ReportJSON outputs statistics in JSON format
func (r *Reporter) ReportJSON(metrics *StatisticsMetrics, path string) error {
	data, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal statistics: %w", err)
	}

	if path != "" {
		if err := os.WriteFile(path, data, 0644); err != nil {
			return fmt.Errorf("failed to write JSON file: %w", err)
		}
		color.Green("✅ Statistics exported to: %s\n", path)
	} else {
		fmt.Println(string(data))
	}

	return nil
}

// ReportMarkdown outputs statistics in Markdown format
func (r *Reporter) ReportMarkdown(metrics *StatisticsMetrics, path string) error {
	var sb strings.Builder

	sb.WriteString("# DoPlan Statistics\n\n")
	sb.WriteString(fmt.Sprintf("*Generated: %s*\n\n", time.Now().Format(time.RFC3339)))

	// Velocity Metrics
	if metrics.Velocity != nil {
		sb.WriteString("## Velocity Metrics\n\n")
		sb.WriteString("| Metric | Value |\n")
		sb.WriteString("|--------|-------|\n")
		sb.WriteString(fmt.Sprintf("| Features/day | %.2f |\n", metrics.Velocity.FeaturesPerDay))
		sb.WriteString(fmt.Sprintf("| Features/week | %.2f |\n", metrics.Velocity.FeaturesPerWeek))
		sb.WriteString(fmt.Sprintf("| Commits/day | %.2f |\n", metrics.Velocity.CommitsPerDay))
		sb.WriteString(fmt.Sprintf("| Commits/week | %.2f |\n", metrics.Velocity.CommitsPerWeek))
		sb.WriteString(fmt.Sprintf("| Tasks/day | %.2f |\n", metrics.Velocity.TasksPerDay))
		sb.WriteString(fmt.Sprintf("| PRs/week | %.2f |\n\n", metrics.Velocity.PRsPerWeek))
	}

	// Completion Rates
	if metrics.Completion != nil {
		sb.WriteString("## Completion Rates\n\n")
		sb.WriteString(fmt.Sprintf("- **Overall:** %s %.1f%%\n", renderMarkdownProgressBar(float64(metrics.Completion.Overall)), float64(metrics.Completion.Overall)))
		sb.WriteString(fmt.Sprintf("- **Tasks:** %s %.1f%%\n\n", renderMarkdownProgressBar(float64(metrics.Completion.Tasks)), float64(metrics.Completion.Tasks)))

		if len(metrics.Completion.Phases) > 0 {
			sb.WriteString("### Phases\n\n")
			sb.WriteString("| Phase | Progress |\n|-------|----------|\n")
			for _, phaseID := range sortedKeys(metrics.Completion.Phases) {
				val := float64(metrics.Completion.Phases[phaseID])
				sb.WriteString(fmt.Sprintf("| %s | %s %.1f%% |\n", phaseID, renderMarkdownProgressBar(val), val))
			}
			sb.WriteString("\n")
		}

		if len(metrics.Completion.Features) > 0 {
			sb.WriteString("### Key Features\n\n")
			sb.WriteString("| Feature | Progress |\n|---------|----------|\n")
			keys := sortedKeys(metrics.Completion.Features)
			maxFeatures := minInt(len(keys), 10)
			for _, featureID := range keys[:maxFeatures] {
				val := float64(metrics.Completion.Features[featureID])
				sb.WriteString(fmt.Sprintf("| %s | %s %.1f%% |\n", featureID, renderMarkdownProgressBar(val), val))
			}
			if len(keys) > maxFeatures {
				sb.WriteString(fmt.Sprintf("| ... | %d more features |\n", len(keys)-maxFeatures))
			}
			sb.WriteString("\n")
		}
	}

	// Time Metrics
	if metrics.Time != nil {
		sb.WriteString("## Time Metrics\n\n")
		sb.WriteString(fmt.Sprintf("- **Days since start:** %d\n", metrics.Time.DaysSinceStart))
		if metrics.Time.AvgFeatureTime > 0 {
			sb.WriteString(fmt.Sprintf("- **Avg feature time:** %.1f days\n", metrics.Time.AvgFeatureTime))
		}
		if !metrics.Time.EstimatedCompletion.IsZero() {
			sb.WriteString(fmt.Sprintf("- **Estimated completion:** %s\n", metrics.Time.EstimatedCompletion.Format("2006-01-02")))
		}
		sb.WriteString("\n")
	}

	// Quality Metrics
	if metrics.Quality != nil {
		sb.WriteString("## Quality Metrics\n\n")
		sb.WriteString(fmt.Sprintf("- **PR merge rate:** %.1f%%\n", metrics.Quality.PRMergeRate))
		sb.WriteString(fmt.Sprintf("- **Checkpoint frequency:** %.1f per week\n", metrics.Quality.CheckpointFrequency))
		sb.WriteString("\n")
	}

	// Testing Metrics
	if metrics.Testing != nil {
		sb.WriteString("## Testing Metrics\n\n")
		sb.WriteString(fmt.Sprintf("- **Overall Coverage:** %s %.1f%%\n\n", renderMarkdownProgressBar(metrics.Testing.OverallCoverage), metrics.Testing.OverallCoverage))

		if len(metrics.Testing.Packages) > 0 {
			sb.WriteString("| Package | Coverage |\n|---------|----------|\n")
			maxPackages := minInt(len(metrics.Testing.Packages), 10)
			for _, pkg := range metrics.Testing.Packages[:maxPackages] {
				sb.WriteString(fmt.Sprintf("| %s | %s %.1f%% |\n", pkg.Name, renderMarkdownProgressBar(pkg.Coverage), pkg.Coverage))
			}
			if len(metrics.Testing.Packages) > maxPackages {
				sb.WriteString(fmt.Sprintf("| ... | %d more packages |\n", len(metrics.Testing.Packages)-maxPackages))
			}
			sb.WriteString("\n")
		}
	}

	content := sb.String()

	if path != "" {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write Markdown file: %w", err)
		}
		color.Green("✅ Statistics exported to: %s\n", path)
	} else {
		fmt.Print(content)
	}

	return nil
}

// ReportHTML outputs statistics in HTML format
func (r *Reporter) ReportHTML(metrics *StatisticsMetrics, path string) error {
	var sb strings.Builder

	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html>\n<head>\n")
	sb.WriteString("<title>DoPlan Statistics</title>\n")
	sb.WriteString("<style>\n")
	sb.WriteString("body { font-family: Arial, sans-serif; margin: 20px; }\n")
	sb.WriteString("h1 { color: #333; }\n")
	sb.WriteString("h2 { color: #666; margin-top: 30px; }\n")
	sb.WriteString("table { border-collapse: collapse; width: 100%; margin: 20px 0; }\n")
	sb.WriteString("th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }\n")
	sb.WriteString("th { background-color: #f2f2f2; }\n")
	sb.WriteString(".progress-group { margin: 12px 0; }\n")
	sb.WriteString(".progress-label { font-weight: bold; display: block; margin-bottom: 4px; }\n")
	sb.WriteString(".progress-bar { background: #e5e7eb; border-radius: 9999px; overflow: hidden; height: 16px; }\n")
	sb.WriteString(".progress-bar span { display: block; height: 100%; background: linear-gradient(90deg,#10b981,#3b82f6); }\n")
	sb.WriteString("</style>\n")
	sb.WriteString("</head>\n<body>\n")
	sb.WriteString("<h1>📊 DoPlan Statistics</h1>\n")
	sb.WriteString(fmt.Sprintf("<p><em>Generated: %s</em></p>\n", time.Now().Format(time.RFC3339)))

	// Velocity Metrics
	if metrics.Velocity != nil {
		sb.WriteString("<h2>Velocity Metrics</h2>\n")
		sb.WriteString("<table>\n")
		sb.WriteString("<tr><th>Metric</th><th>Value</th></tr>\n")
		sb.WriteString(fmt.Sprintf("<tr><td>Features/day</td><td>%.2f</td></tr>\n", metrics.Velocity.FeaturesPerDay))
		sb.WriteString(fmt.Sprintf("<tr><td>Features/week</td><td>%.2f</td></tr>\n", metrics.Velocity.FeaturesPerWeek))
		sb.WriteString(fmt.Sprintf("<tr><td>Commits/day</td><td>%.2f</td></tr>\n", metrics.Velocity.CommitsPerDay))
		sb.WriteString(fmt.Sprintf("<tr><td>Commits/week</td><td>%.2f</td></tr>\n", metrics.Velocity.CommitsPerWeek))
		sb.WriteString(fmt.Sprintf("<tr><td>Tasks/day</td><td>%.2f</td></tr>\n", metrics.Velocity.TasksPerDay))
		sb.WriteString(fmt.Sprintf("<tr><td>PRs/week</td><td>%.2f</td></tr>\n", metrics.Velocity.PRsPerWeek))
		sb.WriteString("</table>\n")
	}

	// Completion Rates
	if metrics.Completion != nil {
		sb.WriteString("<h2>Completion Rates</h2>\n")
		sb.WriteString(fmt.Sprintf("<div class=\"progress-group\"><span class=\"progress-label\">Overall</span>%s<p>%.1f%%</p></div>\n", renderHTMLProgressBar(float64(metrics.Completion.Overall)), float64(metrics.Completion.Overall)))
		sb.WriteString(fmt.Sprintf("<div class=\"progress-group\"><span class=\"progress-label\">Tasks</span>%s<p>%.1f%%</p></div>\n", renderHTMLProgressBar(float64(metrics.Completion.Tasks)), float64(metrics.Completion.Tasks)))

		if len(metrics.Completion.Phases) > 0 {
			sb.WriteString("<h3>Phases</h3>\n")
			sb.WriteString("<table><tr><th>Phase</th><th>Progress</th></tr>\n")
			for _, phaseID := range sortedKeys(metrics.Completion.Phases) {
				val := float64(metrics.Completion.Phases[phaseID])
				sb.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s %.1f%%</td></tr>\n", phaseID, renderHTMLProgressBarInline(val), val))
			}
			sb.WriteString("</table>\n")
		}
	}

	// Time Metrics
	if metrics.Time != nil {
		sb.WriteString("<h2>Time Metrics</h2>\n")
		sb.WriteString(fmt.Sprintf("<p><strong>Days since start:</strong> %d</p>\n", metrics.Time.DaysSinceStart))
		if metrics.Time.AvgFeatureTime > 0 {
			sb.WriteString(fmt.Sprintf("<p><strong>Avg feature time:</strong> %.1f days</p>\n", metrics.Time.AvgFeatureTime))
		}
		if !metrics.Time.EstimatedCompletion.IsZero() {
			sb.WriteString(fmt.Sprintf("<p><strong>Estimated completion:</strong> %s</p>\n", metrics.Time.EstimatedCompletion.Format("2006-01-02")))
		}
	}

	// Quality Metrics
	if metrics.Quality != nil {
		sb.WriteString("<h2>Quality Metrics</h2>\n")
		sb.WriteString(fmt.Sprintf("<p><strong>PR merge rate:</strong> %.1f%%</p>\n", metrics.Quality.PRMergeRate))
		sb.WriteString(fmt.Sprintf("<p><strong>Checkpoint frequency:</strong> %.1f per week</p>\n", metrics.Quality.CheckpointFrequency))
	}

	// Testing Metrics
	if metrics.Testing != nil {
		sb.WriteString("<h2>Testing Metrics</h2>\n")
		sb.WriteString(fmt.Sprintf("<div class=\"progress-group\"><span class=\"progress-label\">Overall Coverage</span>%s<p>%.1f%%</p></div>\n", renderHTMLProgressBar(metrics.Testing.OverallCoverage), metrics.Testing.OverallCoverage))

		if len(metrics.Testing.Packages) > 0 {
			sb.WriteString("<table><tr><th>Package</th><th>Coverage</th></tr>\n")
			maxPackages := minInt(len(metrics.Testing.Packages), 10)
			for _, pkg := range metrics.Testing.Packages[:maxPackages] {
				sb.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s %.1f%%</td></tr>\n", pkg.Name, renderHTMLProgressBarInline(pkg.Coverage), pkg.Coverage))
			}
			if len(metrics.Testing.Packages) > maxPackages {
				sb.WriteString(fmt.Sprintf("<tr><td>...</td><td>%d more packages</td></tr>\n", len(metrics.Testing.Packages)-maxPackages))
			}
			sb.WriteString("</table>\n")
		}
	}

	sb.WriteString("</body>\n</html>\n")

	content := sb.String()

	if path != "" {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write HTML file: %w", err)
		}
		color.Green("✅ Statistics exported to: %s\n", path)
	} else {
		fmt.Print(content)
	}

	return nil
}

const cliProgressBarWidth = 24

func (r *Reporter) printVelocityCLI(velocity *VelocityMetrics) {
	if velocity == nil {
		return
	}
	fmt.Println(color.YellowString("Velocity Metrics:"))
	fmt.Printf("  Features/day:    %.2f\n", velocity.FeaturesPerDay)
	fmt.Printf("  Features/week:   %.2f\n", velocity.FeaturesPerWeek)
	fmt.Printf("  Commits/day:     %.2f\n", velocity.CommitsPerDay)
	fmt.Printf("  Commits/week:    %.2f\n", velocity.CommitsPerWeek)
	fmt.Printf("  Tasks/day:       %.2f\n", velocity.TasksPerDay)
	fmt.Printf("  PRs/week:        %.2f\n\n", velocity.PRsPerWeek)
}

func (r *Reporter) printCompletionCLI(completion *CompletionRates, animations bool) {
	if completion == nil {
		return
	}

	fmt.Println(color.YellowString("Completion Rates:"))
	r.printCLIProgressBar("Overall", float64(completion.Overall), 2, animations)
	r.printCLIProgressBar("Tasks", float64(completion.Tasks), 2, animations)

	if len(completion.Phases) > 0 {
		fmt.Println("  Phases:")
		for _, phaseID := range sortedKeys(completion.Phases) {
			r.printCLIProgressBar(phaseID, float64(completion.Phases[phaseID]), 4, animations)
		}
	}

	if len(completion.Features) > 0 {
		fmt.Println("  Key Features:")
		keys := sortedKeys(completion.Features)
		maxFeatures := minInt(len(keys), 5)
		for _, featureID := range keys[:maxFeatures] {
			r.printCLIProgressBar(featureID, float64(completion.Features[featureID]), 4, animations)
		}
		if len(keys) > maxFeatures {
			fmt.Printf("    ...and %d more features\n", len(keys)-maxFeatures)
		}
	}

	fmt.Println()
}

func (r *Reporter) printTimeCLI(timeMetrics *TimeMetrics) {
	if timeMetrics == nil {
		return
	}

	fmt.Println(color.YellowString("Time Metrics:"))
	fmt.Printf("  Days since start: %d\n", timeMetrics.DaysSinceStart)
	if timeMetrics.AvgFeatureTime > 0 {
		fmt.Printf("  Avg feature time: %.1f days\n", timeMetrics.AvgFeatureTime)
	}
	if timeMetrics.AvgPhaseTime > 0 {
		fmt.Printf("  Avg phase time:   %.1f days\n", timeMetrics.AvgPhaseTime)
	}
	if !timeMetrics.EstimatedCompletion.IsZero() {
		fmt.Printf("  Estimated completion: %s\n", timeMetrics.EstimatedCompletion.Format("2006-01-02"))
	}
	fmt.Println()
}

func (r *Reporter) printQualityCLI(quality *QualityMetrics) {
	if quality == nil {
		return
	}

	fmt.Println(color.YellowString("Quality Metrics:"))
	fmt.Printf("  PR merge rate:        %.1f%%\n", quality.PRMergeRate)
	fmt.Printf("  Checkpoint frequency: %.1f per week\n", quality.CheckpointFrequency)
	if quality.AvgBranchLifetime > 0 {
		fmt.Printf("  Avg branch lifetime:  %.1f days\n", quality.AvgBranchLifetime)
	}
	fmt.Println()
}

func (r *Reporter) printTestingCLI(testing *TestingMetrics, animations bool) {
	if testing == nil {
		return
	}

	fmt.Println(color.YellowString("Testing Metrics:"))
	r.printCLIProgressBar("Coverage", testing.OverallCoverage, 2, animations)

	if len(testing.Packages) > 0 {
		fmt.Println("  Packages:")
		maxPackages := minInt(len(testing.Packages), 5)
		for _, pkg := range testing.Packages[:maxPackages] {
			r.printCLIProgressBar(pkg.Name, pkg.Coverage, 4, animations)
		}
		if len(testing.Packages) > maxPackages {
			fmt.Printf("    ...and %d more packages\n", len(testing.Packages)-maxPackages)
		}
	}

	fmt.Println()
}

func (r *Reporter) printTrendsCLI(trends *Trends) {
	if trends == nil {
		return
	}

	fmt.Println(color.YellowString("Trends:"))
	fmt.Printf("  Velocity:   %s", trends.VelocityTrend)
	if trends.VelocityChange != 0 {
		fmt.Printf(" (%.1f%%)", trends.VelocityChange)
	}
	fmt.Println()
	fmt.Printf("  Completion: %s", trends.CompletionTrend)
	if trends.CompletionChange != 0 {
		fmt.Printf(" (%.1f%%)", trends.CompletionChange)
	}
	fmt.Println()
}

func (r *Reporter) printCLIProgressBar(label string, percent float64, indent int, animate bool) {
	indentStr := strings.Repeat(" ", indent)
	percent = clampPercent(percent)
	linePrefix := fmt.Sprintf("%s%s:", indentStr, label)

	if animate && percent > 0 {
		animateProgressBar(linePrefix, percent)
		return
	}

	fmt.Printf("%s %s %5.1f%%\n", linePrefix, renderProgressBar(cliProgressBarWidth, percent), percent)
}

func animateProgressBar(prefix string, target float64) {
	target = clampPercent(target)
	stepCount := int(math.Max(5, math.Round(target/5)))
	if stepCount < 5 {
		stepCount = 5
	}
	increment := target / float64(stepCount)
	current := 0.0

	for current <= target {
		fmt.Printf("\r%s %s %5.1f%%", prefix, renderProgressBar(cliProgressBarWidth, current), current)
		time.Sleep(40 * time.Millisecond)
		if current >= target {
			break
		}
		current += increment
		if current > target {
			current = target
		}
	}
	fmt.Printf("\r%s %s %5.1f%%\n", prefix, renderProgressBar(cliProgressBarWidth, target), target)
}

func renderProgressBar(width int, percent float64) string {
	percent = clampPercent(percent)
	filled := int(math.Round((percent / 100) * float64(width)))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	var sb strings.Builder
	sb.Grow(width + 2)
	sb.WriteRune('[')
	for i := 0; i < width; i++ {
		if i < filled {
			sb.WriteRune('█')
		} else {
			sb.WriteRune('░')
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

func renderMarkdownProgressBar(percent float64) string {
	return renderProgressBar(20, percent)
}

func renderHTMLProgressBar(percent float64) string {
	return fmt.Sprintf("<div class=\"progress-bar\"><span style=\"width: %.1f%%\"></span></div>", clampPercent(percent))
}

func renderHTMLProgressBarInline(percent float64) string {
	return fmt.Sprintf("<div class=\"progress-bar\"><span style=\"width: %.1f%%\"></span></div>", clampPercent(percent))
}

func clampPercent(value float64) float64 {
	switch {
	case value < 0:
		return 0
	case value > 100:
		return 100
	default:
		return value
	}
}

func sortedKeys(data map[string]int) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
