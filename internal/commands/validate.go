package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/validator"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate project structure",
		Long:  "Validate project structure, configuration, and consistency",
		RunE:  runValidate,
	}

	cmd.Flags().BoolP("fix", "f", false, "Automatically fix issues where possible")

	return cmd
}

func runValidate(cmd *cobra.Command, args []string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return doplanerror.NewIOError("IO001", "Failed to get current directory").WithCause(err)
	}

	errLogger := doplanerror.NewLogger(projectRoot, doplanerror.LogLevelInfo)
	errHandler := doplanerror.NewHandler(errLogger)

	if !config.IsInstalled(projectRoot) {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		errHandler.PrintError(doplanerror.ErrConfigNotFound(configPath))
		return nil
	}

	color.Blue("🔍 Validating project structure...\n\n")

	val := validator.NewValidator(projectRoot)
	issues, err := val.Validate()
	if err != nil {
		return errHandler.Handle(doplanerror.NewValidationError("VAL006", "Failed to validate project").WithCause(err))
	}

	if len(issues) == 0 {
		color.Green("✅ All checks passed! Project structure is valid.\n")
		return nil
	}

	// Group issues by level
	errors := []validator.Issue{}
	warnings := []validator.Issue{}
	infos := []validator.Issue{}

	for _, issue := range issues {
		switch issue.Level {
		case "error":
			errors = append(errors, issue)
		case "warning":
			warnings = append(warnings, issue)
		case "info":
			infos = append(infos, issue)
		}
	}

	// Print errors
	if len(errors) > 0 {
		color.Red("❌ Errors (%d):\n", len(errors))
		for _, issue := range errors {
			fmt.Printf("  • %s\n", issue.Message)
			if issue.Path != "" {
				fmt.Printf("    Path: %s\n", issue.Path)
			}
			if issue.Fix != "" {
				fmt.Printf("    Fix:  %s\n", issue.Fix)
			}
			fmt.Println()
		}
	}

	// Print warnings
	if len(warnings) > 0 {
		color.Yellow("⚠️  Warnings (%d):\n", len(warnings))
		for _, issue := range warnings {
			fmt.Printf("  • %s\n", issue.Message)
			if issue.Path != "" {
				fmt.Printf("    Path: %s\n", issue.Path)
			}
			if issue.Fix != "" {
				fmt.Printf("    Fix:  %s\n", issue.Fix)
			}
			fmt.Println()
		}
	}

	// Print info
	if len(infos) > 0 {
		color.Cyan("ℹ️  Info (%d):\n", len(infos))
		for _, issue := range infos {
			fmt.Printf("  • %s\n", issue.Message)
			if issue.Path != "" {
				fmt.Printf("    Path: %s\n", issue.Path)
			}
			fmt.Println()
		}
	}

	// Auto-fix if requested
	autoFix, _ := cmd.Flags().GetBool("fix")
	if autoFix && len(errors) > 0 {
		color.Cyan("\n🔧 Attempting to auto-fix errors...\n")
		if err := val.AutoFix(errors); err != nil {
			color.Yellow("⚠️  Some issues could not be auto-fixed\n")
		}
	}

	// Summary
	total := len(issues)
	if len(errors) > 0 {
		color.Red("\n❌ Validation failed with %d issue(s)\n", total)
		return fmt.Errorf("validation failed")
	}

	if len(warnings) > 0 {
		color.Yellow("\n⚠️  Validation passed with %d warning(s)\n", len(warnings))
		return nil
	}

	color.Green("\n✅ Validation passed!\n")
	return nil
}
