package generators

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContextGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewContextGenerator(projectRoot)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
}

func TestContextGenerator_Generate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewContextGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	// Check CONTEXT.md file
	contextPath := filepath.Join(projectRoot, "CONTEXT.md")
	assert.FileExists(t, contextPath)

	content, err := os.ReadFile(contextPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Technology Stack")
}

func TestContextGenerator_Generate_WithGoProject(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create go.mod
	goModPath := filepath.Join(projectRoot, "go.mod")
	require.NoError(t, os.WriteFile(goModPath, []byte("module test\n\ngo 1.21\n"), 0644))

	gen := NewContextGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	contextPath := filepath.Join(projectRoot, "CONTEXT.md")
	content, err := os.ReadFile(contextPath)
	require.NoError(t, err)
	contentStr := string(content)
	assert.Contains(t, contentStr, "Go")
	assert.Contains(t, contentStr, "go.dev/doc")
}

func TestContextGenerator_Generate_WithNodeProject(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create package.json
	packageJsonPath := filepath.Join(projectRoot, "package.json")
	require.NoError(t, os.WriteFile(packageJsonPath, []byte(`{"name": "test", "version": "1.0.0"}`), 0644))

	gen := NewContextGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	contextPath := filepath.Join(projectRoot, "CONTEXT.md")
	content, err := os.ReadFile(contextPath)
	require.NoError(t, err)
	contentStr := string(content)
	assert.Contains(t, contentStr, "JavaScript")
}

func TestContextGenerator_detectTechStack(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewContextGenerator(projectRoot)
	stack := gen.detectTechStack()

	assert.NotNil(t, stack)
	assert.NotNil(t, stack.Languages)
	assert.NotNil(t, stack.Frameworks)
	assert.NotNil(t, stack.CLIs)
	assert.NotNil(t, stack.Services)
	assert.NotNil(t, stack.Databases)
	assert.NotNil(t, stack.Tools)
}

func TestContextGenerator_detectTechStack_Go(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create go.mod
	goModPath := filepath.Join(projectRoot, "go.mod")
	require.NoError(t, os.WriteFile(goModPath, []byte("module test\n\ngo 1.21\n"), 0644))

	gen := NewContextGenerator(projectRoot)
	stack := gen.detectTechStack()

	// Should detect Go
	found := false
	for _, lang := range stack.Languages {
		if lang.Name == "Go" {
			found = true
			assert.Equal(t, "1.21", lang.Version)
			break
		}
	}
	assert.True(t, found, "Go should be detected")
}

func TestContextGenerator_detectTechStack_Node(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create package.json
	packageJsonPath := filepath.Join(projectRoot, "package.json")
	require.NoError(t, os.WriteFile(packageJsonPath, []byte(`{"name": "test"}`), 0644))

	gen := NewContextGenerator(projectRoot)
	stack := gen.detectTechStack()

	// Should detect JavaScript/TypeScript
	found := false
	for _, lang := range stack.Languages {
		if lang.Name == "JavaScript/TypeScript" {
			found = true
			break
		}
	}
	assert.True(t, found, "JavaScript/TypeScript should be detected")
}

func TestContextGenerator_fileExists(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewContextGenerator(projectRoot)

	// Non-existent file
	assert.False(t, gen.fileExists("nonexistent.txt"))

	// Create a file
	testFile := filepath.Join(projectRoot, "test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("test"), 0644))
	assert.True(t, gen.fileExists("test.txt"))
}

func TestContextGenerator_extractGoVersion(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create go.mod with version
	goModPath := filepath.Join(projectRoot, "go.mod")
	require.NoError(t, os.WriteFile(goModPath, []byte("module test\n\ngo 1.21\n"), 0644))

	gen := NewContextGenerator(projectRoot)
	version := gen.extractGoVersion()

	assert.Equal(t, "1.21", version)
}

func TestContextGenerator_extractNodeVersion(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Create package.json
	packageJsonPath := filepath.Join(projectRoot, "package.json")
	require.NoError(t, os.WriteFile(packageJsonPath, []byte(`{"name": "test", "engines": {"node": ">=18.0.0"}}`), 0644))

	gen := NewContextGenerator(projectRoot)
	version := gen.extractNodeVersion()

	// Should extract version or return default
	assert.NotEmpty(t, version)
}

func TestContextGenerator_generateContextContent(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewContextGenerator(projectRoot)
	stack := gen.detectTechStack()

	content := gen.generateContextContent(stack)

	assert.Contains(t, content, "Technology Stack")
	// Sections only appear if there are items, so check for CLIs which is always present
	assert.Contains(t, content, "CLIs and Development Tools")
	assert.Contains(t, content, "DoPlan CLI")
}
