package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/DoPlan-dev/CLI/internal/statistics"
)

func NewStatsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Show project statistics and insights",
		Long:  "Display comprehensive project statistics including velocity, completion rates, time metrics, and quality indicators",
		RunE:  runStats,
	}

	cmd.Flags().StringP("format", "f", "table", "Output format: table, json, html, markdown")
	cmd.Flags().String("export", "", "Export to file path")
	cmd.Flags().String("since", "", "Show stats since date/duration (e.g., '7d', '2025-01-01')")
	cmd.Flags().String("range", "", "Show stats for date range (e.g., '2025-01-01:2025-01-15')")
	cmd.Flags().String("metrics", "all", "Show specific metrics: velocity, completion, time, quality, all")
	cmd.Flags().Bool("trends", false, "Include trend analysis")

	return cmd
}

func runStats(cmd *cobra.Command, args []string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Check if installed
	if !config.IsInstalled(projectRoot) {
		errHandler := doplanerror.NewHandler(nil)
		errHandler.PrintError(doplanerror.ErrConfigNotFound(""))
		return nil
	}

	// Initialize error handler
	errLogger := doplanerror.NewLogger(projectRoot, doplanerror.LogLevelInfo)
	errHandler := doplanerror.NewHandler(errLogger)

	// Collect statistics
	color.Blue("Collecting statistics...\n")

	collector := statistics.NewCollector(projectRoot)
	data, err := collector.Collect()
	if err != nil {
		return errHandler.Handle(err)
	}

	// Load state and GitHub data for calculations
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		return errHandler.Handle(err)
	}

	githubSync := github.NewGitHubSync(projectRoot)
	githubData, err := githubSync.LoadData()
	if err != nil {
		githubData = &github.GitHubData{}
	}

	// Get project start date from config
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		return errHandler.Handle(err)
	}

	projectStartDate := cfg.InstalledAt
	if projectStartDate.IsZero() {
		projectStartDate = time.Now()
	}

	// Handle time range filtering
	sinceFlag, _ := cmd.Flags().GetString("since")
	rangeFlag, _ := cmd.Flags().GetString("range")
	metricsFlag, _ := cmd.Flags().GetString("metrics")

	var historicalData []*statistics.HistoricalData
	var metrics *statistics.StatisticsMetrics

	// If time range is specified, load from historical data
	if sinceFlag != "" || rangeFlag != "" {
		storage := statistics.NewStorage(projectRoot)

		if sinceFlag != "" {
			sinceTime, err := parseTimeInput(sinceFlag)
			if err != nil {
				return errHandler.Handle(doplanerror.NewValidationError("VAL002", "Invalid 'since' time format").WithDetails(err.Error()))
			}
			historicalData, err = storage.LoadSince(sinceTime)
			if err != nil {
				return errHandler.Handle(doplanerror.NewIOError("IO005", "Failed to load historical data").WithCause(err))
			}
		}

		if rangeFlag != "" {
			start, end, err := parseRangeInput(rangeFlag)
			if err != nil {
				return errHandler.Handle(doplanerror.NewValidationError("VAL003", "Invalid 'range' format").WithDetails(err.Error()))
			}
			historicalData, err = storage.LoadRange(start, end)
			if err != nil {
				return errHandler.Handle(doplanerror.NewIOError("IO005", "Failed to load historical data").WithCause(err))
			}
		}

		// Calculate metrics from historical data if available
		if len(historicalData) > 0 {
			metrics = aggregateHistoricalMetrics(historicalData)
		}
	}

	// If no historical data or no time range specified, calculate from current data
	if metrics == nil {
		calculator := statistics.NewCalculator(projectStartDate)
		metrics = calculator.Calculate(data, state, githubData)
	}

	// Calculate trends if requested
	showTrends, _ := cmd.Flags().GetBool("trends")
	if showTrends {
		storage := statistics.NewStorage(projectRoot)
		history, err := storage.LoadAll()
		if err == nil && len(history) > 0 {
			trendCalculator := statistics.NewTrendCalculator()
			metrics.Trends = trendCalculator.CalculateTrends(metrics, history)
		}
	}

	// Filter metrics if requested
	if metricsFlag != "all" {
		metrics = filterMetrics(metrics, metricsFlag)
	}

	// Save to storage for historical tracking (only if not using historical data)
	if sinceFlag == "" && rangeFlag == "" {
		storage := statistics.NewStorage(projectRoot)
		if err := storage.Save(metrics, data); err != nil {
			// Log but don't fail the command
			errHandler.PrintError(err)
		}
	}

	// Get format and export options
	format, _ := cmd.Flags().GetString("format")
	exportPath, _ := cmd.Flags().GetString("export")

	// Report statistics
	reporter := statistics.NewReporter()

	switch format {
	case "json":
		return reporter.ReportJSON(metrics, exportPath)
	case "html":
		return reporter.ReportHTML(metrics, exportPath)
	case "markdown":
		return reporter.ReportMarkdown(metrics, exportPath)
	case "table":
		fallthrough
	default:
		if err := reporter.ReportCLI(metrics); err != nil {
			return errHandler.Handle(err)
		}
		if exportPath != "" {
			// Also export to file if specified
			return reporter.ReportMarkdown(metrics, exportPath)
		}
		return nil
	}
}

