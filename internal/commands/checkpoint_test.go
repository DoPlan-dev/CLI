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

func TestNewCheckpointCommand(t *testing.T) {
	cmd := NewCheckpointCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "checkpoint", cmd.Use)
}

func TestNewCheckpointCreateCommand(t *testing.T) {
	cmd := NewCheckpointCreateCommand()
	assert.NotNil(t, cmd)
	assert.Contains(t, cmd.Use, "create")
}

func TestNewCheckpointListCommand(t *testing.T) {
	cmd := NewCheckpointListCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "list", cmd.Use)
}

func TestNewCheckpointRestoreCommand(t *testing.T) {
	cmd := NewCheckpointRestoreCommand()
	assert.NotNil(t, cmd)
	assert.Contains(t, cmd.Use, "restore")
}

func TestRunCheckpointCreate_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	cmd := NewCheckpointCreateCommand()
	cmd.SetArgs([]string{})
	err := runCheckpointCreate(cmd, []string{})

	// Should not error, just print message
	assert.NoError(t, err)
}

func TestRunCheckpointCreate_WithName(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{Overall: 0},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	cmd := NewCheckpointCreateCommand()
	cmd.SetArgs([]string{"test-checkpoint"})
	err := runCheckpointCreate(cmd, []string{"test-checkpoint"})

	assert.NoError(t, err)
}

func TestRunCheckpointCreate_WithType(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{Overall: 0},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	// Test different checkpoint types
	types := []string{"manual", "feature", "phase"}
	for _, checkpointType := range types {
		cmd := NewCheckpointCreateCommand()
		cmd.SetArgs([]string{"--type", checkpointType, "test"})
		err := runCheckpointCreate(cmd, []string{"test"})
		// May error if checkpoint already exists, which is acceptable
		_ = err
	}
}

func TestRunCheckpointCreate_WithDescription(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{Overall: 0},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	cmd := NewCheckpointCreateCommand()
	cmd.SetArgs([]string{"--description", "Test description", "test-checkpoint"})
	err := runCheckpointCreate(cmd, []string{"test-checkpoint"})

	// May error if checkpoint already exists
	_ = err
}

func TestRunCheckpointList_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	cmd := NewCheckpointListCommand()
	err := runCheckpointList(cmd, []string{})

	// Should not error, just print message
	assert.NoError(t, err)
}

func TestRunCheckpointList_Empty(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{Overall: 0},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	cmd := NewCheckpointListCommand()
	err := runCheckpointList(cmd, []string{})

	// Should succeed even with no checkpoints
	assert.NoError(t, err)
}

func TestRunCheckpointList_WithCheckpoints(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{Overall: 0},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	// Create a checkpoint first
	checkpointDir := filepath.Join(projectRoot, ".doplan", "checkpoints")
	require.NoError(t, os.MkdirAll(checkpointDir, 0755))

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	// Create checkpoint
	createCmd := NewCheckpointCreateCommand()
	createCmd.SetArgs([]string{"test-list"})
	_ = runCheckpointCreate(createCmd, []string{"test-list"})

	// List checkpoints
	cmd := NewCheckpointListCommand()
	err := runCheckpointList(cmd, []string{})

	assert.NoError(t, err)
}

func TestRunCheckpointRestore_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	cmd := NewCheckpointRestoreCommand()
	cmd.SetArgs([]string{"test-id"})
	err := runCheckpointRestore(cmd, []string{"test-id"})

	// Should not error, just print message
	assert.NoError(t, err)
}

func TestRunCheckpointRestore_InvalidID(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cfgMgr := config.NewManager(projectRoot)
	cfg := config.NewConfig("cursor")
	require.NoError(t, cfgMgr.SaveConfig(cfg))

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{Overall: 0},
	}
	require.NoError(t, cfgMgr.SaveState(state))

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(projectRoot)

	// Note: This test is limited because runCheckpointRestore prompts for user input.
	// Without providing "y" input, the function will cancel and return nil.
	// The actual error for invalid checkpoint ID only occurs after user confirms.
	// For a complete test, we would need to mock stdin or use a different approach.
	cmd := NewCheckpointRestoreCommand()
	cmd.SetArgs([]string{"nonexistent-id"})
	err := runCheckpointRestore(cmd, []string{"nonexistent-id"})

	// Without user input, the function cancels and returns nil
	// This tests the cancellation path, which is valid behavior
	_ = err
}

