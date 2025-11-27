package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGetAgentTemplate tests loading agent template from embedded files
func TestGetAgentTemplate(t *testing.T) {
	templateStr, err := GetAgentTemplate()
	if err != nil {
		t.Fatalf("GetAgentTemplate() error = %v", err)
	}
	if templateStr == "" {
		t.Error("GetAgentTemplate() should return non-empty template")
	}
	if !strings.Contains(templateStr, "{{.Name}}") {
		t.Error("Template should contain {{.Name}}")
	}
	if !strings.Contains(templateStr, "{{.Role}}") {
		t.Error("Template should contain {{.Role}}")
	}
	if !strings.Contains(templateStr, "{{.SystemPrompt}}") {
		t.Error("Template should contain {{.SystemPrompt}}")
	}
}

// TestGetCommandTemplate tests loading command template from embedded files
func TestGetCommandTemplate(t *testing.T) {
	templateStr, err := GetCommandTemplate()
	if err != nil {
		t.Fatalf("GetCommandTemplate() error = %v", err)
	}
	if templateStr == "" {
		t.Error("GetCommandTemplate() should return non-empty template")
	}
	if !strings.Contains(templateStr, "{{.Name}}") {
		t.Error("Template should contain {{.Name}}")
	}
	if !strings.Contains(templateStr, "{{.Trigger}}") {
		t.Error("Template should contain {{.Trigger}}")
	}
	if !strings.Contains(templateStr, "{{.Action}}") {
		t.Error("Template should contain {{.Action}}")
	}
}

// TestRenderAgentMarkdownFileBased tests rendering agent with file-based template
func TestRenderAgentMarkdownFileBased(t *testing.T) {
	agent := &Agent{
		Name:         "Test Agent",
		Role:         "Test Role",
		SystemPrompt: "Test system prompt",
		ReportsTo:    "Manager",
		Manages:      []string{"Agent1", "Agent2"},
		Responsibilities: []string{
			"Responsibility 1",
			"Responsibility 2",
		},
	}

	rendered, err := RenderAgentMarkdownFileBased(agent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdownFileBased() error = %v", err)
	}
	if !strings.Contains(rendered, "Test Agent") {
		t.Error("Rendered should contain Test Agent")
	}
	if !strings.Contains(rendered, "Test Role") {
		t.Error("Rendered should contain Test Role")
	}
	if !strings.Contains(rendered, "Test system prompt") {
		t.Error("Rendered should contain Test system prompt")
	}
	if !strings.Contains(rendered, "Manager") {
		t.Error("Rendered should contain Manager")
	}
}

// TestRenderAgentMarkdownFileBased_NilAgent tests error handling for nil agent
func TestRenderAgentMarkdownFileBased_NilAgent(t *testing.T) {
	_, err := RenderAgentMarkdownFileBased(nil)
	if err == nil {
		t.Error("RenderAgentMarkdownFileBased() should return error for nil agent")
	}
	if !strings.Contains(err.Error(), "cannot be nil") {
		t.Errorf("Error should mention 'cannot be nil', got: %v", err)
	}
}

// TestRenderCommandMarkdownFileBased tests rendering command with file-based template
func TestRenderCommandMarkdownFileBased(t *testing.T) {
	cmd := &Command{
		Name:        "test",
		Trigger:     "/test",
		Description: "Test command",
		Action:      "Test action",
		Category:    "core",
		AgentInvolvement: []string{
			"Agent1",
			"Agent2",
		},
		Examples: []string{
			"/test example",
		},
		FilesRead: []string{
			"file1.md",
		},
		FilesModified: []string{
			"file2.md",
		},
	}

	rendered, err := RenderCommandMarkdownFileBased(cmd)
	if err != nil {
		t.Fatalf("RenderCommandMarkdownFileBased() error = %v", err)
	}
	if !strings.Contains(rendered, "/test") {
		t.Error("Rendered should contain /test")
	}
	if !strings.Contains(rendered, "Test action") {
		t.Error("Rendered should contain Test action")
	}
	// Command template doesn't include Description field, only Trigger and Action
}

// TestRenderCommandMarkdownFileBased_NilCommand tests error handling for nil command
func TestRenderCommandMarkdownFileBased_NilCommand(t *testing.T) {
	_, err := RenderCommandMarkdownFileBased(nil)
	if err == nil {
		t.Error("RenderCommandMarkdownFileBased() should return error for nil command")
	}
	if !strings.Contains(err.Error(), "cannot be nil") {
		t.Errorf("Error should mention 'cannot be nil', got: %v", err)
	}
}

