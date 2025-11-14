package generators

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContractsGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}

	gen := NewContractsGenerator(projectRoot, state)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
	assert.Equal(t, state, gen.state)
}

func TestContractsGenerator_Generate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Idea: &models.Idea{
			Name: "Test Project",
		},
		Features: []models.Feature{
			{
				Name:        "Feature 1",
				Description: "Test feature",
			},
		},
	}

	gen := NewContractsGenerator(projectRoot, state)
	err := gen.Generate()
	require.NoError(t, err)

	// Check contracts directory
	contractsDir := filepath.Join(projectRoot, "doplan", "contracts")
	assert.DirExists(t, contractsDir)

	// Check API spec
	apiSpecPath := filepath.Join(contractsDir, "api-spec.json")
	assert.FileExists(t, apiSpecPath)

	data, err := os.ReadFile(apiSpecPath)
	require.NoError(t, err)

	var spec map[string]interface{}
	err = json.Unmarshal(data, &spec)
	require.NoError(t, err)
	assert.Equal(t, "3.0.0", spec["openapi"])

	info := spec["info"].(map[string]interface{})
	assert.Equal(t, "Test Project", info["title"])

	// Check data model
	dataModelPath := filepath.Join(contractsDir, "data-model.md")
	assert.FileExists(t, dataModelPath)

	content, err := os.ReadFile(dataModelPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Data Model Documentation")
	assert.Contains(t, string(content), "Feature 1")
}

func TestContractsGenerator_Generate_EmptyState(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}

	gen := NewContractsGenerator(projectRoot, state)
	err := gen.Generate()
	require.NoError(t, err)

	apiSpecPath := filepath.Join(projectRoot, "doplan", "contracts", "api-spec.json")
	assert.FileExists(t, apiSpecPath)

	dataModelPath := filepath.Join(projectRoot, "doplan", "contracts", "data-model.md")
	assert.FileExists(t, dataModelPath)

	content, err := os.ReadFile(dataModelPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "No data models defined yet")
}

func TestContractsGenerator_generateAPISpec(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Idea: &models.Idea{
			Name: "Test Project",
		},
	}

	gen := NewContractsGenerator(projectRoot, state)
	contractsDir := filepath.Join(projectRoot, "contracts")
	require.NoError(t, os.MkdirAll(contractsDir, 0755))

	err := gen.generateAPISpec(contractsDir)
	require.NoError(t, err)

	apiSpecPath := filepath.Join(contractsDir, "api-spec.json")
	assert.FileExists(t, apiSpecPath)

	data, err := os.ReadFile(apiSpecPath)
	require.NoError(t, err)

	var spec map[string]interface{}
	err = json.Unmarshal(data, &spec)
	require.NoError(t, err)
	assert.Equal(t, "3.0.0", spec["openapi"])

	info := spec["info"].(map[string]interface{})
	assert.Equal(t, "Test Project", info["title"])
	assert.Equal(t, "1.0.0", info["version"])

	// Check servers
	servers := spec["servers"].([]interface{})
	assert.Len(t, servers, 1)
}

func TestContractsGenerator_generateDataModel(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Features: []models.Feature{
			{
				Name:        "Feature 1",
				Description: "Test feature description",
			},
			{
				Name:        "Feature 2",
				Description: "Another feature",
			},
		},
	}

	gen := NewContractsGenerator(projectRoot, state)
	contractsDir := filepath.Join(projectRoot, "contracts")
	require.NoError(t, os.MkdirAll(contractsDir, 0755))

	err := gen.generateDataModel(contractsDir)
	require.NoError(t, err)

	dataModelPath := filepath.Join(contractsDir, "data-model.md")
	assert.FileExists(t, dataModelPath)

	content, err := os.ReadFile(dataModelPath)
	require.NoError(t, err)
	contentStr := string(content)
	assert.Contains(t, contentStr, "Data Model Documentation")
	assert.Contains(t, contentStr, "Feature 1")
	assert.Contains(t, contentStr, "Feature 2")
	assert.Contains(t, contentStr, "Test feature description")
}

func TestContractsGenerator_getProjectName(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// With idea name
	state := &models.State{
		Idea: &models.Idea{
			Name: "My Project",
		},
	}
	gen := NewContractsGenerator(projectRoot, state)
	assert.Equal(t, "My Project", gen.getProjectName())

	// Without idea
	gen2 := NewContractsGenerator(projectRoot, &models.State{})
	assert.Equal(t, "Untitled Project", gen2.getProjectName())

	// With empty idea name
	gen3 := NewContractsGenerator(projectRoot, &models.State{
		Idea: &models.Idea{},
	})
	assert.Equal(t, "Untitled Project", gen3.getProjectName())
}
