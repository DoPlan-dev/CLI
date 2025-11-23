package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/doplan/cli/internal/generator"
	"github.com/doplan/cli/pkg/models"
)

// Color definitions matching the design system
var (
	// Primary colors
	purple = lipgloss.Color("#A855F7")
	pink   = lipgloss.Color("#EC4899")

	// Status colors
	green  = lipgloss.Color("#10B981")
	red    = lipgloss.Color("#EF4444")
	blue   = lipgloss.Color("#3B82F6")
	yellow = lipgloss.Color("#F59E0B")
	white  = lipgloss.Color("#FFFFFF")

	// Neutral colors
	gray800 = lipgloss.Color("#1F2937")
	gray900 = lipgloss.Color("#111827")
	gray600 = lipgloss.Color("#4B5563")
	gray400 = lipgloss.Color("#9CA3AF")
)

// StepStatus represents the status of a generation step
type StepStatus int

const (
	stepPending StepStatus = iota
	stepInProgress
	stepCompleted
	stepFailed
)

// GenerationStep represents a single step in the project generation process
type GenerationStep struct {
	Name   string
	Status StepStatus
}

// wizardState represents the current state of the wizard
type wizardState int

const (
	stateWelcome wizardState = iota
	stateProjectName
	stateIDESelection
	stateGenerating
	stateSuccess
	stateError
)

// Model represents the wizard model
type Model struct {
	state            wizardState
	width            int
	textInput        textinput.Model
	projectName      string
	validationErr    string
	selectedIDEIndex int
	selectedIDE      string
	steps            []GenerationStep
	spinnerFrame     int
	errorMessage     string
	errorSuggestion  string
	previousState    wizardState   // Store previous state for recovery
	stateHistory     []wizardState // History for back navigation
	generationDone   bool          // Track if generation has completed
	generationErr    error         // Store generation error if any
	generationChan   chan tea.Msg  // Channel for generation results
}

// generationCompleteMsg is sent when generation completes
type generationCompleteMsg struct {
	err error
}

// stepProgressMsg is sent when a step completes
type stepProgressMsg struct {
	stepIndex int
}

// IDEInfo contains information about an IDE option
type IDEInfo struct {
	Name        string
	Description string
	Recommended bool
}

// InitialModel returns the initial model for the wizard
func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "my-awesome-project"
	ti.CharLimit = 50
	ti.Width = 40
	ti.Focus()

	return Model{
		state:            stateWelcome,
		width:            80, // Default width
		textInput:        ti,
		projectName:      "",
		selectedIDEIndex: 0, // Default to first IDE (Cursor)
		selectedIDE:      "",
		steps:            getGenerationSteps(),
		spinnerFrame:     0,
		errorMessage:     "",
		errorSuggestion:  "",
		previousState:    stateWelcome,
		stateHistory:     []wizardState{stateWelcome},
	}
}

// getGenerationSteps returns the list of generation steps
func getGenerationSteps() []GenerationStep {
	return []GenerationStep{
		{Name: "Creating directory structure", Status: stepPending},
		{Name: "Generating AI agents", Status: stepPending},
		{Name: "Extracting rules library", Status: stepPending},
		{Name: "Creating GitHub workflows", Status: stepPending},
		{Name: "Setting up boilerplate", Status: stepPending},
		{Name: "Generating documentation", Status: stepPending},
	}
}

// tickSpinner is a command that sends a message to update the spinner animation
// Returns a channel that Bubble Tea will receive from - this is non-blocking
func tickSpinner() tea.Msg {
	return time.After(time.Millisecond * 100)
}

