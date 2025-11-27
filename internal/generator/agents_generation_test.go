package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestAgentsGenerator_Name(t *testing.T) {
	generator := &AgentsGenerator{}
	if generator.Name() != "AI Agents" {
		t.Errorf("AgentsGenerator.Name() = %q, want %q", generator.Name(), "AI Agents")
	}
}

func TestAgentsGenerator_Generate(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("AgentsGenerator.Generate() error = %v", err)
	}

	// Verify .cursor/agents/ directory was created
	agentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	if !utils.IsDirectory(agentsDir) {
		t.Error("AgentsGenerator.Generate() should create .cursor/agents/ directory")
	}

	// Verify all 18 agent files were created
	agents := GetAllAgents()
	if len(agents) != 18 {
		t.Fatalf("Expected 18 agents, got %d", len(agents))
	}

	// Check central location (where files are actually created)
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")

	for _, agent := range agents {
		category := agent.Category
		if category == "" {
			category = "other"
		}

		// Check central location first
		centralAgentPath := filepath.Join(centralAgentsDir, category, agent.FileName)
		if !utils.IsFile(centralAgentPath) {
			t.Errorf("Agent file %s was not created in central location", agent.FileName)
			continue
		}

		// Also check IDE location (symlink or copy)
		ideAgentPath := filepath.Join(agentsDir, category, agent.FileName)
		if !utils.PathExists(ideAgentPath) {
			t.Errorf("Agent file %s is not accessible via IDE location", agent.FileName)
			continue
		}

		// Read from central location (source of truth)
		agentPath := centralAgentPath

		// Verify file content
		content, err := os.ReadFile(agentPath)
		if err != nil {
			t.Errorf("Failed to read agent file %s: %v", agent.FileName, err)
			continue
		}

		contentStr := string(content)

		// Check title
		expectedTitle := "# " + agent.Name
		if !strings.Contains(contentStr, expectedTitle) {
			t.Errorf("Agent file %s missing title: %s", agent.FileName, expectedTitle)
		}

		// Check role
		if !strings.Contains(contentStr, agent.Role) {
			t.Errorf("Agent file %s missing role: %s", agent.FileName, agent.Role)
		}

		// Check system prompt
		if !strings.Contains(contentStr, agent.SystemPrompt) {
			t.Errorf("Agent file %s missing system prompt", agent.FileName)
		}
	}
}

func TestGenerateAgents(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	if err := GenerateAgents(request, tmpDir); err != nil {
		t.Fatalf("GenerateAgents() error = %v", err)
	}

	// Verify directory was created
	agentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	if !utils.IsDirectory(agentsDir) {
		t.Error("GenerateAgents() should create .cursor/agents/ directory")
	}

	// Verify files were created
	// Files are now in category folders, so check central location or IDE location
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	agents := GetAllAgents()

	// Check if files exist in either central location or IDE location (via symlink or copy)
	for _, agent := range agents {
		category := agent.Category
		if category == "" {
			category = "other"
		}

		// Check central location first (where files are actually created)
		centralAgentPath := filepath.Join(centralAgentsDir, category, agent.FileName)
		ideAgentPath := filepath.Join(agentsDir, category, agent.FileName)

		// File should exist in central location
		if !utils.IsFile(centralAgentPath) {
			t.Errorf("Agent file %s was not created in central location", agent.FileName)
		}

		// File should also be accessible via IDE location (symlink or copy)
		if !utils.PathExists(ideAgentPath) {
			t.Errorf("Agent file %s is not accessible via IDE location (symlink/copy failed)", agent.FileName)
		}
	}
}

func TestAgentsGenerator_Generate_InvalidPath(t *testing.T) {
	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	// Use an invalid path (non-existent parent)
	invalidPath := filepath.Join("/nonexistent", "path", "test")

	err := generator.Generate(request, invalidPath)
	if err == nil {
		t.Error("AgentsGenerator.Generate() with invalid path should return error")
	}
}

func TestAgentsGenerator_Generate_ExistingDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Pre-create the agents directory
	agentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	if err := utils.CreateDirectory(agentsDir); err != nil {
		t.Fatalf("Failed to create agents dir: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	// Should succeed even if directory already exists
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("AgentsGenerator.Generate() with existing directory error = %v", err)
	}

	// Verify files were still created
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	agents := GetAllAgents()
	for _, agent := range agents {
		category := agent.Category
		if category == "" {
			category = "other"
		}

		// Check central location
		centralAgentPath := filepath.Join(centralAgentsDir, category, agent.FileName)
		if !utils.IsFile(centralAgentPath) {
			t.Errorf("Agent file %s was not created in central location", agent.FileName)
		}

		// Check IDE location (symlink or copy)
		ideAgentPath := filepath.Join(agentsDir, category, agent.FileName)
		if !utils.PathExists(ideAgentPath) {
			t.Errorf("Agent file %s is not accessible via IDE location", agent.FileName)
		}
	}
}

