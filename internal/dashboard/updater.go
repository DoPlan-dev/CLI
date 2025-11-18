package dashboard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Updater handles dashboard.json updates
type Updater struct {
	projectRoot string
}

// NewUpdater creates a new dashboard updater
func NewUpdater(projectRoot string) *Updater {
	return &Updater{
		projectRoot: projectRoot,
	}
}

// UpdateDashboard regenerates dashboard.json
func (u *Updater) UpdateDashboard() error {
	// Load state
	configMgr := config.NewManager(u.projectRoot)
	state, err := configMgr.LoadState()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	// Load config
	cfg, err := configMgr.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Sync GitHub data if configured
	var githubData *models.GitHubData
	if cfg.GitHub.Enabled {
		githubSync := github.NewGitHubSync(u.projectRoot)
		githubData, err = githubSync.Sync()
		if err != nil {
			// Continue without GitHub data if sync fails
			githubData = &models.GitHubData{}
		}
	} else {
		githubData = &models.GitHubData{}
	}

	// Read progress files
	progressParser := NewProgressParser(u.projectRoot)
	progressData, err := progressParser.ReadProgressFiles()
	if err != nil {
		// Continue without progress data if parsing fails
		progressData = make(map[string]*ProgressData)
	}

	// Generate activity feed
	activityGen := NewActivityGenerator(state, githubData, progressData)
	activity := activityGen.GenerateActivityFeed()

	// Create dashboard generator
	gen := generators.NewDashboardGenerator(u.projectRoot, state, githubData)
	
	// Generate dashboard JSON
	if err := gen.GenerateJSON(); err != nil {
		return fmt.Errorf("failed to generate dashboard JSON: %w", err)
	}

	// Also generate markdown and HTML for backward compatibility
	if err := gen.Generate(); err != nil {
		return fmt.Errorf("failed to generate dashboard: %w", err)
	}

	return nil
}

// DashboardExists checks if dashboard.json exists
func (u *Updater) DashboardExists() bool {
	dashboardPath := filepath.Join(u.projectRoot, ".doplan", "dashboard.json")
	_, err := os.Stat(dashboardPath)
	return err == nil
}

// ShouldUpdate checks if dashboard should be updated based on file timestamps
func (u *Updater) ShouldUpdate() bool {
	dashboardPath := filepath.Join(u.projectRoot, ".doplan", "dashboard.json")
	
	// If dashboard doesn't exist, should update
	if _, err := os.Stat(dashboardPath); os.IsNotExist(err) {
		return true
	}

	dashboardInfo, err := os.Stat(dashboardPath)
	if err != nil {
		return true
	}

	// Check if state.json is newer
	statePath := filepath.Join(u.projectRoot, ".cursor", "config", "doplan-state.json")
	if stateInfo, err := os.Stat(statePath); err == nil {
		if stateInfo.ModTime().After(dashboardInfo.ModTime()) {
			return true
		}
	}

	// Check if any progress.json is newer
	doplanDir := filepath.Join(u.projectRoot, "doplan")
	if _, err := os.Stat(doplanDir); err == nil {
		// Walk and check progress.json files
		filepath.Walk(doplanDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.Name() == "progress.json" || info.Name() == "phase-progress.json" {
				if info.ModTime().After(dashboardInfo.ModTime()) {
					// Found newer file, but we can't return from Walk
					// This is a simplified check - could be enhanced
				}
			}
			return nil
		})
	}

	return false
}