// getIDEInfo returns information about all supported IDEs
func getIDEInfo() []IDEInfo {
	return []IDEInfo{
		{
			Name:        "Cursor",
			Description: "AI-powered code editor with deep GitHub Copilot integration",
			Recommended: true,
		},
		{
			Name:        "Claude Code",
			Description: "Anthropic's AI coding assistant with advanced reasoning",
			Recommended: true,
		},
		{
			Name:        "Antigravity",
			Description: "Lightweight AI coding assistant with fast performance",
			Recommended: false,
		},
		{
			Name:        "Windsurf",
			Description: "Modern IDE with built-in AI pair programming",
			Recommended: false,
		},
		{
			Name:        "Cline",
			Description: "AI-powered development environment with code completion",
			Recommended: false,
		},
		{
			Name:        "OpenCode",
			Description: "Open-source AI coding assistant with extensible plugins",
			Recommended: false,
		},
	}
}

// canTransitionTo checks if a state transition is valid
func canTransitionTo(from, to wizardState) bool {
	// Define valid state transitions
	validTransitions := map[wizardState][]wizardState{
		stateWelcome:      {stateProjectName, stateError},
		stateProjectName:  {stateIDESelection, stateError},
		stateIDESelection: {stateGenerating, stateError},
		stateGenerating:   {stateSuccess, stateError},
		stateSuccess:      {},                                                                   // Terminal state
		stateError:        {stateWelcome, stateProjectName, stateIDESelection, stateGenerating}, // Can recover to any previous state
	}

	allowed, exists := validTransitions[from]
	if !exists {
		return false
	}

	// Check if 'to' is in the allowed list
	for _, allowedState := range allowed {
		if allowedState == to {
			return true
		}
	}

	return false
}

// transitionToState safely transitions to a new state
func (m *Model) transitionToState(newState wizardState) bool {
	if canTransitionTo(m.state, newState) {
		// Add current state to history (except for error state which doesn't count)
		if m.state != stateError && newState != stateError {
			m.stateHistory = append(m.stateHistory, m.state)
		}
		m.state = newState
		return true
	}
	return false
}

// goBack navigates to the previous state in history
func (m *Model) goBack() bool {
	if len(m.stateHistory) > 1 {
		// Remove current state from history
		m.stateHistory = m.stateHistory[:len(m.stateHistory)-1]
		// Get previous state
		previousState := m.stateHistory[len(m.stateHistory)-1]
		m.state = previousState
		return true
	}
	return false
}

// getCurrentStepNumber returns the current step number (1-based)
func (m Model) getCurrentStepNumber() int {
	switch m.state {
	case stateWelcome:
		return 0
	case stateProjectName:
		return 1
	case stateIDESelection:
		return 2
	case stateGenerating:
		return 3
	case stateSuccess:
		return 4
	default:
		return 0
	}
}

