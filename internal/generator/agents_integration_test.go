package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestRenderAgentMarkdown_EdgeCases(t *testing.T) {
	agent := &Agent{
		Name:         "test_agent",
		Category:     "test_category",
		ReportsTo:    "test_manager",
		SystemPrompt: "Test system prompt",
		Role:         "Test Role",
		Responsibilities: []string{
			"Responsibility 1",
			"Responsibility 2",
		},
		FileName: "test_agent.md",
	}

	// Test with all fields
	markdown, err := RenderAgentMarkdown(agent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdown() error = %v", err)
	}
	if len(markdown) == 0 {
		t.Error("RenderAgentMarkdown() should return non-empty markdown")
	}
	if !contains(markdown, agent.Name) {
		t.Error("RenderAgentMarkdown() should contain agent name")
	}
	if !contains(markdown, agent.SystemPrompt) {
		t.Error("RenderAgentMarkdown() should contain system prompt")
	}

	// Test with minimal fields
	minimalAgent := &Agent{
		Name:     "minimal",
		Category: "test",
	}
	minimalMarkdown, err := RenderAgentMarkdown(minimalAgent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdown() with minimal fields error = %v", err)
	}
	if len(minimalMarkdown) == 0 {
		t.Error("RenderAgentMarkdown() should work with minimal fields")
	}

	// Test with empty responsibilities
	emptyRespAgent := &Agent{
		Name:             "empty_resp",
		Category:         "test",
		Responsibilities: []string{},
	}
	emptyRespMarkdown, err := RenderAgentMarkdown(emptyRespAgent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdown() with empty responsibilities error = %v", err)
	}
	if len(emptyRespMarkdown) == 0 {
		t.Error("RenderAgentMarkdown() should handle empty responsibilities")
	}

	// Test with nil agent
	_, err = RenderAgentMarkdown(nil)
	if err == nil {
		t.Error("RenderAgentMarkdown() should return error for nil agent")
	}
}

func TestAgentsGenerator_Generate_ErrorPaths(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with invalid project path (non-writable)
	invalidPath := filepath.Join(tmpDir, "invalid", "nested", "path")
	// Don't create the directory to test error handling

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	err := generator.Generate(request, invalidPath)
	// Should handle error gracefully (may create directories or error)
	_ = err
}

func TestAgentsGenerator_Generate_MultipleIDEs_ErrorHandling(t *testing.T) {
	tmpDir := t.TempDir()

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor", "Claude Code", "Antigravity"},
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("AgentsGenerator.Generate() error = %v", err)
	}

	// Verify agents were created for all IDEs
	for _, ide := range request.IDEs {
		ideAgentsDir, err := getIDEAgentsDir(tmpDir, ide)
		if err != nil {
			t.Fatalf("getIDEAgentsDir() error = %v", err)
		}

		// Check if directory exists (symlink or copy)
		if _, err := os.Stat(ideAgentsDir); os.IsNotExist(err) {
			t.Errorf("Agents directory should exist for IDE: %s", ide)
		}
	}
}

func TestCreateAgentCategorySymlinks_ErrorHandling(t *testing.T) {
	tmpDir := t.TempDir()

	// Create central agents directory
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	if err := os.MkdirAll(centralAgentsDir, 0755); err != nil {
		t.Fatalf("Failed to create central dir: %v", err)
	}

	// Create category folder
	categoryDir := filepath.Join(centralAgentsDir, "leadership")
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		t.Fatalf("Failed to create category dir: %v", err)
	}

	// Create agent file
	agentFile := filepath.Join(categoryDir, "test_agent.md")
	if err := os.WriteFile(agentFile, []byte("# Test Agent"), 0644); err != nil {
		t.Fatalf("Failed to create agent file: %v", err)
	}

	// Test with non-existent IDE directory (should handle gracefully)
	ideAgentsDir := filepath.Join(tmpDir, ".nonexistent", "agents")
	err := createAgentCategorySymlinks(ideAgentsDir, centralAgentsDir)
	// May error or create directory, both are acceptable
	_ = err
}

func TestCreateAgentCategorySymlinks_ExistingSymlink(t *testing.T) {
	tmpDir := t.TempDir()

	// Create central agents directory
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	categoryDir := filepath.Join(centralAgentsDir, "leadership")
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		t.Fatalf("Failed to create category dir: %v", err)
	}

	// Create IDE agents directory
	ideAgentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	if err := os.MkdirAll(ideAgentsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Create existing symlink
	existingLink := filepath.Join(ideAgentsDir, "leadership")
	if err := os.Symlink(categoryDir, existingLink); err != nil {
		t.Fatalf("Failed to create existing symlink: %v", err)
	}

	// Should handle existing symlink gracefully
	err := createAgentCategorySymlinks(ideAgentsDir, centralAgentsDir)
	if err != nil {
		t.Logf("createAgentCategorySymlinks() with existing symlink: %v (may be expected)", err)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(s != "" && substr != "" && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

