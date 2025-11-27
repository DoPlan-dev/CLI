package generator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestPlanGenerator_Name(t *testing.T) {
	generator := &PlanGenerator{}
	if got := generator.Name(); got != ".do Structure" {
		t.Errorf("PlanGenerator.Name() = %v, want %v", got, ".do Structure")
	}
}

func TestPlanGenerator_Generate(t *testing.T) {
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

	generator := &PlanGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("PlanGenerator.Generate() error = %v", err)
	}

	// Verify .do/system directory was created
	systemDir := filepath.Join(tmpDir, ".do", "system")
	if _, err := os.Stat(systemDir); os.IsNotExist(err) {
		t.Error("PlanGenerator.Generate() should create .do/system directory")
	}

	// Verify all required files exist
	requiredFiles := []string{
		"IDEA.md",
		"BRAINSTORM.md",
		"PRD.md",
		"ARCHITECTURE.md",
		"DESIGN_SYSTEM.md",
	}

	for _, file := range requiredFiles {
		filePath := filepath.Join(systemDir, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("PlanGenerator.Generate() should create %s", file)
		}
	}

	// Verify TASKS.md exists
	tasksPath := filepath.Join(tmpDir, ".do", "plan", "TASKS.md")
	if _, err := os.Stat(tasksPath); os.IsNotExist(err) {
		t.Error("PlanGenerator.Generate() should create TASKS.md")
	}

	// Verify active_state.json exists and is valid
	statePath := filepath.Join(tmpDir, ".do", "system", "history", "active_state.json")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		t.Error("PlanGenerator.Generate() should create active_state.json")
	}

	// Verify active_state.json is valid JSON
	content, err := os.ReadFile(statePath)
	if err != nil {
		t.Fatalf("Failed to read active_state.json: %v", err)
	}

	var state map[string]interface{}
	if err := json.Unmarshal(content, &state); err != nil {
		t.Errorf("active_state.json is not valid JSON: %v", err)
	}

	// Verify required fields
	if state["phase"] == nil {
		t.Error("active_state.json should have 'phase' field")
	}
	if state["locked"] == nil {
		t.Error("active_state.json should have 'locked' field")
	}
}

func TestPlanGenerator_Generate_FileContent(t *testing.T) {
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

	generator := &PlanGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("PlanGenerator.Generate() error = %v", err)
	}

	// Verify PRD.md contains project name
	prdPath := filepath.Join(tmpDir, ".do", "system", "PRD.md")
	content, err := os.ReadFile(prdPath)
	if err != nil {
		t.Fatalf("Failed to read PRD.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, request.ProjectName) {
		t.Errorf("PRD.md should contain project name %s", request.ProjectName)
	}

	// Verify IDEA.md has expected structure
	ideaPath := filepath.Join(tmpDir, ".do", "system", "IDEA.md")
	ideaContent, err := os.ReadFile(ideaPath)
	if err != nil {
		t.Fatalf("Failed to read IDEA.md: %v", err)
	}

	ideaStr := string(ideaContent)
	if !strings.Contains(ideaStr, "# Project Idea") {
		t.Error("IDEA.md should contain '# Project Idea' header")
	}
}

func TestGeneratePlan(t *testing.T) {
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

	if err := GeneratePlan(request, tmpDir); err != nil {
		t.Fatalf("GeneratePlan() error = %v", err)
	}

	// Verify directory was created
	systemDir := filepath.Join(tmpDir, ".do", "system")
	if _, err := os.Stat(systemDir); os.IsNotExist(err) {
		t.Error("GeneratePlan() should create .do/system directory")
	}
}
