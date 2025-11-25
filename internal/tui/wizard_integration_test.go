package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// TestWizardIntegration_EndToEnd tests the complete wizard flow from start to finish
func TestWizardIntegration_EndToEnd(t *testing.T) {
	model := InitialModel()

	// Step 1: Welcome screen → Project Name
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := model.Update(msg)
	model = newModel.(Model)
	if model.state != stateProjectName {
		t.Errorf("Step 1: state = %v, want %v", model.state, stateProjectName)
	}

	// Step 2: Enter project name
	model.textInput.SetValue("integration-test-project")
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)
	if model.state != stateIDESelection {
		t.Errorf("Step 2: state = %v, want %v", model.state, stateIDESelection)
	}
	if model.projectName != "integration-test-project" {
		t.Errorf("Step 2: projectName = %q, want %q", model.projectName, "integration-test-project")
	}

	// Step 3: Navigate IDE selection
	msg = tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)
	if model.selectedIDEIndex != 1 {
		t.Errorf("Step 3: selectedIDEIndex = %d, want 1", model.selectedIDEIndex)
	}

	// Step 4: Select IDE
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)
	if model.state != stateGenerating {
		t.Errorf("Step 4: state = %v, want %v", model.state, stateGenerating)
	}

	// Step 5: Simulate generation completion by sending generationCompleteMsg
	// In real usage, this would be sent by the async generation goroutine
	genCompleteMsg := generationCompleteMsg{err: nil}
	newModel, _ = model.Update(genCompleteMsg)
	model = newModel.(Model)
	if model.state != stateSuccess {
		t.Errorf("Step 5: state = %v, want %v", model.state, stateSuccess)
	}

	// Step 6: Verify project request can be extracted
	request := model.toProjectRequest()
	if request.ProjectName != "integration-test-project" {
		t.Errorf("toProjectRequest() ProjectName = %q, want %q", request.ProjectName, "integration-test-project")
	}
	if request.IDE == "" {
		t.Error("toProjectRequest() IDE should not be empty")
	}
	if len(request.IDEs) == 0 {
		t.Error("toProjectRequest() should include at least one IDE selection")
	}
}

// TestWizardIntegration_BackNavigation tests back navigation through the wizard
func TestWizardIntegration_BackNavigation(t *testing.T) {
	model := InitialModel()

	// Manually set up state history for testing
	model.state = stateIDESelection
	model.stateHistory = []wizardState{stateWelcome, stateProjectName, stateIDESelection}

	// Go back from IDE selection
	msg := tea.KeyMsg{Type: tea.KeyEsc}
	newModel, _ := model.Update(msg)
	model = newModel.(Model)

	if model.state != stateProjectName {
		t.Errorf("Back navigation from IDE: state = %v, want %v", model.state, stateProjectName)
	}

	// Go back from project name
	msg = tea.KeyMsg{Type: tea.KeyEsc}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)

	if model.state != stateWelcome {
		t.Errorf("Back navigation from project name: state = %v, want %v", model.state, stateWelcome)
	}
}

// TestWizardIntegration_ErrorRecovery tests error handling and recovery
func TestWizardIntegration_ErrorRecovery(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName

	// Set an error
	model.setError("Test error", "Test suggestion")

	if model.state != stateError {
		t.Errorf("setError() state = %v, want %v", model.state, stateError)
	}

	// Test retry
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}}
	newModel, _ := model.Update(msg)
	model = newModel.(Model)

	if model.state != stateProjectName {
		t.Errorf("Retry: state = %v, want %v", model.state, stateProjectName)
	}

	// Test back to welcome
	model.setError("Test error", "Test suggestion")
	msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'b'}}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)

	if model.state != stateWelcome {
		t.Errorf("Back from error: state = %v, want %v", model.state, stateWelcome)
	}
}

// TestWizardIntegration_ValidationFlow tests the validation flow
func TestWizardIntegration_ValidationFlow(t *testing.T) {
	model := InitialModel()

	// Navigate to project name
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := model.Update(msg)
	model = newModel.(Model)

	// Try to proceed with invalid project name
	model.textInput.SetValue("invalid@name")
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)

	// Should stay on project name screen
	if model.state != stateProjectName {
		t.Errorf("Invalid name: state = %v, want %v", model.state, stateProjectName)
	}

	// Should have validation error
	if model.validationErr == "" {
		t.Error("Invalid name should set validation error")
	}

	// Fix the name
	model.textInput.SetValue("valid-name")
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)

	// Should proceed to IDE selection
	if model.state != stateIDESelection {
		t.Errorf("Valid name: state = %v, want %v", model.state, stateIDESelection)
	}
}
