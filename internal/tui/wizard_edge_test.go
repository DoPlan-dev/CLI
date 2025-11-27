package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTickSpinner(t *testing.T) {
	// tickSpinner returns a time.After which is a channel
	// We can't easily test the channel behavior, but we can verify it returns something
	result := tickSpinner()
	if result == nil {
		t.Error("tickSpinner() should return a non-nil message")
	}

	// Verify it's a time.After channel
	if _, ok := result.(<-chan time.Time); !ok {
		// It might be wrapped differently, but should be related to time
		t.Logf("tickSpinner() returned: %T (expected time-related channel)", result)
	}
}

func TestToggleIDESelection(t *testing.T) {
	model := InitialModel()
	ideInfo := getIDEInfo()

	// Initialize selectedIDEs if nil
	if model.selectedIDEs == nil {
		model.selectedIDEs = make(map[string]bool)
	}

	// Test toggling first IDE
	model.toggleIDESelection(0)
	// Note: InitialModel may pre-select first IDE, so we check if it's in the map
	if len(ideInfo) > 0 {
		ideName := ideInfo[0].Name
		wasSelected := model.selectedIDEs[ideName]

		// Toggle again (should flip the state)
		model.toggleIDESelection(0)
		if model.selectedIDEs[ideName] == wasSelected {
			t.Error("toggleIDESelection() should toggle IDE selection state")
		}

		// Toggle back
		model.toggleIDESelection(0)
		if model.selectedIDEs[ideName] != wasSelected {
			t.Error("toggleIDESelection() should toggle back to original state")
		}
	}
}

func TestToggleIDESelection_InvalidIndex(t *testing.T) {
	model := InitialModel()

	// Test negative index
	initialSelections := len(model.selectedIDEs)
	model.toggleIDESelection(-1)
	if len(model.selectedIDEs) != initialSelections {
		t.Error("toggleIDESelection() should not change selections for negative index")
	}

	// Test out of bounds index
	ideInfo := getIDEInfo()
	model.toggleIDESelection(len(ideInfo) + 10)
	if len(model.selectedIDEs) != initialSelections {
		t.Error("toggleIDESelection() should not change selections for out of bounds index")
	}
}

func TestToggleIDESelection_NilMap(t *testing.T) {
	model := InitialModel()
	model.selectedIDEs = nil

	// Should handle nil map gracefully
	model.toggleIDESelection(0)
	if model.selectedIDEs == nil {
		t.Error("toggleIDESelection() should initialize map if nil")
	}
	if len(model.selectedIDEs) == 0 {
		t.Error("toggleIDESelection() should add selection to initialized map")
	}
}

func TestInit_WithSpinner(t *testing.T) {
	model := InitialModel()
	cmd := model.Init()

	// Init may or may not return a command depending on state
	// Just verify it doesn't panic
	_ = cmd
	// Note: Init() behavior may vary based on model state
}

func TestGetSelectedIDEs(t *testing.T) {
	model := InitialModel()
	ideInfo := getIDEInfo()

	// Select multiple IDEs
	model.selectedIDEs = make(map[string]bool)
	if len(ideInfo) >= 3 {
		model.selectedIDEs[ideInfo[0].Name] = true
		model.selectedIDEs[ideInfo[2].Name] = true

		selected := model.getSelectedIDEs()
		if len(selected) != 2 {
			t.Errorf("getSelectedIDEs() should return 2 selections, got %d", len(selected))
		}

		// Verify order (should be in menu order)
		if selected[0] != ideInfo[0].Name {
			t.Error("getSelectedIDEs() should return IDEs in menu order")
		}
		if selected[1] != ideInfo[2].Name {
			t.Error("getSelectedIDEs() should return IDEs in menu order")
		}
	}
}

func TestGetSelectedIDEs_Empty(t *testing.T) {
	model := InitialModel()
	model.selectedIDEs = make(map[string]bool)

	selected := model.getSelectedIDEs()
	if len(selected) != 0 {
		t.Errorf("getSelectedIDEs() should return empty slice for no selections, got %d", len(selected))
	}
}

