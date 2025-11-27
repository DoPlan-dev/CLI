package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// CreateDirectory creates a directory and all necessary parent directories.
// Returns an error if the directory cannot be created or already exists as a file.
func CreateDirectory(path string) error {
	// Sanitize path first
	sanitized, err := SanitizePath(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	// Check if path already exists
	info, err := os.Stat(sanitized)
	if err == nil {
		// Path exists
		if !info.IsDir() {
			return fmt.Errorf("path exists but is not a directory: %s", sanitized)
		}
		// Directory already exists, that's okay
		return nil
	}

	// Path doesn't exist, create it
	if err := os.MkdirAll(sanitized, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", sanitized, err)
	}

	return nil
}

// WriteFile writes data to a file atomically (using temp file then rename).
// Creates parent directories if they don't exist.
// Returns an error if the file cannot be written.
func WriteFile(path string, data []byte) error {
	// Sanitize path first
	sanitized, err := SanitizePath(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	// Create parent directories
	parentDir := filepath.Dir(sanitized)
	if err := CreateDirectory(parentDir); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	// Write atomically: create temp file, write data, rename
	tempFile := sanitized + ".tmp"

	// Remove temp file if it exists (cleanup from previous failed write)
	_ = os.Remove(tempFile)

	// Write to temp file
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	// Rename temp file to final file (atomic operation)
	if err := os.Rename(tempFile, sanitized); err != nil {
		// Clean up temp file on error
		_ = os.Remove(tempFile)
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

// ValidatePath validates that a path is safe to use.
// Checks for:
// - Empty paths
// - Path traversal attempts (..)
// - Invalid characters
// Returns the sanitized path and an error if invalid.
func ValidatePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	// Check for path traversal
	if strings.Contains(path, "..") {
		return "", fmt.Errorf("path contains '..' which is not allowed")
	}

	// Check for absolute paths (we want relative paths)
	if filepath.IsAbs(path) {
		// Allow absolute paths but warn - in practice we'll use relative
		// For now, we'll allow it but could restrict if needed
	}

	// Sanitize the path
	sanitized, err := SanitizePath(path)
	if err != nil {
		return "", err
	}

	return sanitized, nil
}

// SanitizePath sanitizes a file path by:
// - Cleaning the path (removing redundant separators, etc.)
// - Removing invalid characters
// - Normalizing separators
func SanitizePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	// Clean the path (removes redundant separators, resolves . and ..)
	cleaned := filepath.Clean(path)

	// Check for invalid characters in each component
	components := strings.Split(cleaned, string(filepath.Separator))
	for _, component := range components {
		if component == "" {
			continue // Empty components are okay (leading/trailing separators)
		}

		// Check for invalid characters
		for _, r := range component {
			// Allow: letters, digits, hyphens, underscores, dots
			if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '.') {
				return "", fmt.Errorf("path contains invalid character '%c' in component '%s'", r, component)
			}
		}

		// Don't allow components that are just dots (except . and .. which filepath.Clean handles)
		if component == "." || component == ".." {
			// filepath.Clean should have resolved these, but check anyway
			return "", fmt.Errorf("path contains unresolved '.' or '..'")
		}
	}

	return cleaned, nil
}

// CheckPermissions checks if we have the necessary permissions to:
// - Read from a directory (if it exists)
// - Write to a directory (if it exists)
// - Create a directory (if it doesn't exist)
// Returns an error if permissions are insufficient.
func CheckPermissions(path string) error {
	// Check if path exists
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Path doesn't exist, check parent directory
		parentDir := filepath.Dir(path)
		// Ensure parent directory exists before checking permissions
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			// Parent doesn't exist, check its parent recursively
			return CheckPermissions(parentDir)
		}
		return checkWritePermissions(parentDir)
	}
	if err != nil {
		return fmt.Errorf("failed to check path: %w", err)
	}

	// Path exists
	if info.IsDir() {
		// Check if we can write to the directory
		return checkWritePermissions(path)
	}

	// Path is a file, check parent directory
	parentDir := filepath.Dir(path)
	return checkWritePermissions(parentDir)
}

// checkWritePermissions checks if we can write to a directory
func checkWritePermissions(dirPath string) error {
	// Try to create a test file
	testFile := filepath.Join(dirPath, ".doplan_test_write")

	// Remove test file if it exists
	_ = os.Remove(testFile)

	// Try to create the test file
	file, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("insufficient permissions to write to %s: %w. Hint: run `npx --yes @doplan-dev/cli goplan access all` (or target `.do/system` / `docs`) from your project root, then try again.", dirPath, err)
	}

	// Clean up
	file.Close()
	_ = os.Remove(testFile)

	return nil
}

// PathExists checks if a path exists (file or directory)
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDirectory checks if a path exists and is a directory
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile checks if a path exists and is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// CreateSymlink creates a symbolic link from linkPath to targetPath.
// On Windows, this may require administrator privileges.
// Returns an error if the symlink cannot be created.
func CreateSymlink(linkPath, targetPath string) error {
	// Ensure parent directory exists
	parentDir := filepath.Dir(linkPath)
	if err := CreateDirectory(parentDir); err != nil {
		return fmt.Errorf("failed to create parent directory for symlink: %w", err)
	}

	// Remove existing link/file if it exists
	if PathExists(linkPath) {
		if err := os.Remove(linkPath); err != nil {
			return fmt.Errorf("failed to remove existing path before creating symlink: %w", err)
		}
	}

	// Calculate relative path for better portability
	relTarget, err := filepath.Rel(filepath.Dir(linkPath), targetPath)
	if err != nil {
		// If relative path calculation fails, use absolute path
		relTarget = targetPath
	}

	// Create the symlink
	if err := os.Symlink(relTarget, linkPath); err != nil {
		return fmt.Errorf("failed to create symlink from %s to %s: %w", linkPath, relTarget, err)
	}

	return nil
}

// IsSymlink checks if a path is a symbolic link
func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

// CopyDirectory recursively copies a directory from src to dst
func CopyDirectory(src, dst string) error {
	// Create destination directory
	if err := CreateDirectory(dst); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Walk source directory
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return CreateDirectory(dstPath)
		}

		// Read and copy file
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return WriteFile(dstPath, data)
	})
}
