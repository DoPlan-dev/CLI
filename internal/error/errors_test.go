package error

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("VAL001", "Test validation error")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryValidation, err.Category)
	assert.Equal(t, "VAL001", err.Code)
	assert.Equal(t, "Test validation error", err.Message)
	assert.True(t, err.Recoverable)
}

func TestNewIOError(t *testing.T) {
	err := NewIOError("IO001", "File not found")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryIO, err.Category)
	assert.Equal(t, "IO001", err.Code)
	assert.Equal(t, "File not found", err.Message)
}

func TestNewConfigError(t *testing.T) {
	err := NewConfigError("CFG001", "Config not found")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryConfig, err.Category)
	assert.Equal(t, "CFG001", err.Code)
}

func TestDoPlanError_Error(t *testing.T) {
	err := &DoPlanError{
		Code:    "TEST001",
		Message: "Test error",
		Details: "Additional details",
	}

	errorMsg := err.Error()
	assert.Contains(t, errorMsg, "Test error")
	assert.Contains(t, errorMsg, "Additional details")
}

func TestDoPlanError_WithCause(t *testing.T) {
	originalErr := errors.New("original error")
	err := NewValidationError("VAL001", "Test error").WithCause(originalErr)

	assert.Equal(t, originalErr, err.Cause)
	assert.Equal(t, originalErr, err.Unwrap())
}

func TestDoPlanError_WithDetails(t *testing.T) {
	err := NewValidationError("VAL001", "Test error").WithDetails("More info")

	assert.Equal(t, "More info", err.Details)
}

func TestDoPlanError_WithPath(t *testing.T) {
	err := NewIOError("IO001", "File error").WithPath("/path/to/file")

	assert.Equal(t, "/path/to/file", err.Path)
}

func TestDoPlanError_WithSuggestion(t *testing.T) {
	err := NewConfigError("CFG001", "Config error").WithSuggestion("Run doplan install")

	assert.Equal(t, "Run doplan install", err.Suggestion)
}

func TestDoPlanError_WithFix(t *testing.T) {
	err := NewConfigError("CFG001", "Config error").WithFix("doplan install")

	assert.Equal(t, "doplan install", err.Fix)
}

func TestErrConfigNotFound(t *testing.T) {
	err := ErrConfigNotFound("/path/to/config")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryConfig, err.Category)
	assert.Equal(t, "CFG001", err.Code)
	assert.Equal(t, "/path/to/config", err.Path)
	assert.Contains(t, err.Suggestion, "doplan install")
}

func TestErrStateNotFound(t *testing.T) {
	err := ErrStateNotFound("/path/to/state")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryState, err.Category)
	assert.Equal(t, "STA001", err.Code)
}

func TestErrGitHubCLINotFound(t *testing.T) {
	err := ErrGitHubCLINotFound()

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryGitHub, err.Category)
	assert.Equal(t, "GH001", err.Code)
	assert.Contains(t, err.Suggestion, "brew install gh")
}

func TestErrFileNotFound(t *testing.T) {
	err := ErrFileNotFound("/path/to/file")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryIO, err.Category)
	assert.Equal(t, "IO001", err.Code)
	assert.Equal(t, "/path/to/file", err.Path)
}

func TestErrDirectoryNotFound(t *testing.T) {
	err := ErrDirectoryNotFound("/path/to/dir")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorCategoryIO, err.Category)
	assert.Equal(t, "IO004", err.Code)
	assert.Contains(t, err.Suggestion, "doplan install")
}

