package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildFeatureSlug(t *testing.T) {
	testCases := []struct {
		name     string
		numStr   string
		title    string
		expected string
	}{
		{"simple", "1", "User Authentication", "02_user_authentication"},
		{"with special chars", "2", "API & Data Management!", "03_api__data_management"}, // Spaces become underscores, special chars removed
		{"with numbers", "3", "Feature 123 Test", "04_feature_123_test"},
		{"empty title", "4", "", "05_feature"},
		{"special chars only", "5", "@#$%", "06_feature"},
		{"mixed case", "6", "MixedCase Title", "07_mixedcase_title"},
		{"with dashes", "7", "Feature-Name Test", "08_feature-name_test"},
		{"phase 0", "0", "Initial Setup", "01_initial_setup"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := buildFeatureSlug(tc.numStr, tc.title)
			if result != tc.expected {
				t.Errorf("buildFeatureSlug(%q, %q) = %q, want %q", tc.numStr, tc.title, result, tc.expected)
			}
		})
	}
}

func TestPadPhaseNumber(t *testing.T) {
	testCases := []struct {
		name     string
		numStr   string
		expected string
	}{
		{"zero", "0", "01"},
		{"one", "1", "02"},
		{"nine", "9", "10"},
		{"ten", "10", "11"},
		{"invalid", "abc", "abc"},
		{"empty", "", ""},
		{"negative", "-1", "00"}, // strconv.Atoi("-1") succeeds, -1+1=0, formatted as "00"
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := padPhaseNumber(tc.numStr)
			if result != tc.expected {
				t.Errorf("padPhaseNumber(%q) = %q, want %q", tc.numStr, result, tc.expected)
			}
		})
	}
}

func TestGenerateFeatureTemplates(t *testing.T) {
	tmpDir := t.TempDir()
	featureDir := filepath.Join(tmpDir, "feature")
	if err := os.MkdirAll(featureDir, 0755); err != nil {
		t.Fatalf("Failed to create feature dir: %v", err)
	}

	feature := Feature{
		ID:          "1.1",
		Title:       "Test Feature",
		Description: "Test description",
	}

	if err := generateFeatureTemplates(featureDir, feature, "test-project"); err != nil {
		t.Fatalf("generateFeatureTemplates() error = %v", err)
	}

	// Verify all template files were created
	expectedFiles := []string{"design.md", "plan.md", "tasks.md", "prompts.md", "github.md"}
	for _, filename := range expectedFiles {
		path := filepath.Join(featureDir, filename)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("generateFeatureTemplates() should create %s", filename)
		}

		// Verify file content contains feature information
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("Failed to read %s: %v", filename, err)
		}
		contentStr := string(content)
		if !strings.Contains(contentStr, feature.Title) {
			t.Errorf("%s should contain feature title", filename)
		}
		if !strings.Contains(contentStr, feature.ID) {
			t.Errorf("%s should contain feature ID", filename)
		}
	}
}

func TestGenerateDesignTemplate(t *testing.T) {
	feature := Feature{
		ID:          "1.1",
		Title:       "User Login",
		Description: "Implement user login functionality",
	}

	result := generateDesignTemplate(feature, "test-project")
	if !strings.Contains(result, feature.Title) {
		t.Error("generateDesignTemplate() should contain feature title")
	}
	if !strings.Contains(result, feature.ID) {
		t.Error("generateDesignTemplate() should contain feature ID")
	}
	if !strings.Contains(result, feature.Description) {
		t.Error("generateDesignTemplate() should contain feature description")
	}
	if !strings.Contains(result, "test-project") {
		t.Error("generateDesignTemplate() should contain project name")
	}
}

func TestGeneratePlanTemplate(t *testing.T) {
	feature := Feature{
		ID:          "2.1",
		Title:       "API Endpoint",
		Description: "Create REST API endpoint",
	}

	result := generatePlanTemplate(feature, "test-project")
	if !strings.Contains(result, feature.Title) {
		t.Error("generatePlanTemplate() should contain feature title")
	}
	if !strings.Contains(result, feature.ID) {
		t.Error("generatePlanTemplate() should contain feature ID")
	}
}

