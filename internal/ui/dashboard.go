package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/github"
	"github.com/DoPlan-dev/CLI/internal/statistics"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#667eea")).
			Bold(true).
			Padding(0, 1)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#667eea")).
				Bold(true)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))
)

type DashboardModel struct {
	width      int
	height     int
	state      *models.State
	githubData *github.GitHubData
	config     *models.Config
	statistics *statistics.StatisticsMetrics

	// Views
	currentView string // "dashboard", "phases", "features", "github", "config", "stats"

	// Dashboard view
	overallProgress progress.Model
	phaseList       list.Model
	featureList     list.Model

	// Navigation
	selectedPhase   int
	selectedFeature int

	// Loading
	spinner spinner.Model
	loading bool
}

func NewDashboardModel() *DashboardModel {
	p := progress.New(progress.WithScaledGradient("#667eea", "#764ba2"))
	p.Width = 50

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#667eea"))

	return &DashboardModel{
		currentView:     "dashboard",
		overallProgress: p,
		spinner:         s,
		loading:         true,
	}
}

type loadDataMsg struct {
	state      *models.State
	githubData *github.GitHubData
	config     *models.Config
	statistics *statistics.StatisticsMetrics
	err        error
}

func loadDataCmd() tea.Msg {
	projectRoot, _ := os.Getwd()
	cfgMgr := config.NewManager(projectRoot)

	state, err := cfgMgr.LoadState()
	if err != nil {
		return loadDataMsg{err: err}
	}

	githubSync := github.NewGitHubSync(projectRoot)
	githubData, _ := githubSync.LoadData()

	cfg, _ := cfgMgr.LoadConfig()

	// Load statistics
	var stats *statistics.StatisticsMetrics
	collector := statistics.NewCollector(projectRoot)
	data, err := collector.Collect()
	if err == nil {
		projectStartDate := cfg.InstalledAt
		if projectStartDate.IsZero() {
			projectStartDate = time.Now()
		}
		calculator := statistics.NewCalculator(projectStartDate)
		stats = calculator.Calculate(data, state, githubData)
	}

	return loadDataMsg{
		state:      state,
		githubData: githubData,
		config:     cfg,
		statistics: stats,
	}
}

func (m *DashboardModel) Init() tea.Cmd {
	return tea.Batch(
		loadDataCmd,
		m.spinner.Tick,
	)
}

func (m *DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case loadDataMsg:
		if msg.err != nil {
			m.loading = false
			return m, nil
		}
		m.state = msg.state
		m.githubData = msg.githubData
		m.config = msg.config
		m.statistics = msg.statistics
		m.loading = false
		m.setupLists()
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.currentView = "dashboard"
			return m, nil
		case "2":
			m.currentView = "phases"
			return m, nil
		case "3":
			m.currentView = "features"
			return m, nil
		case "4":
			m.currentView = "github"
			return m, nil
		case "5":
			m.currentView = "config"
			return m, nil
		case "6":
			m.currentView = "stats"
			return m, nil
		case "r":
			m.loading = true
			return m, loadDataCmd
		}
	}

	// Delegate to current view
	if m.loading {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	switch m.currentView {
	case "dashboard":
		return m.updateDashboard(msg)
	case "phases":
		return m.updatePhases(msg)
	case "features":
		return m.updateFeatures(msg)
	case "github":
		return m.updateGitHub(msg)
	case "config":
		return m.updateConfig(msg)
	case "stats":
		return m.updateStats(msg)
	}

	return m, nil
}

func (m *DashboardModel) View() string {
	if m.loading {
		return m.renderLoading()
	}

	header := RenderBestDoPlanHeader(m.width, "dev")
	menu := m.renderMenu()

	var content string
	switch m.currentView {
	case "dashboard":
		content = m.renderDashboard()
	case "phases":
		content = m.renderPhases()
	case "features":
		content = m.renderFeatures()
	case "github":
		content = m.renderGitHub()
	case "config":
		content = m.renderConfig()
	case "stats":
		content = m.renderStats()
	}

	footer := m.renderFooter()

	body := lipgloss.JoinVertical(lipgloss.Left, menu, content)
	body = lipgloss.NewStyle().
		Width(m.width - 4).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer) - 5).
		Padding(1, 2).
		Render(body)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)
}

