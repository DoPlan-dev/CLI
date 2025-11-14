package github

import (
	"fmt"
	"os/exec"
	"strings"
)

// PRManager handles pull request operations
type PRManager struct {
	repoPath string
}

// NewPRManager creates a new PR manager
func NewPRManager(repoPath string) *PRManager {
	return &PRManager{
		repoPath: repoPath,
	}
}

// CreatePullRequest creates a pull request using GitHub CLI
func (prm *PRManager) CreatePullRequest(branchName, title, body, baseBranch string) (string, error) {
	// Use GitHub CLI to create PR
	cmd := exec.Command("gh", "pr", "create",
		"--title", title,
		"--body", body,
		"--base", baseBranch,
		"--head", branchName,
	)
	cmd.Dir = prm.repoPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to create PR: %s, output: %s", err, string(output))
	}

	// Extract PR URL from output
	outputStr := string(output)
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		if strings.Contains(line, "https://github.com") {
			return strings.TrimSpace(line), nil
		}
	}

	return outputStr, nil
}

// GeneratePRBody generates PR body from feature information
func GeneratePRBody(featureName, planPath, designPath, tasksPath string) string {
	var body strings.Builder

	body.WriteString(fmt.Sprintf("## Feature: %s\n\n", featureName))
	body.WriteString("This PR implements the feature as planned.\n\n")

	body.WriteString("### Planning Documents\n\n")
	if planPath != "" {
		body.WriteString(fmt.Sprintf("- [Plan](%s)\n", planPath))
	}
	if designPath != "" {
		body.WriteString(fmt.Sprintf("- [Design](%s)\n", designPath))
	}
	if tasksPath != "" {
		body.WriteString(fmt.Sprintf("- [Tasks](%s)\n", tasksPath))
	}

	body.WriteString("\n### Checklist\n\n")
	body.WriteString("- [ ] Code implemented\n")
	body.WriteString("- [ ] Tests written\n")
	body.WriteString("- [ ] Documentation updated\n")
	body.WriteString("- [ ] Ready for review\n")

	return body.String()
}

// CheckFeatureCompletion checks if a feature is complete
func CheckFeatureCompletion(tasksPath string) (bool, error) {
	// Read tasks.md file and check if all tasks are completed
	// This is a simplified version - in real implementation, parse the markdown
	return false, fmt.Errorf("not implemented")
}
