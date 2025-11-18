package wizard

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/context"
	doplanerror "github.com/DoPlan-dev/CLI/internal/error"
	"github.com/DoPlan-dev/CLI/internal/integration"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

// AdoptProjectWizard handles the project adoption wizard
type AdoptProjectWizard struct {
	projectRoot string
	model       *adoptProjectModel
}

// NewAdoptProjectWizard creates a new adopt project wizard
func NewAdoptProjectWizard(projectRoot string) *AdoptProjectWizard {
	return &AdoptProjectWizard{
		projectRoot: projectRoot,
		model:       newAdoptProjectModel(projectRoot),
	}
}

// Run starts the wizard
func (w *AdoptProjectWizard) Run() error {
	p := tea.NewProgram(w.model, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

type adoptScreen int

const (
	adoptScreenFound adoptScreen = iota
	adoptScreenAnalysis
	adoptScreenOptions
	adoptScreenGitHub
	adoptScreenIDE
	adoptScreenProgress
	adoptScreenPlanPreview
	adoptScreenConfirmation
)

type adoptProjectModel struct {
	width  int
	height int

	projectRoot string
	currentScreen adoptScreen

	// Project data
	analysis      *context.ProjectAnalysis
	adoptionOption string
	githubRepo    string
	ide           string

	// UI components
	textInput    textinput.Model
	optionsList  list.Model
	ideList      list.Model
	spinner      spinner.Model
	loading      bool
	loadingMsg   string

	// Analysis progress
	analysisStep int
	analysisSteps []string
	err          error
}

func newAdoptProjectModel(projectRoot string) *adoptProjectModel {
	options := []list.Item{
		optionItem{name: "analyze", desc: "Analyze & generate plan (recommended)"},
		optionItem{name: "import", desc: "Import existing docs"},
		optionItem{name: "fresh", desc: "Start fresh"},
	}

	optionsList := list.New(options, list.NewDefaultDelegate(), 0, 0)
	optionsList.Title = "Adoption Options"
	optionsList.SetShowStatusBar(false)
	optionsList.SetFilteringEnabled(false)

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

	ti := textinput.New()
	ti.Placeholder = "https://github.com/username/repo"
	ti.CharLimit = 200
	ti.Width = 50

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#667eea"))

	return &adoptProjectModel{
		projectRoot:   projectRoot,
		currentScreen: adoptScreenFound,
		textInput:     ti,
		optionsList:   optionsList,
		ideList:       ideList,
		spinner:       s,
		analysisSteps: []string{
			"Analyzing project structure",
			"Detecting phases and features",
			"Generating plan",
			"Setting up DoPlan structure",
		},
	}
}

func (m *adoptProjectModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		textinput.Blink,
		m.analyzeProject,
	)
}

func (m *adoptProjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.optionsList.SetWidth(msg.Width - 4)
		m.optionsList.SetHeight(msg.Height - 10)
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

	case analysisCompleteMsg:
		m.analysis = msg.analysis
		m.err = msg.err
		if m.err == nil {
			m.currentScreen = adoptScreenAnalysis
		}
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case analysisStepMsg:
		m.analysisStep = msg.step
		if msg.done {
			m.loading = false
			m.currentScreen = adoptScreenPlanPreview
			return m, nil
		}
		return m, m.nextAnalysisStep()

	case analysisErrorMsg:
		m.err = msg.err
		m.loading = false
		return m, nil
	}

	return m, nil
}

func (m *adoptProjectModel) updateScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case adoptScreenFound:
		return m.updateFound(msg)
	case adoptScreenAnalysis:
		return m.updateAnalysis(msg)
	case adoptScreenOptions:
		return m.updateOptions(msg)
	case adoptScreenGitHub:
		return m.updateGitHub(msg)
	case adoptScreenIDE:
		return m.updateIDE(msg)
	case adoptScreenProgress:
		return m.updateProgress(msg)
	case adoptScreenPlanPreview:
		return m.updatePlanPreview(msg)
	case adoptScreenConfirmation:
		return m.updateConfirmation(msg)
	}
	return m, nil
}

