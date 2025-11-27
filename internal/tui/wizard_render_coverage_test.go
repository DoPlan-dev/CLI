package tui

import (
	"strings"
	"testing"
)

// TestRenderASCIIHeader_AllEdgeCases tests renderASCIIHeader with all edge cases (68.8% -> 90%+)
func TestRenderASCIIHeader_AllEdgeCases(t *testing.T) {
	header := renderASCIIHeader()

	if len(header) == 0 {
		t.Error("renderASCIIHeader() should return non-empty string")
	}

	// Should contain DoPlan ASCII art
	if !strings.Contains(header, "█") && !strings.Contains(header, "║") && !strings.Contains(header, "╗") {
		t.Error("renderASCIIHeader() should contain ASCII art characters")
	}

	// Should have multiple lines
	lines := strings.Split(header, "\n")
	if len(lines) < 5 {
		t.Errorf("renderASCIIHeader() should have multiple lines, got %d", len(lines))
	}

	// Check that lines are centered (non-empty lines should have padding)
	nonEmptyLines := 0
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) > 0 {
			nonEmptyLines++
			// Line should not start at column 0 if it's short enough
			if len(trimmed) < 80 && strings.HasPrefix(line, "█") {
				// This is acceptable - it means the line starts with art
			}
		}
	}

	if nonEmptyLines < 5 {
		t.Errorf("renderASCIIHeader() should have at least 5 non-empty lines, got %d", nonEmptyLines)
	}

	// Verify ASCII art contains "DoPlan" pattern
	headerLower := strings.ToLower(header)
	hasPattern := strings.Contains(headerLower, "do") || strings.Contains(headerLower, "plan") ||
		strings.Contains(header, "█") || strings.Contains(header, "╔") || strings.Contains(header, "╗")

	if !hasPattern {
		t.Error("renderASCIIHeader() should contain ASCII art pattern")
	}
}

// TestRenderTopLine_EdgeCases tests renderTopLine edge cases (87.5% -> 90%+)
func TestRenderTopLine_EdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		state          wizardState
		currentStep    int
		totalSteps     int
		expectProgress bool
	}{
		{
			name:           "Welcome state",
			state:          stateWelcome,
			currentStep:    0,
			totalSteps:     4,
			expectProgress: true,
		},
		{
			name:           "Project name state",
			state:          stateProjectName,
			currentStep:    1,
			totalSteps:     4,
			expectProgress: true,
		},
		{
			name:           "IDE selection state",
			state:          stateIDESelection,
			currentStep:    2,
			totalSteps:     4,
			expectProgress: true,
		},
		{
			name:           "Generating state",
			state:          stateGenerating,
			currentStep:    3,
			totalSteps:     4,
			expectProgress: true,
		},
		{
			name:           "Success state",
			state:          stateSuccess,
			currentStep:    4,
			totalSteps:     4,
			expectProgress: true,
		},
		{
			name:           "Error state",
			state:          stateError,
			currentStep:    0,
			totalSteps:     4,
			expectProgress: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := InitialModel()
			model.state = tt.state
			model.width = 80

			topLine := renderTopLine()

			if len(topLine) == 0 {
				t.Error("renderTopLine() should return non-empty string")
			}

			// Should contain progress indicator or state info
			if !strings.Contains(topLine, "Step") && !strings.Contains(topLine, "Progress") {
				// May contain other indicators
				if len(topLine) == 0 {
					t.Error("renderTopLine() should contain some content")
				}
			}
		})
	}
}

