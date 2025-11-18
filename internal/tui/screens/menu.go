package screens

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	width  int
	height int
	list   list.Model
}

type menuItem struct {
	id          string
	title       string
	description string
	action      string
}

func (i menuItem) FilterValue() string { return i.title }
func (i menuItem) Title() string       { return i.title }
func (i menuItem) Description() string { return i.description }

func NewMenuModel() *MenuModel {
	items := []list.Item{
		menuItem{
			id:          "dashboard",
			title:       "📊 View Dashboard",
			description: "View project progress and statistics",
			action:      "dashboard",
		},
		menuItem{
			id:          "run",
			title:       "▶️  Run Dev Server",
			description: "Auto-detect and run development server",
			action:      "run",
		},
		menuItem{
			id:          "undo",
			title:       "↩️  Undo Last Action",
			description: "Revert the last DoPlan action",
			action:      "undo",
		},
		menuItem{
			id:          "deploy",
			title:       "🚀 Deploy Project",
			description: "Deploy to Vercel, Netlify, Railway, etc.",
			action:      "deploy",
		},
		menuItem{
			id:          "publish",
			title:       "📦 Publish Package",
			description: "Publish to npm, Homebrew, Scoop, Winget",
			action:      "publish",
		},
		menuItem{
			id:          "create",
			title:       "✨ Create New Project",
			description: "Start a new DoPlan project",
			action:      "create",
		},
		menuItem{
			id:          "security",
			title:       "🛡️  Run Security Scan",
			description: "Scan for vulnerabilities and security issues",
			action:      "security",
		},
		menuItem{
			id:          "fix",
			title:       "🩹 Auto-fix Issues",
			description: "AI-powered auto-fix for common issues",
			action:      "fix",
		},
		menuItem{
			id:          "discuss",
			title:       "💬 Discuss Idea",
			description: "Refine your project idea",
			action:      "discuss",
		},
		menuItem{
			id:          "generate",
			title:       "📝 Generate Documents",
			description: "Generate PRD, contracts, and documentation",
			action:      "generate",
		},
		menuItem{
			id:          "plan",
			title:       "🗺️  Create Plan",
			description: "Create phase and feature structure",
			action:      "plan",
		},
		menuItem{
			id:          "progress",
			title:       "📈 Update Progress",
			description: "Update progress tracking files",
			action:      "progress",
		},
		menuItem{
			id:          "keys",
			title:       "🔑 Manage API Keys",
			description: "Detect, validate, and manage API keys",
			action:      "keys",
		},
		menuItem{
			id:          "design",
			title:       "🎨 Apply Design / DPR",
			description: "Generate design system and tokens",
			action:      "design",
		},
		menuItem{
			id:          "integration",
			title:       "⚙️  Setup AI/IDE Integration",
			description: "Configure IDE integration (Cursor, VS Code, etc.)",
			action:      "integration",
		},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "DoPlan - Main Menu"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Padding(0, 1)

	return &MenuModel{
		list: l,
	}
}

func (m *MenuModel) Init() tea.Cmd {
	return nil
}

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			selected := m.list.SelectedItem()
			if selected != nil {
				item := selected.(menuItem)
				return m, func() tea.Msg {
					return MenuActionMsg{Action: item.action}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *MenuModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("↑/↓: Navigate | Enter: Select | /: Filter | q: Quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.list.View(),
		"",
		strings.Repeat("─", m.width-4),
		"",
		help,
	)
}

// MenuActionMsg is sent when a menu item is selected
type MenuActionMsg struct {
	Action string
}

// ErrorMsg is sent when an action fails
type ErrorMsg struct {
	Error error
}

// SuccessMsg is sent when an action succeeds
type SuccessMsg struct {
	Message string
	Action  string // The action that was completed (e.g., "plan_complete", "feature_implemented")
}

// RecommendationMsg is sent to display a recommended next step
type RecommendationMsg struct {
	Title       string
	Description string
}

