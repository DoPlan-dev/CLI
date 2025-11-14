package error

import "fmt"

// ErrorCategory represents the type of error
type ErrorCategory string

const (
	ErrorCategoryValidation ErrorCategory = "validation"
	ErrorCategoryIO         ErrorCategory = "io"
	ErrorCategoryNetwork    ErrorCategory = "network"
	ErrorCategoryConfig     ErrorCategory = "config"
	ErrorCategoryGit        ErrorCategory = "git"
	ErrorCategoryGitHub     ErrorCategory = "github"
	ErrorCategoryState      ErrorCategory = "state"
	ErrorCategoryTemplate   ErrorCategory = "template"
	ErrorCategoryUnknown    ErrorCategory = "unknown"
)

// DoPlanError is the base error type
type DoPlanError struct {
	Category    ErrorCategory `json:"category"`
	Code        string        `json:"code"`
	Message     string        `json:"message"`
	Details     string        `json:"details,omitempty"`
	Path        string        `json:"path,omitempty"`
	Suggestion  string        `json:"suggestion,omitempty"`
	Fix         string        `json:"fix,omitempty"`
	Recoverable bool          `json:"recoverable"`
	Cause       error         `json:"-"`
}

// Error implements error interface
func (e *DoPlanError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *DoPlanError) Unwrap() error {
	return e.Cause
}

// WithCause sets the underlying cause
func (e *DoPlanError) WithCause(cause error) *DoPlanError {
	e.Cause = cause
	return e
}

// WithDetails adds details to the error
func (e *DoPlanError) WithDetails(details string) *DoPlanError {
	e.Details = details
	return e
}

// WithPath adds a file path to the error
func (e *DoPlanError) WithPath(path string) *DoPlanError {
	e.Path = path
	return e
}

// WithSuggestion adds a suggestion to the error
func (e *DoPlanError) WithSuggestion(suggestion string) *DoPlanError {
	e.Suggestion = suggestion
	return e
}

// WithFix adds a fix command to the error
func (e *DoPlanError) WithFix(fix string) *DoPlanError {
	e.Fix = fix
	return e
}

// NewValidationError creates a validation error
func NewValidationError(code, message string) *DoPlanError {
	return &DoPlanError{
		Category:    ErrorCategoryValidation,
		Code:        code,
		Message:     message,
		Recoverable: true,
	}
}

// NewIOError creates an IO error
func NewIOError(code, message string) *DoPlanError {
	return &DoPlanError{
		Category:    ErrorCategoryIO,
		Code:        code,
		Message:     message,
		Recoverable: true,
	}
}

// NewConfigError creates a configuration error
func NewConfigError(code, message string) *DoPlanError {
	return &DoPlanError{
		Category:    ErrorCategoryConfig,
		Code:        code,
		Message:     message,
		Recoverable: true,
	}
}

// NewGitHubError creates a GitHub error
func NewGitHubError(code, message string) *DoPlanError {
	return &DoPlanError{
		Category:    ErrorCategoryGitHub,
		Code:        code,
		Message:     message,
		Recoverable: true,
	}
}

// NewStateError creates a state error
func NewStateError(code, message string) *DoPlanError {
	return &DoPlanError{
		Category:    ErrorCategoryState,
		Code:        code,
		Message:     message,
		Recoverable: true,
	}
}

// Common error constructors

// ErrConfigNotFound returns a configuration not found error
func ErrConfigNotFound(path string) *DoPlanError {
	return NewConfigError("CFG001", "Configuration file not found").
		WithPath(path).
		WithSuggestion("Run 'doplan install' to set up DoPlan in this project").
		WithFix("doplan install")
}

// ErrStateNotFound returns a state file not found error
func ErrStateNotFound(path string) *DoPlanError {
	return NewStateError("STA001", "State file not found").
		WithPath(path).
		WithSuggestion("Run 'doplan install' to initialize the project state").
		WithFix("doplan install")
}

// ErrGitHubCLINotFound returns a GitHub CLI not found error
func ErrGitHubCLINotFound() *DoPlanError {
	return NewGitHubError("GH001", "GitHub CLI not found").
		WithSuggestion("Install GitHub CLI:\n  macOS: brew install gh\n  Linux: See https://cli.github.com/manual/installation\n  Windows: winget install GitHub.cli").
		WithFix("brew install gh")
}

// ErrFileNotFound returns a file not found error
func ErrFileNotFound(path string) *DoPlanError {
	return NewIOError("IO001", "File not found").
		WithPath(path).
		WithSuggestion("Check that the file exists and the path is correct")
}

// ErrDirectoryNotFound returns a directory not found error
func ErrDirectoryNotFound(path string) *DoPlanError {
	return NewIOError("IO004", "Directory not found").
		WithPath(path).
		WithSuggestion("Check that the directory exists or run 'doplan install' to create the required structure").
		WithFix("doplan install")
}
