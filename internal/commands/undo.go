package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// UndoLastAction reverts the last DoPlan action using action history
func UndoLastAction() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Load action history
	history, err := loadActionHistory(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to load action history: %w", err)
	}

	if len(history.Actions) == 0 {
		return fmt.Errorf("no actions to undo")
	}

	// Get last action
	lastAction := history.Actions[len(history.Actions)-1]

	fmt.Printf("↩️  Undoing: %s\n", lastAction.Description)
	fmt.Printf("⏰ Time: %s\n\n", lastAction.Timestamp.Format(time.RFC3339))

	// Undo based on action type
	switch lastAction.Type {
	case "file_create":
		return undoFileCreate(projectRoot, lastAction)
	case "file_modify":
		return undoFileModify(projectRoot, lastAction)
	case "git_commit":
		return undoGitCommit(projectRoot, lastAction)
	case "config_change":
		return undoConfigChange(projectRoot, lastAction)
	default:
		return fmt.Errorf("unsupported action type: %s", lastAction.Type)
	}
}

// ActionHistory stores action history
type ActionHistory struct {
	Actions []ActionRecord `json:"actions"`
}

// ActionRecord represents a single action
type ActionRecord struct {
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Timestamp   time.Time              `json:"timestamp"`
	Details     map[string]interface{} `json:"details"`
	Checkpoint  string                 `json:"checkpoint,omitempty"`
}

func loadActionHistory(projectRoot string) (*ActionHistory, error) {
	historyPath := filepath.Join(projectRoot, ".doplan", "state.json")
	data, err := os.ReadFile(historyPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &ActionHistory{Actions: []ActionRecord{}}, nil
		}
		return nil, err
	}

	var state models.State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	// Extract action history from state (if available)
	// For now, return empty history - will be populated as actions are performed
	return &ActionHistory{Actions: []ActionRecord{}}, nil
}

func undoFileCreate(projectRoot string, action ActionRecord) error {
	filePath, ok := action.Details["file"].(string)
	if !ok {
		return fmt.Errorf("invalid action details: missing file path")
	}

	fullPath := filepath.Join(projectRoot, filePath)
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to remove file: %w", err)
	}

	fmt.Printf("✓ Removed file: %s\n", filePath)
	return nil
}

func undoFileModify(projectRoot string, action ActionRecord) error {
	filePath, ok := action.Details["file"].(string)
	if !ok {
		return fmt.Errorf("invalid action details: missing file path")
	}

	checkpoint, ok := action.Details["checkpoint"].(string)
	if !ok || checkpoint == "" {
		return fmt.Errorf("no checkpoint available for file modification")
	}

	// Restore from checkpoint
	checkpointPath := filepath.Join(projectRoot, ".doplan", "checkpoints", checkpoint, filePath)
	if _, err := os.Stat(checkpointPath); err != nil {
		return fmt.Errorf("checkpoint file not found: %w", err)
	}

	data, err := os.ReadFile(checkpointPath)
	if err != nil {
		return fmt.Errorf("failed to read checkpoint: %w", err)
	}

	fullPath := filepath.Join(projectRoot, filePath)
	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return fmt.Errorf("failed to restore file: %w", err)
	}

	fmt.Printf("✓ Restored file: %s\n", filePath)
	return nil
}

func undoGitCommit(projectRoot string, action ActionRecord) error {
	commitHash, ok := action.Details["commit"].(string)
	if !ok {
		return fmt.Errorf("invalid action details: missing commit hash")
	}

	cmd := exec.Command("git", "revert", "--no-edit", commitHash)
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to revert commit: %w\nOutput: %s", err, string(output))
	}

	fmt.Printf("✓ Reverted commit: %s\n", commitHash[:8])
	return nil
}

func undoConfigChange(projectRoot string, action ActionRecord) error {
	// Restore config from checkpoint or previous state
	checkpoint, ok := action.Details["checkpoint"].(string)
	if !ok || checkpoint == "" {
		return fmt.Errorf("no checkpoint available for config change")
	}

	checkpointPath := filepath.Join(projectRoot, ".doplan", "checkpoints", checkpoint, "config.yaml")
	if _, err := os.Stat(checkpointPath); err != nil {
		return fmt.Errorf("checkpoint config not found: %w", err)
	}

	data, err := os.ReadFile(checkpointPath)
	if err != nil {
		return fmt.Errorf("failed to read checkpoint config: %w", err)
	}

	configPath := filepath.Join(projectRoot, ".doplan", "config.yaml")
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to restore config: %w", err)
	}

	fmt.Printf("✓ Restored configuration\n")
	return nil
}

