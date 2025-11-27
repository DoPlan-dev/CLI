package generator

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

type mockGenerator struct {
	name   string
	err    error
	onRun  func(projectPath string)
	called bool
}

func (m *mockGenerator) Generate(_ *models.ProjectRequest, projectPath string) error {
	m.called = true
	if m.onRun != nil {
		m.onRun(projectPath)
	}
	return m.err
}

func (m *mockGenerator) Name() string {
	return m.name
}

func withTempCWD(t *testing.T, dir string) func() {
	t.Helper()
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}
	return func() {
		_ = os.Chdir(oldDir)
	}
}

func TestOrchestrate_UsesGenerationSteps(t *testing.T) {
	t.Cleanup(func() { generationStepFactory = defaultGenerationSteps })

	successGen := &mockGenerator{name: "success"}
	generationStepFactory = func() []GenerationStep {
		return []GenerationStep{{Generator: successGen, Name: "mock"}}
	}

	tmp := t.TempDir()
	restore := withTempCWD(t, tmp)
	defer restore()

	req := &models.ProjectRequest{
		ProjectName: "orchestrate_success",
		IDEs:        []string{"Cursor"},
	}

	if err := Orchestrate(req); err != nil {
		t.Fatalf("Orchestrate() error = %v", err)
	}
	if !successGen.called {
		t.Fatal("expected mock generator to be called")
	}
	projectPath := filepath.Join(tmp, "orchestrate_success")
	if _, err := os.Stat(projectPath); err != nil {
		t.Fatalf("expected project directory to exist: %v", err)
	}
}

func TestOrchestrate_RollbackOnFailure(t *testing.T) {
	t.Cleanup(func() { generationStepFactory = defaultGenerationSteps })

	failGen := &mockGenerator{
		name: "fail",
		err:  errors.New("boom"),
		onRun: func(projectPath string) {
			// create dummy file to verify rollback cleans up
			_ = os.WriteFile(filepath.Join(projectPath, "temp.txt"), []byte("temp"), 0644)
		},
	}
	generationStepFactory = func() []GenerationStep {
		return []GenerationStep{{Generator: failGen, Name: "mock"}}
	}

	tmp := t.TempDir()
	restore := withTempCWD(t, tmp)
	defer restore()

	req := &models.ProjectRequest{
		ProjectName: "orchestrate_failure",
		IDEs:        []string{"Cursor"},
	}

	err := Orchestrate(req)
	if err == nil {
		t.Fatal("Orchestrate() expected error, got nil")
	}
	if !failGen.called {
		t.Fatal("expected mock generator to be called")
	}

	projectPath := filepath.Join(tmp, "orchestrate_failure")
	if _, statErr := os.Stat(projectPath); !os.IsNotExist(statErr) {
		t.Fatalf("expected project directory to be removed, got err=%v", statErr)
	}
}
func TestOrchestrate_ValidRequest(t *testing.T) {
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

	// Create valid request
	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	// Should succeed (no generators to run yet, but structure is created)
	err = Orchestrate(request)
	if err != nil {
		t.Errorf("Orchestrate() with valid request error = %v, want nil", err)
	}

	// Verify project directory was created
	projectPath := filepath.Join(tmpDir, "test-project")
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		t.Error("Orchestrate() should create project directory")
	}
}

func TestOrchestrate_InvalidRequest(t *testing.T) {
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

	testCases := []struct {
		name    string
		request *models.ProjectRequest
		wantErr bool
	}{
		{
			name: "empty project name",
			request: &models.ProjectRequest{
				ProjectName: "",
				IDE:         "Cursor",
				ProjectType: "Fullstack",
			},
			wantErr: true,
		},
		{
			name: "empty IDE",
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDE:         "",
				ProjectType: "Fullstack",
			},
			wantErr: true,
		},
		{
			name: "invalid project name",
			request: &models.ProjectRequest{
				ProjectName: "test@project",
				IDE:         "Cursor",
				ProjectType: "Fullstack",
			},
			wantErr: true,
		},
		{
			name: "unsupported IDE",
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDE:         "UnsupportedIDE",
				ProjectType: "Fullstack",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Orchestrate(tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("Orchestrate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestOrchestrate_ExistingDirectory(t *testing.T) {
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

	// Create existing directory
	projectName := "existing-project"
	projectPath := filepath.Join(tmpDir, projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		t.Fatalf("Failed to create existing directory: %v", err)
	}

	// Try to orchestrate with existing directory
	request := &models.ProjectRequest{
		ProjectName: projectName,
		IDE:         "Cursor",
		ProjectType: "Fullstack",
	}

	err = Orchestrate(request)
	if err == nil {
		t.Error("Orchestrate() should fail when directory already exists")
	}
}

func TestGenerationContext_Tracking(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := &GenerationContext{
		Request: &models.ProjectRequest{
			ProjectName: "test",
			IDE:         "Cursor",
			ProjectType: "Fullstack",
		},
		ProjectPath:  tmpDir,
		CreatedDirs:  []string{},
		CreatedFiles: []string{},
	}

	// Test tracking directories
	dir1 := filepath.Join(tmpDir, "dir1")
	dir2 := filepath.Join(tmpDir, "dir2")
	trackCreatedDir(ctx, dir1)
	trackCreatedDir(ctx, dir2)

	if len(ctx.CreatedDirs) != 2 {
		t.Errorf("trackCreatedDir() CreatedDirs length = %d, want 2", len(ctx.CreatedDirs))
	}
	if ctx.CreatedDirs[0] != dir1 || ctx.CreatedDirs[1] != dir2 {
		t.Errorf("trackCreatedDir() CreatedDirs = %v, want [%s, %s]", ctx.CreatedDirs, dir1, dir2)
	}

	// Test tracking files
	file1 := filepath.Join(tmpDir, "file1.txt")
	file2 := filepath.Join(tmpDir, "file2.txt")
	trackCreatedFile(ctx, file1)
	trackCreatedFile(ctx, file2)

	if len(ctx.CreatedFiles) != 2 {
		t.Errorf("trackCreatedFile() CreatedFiles length = %d, want 2", len(ctx.CreatedFiles))
	}
	if ctx.CreatedFiles[0] != file1 || ctx.CreatedFiles[1] != file2 {
		t.Errorf("trackCreatedFile() CreatedFiles = %v, want [%s, %s]", ctx.CreatedFiles, file1, file2)
	}
}

func TestRollback(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test files and directories
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

	// Create context with tracked items
	ctx := &GenerationContext{
		CreatedDirs:  []string{testDir1, testDir2},
		CreatedFiles: []string{testFile1, testFile2},
	}

	// Rollback
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

func TestRollback_NonExistentItems(t *testing.T) {
	tmpDir := t.TempDir()
	// Create context with non-existent paths
	ctx := &GenerationContext{
		CreatedDirs:  []string{filepath.Join(tmpDir, "non", "existent", "dir")},
		CreatedFiles: []string{filepath.Join(tmpDir, "non", "existent", "file.txt")},
	}

	// Rollback should not error on non-existent items
	if err := rollback(ctx); err != nil {
		t.Errorf("rollback() on non-existent items error = %v, want nil", err)
	}
}

func TestRollback_EmptyContext(t *testing.T) {
	ctx := &GenerationContext{
		CreatedDirs:  []string{},
		CreatedFiles: []string{},
	}

	// Rollback should succeed with empty context
	if err := rollback(ctx); err != nil {
		t.Errorf("rollback() on empty context error = %v, want nil", err)
	}
}
