package tui

import (
	"fmt"

	"github.com/DoPlan-dev/CLI/internal/tui/screens"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// App represents the main TUI application
type App struct {
	dashboard *screens.DashboardModel
}

// NewApp creates a new TUI app
func NewApp() *App {
	return &App{
		dashboard: screens.NewDashboardModel(),
	}
}

// Init initializes the app
func (a *App) Init() tea.Cmd {
	return a.dashboard.Init()
}

// Update handles messages
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var model tea.Model
	model, cmd = a.dashboard.Update(msg)
	if dm, ok := model.(*screens.DashboardModel); ok {
		a.dashboard = dm
	}
	return a, cmd
}

// View renders the UI
func (a *App) View() string {
	return a.dashboard.View()
}

// Run starts the TUI
func Run() error {
	p := tea.NewProgram(NewApp(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// RenderBestDoPlanHeader renders the DoPlan header with ASCII art
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

	title := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("DoPlan")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#999999", Light: "#666666"}).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("Project Workflow Manager")

	versionText := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#666666", Light: "#999999"}).
		Width(width - 4).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("v%s", version))

	bottomBorder := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Dark: "#667eea", Light: "#764ba2"}).
		Width(width).
		Render("╚" + repeatString("═", width-2) + "╝")

	return lipgloss.JoinVertical(lipgloss.Center,
		topBorder,
		"",
		styledLogo,
		"",
		title,
		subtitle,
		versionText,
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
