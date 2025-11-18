package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// ManifestGenerator generates the plan manifest file
type ManifestGenerator struct {
	projectRoot string
	state       *models.State
}

// NewManifestGenerator creates a new manifest generator
func NewManifestGenerator(projectRoot string, state *models.State) *ManifestGenerator {
	return &ManifestGenerator{
		projectRoot: projectRoot,
		state:       state,
	}
}

// Generate creates or updates the plan manifest file
func (g *ManifestGenerator) Generate() error {
	manifestPath := filepath.Join(g.projectRoot, "doplan", "plan", "plan-manifest.md")

	// Ensure plan directory exists
	planDir := filepath.Join(g.projectRoot, "doplan", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		return fmt.Errorf("failed to create plan directory: %w", err)
	}

	// Check if manifest already exists
	if _, err := os.Stat(manifestPath); err == nil {
		// Manifest exists, check if it needs updating
		existingManifest, err := g.loadManifest(manifestPath)
		if err == nil && existingManifest != nil {
			// Check if the structure matches the current state
			if g.manifestMatchesState(existingManifest) {
				// Manifest is up to date, don't overwrite
				return nil
			}
		}
		// Manifest exists but doesn't match, update it
	}

	// Generate manifest content
	content := g.generateManifestContent()

	// Ensure doplan directory exists
	doplanDir := filepath.Join(g.projectRoot, "doplan")
	if err := os.MkdirAll(doplanDir, 0755); err != nil {
		return fmt.Errorf("failed to create doplan directory: %w", err)
	}

	// Write manifest file
	return os.WriteFile(manifestPath, []byte(content), 0644)
}

// loadManifest loads an existing manifest file (for comparison)
func (g *ManifestGenerator) loadManifest(path string) (*ManifestStructure, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Simple parser to extract phase/feature names from manifest
	manifest := &ManifestStructure{
		Phases: make([]ManifestPhase, 0),
	}

	lines := strings.Split(string(content), "\n")
	var currentPhase *ManifestPhase

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Phase pattern: "## Phase 01: {name} (`doplan/01-{slug}/`)"
		if strings.HasPrefix(line, "## Phase ") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				phaseName := strings.TrimSpace(parts[1])
				// Extract slug from backticks
				slugStart := strings.Index(phaseName, "`doplan/")
				if slugStart >= 0 {
					phaseName = strings.TrimSpace(phaseName[:slugStart])
					slugPart := phaseName[slugStart+8:]
					slugEnd := strings.Index(slugPart, "/`")
					if slugEnd >= 0 {
						slug := slugPart[:slugEnd]
						// Remove number prefix
						slugParts := strings.SplitN(slug, "-", 2)
						if len(slugParts) == 2 {
							currentPhase = &ManifestPhase{
								Name: phaseName,
								Slug: slugParts[1],
							}
							manifest.Phases = append(manifest.Phases, *currentPhase)
						}
					}
				}
			}
		}

		// Feature pattern: "- Feature 01: {name} (`01-{slug}/`)"
		if currentPhase != nil && strings.HasPrefix(line, "- Feature ") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				featureName := strings.TrimSpace(parts[1])
				// Extract slug from backticks
				slugStart := strings.Index(featureName, "`")
				if slugStart >= 0 {
					featureName = strings.TrimSpace(featureName[:slugStart])
					slugPart := featureName[slugStart+1:]
					slugEnd := strings.Index(slugPart, "/`")
					if slugEnd >= 0 {
						slug := slugPart[:slugEnd]
						// Remove number prefix
						slugParts := strings.SplitN(slug, "-", 2)
						if len(slugParts) == 2 {
							currentPhase.Features = append(currentPhase.Features, ManifestFeature{
								Name: featureName,
								Slug: slugParts[1],
							})
						}
					}
				}
			}
		}
	}

	return manifest, nil
}

// manifestMatchesState checks if manifest matches current state
func (g *ManifestGenerator) manifestMatchesState(manifest *ManifestStructure) bool {
	if len(manifest.Phases) != len(g.state.Phases) {
		return false
	}

	for i, phase := range g.state.Phases {
		if i >= len(manifest.Phases) {
			return false
		}

		manifestPhase := manifest.Phases[i]
		expectedSlug := utils.Slugify(phase.Name)

		if manifestPhase.Slug != expectedSlug {
			return false
		}

		// Check features
		if len(manifestPhase.Features) != len(phase.Features) {
			return false
		}

		for j, featureID := range phase.Features {
			feature := g.findFeature(featureID)
			if feature == nil {
				continue
			}

			if j >= len(manifestPhase.Features) {
				return false
			}

			manifestFeature := manifestPhase.Features[j]
			expectedSlug := utils.Slugify(feature.Name)

			if manifestFeature.Slug != expectedSlug {
				return false
			}
		}
	}

	return true
}

