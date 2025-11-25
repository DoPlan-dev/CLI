package generator

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/rules"
	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// RulesGenerator generates the rules library by extracting embedded rules
type RulesGenerator struct{}

// Name returns the name of the generator
func (g *RulesGenerator) Name() string {
	return "Rules Library"
}

// Generate extracts all embedded rules to .cursor/rules/library/ in the project
func (g *RulesGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	rulesDir := filepath.Join(projectPath, ".cursor", "rules", "library")

	// Create .cursor/rules/library directory
	if err := utils.CreateDirectory(rulesDir); err != nil {
		return fmt.Errorf("failed to create .cursor/rules/library directory: %w", err)
	}

	// Extract all embedded rules
	if err := ExtractRules(rulesDir); err != nil {
		return fmt.Errorf("failed to extract rules: %w", err)
	}

	return nil
}

// ExtractRules extracts all embedded rules to the target directory, maintaining structure
func ExtractRules(targetDir string) error {
	// Validate target directory
	if targetDir == "" {
		return fmt.Errorf("target directory cannot be empty")
	}

	// Walk through embedded rules and extract them
	err := rules.WalkDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking rules: %w", err)
		}

		// Remove "library/" prefix from path
		relPath := path
		if len(path) > 8 && path[:8] == "library/" {
			relPath = path[8:]
		}

		targetPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			// Create directory
			if err := utils.CreateDirectory(targetPath); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
		} else {
			// Read file from embedded FS (with automatic decompression if needed)
			content, err := rules.ReadFileDecompressed(relPath)
			if err != nil {
				return fmt.Errorf("failed to read embedded file %s: %w", relPath, err)
			}

			// Write file to target
			if err := utils.WriteFile(targetPath, content); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error extracting rules: %w", err)
	}

	return nil
}

// GenerateRules is a convenience function that creates a RulesGenerator and generates rules
func GenerateRules(request *models.ProjectRequest, projectPath string) error {
	generator := &RulesGenerator{}
	return generator.Generate(request, projectPath)
}
