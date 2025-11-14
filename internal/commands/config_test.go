package commands

import (
	"os"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfigCommand(t *testing.T) {
	cmd := NewConfigCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "config", cmd.Use)
}

func TestNewConfigShowCommand(t *testing.T) {
	cmd := NewConfigShowCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "show", cmd.Use)
}

func TestRunConfigShow_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	cmd := NewConfigShowCommand()
	err := cmd.Execute()
	assert.NoError(t, err) // Should not error, just print message
}

func TestRunConfigShow_Installed(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	// Create config
	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := NewConfigShowCommand()
	err = cmd.Execute()
	assert.NoError(t, err)
}

func TestRunConfigShow_JSONFormat(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := NewConfigShowCommand()
	cmd.SetArgs([]string{"--format", "json"})
	err = cmd.Execute()
	assert.NoError(t, err)
}

func TestNewConfigSetCommand(t *testing.T) {
	cmd := NewConfigSetCommand()
	assert.NotNil(t, cmd)
	assert.Contains(t, cmd.Use, "set")
}

func TestRunConfigSet(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	// Create config
	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
		GitHub: models.GitHubConfig{
			Enabled: false,
		},
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := NewConfigSetCommand()
	cmd.SetArgs([]string{"github.enabled", "true"})
	err = cmd.Execute()
	assert.NoError(t, err)

	// Verify update - create a new manager to avoid cache issues
	cfgMgr2 := config.NewManager(projectRoot)
	updated, err := cfgMgr2.LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.True(t, updated.GitHub.Enabled)
}

func TestNewConfigResetCommand(t *testing.T) {
	cmd := NewConfigResetCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "reset", cmd.Use)
}

func TestRunConfigReset(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := NewConfigResetCommand()
	err = cmd.Execute()
	assert.NoError(t, err)
}

func TestNewConfigValidateCommand(t *testing.T) {
	cmd := NewConfigValidateCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "validate", cmd.Use)
}

func TestRunConfigValidate(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := NewConfigValidateCommand()
	err = cmd.Execute()
	assert.NoError(t, err)
}
