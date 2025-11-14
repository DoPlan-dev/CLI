package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewGitHubCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github",
		Short: "Update GitHub data (branches, commits, pushes)",
		Long:  "Sync GitHub information and update dashboard",
		RunE:  runGitHub,
	}

	return cmd
}

func runGitHub(cmd *cobra.Command, args []string) error {
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

	color.Blue("Syncing GitHub data...\n")

	// Sync GitHub data
	githubSync := github.NewGitHubSync(projectRoot)
	data, err := githubSync.Sync()
	if err != nil {
		return errHandler.Handle(doplanerror.NewGitHubError("GH002", "Failed to sync GitHub data").WithCause(err))
	}

	color.Green("✅ GitHub data synced successfully!\n\n")

	// Display summary
	fmt.Printf("Branches: %d\n", len(data.Branches))
	fmt.Printf("Commits: %d\n", len(data.Commits))
	fmt.Printf("Pull Requests: %d\n", len(data.PRs))

	// Check for auto-PR creation
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err == nil {
		autoPRMgr := github.NewAutoPRManager(projectRoot)
		if err := autoPRMgr.WatchFeatures(state); err != nil {
			errHandler.PrintError(doplanerror.NewGitHubError("GH003", "Failed to check for auto-PR creation").WithCause(err))
		}
	}

	color.Cyan("\nRun 'doplan dashboard' to see the updated dashboard.")

	return nil
}
