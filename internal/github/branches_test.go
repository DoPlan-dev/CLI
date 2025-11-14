package github

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateBranchName(t *testing.T) {
	tests := []struct {
		name        string
		phaseID     string
		featureID   string
		featureName string
		expected    string
	}{
		{
			name:        "simple name",
			phaseID:     "01-phase",
			featureID:   "01-feature",
			featureName: "User Authentication",
			expected:    "feature/01-phase-01-feature-user-authentication",
		},
		{
			name:        "with underscores",
			phaseID:     "02-phase",
			featureID:   "02-feature",
			featureName: "user_auth",
			expected:    "feature/02-phase-02-feature-user-auth",
		},
		{
			name:        "with special characters",
			phaseID:     "03-phase",
			featureID:   "03-feature",
			featureName: "User@Auth#123",
			expected:    "feature/03-phase-03-feature-userauth123",
		},
		{
			name:        "all lowercase",
			phaseID:     "01-phase",
			featureID:   "01-feature",
			featureName: "test feature",
			expected:    "feature/01-phase-01-feature-test-feature",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateBranchName(tt.phaseID, tt.featureID, tt.featureName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewBranchManager(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Initialize git repo
	_, err := git.PlainInit(projectRoot, false)
	require.NoError(t, err)

	mgr, err := NewBranchManager(projectRoot)
	require.NoError(t, err)
	assert.NotNil(t, mgr)
	assert.Equal(t, projectRoot, mgr.repoPath)
	assert.NotNil(t, mgr.repo)
}

func TestNewBranchManager_NoRepo(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	_, err := NewBranchManager(projectRoot)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open repository")
}

func TestBranchManager_CreateFeatureBranch(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Initialize git repo
	repo, err := git.PlainInit(projectRoot, false)
	require.NoError(t, err)

	// Create initial commit
	w, err := repo.Worktree()
	require.NoError(t, err)

	readmePath := filepath.Join(projectRoot, "README.md")
	require.NoError(t, os.WriteFile(readmePath, []byte("# Test\n"), 0644))

	_, err = w.Add("README.md")
	require.NoError(t, err)

	_, err = w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@example.com",
		},
	})
	require.NoError(t, err)

	mgr, err := NewBranchManager(projectRoot)
	require.NoError(t, err)

	branchName := "feature/test-branch"
	err = mgr.CreateFeatureBranch(branchName)
	require.NoError(t, err)

	// Verify branch exists
	branches, err := repo.Branches()
	require.NoError(t, err)

	found := false
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == branchName {
			found = true
		}
		return nil
	})
	require.NoError(t, err)
	assert.True(t, found, "Branch should exist")
}

func TestBranchManager_CreateFeatureBranch_AlreadyExists(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Initialize git repo
	repo, err := git.PlainInit(projectRoot, false)
	require.NoError(t, err)

	// Create initial commit
	w, err := repo.Worktree()
	require.NoError(t, err)

	readmePath := filepath.Join(projectRoot, "README.md")
	require.NoError(t, os.WriteFile(readmePath, []byte("# Test\n"), 0644))

	_, err = w.Add("README.md")
	require.NoError(t, err)

	_, err = w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@example.com",
		},
	})
	require.NoError(t, err)

	mgr, err := NewBranchManager(projectRoot)
	require.NoError(t, err)

	branchName := "feature/test-branch"
	err = mgr.CreateFeatureBranch(branchName)
	require.NoError(t, err)

	// Try to create again
	err = mgr.CreateFeatureBranch(branchName)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

func TestBranchManager_BranchExists(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Initialize git repo
	repo, err := git.PlainInit(projectRoot, false)
	require.NoError(t, err)

	// Create initial commit
	w, err := repo.Worktree()
	require.NoError(t, err)

	readmePath := filepath.Join(projectRoot, "README.md")
	require.NoError(t, os.WriteFile(readmePath, []byte("# Test\n"), 0644))

	_, err = w.Add("README.md")
	require.NoError(t, err)

	_, err = w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@example.com",
		},
	})
	require.NoError(t, err)

	mgr, err := NewBranchManager(projectRoot)
	require.NoError(t, err)

	// Create a branch
	branchName := "feature/test-branch"
	err = mgr.CreateFeatureBranch(branchName)
	require.NoError(t, err)

	// Check if branch exists
	exists, err := mgr.BranchExists(branchName)
	require.NoError(t, err)
	assert.True(t, exists)

	// Check non-existent branch
	exists, err = mgr.BranchExists("feature/non-existent")
	require.NoError(t, err)
	assert.False(t, exists)
}

func TestBranchManager_GetCurrentBranch(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	// Initialize git repo
	repo, err := git.PlainInit(projectRoot, false)
	require.NoError(t, err)

	// Create initial commit
	w, err := repo.Worktree()
	require.NoError(t, err)

	readmePath := filepath.Join(projectRoot, "README.md")
	require.NoError(t, os.WriteFile(readmePath, []byte("# Test\n"), 0644))

	_, err = w.Add("README.md")
	require.NoError(t, err)

	_, err = w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@example.com",
		},
	})
	require.NoError(t, err)

	mgr, err := NewBranchManager(projectRoot)
	require.NoError(t, err)

	// Get current branch (should be main or master)
	branch, err := mgr.GetCurrentBranch()
	require.NoError(t, err)
	assert.NotEmpty(t, branch)
	assert.Contains(t, []string{"main", "master"}, branch)
}
