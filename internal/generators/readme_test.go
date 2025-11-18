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

	// Check project-first structure
	assert.Contains(t, content, "🚀 Quick Start")
	assert.Contains(t, content, "📋 Features")
	assert.Contains(t, content, "🛠️ Tech Stack")
	assert.Contains(t, content, "📁 Project Structure")
	assert.Contains(t, content, "🔑 Environment Variables")
	assert.Contains(t, content, "📚 Documentation")
	// Check DoPlan section is in collapsible details
	assert.Contains(t, content, "<details>")
	assert.Contains(t, content, "Powered by DoPlan")
	assert.Contains(t, content, "/Discuss")
	assert.Contains(t, content, "doplan dashboard")
}

func TestREADMEGenerator_Generate_ContainsWorkflow(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewREADMEGenerator(projectRoot)
	content := gen.Generate()

	// Check project-first sections are present (workflow steps removed in new structure)
	assert.Contains(t, content, "Quick Start")
	assert.Contains(t, content, "Features")
	assert.Contains(t, content, "Tech Stack")
	assert.Contains(t, content, "Project Structure")
}

func TestREADMEGenerator_Generate_ContainsCommands(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewREADMEGenerator(projectRoot)
	content := gen.Generate()

	// Check DoPlan commands are in collapsible section
	assert.Contains(t, content, "/Discuss")
	assert.Contains(t, content, "/Generate")
	assert.Contains(t, content, "/Plan")
	assert.Contains(t, content, "doplan dashboard")
	assert.Contains(t, content, "doplan progress")
	assert.Contains(t, content, "doplan github")
}
