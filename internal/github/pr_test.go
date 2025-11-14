package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DoPlan-dev/CLI/test/helpers"
)

func TestNewPRManager(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	mgr := NewPRManager(projectRoot)
	assert.NotNil(t, mgr)
	assert.Equal(t, projectRoot, mgr.repoPath)
}

func TestGeneratePRBody(t *testing.T) {
	body := GeneratePRBody(
		"Test Feature",
		"doplan/01-phase/01-Feature/plan.md",
		"doplan/01-phase/01-Feature/design.md",
		"doplan/01-phase/01-Feature/tasks.md",
	)

	assert.Contains(t, body, "Test Feature")
	assert.Contains(t, body, "plan.md")
	assert.Contains(t, body, "design.md")
	assert.Contains(t, body, "tasks.md")
	assert.Contains(t, body, "Code implemented")
	assert.Contains(t, body, "Tests written")
}

func TestGeneratePRBody_NoPaths(t *testing.T) {
	body := GeneratePRBody("Test Feature", "", "", "")

	assert.Contains(t, body, "Test Feature")
	assert.NotContains(t, body, "plan.md")
}

func TestCheckFeatureCompletion(t *testing.T) {
	// This function is not implemented, so we just test it returns an error
	_, err := CheckFeatureCompletion("path/to/tasks.md")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestPRManager_CreatePullRequest(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	mgr := NewPRManager(projectRoot)

	// CreatePullRequest requires GitHub CLI, which may not be available
	// Test that it handles the error gracefully
	_, err := mgr.CreatePullRequest("feature/test", "Test PR", "Test body", "main")
	// Should error because GitHub CLI is not available or not authenticated
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create PR")
}

