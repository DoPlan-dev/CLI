package generator

import (
	"testing"
)

func TestGetAllAgents(t *testing.T) {
	agents := GetAllAgents()

	// Should have exactly 18 agents
	if len(agents) != 18 {
		t.Errorf("GetAllAgents() returned %d agents, want 18", len(agents))
	}

	// Check for required fields
	requiredAgents := []string{
		"Project Orchestrator",
		"Product Manager",
		"Engineering Lead",
		"System Architect",
		"Frontend Lead",
		"Backend Lead",
		"DevOps Engineer",
		"Security Lead",
		"Performance Engineer",
		"Design & UX Manager",
		"UI/UX Designer",
		"QA & Reliability Manager",
		"QA Engineer",
		"Release & Growth Manager",
		"Release Captain",
		"Growth Coach",
		"Documentation Lead",
		"Documentation Writer",
	}

	foundAgents := make(map[string]bool)
	for _, agent := range agents {
		foundAgents[agent.Name] = true

		// Validate required fields
		if agent.Name == "" {
			t.Error("Agent name cannot be empty")
		}
		if agent.Role == "" {
			t.Errorf("Agent %s: Role cannot be empty", agent.Name)
		}
		if agent.SystemPrompt == "" {
			t.Errorf("Agent %s: SystemPrompt cannot be empty", agent.Name)
		}
		if agent.FileName == "" {
			t.Errorf("Agent %s: FileName cannot be empty", agent.Name)
		}
	}

	// Check all required agents are present
	for _, name := range requiredAgents {
		if !foundAgents[name] {
			t.Errorf("Required agent %s not found", name)
		}
	}
}

func TestGetAgentByName(t *testing.T) {
	testCases := []struct {
		name     string
		wantName string
		wantNil  bool
	}{
		{"Project Orchestrator", "Project Orchestrator", false},
		{"Product Manager", "Product Manager", false},
		{"Engineering Lead", "Engineering Lead", false},
		{"Nonexistent Agent", "", true},
		{"", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			agent := GetAgentByName(tc.name)
			if tc.wantNil {
				if agent != nil {
					t.Errorf("GetAgentByName(%q) = %v, want nil", tc.name, agent)
				}
			} else {
				if agent == nil {
					t.Errorf("GetAgentByName(%q) = nil, want agent", tc.name)
				} else if agent.Name != tc.wantName {
					t.Errorf("GetAgentByName(%q).Name = %q, want %q", tc.name, agent.Name, tc.wantName)
				}
			}
		})
	}
}

func TestGetAgentsByManager(t *testing.T) {
	// Test Project Orchestrator's direct reports
	reports := GetAgentsByManager("Project Orchestrator")
	expectedReports := []string{
		"Product Manager",
		"Engineering Lead",
		"Design & UX Manager",
		"QA & Reliability Manager",
		"Release & Growth Manager",
		"Documentation Lead",
	}

	if len(reports) != len(expectedReports) {
		t.Errorf("GetAgentsByManager(Project Orchestrator) returned %d agents, want %d", len(reports), len(expectedReports))
	}

	foundReports := make(map[string]bool)
	for _, agent := range reports {
		foundReports[agent.Name] = true
	}

	for _, expected := range expectedReports {
		if !foundReports[expected] {
			t.Errorf("Expected report %s not found for Project Orchestrator", expected)
		}
	}

	// Test Engineering Lead's direct reports
	engReports := GetAgentsByManager("Engineering Lead")
	expectedEngReports := []string{
		"System Architect",
		"Frontend Lead",
		"Backend Lead",
		"DevOps Engineer",
		"Security Lead",
		"Performance Engineer",
	}

	if len(engReports) != len(expectedEngReports) {
		t.Errorf("GetAgentsByManager(Engineering Lead) returned %d agents, want %d", len(engReports), len(expectedEngReports))
	}

	// Test agent with no reports
	noReports := GetAgentsByManager("Product Manager")
	if len(noReports) != 0 {
		t.Errorf("GetAgentsByManager(Product Manager) returned %d agents, want 0", len(noReports))
	}
}

func TestAgentHierarchy(t *testing.T) {
	agents := GetAllAgents()

	// Project Orchestrator should not report to anyone
	orchestrator := GetAgentByName("Project Orchestrator")
	if orchestrator == nil {
		t.Fatal("Project Orchestrator not found")
	}
	if orchestrator.ReportsTo != "" {
		t.Errorf("Project Orchestrator reports to %q, want empty", orchestrator.ReportsTo)
	}

	// All other agents should report to someone
	for _, agent := range agents {
		if agent.Name == "Project Orchestrator" {
			continue
		}
		if agent.ReportsTo == "" {
			t.Errorf("Agent %s has no ReportsTo", agent.Name)
		}
		// Verify the manager exists
		manager := GetAgentByName(agent.ReportsTo)
		if manager == nil {
			t.Errorf("Agent %s reports to %q, but that agent doesn't exist", agent.Name, agent.ReportsTo)
		}
	}

	// Verify Manages relationships are bidirectional
	for _, agent := range agents {
		for _, managedName := range agent.Manages {
			managed := GetAgentByName(managedName)
			if managed == nil {
				t.Errorf("Agent %s manages %q, but that agent doesn't exist", agent.Name, managedName)
			} else if managed.ReportsTo != agent.Name {
				t.Errorf("Agent %s manages %q, but %q reports to %q instead", agent.Name, managedName, managedName, managed.ReportsTo)
			}
		}
	}
}

func TestAgentFileNames(t *testing.T) {
	agents := GetAllAgents()

	// Check all file names are unique
	fileNames := make(map[string]string)
	for _, agent := range agents {
		if existing, exists := fileNames[agent.FileName]; exists {
			t.Errorf("Duplicate filename %q: used by both %s and %s", agent.FileName, existing, agent.Name)
		}
		fileNames[agent.FileName] = agent.Name

		// File names should end with .md
		if len(agent.FileName) < 3 || agent.FileName[len(agent.FileName)-3:] != ".md" {
			t.Errorf("Agent %s: FileName %q should end with .md", agent.Name, agent.FileName)
		}
	}
}

func TestAgentResponsibilities(t *testing.T) {
	agents := GetAllAgents()

	for _, agent := range agents {
		if len(agent.Responsibilities) == 0 {
			t.Errorf("Agent %s has no responsibilities", agent.Name)
		}

		// Each responsibility should be non-empty
		for i, resp := range agent.Responsibilities {
			if resp == "" {
				t.Errorf("Agent %s: Responsibility %d is empty", agent.Name, i)
			}
		}
	}
}
