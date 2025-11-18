package context

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// ProjectState represents the detected state of a project
type ProjectState string

const (
	StateEmptyFolder          ProjectState = "EmptyFolder"
	StateExistingCodeNoDoPlan ProjectState = "ExistingCodeNoDoPlan"
	StateDoPlanInstalled      ProjectState = "DoPlanInstalled"
	StateOldDoPlanStructure   ProjectState = "OldDoPlanStructure"
	StateNewDoPlanStructure   ProjectState = "NewDoPlanStructure"
	StateInsideFeature        ProjectState = "InsideFeature"
	StateInsidePhase          ProjectState = "InsidePhase"
)

// ContextDetails contains detailed context information
type ContextDetails struct {
	State         ProjectState `json:"state"`
	ProjectRoot   string       `json:"projectRoot"`
	CurrentPath   string       `json:"currentPath"`
	PhaseID       string       `json:"phaseId,omitempty"`
	PhaseName     string       `json:"phaseName,omitempty"`
	FeatureID     string       `json:"featureId,omitempty"`
	FeatureName   string       `json:"featureName,omitempty"`
	IsGitRepo     bool         `json:"isGitRepo"`
	GitBranch     string       `json:"gitBranch,omitempty"`
	DashboardPath string       `json:"dashboardPath,omitempty"`
}

// Detector detects the current state of a project
type Detector struct {
	projectRoot string
}

// NewDetector creates a new context detector
func NewDetector(projectRoot string) *Detector {
	return &Detector{
		projectRoot: projectRoot,
	}
}

// DetectProjectState detects the overall project state
func (d *Detector) DetectProjectState() (ProjectState, error) {
	// Check for new structure first
	newConfigPath := filepath.Join(d.projectRoot, ".doplan", "config.yaml")
	if _, err := os.Stat(newConfigPath); err == nil {
		// Check if inside feature/phase
		if state := d.detectContext(); state != "" {
			return state, nil
		}
		return StateNewDoPlanStructure, nil
	}

	// Check for old structure
	hasOld, err := d.DetectOldStructure()
	if err != nil {
		return "", err
	}
	if hasOld {
		return StateOldDoPlanStructure, nil
	}

	// Check for existing code without DoPlan
	if d.hasProjectFiles() {
		return StateExistingCodeNoDoPlan, nil
	}

	return StateEmptyFolder, nil
}

// DetectOldStructure checks if project uses old v0.0.17 structure
func (d *Detector) DetectOldStructure() (bool, error) {
	// Check for old config location
	oldConfigPath := filepath.Join(d.projectRoot, ".cursor", "config", "doplan-config.json")
	if _, err := os.Stat(oldConfigPath); err == nil {
		return true, nil
	}

	// Check for old folder structure
	doplanDir := filepath.Join(d.projectRoot, "doplan")
	if _, err := os.Stat(doplanDir); err != nil {
		return false, nil
	}

	// Check for old phase folder pattern
	oldPhasePattern := regexp.MustCompile(`^\d+-phase$`)
	entries, err := os.ReadDir(doplanDir)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if entry.IsDir() && oldPhasePattern.MatchString(entry.Name()) {
			return true, nil
		}
	}

	return false, nil
}

