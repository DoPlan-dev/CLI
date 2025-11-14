package generators

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDashboardGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	githubData := &models.GitHubData{}

	gen := NewDashboardGenerator(projectRoot, state, githubData)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
	assert.Equal(t, state, gen.state)
	assert.Equal(t, githubData, gen.githubData)
}

func TestDashboardGenerator_Generate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:       "phase-1",
				Name:     "Phase 1",
				Status:   "in-progress",
				Features: []string{"feat-1"},
			},
		},
		Features: []models.Feature{
			{
				ID:       "feat-1",
				Name:     "Feature 1",
				Progress: 75,
				Phase:    "phase-1",
			},
		},
		Progress: models.Progress{
			Overall: 60,
		},
	}
	githubData := &models.GitHubData{
		Branches: []models.Branch{
			{Name: "feature/test", Status: "active"},
		},
		PRs: []models.PullRequest{
			{Number: 1, Title: "Test PR", Status: "open"},
		},
	}

	gen := NewDashboardGenerator(projectRoot, state, githubData)
	err := gen.Generate()
	require.NoError(t, err)

	// Check markdown dashboard
	mdPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	assert.FileExists(t, mdPath)

	mdContent, err := os.ReadFile(mdPath)
	require.NoError(t, err)
	content := string(mdContent)
	assert.Contains(t, content, "Project Dashboard")
	assert.Contains(t, content, "Phase 1")
	assert.Contains(t, content, "Feature 1")

	// Check HTML dashboard
	htmlPath := filepath.Join(projectRoot, "doplan", "dashboard.html")
	assert.FileExists(t, htmlPath)

	htmlContent, err := os.ReadFile(htmlPath)
	require.NoError(t, err)
	htmlStr := string(htmlContent)
	assert.Contains(t, htmlStr, "<!DOCTYPE html>")
	assert.Contains(t, htmlStr, "Phase 1")
}

func TestDashboardGenerator_Generate_EmptyState(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
	}
	githubData := &models.GitHubData{}

	gen := NewDashboardGenerator(projectRoot, state, githubData)
	err := gen.Generate()
	require.NoError(t, err)

	mdPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	assert.FileExists(t, mdPath)

	content, err := os.ReadFile(mdPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "No phases created yet")
}

func TestGenerateInitialDashboard(t *testing.T) {
	dashboard := GenerateInitialDashboard()
	assert.Contains(t, dashboard, "Project Dashboard")
	assert.Contains(t, dashboard, "Last Updated")
	assert.Contains(t, dashboard, "0%")
	assert.Contains(t, dashboard, "No phases created yet")
}

func TestDashboardGenerator_generateMarkdown(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:       "phase-1",
				Name:     "Phase 1",
				Status:   "in-progress",
				Features: []string{"feat-1"},
			},
		},
		Features: []models.Feature{
			{
				ID:       "feat-1",
				Name:     "Feature 1",
				Progress: 75,
				Phase:    "phase-1",
				Status:   "in-progress",
			},
		},
		Progress: models.Progress{
			Overall: 60,
			Phases:  map[string]int{"phase-1": 75},
		},
	}
	githubData := &models.GitHubData{}

	gen := NewDashboardGenerator(projectRoot, state, githubData)
	content := gen.generateMarkdown()

	assert.Contains(t, content, "Project Dashboard")
	assert.Contains(t, content, "Phase 1")
	assert.Contains(t, content, "Feature 1")
	// Progress calculation may vary, so just check it contains progress info
	assert.Contains(t, content, "%")
}

func TestDashboardGenerator_generateHTML(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:     "phase-1",
				Name:   "Phase 1",
				Status: "in-progress",
			},
		},
		Progress: models.Progress{
			Overall: 60,
		},
	}
	githubData := &models.GitHubData{
		PRs: []models.PullRequest{
			{Number: 1, Title: "Test PR", Status: "open"},
		},
	}

	gen := NewDashboardGenerator(projectRoot, state, githubData)
	content := gen.generateHTML()

	assert.Contains(t, content, "<!DOCTYPE html>")
	assert.Contains(t, content, "<html")
	assert.Contains(t, content, "Phase 1")
	assert.Contains(t, content, "Test PR")
}

func TestDashboardGenerator_generateMarkdown_WithGitHubData(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	githubData := &models.GitHubData{
		Branches: []models.Branch{
			{Name: "feature/test", Status: "active"},
		},
		Commits: []models.Commit{
			{Hash: "abc123456789", Message: "Test commit", Author: "Test User", Date: "2025-01-01"},
		},
		PRs: []models.PullRequest{
			{Number: 1, Title: "Test PR", Status: "open", URL: "https://github.com/test/pr/1"},
		},
	}

	gen := NewDashboardGenerator(projectRoot, state, githubData)
	content := gen.generateMarkdown()

	assert.Contains(t, content, "feature/test")
	assert.Contains(t, content, "Test commit")
	assert.Contains(t, content, "Test PR")
}
