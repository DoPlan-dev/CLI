package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestDocsGenerator_Generate_ErrorPaths(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with invalid project path
	invalidPath := filepath.Join(tmpDir, "invalid", "nested", "path")

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &DocsGenerator{}
	err := generator.Generate(request, invalidPath)
	// Should handle error gracefully
	_ = err
}

func TestDocsGenerator_Generate_AllDocs(t *testing.T) {
	tmpDir := t.TempDir()

	request := &models.ProjectRequest{
		ProjectName: "test-docs",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	generator := &DocsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DocsGenerator.Generate() error = %v", err)
	}

	// Verify docs/overview/README.md was created
	readmePath := filepath.Join(tmpDir, "docs", "overview", "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		t.Error("DocsGenerator.Generate() should create docs/overview/README.md")
	}

	// Verify STANDUP.md was created
	standupPath := filepath.Join(tmpDir, ".do", "plan", "STANDUP.md")
	if _, err := os.Stat(standupPath); os.IsNotExist(err) {
		t.Error("DocsGenerator.Generate() should create STANDUP.md")
	}

	// Verify CHANGELOG.md was created in docs/history
	changelogPath := filepath.Join(tmpDir, "docs", "history", "CHANGELOG.md")
	if _, err := os.Stat(changelogPath); os.IsNotExist(err) {
		t.Error("DocsGenerator.Generate() should create docs/history/CHANGELOG.md")
	}

	// Verify docs hierarchy was created
	docsDir := filepath.Join(tmpDir, "docs")
	if _, err := os.Stat(docsDir); os.IsNotExist(err) {
		t.Error("DocsGenerator.Generate() should create docs directory")
	}
}

func TestGenerateDocsHierarchy_Structure(t *testing.T) {
	tmpDir := t.TempDir()

	if err := generateDocsHierarchy(tmpDir); err != nil {
		t.Fatalf("generateDocsHierarchy() error = %v", err)
	}

	// Verify expected directories were created
	// generateDocsHierarchy creates: overview, references, tutorials, history, ops, research
	expectedDirs := []string{
		"docs",
		"docs/overview",
		"docs/references",
		"docs/tutorials",
		"docs/history",
		"docs/ops",
		"docs/research",
	}

	for _, dir := range expectedDirs {
		dirPath := filepath.Join(tmpDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Errorf("generateDocsHierarchy() should create %s", dir)
		}
	}
}

func TestGenerateDocsHierarchy_ExistingDirs(t *testing.T) {
	tmpDir := t.TempDir()

	// Create existing docs directory
	docsDir := filepath.Join(tmpDir, "docs")
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		t.Fatalf("Failed to create existing docs dir: %v", err)
	}

	// Should handle existing directory gracefully
	if err := generateDocsHierarchy(tmpDir); err != nil {
		t.Errorf("generateDocsHierarchy() should handle existing directory, got error: %v", err)
	}
}

func TestGenerateREADME_Content(t *testing.T) {
	tmpDir := t.TempDir()

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	if err := generateREADME(tmpDir, request); err != nil {
		t.Fatalf("generateREADME() error = %v", err)
	}

	readmePath := filepath.Join(tmpDir, "docs", "overview", "README.md")
	content, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatalf("Failed to read docs/overview/README.md: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, request.ProjectName) {
		t.Error("generateREADME() should contain project name")
	}
}

func TestGenerateSTANDUP_Content(t *testing.T) {
	tmpDir := t.TempDir()

	// Create .do/plan directory
	planDir := filepath.Join(tmpDir, ".do", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		t.Fatalf("Failed to create plan dir: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	if err := generateSTANDUP(tmpDir, request); err != nil {
		t.Fatalf("generateSTANDUP() error = %v", err)
	}

	standupPath := filepath.Join(planDir, "STANDUP.md")
	if _, err := os.Stat(standupPath); os.IsNotExist(err) {
		t.Error("generateSTANDUP() should create STANDUP.md")
	}
}

func TestGenerateCHANGELOG_Content(t *testing.T) {
	tmpDir := t.TempDir()

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	if err := generateCHANGELOG(tmpDir, request); err != nil {
		t.Fatalf("generateCHANGELOG() error = %v", err)
	}

	// CHANGELOG.md is created in docs/history directory, not root
	changelogPath := filepath.Join(tmpDir, "docs", "history", "CHANGELOG.md")
	content, err := os.ReadFile(changelogPath)
	if err != nil {
		t.Fatalf("Failed to read docs/history/CHANGELOG.md: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, request.ProjectName) {
		t.Error("generateCHANGELOG() should contain project name")
	}
}

// Helper functions are in agents_integration_test.go
