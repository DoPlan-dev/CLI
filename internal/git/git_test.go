package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGetTaskBranchName(t *testing.T) {
	tests := []struct {
		name     string
		taskID   string
		expected string
	}{
		{"simple task", "1.1", "task/1-1"},
		{"multi-dot task", "5.2.3", "task/5-2-3"},
		{"single number", "1", "task/1"},
		{"empty", "", "task/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTaskBranchName(tt.taskID)
			if result != tt.expected {
				t.Errorf("GetTaskBranchName(%q) = %q, want %q", tt.taskID, result, tt.expected)
			}
		})
	}
}

func TestIsGitRepository(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()

	// Should return false for non-git directory
	if IsGitRepository(tmpDir) {
		t.Error("IsGitRepository should return false for non-git directory")
	}

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Should return true for git directory
	if !IsGitRepository(tmpDir) {
		t.Error("IsGitRepository should return true for git directory")
	}
}

func TestGetCurrentBranch(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Create initial commit to establish branch
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tmpDir
	cmd.Run()

	// Create a file and commit
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = tmpDir
	cmd.Run()

	// Test getting current branch
	branch, err := GetCurrentBranch(tmpDir)
	if err != nil {
		t.Fatalf("GetCurrentBranch failed: %v", err)
	}

	// Should be main or master depending on git version
	if branch != "main" && branch != "master" {
		t.Errorf("Expected branch to be 'main' or 'master', got %q", branch)
	}
}

func TestGetCurrentBranch_NonGitDir(t *testing.T) {
	tmpDir := t.TempDir()

	_, err := GetCurrentBranch(tmpDir)
	if err == nil {
		t.Error("GetCurrentBranch should fail for non-git directory")
	}
}

func TestBranchExists(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Setup git config
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tmpDir
	cmd.Run()

	// Create initial commit
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = tmpDir
	cmd.Run()

	// Test non-existent branch
	exists, err := BranchExists(tmpDir, "nonexistent-branch")
	if err != nil {
		t.Fatalf("BranchExists failed: %v", err)
	}
	if exists {
		t.Error("BranchExists should return false for non-existent branch")
	}

	// Create and test existing branch
	cmd = exec.Command("git", "checkout", "-b", "test-branch")
	cmd.Dir = tmpDir
	cmd.Run()

	exists, err = BranchExists(tmpDir, "test-branch")
	if err != nil {
		t.Fatalf("BranchExists failed: %v", err)
	}
	if !exists {
		t.Error("BranchExists should return true for existing branch")
	}
}

func TestGetDefaultBranch(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Setup git config
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tmpDir
	cmd.Run()

	// Create initial commit
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = tmpDir
	cmd.Run()

	// Test getting default branch
	branch, err := GetDefaultBranch(tmpDir)
	if err != nil {
		t.Fatalf("GetDefaultBranch failed: %v", err)
	}

	// Should be main or master
	if branch != "main" && branch != "master" {
		t.Errorf("Expected default branch to be 'main' or 'master', got %q", branch)
	}
}

func TestGetDefaultBranch_NoCommits(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo without commits
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	_, err := GetDefaultBranch(tmpDir)
	if err == nil {
		t.Error("GetDefaultBranch should fail when no commits exist")
	}
}

func TestIsCleanWorkingTree(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Setup git config
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tmpDir
	cmd.Run()

	// Create initial commit
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = tmpDir
	cmd.Run()

	// Should be clean after commit
	clean, err := IsCleanWorkingTree(tmpDir)
	if err != nil {
		t.Fatalf("IsCleanWorkingTree failed: %v", err)
	}
	if !clean {
		t.Error("IsCleanWorkingTree should return true for clean working tree")
	}

	// Create uncommitted file
	uncommittedFile := filepath.Join(tmpDir, "dirty.txt")
	os.WriteFile(uncommittedFile, []byte("dirty"), 0644)

	// Should be dirty now
	clean, err = IsCleanWorkingTree(tmpDir)
	if err != nil {
		t.Fatalf("IsCleanWorkingTree failed: %v", err)
	}
	if clean {
		t.Error("IsCleanWorkingTree should return false for dirty working tree")
	}
}

func TestCreateAndCheckoutBranch(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Setup git config
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tmpDir
	cmd.Run()

	// Create initial commit
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tmpDir
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = tmpDir
	cmd.Run()

	// Test creating new branch
	err := CreateAndCheckoutBranch(tmpDir, "new-branch")
	if err != nil {
		t.Fatalf("CreateAndCheckoutBranch failed: %v", err)
	}

	// Verify we're on the new branch
	branch, err := GetCurrentBranch(tmpDir)
	if err != nil {
		t.Fatalf("GetCurrentBranch failed: %v", err)
	}
	if branch != "new-branch" {
		t.Errorf("Expected to be on 'new-branch', got %q", branch)
	}

	// Test checking out existing branch
	err = CreateAndCheckoutBranch(tmpDir, "new-branch")
	if err != nil {
		t.Fatalf("CreateAndCheckoutBranch failed on existing branch: %v", err)
	}
}

func TestGetRemoteURL(t *testing.T) {
	tmpDir := t.TempDir()

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Add a remote
	cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/test/repo.git")
	cmd.Dir = tmpDir
	cmd.Run()

	// Test getting remote URL
	url, err := GetRemoteURL(tmpDir, "origin")
	if err != nil {
		t.Fatalf("GetRemoteURL failed: %v", err)
	}
	if url != "https://github.com/test/repo.git" {
		t.Errorf("Expected remote URL 'https://github.com/test/repo.git', got %q", url)
	}

	// Test non-existent remote
	_, err = GetRemoteURL(tmpDir, "nonexistent")
	if err == nil {
		t.Error("GetRemoteURL should fail for non-existent remote")
	}
}

