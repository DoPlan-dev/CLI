package error

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// LogLevel represents logging verbosity
type LogLevel string

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warning"
	LogLevelError   LogLevel = "error"
)

// Logger handles error logging
type Logger struct {
	logFile   string
	level     LogLevel
	logs      []*ErrorLog
	buffer    []*ErrorLog
	bufferMu  sync.Mutex
	flushSize int
	flushTime time.Duration
	lastFlush time.Time
}

// ErrorLog represents a logged error
type ErrorLog struct {
	Timestamp time.Time              `json:"timestamp"`
	Category  ErrorCategory          `json:"category"`
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Details   string                 `json:"details,omitempty"`
	Path      string                 `json:"path,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
	Stack     string                 `json:"stack,omitempty"`
}

// NewLogger creates a new error logger
func NewLogger(projectRoot string, level LogLevel) *Logger {
	logDir := filepath.Join(projectRoot, ".doplan")
	logFile := filepath.Join(logDir, "errors.json")

	return &Logger{
		logFile:   logFile,
		level:     level,
		logs:      []*ErrorLog{},
		buffer:    []*ErrorLog{},
		flushSize: 10,
		flushTime: 100 * time.Millisecond,
		lastFlush: time.Now(),
	}
}

// Log logs an error
func (l *Logger) Log(err error, category ErrorCategory) {
	logEntry := &ErrorLog{
		Timestamp: time.Now(),
		Category:  category,
		Message:   err.Error(),
	}

	// Extract DoPlanError details if available
	if doplanErr, ok := err.(*DoPlanError); ok {
		logEntry.Code = doplanErr.Code
		logEntry.Details = doplanErr.Details
		logEntry.Path = doplanErr.Path
	}

	l.logs = append(l.logs, logEntry)

	// Add to buffer for batch writing
	l.bufferMu.Lock()
	l.buffer = append(l.buffer, logEntry)
	shouldFlush := len(l.buffer) >= l.flushSize || time.Since(l.lastFlush) >= l.flushTime
	l.bufferMu.Unlock()

	// Flush if needed
	if shouldFlush {
		l.flush()
	}
}

// LogWithContext logs an error with additional context
func (l *Logger) LogWithContext(err error, category ErrorCategory, context map[string]interface{}) {
	logEntry := &ErrorLog{
		Timestamp: time.Now(),
		Category:  category,
		Message:   err.Error(),
		Context:   context,
	}

	// Extract DoPlanError details if available
	if doplanErr, ok := err.(*DoPlanError); ok {
		logEntry.Code = doplanErr.Code
		logEntry.Details = doplanErr.Details
		logEntry.Path = doplanErr.Path
	}

	l.logs = append(l.logs, logEntry)

	// Add to buffer for batch writing
	l.bufferMu.Lock()
	l.buffer = append(l.buffer, logEntry)
	shouldFlush := len(l.buffer) >= l.flushSize || time.Since(l.lastFlush) >= l.flushTime
	l.bufferMu.Unlock()

	// Flush if needed
	if shouldFlush {
		l.flush()
	}
}

// GetLogs returns logs since a given time
func (l *Logger) GetLogs(since time.Time) ([]*ErrorLog, error) {
	// Load logs from file
	logs, err := l.load()
	if err != nil {
		return nil, err
	}

	// Filter by time
	filtered := []*ErrorLog{}
	for _, log := range logs {
		if log.Timestamp.After(since) {
			filtered = append(filtered, log)
		}
	}

	return filtered, nil
}

// ClearLogs clears all logs
func (l *Logger) ClearLogs() error {
	l.logs = []*ErrorLog{}
	l.bufferMu.Lock()
	l.buffer = []*ErrorLog{}
	l.bufferMu.Unlock()
	return l.persist()
}

// flush writes buffered logs to file
func (l *Logger) flush() {
	l.bufferMu.Lock()
	if len(l.buffer) == 0 {
		l.bufferMu.Unlock()
		return
	}

	// Copy buffer and clear it
	bufferCopy := make([]*ErrorLog, len(l.buffer))
	copy(bufferCopy, l.buffer)
	l.buffer = []*ErrorLog{}
	l.lastFlush = time.Now()
	l.bufferMu.Unlock()

	// Add buffered logs to main logs and persist
	l.logs = append(l.logs, bufferCopy...)
	l.persist()
}

// persist writes logs to file
func (l *Logger) persist() error {
	// Ensure directory exists
	dir := filepath.Dir(l.logFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	// Load existing logs
	existing, _ := l.load()

	// Merge with new logs
	allLogs := append(existing, l.logs...)

	// Keep only last 1000 logs
	if len(allLogs) > 1000 {
		allLogs = allLogs[len(allLogs)-1000:]
	}

	// Write to file
	data, err := json.MarshalIndent(allLogs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal logs: %w", err)
	}

	if err := os.WriteFile(l.logFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write log file: %w", err)
	}

	return nil
}

// load reads logs from file
func (l *Logger) load() ([]*ErrorLog, error) {
	if _, err := os.Stat(l.logFile); os.IsNotExist(err) {
		return []*ErrorLog{}, nil
	}

	data, err := os.ReadFile(l.logFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read log file: %w", err)
	}

	var logs []*ErrorLog
	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal logs: %w", err)
	}

	return logs, nil
}