// TestLoadBrainstormPhaseTemplate tests loading brainstorm phase templates
func TestLoadBrainstormPhaseTemplate(t *testing.T) {
	phases := []string{"01", "02", "03", "04", "05", "06"}

	for _, phase := range phases {
		t.Run("Phase_"+phase, func(t *testing.T) {
			content, err := LoadBrainstormPhaseTemplate(phase)
			if err != nil {
				t.Fatalf("LoadBrainstormPhaseTemplate(%s) error = %v", phase, err)
			}
			if content == "" {
				t.Error("LoadBrainstormPhaseTemplate() should return non-empty content")
			}
			if !strings.Contains(content, "Phase "+phase) {
				t.Errorf("Content should contain 'Phase %s', got: %s", phase, content[:50])
			}
		})
	}
}

// TestLoadBrainstormPhaseTemplate_InvalidPhase tests error handling for invalid phase
func TestLoadBrainstormPhaseTemplate_InvalidPhase(t *testing.T) {
	_, err := LoadBrainstormPhaseTemplate("99")
	if err == nil {
		t.Error("LoadBrainstormPhaseTemplate() should return error for invalid phase")
	}
}

// TestLoadBrainstormConfirmationTemplate tests loading confirmation template
func TestLoadBrainstormConfirmationTemplate(t *testing.T) {
	content, err := LoadBrainstormConfirmationTemplate()
	if err != nil {
		t.Fatalf("LoadBrainstormConfirmationTemplate() error = %v", err)
	}
	if content == "" {
		t.Error("LoadBrainstormConfirmationTemplate() should return non-empty content")
	}
	if !strings.Contains(content, "Brainstorm Session Summary") {
		t.Error("Content should contain 'Brainstorm Session Summary'")
	}
	if !strings.Contains(content, "Review & Confirm") {
		t.Error("Content should contain 'Review & Confirm'")
	}
}

// TestLoadBrainstormOutputTemplate tests loading brainstorm output template
func TestLoadBrainstormOutputTemplate(t *testing.T) {
	content, err := LoadBrainstormOutputTemplate()
	if err != nil {
		t.Fatalf("LoadBrainstormOutputTemplate() error = %v", err)
	}
	if content == "" {
		t.Error("LoadBrainstormOutputTemplate() should return non-empty content")
	}
	if !strings.Contains(content, "Brainstorm Session") {
		t.Error("Content should contain 'Brainstorm Session'")
	}
}

// TestExtractTemplates tests extracting templates to a directory
func TestExtractTemplates(t *testing.T) {
	tmpDir := t.TempDir()

	err := ExtractTemplates(tmpDir)
	if err != nil {
		t.Fatalf("ExtractTemplates() error = %v", err)
	}

	// Verify directory structure was created
	expectedDirs := []string{
		"agents",
		"commands",
		"documents/brainstorm",
	}

	for _, dir := range expectedDirs {
		dirPath := filepath.Join(tmpDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Errorf("Directory %s should exist", dirPath)
		}
	}

	// Verify agent template exists
	agentTemplatePath := filepath.Join(tmpDir, "agents", "agent.md.tmpl")
	if _, err := os.Stat(agentTemplatePath); os.IsNotExist(err) {
		t.Error("Agent template should exist")
	}

	content, err := os.ReadFile(agentTemplatePath)
	if err != nil {
		t.Fatalf("Failed to read agent template: %v", err)
	}
	if !strings.Contains(string(content), "{{.Name}}") {
		t.Error("Agent template should contain {{.Name}}")
	}

	// Verify command template exists
	commandTemplatePath := filepath.Join(tmpDir, "commands", "command.md.tmpl")
	if _, err := os.Stat(commandTemplatePath); os.IsNotExist(err) {
		t.Error("Command template should exist")
	}

	content, err = os.ReadFile(commandTemplatePath)
	if err != nil {
		t.Fatalf("Failed to read command template: %v", err)
	}
	if !strings.Contains(string(content), "{{.Name}}") {
		t.Error("Command template should contain {{.Name}}")
	}

	// Verify brainstorm phase templates exist
	phaseFiles := []string{
		"phase-01-vision.md",
		"phase-02-audience.md",
		"phase-03-experience.md",
		"phase-04-content.md",
		"phase-05-marketing.md",
		"phase-06-delivery.md",
	}

	for _, phaseFile := range phaseFiles {
		phasePath := filepath.Join(tmpDir, "documents", "brainstorm", phaseFile)
		if _, err := os.Stat(phasePath); os.IsNotExist(err) {
			t.Errorf("Phase template %s should exist", phasePath)
		}
	}

	// Verify confirmation template exists
	confirmationPath := filepath.Join(tmpDir, "documents", "brainstorm", "CONFIRMATION_TEMPLATE.md")
	if _, err := os.Stat(confirmationPath); os.IsNotExist(err) {
		t.Error("Confirmation template should exist")
	}

	// Verify output template exists
	outputPath := filepath.Join(tmpDir, "documents", "brainstorm", "TEMPLATE_BRAINSTORM.md")
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("Output template should exist")
	}
}

