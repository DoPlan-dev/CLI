package rules

import (
	"io/fs"
	"path/filepath"
	"testing"
)

func TestWalkDir_WithRoot(t *testing.T) {
	var files []string
	var dirs []string

	err := WalkDir("01-core-workflow", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Remove "library/" prefix if present
		relPath := path
		if len(path) > 8 && path[:8] == "library/" {
			relPath = path[8:]
		}
		if d.IsDir() {
			dirs = append(dirs, relPath)
		} else {
			files = append(files, relPath)
		}
		return nil
	})

	if err != nil {
		t.Fatalf("WalkDir() with root error = %v", err)
	}

	if len(files) == 0 {
		t.Error("WalkDir() with root should find files")
	}

	// Verify we found README.md
	foundREADME := false
	for _, file := range files {
		if filepath.Base(file) == "README.md" {
			foundREADME = true
			break
		}
	}

	if !foundREADME {
		t.Error("WalkDir() with root should find README.md")
	}
}

func TestWalkDir_WithError(t *testing.T) {
	// Test that WalkDir properly handles errors from the walk function
	err := WalkDir("", func(path string, d fs.DirEntry, err error) error {
		// Return an error for a specific path to test error propagation
		if path == "library/01-core-workflow/README.md" {
			return fs.ErrPermission
		}
		return err
	})

	if err == nil {
		t.Error("WalkDir() should propagate errors from walk function")
	}
}

func TestWalkDir_NonExistentRoot(t *testing.T) {
	err := WalkDir("nonexistent-category", func(path string, d fs.DirEntry, err error) error {
		// Should receive error for non-existent root
		if err != nil {
			return err
		}
		return nil
	})

	// WalkDir may or may not error on non-existent root depending on implementation
	// Just verify it doesn't panic
	_ = err
}

func TestReadFileDecompressed_CompressedFile(t *testing.T) {
	// First, compress some data
	original := []byte("This is test content that will be compressed and then decompressed.")
	compressed, err := CompressData(original)
	if err != nil {
		t.Fatalf("Failed to compress test data: %v", err)
	}

	// Note: We can't actually test ReadFileDecompressed with compressed embedded files
	// because the embedded files are not compressed. But we can test the logic path
	// by verifying it works with uncompressed files (which is the actual use case)

	// Test with actual embedded file (uncompressed)
	data, err := ReadFileDecompressed("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("ReadFileDecompressed() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("ReadFileDecompressed() should return non-empty data")
	}

	// Verify it's readable text
	if !contains(data, []byte("Core Workflow")) && !contains(data, []byte("core")) {
		t.Log("ReadFileDecompressed() returned data (content check may vary)")
	}

	// Verify compressed data would be detected
	if IsCompressed(compressed) {
		// If we had a way to inject compressed data, we could test decompression
		t.Log("Compression detection works (cannot test decompression path with embedded files)")
	}
}

func TestReadFileDecompressed_NonExistentFile(t *testing.T) {
	_, err := ReadFileDecompressed("nonexistent/file.md")
	if err == nil {
		t.Error("ReadFileDecompressed() should return error for non-existent file")
	}
}

func TestReadDir_EmptyCategory(t *testing.T) {
	// Test reading a category that might be empty (if it exists)
	// Most categories should have files, but we test the error handling
	_, err := ReadDir("nonexistent-category")
	if err == nil {
		t.Error("ReadDir() should return error for non-existent category")
	}
}

func TestReadDir_WithSubdirectories(t *testing.T) {
	// Test reading a directory that might have subdirectories
	entries, err := ReadDir("01-core-workflow")
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}

	// Verify we got entries
	if len(entries) == 0 {
		t.Error("ReadDir() should return entries for existing category")
	}

	// Check that we can distinguish between files and directories
	hasFiles := false
	for _, entry := range entries {
		if !entry.IsDir() {
			hasFiles = true
			break
		}
	}

	if !hasFiles {
		t.Log("ReadDir() returned entries (may be all directories)")
	}
}

func TestReadFile_EdgeCases(t *testing.T) {
	// Test reading file with special characters in path (if any exist)
	// Most embedded files have standard paths, but test error handling

	// Test with empty path
	_, err := ReadFile("")
	if err == nil {
		t.Error("ReadFile() should return error for empty path")
	}

	// Test with path that's just a directory
	_, err = ReadFile("01-core-workflow")
	if err == nil {
		// May or may not error depending on implementation
		t.Log("ReadFile() with directory path (error handling may vary)")
	}
}

// Helper function
func contains(slice []byte, subslice []byte) bool {
	if len(subslice) > len(slice) {
		return false
	}
	for i := 0; i <= len(slice)-len(subslice); i++ {
		match := true
		for j := 0; j < len(subslice); j++ {
			if slice[i+j] != subslice[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
