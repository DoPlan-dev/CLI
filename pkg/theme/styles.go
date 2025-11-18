package theme

import "github.com/charmbracelet/lipgloss"

var (
	// HeaderStyle is used for section headers
	HeaderStyle = lipgloss.NewStyle().
			Foreground(Primary()).
			Bold(true).
			Padding(0, 1)

	// CardStyle is used for card containers
	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Border()).
			Padding(1, 2).
			Margin(1, 0)

	// ButtonStyle is used for buttons
	ButtonStyle = lipgloss.NewStyle().
			Foreground(Text()).
			Background(Primary()).
			Padding(0, 2).
			Margin(1, 0)

	// ButtonActiveStyle is used for active/selected buttons
	ButtonActiveStyle = lipgloss.NewStyle().
				Foreground(Text()).
				Background(Secondary()).
				Padding(0, 2).
				Margin(1, 0)

	// ProgressFilled is used for filled progress bar segments
	ProgressFilled = "█"

	// ProgressEmpty is used for empty progress bar segments
	ProgressEmpty = "░"

	// SuccessStyle is used for success messages
	SuccessStyle = lipgloss.NewStyle().
			Foreground(Success()).
			Bold(true)

	// ErrorStyle is used for error messages
	ErrorStyle = lipgloss.NewStyle().
			Foreground(Error()).
			Bold(true)

	// WarningStyle is used for warning messages
	WarningStyle = lipgloss.NewStyle().
			Foreground(Warning()).
			Bold(true)

	// HelpStyle is used for help text
	HelpStyle = lipgloss.NewStyle().
			Foreground(TextDim()).
			Italic(true)

	// TextStyle is used for normal text
	TextStyle = lipgloss.NewStyle().
			Foreground(Text())

	// TextDimStyle is used for dimmed text
	TextDimStyle = lipgloss.NewStyle().
			Foreground(TextDim())
)

// RenderProgressBar renders a progress bar with the given percentage
func RenderProgressBar(percent int, width int) string {
	filled := int(float64(width) * float64(percent) / 100.0)
	empty := width - filled
	
	result := ""
	for i := 0; i < filled; i++ {
		result += ProgressFilled
	}
	for i := 0; i < empty; i++ {
		result += ProgressEmpty
	}
	
	return lipgloss.NewStyle().
		Foreground(Primary()).
		Render(result)
}

