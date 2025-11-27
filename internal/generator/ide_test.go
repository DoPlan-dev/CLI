package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
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
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify .cursorrules was created
	cursorrulesPath := filepath.Join(tmpDir, ".cursorrules")
	if _, err := os.Stat(cursorrulesPath); os.IsNotExist(err) {
		t.Error("IDEGenerator.Generate() should create .cursorrules for Cursor")
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
		IDEs:        []string{"Claude Code"},
		ProjectType: "Fullstack",
	}

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify CLAUDE.md was created in docs/
	claudePath := filepath.Join(tmpDir, "docs", "CLAUDE.md")
	if _, err := os.Stat(claudePath); os.IsNotExist(err) {
		t.Error("IDEGenerator.Generate() should create docs/CLAUDE.md for Claude Code")
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
		IDEs:        []string{"Claude Code"},
		ProjectType: "Fullstack",
	}

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify CLAUDE.md content
	claudePath := filepath.Join(tmpDir, "docs", "CLAUDE.md")
	content, err := os.ReadFile(claudePath)
	if err != nil {
		t.Fatalf("Failed to read docs/CLAUDE.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Agent Hierarchy") {
		t.Error("CLAUDE.md should contain 'Agent Hierarchy'")
	}
	if !strings.Contains(contentStr, "Commands") {
		t.Error("CLAUDE.md should contain 'Commands'")
	}
	if !strings.Contains(contentStr, "Rules") {
		t.Error("CLAUDE.md should contain 'Rules'")
	}
}

func TestIDEGenerator_Generate_MultipleIDEs(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor", "Claude Code", "Windsurf", "Antigravity", "Cline", "OpenCode"},
		ProjectType: "Fullstack",
	}

	generator := &IDEGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("IDEGenerator.Generate() error = %v", err)
	}

	// Verify all IDE configs were created
	expectedFiles := map[string]string{
		"Cursor":      filepath.Join(tmpDir, ".cursorrules"),
		"Claude Code": filepath.Join(tmpDir, "docs", "CLAUDE.md"),
		"Windsurf":    filepath.Join(tmpDir, ".windsurfrules"),
		"Antigravity": filepath.Join(tmpDir, ".antigravity", "config.md"),
		"Cline":       filepath.Join(tmpDir, ".cline", "config.md"),
		"OpenCode":    filepath.Join(tmpDir, "opencode.json"),
	}

	for ide, path := range expectedFiles {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("IDEGenerator.Generate() should create config for %s at %s", ide, path)
		}
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
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	if err := GenerateIDEConfigs(request, tmpDir); err != nil {
		t.Fatalf("GenerateIDEConfigs() error = %v", err)
	}

	// Verify .cursorrules was created for Cursor
	cursorrulesPath := filepath.Join(tmpDir, ".cursorrules")
	if _, err := os.Stat(cursorrulesPath); os.IsNotExist(err) {
		t.Error("GenerateIDEConfigs() should create .cursorrules for Cursor")
	}
}
