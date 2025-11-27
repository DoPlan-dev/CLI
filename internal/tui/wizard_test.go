package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInitialModel(t *testing.T) {
	model := InitialModel()

	if model.state != stateWelcome {
		t.Errorf("InitialModel() state = %v, want %v", model.state, stateWelcome)
	}

	if model.width != 80 {
		t.Errorf("InitialModel() width = %v, want 80", model.width)
	}
}

func TestModel_Init(t *testing.T) {
	model := InitialModel()
	cmd := model.Init()

	if cmd != nil {
		t.Error("Init() should return nil command initially")
	}
}

func TestModel_Update_WindowSize(t *testing.T) {
	model := InitialModel()
	msg := tea.WindowSizeMsg{Width: 120, Height: 40}

	newModel, cmd := model.Update(msg)

	if newModel.(Model).width != 120 {
		t.Errorf("Update(WindowSizeMsg) width = %v, want 120", newModel.(Model).width)
	}

	if cmd != nil {
		t.Error("Update(WindowSizeMsg) should return nil command")
	}
}

func TestModel_Update_QuitKey(t *testing.T) {
	model := InitialModel()
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}

	newModel, cmd := model.Update(msg)

	if cmd == nil {
		t.Error("Update(KeyMsg 'q') should return quit command")
	}

	// Check if it's a quit command
	if _, ok := cmd().(tea.QuitMsg); !ok {
		// If not a quit command, check if it's nil (which is also acceptable for testing)
		if cmd != nil {
			t.Error("Update(KeyMsg 'q') should return quit command")
		}
	}

	_ = newModel
}

func TestModel_Update_EnterKey_Welcome(t *testing.T) {
	model := InitialModel()
	model.state = stateWelcome
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	newModel, cmd := model.Update(msg)

	// Enter on welcome screen should transition to project name input
	updatedModel := newModel.(Model)
	if updatedModel.state != stateProjectName {
		t.Errorf("Update(Enter on welcome) state = %v, want %v", updatedModel.state, stateProjectName)
	}

	_ = cmd
}

func TestModel_Update_EnterKey_ProjectName(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.textInput.SetValue("valid-project")
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	newModel, _ := model.Update(msg)

	// Enter on project name with valid input should transition to IDE selection
	updatedModel := newModel.(Model)
	if updatedModel.projectName != "valid-project" {
		t.Errorf("Update(Enter on valid project name) projectName = %q, want %q", updatedModel.projectName, "valid-project")
	}
	if updatedModel.state != stateIDESelection {
		t.Errorf("Update(Enter on valid project name) state = %v, want %v", updatedModel.state, stateIDESelection)
	}
}

func TestModel_Update_EnterKey_ProjectName_Invalid(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.textInput.SetValue("invalid project name!")
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	newModel, _ := model.Update(msg)

	// Enter on project name with invalid input should not proceed
	updatedModel := newModel.(Model)
	if updatedModel.state != stateProjectName {
		t.Errorf("Update(Enter on invalid project name) state = %v, want %v", updatedModel.state, stateProjectName)
	}
	if updatedModel.validationErr == "" {
		t.Error("Update(Enter on invalid project name) should set validation error")
	}
}

func TestModel_View_Welcome(t *testing.T) {
	model := InitialModel()
	model.width = 80

	view := model.View()

	if len(view) == 0 {
		t.Error("View() should return non-empty string for welcome screen")
	}

	// Check that view contains expected elements
	if !strings.Contains(view, "DoPlan CLI") {
		t.Error("View() should contain 'DoPlan CLI'")
	}

	if !strings.Contains(view, "[>]") {
		t.Error("View() should contain welcome indicator")
	}

	if !strings.Contains(view, "Enter") {
		t.Error("View() should contain 'Enter' instruction")
	}
}

func TestModel_View_UnknownState(t *testing.T) {
	model := InitialModel()
	model.state = wizardState(999) // Unknown state

	view := model.View()

	if view != "" {
		t.Errorf("View() for unknown state should return empty string, got %q", view)
	}
}

