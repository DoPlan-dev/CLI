package tui

import (
	"fmt"

	"github.com/DoPlan-dev/CLI/internal/tui/screens"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// App represents the main TUI application
type App struct {
	width     int
	height    int
	current   string // "menu", "dashboard", or action name
	menu      *screens.MenuModel
	dashboard *screens.DashboardModel
	executor  CommandExecutor // Command executor to avoid import cycles
}

// NewApp creates a new TUI app
func NewApp(executor CommandExecutor) *App {
	if executor == nil {
		executor = &DefaultExecutor{}
	}
	return &App{
		current:   "menu",
		menu:      screens.NewMenuModel(),
		dashboard: screens.NewDashboardModel(),
		executor:  executor,
	}
}

// Init initializes the app
func (a *App) Init() tea.Cmd {
	return a.menu.Init()
}

// Update handles messages
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		if a.menu != nil {
			var cmd tea.Cmd
			_, cmd = a.menu.Update(msg)
			return a, cmd
		}
		if a.dashboard != nil {
			var cmd tea.Cmd
			_, cmd = a.dashboard.Update(msg)
			return a, cmd
		}
		return a, nil

	case screens.MenuActionMsg:
		switch msg.Action {
		case "dashboard":
			a.current = "dashboard"
			return a, a.dashboard.Init()
		case "run":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.RunDevServer(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Dev server started"}
				},
			)
		case "undo":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.UndoLastAction(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Last action undone"}
				},
			)
		case "create":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.CreateNewProject(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "New project created"}
				},
			)
		case "deploy":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.DeployProject(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Deployment started"}
				},
			)
		case "publish":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.PublishPackage(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Package publishing started"}
				},
			)
		case "security":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.RunSecurityScan(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Security scan completed"}
				},
			)
		case "fix":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.AutoFix(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Auto-fix completed"}
				},
			)
		case "discuss":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.DiscussIdea(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Idea discussion completed"}
				},
			)
		case "generate":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.GenerateDocuments(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Documents generated"}
				},
			)
		case "plan":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.CreatePlan(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Plan created"}
				},
			)
		case "progress":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.UpdateProgress(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Progress updated"}
				},
			)
		case "design":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.ApplyDesign(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "Design system generated"}
				},
			)
		case "keys":
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.ManageAPIKeys(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "API keys managed"}
				},
			)
		case "integration":
			// This will be implemented in later phases
			return a, tea.Sequence(
				tea.Printf("Action '%s' - Coming soon in v0.0.19-beta!\n", msg.Action),
				func() tea.Msg {
					return screens.MenuActionMsg{Action: "menu"}
				},
			)
		case "menu":
			a.current = "menu"
			return a, nil
		default:
			return a, nil
		}

	case screens.ErrorMsg:
		// Show error and return to menu
		fmt.Printf("❌ Error: %v\n", msg.Error)
		a.current = "menu"
		return a, nil

	case screens.SuccessMsg:
		// Show success message and return to menu
		fmt.Printf("✅ %s\n", msg.Message)
		a.current = "menu"
		return a, nil

		case screens.BackToMenuMsg:
			a.current = "menu"
			return a, nil
		case screens.OpenKeysManagementMsg:
			// Open keys management
			return a, tea.Sequence(
				func() tea.Msg {
					if err := a.executor.ManageAPIKeys(); err != nil {
						return screens.ErrorMsg{Error: err}
					}
					return screens.SuccessMsg{Message: "API keys managed"}
				},
			)
	}

	// Delegate to current screen
	switch a.current {
	case "menu":
		if a.menu != nil {
			var cmd tea.Cmd
			var model tea.Model
			model, cmd = a.menu.Update(msg)
			if mm, ok := model.(*screens.MenuModel); ok {
				a.menu = mm
			}
			// Check if menu sent an action message (will be handled in next Update call)
			return a, cmd
		}
	case "dashboard":
		if a.dashboard != nil {
			var cmd tea.Cmd
			var model tea.Model
			model, cmd = a.dashboard.Update(msg)
			if dm, ok := model.(*screens.DashboardModel); ok {
				a.dashboard = dm
			}
			return a, cmd
		}
	}

	return a, nil
}

// View renders the UI
func (a *App) View() string {
	if a.width == 0 {
		return "Loading..."
	}

	header := RenderBestDoPlanHeader(a.width, "dev")

	var content string
	switch a.current {
	case "menu":
		if a.menu != nil {
			content = a.menu.View()
		}
	case "dashboard":
		if a.dashboard != nil {
			content = a.dashboard.View()
		}
	default:
		content = "Unknown view"
	}

	body := lipgloss.NewStyle().
		Width(a.width-4).
		Height(a.height-lipgloss.Height(header)-5).
		Padding(1, 2).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
	)
}

// Run starts the TUI with default executor
func Run() error {
	return RunWithExecutor(nil)
}

// RunWithExecutor starts the TUI with a custom command executor
func RunWithExecutor(executor CommandExecutor) error {
	p := tea.NewProgram(NewApp(executor), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// RenderBestDoPlanHeader renders the DoPlan header with ASCII art only
func RenderBestDoPlanHeader(width int, version string) string {
	topBorder := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#667eea", Light: "#764ba2"}).
		Width(width).
		Render("╔" + repeatString("═", width-2) + "╗")

	logo := `
  ██████╗░░█████╗░██████╗░██╗░░░░░░█████╗░███╗░░██╗
  ██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔══██╗████╗░██║
  ██║░░██║██║░░██║██████╔╝██║░░░░░███████║██╔██╗██║
  ██║░░██║██║░░██║██╔═══╝░██║░░░░░██╔══██║██║╚████║
  ██████╔╝╚█████╔╝██║░░░░░███████╗██║░░██║██║░╚███║
  ╚═════╝░░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚══╝
`

	styledLogo := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#667eea", Light: "#764ba2"}).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render(logo)

	bottomBorder := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#667eea", Light: "#764ba2"}).
		Width(width).
		Render("╚" + repeatString("═", width-2) + "╝")

	return lipgloss.JoinVertical(lipgloss.Center,
		topBorder,
		"",
		styledLogo,
		"",
		bottomBorder,
	)
}

func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
