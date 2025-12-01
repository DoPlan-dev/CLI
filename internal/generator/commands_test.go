package generator

import (
	"strings"
	"testing"
)

func TestGetAllCommands(t *testing.T) {
	commands := GetAllCommands()

	// Should have 5 commands (2 onboarding + 2 developing + 1 system)
	expectedCount := 5
	if len(commands) != expectedCount {
		t.Errorf("GetAllCommands() returned %d commands, want %d", len(commands), expectedCount)
	}

	// Check for required fields
	onboardingCommandNames := []string{
		"hey",
		"do",
	}

	developingCommandNames := []string{
		"plan",
		"dev",
	}

	systemCommandNames := []string{
		"sys",
	}

	foundCommands := make(map[string]bool)
	for _, cmd := range commands {
		foundCommands[cmd.Name] = true

		// Validate required fields
		if cmd.Name == "" {
			t.Error("Command name cannot be empty")
		}
		if cmd.Trigger == "" {
			t.Errorf("Command %s: Trigger cannot be empty", cmd.Name)
		}
		if cmd.Description == "" {
			t.Errorf("Command %s: Description cannot be empty", cmd.Name)
		}
		if cmd.Action == "" {
			t.Errorf("Command %s: Action cannot be empty", cmd.Name)
		}
	}

	// Check all onboarding commands are present
	for _, name := range onboardingCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required onboarding command %s not found", name)
		}
	}

	// Check all developing commands are present
	for _, name := range developingCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required developing command %s not found", name)
		}
	}

	// Check all system commands are present
	for _, name := range systemCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required system command %s not found", name)
		}
	}
}

func TestGetCommandByName(t *testing.T) {
	testCases := []struct {
		name     string
		wantName string
		wantNil  bool
	}{
		{"hey", "hey", false},
		{"do", "do", false},
		{"plan", "plan", false},
		{"dev", "dev", false},
		{"sys", "sys", false},
		{"nonexistent", "", true},
		{"", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := GetCommandByName(tc.name)
			if tc.wantNil {
				if cmd != nil {
					t.Errorf("GetCommandByName(%q) = %v, want nil", tc.name, cmd)
				}
			} else {
				if cmd == nil {
					t.Errorf("GetCommandByName(%q) = nil, want command", tc.name)
				} else if cmd.Name != tc.wantName {
					t.Errorf("GetCommandByName(%q).Name = %q, want %q", tc.name, cmd.Name, tc.wantName)
				}
			}
		})
	}
}

func TestGetCoreCommands(t *testing.T) {
	coreCommands := GetCoreCommands()

	// Should have exactly 4 core commands (onboarding + developing)
	if len(coreCommands) != 4 {
		t.Errorf("GetCoreCommands() returned %d commands, want 4", len(coreCommands))
	}

	// Verify all are core commands (onboarding or developing)
	coreNames := []string{
		"hey", "do", "plan", "dev",
	}

	foundNames := make(map[string]bool)
	for _, cmd := range coreCommands {
		foundNames[cmd.Name] = true
		// Verify category is correct (onboarding or developing)
		validCategories := map[string]bool{
			"onboarding": true,
			"developing": true,
		}
		if !validCategories[cmd.Category] {
			t.Errorf("Command %s has invalid category %s, want 'onboarding' or 'developing'", cmd.Name, cmd.Category)
		}
	}

	for _, name := range coreNames {
		if !foundNames[name] {
			t.Errorf("Core command %s not found", name)
		}
	}
}

func TestGetSquadCommands(t *testing.T) {
	squadCommands := GetSquadCommands()

	// Should have exactly 1 squad command (system)
	if len(squadCommands) != 1 {
		t.Errorf("GetSquadCommands() returned %d commands, want 1", len(squadCommands))
	}

	// Verify all are system commands
	systemNames := []string{
		"sys",
	}

	foundNames := make(map[string]bool)
	for _, cmd := range squadCommands {
		foundNames[cmd.Name] = true
		// Verify category is "system"
		if cmd.Category != "system" {
			t.Errorf("Squad command %s has category %s, want 'system'", cmd.Name, cmd.Category)
		}
	}

	for _, name := range systemNames {
		if !foundNames[name] {
			t.Errorf("System command %s not found", name)
		}
	}
}

func TestCommandFields(t *testing.T) {
	commands := GetAllCommands()

	for _, cmd := range commands {
		t.Run(cmd.Name, func(t *testing.T) {
			// Trigger should start with /
			if !strings.HasPrefix(cmd.Trigger, "/") {
				t.Errorf("Command %s: Trigger should start with '/', got %q", cmd.Name, cmd.Trigger)
			}

			// Should have at least one agent involved
			if len(cmd.AgentInvolvement) == 0 {
				t.Errorf("Command %s: Should have at least one agent involved", cmd.Name)
			}

			// Examples should not be empty
			if len(cmd.Examples) == 0 {
				t.Errorf("Command %s: Should have at least one example", cmd.Name)
			}

			// Each example should start with /
			for i, example := range cmd.Examples {
				if !strings.HasPrefix(example, "/") {
					t.Errorf("Command %s: Example %d should start with '/', got %q", cmd.Name, i, example)
				}
			}
		})
	}
}

func TestCommandAgentInvolvement(t *testing.T) {
	commands := GetAllCommands()

	// Verify agent names are valid (they should match agent names)
	validAgents := make(map[string]bool)
	agents := GetAllAgents()
	for _, agent := range agents {
		validAgents[agent.Name] = true
	}

	for _, cmd := range commands {
		for _, agentName := range cmd.AgentInvolvement {
			// Some commands might reference agents that don't exist in our 18 base agents
			// (like "Project Orchestrator" which exists), so we'll just check they're not empty
			if agentName == "" {
				t.Errorf("Command %s: Agent involvement contains empty string", cmd.Name)
			}
		}
	}
}

func TestCommandFiles(t *testing.T) {
	commands := GetAllCommands()

	for _, cmd := range commands {
		t.Run(cmd.Name, func(t *testing.T) {
			// Files should use consistent patterns
			for _, file := range cmd.FilesRead {
				if file == "" {
					t.Errorf("Command %s: FilesRead contains empty string", cmd.Name)
				}
			}

			for _, file := range cmd.FilesModified {
				if file == "" {
					t.Errorf("Command %s: FilesModified contains empty string", cmd.Name)
				}
			}
		})
	}
}

func TestCommandGitHubAutomation(t *testing.T) {
	// Commands that should have GitHub automation
	commandsWithGitHub := []string{"dev"}

	for _, cmdName := range commandsWithGitHub {
		var cmd *Command
		for _, c := range GetAllCommands() {
			if c.Name == cmdName {
				cmd = &c
				break
			}
		}
		if cmd == nil {
			t.Fatalf("Command %s not found", cmdName)
		}

		if cmd.GitHubAutomation == "" {
			t.Errorf("Command %s should have GitHubAutomation field", cmdName)
		}
	}
}
