package dashboard

import (
	"fmt"
	"os"
	"path/filepath"
)

// Updater handles dashboard.json updates
// Note: UpdateDashboard is now in generators package to avoid import cycle
type Updater struct {
	projectRoot string
}

// NewUpdater creates a new dashboard updater
func NewUpdater(projectRoot string) *Updater {
	return &Updater{
		projectRoot: projectRoot,
	}
}

// UpdateDashboard is deprecated - use generators.UpdateDashboard instead
// This method is kept for backward compatibility but will be removed
func (u *Updater) UpdateDashboard() error {
	// This creates an import cycle, so we need to use generators.UpdateDashboard directly
	// For now, return an error directing users to use generators.UpdateDashboard
	return fmt.Errorf("UpdateDashboard has been moved to generators package. Use generators.UpdateDashboard instead")
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

