package cli

import (
	"fmt"
	"io"
	"strings"
)

// RenderProgressBar renders a simple text-based progress bar for CLI output
// percentage: 0-100
// width: number of characters for the bar (default 30 if <= 0)
func RenderProgressBar(percentage int, width int) string {
	// Clamp percentage
	if percentage < 0 {
		percentage = 0
	} else if percentage > 100 {
		percentage = 100
	}

	// Default width
	if width <= 0 {
		width = 30
	}

	// Calculate filled and unfilled
	filled := (percentage * width) / 100
	unfilled := width - filled

	// Build bar
	filledBar := strings.Repeat("█", filled)
	unfilledBar := strings.Repeat("░", unfilled)

	return fmt.Sprintf("[%s%s] %d%%", filledBar, unfilledBar, percentage)
}

// DisplayPhaseProgress displays progress for a phase with a progress bar
func DisplayPhaseProgress(out io.Writer, phaseName string, currentStep, totalSteps int) {
	percentage := 0
	if totalSteps > 0 {
		percentage = (currentStep * 100) / totalSteps
	}

	progressBar := RenderProgressBar(percentage, 30)
	fmt.Fprintf(out, "\n📊 Phase: %s\n", phaseName)
	fmt.Fprintf(out, "   %s\n", progressBar)
	fmt.Fprintf(out, "   Step %d of %d\n\n", currentStep, totalSteps)
}

// DisplayFeatureProgress displays progress for a feature with a progress bar
func DisplayFeatureProgress(out io.Writer, featureName string, completedTasks, totalTasks int) {
	percentage := 0
	if totalTasks > 0 {
		percentage = (completedTasks * 100) / totalTasks
	}

	progressBar := RenderProgressBar(percentage, 30)
	fmt.Fprintf(out, "\n🚀 Feature: %s\n", featureName)
	fmt.Fprintf(out, "   %s\n", progressBar)
	fmt.Fprintf(out, "   Tasks: %d/%d completed\n\n", completedTasks, totalTasks)
}
