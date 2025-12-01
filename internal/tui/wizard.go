package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/generator"
	"github.com/DoPlan-dev/CLI/internal/version"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Color definitions - monochrome design
var (
	// Single color scheme (white/gray for monochrome)
	primary   = lipgloss.Color("#FFFFFF") // White for primary text
	secondary = lipgloss.Color("#CCCCCC") // Light gray for secondary text
	tertiary  = lipgloss.Color("#888888") // Medium gray for tertiary text
	dim       = lipgloss.Color("#666666") // Dim gray for less important text
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
	selectedIDEs     map[string]bool
	ideSelectionErr  string
	steps            []GenerationStep
	spinnerFrame     int
	errorMessage     string
	errorSuggestion  string
	previousState    wizardState   // Store previous state for recovery
	stateHistory     []wizardState // History for back navigation
	generationDone   bool          // Track if generation has completed
	generationErr    error         // Store generation error if any
}

// generationCompleteMsg is sent when generation completes
type generationCompleteMsg struct {
	err error
}

// stepProgressMsg is sent when a step completes
type stepProgressMsg struct {
	stepIndex int
}

// programRef is a global reference to the program for sending messages from goroutines
// This is safe because Run() is only called once and clears it after
var programRef *tea.Program

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

	ideSelections := make(map[string]bool)
	if ideOptions := getIDEInfo(); len(ideOptions) > 0 {
		ideSelections[ideOptions[0].Name] = true
	}

	return Model{
		state:            stateWelcome,
		width:            80, // Default width
		textInput:        ti,
		projectName:      "",
		selectedIDEIndex: 0, // Default to first IDE (Cursor)
		selectedIDE:      "",
		selectedIDEs:     ideSelections,
		ideSelectionErr:  "",
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

		case "ctrl+r":
			// Retry: go back to previous state
			if m.state == stateError {
				m.state = m.previousState
				m.errorMessage = ""
				m.errorSuggestion = ""
				return m, nil
			}
			return m, nil

		case "ctrl+b":
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

		case "esc":
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
				// Require at least one IDE selection
				selectedIDEs := m.getSelectedIDEs()
				if len(selectedIDEs) == 0 {
					m.ideSelectionErr = "Select at least one IDE (press Space to toggle)"
					return m, nil
				}
				m.selectedIDE = selectedIDEs[0]
				m.ideSelectionErr = ""
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
				return m, nil
			} else if m.state == stateSuccess {
				// Enter on success screen exits
				return m, tea.Quit
			}
			return m, nil

		case "up", "ctrl+k":
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

		case "down", "ctrl+j":
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

		case " ":
			if m.state == stateIDESelection {
				m.toggleIDESelection(m.selectedIDEIndex)
				m.ideSelectionErr = ""
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
			// Continue spinner animation while generating
			// The generation goroutine will send a message directly when done
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
		Foreground(tertiary).
		Align(lipgloss.Center).
		MarginBottom(1).
		Italic(true)

	return progressStyle.Render(fmt.Sprintf("Step %d of %d", currentStep, totalSteps))
}

// extractCleanVersion extracts just the version number from version strings like "v1.3.0-11-g7e8015b-dirty"
func extractCleanVersion(versionStr string) string {
	if versionStr == "dev" {
		return "dev"
	}
	
	// Remove "v" prefix if present
	if strings.HasPrefix(versionStr, "v") {
		versionStr = versionStr[1:]
	}
	
	// Extract just the version number (stop at first non-version character like "-")
	parts := strings.Split(versionStr, "-")
	versionNum := parts[0]
	
	// Format as "v 1.3.0" (with space after v)
	return fmt.Sprintf("v %s", versionNum)
}

// getVersionNumber extracts just the version number without "v " prefix
func getVersionNumber(versionStr string) string {
	if versionStr == "dev" {
		return "dev"
	}
	
	// Remove "v" prefix if present
	if strings.HasPrefix(versionStr, "v") {
		versionStr = versionStr[1:]
	}
	
	// Extract just the version number (stop at first non-version character like "-")
	parts := strings.Split(versionStr, "-")
	return parts[0]
}

