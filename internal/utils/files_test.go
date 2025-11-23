package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSanitizePath(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"valid path", "my-project/src", false},
		{"valid nested path", "my-project/src/components", false},
		{"path with hyphens", "my-awesome-project", false},
		{"path with underscores", "my_awesome_project", false},
		{"path with dots", "my.project", false},
		{"empty path", "", true},
		{"path with invalid char", "my@project", true},
		{"path with space", "my project", true},
		{"path with slash", "my/project", false}, // filepath.Clean handles this
		{"relative path", "./my-project", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := SanitizePath(tc.path)
			if (err != nil) != tc.wantErr {
				t.Errorf("SanitizePath(%q) error = %v, wantErr %v", tc.path, err, tc.wantErr)
			}
		})
	}
}

func TestValidatePath(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"valid path", "my-project/src", false},
		{"empty path", "", true},
		{"path traversal", "../parent", true},
		{"path traversal in middle", "my/../project", true},
		{"valid nested", "my-project/.cursor/agents", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ValidatePath(tc.path)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidatePath(%q) error = %v, wantErr %v", tc.path, err, tc.wantErr)
			}
		})
	}
}

func TestCreateDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("create new directory", func(t *testing.T) {
		dir := filepath.Join(tmpDir, "new-dir")
		if err := CreateDirectory(dir); err != nil {
			t.Errorf("CreateDirectory() error = %v", err)
		}

		if !IsDirectory(dir) {
			t.Error("CreateDirectory() should create the directory")
		}
	})

	t.Run("create nested directories", func(t *testing.T) {
		dir := filepath.Join(tmpDir, "nested", "deep", "path")
		if err := CreateDirectory(dir); err != nil {
			t.Errorf("CreateDirectory() error = %v", err)
		}

		if !IsDirectory(dir) {
			t.Error("CreateDirectory() should create nested directories")
		}
	})

	t.Run("create existing directory", func(t *testing.T) {
		dir := filepath.Join(tmpDir, "existing")
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create test dir: %v", err)
		}

		// Should not error if directory already exists
		if err := CreateDirectory(dir); err != nil {
			t.Errorf("CreateDirectory() on existing dir error = %v", err)
		}
	})

	t.Run("fail on existing file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "existing-file")
		if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		if err := CreateDirectory(file); err == nil {
			t.Error("CreateDirectory() on existing file should error")
		}
	})
}

func TestWriteFile(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("write new file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "test.txt")
		data := []byte("test content")

		if err := WriteFile(file, data); err != nil {
			t.Errorf("WriteFile() error = %v", err)
		}

		// Verify file was written
		readData, err := os.ReadFile(file)
		if err != nil {
			t.Errorf("Failed to read written file: %v", err)
		}

		if string(readData) != string(data) {
			t.Errorf("WriteFile() content = %q, want %q", string(readData), string(data))
		}
	})

	t.Run("write file in new directory", func(t *testing.T) {
		file := filepath.Join(tmpDir, "new-dir", "test.txt")
		data := []byte("test content")

		if err := WriteFile(file, data); err != nil {
			t.Errorf("WriteFile() error = %v", err)
		}

		if !IsFile(file) {
			t.Error("WriteFile() should create parent directories and file")
		}
	})

	t.Run("overwrite existing file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "overwrite.txt")
		oldData := []byte("old content")
		newData := []byte("new content")

		// Create initial file
		if err := os.WriteFile(file, oldData, 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		// Overwrite it
		if err := WriteFile(file, newData); err != nil {
			t.Errorf("WriteFile() error = %v", err)
		}

		// Verify new content
		readData, err := os.ReadFile(file)
		if err != nil {
			t.Errorf("Failed to read overwritten file: %v", err)
		}

		if string(readData) != string(newData) {
			t.Errorf("WriteFile() overwrite content = %q, want %q", string(readData), string(newData))
		}
	})
}

func TestCheckPermissions(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("check permissions on writable directory", func(t *testing.T) {
		if err := CheckPermissions(tmpDir); err != nil {
			t.Errorf("CheckPermissions() on writable dir error = %v", err)
		}
	})

	t.Run("check permissions on non-existent path", func(t *testing.T) {
		path := filepath.Join(tmpDir, "non-existent", "file.txt")
		if err := CheckPermissions(path); err != nil {
			t.Errorf("CheckPermissions() on non-existent path error = %v", err)
		}
	})
}

func TestPathExists(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("existing path", func(t *testing.T) {
		if !PathExists(tmpDir) {
			t.Error("PathExists() should return true for existing path")
		}
	})

	t.Run("non-existent path", func(t *testing.T) {
		path := filepath.Join(tmpDir, "non-existent")
		if PathExists(path) {
			t.Error("PathExists() should return false for non-existent path")
		}
	})
}

func TestIsDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("existing directory", func(t *testing.T) {
		if !IsDirectory(tmpDir) {
			t.Error("IsDirectory() should return true for directory")
		}
	})

	t.Run("existing file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "test.txt")
		if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		if IsDirectory(file) {
			t.Error("IsDirectory() should return false for file")
		}
	})

	t.Run("non-existent path", func(t *testing.T) {
		path := filepath.Join(tmpDir, "non-existent")
		if IsDirectory(path) {
			t.Error("IsDirectory() should return false for non-existent path")
		}
	})
}

func TestIsFile(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("existing file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "test.txt")
		if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		if !IsFile(file) {
			t.Error("IsFile() should return true for file")
		}
	})

	t.Run("existing directory", func(t *testing.T) {
		if IsFile(tmpDir) {
			t.Error("IsFile() should return false for directory")
		}
	})

	t.Run("non-existent path", func(t *testing.T) {
		path := filepath.Join(tmpDir, "non-existent")
		if IsFile(path) {
			t.Error("IsFile() should return false for non-existent path")
		}
	})
}

