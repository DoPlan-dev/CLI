package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestCommandsGenerator_Name(t *testing.T) {
	generator := &CommandsGenerator{}
	if got := generator.Name(); got != "Commands" {
		t.Errorf("CommandsGenerator.Name() = %v, want %v", got, "Commands")
	}
}

func TestCommandsGenerator_Generate(t *testing.T) {
	// Create temporary directory
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

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	// Verify .cursor/commands directory was created
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if _, err := os.Stat(commandsDir); os.IsNotExist(err) {
		t.Error("CommandsGenerator.Generate() should create .cursor/commands directory")
	}

	// Verify all command files were created
	commands := GetAllCommands()
	for _, cmd := range commands {
		commandPath := filepath.Join(commandsDir, cmd.Name+".md")
		if _, err := os.Stat(commandPath); os.IsNotExist(err) {
			t.Errorf("CommandsGenerator.Generate() should create file %s", cmd.Name+".md")
		}
	}
}

func TestCommandsGenerator_Generate_FileContent(t *testing.T) {
	// Create temporary directory
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

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	// Verify file content for a specific command
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	tellPath := filepath.Join(commandsDir, "tell.md")

	content, err := os.ReadFile(tellPath)
	if err != nil {
		t.Fatalf("Failed to read tell.md: %v", err)
	}

	contentStr := string(content)

	// Check for required sections
	if !strings.Contains(contentStr, "# /tell") {
		t.Error("tell.md should contain command title")
	}
	if !strings.Contains(contentStr, "## Trigger") {
		t.Error("tell.md should contain Trigger section")
	}
	if !strings.Contains(contentStr, "## Action") {
		t.Error("tell.md should contain Action section")
	}
	if !strings.Contains(contentStr, "## Agent Involvement") {
		t.Error("tell.md should contain Agent Involvement section")
	}
}

func TestCommandsGenerator_Generate_AllCommandsContent(t *testing.T) {
	// Create temporary directory
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

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	commands := GetAllCommands()

	// Verify each command file has correct content
	for _, cmd := range commands {
		t.Run(cmd.Name, func(t *testing.T) {
			commandPath := filepath.Join(commandsDir, cmd.Name+".md")
			content, err := os.ReadFile(commandPath)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", cmd.Name+".md", err)
			}

			contentStr := string(content)

			// Check for command title
			expectedTitle := "# /" + cmd.Name
			if !strings.Contains(contentStr, expectedTitle) {
				t.Errorf("%s should contain title %s", cmd.Name+".md", expectedTitle)
			}

			// Check for required sections
			if !strings.Contains(contentStr, "## Trigger") {
				t.Errorf("%s should contain Trigger section", cmd.Name+".md")
			}
			if !strings.Contains(contentStr, "## Action") {
				t.Errorf("%s should contain Action section", cmd.Name+".md")
			}
			if !strings.Contains(contentStr, "## Agent Involvement") {
				t.Errorf("%s should contain Agent Involvement section", cmd.Name+".md")
			}

			// Check agent names are included
			for _, agent := range cmd.AgentInvolvement {
				if !strings.Contains(contentStr, agent) {
					t.Errorf("%s should contain agent name %s", cmd.Name+".md", agent)
				}
			}
		})
	}
}

func TestCommandsGenerator_Generate_InvalidPath(t *testing.T) {
	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}

	// Try to generate in a non-existent parent directory
	invalidPath := filepath.Join("/nonexistent", "path", "test")
	err := generator.Generate(request, invalidPath)
	if err == nil {
		t.Error("CommandsGenerator.Generate() should return error for invalid path")
	}
}

func TestCommandsGenerator_Generate_ExistingDirectory(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Pre-create .cursor/commands directory
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		t.Fatalf("Failed to create commands directory: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}
	// Should succeed even if directory already exists
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Errorf("CommandsGenerator.Generate() error = %v, want nil (existing directory should be OK)", err)
	}
}

func TestGenerateCommands(t *testing.T) {
	// Create temporary directory
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

	if err := GenerateCommands(request, tmpDir); err != nil {
		t.Fatalf("GenerateCommands() error = %v", err)
	}

	// Verify directory was created
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if _, err := os.Stat(commandsDir); os.IsNotExist(err) {
		t.Error("GenerateCommands() should create .cursor/commands directory")
	}
}
