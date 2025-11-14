package error

import (
	"fmt"
	"os"
	"strings"
)

// Handler processes and formats errors
type Handler struct {
	logger *Logger
}

// NewHandler creates a new error handler
func NewHandler(logger *Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// Handle processes an error and returns a formatted error
func (h *Handler) Handle(err error) error {
	if err == nil {
		return nil
	}

	// Categorize error
	category := h.categorize(err)

	// Log error
	if h.logger != nil {
		h.logger.Log(err, category)
	}

	// Format for user
	formatted := h.Format(err)

	// Suggest fix if recoverable
	if doplanErr, ok := err.(*DoPlanError); ok && doplanErr.Recoverable {
		if doplanErr.Suggestion != "" {
			formatted += "\n\n💡 " + doplanErr.Suggestion
		}
		if doplanErr.Fix != "" {
			formatted += "\n\n🔧 Fix: " + doplanErr.Fix
		}
	}

	return fmt.Errorf("%s", formatted)
}

// Format formats error for display
func (h *Handler) Format(err error) string {
	if doplanErr, ok := err.(*DoPlanError); ok {
		return h.formatDoPlanError(doplanErr)
	}
	return h.formatGenericError(err)
}

// formatDoPlanError formats a DoPlanError
func (h *Handler) formatDoPlanError(err *DoPlanError) string {
	var sb strings.Builder

	// Error icon and message
	sb.WriteString(fmt.Sprintf("❌ %s\n", err.Message))

	// Details
	if err.Details != "" {
		sb.WriteString(fmt.Sprintf("\n%s", err.Details))
	}

	// Path
	if err.Path != "" {
		sb.WriteString(fmt.Sprintf("\n\n📍 Path: %s", err.Path))
	}

	// Code
	sb.WriteString(fmt.Sprintf("\n📋 Code: %s", err.Code))

	return sb.String()
}

// formatGenericError formats a generic error
func (h *Handler) formatGenericError(err error) string {
	return fmt.Sprintf("❌ %s", err.Error())
}

// categorize determines the error category
func (h *Handler) categorize(err error) ErrorCategory {
	if doplanErr, ok := err.(*DoPlanError); ok {
		return doplanErr.Category
	}

	// Try to categorize based on error message
	errMsg := strings.ToLower(err.Error())

	switch {
	case strings.Contains(errMsg, "file") || strings.Contains(errMsg, "directory"):
		return ErrorCategoryIO
	case strings.Contains(errMsg, "config") || strings.Contains(errMsg, "configuration"):
		return ErrorCategoryConfig
	case strings.Contains(errMsg, "github") || strings.Contains(errMsg, "git"):
		return ErrorCategoryGitHub
	case strings.Contains(errMsg, "state"):
		return ErrorCategoryState
	case strings.Contains(errMsg, "network") || strings.Contains(errMsg, "http"):
		return ErrorCategoryNetwork
	default:
		return ErrorCategoryUnknown
	}
}

// CanRecover checks if an error is recoverable
func (h *Handler) CanRecover(err error) bool {
	if doplanErr, ok := err.(*DoPlanError); ok {
		return doplanErr.Recoverable
	}
	return false
}

// PrintError prints an error to stderr with formatting
func (h *Handler) PrintError(err error) {
	formatted := h.Format(err)
	fmt.Fprintf(os.Stderr, "%s\n", formatted)

	if doplanErr, ok := err.(*DoPlanError); ok {
		if doplanErr.Suggestion != "" {
			fmt.Fprintf(os.Stderr, "\n💡 %s\n", doplanErr.Suggestion)
		}
		if doplanErr.Fix != "" {
			fmt.Fprintf(os.Stderr, "🔧 Fix: %s\n", doplanErr.Fix)
		}
	}
}