// getTotalSteps returns the total number of steps
func getTotalSteps() int {
	return 4 // Welcome, Project Name, IDE Selection, Generating (Success is terminal)
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	// Start spinner animation if in generating state
	if m.state == stateGenerating {
		return tickSpinner
	}
	return nil
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "r":
			// Retry: go back to previous state
			if m.state == stateError {
				m.state = m.previousState
				m.errorMessage = ""
				m.errorSuggestion = ""
				return m, nil
			}
			return m, nil

		case "b":
			// Go back: return to welcome screen
			if m.state == stateError {
				m.state = stateWelcome
				m.errorMessage = ""
				m.errorSuggestion = ""
				m.previousState = stateWelcome
				m.stateHistory = []wizardState{stateWelcome}
				return m, nil
			}
			return m, nil

		case "esc", "backspace":
			// Back navigation: go to previous state
			if m.state != stateWelcome && m.state != stateError && m.state != stateGenerating && m.state != stateSuccess {
				if m.goBack() {
					// Reset state-specific data when going back
					if m.state == stateProjectName {
						m.textInput.Focus()
						return m, textinput.Blink
					}
					return m, nil
				}
			}
			return m, nil

		case "enter":
			if m.state == stateWelcome {
				// Transition to project name input
				if m.transitionToState(stateProjectName) {
					m.textInput.Focus()
					return m, textinput.Blink
				}
				return m, nil
			} else if m.state == stateProjectName {
				// Validate and proceed if valid
				if m.validateProjectName(m.textInput.Value()) {
					m.projectName = m.textInput.Value()
					// Transition to IDE selection
					if m.transitionToState(stateIDESelection) {
						return m, nil
					}
				}
				return m, nil
			} else if m.state == stateIDESelection {
				// Select IDE and proceed to generating state
				ideInfo := getIDEInfo()
				if m.selectedIDEIndex >= 0 && m.selectedIDEIndex < len(ideInfo) {
					m.selectedIDE = ideInfo[m.selectedIDEIndex].Name
					// Transition to generating state
					if m.transitionToState(stateGenerating) {
						// Initialize steps
						m.steps = getGenerationSteps()
						// Start with first step in progress
						if len(m.steps) > 0 {
							m.steps[0].Status = stepInProgress
						}
						// Start generation using a command that returns immediately
						// The command will start the goroutine and return right away
						return m, tea.Batch(
							tickSpinner,
							startGenerationAsync(m),
						)
					}
				}
				return m, nil
			} else if m.state == stateSuccess {
				// Enter on success screen exits
				return m, tea.Quit
			}
			return m, nil

		case "up", "k":
			if m.state == stateIDESelection {
				if m.selectedIDEIndex > 0 {
					m.selectedIDEIndex--
				} else {
					// Wrap to bottom
					m.selectedIDEIndex = len(getIDEInfo()) - 1
				}
				return m, nil
			}
			return m, nil

		case "down", "j":
			if m.state == stateIDESelection {
				ideInfo := getIDEInfo()
				if m.selectedIDEIndex < len(ideInfo)-1 {
					m.selectedIDEIndex++
				} else {
					// Wrap to top
					m.selectedIDEIndex = 0
				}
				return m, nil
			}
			return m, nil

		default:
			// Handle text input in project name state
			if m.state == stateProjectName {
				m.textInput, cmd = m.textInput.Update(msg)
				// Real-time validation
				m.validateProjectName(m.textInput.Value())
				return m, cmd
			}
			return m, nil
		}

	case time.Time:
		// Handle spinner tick
		if m.state == stateGenerating {
			m.spinnerFrame++

			// Check for generation completion (non-blocking)
			if m.generationChan != nil {
				select {
				case msg := <-m.generationChan:
					// Generation completed
					if genMsg, ok := msg.(generationCompleteMsg); ok {
						m.generationDone = true
						m.generationErr = genMsg.err
						m.generationChan = nil // Close channel

						if genMsg.err != nil {
							// Transition to error state
							m.state = stateError
							m.errorMessage = genMsg.err.Error()
							m.errorSuggestion = "Please check the error message above and try again."
							return m, nil
						} else {
							// Mark all steps as completed and transition to success
							for i := range m.steps {
								m.steps[i].Status = stepCompleted
							}
							m.state = stateSuccess
							return m, nil
						}
					}
				default:
					// No message yet, continue spinner
				}
			}

			// Continue spinner animation while generating
			return m, tickSpinner
		}
		return m, nil

	case generationCompleteMsg:
		// Generation completed
		m.generationDone = true
		m.generationErr = msg.err
		if msg.err != nil {
			// Transition to error state
			m.state = stateError
			m.errorMessage = msg.err.Error()
			m.errorSuggestion = "Please check the error message above and try again."
			return m, nil
		} else {
			// Mark all steps as completed
			for i := range m.steps {
				m.steps[i].Status = stepCompleted
			}
			// Transition to success state
			m.state = stateSuccess
			return m, nil
		}

	case stepProgressMsg:
		// Update step progress
		if msg.stepIndex < len(m.steps) {
			// Mark current step as completed
			if msg.stepIndex > 0 {
				m.steps[msg.stepIndex-1].Status = stepCompleted
			}
			// Mark next step as in progress
			if msg.stepIndex < len(m.steps) {
				m.steps[msg.stepIndex].Status = stepInProgress
			}
		}
		return m, nil

	case setupGenerationChanMsg:
		// Store the channel for checking results in tick handler
		m.generationChan = msg.ch
		// Return immediately - tick handler will check the channel
		return m, nil

	default:
		// Handle text input messages
		if m.state == stateProjectName {
			m.textInput, cmd = m.textInput.Update(msg)
			// Real-time validation
			m.validateProjectName(m.textInput.Value())
			return m, cmd
		}
		return m, nil
	}
}

