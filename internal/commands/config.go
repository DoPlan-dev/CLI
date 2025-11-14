package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage DoPlan configuration",
		Long:  "View and modify DoPlan configuration settings",
	}

	cmd.AddCommand(NewConfigShowCommand())
	cmd.AddCommand(NewConfigSetCommand())
	cmd.AddCommand(NewConfigResetCommand())
	cmd.AddCommand(NewConfigValidateCommand())

	return cmd
}

func NewConfigShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		Long:  "Display the current DoPlan configuration",
		RunE:  runConfigShow,
	}

	cmd.Flags().StringP("format", "f", "table", "Output format (table, json)")

	return cmd
}

func runConfigShow(cmd *cobra.Command, args []string) error {
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

	cfgMgr := config.NewManager(projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		return errHandler.Handle(doplanerror.ErrConfigNotFound(configPath).WithCause(err))
	}

	format, _ := cmd.Flags().GetString("format")

	switch format {
	case "json":
		data, _ := json.MarshalIndent(cfg, "", "  ")
		fmt.Println(string(data))
	default:
		printConfigTable(cfg)
	}

	return nil
}

func printConfigTable(cfg *models.Config) {
	color.Cyan("\n📋 DoPlan Configuration\n")
	fmt.Printf("IDE:         %s\n", cfg.IDE)
	fmt.Printf("Version:     %s\n", cfg.Version)
	fmt.Printf("Installed:   %s\n", formatBool(cfg.Installed))
	fmt.Printf("Installed At: %s\n", cfg.InstalledAt.Format("2006-01-02 15:04:05"))

	color.Cyan("\n🔗 GitHub Settings\n")
	fmt.Printf("Enabled:     %s\n", formatBool(cfg.GitHub.Enabled))
	fmt.Printf("Auto Branch: %s\n", formatBool(cfg.GitHub.AutoBranch))
	fmt.Printf("Auto PR:     %s\n", formatBool(cfg.GitHub.AutoPR))

	color.Cyan("\n💾 Checkpoint Settings\n")
	fmt.Printf("Auto Feature:  %s\n", formatBool(cfg.Checkpoint.AutoFeature))
	fmt.Printf("Auto Phase:    %s\n", formatBool(cfg.Checkpoint.AutoPhase))
	fmt.Printf("Auto Complete: %s\n", formatBool(cfg.Checkpoint.AutoComplete))

	color.Cyan("\n📊 State\n")
	fmt.Printf("Current Phase:   %s\n", cfg.State.CurrentPhase)
	fmt.Printf("Current Feature: %s\n", cfg.State.CurrentFeature)
	fmt.Printf("Idea Captured:   %s\n", formatBool(cfg.State.IdeaCaptured))
	fmt.Printf("PRD Generated:   %s\n", formatBool(cfg.State.PRDGenerated))
	fmt.Printf("Plan Generated:  %s\n", formatBool(cfg.State.PlanGenerated))
	fmt.Println()
}

func formatBool(b bool) string {
	if b {
		return color.GreenString("✓")
	}
	return color.RedString("✗")
}

func NewConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set configuration value",
		Long:  "Set a configuration value. Use 'true' or 'false' for boolean values.",
		RunE:  runConfigSet,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func runConfigSet(cmd *cobra.Command, args []string) error {
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

	cfgMgr := config.NewManager(projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		return errHandler.Handle(doplanerror.ErrConfigNotFound(configPath).WithCause(err))
	}

	key := args[0]
	value := args[1]

	updated := false

	switch key {
	case "github.enabled":
		cfg.GitHub.Enabled = value == "true"
		updated = true
	case "github.autoBranch":
		cfg.GitHub.AutoBranch = value == "true"
		updated = true
	case "github.autoPR":
		cfg.GitHub.AutoPR = value == "true"
		updated = true
	case "checkpoint.autoFeature":
		cfg.Checkpoint.AutoFeature = value == "true"
		updated = true
	case "checkpoint.autoPhase":
		cfg.Checkpoint.AutoPhase = value == "true"
		updated = true
	case "checkpoint.autoComplete":
		cfg.Checkpoint.AutoComplete = value == "true"
		updated = true
	default:
		return errHandler.Handle(doplanerror.NewValidationError("VAL004", "Unknown config key").WithDetails(fmt.Sprintf("Key: %s\n\nAvailable keys:\n  github.enabled\n  github.autoBranch\n  github.autoPR\n  checkpoint.autoFeature\n  checkpoint.autoPhase\n  checkpoint.autoComplete", key)))
	}

	if !updated {
		return errHandler.Handle(doplanerror.NewValidationError("VAL005", "Failed to update config"))
	}

	if err := cfgMgr.SaveConfig(cfg); err != nil {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to save config").WithPath(configPath).WithCause(err))
	}

	color.Green("✅ Configuration updated: %s = %s\n", key, value)

	return nil
}

func NewConfigResetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset configuration to defaults",
		Long:  "Reset configuration to default values (keeps IDE setting)",
		RunE:  runConfigReset,
	}
}

func runConfigReset(cmd *cobra.Command, args []string) error {
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

	cfgMgr := config.NewManager(projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		return errHandler.Handle(doplanerror.ErrConfigNotFound(configPath).WithCause(err))
	}

	// Keep IDE setting, reset others
	ide := cfg.IDE
	cfg = config.NewConfig(ide)
	cfg.IDE = ide

	if err := cfgMgr.SaveConfig(cfg); err != nil {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to save config").WithPath(configPath).WithCause(err))
	}

	color.Green("✅ Configuration reset to defaults\n")

	return nil
}

func NewConfigValidateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "Validate configuration",
		Long:  "Validate configuration for errors and inconsistencies",
		RunE:  runConfigValidate,
	}
}

func runConfigValidate(cmd *cobra.Command, args []string) error {
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

	cfgMgr := config.NewManager(projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
		return errHandler.Handle(doplanerror.ErrConfigNotFound(configPath).WithCause(err))
	}

	issues := validateConfig(cfg)

	if len(issues) == 0 {
		color.Green("✅ Configuration is valid\n")
		return nil
	}

	color.Yellow("⚠️  Configuration issues found:\n")
	for _, issue := range issues {
		color.Red("  - %s\n", issue)
	}

	return nil
}

func validateConfig(cfg *models.Config) []string {
	var issues []string

	if cfg.IDE == "" {
		issues = append(issues, "IDE not set")
	}

	if cfg.GitHub.Enabled && cfg.GitHub.AutoPR && !cfg.GitHub.AutoBranch {
		issues = append(issues, "AutoPR requires AutoBranch to be enabled")
	}

	if cfg.GitHub.AutoPR && !cfg.GitHub.Enabled {
		issues = append(issues, "AutoPR requires GitHub to be enabled")
	}

	return issues
}

