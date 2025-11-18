# Error Handling - Comprehensive Error Scenarios & Recovery

This document provides comprehensive error handling strategies for all v0.0.18-beta features, including error scenarios, recovery procedures, and user-friendly error messages.

## 🎯 Error Handling Philosophy

### Principles
1. **Never crash silently** - All errors must be caught and handled
2. **User-friendly messages** - Errors should be understandable by non-technical users
3. **Recovery options** - Always provide a way to recover from errors
4. **Detailed logging** - Log errors with context for debugging
5. **Graceful degradation** - Fall back to simpler functionality when possible

### Error Categories
- **User Errors**: Invalid input, missing requirements
- **System Errors**: File system, network, permissions
- **Configuration Errors**: Invalid config, missing files
- **Migration Errors**: Data corruption, incomplete migration
- **Integration Errors**: IDE setup failures, API failures

## 📋 Error Code System

### Error Code Format
```
[CATEGORY][NUMBER] - [Description]

Categories:
- USR = User Error (1xxx)
- SYS = System Error (2xxx)
- CFG = Configuration Error (3xxx)
- MIG = Migration Error (4xxx)
- INT = Integration Error (5xxx)
- TUI = TUI Error (6xxx)
```

### Error Code Registry

#### User Errors (USR1xxx)
```
USR1001 - Invalid project name format
USR1002 - Project name too short/long
USR1003 - Invalid GitHub URL
USR1004 - GitHub repository not accessible
USR1005 - IDE not selected
USR1006 - Template not selected
USR1007 - Required field missing
```

#### System Errors (SYS2xxx)
```
SYS2001 - File system permission denied
SYS2002 - Disk space insufficient
SYS2003 - Directory creation failed
SYS2004 - File read failed
SYS2005 - File write failed
SYS2006 - Network connection failed
SYS2007 - Process execution failed
```

#### Configuration Errors (CFG3xxx)
```
CFG3001 - Config file not found
CFG3002 - Config file invalid format
CFG3003 - Config file missing required field
CFG3004 - Config validation failed
CFG3005 - State file corrupted
CFG3006 - Dashboard JSON invalid
```

#### Migration Errors (MIG4xxx)
```
MIG4001 - Old structure not detected
MIG4002 - Backup creation failed
MIG4003 - Config migration failed
MIG4004 - Folder migration failed
MIG4005 - Reference update failed
MIG4006 - Migration rollback failed
MIG4007 - Migration verification failed
```

#### Integration Errors (INT5xxx)
```
INT5001 - IDE not detected
INT5002 - IDE setup failed
INT5003 - Symlink creation failed
INT5004 - IDE config generation failed
INT5005 - IDE verification failed
INT5006 - GitHub API authentication failed
INT5007 - GitHub repository access denied
```

#### TUI Errors (TUI6xxx)
```
TUI6001 - Terminal too small
TUI6002 - TUI initialization failed
TUI6003 - Dashboard load failed
TUI6004 - Wizard navigation error
TUI6005 - Input validation failed
```

## 🔧 Error Handling Implementation

### Error Type Definition
```go
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
```

### Error Factory Functions
```go
// User Errors
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

func ErrGitHubURLInvalid(url string) *DoPlanError {
    return &DoPlanError{
        Code:    "USR1003",
        Message: "Invalid GitHub repository URL",
        Details: fmt.Sprintf("URL '%s' is not a valid GitHub repository", url),
        Fix: `Please provide a valid GitHub URL in one of these formats:
- https://github.com/username/repository
- git@github.com:username/repository.git
- username/repository

Make sure the repository exists and you have access to it.`,
        Timestamp: time.Now(),
        Context: map[string]interface{}{
            "url": url,
        },
    }
}

// System Errors
func ErrPermissionDenied(path string) *DoPlanError {
    return &DoPlanError{
        Code:    "SYS2001",
        Message: "Permission denied",
        Details: fmt.Sprintf("Cannot access '%s' - permission denied", path),
        Fix: `Please check:
1. You have read/write permissions for this directory
2. The directory is not locked by another process
3. You're not running in a restricted environment

Try running with appropriate permissions or changing the directory.`,
        Timestamp: time.Now(),
        Context: map[string]interface{}{
            "path": path,
        },
    }
}

// Configuration Errors
func ErrConfigNotFound(path string) *DoPlanError {
    return &DoPlanError{
        Code:    "CFG3001",
        Message: "Configuration file not found",
        Details: fmt.Sprintf("Config file not found at '%s'", path),
        Fix: `The DoPlan configuration file is missing. This usually means:
1. DoPlan hasn't been initialized in this project
2. The configuration was deleted

