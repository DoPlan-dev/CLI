package github

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// CommitManager handles Git commit operations
type CommitManager struct {
	repoPath string
	repo     *git.Repository
}

// NewCommitManager creates a new commit manager
func NewCommitManager(repoPath string) (*CommitManager, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}
	
	return &CommitManager{
		repoPath: repoPath,
		repo:     repo,
	}, nil
}

// CommitFiles commits files with a message
func (cm *CommitManager) CommitFiles(message string, paths []string) error {
	worktree, err := cm.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}
	
	// Add files
	for _, path := range paths {
		_, err := worktree.Add(path)
		if err != nil {
			return fmt.Errorf("failed to add file %s: %w", path, err)
		}
	}
	
	// Commit
	_, err = worktree.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "DoPlan",
			Email: "doplan@example.com",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	
	return nil
}

// PushBranch pushes the current branch to remote
func (cm *CommitManager) PushBranch(branchName string) error {
	cmd := exec.Command("git", "push", "origin", branchName)
	cmd.Dir = cm.repoPath
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to push branch: %s, output: %s", err, string(output))
	}
	
	return nil
}

// AutoCommitAndPush commits and pushes automatically
func (cm *CommitManager) AutoCommitAndPush(message string, paths []string) error {
	// Get current branch
	headRef, err := cm.repo.Head()
	if err != nil {
		return fmt.Errorf("failed to get HEAD: %w", err)
	}
	branchName := headRef.Name().Short()
	
	// Commit
	if err := cm.CommitFiles(message, paths); err != nil {
		return err
	}
	
	// Push
	if err := cm.PushBranch(branchName); err != nil {
		return err
	}
	
	return nil
}

// FormatCommitMessage formats a commit message according to conventional commits
func FormatCommitMessage(commitType, scope, description string) string {
	var parts []string
	
	if commitType != "" {
		parts = append(parts, commitType)
	}
	
	if scope != "" {
		parts = append(parts, fmt.Sprintf("(%s)", scope))
	}
	
	if len(parts) > 0 {
		parts = append(parts, ":")
	}
	
	if description != "" {
		parts = append(parts, description)
	}
	
	return strings.Join(parts, " ")
}