func (m *adoptProjectModel) updateFound(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", " ":
		if m.analysis != nil {
			m.currentScreen = adoptScreenAnalysis
		}
		return m, nil
	}
	return m, nil
}

func (m *adoptProjectModel) updateAnalysis(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", " ":
		m.currentScreen = adoptScreenOptions
		return m, nil
	case "esc":
		return m, tea.Quit
	}
	return m, nil
}

func (m *adoptProjectModel) updateOptions(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		selected := m.optionsList.SelectedItem()
		if selected != nil {
			m.adoptionOption = selected.(optionItem).name
			m.currentScreen = adoptScreenGitHub
		}
		return m, nil
	case "esc":
		m.currentScreen = adoptScreenAnalysis
		return m, nil
	}

	var cmd tea.Cmd
	m.optionsList, cmd = m.optionsList.Update(msg)
	return m, cmd
}

func (m *adoptProjectModel) updateGitHub(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		repo := strings.TrimSpace(m.textInput.Value())
		if repo == "" {
			m.githubRepo = ""
		} else {
			m.githubRepo = repo
		}
		m.currentScreen = adoptScreenIDE
		return m, nil
	case "esc":
		m.currentScreen = adoptScreenOptions
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	m.githubRepo = m.textInput.Value()
	return m, cmd
}

func (m *adoptProjectModel) updateIDE(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		selected := m.ideList.SelectedItem()
		if selected != nil {
			m.ide = selected.(ideItem).name
			m.currentScreen = adoptScreenProgress
			m.loading = true
			m.analysisStep = 0
			return m, m.startAdoption()
		}
		return m, nil
	case "esc":
		m.currentScreen = adoptScreenGitHub
		return m, nil
	}

	var cmd tea.Cmd
	m.ideList, cmd = m.ideList.Update(msg)
	return m, cmd
}

func (m *adoptProjectModel) updateProgress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Progress screen is read-only during loading
	return m, nil
}

func (m *adoptProjectModel) updatePlanPreview(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "y":
		m.currentScreen = adoptScreenConfirmation
		return m, nil
	case "n", "esc":
		return m, tea.Quit
	}
	return m, nil
}

func (m *adoptProjectModel) updateConfirmation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "q", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m *adoptProjectModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var view string
	switch m.currentScreen {
	case adoptScreenFound:
		view = m.renderFound()
	case adoptScreenAnalysis:
		view = m.renderAnalysis()
	case adoptScreenOptions:
		view = m.renderOptions()
	case adoptScreenGitHub:
		view = m.renderGitHub()
	case adoptScreenIDE:
		view = m.renderIDE()
	case adoptScreenProgress:
		view = m.renderProgress()
	case adoptScreenPlanPreview:
		view = m.renderPlanPreview()
	case adoptScreenConfirmation:
		view = m.renderConfirmation()
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, view)
}

func (m *adoptProjectModel) renderFound() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("Found Existing Project!")

	message := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("DoPlan detected an existing project in this directory.\n" +
			"Let's analyze it and set up DoPlan integration.")

	if m.analysis == nil {
		loading := lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			"",
			message,
			"",
			m.spinner.View()+" Analyzing project...",
		)
		return loading
	}

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		message,
		"",
		help,
	)
}

