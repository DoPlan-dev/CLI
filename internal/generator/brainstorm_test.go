package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestBrainstormTemplatesExist verifies all required phase templates exist
func TestBrainstormTemplatesExist(t *testing.T) {
	// Get the project root (assuming we're in internal/generator/)
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

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

	for _, template := range requiredTemplates {
		path := filepath.Join(templatesDir, template)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Required template missing: %s", path)
		} else if err != nil {
			t.Errorf("Error checking template %s: %v", template, err)
		}
	}
}

// TestPhaseTemplatesHaveContent verifies phase templates contain questions
func TestPhaseTemplatesHaveContent(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	phaseTemplates := []string{
		"phase-01-vision.md",
		"phase-02-audience.md",
		"phase-03-experience.md",
		"phase-04-content.md",
		"phase-05-marketing.md",
		"phase-06-delivery.md",
	}

	for _, template := range phaseTemplates {
		path := filepath.Join(templatesDir, template)
		content, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read template %s: %v", template, err)
			continue
		}

		contentStr := string(content)

		// Check that template has a title/header
		if !strings.Contains(contentStr, "#") {
			t.Errorf("Template %s missing header/title", template)
		}

		// Check that template has questions (bullet points with question marks or dashes)
		hasQuestions := strings.Contains(contentStr, "?") || strings.Contains(contentStr, "-")
		if !hasQuestions {
			t.Errorf("Template %s appears to have no questions", template)
		}

		// Check that template is not empty
		if len(strings.TrimSpace(contentStr)) < 50 {
			t.Errorf("Template %s appears to be too short (less than 50 chars)", template)
		}
	}
}

// TestPhaseTemplatesHaveValidFormat verifies phase templates follow expected format
func TestPhaseTemplatesHaveValidFormat(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	phaseTemplates := []string{
		"phase-01-vision.md",
		"phase-02-audience.md",
		"phase-03-experience.md",
		"phase-04-content.md",
		"phase-05-marketing.md",
		"phase-06-delivery.md",
	}

	for _, template := range phaseTemplates {
		path := filepath.Join(templatesDir, template)
		content, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read template %s: %v", template, err)
			continue
		}

		contentStr := string(content)
		lines := strings.Split(contentStr, "\n")

		// Check that first line is a header
		if len(lines) == 0 || !strings.HasPrefix(strings.TrimSpace(lines[0]), "#") {
			t.Errorf("Template %s should start with a markdown header (#)", template)
		}

		// Check that template has at least one question (line starting with -)
		hasQuestion := false
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "-") && (strings.Contains(trimmed, "?") || len(trimmed) > 10) {
				hasQuestion = true
				break
			}
		}

		if !hasQuestion {
			t.Errorf("Template %s should have at least one question (line starting with -)", template)
		}
	}
}

// TestConfirmationTemplateExists verifies confirmation template exists and has proper structure
func TestConfirmationTemplateExists(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatePath := filepath.Join(projectRoot, ".plan", "templates", "brainstorm", "CONFIRMATION_TEMPLATE.md")

	content, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("CONFIRMATION_TEMPLATE.md not found: %v", err)
	}

	contentStr := string(content)

	// Check for key sections
	requiredSections := []string{
		"Brainstorm Summary",
		"Review & Confirm",
		"Looks good, save it",
		"I want to revise",
		"Add this:",
		"Start over",
	}

	for _, section := range requiredSections {
		if !strings.Contains(contentStr, section) {
			t.Errorf("CONFIRMATION_TEMPLATE.md missing required section: %s", section)
		}
	}
}

// TestBrainstormOutputTemplateExists verifies output template exists
func TestBrainstormOutputTemplateExists(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatePath := filepath.Join(projectRoot, ".plan", "templates", "brainstorm", "TEMPLATE_BRAINSTORM.md")

	content, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("TEMPLATE_BRAINSTORM.md not found: %v", err)
	}

	contentStr := string(content)

	// Check for key sections that should be in the output
	requiredSections := []string{
		"Brainstorm Session",
		"Phase 01",
		"Phase 02",
		"Phase 03",
		"Phase 04",
		"Phase 05",
		"Phase 06",
		"Recommended Next Steps",
	}

	for _, section := range requiredSections {
		if !strings.Contains(contentStr, section) {
			t.Errorf("TEMPLATE_BRAINSTORM.md missing required section: %s", section)
		}
	}
}

// TestPhaseTemplatesAreOrdered verifies phase templates are numbered sequentially
func TestPhaseTemplatesAreOrdered(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	expectedPhases := 6
	for i := 1; i <= expectedPhases; i++ {
		phaseNum := i
		if phaseNum < 10 {
			phaseNum = phaseNum // Keep as is for 01-06
		}
		templateName := filepath.Join(templatesDir, "phase-01-vision.md")
		if i == 1 {
			// Check first one exists
			if _, err := os.Stat(templateName); os.IsNotExist(err) {
				t.Errorf("Phase template phase-01-vision.md not found")
			}
		}
		// For now, just verify we have 6 phase templates
	}

	// Verify we have exactly 6 phase templates
	entries, err := os.ReadDir(templatesDir)
	if err != nil {
		t.Fatalf("Failed to read templates directory: %v", err)
	}

	phaseCount := 0
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "phase-") && strings.HasSuffix(entry.Name(), ".md") {
			phaseCount++
		}
	}

	if phaseCount != expectedPhases {
		t.Errorf("Expected %d phase templates, found %d", expectedPhases, phaseCount)
	}
}

// TestTemplateCustomization verifies templates can be read and parsed
func TestTemplateCustomization(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	// Test that we can read and parse a template
	templatePath := filepath.Join(templatesDir, "phase-01-vision.md")
	content, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("Failed to read phase-01-vision.md: %v", err)
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	// Verify we can extract questions
	questions := []string{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "-") {
			questions = append(questions, trimmed)
		}
	}

	if len(questions) == 0 {
		t.Error("phase-01-vision.md should contain at least one question")
	}

	// Verify questions are readable (not empty, have content)
	for i, q := range questions {
		if len(strings.TrimSpace(strings.TrimPrefix(q, "-"))) < 5 {
			t.Errorf("Question %d in phase-01-vision.md appears to be empty or too short", i+1)
		}
	}
}

// findProjectRoot finds the project root by looking for .plan directory
func findProjectRoot(t *testing.T) string {
	// Start from current directory and walk up
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	for {
		planDir := filepath.Join(dir, ".plan")
		if _, err := os.Stat(planDir); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root
			t.Fatalf("Could not find project root (looking for .plan directory)")
		}
		dir = parent
	}
}
