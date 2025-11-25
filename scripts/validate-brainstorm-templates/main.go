package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Find project root
	projectRoot := findProjectRoot()
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	fmt.Println("🔍 Validating Brainstorm Template System")
	fmt.Println("=======================================")
	fmt.Println()

	// Check all required templates exist
	fmt.Println("📋 Checking Required Templates...")
	requiredTemplates := []string{
		"phase-01-vision.md",
		"phase-02-audience.md",
		"phase-03-experience.md",
		"phase-04-content.md",
		"phase-05-marketing.md",
		"phase-06-delivery.md",
		"CONFIRMATION_TEMPLATE.md",
		"TEMPLATE_BRAINSTORM.md",
		"README.md",
	}

	allExist := true
	for _, template := range requiredTemplates {
		path := filepath.Join(templatesDir, template)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("  ❌ Missing: %s\n", template)
			allExist = false
		} else {
			fmt.Printf("  ✅ Found: %s\n", template)
		}
	}

	if !allExist {
		fmt.Println("\n❌ Some required templates are missing!")
		os.Exit(1)
	}

	// Validate phase templates
	fmt.Println("\n📝 Validating Phase Templates...")
	phaseCount := 0
	totalQuestions := 0

	for i := 1; i <= 6; i++ {
		phaseFile := fmt.Sprintf("phase-%02d-*.md", i)
		matches, _ := filepath.Glob(filepath.Join(templatesDir, phaseFile))
		if len(matches) > 0 {
			phaseCount++
			phasePath := matches[0]
			content, err := os.ReadFile(phasePath)
			if err != nil {
				fmt.Printf("  ❌ Error reading %s: %v\n", filepath.Base(phasePath), err)
				continue
			}

			questions := extractQuestions(string(content))
			totalQuestions += len(questions)
			fmt.Printf("  ✅ %s: %d questions\n", filepath.Base(phasePath), len(questions))
		}
	}

	if phaseCount != 6 {
		fmt.Printf("\n❌ Expected 6 phase templates, found %d\n", phaseCount)
		os.Exit(1)
	}

	fmt.Printf("\n📊 Summary: %d phases, %d total questions\n", phaseCount, totalQuestions)

	// Validate confirmation template
	fmt.Println("\n✅ Validating Confirmation Template...")
	confPath := filepath.Join(templatesDir, "CONFIRMATION_TEMPLATE.md")
	confContent, err := os.ReadFile(confPath)
	if err == nil {
		confStr := string(confContent)
		requiredSections := []string{
			"Review & Confirm",
			"Looks good, save it",
			"I want to revise",
		}
		allSections := true
		for _, section := range requiredSections {
			if !strings.Contains(confStr, section) {
				fmt.Printf("  ❌ Missing section: %s\n", section)
				allSections = false
			}
		}
		if allSections {
			fmt.Println("  ✅ All required sections present")
		}
	}

	// Validate output template
	fmt.Println("\n📄 Validating Output Template...")
	outputPath := filepath.Join(templatesDir, "TEMPLATE_BRAINSTORM.md")
	outputContent, err := os.ReadFile(outputPath)
	if err == nil {
		outputStr := string(outputContent)
		requiredPhases := []string{
			"Phase 01:",
			"Phase 02:",
			"Phase 03:",
			"Phase 04:",
			"Phase 05:",
			"Phase 06:",
		}
		allPhases := true
		for _, phase := range requiredPhases {
			if !strings.Contains(outputStr, phase) {
				fmt.Printf("  ❌ Missing phase: %s\n", phase)
				allPhases = false
			}
		}
		if allPhases {
			fmt.Println("  ✅ All phase sections present")
		}
	}

	fmt.Println("\n✨ Template system validation complete!")
	fmt.Println("\n💡 To customize templates, edit files in:")
	fmt.Printf("   %s\n", templatesDir)
}

func extractQuestions(content string) []string {
	questions := []string{}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "-") && len(trimmed) > 2 {
			questions = append(questions, trimmed)
		}
	}

	return questions
}

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		planDir := filepath.Join(dir, ".plan")
		if _, err := os.Stat(planDir); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}
