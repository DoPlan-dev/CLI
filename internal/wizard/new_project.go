package wizard

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/config"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/integration"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

// NewProjectWizard handles the new project creation wizard
type NewProjectWizard struct {
	projectRoot string
	model       *newProjectModel
}

// NewNewProjectWizard creates a new project wizard
func NewNewProjectWizard(projectRoot string) *NewProjectWizard {
	return &NewProjectWizard{
		projectRoot: projectRoot,
		model:       newNewProjectModel(),
	}
}

// Run starts the wizard
func (w *NewProjectWizard) Run() error {
	p := tea.NewProgram(w.model, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

type wizardScreen int

const (
	screenWelcome wizardScreen = iota
	screenProjectName
	screenTemplateChoice
	screenTemplate
	screenGitHubType
	screenGitHubUsername
	screenGitHubOrg
	screenGitHubRepoName
	screenGitHub
	screenIDE
	screenInstall
	screenSuccess
)

type newProjectModel struct {
	width  int
	height int

	currentScreen wizardScreen

	// Project data
	projectName    string
	useTemplate    bool
	template       string
	githubRepo     string
	githubRepoType string // "personal" or "organization"
	githubUsername string
	githubOrg      string
	githubRepoName string
	ide            string

	// UI components
	textInput          textinput.Model
	templateChoiceList list.Model
	templateList       list.Model
	githubTypeList     list.Model
	ideList            list.Model
	spinner            spinner.Model
	loading            bool
	loadingMsg         string

	// Installation progress
	installStep  int
	installSteps []string
	err          error
}

func newNewProjectModel() *newProjectModel {
	ti := textinput.New()
	ti.Placeholder = "my-awesome-project"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 50

	templates := []list.Item{
		templateItem{name: "website-frontend", desc: "01- Website (Frontend)"},
		templateItem{name: "website-admin", desc: "02- Website & Admin Dashboard (Frontend & Backend)"},
		templateItem{name: "web-app", desc: "03- Web Application"},
		templateItem{name: "mobile", desc: "04- Mobile Application"},
		templateItem{name: "micro-saas", desc: "05- Micro SaaS"},
		templateItem{name: "saas", desc: "06- SaaS"},
		templateItem{name: "web-game", desc: "07- Web Game"},
		templateItem{name: "cli", desc: "08- CLI"},
		templateItem{name: "chrome-ext", desc: "09- Chrome Extension"},
		templateItem{name: "ai-agent", desc: "10- AI Agent"},
		templateItem{name: "landing", desc: "Landing Page (marketing website)"},
		templateItem{name: "electron", desc: "Electron Desktop App"},
		templateItem{name: "api", desc: "API Service (REST/GraphQL backend)"},
		templateItem{name: "desktop-app", desc: "Desktop Application"},
		templateItem{name: "fullstack", desc: "Full Stack Application"},
		templateItem{name: "backend-service", desc: "Backend Service"},
		templateItem{name: "data-pipeline", desc: "Data Pipeline"},
		templateItem{name: "ml-project", desc: "Machine Learning Project"},
	}

	templateChoiceOptions := []list.Item{
		templateChoiceItem{name: "use-template", desc: "Use Template"},
		templateChoiceItem{name: "start-without", desc: "Start without Template"},
	}

	templateChoiceList := list.New(templateChoiceOptions, list.NewDefaultDelegate(), 0, 0)
	templateChoiceList.Title = "Choose Template Option"
	templateChoiceList.SetShowStatusBar(false)
	templateChoiceList.SetFilteringEnabled(false)

	templateList := list.New(templates, list.NewDefaultDelegate(), 0, 0)
	templateList.Title = "Select Project Template"
	templateList.SetShowStatusBar(false)
	templateList.SetFilteringEnabled(true)

	githubTypes := []list.Item{
		githubTypeItem{name: "personal", desc: "Personal"},
		githubTypeItem{name: "organization", desc: "Organization"},
	}

	githubTypeList := list.New(githubTypes, list.NewDefaultDelegate(), 0, 0)
	githubTypeList.Title = "Choose Repository Type"
	githubTypeList.SetShowStatusBar(false)
	githubTypeList.SetFilteringEnabled(false)

	ides := []list.Item{
		ideItem{name: "cursor", desc: "Cursor - AI-powered code editor"},
		ideItem{name: "kiro", desc: "Kiro - AI development environment"},
		ideItem{name: "copilot", desc: "VS Code + GitHub Copilot"},
		ideItem{name: "windsurf", desc: "Windsurf - AI code editor"},
		ideItem{name: "qoder", desc: "Qoder - AI coding assistant"},
		ideItem{name: "gemini", desc: "Gemini CLI"},
		ideItem{name: "claude", desc: "Claude CLI"},
		ideItem{name: "other", desc: "Other / Generic setup"},
	}

	ideList := list.New(ides, list.NewDefaultDelegate(), 0, 0)
	ideList.Title = "Select Your IDE / AI Tool"
	ideList.SetShowStatusBar(false)
	ideList.SetFilteringEnabled(true)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#667eea"))

	return &newProjectModel{
		currentScreen:      screenWelcome,
		textInput:          ti,
		templateChoiceList: templateChoiceList,
		templateList:       templateList,
		githubTypeList:     githubTypeList,
		ideList:            ideList,
		spinner:            s,
		installSteps: []string{
			"Creating project structure",
			"Setting up GitHub integration",
			"Configuring IDE integration",
			"Generating agents, rules, and commands",
			"Generating initial files",
		},
	}
}

func (m *newProjectModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		textinput.Blink,
	)
}

func (m *newProjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.templateChoiceList.SetWidth(msg.Width - 4)
		m.templateChoiceList.SetHeight(msg.Height - 10)
		m.templateList.SetWidth(msg.Width - 4)
		m.templateList.SetHeight(msg.Height - 10)
		m.githubTypeList.SetWidth(msg.Width - 4)
		m.githubTypeList.SetHeight(msg.Height - 10)
		m.ideList.SetWidth(msg.Width - 4)
		m.ideList.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		if m.loading {
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		return m.updateScreen(msg)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case installStepMsg:
		m.installStep = msg.step
		if msg.done {
			m.loading = false
			m.currentScreen = screenSuccess
			return m, nil
		}
		return m, m.nextInstallStep()

	case installErrorMsg:
		m.err = msg.err
		m.loading = false
		return m, nil
	}

	return m, nil
}

func (m *newProjectModel) updateScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case screenWelcome:
		return m.updateWelcome(msg)
	case screenProjectName:
		return m.updateProjectName(msg)
	case screenTemplateChoice:
		return m.updateTemplateChoice(msg)
	case screenTemplate:
		return m.updateTemplate(msg)
	case screenGitHubType:
		return m.updateGitHubType(msg)
	case screenGitHubUsername:
		return m.updateGitHubUsername(msg)
	case screenGitHubOrg:
		return m.updateGitHubOrg(msg)
	case screenGitHubRepoName:
		return m.updateGitHubRepoName(msg)
	case screenGitHub:
		return m.updateGitHub(msg)
	case screenIDE:
		return m.updateIDE(msg)
	case screenSuccess:
		return m.updateSuccess(msg)
	}
	return m, nil
}

