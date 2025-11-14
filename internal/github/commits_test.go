package github

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCommitManager(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)

	cmd := exec.Command("git", "init")
	cmd.Dir = projectRoot
	require.NoError(t, cmd.Run())

	manager, err := NewCommitManager(projectRoot)
	if err == nil {
		assert.NotNil(t, manager)
	}
}

func TestCommitManager_CommitFiles(t *testing.T) {
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

	// Create test file
	testFile := filepath.Join(projectRoot, "test.txt")
	err := os.WriteFile(testFile, []byte("test"), 0644)
	require.NoError(t, err)

	manager, err := NewCommitManager(projectRoot)
	if err == nil {
		err = manager.CommitFiles("Test commit", []string{"test.txt"})
		assert.NoError(t, err)
	}
}

func TestFormatCommitMessage(t *testing.T) {
	tests := []struct {
		commitType  string
		scope       string
		description string
		expected    string
	}{
		{"feat", "api", "add endpoint", "feat (api) : add endpoint"},
		{"fix", "", "bug fix", "fix : bug fix"},
		{"", "", "simple message", "simple message"},
	}

	for _, tt := range tests {
		result := FormatCommitMessage(tt.commitType, tt.scope, tt.description)
		assert.Equal(t, tt.expected, result)
	}
}

func TestCommitManager_PushBranch(t *testing.T) {
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

	// Create initial commit
	testFile := filepath.Join(projectRoot, "test.txt")
	err := os.WriteFile(testFile, []byte("test"), 0644)
	require.NoError(t, err)

	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = projectRoot
	cmd.Run()

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = projectRoot
	cmd.Run()

	manager, err := NewCommitManager(projectRoot)
	if err == nil {
		// PushBranch will fail without a remote, but we can test the error handling
		err = manager.PushBranch("main")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to push branch")
	}
}

func TestCommitManager_AutoCommitAndPush(t *testing.T) {
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

	// Create initial commit
	testFile := filepath.Join(projectRoot, "test.txt")
	err := os.WriteFile(testFile, []byte("test"), 0644)
	require.NoError(t, err)

	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = projectRoot
	cmd.Run()

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = projectRoot
	cmd.Run()

	manager, err := NewCommitManager(projectRoot)
	if err == nil {
		// Create a new file
		newFile := filepath.Join(projectRoot, "new.txt")
		require.NoError(t, os.WriteFile(newFile, []byte("new content"), 0644))

		// AutoCommitAndPush will fail on push without remote, but commit should succeed
		err = manager.AutoCommitAndPush("test commit", []string{"new.txt"})
		// The commit part should succeed, but push will fail
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to push branch")
	}
}
