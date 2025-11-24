package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestRulesGenerator_Name(t *testing.T) {
	generator := &RulesGenerator{}
	if got := generator.Name(); got != "Rules Library" {
		t.Errorf("RulesGenerator.Name() = %v, want %v", got, "Rules Library")
	}
}

func TestRulesGenerator_Generate(t *testing.T) {
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

	generator := &RulesGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("RulesGenerator.Generate() error = %v", err)
	}

	// Verify .cursor/rules/library directory was created
	rulesDir := filepath.Join(tmpDir, ".cursor", "rules", "library")
	if _, err := os.Stat(rulesDir); os.IsNotExist(err) {
		t.Error("RulesGenerator.Generate() should create .cursor/rules/library directory")
	}

	// Verify some expected files exist
	expectedFiles := []string{
		"01-core-workflow/README.md",
		"03-languages/go.md",
		"04-frameworks/nextjs.md",
		"08-testing/jest.md",
	}

	for _, expectedFile := range expectedFiles {
		filePath := filepath.Join(rulesDir, expectedFile)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("RulesGenerator.Generate() should create file %s", expectedFile)
		}
	}
}

func TestExtractRules(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Extract rules
	if err := ExtractRules(tmpDir); err != nil {
		t.Fatalf("ExtractRules() error = %v", err)
	}

	// Verify directory structure was maintained
	categories := []string{
		"01-core-workflow",
		"02-ai-agents",
		"03-languages",
		"04-frameworks",
		"05-ui-libraries",
		"06-cloud-infrastructure",
		"07-databases",
		"08-testing",
		"09-devops-ci-cd",
		"10-code-quality",
		"11-documentation",
		"12-security",
		"13-development-practices",
		"14-mcp-tools",
		"15-project-specific",
	}

	for _, category := range categories {
		categoryPath := filepath.Join(tmpDir, category)
		if _, err := os.Stat(categoryPath); os.IsNotExist(err) {
			t.Errorf("ExtractRules() should create category directory %s", category)
		}

		// Verify README.md exists in each category
		readmePath := filepath.Join(categoryPath, "README.md")
		if _, err := os.Stat(readmePath); os.IsNotExist(err) {
			t.Errorf("ExtractRules() should create README.md in %s", category)
		}
	}
}

func TestExtractRules_FileContent(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Extract rules
	if err := ExtractRules(tmpDir); err != nil {
		t.Fatalf("ExtractRules() error = %v", err)
	}

	// Verify file content for a known file
	goRulesPath := filepath.Join(tmpDir, "03-languages", "go.md")
	content, err := os.ReadFile(goRulesPath)
	if err != nil {
		t.Fatalf("Failed to read go.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Go Language Rules") {
		t.Error("Extracted go.md should contain 'Go Language Rules'")
	}
}

func TestExtractRules_AllFiles(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Extract rules
	if err := ExtractRules(tmpDir); err != nil {
		t.Fatalf("ExtractRules() error = %v", err)
	}

	// Count extracted files
	var fileCount int
	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileCount++
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Error walking extracted rules: %v", err)
	}

	// Should have at least the README files and some rules files
	if fileCount < 15 {
		t.Errorf("ExtractRules() should extract at least 15 files (one README per category), got %d", fileCount)
	}
}

func TestExtractRules_InvalidPath(t *testing.T) {
	// Try to extract to a non-existent parent directory
	invalidPath := filepath.Join("/nonexistent", "path", "test")
	err := ExtractRules(invalidPath)
	if err == nil {
		t.Error("ExtractRules() should return error for invalid path")
	}
}

func TestExtractRules_EmptyPath(t *testing.T) {
	err := ExtractRules("")
	if err == nil {
		t.Error("ExtractRules() should return error for empty path")
	}
}

func TestExtractRules_ExistingDirectory(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Pre-create some directories
	categoryDir := filepath.Join(tmpDir, "01-core-workflow")
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		t.Fatalf("Failed to create category directory: %v", err)
	}

	// Extract rules - should succeed even if directory already exists
	if err := ExtractRules(tmpDir); err != nil {
		t.Errorf("ExtractRules() error = %v, want nil (existing directory should be OK)", err)
	}
}

func TestGenerateRules(t *testing.T) {
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

	if err := GenerateRules(request, tmpDir); err != nil {
		t.Fatalf("GenerateRules() error = %v", err)
	}

	// Verify directory was created
	rulesDir := filepath.Join(tmpDir, ".cursor", "rules", "library")
	if _, err := os.Stat(rulesDir); os.IsNotExist(err) {
		t.Error("GenerateRules() should create .cursor/rules/library directory")
	}
}
