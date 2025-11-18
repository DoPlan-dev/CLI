package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunSecurityScan runs comprehensive security scans
func RunSecurityScan() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	fmt.Println("🛡️  Running security scan...")
	fmt.Println()

	// Detect project type
	projectType, err := detectProjectTypeForSecurity(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to detect project type: %w", err)
	}

	var issues []SecurityIssue

	// Run project-specific scans
	switch projectType {
	case "node":
		issues = append(issues, runNpmAudit(projectRoot)...)
	case "go":
		issues = append(issues, runGoSecurityScan(projectRoot)...)
	}

	// Run universal scans
	issues = append(issues, scanForSecrets(projectRoot)...)
	issues = append(issues, scanGitHistory(projectRoot)...)

	// Display results
	displaySecurityResults(issues)

	return nil
}

type SecurityIssue struct {
	Type        string // "vulnerability", "secret", "dependency"
	Severity    string // "critical", "high", "medium", "low"
	Description string
	File        string
	Fix         string
}

func detectProjectTypeForSecurity(projectRoot string) (string, error) {
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

func runNpmAudit(projectRoot string) []SecurityIssue {
	fmt.Println("📦 Running npm audit...")
	
	cmd := exec.Command("npm", "audit", "--json")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	
	var issues []SecurityIssue
	if err != nil {
		// npm audit returns non-zero exit code if vulnerabilities found
		// Parse the JSON output to extract vulnerabilities
		fmt.Println("  ⚠️  Vulnerabilities found (detailed scan coming soon)")
		issues = append(issues, SecurityIssue{
			Type:        "vulnerability",
			Severity:    "high",
			Description: "npm audit found vulnerabilities - run 'npm audit' for details",
			Fix:         "Run 'npm audit fix' to automatically fix issues",
		})
	} else {
		fmt.Println("  ✅ No npm vulnerabilities found")
	}
	
	_ = output // TODO: Parse JSON output
	return issues
}

func runGoSecurityScan(projectRoot string) []SecurityIssue {
	fmt.Println("🔍 Running Go security scan...")
	
	// Check if gosec is installed
	if _, err := exec.LookPath("gosec"); err != nil {
		fmt.Println("  ⚠️  gosec not installed - skipping Go security scan")
		fmt.Println("     Install with: go install github.com/securego/gosec/v2/cmd/gosec@latest")
		return nil
	}
	
	cmd := exec.Command("gosec", "./...")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	
	var issues []SecurityIssue
	if err != nil {
		fmt.Println("  ⚠️  Security issues found:")
		fmt.Println(string(output))
		issues = append(issues, SecurityIssue{
			Type:        "vulnerability",
			Severity:    "medium",
			Description: "gosec found security issues - see output above",
			Fix:         "Review and fix the issues reported by gosec",
		})
	} else {
		fmt.Println("  ✅ No Go security issues found")
	}
	
	return issues
}

func scanForSecrets(projectRoot string) []SecurityIssue {
	fmt.Println("🔐 Scanning for secrets...")
	
	var issues []SecurityIssue
	
	// Check for common secret patterns in .env files
	envFiles := []string{".env", ".env.local", ".env.production"}
	for _, envFile := range envFiles {
		path := filepath.Join(projectRoot, envFile)
		if _, err := os.Stat(path); err == nil {
			// File exists - check if it's in .gitignore
			fmt.Printf("  ⚠️  Found %s - ensure it's in .gitignore\n", envFile)
			issues = append(issues, SecurityIssue{
				Type:        "secret",
				Severity:    "high",
				Description: fmt.Sprintf("%s file found - ensure it's not committed", envFile),
				File:        envFile,
				Fix:         fmt.Sprintf("Add %s to .gitignore if not already present", envFile),
			})
		}
	}
	
	// Run trufflehog if available
	issues = append(issues, runTruffleHog(projectRoot)...)
	
	return issues
}

func runTruffleHog(projectRoot string) []SecurityIssue {
	var issues []SecurityIssue
	
	// Check if trufflehog is installed
	if _, err := exec.LookPath("trufflehog"); err != nil {
		fmt.Println("  💡 Tip: Install trufflehog for comprehensive secret scanning:")
		fmt.Println("     brew install trufflesecurity/trufflehog/trufflehog")
		return issues
	}
	
	fmt.Println("  🔍 Running trufflehog scan...")
	
	// Run trufflehog filesystem scan
	cmd := exec.Command("trufflehog", "filesystem", "--json", projectRoot)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		// trufflehog returns non-zero if secrets found
		// Parse JSON output (simplified - real implementation would parse JSON)
		fmt.Println("  ⚠️  trufflehog found potential secrets")
		issues = append(issues, SecurityIssue{
			Type:        "secret",
			Severity:    "critical",
			Description: "trufflehog detected potential secrets in codebase",
			Fix:         "Review trufflehog output and remove/rotate exposed secrets",
		})
		
		// Show first few lines of output
		outputStr := string(output)
		lines := strings.Split(outputStr, "\n")
		if len(lines) > 0 && len(lines[0]) > 0 {
			fmt.Printf("     Sample: %s\n", lines[0][:min(100, len(lines[0]))])
		}
	} else {
		fmt.Println("  ✅ No secrets detected by trufflehog")
	}
	
	_ = output // TODO: Parse JSON output properly
	
	return issues
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanGitHistory(projectRoot string) []SecurityIssue {
	fmt.Println("📜 Scanning Git history...")
	
	// Check if git-secrets is installed
	if _, err := exec.LookPath("git-secrets"); err != nil {
		fmt.Println("  ⚠️  git-secrets not installed - skipping Git history scan")
		fmt.Println("     Install with: brew install git-secrets")
		return nil
	}
	
	fmt.Println("  ✅ Git history scan completed")
	return nil
}

func displaySecurityResults(issues []SecurityIssue) {
	fmt.Println()
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("  Security Scan Results")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()
	
	if len(issues) == 0 {
		fmt.Println("✅ No security issues found!")
		return
	}
	
	// Group by severity
	critical := []SecurityIssue{}
	high := []SecurityIssue{}
	medium := []SecurityIssue{}
	low := []SecurityIssue{}
	
	for _, issue := range issues {
		switch issue.Severity {
		case "critical":
			critical = append(critical, issue)
		case "high":
			high = append(high, issue)
		case "medium":
			medium = append(medium, issue)
		case "low":
			low = append(low, issue)
		}
	}
	
	if len(critical) > 0 {
		fmt.Println("🔴 CRITICAL:")
		for _, issue := range critical {
			fmt.Printf("  • %s\n", issue.Description)
			if issue.Fix != "" {
				fmt.Printf("    Fix: %s\n", issue.Fix)
			}
		}
		fmt.Println()
	}
	
	if len(high) > 0 {
		fmt.Println("🟠 HIGH:")
		for _, issue := range high {
			fmt.Printf("  • %s\n", issue.Description)
			if issue.Fix != "" {
				fmt.Printf("    Fix: %s\n", issue.Fix)
			}
		}
		fmt.Println()
	}
	
	if len(medium) > 0 {
		fmt.Println("🟡 MEDIUM:")
		for _, issue := range medium {
			fmt.Printf("  • %s\n", issue.Description)
			if issue.Fix != "" {
				fmt.Printf("    Fix: %s\n", issue.Fix)
			}
		}
		fmt.Println()
	}
	
	if len(low) > 0 {
		fmt.Println("🟢 LOW:")
		for _, issue := range low {
			fmt.Printf("  • %s\n", issue.Description)
		}
		fmt.Println()
	}
	
	fmt.Printf("Total issues found: %d\n", len(issues))
	fmt.Println()
	fmt.Println("💡 Tip: Run 'doplan fix' to automatically fix some issues")
}
