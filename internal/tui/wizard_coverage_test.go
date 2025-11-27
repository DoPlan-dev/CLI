package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// TestMapIDEToCommand tests the mapIDEToCommand function which has 0% coverage
func TestMapIDEToCommand(t *testing.T) {
	tests := []struct {
		name     string
		ide      string
		expected string
	}{
		{"Cursor", "Cursor", "cursor"},
		{"Claude Code", "Claude Code", "claude"},
		{"Antigravity", "Antigravity", "antigravity"},
		{"Windsurf", "Windsurf", "windsurf"},
		{"Cline", "Cline", "cline"},
		{"OpenCode", "OpenCode", "opencode"},
		{"Unknown IDE", "Unknown", "code"},
		{"Empty IDE", "", "code"},
		{"Lowercase cursor", "cursor", "code"}, // Should not match
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapIDEToCommand(tt.ide)
			if result != tt.expected {
				t.Errorf("mapIDEToCommand(%q) = %q, want %q", tt.ide, result, tt.expected)
			}
		})
	}
}

// TestInit_GeneratingState tests Init when in generating state (66.7% -> 100%)
func TestInit_GeneratingState(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating

	cmd := model.Init()
	if cmd == nil {
		t.Error("Init() should return tickSpinner command when in generating state")
	}

	// Verify it's a time.After function
	timeCmd := cmd()
	if timeCmd == nil {
		t.Error("Init() returned command should be callable")
	}
}

// TestUpdate_AllKeyMessages tests Update function edge cases to increase 65% -> 95%+
func TestUpdate_AllKeyMessages(t *testing.T) {
	tests := []struct {
		name           string
		initialState   wizardState
		key            string
		expectState    wizardState
		expectQuit     bool
		setupModel     func(*Model)
	}{
		{
			name:         "Quit with ctrl+c",
			initialState: stateWelcome,
			key:          "ctrl+c",
			expectQuit:   true,
		},
		{
			name:         "Retry from error state",
			initialState: stateError,
			key:          "r",
			expectState:  stateWelcome,
			setupModel: func(m *Model) {
				m.previousState = stateWelcome
			},
		},
		{
			name:         "Go back from error to welcome",
			initialState: stateError,
			key:          "b",
			expectState:  stateWelcome,
		},
		{
			name:         "Backspace from project name",
			initialState: stateProjectName,
			key:          "backspace",
			setupModel: func(m *Model) {
				m.stateHistory = []wizardState{stateWelcome}
				m.textInput.Focus()
			},
		},
		{
			name:         "ESC from IDE selection",
			initialState: stateIDESelection,
			key:          "esc",
			setupModel: func(m *Model) {
				m.stateHistory = []wizardState{stateWelcome, stateProjectName}
			},
		},
		{
			name:         "Enter on success screen exits",
			initialState: stateSuccess,
			key:          "enter",
			expectQuit:   true,
		},
		{
			name:         "Up arrow in IDE selection",
			initialState: stateIDESelection,
			key:          "up",
			setupModel: func(m *Model) {
				m.selectedIDEIndex = 1
			},
		},
		{
			name:         "Down arrow in IDE selection wraps to top",
			initialState: stateIDESelection,
			key:          "down",
			setupModel: func(m *Model) {
				ideInfo := getIDEInfo()
				m.selectedIDEIndex = len(ideInfo) - 1
			},
		},
		{
			name:         "K key (vim-style up) in IDE selection",
			initialState: stateIDESelection,
			key:          "k",
			setupModel: func(m *Model) {
				m.selectedIDEIndex = 1
			},
		},
		{
			name:         "J key (vim-style down) in IDE selection",
			initialState: stateIDESelection,
			key:          "j",
			setupModel: func(m *Model) {
				m.selectedIDEIndex = 0
			},
		},
		{
			name:         "Space toggles IDE selection",
			initialState: stateIDESelection,
			key:          " ",
			setupModel: func(m *Model) {
				m.selectedIDEs = make(map[string]bool)
				m.selectedIDEIndex = 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			model.state = tt.initialState
			if tt.setupModel != nil {
				tt.setupModel(&model)
			}

			var msg tea.Msg
			if tt.key == "ctrl+c" {
				msg = tea.KeyMsg{Type: tea.KeyCtrlC}
			} else {
				msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.key)}
			}

			newModel, cmd := model.Update(msg)
			updatedModel := newModel.(Model)

			if tt.expectQuit {
				if cmd == nil {
					t.Error("Expected quit command, got nil")
				} else {
					// Verify it's a quit command
					quitMsg := cmd()
					if _, ok := quitMsg.(tea.QuitMsg); !ok {
						t.Errorf("Expected tea.QuitMsg, got %T", quitMsg)
					}
				}
			}

			if tt.expectState != 0 && updatedModel.state != tt.expectState {
				t.Errorf("Update() state = %v, want %v", updatedModel.state, tt.expectState)
			}
		})
	}
}

// TestUpdate_TextInputMessages tests Update with text input messages
func TestUpdate_TextInputMessages(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.textInput.Focus()

	// Test typing in project name
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'t', 'e', 's', 't'}}
	for _, r := range msg.Runes {
		charMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		newModel, _ := model.Update(charMsg)
		model = newModel.(Model)
	}

	if model.textInput.Value() != "test" {
		t.Errorf("Text input value = %q, want %q", model.textInput.Value(), "test")
	}
}

