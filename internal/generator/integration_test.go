package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// TestFullProjectGeneration tests the complete project generation flow
func TestFullProjectGeneration(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-integration-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	// Run full orchestration
	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() error = %v", err)
	}

	projectPath := filepath.Join(tmpDir, "test-project")

	// Verify all major components were generated
	verifyAgentsGenerated(t, projectPath)
	verifyCommandsGenerated(t, projectPath)
	verifyRulesGenerated(t, projectPath)
	verifyPlanGenerated(t, projectPath)
	verifyGitHubWorkflowsGenerated(t, projectPath)
	verifyDocumentationGenerated(t, projectPath)
	verifyIDEConfigsGenerated(t, projectPath)
}

func verifyAgentsGenerated(t *testing.T, projectPath string) {
	// Check central location (where files are actually created)
	centralAgentsDir := filepath.Join(projectPath, ".do", "core", "agents")
	if _, err := os.Stat(centralAgentsDir); os.IsNotExist(err) {
		t.Error("Central agents directory should be created")
		return
	}

	// Check for some expected agent files (using actual file names from agents.go)
	// Files are now in category folders
	expectedAgents := map[string]string{
		"project_orchestrator.md": "leadership",
		"product_manager.md":      "product",
		"engineering_lead.md":     "engineering",
	}

	for agentFile, category := range expectedAgents {
		agentPath := filepath.Join(centralAgentsDir, category, agentFile)
		if _, err := os.Stat(agentPath); os.IsNotExist(err) {
			t.Errorf("Agent file %s should be generated in category %s", agentFile, category)
		}
	}
}

func verifyCommandsGenerated(t *testing.T, projectPath string) {
	// Check central location (where files are actually created)
	centralCommandsDir := filepath.Join(projectPath, ".do", "core", "commands")
	if _, err := os.Stat(centralCommandsDir); os.IsNotExist(err) {
		t.Error("Central commands directory should be created")
		return
	}

	// Check for some expected command files (now in category folders)
	expectedCommands := map[string]string{
		"tell.md":  "core",
		"build.md": "core",
		"write.md": "core",
	}

	for cmdFile, category := range expectedCommands {
		cmdPath := filepath.Join(centralCommandsDir, category, cmdFile)
		if _, err := os.Stat(cmdPath); os.IsNotExist(err) {
			t.Errorf("Command file %s should be generated in category %s", cmdFile, category)
		}
	}
}

func verifyRulesGenerated(t *testing.T, projectPath string) {
	// Check central location (where rules are actually extracted)
	centralRulesDir := filepath.Join(projectPath, ".do", "core", "library")
	if _, err := os.Stat(centralRulesDir); os.IsNotExist(err) {
		t.Error("Central rules library directory should be created")
		return
	}

	// Check for category directories in central location
	expectedCategories := []string{
		"01-core-workflow",
		"03-languages",
		"04-frameworks",
	}

	for _, category := range expectedCategories {
		categoryPath := filepath.Join(centralRulesDir, category)
		if _, err := os.Stat(categoryPath); os.IsNotExist(err) {
			t.Errorf("Rules category %s should be generated in central location", category)
		}
	}

	// Also verify IDE location is accessible (symlink or copy)
	ideRulesDir := filepath.Join(projectPath, ".cursor", "rules", "library")
	if _, err := os.Stat(ideRulesDir); os.IsNotExist(err) {
		t.Error("IDE rules library directory should be created (symlink or copy)")
	}
}

func verifyPlanGenerated(t *testing.T, projectPath string) {
	planDir := filepath.Join(projectPath, ".do")
	if _, err := os.Stat(planDir); os.IsNotExist(err) {
		t.Error(".do directory should be created")
		return
	}

	// Check for system directory
	systemDir := filepath.Join(planDir, "system")
	if _, err := os.Stat(systemDir); os.IsNotExist(err) {
		t.Error(".do/system directory should be created")
	}

	// Check for required files
	requiredFiles := []string{
		"IDEA.md",
		"PRD.md",
		"ARCHITECTURE.md",
		"TASKS.md",
		"active_state.json",
	}

	for _, file := range requiredFiles {
		var filePath string
		if file == "TASKS.md" {
			filePath = filepath.Join(planDir, "plan", file)
		} else if file == "active_state.json" {
			filePath = filepath.Join(planDir, "system", "history", file)
		} else {
			filePath = filepath.Join(systemDir, file)
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Plan file %s should be generated", file)
		}
	}
}

