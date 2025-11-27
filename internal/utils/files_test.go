package utils

import (
	"os"
	"path/filepath"
	"runtime"
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

func TestWriteFileInvalidPath(t *testing.T) {
	if err := WriteFile("", []byte("data")); err == nil {
		t.Error("WriteFile() should fail for empty path")
	}
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

	t.Run("check permissions on read-only dir", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("file permissions unreliable on windows")
		}
		dir := filepath.Join(tmpDir, "readonly")
		if err := os.MkdirAll(dir, 0o555); err != nil {
			t.Fatalf("Failed to create readonly dir: %v", err)
		}
		if err := CheckPermissions(dir); err == nil {
			t.Error("CheckPermissions() should fail for read-only directory")
		}
	})
}

func TestCheckPermissions_ReadOnlyParent(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod semantics differ on Windows")
	}

	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	readonlyDir := filepath.Join(tmpDir, "readonly")
	if err := os.MkdirAll(readonlyDir, 0o755); err != nil {
		t.Fatalf("Failed to create dir: %v", err)
	}
	filePath := filepath.Join(readonlyDir, "file.txt")
	if err := os.WriteFile(filePath, []byte("data"), 0o644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if err := os.Chmod(readonlyDir, 0o555); err != nil {
		t.Fatalf("chmod failed: %v", err)
	}
	defer os.Chmod(readonlyDir, 0o755)

	if err := CheckPermissions(filePath); err == nil {
		t.Error("CheckPermissions() should fail when parent dir is read-only")
	}
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

func TestCreateSymlink(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("create symlink to file", func(t *testing.T) {
		// Create target file
		targetFile := filepath.Join(tmpDir, "target.txt")
		if err := os.WriteFile(targetFile, []byte("target content"), 0644); err != nil {
			t.Fatalf("Failed to create target file: %v", err)
		}

		// Create symlink
		linkPath := filepath.Join(tmpDir, "link.txt")
		if err := CreateSymlink(linkPath, targetFile); err != nil {
			t.Errorf("CreateSymlink() error = %v", err)
		}

		// Verify symlink exists
		if !IsSymlink(linkPath) {
			t.Error("CreateSymlink() should create a symlink")
		}

		// Verify symlink points to target
		linkTarget, err := os.Readlink(linkPath)
		if err != nil {
			t.Errorf("Failed to read symlink: %v", err)
		}
		if linkTarget == "" {
			t.Error("Symlink should have a target")
		}
	})

	t.Run("create symlink to directory", func(t *testing.T) {
		// Create target directory
		targetDir := filepath.Join(tmpDir, "target-dir")
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			t.Fatalf("Failed to create target directory: %v", err)
		}

		// Create symlink
		linkPath := filepath.Join(tmpDir, "link-dir")
		if err := CreateSymlink(linkPath, targetDir); err != nil {
			t.Errorf("CreateSymlink() to directory error = %v", err)
		}

		// Verify symlink exists
		if !IsSymlink(linkPath) {
			t.Error("CreateSymlink() should create a symlink to directory")
		}
	})

	t.Run("overwrite existing symlink", func(t *testing.T) {
		// Create first target
		target1 := filepath.Join(tmpDir, "target1.txt")
		if err := os.WriteFile(target1, []byte("target1"), 0644); err != nil {
			t.Fatalf("Failed to create target1: %v", err)
		}

		// Create second target
		target2 := filepath.Join(tmpDir, "target2.txt")
		if err := os.WriteFile(target2, []byte("target2"), 0644); err != nil {
			t.Fatalf("Failed to create target2: %v", err)
		}

		// Create initial symlink
		linkPath := filepath.Join(tmpDir, "overwrite-link.txt")
		if err := CreateSymlink(linkPath, target1); err != nil {
			t.Fatalf("Failed to create initial symlink: %v", err)
		}

		// Overwrite with new target
		if err := CreateSymlink(linkPath, target2); err != nil {
			t.Errorf("CreateSymlink() overwrite error = %v", err)
		}

		// Verify symlink points to new target
		linkTarget, err := os.Readlink(linkPath)
		if err != nil {
			t.Errorf("Failed to read overwritten symlink: %v", err)
		}
		if linkTarget == "" {
			t.Error("Overwritten symlink should have a target")
		}
	})

	t.Run("create symlink with non-existent parent", func(t *testing.T) {
		targetFile := filepath.Join(tmpDir, "target.txt")
		if err := os.WriteFile(targetFile, []byte("target"), 0644); err != nil {
			t.Fatalf("Failed to create target: %v", err)
		}

		// Create symlink in nested directory
		linkPath := filepath.Join(tmpDir, "nested", "deep", "link.txt")
		if err := CreateSymlink(linkPath, targetFile); err != nil {
			t.Errorf("CreateSymlink() with nested path error = %v", err)
		}

		if !IsSymlink(linkPath) {
			t.Error("CreateSymlink() should create parent directories and symlink")
		}
	})

	t.Run("fail when parent is file", func(t *testing.T) {
		targetFile := filepath.Join(tmpDir, "target.txt")
		if err := os.WriteFile(targetFile, []byte("target"), 0644); err != nil {
			t.Fatalf("Failed to create target: %v", err)
		}
		parent := filepath.Join(tmpDir, "file-parent")
		if err := os.WriteFile(parent, []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to create parent file: %v", err)
		}
		linkPath := filepath.Join(parent, "link.txt")
		if err := CreateSymlink(linkPath, targetFile); err == nil {
			t.Error("CreateSymlink() should fail when parent path is not a directory")
		}
	})
}

