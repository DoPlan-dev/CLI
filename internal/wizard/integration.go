package wizard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/integration"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))
)

type integrationModel struct {
	width     int
	height    int
	list      list.Model
	projectRoot string
	selectedIDE string
	status     string
	err        error
}

type integrationItem struct {
	id          string
	title       string
	description string
}

func (i integrationItem) FilterValue() string { return i.title }
func (i integrationItem) Title() string       { return i.title }
func (i integrationItem) Description() string { return i.description }

func RunIntegrationWizard(projectRoot string) error {
	items := []list.Item{
		integrationItem{
			id:          "cursor",
			title:       "Cursor",
			description: "AI-powered code editor with built-in chat",
		},
		integrationItem{
			id:          "vscode",
			title:       "VS Code + Copilot",
			description: "Visual Studio Code with GitHub Copilot",
		},
		integrationItem{
			id:          "gemini",
			title:       "Gemini CLI",
			description: "Google's Gemini CLI tool",
		},
		integrationItem{
			id:          "claude",
			title:       "Claude Code",
			description: "Anthropic's Claude Code",
		},
		integrationItem{
			id:          "codex",
			title:       "Codex CLI",
			description: "Codex command-line interface",
		},
		integrationItem{
			id:          "opencode",
			title:       "OpenCode",
			description: "OpenCode IDE integration",
		},
		integrationItem{
			id:          "qwen",
			title:       "Qwen Code",
			description: "Qwen Code IDE",
		},
		integrationItem{
			id:          "other",
			title:       "Other / Manual Setup",
			description: "View setup guide for other IDEs",
		},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select IDE for Integration"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.HelpStyle = helpStyle

	m := integrationModel{
		list:        l,
		projectRoot: projectRoot,
		status:      "Select an IDE to set up integration...",
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func (m integrationModel) Init() tea.Cmd {
	return nil
}

func (m integrationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(msg.Height - 6)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			selected := m.list.SelectedItem()
			if item, ok := selected.(integrationItem); ok {
				m.selectedIDE = item.id
				if item.id == "other" {
					m.status = "Opening setup guide..."
					return m, tea.Sequence(
						func() tea.Msg {
							// Show guide
							showOtherIDEGuide(m.projectRoot)
							return tea.Quit()
						},
					)
				}
				// Setup integration
				m.status = fmt.Sprintf("Setting up %s integration...", item.title)
				return m, tea.Sequence(
					func() tea.Msg {
						if err := setupIDEIntegration(m.projectRoot, item.id); err != nil {
							return errorMsg{err: err}
						}
						return successMsg{message: fmt.Sprintf("✅ %s integration set up successfully!", item.title)}
					},
				)
			}
		}

	case errorMsg:
		m.err = msg.err
		m.status = fmt.Sprintf("❌ Error: %v", msg.err)
		return m, tea.Sequence(
			tea.Printf("Press any key to continue..."),
			func() tea.Msg {
				return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")}
			},
		)

	case successMsg:
		m.status = msg.message
		return m, tea.Sequence(
			tea.Printf("\n%s\n\nPress any key to continue...", msg.message),
			func() tea.Msg {
				return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")}
			},
		)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m integrationModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	header := titleStyle.Render("⚙️  Setup AI/IDE Integration")
	help := helpStyle.Render("\n↑/↓: Navigate  Enter: Select  q: Quit")

	var status string
	if m.status != "" {
		status = "\n" + m.status + "\n"
	}

	var errorMsg string
	if m.err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ef4444")).
			Bold(true)
		errorMsg = "\n" + errorStyle.Render(fmt.Sprintf("Error: %v", m.err)) + "\n"
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		status,
		m.list.View(),
		errorMsg,
		help,
	)
}

type errorMsg struct {
	err error
}

type successMsg struct {
	message string
}

func setupIDEIntegration(projectRoot, ide string) error {
	// Check if already installed
	if !config.IsInstalled(projectRoot) {
		return fmt.Errorf("DoPlan is not installed in this project. Run 'doplan install' first")
	}

	// Setup IDE integration
	if err := integration.SetupIDE(projectRoot, ide); err != nil {
		return fmt.Errorf("failed to setup %s integration: %w", ide, err)
	}

	// Regenerate agents and rules if needed
	agentsGen := generators.NewAgentsGenerator(projectRoot)
	if err := agentsGen.Generate(); err != nil {
		return fmt.Errorf("failed to regenerate agents: %w", err)
	}

	rulesGen := generators.NewRulesGenerator(projectRoot)
	if err := rulesGen.Generate(); err != nil {
		return fmt.Errorf("failed to regenerate rules: %w", err)
	}

	// Verify integration
	if err := integration.VerifyIDE(projectRoot, ide); err != nil {
		return fmt.Errorf("verification failed: %w", err)
	}

	return nil
}

func showOtherIDEGuide(projectRoot string) {
	guidePath := filepath.Join(projectRoot, ".doplan", "guides", "IDE_INTEGRATION.md")
	
	// Read guide if it exists
	if content, err := os.ReadFile(guidePath); err == nil {
		fmt.Println(string(content))
	} else {
		// Show basic instructions
		fmt.Println("\n📚 IDE Integration Guide")
		fmt.Println("========================")
		fmt.Println("\nFor IDEs not listed, you can manually set up DoPlan integration:")
		fmt.Println("\n1. Create the following directories:")
		fmt.Println("   - .doplan/ai/agents/")
		fmt.Println("   - .doplan/ai/rules/")
		fmt.Println("   - .doplan/ai/commands/")
		fmt.Println("\n2. Run 'doplan install' to generate agents, rules, and commands")
		fmt.Println("\n3. Configure your IDE to reference these directories")
		fmt.Println("\n4. See .doplan/guides/IDE_INTEGRATION.md for detailed instructions")
		fmt.Println()
	}
}