// findFeature finds a feature by ID
func (g *ManifestGenerator) findFeature(featureID string) *models.Feature {
	for _, feature := range g.state.Features {
		if feature.ID == featureID {
			return &feature
		}
	}
	return nil
}

// generateManifestContent generates the manifest markdown content
func (g *ManifestGenerator) generateManifestContent() string {
	var content strings.Builder

	content.WriteString("# Project Plan Manifest\n\n")
	content.WriteString("**This file defines the exact folder structure for phases and features.**\n\n")
	content.WriteString("⚠️ **IMPORTANT:** This manifest is the source of truth for folder names. ")
	content.WriteString("If you need to regenerate the plan structure, use this manifest to ensure consistency.\n\n")
	content.WriteString("---\n\n")

	for i, phase := range g.state.Phases {
		phaseNum := fmt.Sprintf("%02d", i+1)
		phaseSlug := utils.Slugify(phase.Name)
		phaseDir := fmt.Sprintf("plan/01-phases/%02d-%s", i+1, phaseSlug)

		content.WriteString(fmt.Sprintf("## Phase %s: %s (`doplan/%s/`)\n\n", phaseNum, phase.Name, phaseDir))
		content.WriteString(fmt.Sprintf("**Description:** %s\n\n", phase.Description))

		if len(phase.Objectives) > 0 {
			content.WriteString("**Objectives:**\n")
			for _, obj := range phase.Objectives {
				content.WriteString(fmt.Sprintf("- %s\n", obj))
			}
			content.WriteString("\n")
		}

		content.WriteString("**Features:**\n\n")

		for j, featureID := range phase.Features {
			feature := g.findFeature(featureID)
			if feature == nil {
				continue
			}

			featureNum := fmt.Sprintf("%02d", j+1)
			featureSlug := utils.Slugify(feature.Name)
			featureDir := fmt.Sprintf("%02d-%s", j+1, featureSlug)

			content.WriteString(fmt.Sprintf("- Feature %s: %s (`%s/`)\n", featureNum, feature.Name, featureDir))
			content.WriteString(fmt.Sprintf("  - Description: %s\n", feature.Description))

			if len(feature.Objectives) > 0 {
				content.WriteString("  - Objectives:\n")
				for _, obj := range feature.Objectives {
					content.WriteString(fmt.Sprintf("    - %s\n", obj))
				}
			}
			content.WriteString("\n")
		}

		content.WriteString("---\n\n")
	}

	content.WriteString("## Folder Structure Summary\n\n")
	content.WriteString("```\n")
	content.WriteString("doplan/\n")
	content.WriteString("└── plan/\n")
	content.WriteString("    └── 01-phases/\n")

	for i, phase := range g.state.Phases {
		phaseSlug := utils.Slugify(phase.Name)
		phaseDir := fmt.Sprintf("%02d-%s", i+1, phaseSlug)

		isLastPhase := i == len(g.state.Phases)-1
		prefix := "    │   ├──"
		if isLastPhase {
			prefix = "    │   └──"
		}
		content.WriteString(fmt.Sprintf("%s %s/\n", prefix, phaseDir))
		if isLastPhase {
			content.WriteString("    │       ├── phase-plan.md\n")
			content.WriteString("    │       ├── phase-progress.json\n")
		} else {
			content.WriteString("    │   │   ├── phase-plan.md\n")
			content.WriteString("    │   │   ├── phase-progress.json\n")
		}

		for j, featureID := range phase.Features {
			feature := g.findFeature(featureID)
			if feature == nil {
				continue
			}

			featureSlug := utils.Slugify(feature.Name)
			featureDir := fmt.Sprintf("%02d-%s", j+1, featureSlug)

			isLastFeature := j == len(phase.Features)-1
			if isLastPhase {
				if isLastFeature {
					prefix = "    │       └──"
					content.WriteString(fmt.Sprintf("%s %s/\n", prefix, featureDir))
					content.WriteString("    │           ├── plan.md\n")
					content.WriteString("    │           ├── design.md\n")
					content.WriteString("    │           ├── tasks.md\n")
					content.WriteString("    │           └── progress.json\n")
				} else {
					prefix = "    │       ├──"
					content.WriteString(fmt.Sprintf("%s %s/\n", prefix, featureDir))
					content.WriteString("    │       │   ├── plan.md\n")
					content.WriteString("    │       │   ├── design.md\n")
					content.WriteString("    │       │   ├── tasks.md\n")
					content.WriteString("    │       │   └── progress.json\n")
				}
			} else {
				if isLastFeature {
					prefix = "    │   │   └──"
					content.WriteString(fmt.Sprintf("%s %s/\n", prefix, featureDir))
					content.WriteString("    │   │       ├── plan.md\n")
					content.WriteString("    │   │       ├── design.md\n")
					content.WriteString("    │   │       ├── tasks.md\n")
					content.WriteString("    │   │       └── progress.json\n")
				} else {
					prefix = "    │   │   ├──"
					content.WriteString(fmt.Sprintf("%s %s/\n", prefix, featureDir))
					content.WriteString("    │   │   │   ├── plan.md\n")
					content.WriteString("    │   │   │   ├── design.md\n")
					content.WriteString("    │   │   │   ├── tasks.md\n")
					content.WriteString("    │   │   │   └── progress.json\n")
				}
			}
		}
	}

	content.WriteString("    └── plan-manifest.md\n")
	content.WriteString("```\n\n")

	content.WriteString("---\n\n")
	content.WriteString("*This manifest was generated automatically. Do not edit manually unless you understand the implications.*\n")

	return content.String()
}

