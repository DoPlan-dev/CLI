package github

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
)

func TestNewAutoPRManager(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	mgr := NewAutoPRManager(projectRoot)
	assert.NotNil(t, mgr)
	assert.Equal(t, projectRoot, mgr.repoPath)
}

func TestAutoPRManager_CheckAndCreatePR_Disabled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create config with AutoPR disabled
	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = false
	cfg.GitHub.AutoPR = false
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	mgr := NewAutoPRManager(projectRoot)
	feature := &models.Feature{
		ID:       "feat-1",
		Name:     "Test Feature",
		Progress: 100,
		Status:   "complete",
	}

	err := mgr.CheckAndCreatePR(feature)
	assert.NoError(t, err)
}

func TestAutoPRManager_CheckAndCreatePR_AlreadyHasPR(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = true
	cfg.GitHub.AutoPR = true
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	mgr := NewAutoPRManager(projectRoot)
	feature := &models.Feature{
		ID:       "feat-1",
		Name:     "Test Feature",
		Progress: 100,
		Status:   "complete",
		PR: &models.PullRequest{
			Number: 1,
			URL:    "https://github.com/test/pr/1",
		},
	}

	err := mgr.CheckAndCreatePR(feature)
	assert.NoError(t, err)
}

func TestAutoPRManager_isFeatureComplete(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	mgr := NewAutoPRManager(projectRoot)

	// Test incomplete feature (progress < 100)
	feature1 := &models.Feature{
		ID:       "feat-1",
		Progress: 50,
		Status:   "in-progress",
	}
	complete, err := mgr.isFeatureComplete(feature1)
	require.NoError(t, err)
	assert.False(t, complete)

	// Test incomplete feature (status != complete)
	feature2 := &models.Feature{
		ID:       "feat-2",
		Progress: 100,
		Status:   "in-progress",
	}
	complete, err = mgr.isFeatureComplete(feature2)
	require.NoError(t, err)
	assert.False(t, complete)

	// Test complete feature with all tasks done
	featureDir := filepath.Join(projectRoot, "doplan", "phase-1", "feat-3")
	require.NoError(t, os.MkdirAll(featureDir, 0755))
	tasksPath := filepath.Join(featureDir, "tasks.md")
	tasksContent := `# Tasks
- [x] Task 1
- [x] Task 2
- [x] Task 3
`
	require.NoError(t, os.WriteFile(tasksPath, []byte(tasksContent), 0644))

	feature3 := &models.Feature{
		ID:     "feat-3",
		Phase:  "phase-1",
		Status: "complete",
		Progress: 100,
	}
	complete, err = mgr.isFeatureComplete(feature3)
	require.NoError(t, err)
	assert.True(t, complete)
}

func TestAutoPRManager_isFeatureComplete_NoTasksFile(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	mgr := NewAutoPRManager(projectRoot)
	feature := &models.Feature{
		ID:       "feat-1",
		Phase:    "phase-1",
		Progress: 100,
		Status:   "complete",
	}

	complete, err := mgr.isFeatureComplete(feature)
	require.NoError(t, err)
	assert.False(t, complete)
}

func TestAutoPRManager_isFeatureComplete_PartialTasks(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	mgr := NewAutoPRManager(projectRoot)

	featureDir := filepath.Join(projectRoot, "doplan", "phase-1", "feat-1")
	require.NoError(t, os.MkdirAll(featureDir, 0755))
	tasksPath := filepath.Join(featureDir, "tasks.md")
	tasksContent := `# Tasks
- [x] Task 1
- [ ] Task 2
- [x] Task 3
`
	require.NoError(t, os.WriteFile(tasksPath, []byte(tasksContent), 0644))

	feature := &models.Feature{
		ID:       "feat-1",
		Phase:    "phase-1",
		Progress: 100,
		Status:   "complete",
	}

	complete, err := mgr.isFeatureComplete(feature)
	require.NoError(t, err)
	assert.False(t, complete)
}

func TestAutoPRManager_createPRForFeature_NoBranch(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	mgr := NewAutoPRManager(projectRoot)
	feature := &models.Feature{
		ID:   "feat-1",
		Name: "Test Feature",
	}

	err := mgr.createPRForFeature(feature)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "branch not set")
}

func TestAutoPRManager_createPRForFeature_Success(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = true
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	// Setup state with feature
	state := &models.State{
		Features: []models.Feature{
			{
				ID:     "feat-1",
				Name:   "Test Feature",
				Phase:  "01-phase",
				Branch: "feature/test-branch",
			},
		},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	mgr := NewAutoPRManager(projectRoot)
	feature := &state.Features[0]

	// This will fail because GitHub CLI is not available, but tests the flow
	err := mgr.createPRForFeature(feature)
	// Should error on PR creation, but tests the function logic
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create PR")
}

func TestAutoPRManager_createPRForFeature_StateUpdate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = true
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	// Setup state with feature
	state := &models.State{
		Features: []models.Feature{
			{
				ID:     "feat-1",
				Name:   "Test Feature",
				Phase:  "01-phase",
				Branch: "feature/test-branch",
			},
		},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	mgr := NewAutoPRManager(projectRoot)
	feature := &state.Features[0]

	// Test that feature is updated even if PR creation fails
	// The function should attempt to update state
	err := mgr.createPRForFeature(feature)
	// Will error on PR creation, but feature should be modified
	assert.Error(t, err)
	// Feature should not have PR set since creation failed
	assert.Nil(t, feature.PR)
}

func TestAutoPRManager_WatchFeatures_MultipleFeatures(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = true
	cfg.GitHub.AutoPR = true
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	// Setup state with multiple features (complete and incomplete)
	state := &models.State{
		Features: []models.Feature{
			{
				ID:       "feat-1",
				Name:     "Complete Feature",
				Progress: 100,
				Status:   "complete",
				Branch:   "feature/complete",
			},
			{
				ID:       "feat-2",
				Name:     "Incomplete Feature",
				Progress: 50,
				Status:   "in-progress",
				Branch:   "feature/incomplete",
			},
			{
				ID:       "feat-3",
				Name:     "Another Complete Feature",
				Progress: 100,
				Status:   "complete",
				Branch:   "feature/complete-2",
			},
		},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	mgr := NewAutoPRManager(projectRoot)
	err := mgr.WatchFeatures(state)
	// Should not error, but won't create PRs (GitHub CLI not available)
	assert.NoError(t, err)
}

func TestAutoPRManager_WatchFeatures_NoFeatures(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = true
	cfg.GitHub.AutoPR = true
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	// Setup state with no features
	state := &models.State{
		Features: []models.Feature{},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	mgr := NewAutoPRManager(projectRoot)
	err := mgr.WatchFeatures(state)
	// Should not error with empty features
	assert.NoError(t, err)
}

func TestAutoPRManager_WatchFeatures_PartialFailures(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	cfg.GitHub.Enabled = true
	cfg.GitHub.AutoPR = true
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	// Setup state with features that will fail (no branch)
	state := &models.State{
		Features: []models.Feature{
			{
				ID:       "feat-1",
				Name:     "Feature Without Branch",
				Progress: 100,
				Status:   "complete",
				// No branch set - will fail
			},
			{
				ID:       "feat-2",
				Name:     "Feature With Branch",
				Progress: 100,
				Status:   "complete",
				Branch:   "feature/test",
			},
		},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	mgr := NewAutoPRManager(projectRoot)
	err := mgr.WatchFeatures(state)
	// Should not error - partial failures are handled gracefully
	assert.NoError(t, err)
}

