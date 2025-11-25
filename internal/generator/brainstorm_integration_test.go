package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestLoadPhaseTemplates simulates loading all phase templates as the /improve command would
func TestLoadPhaseTemplates(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	// Simulate what /improve command does: load all phase templates in order
	phaseFiles := []string{
		"phase-01-vision.md",
		"phase-02-audience.md",
		"phase-03-experience.md",
		"phase-04-content.md",
		"phase-05-marketing.md",
		"phase-06-delivery.md",
	}

	phases := make(map[string]string)
	questionsByPhase := make(map[string][]string)

	for _, phaseFile := range phaseFiles {
		path := filepath.Join(templatesDir, phaseFile)
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("Failed to load phase template %s: %v", phaseFile, err)
		}

		contentStr := string(content)
		phases[phaseFile] = contentStr

		// Extract questions (lines starting with -)
		questions := extractQuestions(contentStr)
		questionsByPhase[phaseFile] = questions

		// Verify we extracted questions
		if len(questions) == 0 {
			t.Errorf("Phase %s has no extractable questions", phaseFile)
		}

		// Verify questions are meaningful (not just dashes)
		for i, q := range questions {
			cleanQ := strings.TrimSpace(strings.TrimPrefix(q, "-"))
			if len(cleanQ) < 10 {
				t.Errorf("Phase %s question %d is too short: %q", phaseFile, i+1, cleanQ)
			}
		}
	}

	// Verify we loaded all 6 phases
	if len(phases) != 6 {
		t.Errorf("Expected 6 phases, loaded %d", len(phases))
	}

	// Verify total question count is reasonable (at least 3 per phase)
	totalQuestions := 0
	for phase, questions := range questionsByPhase {
		totalQuestions += len(questions)
		if len(questions) < 3 {
			t.Errorf("Phase %s has fewer than 3 questions (%d)", phase, len(questions))
		}
	}

	if totalQuestions < 18 {
		t.Errorf("Expected at least 18 total questions across all phases, found %d", totalQuestions)
	}

	t.Logf("Successfully loaded %d phases with %d total questions", len(phases), totalQuestions)
}

// TestLoadConfirmationTemplate simulates loading the confirmation template
func TestLoadConfirmationTemplate(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatePath := filepath.Join(projectRoot, ".plan", "templates", "brainstorm", "CONFIRMATION_TEMPLATE.md")

	content, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("Failed to load confirmation template: %v", err)
	}

	contentStr := string(content)

	// Verify template has placeholders that can be filled
	placeholders := []string{
		"[Project Name]",
		"[YYYY-MM-DD]",
		"[X minutes]",
		"[Answer]",
	}

	for _, placeholder := range placeholders {
		if !strings.Contains(contentStr, placeholder) {
			t.Errorf("Confirmation template missing placeholder: %s", placeholder)
		}
	}

	// Verify template has all required sections
	sections := []string{
		"Phase 01",
		"Phase 02",
		"Phase 03",
		"Phase 04",
		"Phase 05",
		"Phase 06",
		"Review & Confirm",
	}

	for _, section := range sections {
		if !strings.Contains(contentStr, section) {
			t.Errorf("Confirmation template missing section: %s", section)
		}
	}

	t.Log("Confirmation template loaded and validated successfully")
}

// TestLoadOutputTemplate simulates loading the output template for BRAINSTORM.md
func TestLoadOutputTemplate(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatePath := filepath.Join(projectRoot, ".plan", "templates", "brainstorm", "TEMPLATE_BRAINSTORM.md")

	content, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("Failed to load output template: %v", err)
	}

	contentStr := string(content)

	// Verify template has all phase sections
	phases := []string{
		"Phase 01: Vision & Outcomes",
		"Phase 02: Audience & Differentiation",
		"Phase 03: Experience, UI/UX & Tech",
		"Phase 04: Content & SEO",
		"Phase 05: Marketing & Growth",
		"Phase 06: Delivery, Ops & Risks",
	}

	for _, phase := range phases {
		if !strings.Contains(contentStr, phase) {
			t.Errorf("Output template missing phase section: %s", phase)
		}
	}

	// Verify template has metadata placeholders
	metadata := []string{
		"[YYYY-MM-DD]",
		"[Project Name]",
	}

	for _, meta := range metadata {
		if !strings.Contains(contentStr, meta) {
			t.Errorf("Output template missing metadata placeholder: %s", meta)
		}
	}

	t.Log("Output template loaded and validated successfully")
}

// TestTemplateQuestionExtraction tests the question extraction logic
func TestTemplateQuestionExtraction(t *testing.T) {
	sampleTemplate := `# Phase 01 · Vision & Outcomes

- What problem are we solving and why now?
- What does success look like in 3, 6, and 12 months?
- What constraints should we respect?
`

	questions := extractQuestions(sampleTemplate)

	if len(questions) != 3 {
		t.Errorf("Expected 3 questions, extracted %d", len(questions))
	}

	expectedQuestions := []string{
		"- What problem are we solving and why now?",
		"- What does success look like in 3, 6, and 12 months?",
		"- What constraints should we respect?",
	}

	for i, expected := range expectedQuestions {
		if i >= len(questions) {
			t.Errorf("Missing question %d", i+1)
			continue
		}
		if strings.TrimSpace(questions[i]) != expected {
			t.Errorf("Question %d mismatch:\nExpected: %q\nGot: %q", i+1, expected, questions[i])
		}
	}
}

// TestTemplateCustomizationWorkflow simulates customizing a template
func TestTemplateCustomizationWorkflow(t *testing.T) {
	projectRoot := findProjectRoot(t)
	templatesDir := filepath.Join(projectRoot, ".plan", "templates", "brainstorm")

	// Test that we can read, modify, and validate a template
	originalPath := filepath.Join(templatesDir, "phase-01-vision.md")
	originalContent, err := os.ReadFile(originalPath)
	if err != nil {
		t.Fatalf("Failed to read original template: %v", err)
	}

	originalStr := string(originalContent)
	originalQuestions := extractQuestions(originalStr)

	// Simulate adding a custom question
	customQuestion := "- What is your budget range for this project?"
	modifiedContent := originalStr + "\n" + customQuestion

	// Verify we can extract the new question
	modifiedQuestions := extractQuestions(modifiedContent)
	if len(modifiedQuestions) != len(originalQuestions)+1 {
		t.Errorf("Expected %d questions after adding one, got %d", len(originalQuestions)+1, len(modifiedQuestions))
	}

	// Verify the custom question is included
	found := false
	for _, q := range modifiedQuestions {
		if strings.Contains(q, "budget range") {
			found = true
			break
		}
	}
	if !found {
		t.Error("Custom question not found in modified template")
	}

	t.Log("Template customization workflow validated successfully")
}

// extractQuestions extracts questions (lines starting with -) from template content
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
