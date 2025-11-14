package error

import (
	"os"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	logger := NewLogger(helpers.CreateTempProject(t), LogLevelInfo)
	handler := NewHandler(logger)

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.logger)
}

func TestHandler_Format_DoPlanError(t *testing.T) {
	handler := NewHandler(nil)
	err := NewValidationError("VAL001", "Test error").
		WithDetails("Additional details").
		WithPath("/path/to/file")

	formatted := handler.Format(err)

	assert.Contains(t, formatted, "Test error")
	assert.Contains(t, formatted, "Additional details")
	assert.Contains(t, formatted, "/path/to/file")
	assert.Contains(t, formatted, "VAL001")
}

func TestHandler_Format_GenericError(t *testing.T) {
	handler := NewHandler(nil)
	err := &testError{message: "Generic error"}

	formatted := handler.Format(err)

	assert.Contains(t, formatted, "Generic error")
}

func TestHandler_Categorize_DoPlanError(t *testing.T) {
	handler := NewHandler(nil)
	err := NewValidationError("VAL001", "Test")

	category := handler.categorize(err)
	assert.Equal(t, ErrorCategoryValidation, category)
}

func TestHandler_Categorize_GenericError(t *testing.T) {
	handler := NewHandler(nil)

	tests := []struct {
		err      error
		expected ErrorCategory
	}{
		{&testError{message: "file not found"}, ErrorCategoryIO},
		{&testError{message: "config error"}, ErrorCategoryConfig},
		{&testError{message: "github api error"}, ErrorCategoryGitHub},
		{&testError{message: "state inconsistent"}, ErrorCategoryState},
		{&testError{message: "network timeout"}, ErrorCategoryNetwork},
		{&testError{message: "unknown error"}, ErrorCategoryUnknown},
	}

	for _, tt := range tests {
		category := handler.categorize(tt.err)
		assert.Equal(t, tt.expected, category, "Error: %s", tt.err.Error())
	}
}

func TestHandler_CanRecover(t *testing.T) {
	handler := NewHandler(nil)

	recoverableErr := NewValidationError("VAL001", "Test")
	nonRecoverableErr := &testError{message: "Fatal"}

	assert.True(t, handler.CanRecover(recoverableErr))
	assert.False(t, handler.CanRecover(nonRecoverableErr))
}

func TestHandler_PrintError(t *testing.T) {
	handler := NewHandler(nil)
	err := NewValidationError("VAL001", "Test error").
		WithSuggestion("Run this command").
		WithFix("doplan install")

	// Capture stderr
	oldStderr := os.Stderr
	defer func() { os.Stderr = oldStderr }()

	// This test just ensures it doesn't panic
	handler.PrintError(err)
}

func TestHandler_Handle(t *testing.T) {
	logger := NewLogger(helpers.CreateTempProject(t), LogLevelInfo)
	handler := NewHandler(logger)

	err := NewValidationError("VAL001", "Test error")
	handledErr := handler.Handle(err)

	assert.Error(t, handledErr)
	assert.Contains(t, handledErr.Error(), "Test error")
}

func TestHandler_Handle_Nil(t *testing.T) {
	handler := NewHandler(nil)
	err := handler.Handle(nil)
	assert.NoError(t, err)
}

func TestHandler_Handle_WithSuggestion(t *testing.T) {
	handler := NewHandler(nil)
	err := NewValidationError("VAL001", "Test error").
		WithSuggestion("Run this command").
		WithFix("doplan install")

	handledErr := handler.Handle(err)
	assert.Error(t, handledErr)
	assert.Contains(t, handledErr.Error(), "Test error")
	assert.Contains(t, handledErr.Error(), "Run this command")
}

func TestHandler_formatDoPlanError_WithDetails(t *testing.T) {
	handler := NewHandler(nil)
	err := NewValidationError("VAL001", "Test error").
		WithDetails("Additional details").
		WithPath("/path/to/file")

	formatted := handler.Format(err)
	assert.Contains(t, formatted, "Test error")
	assert.Contains(t, formatted, "Additional details")
	assert.Contains(t, formatted, "/path/to/file")
	assert.Contains(t, formatted, "VAL001")
}

func TestHandler_formatDoPlanError_NoDetails(t *testing.T) {
	handler := NewHandler(nil)
	err := NewValidationError("VAL001", "Test error")

	formatted := handler.Format(err)
	assert.Contains(t, formatted, "Test error")
	assert.Contains(t, formatted, "VAL001")
}

func TestHandler_categorize_NetworkError(t *testing.T) {
	handler := NewHandler(nil)
	err := &testError{message: "network timeout error"}

	category := handler.categorize(err)
	assert.Equal(t, ErrorCategoryNetwork, category)
}

// testError is a simple error for testing
type testError struct {
	message string
}

func (e *testError) Error() string {
	return e.message
}