func (m *newProjectModel) updateWelcome(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", " ":
		m.currentScreen = screenProjectName
		return m, textinput.Blink
	}
	return m, nil
}

func (m *newProjectModel) updateTemplateChoice(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		selected := m.templateChoiceList.SelectedItem()
		if selected != nil {
			choice := selected.(templateChoiceItem).name
			if choice == "use-template" {
				m.useTemplate = true
				m.currentScreen = screenTemplate
			} else {
				m.useTemplate = false
				m.template = ""
				m.currentScreen = screenGitHub
			}
		}
		return m, nil
	case "esc":
		m.currentScreen = screenProjectName
		return m, nil
	}

	var cmd tea.Cmd
	m.templateChoiceList, cmd = m.templateChoiceList.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateProjectName(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if strings.TrimSpace(m.textInput.Value()) == "" {
			return m, nil
		}
		m.projectName = strings.TrimSpace(m.textInput.Value())
		m.currentScreen = screenTemplateChoice
		return m, nil
	case "esc":
		m.currentScreen = screenWelcome
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateTemplate(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		selected := m.templateList.SelectedItem()
		if selected != nil {
			m.template = selected.(templateItem).name
			m.currentScreen = screenGitHubType
		}
		return m, nil
	case "esc":
		m.currentScreen = screenTemplateChoice
		return m, nil
	}

	var cmd tea.Cmd
	m.templateList, cmd = m.templateList.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateGitHubType(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		selected := m.githubTypeList.SelectedItem()
		if selected != nil {
			m.githubRepoType = selected.(githubTypeItem).name
			if m.githubRepoType == "personal" {
				m.currentScreen = screenGitHubUsername
			} else {
				m.currentScreen = screenGitHubOrg
			}
			m.textInput.SetValue("")
			m.textInput.Focus()
		}
		return m, textinput.Blink
	case "esc":
		if m.useTemplate {
			m.currentScreen = screenTemplate
		} else {
			m.currentScreen = screenTemplateChoice
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.githubTypeList, cmd = m.githubTypeList.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateGitHubUsername(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		username := strings.TrimSpace(m.textInput.Value())
		if username == "" {
			return m, nil
		}
		m.githubUsername = username
		m.currentScreen = screenGitHubRepoName
		m.textInput.SetValue("")
		m.textInput.Placeholder = "repository-name"
		return m, textinput.Blink
	case "esc":
		m.currentScreen = screenGitHubType
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateGitHubOrg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		org := strings.TrimSpace(m.textInput.Value())
		if org == "" {
			return m, nil
		}
		m.githubOrg = org
		m.currentScreen = screenGitHubRepoName
		m.textInput.SetValue("")
		m.textInput.Placeholder = "repository-name"
		return m, textinput.Blink
	case "esc":
		m.currentScreen = screenGitHubType
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateGitHubRepoName(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		repoName := strings.TrimSpace(m.textInput.Value())
		if repoName == "" {
			return m, nil
		}
		m.githubRepoName = repoName
		// Build full URL
		if m.githubRepoType == "personal" {
			m.githubRepo = fmt.Sprintf("https://github.com/%s/%s", m.githubUsername, m.githubRepoName)
		} else {
			m.githubRepo = fmt.Sprintf("https://github.com/%s/%s", m.githubOrg, m.githubRepoName)
		}
		m.currentScreen = screenGitHub
		return m, nil
	case "esc":
		if m.githubRepoType == "personal" {
			m.currentScreen = screenGitHubUsername
		} else {
			m.currentScreen = screenGitHubOrg
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateGitHub(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.currentScreen = screenIDE
		return m, nil
	case "esc":
		m.currentScreen = screenGitHubRepoName
		return m, nil
	}
	return m, nil
}

func (m *newProjectModel) updateIDE(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		selected := m.ideList.SelectedItem()
		if selected != nil {
			m.ide = selected.(ideItem).name
			m.currentScreen = screenInstall
			m.loading = true
			m.installStep = 0
			return m, m.startInstallation()
		}
		return m, nil
	case "esc":
		m.currentScreen = screenGitHub
		return m, nil
	}

	var cmd tea.Cmd
	m.ideList, cmd = m.ideList.Update(msg)
	return m, cmd
}

func (m *newProjectModel) updateSuccess(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "q", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m *newProjectModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Always render header
	header := m.renderHeader()

	var view string
	switch m.currentScreen {
	case screenWelcome:
		view = m.renderWelcome()
	case screenProjectName:
		view = m.renderProjectName()
	case screenTemplateChoice:
		view = m.renderTemplateChoice()
	case screenTemplate:
		view = m.renderTemplate()
	case screenGitHubType:
		view = m.renderGitHubType()
	case screenGitHubUsername:
		view = m.renderGitHubUsername()
	case screenGitHubOrg:
		view = m.renderGitHubOrg()
	case screenGitHubRepoName:
		view = m.renderGitHubRepoName()
	case screenGitHub:
		view = m.renderGitHub()
	case screenIDE:
		view = m.renderIDE()
	case screenInstall:
		view = m.renderInstall()
	case screenSuccess:
		view = m.renderSuccess()
	}

	// Combine header and content
	content := lipgloss.JoinVertical(lipgloss.Center, header, "", view)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

func (m *newProjectModel) renderHeader() string {
	topBorder := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Width(m.width).
		Render("╔" + strings.Repeat("═", m.width-2) + "╗")

	logo := `
  ██████╗░░█████╗░██████╗░██╗░░░░░░█████╗░███╗░░██╗
  ██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔══██╗████╗░██║
  ██║░░██║██║░░██║██████╔╝██║░░░░░███████║██╔██╗██║
  ██║░░██║██║░░██║██╔═══╝░██║░░░░░██╔══██║██║╚████║
  ██████╔╝╚█████╔╝██║░░░░░███████╗██║░░██║██║░╚███║
  ╚═════╝░░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚══╝
`

	styledLogo := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Width(m.width - 4).
		Align(lipgloss.Center).
		Render(logo)

	bottomBorder := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Width(m.width).
		Render("╚" + strings.Repeat("═", m.width-2) + "╝")

	return lipgloss.JoinVertical(lipgloss.Center,
		topBorder,
		"",
		styledLogo,
		"",
		bottomBorder,
	)
}

func (m *newProjectModel) renderWelcome() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Align(lipgloss.Center).
		Width(m.width - 4).
		Render("Welcome to DoPlan!")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999")).
		Align(lipgloss.Center).
		Width(m.width - 4).
		Render("Let's set up your new project")

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Align(lipgloss.Center).
		Width(m.width - 4).
		Render("Press Enter to continue")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		subtitle,
		"",
		help,
	)
}

func (m *newProjectModel) renderProjectName() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("Project Name")

	prompt := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("Enter your project name:")

	input := m.textInput.View()

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue, Esc to go back")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		prompt,
		"",
		input,
		"",
		help,
	)
}

func (m *newProjectModel) renderTemplateChoice() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("Choose Template Option"),
		"",
		m.templateChoiceList.View(),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Render("Press Enter to select, Esc to go back"),
	)
}