func TestRenderWelcome(t *testing.T) {
	model := InitialModel()
	model.width = 80

	welcome := model.renderWelcome()

	if len(welcome) == 0 {
		t.Error("renderWelcome() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(welcome, "DoPlan CLI") {
		t.Error("renderWelcome() should contain 'DoPlan CLI'")
	}

	if !strings.Contains(welcome, "[>]") {
		t.Error("renderWelcome() should contain welcome indicator")
	}
}

func TestModel_View_ProjectName(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.width = 80

	view := model.View()

	if len(view) == 0 {
		t.Error("View() should return non-empty string for project name screen")
	}

	// Check that view contains expected elements
	if !strings.Contains(view, "Project Name") {
		t.Error("View() should contain 'Project Name'")
	}

	if !strings.Contains(view, "[*]") {
		t.Error("View() should contain project name indicator")
	}
}

func TestValidateProjectName(t *testing.T) {
	model := InitialModel()

	// Test valid names
	validNames := []string{
		"my-project",
		"my_project",
		"myproject",
		"myProject123",
		"project-123",
		"project_123",
	}

	for _, name := range validNames {
		model.textInput.SetValue(name)
		if !model.validateProjectName(name) {
			t.Errorf("validateProjectName(%q) = false, want true", name)
		}
		if model.validationErr != "" {
			t.Errorf("validateProjectName(%q) should not set error, got %q", name, model.validationErr)
		}
	}

	// Test invalid names
	invalidNames := []string{
		"my project", // space
		"my.project", // dot
		"my@project", // @ symbol
		"my#project", // # symbol
		"",           // empty
	}

	for _, name := range invalidNames {
		model.textInput.SetValue(name)
		if name != "" && model.validateProjectName(name) {
			t.Errorf("validateProjectName(%q) = true, want false", name)
		}
		if name != "" && model.validationErr == "" {
			t.Errorf("validateProjectName(%q) should set error", name)
		}
	}
}

func TestRenderProjectName_ValidInput(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.width = 80
	model.textInput.SetValue("valid-project")
	model.validateProjectName("valid-project")

	view := model.renderProjectName()

	if len(view) == 0 {
		t.Error("renderProjectName() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(view, "Project Name") {
		t.Error("renderProjectName() should contain 'Project Name'")
	}

	if !strings.Contains(view, "valid-project") {
		t.Error("renderProjectName() should contain input value")
	}
}

func TestRenderProjectName_InvalidInput(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.width = 80
	model.textInput.SetValue("invalid project")
	model.validateProjectName("invalid project")

	view := model.renderProjectName()

	// Should show error message
	if !strings.Contains(view, "[X]") {
		t.Error("renderProjectName() with invalid input should contain error indicator")
	}

	if model.validationErr == "" {
		t.Error("renderProjectName() with invalid input should have validation error")
	}
}

func TestGetIDEInfo(t *testing.T) {
	ideInfo := getIDEInfo()

	if len(ideInfo) != 6 {
		t.Errorf("getIDEInfo() returned %d IDEs, want 6", len(ideInfo))
	}

	// Check that Cursor and Claude Code are recommended
	cursorRecommended := false
	claudeRecommended := false
	for _, ide := range ideInfo {
		if ide.Name == "Cursor" && ide.Recommended {
			cursorRecommended = true
		}
		if ide.Name == "Claude Code" && ide.Recommended {
			claudeRecommended = true
		}
	}

	if !cursorRecommended {
		t.Error("Cursor should be marked as recommended")
	}
	if !claudeRecommended {
		t.Error("Claude Code should be marked as recommended")
	}
}

func TestModel_Update_ArrowKeys_IDESelection(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	model.selectedIDEIndex = 0

	// Test down arrow
	msg := tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ := model.Update(msg)
	updatedModel := newModel.(Model)
	if updatedModel.selectedIDEIndex != 1 {
		t.Errorf("Update(Down arrow) selectedIDEIndex = %d, want 1", updatedModel.selectedIDEIndex)
	}

	// Test up arrow
	model.selectedIDEIndex = 1
	msg = tea.KeyMsg{Type: tea.KeyUp}
	newModel, _ = model.Update(msg)
	updatedModel = newModel.(Model)
	if updatedModel.selectedIDEIndex != 0 {
		t.Errorf("Update(Up arrow) selectedIDEIndex = %d, want 0", updatedModel.selectedIDEIndex)
	}

	// Test wrap to bottom
	model.selectedIDEIndex = 0
	msg = tea.KeyMsg{Type: tea.KeyUp}
	newModel, _ = model.Update(msg)
	updatedModel = newModel.(Model)
	ideInfo := getIDEInfo()
	expectedIndex := len(ideInfo) - 1
	if updatedModel.selectedIDEIndex != expectedIndex {
		t.Errorf("Update(Up arrow from top) selectedIDEIndex = %d, want %d", updatedModel.selectedIDEIndex, expectedIndex)
	}

	// Test wrap to top
	model.selectedIDEIndex = len(ideInfo) - 1
	msg = tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ = model.Update(msg)
	updatedModel = newModel.(Model)
	if updatedModel.selectedIDEIndex != 0 {
		t.Errorf("Update(Down arrow from bottom) selectedIDEIndex = %d, want 0", updatedModel.selectedIDEIndex)
	}
}

func TestModel_Update_EnterKey_IDESelection(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	model.selectedIDEIndex = 0
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	newModel, _ := model.Update(msg)

	updatedModel := newModel.(Model)
	ideInfo := getIDEInfo()
	expectedIDE := ideInfo[0].Name
	if updatedModel.selectedIDE != expectedIDE {
		t.Errorf("Update(Enter on IDE selection) selectedIDE = %q, want %q", updatedModel.selectedIDE, expectedIDE)
	}
}

func TestModel_View_IDESelection(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	model.width = 80

	view := model.View()

	if len(view) == 0 {
		t.Error("View() should return non-empty string for IDE selection screen")
	}

	// Check that view contains expected elements
	if !strings.Contains(view, "Select Your IDE") {
		t.Error("View() should contain 'Select Your IDE'")
	}

	if !strings.Contains(view, "[*]") {
		t.Error("View() should contain IDE selection indicator")
	}

	// Check for IDE names
	ideInfo := getIDEInfo()
	for _, ide := range ideInfo {
		if !strings.Contains(view, ide.Name) {
			t.Errorf("View() should contain IDE name %q", ide.Name)
		}
	}
}

func TestRenderIDESelection(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	model.width = 80
	model.selectedIDEIndex = 0

	view := model.renderIDESelection()

	if len(view) == 0 {
		t.Error("renderIDESelection() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(view, "Select Your IDE") {
		t.Error("renderIDESelection() should contain 'Select Your IDE'")
	}

	if !strings.Contains(view, "[*]") {
		t.Error("renderIDESelection() should contain IDE selection indicator")
	}

	// Check for checkbox selections
	if !strings.Contains(view, "☑") {
		t.Error("renderIDESelection() should contain selected checkbox (☑)")
	}

	// Check for recommended indicator
	if !strings.Contains(view, "[*]") {
		t.Error("renderIDESelection() should contain recommended indicator")
	}
}

func TestRenderIDESelection_SelectedItem(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	model.width = 80
	model.selectedIDEIndex = 1 // Select second IDE

	view := model.renderIDESelection()

	ideInfo := getIDEInfo()
	selectedIDE := ideInfo[1].Name

	// Check that selected IDE is in the view
	if !strings.Contains(view, selectedIDE) {
		t.Errorf("renderIDESelection() should contain selected IDE %q", selectedIDE)
	}
}

func TestGetGenerationSteps(t *testing.T) {
	steps := getGenerationSteps()

	if len(steps) != 6 {
		t.Errorf("getGenerationSteps() returned %d steps, want 6", len(steps))
	}

	// Check that all steps start as pending
	for i, step := range steps {
		if step.Status != stepPending {
			t.Errorf("getGenerationSteps() step %d (%q) status = %v, want %v", i, step.Name, step.Status, stepPending)
		}
	}
}

func TestModel_Update_EnterKey_IDESelection_TransitionsToGenerating(t *testing.T) {
	model := InitialModel()
	model.state = stateIDESelection
	model.selectedIDEIndex = 0
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	newModel, cmd := model.Update(msg)

	updatedModel := newModel.(Model)
	if updatedModel.state != stateGenerating {
		t.Errorf("Update(Enter on IDE selection) state = %v, want %v", updatedModel.state, stateGenerating)
	}

	// Check that steps are initialized
	if len(updatedModel.steps) == 0 {
		t.Error("Update(Enter on IDE selection) should initialize steps")
	}

	// Check that first step is in progress
	if len(updatedModel.steps) > 0 && updatedModel.steps[0].Status != stepInProgress {
		t.Errorf("Update(Enter on IDE selection) first step status = %v, want %v", updatedModel.steps[0].Status, stepInProgress)
	}

	// Check that spinner command is returned
	if cmd == nil {
		t.Error("Update(Enter on IDE selection) should return spinner command")
	}

	_ = cmd
}

func TestModel_View_Generating(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.width = 80
	model.steps = getGenerationSteps()

	view := model.View()

	if len(view) == 0 {
		t.Error("View() should return non-empty string for generating screen")
	}

	// Check that view contains expected elements
	if !strings.Contains(view, "Generating") {
		t.Error("View() should contain 'Generating'")
	}

	if !strings.Contains(view, "[...]") {
		t.Error("View() should contain generating indicator")
	}
}

func TestRenderGenerating(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.width = 80
	model.steps = getGenerationSteps()

	view := model.renderGenerating()

	if len(view) == 0 {
		t.Error("renderGenerating() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(view, "Generating") {
		t.Error("renderGenerating() should contain 'Generating'")
	}

	if !strings.Contains(view, "[...]") {
		t.Error("renderGenerating() should contain generating indicator")
	}

	// Check for step names
	for _, step := range model.steps {
		if !strings.Contains(view, step.Name) {
			t.Errorf("renderGenerating() should contain step name %q", step.Name)
		}
	}
}

func TestGetStepIcon(t *testing.T) {
	model := InitialModel()
	model.spinnerFrame = 0

	// Test completed
	step := GenerationStep{Name: "Test", Status: stepCompleted}
	icon := model.getStepIcon(step)
	if icon != "[+]" {
		t.Errorf("getStepIcon(completed) = %q, want %q", icon, "[+]")
	}

	// Test in progress
	step.Status = stepInProgress
	icon = model.getStepIcon(step)
	if !strings.Contains(icon, "[...]") {
		t.Errorf("getStepIcon(inProgress) should contain [...], got %q", icon)
	}

	// Test failed
	step.Status = stepFailed
	icon = model.getStepIcon(step)
	if icon != "[X]" {
		t.Errorf("getStepIcon(failed) = %q, want %q", icon, "[X]")
	}

	// Test pending
	step.Status = stepPending
	icon = model.getStepIcon(step)
	if icon != "[-]" {
		t.Errorf("getStepIcon(pending) = %q, want %q", icon, "[-]")
	}
}

func TestGetStepColor(t *testing.T) {
	model := InitialModel()

	// Test completed
	step := GenerationStep{Name: "Test", Status: stepCompleted}
	color := model.getStepColor(step)
	if color != primary {
		t.Errorf("getStepColor(completed) = %v, want %v", color, primary)
	}

	// Test in progress
	step.Status = stepInProgress
	color = model.getStepColor(step)
	if color != secondary {
		t.Errorf("getStepColor(inProgress) = %v, want %v", color, secondary)
	}

	// Test failed
	step.Status = stepFailed
	color = model.getStepColor(step)
	if color != primary {
		t.Errorf("getStepColor(failed) = %v, want %v", color, primary)
	}

	// Test pending
	step.Status = stepPending
	color = model.getStepColor(step)
	if color != tertiary {
		t.Errorf("getStepColor(pending) = %v, want %v", color, tertiary)
	}
}

func TestCalculateProgress(t *testing.T) {
	model := InitialModel()

	// Test empty steps
	model.steps = []GenerationStep{}
	progress := model.calculateProgress()
	if progress != 0 {
		t.Errorf("calculateProgress() with empty steps = %d, want 0", progress)
	}

	// Test all pending
	model.steps = getGenerationSteps()
	progress = model.calculateProgress()
	if progress != 0 {
		t.Errorf("calculateProgress() with all pending = %d, want 0", progress)
	}

	// Test half completed
	model.steps = []GenerationStep{
		{Name: "Step 1", Status: stepCompleted},
		{Name: "Step 2", Status: stepCompleted},
		{Name: "Step 3", Status: stepPending},
		{Name: "Step 4", Status: stepPending},
	}
	progress = model.calculateProgress()
	if progress != 50 {
		t.Errorf("calculateProgress() with 2/4 completed = %d, want 50", progress)
	}

	// Test all completed
	for i := range model.steps {
		model.steps[i].Status = stepCompleted
	}
	progress = model.calculateProgress()
	if progress != 100 {
		t.Errorf("calculateProgress() with all completed = %d, want 100", progress)
	}
}

func TestRenderProgressBar(t *testing.T) {
	model := InitialModel()

	// Test 0%
	bar := model.renderProgressBar(0, 10)
	if !strings.Contains(bar, "0%") {
		t.Error("renderProgressBar(0%) should contain '0%'")
	}

	// Test 50%
	bar = model.renderProgressBar(50, 10)
	if !strings.Contains(bar, "50%") {
		t.Error("renderProgressBar(50%) should contain '50%'")
	}

	// Test 100%
	bar = model.renderProgressBar(100, 10)
	if !strings.Contains(bar, "100%") {
		t.Error("renderProgressBar(100%) should contain '100%'")
	}
}

func TestGetSpinnerChar(t *testing.T) {
	model := InitialModel()

	// Test spinner cycles through characters
	chars := make(map[string]bool)
	for i := 0; i < 20; i++ {
		model.spinnerFrame = i
		char := model.getSpinnerChar()
		chars[char] = true
	}

	// Should have multiple different characters (ASCII spinner has 4: |, /, -, \)
	if len(chars) < 4 {
		t.Errorf("getSpinnerChar() should cycle through multiple characters, got %d unique", len(chars))
	}
}

func TestModel_Update_SpinnerTick(t *testing.T) {
	model := InitialModel()
	model.state = stateGenerating
	model.spinnerFrame = 0

	msg := time.Now()
	newModel, cmd := model.Update(msg)

	updatedModel := newModel.(Model)
	if updatedModel.spinnerFrame <= model.spinnerFrame {
		t.Errorf("Update(spinner tick) spinnerFrame = %d, want > %d", updatedModel.spinnerFrame, model.spinnerFrame)
	}

	// Should return spinner command to continue animation
	if cmd == nil {
		t.Error("Update(spinner tick) should return spinner command")
	}

	_ = cmd
}

func TestModel_View_Success(t *testing.T) {
	model := InitialModel()
	model.state = stateSuccess
	model.width = 80
	model.projectName = "test-project"
	model.selectedIDE = "Cursor"

	view := model.View()

	if len(view) == 0 {
		t.Error("View() should return non-empty string for success screen")
	}

	// Check that view contains expected elements
	viewLower := strings.ToLower(view)
	if !strings.Contains(viewLower, "successfully") {
		t.Error("View() should contain 'successfully'")
	}

	// Check for success indicators (text-based, not symbols)
	successIndicators := []string{"successfully", "project created", "ready to go"}
	found := false
	for _, indicator := range successIndicators {
		if strings.Contains(viewLower, indicator) {
			found = true
			break
		}
	}
	if !found {
		t.Error("View() should contain success indicator")
	}
}

func TestRenderSuccess(t *testing.T) {
	model := InitialModel()
	model.state = stateSuccess
	model.width = 80
	model.projectName = "test-project"
	model.selectedIDE = "Cursor"

	view := model.renderSuccess()

	if len(view) == 0 {
		t.Error("renderSuccess() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(view, "Project created successfully") {
		t.Error("renderSuccess() should contain 'Project created successfully'")
	}

	if !strings.Contains(view, "test-project") {
		t.Error("renderSuccess() should contain project name")
	}

	if !strings.Contains(view, "/tell") {
		t.Error("renderSuccess() should contain '/tell' instruction")
	}
}

func TestGetProjectStructureTree(t *testing.T) {
	model := InitialModel()
	model.projectName = "my-awesome-app"

	tree := model.getProjectStructureTree()

	if len(tree) == 0 {
		t.Error("getProjectStructureTree() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(tree, "my-awesome-app") {
		t.Error("getProjectStructureTree() should contain project name")
	}

	if !strings.Contains(tree, ".cursor/") {
		t.Error("getProjectStructureTree() should contain .cursor directory")
	}

	if !strings.Contains(tree, ".do/") {
		t.Error("getProjectStructureTree() should contain .do directory")
	}

	if !strings.Contains(tree, ".github/") {
		t.Error("getProjectStructureTree() should contain .github directory")
	}

	if !strings.Contains(tree, "[+]") {
		t.Error("getProjectStructureTree() should contain success indicator")
	}
}

func TestRenderSuccess_NextSteps(t *testing.T) {
	model := InitialModel()
	model.state = stateSuccess
	model.width = 80
	model.projectName = "my-project"
	model.selectedIDEs = map[string]bool{"Cursor": true}
	model.selectedIDE = "Cursor"

	view := model.renderSuccess()

	// Check for next steps
	if !strings.Contains(view, "now open my-project inside Cursor") {
		t.Error("renderSuccess() should contain 'now open my-project inside Cursor'")
	}

	if !strings.Contains(view, "/tell") {
		t.Error("renderSuccess() should contain '/tell' instruction")
	}
}

func TestRenderSuccess_DifferentIDEs(t *testing.T) {
	testCases := []struct {
		ide      string
		checkFor string
	}{
		{"Cursor", "Cursor"},
		{"Claude Code", "Claude Code"},
		{"Antigravity", "Antigravity"},
		{"Windsurf", "Windsurf"},
		{"Cline", "Cline"},
		{"OpenCode", "OpenCode"},
		{"", "your IDE"}, // Default
	}

	for _, tc := range testCases {
		model := InitialModel()
		model.state = stateSuccess
		model.width = 80
		model.projectName = "test-project"
		model.selectedIDEs = map[string]bool{}
		if tc.ide != "" {
			model.selectedIDEs[tc.ide] = true
		}
		model.selectedIDE = tc.ide

		view := model.renderSuccess()

		if !strings.Contains(view, tc.checkFor) {
			t.Errorf("renderSuccess() with IDE %q should contain %q", tc.ide, tc.checkFor)
		}
	}
}

func TestModel_Update_EnterKey_Success(t *testing.T) {
	model := InitialModel()
	model.state = stateSuccess
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	newModel, cmd := model.Update(msg)

	// Enter on success screen should quit
	if cmd == nil {
		t.Error("Update(Enter on success) should return quit command")
	}

	_ = newModel
}

func TestRenderSuccess_Checkmarks(t *testing.T) {
	model := InitialModel()
	model.state = stateSuccess
	model.width = 80
	model.projectName = "test-project"

	view := model.renderSuccess()

	// Should contain success indicators - check for success-related text
	successIndicators := []string{
		"successfully",
		"Project created",
		"ready to go",
	}

	found := false
	viewLower := strings.ToLower(view)
	for _, indicator := range successIndicators {
		if strings.Contains(viewLower, strings.ToLower(indicator)) {
			found = true
			break
		}
	}

	if !found {
		t.Error("renderSuccess() should contain at least one success indicator")
	}
}

func TestModel_SetError(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.setError("Test error", "Test suggestion")

	if model.state != stateError {
		t.Errorf("setError() state = %v, want %v", model.state, stateError)
	}

	if model.errorMessage != "Test error" {
		t.Errorf("setError() errorMessage = %q, want %q", model.errorMessage, "Test error")
	}

	if model.errorSuggestion != "Test suggestion" {
		t.Errorf("setError() errorSuggestion = %q, want %q", model.errorSuggestion, "Test suggestion")
	}

	if model.previousState != stateProjectName {
		t.Errorf("setError() previousState = %v, want %v", model.previousState, stateProjectName)
	}
}

func TestModel_View_Error(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.width = 80
	model.errorMessage = "Directory 'my-project' already exists"
	model.errorSuggestion = "Choose a different name or delete the existing directory"

	view := model.View()

	if len(view) == 0 {
		t.Error("View() should return non-empty string for error screen")
	}

	// Check that view contains expected elements
	if !strings.Contains(view, "Error") {
		t.Error("View() should contain 'Error'")
	}

	if !strings.Contains(view, "[X]") {
		t.Error("View() should contain error indicator")
	}
}

func TestRenderError(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.width = 80
	model.errorMessage = "Directory 'my-project' already exists"
	model.errorSuggestion = "Choose a different name or delete the existing directory"

	view := model.renderError()

	if len(view) == 0 {
		t.Error("renderError() should return non-empty string")
	}

	// Check for key elements
	if !strings.Contains(view, "Error") {
		t.Error("renderError() should contain 'Error'")
	}

	if !strings.Contains(view, "[X]") {
		t.Error("renderError() should contain error indicator")
	}

	if !strings.Contains(view, "my-project") {
		t.Error("renderError() should contain error message")
	}

	if !strings.Contains(view, "Choose a different name") {
		t.Error("renderError() should contain suggestion")
	}
}

func TestRenderError_RecoveryOptions(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.width = 80
	model.errorMessage = "Test error"
	model.errorSuggestion = "Test suggestion"

	view := model.renderError()

	// Check for recovery options
	if !strings.Contains(view, "retry") {
		t.Error("renderError() should contain 'retry' option")
	}

	if !strings.Contains(view, "go back") {
		t.Error("renderError() should contain 'go back' option")
	}

	if !strings.Contains(view, "quit") {
		t.Error("renderError() should contain 'quit' option")
	}
}

func TestModel_Update_RetryKey(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.previousState = stateProjectName
	model.errorMessage = "Test error"
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}}

	newModel, _ := model.Update(msg)

	updatedModel := newModel.(Model)
	if updatedModel.state != stateProjectName {
		t.Errorf("Update('r' key) state = %v, want %v", updatedModel.state, stateProjectName)
	}

	if updatedModel.errorMessage != "" {
		t.Error("Update('r' key) should clear error message")
	}
}

func TestModel_Update_BackKey(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.previousState = stateIDESelection
	model.errorMessage = "Test error"
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'b'}}

	newModel, _ := model.Update(msg)

	updatedModel := newModel.(Model)
	if updatedModel.state != stateWelcome {
		t.Errorf("Update('b' key) state = %v, want %v", updatedModel.state, stateWelcome)
	}

	if updatedModel.errorMessage != "" {
		t.Error("Update('b' key) should clear error message")
	}
}

func TestRenderError_WithoutSuggestion(t *testing.T) {
	model := InitialModel()
	model.state = stateError
	model.width = 80
	model.errorMessage = "Test error"
	model.errorSuggestion = "" // No suggestion

	view := model.renderError()

	// Should still render error message
	if !strings.Contains(view, "Test error") {
		t.Error("renderError() should contain error message even without suggestion")
	}

	// Should not contain suggestion icon if no suggestion
	if strings.Contains(view, "💡") {
		t.Error("renderError() should not contain suggestion icon when suggestion is empty")
	}
}

func TestRenderError_CommonErrors(t *testing.T) {
	testCases := []struct {
		name       string
		message    string
		suggestion string
	}{
		{
			name:       "Directory exists",
			message:    "Directory 'my-project' already exists",
			suggestion: "Choose a different name or delete the existing directory",
		},
		{
			name:       "Permission denied",
			message:    "Permission denied: cannot create directory",
			suggestion: "Check your file permissions or run with appropriate privileges",
		},
		{
			name:       "Invalid path",
			message:    "Invalid path: contains invalid characters",
			suggestion: "Use only alphanumeric characters, hyphens, and underscores",
		},
		{
			name:       "Network error",
			message:    "Failed to download dependencies",
			suggestion: "Check your internet connection and try again",
		},
	}

	for _, tc := range testCases {
		model := InitialModel()
		model.state = stateError
		model.width = 80
		model.errorMessage = tc.message
		model.errorSuggestion = tc.suggestion

		view := model.renderError()

		if !strings.Contains(view, tc.message) {
			t.Errorf("renderError() for %s should contain error message", tc.name)
		}

		if !strings.Contains(view, tc.suggestion) {
			t.Errorf("renderError() for %s should contain suggestion", tc.name)
		}
	}
}

func TestCanTransitionTo(t *testing.T) {
	// Test valid transitions
	validTransitions := []struct {
		from wizardState
		to   wizardState
		want bool
	}{
		{stateWelcome, stateProjectName, true},
		{stateProjectName, stateIDESelection, true},
		{stateIDESelection, stateGenerating, true},
		{stateGenerating, stateSuccess, true},
		{stateWelcome, stateError, true},
		{stateProjectName, stateError, true},
		{stateIDESelection, stateError, true},
		{stateGenerating, stateError, true},
		{stateError, stateWelcome, true},
		{stateError, stateProjectName, true},
	}

	for _, tt := range validTransitions {
		if got := canTransitionTo(tt.from, tt.to); got != tt.want {
			t.Errorf("canTransitionTo(%v, %v) = %v, want %v", tt.from, tt.to, got, tt.want)
		}
	}

	// Test invalid transitions
	invalidTransitions := []struct {
		from wizardState
		to   wizardState
		want bool
	}{
		{stateWelcome, stateIDESelection, false}, // Can't skip steps
		{stateProjectName, stateGenerating, false},
		{stateSuccess, stateWelcome, false},        // Terminal state
		{stateGenerating, stateProjectName, false}, // Can't go backwards directly
	}

	for _, tt := range invalidTransitions {
		if got := canTransitionTo(tt.from, tt.to); got != tt.want {
			t.Errorf("canTransitionTo(%v, %v) = %v, want %v", tt.from, tt.to, got, tt.want)
		}
	}
}

func TestTransitionToState(t *testing.T) {
	model := InitialModel()
	model.state = stateWelcome

	// Test valid transition
	if !model.transitionToState(stateProjectName) {
		t.Error("transitionToState() should succeed for valid transition")
	}
	if model.state != stateProjectName {
		t.Errorf("transitionToState() state = %v, want %v", model.state, stateProjectName)
	}
	if len(model.stateHistory) != 2 {
		t.Errorf("transitionToState() stateHistory length = %d, want 2", len(model.stateHistory))
	}

	// Test invalid transition
	model.state = stateWelcome
	if model.transitionToState(stateIDESelection) {
		t.Error("transitionToState() should fail for invalid transition")
	}
	if model.state != stateWelcome {
		t.Errorf("transitionToState() should not change state for invalid transition")
	}
}

func TestGoBack(t *testing.T) {
	model := InitialModel()
	model.state = stateWelcome
	model.stateHistory = []wizardState{stateWelcome, stateProjectName, stateIDESelection}
	model.state = stateIDESelection

	// Test going back
	if !model.goBack() {
		t.Error("goBack() should succeed when history exists")
	}
	if model.state != stateProjectName {
		t.Errorf("goBack() state = %v, want %v", model.state, stateProjectName)
	}
	if len(model.stateHistory) != 2 {
		t.Errorf("goBack() stateHistory length = %d, want 2", len(model.stateHistory))
	}

	// Test going back from first state
	model.stateHistory = []wizardState{stateWelcome}
	model.state = stateWelcome
	if model.goBack() {
		t.Error("goBack() should fail when at first state")
	}
}

func TestGetCurrentStepNumber(t *testing.T) {
	testCases := []struct {
		state wizardState
		want  int
	}{
		{stateWelcome, 0},
		{stateProjectName, 1},
		{stateIDESelection, 2},
		{stateGenerating, 3},
		{stateSuccess, 4},
		{stateError, 0},
	}

	for _, tc := range testCases {
		model := InitialModel()
		model.state = tc.state
		if got := model.getCurrentStepNumber(); got != tc.want {
			t.Errorf("getCurrentStepNumber() for state %v = %d, want %d", tc.state, got, tc.want)
		}
	}
}

func TestGetTotalSteps(t *testing.T) {
	if got := getTotalSteps(); got != 4 {
		t.Errorf("getTotalSteps() = %d, want 4", got)
	}
}

func TestModel_Update_BackNavigation(t *testing.T) {
	model := InitialModel()
	model.state = stateProjectName
	model.stateHistory = []wizardState{stateWelcome, stateProjectName}

	// Test Esc key
	msg := tea.KeyMsg{Type: tea.KeyEsc}
	newModel, _ := model.Update(msg)

	updatedModel := newModel.(Model)
	if updatedModel.state != stateWelcome {
		t.Errorf("Update(Esc) state = %v, want %v", updatedModel.state, stateWelcome)
	}

	// Test Backspace key
	model.state = stateIDESelection
	model.stateHistory = []wizardState{stateWelcome, stateProjectName, stateIDESelection}
	msg = tea.KeyMsg{Type: tea.KeyBackspace}
	newModel, _ = model.Update(msg)

	updatedModel = newModel.(Model)
	if updatedModel.state != stateProjectName {
		t.Errorf("Update(Backspace) state = %v, want %v", updatedModel.state, stateProjectName)
	}
}

func TestFullWizardFlow(t *testing.T) {
	model := InitialModel()

	// Step 1: Welcome → Project Name
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := model.Update(msg)
	model = newModel.(Model)
	if model.state != stateProjectName {
		t.Errorf("Step 1: state = %v, want %v", model.state, stateProjectName)
	}

	// Step 2: Enter project name
	model.textInput.SetValue("test-project")
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)
	if model.state != stateIDESelection {
		t.Errorf("Step 2: state = %v, want %v", model.state, stateIDESelection)
	}
	if model.projectName != "test-project" {
		t.Errorf("Step 2: projectName = %q, want %q", model.projectName, "test-project")
	}

	// Step 3: Select IDE
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ = model.Update(msg)
	model = newModel.(Model)
	if model.state != stateGenerating {
		t.Errorf("Step 3: state = %v, want %v", model.state, stateGenerating)
	}

	// Step 4: Simulate generation completion by sending generationCompleteMsg
	// In real usage, this would be sent by the async generation goroutine
	genCompleteMsg := generationCompleteMsg{err: nil}
	newModel, _ = model.Update(genCompleteMsg)
	model = newModel.(Model)
	if model.state != stateSuccess {
		t.Errorf("Step 4: state = %v, want %v", model.state, stateSuccess)
	}
}

func TestRenderProgressIndicator(t *testing.T) {
	testCases := []struct {
		state wizardState
		want  bool // Should show progress indicator
	}{
		{stateWelcome, false},
		{stateProjectName, true},
		{stateIDESelection, true},
		{stateGenerating, true},
		{stateSuccess, false},
		{stateError, false},
	}

	for _, tc := range testCases {
		model := InitialModel()
		model.state = tc.state

		indicator := model.renderProgressIndicator()
		hasIndicator := indicator != ""

		if hasIndicator != tc.want {
			t.Errorf("renderProgressIndicator() for state %v = %v, want %v", tc.state, hasIndicator, tc.want)
		}
	}
}

func TestToProjectRequest(t *testing.T) {
	model := InitialModel()
	model.projectName = "test-project"
	model.selectedIDEs = map[string]bool{"Cursor": true}

	request := model.toProjectRequest()
	if request == nil {
		t.Fatal("toProjectRequest() should not return nil")
	}

	if request.ProjectName != "test-project" {
		t.Errorf("toProjectRequest() ProjectName = %q, want %q", request.ProjectName, "test-project")
	}

	if request.IDE != "Cursor" {
		t.Errorf("toProjectRequest() IDE = %q, want %q", request.IDE, "Cursor")
	}
	if len(request.IDEs) != 1 || request.IDEs[0] != "Cursor" {
		t.Errorf("toProjectRequest() IDEs = %v, want [Cursor]", request.IDEs)
	}

	if request.ProjectType != "Fullstack" {
		t.Errorf("toProjectRequest() ProjectType = %q, want %q", request.ProjectType, "Fullstack")
	}
}

func TestToProjectRequest_Validation(t *testing.T) {
	model := InitialModel()
	model.projectName = "test-project"
	model.selectedIDEs = map[string]bool{"Cursor": true}

	request := model.toProjectRequest()

	// Validate the request
	if err := request.Validate(); err != nil {
		t.Errorf("toProjectRequest() should return valid request, got error: %v", err)
	}
}

func TestToProjectRequest_EmptyFields(t *testing.T) {
	model := InitialModel()
	// Leave fields empty

	request := model.toProjectRequest()

	if request.ProjectName != "" {
		t.Errorf("toProjectRequest() ProjectName = %q, want empty", request.ProjectName)
	}

	if request.IDE != "Cursor" {
		t.Errorf("toProjectRequest() IDE = %q, want %q", request.IDE, "Cursor")
	}

	// Should still have default project type
	if request.ProjectType != "Fullstack" {
		t.Errorf("toProjectRequest() ProjectType = %q, want %q", request.ProjectType, "Fullstack")
	}
}

func TestToProjectRequest_AllIDEs(t *testing.T) {
	ideInfo := getIDEInfo()

	for _, ide := range ideInfo {
		model := InitialModel()
		model.projectName = "test-project"
		model.selectedIDEs = map[string]bool{ide.Name: true}

		request := model.toProjectRequest()

		if request.IDE != ide.Name {
			t.Errorf("toProjectRequest() with IDE %q = %q, want %q", ide.Name, request.IDE, ide.Name)
		}
		if len(request.IDEs) != 1 || request.IDEs[0] != ide.Name {
			t.Errorf("toProjectRequest() with IDE %q returned IDEs %v", ide.Name, request.IDEs)
		}

		// Should validate successfully
		if err := request.Validate(); err != nil {
			t.Errorf("toProjectRequest() with IDE %q should be valid, got error: %v", ide.Name, err)
		}
	}
}
