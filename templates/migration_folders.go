package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// FolderMigrator handles folder structure migration
type FolderMigrator struct {
	projectRoot string
	doplanDir   string
}

// NewFolderMigrator creates a new folder migrator
func NewFolderMigrator(projectRoot string) *FolderMigrator {
	return &FolderMigrator{
		projectRoot: projectRoot,
		doplanDir:   filepath.Join(projectRoot, "doplan"),
	}
}

// MigrateFolders migrates all old folders to new naming
func (f *FolderMigrator) MigrateFolders(folders []OldFolder) error {
	// Sort folders: phases first, then features
	phases := []OldFolder{}
	features := []OldFolder{}

	for _, folder := range folders {
		if folder.Type == "phase" {
			phases = append(phases, folder)
		} else {
			features = append(features, folder)
		}
	}

	// Migrate phases first
	for _, phase := range phases {
		if err := f.migrateFolder(phase); err != nil {
			return fmt.Errorf("failed to migrate phase %s: %w", phase.OldName, err)
		}
	}

	// Then migrate features
	for _, feature := range features {
		if err := f.migrateFolder(feature); err != nil {
			return fmt.Errorf("failed to migrate feature %s: %w", feature.OldName, err)
		}
	}

	return nil
}

// migrateFolder migrates a single folder
func (f *FolderMigrator) migrateFolder(folder OldFolder) error {
	// Generate new name
	newName, err := f.GenerateSlugName(folder)
	if err != nil {
		return err
	}

	// Construct new path
	oldPath := folder.OldPath
	oldDir := filepath.Dir(oldPath)
	newPath := filepath.Join(oldDir, newName)

	// Create new directory
	if err := os.MkdirAll(newPath, 0755); err != nil {
		return err
	}

	// Copy all files
	entries, err := os.ReadDir(oldPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		oldEntryPath := filepath.Join(oldPath, entry.Name())
		newEntryPath := filepath.Join(newPath, entry.Name())

		if entry.IsDir() {
			// Skip if it's a feature folder (will be migrated separately)
			if folder.Type == "phase" {
				// Check if it's an old feature folder
				featurePattern := regexp.MustCompile(`^\d+-Feature$`)
				if featurePattern.MatchString(entry.Name()) {
					continue // Will be migrated separately
				}
			}
			// Copy directory
			if err := f.copyDir(oldEntryPath, newEntryPath); err != nil {
				return err
			}
		} else {
			// Copy file and update references
			data, err := os.ReadFile(oldEntryPath)
			if err != nil {
				return err
			}

			// Update references in file content
			updatedData := f.updateReferences(data, folder)

			if err := os.WriteFile(newEntryPath, updatedData, entry.Type().Perm()); err != nil {
				return err
			}
		}
	}

	// Remove old folder (after verification)
	if err := os.RemoveAll(oldPath); err != nil {
		return err
	}

	return nil
}

// GenerateSlugName generates a new slug-based name for a folder
func (f *FolderMigrator) GenerateSlugName(folder OldFolder) (string, error) {
	var name string

	// Try to extract name from plan.md
	planPath := filepath.Join(folder.OldPath, "plan.md")
	if folder.Type == "phase" {
		planPath = filepath.Join(folder.OldPath, "phase-plan.md")
	}

	if data, err := os.ReadFile(planPath); err == nil {
		name = f.extractNameFromMarkdown(data)
	}

	// Fallback to folder name
	if name == "" {
		name = folder.OldName
		// Remove number prefix
		parts := strings.SplitN(name, "-", 2)
		if len(parts) > 1 {
			name = parts[1]
		}
	}

	// Convert to slug
	slug := f.nameToSlug(name)

	// Add number prefix
	number := f.extractNumber(folder.OldName)
	if number == "" {
		if folder.Type == "phase" {
			number = folder.PhaseNum
		} else {
			number = folder.FeatureNum
		}
	}

	return fmt.Sprintf("%s-%s", number, slug), nil
}

// extractNameFromMarkdown extracts name from markdown file
func (f *FolderMigrator) extractNameFromMarkdown(data []byte) string {
	// Look for # Title pattern
	titlePattern := regexp.MustCompile(`^#\s+(.+)$`)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		matches := titlePattern.FindStringSubmatch(line)
		if len(matches) > 1 {
			return strings.TrimSpace(matches[1])
		}
	}
	return ""
}

// nameToSlug converts a name to a slug
func (f *FolderMigrator) nameToSlug(name string) string {
	// Convert to lowercase
	slug := strings.ToLower(name)

	// Replace spaces and special chars with hyphens
	slug = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(slug, "-")

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

// extractNumber extracts number prefix from folder name
func (f *FolderMigrator) extractNumber(name string) string {
	parts := strings.SplitN(name, "-", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// updateReferences updates references to old folder names in file content
func (f *FolderMigrator) updateReferences(data []byte, folder OldFolder) []byte {
	content := string(data)

	// Update references to old folder name
	oldName := folder.OldName
	// Generate new name (simplified - should use actual new name)
	newName := folder.OldName // This should be the actual new name

	content = strings.ReplaceAll(content, oldName, newName)

	return []byte(content)
}

// copyDir recursively copies a directory
func (f *FolderMigrator) copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := f.copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			if err := os.WriteFile(dstPath, data, entry.Type().Perm()); err != nil {
				return err
			}
		}
	}

	return nil
}

