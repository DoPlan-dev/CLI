package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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
// Supports both new YAML format (.doplan/config.yaml) and old JSON format (.cursor/config/doplan-config.json)
func (m *Manager) LoadConfig() (*models.Config, error) {
	// Check cache first
	if cached := m.cache.GetConfig(); cached != nil {
		return cached, nil
	}

	// Try new YAML format first
	newConfigPath := filepath.Join(m.projectRoot, ".doplan", "config.yaml")
	if _, err := os.Stat(newConfigPath); err == nil {
		cfg, err := m.loadConfigYAML(newConfigPath)
		if err == nil {
			m.cache.SetConfig(cfg)
			return cfg, nil
		}
		// If YAML load fails, fall through to old format
	}

	// Fallback to old JSON format
	oldConfigPath := filepath.Join(m.projectRoot, ".cursor", "config", "doplan-config.json")
	data, err := os.ReadFile(oldConfigPath)
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

// loadConfigYAML loads configuration from YAML file
func (m *Manager) loadConfigYAML(path string) (*models.Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read YAML config: %w", err)
	}

	// Map YAML structure to Config model
	cfg := &models.Config{
		IDE:         viper.GetString("project.ide"),
		Installed:   true,
		InstalledAt: time.Now(), // Use current time as default (config may not have this field yet)
		Version:     viper.GetString("project.version"),
		GitHub: models.GitHubConfig{
			Enabled:    viper.GetBool("github.enabled"),
			AutoBranch: viper.GetBool("github.autoBranch"),
			AutoPR:     viper.GetBool("github.autoPR"),
		},
		Checkpoint: models.CheckpointConfig{
			AutoFeature:  true, // Defaults
			AutoPhase:    true,
			AutoComplete: true,
		},
		State: models.StateConfig{
			IdeaCaptured:  false,
			PRDGenerated:  false,
			PlanGenerated: false,
		},
	}

	return cfg, nil
}

// SaveConfigV2 saves configuration in new YAML format
func (m *Manager) SaveConfigV2(cfg *models.Config) error {
	configPath := filepath.Join(m.projectRoot, ".doplan", "config.yaml")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Convert to YAML structure
	yamlConfig := map[string]interface{}{
		"project": map[string]interface{}{
			"name":    "", // Project name not stored in config yet
			"type":    "", // Project type not stored in config yet
			"version": cfg.Version,
			"ide":     cfg.IDE,
		},
		"github": map[string]interface{}{
			"repository": "", // GitHub repository URL not stored in config yet
			"enabled":    cfg.GitHub.Enabled,
			"autoBranch": cfg.GitHub.AutoBranch,
			"autoPR":     cfg.GitHub.AutoPR,
		},
		"design": map[string]interface{}{
			"hasPreferences": false,
			"tokensPath":     "doplan/design/design-tokens.json",
		},
		"security": map[string]interface{}{
			"lastScan": nil,
			"autoFix":  false,
		},
		"apis": map[string]interface{}{
			"configured": []string{},
			"required":   []string{},
		},
		"tui": map[string]interface{}{
			"theme":      "default",
			"animations": true,
		},
	}

	data, err := yaml.Marshal(yamlConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}

	// Update cache
	m.cache.SetConfig(cfg)

	return nil
}

// SaveState saves state to file (invalidates cache)
// Uses new location (.doplan/state.json) but supports old location for migration
func (m *Manager) SaveState(state *models.State) error {
	// Use new location
	statePath := filepath.Join(m.projectRoot, ".doplan", "state.json")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(statePath), 0755); err != nil {
		return fmt.Errorf("failed to create .doplan directory: %w", err)
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
// Supports both new location (.doplan/state.json) and old location (.cursor/config/doplan-state.json)
func (m *Manager) LoadState() (*models.State, error) {
	// Check cache first
	if cached := m.cache.GetState(); cached != nil {
		return cached, nil
	}

	// Try new location first
	newStatePath := filepath.Join(m.projectRoot, ".doplan", "state.json")
	data, err := os.ReadFile(newStatePath)
	if err != nil {
		// Fallback to old location
		oldStatePath := filepath.Join(m.projectRoot, ".cursor", "config", "doplan-state.json")
		data, err = os.ReadFile(oldStatePath)
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
// Checks both new location (.doplan/config.yaml) and old location (.cursor/config/doplan-config.json)
func IsInstalled(projectRoot string) bool {
	// Check new location first
	newConfigPath := filepath.Join(projectRoot, ".doplan", "config.yaml")
	if _, err := os.Stat(newConfigPath); err == nil {
		return true
	}

	// Check old location
	oldConfigPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
	_, err := os.Stat(oldConfigPath)
	return err == nil
}
