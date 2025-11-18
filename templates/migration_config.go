//go:build ignore
// +build ignore

package migration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ConfigMigrator handles configuration migration
type ConfigMigrator struct {
	projectRoot string
}

// NewConfigMigrator creates a new config migrator
func NewConfigMigrator(projectRoot string) *ConfigMigrator {
	return &ConfigMigrator{
		projectRoot: projectRoot,
	}
}

// MigrateConfig migrates config from old JSON to new YAML format
func (c *ConfigMigrator) MigrateConfig() error {
	oldConfigPath := filepath.Join(c.projectRoot, ".cursor", "config", "doplan-config.json")
	newConfigPath := filepath.Join(c.projectRoot, ".doplan", "config.yaml")

	// Load old config
	oldConfig, err := c.loadOldConfig(oldConfigPath)
	if err != nil {
		return fmt.Errorf("failed to load old config: %w", err)
	}

	// Convert to new format
	newConfig, err := c.convertConfig(oldConfig)
	if err != nil {
		return fmt.Errorf("failed to convert config: %w", err)
	}

	// Ensure .doplan directory exists
	if err := os.MkdirAll(filepath.Dir(newConfigPath), 0755); err != nil {
		return fmt.Errorf("failed to create .doplan directory: %w", err)
	}

	// Save new config
	if err := c.saveNewConfig(newConfigPath, newConfig); err != nil {
		return fmt.Errorf("failed to save new config: %w", err)
	}

	return nil
}

// loadOldConfig loads the old JSON config
func (c *ConfigMigrator) loadOldConfig(path string) (*OldConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config OldConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// convertConfig converts old config to new format
func (c *ConfigMigrator) convertConfig(old *OldConfig) (*NewConfig, error) {
	newConfig := &NewConfig{
		Project: ProjectConfig{
			Name:    c.detectProjectName(),
			Type:    c.detectProjectType(),
			Version: old.Version,
			IDE:     old.IDE,
		},
		Github: GithubConfig{
			Repository: "", // Will be prompted if missing
			Enabled:    old.Github.Enabled,
			AutoBranch: old.Github.AutoBranch,
			AutoPR:     old.Github.AutoPR,
		},
		Design: DesignConfig{
			HasPreferences: false,
			TokensPath:     "doplan/design/design-tokens.json",
		},
		Security: SecurityConfig{
			LastScan: nil,
			AutoFix:  false,
		},
		APIs: APIsConfig{
			Configured: []string{},
			Required:   []string{},
		},
		TUI: TUIConfig{
			Theme:      "default",
			Animations: true,
		},
	}

	return newConfig, nil
}

// detectProjectName detects project name from directory
func (c *ConfigMigrator) detectProjectName() string {
	base := filepath.Base(c.projectRoot)
	// Convert to slug
	return strings.ToLower(strings.ReplaceAll(base, " ", "-"))
}

// detectProjectType detects project type from files
func (c *ConfigMigrator) detectProjectType() string {
	// Check for common project files
	if _, err := os.Stat(filepath.Join(c.projectRoot, "package.json")); err == nil {
		return "web"
	}
	if _, err := os.Stat(filepath.Join(c.projectRoot, "go.mod")); err == nil {
		return "api"
	}
	return "web" // default
}

// saveNewConfig saves the new YAML config
func (c *ConfigMigrator) saveNewConfig(path string, config *NewConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// OldConfig represents the old JSON config structure
type OldConfig struct {
	IDE     string `json:"ide"`
	Version string `json:"version"`
	Github  struct {
		Enabled   bool `json:"enabled"`
		AutoBranch bool `json:"autoBranch"`
		AutoPR    bool `json:"autoPR"`
	} `json:"github"`
}

// NewConfig represents the new YAML config structure
type NewConfig struct {
	Project  ProjectConfig  `yaml:"project"`
	Github   GithubConfig   `yaml:"github"`
	Design   DesignConfig   `yaml:"design"`
	Security SecurityConfig `yaml:"security"`
	APIs     APIsConfig     `yaml:"apis"`
	TUI      TUIConfig      `yaml:"tui"`
}

type ProjectConfig struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
	IDE     string `yaml:"ide"`
}

type GithubConfig struct {
	Repository string `yaml:"repository"`
	Enabled    bool   `yaml:"enabled"`
	AutoBranch  bool   `yaml:"autoBranch"`
	AutoPR     bool   `yaml:"autoPR"`
}

type DesignConfig struct {
	HasPreferences bool   `yaml:"hasPreferences"`
	TokensPath     string `yaml:"tokensPath"`
}

type SecurityConfig struct {
	LastScan *string `yaml:"lastScan"`
	AutoFix  bool    `yaml:"autoFix"`
}

type APIsConfig struct {
	Configured []string `yaml:"configured"`
	Required   []string `yaml:"required"`
}

type TUIConfig struct {
	Theme      string `yaml:"theme"`
	Animations bool   `yaml:"animations"`
}

