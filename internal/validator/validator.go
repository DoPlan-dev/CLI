package validator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/DoPlan-dev/CLI/internal/config"
)

// Issue represents a validation issue
type Issue struct {
	Level   string // "error", "warning", "info"
	Type    string // "missing_file", "invalid_structure", "inconsistent_data"
	Message string
	Path    string
	Fix     string // Suggested fix
}

// Validator validates project structure and configuration
type Validator struct {
	projectRoot string
	issues      []Issue
}

// NewValidator creates a new validator
func NewValidator(projectRoot string) *Validator {
	return &Validator{
		projectRoot: projectRoot,
		issues:      []Issue{},
	}
}

// Validate performs all validation checks
func (v *Validator) Validate() ([]Issue, error) {
	v.issues = []Issue{}

	// Validate installation
	v.validateInstallation()

	// Validate structure
	v.validateStructure()

	// Validate state consistency
	v.validateStateConsistency()

	// Validate GitHub integration
	v.validateGitHub()

	return v.issues, nil
}

func (v *Validator) validateInstallation() {
	// Check config file
	configPath := filepath.Join(v.projectRoot, ".cursor", "config", "doplan-config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		v.addIssue("error", "missing_file", "Configuration file not found", configPath, "Run 'doplan install'")
	}

	// Check state file
	statePath := filepath.Join(v.projectRoot, ".cursor", "config", "doplan-state.json")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		v.addIssue("warning", "missing_file", "State file not found", statePath, "State will be created on first use")
	}
}

func (v *Validator) validateStructure() {
	doplanDir := filepath.Join(v.projectRoot, "doplan")

	// Check doplan directory
	if _, err := os.Stat(doplanDir); os.IsNotExist(err) {
		v.addIssue("error", "missing_directory", "doplan directory not found", doplanDir, "Run 'doplan install'")
		return
	}

	// Validate phase structure
	v.validatePhases(doplanDir)
}

func (v *Validator) validatePhases(doplanDir string) {
	entries, err := os.ReadDir(doplanDir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), "-phase") {
			continue
		}

		phaseDir := filepath.Join(doplanDir, entry.Name())

		// Check phase-plan.md
		planPath := filepath.Join(phaseDir, "phase-plan.md")
		if _, err := os.Stat(planPath); os.IsNotExist(err) {
			v.addIssue("warning", "missing_file", fmt.Sprintf("phase-plan.md missing for %s", entry.Name()), planPath, "Create phase-plan.md")
		}

		// Check phase-progress.json
		progressPath := filepath.Join(phaseDir, "phase-progress.json")
		if _, err := os.Stat(progressPath); os.IsNotExist(err) {
			v.addIssue("warning", "missing_file", fmt.Sprintf("phase-progress.json missing for %s", entry.Name()), progressPath, "Run 'doplan progress'")
		}

		// Validate features
		v.validateFeatures(phaseDir)
	}
}

func (v *Validator) validateFeatures(phaseDir string) {
	entries, err := os.ReadDir(phaseDir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if !strings.Contains(entry.Name(), "Feature") {
			continue
		}

		featureDir := filepath.Join(phaseDir, entry.Name())

		// Check required files
		requiredFiles := []string{"plan.md", "design.md", "tasks.md", "progress.json"}
		for _, file := range requiredFiles {
			filePath := filepath.Join(featureDir, file)
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				v.addIssue("warning", "missing_file", fmt.Sprintf("%s missing for feature %s", file, entry.Name()), filePath, fmt.Sprintf("Create %s", file))
			}
		}
	}
}

func (v *Validator) validateStateConsistency() {
	cfgMgr := config.NewManager(v.projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		v.addIssue("error", "invalid_structure", "Failed to load state", "", "Check state file format")
		return
	}

	// Check if phases in state match directories
	doplanDir := filepath.Join(v.projectRoot, "doplan")
	for _, phase := range state.Phases {
		phaseDir := filepath.Join(doplanDir, phase.ID)
		if _, err := os.Stat(phaseDir); os.IsNotExist(err) {
			v.addIssue("warning", "inconsistent_data", fmt.Sprintf("Phase %s in state but directory missing", phase.ID), phaseDir, "Create directory or remove from state")
		}
	}
}

func (v *Validator) validateGitHub() {
	cfgMgr := config.NewManager(v.projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil || cfg == nil {
		return
	}

	if !cfg.GitHub.Enabled {
		return
	}

	// Check if gh CLI is installed
	if !isCommandAvailable("gh") {
		v.addIssue("error", "missing_dependency", "GitHub CLI (gh) not found", "", "Install GitHub CLI: https://cli.github.com/")
	}

	// Check if git repo is initialized
	gitDir := filepath.Join(v.projectRoot, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		v.addIssue("warning", "missing_directory", "Git repository not initialized", gitDir, "Run 'git init'")
	}
}

func (v *Validator) addIssue(level, issueType, message, path, fix string) {
	v.issues = append(v.issues, Issue{
		Level:   level,
		Type:    issueType,
		Message: message,
		Path:    path,
		Fix:     fix,
	})
}

func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// AutoFix attempts to fix issues automatically
func (v *Validator) AutoFix(issues []Issue) error {
	fixed := 0
	for _, issue := range issues {
		if issue.Level == "error" {
			if err := v.fixIssue(issue); err != nil {
				color.Yellow("⚠️  Could not auto-fix: %s\n", issue.Message)
			} else {
				color.Green("✅ Fixed: %s\n", issue.Message)
				fixed++
			}
		}
	}
	
	if fixed > 0 {
		color.Green("\n✅ Auto-fixed %d issue(s)\n", fixed)
	}
	
	return nil
}

func (v *Validator) fixIssue(issue Issue) error {
	switch issue.Type {
	case "missing_directory":
		return os.MkdirAll(issue.Path, 0755)
	case "missing_file":
		// Create empty file or use template
		if strings.HasSuffix(issue.Path, ".json") {
			return os.WriteFile(issue.Path, []byte("{}"), 0644)
		}
		return os.WriteFile(issue.Path, []byte(""), 0644)
	}
	return fmt.Errorf("cannot auto-fix issue type: %s", issue.Type)
}

