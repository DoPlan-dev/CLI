package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewDashboardCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Show project dashboard",
		Long:  "Display the project dashboard with progress and GitHub activity",
		RunE:  runDashboard,
	}

	return cmd
}

func runDashboard(cmd *cobra.Command, args []string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return doplanerror.NewIOError("IO001", "Failed to get current directory").WithCause(err)
	}

	// Initialize error handler
	errLogger := doplanerror.NewLogger(projectRoot, doplanerror.LogLevelInfo)
	errHandler := doplanerror.NewHandler(errLogger)

	// Check if installed
	if !config.IsInstalled(projectRoot) {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		errHandler.PrintError(doplanerror.ErrConfigNotFound(configPath))
		return nil
	}

	// Load state and GitHub data
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		statePath := filepath.Join(projectRoot, ".cursor", "config", "doplan-state.json")
		return errHandler.Handle(doplanerror.ErrStateNotFound(statePath).WithCause(err))
	}

	githubSync := github.NewGitHubSync(projectRoot)
	githubData, err := githubSync.LoadData()
	if err != nil {
		// Non-critical error, continue with empty data
		githubData = &github.GitHubData{}
	}

	// Generate dashboard
	dashboardGen := generators.NewDashboardGenerator(projectRoot, state, githubData)
	if err := dashboardGen.Generate(); err != nil {
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to generate dashboard").WithPath(filepath.Join(projectRoot, "doplan")).WithCause(err))
	}

	// Display markdown dashboard
	dashboardPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	content, err := os.ReadFile(dashboardPath)
	if err != nil {
		return errHandler.Handle(doplanerror.ErrFileNotFound(dashboardPath).WithCause(err))
	}

	fmt.Println(string(content))

	// Offer to open HTML dashboard
	color.Cyan("\n💡 Tip: Open 'doplan/dashboard.html' in your browser for a visual dashboard!")

	// Try to open HTML dashboard if on macOS/Linux
	htmlPath := fmt.Sprintf("%s/doplan/dashboard.html", projectRoot)
	if _, err := os.Stat(htmlPath); err == nil {
		color.Yellow("\nWould you like to open the HTML dashboard? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			exec.Command("open", htmlPath).Start() // macOS
		}
	}

	return nil
}