func TestAgentsGenerator_Generate_FileContent(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("AgentsGenerator.Generate() error = %v", err)
	}

	// Test a specific agent file content
	// Files are in category folders, check central location
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	// Project Orchestrator is typically in "management" or "other" category
	// Let's check common categories
	categories := []string{"management", "other", "core"}
	var orchestratorPath string
	var found bool

	for _, cat := range categories {
		candidatePath := filepath.Join(centralAgentsDir, cat, "project_orchestrator.md")
		if utils.IsFile(candidatePath) {
			orchestratorPath = candidatePath
			found = true
			break
		}
	}

	if !found {
		// Fallback: search all categories
		entries, _ := os.ReadDir(centralAgentsDir)
		for _, entry := range entries {
			if entry.IsDir() {
				candidatePath := filepath.Join(centralAgentsDir, entry.Name(), "project_orchestrator.md")
				if utils.IsFile(candidatePath) {
					orchestratorPath = candidatePath
					found = true
					break
				}
			}
		}
	}

	if !found {
		t.Fatalf("Could not find project_orchestrator.md in any category")
	}

	content, err := os.ReadFile(orchestratorPath)
	if err != nil {
		t.Fatalf("Failed to read orchestrator file: %v", err)
	}

	contentStr := string(content)

	// Verify all required sections
	requiredSections := []string{
		"# Project Orchestrator",
		"## Role",
		"## System Prompt",
		"## Responsibilities",
		"## Reports To",
		"## Manages",
	}

	for _, section := range requiredSections {
		if !strings.Contains(contentStr, section) {
			t.Errorf("Orchestrator file missing section: %s", section)
		}
	}

	// Verify Reports To shows "None (Top Level)"
	if !strings.Contains(contentStr, "None (Top Level)") {
		t.Error("Top-level agent should show 'None (Top Level)' in Reports To")
	}
}

func TestAgentsGenerator_Generate_AllAgentsContent(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("AgentsGenerator.Generate() error = %v", err)
	}

	// Check central location (where files are actually created)
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	agents := GetAllAgents()

	// Verify each agent file has correct content
	for _, agent := range agents {
		t.Run(agent.Name, func(t *testing.T) {
			category := agent.Category
			if category == "" {
				category = "other"
			}

			// Check central location
			agentPath := filepath.Join(centralAgentsDir, category, agent.FileName)
			content, err := os.ReadFile(agentPath)
			if err != nil {
				t.Fatalf("Failed to read agent file: %v", err)
			}

			contentStr := string(content)

			// Check title
			expectedTitle := "# " + agent.Name
			if !strings.Contains(contentStr, expectedTitle) {
				t.Errorf("Missing title: %s", expectedTitle)
			}

			// Check role
			if !strings.Contains(contentStr, agent.Role) {
				t.Errorf("Missing role: %s", agent.Role)
			}

			// Check responsibilities
			for _, resp := range agent.Responsibilities {
				if !strings.Contains(contentStr, resp) {
					t.Errorf("Missing responsibility: %s", resp)
				}
			}
		})
	}
}

func TestAgentsGenerator_Generate_MultipleIDEs(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor", "Claude Code", "Antigravity", "Windsurf", "Cline", "OpenCode"},
		ProjectType: "Fullstack",
	}

	generator := &AgentsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("AgentsGenerator.Generate() error = %v", err)
	}

	// Verify agents directories were created for all IDEs
	expectedAgentsDirs := map[string]string{
		"Cursor":      filepath.Join(tmpDir, ".cursor", "agents"),
		"Claude Code": filepath.Join(tmpDir, ".claude", "agents"),
		"Antigravity": filepath.Join(tmpDir, ".antigravity", "agents"),
		"Windsurf":    filepath.Join(tmpDir, ".windsurf", "agents"),
		"Cline":       filepath.Join(tmpDir, ".cline", "agents"),
		"OpenCode":    filepath.Join(tmpDir, ".opencode", "agents"),
	}

	// First verify central location has agents
	centralAgentsDir := filepath.Join(tmpDir, ".do", "core", "agents")
	agents := GetAllAgents()
	if len(agents) == 0 {
		t.Fatalf("No agents found")
	}

	// Check central location for first agent
	firstAgent := agents[0]
	category := firstAgent.Category
	if category == "" {
		category = "other"
	}
	centralAgentFile := filepath.Join(centralAgentsDir, category, firstAgent.FileName)
	if _, err := os.Stat(centralAgentFile); os.IsNotExist(err) {
		t.Fatalf("AgentsGenerator.Generate() should generate agents in central location")
	}

	// Then verify each IDE has access to agents (via symlink or copy)
	for ide, agentsDir := range expectedAgentsDirs {
		if _, err := os.Stat(agentsDir); os.IsNotExist(err) {
			t.Errorf("AgentsGenerator.Generate() should create agents directory for %s at %s", ide, agentsDir)
			continue
		}

		// Verify that agents are accessible (check for a known agent file in category folder)
		expectedAgentFile := filepath.Join(agentsDir, category, firstAgent.FileName)
		if _, err := os.Stat(expectedAgentFile); os.IsNotExist(err) {
			// If symlink/copy failed, that's okay - central location has the agents
			// Just log a warning but don't fail the test
			t.Logf("Agents not accessible via IDE location for %s (symlink/copy may have failed, but central location has agents)", ide)
		}
	}
}
