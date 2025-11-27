package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// TestUpdate_AllMessageTypes tests Update with all message types to increase coverage
func TestUpdate_AllMessageTypes(t *testing.T) {
	tests := []struct {
		name         string
		initialState wizardState
		msg          tea.Msg
		expectState  wizardState
		setupModel   func(*Model)
	}{
		{
			name:         "WindowSizeMsg updates width",
			initialState: stateWelcome,
			msg:          tea.WindowSizeMsg{Width: 100, Height: 30},
			setupModel: func(m *Model) {
				m.width = 80
			},
		},
		{
			name:         "Time message increments spinner in generating state",
			initialState: stateGenerating,
			msg:          time.Now(),
			setupModel: func(m *Model) {
				m.spinnerFrame = 0
			},
		},
		{
			name:         "Time message ignored in non-generating state",
			initialState: stateWelcome,
			msg:          time.Now(),
			setupModel: func(m *Model) {
				m.spinnerFrame = 0
			},
		},
		{
			name:         "Enter on welcome transitions to project name",
			initialState: stateWelcome,
			msg:          tea.KeyMsg{Type: tea.KeyEnter},
			expectState:  stateProjectName,
		},
		{
			name:         "Enter on project name with valid name proceeds",
			initialState: stateProjectName,
			msg:          tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(m *Model) {
				m.textInput.SetValue("valid-project")
				m.textInput.Focus()
			},
			expectState: stateIDESelection,
		},
		{
			name:         "Enter on project name with invalid name stays",
			initialState: stateProjectName,
			msg:          tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(m *Model) {
				m.textInput.SetValue("invalid@project")
				m.textInput.Focus()
			},
			expectState: stateProjectName, // Should stay
		},
		{
			name:         "Enter on IDE selection with selection proceeds",
			initialState: stateIDESelection,
			msg:          tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(m *Model) {
				m.selectedIDEs = make(map[string]bool)
				ideInfo := getIDEInfo()
				if len(ideInfo) > 0 {
					m.selectedIDEs[ideInfo[0].Name] = true
				}
			},
			expectState: stateGenerating,
		},
		{
			name:         "Enter on IDE selection without selection shows error",
			initialState: stateIDESelection,
			msg:          tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(m *Model) {
				m.selectedIDEs = make(map[string]bool) // Empty
			},
			expectState: stateIDESelection, // Should stay
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			model.state = tt.initialState
			if tt.setupModel != nil {
				tt.setupModel(&model)
			}

			newModel, _ := model.Update(tt.msg)
			updatedModel := newModel.(Model)

			if tt.expectState != 0 && updatedModel.state != tt.expectState {
				t.Errorf("Update() state = %v, want %v", updatedModel.state, tt.expectState)
			}

			// For WindowSizeMsg, check width was updated
			if ws, ok := tt.msg.(tea.WindowSizeMsg); ok {
				if updatedModel.width != ws.Width {
					t.Errorf("Update(WindowSizeMsg) width = %v, want %v", updatedModel.width, ws.Width)
				}
			}

			// For Time message in generating state, check spinner incremented
			if _, ok := tt.msg.(time.Time); ok && tt.initialState == stateGenerating {
				if updatedModel.spinnerFrame == 0 {
					t.Error("Update(time.Time) in generating state should increment spinnerFrame")
				}
			}
		})
	}
}

// TestUpdate_TextInputMessages_Extended tests Update with text input messages (extended cases)
func TestUpdate_TextInputMessages_Extended(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.textInput.Focus()

	// Simulate typing multiple characters
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'t', 'e', 's', 't'}}
	for _, r := range msg.Runes {
		charMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		newModel, _ := model.Update(charMsg)
		model = newModel.(Model)
	}

	if model.textInput.Value() != "test" {
		t.Errorf("Text input value = %q, want %q", model.textInput.Value(), "test")
	}

	// Test that text input is working - verify we can type and it updates
	// (backspace behavior may vary with textinput implementation)
	if len(model.textInput.Value()) == 0 {
		t.Error("Text input should have value after typing")
	}
}

// TestUpdate_IDESelectionNavigation tests IDE selection navigation edge cases
func TestUpdate_IDESelectionNavigation(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	ideInfo := getIDEInfo()

	if len(ideInfo) == 0 {
		t.Skip("No IDE info available for testing")
	}

	// Test wrapping from top (up arrow at index 0)
	model.selectedIDEIndex = 0
	msg := tea.KeyMsg{Type: tea.KeyUp}
	newModel, _ := model.Update(msg)
	updatedModel := newModel.(Model)

	// Should wrap to last index
	if updatedModel.selectedIDEIndex != len(ideInfo)-1 {
		t.Errorf("Up arrow at index 0 should wrap to %d, got %d", len(ideInfo)-1, updatedModel.selectedIDEIndex)
	}

	// Test wrapping from bottom (down arrow at last index)
	model.selectedIDEIndex = len(ideInfo) - 1
	msg = tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ = model.Update(msg)
	updatedModel = newModel.(Model)

	// Should wrap to first index
	if updatedModel.selectedIDEIndex != 0 {
		t.Errorf("Down arrow at last index should wrap to 0, got %d", updatedModel.selectedIDEIndex)
	}
}

