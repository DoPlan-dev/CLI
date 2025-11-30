package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"encoding/json"
	"time"

	"github.com/spf13/cobra"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/internal/version"
)

// newSysBackupCommand creates the /sys backup subcommand
func newSysBackupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backup",
		Short: "Create compressed backup with flexible options",
		Long: `Create a compressed backup with granular control over what gets backed up.

Backup types:
  project       - Project files only (for DoPlan updates)
  plan          - Planning folder only (.do/plan/)
  project-plan  - Project files + planning folder
  full          - Everything (complete backup)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectPath, _ := os.Getwd()

			backupTypeFlag, _ := cmd.Flags().GetString("type")
			description, _ := cmd.Flags().GetString("description")

			var backupType BackupType

			if backupTypeFlag != "" {
				backupType = BackupType(backupTypeFlag)
				if !isValidBackupType(backupType) {
					return fmt.Errorf("invalid backup type: %s. Valid types: project, plan, project-plan, full", backupTypeFlag)
				}
			} else {
				// Interactive selection
				selected, err := selectBackupTypeInteractively(cmd.OutOrStdout(), cmd.InOrStdin())
				if err != nil {
					return err
				}
				backupType = selected
			}

			// Create backup
			backupPath, err := CreateBackup(projectPath, backupType, description)
			if err != nil {
				return fmt.Errorf("failed to create backup: %w", err)
			}

			// Get file size
			info, _ := os.Stat(backupPath)
			sizeMB := float64(info.Size()) / (1024 * 1024)

			fmt.Fprintf(cmd.OutOrStdout(), "✅ Backup created: %s\n", filepath.Base(backupPath))
			fmt.Fprintf(cmd.OutOrStdout(), "   Location: %s\n", backupPath)
			fmt.Fprintf(cmd.OutOrStdout(), "   Type: %s\n", backupType)
			fmt.Fprintf(cmd.OutOrStdout(), "   Size: %.2f MB\n", sizeMB)
			if description != "" {
				fmt.Fprintf(cmd.OutOrStdout(), "   Description: %s\n", description)
			}

			return nil
		},
	}

	cmd.Flags().StringP("type", "t", "", "Backup type: project, plan, project-plan, full")
	cmd.Flags().StringP("description", "d", "", "Backup description")

	return cmd
}

// newSysRestoreCommand creates the /sys restore subcommand
func newSysRestoreCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore project from compressed backup",
		Long: `Restore a project from a compressed backup file.