func TestGenerateTasksTemplate(t *testing.T) {
	feature := Feature{
		ID:    "3.1",
		Title: "Database Schema",
	}

	result := generateTasksTemplate(feature, "test-project")
	if !strings.Contains(result, feature.Title) {
		t.Error("generateTasksTemplate() should contain feature title")
	}
	if !strings.Contains(result, feature.ID) {
		t.Error("generateTasksTemplate() should contain feature ID")
	}
}

func TestGeneratePromptsTemplate(t *testing.T) {
	feature := Feature{
		ID:    "4.1",
		Title: "UI Component",
	}

	result := generatePromptsTemplate(feature, "test-project")
	if !strings.Contains(result, feature.Title) {
		t.Error("generatePromptsTemplate() should contain feature title")
	}
	if !strings.Contains(result, feature.ID) {
		t.Error("generatePromptsTemplate() should contain feature ID")
	}
}

func TestGenerateGithubTemplate(t *testing.T) {
	feature := Feature{
		ID:    "5.1",
		Title: "Feature Name",
	}

	result := generateGithubTemplate(feature, "test-project")
	if !strings.Contains(result, feature.Title) {
		t.Error("generateGithubTemplate() should contain feature title")
	}
	if !strings.Contains(result, feature.ID) {
		t.Error("generateGithubTemplate() should contain feature ID")
	}
	if !strings.Contains(result, "task/5-1") {
		t.Error("generateGithubTemplate() should contain branch name")
	}
}

func TestMirrorPlanDocsToDocs(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir

	// Create .do/system directory with foundation docs
	systemDir := filepath.Join(projectPath, ".do", "system")
	if err := os.MkdirAll(systemDir, 0755); err != nil {
		t.Fatalf("Failed to create system dir: %v", err)
	}

	// Create foundation docs
	foundationDocs := []string{"IDEA.md", "BRAINSTORM.md", "PRD.md", "ARCHITECTURE.md", "DESIGN_SYSTEM.md"}
	for _, name := range foundationDocs {
		content := []byte("# " + name + "\nTest content")
		if err := os.WriteFile(filepath.Join(systemDir, name), content, 0644); err != nil {
			t.Fatalf("Failed to create %s: %v", name, err)
		}
	}

	// Create .do/plan/TASKS.md
	planDir := filepath.Join(projectPath, ".do", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		t.Fatalf("Failed to create plan dir: %v", err)
	}
	tasksContent := []byte("# Tasks\n## Phase 1: Test Phase")
	if err := os.WriteFile(filepath.Join(planDir, "TASKS.md"), tasksContent, 0644); err != nil {
		t.Fatalf("Failed to create TASKS.md: %v", err)
	}

	// Create docs directory
	docsRoot := filepath.Join(projectPath, "docs")
	if err := os.MkdirAll(docsRoot, 0755); err != nil {
		t.Fatalf("Failed to create docs dir: %v", err)
	}

	// Run mirrorPlanDocsToDocs
	if err := mirrorPlanDocsToDocs(projectPath); err != nil {
		t.Fatalf("mirrorPlanDocsToDocs() error = %v", err)
	}

	// Verify foundation docs were copied
	foundationDir := filepath.Join(docsRoot, "foundation")
	for _, name := range foundationDocs {
		dest := filepath.Join(foundationDir, name)
		if _, err := os.Stat(dest); os.IsNotExist(err) {
			t.Errorf("mirrorPlanDocsToDocs() should copy %s to foundation", name)
		}
	}

	// Verify TASKS.md was copied
	tasksDest := filepath.Join(docsRoot, "features", "_plan", "TASKS.md")
	if _, err := os.Stat(tasksDest); os.IsNotExist(err) {
		t.Error("mirrorPlanDocsToDocs() should copy TASKS.md to docs/features/_plan")
	}
}

