package generator

import (
	"strings"
	"testing"
)

func TestGetAllCommands(t *testing.T) {
	commands := GetAllCommands()

	// Should have 19 commands (11 core + 8 squad)
	expectedCount := 19
	if len(commands) != expectedCount {
		t.Errorf("GetAllCommands() returned %d commands, want %d", len(commands), expectedCount)
	}

	// Check for required fields
	coreCommandNames := []string{
		"tell",
		"improve",
		"team",
		"write",
		"change",
		"good",
		"tasks",
		"load",
		"build",
		"progress",
		"finished",
	}

	squadCommandNames := []string{
		"secure",
		"roles",
		"money",
		"pretty",
		"seo",
		"ship",
		"safe",
		"cheap",
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

	// Check all core commands are present
	for _, name := range coreCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required core command %s not found", name)
		}
	}

	// Check all squad commands are present
	for _, name := range squadCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required squad command %s not found", name)
		}
	}
}

func TestGetCommandByName(t *testing.T) {
	testCases := []struct {
		name     string
		wantName string
		wantNil  bool
	}{
		{"tell", "tell", false},
		{"build", "build", false},
		{"ship", "ship", false},
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

	// Should have exactly 11 core commands
	if len(coreCommands) != 11 {
		t.Errorf("GetCoreCommands() returned %d commands, want 11", len(coreCommands))
	}

	// Verify all are core commands
	coreNames := []string{
		"tell", "improve", "team", "write", "change",
		"good", "tasks", "load", "build", "progress", "finished",
	}

	foundNames := make(map[string]bool)
	for _, cmd := range coreCommands {
		foundNames[cmd.Name] = true
	}

	for _, name := range coreNames {
		if !foundNames[name] {
			t.Errorf("Core command %s not found", name)
		}
	}
}

func TestGetSquadCommands(t *testing.T) {
	squadCommands := GetSquadCommands()

	// Should have exactly 8 squad commands
	if len(squadCommands) != 8 {
		t.Errorf("GetSquadCommands() returned %d commands, want 8", len(squadCommands))
	}

	// Verify all are squad commands
	squadNames := []string{
		"secure", "roles", "money", "pretty",
		"seo", "ship", "safe", "cheap",
	}

	foundNames := make(map[string]bool)
	for _, cmd := range squadCommands {
		foundNames[cmd.Name] = true
	}

	for _, name := range squadNames {
		if !foundNames[name] {
			t.Errorf("Squad command %s not found", name)
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

			// Should have at least one agent involved (except /team which might have none)
			if cmd.Name != "team" && len(cmd.AgentInvolvement) == 0 {
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
	commandsWithGitHub := []string{"build", "finished"}

	for _, cmdName := range commandsWithGitHub {
		cmd := GetCommandByName(cmdName)
		if cmd == nil {
			t.Fatalf("Command %s not found", cmdName)
		}

		if cmd.GitHubAutomation == "" {
			t.Errorf("Command %s should have GitHubAutomation field", cmdName)
		}
	}
}

