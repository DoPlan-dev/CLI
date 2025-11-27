package generator

import (
	"fmt"
	"io/fs"
	"os"
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

// Generate extracts all embedded rules to a central location and creates symlinks in IDE-specific directories
func (g *RulesGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// Ensure IDEs list is populated (for backward compatibility)
	if len(request.IDEs) == 0 && request.IDE != "" {
		request.IDEs = []string{request.IDE}
	}

	// Central location for rules library: .do/core/library/
	centralRulesDir := filepath.Join(projectPath, ".do", "core", "library")

	// Extract rules to central location (only once)
	if err := utils.CreateDirectory(centralRulesDir); err != nil {
		return fmt.Errorf("failed to create central rules directory: %w", err)
	}

	// Check if central library already exists and has content
	hasContent := false
	if entries, err := os.ReadDir(centralRulesDir); err == nil && len(entries) > 0 {
		hasContent = true
	}

	// Only extract if central library is empty
	if !hasContent {
		if err := ExtractRules(centralRulesDir); err != nil {
			return fmt.Errorf("failed to extract rules to central location: %w", err)
		}
	}

	// Create symlinks in each IDE's rules directory pointing to central location
	for _, ide := range request.IDEs {
		ideRulesDir, err := getIDERulesDir(projectPath, ide)
		if err != nil {
			return fmt.Errorf("failed to get rules directory for %s: %w", ide, err)
		}

		// Ensure parent directory exists
		parentDir := filepath.Dir(ideRulesDir)
		if err := utils.CreateDirectory(parentDir); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", ide, err)
		}

		// Ensure IDE rules directory exists
		if err := utils.CreateDirectory(ideRulesDir); err != nil {
			return fmt.Errorf("failed to create rules directory for %s: %w", ide, err)
		}

		// Create symlinks for each library folder
		if err := createLibraryFolderSymlinks(ideRulesDir, centralRulesDir); err != nil {
			// Fallback: copy rules if symlink creation fails (e.g., on Windows without admin)
			if err := copyLibraryFolders(ideRulesDir, centralRulesDir); err != nil {
				return fmt.Errorf("failed to create rules for %s (symlink and copy both failed): %w", ide, err)
			}
		}
	}

	return nil
}

// createLibraryFolderSymlinks creates symlinks for each folder in the library
// Returns error if symlink creation fails (caller should fallback to copying)
func createLibraryFolderSymlinks(ideRulesDir, centralRulesDir string) error {
	// Read all folders in the central library directory
	entries, err := os.ReadDir(centralRulesDir)
	if err != nil {
		return fmt.Errorf("failed to read central library directory: %w", err)
	}

	var firstError error
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		folderName := entry.Name()
		centralFolderPath := filepath.Join(centralRulesDir, folderName)
		ideFolderPath := filepath.Join(ideRulesDir, folderName)

		// Remove existing symlink/directory if it exists
		if utils.PathExists(ideFolderPath) {
			// Check if it's already a symlink pointing to the right place
			if utils.IsSymlink(ideFolderPath) {
				target, err := os.Readlink(ideFolderPath)
				if err == nil {
					// Resolve to absolute path for comparison
					absTarget, _ := filepath.Abs(target)
					absCentral, _ := filepath.Abs(centralFolderPath)
					if absTarget == absCentral || filepath.Clean(absTarget) == filepath.Clean(absCentral) {
						// Already correctly linked, skip
						continue
					}
				}
			}
			// Remove existing directory/link
			if err := os.RemoveAll(ideFolderPath); err != nil {
				if firstError == nil {
					firstError = fmt.Errorf("failed to remove existing folder %s: %w", folderName, err)
				}
				continue
			}
		}

		// Create symlink for this folder
		if err := utils.CreateSymlink(ideFolderPath, centralFolderPath); err != nil {
			if firstError == nil {
				firstError = fmt.Errorf("failed to create symlink for folder %s: %w", folderName, err)
			}
			// Continue trying other folders
			continue
		}
	}

	return firstError
}

// copyLibraryFolders copies library folders from central location to IDE directory (fallback)
func copyLibraryFolders(ideRulesDir, centralRulesDir string) error {
	// Read all folders in the central library directory
	entries, err := os.ReadDir(centralRulesDir)
	if err != nil {
		return fmt.Errorf("failed to read central library directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		folderName := entry.Name()
		centralFolderPath := filepath.Join(centralRulesDir, folderName)
		ideFolderPath := filepath.Join(ideRulesDir, folderName)

		// Copy the entire folder recursively
		if err := utils.CopyDirectory(centralFolderPath, ideFolderPath); err != nil {
			return fmt.Errorf("failed to copy folder %s: %w", folderName, err)
		}
	}

	return nil
}

// getIDERulesDir returns the rules directory path for the given IDE
func getIDERulesDir(projectPath, ide string) (string, error) {
	switch ide {
	case "Cursor":
		return filepath.Join(projectPath, ".cursor", "rules"), nil
	case "Claude Code":
		return filepath.Join(projectPath, ".claude", "rules"), nil
	case "Antigravity":
		return filepath.Join(projectPath, ".antigravity", "rules"), nil
	case "Windsurf":
		return filepath.Join(projectPath, ".windsurf", "rules"), nil
	case "Cline":
		return filepath.Join(projectPath, ".cline", "rules"), nil
	case "OpenCode":
		return filepath.Join(projectPath, ".opencode", "rules"), nil
	default:
		return "", fmt.Errorf("unsupported IDE: %s", ide)
	}
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
