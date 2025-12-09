package timetracker

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Entry represents a tracked command execution window.
type Entry struct {
	Project        string            `json:"project"`
	Command        string            `json:"command"`
	Phase          string            `json:"phase"`
	Feature        string            `json:"feature,omitempty"` // Feature name for feature-specific tracking
	TaskID         string            `json:"task_id,omitempty"` // Task ID if applicable
	Args           []string          `json:"args,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	StartedAt      time.Time         `json:"started_at"`
	EndedAt        time.Time         `json:"ended_at"`
	DurationMillis int64             `json:"duration_ms"`
	Status         string            `json:"status"`
	Error          string            `json:"error,omitempty"`
}

// Tracker handles writing time tracking entries to disk.
type Tracker struct {
	logPath string
	entry   *Entry
}

// New creates a tracker that writes to .do/system/time-tracker.jsonl inside projectPath.
func New(projectPath string) (*Tracker, error) {
	if projectPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		projectPath = cwd
	}

	logDir := filepath.Join(projectPath, ".do", "system")
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		return nil, err
	}

	return &Tracker{
		logPath: filepath.Join(logDir, "time-tracker.jsonl"),
	}, nil
}

// Start begins tracking for a command/phase pair.
func (t *Tracker) Start(projectPath, phase, command string, args []string, metadata map[string]string) {
	// Extract feature name from metadata if available
	feature := ""
	if metadata != nil {
		if f, ok := metadata["feature"]; ok {
			feature = f
		}
	}

	// Extract task ID from metadata if available
	taskID := ""
	if metadata != nil {
		if tid, ok := metadata["task_id"]; ok {
			taskID = tid
		}
	}

	t.entry = &Entry{
		Project:   projectPath,
		Command:   command,
		Phase:     phase,
		Feature:   feature,
		TaskID:    taskID,
		Args:      args,
		Metadata:  metadata,
		StartedAt: time.Now().UTC(),
	}
}

// UpdateMetadata updates the metadata of the current entry
func (t *Tracker) UpdateMetadata(key, value string) {
	if t.entry == nil {
		return
	}
	if t.entry.Metadata == nil {
		t.entry.Metadata = make(map[string]string)
	}
	t.entry.Metadata[key] = value

	// Also update direct fields if applicable
	if key == "feature" {
		t.entry.Feature = value
	}
	if key == "task_id" {
		t.entry.TaskID = value
	}
}

// Stop finalizes the current entry and writes it to disk.
func (t *Tracker) Stop(success bool, err error) error {
	if t.entry == nil {
		return nil
	}

	t.entry.EndedAt = time.Now().UTC()
	t.entry.DurationMillis = t.entry.EndedAt.Sub(t.entry.StartedAt).Milliseconds()
	if success {
		t.entry.Status = "completed"
	} else {
		t.entry.Status = "failed"
	}
	if err != nil {
		t.entry.Error = err.Error()
	}

	data, marshalErr := json.Marshal(t.entry)
	if marshalErr != nil {
		return marshalErr
	}

	f, openErr := os.OpenFile(t.logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if openErr != nil {
		return openErr
	}
	defer f.Close()

	if _, writeErr := f.Write(append(data, '\n')); writeErr != nil {
		return writeErr
	}

	t.entry = nil
	return nil
}

// GetFeatureTime returns total time spent on a specific feature
func GetFeatureTime(projectPath, featureName string) (time.Duration, error) {
	entries, err := readEntries(projectPath)
	if err != nil {
		return 0, err
	}

	var totalDuration time.Duration
	for _, entry := range entries {
		if entry.Feature == featureName && entry.Status == "completed" {
			totalDuration += time.Duration(entry.DurationMillis) * time.Millisecond
		}
	}

	return totalDuration, nil
}

// GetPhaseTime returns total time spent in a specific phase
func GetPhaseTime(projectPath, phaseName string) (time.Duration, error) {
	entries, err := readEntries(projectPath)
	if err != nil {
		return 0, err
	}

	var totalDuration time.Duration
	for _, entry := range entries {
		if entry.Phase == phaseName && entry.Status == "completed" {
			totalDuration += time.Duration(entry.DurationMillis) * time.Millisecond
		}
	}

	return totalDuration, nil
}

// readEntries reads all time tracking entries from the log file
func readEntries(projectPath string) ([]Entry, error) {
	if projectPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		projectPath = cwd
	}

	logPath := filepath.Join(projectPath, ".do", "system", "time-tracker.jsonl")

	// Check if file exists
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		return []Entry{}, nil
	}

	data, err := os.ReadFile(logPath)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var entry Entry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			continue // Skip invalid entries
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
