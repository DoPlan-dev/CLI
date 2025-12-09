package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestDashboardGenerator_Name(t *testing.T) {
	generator := &DashboardGenerator{}
	if got := generator.Name(); got != "Dashboard" {
		t.Errorf("DashboardGenerator.Name() = %v, want %v", got, "Dashboard")
	}
}

func TestDashboardGenerator_Generate(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-dashboard-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-dashboard-project",
		IDE:         "Cursor",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &DashboardGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DashboardGenerator.Generate() error = %v", err)
	}

	// Verify dashboard directory was created
	dashboardDir := filepath.Join(tmpDir, "dashboard")
	if _, err := os.Stat(dashboardDir); os.IsNotExist(err) {
		t.Error("DashboardGenerator.Generate() should create dashboard directory")
	}

	// Verify data directory was created
	dataDir := filepath.Join(dashboardDir, "data")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		t.Error("DashboardGenerator.Generate() should create dashboard/data directory")
	}

	// Verify all HTML pages were created
	expectedPages := []string{
		"index.html",
		"plan.html",
		"meetings.html",
		"achievements.html",
		"settings.html",
	}

	for _, page := range expectedPages {
		pagePath := filepath.Join(dashboardDir, page)
		if _, err := os.Stat(pagePath); os.IsNotExist(err) {
			t.Errorf("DashboardGenerator.Generate() should create %s", page)
		}
	}

	// Verify project.json was created
	projectDataPath := filepath.Join(dataDir, "project.json")
	if _, err := os.Stat(projectDataPath); os.IsNotExist(err) {
		t.Error("DashboardGenerator.Generate() should create dashboard/data/project.json")
	}
}

func TestDashboardGenerator_Generate_IndexContent(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-dashboard-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "MyTestProject",
		IDE:         "Cursor",
		IDEs:        []string{"Cursor", "VSCode"},
		ProjectType: "Fullstack",
	}

	generator := &DashboardGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DashboardGenerator.Generate() error = %v", err)
	}

	// Verify index.html content
	indexPath := filepath.Join(tmpDir, "dashboard", "index.html")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("Failed to read index.html: %v", err)
	}

	contentStr := string(content)

	// Check for project name
	if !strings.Contains(contentStr, request.ProjectName) {
		t.Errorf("index.html should contain project name %s", request.ProjectName)
	}

	// Check for project type
	if !strings.Contains(contentStr, request.ProjectType) {
		t.Errorf("index.html should contain project type %s", request.ProjectType)
	}

	// Check for DoPlan branding
	if !strings.Contains(contentStr, "DoPlan") {
		t.Error("index.html should contain 'DoPlan' branding")
	}

	// Check for required JavaScript functionality
	if !strings.Contains(contentStr, "loadProgress") {
		t.Error("index.html should contain loadProgress function")
	}

	// Check for Bootstrap and Tabler CSS
	if !strings.Contains(contentStr, "bootstrap") {
		t.Error("index.html should include Bootstrap CSS")
	}
	if !strings.Contains(contentStr, "tabler") {
		t.Error("index.html should include Tabler CSS")
	}
}

func TestDashboardGenerator_Generate_AllPagesContent(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-dashboard-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "TestProject",
		IDE:         "Cursor",
		IDEs:        []string{"Cursor"},
		ProjectType: "Backend",
	}

	generator := &DashboardGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DashboardGenerator.Generate() error = %v", err)
	}

	// Test each page has required content
	pages := map[string][]string{
		"plan.html":         {"Project Plan", "TASKS.md", "marked.parse"},
		"meetings.html":     {"Meetings & Rituals", "loadMeetings"},
		"achievements.html": {"Achievements", "loadAchievements", "canvas-confetti"},
		"settings.html":     {"Settings", "loadAgents", "loadMemoryInsights"},
	}

	for page, requiredStrings := range pages {
		pagePath := filepath.Join(tmpDir, "dashboard", page)
		content, err := os.ReadFile(pagePath)
		if err != nil {
			t.Fatalf("Failed to read %s: %v", page, err)
		}

		contentStr := string(content)
		for _, required := range requiredStrings {
			if !strings.Contains(contentStr, required) {
				t.Errorf("%s should contain '%s'", page, required)
			}
		}

		// All pages should have navigation sidebar
		if !strings.Contains(contentStr, "DoPlan") {
			t.Errorf("%s should contain navigation sidebar with DoPlan branding", page)
		}
	}
}

func TestDashboardGenerator_Generate_ProjectData(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-dashboard-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "DataTestProject",
		IDE:         "Cursor",
		IDEs:        []string{"Cursor", "VSCode", "IntelliJ"},
		ProjectType: "Frontend",
	}

	generator := &DashboardGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("DashboardGenerator.Generate() error = %v", err)
	}

	// Verify project.json content
	projectDataPath := filepath.Join(tmpDir, "dashboard", "data", "project.json")
	content, err := os.ReadFile(projectDataPath)
	if err != nil {
		t.Fatalf("Failed to read project.json: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, request.ProjectName) {
		t.Errorf("project.json should contain project name %s", request.ProjectName)
	}
	if !strings.Contains(contentStr, request.ProjectType) {
		t.Errorf("project.json should contain project type %s", request.ProjectType)
	}
}

func TestDashboardGenerator_Integration(t *testing.T) {
	// Test that dashboard generator works as part of full orchestration
	tmpDir, err := os.MkdirTemp("", "doplan-integration-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "integration-test",
		IDE:         "Cursor",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	// Run full orchestration
	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() error = %v", err)
	}

	projectPath := filepath.Join(tmpDir, "integration-test")

	// Verify dashboard was generated
	dashboardDir := filepath.Join(projectPath, "dashboard")
	if _, err := os.Stat(dashboardDir); os.IsNotExist(err) {
		t.Error("Dashboard should be generated as part of full orchestration")
	}

	// Verify index.html exists
	indexPath := filepath.Join(dashboardDir, "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Error("index.html should be generated as part of full orchestration")
	}
}