func (m *DashboardModel) renderLoading() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		RenderBestDoPlanHeader(m.width, "dev"),
		"",
		m.spinner.View()+" Loading...",
	)
}

func (m *DashboardModel) renderMenu() string {
	views := []string{"Dashboard", "Phases", "Features", "GitHub", "Config", "Stats"}
	menuItems := []string{}

	for i, view := range views {
		key := fmt.Sprintf("%d", i+1)
		style := normalItemStyle
		if m.currentView == strings.ToLower(view) {
			style = selectedItemStyle
		}
		menuItems = append(menuItems, style.Render(fmt.Sprintf("[%s] %s", key, view)))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, menuItems...) + "\n" + strings.Repeat("─", m.width-4)
}

func (m *DashboardModel) renderDashboard() string {
	if m.state == nil {
		return "No project data available"
	}

	var sections []string

	// Overall progress
	sections = append(sections, titleStyle.Render("Overall Progress"))
	progressBar := m.overallProgress.ViewAs(float64(m.state.Progress.Overall) / 100)
	sections = append(sections, fmt.Sprintf("%d%% %s", m.state.Progress.Overall, progressBar))
	sections = append(sections, "")

	// Phase summary
	sections = append(sections, titleStyle.Render("Phases"))
	if len(m.state.Phases) == 0 {
		sections = append(sections, "  No phases defined")
	} else {
		for _, phase := range m.state.Phases {
			progress := m.state.Progress.Phases[phase.ID]
			status := "○"
			if phase.Status == "complete" {
				status = "✓"
			} else if phase.Status == "in-progress" {
				status = "→"
			}
			sections = append(sections, fmt.Sprintf("  %s %s (%d%%)", status, phase.Name, progress))
		}
	}
	sections = append(sections, "")

	// Feature summary
	sections = append(sections, titleStyle.Render("Recent Features"))
	if len(m.state.Features) == 0 {
		sections = append(sections, "  No features defined")
	} else {
		count := 0
		for _, feature := range m.state.Features {
			if count >= 5 {
				break
			}
			status := "○"
			if feature.Status == "complete" {
				status = "✓"
			} else if feature.Status == "in-progress" {
				status = "→"
			}
			sections = append(sections, fmt.Sprintf("  %s %s (%d%%)", status, feature.Name, feature.Progress))
			count++
		}
	}

	return strings.Join(sections, "\n")
}

func (m *DashboardModel) renderPhases() string {
	if m.state == nil || len(m.state.Phases) == 0 {
		return "No phases defined"
	}

	var sections []string
	sections = append(sections, titleStyle.Render("Project Phases"))
	sections = append(sections, "")

	for _, phase := range m.state.Phases {
		progress := m.state.Progress.Phases[phase.ID]
		sections = append(sections, fmt.Sprintf("Phase: %s", phase.Name))
		sections = append(sections, fmt.Sprintf("  Status: %s", phase.Status))
		sections = append(sections, fmt.Sprintf("  Progress: %d%%", progress))
		if phase.Description != "" {
			sections = append(sections, fmt.Sprintf("  Description: %s", phase.Description))
		}
		sections = append(sections, "")
	}

	return strings.Join(sections, "\n")
}

func (m *DashboardModel) renderFeatures() string {
	if m.state == nil || len(m.state.Features) == 0 {
		return "No features defined"
	}

	var sections []string
	sections = append(sections, titleStyle.Render("Project Features"))
	sections = append(sections, "")

	for _, feature := range m.state.Features {
		sections = append(sections, fmt.Sprintf("Feature: %s", feature.Name))
		sections = append(sections, fmt.Sprintf("  Phase: %s", feature.Phase))
		sections = append(sections, fmt.Sprintf("  Status: %s", feature.Status))
		sections = append(sections, fmt.Sprintf("  Progress: %d%%", feature.Progress))
		if feature.Branch != "" {
			sections = append(sections, fmt.Sprintf("  Branch: %s", feature.Branch))
		}
		if feature.PR != nil {
			sections = append(sections, fmt.Sprintf("  PR: %s", feature.PR.URL))
		}
		sections = append(sections, "")
	}

	return strings.Join(sections, "\n")
}

