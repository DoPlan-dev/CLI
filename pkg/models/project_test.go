package models

import (
	"encoding/json"
	"testing"
)

func TestProjectRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request ProjectRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			request: ProjectRequest{
				ProjectName: "my-awesome-app",
				IDE:         "Cursor",
				ProjectType: "Fullstack",
			},
			wantErr: false,
		},
		{
			name: "valid request with default project type",
			request: ProjectRequest{
				ProjectName: "test-project",
				IDE:         "Claude Code",
			},
			wantErr: false,
		},
		{
			name: "empty project name",
			request: ProjectRequest{
				ProjectName: "",
				IDE:         "Cursor",
			},
			wantErr: true,
			errMsg:  "project name is required",
		},
		{
			name: "invalid project name with spaces",
			request: ProjectRequest{
				ProjectName: "my awesome app",
				IDE:         "Cursor",
			},
			wantErr: true,
			errMsg:  "invalid",
		},
		{
			name: "invalid project name with special characters",
			request: ProjectRequest{
				ProjectName: "my@project",
				IDE:         "Cursor",
			},
			wantErr: true,
			errMsg:  "invalid",
		},
		{
			name: "empty IDE",
			request: ProjectRequest{
				ProjectName: "test-project",
				IDE:         "",
			},
			wantErr: true,
			errMsg:  "at least one IDE",
		},
		{
			name: "valid with underscores",
			request: ProjectRequest{
				ProjectName: "my_awesome_app",
				IDE:         "Cursor",
			},
			wantErr: false,
		},
		{
			name: "valid with numbers",
			request: ProjectRequest{
				ProjectName: "project123",
				IDE:         "Cursor",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != "" {
				if err == nil || !contains(err.Error(), tt.errMsg) {
					t.Errorf("ProjectRequest.Validate() error = %v, want error containing %v", err, tt.errMsg)
				}
			}
			// Check default project type is set
			if !tt.wantErr && tt.request.ProjectType == "" {
				if tt.request.ProjectType != "Fullstack" {
					t.Errorf("ProjectRequest.Validate() should set default ProjectType to 'Fullstack', got %v", tt.request.ProjectType)
				}
			}
		})
	}
}

func TestProjectRequestValidateIDEs(t *testing.T) {
	req := ProjectRequest{
		ProjectName: "multi-ide",
		IDEs:        []string{"Cursor", "cursor", "Claude Code"},
	}
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate() returned error for multi IDEs: %v", err)
	}
	if len(req.IDEs) != 2 {
		t.Fatalf("expected 2 normalized IDEs, got %v", req.IDEs)
	}
	if req.IDEs[0] != "Cursor" || req.IDEs[1] != "Claude Code" {
		t.Errorf("unexpected IDE order: %v", req.IDEs)
	}
	if req.IDE != "Cursor" {
		t.Errorf("primary IDE should be first entry, got %q", req.IDE)
	}

	invalid := ProjectRequest{
		ProjectName: "bad-ide",
		IDEs:        []string{"UnknownIDE"},
	}
	if err := invalid.Validate(); err == nil {
		t.Fatal("expected error for unsupported IDE")
	}
}

