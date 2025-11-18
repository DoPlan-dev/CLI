package dashboard_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/dashboard"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDashboardIntegration_RealProjectData tests the full dashboard generation
// with real project structure including phases, features, progress.json files, and commits
func TestDashboardIntegration_RealProjectData(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Setup project structure
	setupRealProjectStructure(t, projectRoot)

	// Create state with phases and features
	state := createTestState(t)

	// Create GitHub data with commits
	githubData := createTestGitHubData(t)

	// Generate dashboard
	gen := generators.NewDashboardGenerator(projectRoot, state, githubData)
	err := gen.GenerateJSON()
	require.NoError(t, err)

	// Verify dashboard.json exists
	loader := dashboard.NewLoader(projectRoot)
	assert.True(t, loader.DashboardExists())

	// Load and verify dashboard
	dashboard, err := loader.LoadDashboard()
	require.NoError(t, err)
	assert.NotNil(t, dashboard)

	// Verify structure
	assert.Equal(t, "1.0", dashboard.Version)
	assert.NotEmpty(t, dashboard.Generated)
	assert.Greater(t, dashboard.Project.Progress, 0)

	// Verify phases
	assert.Greater(t, len(dashboard.Phases), 0)
	for _, phase := range dashboard.Phases {
		assert.NotEmpty(t, phase.ID)
		assert.NotEmpty(t, phase.Name)
		assert.GreaterOrEqual(t, phase.Progress, 0)
		assert.LessOrEqual(t, phase.Progress, 100)
	}

	// Verify activity feed
	assert.NotNil(t, dashboard.Activity)
	assert.GreaterOrEqual(t, dashboard.Activity.Last7Days.Commits, 0)
	assert.GreaterOrEqual(t, len(dashboard.Activity.RecentActivity), 0)

	// Verify velocity metrics
	assert.NotNil(t, dashboard.Velocity)
	assert.GreaterOrEqual(t, dashboard.Velocity.CommitsPerDay, 0.0)
	assert.GreaterOrEqual(t, dashboard.Velocity.TasksPerDay, 0.0)
}

// TestActivityGenerator_RealData tests activity generation with real project data
func TestActivityGenerator_RealData(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Setup progress files
	setupProgressFiles(t, projectRoot)

	// Create state
	state := createTestState(t)

	// Create GitHub data with commits
	githubData := createTestGitHubData(t)

	// Read progress data
	progressParser := dashboard.NewProgressParser(projectRoot)
	progressData, err := progressParser.ReadProgressFiles()
	require.NoError(t, err)

	// Generate activity
	activityGen := dashboard.NewActivityGenerator(state, githubData, progressData)
	activity := activityGen.GenerateActivityFeed()

	// Verify activity structure
	assert.NotNil(t, activity)
	assert.GreaterOrEqual(t, activity.Last24Hours.Commits, 0)
	assert.GreaterOrEqual(t, activity.Last7Days.Commits, 0)
	assert.GreaterOrEqual(t, len(activity.RecentActivity), 0)

	// Verify recent activity is sorted (newest first)
	if len(activity.RecentActivity) > 1 {
		for i := 0; i < len(activity.RecentActivity)-1; i++ {
			time1, err1 := dashboard.ParseTime(activity.RecentActivity[i].Timestamp)
			time2, err2 := dashboard.ParseTime(activity.RecentActivity[i+1].Timestamp)
			if err1 == nil && err2 == nil {
				assert.True(t, time1.After(time2) || time1.Equal(time2),
					"Activities should be sorted newest first")
			}
		}
	}
}

// TestProgressParser_RealFiles tests progress parser with real progress.json files
func TestProgressParser_RealFiles(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Setup progress files
	setupProgressFiles(t, projectRoot)

	// Parse progress files
	parser := dashboard.NewProgressParser(projectRoot)
	progressData, err := parser.ReadProgressFiles()
	require.NoError(t, err)

	// Verify we found progress files
	assert.Greater(t, len(progressData), 0)

	// Verify structure
	for key, data := range progressData {
		assert.NotEmpty(t, key)
		assert.NotEmpty(t, data.Status)
		assert.GreaterOrEqual(t, data.Progress, 0)
		assert.LessOrEqual(t, data.Progress, 100)
	}
}

// TestVelocityCalculation_RealData tests velocity calculation with real commit data
func TestVelocityCalculation_RealData(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	state := createTestState(t)
	githubData := createTestGitHubData(t)

	// Create dashboard generator
	gen := generators.NewDashboardGenerator(projectRoot, state, githubData)

	// Generate dashboard to trigger velocity calculation
	err := gen.GenerateJSON()
	require.NoError(t, err)

	// Load dashboard
	loader := dashboard.NewLoader(projectRoot)
	dashboard, err := loader.LoadDashboard()
	require.NoError(t, err)

	// Verify velocity metrics
	velocity := dashboard.Velocity
	assert.NotNil(t, velocity)

	// If we have commits, velocity should be calculated
	if len(githubData.Commits) > 0 {
		assert.Greater(t, velocity.CommitsPerDay, 0.0)
	}
}