// renderTopLine renders the top line with application name and version
func renderTopLine() string {
	versionStr := version.GetVersion()
	versionNum := getVersionNumber(versionStr)
	appNameWithVersion := fmt.Sprintf("DoPlan CLI %s", versionNum)
	
	topStyle := lipgloss.NewStyle().
		Foreground(tertiary).
		Width(80)

	// Format: "DoPlan CLI 1.3.0"
	line := appNameWithVersion
	return topStyle.Render(line)
}

// renderASCIIHeader renders the DoPlan ASCII art header
func renderASCIIHeader() string {
	asciiArt := `██████╗░░█████╗░██████╗░██╗░░░░░░█████╗░███╗░░██╗
██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔══██╗████╗░██║
██║░░██║██║░░██║██████╔╝██║░░░░░███████║██╔██╗██║
██║░░██║██║░░██║██╔═══╝░██║░░░░░██╔══██║██║╚████║
██████╔╝╚█████╔╝██║░░░░░███████╗██║░░██║██║░╚███║
╚═════╝░░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚══╝`

	// Center the ASCII art within 80-character width
	// Each line is 49 characters, so we need 15 spaces on each side (80-49)/2 = 15.5 -> 15
	lines := strings.Split(asciiArt, "\n")
	var centeredLines []string
	for _, line := range lines {
		if len(line) > 0 {
			// Center each line by padding with spaces
			lineLen := len(line)
			if lineLen < 80 {
				paddingLen := (80 - lineLen) / 2
				if paddingLen > 0 {
					padding := strings.Repeat(" ", paddingLen)
					centeredLines = append(centeredLines, padding+line)
				} else {
					centeredLines = append(centeredLines, line)
				}
			} else {
				// Line is too long, just use it as-is
				centeredLines = append(centeredLines, line)
			}
		}
	}
	centeredArt := strings.Join(centeredLines, "\n")

	headerStyle := lipgloss.NewStyle().
		Foreground(primary).
		Width(80).
		Align(lipgloss.Center).
		MarginTop(2). // ~30px spacing above ASCII art header
		MarginBottom(0)

	return headerStyle.Render(centeredArt)
}

// renderBody renders the body section with content
func renderBody(content string) string {
	bodyStyle := lipgloss.NewStyle().
		Width(80).
		Align(lipgloss.Center).
		Padding(1, 0)

	return bodyStyle.Render(content)
}

// renderFooter renders the footer with navigation instructions
func renderFooter() string {
	footerStyle := lipgloss.NewStyle().
		Foreground(dim).
		Width(80).
		Align(lipgloss.Center).
		MarginTop(1)

	// Footer format matching footer.md exactly (with 4 trailing spaces)
	footer := "|    ↑/↓  Nav     |    Space / Select    |     Enter / Apply     |    "
	return footerStyle.Render(footer)
}

// renderFullTUI renders the complete TUI layout according to documentation
func renderFullTUI(bodyContent string) string {
	topLine := renderTopLine()
	asciiHeader := renderASCIIHeader()
	body := renderBody(bodyContent)
	footer := renderFooter()

	// Top separator: 80 underscores (matching Top.md)
	topSeparator := strings.Repeat("_", 80)

	return fmt.Sprintf("%s\n%s\n%s\n%s\n\n%s\n%s",
		topLine,
		topSeparator,
		asciiHeader,
		body,
		topSeparator,
		footer,
	)
}

