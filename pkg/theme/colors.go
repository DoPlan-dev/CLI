package theme

import "github.com/charmbracelet/lipgloss"

// Color constants for the DoPlan theme
const (
	ColorPrimary   = "#667eea"
	ColorSecondary = "#764ba2"
	ColorSuccess   = "#10b981"
	ColorWarning   = "#f59e0b"
	ColorError     = "#ef4444"
	ColorText      = "#ffffff"
	ColorTextDim   = "#999999"
	ColorBorder    = "#333333"
)

// Primary returns the primary color
func Primary() lipgloss.Color {
	return lipgloss.Color(ColorPrimary)
}

// Secondary returns the secondary color
func Secondary() lipgloss.Color {
	return lipgloss.Color(ColorSecondary)
}

// Success returns the success color
func Success() lipgloss.Color {
	return lipgloss.Color(ColorSuccess)
}

// Warning returns the warning color
func Warning() lipgloss.Color {
	return lipgloss.Color(ColorWarning)
}

// Error returns the error color
func Error() lipgloss.Color {
	return lipgloss.Color(ColorError)
}

// Text returns the text color
func Text() lipgloss.Color {
	return lipgloss.Color(ColorText)
}

// TextDim returns the dimmed text color
func TextDim() lipgloss.Color {
	return lipgloss.Color(ColorTextDim)
}

// Border returns the border color
func Border() lipgloss.Color {
	return lipgloss.Color(ColorBorder)
}

