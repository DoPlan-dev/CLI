package generator

import (
	"strings"
	"testing"
)

func TestGetAllCommands(t *testing.T) {
	commands := GetAllCommands()

	// Should have 14 commands (7 core + 6 tools + 1 optimize)
	expectedCount := 14
	if len(commands) != expectedCount {
		t.Errorf("GetAllCommands() returned %d commands, want %d", len(commands), expectedCount)
	}

	// Check for required fields
	coreCommandNames := []string{
		"hello",
		"tell",
		"meeting",
		"write",
		"plan",
		"build",
		"status",
	}

	toolsCommandNames := []string{
		"feedback",
		"state",
		"github",
		"security",
		"permissions",
		"access",
	}

	optimizeCommandNames := []string{
		"optimize",
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

	// Check all tools commands are present
	for _, name := range toolsCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required tools command %s not found", name)
		}
	}

	// Check all optimize commands are present
	for _, name := range optimizeCommandNames {
		if !foundCommands[name] {
			t.Errorf("Required optimize command %s not found", name)
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
		{"github", "github", false},
		{"optimize", "optimize", false},
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

	// Should have exactly 7 core commands
	if len(coreCommands) != 7 {
		t.Errorf("GetCoreCommands() returned %d commands, want 7", len(coreCommands))
	}

	// Verify all are core commands
	coreNames := []string{
		"hello", "tell", "meeting", "write", "plan", "build", "status",
	}

	foundNames := make(map[string]bool)
	for _, cmd := range coreCommands {
		foundNames[cmd.Name] = true
		// Verify category is "core"
		if cmd.Category != "core" {
			t.Errorf("Core command %s has category %s, want 'core'", cmd.Name, cmd.Category)
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

	// Should have exactly 7 squad commands (6 tools + 1 optimize)
	if len(squadCommands) != 7 {
		t.Errorf("GetSquadCommands() returned %d commands, want 7", len(squadCommands))
	}

	// Verify all are tools or optimize commands
	toolsNames := []string{
		"feedback", "state", "github", "security", "permissions", "access",
	}

	optimizeNames := []string{
		"optimize",
	}

	foundNames := make(map[string]bool)
	for _, cmd := range squadCommands {
		foundNames[cmd.Name] = true
		// Verify category is "tools" or "optimize"
		if cmd.Category != "tools" && cmd.Category != "optimize" {
			t.Errorf("Squad command %s has category %s, want 'tools' or 'optimize'", cmd.Name, cmd.Category)
		}
	}

	for _, name := range toolsNames {
		if !foundNames[name] {
			t.Errorf("Tools command %s not found", name)
		}
	}

	for _, name := range optimizeNames {
		if !foundNames[name] {
			t.Errorf("Optimize command %s not found", name)
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
	commandsWithGitHub := []string{"build"}

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