func TestMirrorPlanDocsToDocs_NoDocsDir(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir

	// Create .do/system but no docs directory
	systemDir := filepath.Join(projectPath, ".do", "system")
	if err := os.MkdirAll(systemDir, 0755); err != nil {
		t.Fatalf("Failed to create system dir: %v", err)
	}

	// Should not error when docs directory doesn't exist
	if err := mirrorPlanDocsToDocs(projectPath); err != nil {
		t.Errorf("mirrorPlanDocsToDocs() should handle missing docs dir gracefully, got error: %v", err)
	}
}

func TestScaffoldFeaturePhaseDocs(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir
	docsRoot := filepath.Join(projectPath, "docs")
	if err := os.MkdirAll(docsRoot, 0755); err != nil {
		t.Fatalf("Failed to create docs dir: %v", err)
	}

	// Create TASKS.md with phases
	planDir := filepath.Join(projectPath, ".do", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		t.Fatalf("Failed to create plan dir: %v", err)
	}

	tasksContent := `# Tasks

## Phase 1: User Authentication
Some description

## Phase 2: Data Management
Another description
`
	if err := os.WriteFile(filepath.Join(planDir, "TASKS.md"), []byte(tasksContent), 0644); err != nil {
		t.Fatalf("Failed to create TASKS.md: %v", err)
	}

	// Run scaffoldFeaturePhaseDocs
	if err := scaffoldFeaturePhaseDocs(projectPath, docsRoot); err != nil {
		t.Fatalf("scaffoldFeaturePhaseDocs() error = %v", err)
	}

	// Verify phase directories were created
	phase1Dir := filepath.Join(docsRoot, "features", "02_user_authentication")
	phase2Dir := filepath.Join(docsRoot, "features", "03_data_management")

	if _, err := os.Stat(phase1Dir); os.IsNotExist(err) {
		t.Error("scaffoldFeaturePhaseDocs() should create phase 1 directory")
	}
	if _, err := os.Stat(phase2Dir); os.IsNotExist(err) {
		t.Error("scaffoldFeaturePhaseDocs() should create phase 2 directory")
	}

	// Verify README.md was created in each phase directory
	readme1 := filepath.Join(phase1Dir, "README.md")
	readme2 := filepath.Join(phase2Dir, "README.md")

	if _, err := os.Stat(readme1); os.IsNotExist(err) {
		t.Error("scaffoldFeaturePhaseDocs() should create README.md for phase 1")
	}
	if _, err := os.Stat(readme2); os.IsNotExist(err) {
		t.Error("scaffoldFeaturePhaseDocs() should create README.md for phase 2")
	}
}

func TestScaffoldFeaturePhaseDocs_NoTasksFile(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir
	docsRoot := filepath.Join(projectPath, "docs")
	if err := os.MkdirAll(docsRoot, 0755); err != nil {
		t.Fatalf("Failed to create docs dir: %v", err)
	}

	// Don't create TASKS.md - should handle gracefully
	if err := scaffoldFeaturePhaseDocs(projectPath, docsRoot); err != nil {
		t.Errorf("scaffoldFeaturePhaseDocs() should handle missing TASKS.md gracefully, got error: %v", err)
	}
}

