package cli

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/tui"
	"github.com/DoPlan-dev/CLI/internal/version"
	"github.com/spf13/cobra"
)

const (
	// AppName is the name of the application
	AppName = "doplan"
)

// Version returns the version of the application
func Version() string {
	return version.GetVersion()
}

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   AppName,
		Short: "DoPlan CLI - Zero-install project generator with AI agency",
		Long: `DoPlan CLI is a zero-install, pure-Go command-line tool that instantly 
generates professional project structures with a complete hierarchical AI agency system.

Users can bootstrap production-ready projects in seconds with full automation, 
intelligent agents, and comprehensive rules libraries.

Examples:
  doplan              # Run interactive wizard to create a new project
  doplan --version    # Show version information
  doplan --help       # Show help information`,
		Version: Version(),
		// Default action: run the wizard
		Run: func(cmd *cobra.Command, args []string) {
			// Run the TUI wizard
			if err := runWizard(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// init initializes the root command with flags and configuration
func init() {
	// Add version flag
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and exit")

	// Override version template to show custom version format
	rootCmd.SetVersionTemplate(fmt.Sprintf("%s version {{.Version}}\n", AppName))
}

// GetRootCmd returns the root command (useful for testing)
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// runWizard runs the TUI wizard and handles the returned project data
func runWizard() error {
	request, err := tui.Run()
	if err != nil {
		return err
	}

	// If user quit without completing, exit gracefully
	if request == nil {
		fmt.Println("Wizard cancelled. No project created.")
		return nil
	}

	// Generation is now handled inside the wizard, so if we get here,
	// the project was successfully created
	fmt.Printf("\n✨ Project '%s' created successfully!\n", request.ProjectName)
	fmt.Printf("Open with: %s ./%s\n", getIDECommand(request.IDE), request.ProjectName)
	fmt.Printf("Then type /tell to begin\n")

	return nil
}

// getIDECommand returns the command to open a project in the selected IDE
func getIDECommand(ide string) string {
	switch ide {
	case "Cursor":
		return "cursor"
	case "Claude Code":
		return "claude"
	case "Antigravity":
		return "antigravity"
	case "Windsurf":
		return "windsurf"
	case "Cline":
		return "cline"
	case "OpenCode":
		return "opencode"
	default:
		return "code"
	}
}