To fix:
- Run 'doplan' to initialize the project
- Or restore from backup if you have one`,
        Timestamp: time.Now(),
        Context: map[string]interface{}{
            "path": path,
        },
    }
}

// Migration Errors
func ErrMigrationBackupFailed(cause error) *DoPlanError {
    return &DoPlanError{
        Code:    "MIG4002",
        Message: "Failed to create migration backup",
        Details: fmt.Sprintf("Backup creation failed: %v", cause),
        Fix: `Migration cannot proceed without a backup. Please:

1. Check disk space (need at least 100 MB free)
2. Verify write permissions in the project directory
3. Ensure no other process is locking files
4. Try again, or manually create a backup first`,
        Timestamp: time.Now(),
        Cause:    cause,
    }
}

// Integration Errors
func ErrIDESetupFailed(ide string, cause error) *DoPlanError {
    return &DoPlanError{
        Code:    "INT5002",
        Message: fmt.Sprintf("Failed to set up %s integration", ide),
        Details: fmt.Sprintf("IDE setup failed: %v", cause),
        Fix: fmt.Sprintf(`Failed to set up %s integration. Please:

1. Verify %s is installed and accessible
2. Check file permissions in the project directory
3. Ensure .doplan/ai/ directory exists
4. Try running setup again
5. Check the guide: .doplan/guides/%s_setup.md`, ide, ide, ide),
        Timestamp: time.Now(),
        Cause:    cause,
        Context: map[string]interface{}{
            "ide": ide,
        },
    }
}

// TUI Errors
func ErrTerminalTooSmall(width, height, minWidth, minHeight int) *DoPlanError {
    return &DoPlanError{
        Code:    "TUI6001",
        Message: "Terminal window is too small",
        Details: fmt.Sprintf("Terminal size: %dx%d, Minimum required: %dx%d", width, height, minWidth, minHeight),
        Fix: `Please resize your terminal window:
- Minimum size: 80x24 characters
- Recommended: 120x30 characters

Or use the CLI commands instead:
  doplan dashboard
  doplan config show`,
        Timestamp: time.Now(),
        Context: map[string]interface{}{
            "width":      width,
            "height":     height,
            "minWidth":   minWidth,
            "minHeight":  minHeight,
        },
    }
}
```

## 🛡️ Error Recovery Strategies

### Recovery Strategy 1: User Input Validation

**Scenario**: Invalid project name entered

**Error**: `USR1001 - Invalid project name format`

**Recovery**:
```go
func handleProjectNameInput(input string) (string, error) {
    // Validate
    if err := validateProjectName(input); err != nil {
        // Show error in TUI
        showError(err)
        
        // Suggest correction
        suggestion := suggestProjectName(input)
        if suggestion != "" {
            showSuggestion(fmt.Sprintf("Did you mean: %s?", suggestion))
        }
        
        // Allow retry
        return "", err
    }
    
    return input, nil
}
```

### Recovery Strategy 2: File System Errors

**Scenario**: Permission denied when creating directory

**Error**: `SYS2001 - Permission denied`

**Recovery**:
```go
func createDirectoryWithRetry(path string) error {
    maxRetries := 3
    
    for i := 0; i < maxRetries; i++ {
        err := os.MkdirAll(path, 0755)
        if err == nil {
            return nil
        }
        
        // Check error type
        if os.IsPermission(err) {
            // Suggest fix
            showError(ErrPermissionDenied(path))
            showSuggestion("Try running with appropriate permissions")
            return err
        }
        
        // Retry for transient errors
        if i < maxRetries-1 {
            time.Sleep(time.Second * time.Duration(i+1))
            continue
        }
        
        return err
    }
    
    return nil
}
```

### Recovery Strategy 3: Migration Rollback

**Scenario**: Migration fails mid-process

**Error**: `MIG4004 - Folder migration failed`

**Recovery**:
```go
func handleMigrationFailure(migrator *migration.Migrator, backupPath string) error {
    // Show error
    showError(ErrMigrationFolderFailed("", "", nil))
    
    // Offer rollback
    showDialog(Dialog{
        Title:   "Migration Failed",
        Message: "Migration encountered an error. Would you like to rollback?",
        Options: []DialogOption{
            {Label: "Rollback", Action: func() {
                if err := migrator.Rollback(backupPath); err != nil {
                    showError(ErrMigrationRollbackFailed(err))
                } else {
                    showSuccess("Migration rolled back successfully")
                }
            }},
            {Label: "Keep Partial", Action: func() {
                showWarning("Partial migration kept. Manual cleanup may be needed.")
            }},
            {Label: "Retry", Action: func() {
                // Retry migration
            }},
        },
    })
    
    return nil
}
```

### Recovery Strategy 4: Config Fallback

**Scenario**: Config file is corrupted

**Error**: `CFG3002 - Config file invalid`