func (m *DashboardModel) renderGitHub() string {
	if m.githubData == nil {
		return "No GitHub data available"
	}

	var sections []string
	sections = append(sections, titleStyle.Render("GitHub Activity"))
	sections = append(sections, "")

	sections = append(sections, fmt.Sprintf("Branches: %d", len(m.githubData.Branches)))
	sections = append(sections, fmt.Sprintf("Commits: %d", len(m.githubData.Commits)))
	sections = append(sections, fmt.Sprintf("Pull Requests: %d", len(m.githubData.PRs)))
	sections = append(sections, "")

	if len(m.githubData.PRs) > 0 {
		sections = append(sections, titleStyle.Render("Recent PRs"))
		for i, pr := range m.githubData.PRs {
			if i >= 5 {
				break
			}
			sections = append(sections, fmt.Sprintf("  #%d: %s [%s]", pr.Number, pr.Title, pr.Status))
		}
	}

	return strings.Join(sections, "\n")
}

func (m *DashboardModel) renderConfig() string {
	if m.config == nil {
		return "No configuration available"
	}

	var sections []string
	sections = append(sections, titleStyle.Render("Configuration"))
	sections = append(sections, "")
	sections = append(sections, fmt.Sprintf("IDE: %s", m.config.IDE))
	sections = append(sections, fmt.Sprintf("Version: %s", m.config.Version))
	sections = append(sections, "")
	sections = append(sections, titleStyle.Render("GitHub"))
	sections = append(sections, fmt.Sprintf("  Enabled: %v", m.config.GitHub.Enabled))
	sections = append(sections, fmt.Sprintf("  Auto Branch: %v", m.config.GitHub.AutoBranch))
	sections = append(sections, fmt.Sprintf("  Auto PR: %v", m.config.GitHub.AutoPR))
	sections = append(sections, "")
	sections = append(sections, titleStyle.Render("Checkpoints"))
	sections = append(sections, fmt.Sprintf("  Auto Feature: %v", m.config.Checkpoint.AutoFeature))
	sections = append(sections, fmt.Sprintf("  Auto Phase: %v", m.config.Checkpoint.AutoPhase))
	sections = append(sections, fmt.Sprintf("  Auto Complete: %v", m.config.Checkpoint.AutoComplete))

	return strings.Join(sections, "\n")
}

