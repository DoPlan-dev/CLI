package animations

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SpinnerFrames are the frames for the spinner animation
var SpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// NewSpinner creates a new spinner with DoPlan styling
func NewSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: SpinnerFrames,
		FPS:    80, // 80ms per frame
	}
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#667eea"))
	return s
}

// SpinnerModel wraps the spinner with additional functionality
type SpinnerModel struct {
	spinner spinner.Model
	message string
}

// NewSpinnerModel creates a new spinner model
func NewSpinnerModel(message string) SpinnerModel {
	return SpinnerModel{
		spinner: NewSpinner(),
		message: message,
	}
}

// Init initializes the spinner
func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update updates the spinner
func (m SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// View renders the spinner
func (m SpinnerModel) View() string {
	return m.spinner.View() + " " + m.message
}
