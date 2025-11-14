package template

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Manager manages templates
type Manager struct {
	templatesDir string
	configPath   string
}

// TemplateConfig stores template configuration
type TemplateConfig struct {
	DefaultPlan   string            `json:"defaultPlan"`
	DefaultDesign string            `json:"defaultDesign"`
	DefaultTasks  string            `json:"defaultTasks"`
	Templates     map[string]string `json:"templates"` // name -> file path
}

// NewManager creates a new template manager
func NewManager(projectRoot string) *Manager {
	return &Manager{
		templatesDir: filepath.Join(projectRoot, "doplan", "templates"),
		configPath:   filepath.Join(projectRoot, ".doplan", "templates.json"),
	}
}

// ListTemplates lists all available templates
func (m *Manager) ListTemplates() ([]string, error) {
	if _, err := os.Stat(m.templatesDir); os.IsNotExist(err) {
		return []string{}, nil
	}

	entries, err := os.ReadDir(m.templatesDir)
	if err != nil {
		return nil, err
	}

	var templates []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			templates = append(templates, entry.Name())
		}
	}

	return templates, nil
}

// GetTemplate returns template content
func (m *Manager) GetTemplate(name string) (string, error) {
	path := filepath.Join(m.templatesDir, name)
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("template not found: %s", name)
	}
	return string(content), nil
}

// AddTemplate adds a new template
func (m *Manager) AddTemplate(name string, content string) error {
	if err := os.MkdirAll(m.templatesDir, 0755); err != nil {
		return err
	}

	path := filepath.Join(m.templatesDir, name)
	return os.WriteFile(path, []byte(content), 0644)
}

// RemoveTemplate removes a template
func (m *Manager) RemoveTemplate(name string) error {
	path := filepath.Join(m.templatesDir, name)
	return os.Remove(path)
}

// LoadConfig loads template configuration
func (m *Manager) LoadConfig() (*TemplateConfig, error) {
	if _, err := os.Stat(m.configPath); os.IsNotExist(err) {
		return &TemplateConfig{
			DefaultPlan:   "plan-template.md",
			DefaultDesign: "design-template.md",
			DefaultTasks:  "tasks-template.md",
			Templates:     make(map[string]string),
		}, nil
	}

	data, err := os.ReadFile(m.configPath)
	if err != nil {
		return nil, err
	}

	var config TemplateConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if config.Templates == nil {
		config.Templates = make(map[string]string)
	}

	return &config, nil
}

// SaveConfig saves template configuration
func (m *Manager) SaveConfig(config *TemplateConfig) error {
	if err := os.MkdirAll(filepath.Dir(m.configPath), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.configPath, data, 0644)
}

// GetDefaultTemplate returns the default template for a type
func (m *Manager) GetDefaultTemplate(templateType string) (string, error) {
	config, err := m.LoadConfig()
	if err != nil {
		return "", err
	}

	switch templateType {
	case "plan":
		return config.DefaultPlan, nil
	case "design":
		return config.DefaultDesign, nil
	case "tasks":
		return config.DefaultTasks, nil
	default:
		return "", fmt.Errorf("unknown template type: %s", templateType)
	}
}
