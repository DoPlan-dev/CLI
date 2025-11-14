package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/DoPlan-dev/CLI/internal/commands"
	"github.com/DoPlan-dev/CLI/internal/ui"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "doplan",
		Short: "DoPlan - Project Workflow Manager",
		Long: `DoPlan automates your project workflow from idea to deployment.
Combines Spec-Kit and BMAD-METHOD methodologies.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
	}

	// Add commands
	rootCmd.AddCommand(commands.NewInstallCommand())
	rootCmd.AddCommand(commands.NewDashboardCommand())
	rootCmd.AddCommand(commands.NewGitHubCommand())
	rootCmd.AddCommand(commands.NewProgressCommand())
	rootCmd.AddCommand(commands.NewConfigCommand())
	rootCmd.AddCommand(commands.NewCheckpointCommand())
	rootCmd.AddCommand(commands.NewValidateCommand())
	rootCmd.AddCommand(commands.NewTemplatesCommand())
	rootCmd.AddCommand(commands.NewStatsCommand())

	// TUI mode flag
	tuiFlag := rootCmd.PersistentFlags().Bool("tui", false, "Run in TUI mode")
	
	// Check for TUI mode
	if len(os.Args) > 1 && os.Args[1] == "--tui" {
		if err := ui.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}
	
	// Check flag value if parsed
	if tuiFlag != nil && *tuiFlag {
		if err := ui.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

