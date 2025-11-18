package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogLevel represents the log level
type LogLevel string

const (
	LogLevelDebug   LogLevel = "DEBUG"
	LogLevelInfo    LogLevel = "INFO"
	LogLevelWarning LogLevel = "WARNING"
	LogLevelError   LogLevel = "ERROR"
)

// Logger handles logging for DoPlan
type Logger struct {
	projectRoot string
	logDir      string
	logFile     *os.File
	level       LogLevel
}

// NewLogger creates a new logger
func NewLogger(projectRoot string, level LogLevel) (*Logger, error) {
	logDir := filepath.Join(projectRoot, ".doplan", "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	logPath := filepath.Join(logDir, "doplan.log")
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Rotate logs if needed
	if err := rotateLogs(logDir); err != nil {
		// Non-fatal, continue
	}

	return &Logger{
		projectRoot: projectRoot,
		logDir:      logDir,
		logFile:     logFile,
		level:       level,
	}, nil
}

// rotateLogs rotates log files, keeping the last 10
func rotateLogs(logDir string) error {
	logPath := filepath.Join(logDir, "doplan.log")
	
	// Check if log file exists and is large enough to rotate
	info, err := os.Stat(logPath)
	if err != nil || info.Size() < 10*1024*1024 { // 10MB
		return nil
	}

	// Rotate existing logs
	for i := 9; i >= 1; i-- {
		oldPath := filepath.Join(logDir, fmt.Sprintf("doplan.log.%d", i))
		newPath := filepath.Join(logDir, fmt.Sprintf("doplan.log.%d", i+1))
		
		if _, err := os.Stat(oldPath); err == nil {
			os.Rename(oldPath, newPath)
		}
	}

	// Move current log to .1
	currentLog := filepath.Join(logDir, "doplan.log.1")
	return os.Rename(logPath, currentLog)
}

// Log writes a log entry
func (l *Logger) Log(level LogLevel, message string, args ...interface{}) {
	if !l.shouldLog(level) {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf(message, args...)
	logEntry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, formattedMessage)

	l.logFile.WriteString(logEntry)
	l.logFile.Sync()
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...interface{}) {
	l.Log(LogLevelDebug, message, args...)
}

// Info logs an info message
func (l *Logger) Info(message string, args ...interface{}) {
	l.Log(LogLevelInfo, message, args...)
}

// Warning logs a warning message
func (l *Logger) Warning(message string, args ...interface{}) {
	l.Log(LogLevelWarning, message, args...)
}

// Error logs an error message
func (l *Logger) Error(message string, args ...interface{}) {
	l.Log(LogLevelError, message, args...)
}

// shouldLog checks if a log level should be logged
func (l *Logger) shouldLog(level LogLevel) bool {
	levels := map[LogLevel]int{
		LogLevelDebug:   0,
		LogLevelInfo:    1,
		LogLevelWarning: 2,
		LogLevelError:   3,
	}

	return levels[level] >= levels[l.level]
}

// Close closes the log file
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

