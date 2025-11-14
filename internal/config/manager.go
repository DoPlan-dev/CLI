package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Manager handles configuration and state management
type Manager struct {
	projectRoot string
	cache       *CacheManager
}

// NewManager creates a new configuration manager
func NewManager(projectRoot string) *Manager {
	return &Manager{
		projectRoot: projectRoot,
		cache:       NewCacheManager(),
	}
}

// NewConfig creates a new configuration
func NewConfig(ide string) *models.Config {
	return &models.Config{
		IDE:         ide,
		Installed:   true,
		InstalledAt: time.Now(),
		Version:     "1.0.0",
		GitHub: models.GitHubConfig{
			Enabled:    false,
			AutoBranch: true,
			AutoPR:     true,
		},
		Checkpoint: models.CheckpointConfig{
			AutoFeature:  true,
			AutoPhase:    true,
			AutoComplete: true,
		},
		State: models.StateConfig{
			IdeaCaptured:  false,
			PRDGenerated:  false,
			PlanGenerated: false,
		},
	}
}

// SaveConfig saves configuration to file (invalidates cache)
func (m *Manager) SaveConfig(cfg *models.Config) error {
	configPath := filepath.Join(m.projectRoot, ".cursor", "config", "doplan-config.json")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}

	// Update cache
	m.cache.SetConfig(cfg)

	return nil
}

// LoadConfig loads configuration from file (with caching)
func (m *Manager) LoadConfig() (*models.Config, error) {
	// Check cache first
	if cached := m.cache.GetConfig(); cached != nil {
		return cached, nil
	}

	configPath := filepath.Join(m.projectRoot, ".cursor", "config", "doplan-config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg models.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Store in cache
	m.cache.SetConfig(&cfg)

	return &cfg, nil
}

// SaveState saves state to file (invalidates cache)
func (m *Manager) SaveState(state *models.State) error {
	statePath := filepath.Join(m.projectRoot, ".cursor", "config", "doplan-state.json")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(statePath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}

	if err := os.WriteFile(statePath, data, 0644); err != nil {
		return err
	}

	// Update cache
	m.cache.SetState(state)

	return nil
}

// LoadState loads state from file (with caching)
func (m *Manager) LoadState() (*models.State, error) {
	// Check cache first
	if cached := m.cache.GetState(); cached != nil {
		return cached, nil
	}

	statePath := filepath.Join(m.projectRoot, ".cursor", "config", "doplan-state.json")

	data, err := os.ReadFile(statePath)
	if err != nil {
		if os.IsNotExist(err) {
			emptyState := &models.State{
				Progress: models.Progress{
					Overall: 0,
					Phases:  make(map[string]int),
				},
			}
			// Cache empty state
			m.cache.SetState(emptyState)
			return emptyState, nil
		}
		return nil, fmt.Errorf("failed to read state: %w", err)
	}

	var state models.State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("failed to unmarshal state: %w", err)
	}

	// Store in cache
	m.cache.SetState(&state)

	return &state, nil
}

// IsInstalled checks if DoPlan is installed
func IsInstalled(projectRoot string) bool {
	configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
	_, err := os.Stat(configPath)
	return err == nil
}
