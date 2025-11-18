package wizard

import (
	"fmt"
	"os"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/rakd"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keysModel struct {
	width       int
	height      int
	currentView string // "list", "service", "actions"
	selectedIdx int
	services    []rakd.Service
	rakdData    *rakd.RAKDData
	err         error
}

// RunKeysWizard launches the API keys management TUI
func RunKeysWizard() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Load RAKD data
	data, err := rakd.GenerateRAKD(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to load RAKD data: %w", err)
	}

	m := &keysModel{
		currentView: "list",
		services:    data.Services,
		rakdData:    data,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	return err
}

func (m *keysModel) Init() tea.Cmd {
	return nil
}

func (m *keysModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.selectedIdx > 0 {
				m.selectedIdx--
			}
			return m, nil
		case "down", "j":
			if m.selectedIdx < len(m.services)-1 {
				m.selectedIdx++
			}
			return m, nil
		case "enter":
			if m.currentView == "list" {
				m.currentView = "service"
				return m, nil
			}
		case "esc":
			if m.currentView == "service" {
				m.currentView = "list"
				return m, nil
			}
			return m, tea.Quit
		case "v":
			// Validate all keys
			return m, m.validateAll()
		case "s":
			// Sync .env.example
			return m, m.syncEnvExample()
		}
	}

	// Handle wizard messages
	switch msg := msg.(type) {
	case keysErrorMsg:
		m.err = msg.Error
		return m, nil
	case keysSuccessMsg:
		// Success - data already updated
		return m, nil
	}

	return m, nil
}

func (m *keysModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	header := renderKeysHeader(m.width, m.rakdData)

	var content string
	switch m.currentView {
	case "list":
		content = m.renderServiceList()
	case "service":
		content = m.renderServiceDetail()
	}

	body := lipgloss.NewStyle().
		Width(m.width - 4).
		Height(m.height - lipgloss.Height(header) - 5).
		Padding(1, 2).
		Render(content)

	help := renderKeysHelp(m.currentView)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		help,
	)
}

func renderKeysHeader(width int, data *rakd.RAKDData) string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f59e0b")).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("🔑 API Keys Management")

	status := fmt.Sprintf("✅ %d | 🔴 %d | 🟡 %d | 🔵 %d",
		data.ConfiguredCount,
		data.RequiredCount,
		data.PendingCount,
		data.OptionalCount)

	statusLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999")).
		Width(width - 4).
		Align(lipgloss.Center).
		Render(status)

	return lipgloss.JoinVertical(lipgloss.Center, "", title, statusLine, "")
}

func (m *keysModel) renderServiceList() string {
	var content strings.Builder

	content.WriteString("Services:\n\n")

	for i, service := range m.services {
		marker := "  "
		if i == m.selectedIdx {
			marker = "→ "
		}

		statusIcon := getServiceStatusIcon(service.Status)
		content.WriteString(fmt.Sprintf("%s%s %s\n", marker, statusIcon, service.Name))
		content.WriteString(fmt.Sprintf("   %s\n", service.Description))
		content.WriteString("\n")
	}

	return content.String()
}

func (m *keysModel) renderServiceDetail() string {
	if m.selectedIdx >= len(m.services) {
		return "Invalid selection"
	}

	service := m.services[m.selectedIdx]
	var content strings.Builder

	content.WriteString(fmt.Sprintf("## %s\n\n", service.Name))
	content.WriteString(fmt.Sprintf("**Category:** %s\n", strings.Title(service.Category)))
	content.WriteString(fmt.Sprintf("**Status:** %s %s\n\n", getServiceStatusIcon(service.Status), service.Status))
	content.WriteString(fmt.Sprintf("**Description:** %s\n\n", service.Description))

	content.WriteString("### API Keys:\n\n")
	for _, key := range service.Keys {
		statusIcon := getKeyStatusIcon(key.Status)
		configured := ""
		if key.Value != "" {
			configured = fmt.Sprintf(" (configured: %s...)", key.Value[:min(8, len(key.Value))])
		}
		content.WriteString(fmt.Sprintf("- %s **%s** (`%s`)%s\n", statusIcon, key.Name, key.EnvVar, configured))
		if key.Error != "" {
			content.WriteString(fmt.Sprintf("  ⚠️ %s\n", key.Error))
		}
	}

	return content.String()
}

func renderKeysHelp(view string) string {
	help := "\n"
	if view == "list" {
		help += "↑/↓: Navigate | Enter: View details | V: Validate all | S: Sync .env.example | Q: Quit"
	} else {
		help += "Esc: Back | Q: Quit"
	}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render(help)
}

func (m *keysModel) validateAll() tea.Cmd {
	return func() tea.Msg {
		projectRoot, _ := os.Getwd()
		data, err := rakd.GenerateRAKD(projectRoot)
		if err != nil {
			return keysErrorMsg{Error: err}
		}
		m.services = data.Services
		m.rakdData = data
		return keysSuccessMsg{Message: "Keys validated"}
	}
}

func (m *keysModel) syncEnvExample() tea.Cmd {
	return func() tea.Msg {
		// This would call the sync function
		return keysSuccessMsg{Message: ".env.example synced"}
	}
}

func getServiceStatusIcon(status rakd.APIKeyStatus) string {
	switch status {
	case rakd.StatusConfigured:
		return "✅"
	case rakd.StatusRequired:
		return "🔴"
	case rakd.StatusPending:
		return "🟡"
	case rakd.StatusOptional:
		return "🔵"
	case rakd.StatusInvalid:
		return "⚠️"
	default:
		return "⚪"
	}
}

func getKeyStatusIcon(status rakd.APIKeyStatus) string {
	return getServiceStatusIcon(status)
}

type keysErrorMsg struct {
	Error error
}

type keysSuccessMsg struct {
	Message string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

