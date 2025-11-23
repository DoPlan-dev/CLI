package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/doplan/cli/internal/utils"
	"github.com/doplan/cli/pkg/models"
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

	for _, agent := range agents {
		agentPath := filepath.Join(agentsDir, agent.FileName)
		if !utils.IsFile(agentPath) {
			t.Errorf("Agent file %s was not created", agent.FileName)
		}

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
	agents := GetAllAgents()
	for _, agent := range agents {
		agentPath := filepath.Join(agentsDir, agent.FileName)
		if !utils.IsFile(agentPath) {
			t.Errorf("Agent file %s was not created", agent.FileName)
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
	agents := GetAllAgents()
	for _, agent := range agents {
		agentPath := filepath.Join(agentsDir, agent.FileName)
		if !utils.IsFile(agentPath) {
			t.Errorf("Agent file %s was not created", agent.FileName)
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
	agentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	orchestratorPath := filepath.Join(agentsDir, "project_orchestrator.md")

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

	agentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	agents := GetAllAgents()

	// Verify each agent file has correct content
	for _, agent := range agents {
		t.Run(agent.Name, func(t *testing.T) {
			agentPath := filepath.Join(agentsDir, agent.FileName)
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

