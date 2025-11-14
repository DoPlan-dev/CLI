package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/checkpoint"
	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

func NewProgressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "progress",
		Short: "Update progress tracking",
		Long:  "Update all progress tracking files and regenerate dashboard",
		RunE:  runProgress,
	}

	return cmd
}

func runProgress(cmd *cobra.Command, args []string) error {
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

	color.Blue("Updating progress tracking...\n")

	// Load state
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		statePath := filepath.Join(projectRoot, ".cursor", "config", "doplan-state.json")
		return errHandler.Handle(doplanerror.ErrStateNotFound(statePath).WithCause(err))
	}

	// Scan feature directories and update progress
	doplanDir := filepath.Join(projectRoot, "doplan")
	if err := updateProgressFromTasks(doplanDir, state); err != nil {
		color.Yellow("Warning: Failed to update progress from tasks: %v", err)
	}

	// Save updated state
	if err := cfgMgr.SaveState(state); err != nil {
		return fmt.Errorf("failed to save state: %w", err)
	}

	// Sync GitHub data
	githubSync := github.NewGitHubSync(projectRoot)
	githubData, err := githubSync.LoadData()
	if err != nil {
		githubData = &github.GitHubData{}
	}

	// Regenerate dashboard
	dashboardGen := generators.NewDashboardGenerator(projectRoot, state, githubData)
	if err := dashboardGen.Generate(); err != nil {
		return fmt.Errorf("failed to regenerate dashboard: %w", err)
	}

	// Auto-create checkpoints for completed features/phases
	checkpointMgr := checkpoint.NewCheckpointManager(projectRoot)
	for i := range state.Features {
		feature := &state.Features[i]
		if feature.Progress == 100 && feature.Status == "complete" {
			// Check if checkpoint already exists
			if feature.CheckpointID == "" {
				if err := checkpointMgr.AutoCreateFeatureCheckpoint(feature); err != nil {
					color.Yellow("⚠️  Failed to create checkpoint for feature '%s': %v\n", feature.Name, err)
				}
			}
		}
	}

	// Check for auto-PR creation
	autoPRMgr := github.NewAutoPRManager(projectRoot)
	if err := autoPRMgr.WatchFeatures(state); err != nil {
		color.Yellow("⚠️  Failed to check for auto-PR creation: %v\n", err)
	}

	color.Green("✅ Progress updated and dashboard regenerated!\n")
	color.Cyan("Run 'doplan dashboard' to view the updated dashboard.")

	return nil
}

func updateProgressFromTasks(doplanDir string, state *models.State) error {
	// Walk through phase directories
	return filepath.Walk(doplanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if this is a tasks.md file
		if info.Name() == "tasks.md" {
			// Read tasks.md and count completed tasks
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Count completed tasks (lines with [x])
			lines := strings.Split(string(content), "\n")
			totalTasks := 0
			completedTasks := 0

			for _, line := range lines {
				if strings.Contains(line, "- [") {
					totalTasks++
					if strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]") {
						completedTasks++
					}
				}
			}

			// Calculate progress percentage
			progress := 0
			if totalTasks > 0 {
				progress = (completedTasks * 100) / totalTasks
			}

			// Find the feature and update its progress
			featureDir := filepath.Dir(path)
			featureName := filepath.Base(featureDir)
			
			// Try to match feature by directory name
			for i := range state.Features {
				if strings.Contains(strings.ToLower(state.Features[i].Name), strings.ToLower(featureName)) ||
					strings.Contains(strings.ToLower(featureName), strings.ToLower(state.Features[i].Name)) {
					state.Features[i].Progress = progress
					if progress == 100 {
						state.Features[i].Status = "complete"
					} else if progress > 0 {
						state.Features[i].Status = "in-progress"
					}
					break
				}
			}
		}

		return nil
	})
}