func (m *adoptProjectModel) renderAnalysis() string {
	if m.analysis == nil {
		return "No analysis data available"
	}

	var sections []string

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("Project Analysis Results")

	sections = append(sections, title)
	sections = append(sections, "")

	// Tech Stack
	if len(m.analysis.TechStack) > 0 {
		sections = append(sections, lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Bold(true).
			Render("Technology Stack:"))
		for _, tech := range m.analysis.TechStack {
			sections = append(sections, fmt.Sprintf("  • %s", tech))
		}
		sections = append(sections, "")
	}

	// Potential Phases
	if len(m.analysis.PotentialPhases) > 0 {
		sections = append(sections, lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Bold(true).
			Render("Detected Phases:"))
		for _, phase := range m.analysis.PotentialPhases {
			sections = append(sections, fmt.Sprintf("  • %s (%d features)", phase.Name, len(phase.Features)))
		}
		sections = append(sections, "")
	}

	// TODOs
	if len(m.analysis.TODOs) > 0 {
		sections = append(sections, lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Bold(true).
			Render(fmt.Sprintf("Found %d TODO items", len(m.analysis.TODOs))))
		sections = append(sections, "")
	}

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to continue, Esc to quit")

	sections = append(sections, help)

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m *adoptProjectModel) renderOptions() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("How would you like to proceed?"),
		"",
		m.optionsList.View(),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Render("Press Enter to select, Esc to go back"),
	)
}

func (m *adoptProjectModel) renderGitHub() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("GitHub Repository")

	warning := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f59e0b")).
		Render("⚠️  GitHub repository is required for full functionality")

	prompt := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("Enter GitHub repository URL (or press Enter to skip):")

	if m.textInput.Placeholder != "https://github.com/username/repo" {
		m.textInput.Placeholder = "https://github.com/username/repo"
		m.textInput.SetValue(m.githubRepo)
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
		warning,
		"",
		prompt,
		"",
		input,
		"",
		help,
	)
}

func (m *adoptProjectModel) renderIDE() string {
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

func (m *adoptProjectModel) renderProgress() string {
	var progress strings.Builder
	for i, step := range m.analysisSteps {
		status := "○"
		if i < m.analysisStep {
			status = "✓"
		} else if i == m.analysisStep {
			status = m.spinner.View()
		}
		progress.WriteString(fmt.Sprintf("%s %s\n", status, step))
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Render("Adopting Project..."),
		"",
		progress.String(),
	)
}

func (m *adoptProjectModel) renderPlanPreview() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Render("Generated Plan Preview")

	var content strings.Builder
	content.WriteString("The following structure will be created:\n\n")

	if m.analysis != nil && len(m.analysis.PotentialPhases) > 0 {
		for _, phase := range m.analysis.PotentialPhases {
			content.WriteString(fmt.Sprintf("Phase: %s\n", phase.Name))
			for _, feature := range phase.Features {
				content.WriteString(fmt.Sprintf("  • %s\n", feature))
			}
			content.WriteString("\n")
		}
	} else {
		content.WriteString("No phases detected. A fresh structure will be created.\n")
	}

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Render("Press Enter to confirm, N/Esc to cancel")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		content.String(),
		help,
	)
}

