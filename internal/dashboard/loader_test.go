package dashboard

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLoader(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	
	loader := NewLoader(projectRoot)
	assert.NotNil(t, loader)
	assert.Equal(t, projectRoot, loader.projectRoot)
}

func TestLoader_DashboardExists(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	loader := NewLoader(projectRoot)
	
	// Should not exist initially
	assert.False(t, loader.DashboardExists())
	
	// Create dashboard.json
	dashboardDir := filepath.Join(projectRoot, ".doplan")
	err := os.MkdirAll(dashboardDir, 0755)
	require.NoError(t, err)
	
	dashboardPath := filepath.Join(dashboardDir, "dashboard.json")
	dashboard := &generators.DashboardJSON{
		Version:   "1.0",
		Generated: time.Now().Format(time.RFC3339),
	}
	
	data, err := json.Marshal(dashboard)
	require.NoError(t, err)
	
	err = os.WriteFile(dashboardPath, data, 0644)
	require.NoError(t, err)
	
	// Should exist now
	assert.True(t, loader.DashboardExists())
}

func TestLoader_LoadDashboard(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	loader := NewLoader(projectRoot)
	
	// Test loading non-existent dashboard
	_, err := loader.LoadDashboard()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dashboard.json not found")
	
	// Create dashboard.json
	dashboardDir := filepath.Join(projectRoot, ".doplan")
	err = os.MkdirAll(dashboardDir, 0755)
	require.NoError(t, err)
	
	dashboardPath := filepath.Join(dashboardDir, "dashboard.json")
	now := time.Now()
	dashboard := &generators.DashboardJSON{
		Version:   "1.0",
		Generated: now.Format(time.RFC3339),
		Project: generators.ProjectJSON{
			Name:     "Test Project",
			Progress: 50,
			Status:   "in-progress",
		},
		Phases: []generators.PhaseJSON{
			{
				ID:       "01-phase-1",
				Name:     "Phase 1",
				Status:   "in-progress",
				Progress: 60,
			},
		},
		Summary: generators.SummaryJSON{
			TotalPhases:   1,
			InProgress:    1,
			TotalFeatures: 2,
		},
	}
	
	data, err := json.Marshal(dashboard)
	require.NoError(t, err)
	
	err = os.WriteFile(dashboardPath, data, 0644)
	require.NoError(t, err)
	
	// Load dashboard
	loaded, err := loader.LoadDashboard()
	require.NoError(t, err)
	assert.NotNil(t, loaded)
	assert.Equal(t, "1.0", loaded.Version)
	assert.Equal(t, "Test Project", loaded.Project.Name)
	assert.Equal(t, 50, loaded.Project.Progress)
	assert.Equal(t, 1, len(loaded.Phases))
	assert.Equal(t, "Phase 1", loaded.Phases[0].Name)
}

func TestLoader_GetLastUpdateTime(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	loader := NewLoader(projectRoot)
	
	// Test with non-existent dashboard
	_, err := loader.GetLastUpdateTime()
	assert.Error(t, err)
	
	// Create dashboard.json with timestamp
	dashboardDir := filepath.Join(projectRoot, ".doplan")
	err = os.MkdirAll(dashboardDir, 0755)
	require.NoError(t, err)
	
	dashboardPath := filepath.Join(dashboardDir, "dashboard.json")
	now := time.Now().Truncate(time.Second)
	dashboard := &generators.DashboardJSON{
		Version:   "1.0",
		Generated: now.Format(time.RFC3339),
	}
	
	data, err := json.Marshal(dashboard)
	require.NoError(t, err)
	
	err = os.WriteFile(dashboardPath, data, 0644)
	require.NoError(t, err)
	
	// Get last update time
	updateTime, err := loader.GetLastUpdateTime()
	require.NoError(t, err)
	assert.WithinDuration(t, now, updateTime, time.Second)
}

func TestLoader_LoadDashboard_InvalidJSON(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	loader := NewLoader(projectRoot)
	
	// Create invalid JSON file
	dashboardDir := filepath.Join(projectRoot, ".doplan")
	err := os.MkdirAll(dashboardDir, 0755)
	require.NoError(t, err)
	
	dashboardPath := filepath.Join(dashboardDir, "dashboard.json")
	err = os.WriteFile(dashboardPath, []byte("invalid json"), 0644)
	require.NoError(t, err)
	
	// Try to load
	_, err = loader.LoadDashboard()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse dashboard.json")
}

func TestLoader_LoadDashboard_CompleteStructure(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	loader := NewLoader(projectRoot)
	
	// Create complete dashboard.json
	dashboardDir := filepath.Join(projectRoot, ".doplan")
	err := os.MkdirAll(dashboardDir, 0755)
	require.NoError(t, err)
	
	dashboardPath := filepath.Join(dashboardDir, "dashboard.json")
	dashboard := &generators.DashboardJSON{
		Version:   "1.0",
		Generated: time.Now().Format(time.RFC3339),
		Project: generators.ProjectJSON{
			Name:        "Complete Project",
			Description: "A complete test project",
			Version:     "1.0.0",
			Progress:    75,
			Status:      "in-progress",
		},
		GitHub: generators.GitHubJSON{
			Repository: "user/repo",
			Branch:     "main",
			Commits:    10,
		},
		Phases: []generators.PhaseJSON{
			{
				ID:          "01-phase-1",
				Name:        "Phase 1",
				Description: "First phase",
				Status:      "complete",
				Progress:    100,
				Features: []generators.FeatureJSON{
					{
						ID:       "01-feature-1",
						Name:     "Feature 1",
						Status:   "complete",
						Progress: 100,
					},
				},
				Stats: generators.PhaseStatsJSON{
					TotalFeatures:  1,
					Completed:      1,
					TotalTasks:     5,
					CompletedTasks: 5,
				},
			},
		},
		Summary: generators.SummaryJSON{
			TotalPhases:    1,
			Completed:      1,
			TotalFeatures:  1,
			TotalTasks:     5,
			CompletedTasks: 5,
		},
		Activity: generators.ActivityJSON{
			Last24Hours: generators.ActivityPeriodJSON{
				Commits:        2,
				TasksCompleted: 1,
			},
			Last7Days: generators.ActivityPeriodJSON{
				Commits:        5,
				TasksCompleted: 3,
			},
		},
		APIKeys: generators.APIKeysJSON{
			Total:      3,
			Configured: 2,
			Pending:    1,
		},
		Velocity: generators.VelocityJSON{
			TasksPerDay:   1.5,
			CommitsPerDay: 2.0,
		},
	}
	
	data, err := json.Marshal(dashboard)
	require.NoError(t, err)
	
	err = os.WriteFile(dashboardPath, data, 0644)
	require.NoError(t, err)
	
	// Load and verify
	loaded, err := loader.LoadDashboard()
	require.NoError(t, err)
	assert.Equal(t, "Complete Project", loaded.Project.Name)
	assert.Equal(t, 75, loaded.Project.Progress)
	assert.Equal(t, 1, len(loaded.Phases))
	assert.Equal(t, 1, len(loaded.Phases[0].Features))
	assert.Equal(t, "user/repo", loaded.GitHub.Repository)
	assert.Equal(t, 3, loaded.APIKeys.Total)
	assert.Equal(t, 2, loaded.APIKeys.Configured)
}

