package context

import (
	"os"
	"path/filepath"
	"regexp"
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

// detectContext checks if we're inside a feature or phase directory
func (d *Detector) detectContext() ProjectState {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Check if we're inside doplan directory
	if !filepath.HasPrefix(cwd, filepath.Join(d.projectRoot, "doplan")) {
		return ""
	}

	// Check for feature pattern (##-slug-name)
	featurePattern := regexp.MustCompile(`\d+-\w+(-\w+)*$`)
	phasePattern := regexp.MustCompile(`\d+-\w+(-\w+)*$`)

	relPath, err := filepath.Rel(filepath.Join(d.projectRoot, "doplan"), cwd)
	if err != nil {
		return ""
	}

	parts := filepath.SplitList(relPath)
	if len(parts) >= 2 {
		// Inside a feature (phase/feature)
		if phasePattern.MatchString(parts[0]) && featurePattern.MatchString(parts[1]) {
			return StateInsideFeature
		}
	} else if len(parts) == 1 {
		// Inside a phase
		if phasePattern.MatchString(parts[0]) {
			return StateInsidePhase
		}
	}

	return ""
}

