package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ProjectRequest represents a request to create a new project
type ProjectRequest struct {
	ProjectName string `json:"project_name"`
	IDE         string `json:"ide"`
	ProjectType string `json:"project_type"`
}

// Validate validates the ProjectRequest and returns an error if invalid
func (r *ProjectRequest) Validate() error {
	if r.ProjectName == "" {
		return errors.New("project name is required")
	}

	// Validate project name: alphanumeric + hyphens/underscores only
	if !IsValidProjectName(r.ProjectName) {
		return fmt.Errorf("project name '%s' is invalid: must contain only alphanumeric characters, hyphens, and underscores", r.ProjectName)
	}

	if r.IDE == "" {
		return errors.New("IDE is required")
	}

	// Set default project type if not specified
	if r.ProjectType == "" {
		r.ProjectType = "Fullstack"
	}

	return nil
}

// IsValidProjectName checks if a project name is valid
// Valid names contain only alphanumeric characters, hyphens, and underscores
func IsValidProjectName(name string) bool {
	if len(name) == 0 {
		return false
	}

	for _, char := range name {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' ||
			char == '_') {
			return false
		}
	}

	return true
}

// ProjectState represents the current state of a project
type ProjectState struct {
	Phase      string `json:"phase"`
	ActiveTask *int   `json:"active_task,omitempty"`
	Completed  []int  `json:"completed"`
	Locked     bool   `json:"locked"`
}

// Valid phases for project state
const (
	PhaseIdea        = "idea"
	PhaseBrainstorm  = "brainstorm"
	PhaseWriting     = "writing"
	PhaseApproved    = "approved"
	PhaseTasks       = "tasks"
	PhaseBuilding    = "building"
)

// Validate validates the ProjectState and returns an error if invalid
func (s *ProjectState) Validate() error {
	validPhases := map[string]bool{
		PhaseIdea:       true,
		PhaseBrainstorm: true,
		PhaseWriting:    true,
		PhaseApproved:   true,
		PhaseTasks:      true,
		PhaseBuilding:   true,
	}

	if s.Phase == "" {
		return errors.New("phase is required")
	}

	if !validPhases[s.Phase] {
		return fmt.Errorf("invalid phase: %s (must be one of: idea, brainstorm, writing, approved, tasks, building)", s.Phase)
	}

	// Validate completed tasks are not negative
	for _, taskID := range s.Completed {
		if taskID < 0 {
			return fmt.Errorf("completed task ID cannot be negative: %d", taskID)
		}
	}

	// Validate active task is not negative if set
	if s.ActiveTask != nil && *s.ActiveTask < 0 {
		return fmt.Errorf("active task ID cannot be negative: %d", *s.ActiveTask)
	}

	return nil
}

// NewProjectState creates a new ProjectState with default values
func NewProjectState(phase string) *ProjectState {
	return &ProjectState{
		Phase:      phase,
		ActiveTask: nil,
		Completed: []int{},
		Locked:     false,
	}
}

// IsTaskCompleted checks if a task ID is in the completed list
func (s *ProjectState) IsTaskCompleted(taskID int) bool {
	for _, id := range s.Completed {
		if id == taskID {
			return true
		}
	}
	return false
}

// MarkTaskCompleted adds a task ID to the completed list if not already present
func (s *ProjectState) MarkTaskCompleted(taskID int) {
	if !s.IsTaskCompleted(taskID) {
		s.Completed = append(s.Completed, taskID)
	}
}

// SetActiveTask sets the active task ID
func (s *ProjectState) SetActiveTask(taskID *int) {
	s.ActiveTask = taskID
}

// ClearActiveTask clears the active task
func (s *ProjectState) ClearActiveTask() {
	s.ActiveTask = nil
}

// ToJSON converts ProjectState to JSON bytes
func (s *ProjectState) ToJSON() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

// FromJSON creates a ProjectState from JSON bytes
func ProjectStateFromJSON(data []byte) (*ProjectState, error) {
	var state ProjectState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("failed to unmarshal project state: %w", err)
	}

	if err := state.Validate(); err != nil {
		return nil, fmt.Errorf("invalid project state: %w", err)
	}

	return &state, nil
}

// GetSupportedIDEs returns a list of supported IDE names
func GetSupportedIDEs() []string {
	return []string{
		"Cursor",
		"Claude Code",
		"Antigravity",
		"Windsurf",
		"Cline",
		"OpenCode",
	}
}

// IsValidIDE checks if the given IDE name is supported
func IsValidIDE(ide string) bool {
	ide = strings.TrimSpace(ide)
	supported := GetSupportedIDEs()
	for _, supportedIDE := range supported {
		if strings.EqualFold(ide, supportedIDE) {
			return true
		}
	}
	return false
}