func (m *newProjectModel) renderTemplate() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("Select Project Template"),
		"",
		m.templateList.View(),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Render("Press Enter to select, Esc to go back"),
	)
}

func (m *newProjectModel) renderGitHubType() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("GitHub Repository Setup"),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f59e0b")).
			Render("⚠️  GitHub repository is required for full functionality"),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Render("Choose from: Personal / Organization"),
		"",
		m.githubTypeList.View(),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Render("Press Enter to select, Esc to go back"),
	)
}

func (m *newProjectModel) renderGitHubUsername() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("GitHub Username")

	prompt := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("GitHub Username:")

	if m.textInput.Placeholder != "username" {
		m.textInput.Placeholder = "username"
		m.textInput.SetValue(m.githubUsername)
		m.textInput.Focus()
	}
	input := m.textInput.View()

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue, Esc to go back")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		prompt,
		"",
		input,
		"",
		help,
	)
}

func (m *newProjectModel) renderGitHubOrg() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("GitHub Organization")

	prompt := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("GitHub Organization:")

	if m.textInput.Placeholder != "organization" {
		m.textInput.Placeholder = "organization"
		m.textInput.SetValue(m.githubOrg)
		m.textInput.Focus()
	}
	input := m.textInput.View()

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue, Esc to go back")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		prompt,
		"",
		input,
		"",
		help,
	)
}