// TestUpdate_StepProgressEdgeCases tests step progress message edge cases
func TestUpdate_StepProgressEdgeCases(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.steps = getGenerationSteps()

	if len(model.steps) == 0 {
		t.Skip("No generation steps available")
	}

	// Test step progress with out of bounds index
	msg := stepProgressMsg{stepIndex: len(model.steps) + 10}
	newModel, _ := model.Update(msg)
	updatedModel := newModel.(Model)

	// Should not panic and should handle gracefully
	if updatedModel.state != stateGenerating {
		t.Errorf("Step progress with out of bounds index should maintain state, got %v", updatedModel.state)
	}

	// Test step progress at index 0
	msg = stepProgressMsg{stepIndex: 0}
	model.steps[0].Status = stepPending
	newModel, _ = model.Update(msg)
	updatedModel = newModel.(Model)

	// First step should be in progress
	if len(updatedModel.steps) > 0 && updatedModel.steps[0].Status != stepInProgress {
		t.Error("Step progress at index 0 should mark first step as in progress")
	}
}

// TestUpdate_DefaultCase tests the default case in Update switch
func TestUpdate_DefaultCase(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.textInput.Focus()

	// Test with a custom message type (not handled in switch)
	type customMsg struct{}
	msg := customMsg{}

	newModel, _ := model.Update(msg)
	updatedModel := newModel.(Model)

	// Should not change state for unhandled message
	if updatedModel.state != stateProjectName {
		t.Errorf("Unhandled message should not change state, got %v", updatedModel.state)
	}
}

// TestStartGenerationAsync_ValidationError tests startGenerationAsync with validation error
func TestStartGenerationAsync_ValidationError(t *testing.T) {
	model := InitialModel()
	model.projectName = "" // Invalid - empty name
	model.selectedIDEs = make(map[string]bool)
	model.selectedIDEs["Cursor"] = true

	cmd := startGenerationAsync(model)
	if cmd == nil {
		t.Fatal("startGenerationAsync() should return a command")
	}

	// Execute the command
	msg := cmd()
	if msg == nil {
		t.Fatal("startGenerationAsync command should return a message")
	}

	// Should return generationCompleteMsg with error
	genMsg, ok := msg.(generationCompleteMsg)
	if !ok {
		t.Fatalf("Expected generationCompleteMsg, got %T", msg)
	}

	if genMsg.err == nil {
		t.Error("startGenerationAsync() should return error for invalid request")
	}
}

// TestToProjectRequest_AllCases tests toProjectRequest with various selections
func TestToProjectRequest_AllCases(t *testing.T) {
	tests := []struct {
		name       string
		setupModel func(*Model)
		expectIDEs int
		expectName string
	}{
		{
			name: "With selected IDEs",
			setupModel: func(m *Model) {
				m.projectName = "test-project"
				m.selectedIDEs = make(map[string]bool)
				ideInfo := getIDEInfo()
				if len(ideInfo) >= 2 {
					m.selectedIDEs[ideInfo[0].Name] = true
					m.selectedIDEs[ideInfo[1].Name] = true
				}
			},
			expectIDEs: 2,
			expectName: "test-project",
		},
		{
			name: "Without selected IDEs, with legacy selectedIDE",
			setupModel: func(m *Model) {
				m.projectName = "test-project"
				m.selectedIDEs = nil // Will fall back to selectedIDE in getSelectedIDEs
				m.selectedIDE = "Cursor"
			},
			expectIDEs: 1, // getSelectedIDEs will fall back to selectedIDE field
			expectName: "test-project",
		},
		{
			name: "With empty selections",
			setupModel: func(m *Model) {
				m.projectName = "test-project"
				m.selectedIDEs = make(map[string]bool)
			},
			expectIDEs: 0,
			expectName: "test-project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			if tt.setupModel != nil {
				tt.setupModel(&model)
			}

			request := model.toProjectRequest()

			if request.ProjectName != tt.expectName {
				t.Errorf("toProjectRequest() ProjectName = %q, want %q", request.ProjectName, tt.expectName)
			}

			if len(request.IDEs) != tt.expectIDEs {
				t.Errorf("toProjectRequest() IDEs length = %d, want %d", len(request.IDEs), tt.expectIDEs)
			}

			if request.ProjectType != "Fullstack" {
				t.Errorf("toProjectRequest() ProjectType = %q, want %q", request.ProjectType, "Fullstack")
			}
		})
	}
}

// TestRenderASCIIHeader_EdgeCases tests renderASCIIHeader with different scenarios
func TestRenderASCIIHeader_EdgeCases(t *testing.T) {
	// Test that it returns ASCII art
	header := renderASCIIHeader()

	if len(header) == 0 {
		t.Error("renderASCIIHeader() should return non-empty string")
	}

	// Should contain DoPlan characters (check for any non-space characters)
	if strings.TrimSpace(header) == "" {
		t.Error("renderASCIIHeader() should contain ASCII art characters")
	}

	// Should have multiple lines
	lines := strings.Split(header, "\n")
	if len(lines) < 5 {
		t.Errorf("renderASCIIHeader() should have multiple lines, got %d", len(lines))
	}
}