The restore operation will:
- Verify backup integrity
- Check version compatibility
- Restore files based on backup type
- Handle conflicts intelligently`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectPath, _ := os.Getwd()
			dryRun, _ := cmd.Flags().GetBool("dry-run")

			var backupPath string

			if len(args) > 0 {
				// Backup filename provided
				backupDir := filepath.Join(projectPath, ".do", "backup")
				backupPath = filepath.Join(backupDir, args[0])
				if !strings.HasSuffix(args[0], ".doplan") {
					backupPath += ".doplan"
				}
			} else {
				// Interactive selection
				selected, err := selectBackupInteractively(projectPath, cmd.OutOrStdout(), cmd.InOrStdin())
				if err != nil {
					return err
				}
				backupPath = selected
			}

			// Verify backup exists
			if !utils.PathExists(backupPath) {
				return fmt.Errorf("backup file not found: %s", backupPath)
			}

			// Extract manifest for preview
			manifest, err := ExtractManifest(backupPath)
			if err != nil {
				return fmt.Errorf("failed to read backup: %w", err)
			}

			// Show backup info
			fmt.Fprintf(cmd.OutOrStdout(), "📦 Backup Information:\n")
			fmt.Fprintf(cmd.OutOrStdout(), "   File: %s\n", filepath.Base(backupPath))
			fmt.Fprintf(cmd.OutOrStdout(), "   Type: %s\n", manifest.BackupType)
			fmt.Fprintf(cmd.OutOrStdout(), "   Created: %s\n", manifest.Timestamp)
			if manifest.Description != "" {
				fmt.Fprintf(cmd.OutOrStdout(), "   Description: %s\n", manifest.Description)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "   Version: %s\n", manifest.DoplanVersion)
			fmt.Fprintf(cmd.OutOrStdout(), "   Files: %d\n", manifest.FileCount)
			fmt.Fprintf(cmd.OutOrStdout(), "\n")

			// Check version compatibility
			currentVersion := version.GetVersion()
			compatible, msg := CheckVersionCompatibility(manifest.DoplanVersion, currentVersion)
			if !compatible {
				fmt.Fprintf(cmd.OutOrStdout(), "⚠️  Warning: %s\n", msg)
				fmt.Fprintf(cmd.OutOrStdout(), "\n")
				fmt.Fprintf(cmd.OutOrStdout(), "Continue anyway? (yes/no): ")
				reader := bufio.NewReader(cmd.InOrStdin())
				response, _ := reader.ReadString('\n')
				if strings.TrimSpace(strings.ToLower(response)) != "yes" {
					return fmt.Errorf("restore cancelled")
				}
			} else if msg != "" {
				fmt.Fprintf(cmd.OutOrStdout(), "ℹ️  Note: %s\n", msg)
			}

			// Verify backup integrity
			fmt.Fprintf(cmd.OutOrStdout(), "🔍 Verifying backup integrity...\n")
			if err := VerifyBackup(backupPath); err != nil {
				return fmt.Errorf("backup verification failed: %w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "✅ Backup integrity verified\n")
			fmt.Fprintf(cmd.OutOrStdout(), "\n")

			// Create safety backup before restore
			if !dryRun {
				fmt.Fprintf(cmd.OutOrStdout(), "💾 Creating safety backup of current state...\n")
				safetyBackup, err := CreateBackup(projectPath, BackupTypeFull, "Safety backup before restore")
				if err != nil {
					return fmt.Errorf("failed to create safety backup: %w", err)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "✅ Safety backup created: %s\n", filepath.Base(safetyBackup))
				fmt.Fprintf(cmd.OutOrStdout(), "\n")
			}

			// Restore
			options := RestoreOptions{
				DryRun:        dryRun,
				BackupPath:    backupPath,
				ProjectPath:   projectPath,
				MergeStrategy: "overwrite", // Default to overwrite, can be enhanced
			}

			result, err := RestoreBackup(options)
			if err != nil {
				return fmt.Errorf("restore failed: %w", err)
			}

			// Show results
			if dryRun {
				fmt.Fprintf(cmd.OutOrStdout(), "🔍 Dry Run Results:\n")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "✅ Restore Complete!\n")
			}
			fmt.Fprintf(cmd.OutOrStdout(), "   Files restored: %d\n", result.FilesRestored)
			if result.FilesSkipped > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "   Files skipped: %d\n", result.FilesSkipped)
			}
			if result.FilesMerged > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "   Files merged: %d\n", result.FilesMerged)
			}
			if len(result.Conflicts) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "   Conflicts: %d\n", len(result.Conflicts))
			}
			if len(result.Warnings) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "\n⚠️  Warnings:\n")
				for _, warning := range result.Warnings {
					fmt.Fprintf(cmd.OutOrStdout(), "   • %s\n", warning)
				}
			}

			return nil
		},
	}

	cmd.Flags().Bool("dry-run", false, "Preview what would be restored without actually restoring")

	return cmd
}

// newSysMigrateCommand creates the /sys migrate subcommand
func newSysMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Guided migration assistant for DoPlan updates",
		Long: `Guided migration assistant that helps you update DoPlan while preserving your work.

The migration assistant will:
- Detect old DoPlan structure
- Suggest appropriate backup type
- Guide you through the update process
- Help restore your project after update`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectPath, _ := os.Getwd()
			return RunInteractiveMigration(projectPath, cmd.OutOrStdout(), cmd.InOrStdin())
		},
	}

	return cmd
}

// newSysMemoryCommand creates the /sys memory subcommand
func newSysMemoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memory",
		Short: "Export or restore memory card",
		Long: `Manage memory card data for transferring user preferences and engagement data between projects.