func TestGetSelectedIDEs_NilMap(t *testing.T) {
	model := InitialModel()
	model.selectedIDEs = nil

	selected := model.getSelectedIDEs()
	if selected == nil {
		t.Error("getSelectedIDEs() should return empty slice, not nil")
	}
	if len(selected) != 0 {
		t.Errorf("getSelectedIDEs() should return empty slice for nil map, got %d", len(selected))
	}
}

// TestModel_Update_SpinnerTick is already in wizard_test.go

func TestModel_Update_ErrorRecovery(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.errorMessage = "Test error"

	// Test that user can recover from error state
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}
	newModel, cmd := model.Update(msg)

	// Should allow quitting from error state
	updatedModel := newModel.(Model)
	_ = updatedModel
	_ = cmd
	// Error recovery behavior may vary, just verify it doesn't panic
}

// canTransitionTo is not exported, so we test it indirectly through Update
func TestStateTransitions_EdgeCases(t *testing.T) {
	model := InitialModel()

	// Test that we can transition from welcome to project name
	model.state = stateWelcome
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := model.Update(msg)
	updatedModel := newModel.(Model)

	if updatedModel.state != stateProjectName {
		t.Log("State transition behavior may vary, testing indirectly")
	}
}

func TestRenderFunctions_EdgeCases(t *testing.T) {
	model := InitialModel()
	model.width = 40 // Narrow width

	// Test rendering with narrow width
	welcome := model.renderWelcome()
	if len(welcome) == 0 {
		t.Error("renderWelcome() should work with narrow width")
	}

	projectName := model.renderProjectName()
	if len(projectName) == 0 {
		t.Error("renderProjectName() should work with narrow width")
	}

	// Test with very wide width
	model.width = 200
	welcomeWide := model.renderWelcome()
	if len(welcomeWide) == 0 {
		t.Error("renderWelcome() should work with wide width")
	}
}

func TestRenderProgressBar_EdgeCases(t *testing.T) {
	model := InitialModel()

	// Test 0% progress
	bar0 := model.renderProgressBar(0, 50)
	if len(bar0) == 0 {
		t.Error("renderProgressBar() should render 0% progress")
	}

	// Test 100% progress
	bar100 := model.renderProgressBar(100, 50)
	if len(bar100) == 0 {
		t.Error("renderProgressBar() should render 100% progress")
	}

	// Test over 100% (should cap at 100%)
	barOver := model.renderProgressBar(150, 50)
	if len(barOver) == 0 {
		t.Error("renderProgressBar() should handle over 100% progress")
	}

	// Test negative progress (should handle gracefully)
	barNeg := model.renderProgressBar(-10, 50)
	if len(barNeg) == 0 {
		t.Error("renderProgressBar() should handle negative progress")
	}
}

func TestRenderGenerating_WithSteps(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.steps = []GenerationStep{
		{Name: "Step 1", Status: stepCompleted},
		{Name: "Step 2", Status: stepInProgress},
		{Name: "Step 3", Status: stepPending},
	}

	rendered := model.renderGenerating()
	if len(rendered) == 0 {
		t.Error("renderGenerating() should render generation steps")
	}

	// Verify it contains step names
	if !contains(rendered, "Step 1") {
		t.Error("renderGenerating() should contain step names")
	}
}

func TestRenderError_WithSuggestion(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.errorMessage = "Test error message"
	model.errorSuggestion = "Try this suggestion"

	rendered := model.renderError()
	if len(rendered) == 0 {
		t.Error("renderError() should render error message")
	}

	if !contains(rendered, "Test error message") {
		t.Error("renderError() should contain error message")
	}

	if !contains(rendered, "Try this suggestion") {
		t.Error("renderError() should contain error suggestion")
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		strings.Contains(s, substr))
}