// ManifestStructure represents the parsed manifest structure
type ManifestStructure struct {
	Phases []ManifestPhase
}

// ManifestPhase represents a phase in the manifest
type ManifestPhase struct {
	Name     string
	Slug     string
	Features []ManifestFeature
}

// ManifestFeature represents a feature in the manifest
type ManifestFeature struct {
	Name string
	Slug string
}

// LoadManifest loads the manifest file and returns phase/feature folder names
func LoadManifest(projectRoot string) (map[string]string, map[string]string, error) {
	manifestPath := filepath.Join(projectRoot, "doplan", "plan", "plan-manifest.md")

	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("manifest file not found: %s", manifestPath)
	}

	content, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read manifest: %w", err)
	}

	phaseFolders := make(map[string]string)   // phase ID -> folder name
	featureFolders := make(map[string]string) // feature ID -> folder name

	lines := strings.Split(string(content), "\n")
	var currentPhaseID string

	for i, line := range lines {
		line = strings.TrimSpace(line)

		// Extract phase folder name
		// Pattern: "## Phase 01: {name} (`doplan/01-{slug}/`)"
		if strings.HasPrefix(line, "## Phase ") {
			// Extract phase number
			phaseNumStart := strings.Index(line, "Phase ") + 6
			phaseNumEnd := strings.Index(line[phaseNumStart:], ":")
			if phaseNumEnd > 0 {
				phaseNum := line[phaseNumStart : phaseNumStart+phaseNumEnd]

				// Find the folder name in backticks
				backtickStart := strings.Index(line, "`doplan/")
				if backtickStart >= 0 {
					folderStart := backtickStart + 8
					folderEnd := strings.Index(line[folderStart:], "/`")
					if folderEnd > 0 {
						folderName := line[folderStart : folderStart+folderEnd]

						// Find phase ID from state (we'll need to match by index)
						// For now, store by number
						phaseFolders[phaseNum] = folderName
					}
				}
			}
		}

		// Extract feature folder name
		// Pattern: "- Feature 01: {name} (`01-{slug}/`)"
		if strings.HasPrefix(line, "- Feature ") && currentPhaseID != "" {
			// Extract feature number
			featureNumStart := strings.Index(line, "Feature ") + 8
			featureNumEnd := strings.Index(line[featureNumStart:], ":")
			if featureNumEnd > 0 {
				featureNum := line[featureNumStart : featureNumStart+featureNumEnd]

				// Find the folder name in backticks
				backtickStart := strings.Index(line, "`")
				if backtickStart >= 0 {
					folderStart := backtickStart + 1
					folderEnd := strings.Index(line[folderStart:], "/`")
					if folderEnd > 0 {
						folderName := line[folderStart : folderStart+folderEnd]
						// Store by phase-feature combo
						featureFolders[currentPhaseID+"-"+featureNum] = folderName
					}
				}
			}
		}

		// Reset phase context when we hit a new phase or end
		if strings.HasPrefix(line, "---") && i > 0 {
			// We might be moving to a new phase, but we'll handle it above
		}
	}

	return phaseFolders, featureFolders, nil
}
