package error

import (
	"fmt"
	"time"
)

// DoPlanError represents a DoPlan-specific error
type DoPlanError struct {
	Code      string
	Message   string
	Details   string
	Fix       string
	Timestamp time.Time
	Cause     error
	Context   map[string]interface{}
}

// Error implements the error interface
func (e *DoPlanError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// UserMessage returns a user-friendly error message
func (e *DoPlanError) UserMessage() string {
	msg := fmt.Sprintf("%s\n\n", e.Message)
	if e.Details != "" {
		msg += fmt.Sprintf("Details: %s\n", e.Details)
	}
	if e.Fix != "" {
		msg += fmt.Sprintf("\nHow to fix:\n%s\n", e.Fix)
	}
	if e.Code != "" {
		msg += fmt.Sprintf("\nError Code: %s\n", e.Code)
	}
	return msg
}

// WithCause adds a cause to the error
func (e *DoPlanError) WithCause(cause error) *DoPlanError {
	e.Cause = cause
	return e
}

// WithContext adds context to the error
func (e *DoPlanError) WithContext(key string, value interface{}) *DoPlanError {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
	return e
}

// Error factory functions
func ErrInvalidProjectName(name string) *DoPlanError {
	return &DoPlanError{
		Code:    "USR1001",
		Message: "Invalid project name format",
		Details: fmt.Sprintf("Project name '%s' contains invalid characters", name),
		Fix: `Project names must:
- Use only lowercase letters, numbers, and hyphens
- Be 3-50 characters long
- Not contain spaces or special characters

Example: "my-awesome-project"`,
		Timestamp: time.Now(),
		Context: map[string]interface{}{
			"name": name,
		},
	}
}

// Add more error factory functions as needed