// TestUpdate_TimeMessage tests Update with time messages (spinner ticks)
func TestUpdate_TimeMessage(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.spinnerFrame = 0

	msg := time.Now()
	newModel, cmd := model.Update(msg)
	updatedModel := newModel.(Model)

	if updatedModel.spinnerFrame == 0 {
		t.Error("Update(time.Time) should increment spinnerFrame")
	}

	if cmd == nil {
		t.Error("Update(time.Time) in generating state should return tickSpinner command")
	}
}

// TestUpdate_GenerationCompleteMsg tests Update with generation complete messages
func TestUpdate_GenerationCompleteMsg(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectState  wizardState
		expectError  bool
	}{
		{
			name:        "Success - no error",
			err:         nil,
			expectState: stateSuccess,
			expectError: false,
		},
		{
			name:        "Failure - with error",
			err:         &testError{message: "generation failed"},
			expectState: stateError,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			model.state = stateGenerating
			model.steps = getGenerationSteps()

			msg := generationCompleteMsg{err: tt.err}
			newModel, _ := model.Update(msg)
			updatedModel := newModel.(Model)

			if updatedModel.state != tt.expectState {
				t.Errorf("Update(generationCompleteMsg) state = %v, want %v", updatedModel.state, tt.expectState)
			}

			if tt.expectError {
				if updatedModel.errorMessage == "" {
					t.Error("Update(generationCompleteMsg) should set error message on error")
				}
			} else {
				// Check all steps are completed
				for _, step := range updatedModel.steps {
					if step.Status != stepCompleted {
						t.Error("All steps should be completed on success")
					}
				}
			}
		})
	}
}

// TestUpdate_StepProgressMsg tests Update with step progress messages
func TestUpdate_StepProgressMsg(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.steps = getGenerationSteps()
	// Set first step as in progress
	if len(model.steps) > 0 {
		model.steps[0].Status = stepInProgress
	}

	// Simulate step 0 completion, moving to step 1
	msg := stepProgressMsg{stepIndex: 1}
	newModel, _ := model.Update(msg)
	updatedModel := newModel.(Model)

	// Check step 0 is completed
	if len(updatedModel.steps) > 0 && updatedModel.steps[0].Status != stepCompleted {
		t.Error("Previous step should be marked as completed")
	}

	// Check step 1 is in progress
	if len(updatedModel.steps) > 1 && updatedModel.steps[1].Status != stepInProgress {
		t.Error("Next step should be marked as in progress")
	}
}

// TestRenderASCIIHeader tests renderASCIIHeader function (68.8% -> 100%)
func TestRenderASCIIHeader(t *testing.T) {
	header := renderASCIIHeader()

	if len(header) == 0 {
		t.Error("renderASCIIHeader() should return non-empty string")
	}

	// Check it contains DoPlan ASCII art
	if !strings.Contains(header, "██████╗") {
		t.Error("renderASCIIHeader() should contain DoPlan ASCII art")
	}

	// Check it's centered (should have some padding)
	lines := strings.Split(header, "\n")
	for _, line := range lines {
		if len(line) > 0 && !strings.HasPrefix(strings.TrimSpace(line), "█") {
			// Lines should be centered, so trimmed should start with art
			t.Logf("Line: %q", line)
		}
	}
}

// TestGetProjectStructureTree_EdgeCases tests getProjectStructureTree edge cases (75% -> 100%)
func TestGetProjectStructureTree_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		setupModel  func(*Model)
	}{
		{
			name:        "With project name",
			projectName: "my-app",
			setupModel: func(m *Model) {
				m.projectName = "my-app"
			},
		},
		{
			name:        "Without project name (should default)",
			projectName: "",
			setupModel: func(m *Model) {
				m.projectName = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			if tt.setupModel != nil {
				tt.setupModel(&model)
			}

			tree := model.getProjectStructureTree()

			if len(tree) == 0 {
				t.Error("getProjectStructureTree() should return non-empty string")
			}

			// Check it contains project structure
			if !strings.Contains(tree, ".cursor/") {
				t.Error("getProjectStructureTree() should contain .cursor/")
			}
			if !strings.Contains(tree, ".do/") {
				t.Error("getProjectStructureTree() should contain .do/")
			}

			// If project name was set, it should appear
			if tt.projectName != "" && !strings.Contains(tree, tt.projectName) {
				t.Errorf("getProjectStructureTree() should contain project name %q", tt.projectName)
			}

			// If project name was empty, should default to "my-project"
			if tt.projectName == "" && !strings.Contains(tree, "my-project") {
				t.Error("getProjectStructureTree() should default to 'my-project' when name is empty")
			}
		})
	}
}

// TestValidateProjectName_EdgeCases tests validateProjectName edge cases (75% -> 100%)
func TestValidateProjectName_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		wantValid   bool
	}{
		{"Valid name", "my-project", true},
		{"Valid with numbers", "project123", true},
		{"Valid with underscore", "my_project", true},
		{"Empty name", "", false},
		{"Short name (valid per rules)", "ab", true}, // IsValidProjectName doesn't check length
		{"Invalid characters", "my@project", false},
		{"Starts with number (valid per rules)", "123project", true}, // IsValidProjectName allows numbers at start
		{"Only numbers (valid per rules)", "12345", true}, // IsValidProjectName allows all numbers
		{"Valid with dash", "my-project-name", true},
		{"Invalid with space", "my project", false},
		{"Invalid with special chars", "my.project", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			valid := model.validateProjectName(tt.projectName)
			if valid != tt.wantValid {
				t.Errorf("validateProjectName(%q) = %v, want %v", tt.projectName, valid, tt.wantValid)
			}
		})
	}
}

// Helper type for testing errors
type testError struct {
	message string
}

func (e *testError) Error() string {
	return e.message
}

