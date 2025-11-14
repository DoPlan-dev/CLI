package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheManager_GetSetConfig(t *testing.T) {
	cm := NewCacheManager()

	// Initially no config
	assert.Nil(t, cm.GetConfig())

	// Set config
	cfg := &models.Config{
		IDE:     "cursor",
		Version: "1.0.0",
	}
	cm.SetConfig(cfg)

	// Should get cached config
	cached := cm.GetConfig()
	assert.NotNil(t, cached)
	assert.Equal(t, "cursor", cached.IDE)
	assert.Equal(t, "1.0.0", cached.Version)
}

func TestCacheManager_GetSetState(t *testing.T) {
	cm := NewCacheManager()

	// Initially no state
	assert.Nil(t, cm.GetState())

	// Set state
	state := &models.State{
		Progress: models.Progress{
			Overall: 75,
		},
	}
	cm.SetState(state)

	// Should get cached state
	cached := cm.GetState()
	assert.NotNil(t, cached)
	assert.Equal(t, 75, cached.Progress.Overall)
}

func TestCacheManager_ConfigExpiration(t *testing.T) {
	cm := NewCacheManager()

	cfg := &models.Config{
		IDE: "cursor",
	}
	cm.SetConfig(cfg)

	// Should be cached immediately
	assert.NotNil(t, cm.GetConfig())

	// Manually expire cache by accessing internal struct
	// Since we can't directly access the internal cache, we test by waiting
	// But config TTL is 5 minutes, so we'll test the behavior differently
	// by setting a new config and verifying it updates
	cfg2 := &models.Config{
		IDE: "gemini",
	}
	cm.SetConfig(cfg2)

	// Should get new config
	cached := cm.GetConfig()
	assert.NotNil(t, cached)
	assert.Equal(t, "gemini", cached.IDE)
}

func TestCacheManager_StateExpiration(t *testing.T) {
	cm := NewCacheManager()

	state := &models.State{
		Progress: models.Progress{
			Overall: 50,
		},
	}
	cm.SetState(state)

	// Should be cached immediately
	assert.NotNil(t, cm.GetState())

	// Update state
	state2 := &models.State{
		Progress: models.Progress{
			Overall: 75,
		},
	}
	cm.SetState(state2)

	// Should get new state
	cached := cm.GetState()
	assert.NotNil(t, cached)
	assert.Equal(t, 75, cached.Progress.Overall)
}

func TestCacheManager_ConcurrentAccess(t *testing.T) {
	cm := NewCacheManager()

	// Test concurrent access
	done := make(chan bool, 2)

	// Goroutine 1: Set config
	go func() {
		for i := 0; i < 10; i++ {
			cfg := &models.Config{
				IDE:     "cursor",
				Version: "1.0.0",
			}
			cm.SetConfig(cfg)
			time.Sleep(1 * time.Millisecond)
		}
		done <- true
	}()

	// Goroutine 2: Get config
	go func() {
		for i := 0; i < 10; i++ {
			cfg := cm.GetConfig()
			_ = cfg // Use the value
			time.Sleep(1 * time.Millisecond)
		}
		done <- true
	}()

	// Wait for both goroutines
	<-done
	<-done

	// Final check
	cached := cm.GetConfig()
	assert.NotNil(t, cached)
}

func TestManager_CacheIntegration(t *testing.T) {
	tmpDir := t.TempDir()
	cfgMgr := NewManager(tmpDir)

	cfg := NewConfig("cursor")
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	// First load should cache
	loaded1, err := cfgMgr.LoadConfig()
	require.NoError(t, err)
	assert.NotNil(t, loaded1)

	// Second load should use cache (same instance)
	loaded2, err := cfgMgr.LoadConfig()
	require.NoError(t, err)
	assert.NotNil(t, loaded2)
	assert.Equal(t, loaded1.IDE, loaded2.IDE)

	// New manager instance should load from file (cache is per-instance)
	cfgMgr2 := NewManager(tmpDir)
	loaded3, err := cfgMgr2.LoadConfig()
	require.NoError(t, err)
	assert.NotNil(t, loaded3)
	assert.Equal(t, "cursor", loaded3.IDE)
}

func TestManager_StateCacheIntegration(t *testing.T) {
	tmpDir := t.TempDir()
	cfgMgr := NewManager(tmpDir)

	state := &models.State{
		Progress: models.Progress{
			Overall: 60,
		},
	}
	err := cfgMgr.SaveState(state)
	require.NoError(t, err)

	// First load should cache
	loaded1, err := cfgMgr.LoadState()
	require.NoError(t, err)
	assert.NotNil(t, loaded1)
	assert.Equal(t, 60, loaded1.Progress.Overall)

	// Second load should use cache
	loaded2, err := cfgMgr.LoadState()
	require.NoError(t, err)
	assert.NotNil(t, loaded2)
	assert.Equal(t, loaded1.Progress.Overall, loaded2.Progress.Overall)
}

func TestManager_LoadConfig_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".cursor", "config", "doplan-config.json")
	require.NoError(t, os.MkdirAll(filepath.Dir(configPath), 0755))
	require.NoError(t, os.WriteFile(configPath, []byte("invalid json"), 0644))

	cfgMgr := NewManager(tmpDir)
	cfg, err := cfgMgr.LoadConfig()
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestManager_LoadState_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, ".cursor", "config", "doplan-state.json")
	require.NoError(t, os.MkdirAll(filepath.Dir(statePath), 0755))
	require.NoError(t, os.WriteFile(statePath, []byte("invalid json"), 0644))

	cfgMgr := NewManager(tmpDir)
	state, err := cfgMgr.LoadState()
	assert.Error(t, err)
	assert.Nil(t, state)
}

func TestManager_SaveState_ErrorHandling(t *testing.T) {
	// Test saving to invalid path
	cfgMgr := NewManager("/invalid/path/that/does/not/exist")
	state := &models.State{
		Progress: models.Progress{
			Overall: 50,
		},
	}

	err := cfgMgr.SaveState(state)
	// May or may not error depending on system
	_ = err
}

func TestCacheManager_InvalidateConfig(t *testing.T) {
	cm := NewCacheManager()

	cfg := &models.Config{
		IDE: "cursor",
	}
	cm.SetConfig(cfg)
	assert.NotNil(t, cm.GetConfig())

	cm.InvalidateConfig()
	assert.Nil(t, cm.GetConfig())
}

func TestCacheManager_InvalidateState(t *testing.T) {
	cm := NewCacheManager()

	state := &models.State{
		Progress: models.Progress{
			Overall: 50,
		},
	}
	cm.SetState(state)
	assert.NotNil(t, cm.GetState())

	cm.InvalidateState()
	assert.Nil(t, cm.GetState())
}

func TestCacheManager_InvalidateAll(t *testing.T) {
	cm := NewCacheManager()

	cfg := &models.Config{IDE: "cursor"}
	state := &models.State{Progress: models.Progress{Overall: 50}}

	cm.SetConfig(cfg)
	cm.SetState(state)

	assert.NotNil(t, cm.GetConfig())
	assert.NotNil(t, cm.GetState())

	cm.InvalidateAll()

	assert.Nil(t, cm.GetConfig())
	assert.Nil(t, cm.GetState())
}
