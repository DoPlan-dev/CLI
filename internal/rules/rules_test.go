package rules

import (
	"io/fs"
	"path/filepath"
	"testing"
)

func TestLibrary(t *testing.T) {
	lib := Library()
	if lib == nil {
		t.Fatal("Library() should not return nil")
	}

	// Try to read a known file (paths in embed.FS are relative to the embed directive)
	_, err := lib.Open("library/01-core-workflow/README.md")
	if err != nil {
		t.Errorf("Library() should contain library/01-core-workflow/README.md: %v", err)
	}
}

func TestReadFile(t *testing.T) {
	// Test reading a known file
	content, err := ReadFile("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if len(content) == 0 {
		t.Error("ReadFile() should return non-empty content")
	}

	// Test reading a non-existent file
	_, err = ReadFile("nonexistent/file.md")
	if err == nil {
		t.Error("ReadFile() should return error for non-existent file")
	}
}

func TestReadDir(t *testing.T) {
	// Test reading a known directory
	entries, err := ReadDir("01-core-workflow")
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}

	if len(entries) == 0 {
		t.Error("ReadDir() should return non-empty entries")
	}

	// Verify we can find README.md
	foundREADME := false
	for _, entry := range entries {
		if entry.Name() == "README.md" {
			foundREADME = true
			break
		}
	}

	if !foundREADME {
		t.Error("ReadDir() should contain README.md")
	}

	// Test reading a non-existent directory
	_, err = ReadDir("nonexistent")
	if err == nil {
		t.Error("ReadDir() should return error for non-existent directory")
	}
}

func TestWalkDir(t *testing.T) {
	var files []string

	err := WalkDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Remove "library/" prefix from path for comparison
		relPath := path
		if len(path) > 8 && path[:8] == "library/" {
			relPath = path[8:]
		}
		if !d.IsDir() {
			files = append(files, relPath)
		}
		return nil
	})

	if err != nil {
		t.Fatalf("WalkDir() error = %v", err)
	}

	if len(files) == 0 {
		t.Error("WalkDir() should find files")
	}

	// Verify we found some expected files
	foundREADME := false
	for _, file := range files {
		if filepath.Base(file) == "README.md" {
			foundREADME = true
			break
		}
	}

	if !foundREADME {
		t.Error("WalkDir() should find README.md files")
	}
}

func TestEmbeddedRules_AllCategories(t *testing.T) {
	// Verify all 15 category directories exist
	categories := []string{
		"01-core-workflow",
		"02-ai-agents",
		"03-languages",
		"04-frameworks",
		"05-ui-libraries",
		"06-cloud-infrastructure",
		"07-databases",
		"08-testing",
		"09-devops-ci-cd",
		"10-code-quality",
		"11-documentation",
		"12-security",
		"13-development-practices",
		"14-mcp-tools",
		"15-project-specific",
	}

	for _, category := range categories {
		t.Run(category, func(t *testing.T) {
			entries, err := ReadDir(category)
			if err != nil {
				t.Fatalf("ReadDir(%s) error = %v", category, err)
			}

			if len(entries) == 0 {
				t.Errorf("Category %s should contain files", category)
			}

			// Verify README.md exists
			foundREADME := false
			for _, entry := range entries {
				if entry.Name() == "README.md" {
					foundREADME = true
					break
				}
			}

			if !foundREADME {
				t.Errorf("Category %s should contain README.md", category)
			}
		})
	}
}
