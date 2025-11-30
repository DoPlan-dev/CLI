package cli

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/internal/version"
)

// RestoreOptions contains options for restore operations
type RestoreOptions struct {
	DryRun         bool
	MergeStrategy  string // "overwrite", "merge", "skip"
	BackupPath     string
	ProjectPath    string
	SelectedBackup string
}

// RestoreResult contains information about the restore operation
type RestoreResult struct {
	FilesRestored int
	FilesSkipped  int
	FilesMerged   int
	Conflicts     []string
	Warnings      []string
}

// RestoreBackup restores files from a backup archive
func RestoreBackup(options RestoreOptions) (*RestoreResult, error) {
	result := &RestoreResult{
		Conflicts: []string{},
		Warnings:  []string{},
	}

	// Extract manifest first
	manifest, err := ExtractManifest(options.BackupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract manifest: %w", err)
	}

	// Verify backup integrity
	if err := VerifyBackup(options.BackupPath); err != nil {
		return nil, fmt.Errorf("backup verification failed: %w", err)
	}

	// Check version compatibility
	currentVersion := version.GetVersion()
	compatible, msg := CheckVersionCompatibility(manifest.DoplanVersion, currentVersion)
	if !compatible {
		result.Warnings = append(result.Warnings, msg)
	} else if msg != "" {
		result.Warnings = append(result.Warnings, msg)
	}

	// Open backup file
	file, err := os.Open(options.BackupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open backup file: %w", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	// Track what we're restoring
	filesToRestore := map[string]bool{}
	conflictFiles := []string{}

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read archive: %w", err)
		}

		// Skip manifest (we already extracted it)
		if header.Name == "BACKUP_MANIFEST.json" {
			continue
		}

		// Determine restore path based on backup type
		restorePath, shouldRestore := getRestorePath(header.Name, options.ProjectPath, manifest.BackupType)
		if !shouldRestore {
			continue
		}

		// Check for conflicts
		if utils.PathExists(restorePath) {
			conflictFiles = append(conflictFiles, restorePath)
			if options.DryRun {
				result.Conflicts = append(result.Conflicts, restorePath)
				continue
			}

			// Handle conflict based on merge strategy
			switch options.MergeStrategy {
			case "skip":
				result.FilesSkipped++
				continue
			case "merge":
				// Special handling for .do/plan/ directory
				if strings.HasPrefix(header.Name, ".do/plan/") {
					if err := mergePlanFile(restorePath, tarReader, header); err != nil {
						result.Warnings = append(result.Warnings, fmt.Sprintf("Failed to merge %s: %v", restorePath, err))
						continue
					}
					result.FilesMerged++
					continue
				}
				// For other files, skip on conflict (could enhance later)
				result.FilesSkipped++
				continue
			case "overwrite":
				// Will overwrite below
			default:
				result.FilesSkipped++
				continue
			}
		}

		if options.DryRun {
			filesToRestore[restorePath] = true
			continue
		}

		// Restore file
		if err := restoreFile(tarReader, header, restorePath); err != nil {
			result.Warnings = append(result.Warnings, fmt.Sprintf("Failed to restore %s: %v", restorePath, err))
			continue
		}

		result.FilesRestored++
	}

	if options.DryRun {
		result.FilesRestored = len(filesToRestore)
		result.Conflicts = conflictFiles
	}

	return result, nil
}

// getRestorePath determines where to restore a file based on backup type
func getRestorePath(archivePath, projectPath string, backupType BackupType) (string, bool) {
	switch backupType {
	case BackupTypeProject:
		// Restore project files only, skip .do/ and .cursor/
		if strings.HasPrefix(archivePath, ".do/") || strings.HasPrefix(archivePath, ".cursor/") {
			return "", false
		}
		return filepath.Join(projectPath, archivePath), true

	case BackupTypePlan:
		// Restore only .do/plan/ directory
		if !strings.HasPrefix(archivePath, ".do/plan/") {
			return "", false
		}
		return filepath.Join(projectPath, archivePath), true

	case BackupTypeProjectPlan:
		// Restore project files + .do/plan/, skip .do/core/ and .do/system/
		if strings.HasPrefix(archivePath, ".do/core/") || strings.HasPrefix(archivePath, ".do/system/") {
			return "", false
		}
		if strings.HasPrefix(archivePath, ".cursor/") {
			return "", false
		}
		return filepath.Join(projectPath, archivePath), true

	case BackupTypeFull:
		// Restore everything
		return filepath.Join(projectPath, archivePath), true

	default:
		return "", false
	}
}

// restoreFile restores a single file from the archive
func restoreFile(tarReader *tar.Reader, header *tar.Header, restorePath string) error {
	// Create parent directory
	parentDir := filepath.Dir(restorePath)
	if err := utils.CreateDirectory(parentDir); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	// Create file
	file, err := os.Create(restorePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Copy data
	if _, err := io.Copy(file, tarReader); err != nil {
		return fmt.Errorf("failed to copy file data: %w", err)
	}

	// Restore file permissions
	if err := os.Chmod(restorePath, os.FileMode(header.Mode)); err != nil {
		// Non-critical error, just warn
	}

	return nil
}

// mergePlanFile merges a planning file (for .do/plan/ directory)
func mergePlanFile(existingPath string, tarReader *tar.Reader, header *tar.Header) error {
	// Read existing file
	existingData, err := os.ReadFile(existingPath)
	if err != nil {
		// If we can't read existing, just overwrite
		return restoreFile(tarReader, header, existingPath)
	}

	// Read new file
	newData, err := io.ReadAll(tarReader)
	if err != nil {
		return fmt.Errorf("failed to read new file: %w", err)
	}

	// For now, simple merge strategy: append new content if different
	// TODO: Implement smarter merge for markdown files
	if string(existingData) != string(newData) {
		// Create backup of existing
		backupPath := existingPath + ".backup"
		if err := utils.WriteFile(backupPath, existingData); err != nil {
			return fmt.Errorf("failed to backup existing file: %w", err)
		}

		// For TASKS.md, we might want smarter merging
		// For now, overwrite with new version
		return utils.WriteFile(existingPath, newData)
	}

	// Files are identical, no merge needed
	return nil
}

// DetectOldStructure detects if project has old DoPlan structure
func DetectOldStructure(projectPath string) (bool, []string) {
	indicators := []string{}
	hasOldStructure := false

	// Check for .plan/ directory (old structure)
	if utils.PathExists(filepath.Join(projectPath, ".plan")) {
		hasOldStructure = true
		indicators = append(indicators, ".plan/ directory found (old structure)")
	}

	// Check for old command structure in .cursor/
	cursorDir := filepath.Join(projectPath, ".cursor", "commands")
	if utils.PathExists(cursorDir) {
		// Check if it's not symlinked (old structure has real files)
		if !utils.IsSymlink(cursorDir) {
			hasOldStructure = true
			indicators = append(indicators, ".cursor/commands contains real files (should be symlinks)")
		}
	}

	// Check for missing .do/ structure
	if !utils.PathExists(filepath.Join(projectPath, ".do", "core")) {
		hasOldStructure = true
		indicators = append(indicators, ".do/core/ directory missing (new structure)")
	}

	return hasOldStructure, indicators
}

// SuggestBackupType suggests the best backup type based on project state
func SuggestBackupType(projectPath string, hasOldStructure bool) BackupType {
	if hasOldStructure {
		// For migration, suggest project-plan to preserve work + planning
		return BackupTypeProjectPlan
	}

	// For updates, suggest project to preserve work only
	return BackupTypeProject
}
