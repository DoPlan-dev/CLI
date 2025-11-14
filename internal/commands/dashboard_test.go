package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDashboardCommand(t *testing.T) {
	cmd := NewDashboardCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "dashboard", cmd.Use)
	assert.NotEmpty(t, cmd.Short)
	assert.NotEmpty(t, cmd.Long)
}

func TestRunDashboard_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Change to project root
	oldDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(oldDir)

	err = os.Chdir(projectRoot)
	require.NoError(t, err)

	// Run dashboard command
	err = runDashboard(nil, nil)
	// Should not error, just print message
	assert.NoError(t, err)
}

func TestRunDashboard_Installed(t *testing.T) {
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
	}
	require.NoError(t, err)

	// Change to project root
	oldDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(oldDir)

	err = os.Chdir(projectRoot)
	require.NoError(t, err)

	// Run dashboard command
	err = runDashboard(nil, nil)
	require.NoError(t, err)

	// Check dashboard was generated
	dashboardPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	assert.FileExists(t, dashboardPath)
}