func (m *DashboardModel) renderStats() string {
	if m.statistics == nil {
		return "No statistics available. Run 'doplan stats' to generate statistics."
	}

	var sections []string
	sections = append(sections, titleStyle.Render("📊 Statistics"))

	// Velocity Metrics
	if m.statistics.Velocity != nil {
		sections = append(sections, "")
		sections = append(sections, titleStyle.Render("Velocity"))
		sections = append(sections, fmt.Sprintf("  Features/day:  %.2f", m.statistics.Velocity.FeaturesPerDay))
		sections = append(sections, fmt.Sprintf("  Commits/day:   %.2f", m.statistics.Velocity.CommitsPerDay))
		sections = append(sections, fmt.Sprintf("  Tasks/day:     %.2f", m.statistics.Velocity.TasksPerDay))
		sections = append(sections, fmt.Sprintf("  PRs/week:      %.2f", m.statistics.Velocity.PRsPerWeek))
	}

	// Completion Rates
	if m.statistics.Completion != nil {
		sections = append(sections, "")
		sections = append(sections, titleStyle.Render("Completion"))
		sections = append(sections, fmt.Sprintf("  Overall:       %d%%", m.statistics.Completion.Overall))
		sections = append(sections, fmt.Sprintf("  Tasks:         %d%%", m.statistics.Completion.Tasks))
		if len(m.statistics.Completion.Phases) > 0 {
			sections = append(sections, "  Phases:")
			for phaseID, progress := range m.statistics.Completion.Phases {
				sections = append(sections, fmt.Sprintf("    %s: %d%%", phaseID, progress))
			}
		}
	}

	// Time Metrics
	if m.statistics.Time != nil {
		sections = append(sections, "")
		sections = append(sections, titleStyle.Render("Time"))
		sections = append(sections, fmt.Sprintf("  Days since start: %d", m.statistics.Time.DaysSinceStart))
		if m.statistics.Time.AvgFeatureTime > 0 {
			sections = append(sections, fmt.Sprintf("  Avg feature time: %.1f days", m.statistics.Time.AvgFeatureTime))
		}
		if !m.statistics.Time.EstimatedCompletion.IsZero() {
			sections = append(sections, fmt.Sprintf("  Est. completion:  %s", m.statistics.Time.EstimatedCompletion.Format("2006-01-02")))
		}
	}

	// Quality Metrics
	if m.statistics.Quality != nil {
		sections = append(sections, "")
		sections = append(sections, titleStyle.Render("Quality"))
		sections = append(sections, fmt.Sprintf("  PR merge rate:    %.1f%%", m.statistics.Quality.PRMergeRate))
		sections = append(sections, fmt.Sprintf("  Checkpoint freq:  %.1f/week", m.statistics.Quality.CheckpointFrequency))
	}

	// Trends
	if m.statistics.Trends != nil {
		sections = append(sections, "")
		sections = append(sections, titleStyle.Render("Trends"))
		sections = append(sections, fmt.Sprintf("  Velocity:   %s", m.statistics.Trends.VelocityTrend))
		if m.statistics.Trends.VelocityChange != 0 {
			sections = append(sections, fmt.Sprintf("    Change: %.1f%%", m.statistics.Trends.VelocityChange))
		}
		sections = append(sections, fmt.Sprintf("  Completion: %s", m.statistics.Trends.CompletionTrend))
		if m.statistics.Trends.CompletionChange != 0 {
			sections = append(sections, fmt.Sprintf("    Change: %.1f%%", m.statistics.Trends.CompletionChange))
		}
		sections = append(sections, fmt.Sprintf("  Quality:    %s", m.statistics.Trends.QualityTrend))
	}

	return strings.Join(sections, "\n")
}

func (m *DashboardModel) updateStats(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Stats view is read-only, just handle navigation
	return m, nil
}

func (m *DashboardModel) renderFooter() string {
	help := helpStyle.Render("Press [1-6] to switch views | [r] to refresh | [q] to quit")
	return strings.Repeat("─", m.width-4) + "\n" + help
}

func (m *DashboardModel) setupLists() {
	// Setup phase list
	if m.state != nil && len(m.state.Phases) > 0 {
		items := []list.Item{}
		for _, phase := range m.state.Phases {
			items = append(items, phaseItem{phase: phase})
		}
		m.phaseList = list.New(items, list.NewDefaultDelegate(), m.width-10, m.height-10)
		m.phaseList.Title = "Phases"
	}

	// Setup feature list
	if m.state != nil && len(m.state.Features) > 0 {
		items := []list.Item{}
		for _, feature := range m.state.Features {
			items = append(items, featureItem{feature: feature})
		}
		m.featureList = list.New(items, list.NewDefaultDelegate(), m.width-10, m.height-10)
		m.featureList.Title = "Features"
	}
}

func (m *DashboardModel) updateDashboard(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *DashboardModel) updatePhases(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.phaseList.Items() != nil && len(m.phaseList.Items()) > 0 {
		var cmd tea.Cmd
		m.phaseList, cmd = m.phaseList.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *DashboardModel) updateFeatures(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.featureList.Items() != nil && len(m.featureList.Items()) > 0 {
		var cmd tea.Cmd
		m.featureList, cmd = m.featureList.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *DashboardModel) updateGitHub(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *DashboardModel) updateConfig(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// List items
type phaseItem struct {
	phase models.Phase
}

func (i phaseItem) FilterValue() string { return i.phase.Name }
func (i phaseItem) Title() string       { return i.phase.Name }
func (i phaseItem) Description() string {
	return fmt.Sprintf("%s - %d%%", i.phase.Status, 0)
}

type featureItem struct {
	feature models.Feature
}

func (i featureItem) FilterValue() string { return i.feature.Name }
func (i featureItem) Title() string       { return i.feature.Name }
func (i featureItem) Description() string {
	return fmt.Sprintf("%s - %d%%", i.feature.Status, i.feature.Progress)
}

