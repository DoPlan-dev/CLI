package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestGitHubGenerator_Name(t *testing.T) {
	generator := &GitHubGenerator{}
	if got := generator.Name(); got != "GitHub Workflows" {
		t.Errorf("GitHubGenerator.Name() = %v, want %v", got, "GitHub Workflows")
	}
}

func TestGitHubGenerator_Generate(t *testing.T) {
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

	generator := &GitHubGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("GitHubGenerator.Generate() error = %v", err)
	}

	// Verify .github/workflows directory was created
	workflowsDir := filepath.Join(tmpDir, ".github", "workflows")
	if _, err := os.Stat(workflowsDir); os.IsNotExist(err) {
		t.Error("GitHubGenerator.Generate() should create .github/workflows directory")
	}

	// Verify all workflow files exist
	requiredFiles := []string{
		"ci.yml",
		"release.yml",
		"changelog.yml",
		"branch-protection.yml",
	}

	for _, file := range requiredFiles {
		filePath := filepath.Join(workflowsDir, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("GitHubGenerator.Generate() should create %s", file)
		}
	}
}

func TestGitHubGenerator_Generate_FileContent(t *testing.T) {
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

	generator := &GitHubGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("GitHubGenerator.Generate() error = %v", err)
	}

	// Verify CI workflow content
	ciPath := filepath.Join(tmpDir, ".github", "workflows", "ci.yml")
	content, err := os.ReadFile(ciPath)
	if err != nil {
		t.Fatalf("Failed to read ci.yml: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "name: CI") {
		t.Error("ci.yml should contain 'name: CI'")
	}
	if !strings.Contains(contentStr, "test:") {
		t.Error("ci.yml should contain test job")
	}
	if !strings.Contains(contentStr, "lint:") {
		t.Error("ci.yml should contain lint job")
	}
	if !strings.Contains(contentStr, "build:") {
		t.Error("ci.yml should contain build job")
	}

	// Verify Release workflow content
	releasePath := filepath.Join(tmpDir, ".github", "workflows", "release.yml")
	releaseContent, err := os.ReadFile(releasePath)
	if err != nil {
		t.Fatalf("Failed to read release.yml: %v", err)
	}

	releaseStr := string(releaseContent)
	if !strings.Contains(releaseStr, "name: Release") {
		t.Error("release.yml should contain 'name: Release'")
	}
	if !strings.Contains(releaseStr, "tags:") {
		t.Error("release.yml should trigger on tags")
	}
}

func TestGenerateGitHubWorkflows(t *testing.T) {
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

	if err := GenerateGitHubWorkflows(request, tmpDir); err != nil {
		t.Fatalf("GenerateGitHubWorkflows() error = %v", err)
	}

	// Verify directory was created
	workflowsDir := filepath.Join(tmpDir, ".github", "workflows")
	if _, err := os.Stat(workflowsDir); os.IsNotExist(err) {
		t.Error("GenerateGitHubWorkflows() should create .github/workflows directory")
	}
}
