package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyLibraryFolders(t *testing.T) {
	tmpDir := t.TempDir()

	// Create central library directory with category folders
	centralRulesDir := filepath.Join(tmpDir, "central", "library")
	category1Dir := filepath.Join(centralRulesDir, "01-core-workflow")
	category2Dir := filepath.Join(centralRulesDir, "02-ai-agents")

	if err := os.MkdirAll(category1Dir, 0755); err != nil {
		t.Fatalf("Failed to create category1 dir: %v", err)
	}
	if err := os.MkdirAll(category2Dir, 0755); err != nil {
		t.Fatalf("Failed to create category2 dir: %v", err)
	}

	// Create rule files in categories
	rule1File := filepath.Join(category1Dir, "README.md")
	rule2File := filepath.Join(category2Dir, "README.md")
	if err := os.WriteFile(rule1File, []byte("# Core Workflow"), 0644); err != nil {
		t.Fatalf("Failed to create rule1 file: %v", err)
	}
	if err := os.WriteFile(rule2File, []byte("# AI Agents"), 0644); err != nil {
		t.Fatalf("Failed to create rule2 file: %v", err)
	}

	// Create IDE rules directory
	ideRulesDir := filepath.Join(tmpDir, "ide", "rules")
	if err := os.MkdirAll(ideRulesDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE rules dir: %v", err)
	}

	// Test copyLibraryFolders
	if err := copyLibraryFolders(ideRulesDir, centralRulesDir); err != nil {
		t.Fatalf("copyLibraryFolders() error = %v", err)
	}

	// Verify category folders were copied
	ideCategory1Dir := filepath.Join(ideRulesDir, "01-core-workflow")
	ideCategory2Dir := filepath.Join(ideRulesDir, "02-ai-agents")

	if _, err := os.Stat(ideCategory1Dir); os.IsNotExist(err) {
		t.Error("copyLibraryFolders() should copy category folders")
	}
	if _, err := os.Stat(ideCategory2Dir); os.IsNotExist(err) {
		t.Error("copyLibraryFolders() should copy all category folders")
	}

	// Verify rule files were copied
	ideRule1File := filepath.Join(ideCategory1Dir, "README.md")
	ideRule2File := filepath.Join(ideCategory2Dir, "README.md")

	if _, err := os.Stat(ideRule1File); os.IsNotExist(err) {
		t.Error("copyLibraryFolders() should copy rule files")
	}
	if _, err := os.Stat(ideRule2File); os.IsNotExist(err) {
		t.Error("copyLibraryFolders() should copy all rule files")
	}

	// Verify file content
	content1, err := os.ReadFile(ideRule1File)
	if err != nil {
		t.Fatalf("Failed to read copied rule1 file: %v", err)
	}
	if string(content1) != "# Core Workflow" {
		t.Error("copyLibraryFolders() should preserve file content")
	}
}

func TestCopyLibraryFolders_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	centralRulesDir := filepath.Join(tmpDir, "central", "library")
	if err := os.MkdirAll(centralRulesDir, 0755); err != nil {
		t.Fatalf("Failed to create central dir: %v", err)
	}

	ideRulesDir := filepath.Join(tmpDir, "ide", "rules")
	if err := os.MkdirAll(ideRulesDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Should not error on empty directory
	if err := copyLibraryFolders(ideRulesDir, centralRulesDir); err != nil {
		t.Errorf("copyLibraryFolders() should handle empty directory, got error: %v", err)
	}
}

func TestCopyLibraryFolders_NonExistentCentralDir(t *testing.T) {
	tmpDir := t.TempDir()

	centralRulesDir := filepath.Join(tmpDir, "nonexistent", "library")
	ideRulesDir := filepath.Join(tmpDir, "ide", "rules")
	if err := os.MkdirAll(ideRulesDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Should return error for non-existent central directory
	if err := copyLibraryFolders(ideRulesDir, centralRulesDir); err == nil {
		t.Error("copyLibraryFolders() should return error for non-existent central directory")
	}
}

func TestCopyLibraryFolders_WithNestedFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create central library with nested structure
	centralRulesDir := filepath.Join(tmpDir, "central", "library")
	categoryDir := filepath.Join(centralRulesDir, "01-core-workflow")
	nestedDir := filepath.Join(categoryDir, "subfolder")

	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}

	// Create nested file
	nestedFile := filepath.Join(nestedDir, "nested.md")
	if err := os.WriteFile(nestedFile, []byte("# Nested Content"), 0644); err != nil {
		t.Fatalf("Failed to create nested file: %v", err)
	}

	ideRulesDir := filepath.Join(tmpDir, "ide", "rules")
	if err := os.MkdirAll(ideRulesDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Test copyLibraryFolders
	if err := copyLibraryFolders(ideRulesDir, centralRulesDir); err != nil {
		t.Fatalf("copyLibraryFolders() error = %v", err)
	}

	// Verify nested structure was copied
	ideNestedFile := filepath.Join(ideRulesDir, "01-core-workflow", "subfolder", "nested.md")
	if _, err := os.Stat(ideNestedFile); os.IsNotExist(err) {
		t.Error("copyLibraryFolders() should copy nested files")
	}
}
