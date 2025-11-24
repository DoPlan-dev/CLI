package generator

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// TestEndToEnd_CompleteWizardFlow tests the complete wizard flow
// This simulates the full user journey from wizard to project generation
func TestEndToEnd_CompleteWizardFlow(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-e2e-*")
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

	// Simulate wizard flow by creating a ProjectRequest
	// (In real usage, this would come from the TUI wizard)
	request := &models.ProjectRequest{
		ProjectName: "e2e-test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	// Validate request (as wizard would)
	if err := request.Validate(); err != nil {
		t.Fatalf("Request validation failed: %v", err)
	}

	// Run full orchestration (this is what happens after wizard completes)
	startTime := time.Now()
	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() error = %v", err)
	}
	generationTime := time.Since(startTime)

	// Verify generation time < 5 seconds
	if generationTime > 5*time.Second {
		t.Errorf("Generation took %v, expected < 5 seconds", generationTime)
	}

	projectPath := filepath.Join(tmpDir, "e2e-test-project")

	// Comprehensive verification of all generated components
	verifyCompleteProjectStructure(t, projectPath)
}

// TestEndToEnd_AllFilesGenerated verifies all expected files are generated correctly
func TestEndToEnd_AllFilesGenerated(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-files-*")
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
		ProjectName: "files-test",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() error = %v", err)
	}

	projectPath := filepath.Join(tmpDir, "files-test")

	// Verify all expected files exist
	expectedFiles := []string{
		// Root documentation
		"README.md",
		"STANDUP.md",
		"CHANGELOG.md",
		// IDE configs
		".cursorrules",
		// Boilerplate
		"package.json",
		"tsconfig.json",
		"tailwind.config.ts",
		".eslintrc.json",
		// App structure
		"src/app/layout.tsx",
		"src/app/page.tsx",
		"src/app/globals.css",
		// Agents
		".cursor/agents/project_orchestrator.md",
		".cursor/agents/product_manager.md",
		".cursor/agents/engineering_lead.md",
		// Commands
		".cursor/commands/tell.md",
		".cursor/commands/build.md",
		".cursor/commands/write.md",
		// Rules
		".cursor/rules/README.md",
		".cursor/rules/library/01-core-workflow/README.md",
		".cursor/rules/library/03-languages/go.md",
		// Plan
		".plan/00_System/IDEA.md",
		".plan/00_System/PRD.md",
		".plan/00_System/ARCHITECTURE.md",
		".plan/TASKS.md",
		".plan/active_state.json",
		// GitHub workflows
		".github/workflows/ci.yml",
		".github/workflows/release.yml",
		".github/workflows/changelog.yml",
		".github/workflows/branch-protection.yml",
	}

	for _, file := range expectedFiles {
		filePath := filepath.Join(projectPath, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected file %s was not generated", file)
		}
	}
}

// TestEndToEnd_GenerationTime verifies generation completes in < 5 seconds
func TestEndToEnd_GenerationTime(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-timing-*")
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
		ProjectName: "timing-test",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	startTime := time.Now()
	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() error = %v", err)
	}
	duration := time.Since(startTime)

	if duration > 5*time.Second {
		t.Errorf("Generation took %v, expected < 5 seconds", duration)
	}

	t.Logf("Generation completed in %v (target: < 5 seconds)", duration)
}

// TestEndToEnd_CrossPlatformPaths verifies path handling works across platforms
func TestEndToEnd_CrossPlatformPaths(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-crossplatform-*")
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

	// Test with project name that could cause issues on different platforms
	testCases := []struct {
		name        string
		projectName string
	}{
		{"simple", "test-project"},
		{"with-underscore", "test_project"},
		{"with-numbers", "test123"},
		{"mixed-case", "TestProject"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := &models.ProjectRequest{
				ProjectName: tc.projectName,
				IDE:         "Cursor",
				ProjectType: "Fullstack",
			}

			if err := Orchestrate(request); err != nil {
				t.Errorf("Orchestrate() with project name %q error = %v", tc.projectName, err)
				return
			}

			projectPath := filepath.Join(tmpDir, tc.projectName)
			if _, err := os.Stat(projectPath); os.IsNotExist(err) {
				t.Errorf("Project directory %s was not created", projectPath)
			}

			// Clean up for next test
			os.RemoveAll(projectPath)
		})
	}
}

