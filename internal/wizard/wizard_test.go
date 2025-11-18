package wizard

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunNewProjectWizard(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)
	
	err = os.Chdir(projectRoot)
	require.NoError(t, err)

	// Test that wizard can be created
	wizard := NewNewProjectWizard(projectRoot)
	assert.NotNil(t, wizard)
	assert.Equal(t, projectRoot, wizard.projectRoot)
}

func TestRunAdoptProjectWizard(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	
	// Create some project files to simulate existing project
	helpers.WriteTestFile(t, projectRoot, "package.json", []byte(`{"name": "test-project"}`))
	
	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)
	
	err = os.Chdir(projectRoot)
	require.NoError(t, err)

	// Test that wizard can be created
	wizard := NewAdoptProjectWizard(projectRoot)
	assert.NotNil(t, wizard)
	assert.Equal(t, projectRoot, wizard.projectRoot)
}

func TestNewNewProjectWizard(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	
	wizard := NewNewProjectWizard(projectRoot)
	assert.NotNil(t, wizard)
	assert.NotNil(t, wizard.model)
	assert.Equal(t, projectRoot, wizard.projectRoot)
}

func TestNewAdoptProjectWizard(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	
	wizard := NewAdoptProjectWizard(projectRoot)
	assert.NotNil(t, wizard)
	assert.NotNil(t, wizard.model)
	assert.Equal(t, projectRoot, wizard.projectRoot)
}

func TestNewProjectModel_Initialization(t *testing.T) {
	model := newNewProjectModel()
	
	assert.NotNil(t, model)
	assert.Equal(t, screenWelcome, model.currentScreen)
	assert.NotNil(t, model.textInput)
	assert.NotNil(t, model.templateList)
	assert.NotNil(t, model.ideList)
	assert.NotNil(t, model.spinner)
	assert.Equal(t, 4, len(model.installSteps))
}

func TestAdoptProjectModel_Initialization(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	model := newAdoptProjectModel(projectRoot)
	
	assert.NotNil(t, model)
	assert.Equal(t, adoptScreenFound, model.currentScreen)
	assert.Equal(t, projectRoot, model.projectRoot)
	assert.NotNil(t, model.textInput)
	assert.NotNil(t, model.optionsList)
	assert.NotNil(t, model.ideList)
	assert.NotNil(t, model.spinner)
	assert.Equal(t, 4, len(model.analysisSteps))
}

func TestNewProjectModel_CreateProjectStructure(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	model := newNewProjectModel()
	
	err := model.createProjectStructure(projectRoot)
	require.NoError(t, err)
	
	// Check directories were created
	assert.DirExists(t, filepath.Join(projectRoot, ".doplan"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "contracts"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "templates"))
}

func TestAdoptProjectModel_CreateProjectStructure(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	model := newAdoptProjectModel(projectRoot)
	
	err := model.createProjectStructure()
	require.NoError(t, err)
	
	// Check directories were created
	assert.DirExists(t, filepath.Join(projectRoot, ".doplan"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "contracts"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "templates"))
}

func TestNewProjectModel_SaveConfigYAML(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	model := newNewProjectModel()
	model.projectName = "test-project"
	model.template = "saas"
	model.ide = "cursor"
	model.githubRepo = "https://github.com/user/repo"
	
	// Create .doplan directory
	err := os.MkdirAll(filepath.Join(projectRoot, ".doplan"), 0755)
	require.NoError(t, err)
	
	// Create a config
	cfg := config.NewConfig("cursor")
	
	err = model.saveConfigYAML(projectRoot, cfg, model.githubRepo)
	require.NoError(t, err)
	
	// Check that config file was created
	configPath := filepath.Join(projectRoot, ".doplan", "config.yaml")
	assert.FileExists(t, configPath)
}

func TestAdoptProjectModel_SaveConfigYAML(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	model := newAdoptProjectModel(projectRoot)
	model.ide = "cursor"
	model.githubRepo = "https://github.com/user/repo"
	
	// Create .doplan directory
	err := os.MkdirAll(filepath.Join(projectRoot, ".doplan"), 0755)
	require.NoError(t, err)
	
	// Create a config
	cfg := config.NewConfig("cursor")
	
	err = model.saveConfigYAML(cfg, model.githubRepo)
	require.NoError(t, err)
	
	// Check config file was created
	configPath := filepath.Join(projectRoot, ".doplan", "config.yaml")
	assert.FileExists(t, configPath)
}