// renderWelcome renders the welcome screen
func (m Model) renderWelcome() string {
	// Define styles for body content
	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginBottom(3)

	instructionStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	// Body content
	bodyContent := fmt.Sprintf("%s\n%s\n%s",
		titleStyle.Render("[>] Welcome to DoPlan CLI"),
		subtitleStyle.Render("Create professional projects in seconds"),
		instructionStyle.Render("Enter to Start Installing"),
	)

	// Use the documented TUI layout
	return renderFullTUI(bodyContent)
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
	// Define styles for body content
	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginBottom(3)

	errorStyle := lipgloss.NewStyle().
		Foreground(primary).
		MarginTop(1)

	// Input field style
	inputStyle := lipgloss.NewStyle().
		Width(40).
		Align(lipgloss.Center)

	// Title
	title := titleStyle.Render("[*] Project Name")

	// Subtitle
	subtitle := subtitleStyle.Render("Enter your project name")

	// Input field
	inputField := inputStyle.Render(m.textInput.View())

	// Error message
	errorMsg := ""
	if m.validationErr != "" {
		errorMsg = errorStyle.Render("[X] " + m.validationErr)
	}

	// Body content
	bodyContent := fmt.Sprintf("%s\n%s\n%s\n%s",
		title,
		subtitle,
		inputField,
		errorMsg,
	)

	// Use the documented TUI layout
	return renderFullTUI(bodyContent)
}

// toggleIDESelection toggles the selection for the IDE at the given index.
func (m *Model) toggleIDESelection(index int) {
	ideInfo := getIDEInfo()
	if index < 0 || index >= len(ideInfo) {
		return
	}
	if m.selectedIDEs == nil {
		m.selectedIDEs = make(map[string]bool)
	}
	name := ideInfo[index].Name
	if m.selectedIDEs[name] {
		delete(m.selectedIDEs, name)
	} else {
		m.selectedIDEs[name] = true
	}
}

// getSelectedIDEs returns the current IDE selections in menu order.
func (m Model) getSelectedIDEs() []string {
	selections := make([]string, 0) // Initialize as non-nil empty slice
	ideInfo := getIDEInfo()
	for _, ide := range ideInfo {
		if m.selectedIDEs != nil && m.selectedIDEs[ide.Name] {
			selections = append(selections, ide.Name)
		}
	}

	// Fallback to legacy single-select field for test scenarios.
	if len(selections) == 0 && m.selectedIDE != "" {
		selections = append(selections, m.selectedIDE)
	}

	return selections
}

// renderIDESelection renders the IDE selection menu screen
func (m Model) renderIDESelection() string {
	ideInfo := getIDEInfo()

	// Define styles for body content
	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginBottom(3)

	// Title
	title := titleStyle.Render("[*] Select Your IDE")

	// Subtitle
	subtitle := subtitleStyle.Render("Choose your preferred development environment")

	// Build menu items
	var menuItems []string
	for i, ide := range ideInfo {
		// Checkbox (☐ for unselected, ☑ for selected)
		checkbox := "☐"
		if m.selectedIDEs != nil && m.selectedIDEs[ide.Name] {
			checkbox = "☑"
		}

		// Recommended indicator
		recommended := ""
		if ide.Recommended {
			recommended = " [*]"
		}

		// Item text
		itemText := fmt.Sprintf("%s %s%s", checkbox, ide.Name, recommended)

		// Style based on selection
		if i == m.selectedIDEIndex {
			// Selected item: inverted (white on black equivalent)
			itemStyle := lipgloss.NewStyle().
				Foreground(primary).
				Bold(true).
				Padding(0, 1).
				Margin(0, 0, 0, 0).
				Align(lipgloss.Left)
			itemText = itemStyle.Render(itemText)
		} else {
			// Unselected item: dim text
			itemStyle := lipgloss.NewStyle().
				Foreground(tertiary).
				Margin(0, 0, 0, 0).
				Align(lipgloss.Left)
			itemText = itemStyle.Render(itemText)
		}

		menuItems = append(menuItems, itemText)
	}

	// Combine menu items
	menu := ""
	for i, item := range menuItems {
		menu += item
		if i < len(menuItems)-1 {
			menu += "\n\n"
		}
	}

	// Create left-aligned container for menu with 40px (5 chars) left padding
	// Add left padding directly to the menu string
	padding := strings.Repeat(" ", 5) // 5 spaces = ~40px
	menuLines := strings.Split(menu, "\n")
	var paddedMenuLines []string
	for _, line := range menuLines {
		if line != "" {
			paddedMenuLines = append(paddedMenuLines, padding+line)
		} else {
			paddedMenuLines = append(paddedMenuLines, line)
		}
	}
	paddedMenu := strings.Join(paddedMenuLines, "\n")

	errorStyle := lipgloss.NewStyle().
		Foreground(primary).
		MarginTop(1)
	errorMsg := ""
	if m.ideSelectionErr != "" {
		errorMsg = errorStyle.Render("[X] " + m.ideSelectionErr)
	}

	// Body content - title and subtitle centered, menu left-aligned
	bodyContent := fmt.Sprintf("%s\n%s\n\n%s\n%s",
		title,
		subtitle,
		paddedMenu,
		errorMsg,
	)

	// Use the documented TUI layout
	return renderFullTUI(bodyContent)
}

