package commands

import (
	"os"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
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
