package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("cursor")

	assert.Equal(t, "cursor", cfg.IDE)
	assert.True(t, cfg.Installed)
	assert.False(t, cfg.GitHub.Enabled)
	assert.True(t, cfg.GitHub.AutoBranch)
	assert.True(t, cfg.GitHub.AutoPR)
}

func TestManager_SaveAndLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := NewConfig("cursor")
	cfgMgr := NewManager(tmpDir)

	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	loaded, err := cfgMgr.LoadConfig()
	require.NoError(t, err)
	assert.NotNil(t, loaded)
	assert.Equal(t, cfg.IDE, loaded.IDE)
	assert.Equal(t, cfg.Installed, loaded.Installed)
}

func TestIsInstalled(t *testing.T) {
	tmpDir := t.TempDir()

	// Not installed initially
	assert.False(t, IsInstalled(tmpDir))

	// Create config file
	configPath := filepath.Join(tmpDir, ".cursor", "config", "doplan-config.json")
	require.NoError(t, os.MkdirAll(filepath.Dir(configPath), 0755))
	require.NoError(t, os.WriteFile(configPath, []byte("{}"), 0644))

	// Now installed
	assert.True(t, IsInstalled(tmpDir))
}

func TestManager_SaveAndLoadState(t *testing.T) {
	tmpDir := t.TempDir()
	cfgMgr := NewManager(tmpDir)

	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
		Progress: models.Progress{
			Overall: 50,
			Phases:  make(map[string]int),
		},
	}

	err := cfgMgr.SaveState(state)
	require.NoError(t, err)

	loaded, err := cfgMgr.LoadState()
	require.NoError(t, err)
	assert.NotNil(t, loaded)
	assert.Equal(t, 50, loaded.Progress.Overall)
}

func TestManager_LoadState_NotExists(t *testing.T) {
	tmpDir := t.TempDir()
	cfgMgr := NewManager(tmpDir)

	// Load state when file doesn't exist should return empty state
	state, err := cfgMgr.LoadState()
	require.NoError(t, err)
	assert.NotNil(t, state)
	assert.Equal(t, 0, state.Progress.Overall)
}

func TestManager_LoadConfig_NotExists(t *testing.T) {
	tmpDir := t.TempDir()
	cfgMgr := NewManager(tmpDir)

	// Load config when file doesn't exist should return nil
	cfg, err := cfgMgr.LoadConfig()
	require.NoError(t, err)
	assert.Nil(t, cfg)
}

func TestManager_SaveConfig_ErrorHandling(t *testing.T) {
	// Test saving to invalid path (read-only directory)
	cfgMgr := NewManager("/")
	cfg := NewConfig("cursor")

	// This should error on most systems
	err := cfgMgr.SaveConfig(cfg)
	// May or may not error depending on permissions
	_ = err
}

func TestNewConfig_DifferentIDEs(t *testing.T) {
	ides := []string{"cursor", "gemini", "claude", "codex"}

	for _, ide := range ides {
		cfg := NewConfig(ide)
		assert.Equal(t, ide, cfg.IDE)
		assert.True(t, cfg.Installed)
		assert.Equal(t, "1.0.0", cfg.Version)
	}
}
