package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/template"
)

func NewTemplatesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "templates",
		Short: "Manage document templates",
		Long:  "Manage templates for feature plans, designs, and tasks",
	}

	cmd.AddCommand(NewTemplatesListCommand())
	cmd.AddCommand(NewTemplatesShowCommand())
	cmd.AddCommand(NewTemplatesAddCommand())
	cmd.AddCommand(NewTemplatesEditCommand())
	cmd.AddCommand(NewTemplatesUseCommand())
	cmd.AddCommand(NewTemplatesRemoveCommand())

	return cmd
}

func NewTemplatesListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all templates",
		RunE:  runTemplatesList,
	}
}

func runTemplatesList(cmd *cobra.Command, args []string) error {
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

	mgr := template.NewManager(projectRoot)
	templates, err := mgr.ListTemplates()
	if err != nil {
		templateDir := filepath.Join(projectRoot, "doplan", "templates")
		return errHandler.Handle(doplanerror.NewIOError("IO005", "Failed to list templates").WithPath(templateDir).WithCause(err))
	}

	if len(templates) == 0 {
		color.Yellow("No templates found.")
		color.Cyan("Templates will be created during 'doplan install'")
		return nil
	}

	cfg, _ := mgr.LoadConfig()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Template\tType\tDefault")
	fmt.Fprintln(w, "---\t---\t---")

	for _, tmpl := range templates {
		tmplType := "custom"
		if strings.Contains(tmpl, "plan") {
			tmplType = "plan"
		} else if strings.Contains(tmpl, "design") {
			tmplType = "design"
		} else if strings.Contains(tmpl, "task") {
			tmplType = "tasks"
		}

		isDefault := ""
		if cfg != nil {
			if tmpl == cfg.DefaultPlan || tmpl == cfg.DefaultDesign || tmpl == cfg.DefaultTasks {
				isDefault = "✓"
			}
		}

		fmt.Fprintf(w, "%s\t%s\t%s\n", tmpl, tmplType, isDefault)
	}

	w.Flush()
	return nil
}

func NewTemplatesShowCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show <template-name>",
		Short: "Show template content",
		Args:  cobra.ExactArgs(1),
		RunE:  runTemplatesShow,
	}
}

func runTemplatesShow(cmd *cobra.Command, args []string) error {
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

	mgr := template.NewManager(projectRoot)
	content, err := mgr.GetTemplate(args[0])
	if err != nil {
		templatePath := filepath.Join(projectRoot, "doplan", "templates", args[0])
		return errHandler.Handle(doplanerror.ErrFileNotFound(templatePath).WithCause(err))
	}

	fmt.Println(content)
	return nil
}

func NewTemplatesAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <name> <file-path>",
		Short: "Add a template from file",
		Args:  cobra.ExactArgs(2),
		RunE:  runTemplatesAdd,
	}
	return cmd
}

func runTemplatesAdd(cmd *cobra.Command, args []string) error {
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

	name := args[0]
	filePath := args[1]

	// Ensure .md extension
	if !strings.HasSuffix(name, ".md") {
		name += ".md"
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return errHandler.Handle(doplanerror.ErrFileNotFound(filePath).WithCause(err))
	}

	mgr := template.NewManager(projectRoot)
	if err := mgr.AddTemplate(name, string(content)); err != nil {
		templateDir := filepath.Join(projectRoot, "doplan", "templates")
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to add template").WithPath(templateDir).WithCause(err))
	}

	color.Green("✅ Template '%s' added successfully\n", name)
	return nil
}

func NewTemplatesEditCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "edit <template-name>",
		Short: "Edit template (opens in default editor)",
		Args:  cobra.ExactArgs(1),
		RunE:  runTemplatesEdit,
	}
}