func (m *adoptProjectModel) renderConfirmation() string {
	success := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10b981")).
		Bold(true).
		Render("✅ Project adopted successfully!")

	nextSteps := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Render("Next steps:\n" +
			"1. Run 'doplan dashboard' to view your project\n" +
			"2. Review the generated plan\n" +
			"3. Use /Discuss in your IDE to refine your idea")

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

// Analysis logic
type analysisCompleteMsg struct {
	analysis *context.ProjectAnalysis
	err      error
}

type analysisStepMsg struct {
	step int
	done bool
}

type analysisErrorMsg struct {
	err error
}

func (m *adoptProjectModel) analyzeProject() tea.Msg {
	analyzer := context.NewAnalyzer(m.projectRoot)
	analysis, err := analyzer.Analyze()
	return analysisCompleteMsg{analysis: analysis, err: err}
}

func (m *adoptProjectModel) startAdoption() tea.Cmd {
	return tea.Batch(
		m.nextAnalysisStep(),
		m.spinner.Tick,
	)
}

func (m *adoptProjectModel) nextAnalysisStep() tea.Cmd {
	return func() tea.Msg {
		// Use a small delay to show spinner
		time.Sleep(500 * time.Millisecond)

		switch m.analysisStep {
		case 0:
			// Analyze project structure
			if err := m.createProjectStructure(); err != nil {
				return analysisErrorMsg{err: err}
			}
			return analysisStepMsg{step: 1, done: false}

		case 1:
			// Detect phases/features
			if err := m.detectPhasesFeatures(); err != nil {
				return analysisErrorMsg{err: err}
			}
			return analysisStepMsg{step: 2, done: false}

		case 2:
			// Generate plan
			if err := m.generatePlan(); err != nil {
				return analysisErrorMsg{err: err}
			}
			return analysisStepMsg{step: 3, done: false}

		case 3:
			// Setup DoPlan structure
			if err := m.setupDoPlan(); err != nil {
				return analysisErrorMsg{err: err}
			}
			return analysisStepMsg{step: 4, done: true}
		}

		return analysisStepMsg{step: m.analysisStep, done: true}
	}
}

func (m *adoptProjectModel) createProjectStructure() error {
	dirs := []string{
		filepath.Join(m.projectRoot, ".doplan"),
		filepath.Join(m.projectRoot, "doplan", "contracts"),
		filepath.Join(m.projectRoot, "doplan", "templates"),
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

func (m *adoptProjectModel) detectPhasesFeatures() error {
	// Phases/features detection is already done in analysis
	// This step can be used for additional processing if needed
	return nil
}

func (m *adoptProjectModel) generatePlan() error {
	// Generate plan based on analysis
	// For now, we'll create a basic state structure
	cfgMgr := config.NewManager(m.projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		// Create new state
		state = &models.State{
			Phases:   []models.Phase{},
			Features: []models.Feature{},
			Progress: models.Progress{
				Overall: 0,
				Phases:  make(map[string]int),
			},
		}

		// Convert analysis phases to state phases
		if m.analysis != nil {
			for i, phase := range m.analysis.PotentialPhases {
				phaseID := fmt.Sprintf("%02d-%s", i+1, strings.ToLower(strings.ReplaceAll(phase.Name, " ", "-")))
				state.Phases = append(state.Phases, models.Phase{
					ID:          phaseID,
					Name:        phase.Name,
					Status:      "todo",
					Description: fmt.Sprintf("Phase detected from existing structure"),
					Features:    []string{},
				})
			}
		}

		if err := cfgMgr.SaveState(state); err != nil {
			return err
		}
	}

	return nil
}

func (m *adoptProjectModel) setupDoPlan() error {
	// Setup GitHub
	if err := m.setupGitHub(); err != nil {
		return err
	}

	// Setup IDE
	if err := integration.SetupIDE(m.projectRoot, m.ide); err != nil {
		return err
	}

	return nil
}

func (m *adoptProjectModel) setupGitHub() error {
	cfg := config.NewConfig(m.ide)

	if m.githubRepo != "" && m.githubRepo != "skip" {
		cfg.GitHub.Enabled = true
	}

	return m.saveConfigYAML(cfg, m.githubRepo)
}

func (m *adoptProjectModel) saveConfigYAML(cfg *models.Config, githubRepo string) error {
	configPath := filepath.Join(m.projectRoot, ".doplan", "config.yaml")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return doplanerror.NewIOError("IO003", "Failed to create config directory").
			WithPath(filepath.Dir(configPath)).
			WithCause(err).
			WithSuggestion("Check file system permissions")
	}

	// Get project name from directory
	projectName := filepath.Base(m.projectRoot)

	// Convert to YAML structure
	yamlConfig := map[string]interface{}{
		"project": map[string]interface{}{
			"name":    projectName,
			"type":    "existing",
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
type optionItem struct {
	name string
	desc string
}

func (i optionItem) FilterValue() string { return i.name }
func (i optionItem) Title() string       { return i.name }
func (i optionItem) Description() string { return i.desc }