func (m *newProjectModel) renderGitHubRepoName() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("Repository Name")

	prompt := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("Repository Name:")

	if m.textInput.Placeholder != "repository-name" {
		m.textInput.Placeholder = "repository-name"
		m.textInput.SetValue(m.githubRepoName)
		m.textInput.Focus()
	}
	input := m.textInput.View()

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue, Esc to go back")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		prompt,
		"",
		input,
		"",
		help,
	)
}

func (m *newProjectModel) renderGitHub() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("GitHub Repository")

	urlStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10b981")).
		Bold(true).
		Render(fmt.Sprintf("Repository URL: %s", m.githubRepo))

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue, Esc to go back")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		urlStyle,
		"",
		help,
	)
}

func (m *newProjectModel) renderIDE() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("Select Your IDE / AI Tool"),
		"",
		m.ideList.View(),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Render("Press Enter to select, Esc to go back"),
	)
}

func (m *newProjectModel) renderInstall() string {
	var progress strings.Builder
	for i, step := range m.installSteps {
		status := "○"
		if i < m.installStep {
			status = "✓"
		} else if i == m.installStep {
			status = m.spinner.View()
		}
		progress.WriteString(fmt.Sprintf("%s %s\n", status, step))
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("Installing DoPlan..."),
		"",
		progress.String(),
	)
}

