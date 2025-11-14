package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTemplatesCommand(t *testing.T) {
	cmd := NewTemplatesCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "templates", cmd.Use)
}

func TestNewTemplatesListCommand(t *testing.T) {
	cmd := NewTemplatesListCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "list", cmd.Use)
}

func TestRunTemplatesList(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	// Create template
	templateDir := filepath.Join(projectRoot, "doplan", "templates")
	err := os.MkdirAll(templateDir, 0755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(templateDir, "test-template.md"), []byte("# Test"), 0644)
	require.NoError(t, err)

	cmd := NewTemplatesListCommand()
	err = cmd.Execute()
	assert.NoError(t, err)
}
