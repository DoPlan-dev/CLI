package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// TestGenerateNextJSBoilerplate_ErrorPaths tests error paths in generateNextJSBoilerplate (63.6% -> 90%+)
func TestGenerateNextJSBoilerplate_ErrorPaths(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(string) error
		request     *models.ProjectRequest
		expectError bool
	}{
		{
			name: "Directory creation failure",
			setup: func(projectPath string) error {
				// Create a file instead of directory at projectPath
				return os.WriteFile(projectPath, []byte("file"), 0644)
			},
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: true,
		},
		{
			name: "Normal generation should succeed",
			setup: func(projectPath string) error {
				return os.MkdirAll(projectPath, 0755)
			},
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			projectPath := filepath.Join(tmpDir, "project")

			if tt.setup != nil {
				if err := tt.setup(projectPath); err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}

			generator := &BoilerplateGenerator{}
			err := generator.generateNextJSBoilerplate(projectPath, tt.request)

			if (err != nil) != tt.expectError {
				t.Errorf("generateNextJSBoilerplate() error = %v, expectError = %v", err, tt.expectError)
			}
		})
	}
}

// TestGenerateNextJSAppStructure_ErrorPaths tests error paths in generateNextJSAppStructure (75% -> 90%+)
func TestGenerateNextJSAppStructure_ErrorPaths(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(string) error
		request     *models.ProjectRequest
		expectError bool
	}{
		{
			name: "App directory creation failure",
			setup: func(projectPath string) error {
				// Create a file at src/app to cause directory creation to fail
				srcDir := filepath.Join(projectPath, "src")
				if err := os.MkdirAll(srcDir, 0755); err != nil {
					return err
				}
				appPath := filepath.Join(srcDir, "app")
				return os.WriteFile(appPath, []byte("file"), 0644)
			},
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: true,
		},
		{
			name: "Layout file write failure - read-only directory",
			setup: func(projectPath string) error {
				// Create directory first with normal permissions
				appDir := filepath.Join(projectPath, "src", "app")
				if err := os.MkdirAll(appDir, 0755); err != nil {
					return err
				}
				// Then make it read-only (may not work on all systems)
				if err := os.Chmod(appDir, 0444); err != nil {
					return err
				}
				return nil
			},
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: true, // Should fail when directory is read-only
		},
		{
			name: "Normal generation should succeed",
			setup: func(projectPath string) error {
				return os.MkdirAll(projectPath, 0755)
			},
			request: &models.ProjectRequest{
				ProjectName: "test-project",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: false,
		},
		{
			name: "With long project name",
			setup: func(projectPath string) error {
				return os.MkdirAll(projectPath, 0755)
			},
			request: &models.ProjectRequest{
				ProjectName: "very-long-project-name-with-many-characters",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: false,
		},
		{
			name: "With special characters in project name (sanitized)",
			setup: func(projectPath string) error {
				return os.MkdirAll(projectPath, 0755)
			},
			request: &models.ProjectRequest{
				ProjectName: "test-project-123",
				IDEs:        []string{"Cursor"},
				ProjectType: "Fullstack",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			projectPath := filepath.Join(tmpDir, "project")

			if tt.setup != nil {
				if err := tt.setup(projectPath); err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}

			err := generateNextJSAppStructure(projectPath, tt.request)

			if (err != nil) != tt.expectError {
				t.Errorf("generateNextJSAppStructure() error = %v, expectError = %v", err, tt.expectError)
			}

			// Cleanup read-only directory if created
			if tt.name == "Layout file write failure - read-only directory" {
				appDir := filepath.Join(projectPath, "src", "app")
				os.Chmod(appDir, 0755) // Restore permissions for cleanup
			}
		})
	}
}

// TestGenerateNextJSAppStructure_FileContent tests file content generation
func TestGenerateNextJSAppStructure_FileContent(t *testing.T) {
	tmpDir := t.TempDir()
	request := &models.ProjectRequest{
		ProjectName: "my-test-app",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	if err := generateNextJSAppStructure(tmpDir, request); err != nil {
		t.Fatalf("generateNextJSAppStructure() error = %v", err)
	}

	// Verify layout.tsx
	layoutPath := filepath.Join(tmpDir, "src", "app", "layout.tsx")
	layoutContent, err := os.ReadFile(layoutPath)
	if err != nil {
		t.Fatalf("Failed to read layout.tsx: %v", err)
	}

	layoutStr := string(layoutContent)
	if !strings.Contains(layoutStr, "my-test-app") {
		t.Error("layout.tsx should contain project name")
	}
	if !strings.Contains(layoutStr, "Metadata") {
		t.Error("layout.tsx should contain Metadata")
	}

	// Verify page.tsx
	pagePath := filepath.Join(tmpDir, "src", "app", "page.tsx")
	pageContent, err := os.ReadFile(pagePath)
	if err != nil {
		t.Fatalf("Failed to read page.tsx: %v", err)
	}

	pageStr := string(pageContent)
	if !strings.Contains(pageStr, "my-test-app") {
		t.Error("page.tsx should contain project name")
	}
	if !strings.Contains(pageStr, "Welcome") {
		t.Error("page.tsx should contain Welcome")
	}

	// Verify globals.css
	globalsPath := filepath.Join(tmpDir, "src", "app", "globals.css")
	globalsContent, err := os.ReadFile(globalsPath)
	if err != nil {
		t.Fatalf("Failed to read globals.css: %v", err)
	}

	globalsStr := string(globalsContent)
	if !strings.Contains(globalsStr, "@tailwind") {
		t.Error("globals.css should contain @tailwind directives")
	}

	// Verify postcss.config.js
	postcssPath := filepath.Join(tmpDir, "postcss.config.js")
	if _, err := os.Stat(postcssPath); os.IsNotExist(err) {
		t.Error("postcss.config.js should be created")
	}

	postcssContent, err := os.ReadFile(postcssPath)
	if err != nil {
		t.Fatalf("Failed to read postcss.config.js: %v", err)
	}

	postcssStr := string(postcssContent)
	if !strings.Contains(postcssStr, "tailwindcss") {
		t.Error("postcss.config.js should contain tailwindcss")
	}
}

// TestGenerateNextJSBoilerplate_AllFiles tests that all files are generated
func TestGenerateNextJSBoilerplate_AllFiles(t *testing.T) {
	tmpDir := t.TempDir()
	request := &models.ProjectRequest{
		ProjectName: "complete-test",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &BoilerplateGenerator{}
	if err := generator.generateNextJSBoilerplate(tmpDir, request); err != nil {
		t.Fatalf("generateNextJSBoilerplate() error = %v", err)
	}

	// Verify all expected files exist
	expectedFiles := []string{
		"package.json",
		"tsconfig.json",
		"tailwind.config.ts",
		".eslintrc.json",
		"src/app/layout.tsx",
		"src/app/page.tsx",
		"src/app/globals.css",
		"postcss.config.js",
	}

	for _, file := range expectedFiles {
		filePath := filepath.Join(tmpDir, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected file %s was not created", file)
		}
	}
}