func runTemplatesEdit(cmd *cobra.Command, args []string) error {
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

	templatePath := filepath.Join(projectRoot, "doplan", "templates", args[0])

	// Check if template exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return errHandler.Handle(doplanerror.ErrFileNotFound(templatePath).WithCause(err))
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		// Try common editors
		editors := []string{"nano", "vim", "vi", "code", "subl"}
		for _, e := range editors {
			if _, err := exec.LookPath(e); err == nil {
				editor = e
				break
			}
		}
		if editor == "" {
			return errHandler.Handle(doplanerror.NewIOError("IO001", "No editor found").WithSuggestion("Set EDITOR environment variable or install a text editor"))
		}
	}

	color.Cyan("Opening template in %s...\n", editor)
	color.Yellow("Template path: %s\n", templatePath)

	// Open editor
	cmdExec := exec.Command(editor, templatePath)
	cmdExec.Stdin = os.Stdin
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr

	if err := cmdExec.Run(); err != nil {
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to open editor").WithPath(templatePath).WithCause(err))
	}

	color.Green("✅ Template edited\n")
	return nil
}

func NewTemplatesUseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use <template-name> [--for plan|design|tasks]",
		Short: "Set default template",
		Args:  cobra.ExactArgs(1),
		RunE:  runTemplatesUse,
	}
	cmd.Flags().String("for", "", "Template type (plan, design, tasks)")
	return cmd
}

func runTemplatesUse(cmd *cobra.Command, args []string) error {
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

	templateName := args[0]
	templateFor, _ := cmd.Flags().GetString("for")

	mgr := template.NewManager(projectRoot)

	// Verify template exists
	if _, err := mgr.GetTemplate(templateName); err != nil {
		templatePath := filepath.Join(projectRoot, "doplan", "templates", templateName)
		return errHandler.Handle(doplanerror.ErrFileNotFound(templatePath).WithCause(err))
	}

	cfg, err := mgr.LoadConfig()
	if err != nil {
		configPath := filepath.Join(projectRoot, ".doplan", "templates.json")
		return errHandler.Handle(doplanerror.NewIOError("IO005", "Failed to load template config").WithPath(configPath).WithCause(err))
	}

	if templateFor != "" {
		switch templateFor {
		case "plan":
			cfg.DefaultPlan = templateName
		case "design":
			cfg.DefaultDesign = templateName
		case "tasks":
			cfg.DefaultTasks = templateName
		default:
			return errHandler.Handle(doplanerror.NewValidationError("VAL007", "Invalid template type").WithDetails(fmt.Sprintf("Type: %s (use: plan, design, tasks)", templateFor)))
		}
	} else {
		// Auto-detect type
		nameLower := strings.ToLower(templateName)
		if strings.Contains(nameLower, "plan") {
			cfg.DefaultPlan = templateName
			templateFor = "plan"
		} else if strings.Contains(nameLower, "design") {
			cfg.DefaultDesign = templateName
			templateFor = "design"
		} else if strings.Contains(nameLower, "task") {
			cfg.DefaultTasks = templateName
			templateFor = "tasks"
		} else {
			return errHandler.Handle(doplanerror.NewValidationError("VAL008", "Cannot auto-detect template type").WithSuggestion("Use --for flag to specify type"))
		}
	}

	if err := mgr.SaveConfig(cfg); err != nil {
		configPath := filepath.Join(projectRoot, ".doplan", "templates.json")
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to save template config").WithPath(configPath).WithCause(err))
	}

	color.Green("✅ Default %s template set: %s\n", templateFor, templateName)
	return nil
}

func NewTemplatesRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "remove <template-name>",
		Short: "Remove a template",
		Args:  cobra.ExactArgs(1),
		RunE:  runTemplatesRemove,
	}
}

func runTemplatesRemove(cmd *cobra.Command, args []string) error {
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

	mgr := template.NewManager(projectRoot)

	// Check if it's a default template
	cfg, _ := mgr.LoadConfig()
	if cfg != nil {
		if args[0] == cfg.DefaultPlan || args[0] == cfg.DefaultDesign || args[0] == cfg.DefaultTasks {
			color.Yellow("⚠️  This is a default template. Removing it may cause issues.")
			color.Yellow("Continue? (y/n): ")
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) != "y" {
				color.Cyan("Removal cancelled.")
				return nil
			}
		}
	}

	if err := mgr.RemoveTemplate(args[0]); err != nil {
		templatePath := filepath.Join(projectRoot, "doplan", "templates", args[0])
		return errHandler.Handle(doplanerror.NewIOError("IO006", "Failed to remove template").WithPath(templatePath).WithCause(err))
	}

	color.Green("✅ Template '%s' removed\n", args[0])
	return nil
}

