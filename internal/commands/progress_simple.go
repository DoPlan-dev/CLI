package commands

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// UpdateProgress updates all progress tracking files
func UpdateProgress() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	fmt.Println("📈 Updating progress...")
	fmt.Println()

	// Load state
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	// Load GitHub data (can be nil if GitHub not configured)
	var githubData *models.GitHubData
	githubSync := github.NewGitHubSync(projectRoot)
	if data, err := githubSync.LoadData(); err == nil {
		githubData = data
	}

	// Regenerate dashboard
	dashboardGenerator := generators.NewDashboardGenerator(projectRoot, state, githubData)
	if err := dashboardGenerator.Generate(); err != nil {
		return fmt.Errorf("failed to regenerate dashboard: %w", err)
	}

	fmt.Println("✅ Progress updated successfully!")
	fmt.Println("📊 Dashboard regenerated")
	fmt.Println("   • .doplan/dashboard.json")
	fmt.Println("   • .doplan/dashboard.md")

	return nil
}