func verifyGitHubWorkflowsGenerated(t *testing.T, projectPath string) {
	workflowsDir := filepath.Join(projectPath, ".github", "workflows")
	if _, err := os.Stat(workflowsDir); os.IsNotExist(err) {
		t.Error(".github/workflows directory should be created")
		return
	}

	// Check for all workflow files
	expectedWorkflows := []string{
		"ci.yml",
		"release.yml",
		"changelog.yml",
		"branch-protection.yml",
	}

	for _, workflow := range expectedWorkflows {
		workflowPath := filepath.Join(workflowsDir, workflow)
		if _, err := os.Stat(workflowPath); os.IsNotExist(err) {
			t.Errorf("Workflow file %s should be generated", workflow)
		}
	}
}

// TestGeneratorPipelineOrder tests that generators run in the correct order
func TestGeneratorPipelineOrder(t *testing.T) {
	// This test verifies that the pipeline order is correct
	// Agents, Commands, Rules, Plan, GitHub should all run successfully
	// We can't easily test order without mocking, but we can verify all run

	tmpDir, err := os.MkdirTemp("", "doplan-order-*")
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
		ProjectName: "order-test",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	// All generators should run without errors
	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() should run all generators successfully: %v", err)
	}

	// Verify all generators produced output
	projectPath := filepath.Join(tmpDir, "order-test")

	checks := []struct {
		name string
		path string
	}{
		{"Agents", filepath.Join(projectPath, ".cursor", "agents")},
		{"Commands", filepath.Join(projectPath, ".cursor", "commands")},
		{"Rules", filepath.Join(projectPath, ".cursor", "rules", "library")},
		{"Plan", filepath.Join(projectPath, ".do")},
		{"GitHub", filepath.Join(projectPath, ".github", "workflows")},
	}

	for _, check := range checks {
		if _, err := os.Stat(check.path); os.IsNotExist(err) {
			t.Errorf("%s generator should create %s", check.name, check.path)
		}
	}
}

func verifyDocumentationGenerated(t *testing.T, projectPath string) {
	// Check for docs/overview/README.md
	readmePath := filepath.Join(projectPath, "docs", "overview", "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		t.Error("Documentation generator should create docs/overview/README.md")
	}

	// Check for STANDUP.md in .do/plan/
	standupPath := filepath.Join(projectPath, ".do", "plan", "STANDUP.md")
	if _, err := os.Stat(standupPath); os.IsNotExist(err) {
		t.Error("Documentation generator should create .do/plan/STANDUP.md")
	}

	// Check for CHANGELOG.md in docs/history/
	changelogPath := filepath.Join(projectPath, "docs", "history", "CHANGELOG.md")
	if _, err := os.Stat(changelogPath); os.IsNotExist(err) {
		t.Error("Documentation generator should create docs/history/CHANGELOG.md")
	}

	// Check for rules README
	rulesReadmePath := filepath.Join(projectPath, ".cursor", "rules", "README.md")
	if _, err := os.Stat(rulesReadmePath); os.IsNotExist(err) {
		t.Error("Documentation generator should create .cursor/rules/README.md")
	}
}

func verifyIDEConfigsGenerated(t *testing.T, projectPath string) {
	// Check for .cursorrules (default IDE config for Cursor)
	cursorrulesPath := filepath.Join(projectPath, ".cursorrules")
	if _, err := os.Stat(cursorrulesPath); os.IsNotExist(err) {
		t.Error("IDE generator should create .cursorrules for Cursor")
	}
}
