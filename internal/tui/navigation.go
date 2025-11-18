package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/manifoldco/promptui"
)

var (
	headerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{
			Dark:  "#667eea",
			Light: "#764ba2",
		}).
		Bold(true).
		Align(lipgloss.Center).
		Padding(1, 2)
)

// ShowHeader displays the DoPlan header
func ShowHeader() {
	logo := `
  ██████╗░░█████╗░██████╗░██╗░░░░░░█████╗░███╗░░██╗
  ██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔══██╗████╗░██║
  ██║░░██║██║░░██║██████╔╝██║░░░░░███████║██╔██╗██║
  ██║░░██║██║░░██║██╔═══╝░██║░░░░░██╔══██║██║╚████║
  ██████╔╝╚█████╔╝██║░░░░░███████╗██║░░██║██║░╚███║
  ╚═════╝░░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚══╝
`
	fmt.Print(headerStyle.Render(logo))
	fmt.Println()
}

// ShowInstallMenu displays the installation menu
func ShowInstallMenu() (string, error) {
	ShowHeader()

	prompt := promptui.Select{
		Label: "What are you using to develop your application?",
		Items: []string{
			"Cursor",
			"Gemini CLI",
			"Claude CLI",
			"Codex",
			"OpenCode",
			"Qwen Code",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// Convert to lowercase with dashes
	ideMap := map[string]string{
		"Cursor":     "cursor",
		"Gemini CLI": "gemini",
		"Claude CLI": "claude",
		"Codex":      "codex",
		"OpenCode":   "opencode",
		"Qwen Code":  "qwen",
		"Back":       "back",
	}

	return ideMap[result], nil
}

// ConfirmReinstall asks user to confirm reinstallation
func ConfirmReinstall() (bool, error) {
	prompt := promptui.Prompt{
		Label:     "DoPlan is already installed. Reinstall? This will overwrite existing configuration",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return result == "y" || result == "Y", nil
}
