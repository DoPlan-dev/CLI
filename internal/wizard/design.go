package wizard

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/dpr"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type designModel struct {
	width       int
	height      int
	currentStep int
	totalSteps  int
	projectRoot string
	dprData     *dpr.DPRData
	err         error
}

type designStep int

const (
	stepDesignWelcome designStep = iota
	stepDesignQuestionnaire
	stepDesignGenerating
	stepDesignSuccess
)

// RunDesignWizard launches the design system (DPR) wizard
func RunDesignWizard() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	m := &designModel{
		currentStep: int(stepDesignWelcome),
		totalSteps:  4,
		projectRoot: projectRoot,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	return err
}

func (m *designModel) Init() tea.Cmd {
	return nil
}

func (m *designModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			switch designStep(m.currentStep) {
			case stepDesignWelcome:
				m.currentStep = int(stepDesignQuestionnaire)
				return m, nil
			case stepDesignQuestionnaire:
				// Run questionnaire
				m.currentStep = int(stepDesignGenerating)
				return m, m.runQuestionnaire()
			case stepDesignSuccess:
				return m, tea.Quit
			}
		case "esc":
			if m.currentStep > 0 {
				m.currentStep--
				return m, nil
			}
			return m, tea.Quit
		}
	}

	// Handle design wizard messages
	switch msg := msg.(type) {
	case designErrorMsg:
		m.err = msg.Error
		m.currentStep = int(stepDesignGenerating) // Stay on generating screen to show error
		return m, nil
	case designSuccessMsg:
		m.currentStep = int(stepDesignSuccess)
		return m, nil
	}

	return m, nil
}

func (m *designModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	header := renderDesignHeader(m.width)

	var content string
	switch designStep(m.currentStep) {
	case stepDesignWelcome:
		content = m.renderWelcome()
	case stepDesignQuestionnaire:
		content = m.renderQuestionnaire()
	case stepDesignGenerating:
		content = m.renderGenerating()
	case stepDesignSuccess:
		content = m.renderSuccess()
	}

	body := lipgloss.NewStyle().
		Width(m.width - 4).
		Height(m.height - lipgloss.Height(header) - 5).
		Padding(1, 2).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
	)
}

func renderDesignHeader(width int) string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("🎨 Design System (DPR) Generator")

	return lipgloss.JoinVertical(lipgloss.Center, "", title, "")
}

func (m *designModel) renderWelcome() string {
	text := "Welcome to the Design System Generator!\n\n"
	text += "This wizard will help you create a comprehensive Design Preferences & Requirements (DPR) document.\n\n"
	text += "You'll be asked about:\n"
	text += "  • Your target audience\n"
	text += "  • Design style and emotions\n"
	text += "  • Colors and typography\n"
	text += "  • Layout and components\n"
	text += "  • Animation preferences\n"
	text += "  • Accessibility requirements\n\n"
	text += "Press Enter to start the questionnaire..."
	return text
}

func (m *designModel) renderQuestionnaire() string {
	text := "Starting interactive questionnaire...\n\n"
	text += "You'll answer 20-30 questions about your design preferences.\n\n"
	text += "Press Enter to begin..."
	return text
}

func (m *designModel) renderGenerating() string {
	text := "Generating design system...\n\n"
	if m.err != nil {
		text += fmt.Sprintf("❌ Error: %v\n", m.err)
	} else {
		text += "⏳ Creating:\n"
		text += "  • DPR.md document\n"
		text += "  • design-tokens.json\n"
		text += "  • .doplan/ai/rules/design_rules.mdc\n"
	}
	return text
}

func (m *designModel) renderSuccess() string {
	text := "✅ Design system generated successfully!\n\n"
	text += "Files created:\n"
	text += "  • doplan/design/DPR.md\n"
	text += "  • doplan/design/design-tokens.json\n"
	text += "  • .doplan/ai/rules/design_rules.mdc\n\n"
	text += "💡 Next steps:\n"
	text += "  • Review DPR.md for your design requirements\n"
	text += "  • Use design-tokens.json in your code\n"
	text += "  • AI agents will follow design_rules.mdc\n\n"
	text += "Press Enter to exit..."
	return text
}

func (m *designModel) runQuestionnaire() tea.Cmd {
	return func() tea.Msg {
		// Run the questionnaire in a separate TUI
		// Note: This will quit the current TUI and start a new one
		// For better UX, we could integrate it directly, but for now this works
		dprData, err := dpr.RunQuestionnaire(m.projectRoot)
		if err != nil {
			// If user quit questionnaire, return to menu
			if err.Error() == "user quit" {
				return designErrorMsg{Error: fmt.Errorf("questionnaire cancelled")}
			}
			return designErrorMsg{Error: err}
		}
		m.dprData = dprData

		// Generate DPR.md
		dprGenerator := dpr.NewGenerator(m.projectRoot, dprData)
		if err := dprGenerator.Generate(); err != nil {
			return designErrorMsg{Error: fmt.Errorf("failed to generate DPR.md: %w", err)}
		}

		// Generate design-tokens.json
		tokenGenerator := dpr.NewTokenGenerator(m.projectRoot, dprData)
		if err := tokenGenerator.Generate(); err != nil {
			return designErrorMsg{Error: fmt.Errorf("failed to generate design-tokens.json: %w", err)}
		}

		// Generate design_rules.mdc
		rulesGenerator := dpr.NewCursorRulesGenerator(m.projectRoot, dprData)
		if err := rulesGenerator.Generate(); err != nil {
			return designErrorMsg{Error: fmt.Errorf("failed to generate design_rules.mdc: %w", err)}
		}

		return designSuccessMsg{}
	}
}

type designErrorMsg struct {
	Error error
}

type designSuccessMsg struct{}

