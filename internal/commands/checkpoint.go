package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/DoPlan-dev/CLI/internal/checkpoint"
	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCheckpointCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkpoint",
		Short: "Manage project checkpoints",
		Long:  "Create, list, and restore project checkpoints for features and phases",
	}

	cmd.AddCommand(NewCheckpointCreateCommand())
	cmd.AddCommand(NewCheckpointListCommand())
	cmd.AddCommand(NewCheckpointRestoreCommand())

	return cmd
}

func NewCheckpointCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a checkpoint",
		Long:  "Create a manual checkpoint of the current project state",
		RunE:  runCheckpointCreate,
		Args:  cobra.MaximumNArgs(1),
	}

	cmd.Flags().StringP("type", "t", "manual", "Checkpoint type (manual, feature, phase)")
	cmd.Flags().StringP("description", "d", "", "Checkpoint description")

	return cmd
}

func runCheckpointCreate(cmd *cobra.Command, args []string) error {
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

	checkpointType, _ := cmd.Flags().GetString("type")
	description, _ := cmd.Flags().GetString("description")

	name := "Manual Checkpoint"
	if len(args) > 0 {
		name = args[0]
	}

	if description == "" {
		description = fmt.Sprintf("Manual checkpoint created at %s", name)
	}

	cm := checkpoint.NewCheckpointManager(projectRoot)
	_, err = cm.CreateCheckpoint(checkpointType, name, description)
	if err != nil {
		checkpointDir := filepath.Join(projectRoot, ".doplan", "checkpoints")
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to create checkpoint").WithPath(checkpointDir).WithCause(err))
	}

	color.Green("✅ Checkpoint created successfully!\n")
	return nil
}

func NewCheckpointListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all checkpoints",
		Long:  "List all available checkpoints",
		RunE:  runCheckpointList,
	}
}

func runCheckpointList(cmd *cobra.Command, args []string) error {
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

	cm := checkpoint.NewCheckpointManager(projectRoot)
	checkpoints, err := cm.ListCheckpoints()
	if err != nil {
		checkpointDir := filepath.Join(projectRoot, ".doplan", "checkpoints")
		return errHandler.Handle(doplanerror.NewIOError("IO005", "Failed to list checkpoints").WithPath(checkpointDir).WithCause(err))
	}

	if len(checkpoints) == 0 {
		color.Yellow("No checkpoints found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tType\tName\tCreated At")
	fmt.Fprintln(w, "---\t---\t---\t---")

	for _, cp := range checkpoints {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			cp.ID,
			cp.Type,
			cp.Name,
			cp.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}

	w.Flush()

	return nil
}

func NewCheckpointRestoreCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "restore <checkpoint-id>",
		Short: "Restore a checkpoint",
		Long:  "Restore project state from a checkpoint",
		RunE:  runCheckpointRestore,
		Args:  cobra.ExactArgs(1),
	}
}

func runCheckpointRestore(cmd *cobra.Command, args []string) error {
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

	checkpointID := args[0]

	// Confirm restoration
	color.Yellow("⚠️  This will restore the project to checkpoint: %s\n", checkpointID)
	color.Yellow("Current state will be overwritten. Continue? (y/n): ")

	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) != "y" {
		color.Cyan("Restoration cancelled.")
		return nil
	}

	cm := checkpoint.NewCheckpointManager(projectRoot)
	if err := cm.RestoreCheckpoint(checkpointID); err != nil {
		checkpointDir := filepath.Join(projectRoot, ".doplan", "checkpoints", checkpointID)
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to restore checkpoint").WithPath(checkpointDir).WithCause(err))
	}

	color.Green("✅ Checkpoint restored successfully!")
	color.Cyan("Run 'doplan dashboard' to view the restored state.")

	return nil
}
