package wizard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// WizardState represents the state of a wizard
type WizardState struct {
	CurrentScreen string
	Step          int
	TotalSteps    int
	Data          map[string]interface{}
	Errors        []error
	Width         int
	Height        int
}

// WizardModel is the base model for all wizards
type WizardModel struct {
	state  WizardState
	styles WizardStyles
}

// WizardStyles contains all styling for the wizard
type WizardStyles struct {
	Header      lipgloss.Style
	Body        lipgloss.Style
	Button      lipgloss.Style
	ButtonActive lipgloss.Style
	Error       lipgloss.Style
	Success     lipgloss.Style
	Help        lipgloss.Style
}

// NewWizardStyles creates default wizard styles
func NewWizardStyles() WizardStyles {
	return WizardStyles{
		Header: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Padding(0, 1),
		Body: lipgloss.NewStyle().
			Padding(1, 2),
		Button: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#667eea")).
			Padding(0, 2).
			Margin(1, 0),
		ButtonActive: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#764ba2")).
			Padding(0, 2).
			Margin(1, 0),
		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ef4444")).
			Bold(true),
		Success: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10b981")).
			Bold(true),
		Help: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#999999")).
			Italic(true),
	}
}

// Init initializes the wizard
func (m *WizardModel) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m *WizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.state.Width = msg.Width
		m.state.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the wizard
func (m *WizardModel) View() string {
	return m.renderScreen()
}

// renderScreen renders the current screen
func (m *WizardModel) renderScreen() string {
	switch m.state.CurrentScreen {
	case "welcome":
		return m.renderWelcome()
	case "projectName":
		return m.renderProjectName()
	// Add more screens
	default:
		return "Unknown screen"
	}
}

// renderWelcome renders the welcome screen
func (m *WizardModel) renderWelcome() string {
	header := m.styles.Header.Render("Welcome to DoPlan")
	body := "Welcome message here"
	footer := m.styles.Help.Render("Press Enter to continue, Esc to exit")
	
	return lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		"",
		body,
		"",
		footer,
	)
}

// renderProjectName renders the project name input screen
func (m *WizardModel) renderProjectName() string {
	// Implementation here
	return "Project name screen"
}

