package timetracker

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Entry represents a tracked command execution window.
type Entry struct {
	Project        string            `json:"project"`
	Command        string            `json:"command"`
	Phase          string            `json:"phase"`
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
	t.entry = &Entry{
		Project:   projectPath,
		Command:   command,
		Phase:     phase,
		Args:      args,
		Metadata:  metadata,
		StartedAt: time.Now().UTC(),
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
