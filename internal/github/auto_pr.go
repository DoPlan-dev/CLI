package github

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// AutoPRManager handles automatic PR creation
type AutoPRManager struct {
	repoPath string
	config   *models.Config
}

// NewAutoPRManager creates a new auto PR manager
func NewAutoPRManager(repoPath string) *AutoPRManager {
	cfgMgr := config.NewManager(repoPath)
	cfg, _ := cfgMgr.LoadConfig()

	return &AutoPRManager{
		repoPath: repoPath,
		config:   cfg,
	}
}

// CheckAndCreatePR checks if feature is complete and creates PR if needed
func (aprm *AutoPRManager) CheckAndCreatePR(feature *models.Feature) error {
	if !aprm.config.GitHub.Enabled || !aprm.config.GitHub.AutoPR {
		return nil // AutoPR disabled
	}

	if feature.PR != nil && feature.PR.URL != "" {
		return nil // PR already exists
	}

	// Check if feature is complete
	isComplete, err := aprm.isFeatureComplete(feature)
	if err != nil {
		return fmt.Errorf("failed to check feature completion: %w", err)
	}

	if !isComplete {
		return nil // Feature not complete yet
	}

	// Create PR
	return aprm.createPRForFeature(feature)
}

func (aprm *AutoPRManager) isFeatureComplete(feature *models.Feature) (bool, error) {
	// Check progress
	if feature.Progress < 100 {
		return false, nil
	}

	// Check status
	if feature.Status != "complete" {
		return false, nil
	}

	// Check tasks.md file
	featureDir := filepath.Join(aprm.repoPath, "doplan", feature.Phase, feature.ID)
	tasksPath := filepath.Join(featureDir, "tasks.md")

	if _, err := os.Stat(tasksPath); os.IsNotExist(err) {
		return false, nil
	}

	content, err := os.ReadFile(tasksPath)
	if err != nil {
		return false, err
	}

	lines := strings.Split(string(content), "\n")
	totalTasks := 0
	completedTasks := 0

	for _, line := range lines {
		if strings.Contains(line, "- [") {
			totalTasks++
			if strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]") {
				completedTasks++
			}
		}
	}

	return totalTasks > 0 && completedTasks == totalTasks, nil
}

func (aprm *AutoPRManager) createPRForFeature(feature *models.Feature) error {
	if feature.Branch == "" {
		return fmt.Errorf("feature branch not set")
	}

	prManager := NewPRManager(aprm.repoPath)

	// Generate PR title
	title := fmt.Sprintf("Feature: %s", feature.Name)

	// Generate PR body
	planPath := fmt.Sprintf("doplan/%s/%s/plan.md", feature.Phase, feature.ID)
	designPath := fmt.Sprintf("doplan/%s/%s/design.md", feature.Phase, feature.ID)
	tasksPath := fmt.Sprintf("doplan/%s/%s/tasks.md", feature.Phase, feature.ID)

	body := GeneratePRBody(feature.Name, planPath, designPath, tasksPath)

	// Create PR
	prURL, err := prManager.CreatePullRequest(
		feature.Branch,
		title,
		body,
		"main", // base branch
	)

	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	color.Green("✅ Created PR for feature '%s': %s\n", feature.Name, prURL)

	// Update feature with PR info
	feature.PR = &models.PullRequest{
		URL:    prURL,
		Title:  title,
		Status: "open",
	}

	// Save updated state
	cfgMgr := config.NewManager(aprm.repoPath)
	state, err := cfgMgr.LoadState()
	if err == nil {
		// Update feature in state
		for i := range state.Features {
			if state.Features[i].ID == feature.ID {
				state.Features[i].PR = feature.PR
				cfgMgr.SaveState(state)
				break
			}
		}
	}

	return nil
}

// WatchFeatures monitors features and creates PRs when complete
func (aprm *AutoPRManager) WatchFeatures(state *models.State) error {
	created := 0
	for i := range state.Features {
		feature := &state.Features[i]

		if err := aprm.CheckAndCreatePR(feature); err != nil {
			color.Yellow("⚠️  Failed to create PR for feature '%s': %v\n", feature.Name, err)
			continue
		}

		if feature.PR != nil && feature.PR.URL != "" {
			created++
		}
	}

	if created > 0 {
		color.Green("✅ Created %d PR(s) for completed features\n", created)
	}

	return nil
}