// View renders the UI
func (m Model) View() string {
	switch m.state {
	case stateWelcome:
		return m.renderWelcome()
	case stateProjectName:
		return m.renderProjectName()
	case stateIDESelection:
		return m.renderIDESelection()
	case stateGenerating:
		return m.renderGenerating()
	case stateSuccess:
		return m.renderSuccess()
	case stateError:
		return m.renderError()
	default:
		return ""
	}
}

// renderProgressIndicator renders a progress indicator (Step X of Y)
func (m Model) renderProgressIndicator() string {
	currentStep := m.getCurrentStepNumber()
	totalSteps := getTotalSteps()

	// Only show progress indicator for non-terminal states
	if currentStep == 0 || m.state == stateSuccess || m.state == stateError {
		return ""
	}

	progressStyle := lipgloss.NewStyle().
		Foreground(gray600).
		Align(lipgloss.Center).
		MarginBottom(1).
		Italic(true)

	return progressStyle.Render(fmt.Sprintf("Step %d of %d", currentStep, totalSteps))
}

// renderWelcome renders the welcome screen
func (m Model) renderWelcome() string {
	// Define styles
	titleStyle := lipgloss.NewStyle().
		Foreground(purple).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(pink).
		Align(lipgloss.Center).
		MarginBottom(2)

	instructionStyle := lipgloss.NewStyle().
		Foreground(white).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	helpStyle := lipgloss.NewStyle().
		Foreground(gray800).
		Align(lipgloss.Center).
		MarginTop(1)

	// Border style
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		Padding(2, 4).
		Width(m.width - 4).
		Align(lipgloss.Center)

	// ASCII art header with emoji
	header := fmt.Sprintf("%s\n%s\n%s",
		"🚀",
		titleStyle.Render("DoPlan CLI"),
		subtitleStyle.Render("Create professional projects\nin seconds"),
	)

	// Instructions
	instructions := instructionStyle.Render(
		"Press Enter to continue\n" +
			"Press 'q' to quit",
	)

	// Help text
	help := helpStyle.Render("")

	// Combine all elements
	content := fmt.Sprintf("%s\n\n%s\n%s",
		header,
		instructions,
		help,
	)

	// Apply border
	return lipgloss.Place(
		m.width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		borderStyle.Render(content),
	)
}

// validateProjectName validates the project name and updates validation error
// Returns true if valid, false otherwise
func (m *Model) validateProjectName(name string) bool {
	if name == "" {
		m.validationErr = ""
		return false
	}

	// Use the validation function from models package
	if !models.IsValidProjectName(name) {
		m.validationErr = "Project name must contain only alphanumeric characters, hyphens, and underscores"
		return false
	}

	m.validationErr = ""
	return true
}

