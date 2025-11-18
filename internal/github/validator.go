package github

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
)

// RequireGitHubRepo checks if GitHub repository is configured and accessible
// Returns error if repository is not configured or not accessible
func RequireGitHubRepo(projectRoot string, cmd string) error {
	cfgMgr := config.NewManager(projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		return doplanerror.NewConfigError("CFG001", "Failed to load configuration").
			WithCause(err).
			WithSuggestion("Run 'doplan install' to set up DoPlan")
	}

	// Check YAML config first
	repo := getRepositoryFromConfig(projectRoot)
	if repo == "" {
		// Fallback to old config format
		if cfg != nil && cfg.GitHub.Enabled {
			// Repository might be in old format, but we need it in new format
			return doplanerror.NewValidationError("VAL001", "GitHub repository is required").
				WithSuggestion("Run 'doplan' and select 'Setup GitHub Repository' from the menu").
				WithDetails(fmt.Sprintf("The '%s' action requires a GitHub repository to be configured", cmd))
		}
		return doplanerror.NewValidationError("VAL001", "GitHub repository is required").
			WithSuggestion("Run 'doplan' and select 'Setup GitHub Repository' from the menu").
			WithDetails(fmt.Sprintf("The '%s' action requires a GitHub repository to be configured", cmd))
	}

	// Validate repository format
	if !isValidGitHubRepo(repo) {
		return doplanerror.NewValidationError("VAL002", "Invalid GitHub repository format").
			WithSuggestion("Repository should be in format: 'user/repo' or 'https://github.com/user/repo'").
			WithDetails(fmt.Sprintf("Current value: %s", repo))
	}

	// Check GitHub CLI access if available
	if err := validateGitHubAccess(repo); err != nil {
		return doplanerror.NewValidationError("VAL003", "Cannot access GitHub repository").
			WithCause(err).
			WithSuggestion("Ensure GitHub CLI (gh) is installed and authenticated, or check repository URL")
	}

	return nil
}

// getRepositoryFromConfig extracts repository from YAML config
func getRepositoryFromConfig(projectRoot string) string {
	configPath := filepath.Join(projectRoot, ".doplan", "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return ""
	}

	content := string(data)
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.Contains(line, "repository:") && i > 0 {
			// Check if previous line contains "github:"
			prevLine := ""
			if i > 0 {
				prevLine = lines[i-1]
			}
			if strings.Contains(prevLine, "github:") || strings.Contains(strings.TrimSpace(prevLine), "github") {
				parts := strings.Split(line, "repository:")
				if len(parts) > 1 {
					repo := strings.TrimSpace(strings.Trim(parts[1], `"`))
					if repo != "" && repo != "null" {
						return repo
					}
				}
			}
		}
	}
	return ""
}

// isValidGitHubRepo validates GitHub repository format
func isValidGitHubRepo(repo string) bool {
	// Check for user/repo format
	userRepoPattern := regexp.MustCompile(`^[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+$`)
	if userRepoPattern.MatchString(repo) {
		return true
	}

	// Check for full URL format
	urlPattern := regexp.MustCompile(`^https?://github\.com/[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+(/)?$`)
	if urlPattern.MatchString(repo) {
		return true
	}

	// Check for SSH format
	sshPattern := regexp.MustCompile(`^git@github\.com:[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+\.git$`)
	return sshPattern.MatchString(repo)
}

// validateGitHubAccess checks if GitHub CLI can access the repository
func validateGitHubAccess(repo string) error {
	// Extract user/repo from various formats
	userRepo := extractUserRepo(repo)
	if userRepo == "" {
		return fmt.Errorf("invalid repository format")
	}

	// Check if gh CLI is available
	if _, err := exec.LookPath("gh"); err != nil {
		// GitHub CLI not installed - this is OK, just skip validation
		return nil
	}

	// Try to check repository access
	cmd := exec.Command("gh", "repo", "view", userRepo, "--json", "name")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cannot access repository: %w", err)
	}

	return nil
}

// extractUserRepo extracts user/repo from various GitHub URL formats
func extractUserRepo(repo string) string {
	// Already in user/repo format
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+$`, repo); matched {
		return repo
	}

	// Extract from URL
	urlPattern := regexp.MustCompile(`github\.com/([a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+)`)
	if matches := urlPattern.FindStringSubmatch(repo); len(matches) > 1 {
		return matches[1]
	}

	// Extract from SSH
	sshPattern := regexp.MustCompile(`git@github\.com:([a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+)`)
	if matches := sshPattern.FindStringSubmatch(repo); len(matches) > 1 {
		return strings.TrimSuffix(matches[1], ".git")
	}

	return ""
}

// ActionsRequiringGitHub lists actions that require GitHub repository
var ActionsRequiringGitHub = []string{
	"discuss",
	"plan",
	"implement",
	"feature",
	"progress",
	"deploy",
}