// TestEndToEnd_AllIDEs tests generation with all supported IDEs
func TestEndToEnd_AllIDEs(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-ides-*")
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

	supportedIDEs := models.GetSupportedIDEs()

	for _, ide := range supportedIDEs {
		t.Run(ide, func(t *testing.T) {
			// Sanitize IDE name for project name (remove spaces, special chars)
			projectName := "ide-test-" + sanitizeForProjectName(ide)
			request := &models.ProjectRequest{
				ProjectName: projectName,
				IDE:         ide,
				ProjectType: "Fullstack",
			}

			if err := Orchestrate(request); err != nil {
				t.Errorf("Orchestrate() with IDE %q error = %v", ide, err)
				return
			}

			projectPath := filepath.Join(tmpDir, projectName)

			// Verify IDE-specific configs were generated
			if ide == "Cursor" {
				cursorrulesPath := filepath.Join(projectPath, ".cursorrules")
				if _, err := os.Stat(cursorrulesPath); os.IsNotExist(err) {
					t.Errorf("Expected .cursorrules for Cursor IDE")
				}
			} else if ide == "Claude Code" {
				claudePath := filepath.Join(projectPath, "CLAUDE.md")
				if _, err := os.Stat(claudePath); os.IsNotExist(err) {
					t.Errorf("Expected CLAUDE.md for Claude Code IDE")
				}
			}

			// Clean up for next test
			os.RemoveAll(projectPath)
		})
	}
}

// sanitizeForProjectName converts an IDE name to a valid project name
func sanitizeForProjectName(ide string) string {
	// Replace spaces and special characters with hyphens
	result := ""
	for _, r := range ide {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			result += string(r)
		} else if r == ' ' || r == '-' {
			result += "-"
		}
	}
	return result
}

// verifyCompleteProjectStructure performs comprehensive verification
func verifyCompleteProjectStructure(t *testing.T, projectPath string) {
	// Verify directory structure
	expectedDirs := []string{
		".cursor/agents",
		".cursor/commands",
		".cursor/rules/library",
		".plan/00_System",
		".github/workflows",
		"src/app",
	}

	for _, dir := range expectedDirs {
		dirPath := filepath.Join(projectPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Errorf("Expected directory %s was not created", dir)
		}
	}

	// Verify agents (should have 18 agents)
	agentsDir := filepath.Join(projectPath, ".cursor", "agents")
	agentFiles, err := os.ReadDir(agentsDir)
	if err != nil {
		t.Errorf("Failed to read agents directory: %v", err)
	} else if len(agentFiles) < 18 {
		t.Errorf("Expected at least 18 agent files, found %d", len(agentFiles))
	}

	// Verify commands (should have 11 core + 8 squad commands = 19)
	commandsDir := filepath.Join(projectPath, ".cursor", "commands")
	commandFiles, err := os.ReadDir(commandsDir)
	if err != nil {
		t.Errorf("Failed to read commands directory: %v", err)
	} else if len(commandFiles) < 19 {
		t.Errorf("Expected at least 19 command files, found %d", len(commandFiles))
	}

	// Verify rules library has categories
	rulesDir := filepath.Join(projectPath, ".cursor", "rules", "library")
	ruleCategories, err := os.ReadDir(rulesDir)
	if err != nil {
		t.Errorf("Failed to read rules directory: %v", err)
	} else if len(ruleCategories) < 10 {
		t.Errorf("Expected at least 10 rule categories, found %d", len(ruleCategories))
	}

	// Verify all 4 GitHub workflows
	workflowsDir := filepath.Join(projectPath, ".github", "workflows")
	workflowFiles, err := os.ReadDir(workflowsDir)
	if err != nil {
		t.Errorf("Failed to read workflows directory: %v", err)
	} else if len(workflowFiles) != 4 {
		t.Errorf("Expected 4 workflow files, found %d", len(workflowFiles))
	}

	// Verify documentation files
	docs := []string{"README.md", "STANDUP.md", "CHANGELOG.md"}
	for _, doc := range docs {
		docPath := filepath.Join(projectPath, doc)
		if _, err := os.Stat(docPath); os.IsNotExist(err) {
			t.Errorf("Expected documentation file %s was not created", doc)
		}
	}

	// Verify boilerplate files
	boilerplate := []string{"package.json", "tsconfig.json", "tailwind.config.ts"}
	for _, bp := range boilerplate {
		bpPath := filepath.Join(projectPath, bp)
		if _, err := os.Stat(bpPath); os.IsNotExist(err) {
			t.Errorf("Expected boilerplate file %s was not created", bp)
		}
	}
}

// TestEndToEnd_PlatformInfo logs platform information for cross-platform testing
func TestEndToEnd_PlatformInfo(t *testing.T) {
	t.Logf("Platform: %s/%s", runtime.GOOS, runtime.GOARCH)
	t.Logf("This test verifies the code works on the current platform")
	t.Logf("For full cross-platform testing, run on:")
	t.Logf("  - macOS (Intel and Apple Silicon)")
	t.Logf("  - Linux (Ubuntu, Debian)")
	t.Logf("  - Windows")

	// Verify filepath operations work correctly
	tmpDir, err := os.MkdirTemp("", "doplan-platform-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Test path joining (should work on all platforms)
	testPath := filepath.Join(tmpDir, "test", "nested", "path")
	if err := os.MkdirAll(testPath, 0755); err != nil {
		t.Errorf("Failed to create nested path: %v", err)
	}

	// Verify path was created correctly
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Errorf("Path %s was not created correctly", testPath)
	}
}
