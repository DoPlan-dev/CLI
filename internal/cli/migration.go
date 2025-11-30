package cli

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

// MigrationGuide provides step-by-step migration guidance
type MigrationGuide struct {
	HasOldStructure bool
	Indicators      []string
	SuggestedType   BackupType
	Steps           []string
}

// DetectAndGuideMigration detects old structure and provides migration guidance
func DetectAndGuideMigration(projectPath string) (*MigrationGuide, error) {
	hasOldStructure, indicators := DetectOldStructure(projectPath)
	suggestedType := SuggestBackupType(projectPath, hasOldStructure)

	guide := &MigrationGuide{
		HasOldStructure: hasOldStructure,
		Indicators:      indicators,
		SuggestedType:   suggestedType,
		Steps:           generateMigrationSteps(hasOldStructure, suggestedType),
	}

	return guide, nil
}

// generateMigrationSteps generates step-by-step migration instructions
func generateMigrationSteps(hasOldStructure bool, backupType BackupType) []string {
	steps := []string{}

	if hasOldStructure {
		steps = append(steps, "⚠️  Old DoPlan structure detected!")
		steps = append(steps, "")
		steps = append(steps, "Step 1: Create backup of your project")
		steps = append(steps, fmt.Sprintf("   Run: /sys backup --type %s --description \"Before migration\"", backupType))
		steps = append(steps, "")
		steps = append(steps, "Step 2: Install/update DoPlan CLI")
		steps = append(steps, "   Run: npx @doplan-dev/cli@latest")
		steps = append(steps, "")
		steps = append(steps, "Step 3: Restore your project files")
		steps = append(steps, "   Run: /sys restore [backup-filename]")
		steps = append(steps, "")
		steps = append(steps, "Step 4: Verify everything works")
		steps = append(steps, "   Test your project and verify all files are intact")
	} else {
		steps = append(steps, "✅ Current DoPlan structure looks up-to-date")
		steps = append(steps, "")
		steps = append(steps, "For future DoPlan updates:")
		steps = append(steps, fmt.Sprintf("  1. Backup: /sys backup --type %s", backupType))
		steps = append(steps, "  2. Update: npx @doplan-dev/cli@latest")
		steps = append(steps, "  3. Restore: /sys restore [backup-filename]")
	}

	return steps
}

// RunInteractiveMigration runs an interactive migration wizard
func RunInteractiveMigration(projectPath string, out io.Writer, in io.Reader) error {
	guide, err := DetectAndGuideMigration(projectPath)
	if err != nil {
		return fmt.Errorf("failed to detect migration status: %w", err)
	}

	fmt.Fprintln(out, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Fprintln(out, "🔄 DoPlan Migration Assistant")
	fmt.Fprintln(out, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Fprintln(out, "")

	if guide.HasOldStructure {
		fmt.Fprintln(out, "⚠️  Old DoPlan structure detected!")
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Indicators found:")
		for _, indicator := range guide.Indicators {
			fmt.Fprintf(out, "  • %s\n", indicator)
		}
		fmt.Fprintln(out, "")
		fmt.Fprintf(out, "Recommended backup type: %s\n", guide.SuggestedType)
		fmt.Fprintln(out, "")
	} else {
		fmt.Fprintln(out, "✅ Current DoPlan structure looks up-to-date!")
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "You can use this tool to prepare for future updates.")
		fmt.Fprintln(out, "")
	}

	// Show steps
	fmt.Fprintln(out, "Migration steps:")
	fmt.Fprintln(out, "")
	for i, step := range guide.Steps {
		fmt.Fprintf(out, "  %s\n", step)
		if i < len(guide.Steps)-1 && step == "" {
			continue
		}
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Would you like to start the migration process? (yes/no)")

	reader := bufio.NewReader(in)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response != "yes" && response != "y" {
		fmt.Fprintln(out, "Migration cancelled.")
		return nil
	}

	// Start migration process
	if guide.HasOldStructure {
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Starting migration...")
		fmt.Fprintf(out, "Creating backup with type: %s\n", guide.SuggestedType)

		// Create backup
		backupPath, err := CreateBackup(projectPath, guide.SuggestedType, "Migration backup - created by migration assistant")
		if err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}

		fmt.Fprintf(out, "✅ Backup created: %s\n", filepath.Base(backupPath))
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Next steps:")
		fmt.Fprintln(out, "  1. Install/update DoPlan: npx @doplan-dev/cli@latest")
		fmt.Fprintf(out, "  2. Restore your project: /sys restore %s\n", filepath.Base(backupPath))
	} else {
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "No migration needed! Your project structure is up-to-date.")
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "For future updates, use:")
		fmt.Fprintf(out, "  /sys backup --type %s\n", guide.SuggestedType)
	}

	return nil
}
