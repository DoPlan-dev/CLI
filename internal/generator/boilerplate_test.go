package generator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestBoilerplateGenerator_Name(t *testing.T) {
	generator := &BoilerplateGenerator{}
	if got := generator.Name(); got != "Boilerplate" {
		t.Errorf("BoilerplateGenerator.Name() = %v, want %v", got, "Boilerplate")
	}
}

func TestBoilerplateGenerator_Generate(t *testing.T) {
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

	generator := &BoilerplateGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("BoilerplateGenerator.Generate() error = %v", err)
	}

	// Verify package.json was created
	packagePath := filepath.Join(tmpDir, "package.json")
	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create package.json")
	}

	// Verify tsconfig.json was created
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if _, err := os.Stat(tsconfigPath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create tsconfig.json")
	}

	// Verify tailwind.config.ts was created
	tailwindPath := filepath.Join(tmpDir, "tailwind.config.ts")
	if _, err := os.Stat(tailwindPath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create tailwind.config.ts")
	}

	// Verify .eslintrc.json was created
	eslintPath := filepath.Join(tmpDir, ".eslintrc.json")
	if _, err := os.Stat(eslintPath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create .eslintrc.json")
	}

	// Verify src/app directory was created
	appDir := filepath.Join(tmpDir, "src", "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create src/app directory")
	}
}

func TestBoilerplateGenerator_Generate_FileContent(t *testing.T) {
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

	generator := &BoilerplateGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("BoilerplateGenerator.Generate() error = %v", err)
	}

	// Verify package.json content
	packagePath := filepath.Join(tmpDir, "package.json")
	content, err := os.ReadFile(packagePath)
	if err != nil {
		t.Fatalf("Failed to read package.json: %v", err)
	}

	var pkg map[string]interface{}
	if err := json.Unmarshal(content, &pkg); err != nil {
		t.Fatalf("package.json is not valid JSON: %v", err)
	}

	if pkg["name"] != request.ProjectName {
		t.Errorf("package.json name = %v, want %v", pkg["name"], request.ProjectName)
	}

	// Verify Next.js app files
	layoutPath := filepath.Join(tmpDir, "src", "app", "layout.tsx")
	if _, err := os.Stat(layoutPath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create layout.tsx")
	}

	pagePath := filepath.Join(tmpDir, "src", "app", "page.tsx")
	if _, err := os.Stat(pagePath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create page.tsx")
	}

	globalsPath := filepath.Join(tmpDir, "src", "app", "globals.css")
	if _, err := os.Stat(globalsPath); os.IsNotExist(err) {
		t.Error("BoilerplateGenerator.Generate() should create globals.css")
	}

	// Verify page.tsx contains project name
	pageContent, err := os.ReadFile(pagePath)
	if err != nil {
		t.Fatalf("Failed to read page.tsx: %v", err)
	}

	pageStr := string(pageContent)
	if !strings.Contains(pageStr, request.ProjectName) {
		t.Errorf("page.tsx should contain project name %s", request.ProjectName)
	}
}

func TestGenerateBoilerplate(t *testing.T) {
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

	if err := GenerateBoilerplate(request, tmpDir); err != nil {
		t.Fatalf("GenerateBoilerplate() error = %v", err)
	}

	// Verify package.json was created
	packagePath := filepath.Join(tmpDir, "package.json")
	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		t.Error("GenerateBoilerplate() should create package.json")
	}
}