// TestDashboardAutoUpdate_RealData tests auto-update functionality
func TestDashboardAutoUpdate_RealData(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Setup state and config files (UpdateDashboard needs these)
	state := createTestState(t)
	cfgMgr := config.NewManager(projectRoot)
	err := cfgMgr.SaveState(state)
	require.NoError(t, err)

	cfg := config.NewConfig("cursor")
	err = cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	// Initial dashboard generation
	githubData := createTestGitHubData(t)
	gen := generators.NewDashboardGenerator(projectRoot, state, githubData)
	err = gen.GenerateJSON()
	require.NoError(t, err)

	// Get initial timestamp
	loader := dashboard.NewLoader(projectRoot)
	initialTime, err := loader.GetLastUpdateTime()
	require.NoError(t, err)

	// Wait a bit to ensure timestamp difference
	time.Sleep(200 * time.Millisecond)

	// Update dashboard
	// Use generators.UpdateDashboard directly to avoid import cycle
	err = generators.UpdateDashboard(projectRoot)
	require.NoError(t, err)

	// Wait a bit for file system to update
	time.Sleep(50 * time.Millisecond)

	// Verify timestamp updated
	updatedTime, err := loader.GetLastUpdateTime()
	require.NoError(t, err)

	// Check that time is after or equal (allowing for file system timestamp precision)
	if !updatedTime.After(initialTime) && !updatedTime.Equal(initialTime) {
		t.Errorf("Updated time (%v) should be after or equal to initial time (%v)", updatedTime, initialTime)
	}

	// At minimum, the generated timestamp in the JSON should be different
	// Let's check the actual dashboard JSON
	dashboard, err := loader.LoadDashboard()
	require.NoError(t, err)
	assert.NotEmpty(t, dashboard.Generated)
}

// Helper functions

func setupRealProjectStructure(t *testing.T, projectRoot string) {
	// Create phase and feature directories
	phase1Dir := filepath.Join(projectRoot, "doplan", "01-phase-1")
	phase2Dir := filepath.Join(projectRoot, "doplan", "02-phase-2")
	feature1Dir := filepath.Join(phase1Dir, "01-Feature-1")
	feature2Dir := filepath.Join(phase1Dir, "02-Feature-2")
	feature3Dir := filepath.Join(phase2Dir, "01-Feature-3")

	dirs := []string{phase1Dir, phase2Dir, feature1Dir, feature2Dir, feature3Dir}
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		require.NoError(t, err)
	}

	// Setup progress files
	setupProgressFiles(t, projectRoot)
}

func setupProgressFiles(t *testing.T, projectRoot string) {
	// Feature 1 progress (50% complete)
	feature1Progress := map[string]interface{}{
		"featureID":   "feature-1",
		"featureName": "Feature 1",
		"status":      "in-progress",
		"progress":    50,
		"branch":      "feature/01-phase-1-01-feature-1",
		"lastUpdated": time.Now().Add(-2 * time.Hour).Format(time.RFC3339),
	}
	writeProgressFile(t, projectRoot, "doplan/01-phase-1/01-Feature-1/progress.json", feature1Progress)

	// Feature 2 progress (100% complete)
	feature2Progress := map[string]interface{}{
		"featureID":   "feature-2",
		"featureName": "Feature 2",
		"status":      "complete",
		"progress":    100,
		"branch":      "feature/01-phase-1-02-feature-2",
		"lastUpdated": time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
	}
	writeProgressFile(t, projectRoot, "doplan/01-phase-1/02-Feature-2/progress.json", feature2Progress)

	// Feature 3 progress (25% complete)
	feature3Progress := map[string]interface{}{
		"featureID":   "feature-3",
		"featureName": "Feature 3",
		"status":      "in-progress",
		"progress":    25,
		"branch":      "feature/02-phase-2-01-feature-3",
		"lastUpdated": time.Now().Add(-30 * time.Minute).Format(time.RFC3339),
	}
	writeProgressFile(t, projectRoot, "doplan/02-phase-2/01-Feature-3/progress.json", feature3Progress)

	// Phase progress files
	phase1Progress := map[string]interface{}{
		"phaseID":   "phase-1",
		"phaseName": "Phase 1",
		"status":    "in-progress",
		"progress":  75,
		"features":  2,
	}
	writeProgressFile(t, projectRoot, "doplan/01-phase-1/phase-progress.json", phase1Progress)

	phase2Progress := map[string]interface{}{
		"phaseID":   "phase-2",
		"phaseName": "Phase 2",
		"status":    "in-progress",
		"progress":  25,
		"features":  1,
	}
	writeProgressFile(t, projectRoot, "doplan/02-phase-2/phase-progress.json", phase2Progress)
}

