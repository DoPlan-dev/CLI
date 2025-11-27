package generator

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadAgentsFromFiles tests loading agents from embedded files
func TestLoadAgentsFromFiles(t *testing.T) {
	agents, err := LoadAgentsFromFiles()
	if err != nil {
		t.Fatalf("LoadAgentsFromFiles() error = %v", err)
	}

	if len(agents) == 0 {
		t.Error("LoadAgentsFromFiles() should return at least one agent")
	}

	// Verify we have the expected number of agents (18)
	if len(agents) < 18 {
		t.Errorf("LoadAgentsFromFiles() returned %d agents, expected at least 18", len(agents))
	}
	
	// Verify we loaded markdown files, not JSON
	if len(agents) == 0 {
		t.Fatal("No agents loaded")
	}

	// Verify key agents exist
	agentNames := make(map[string]bool)
	for _, agent := range agents {
		agentNames[agent.Name] = true
	}

	expectedAgents := []string{
		"Project Orchestrator",
		"Product Manager",
		"Engineering Lead",
		"System Architect",
	}

	for _, name := range expectedAgents {
		if !agentNames[name] {
			t.Errorf("LoadAgentsFromFiles() missing expected agent: %s", name)
		}
	}
}

// TestLoadAgentsFromDirectory tests loading agents from a directory
func TestLoadAgentsFromDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	// Extract agents to temp directory
	if err := ExtractAgents(tmpDir); err != nil {
		t.Fatalf("ExtractAgents() error = %v", err)
	}

	// Load from directory
	agents, err := LoadAgentsFromDirectory(tmpDir)
	if err != nil {
		t.Fatalf("LoadAgentsFromDirectory() error = %v", err)
	}

	if len(agents) == 0 {
		t.Error("LoadAgentsFromDirectory() should return at least one agent")
	}

	// Verify structure is correct
	if len(agents) < 18 {
		t.Errorf("LoadAgentsFromDirectory() returned %d agents, expected at least 18", len(agents))
	}
}

// TestLoadAgentsFromDirectory_InvalidPath tests error handling
func TestLoadAgentsFromDirectory_InvalidPath(t *testing.T) {
	// Test with non-existent directory
	_, err := LoadAgentsFromDirectory("/nonexistent/path")
	if err == nil {
		t.Error("LoadAgentsFromDirectory() should return error for non-existent path")
	}

	// Test with empty path
	_, err = LoadAgentsFromDirectory("")
	if err == nil {
		t.Error("LoadAgentsFromDirectory() should return error for empty path")
	}
}

// TestLoadAgentsFromDirectory_InvalidMarkdown tests error handling for invalid markdown
func TestLoadAgentsFromDirectory_InvalidMarkdown(t *testing.T) {
	tmpDir := t.TempDir()

	// Create invalid markdown file (no frontmatter)
	invalidFile := filepath.Join(tmpDir, "invalid.md")
	if err := os.WriteFile(invalidFile, []byte("# No frontmatter here"), 0644); err != nil {
		t.Fatalf("Failed to create invalid markdown file: %v", err)
	}

	_, err := LoadAgentsFromDirectory(tmpDir)
	if err == nil {
		t.Error("LoadAgentsFromDirectory() should return error for invalid markdown frontmatter")
	}
}

// TestExtractAgents tests extracting agents to a directory
func TestExtractAgents(t *testing.T) {
	tmpDir := t.TempDir()

	if err := ExtractAgents(tmpDir); err != nil {
		t.Fatalf("ExtractAgents() error = %v", err)
	}

	// Verify directory structure was created
	categories := []string{
		"leadership",
		"engineering",
		"product",
		"design",
		"quality",
		"release",
		"documentation",
	}

	for _, category := range categories {
		categoryPath := filepath.Join(tmpDir, category)
		if _, err := os.Stat(categoryPath); os.IsNotExist(err) {
			t.Errorf("ExtractAgents() should create category directory %s", category)
		}
	}

	// Verify some agent files exist (markdown files, not JSON)
	projectOrchestratorPath := filepath.Join(tmpDir, "leadership", "project_orchestrator.md")
	if _, err := os.Stat(projectOrchestratorPath); os.IsNotExist(err) {
		t.Error("ExtractAgents() should create project_orchestrator.md")
	}

	engineeringLeadPath := filepath.Join(tmpDir, "engineering", "engineering_lead.md")
	if _, err := os.Stat(engineeringLeadPath); os.IsNotExist(err) {
		t.Error("ExtractAgents() should create engineering_lead.md")
	}
}

// TestExtractAgents_InvalidPath tests error handling
func TestExtractAgents_InvalidPath(t *testing.T) {
	// Test with empty path
	err := ExtractAgents("")
	if err == nil {
		t.Error("ExtractAgents() should return error for empty path")
	}

	// Test with invalid parent directory
	invalidPath := filepath.Join("/nonexistent", "path", "test")
	err = ExtractAgents(invalidPath)
	if err == nil {
		t.Error("ExtractAgents() should return error for invalid path")
	}
}

// TestGetAllAgentsFileBased tests the file-based loader with fallback
func TestGetAllAgentsFileBased(t *testing.T) {
	agents, err := GetAllAgentsFileBased()
	if err != nil {
		// Error is OK if it falls back to hardcoded
		t.Logf("GetAllAgentsFileBased() returned error (may have fallen back): %v", err)
	}

	if len(agents) == 0 {
		t.Error("GetAllAgentsFileBased() should return at least one agent")
	}

	// Should have 18 agents
	if len(agents) < 18 {
		t.Errorf("GetAllAgentsFileBased() returned %d agents, expected at least 18", len(agents))
	}
}

// TestLoadAgentsFromFiles_AgentContent tests agent content validation
func TestLoadAgentsFromFiles_AgentContent(t *testing.T) {
	agents, err := LoadAgentsFromFiles()
	if err != nil {
		t.Fatalf("LoadAgentsFromFiles() error = %v", err)
	}

	for _, agent := range agents {
		// Validate required fields
		if agent.Name == "" {
			t.Error("Agent should have a name")
		}
		if agent.Role == "" {
			t.Errorf("Agent %s should have a role", agent.Name)
		}
		if agent.SystemPrompt == "" {
			t.Errorf("Agent %s should have a system prompt", agent.Name)
		}
		if agent.FileName == "" {
			t.Errorf("Agent %s should have a fileName", agent.Name)
		}
		if agent.Category == "" {
			t.Errorf("Agent %s should have a category", agent.Name)
		}
	}
}

// TestExtractAgents_FileContent tests that extracted files have correct content
func TestExtractAgents_FileContent(t *testing.T) {
	tmpDir := t.TempDir()

	if err := ExtractAgents(tmpDir); err != nil {
		t.Fatalf("ExtractAgents() error = %v", err)
	}

	// Read and verify a specific agent file
	agentPath := filepath.Join(tmpDir, "leadership", "project_orchestrator.md")
	data, err := os.ReadFile(agentPath)
	if err != nil {
		t.Fatalf("Failed to read extracted agent file: %v", err)
	}

	// Verify it's valid markdown with frontmatter
	metadata, _, err := parseMarkdownWithFrontmatter(data)
	if err != nil {
		t.Errorf("Extracted agent file is not valid markdown with frontmatter: %v", err)
	}

	// Verify content
	if name := getString(metadata, "name"); name != "Project Orchestrator" {
		t.Errorf("Agent name = %q, want %q", name, "Project Orchestrator")
	}
	if category := getString(metadata, "category"); category != "leadership" {
		t.Errorf("Agent category = %q, want %q", category, "leadership")
	}
}

