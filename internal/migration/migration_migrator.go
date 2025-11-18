package migration

import (
	"fmt"
	"os"
	"path/filepath"
)

// Migrator orchestrates the entire migration process
type Migrator struct {
	projectRoot string
	detector    *Detector
	backup      *BackupManager
	config      *ConfigMigrator
	folders     *FolderMigrator
}

// NewMigrator creates a new migrator
func NewMigrator(projectRoot string) *Migrator {
	return &Migrator{
		projectRoot: projectRoot,
		detector:    NewDetector(projectRoot),
		backup:      NewBackupManager(projectRoot),
		config:      NewConfigMigrator(projectRoot),
		folders:     NewFolderMigrator(projectRoot),
	}
}

// MigrationResult contains the result of a migration
type MigrationResult struct {
	Success      bool
	BackupPath   string
	FoldersMigrated int
	Errors       []error
}

// Migrate performs the complete migration
func (m *Migrator) Migrate() (*MigrationResult, error) {
	result := &MigrationResult{
		Success: false,
		Errors:  []error{},
	}

	// Step 1: Detect old structure
	hasOld, err := m.detector.DetectOldStructure()
	if err != nil {
		return nil, fmt.Errorf("failed to detect old structure: %w", err)
	}

	if !hasOld {
		return nil, fmt.Errorf("old structure not detected")
	}

	// Step 2: Create backup
	backupPath, err := m.backup.CreateBackup()
	if err != nil {
		return nil, fmt.Errorf("failed to create backup: %w", err)
	}
	result.BackupPath = backupPath

	// Step 3: Migrate config
	if err := m.config.MigrateConfig(); err != nil {
		result.Errors = append(result.Errors, fmt.Errorf("config migration failed: %w", err))
		// Continue with folder migration
	}

	// Step 4: Detect old folders
	oldFolders, err := m.detector.DetectOldFolders()
	if err != nil {
		return nil, fmt.Errorf("failed to detect old folders: %w", err)
	}

	// Step 5: Migrate folders
	if err := m.folders.MigrateFolders(oldFolders); err != nil {
		result.Errors = append(result.Errors, fmt.Errorf("folder migration failed: %w", err))
	} else {
		result.FoldersMigrated = len(oldFolders)
	}

	// Step 6: Validate migration
	if err := m.Validate(); err != nil {
		result.Errors = append(result.Errors, fmt.Errorf("validation failed: %w", err))
	} else {
		result.Success = true
	}

	return result, nil
}

// Validate validates the migration
func (m *Migrator) Validate() error {
	// Check new config exists
	newConfigPath := filepath.Join(m.projectRoot, ".doplan", "config.yaml")
	if _, err := os.Stat(newConfigPath); err != nil {
		return fmt.Errorf("new config not found: %w", err)
	}

	// Check old folders are gone
	oldFolders, err := m.detector.DetectOldFolders()
	if err != nil {
		return fmt.Errorf("failed to check old folders: %w", err)
	}

	if len(oldFolders) > 0 {
		return fmt.Errorf("old folders still exist: %d folders", len(oldFolders))
	}

	return nil
}

// Rollback rolls back the migration using a backup
func (m *Migrator) Rollback(backupPath string) error {
	return m.backup.RestoreBackup(backupPath)
}

