package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// BackupManager manages backup creation and restoration
type BackupManager struct {
	projectRoot string
}

// NewBackupManager creates a new backup manager
func NewBackupManager(projectRoot string) *BackupManager {
	return &BackupManager{
		projectRoot: projectRoot,
	}
}

// CreateBackup creates a timestamped backup of the old structure
func (b *BackupManager) CreateBackup() (string, error) {
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	backupDir := filepath.Join(b.projectRoot, ".doplan", "backup", timestamp)

	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Backup old config
	oldConfigDir := filepath.Join(b.projectRoot, ".cursor", "config")
	if _, err := os.Stat(oldConfigDir); err == nil {
		backupConfigDir := filepath.Join(backupDir, ".cursor", "config")
		if err := b.copyDir(oldConfigDir, backupConfigDir); err != nil {
			return "", fmt.Errorf("failed to backup config: %w", err)
		}
	}

	// Backup doplan directory
	doplanDir := filepath.Join(b.projectRoot, "doplan")
	if _, err := os.Stat(doplanDir); err == nil {
		backupDoplanDir := filepath.Join(backupDir, "doplan")
		if err := b.copyDir(doplanDir, backupDoplanDir); err != nil {
			return "", fmt.Errorf("failed to backup doplan: %w", err)
		}
	}

	return backupDir, nil
}

// RestoreBackup restores from a backup
func (b *BackupManager) RestoreBackup(backupPath string) error {
	// Restore .cursor/config
	backupConfigDir := filepath.Join(backupPath, ".cursor", "config")
	if _, err := os.Stat(backupConfigDir); err == nil {
		targetConfigDir := filepath.Join(b.projectRoot, ".cursor", "config")
		if err := os.RemoveAll(targetConfigDir); err != nil {
			return fmt.Errorf("failed to remove old config: %w", err)
		}
		if err := b.copyDir(backupConfigDir, targetConfigDir); err != nil {
			return fmt.Errorf("failed to restore config: %w", err)
		}
	}

	// Restore doplan
	backupDoplanDir := filepath.Join(backupPath, "doplan")
	if _, err := os.Stat(backupDoplanDir); err == nil {
		targetDoplanDir := filepath.Join(b.projectRoot, "doplan")
		if err := os.RemoveAll(targetDoplanDir); err != nil {
			return fmt.Errorf("failed to remove old doplan: %w", err)
		}
		if err := b.copyDir(backupDoplanDir, targetDoplanDir); err != nil {
			return fmt.Errorf("failed to restore doplan: %w", err)
		}
	}

	return nil
}

// copyDir recursively copies a directory
func (b *BackupManager) copyDir(src, dst string) error {
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
			if err := b.copyDir(srcPath, dstPath); err != nil {
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