// renderProjectName renders the project name input screen
func (m Model) renderProjectName() string {
	// Define styles
	titleStyle := lipgloss.NewStyle().
		Foreground(purple).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(pink).
		Align(lipgloss.Center).
		MarginBottom(2)

	labelStyle := lipgloss.NewStyle().
		Foreground(white).
		Bold(true).
		MarginBottom(1)

	errorStyle := lipgloss.NewStyle().
		Foreground(red).
		MarginTop(1)

	helpStyle := lipgloss.NewStyle().
		Foreground(gray600).
		MarginTop(1).
		Italic(true)

	// Determine border color based on validation
	borderColor := purple
	if m.textInput.Value() != "" {
		if m.validationErr == "" {
			borderColor = green // Valid input
		} else {
			borderColor = red // Invalid input
		}
	}

	// Border style (changes color based on validation)
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(2, 4).
		Width(m.width - 4).
		Align(lipgloss.Center)

	// Input field style
	inputStyle := lipgloss.NewStyle().
		Width(40).
		Align(lipgloss.Center)

	// Header
	header := fmt.Sprintf("%s\n%s\n%s",
		"📝",
		titleStyle.Render("Project Name"),
		subtitleStyle.Render("Enter your project name"),
	)

	// Label
	label := labelStyle.Render("Project Name:")

	// Input field
	inputField := inputStyle.Render(m.textInput.View())

	// Character count
	charCount := fmt.Sprintf("%d/50", len(m.textInput.Value()))
	charCountStyle := lipgloss.NewStyle().
		Foreground(gray600).
		MarginTop(1)
	charCountDisplay := charCountStyle.Render(charCount)

	// Error message
	errorMsg := ""
	if m.validationErr != "" {
		errorMsg = errorStyle.Render("❌ " + m.validationErr)
	}

	// Progress indicator
	progressIndicator := m.renderProgressIndicator()

	// Help text
	help := helpStyle.Render("Press Enter to continue • Press Esc to go back • Press 'q' to quit")

	// Combine all elements
	content := fmt.Sprintf("%s\n\n%s\n%s\n%s\n%s\n%s\n%s",
		header,
		progressIndicator,
		label,
		inputField,
		charCountDisplay,
		errorMsg,
		help,
	)

	// Apply border
	return lipgloss.Place(
		m.width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		borderStyle.Render(content),
	)
}

// renderIDESelection renders the IDE selection menu screen
func (m Model) renderIDESelection() string {
	ideInfo := getIDEInfo()

	// Define styles
	titleStyle := lipgloss.NewStyle().
		Foreground(purple).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(pink).
		Align(lipgloss.Center).
		MarginBottom(2)

	helpStyle := lipgloss.NewStyle().
		Foreground(gray600).
		MarginTop(2).
		Italic(true)

	// Border style
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		Padding(2, 4).
		Width(m.width - 4).
		Align(lipgloss.Center)

	// Header
	header := fmt.Sprintf("%s\n%s\n%s",
		"💻",
		titleStyle.Render("Select Your IDE"),
		subtitleStyle.Render("Choose your preferred development environment"),
	)

	// Build menu items
	var menuItems []string
	for i, ide := range ideInfo {
		// Radio button (○ for unselected, ● for selected)
		radioButton := "○"
		if i == m.selectedIDEIndex {
			radioButton = "●"
		}

		// Recommended star
		star := ""
		if ide.Recommended {
			star = " ⭐"
		}

		// Item text
		itemText := fmt.Sprintf("%s %s%s", radioButton, ide.Name, star)

		// Style based on selection
		if i == m.selectedIDEIndex {
			// Selected item: purple background, white text
			itemStyle := lipgloss.NewStyle().
				Foreground(white).
				Background(purple).
				Bold(true).
				Padding(0, 1).
				Margin(0, 0, 0, 2)
			itemText = itemStyle.Render(itemText)
		} else {
			// Unselected item: gray text
			itemStyle := lipgloss.NewStyle().
				Foreground(gray600).
				Margin(0, 0, 0, 2)
			itemText = itemStyle.Render(itemText)
		}

		// Description
		descStyle := lipgloss.NewStyle().
			Foreground(gray600).
			Italic(true).
			Margin(0, 0, 0, 6)
		description := descStyle.Render(ide.Description)

		menuItems = append(menuItems, fmt.Sprintf("%s\n%s", itemText, description))
	}

	// Combine menu items
	menu := ""
	for i, item := range menuItems {
		menu += item
		if i < len(menuItems)-1 {
			menu += "\n\n"
		}
	}

	// Progress indicator
	progressIndicator := m.renderProgressIndicator()

	// Help text
	help := helpStyle.Render("Use ↑/↓ to navigate • Press Enter to select • Press Esc to go back • Press 'q' to quit")

	// Combine all elements
	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s",
		header,
		progressIndicator,
		menu,
		help,
	)

	// Apply border
	return lipgloss.Place(
		m.width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		borderStyle.Render(content),
	)
}

