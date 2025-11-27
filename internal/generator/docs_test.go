package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestDocsGenerator_Name(t *testing.T) {
	generator := &DocsGenerator{}
	if got := generator.Name(); got != "Documentation" {
		t.Errorf("DocsGenerator.Name() = %v, want %v", got, "Documentation")
	}
}

func TestDocsGenerator_Generate(t *testing.T) {
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

	generator := &DocsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DocsGenerator.Generate() error = %v", err)
	}

	// Verify README.md was created
	readmePath := filepath.Join(tmpDir, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		t.Error("DocsGenerator.Generate() should create README.md")
	}
}

func TestDocsGenerator_Generate_FileContent(t *testing.T) {
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

	generator := &DocsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DocsGenerator.Generate() error = %v", err)
	}

	// Verify README.md content
	readmePath := filepath.Join(tmpDir, "README.md")
	content, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatalf("Failed to read README.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, request.ProjectName) {
		t.Errorf("README.md should contain project name %s", request.ProjectName)
	}
	if !strings.Contains(contentStr, "# "+request.ProjectName) {
		t.Error("README.md should start with project name as heading")
	}
	if !strings.Contains(contentStr, "/tell") {
		t.Error("README.md should contain command reference")
	}
	if !strings.Contains(contentStr, "Agents") {
		t.Error("README.md should contain agent hierarchy section")
	}
	if !strings.Contains(contentStr, "Workflow") {
		t.Error("README.md should contain workflow section")
	}
}

func TestGenerateDocumentation(t *testing.T) {
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

	if err := GenerateDocumentation(request, tmpDir); err != nil {
		t.Fatalf("GenerateDocumentation() error = %v", err)
	}

	// Verify README.md was created
	readmePath := filepath.Join(tmpDir, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		t.Error("GenerateDocumentation() should create README.md")
	}

	// Verify STANDUP.md was created in .do/plan/
	standupPath := filepath.Join(tmpDir, ".do", "plan", "STANDUP.md")
		if _, err := os.Stat(standupPath); os.IsNotExist(err) {
		t.Error("GenerateDocumentation() should create .do/plan/STANDUP.md")
	}

	// Verify CHANGELOG.md was created in docs/
	changelogPath := filepath.Join(tmpDir, "docs", "CHANGELOG.md")
	if _, err := os.Stat(changelogPath); os.IsNotExist(err) {
		t.Error("GenerateDocumentation() should create docs/CHANGELOG.md")
	}

	// Verify rules README was created
	rulesReadmePath := filepath.Join(tmpDir, ".cursor", "rules", "README.md")
	if _, err := os.Stat(rulesReadmePath); os.IsNotExist(err) {
		t.Error("GenerateDocumentation() should create .cursor/rules/README.md")
	}
}

func TestGenerateSTANDUP(t *testing.T) {
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

	if err := generateSTANDUP(tmpDir, request); err != nil {
		t.Fatalf("generateSTANDUP() error = %v", err)
	}

	standupPath := filepath.Join(tmpDir, ".plan", "STANDUP.md")
	content, err := os.ReadFile(standupPath)
	if err != nil {
		t.Fatalf("Failed to read .plan/STANDUP.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Daily Standup") {
		t.Error("STANDUP.md should contain 'Daily Standup'")
	}
	if !strings.Contains(contentStr, "Yesterday") {
		t.Error("STANDUP.md should contain 'Yesterday' section")
	}
	if !strings.Contains(contentStr, "Today") {
		t.Error("STANDUP.md should contain 'Today' section")
	}
}

func TestGenerateCHANGELOG(t *testing.T) {
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

	if err := generateCHANGELOG(tmpDir, request); err != nil {
		t.Fatalf("generateCHANGELOG() error = %v", err)
	}

	changelogPath := filepath.Join(tmpDir, "docs", "CHANGELOG.md")
	content, err := os.ReadFile(changelogPath)
	if err != nil {
		t.Fatalf("Failed to read docs/CHANGELOG.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Changelog") {
		t.Error("CHANGELOG.md should contain 'Changelog'")
	}
	if !strings.Contains(contentStr, "[Unreleased]") {
		t.Error("CHANGELOG.md should contain '[Unreleased]' section")
	}
	if !strings.Contains(contentStr, "Added") {
		t.Error("CHANGELOG.md should contain 'Added' category")
	}
}

func TestGenerateRulesREADME(t *testing.T) {
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

	if err := generateRulesREADME(tmpDir, request); err != nil {
		t.Fatalf("generateRulesREADME() error = %v", err)
	}

	rulesReadmePath := filepath.Join(tmpDir, ".cursor", "rules", "README.md")
	content, err := os.ReadFile(rulesReadmePath)
	if err != nil {
		t.Fatalf("Failed to read rules README.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Project Rules") {
		t.Error("Rules README should contain 'Project Rules'")
	}
	if !strings.Contains(contentStr, "01-core-workflow") {
		t.Error("Rules README should contain category '01-core-workflow'")
	}
	if !strings.Contains(contentStr, "Referencing Rules") {
		t.Error("Rules README should contain 'Referencing Rules' section")
	}
}
