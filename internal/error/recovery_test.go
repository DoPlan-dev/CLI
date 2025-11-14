package error

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRecoveryManager(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	assert.NotNil(t, rm)
	assert.NotNil(t, rm.handler)
}

func TestRecover_AutoFix(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	// Test directory not found - should create it
	projectRoot := helpers.CreateTempProject(t)
	missingDir := filepath.Join(projectRoot, "missing", "dir")
	err := ErrDirectoryNotFound(missingDir)

	result := rm.Recover(err, RecoveryAutoFix)
	assert.NoError(t, result)

	// Verify directory was created
	_, statErr := os.Stat(missingDir)
	assert.NoError(t, statErr)
}

func TestRecover_Skip(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	err := NewValidationError("VAL001", "Test error")
	result := rm.Recover(err, RecoverySkip)
	assert.NoError(t, result)
}

func TestRecover_NonRecoverable(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	// Create a non-recoverable error
	err := &DoPlanError{
		Code:        "TEST001",
		Message:     "Non-recoverable",
		Recoverable: false,
	}

	result := rm.Recover(err, RecoveryAutoFix)
	assert.Error(t, result)
	assert.Equal(t, err, result)
}

func TestAutoFix_DirectoryNotFound(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	projectRoot := helpers.CreateTempProject(t)
	missingDir := filepath.Join(projectRoot, "new", "directory")
	err := ErrDirectoryNotFound(missingDir)

	result := rm.autoFix(err)
	assert.NoError(t, result)

	// Verify directory exists
	_, statErr := os.Stat(missingDir)
	assert.NoError(t, statErr)
}

func TestAutoFix_FileNotFound(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	projectRoot := helpers.CreateTempProject(t)
	missingFile := filepath.Join(projectRoot, "new", "file.txt")
	err := ErrFileNotFound(missingFile).WithPath(missingFile)

	result := rm.autoFix(err)
	assert.NoError(t, result)

	// Verify parent directory was created
	parentDir := filepath.Dir(missingFile)
	_, statErr := os.Stat(parentDir)
	assert.NoError(t, statErr)
}

func TestAutoFix_StateNotFound(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	projectRoot := helpers.CreateTempProject(t)
	statePath := filepath.Join(projectRoot, ".cursor", "config", "doplan-state.json")
	err := ErrStateNotFound(statePath)

	result := rm.autoFix(err)
	assert.NoError(t, result)

	// Verify state file was created
	_, statErr := os.Stat(statePath)
	assert.NoError(t, statErr)
}

func TestAutoFix_UnknownError(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	err := NewValidationError("VAL999", "Unknown error code")
	result := rm.autoFix(err)
	assert.Error(t, result)
	assert.Equal(t, err, result)
}

func TestFixDirectoryNotFound(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	projectRoot := helpers.CreateTempProject(t)
	missingDir := filepath.Join(projectRoot, "test", "dir")
	doplanErr := ErrDirectoryNotFound(missingDir)

	result := rm.fixDirectoryNotFound(doplanErr)
	assert.NoError(t, result)

	// Verify directory exists
	_, statErr := os.Stat(missingDir)
	assert.NoError(t, statErr)
}

func TestFixStateNotFound(t *testing.T) {
	handler := NewHandler(nil)
	rm := NewRecoveryManager(handler)

	projectRoot := helpers.CreateTempProject(t)
	statePath := filepath.Join(projectRoot, ".cursor", "config", "doplan-state.json")
	doplanErr := ErrStateNotFound(statePath)

	result := rm.fixStateNotFound(doplanErr)
	assert.NoError(t, result)

	// Verify state file exists
	_, statErr := os.Stat(statePath)
	assert.NoError(t, statErr)

	// Verify it has valid JSON content
	content, readErr := os.ReadFile(statePath)
	require.NoError(t, readErr)
	assert.Contains(t, string(content), "phases")
	assert.Contains(t, string(content), "features")
}
