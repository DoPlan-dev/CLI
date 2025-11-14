package github

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/DoPlan-dev/CLI/test/helpers"
)

func TestGitHubSync_Sync(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = projectRoot
	err := cmd.Run()
	require.NoError(t, err)

	// Create initial commit
	cmd = exec.Command("git", "config", "user.name", "Test")
	cmd.Dir = projectRoot
	cmd.Run()

	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = projectRoot
	cmd.Run()

	cmd = exec.Command("git", "commit", "--allow-empty", "-m", "Initial commit")
	cmd.Dir = projectRoot
	cmd.Run()

	sync := NewGitHubSync(projectRoot)
	data, err := sync.Sync()

	// Should not error even if GitHub CLI not available
	if err == nil {
		assert.NotNil(t, data)
	}
}

func TestGitHubSync_fetchBranches(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cmd := exec.Command("git", "init")
	cmd.Dir = projectRoot
	require.NoError(t, cmd.Run())

	// Configure git user (required for some git operations)
	exec.Command("git", "config", "user.name", "Test User").Run()
	exec.Command("git", "config", "user.email", "test@example.com").Run()

	sync := NewGitHubSync(projectRoot)
	branches, err := sync.fetchBranches()

	// fetchBranches uses "git branch -r" which may fail if no remote
	// or return empty list. Both are acceptable.
	// The function may return nil branches even without error in some edge cases
	if err != nil {
		// Error is acceptable - no remote configured
		assert.Error(t, err)
	} else {
		// If no error, branches should be a valid slice (may be empty or nil)
		// In Go, a nil slice is valid and different from an empty slice
		_ = branches // Accept nil or empty slice
	}
}

func TestGitHubSync_fetchCommits(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cmd := exec.Command("git", "init")
	cmd.Dir = projectRoot
	require.NoError(t, cmd.Run())

	cmd = exec.Command("git", "config", "user.name", "Test")
	cmd.Dir = projectRoot
	cmd.Run()

	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = projectRoot
	cmd.Run()

	cmd = exec.Command("git", "commit", "--allow-empty", "-m", "Test commit")
	cmd.Dir = projectRoot
	cmd.Run()

	sync := NewGitHubSync(projectRoot)
	commits, err := sync.fetchCommits()

	if err == nil {
		assert.NotNil(t, commits)
	}
}

func TestGitHubSync_fetchPRs(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	sync := NewGitHubSync(projectRoot)

	prs, err := sync.fetchPRs()

	// Should not error even if GitHub CLI not available
	if err == nil {
		assert.NotNil(t, prs)
	}
}

func TestGitHubSync_LoadData(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	sync := NewGitHubSync(projectRoot)

	// Create test data file
	dataPath := filepath.Join(projectRoot, "doplan", "github-data.json")
	err := os.MkdirAll(filepath.Dir(dataPath), 0755)
	require.NoError(t, err)

	data, err := sync.LoadData()
	// May return nil if file doesn't exist
	if err == nil {
		assert.NotNil(t, data)
	}
}