func TestProjectState_Validate(t *testing.T) {
	tests := []struct {
		name    string
		state   ProjectState
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid state - idea phase",
			state: ProjectState{
				Phase:      PhaseIdea,
				ActiveTask: nil,
				Completed:  []int{},
				Locked:     false,
			},
			wantErr: false,
		},
		{
			name: "valid state - building phase with active task",
			state: ProjectState{
				Phase:      PhaseBuilding,
				ActiveTask: intPtr(1),
				Completed:  []int{1, 2, 3},
				Locked:     true,
			},
			wantErr: false,
		},
		{
			name: "empty phase",
			state: ProjectState{
				Phase: "",
			},
			wantErr: true,
			errMsg:  "phase is required",
		},
		{
			name: "invalid phase",
			state: ProjectState{
				Phase: "invalid-phase",
			},
			wantErr: true,
			errMsg:  "invalid phase",
		},
		{
			name: "negative completed task ID",
			state: ProjectState{
				Phase:     PhaseTasks,
				Completed: []int{-1},
			},
			wantErr: true,
			errMsg:  "cannot be negative",
		},
		{
			name: "negative active task ID",
			state: ProjectState{
				Phase:      PhaseBuilding,
				ActiveTask: intPtr(-1),
			},
			wantErr: true,
			errMsg:  "cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.state.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectState.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != "" {
				if err == nil || !contains(err.Error(), tt.errMsg) {
					t.Errorf("ProjectState.Validate() error = %v, want error containing %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestNewProjectState(t *testing.T) {
	state := NewProjectState(PhaseIdea)

	if state.Phase != PhaseIdea {
		t.Errorf("NewProjectState() Phase = %v, want %v", state.Phase, PhaseIdea)
	}

	if state.ActiveTask != nil {
		t.Errorf("NewProjectState() ActiveTask = %v, want nil", state.ActiveTask)
	}

	if len(state.Completed) != 0 {
		t.Errorf("NewProjectState() Completed = %v, want empty slice", state.Completed)
	}

	if state.Locked {
		t.Errorf("NewProjectState() Locked = %v, want false", state.Locked)
	}
}

func TestProjectState_IsTaskCompleted(t *testing.T) {
	state := &ProjectState{
		Completed: []int{1, 2, 3},
	}

	tests := []struct {
		taskID   int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{0, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := state.IsTaskCompleted(tt.taskID); got != tt.expected {
				t.Errorf("ProjectState.IsTaskCompleted(%d) = %v, want %v", tt.taskID, got, tt.expected)
			}
		})
	}
}

func TestProjectState_MarkTaskCompleted(t *testing.T) {
	state := NewProjectState(PhaseTasks)

	// Mark task 1 as completed
	state.MarkTaskCompleted(1)
	if !state.IsTaskCompleted(1) {
		t.Error("MarkTaskCompleted(1) should mark task 1 as completed")
	}

	// Mark task 1 again (should not duplicate)
	initialLen := len(state.Completed)
	state.MarkTaskCompleted(1)
	if len(state.Completed) != initialLen {
		t.Error("MarkTaskCompleted() should not add duplicate task IDs")
	}

	// Mark task 2 as completed
	state.MarkTaskCompleted(2)
	if !state.IsTaskCompleted(2) {
		t.Error("MarkTaskCompleted(2) should mark task 2 as completed")
	}
}

func TestProjectState_SetActiveTask(t *testing.T) {
	state := NewProjectState(PhaseBuilding)

	// Set active task
	taskID := 5
	state.SetActiveTask(&taskID)
	if state.ActiveTask == nil || *state.ActiveTask != 5 {
		t.Error("SetActiveTask() should set the active task")
	}

	// Clear active task
	state.ClearActiveTask()
	if state.ActiveTask != nil {
		t.Error("ClearActiveTask() should clear the active task")
	}
}

func TestProjectState_ToJSON(t *testing.T) {
	state := &ProjectState{
		Phase:      PhaseTasks,
		ActiveTask: intPtr(1),
		Completed:  []int{1, 2},
		Locked:     true,
	}

	jsonData, err := state.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("ToJSON() should return non-empty JSON")
	}

	// Verify it's valid JSON by unmarshaling
	var unmarshaled ProjectState
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Fatalf("ToJSON() returned invalid JSON: %v", err)
	}

	if unmarshaled.Phase != state.Phase {
		t.Errorf("ToJSON() Phase = %v, want %v", unmarshaled.Phase, state.Phase)
	}
}

func TestProjectStateFromJSON(t *testing.T) {
	jsonData := `{
		"phase": "tasks",
		"active_task": 1,
		"completed": [1, 2],
		"locked": true
	}`

	state, err := ProjectStateFromJSON([]byte(jsonData))
	if err != nil {
		t.Fatalf("ProjectStateFromJSON() error = %v", err)
	}

	if state.Phase != PhaseTasks {
		t.Errorf("ProjectStateFromJSON() Phase = %v, want %v", state.Phase, PhaseTasks)
	}

	if state.ActiveTask == nil || *state.ActiveTask != 1 {
		t.Errorf("ProjectStateFromJSON() ActiveTask = %v, want 1", state.ActiveTask)
	}

	if len(state.Completed) != 2 {
		t.Errorf("ProjectStateFromJSON() Completed length = %v, want 2", len(state.Completed))
	}

	if !state.Locked {
		t.Error("ProjectStateFromJSON() Locked = false, want true")
	}
}

func TestProjectStateFromJSON_Invalid(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name:    "invalid JSON",
			json:    "{ invalid json }",
			wantErr: true,
		},
		{
			name:    "invalid phase",
			json:    `{"phase": "invalid"}`,
			wantErr: true,
		},
		{
			name:    "missing phase",
			json:    `{}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProjectStateFromJSON([]byte(tt.json))
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectStateFromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidIDE(t *testing.T) {
	tests := []struct {
		ide      string
		expected bool
	}{
		{"Cursor", true},
		{"cursor", true}, // case insensitive
		{"CURSOR", true},
		{"Claude Code", true},
		{"claude code", true},
		{"Antigravity", true},
		{"Windsurf", true},
		{"Cline", true},
		{"OpenCode", true},
		{"InvalidIDE", false},
		{"", false},
		{"  Cursor  ", true}, // trimmed
	}

	for _, tt := range tests {
		t.Run(tt.ide, func(t *testing.T) {
			if got := IsValidIDE(tt.ide); got != tt.expected {
				t.Errorf("IsValidIDE(%q) = %v, want %v", tt.ide, got, tt.expected)
			}
		})
	}
}

func TestGetSupportedIDEs(t *testing.T) {
	ides := GetSupportedIDEs()

	expected := []string{"Cursor", "Claude Code", "Antigravity", "Windsurf", "Cline", "OpenCode"}

	if len(ides) != len(expected) {
		t.Errorf("GetSupportedIDEs() length = %v, want %v", len(ides), len(expected))
	}

	for i, ide := range expected {
		if i < len(ides) && ides[i] != ide {
			t.Errorf("GetSupportedIDEs()[%d] = %v, want %v", i, ides[i], ide)
		}
	}
}

// Helper functions

func intPtr(i int) *int {
	return &i
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