// getSpinnerChar returns the spinner character for the current frame
func (m Model) getSpinnerChar() string {
	// ASCII spinner characters
	spinnerChars := []string{"|", "/", "-", "\\", "|", "/", "-", "\\"}
	return spinnerChars[m.spinnerFrame%len(spinnerChars)]
}

// getStepIcon returns the icon for a step based on its status
func (m Model) getStepIcon(step GenerationStep) string {
	switch step.Status {
	case stepCompleted:
		return "[+]"
	case stepInProgress:
		return "[...] " + m.getSpinnerChar()
	case stepFailed:
		return "[X]"
	default:
		return "[-]"
	}
}

// getStepColor returns the color for a step based on its status
func (m Model) getStepColor(step GenerationStep) lipgloss.Color {
	switch step.Status {
	case stepCompleted:
		return primary
	case stepInProgress:
		return secondary
	case stepFailed:
		return primary
	default:
		return tertiary
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
	// Clamp percentage to valid range and ensure width is positive
	if percentage < 0 {
		percentage = 0
	} else if percentage > 100 {
		percentage = 100
	}
	if width < 0 {
		width = 0
	}

	filled := (percentage * width) / 100
	unfilled := width - filled

	// Ensure filled and unfilled are non-negative
	if filled < 0 {
		filled = 0
	}
	if unfilled < 0 {
		unfilled = 0
	}

	filledBar := lipgloss.NewStyle().
		Foreground(primary).
		Render(strings.Repeat("=", filled))

	unfilledBar := lipgloss.NewStyle().
		Foreground(tertiary).
		Render(strings.Repeat("-", unfilled))

	return fmt.Sprintf("[%s%s] %d%%", filledBar, unfilledBar, percentage)
}

// renderGenerating renders the progress/generating screen
func (m Model) renderGenerating() string {
	// Define styles for body content
	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginBottom(3)

	stepStyle := lipgloss.NewStyle().
		Margin(0, 0, 0, 2)

	progressBarStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		MarginTop(2)

	// Title
	title := titleStyle.Render("[...] Generating your project...")

	// Subtitle
	subtitle := subtitleStyle.Render("This will only take a moment")

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

	// Body content
	bodyContent := fmt.Sprintf("%s\n%s\n\n%s\n\n%s",
		title,
		subtitle,
		stepsText,
		progressBar,
	)

	// Use the documented TUI layout
	return renderFullTUI(bodyContent)
}

// getProjectStructureTree returns a tree representation of the generated project structure
func (m Model) getProjectStructureTree() string {
	if m.projectName == "" {
		m.projectName = "my-project"
	}

	tree := fmt.Sprintf(`[+] %s/
  ├── .cursor/
  │   ├── agents/           # 18 full persona prompts
  │   ├── commands/         # Command definitions
  │   └── rules/            # Rules library (15 categories, 1000+ rules)
  │       └── library/       # Embedded rules for all tech stacks
  ├── .do/
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
	// Define styles for body content
	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginBottom(1)

	nextStepStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Margin(0, 0, 0, 0).
		Align(lipgloss.Center)

	// Title
	title := titleStyle.Render("Project created successfully!")

	// Subtitle
	boldSubtitleStyle := subtitleStyle.Copy().Bold(true)
	// Render parts without center alignment to avoid line breaks, then center the whole thing
	subtitleParts := []string{
		subtitleStyle.Copy().Align(lipgloss.Left).Render("Ready to meet "),
		boldSubtitleStyle.Copy().Align(lipgloss.Left).Render("/do /plan /dev"),
		subtitleStyle.Copy().Align(lipgloss.Left).Render(" workflow?"),
	}
	subtitleText := lipgloss.JoinHorizontal(lipgloss.Left, subtitleParts...)
	subtitle := subtitleStyle.Render(subtitleText)

	// Next steps
	nextStep2 := nextStepStyle.Render("in cursor type /hey to meet your instructor,\n or /do followed with your idea.")

	nextSteps := nextStep2

	// Center the next steps section
	centeredNextSteps := lipgloss.NewStyle().
		Width(80).
		Align(lipgloss.Center).
		Render(nextSteps)

	bodyContent := fmt.Sprintf("%s\n%s\n%s",
		title,
		subtitle,
		centeredNextSteps,
	)

	// Use the documented TUI layout
	return renderFullTUI(bodyContent)
}

// mapIDEToCommand maps a friendly IDE name to its CLI command.
func mapIDEToCommand(ide string) string {
	switch ide {
	case "Cursor":
		return "cursor"
	case "Claude Code":
		return "claude"
	case "Antigravity":
		return "antigravity"
	case "Windsurf":
		return "windsurf"
	case "Cline":
		return "cline"
	case "OpenCode":
		return "opencode"
	default:
		return "code"
	}
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
	// Define styles for body content
	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Align(lipgloss.Center).
		MarginBottom(3)

	errorMessageStyle := lipgloss.NewStyle().
		Foreground(primary).
		Margin(1, 0, 1, 2).
		Width(70).
		Align(lipgloss.Left)

	suggestionStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Margin(1, 0, 1, 2).
		Width(70).
		Align(lipgloss.Left).
		Italic(true)

	recoveryTitleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		MarginTop(2).
		MarginBottom(1)

	recoveryOptionStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Margin(0, 0, 0, 4)

	// Title
	title := titleStyle.Render("[X] Error")

	// Subtitle
	subtitle := subtitleStyle.Render("Something went wrong")

	// Error message
	errorMsg := errorMessageStyle.Render(m.errorMessage)

	// Suggestion
	suggestion := ""
	if m.errorSuggestion != "" {
		suggestion = suggestionStyle.Render("[!] " + m.errorSuggestion)
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

	// Body content
	bodyContent := fmt.Sprintf("%s\n%s\n\n%s\n%s\n\n%s",
		title,
		subtitle,
		errorMsg,
		suggestion,
		recovery,
	)

	// Use the documented TUI layout
	return renderFullTUI(bodyContent)
}

// toProjectRequest converts the model to a ProjectRequest
func (m Model) toProjectRequest() *models.ProjectRequest {
	selectedIDEs := m.getSelectedIDEs()
	primaryIDE := ""
	if len(selectedIDEs) > 0 {
		primaryIDE = selectedIDEs[0]
	}
	return &models.ProjectRequest{
		ProjectName: m.projectName,
		IDE:         primaryIDE,
		IDEs:        selectedIDEs,
		ProjectType: "Fullstack", // Default project type
	}
}

// programWithSend wraps tea.Program to allow sending messages from goroutines
type programWithSend struct {
	*tea.Program
}

// Run starts the wizard TUI and returns the ProjectRequest if successful
func Run() (*models.ProjectRequest, error) {
	initialModel := InitialModel()
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	// Store program reference for sending messages from goroutines
	programRef = p

	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running wizard: %v\n", err)
		return nil, err
	}

	// Clear program reference
	programRef = nil

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

		// Start generation in goroutine - send message directly to program when done
		go func() {
			// Run generation in background
			var result tea.Msg
			if err := generator.Orchestrate(request); err != nil {
				result = generationCompleteMsg{err: fmt.Errorf("generation failed: %w", err)}
			} else {
				result = generationCompleteMsg{err: nil}
			}

			// Send result directly to program (non-blocking)
			if programRef != nil {
				programRef.Send(result)
			}
		}()

		// Return nil - we don't need a setup message, the goroutine sends directly
		return nil
	}
}
