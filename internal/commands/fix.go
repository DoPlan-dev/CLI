package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// AutoFix runs AI-powered auto-fix for common issues
func AutoFix() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	fmt.Println("🩹 Running auto-fix...")
	fmt.Println()

	// Detect project type
	projectType, err := detectProjectTypeForFix(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to detect project type: %w", err)
	}

	var fixesApplied []string

	// Run project-specific fixes
	switch projectType {
	case "node":
		fixesApplied = append(fixesApplied, fixNpmIssues(projectRoot)...)
		fixesApplied = append(fixesApplied, fixLintingErrors(projectRoot)...)
	case "go":
		fixesApplied = append(fixesApplied, fixGoIssues(projectRoot)...)
		fixesApplied = append(fixesApplied, fixLintingErrors(projectRoot)...)
	}

	// Run universal fixes
	fixesApplied = append(fixesApplied, fixCommonIssues(projectRoot)...)

	// Display results
	displayFixResults(fixesApplied)

	return nil
}

func detectProjectTypeForFix(projectRoot string) (string, error) {
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		return "node", nil
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
		return "go", nil
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "requirements.txt")); err == nil {
		return "python", nil
	}
	return "unknown", nil
}

func fixNpmIssues(projectRoot string) []string {
	fmt.Println("📦 Fixing npm issues...")
	var fixes []string

	// Run npm audit fix
	cmd := exec.Command("npm", "audit", "fix")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Println("  ✅ npm audit fix completed")
		fixes = append(fixes, "npm audit fix")
	} else {
		// Check if there were vulnerabilities that couldn't be auto-fixed
		if strings.Contains(string(output), "vulnerabilities") {
			fmt.Println("  ⚠️  Some vulnerabilities require manual review")
			fmt.Println("     Run 'npm audit' for details")
		}
	}

	// Update outdated packages (if package.json exists)
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		fmt.Println("  💡 Tip: Run 'npm outdated' to check for outdated packages")
		fmt.Println("     Then run 'npm update' to update them")
	}

	return fixes
}

func fixGoIssues(projectRoot string) []string {
	fmt.Println("🔧 Fixing Go issues...")
	var fixes []string

	// Run go mod tidy
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectRoot
	if err := cmd.Run(); err == nil {
		fmt.Println("  ✅ go mod tidy completed")
		fixes = append(fixes, "go mod tidy")
	}

	// Run go fmt
	cmd = exec.Command("go", "fmt", "./...")
	cmd.Dir = projectRoot
	if err := cmd.Run(); err == nil {
		fmt.Println("  ✅ go fmt completed")
		fixes = append(fixes, "go fmt")
	}

	// Run go vet
	cmd = exec.Command("go", "vet", "./...")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Println("  ✅ go vet completed - no issues found")
	} else {
		fmt.Printf("  ⚠️  go vet found issues:\n%s\n", string(output))
		fmt.Println("     These may require manual fixes")
	}

	return fixes
}

func fixLintingErrors(projectRoot string) []string {
	fmt.Println("🔍 Fixing linting errors...")
	var fixes []string

	// Check for ESLint (Node.js)
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		// Check if eslint is available
		if _, err := exec.LookPath("npx"); err == nil {
			cmd := exec.Command("npx", "eslint", "--fix", ".")
			cmd.Dir = projectRoot
			if err := cmd.Run(); err == nil {
				fmt.Println("  ✅ ESLint auto-fix completed")
				fixes = append(fixes, "eslint --fix")
			} else {
				fmt.Println("  ⚠️  ESLint not configured or no issues to fix")
			}
		}
	}

	// Check for golangci-lint (Go)
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
		if _, err := exec.LookPath("golangci-lint"); err == nil {
			cmd := exec.Command("golangci-lint", "run", "--fix")
			cmd.Dir = projectRoot
			if err := cmd.Run(); err == nil {
				fmt.Println("  ✅ golangci-lint auto-fix completed")
				fixes = append(fixes, "golangci-lint run --fix")
			}
		}
	}

	return fixes
}

func fixCommonIssues(projectRoot string) []string {
	fmt.Println("🔧 Fixing common issues...")
	var fixes []string

	// Ensure .gitignore exists and has common entries
	gitignorePath := filepath.Join(projectRoot, ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		fmt.Println("  ⚠️  .gitignore not found - creating one...")
		commonIgnores := []string{
			"# Dependencies",
			"node_modules/",
			"vendor/",
			"",
			"# Environment",
			".env",
			".env.local",
			".env.production",
			"",
			"# Build",
			"dist/",
			"build/",
			"*.exe",
			"",
			"# IDE",
			".vscode/",
			".idea/",
			"*.swp",
			"",
			"# OS",
			".DS_Store",
			"Thumbs.db",
		}
		content := strings.Join(commonIgnores, "\n")
		if err := os.WriteFile(gitignorePath, []byte(content), 0644); err == nil {
			fmt.Println("  ✅ Created .gitignore")
			fixes = append(fixes, "created .gitignore")
		}
	}

	// Check for .env.example
	envExamplePath := filepath.Join(projectRoot, ".env.example")
	if _, err := os.Stat(envExamplePath); os.IsNotExist(err) {
		envPath := filepath.Join(projectRoot, ".env")
		if _, err := os.Stat(envPath); err == nil {
			fmt.Println("  💡 Tip: Consider creating .env.example for documentation")
		}
	}

	return fixes
}

func displayFixResults(fixesApplied []string) {
	fmt.Println()
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("  Auto-fix Results")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()

	if len(fixesApplied) == 0 {
		fmt.Println("✅ No automatic fixes were needed!")
		fmt.Println()
		fmt.Println("💡 If you're experiencing issues, try:")
		fmt.Println("   • Running 'doplan security' to check for security issues")
		fmt.Println("   • Reviewing linter output manually")
		fmt.Println("   • Checking project documentation")
		fmt.Println("   • Using AI suggestions (coming soon)")
		return
	}

	fmt.Printf("✅ Applied %d fixes:\n\n", len(fixesApplied))
	for i, fix := range fixesApplied {
		fmt.Printf("  %d. %s\n", i+1, fix)
	}

	fmt.Println()
	fmt.Println("💡 Review the changes and commit if satisfied")
	fmt.Println("🤖 AI-powered suggestions coming soon in v0.0.19-beta")
}

// generateAISuggestions generates AI-powered fix suggestions
// This is a placeholder for future AI integration
func generateAISuggestions(projectRoot string, issues []string) []string {
	// TODO: Integrate with AI API (OpenAI, Anthropic, etc.)
	// For now, return empty suggestions
	fmt.Println("🤖 AI suggestions (coming soon)")
	return []string{}
}
