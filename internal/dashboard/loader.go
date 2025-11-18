package dashboard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Loader loads dashboard data from dashboard.json
type Loader struct {
	projectRoot string
}

// NewLoader creates a new dashboard loader
func NewLoader(projectRoot string) *Loader {
	return &Loader{
		projectRoot: projectRoot,
	}
}

// LoadDashboard loads dashboard.json and returns the dashboard data
func (l *Loader) LoadDashboard() (*models.DashboardJSON, error) {
	dashboardPath := filepath.Join(l.projectRoot, ".doplan", "dashboard.json")

	// Check if dashboard.json exists
	if _, err := os.Stat(dashboardPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("dashboard.json not found")
	}

	// Read dashboard.json
	data, err := os.ReadFile(dashboardPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dashboard.json: %w", err)
	}

	// Parse JSON
	var dashboard models.DashboardJSON
	if err := json.Unmarshal(data, &dashboard); err != nil {
		return nil, fmt.Errorf("failed to parse dashboard.json: %w", err)
	}

	return &dashboard, nil
}

// DashboardExists checks if dashboard.json exists
func (l *Loader) DashboardExists() bool {
	dashboardPath := filepath.Join(l.projectRoot, ".doplan", "dashboard.json")
	_, err := os.Stat(dashboardPath)
	return err == nil
}

// GetLastUpdateTime returns the last update time from dashboard.json
func (l *Loader) GetLastUpdateTime() (time.Time, error) {
	dashboard, err := l.LoadDashboard()
	if err != nil {
		return time.Time{}, err
	}

	parsed, err := time.Parse(time.RFC3339, dashboard.Generated)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse generated time: %w", err)
	}

	return parsed, nil
}