// getSpinnerChar returns the spinner character for the current frame
func (m Model) getSpinnerChar() string {
	spinnerChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	return spinnerChars[m.spinnerFrame%len(spinnerChars)]
}

// getStepIcon returns the icon for a step based on its status
func (m Model) getStepIcon(step GenerationStep) string {
	switch step.Status {
	case stepCompleted:
		return "✅"
	case stepInProgress:
		return "⏳ " + m.getSpinnerChar()
	case stepFailed:
		return "❌"
	default:
		return "⏸"
	}
}

// getStepColor returns the color for a step based on its status
func (m Model) getStepColor(step GenerationStep) lipgloss.Color {
	switch step.Status {
	case stepCompleted:
		return green
	case stepInProgress:
		return yellow
	case stepFailed:
		return red
	default:
		return gray600
	}
}

// calculateProgress calculates the progress percentage
func (m Model) calculateProgress() int {
	if len(m.steps) == 0 {
		return 0
	}
	completed := 0
	for _, step := range m.steps {
		if step.Status == stepCompleted {
			completed++
		}
	}
	return (completed * 100) / len(m.steps)
}

// renderProgressBar renders a progress bar
func (m Model) renderProgressBar(percentage int, width int) string {
	filled := (percentage * width) / 100
	unfilled := width - filled

	filledBar := lipgloss.NewStyle().
		Background(purple).
		Foreground(white).
		Render(strings.Repeat("█", filled))

	unfilledBar := lipgloss.NewStyle().
		Background(gray800).
		Foreground(gray600).
		Render(strings.Repeat("░", unfilled))

	return fmt.Sprintf("[%s%s] %d%%", filledBar, unfilledBar, percentage)
}

// renderGenerating renders the progress/generating screen
func (m Model) renderGenerating() string {
	// Define styles
	titleStyle := lipgloss.NewStyle().
		Foreground(purple).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(pink).
		Align(lipgloss.Center).
		MarginBottom(2)

	stepStyle := lipgloss.NewStyle().
		Margin(0, 0, 0, 2)

	progressBarStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		MarginTop(2)

	// Border style
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		Padding(2, 4).
		Width(m.width - 4).
		Align(lipgloss.Center)

	// Header
	header := fmt.Sprintf("%s\n%s\n%s",
		"⚙️",
		titleStyle.Render("Generating your project..."),
		subtitleStyle.Render("This will only take a moment"),
	)

	// Build step list
	var stepLines []string
	for _, step := range m.steps {
		icon := m.getStepIcon(step)
		stepColor := m.getStepColor(step)
		stepTextStyle := stepStyle.Copy().
			Foreground(stepColor)
		stepLine := stepTextStyle.Render(fmt.Sprintf("%s %s", icon, step.Name))
		stepLines = append(stepLines, stepLine)
	}

	stepsText := ""
	for i, line := range stepLines {
		stepsText += line
		if i < len(stepLines)-1 {
			stepsText += "\n"
		}
	}

	// Progress bar
	progress := m.calculateProgress()
	progressBarWidth := 30
	progressBar := progressBarStyle.Render(m.renderProgressBar(progress, progressBarWidth))

	// Combine all elements
	content := fmt.Sprintf("%s\n\n%s\n\n%s",
		header,
		stepsText,
		progressBar,
	)

	// Apply border
	return lipgloss.Place(
		m.width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		borderStyle.Render(content),
	)
}

