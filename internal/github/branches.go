package github

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// BranchManager handles Git branch operations
type BranchManager struct {
	repoPath string
	repo     *git.Repository
}

// NewBranchManager creates a new branch manager
func NewBranchManager(repoPath string) (*BranchManager, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}
	
	return &BranchManager{
		repoPath: repoPath,
		repo:     repo,
	}, nil
}

// GenerateBranchName generates a branch name from phase and feature
func GenerateBranchName(phaseID, featureID, featureName string) string {
	// Convert to kebab-case
	name := strings.ToLower(featureName)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	
	// Remove special characters
	var result strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	
	cleanName := result.String()
	return fmt.Sprintf("feature/%s-%s-%s", phaseID, featureID, cleanName)
}

// CreateFeatureBranch creates a new feature branch
func (bm *BranchManager) CreateFeatureBranch(branchName string) error {
	// Check if branch already exists
	branches, err := bm.repo.Branches()
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}
	
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == branchName {
			return fmt.Errorf("branch %s already exists", branchName)
		}
		return nil
	})
	if err != nil {
		return err
	}
	
	// Get HEAD reference
	headRef, err := bm.repo.Head()
	if err != nil {
		return fmt.Errorf("failed to get HEAD: %w", err)
	}
	
	// Create new branch
	branchRef := plumbing.NewBranchReferenceName(branchName)
	ref := plumbing.NewHashReference(branchRef, headRef.Hash())
	
	if err := bm.repo.Storer.SetReference(ref); err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}
	
	// Checkout the branch using git command
	cmd := exec.Command("git", "checkout", branchName)
	cmd.Dir = bm.repoPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to checkout branch: %w", err)
	}
	
	return nil
}

// BranchExists checks if a branch exists
func (bm *BranchManager) BranchExists(branchName string) (bool, error) {
	branches, err := bm.repo.Branches()
	if err != nil {
		return false, err
	}
	
	exists := false
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == branchName {
			exists = true
		}
		return nil
	})
	
	return exists, err
}

// GetCurrentBranch returns the current branch name
func (bm *BranchManager) GetCurrentBranch() (string, error) {
	headRef, err := bm.repo.Head()
	if err != nil {
		return "", err
	}
	
	return headRef.Name().Short(), nil
}