**Recovery**:
```go
func loadConfigWithFallback(path string) (*Config, error) {
    // Try to load config
    config, err := loadConfig(path)
    if err == nil {
        return config, nil
    }
    
    // Try backup
    backupPath := path + ".backup"
    if config, err := loadConfig(backupPath); err == nil {
        showWarning("Using backup config file")
        return config, nil
    }
    
    // Generate default config
    showWarning("Generating default configuration")
    defaultConfig := generateDefaultConfig()
    
    // Save default config
    if err := saveConfig(path, defaultConfig); err != nil {
        return nil, err
    }
    
    return defaultConfig, nil
}
```

### Recovery Strategy 5: Dashboard Fallback

**Scenario**: Dashboard JSON is corrupted

**Error**: `CFG3006 - Dashboard JSON invalid`

**Recovery**:
```go
func loadDashboardWithFallback(projectRoot string) (*Dashboard, error) {
    // Try JSON first
    jsonPath := filepath.Join(projectRoot, ".doplan", "dashboard.json")
    dashboard, err := loadDashboardJSON(jsonPath)
    if err == nil {
        return dashboard, nil
    }
    
    // Fallback to markdown
    showWarning("Dashboard JSON invalid, using markdown format")
    mdPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
    dashboard, err = loadDashboardMarkdown(mdPath)
    if err == nil {
        return dashboard, nil
    }
    
    // Generate new dashboard
    showWarning("Regenerating dashboard")
    return generateDashboard(projectRoot)
}
```

## 📊 Error Display in TUI

### Error Display Component
```go
func renderError(err *DoPlanError, width int) string {
    style := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#ef4444")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#ef4444")).
        Padding(1, 2).
        Width(width - 4)
    
    content := fmt.Sprintf(
        "❌ %s\n\n%s\n\n%s",
        err.Message,
        err.Details,
        err.Fix,
    )
    
    if err.Code != "" {
        content += fmt.Sprintf("\n\nError Code: %s", err.Code)
    }
    
    return style.Render(content)
}
```

## 📝 Error Logging

### Log Format
```go
type ErrorLog struct {
    Timestamp time.Time              `json:"timestamp"`
    Code      string                 `json:"code"`
    Message   string                 `json:"message"`
    Details   string                 `json:"details"`
    Context   map[string]interface{} `json:"context"`
    Cause     string                 `json:"cause,omitempty"`
    Stack     string                 `json:"stack,omitempty"`
}

func logError(err *DoPlanError) {
    log := ErrorLog{
        Timestamp: err.Timestamp,
        Code:      err.Code,
        Message:   err.Message,
        Details:   err.Details,
        Context:   err.Context,
    }
    
    if err.Cause != nil {
        log.Cause = err.Cause.Error()
        log.Stack = string(debug.Stack())
    }
    
    // Write to log file
    logPath := filepath.Join(projectRoot, ".doplan", "logs", "doplan.log")
    writeLog(logPath, log)
}
```

## 🔄 Error Recovery Workflows

### Workflow 1: Wizard Error Recovery
```
User Input → Validate → Error?
                          ↓ Yes
                    Show Error
                    Show Suggestion
                    Allow Retry
                    ↓
                  User Corrects
                    ↓
                  Validate Again
```

### Workflow 2: Migration Error Recovery
```
Start Migration → Create Backup → Error?
                                    ↓ Yes
                              Show Error
                              Offer Rollback
                              ↓
                        User Chooses
                        ↓
                    Rollback / Retry / Keep
```

## 📋 Error Handling Checklist

### For Each Feature
- [ ] Define error codes
- [ ] Create error factory functions
- [ ] Implement error recovery
- [ ] Add error logging
- [ ] Create user-friendly messages
- [ ] Test error scenarios
- [ ] Document error codes
- [ ] Add error handling tests

## 🚨 Critical Error Scenarios

### Scenario 1: Data Loss Prevention
**Error**: Migration fails after partial completion

**Recovery**:
1. Always create backup first
2. Verify backup before proceeding
3. Use transactions where possible
4. Rollback on any error
5. Never delete old data until verified

### Scenario 2: Corrupted State
**Error**: State file is corrupted

**Recovery**:
1. Try to load backup state
2. Regenerate from project files
3. Use last known good state
4. Warn user about potential data loss

### Scenario 3: Network Failure
**Error**: GitHub API unavailable

**Recovery**:
1. Retry with exponential backoff
2. Cache responses when possible
3. Work offline when possible
4. Queue operations for later

## 📚 Error Code Reference

### Quick Reference Table
| Code | Category | Severity | User Action |
|------|----------|----------|-------------|
| USR1xxx | User Input | Low | Correct input |
| SYS2xxx | System | High | Check permissions/space |
| CFG3xxx | Config | Medium | Fix config or regenerate |
| MIG4xxx | Migration | High | Rollback or manual fix |
| INT5xxx | Integration | Medium | Retry or manual setup |
| TUI6xxx | TUI | Low | Resize terminal or use CLI |

