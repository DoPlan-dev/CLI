# DoPlan v0.0.18-beta Advanced Usage Guide

This guide covers advanced usage patterns, workflows, and techniques for implementing v0.0.18-beta.

## 🎯 Development Workflows

### Workflow 1: Implementing a New Wizard

**Step-by-step process:**

1. **Plan the wizard flow**
   - Review TUI wizard flow documentation
   - Create flow diagram
   - Define state machine
   - List all screens

2. **Generate base code**
   - Use code generator or template
   - Copy `templates/wizard_base.go` to `internal/wizard/new_wizard.go`

3. **Implement screens**
   - For each screen, implement rendering and update logic
   - Add validation
   - Test interactively

4. **Add validation**
   - Implement validation functions
   - Return DoPlanError on failure

5. **Test interactively**
   - Use `./scripts/test-tui-interactive.sh`
   - Test all state transitions

6. **Add unit tests**
   - Test each screen
   - Test state transitions
   - Test validation

### Workflow 2: Implementing IDE Integration

**Step-by-step process:**

1. **Study existing integration**
   - Review IDE integration docs
   - Check existing integrations

2. **Create integration function**
   - Detect IDE installation
   - Create directories
   - Create symlinks or copy files
   - Generate config files
   - Verify integration

3. **Add to supported IDEs**
   - Update `supportedIDEs` map
   - Add setup and verify functions

4. **Create setup guide**
   - Generate guide template
   - Include manual setup instructions

5. **Test integration**
   - Use test scripts
   - Verify in actual IDE

### Workflow 3: Implementing Error Handling

**Step-by-step process:**

1. **Define error code**
   - Add to error code registry
   - Choose appropriate category

2. **Create error factory**
   - Implement factory function
   - Include user-friendly message
   - Include fix instructions

3. **Use in code**
   - Return error with context
   - Add cause if available

4. **Display in TUI**
   - Show error with styling
   - Allow retry or recovery

5. **Test error scenario**
   - Use error scenario tests
   - Verify recovery works

## 🔄 Common Patterns

### Pattern 1: State Management

```go
type WizardState struct {
    CurrentScreen string
    Step          int
    TotalSteps    int
    Data          map[string]interface{}
    Errors        []error
}

// Save state
func (m *WizardModel) saveState() error {
    statePath := filepath.Join(projectRoot, ".doplan", "wizard-state.json")
    data, err := json.Marshal(m.state)
    if err != nil {
        return err
    }
    return os.WriteFile(statePath, data, 0644)
}
```

### Pattern 2: Input Validation

```go
type Validator func(string) error

var validators = map[string]Validator{
    "projectName": func(input string) error {
        if len(input) < 3 {
            return ErrInvalidProjectName(input).WithContext("reason", "too_short")
        }
        if !regexp.MustCompile(`^[a-z0-9-]+$`).MatchString(input) {
            return ErrInvalidProjectName(input).WithContext("reason", "invalid_chars")
        }
        return nil
    },
}
```

### Pattern 3: Async Operations

```go
type loadingMsg struct {
    message string
}

type doneMsg struct {
    result interface{}
    err    error
}

func loadDataAsync() tea.Cmd {
    return func() tea.Msg {
        result, err := performOperation()
        return doneMsg{result: result, err: err}
    }
}
```

## 🎨 Styling Patterns

### Pattern 1: Consistent Theming

```go
var theme = struct {
    Primary   lipgloss.Color
    Secondary lipgloss.Color
    Success   lipgloss.Color
    Warning   lipgloss.Color
    Error     lipgloss.Color
    Text      lipgloss.Color
    TextDim   lipgloss.Color
}{
    Primary:   lipgloss.Color("#667eea"),
    Secondary: lipgloss.Color("#764ba2"),
    Success:   lipgloss.Color("#10b981"),
    Warning:   lipgloss.Color("#f59e0b"),
    Error:     lipgloss.Color("#ef4444"),
    Text:      lipgloss.Color("#ffffff"),
    TextDim:   lipgloss.Color("#999999"),
}
```

### Pattern 2: Responsive Layouts

```go
func (m *Model) calculateLayout(width, height int) Layout {
    layout := Layout{
        Padding: 1,
        Margin:  0,
    }
    
    if width < 80 {
        layout.Padding = 0
        layout.Compact = true
    }
    
    if height < 24 {
        layout.Padding = 0
        layout.ShowHelp = false
    }
    
    return layout
}
```

## 🔍 Debugging Techniques

### Technique 1: Logging

```go
import "log"

// Debug logging
func debugLog(format string, args ...interface{}) {
    if os.Getenv("DOPLAN_DEBUG") == "1" {
        log.Printf("[DEBUG] "+format, args...)
    }
}
```

### Technique 2: State Inspection

```go
// Add debug command
case "d":
    // Dump state to file
    stateJSON, _ := json.MarshalIndent(m.state, "", "  ")
    os.WriteFile("debug-state.json", stateJSON, 0644)
    return m, nil
```

## 🧪 Testing Patterns

### Pattern 1: Table-Driven Tests

```go
func TestProjectNameValidation(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
        errCode string
    }{
        {"valid name", "my-project", false, ""},
        {"too short", "ab", true, "USR1001"},
        {"invalid chars", "My Project", true, "USR1001"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateProjectName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("wantErr = %v, got err = %v", tt.wantErr, err)
            }
        })
    }
}
```

## 📊 Performance Optimization

### Optimization 1: Lazy Loading

```go
type LazyLoader struct {
    data     interface{}
    loaded   bool
    loadFunc func() (interface{}, error)
}

func (l *LazyLoader) Get() (interface{}, error) {
    if !l.loaded {
        data, err := l.loadFunc()
        if err != nil {
            return nil, err
        }
        l.data = data
        l.loaded = true
    }
    return l.data, nil
}
```

### Optimization 2: Caching

```go
type Cache struct {
    data      map[string]interface{}
    ttl       time.Duration
    timestamps map[string]time.Time
}

func (c *Cache) Get(key string) (interface{}, bool) {
    if val, ok := c.data[key]; ok {
        if time.Since(c.timestamps[key]) < c.ttl {
            return val, true
        }
        delete(c.data, key)
        delete(c.timestamps, key)
    }
    return nil, false
}
```

## 🔐 Security Best Practices

### Practice 1: Input Sanitization

```go
func sanitizeInput(input string) string {
    // Remove control characters
    input = strings.Map(func(r rune) rune {
        if unicode.IsControl(r) {
            return -1
        }
        return r
    }, input)
    
    // Trim whitespace
    input = strings.TrimSpace(input)
    
    // Limit length
    if len(input) > 1000 {
        input = input[:1000]
    }
    
    return input
}
```

### Practice 2: Path Validation

```go
func validatePath(path string) error {
    // Check for path traversal
    if strings.Contains(path, "..") {
        return errors.New("invalid path: contains '..'")
    }
    
    // Check for absolute paths outside project
    absPath, err := filepath.Abs(path)
    if err != nil {
        return err
    }
    
    projectRoot, _ := filepath.Abs(".")
    if !strings.HasPrefix(absPath, projectRoot) {
        return errors.New("invalid path: outside project root")
    }
    
    return nil
}
```

