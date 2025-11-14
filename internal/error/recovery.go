package error

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// RecoveryStrategy defines how to recover from errors
type RecoveryStrategy string

const (
	RecoveryAutoFix  RecoveryStrategy = "autofix"
	RecoveryPrompt   RecoveryStrategy = "prompt"
	RecoveryRollback RecoveryStrategy = "rollback"
	RecoverySkip     RecoveryStrategy = "skip"
)

// RecoveryManager handles error recovery
type RecoveryManager struct {
	handler *Handler
}

// NewRecoveryManager creates a new recovery manager
func NewRecoveryManager(handler *Handler) *RecoveryManager {
	return &RecoveryManager{
		handler: handler,
	}
}

// Recover attempts to recover from an error
func (rm *RecoveryManager) Recover(err error, strategy RecoveryStrategy) error {
	if !rm.handler.CanRecover(err) {
		return err
	}

	switch strategy {
	case RecoveryAutoFix:
		return rm.autoFix(err)
	case RecoveryPrompt:
		return rm.promptRecovery(err)
	case RecoveryRollback:
		return rm.rollback(err)
	case RecoverySkip:
		return nil // Skip the error
	default:
		return err
	}
}

// autoFix attempts to automatically fix the error
func (rm *RecoveryManager) autoFix(err error) error {
	doplanErr, ok := err.(*DoPlanError)
	if !ok {
		return err
	}

	switch doplanErr.Code {
	case "CFG001": // Config not found
		return rm.fixConfigNotFound(doplanErr)
	case "IO004": // Directory not found
		return rm.fixDirectoryNotFound(doplanErr)
	case "STA001": // State not found
		return rm.fixStateNotFound(doplanErr)
	case "IO001": // File not found - try to create parent directory
		if doplanErr.Path != "" {
			dir := filepath.Dir(doplanErr.Path)
			if err := os.MkdirAll(dir, 0755); err == nil {
				return nil
			}
		}
		return err
	default:
		return err
	}
}

// fixConfigNotFound attempts to fix missing config by running doplan install
func (rm *RecoveryManager) fixConfigNotFound(err *DoPlanError) error {
	// Try to run doplan install
	cmd := exec.Command("doplan", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if runErr := cmd.Run(); runErr != nil {
		return fmt.Errorf("auto-fix failed: %w (original error: %v)", runErr, err)
	}

	return nil
}

// fixDirectoryNotFound creates the missing directory
func (rm *RecoveryManager) fixDirectoryNotFound(err *DoPlanError) error {
	if err.Path != "" {
		if mkErr := os.MkdirAll(err.Path, 0755); mkErr == nil {
			return nil
		}
	}
	return err
}

// fixStateNotFound initializes a basic state file
func (rm *RecoveryManager) fixStateNotFound(err *DoPlanError) error {
	if err.Path != "" {
		// Create directory if needed
		dir := filepath.Dir(err.Path)
		if mkErr := os.MkdirAll(dir, 0755); mkErr != nil {
			return err
		}

		// Create empty state file
		emptyState := []byte(`{"phases":[],"features":[],"progress":{"overall":0}}`)
		if writeErr := os.WriteFile(err.Path, emptyState, 0644); writeErr == nil {
			return nil
		}
	}
	return err
}

// promptRecovery prompts the user for recovery action
func (rm *RecoveryManager) promptRecovery(err error) error {
	doplanErr, ok := err.(*DoPlanError)
	if !ok {
		return err
	}

	fmt.Fprintf(os.Stderr, "\n⚠️  Error occurred: %s\n", doplanErr.Message)
	if doplanErr.Suggestion != "" {
		fmt.Fprintf(os.Stderr, "💡 Suggestion: %s\n", doplanErr.Suggestion)
	}
	if doplanErr.Fix != "" {
		fmt.Fprintf(os.Stderr, "🔧 Fix command: %s\n", doplanErr.Fix)
	}

	fmt.Fprint(os.Stderr, "\nAttempt auto-fix? (y/n): ")
	var response string
	fmt.Scanln(&response)

	if response == "y" || response == "Y" {
		return rm.autoFix(err)
	}

	return err
}

// rollback attempts to rollback to last checkpoint
func (rm *RecoveryManager) rollback(err error) error {
	// This would require checkpoint integration
	// For now, just return the error
	return fmt.Errorf("rollback not yet implemented: %w", err)
}
