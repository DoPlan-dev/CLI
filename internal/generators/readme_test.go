package generators

import (
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewREADMEGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewREADMEGenerator(projectRoot)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
}

func TestREADMEGenerator_Generate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewREADMEGenerator(projectRoot)
	content := gen.Generate()

	assert.Contains(t, content, "DoPlan - Project Workflow Manager")
	assert.Contains(t, content, "Quick Start")
	assert.Contains(t, content, "/Discuss")
	assert.Contains(t, content, "/Refine")
	assert.Contains(t, content, "/Generate")
	assert.Contains(t, content, "/Plan")
	assert.Contains(t, content, "/Implement")
	assert.Contains(t, content, "doplan dashboard")
	assert.Contains(t, content, ".cursor/commands/")
	assert.Contains(t, content, "doplan/")
}

func TestREADMEGenerator_Generate_ContainsWorkflow(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewREADMEGenerator(projectRoot)
	content := gen.Generate()

	// Check workflow steps are present
	assert.Contains(t, content, "Step 1: Discuss Your Idea")
	assert.Contains(t, content, "Step 2: Refine Your Idea")
	assert.Contains(t, content, "Step 3: Generate Documentation")
	assert.Contains(t, content, "Step 4: Create Your Plan")
}

func TestREADMEGenerator_Generate_ContainsCommands(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewREADMEGenerator(projectRoot)
	content := gen.Generate()

	// Check all commands are mentioned
	assert.Contains(t, content, "/Discuss")
	assert.Contains(t, content, "/Refine")
	assert.Contains(t, content, "/Generate")
	assert.Contains(t, content, "/Plan")
	assert.Contains(t, content, "/Implement")
	assert.Contains(t, content, "/Next")
	assert.Contains(t, content, "/Progress")
	assert.Contains(t, content, "/Dashboard")
}