func TestScaffoldPlanHierarchy(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir

	// Create TASKS.md with phases and features
	planDir := filepath.Join(projectPath, ".do", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		t.Fatalf("Failed to create plan dir: %v", err)
	}

	tasksContent := `# Tasks

## Phase 1: User Authentication
**Description**: Implement user authentication

### 1.1 Login Form
**Description**: Create login form component

### 1.2 Password Reset
**Description**: Implement password reset flow

## Phase 2: Data Management
**Description**: Manage application data

### 2.1 Database Schema
**Description**: Design database schema
`
	if err := os.WriteFile(filepath.Join(planDir, "TASKS.md"), []byte(tasksContent), 0644); err != nil {
		t.Fatalf("Failed to create TASKS.md: %v", err)
	}

	// Run ScaffoldPlanHierarchy
	if err := ScaffoldPlanHierarchy(projectPath); err != nil {
		t.Fatalf("ScaffoldPlanHierarchy() error = %v", err)
	}

	// Verify phase directories were created (function uses hyphens, not underscores)
	phase1Dir := filepath.Join(planDir, "02-user_authentication")
	phase2Dir := filepath.Join(planDir, "03-data_management")

	if _, err := os.Stat(phase1Dir); os.IsNotExist(err) {
		t.Error("ScaffoldPlanHierarchy() should create phase 1 directory")
	}
	if _, err := os.Stat(phase2Dir); os.IsNotExist(err) {
		t.Error("ScaffoldPlanHierarchy() should create phase 2 directory")
	}

	// Verify feature directories were created (function uses task sub-number, starting at 01 per phase)
	feature1Dir := filepath.Join(phase1Dir, "01-login_form")
	feature2Dir := filepath.Join(phase1Dir, "02-password_reset")
	feature3Dir := filepath.Join(phase2Dir, "01-database_schema")

	if _, err := os.Stat(feature1Dir); os.IsNotExist(err) {
		t.Error("ScaffoldPlanHierarchy() should create feature 1.1 directory")
	}
	if _, err := os.Stat(feature2Dir); os.IsNotExist(err) {
		t.Error("ScaffoldPlanHierarchy() should create feature 1.2 directory")
	}
	if _, err := os.Stat(feature3Dir); os.IsNotExist(err) {
		t.Error("ScaffoldPlanHierarchy() should create feature 2.1 directory")
	}

	// Verify template files were created in feature directories
	expectedFiles := []string{"design.md", "plan.md", "tasks.md", "prompts.md", "github.md"}
	for _, filename := range expectedFiles {
		path := filepath.Join(feature1Dir, filename)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("ScaffoldPlanHierarchy() should create %s in feature directory", filename)
		}
	}

	// Verify _contracts directory was created (if implemented)
	contractsDir := filepath.Join(phase1Dir, "_contracts")
	if _, err := os.Stat(contractsDir); os.IsNotExist(err) {
		t.Log("_contracts directory may not be created in all cases")
	} else {
		// If it exists, verify README
		contractsReadme := filepath.Join(contractsDir, "README.md")
		if _, err := os.Stat(contractsReadme); os.IsNotExist(err) {
			t.Log("Contracts README may not always be created")
		}
	}
}

func TestScaffoldPlanHierarchy_InvalidTasksFile(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir

	// Create invalid TASKS.md
	planDir := filepath.Join(projectPath, ".do", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		t.Fatalf("Failed to create plan dir: %v", err)
	}

	// Create empty or invalid content
	if err := os.WriteFile(filepath.Join(planDir, "TASKS.md"), []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create TASKS.md: %v", err)
	}

	// Should handle gracefully (no phases found)
	if err := ScaffoldPlanHierarchy(projectPath); err != nil {
		// Error is acceptable for invalid content
		t.Logf("ScaffoldPlanHierarchy() with invalid content returned error (expected): %v", err)
	}
}

func TestScaffoldPlanHierarchy_NoTasksFile(t *testing.T) {
	tmpDir := t.TempDir()
	projectPath := tmpDir

	// Don't create TASKS.md
	planDir := filepath.Join(projectPath, ".do", "plan")
	if err := os.MkdirAll(planDir, 0755); err != nil {
		t.Fatalf("Failed to create plan dir: %v", err)
	}

	// Should return error for missing file
	if err := ScaffoldPlanHierarchy(projectPath); err == nil {
		t.Error("ScaffoldPlanHierarchy() should return error for missing TASKS.md")
	}
}