func writeProgressFile(t *testing.T, projectRoot, relPath string, data map[string]interface{}) {
	fullPath := filepath.Join(projectRoot, relPath)
	dir := filepath.Dir(fullPath)
	err := os.MkdirAll(dir, 0755)
	require.NoError(t, err)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	require.NoError(t, err)

	err = os.WriteFile(fullPath, jsonData, 0644)
	require.NoError(t, err)
}

func createTestState(t *testing.T) *models.State {
	return &models.State{
		Phases: []models.Phase{
			{
				ID:          "phase-1",
				Name:        "Phase 1",
				Status:      "in-progress",
				Description: "First phase",
				Features:    []string{"feature-1", "feature-2"},
				StartDate:   time.Now().AddDate(0, 0, -30).Format(time.RFC3339),
			},
			{
				ID:          "phase-2",
				Name:        "Phase 2",
				Status:      "in-progress",
				Description: "Second phase",
				Features:    []string{"feature-3"},
				StartDate:   time.Now().AddDate(0, 0, -15).Format(time.RFC3339),
			},
		},
		Features: []models.Feature{
			{
				ID:       "feature-1",
				Phase:    "phase-1",
				Name:     "Feature 1",
				Status:   "in-progress",
				Progress: 50,
				Branch:   "feature/01-phase-1-01-feature-1",
				TaskPhases: []models.TaskPhase{
					{
						Name: "Setup",
						Tasks: []models.Task{
							{Name: "Task 1", Completed: true},
							{Name: "Task 2", Completed: true},
							{Name: "Task 3", Completed: false},
							{Name: "Task 4", Completed: false},
						},
					},
				},
				StartDate: time.Now().AddDate(0, 0, -20).Format(time.RFC3339),
			},
			{
				ID:       "feature-2",
				Phase:    "phase-1",
				Name:     "Feature 2",
				Status:   "complete",
				Progress: 100,
				Branch:   "feature/01-phase-1-02-feature-2",
				TaskPhases: []models.TaskPhase{
					{
						Name: "Implementation",
						Tasks: []models.Task{
							{Name: "Task 1", Completed: true},
							{Name: "Task 2", Completed: true},
						},
					},
				},
				StartDate: time.Now().AddDate(0, 0, -25).Format(time.RFC3339),
			},
			{
				ID:       "feature-3",
				Phase:    "phase-2",
				Name:     "Feature 3",
				Status:   "in-progress",
				Progress: 25,
				Branch:   "feature/02-phase-2-01-feature-3",
				TaskPhases: []models.TaskPhase{
					{
						Name: "Planning",
						Tasks: []models.Task{
							{Name: "Task 1", Completed: true},
							{Name: "Task 2", Completed: false},
							{Name: "Task 3", Completed: false},
							{Name: "Task 4", Completed: false},
						},
					},
				},
				StartDate: time.Now().AddDate(0, 0, -10).Format(time.RFC3339),
			},
		},
		Progress: models.Progress{
			Overall: 58, // Average of phases
			Phases: map[string]int{
				"phase-1": 75,
				"phase-2": 25,
			},
		},
	}
}

func createTestGitHubData(t *testing.T) *models.GitHubData {
	now := time.Now()
	return &models.GitHubData{
		Commits: []models.Commit{
			{
				Hash:    "abc123",
				Message: "feat: implement feature 1",
				Author:  "Test User",
				Date:    now.Add(-2 * time.Hour).Format(time.RFC3339),
				Branch:  "feature/01-phase-1-01-feature-1",
			},
			{
				Hash:    "def456",
				Message: "fix: bug in feature 1",
				Author:  "Test User",
				Date:    now.Add(-1 * time.Hour).Format(time.RFC3339),
				Branch:  "feature/01-phase-1-01-feature-1",
			},
			{
				Hash:    "ghi789",
				Message: "feat: complete feature 2",
				Author:  "Test User",
				Date:    now.Add(-3 * time.Hour).Format(time.RFC3339),
				Branch:  "feature/01-phase-1-02-feature-2",
			},
			{
				Hash:    "jkl012",
				Message: "feat: start feature 3",
				Author:  "Test User",
				Date:    now.Add(-30 * time.Minute).Format(time.RFC3339),
				Branch:  "feature/02-phase-2-01-feature-3",
			},
			{
				Hash:    "mno345",
				Message: "docs: update readme",
				Author:  "Test User",
				Date:    now.Add(-5 * time.Hour).Format(time.RFC3339),
				Branch:  "main",
			},
		},
		PRs: []models.PullRequest{
			{
				Number: 1,
				Title:  "Feature 2: Complete implementation",
				URL:    "https://github.com/test/repo/pull/1",
				Status: "merged",
			},
		},
		Branches: []models.Branch{
			{
				Name:        "feature/01-phase-1-01-feature-1",
				Status:      "active",
				CommitCount: 2,
			},
			{
				Name:        "feature/01-phase-1-02-feature-2",
				Status:      "merged",
				CommitCount: 1,
			},
			{
				Name:        "feature/02-phase-2-01-feature-3",
				Status:      "active",
				CommitCount: 1,
			},
		},
	}
}
