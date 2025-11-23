package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/doplan/cli/pkg/models"
)

func TestIDEGenerator_Name(t *testing.T) {
	generator := &IDEGenerator{}
	if got := generator.Name(); got != "IDE Configs" {
		t.Errorf("IDEGenerator.Name() = %v, want %v", got, "IDE Configs")
	}
}

func TestIDEGenerator_Generate_Cursor(t *testing.T) {
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

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify .cursorrules was created
	cursorRulesPath := filepath.Join(tmpDir, ".cursorrules")
	if _, err := os.Stat(cursorRulesPath); os.IsNotExist(err) {
		t.Error("IDEGenerator.Generate() should create .cursorrules for Cursor")
	}

	// Verify CLAUDE.md was also created
	claudePath := filepath.Join(tmpDir, "CLAUDE.md")
	if _, err := os.Stat(claudePath); os.IsNotExist(err) {
		t.Error("IDEGenerator.Generate() should create CLAUDE.md")
	}
}

func TestIDEGenerator_Generate_ClaudeCode(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Claude Code",
		ProjectType: "Fullstack",
	}

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify CLAUDE.md was created
	claudePath := filepath.Join(tmpDir, "CLAUDE.md")
	if _, err := os.Stat(claudePath); os.IsNotExist(err) {
		t.Error("IDEGenerator.Generate() should create CLAUDE.md for Claude Code")
	}
}

func TestIDEGenerator_Generate_FileContent(t *testing.T) {
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

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify .cursorrules content
	cursorRulesPath := filepath.Join(tmpDir, ".cursorrules")
	content, err := os.ReadFile(cursorRulesPath)
	if err != nil {
		t.Fatalf("Failed to read .cursorrules: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Agent Hierarchy") {
		t.Error(".cursorrules should contain 'Agent Hierarchy'")
	}
	if !strings.Contains(contentStr, "Commands") {
		t.Error(".cursorrules should contain 'Commands'")
	}
	if !strings.Contains(contentStr, "Rules") {
		t.Error(".cursorrules should contain 'Rules'")
	}
}

func TestGenerateIDEConfigs(t *testing.T) {
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

	if err := GenerateIDEConfigs(request, tmpDir); err != nil {
		t.Fatalf("GenerateIDEConfigs() error = %v", err)
	}

	// Verify files were created
	cursorRulesPath := filepath.Join(tmpDir, ".cursorrules")
	if _, err := os.Stat(cursorRulesPath); os.IsNotExist(err) {
		t.Error("GenerateIDEConfigs() should create .cursorrules")
	}
}

