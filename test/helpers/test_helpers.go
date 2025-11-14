package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

// CreateTempProject creates a temporary project directory
func CreateTempProject(t *testing.T) string {
	dir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})
	return dir
}

// SetupTestProject sets up a test project with structure
func SetupTestProject(t *testing.T) string {
	projectRoot := CreateTempProject(t)

	// Create directory structure
	dirs := []string{
		".cursor/commands",
		".cursor/rules",
		".cursor/config",
		"doplan/contracts",
		"doplan/templates",
		".doplan",
	}

	for _, dir := range dirs {
		path := filepath.Join(projectRoot, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	return projectRoot
}

// LoadTestFixture loads a test fixture file
func LoadTestFixture(t *testing.T, name string) []byte {
	path := filepath.Join("test", "fixtures", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load fixture %s: %v", name, err)
	}
	return data
}

// WriteTestFile writes content to a test file
func WriteTestFile(t *testing.T, projectRoot, path string, content []byte) {
	fullPath := filepath.Join(projectRoot, path)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("Failed to create directory for %s: %v", path, err)
	}
	if err := os.WriteFile(fullPath, content, 0644); err != nil {
		t.Fatalf("Failed to write test file %s: %v", path, err)
	}
}
