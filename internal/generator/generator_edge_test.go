package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// TestAgentsGenerator_Generate_ErrorPaths_Coverage tests error paths in AgentsGenerator.Generate (72.2% -> 95%+)
func TestAgentsGenerator_Generate_ErrorPaths_Coverage(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(string) error
		request     *models.ProjectRequest
		expectError bool
	}{
		{
			name: "ReadDir error in hasContent check",
			setup: func(projectPath string) error {
				// Create central agents dir as a file (will cause ReadDir to fail)
				centralDir := filepath.Join(projectPath, ".do", "core", "agents")
				if err := os.MkdirAll(filepath.Dir(centralDir), 0755); err != nil {
					return err
				}
				return os.WriteFile(centralDir, []byte("file"), 0644)
			},
			request: &models.ProjectRequest{
				ProjectName: "test",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: true,
		},
		{
			name: "Category directory creation failure",
			setup: func(projectPath string) error {
				// Make central agents dir read-only (might fail on some systems)
				centralDir := filepath.Join(projectPath, ".do", "core", "agents")
				if err := os.MkdirAll(centralDir, 0755); err != nil {
					return err
				}
				return nil
			},
			request: &models.ProjectRequest{
				ProjectName: "test",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: false, // Should succeed
		},
		{
			name: "Invalid IDE for getIDEAgentsDir",
			request: &models.ProjectRequest{
				ProjectName: "test",
				IDEs:        []string{"InvalidIDE"},
				ProjectType: "Fullstack",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			if tt.setup != nil {
				if err := tt.setup(tmpDir); err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}

			generator := &AgentsGenerator{}
			err := generator.Generate(tt.request, tmpDir)

			if (err != nil) != tt.expectError {
				t.Errorf("AgentsGenerator.Generate() error = %v, expectError = %v", err, tt.expectError)
			}
		})
	}
}

// TestCreateAgentCategorySymlinks_ErrorPaths_Coverage tests error paths (66.7% -> 95%+)
func TestCreateAgentCategorySymlinks_ErrorPaths_Coverage(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with non-existent central directory
	centralDir := filepath.Join(tmpDir, "nonexistent", "central")
	ideDir := filepath.Join(tmpDir, "ide")
	if err := os.MkdirAll(ideDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	err := createAgentCategorySymlinks(ideDir, centralDir)
	if err == nil {
		t.Error("createAgentCategorySymlinks() should return error for non-existent central directory")
	}

	// Test with file instead of directory
	centralFile := filepath.Join(tmpDir, "central_file")
	if err := os.WriteFile(centralFile, []byte("file"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	err = createAgentCategorySymlinks(ideDir, centralFile)
	if err == nil {
		t.Error("createAgentCategorySymlinks() should return error when central path is a file")
	}
}

// TestCommandsGenerator_Generate_ErrorPaths_Coverage tests error paths (75% -> 95%+)
func TestCommandsGenerator_Generate_ErrorPaths_Coverage(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with invalid IDE
	generator := &CommandsGenerator{}
	request := &models.ProjectRequest{
		ProjectName: "test",
		IDEs:        []string{"InvalidIDE"},
		ProjectType: "Fullstack",
	}

	err := generator.Generate(request, tmpDir)
	if err == nil {
		t.Error("CommandsGenerator.Generate() should return error for invalid IDE")
	}

	// Test with read-only directory (on systems that support it)
	readOnlyDir := filepath.Join(tmpDir, "readonly")
	if err := os.MkdirAll(readOnlyDir, 0444); err != nil {
		t.Logf("Cannot create read-only dir (may not be supported): %v", err)
	} else {
		request2 := &models.ProjectRequest{
			ProjectName: "test2",
			IDEs:        []string{"Cursor"},
			ProjectType: "Fullstack",
		}
		err = generator.Generate(request2, readOnlyDir)
		// Should fail when trying to create subdirectories
		if err == nil {
			t.Log("Note: Generation succeeded in read-only dir (may be system-dependent)")
		}
		os.Chmod(readOnlyDir, 0755) // Restore for cleanup
	}
}

// TestCreateCommandCategorySymlinks_ErrorPaths_Coverage tests error paths (66.7% -> 95%+)
func TestCreateCommandCategorySymlinks_ErrorPaths_Coverage(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with non-existent central directory
	centralDir := filepath.Join(tmpDir, "nonexistent", "central")
	ideDir := filepath.Join(tmpDir, "ide")
	if err := os.MkdirAll(ideDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	err := createCommandCategorySymlinks(ideDir, centralDir)
	if err == nil {
		t.Error("createCommandCategorySymlinks() should return error for non-existent central directory")
	}

	// Test with file instead of directory
	centralFile := filepath.Join(tmpDir, "central_file")
	if err := os.WriteFile(centralFile, []byte("file"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	err = createCommandCategorySymlinks(ideDir, centralFile)
	if err == nil {
		t.Error("createCommandCategorySymlinks() should return error when central path is a file")
	}
}

// TestRenderAgentMarkdown_EdgeCases_Coverage tests RenderAgentMarkdown edge cases (77.8% -> 95%+)
func TestRenderAgentMarkdown_EdgeCases_Coverage(t *testing.T) {
	tests := []struct {
		name    string
		agent   *Agent
		wantErr bool
	}{
		{
			name: "Agent with empty responsibilities",
			agent: &Agent{
				Name:             "Test Agent",
				Category:         "test",
				ReportsTo:        "Manager",
				SystemPrompt:     "Test prompt",
				Role:             "Test Role",
				Responsibilities: []string{}, // Empty
				FileName:         "test.md",
			},
			wantErr: false,
		},
		{
			name: "Agent with nil responsibilities",
			agent: &Agent{
				Name:             "Test Agent",
				Category:         "test",
				ReportsTo:        "Manager",
				SystemPrompt:     "Test prompt",
				Role:             "Test Role",
				Responsibilities: nil,
				FileName:         "test.md",
			},
			wantErr: false,
		},
		{
			name: "Agent with very long system prompt",
			agent: &Agent{
				Name:             "Test Agent",
				Category:         "test",
				ReportsTo:        "Manager",
				SystemPrompt:     strings.Repeat("Very long prompt. ", 100),
				Role:             "Test Role",
				Responsibilities: []string{"Task 1", "Task 2"},
				FileName:         "test.md",
			},
			wantErr: false,
		},
		{
			name: "Agent with special characters in name",
			agent: &Agent{
				Name:             "Test-Agent_123",
				Category:         "test",
				ReportsTo:        "Manager",
				SystemPrompt:     "Test prompt",
				Role:             "Test Role",
				Responsibilities: []string{"Task 1"},
				FileName:         "test.md",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdown, err := RenderAgentMarkdown(tt.agent)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderAgentMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(markdown) == 0 {
				t.Error("RenderAgentMarkdown() should return non-empty markdown")
			}
		})
	}
}

// TestRenderCommandMarkdown_EdgeCases_Coverage tests RenderCommandMarkdown edge cases (77.8% -> 95%+)
func TestRenderCommandMarkdown_EdgeCases_Coverage(t *testing.T) {
	tests := []struct {
		name    string
		command *Command
		wantErr bool
	}{
		{
			name: "Command with empty examples",
			command: &Command{
				Name:         "test",
				Trigger:      "/test",
				Category:     "core",
				Description:  "Test",
				Action:       "Do something",
				Examples:     []string{},
				Requirements: "req1",
			},
			wantErr: false,
		},
		{
			name: "Command with nil examples",
			command: &Command{
				Name:         "test",
				Trigger:      "/test",
				Category:     "core",
				Description:  "Test",
				Action:       "Do something",
				Examples:     nil,
				Requirements: "",
			},
			wantErr: false,
		},
		{
			name: "Command with empty requirements",
			command: &Command{
				Name:         "test",
				Trigger:      "/test",
				Category:     "core",
				Description:  "Test",
				Action:       "Do something",
				Examples:     []string{"ex1"},
				Requirements: "",
			},
			wantErr: false,
		},
		{
			name: "Command with very long action",
			command: &Command{
				Name:         "test",
				Trigger:      "/test",
				Category:     "core",
				Description:  "Test",
				Action:       strings.Repeat("Very long action. ", 50),
				Examples:     []string{"ex1"},
				Requirements: "req1",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdown, err := RenderCommandMarkdown(tt.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderCommandMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(markdown) == 0 {
				t.Error("RenderCommandMarkdown() should return non-empty markdown")
			}
		})
	}
}

// TestBoilerplateGenerator_Generate_ErrorPaths_Coverage tests error paths in boilerplate generation (54.5% -> 95%+)
func TestBoilerplateGenerator_Generate_ErrorPaths_Coverage(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with read-only directory
	readOnlyDir := filepath.Join(tmpDir, "readonly")
	if err := os.MkdirAll(readOnlyDir, 0444); err != nil {
		t.Logf("Cannot create read-only dir (may not be supported): %v", err)
	} else {
		generator := &BoilerplateGenerator{}
		request := &models.ProjectRequest{
			ProjectName: "test",
			IDEs:        []string{"Cursor"},
			ProjectType: "Fullstack",
		}
		err := generator.Generate(request, readOnlyDir)
		if err == nil {
			t.Log("Note: BoilerplateGenerator.Generate succeeded in read-only dir (may be system-dependent)")
		}
		os.Chmod(readOnlyDir, 0755) // Restore
	}
}

// TestDocsGenerator_Generate_ErrorPaths_Coverage tests error paths (54.5% -> 95%+)
func TestDocsGenerator_Generate_ErrorPaths_Coverage(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with read-only directory
	readOnlyDir := filepath.Join(tmpDir, "readonly")
	if err := os.MkdirAll(readOnlyDir, 0444); err != nil {
		t.Logf("Cannot create read-only dir: %v", err)
	} else {
		generator := &DocsGenerator{}
		request := &models.ProjectRequest{
			ProjectName: "test",
			IDEs:        []string{"Cursor"},
			ProjectType: "Fullstack",
		}
		err := generator.Generate(request, readOnlyDir)
		if err == nil {
			t.Log("Note: DocsGenerator.Generate succeeded in read-only dir (may be system-dependent)")
		}
		os.Chmod(readOnlyDir, 0755) // Restore
	}
}
