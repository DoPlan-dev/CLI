package tools

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Tool represents an external tool
type Tool struct {
	Name        string
	CheckCmd    []string
	InstallCmd  []string
	InstallURL  string
	Description string
}

// CheckTool checks if a tool is installed
func CheckTool(tool Tool) (bool, error) {
	cmd := exec.Command(tool.CheckCmd[0], tool.CheckCmd[1:]...)
	err := cmd.Run()
	return err == nil, nil
}

// InstallTool installs a tool (shows instructions)
func InstallTool(tool Tool) error {
	if runtime.GOOS == "darwin" {
		fmt.Printf("To install %s on macOS:\n", tool.Name)
		if tool.InstallURL != "" {
			fmt.Printf("  Visit: %s\n", tool.InstallURL)
		} else if len(tool.InstallCmd) > 0 {
			fmt.Printf("  Run: %s\n", tool.InstallCmd[0])
		}
	} else if runtime.GOOS == "linux" {
		fmt.Printf("To install %s on Linux:\n", tool.Name)
		if tool.InstallURL != "" {
			fmt.Printf("  Visit: %s\n", tool.InstallURL)
		} else if len(tool.InstallCmd) > 0 {
			fmt.Printf("  Run: %s\n", tool.InstallCmd[0])
		}
	} else if runtime.GOOS == "windows" {
		fmt.Printf("To install %s on Windows:\n", tool.Name)
		if tool.InstallURL != "" {
			fmt.Printf("  Visit: %s\n", tool.InstallURL)
		} else if len(tool.InstallCmd) > 0 {
			fmt.Printf("  Run: %s\n", tool.InstallCmd[0])
		}
	}
	return nil
}

// Common tools
var (
	GitHubCLI = Tool{
		Name:        "GitHub CLI (gh)",
		CheckCmd:    []string{"gh", "--version"},
		InstallCmd:  []string{"brew", "install", "gh"},
		InstallURL:  "https://cli.github.com/",
		Description: "Required for GitHub integration",
	}

	NodeJS = Tool{
		Name:        "Node.js (npm)",
		CheckCmd:    []string{"npm", "--version"},
		InstallCmd:  []string{"brew", "install", "node"},
		InstallURL:  "https://nodejs.org/",
		Description: "Required for Node.js projects",
	}

	Docker = Tool{
		Name:        "Docker",
		CheckCmd:    []string{"docker", "--version"},
		InstallCmd:  []string{"brew", "install", "docker"},
		InstallURL:  "https://www.docker.com/",
		Description: "Required for containerized deployments",
	}
)
