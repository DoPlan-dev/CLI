package generators

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTemplatesGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewTemplatesGenerator(projectRoot)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
}

func TestTemplatesGenerator_Generate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewTemplatesGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	// Check templates directory
	templatesDir := filepath.Join(projectRoot, "doplan", "templates")
	assert.DirExists(t, templatesDir)

	// Check all template files
	templates := []string{
		"plan-template.md",
		"design-template.md",
		"tasks-template.md",
	}

	for _, template := range templates {
		templatePath := filepath.Join(templatesDir, template)
		assert.FileExists(t, templatePath)

		content, err := os.ReadFile(templatePath)
		require.NoError(t, err)
		assert.NotEmpty(t, content)
	}
}

func TestTemplatesGenerator_Generate_ExistingDirectory(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	templatesDir := filepath.Join(projectRoot, "doplan", "templates")
	require.NoError(t, os.MkdirAll(templatesDir, 0755))

	gen := NewTemplatesGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	// Should still create files
	assert.FileExists(t, filepath.Join(templatesDir, "plan-template.md"))
}

func TestGetPlanTemplate(t *testing.T) {
	template := getPlanTemplate()
	assert.Contains(t, template, "Feature Plan")
	assert.Contains(t, template, "{{.Feature.Name}}")
	assert.Contains(t, template, "{{.Feature.Description}}")
	assert.Contains(t, template, "{{range .Feature.Objectives}}")
}

func TestGetDesignTemplate(t *testing.T) {
	template := getDesignTemplate()
	assert.Contains(t, template, "Feature Design")
	assert.Contains(t, template, "{{.Feature.Name}}")
	assert.Contains(t, template, "Architecture")
}

func TestGetTasksTemplate(t *testing.T) {
	template := getTasksTemplate()
	assert.Contains(t, template, "Feature Tasks")
	assert.Contains(t, template, "{{.Feature.Name}}")
	assert.Contains(t, template, "{{range .Feature.TaskPhases}}")
	assert.Contains(t, template, "{{range .Tasks}}")
}
