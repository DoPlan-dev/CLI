package wizard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/deployment"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type deployModel struct {
	width         int
	height        int
	currentScreen deployScreen
	projectType   string
	platform      string
	envVars       map[string]string
	err           error
}

type deployScreen int

const (
	screenDeployWelcome deployScreen = iota
	screenDeployPlatform
	screenDeployConfig
	screenDeployProgress
	screenDeploySuccess
)

// RunDeployWizard launches the deployment wizard
func RunDeployWizard() error {
	m := &deployModel{
		currentScreen: screenDeployWelcome,
		envVars:       make(map[string]string),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func (m *deployModel) Init() tea.Cmd {
	return nil
}

func (m *deployModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			case screenDeployWelcome:
				m.currentScreen = screenDeployPlatform
				return m, nil
			case screenDeployPlatform:
				// Platform selected, move to config
				m.currentScreen = screenDeployConfig
				return m, nil
			case screenDeployConfig:
				// Start deployment
				m.currentScreen = screenDeployProgress
				return m, m.deploy()
			case screenDeploySuccess:
				return m, tea.Quit
			}
		case "esc":
			if m.currentScreen > screenDeployWelcome {
				m.currentScreen--
				return m, nil
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *deployModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	header := renderDeployHeader(m.width)

	var content string
	switch m.currentScreen {
	case screenDeployWelcome:
		content = m.renderWelcome()
	case screenDeployPlatform:
		content = m.renderPlatformSelection()
	case screenDeployConfig:
		content = m.renderConfig()
	case screenDeployProgress:
		content = m.renderProgress()
	case screenDeploySuccess:
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

func renderDeployHeader(width int) string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("🚀 Deploy Project")

	return lipgloss.JoinVertical(lipgloss.Center, "", title, "")
}

func (m *deployModel) renderWelcome() string {
	text := "Welcome to the Deployment Wizard!\n\n"
	text += "This wizard will help you deploy your project to various platforms.\n\n"
	text += "Press Enter to continue..."
	return text
}

func (m *deployModel) renderPlatformSelection() string {
	projectRoot, _ := os.Getwd()
	detected, _ := deployment.DetectPlatform(projectRoot)

	text := "Select deployment platform:\n\n"
	platforms := []struct {
		id   string
		name string
		desc string
	}{
		{"vercel", "Vercel", "Next.js, React, Static sites"},
		{"netlify", "Netlify", "JAMstack, Static sites"},
		{"railway", "Railway", "Full-stack apps"},
		{"render", "Render", "Docker, Static sites"},
		{"coolify", "Coolify", "Self-hosted"},
		{"docker", "Docker", "Custom deployment"},
	}

	for i, p := range platforms {
		marker := "  "
		if p.id == detected {
			marker = "→ " // Recommended
		}
		text += fmt.Sprintf("%s%d. %s (%s)\n", marker, i+1, p.name, p.desc)
	}

	if detected != "" {
		text += fmt.Sprintf("\n💡 Detected platform: %s (recommended)\n", detected)
	}

	text += "\nPress Enter to continue with detected platform..."
	text += "\n(Full platform selection UI coming soon)"

	return text
}

func (m *deployModel) renderConfig() string {
	text := "Configure deployment:\n\n"
	text += "Environment variables and deployment settings.\n\n"
	text += "Press Enter to deploy (coming soon)..."
	return text
}

func (m *deployModel) renderProgress() string {
	text := "Deploying...\n\n"
	if m.err != nil {
		text += fmt.Sprintf("❌ Error: %v\n", m.err)
	} else {
		text += "⏳ Deployment in progress..."
	}
	return text
}

func (m *deployModel) renderSuccess() string {
	text := "✅ Deployment successful!\n\n"
	text += "Your project has been deployed.\n\n"
	text += "Press Enter to exit..."
	return text
}

func (m *deployModel) deploy() tea.Cmd {
	return func() tea.Msg {
		projectRoot, _ := os.Getwd()

		// Detect project type
		projectType, err := detectProjectTypeForDeploy(projectRoot)
		if err != nil {
			return deployErrorMsg{Error: err}
		}
		m.projectType = projectType

		// If no platform selected, detect best platform
		if m.platform == "" {
			detected, err := deployment.DetectPlatform(projectRoot)
			if err == nil {
				m.platform = detected
			} else {
				m.platform = "vercel" // Default
			}
		}

		// Deploy based on platform
		var deployErr error
		switch m.platform {
		case "vercel":
			deployErr = deployment.DeployToVercel(projectRoot, m.envVars)
		case "netlify":
			deployErr = deployment.DeployToNetlify(projectRoot, m.envVars)
		case "railway":
			deployErr = deployment.DeployToRailway(projectRoot, m.envVars)
		case "render":
			deployErr = deployment.DeployToRender(projectRoot, m.envVars)
		case "coolify":
			deployErr = deployment.DeployToCoolify(projectRoot, m.envVars)
		case "docker", "custom":
			deployErr = deployment.DeployWithDocker(projectRoot, m.envVars)
		default:
			deployErr = fmt.Errorf("unsupported platform: %s", m.platform)
		}

		if deployErr != nil {
			m.err = deployErr
			return deployErrorMsg{Error: deployErr}
		}

		return deploySuccessMsg{}
	}
}

func detectProjectTypeForDeploy(projectRoot string) (string, error) {
	// Check for package.json
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		return "node", nil
	}

	// Check for go.mod
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
		return "go", nil
	}

	// Check for Dockerfile
	if _, err := os.Stat(filepath.Join(projectRoot, "Dockerfile")); err == nil {
		return "docker", nil
	}

	return "unknown", fmt.Errorf("could not detect project type")
}

type deployErrorMsg struct {
	Error error
}

type deploySuccessMsg struct{}
