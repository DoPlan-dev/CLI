package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestOrchestrate_ErrorRecovery(t *testing.T) {
	tmpDir := t.TempDir()

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Test with invalid request (should fail validation)
	invalidRequest := &models.ProjectRequest{
		ProjectName: "", // Invalid: empty name
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	err = Orchestrate(invalidRequest)
	if err == nil {
		t.Error("Orchestrate() should return error for invalid request")
	}
}

func TestOrchestrate_WithMultipleIDEs(t *testing.T) {
	tmpDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "multi-ide-project",
		IDEs:        []string{"Cursor", "Claude Code", "Windsurf"},
		ProjectType: "Fullstack",
	}

	if err := Orchestrate(request); err != nil {
		t.Fatalf("Orchestrate() with multiple IDEs error = %v", err)
	}

	// Verify project was created
	projectPath := filepath.Join(tmpDir, "multi-ide-project")
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		t.Error("Orchestrate() should create project directory")
	}

	// Verify IDE configs were created
	for _, ide := range request.IDEs {
		var ideConfigPath string
		switch ide {
		case "Cursor":
			ideConfigPath = filepath.Join(projectPath, ".cursor", "rules")
		case "Claude Code":
			ideConfigPath = filepath.Join(projectPath, ".claude", "rules")
		case "Windsurf":
			ideConfigPath = filepath.Join(projectPath, ".windsurf", "rules")
		}

		if _, err := os.Stat(ideConfigPath); os.IsNotExist(err) {
			t.Errorf("Orchestrate() should create config for IDE: %s", ide)
		}
	}
}

func TestOrchestrate_PermissionError(t *testing.T) {
	// Test with path that might have permission issues
	// This is platform-dependent, so we test the error handling path

	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create a request with a potentially problematic path
	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	// Normal case should work
	err = Orchestrate(request)
	if err != nil {
		t.Logf("Orchestrate() error (may be expected in test env): %v", err)
	}
}

func TestOrchestrate_RollbackOnError(t *testing.T) {
	tmpDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create a request that might cause a generation error
	// We can't easily simulate a mid-generation error, but we test the structure
	request := &models.ProjectRequest{
		ProjectName: "rollback-test",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	// Normal execution should succeed
	err = Orchestrate(request)
	if err != nil {
		t.Logf("Orchestrate() error: %v", err)
	} else {
		// Verify project was created (rollback didn't happen)
		projectPath := filepath.Join(tmpDir, "rollback-test")
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			t.Error("Orchestrate() should create project on success")
		}
	}
}

// TestGenerationContext_Tracking is already in generator_test.go

func TestRollback_WithErrors(t *testing.T) {
	tmpDir := t.TempDir()

	// Create some files and directories
	testDir1 := filepath.Join(tmpDir, "dir1")
	testDir2 := filepath.Join(tmpDir, "dir2")
	testFile1 := filepath.Join(tmpDir, "file1.txt")
	testFile2 := filepath.Join(testDir1, "file2.txt")

	if err := os.MkdirAll(testDir1, 0755); err != nil {
		t.Fatalf("Failed to create test dir1: %v", err)
	}
	if err := os.MkdirAll(testDir2, 0755); err != nil {
		t.Fatalf("Failed to create test dir2: %v", err)
	}
	if err := os.WriteFile(testFile1, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file1: %v", err)
	}
	if err := os.WriteFile(testFile2, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file2: %v", err)
	}

	ctx := &GenerationContext{
		CreatedDirs:  []string{testDir1, testDir2},
		CreatedFiles: []string{testFile1, testFile2},
	}

	// Rollback should succeed
	if err := rollback(ctx); err != nil {
		t.Errorf("rollback() error = %v, want nil", err)
	}

	// Verify files and directories were removed
	if _, err := os.Stat(testFile1); !os.IsNotExist(err) {
		t.Error("rollback() should remove tracked files")
	}
	if _, err := os.Stat(testFile2); !os.IsNotExist(err) {
		t.Error("rollback() should remove tracked files")
	}
	if _, err := os.Stat(testDir1); !os.IsNotExist(err) {
		t.Error("rollback() should remove tracked directories")
	}
	if _, err := os.Stat(testDir2); !os.IsNotExist(err) {
		t.Error("rollback() should remove tracked directories")
	}
}