// getProjectStructureTree returns a tree representation of the generated project structure
func (m Model) getProjectStructureTree() string {
	if m.projectName == "" {
		m.projectName = "my-project"
	}

	tree := fmt.Sprintf(`✅ %s/
  ├── .cursor/
  │   ├── agents/           # 18 full persona prompts
  │   ├── commands/         # Command definitions
  │   └── rules/            # Rules library (15 categories, 1000+ rules)
  │       └── library/       # Embedded rules for all tech stacks
  ├── .plan/
  │   ├── 00_System/       # IDEA.md, PRD.md, ARCHITECTURE.md, DESIGN_SYSTEM.md
  │   ├── TASKS.md         # Implementation tasks
  │   └── active_state.json # Current project state
  ├── .github/
  │   └── workflows/       # CI/CD, release, changelog workflows
  ├── src/                 # Your source code
  ├── STANDUP.md           # Daily standup notes
  └── README.md            # Project documentation`, m.projectName)

	return tree
}

// renderSuccess renders the success screen
func (m Model) renderSuccess() string {
	// Define styles
	titleStyle := lipgloss.NewStyle().
		Foreground(green).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(pink).
		Align(lipgloss.Center).
		MarginBottom(2)

	treeStyle := lipgloss.NewStyle().
		Foreground(white).
		Margin(1, 0, 1, 2)

	nextStepsTitleStyle := lipgloss.NewStyle().
		Foreground(white).
		Bold(true).
		MarginTop(2).
		MarginBottom(1)

	nextStepStyle := lipgloss.NewStyle().
		Foreground(green).
		Margin(0, 0, 0, 4)

	helpStyle := lipgloss.NewStyle().
		Foreground(gray600).
		MarginTop(2).
		Italic(true)

	// Border style
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(green).
		Padding(2, 4).
		Width(m.width - 4).
		Align(lipgloss.Center)

	// Header
	header := fmt.Sprintf("%s\n%s\n%s",
		"✨",
		titleStyle.Render("Project created successfully!"),
		subtitleStyle.Render("Your project is ready to go"),
	)

	// Project structure tree
	projectTree := treeStyle.Render(m.getProjectStructureTree())

	// Next steps
	nextStepsTitle := nextStepsTitleStyle.Render("Next steps:")

	// Determine IDE command based on selected IDE
	ideCommand := "code"
	if m.selectedIDE != "" {
		switch m.selectedIDE {
		case "Cursor":
			ideCommand = "cursor"
		case "Claude Code":
			ideCommand = "claude"
		case "Antigravity":
			ideCommand = "antigravity"
		case "Windsurf":
			ideCommand = "windsurf"
		case "Cline":
			ideCommand = "cline"
		case "OpenCode":
			ideCommand = "opencode"
		default:
			ideCommand = "code"
		}
	}

	nextStep1 := nextStepStyle.Render(fmt.Sprintf("1. Open with: %s ./%s", ideCommand, m.projectName))
	nextStep2 := nextStepStyle.Render("2. Then type /tell to begin")

	nextSteps := fmt.Sprintf("%s\n%s\n%s",
		nextStepsTitle,
		nextStep1,
		nextStep2,
	)

	// Help text
	help := helpStyle.Render("Press Enter to exit • Press 'q' to quit")

	// Combine all elements
	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s",
		header,
		projectTree,
		nextSteps,
		help,
	)

	// Apply border
	return lipgloss.Place(
		m.width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		borderStyle.Render(content),
	)
}

// setError sets an error message and transitions to error state
func (m *Model) setError(message, suggestion string) {
	m.errorMessage = message
	m.errorSuggestion = suggestion
	m.previousState = m.state
	m.transitionToState(stateError)
}

