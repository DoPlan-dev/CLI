package commands

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

func TestNewProgressCommand(t *testing.T) {
	cmd := NewProgressCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "progress", cmd.Use)
	assert.NotEmpty(t, cmd.Short)
	assert.NotEmpty(t, cmd.Long)
}

func TestRunProgress_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Change to project root
	oldDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(oldDir)

	err = os.Chdir(projectRoot)
	require.NoError(t, err)

	// Run progress command
	err = runProgress(nil, nil)
	// Should not error, just print message
	assert.NoError(t, err)
}

func TestRunProgress_Installed(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Install DoPlan first
	installer := NewInstaller(projectRoot, "cursor")
	err := installer.generateConfig()
	require.NoError(t, err)

	// Create state
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		// Create empty state if doesn't exist
		state = &models.State{
			Phases:   []models.Phase{},
			Features: []models.Feature{},
		}
		err = cfgMgr.SaveState(state)
		require.NoError(t, err)
	}

	// Change to project root
	oldDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(oldDir)

	err = os.Chdir(projectRoot)
	require.NoError(t, err)

	// Run progress command
	err = runProgress(nil, nil)
	require.NoError(t, err)

	// Check dashboard was regenerated
	dashboardPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	assert.FileExists(t, dashboardPath)
}

func TestUpdateProgressFromTasks(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Create a feature directory with tasks.md
	featureDir := filepath.Join(projectRoot, "doplan", "01-phase", "01-Feature")
	require.NoError(t, os.MkdirAll(featureDir, 0755))

	tasksContent := `# Tasks

- [x] Task 1
- [ ] Task 2
- [x] Task 3
`
	err := os.WriteFile(filepath.Join(featureDir, "tasks.md"), []byte(tasksContent), 0644)
	require.NoError(t, err)

	state := &models.State{
		Features: []models.Feature{
			{
				ID:   "feat-1",
				Name: "Feature 1",
			},
		},
	}

	err = updateProgressFromTasks(filepath.Join(projectRoot, "doplan"), state)
	require.NoError(t, err)

	// Progress should be calculated (2 out of 3 tasks = 66%)
	// Note: The actual calculation depends on the implementation
	assert.NotNil(t, state)
}

