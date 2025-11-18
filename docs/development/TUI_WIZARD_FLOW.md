# TUI Wizard Flow - Detailed Screen Mockups & State Transitions

This document provides detailed screen mockups, state transitions, and implementation details for all TUI wizards in v0.0.18-beta.

## 📐 Design Principles

### Visual Style
- **Colors**: Primary Blue (#667eea), Secondary Purple (#764ba2), Success Green (#10b981), Warning Amber (#f59e0b), Error Red (#ef4444)
- **Borders**: Rounded corners, subtle shadows
- **Spacing**: Consistent padding (1-2 characters)
- **Typography**: Bold headers, normal body text, subtle help text

### Interaction Patterns
- **Navigation**: Arrow keys, Tab/Shift+Tab, Enter to confirm, Esc to cancel
- **Input**: Text fields with validation, dropdowns for selections
- **Feedback**: Spinners for loading, progress bars for multi-step operations
- **Errors**: Inline validation with helpful messages

## 🎯 Wizard 1: New Project Wizard

### State Machine
```
[Welcome] → [Project Name] → [Template] → [GitHub] → [IDE] → [Install] → [Success] → [Dashboard]
    ↓            ↓              ↓           ↓         ↓         ↓          ↓
  [Cancel]    [Back]         [Back]      [Back]    [Back]    [Back]    [Exit]
```

### Screen 1: Welcome Screen

**State**: `wizard.welcome`

**Layout**:
```
╔══════════════════════════════════════════════════════════════╗
║                                                              ║
║   ██████╗░░█████╗░██████╗░██╗░░░░░░█████╗░███╗░░██╗          ║
║   ██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔══██╗████╗░██║          ║
║   ██║░░██║██║░░██║██████╔╝██║░░░░░███████║██╔██╗██║          ║
║   ██║░░██║██║░░██║██╔═══╝░██║░░░░░██╔══██║██║╚████║          ║
║   ██████╔╝╚█████╔╝██║░░░░░███████╗██║░░██║██║░╚███║          ║
║   ╚═════╝░░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚══╝          ║
║                                                              ║
║                    Welcome to DoPlan                        ║
║              Transform your ideas into reality              ║
║                                                              ║
║  DoPlan helps you:                                          ║
║  • Structure your project with phases and features          ║
║  • Track progress with beautiful dashboards                 ║
║  • Integrate with your favorite IDE and AI tools           ║
║  • Automate workflows and documentation                     ║
║                                                              ║
║  ┌────────────────────────────────────────────────────┐     ║
║  │  Press [Enter] to start                          │     ║
║  │  Press [Esc] to exit                              │     ║
║  └────────────────────────────────────────────────────┘     ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
```

**State Transitions**:
- `Enter` → `wizard.projectName`
- `Esc` → `exit` (with confirmation)
- `q` → `exit` (with confirmation)

### Screen 2: Project Name Input

**State**: `wizard.projectName`

**Validation Rules**:
- Required field
- Lowercase only
- Alphanumeric + hyphens
- 3-50 characters
- No spaces or special characters

**State Transitions**:
- Valid input + `Enter` → `wizard.template`
- `Tab` → Focus next field (if any)
- `Shift+Tab` → Focus previous field
- `Esc` → `wizard.welcome` (with save prompt)
- `Backspace` on empty → `wizard.welcome`

### Screen 3: Template Gallery

**Templates**:
1. **SaaS Application** - Full-stack web app
2. **Mobile App** - React Native/Flutter
3. **AI Agent** - LLM-powered application
4. **Landing Page** - Marketing website
5. **Chrome Extension** - Browser extension
6. **Electron App** - Desktop application
7. **API Service** - Backend API
8. **CLI Tool** - Command-line tool

**State Transitions**:
- `↑/↓` → Navigate templates
- `Enter` → `wizard.github`
- `←` → `wizard.projectName`
- `Esc` → `wizard.projectName`

### Screen 4: GitHub Repository Setup

**Validation**:
- Must be valid GitHub URL format
- Must be accessible (check with GitHub API)
- Must have write permissions

**State Transitions**:
- Valid URL + `Enter` → `wizard.ide`
- `Create New` → Sub-wizard for repo creation
- `Skip` → Warning dialog → `wizard.ide` (with badge)
- `←` → `wizard.template`

### Screen 5: IDE & AI Selection

**IDEs**:
- Cursor
- VS Code (Copilot)
- Kiro
- Windsurf
- Qoder
- Gemini CLI
- Claude CLI
- Other / Manual Setup

**State Transitions**:
- Selection + `Enter` → `wizard.install`
- `←` → `wizard.github`

### Screen 6: Installation Progress

**Progress Steps**:
1. Create `.doplan/` directory structure
2. Generate initial documentation
3. Set up IDE integration
4. Configure GitHub integration
5. Create initial project state
6. Generate dashboard

**State Transitions**:
- Auto-advance on completion → `wizard.success`
- `Esc` → Cancel dialog → `wizard.ide` (with cleanup)

### Screen 7: Success Screen

**State Transitions**:
- `Dashboard` → Open main TUI dashboard
- `Exit` → Close wizard
- Auto-open dashboard after 3 seconds (optional)

## 🎯 Wizard 2: Project Adoption Wizard

### State Machine
```
[Detection] → [Analysis] → [Options] → [GitHub] → [IDE] → [Analysis] → [Preview] → [Confirm] → [Success]
     ↓            ↓           ↓          ↓         ↓          ↓           ↓          ↓
  [Skip]      [Back]      [Back]     [Back]    [Back]     [Back]      [Back]    [Cancel]
```

### Screen 1: Detection Screen

Shows "Found existing project!" with detected project details.

### Screen 2: Analysis Results

Shows analysis progress and results:
- Tech stack detection
- File structure mapping
- Documentation extraction
- Feature identification

### Screen 3: Adoption Options

Three options:
- Analyze & Generate Plan
- Import Existing Docs
- Start Fresh

## 🎯 Wizard 3: Migration Wizard

### State Machine
```
[Detection] → [Backup] → [Config] → [Folders] → [Dashboard] → [IDE] → [Verify] → [Complete]
     ↓           ↓          ↓          ↓            ↓          ↓         ↓
  [Skip]      [Cancel]   [Back]    [Back]       [Back]     [Back]   [Rollback]
```

### Screen 1: Migration Detection

Shows detected old structure and migration explanation.

### Screen 2: Folder Renaming Preview

Shows old → new folder name mappings with options:
- Auto-rename (recommended)
- Manual rename
- Skip renaming

## 🔄 State Management

### Wizard State Structure
```go
type WizardState struct {
    CurrentScreen string
    ProjectName   string
    Template      string
    GitHubRepo    string
    IDE           string
    Step          int
    TotalSteps    int
    Data          map[string]interface{}
    Errors        []error
}
```

## 🎨 Styling Guide

### Color Palette
```go
const (
    ColorPrimary   = "#667eea"  // Blue
    ColorSecondary = "#764ba2"  // Purple
    ColorSuccess   = "#10b981"  // Green
    ColorWarning   = "#f59e0b"  // Amber
    ColorError     = "#ef4444"  // Red
    ColorText      = "#ffffff"  // White
    ColorTextDim   = "#999999"  // Gray
    ColorBorder    = "#333333"  // Dark gray
)
```

## ⌨️ Keyboard Shortcuts

### Global Shortcuts
- `Esc` - Go back / Cancel
- `Ctrl+C` - Exit (with confirmation)
- `q` - Quick exit (with confirmation)
- `Tab` - Next field / Next option
- `Shift+Tab` - Previous field / Previous option

### Navigation
- `↑/↓` - Navigate lists
- `Enter` - Confirm / Next
- `←/→` - Navigate tabs (if applicable)

## 📱 Responsive Design

### Terminal Size Handling
```go
func (m *WizardModel) handleResize(width, height int) {
    if width < 80 {
        // Show compact layout
        m.layout = "compact"
    } else if width < 120 {
        // Show normal layout
        m.layout = "normal"
    } else {
        // Show wide layout
        m.layout = "wide"
    }
    
    if height < 24 {
        // Reduce padding
        m.padding = 0
    }
}
```

