package commands

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DoPlan-dev/CLI/test/helpers"
)

func TestNewValidateCommand(t *testing.T) {
	cmd := NewValidateCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "validate", cmd.Use)
}

func TestRunValidate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	originalDir, _ := os.Getwd()
	os.Chdir(projectRoot)
	defer os.Chdir(originalDir)

	cmd := NewValidateCommand()
	err := cmd.Execute()
	assert.NoError(t, err)
}