func (m *newProjectModel) renderSuccess() string {
	success := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10b981")).
		Bold(true).
		Render("✅ DoPlan installed successfully!")

	nextSteps := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("Next steps:\n" +
			"1. Run 'doplan dashboard' to view your project\n" +
			"2. Use /Discuss in your IDE to start refining your idea\n" +
			"3. Use /Generate to create project documentation")

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter or Q to exit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		success,
		"",
		nextSteps,
		"",
		help,
	)
}

// Installation logic
type installStepMsg struct {
	step int
	done bool
}

type installErrorMsg struct {
	err error
}

func (m *newProjectModel) startInstallation() tea.Cmd {
	return tea.Batch(
		m.nextInstallStep(),
		m.spinner.Tick,
	)
}

func (m *newProjectModel) nextInstallStep() tea.Cmd {
	return func() tea.Msg {
		// Use a small delay to show spinner
		time.Sleep(500 * time.Millisecond)

		projectRoot, _ := os.Getwd()

		switch m.installStep {
		case 0:
			// Create project structure
			if err := m.createProjectStructure(projectRoot); err != nil {
				return installErrorMsg{err: err}
			}
			return installStepMsg{step: 1, done: false}

		case 1:
			// Setup GitHub
			if err := m.setupGitHub(projectRoot); err != nil {
				return installErrorMsg{err: err}
			}
			return installStepMsg{step: 2, done: false}

		case 2:
			// Setup IDE
			if err := m.setupIDE(projectRoot); err != nil {
				return installErrorMsg{err: err}
			}
			return installStepMsg{step: 3, done: false}

		case 3:
			// Generate agents, rules, and commands
			if err := m.generateAgentsAndRules(projectRoot); err != nil {
				return installErrorMsg{err: err}
			}
			return installStepMsg{step: 4, done: false}

		case 4:
			// Generate initial files
			if err := m.generateInitialFiles(projectRoot); err != nil {
				return installErrorMsg{err: err}
			}
			return installStepMsg{step: 5, done: true}
		}

		return installStepMsg{step: m.installStep, done: true}
	}
}

func (m *newProjectModel) createProjectStructure(projectRoot string) error {
	// Only create .doplan structure - doplan/ folder will be created when phases are added
	dirs := []string{
		filepath.Join(projectRoot, ".doplan", "ai", "agents"),
		filepath.Join(projectRoot, ".doplan", "ai", "rules"),
		filepath.Join(projectRoot, ".doplan", "ai", "commands"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return doplanerror.NewIOError("IO003", "Failed to create directory").
				WithPath(dir).
				WithCause(err).
				WithSuggestion("Check file system permissions")
		}
	}

	return nil
}

func (m *newProjectModel) setupGitHub(projectRoot string) error {
	cfg := config.NewConfig(m.ide)

	if m.githubRepo != "" && m.githubRepo != "skip" {
		cfg.GitHub.Enabled = true
	}

	// Save config with GitHub repo in YAML
	return m.saveConfigYAML(projectRoot, cfg, m.githubRepo)
}