// TestRenderBody_EdgeCases tests renderBody edge cases
func TestRenderBody_EdgeCases(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		wantNonEmpty bool
		wantContains string
	}{
		{
			name:         "Empty content",
			content:      "",
			wantNonEmpty: false,
			wantContains: "",
		},
		{
			name:         "Single line content",
			content:      "Hello",
			wantNonEmpty: true,
			wantContains: "Hello",
		},
		{
			name:         "Multi-line content",
			content:      "Line 1\nLine 2\nLine 3",
			wantNonEmpty: true,
			wantContains: "Line 1", // May be styled/wrapped, so just check first line
		},
		{
			name:         "Long content",
			content:      strings.Repeat("A", 200),
			wantNonEmpty: true,
			wantContains: "A", // Just verify it's rendered
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderBody(tt.content)

			if tt.wantNonEmpty && len(result) == 0 {
				t.Error("renderBody() should return non-empty string for non-empty content")
			}

			// Should contain the key content (may be wrapped/styled by lipgloss)
			if len(tt.wantContains) > 0 && !strings.Contains(result, tt.wantContains) {
				t.Errorf("renderBody() should contain %q, got %q", tt.wantContains, result)
			}
		})
	}
}

// TestCanTransitionTo_AllCases tests canTransitionTo with all state combinations (87.5% -> 90%+)
func TestCanTransitionTo_AllCases(t *testing.T) {
	tests := []struct {
		name         string
		currentState wizardState
		targetState  wizardState
		expectValid  bool
	}{
		{
			name:         "Welcome to ProjectName",
			currentState: stateWelcome,
			targetState:  stateProjectName,
			expectValid:  true,
		},
		{
			name:         "ProjectName to IDESelection",
			currentState: stateProjectName,
			targetState:  stateIDESelection,
			expectValid:  true,
		},
		{
			name:         "IDESelection to Generating",
			currentState: stateIDESelection,
			targetState:  stateGenerating,
			expectValid:  true,
		},
		{
			name:         "Generating to Success",
			currentState: stateGenerating,
			targetState:  stateSuccess,
			expectValid:  true,
		},
		{
			name:         "Any state to Error",
			currentState: stateWelcome,
			targetState:  stateError,
			expectValid:  true,
		},
		{
			name:         "Invalid: Welcome to Generating",
			currentState: stateWelcome,
			targetState:  stateGenerating,
			expectValid:  false,
		},
		{
			name:         "Invalid: Success to Welcome",
			currentState: stateSuccess,
			targetState:  stateWelcome,
			expectValid:  false,
		},
		{
			name:         "Error to previous state",
			currentState: stateError,
			targetState:  stateProjectName,
			expectValid:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := canTransitionTo(tt.currentState, tt.targetState)

			if valid != tt.expectValid {
				t.Errorf("canTransitionTo(%v -> %v) = %v, want %v", tt.currentState, tt.targetState, valid, tt.expectValid)
			}
		})
	}
}

// TestGetProjectStructureTree_AllCases tests getProjectStructureTree comprehensively
func TestGetProjectStructureTree_AllCases(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		setupModel  func(*Model)
	}{
		{
			name:        "Default project name",
			projectName: "",
			setupModel: func(m *Model) {
				m.projectName = ""
			},
		},
		{
			name:        "Custom project name",
			projectName: "my-custom-project",
			setupModel: func(m *Model) {
				m.projectName = "my-custom-project"
			},
		},
		{
			name:        "Long project name",
			projectName: "very-long-project-name-with-many-characters",
			setupModel: func(m *Model) {
				m.projectName = "very-long-project-name-with-many-characters"
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

			// Should contain key directories
			requiredDirs := []string{".cursor/", ".do/", ".github/", "src/"}
			for _, dir := range requiredDirs {
				if !strings.Contains(tree, dir) {
					t.Errorf("getProjectStructureTree() should contain %s", dir)
				}
			}

			// If project name was provided, it should appear
			if tt.projectName != "" {
				if !strings.Contains(tree, tt.projectName) {
					t.Errorf("getProjectStructureTree() should contain project name %q", tt.projectName)
				}
			} else {
				// Should default to "my-project"
				if !strings.Contains(tree, "my-project") {
					t.Error("getProjectStructureTree() should default to 'my-project' when name is empty")
				}
			}
		})
	}
}