The memory card contains:
- User preferences and work style
- Command usage statistics
- Achievements and challenges
- Engagement metrics
- Project history and context`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectPath, _ := os.Getwd()

			action := "export"
			if len(args) > 0 {
				action = args[0]
			}

			if action != "export" && action != "restore" {
				return fmt.Errorf("invalid action: %s. Use 'export' or 'restore'", action)
			}

			if action == "export" {
				return exportMemoryCard(projectPath, cmd.OutOrStdout())
			} else {
				return restoreMemoryCard(projectPath, cmd.OutOrStdout(), cmd.InOrStdin())
			}
		},
	}

	return cmd
}

// Helper functions

func isValidBackupType(backupType BackupType) bool {
	return backupType == BackupTypeProject ||
		backupType == BackupTypePlan ||
		backupType == BackupTypeProjectPlan ||
		backupType == BackupTypeFull
}

func selectBackupTypeInteractively(out io.Writer, in io.Reader) (BackupType, error) {
	fmt.Fprintln(out, "What would you like to backup?")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "1. Project files only (for DoPlan updates) - excludes .do/ and .cursor/")
	fmt.Fprintln(out, "2. Planning folder only (.do/plan/)")
	fmt.Fprintln(out, "3. Project files + planning folder")
	fmt.Fprintln(out, "4. Everything (complete backup)")
	fmt.Fprintln(out, "")
	fmt.Fprint(out, "Select option [1-4]: ")

	reader := bufio.NewReader(in)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	switch response {
	case "1":
		return BackupTypeProject, nil
	case "2":
		return BackupTypePlan, nil
	case "3":
		return BackupTypeProjectPlan, nil
	case "4":
		return BackupTypeFull, nil
	default:
		return "", fmt.Errorf("invalid selection: %s", response)
	}
}

func selectBackupInteractively(projectPath string, out io.Writer, in io.Reader) (string, error) {
	backups, err := ListBackups(projectPath)
	if err != nil {
		return "", fmt.Errorf("failed to list backups: %w", err)
	}

	if len(backups) == 0 {
		return "", fmt.Errorf("no backups found in .do/backup/")
	}

	fmt.Fprintln(out, "Available backups:")
	fmt.Fprintln(out, "")
	for i, backup := range backups {
		fmt.Fprintf(out, "%d. %s\n", i+1, backup.Filename)
		fmt.Fprintf(out, "   Type: %s | Size: %.2f MB | Created: %s\n",
			backup.BackupType,
			float64(backup.Size)/(1024*1024),
			backup.Timestamp.Format("2006-01-02 15:04:05"))
		if backup.Description != "" {
			fmt.Fprintf(out, "   Description: %s\n", backup.Description)
		}
		fmt.Fprintln(out, "")
	}

	fmt.Fprint(out, "Select backup [1-", len(backups), "]: ")

	reader := bufio.NewReader(in)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	index, err := strconv.Atoi(response)
	if err != nil || index < 1 || index > len(backups) {
		return "", fmt.Errorf("invalid selection: %s", response)
	}

	return backups[index-1].Path, nil
}

func exportMemoryCard(projectPath string, out io.Writer) error {
	// Load memory card
	mc, err := LoadMemoryCard()
	if err != nil {
		return fmt.Errorf("failed to load memory card: %w", err)
	}

	// Ensure backup directory exists
	backupDir := filepath.Join(projectPath, ".do", "backup")
	if err := utils.CreateDirectory(backupDir); err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Generate filename
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("memory_card_export-%s.json", timestamp)
	exportPath := filepath.Join(backupDir, filename)

	// Export memory card
	data, err := json.MarshalIndent(mc, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal memory card: %w", err)
	}

	if err := utils.WriteFile(exportPath, data); err != nil {
		return fmt.Errorf("failed to write export file: %w", err)
	}

	fmt.Fprintf(out, "✅ Memory card exported to: %s\n", filepath.Base(exportPath))
	fmt.Fprintf(out, "   Location: %s\n", exportPath)
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "You can restore this in your next project with:\n")
	fmt.Fprintf(out, "   /sys memory restore\n")

	return nil
}

func restoreMemoryCard(projectPath string, out io.Writer, in io.Reader) error {
	// Find memory card exports
	backupDir := filepath.Join(projectPath, ".do", "backup")
	if !utils.PathExists(backupDir) {
		return fmt.Errorf("backup directory not found. No memory card exports available.")
	}

	var exports []string
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return fmt.Errorf("failed to read backup directory: %w", err)
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "memory_card_export-") && strings.HasSuffix(entry.Name(), ".json") {
			exports = append(exports, entry.Name())
		}
	}

	if len(exports) == 0 {
		return fmt.Errorf("no memory card exports found in .do/backup/")
	}

	// Select export
	fmt.Fprintln(out, "Available memory card exports:")
	fmt.Fprintln(out, "")
	for i, export := range exports {
		fmt.Fprintf(out, "%d. %s\n", i+1, export)
	}
	fmt.Fprintln(out, "")
	fmt.Fprint(out, "Select export [1-", len(exports), "]: ")

	reader := bufio.NewReader(in)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	index, err := strconv.Atoi(response)
	if err != nil || index < 1 || index > len(exports) {
		return fmt.Errorf("invalid selection: %s", response)
	}

	selectedExport := filepath.Join(backupDir, exports[index-1])

	// Backup current memory card
	memoryCardPath := filepath.Join(os.Getenv("HOME"), ".doplan", "memory_card.json")
	if utils.PathExists(memoryCardPath) {
		backupPath := memoryCardPath + ".backup." + time.Now().Format("20060102-150405")
		data, _ := os.ReadFile(memoryCardPath)
		utils.WriteFile(backupPath, data)
		fmt.Fprintf(out, "✅ Current memory card backed up to: %s\n", filepath.Base(backupPath))
	}

	// Read export
	exportData, err := os.ReadFile(selectedExport)
	if err != nil {
		return fmt.Errorf("failed to read export file: %w", err)
	}

	// Parse memory card
	var mc MemoryCard
	if err := json.Unmarshal(exportData, &mc); err != nil {
		return fmt.Errorf("failed to parse memory card: %w", err)
	}

	// Restore memory card
	if err := SaveMemoryCard(&mc); err != nil {
		return fmt.Errorf("failed to restore memory card: %w", err)
	}

	fmt.Fprintf(out, "✅ Memory card restored from: %s\n", exports[index-1])
	if utils.PathExists(memoryCardPath + ".backup." + time.Now().Format("20060102-150405")) {
		fmt.Fprintf(out, "   Previous memory card backed up\n")
	}

	return nil
}
