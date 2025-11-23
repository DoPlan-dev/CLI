package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/doplan/cli/pkg/models"
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
	verifyBoilerplateGenerated(t, projectPath)
	verifyIDEConfigsGenerated(t, projectPath)
}

func verifyAgentsGenerated(t *testing.T, projectPath string) {
	agentsDir := filepath.Join(projectPath, ".cursor", "agents")
	if _, err := os.Stat(agentsDir); os.IsNotExist(err) {
		t.Error("Agents directory should be created")
		return
	}

	// Check for some expected agent files (using actual file names from agents.go)
	expectedAgents := []string{
		"project_orchestrator.md",
		"product_manager.md",
		"engineering_lead.md",
	}

	for _, agent := range expectedAgents {
		agentPath := filepath.Join(agentsDir, agent)
		if _, err := os.Stat(agentPath); os.IsNotExist(err) {
			t.Errorf("Agent file %s should be generated", agent)
		}
	}
}

func verifyCommandsGenerated(t *testing.T, projectPath string) {
	commandsDir := filepath.Join(projectPath, ".cursor", "commands")
	if _, err := os.Stat(commandsDir); os.IsNotExist(err) {
		t.Error("Commands directory should be created")
		return
	}

	// Check for some expected command files
	expectedCommands := []string{
		"tell.md",
		"build.md",
		"write.md",
	}

	for _, cmd := range expectedCommands {
		cmdPath := filepath.Join(commandsDir, cmd)
		if _, err := os.Stat(cmdPath); os.IsNotExist(err) {
			t.Errorf("Command file %s should be generated", cmd)
		}
	}
}

func verifyRulesGenerated(t *testing.T, projectPath string) {
	rulesDir := filepath.Join(projectPath, ".cursor", "rules", "library")
	if _, err := os.Stat(rulesDir); os.IsNotExist(err) {
		t.Error("Rules library directory should be created")
		return
	}

	// Check for category directories
	expectedCategories := []string{
		"01-core-workflow",
		"03-languages",
		"04-frameworks",
	}

	for _, category := range expectedCategories {
		categoryPath := filepath.Join(rulesDir, category)
		if _, err := os.Stat(categoryPath); os.IsNotExist(err) {
			t.Errorf("Rules category %s should be generated", category)
		}
	}
}

func verifyPlanGenerated(t *testing.T, projectPath string) {
	planDir := filepath.Join(projectPath, ".plan")
	if _, err := os.Stat(planDir); os.IsNotExist(err) {
		t.Error(".plan directory should be created")
		return
	}

	// Check for system directory
	systemDir := filepath.Join(planDir, "00_System")
	if _, err := os.Stat(systemDir); os.IsNotExist(err) {
		t.Error(".plan/00_System directory should be created")
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
		if file == "TASKS.md" || file == "active_state.json" {
			filePath = filepath.Join(planDir, file)
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
		{"Plan", filepath.Join(projectPath, ".plan")},
		{"GitHub", filepath.Join(projectPath, ".github", "workflows")},
	}

	for _, check := range checks {
		if _, err := os.Stat(check.path); os.IsNotExist(err) {
			t.Errorf("%s generator should create %s", check.name, check.path)
		}
	}
}

func verifyDocumentationGenerated(t *testing.T, projectPath string) {
	// Check for README.md
	readmePath := filepath.Join(projectPath, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		t.Error("Documentation generator should create README.md")
	}

	// Check for STANDUP.md
	standupPath := filepath.Join(projectPath, "STANDUP.md")
	if _, err := os.Stat(standupPath); os.IsNotExist(err) {
		t.Error("Documentation generator should create STANDUP.md")
	}

	// Check for CHANGELOG.md
	changelogPath := filepath.Join(projectPath, "CHANGELOG.md")
	if _, err := os.Stat(changelogPath); os.IsNotExist(err) {
		t.Error("Documentation generator should create CHANGELOG.md")
	}

	// Check for rules README
	rulesReadmePath := filepath.Join(projectPath, ".cursor", "rules", "README.md")
	if _, err := os.Stat(rulesReadmePath); os.IsNotExist(err) {
		t.Error("Documentation generator should create .cursor/rules/README.md")
	}
}

func verifyBoilerplateGenerated(t *testing.T, projectPath string) {
	// Check for package.json (Next.js boilerplate)
	packageJsonPath := filepath.Join(projectPath, "package.json")
	if _, err := os.Stat(packageJsonPath); os.IsNotExist(err) {
		t.Error("Boilerplate generator should create package.json")
	}

	// Check for tsconfig.json
	tsconfigPath := filepath.Join(projectPath, "tsconfig.json")
	if _, err := os.Stat(tsconfigPath); os.IsNotExist(err) {
		t.Error("Boilerplate generator should create tsconfig.json")
	}
}

func verifyIDEConfigsGenerated(t *testing.T, projectPath string) {
	// Check for .cursorrules (Cursor IDE)
	cursorrulesPath := filepath.Join(projectPath, ".cursorrules")
	if _, err := os.Stat(cursorrulesPath); os.IsNotExist(err) {
		t.Error("IDE generator should create .cursorrules")
	}
}

