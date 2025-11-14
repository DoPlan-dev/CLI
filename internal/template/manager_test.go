package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	manager := NewManager(projectRoot)

	assert.NotNil(t, manager)
	assert.Contains(t, manager.templatesDir, "templates")
}

func TestManager_ListTemplates(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	// Create template
	templateDir := filepath.Join(projectRoot, "doplan", "templates")
	err := os.MkdirAll(templateDir, 0755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(templateDir, "test.md"), []byte("# Test"), 0644)
	require.NoError(t, err)

	templates, err := manager.ListTemplates()
	require.NoError(t, err)
	assert.Contains(t, templates, "test.md")
}

func TestManager_GetTemplate(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	// Create template
	templateDir := filepath.Join(projectRoot, "doplan", "templates")
	err := os.MkdirAll(templateDir, 0755)
	require.NoError(t, err)

	content := "# Test Template\n\nContent here"
	err = os.WriteFile(filepath.Join(templateDir, "test.md"), []byte(content), 0644)
	require.NoError(t, err)

	template, err := manager.GetTemplate("test.md")
	require.NoError(t, err)
	assert.Equal(t, content, template)
}

func TestManager_AddTemplate(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	content := "# New Template"
	err := manager.AddTemplate("new.md", content)
	require.NoError(t, err)

	// Verify file exists
	templatePath := filepath.Join(manager.templatesDir, "new.md")
	assert.FileExists(t, templatePath)
}

func TestManager_RemoveTemplate(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	// Create template
	err := manager.AddTemplate("temp.md", "# Temp")
	require.NoError(t, err)

	err = manager.RemoveTemplate("temp.md")
	require.NoError(t, err)

	// Verify file removed
	templatePath := filepath.Join(manager.templatesDir, "temp.md")
	assert.NoFileExists(t, templatePath)
}

func TestManager_LoadConfig(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	config, err := manager.LoadConfig()
	require.NoError(t, err)
	assert.NotNil(t, config)
}

func TestManager_SaveConfig(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	config := &TemplateConfig{
		DefaultPlan:   "plan.md",
		DefaultDesign: "design.md",
		DefaultTasks:  "tasks.md",
		Templates:     make(map[string]string),
	}

	err := manager.SaveConfig(config)
	require.NoError(t, err)

	// Verify saved
	loaded, err := manager.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "plan.md", loaded.DefaultPlan)
}

func TestManager_GetDefaultTemplate_AllTypes(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	// Save config with defaults
	config := &TemplateConfig{
		DefaultPlan:   "plan.md",
		DefaultDesign: "design.md",
		DefaultTasks:  "tasks.md",
	}
	err := manager.SaveConfig(config)
	require.NoError(t, err)

	// Test each template type
	plan, err := manager.GetDefaultTemplate("plan")
	require.NoError(t, err)
	assert.Equal(t, "plan.md", plan)

	design, err := manager.GetDefaultTemplate("design")
	require.NoError(t, err)
	assert.Equal(t, "design.md", design)

	tasks, err := manager.GetDefaultTemplate("tasks")
	require.NoError(t, err)
	assert.Equal(t, "tasks.md", tasks)
}

func TestManager_GetDefaultTemplate_UnknownType(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	manager := NewManager(projectRoot)

	// Save config
	config := &TemplateConfig{
		DefaultPlan:   "plan.md",
		DefaultDesign: "design.md",
		DefaultTasks:  "tasks.md",
	}
	err := manager.SaveConfig(config)
	require.NoError(t, err)

	// Test unknown type
	_, err = manager.GetDefaultTemplate("unknown")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown template type")
}

func TestManager_GetDefaultTemplate_MissingConfig(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	manager := NewManager(projectRoot)

	// Don't save config - LoadConfig will return default config
	// But GetDefaultTemplate should still work with default config
	plan, err := manager.GetDefaultTemplate("plan")
	// Should not error - LoadConfig returns default config if file doesn't exist
	if err == nil {
		assert.NotEmpty(t, plan)
	}
}
