package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetCurrentBranch returns the current Git branch name
func GetCurrentBranch(projectPath string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = projectPath
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// BranchExists checks if a branch exists locally or remotely
func BranchExists(projectPath, branchName string) (bool, error) {
	// Check local branches
	cmd := exec.Command("git", "branch", "--list", branchName)
	cmd.Dir = projectPath
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to check local branch: %w", err)
	}
	if strings.TrimSpace(string(output)) != "" {
		return true, nil
	}

	// Check remote branches
	cmd = exec.Command("git", "branch", "-r", "--list", fmt.Sprintf("origin/%s", branchName))
	cmd.Dir = projectPath
	output, err = cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to check remote branch: %w", err)
	}
	return strings.TrimSpace(string(output)) != "", nil
}

// CreateAndCheckoutBranch creates a new branch and checks it out
// If the branch already exists, it just checks it out
func CreateAndCheckoutBranch(projectPath, branchName string) error {
	exists, err := BranchExists(projectPath, branchName)
	if err != nil {
		return err
	}

	if exists {
		// Branch exists, just checkout
		cmd := exec.Command("git", "checkout", branchName)
		cmd.Dir = projectPath
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to checkout existing branch %s: %w", branchName, err)
		}
		return nil
	}

	// Create and checkout new branch
	cmd := exec.Command("git", "checkout", "-b", branchName)
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create and checkout branch %s: %w", branchName, err)
	}
	return nil
}

// GetTaskBranchName generates a branch name for a task ID
// Format: task/TASK-ID (e.g., task/5.2, task/1.3)
func GetTaskBranchName(taskID string) string {
	// Normalize task ID (remove dots, replace with hyphens if needed)
	taskID = strings.ReplaceAll(taskID, ".", "-")
	return fmt.Sprintf("task/%s", taskID)
}

// GetDefaultBranch returns the default branch (main or master)
func GetDefaultBranch(projectPath string) (string, error) {
	// Try main first
	cmd := exec.Command("git", "rev-parse", "--verify", "main")
	cmd.Dir = projectPath
	if err := cmd.Run(); err == nil {
		return "main", nil
	}

	// Try master
	cmd = exec.Command("git", "rev-parse", "--verify", "master")
	cmd.Dir = projectPath
	if err := cmd.Run(); err == nil {
		return "master", nil
	}

	return "", fmt.Errorf("no default branch found (tried main and master)")
}

// IsCleanWorkingTree checks if the working tree is clean (no uncommitted changes)
func IsCleanWorkingTree(projectPath string) (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = projectPath
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to check git status: %w", err)
	}
	return strings.TrimSpace(string(output)) == "", nil
}

// PushBranch pushes the current branch to remote
func PushBranch(projectPath, branchName string, setUpstream bool) error {
	args := []string{"push"}
	if setUpstream {
		args = append(args, "-u", "origin", branchName)
	} else {
		args = append(args, "origin", branchName)
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push branch %s: %w", branchName, err)
	}
	return nil
}

// GetRemoteURL returns the remote URL for the given remote name (default: origin)
func GetRemoteURL(projectPath, remoteName string) (string, error) {
	cmd := exec.Command("git", "remote", "get-url", remoteName)
	cmd.Dir = projectPath
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get remote URL: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// IsGitRepository checks if the given path is a Git repository
func IsGitRepository(projectPath string) bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Dir = projectPath
	return cmd.Run() == nil
}