func TestCreateSymlink_PermissionDenied(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlink permissions vary on Windows")
	}

	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	targetFile := filepath.Join(tmpDir, "target")
	if err := os.WriteFile(targetFile, []byte("data"), 0o644); err != nil {
		t.Fatalf("Failed to create target: %v", err)
	}

	readonlyDir := filepath.Join(tmpDir, "readonly")
	if err := os.MkdirAll(readonlyDir, 0o555); err != nil {
		t.Fatalf("Failed to create readonly dir: %v", err)
	}
	defer os.Chmod(readonlyDir, 0o755)

	linkPath := filepath.Join(readonlyDir, "link.txt")
	if err := CreateSymlink(linkPath, targetFile); err == nil {
		t.Error("CreateSymlink() should fail when parent dir is read-only")
	}
}

func TestIsSymlink(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("symlink to file", func(t *testing.T) {
		// Create target file
		targetFile := filepath.Join(tmpDir, "target.txt")
		if err := os.WriteFile(targetFile, []byte("target"), 0644); err != nil {
			t.Fatalf("Failed to create target: %v", err)
		}

		// Create symlink
		linkPath := filepath.Join(tmpDir, "link.txt")
		if err := os.Symlink(targetFile, linkPath); err != nil {
			t.Fatalf("Failed to create symlink: %v", err)
		}

		if !IsSymlink(linkPath) {
			t.Error("IsSymlink() should return true for symlink")
		}
	})

	t.Run("regular file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "regular.txt")
		if err := os.WriteFile(file, []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		if IsSymlink(file) {
			t.Error("IsSymlink() should return false for regular file")
		}
	})

	t.Run("regular directory", func(t *testing.T) {
		dir := filepath.Join(tmpDir, "regular-dir")
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}

		if IsSymlink(dir) {
			t.Error("IsSymlink() should return false for regular directory")
		}
	})

	t.Run("non-existent path", func(t *testing.T) {
		path := filepath.Join(tmpDir, "non-existent")
		if IsSymlink(path) {
			t.Error("IsSymlink() should return false for non-existent path")
		}
	})
}

func TestCopyDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("copy simple directory", func(t *testing.T) {
		// Create source directory with files
		srcDir := filepath.Join(tmpDir, "src")
		if err := os.MkdirAll(srcDir, 0755); err != nil {
			t.Fatalf("Failed to create src dir: %v", err)
		}

		file1 := filepath.Join(srcDir, "file1.txt")
		file2 := filepath.Join(srcDir, "file2.txt")
		if err := os.WriteFile(file1, []byte("content1"), 0644); err != nil {
			t.Fatalf("Failed to create file1: %v", err)
		}
		if err := os.WriteFile(file2, []byte("content2"), 0644); err != nil {
			t.Fatalf("Failed to create file2: %v", err)
		}

		// Copy directory
		dstDir := filepath.Join(tmpDir, "dst")
		if err := CopyDirectory(srcDir, dstDir); err != nil {
			t.Errorf("CopyDirectory() error = %v", err)
		}

		// Verify destination exists
		if !IsDirectory(dstDir) {
			t.Error("CopyDirectory() should create destination directory")
		}

		// Verify files were copied
		dstFile1 := filepath.Join(dstDir, "file1.txt")
		dstFile2 := filepath.Join(dstDir, "file2.txt")
		if !IsFile(dstFile1) {
			t.Error("CopyDirectory() should copy files")
		}
		if !IsFile(dstFile2) {
			t.Error("CopyDirectory() should copy all files")
		}

		// Verify file contents
		content1, err := os.ReadFile(dstFile1)
		if err != nil {
			t.Errorf("Failed to read copied file1: %v", err)
		}
		if string(content1) != "content1" {
			t.Errorf("CopyDirectory() file1 content = %q, want %q", string(content1), "content1")
		}
	})

	t.Run("copy nested directory", func(t *testing.T) {
		// Create source with nested structure
		srcDir := filepath.Join(tmpDir, "src-nested")
		nestedDir := filepath.Join(srcDir, "nested", "deep")
		if err := os.MkdirAll(nestedDir, 0755); err != nil {
			t.Fatalf("Failed to create nested dir: %v", err)
		}

		nestedFile := filepath.Join(nestedDir, "nested.txt")
		if err := os.WriteFile(nestedFile, []byte("nested content"), 0644); err != nil {
			t.Fatalf("Failed to create nested file: %v", err)
		}

		// Copy directory
		dstDir := filepath.Join(tmpDir, "dst-nested")
		if err := CopyDirectory(srcDir, dstDir); err != nil {
			t.Errorf("CopyDirectory() nested error = %v", err)
		}

		// Verify nested structure
		dstNestedFile := filepath.Join(dstDir, "nested", "deep", "nested.txt")
		if !IsFile(dstNestedFile) {
			t.Error("CopyDirectory() should copy nested files")
		}

		content, err := os.ReadFile(dstNestedFile)
		if err != nil {
			t.Errorf("Failed to read nested file: %v", err)
		}
		if string(content) != "nested content" {
			t.Errorf("CopyDirectory() nested content = %q, want %q", string(content), "nested content")
		}
	})

	t.Run("copy empty directory", func(t *testing.T) {
		srcDir := filepath.Join(tmpDir, "src-empty")
		if err := os.MkdirAll(srcDir, 0755); err != nil {
			t.Fatalf("Failed to create empty dir: %v", err)
		}

		dstDir := filepath.Join(tmpDir, "dst-empty")
		if err := CopyDirectory(srcDir, dstDir); err != nil {
			t.Errorf("CopyDirectory() empty dir error = %v", err)
		}

		if !IsDirectory(dstDir) {
			t.Error("CopyDirectory() should create destination for empty directory")
		}
	})

	t.Run("copy to existing directory", func(t *testing.T) {
		srcDir := filepath.Join(tmpDir, "src-existing")
		if err := os.MkdirAll(srcDir, 0755); err != nil {
			t.Fatalf("Failed to create src: %v", err)
		}
		if err := os.WriteFile(filepath.Join(srcDir, "file.txt"), []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		dstDir := filepath.Join(tmpDir, "dst-existing")
		if err := os.MkdirAll(dstDir, 0755); err != nil {
			t.Fatalf("Failed to create existing dst: %v", err)
		}

		// Should succeed (overwrites files)
		if err := CopyDirectory(srcDir, dstDir); err != nil {
			t.Errorf("CopyDirectory() to existing dir error = %v", err)
		}

		// Verify file was copied
		dstFile := filepath.Join(dstDir, "file.txt")
		if !IsFile(dstFile) {
			t.Error("CopyDirectory() should copy files to existing directory")
		}
	})

	t.Run("error when source missing", func(t *testing.T) {
		srcDir := filepath.Join(tmpDir, "missing-src")
		dstDir := filepath.Join(tmpDir, "dst-missing-src")
		if err := CopyDirectory(srcDir, dstDir); err == nil {
			t.Error("CopyDirectory() should fail for missing source directory")
		}
	})
}

func TestCopyDirectory_DestinationIsFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	srcDir := filepath.Join(tmpDir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatalf("Failed to create src: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "file.txt"), []byte("data"), 0o644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	dstPath := filepath.Join(tmpDir, "dst-file")
	if err := os.WriteFile(dstPath, []byte("not a dir"), 0o644); err != nil {
		t.Fatalf("Failed to create dst file: %v", err)
	}

	if err := CopyDirectory(srcDir, dstPath); err == nil {
		t.Error("CopyDirectory() should fail when destination is an existing file")
	}
}

func TestCopyDirectory_SourceIsFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	srcFile := filepath.Join(tmpDir, "source-file")
	if err := os.WriteFile(srcFile, []byte("data"), 0o644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	dstDir := filepath.Join(tmpDir, "dst")
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		t.Fatalf("Failed to create dst dir: %v", err)
	}

	if err := CopyDirectory(srcFile, dstDir); err == nil {
		t.Error("CopyDirectory() should fail when source is not a directory")
	}
}
