package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunDevServer auto-detects and runs the development server
func RunDevServer() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Detect project type and package manager
	projectType, packageManager, err := detectProjectType(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to detect project type: %w", err)
	}

	// Determine dev command
	devCommand := getDevCommand(projectType, packageManager)
	if devCommand == nil {
		return fmt.Errorf("could not determine dev command for project type: %s", projectType)
	}

	fmt.Printf("🚀 Starting dev server for %s project...\n", projectType)
	fmt.Printf("📦 Using %s\n", packageManager)
	fmt.Printf("▶️  Running: %s\n\n", strings.Join(devCommand, " "))

	// Run the command
	cmd := exec.Command(devCommand[0], devCommand[1:]...)
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// detectProjectType detects the project type and package manager
func detectProjectType(projectRoot string) (string, string, error) {
	// Check for package.json (Node.js)
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		// Check for yarn.lock or pnpm-lock.yaml
		if _, err := os.Stat(filepath.Join(projectRoot, "yarn.lock")); err == nil {
			return "node", "yarn", nil
		}
		if _, err := os.Stat(filepath.Join(projectRoot, "pnpm-lock.yaml")); err == nil {
			return "node", "pnpm", nil
		}
		return "node", "npm", nil
	}

	// Check for go.mod (Go)
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
		return "go", "go", nil
	}

	// Check for requirements.txt or pyproject.toml (Python)
	if _, err := os.Stat(filepath.Join(projectRoot, "requirements.txt")); err == nil {
		return "python", "pip", nil
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "pyproject.toml")); err == nil {
		return "python", "pip", nil
	}

	// Check for Cargo.toml (Rust)
	if _, err := os.Stat(filepath.Join(projectRoot, "Cargo.toml")); err == nil {
		return "rust", "cargo", nil
	}

	return "unknown", "unknown", fmt.Errorf("could not detect project type")
}

// getDevCommand returns the dev command for the project type
func getDevCommand(projectType, packageManager string) []string {
	switch projectType {
	case "node":
		switch packageManager {
		case "yarn":
			return []string{"yarn", "dev"}
		case "pnpm":
			return []string{"pnpm", "dev"}
		default:
			return []string{"npm", "run", "dev"}
		}
	case "go":
		return []string{"go", "run", "."}
	case "python":
		// Try to detect if it's FastAPI, Flask, etc.
		if _, err := os.Stat("main.py"); err == nil {
			// Check for FastAPI (uvicorn) or Flask
			return []string{"python", "-m", "uvicorn", "main:app", "--reload"}
		}
		return []string{"python", "-m", "flask", "run", "--reload"}
	case "rust":
		return []string{"cargo", "run"}
	default:
		return nil
	}
}
