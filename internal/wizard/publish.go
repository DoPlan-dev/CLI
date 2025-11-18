package wizard

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/publisher"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type publishModel struct {
	width         int
	height        int
	currentScreen publishScreen
	packageType   string
	registry      string
	err           error
}

type publishScreen int

const (
	screenPublishWelcome publishScreen = iota
	screenPublishType
	screenPublishConfig
	screenPublishProgress
	screenPublishSuccess
)

// RunPublishWizard launches the package publishing wizard
func RunPublishWizard() error {
	m := &publishModel{
		currentScreen: screenPublishWelcome,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func (m *publishModel) Init() tea.Cmd {
	return nil
}

func (m *publishModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch m.currentScreen {
			case screenPublishWelcome:
				m.currentScreen = screenPublishType
				// Auto-detect package type
				projectRoot, _ := os.Getwd()
				if detected, err := publisher.DetectPackageType(projectRoot); err == nil {
					m.registry = detected
					m.packageType = detected
				}
				return m, nil
			case screenPublishType:
				m.currentScreen = screenPublishConfig
				return m, nil
			case screenPublishConfig:
				m.currentScreen = screenPublishProgress
				return m, m.publish()
			case screenPublishSuccess:
				return m, tea.Quit
			}
		case "esc":
			if m.currentScreen > screenPublishWelcome {
				m.currentScreen--
				return m, nil
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *publishModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	header := renderPublishHeader(m.width)

	var content string
	switch m.currentScreen {
	case screenPublishWelcome:
		content = m.renderWelcome()
	case screenPublishType:
		content = m.renderTypeSelection()
	case screenPublishConfig:
		content = m.renderConfig()
	case screenPublishProgress:
		content = m.renderProgress()
	case screenPublishSuccess:
		content = m.renderSuccess()
	}

	body := lipgloss.NewStyle().
		Width(m.width-4).
		Height(m.height-lipgloss.Height(header)-5).
		Padding(1, 2).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
	)
}

func renderPublishHeader(width int) string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("📦 Publish Package")

	return lipgloss.JoinVertical(lipgloss.Center, "", title, "")
}

func (m *publishModel) renderWelcome() string {
	text := "Welcome to the Package Publishing Wizard!\n\n"
	text += "This wizard will help you publish your package to various registries.\n\n"
	text += "Press Enter to continue..."
	return text
}

func (m *publishModel) renderTypeSelection() string {
	text := "Select package type:\n\n"
	text += "1. npm (Node.js packages)\n"
	text += "2. Homebrew (macOS/Linux CLI tools)\n"
	text += "3. Scoop (Windows CLI tools)\n"
	text += "4. Winget (Windows apps)\n\n"
	text += "Press Enter to select (coming soon)..."
	return text
}

func (m *publishModel) renderConfig() string {
	text := "Configure package:\n\n"
	text += "Package metadata and registry settings.\n\n"
	text += "Press Enter to publish (coming soon)..."
	return text
}

func (m *publishModel) renderProgress() string {
	text := "Publishing...\n\n"
	if m.err != nil {
		text += fmt.Sprintf("❌ Error: %v\n", m.err)
	} else {
		text += "⏳ Publishing in progress..."
	}
	return text
}

func (m *publishModel) renderSuccess() string {
	text := "✅ Package published successfully!\n\n"
	text += "Your package is now available.\n\n"
	text += "Press Enter to exit..."
	return text
}

func (m *publishModel) publish() tea.Cmd {
	return func() tea.Msg {
		projectRoot, _ := os.Getwd()

		// Detect package type
		packageType, err := publisher.DetectPackageType(projectRoot)
		if err != nil {
			return publishErrorMsg{Error: err}
		}
		m.packageType = packageType

		// If no registry selected, use detected type
		if m.registry == "" {
			m.registry = packageType
		}

		// Publish based on registry
		var publishErr error
		switch m.registry {
		case "npm":
			publishErr = publisher.PublishToNPM(projectRoot)
		case "homebrew":
			publishErr = publisher.PublishToHomebrew(projectRoot)
		case "scoop":
			publishErr = publisher.PublishToScoop(projectRoot)
		case "winget":
			publishErr = publisher.PublishToWinget(projectRoot)
		default:
			publishErr = fmt.Errorf("unsupported registry: %s", m.registry)
		}

		if publishErr != nil {
			m.err = publishErr
			return publishErrorMsg{Error: publishErr}
		}

		return publishSuccessMsg{}
	}
}

type publishErrorMsg struct {
	Error error
}

type publishSuccessMsg struct{}