// TestExtractTemplates_InvalidPath tests error handling for invalid target directory
func TestExtractTemplates_InvalidPath(t *testing.T) {
	err := ExtractTemplates("/non/existent/dir/templates")
	if err == nil {
		t.Error("ExtractTemplates() should return error for invalid path")
	}
	if !strings.Contains(err.Error(), "failed to create directory") {
		t.Errorf("Error should mention 'failed to create directory', got: %v", err)
	}
}

// TestLoadTemplate tests loading a template from embedded files
func TestLoadTemplate(t *testing.T) {
	tmpl, err := LoadTemplate("agents/agent.md.tmpl")
	if err != nil {
		t.Fatalf("LoadTemplate() error = %v", err)
	}
	if tmpl == nil {
		t.Error("LoadTemplate() should return non-nil template")
	}

	// Try to execute the template
	agent := &Agent{
		Name: "Test Agent",
		Role: "Test Role",
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, agent)
	if err != nil {
		t.Fatalf("Template execution error = %v", err)
	}
	if !strings.Contains(buf.String(), "Test Agent") {
		t.Error("Executed template should contain 'Test Agent'")
	}
}

// TestLoadTemplateFromFile tests loading a template from disk
func TestLoadTemplateFromFile(t *testing.T) {
	tmpDir := t.TempDir()
	templatePath := filepath.Join(tmpDir, "test.tmpl")
	templateContent := `# {{.Name}}

{{.Description}}`

	if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}

	tmpl, err := LoadTemplateFromFile(templatePath)
	if err != nil {
		t.Fatalf("LoadTemplateFromFile() error = %v", err)
	}
	if tmpl == nil {
		t.Error("LoadTemplateFromFile() should return non-nil template")
	}

	// Try to execute the template
	data := map[string]string{
		"Name":        "Test",
		"Description": "Description",
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Template execution error = %v", err)
	}
	if !strings.Contains(buf.String(), "Test") {
		t.Error("Executed template should contain 'Test'")
	}
	if !strings.Contains(buf.String(), "Description") {
		t.Error("Executed template should contain 'Description'")
	}
}

// TestLoadTemplateFromFile_InvalidPath tests error handling for invalid file path
func TestLoadTemplateFromFile_InvalidPath(t *testing.T) {
	_, err := LoadTemplateFromFile("/non/existent/path.tmpl")
	if err == nil {
		t.Error("LoadTemplateFromFile() should return error for invalid path")
	}
	if !strings.Contains(err.Error(), "failed to read template file") {
		t.Errorf("Error should mention 'failed to read template file', got: %v", err)
	}
}

// TestRenderAgentMarkdown_FileBasedFallback tests that RenderAgentMarkdown falls back correctly
func TestRenderAgentMarkdown_FileBasedFallback(t *testing.T) {
	agent := &Agent{
		Name:         "Test Agent",
		Role:         "Test Role",
		SystemPrompt: "Test prompt",
	}

	rendered, err := RenderAgentMarkdown(agent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdown() error = %v", err)
	}
	if !strings.Contains(rendered, "Test Agent") {
		t.Error("Rendered should contain 'Test Agent'")
	}
	if !strings.Contains(rendered, "Test Role") {
		t.Error("Rendered should contain 'Test Role'")
	}
	if !strings.Contains(rendered, "Test prompt") {
		t.Error("Rendered should contain 'Test prompt'")
	}
}

// TestRenderCommandMarkdown_FileBasedFallback tests that RenderCommandMarkdown falls back correctly
func TestRenderCommandMarkdown_FileBasedFallback(t *testing.T) {
	cmd := &Command{
		Name:        "test",
		Trigger:     "/test",
		Description: "Test command",
		Action:      "Test action",
	}

	rendered, err := RenderCommandMarkdown(cmd)
	if err != nil {
		t.Fatalf("RenderCommandMarkdown() error = %v", err)
	}
	if !strings.Contains(rendered, "/test") {
		t.Error("Rendered should contain '/test'")
	}
	if !strings.Contains(rendered, "Test action") {
		t.Error("Rendered should contain 'Test action'")
	}
	// Note: Description field is not in the template, only Trigger and Action are rendered
}