// renderError renders the error screen
func (m Model) renderError() string {
	// Define styles
	titleStyle := lipgloss.NewStyle().
		Foreground(red).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(pink).
		Align(lipgloss.Center).
		MarginBottom(2)

	errorMessageStyle := lipgloss.NewStyle().
		Foreground(white).
		Margin(1, 0, 1, 2).
		Width(m.width - 12).
		Align(lipgloss.Left)

	suggestionStyle := lipgloss.NewStyle().
		Foreground(yellow).
		Margin(1, 0, 1, 2).
		Width(m.width - 12).
		Align(lipgloss.Left).
		Italic(true)

	recoveryTitleStyle := lipgloss.NewStyle().
		Foreground(white).
		Bold(true).
		MarginTop(2).
		MarginBottom(1)

	recoveryOptionStyle := lipgloss.NewStyle().
		Foreground(blue).
		Margin(0, 0, 0, 4)

	helpStyle := lipgloss.NewStyle().
		Foreground(gray600).
		MarginTop(2).
		Italic(true)

	// Border style
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(red).
		Padding(2, 4).
		Width(m.width - 4).
		Align(lipgloss.Center)

	// Header
	header := fmt.Sprintf("%s\n%s\n%s",
		"❌",
		titleStyle.Render("Error"),
		subtitleStyle.Render("Something went wrong"),
	)

	// Error message
	errorMsg := errorMessageStyle.Render(m.errorMessage)

	// Suggestion
	suggestion := ""
	if m.errorSuggestion != "" {
		suggestion = suggestionStyle.Render("💡 " + m.errorSuggestion)
	}

	// Recovery options
	recoveryTitle := recoveryTitleStyle.Render("What would you like to do?")

	recoveryOptions := fmt.Sprintf("%s\n%s\n%s",
		recoveryOptionStyle.Render("1. Press 'r' to retry"),
		recoveryOptionStyle.Render("2. Press 'b' to go back"),
		recoveryOptionStyle.Render("3. Press 'q' to quit"),
	)

	recovery := fmt.Sprintf("%s\n%s",
		recoveryTitle,
		recoveryOptions,
	)

	// Help text
	help := helpStyle.Render("Use keyboard shortcuts to recover from this error")

	// Combine all elements
	content := fmt.Sprintf("%s\n\n%s\n%s\n\n%s\n\n%s",
		header,
		errorMsg,
		suggestion,
		recovery,
		help,
	)

	// Apply border
	return lipgloss.Place(
		m.width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		borderStyle.Render(content),
	)
}

// toProjectRequest converts the model to a ProjectRequest
func (m Model) toProjectRequest() *models.ProjectRequest {
	return &models.ProjectRequest{
		ProjectName: m.projectName,
		IDE:         m.selectedIDE,
		ProjectType: "Fullstack", // Default project type
	}
}

// Run starts the wizard TUI and returns the ProjectRequest if successful
func Run() (*models.ProjectRequest, error) {
	initialModel := InitialModel()
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running wizard: %v\n", err)
		return nil, err
	}

	// Cast to Model
	model, ok := finalModel.(Model)
	if !ok {
		return nil, fmt.Errorf("unexpected model type: %T", finalModel)
	}

	// Check if wizard completed successfully
	if model.state == stateSuccess {
		// Extract and validate ProjectRequest
		request := model.toProjectRequest()
		if err := request.Validate(); err != nil {
			return nil, fmt.Errorf("invalid project request: %w", err)
		}
		return request, nil
	}

	// User quit or error occurred
	return nil, nil
}

// startGenerationAsync starts generation in a goroutine and returns immediately
// This is a proper Bubble Tea command that doesn't block
func startGenerationAsync(m Model) tea.Cmd {
	return func() tea.Msg {
		// Create project request
		request := m.toProjectRequest()
		if err := request.Validate(); err != nil {
			return generationCompleteMsg{err: fmt.Errorf("invalid project request: %w", err)}
		}

		// Start generation in goroutine - this MUST return immediately
		// We'll use a channel to get the result, but we check it in tick handler
		resultChan := make(chan tea.Msg, 1)
		
		go func() {
			// Run generation in background
			if err := generator.Orchestrate(request); err != nil {
				resultChan <- generationCompleteMsg{err: fmt.Errorf("generation failed: %w", err)}
			} else {
				resultChan <- generationCompleteMsg{err: nil}
			}
		}()

		// Return immediately with channel setup - this is non-blocking
		return setupGenerationChanMsg{ch: resultChan}
	}
}

// setupGenerationChanMsg sets up the channel for checking generation results
type setupGenerationChanMsg struct {
	ch chan tea.Msg
}
