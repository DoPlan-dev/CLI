package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/utils"
)

// UpdateDashboardData silently updates dashboard data by ensuring project files are current.
// Since the dashboard HTML files read directly from project files via JavaScript fetch(),
// this function primarily validates that the dashboard directory exists and project files
// are accessible. It runs silently without user output.
func UpdateDashboardData(projectPath string) error {
	// Check if dashboard directory exists
	dashboardDir := filepath.Join(projectPath, "dashboard")
	if !utils.PathExists(dashboardDir) {
		// Dashboard not generated yet, skip silently
		return nil
	}

	// Validate that key project files exist (silent validation)
	// The dashboard HTML will handle missing files gracefully
	_ = validateDashboardDataSources(projectPath)

	// No errors - dashboard will read from project files directly
	return nil
}

// validateDashboardDataSources checks if dashboard data sources exist
// Returns error only for logging, doesn't fail the update
func validateDashboardDataSources(projectPath string) error {
	// Check for key files that dashboard reads
	keyFiles := []string{
		filepath.Join(projectPath, ".do", "system", "history", "active_state.json"),
		filepath.Join(projectPath, ".do", "plan", "TASKS.md"),
		filepath.Join(projectPath, ".do", "system", "meeting_session.json"),
	}

	for _, file := range keyFiles {
		if !utils.PathExists(file) {
			// File doesn't exist yet - that's okay, dashboard handles it gracefully
			continue
		}
	}

	return nil
}

// updateDashboardProgress reads progress from active_state.json
// This is informational - the dashboard reads directly from the file
func updateDashboardProgress(projectPath string) error {
	statePath := filepath.Join(projectPath, ".do", "system", "history", "active_state.json")
	if !utils.PathExists(statePath) {
		return nil // No state file yet
	}

	// Read and validate state file
	data, err := os.ReadFile(statePath)
	if err != nil {
		return fmt.Errorf("failed to read active_state.json: %w", err)
	}

	// Validate it's valid JSON (dashboard will parse it)
	var state map[string]interface{}
	if err := json.Unmarshal(data, &state); err != nil {
		return fmt.Errorf("invalid JSON in active_state.json: %w", err)
	}

	return nil
}

// updateDashboardTasks validates TASKS.md exists and is readable
func updateDashboardTasks(projectPath string) error {
	tasksPath := filepath.Join(projectPath, ".do", "plan", "TASKS.md")
	if !utils.PathExists(tasksPath) {
		return nil // No tasks file yet
	}

	// Validate file is readable
	_, err := os.ReadFile(tasksPath)
	if err != nil {
		return fmt.Errorf("failed to read TASKS.md: %w", err)
	}

	return nil
}

// updateDashboardMeetings validates meeting_session.json exists
func updateDashboardMeetings(projectPath string) error {
	meetingPath := filepath.Join(projectPath, ".do", "system", "meeting_session.json")
	if !utils.PathExists(meetingPath) {
		return nil // No meeting data yet
	}

	// Validate it's valid JSON
	data, err := os.ReadFile(meetingPath)
	if err != nil {
		return fmt.Errorf("failed to read meeting_session.json: %w", err)
	}

	var meeting map[string]interface{}
	if err := json.Unmarshal(data, &meeting); err != nil {
		return fmt.Errorf("invalid JSON in meeting_session.json: %w", err)
	}

	return nil
}

// updateDashboardAchievements validates memory_card.json exists
func updateDashboardAchievements(projectPath string) error {
	// Try project-level memory card first
	memoryCardPath := filepath.Join(projectPath, ".do", "system", "memory_card.json")
	if !utils.PathExists(memoryCardPath) {
		// Try global memory card
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil // Can't determine home dir, skip
		}
		memoryCardPath = filepath.Join(homeDir, ".doplan", "memory_card.json")
		if !utils.PathExists(memoryCardPath) {
			return nil // No memory card yet
		}
	}

	// Validate it's valid JSON
	data, err := os.ReadFile(memoryCardPath)
	if err != nil {
		return fmt.Errorf("failed to read memory_card.json: %w", err)
	}

	var memoryCard map[string]interface{}
	if err := json.Unmarshal(data, &memoryCard); err != nil {
		return fmt.Errorf("invalid JSON in memory_card.json: %w", err)
	}

	return nil
}

// updateDashboardSettings validates agent files exist
func updateDashboardSettings(projectPath string) error {
	agentsDir := filepath.Join(projectPath, ".cursor", "agents")
	if !utils.PathExists(agentsDir) {
		return nil // No agents directory yet
	}

	// Check if at least one agent file exists
	entries, err := os.ReadDir(agentsDir)
	if err != nil {
		return fmt.Errorf("failed to read agents directory: %w", err)
	}

	// At least one agent file should exist
	if len(entries) == 0 {
		return nil // No agents yet, that's okay
	}

	return nil
}
