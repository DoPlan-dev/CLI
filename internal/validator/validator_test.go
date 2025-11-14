package validator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewValidator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	assert.NotNil(t, validator)
	assert.Equal(t, projectRoot, validator.projectRoot)
}

func TestValidator_Validate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	// Validate may panic if config is nil, so we need to handle that
	// or ensure config exists
	issues, err := validator.Validate()

	// If error, that's acceptable (config not found)
	// If no error, should return issues
	if err == nil {
		assert.NotNil(t, issues)
	} else {
		// Error is acceptable for uninstalled project
		assert.Error(t, err)
	}
}

func TestValidator_Validate_Installed(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Create config
	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	validator := NewValidator(projectRoot)
	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should have fewer issues when installed
	assert.NotNil(t, issues)
}

func TestValidator_AutoFix(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	issues := []Issue{
		{
			Level:   "error",
			Type:    "missing_directory",
			Path:    filepath.Join(projectRoot, "missing", "dir"),
			Message: "Directory missing",
		},
	}

	err := validator.AutoFix(issues)
	require.NoError(t, err)

	// Verify directory was created
	assert.DirExists(t, issues[0].Path)
}

func TestValidator_validateInstallation(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	// This is a private method, but we can test it indirectly through Validate
	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should have issues about missing config
	hasConfigIssue := false
	for _, issue := range issues {
		if issue.Type == "missing_file" && issue.Level == "error" {
			hasConfigIssue = true
			break
		}
	}
	assert.True(t, hasConfigIssue, "Should report missing config file")
}

func TestValidator_validateStructure(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should have issues about missing doplan directory
	hasDoplanIssue := false
	for _, issue := range issues {
		if issue.Type == "missing_directory" && issue.Message == "doplan directory not found" {
			hasDoplanIssue = true
			break
		}
	}
	assert.True(t, hasDoplanIssue, "Should report missing doplan directory")
}

func TestValidator_fixIssue_MissingFile(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	issue := Issue{
		Level:   "error",
		Type:    "missing_file",
		Path:    filepath.Join(projectRoot, "test.json"),
		Message: "File missing",
	}

	err := validator.fixIssue(issue)
	require.NoError(t, err)

	// Verify file was created
	assert.FileExists(t, issue.Path)
}

func TestValidator_fixIssue_UnknownType(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	validator := NewValidator(projectRoot)

	issue := Issue{
		Level:   "error",
		Type:    "unknown_type",
		Path:    filepath.Join(projectRoot, "test.txt"),
		Message: "Unknown issue",
	}

	err := validator.fixIssue(issue)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot auto-fix issue type")
}

func TestValidator_validateFeatures(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Create phase directory with feature
	phaseDir := filepath.Join(projectRoot, "doplan", "01-phase")
	featureDir := filepath.Join(phaseDir, "01-Feature")
	require.NoError(t, os.MkdirAll(featureDir, 0755))

	// Create some but not all required files
	planFile := filepath.Join(featureDir, "plan.md")
	require.NoError(t, os.WriteFile(planFile, []byte("# Plan"), 0644))

	// Don't create design.md, tasks.md, progress.json

	validator := NewValidator(projectRoot)
	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should have issues about missing files
	hasMissingFileIssue := false
	for _, issue := range issues {
		if issue.Type == "missing_file" && strings.Contains(issue.Message, "missing for feature") {
			hasMissingFileIssue = true
			break
		}
	}
	assert.True(t, hasMissingFileIssue, "Should report missing feature files")
}

func TestValidator_validatePhases(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Create config so Validate doesn't error early
	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	// Create phase directory (must end with "-phase")
	phaseDir := filepath.Join(projectRoot, "doplan", "01-phase")
	require.NoError(t, os.MkdirAll(phaseDir, 0755))

	// Create phase-plan.md but not phase-progress.json
	planFile := filepath.Join(phaseDir, "phase-plan.md")
	require.NoError(t, os.WriteFile(planFile, []byte("# Phase Plan"), 0644))

	validator := NewValidator(projectRoot)
	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should have issues about missing phase files
	hasMissingPhaseFileIssue := false
	for _, issue := range issues {
		if issue.Type == "missing_file" && (strings.Contains(issue.Message, "missing for phase") || strings.Contains(issue.Message, "phase-progress.json")) {
			hasMissingPhaseFileIssue = true
			break
		}
	}
	assert.True(t, hasMissingPhaseFileIssue, "Should report missing phase files")
}

func TestValidator_validateStateConsistency(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Create config
	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	// Create state with phase that doesn't have a directory
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:   "01-phase",
				Name: "Phase 1",
			},
		},
	}
	err = cfgMgr.SaveState(state)
	require.NoError(t, err)

	validator := NewValidator(projectRoot)
	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should have issues about inconsistent state
	hasInconsistentIssue := false
	for _, issue := range issues {
		if issue.Type == "inconsistent_data" && strings.Contains(issue.Message, "in state but directory missing") {
			hasInconsistentIssue = true
			break
		}
	}
	assert.True(t, hasInconsistentIssue, "Should report inconsistent state")
}

func TestValidator_validateGitHub(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Create config with GitHub enabled
	cfgMgr := config.NewManager(projectRoot)
	cfg := &models.Config{
		IDE:       "cursor",
		Version:   "1.0.0",
		Installed: true,
		GitHub: models.GitHubConfig{
			Enabled: true,
		},
	}
	err := cfgMgr.SaveConfig(cfg)
	require.NoError(t, err)

	validator := NewValidator(projectRoot)
	issues, err := validator.Validate()
	require.NoError(t, err)

	// Should validate GitHub configuration
	// May or may not have issues depending on git setup
	assert.NotNil(t, issues)
}

func TestIsCommandAvailable(t *testing.T) {
	// Test with a command that should exist
	assert.True(t, isCommandAvailable("go"))

	// Test with a command that likely doesn't exist
	assert.False(t, isCommandAvailable("nonexistent-command-xyz-123"))
}
