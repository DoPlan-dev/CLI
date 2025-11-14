package commands

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DoPlan-dev/CLI/test/helpers"
)

func TestNewGitHubCommand(t *testing.T) {
	cmd := NewGitHubCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "github", cmd.Use)
}

func TestRunGitHub_NotInstalled(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	cmd := NewGitHubCommand()
	err := cmd.Execute()
	assert.NoError(t, err) // Should handle gracefully
}

