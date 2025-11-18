package migration

import (
	"os"
	"path/filepath"
	"regexp"
)

// Detector detects old DoPlan structure
type Detector struct {
	projectRoot string
}

// NewDetector creates a new migration detector
func NewDetector(projectRoot string) *Detector {
	return &Detector{
		projectRoot: projectRoot,
	}
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

// DetectOldFolders finds all old-style folders
func (d *Detector) DetectOldFolders() ([]OldFolder, error) {
	var folders []OldFolder
	doplanDir := filepath.Join(d.projectRoot, "doplan")

	phasePattern := regexp.MustCompile(`^(\d+)-phase$`)

	entries, err := os.ReadDir(doplanDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		matches := phasePattern.FindStringSubmatch(entry.Name())
		if len(matches) > 0 {
			phasePath := filepath.Join(doplanDir, entry.Name())
			phaseFolders, err := d.detectPhaseFolders(phasePath, matches[1])
			if err != nil {
				return nil, err
			}
			folders = append(folders, phaseFolders...)
		}
	}

	return folders, nil
}

// detectPhaseFolders detects features within a phase
func (d *Detector) detectPhaseFolders(phasePath, phaseNum string) ([]OldFolder, error) {
	var folders []OldFolder
	featurePattern := regexp.MustCompile(`^(\d+)-Feature$`)

	entries, err := os.ReadDir(phasePath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		matches := featurePattern.FindStringSubmatch(entry.Name())
		if len(matches) > 0 {
			folders = append(folders, OldFolder{
				Type:     "feature",
				OldPath:  filepath.Join(phasePath, entry.Name()),
				OldName:  entry.Name(),
				PhaseNum: phaseNum,
				FeatureNum: matches[1],
			})
		}
	}

	// Add phase folder
	folders = append(folders, OldFolder{
		Type:    "phase",
		OldPath: phasePath,
		OldName: filepath.Base(phasePath),
		PhaseNum: phaseNum,
	})

	return folders, nil
}

// OldFolder represents an old-style folder
type OldFolder struct {
	Type       string
	OldPath    string
	OldName    string
	PhaseNum   string
	FeatureNum string
}