func (m *newProjectModel) setupIDE(projectRoot string) error {
	return integration.SetupIDE(projectRoot, m.ide)
}

func (m *newProjectModel) generateAgentsAndRules(projectRoot string) error {
	// Generate agents
	agentsGen := generators.NewAgentsGenerator(projectRoot)
	if err := agentsGen.Generate(); err != nil {
		return fmt.Errorf("failed to generate agents: %w", err)
	}

	// Generate rules
	rulesGen := generators.NewRulesGenerator(projectRoot)
	if err := rulesGen.Generate(); err != nil {
		return fmt.Errorf("failed to generate rules: %w", err)
	}

	// Generate commands
	commandsGen := generators.NewCommandsGenerator(projectRoot)
	if err := commandsGen.Generate(); err != nil {
		return fmt.Errorf("failed to generate commands: %w", err)
	}

	return nil
}

func (m *newProjectModel) generateInitialFiles(projectRoot string) error {
	// Generate initial dashboard
	cfgMgr := config.NewManager(projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		// Create new state if it doesn't exist
		state = &models.State{
			Phases:   []models.Phase{},
			Features: []models.Feature{},
			Progress: models.Progress{
				Overall: 0,
				Phases:  make(map[string]int),
			},
		}
		if err := cfgMgr.SaveState(state); err != nil {
			return err
		}
	}

	// Generate dashboard.json will be done by dashboard generator
	// For now, just ensure basic structure exists
	return nil
}

func (m *newProjectModel) saveConfigYAML(projectRoot string, cfg *models.Config, githubRepo string) error {
	configPath := filepath.Join(projectRoot, ".doplan", "config.yaml")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return doplanerror.NewIOError("IO003", "Failed to create config directory").
			WithPath(filepath.Dir(configPath)).
			WithCause(err).
			WithSuggestion("Check file system permissions")
	}

	// Convert to YAML structure
	yamlConfig := map[string]interface{}{
		"project": map[string]interface{}{
			"name":    m.projectName,
			"type":    m.template,
			"version": cfg.Version,
			"ide":     cfg.IDE,
		},
		"github": map[string]interface{}{
			"repository": githubRepo,
			"enabled":    githubRepo != "" && githubRepo != "skip",
			"autoBranch": cfg.GitHub.AutoBranch,
			"autoPR":     cfg.GitHub.AutoPR,
		},
		"design": map[string]interface{}{
			"hasPreferences": false,
			"tokensPath":     "doplan/design/design-tokens.json",
		},
		"security": map[string]interface{}{
			"lastScan": nil,
			"autoFix":  false,
		},
		"apis": map[string]interface{}{
			"configured": []string{},
			"required":   []string{},
		},
		"tui": map[string]interface{}{
			"theme":      "default",
			"animations": true,
		},
	}

	data, err := yaml.Marshal(yamlConfig)
	if err != nil {
		return doplanerror.NewConfigError("CFG002", "Failed to marshal YAML config").
			WithCause(err).
			WithSuggestion("Check YAML configuration structure")
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}

	return nil
}

// List items
type templateItem struct {
	name string
	desc string
}

func (i templateItem) FilterValue() string { return i.name }
func (i templateItem) Title() string       { return i.name }
func (i templateItem) Description() string { return i.desc }

type templateChoiceItem struct {
	name string
	desc string
}

func (i templateChoiceItem) FilterValue() string { return i.name }
func (i templateChoiceItem) Title() string       { return i.name }
func (i templateChoiceItem) Description() string { return i.desc }

type githubTypeItem struct {
	name string
	desc string
}

func (i githubTypeItem) FilterValue() string { return i.name }
func (i githubTypeItem) Title() string       { return i.name }
func (i githubTypeItem) Description() string { return i.desc }

type ideItem struct {
	name string
	desc string
}

func (i ideItem) FilterValue() string { return i.name }
func (i ideItem) Title() string       { return i.name }
func (i ideItem) Description() string { return i.desc }