// hasProjectFiles checks if directory contains project files
func (d *Detector) hasProjectFiles() bool {
	projectFiles := []string{
		"package.json",
		"go.mod",
		"requirements.txt",
		"Cargo.toml",
		"pom.xml",
		"build.gradle",
		"pyproject.toml",
	}

	for _, file := range projectFiles {
		path := filepath.Join(d.projectRoot, file)
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	return false
}

// DetectContextDetails detects the full context with details
func (d *Detector) DetectContextDetails() (*ContextDetails, error) {
	state, err := d.DetectProjectState()
	if err != nil {
		return nil, err
	}

	cwd, _ := os.Getwd()
	details := &ContextDetails{
		State:       state,
		ProjectRoot: d.projectRoot,
		CurrentPath: cwd,
	}

	// Check git repository
	details.IsGitRepo = d.isGitRepository()
	if details.IsGitRepo {
		details.GitBranch = d.getGitBranch()
	}

	// If inside feature or phase, load details from dashboard
	if state == StateInsideFeature || state == StateInsidePhase {
		if err := d.loadContextFromDashboard(details); err != nil {
			// Fallback to pattern matching if dashboard not available
			d.loadContextFromPath(details)
		}
	}

	return details, nil
}

// detectContext checks if we're inside a feature or phase directory
func (d *Detector) detectContext() ProjectState {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Check if we're inside doplan directory
	doplanPath := filepath.Join(d.projectRoot, "doplan")
	if !strings.HasPrefix(cwd, doplanPath) {
		return ""
	}

	// Check for feature pattern (##-slug-name)
	featurePattern := regexp.MustCompile(`^\d+-\w+(-\w+)*$`)
	phasePattern := regexp.MustCompile(`^\d+-\w+(-\w+)*$`)

	relPath, err := filepath.Rel(doplanPath, cwd)
	if err != nil {
		return ""
	}

	parts := strings.Split(relPath, string(filepath.Separator))
	if len(parts) >= 2 {
		// Inside a feature (phase/feature)
		if phasePattern.MatchString(parts[0]) && featurePattern.MatchString(parts[1]) {
			return StateInsideFeature
		}
	} else if len(parts) == 1 && parts[0] != "." {
		// Inside a phase
		if phasePattern.MatchString(parts[0]) {
			return StateInsidePhase
		}
	}

	return ""
}

// isGitRepository checks if the project root is a git repository
func (d *Detector) isGitRepository() bool {
	gitDir := filepath.Join(d.projectRoot, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return true
	}
	return false
}

// getGitBranch gets the current git branch
func (d *Detector) getGitBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = d.projectRoot
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

// loadContextFromDashboard loads context details from dashboard.json
func (d *Detector) loadContextFromDashboard(details *ContextDetails) error {
	dashboardPath := filepath.Join(d.projectRoot, ".doplan", "dashboard.json")
	details.DashboardPath = dashboardPath

	data, err := os.ReadFile(dashboardPath)
	if err != nil {
		return err
	}

	var dashboard struct {
		Phases []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Features []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"features"`
		} `json:"phases"`
	}

	if err := json.Unmarshal(data, &dashboard); err != nil {
		return err
	}

	cwd, _ := os.Getwd()
	relPath, err := filepath.Rel(filepath.Join(d.projectRoot, "doplan"), cwd)
	if err != nil {
		return err
	}

	parts := strings.Split(relPath, string(filepath.Separator))
	if len(parts) >= 2 {
		// Inside a feature
		phaseSlug := parts[0]
		featureSlug := parts[1]

		// Find matching phase and feature
		for _, phase := range dashboard.Phases {
			// Extract slug from phase ID (format: "01-phase-slug")
			if strings.HasSuffix(phase.ID, phaseSlug) || phase.ID == phaseSlug {
				details.PhaseID = phase.ID
				details.PhaseName = phase.Name

				// Find feature
				for _, feature := range phase.Features {
					if strings.HasSuffix(feature.ID, featureSlug) || feature.ID == featureSlug {
						details.FeatureID = feature.ID
						details.FeatureName = feature.Name
						return nil
					}
				}
			}
		}
	} else if len(parts) == 1 && parts[0] != "." {
		// Inside a phase
		phaseSlug := parts[0]
		for _, phase := range dashboard.Phases {
			if strings.HasSuffix(phase.ID, phaseSlug) || phase.ID == phaseSlug {
				details.PhaseID = phase.ID
				details.PhaseName = phase.Name
				return nil
			}
		}
	}

	return nil
}

// loadContextFromPath loads context from path pattern matching (fallback)
func (d *Detector) loadContextFromPath(details *ContextDetails) {
	cwd, _ := os.Getwd()
	relPath, err := filepath.Rel(filepath.Join(d.projectRoot, "doplan"), cwd)
	if err != nil {
		return
	}

	parts := strings.Split(relPath, string(filepath.Separator))
	phasePattern := regexp.MustCompile(`^(\d+)-(.+)$`)
	featurePattern := regexp.MustCompile(`^(\d+)-(.+)$`)

	if len(parts) >= 2 {
		// Inside a feature
		if matches := phasePattern.FindStringSubmatch(parts[0]); len(matches) > 2 {
			details.PhaseID = matches[1]
			details.PhaseName = matches[2]
		}
		if matches := featurePattern.FindStringSubmatch(parts[1]); len(matches) > 2 {
			details.FeatureID = matches[1]
			details.FeatureName = matches[2]
		}
	} else if len(parts) == 1 && parts[0] != "." {
		// Inside a phase
		if matches := phasePattern.FindStringSubmatch(parts[0]); len(matches) > 2 {
			details.PhaseID = matches[1]
			details.PhaseName = matches[2]
		}
	}
}

