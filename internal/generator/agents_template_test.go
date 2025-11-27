package generator

import (
	"strings"
	"testing"
)

func TestRenderAgentMarkdown(t *testing.T) {
	agent := GetAgentByName("Project Orchestrator")
	if agent == nil {
		t.Fatal("Project Orchestrator agent not found")
	}

	markdown, err := RenderAgentMarkdown(agent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdown() error = %v", err)
	}

	// Check required sections
	requiredSections := []string{
		"# Project Orchestrator",
		"## Role",
		"## System Prompt",
		"## Responsibilities",
		"## Reports To",
		"## Manages",
	}

	for _, section := range requiredSections {
		if !strings.Contains(markdown, section) {
			t.Errorf("Rendered markdown missing section: %s", section)
		}
	}

	// Check role is included
	if !strings.Contains(markdown, agent.Role) {
		t.Errorf("Rendered markdown missing role: %s", agent.Role)
	}

	// Check system prompt is included
	if !strings.Contains(markdown, "Project Orchestrator") {
		t.Errorf("Rendered markdown missing system prompt content")
	}

	// Check responsibilities are included
	for _, resp := range agent.Responsibilities {
		if !strings.Contains(markdown, resp) {
			t.Errorf("Rendered markdown missing responsibility: %s", resp)
		}
	}
}

func TestRenderAgentMarkdown_NilAgent(t *testing.T) {
	_, err := RenderAgentMarkdown(nil)
	if err == nil {
		t.Error("RenderAgentMarkdown(nil) should return error")
	}
}

func TestRenderAgentMarkdown_AllAgents(t *testing.T) {
	agents := GetAllAgents()

	for _, agent := range agents {
		t.Run(agent.Name, func(t *testing.T) {
			markdown, err := RenderAgentMarkdown(&agent)
			if err != nil {
				t.Fatalf("RenderAgentMarkdown() error = %v", err)
			}

			// Check title
			expectedTitle := "# " + agent.Name
			if !strings.Contains(markdown, expectedTitle) {
				t.Errorf("Rendered markdown missing title: %s", expectedTitle)
			}

			// Check role
			if !strings.Contains(markdown, agent.Role) {
				t.Errorf("Rendered markdown missing role: %s", agent.Role)
			}

			// Check system prompt
			if !strings.Contains(markdown, agent.SystemPrompt) {
				t.Errorf("Rendered markdown missing system prompt")
			}

			// Check Reports To section
			if agent.ReportsTo == "" {
				if !strings.Contains(markdown, "None (Top Level)") {
					t.Error("Top-level agent should show 'None (Top Level)' in Reports To")
				}
			} else {
				if !strings.Contains(markdown, agent.ReportsTo) {
					t.Errorf("Rendered markdown missing ReportsTo: %s", agent.ReportsTo)
				}
			}

			// Check Manages section
			if len(agent.Manages) == 0 {
				if !strings.Contains(markdown, "## Manages\nNone") {
					t.Error("Agent with no subordinates should show 'None' in Manages")
				}
			} else {
				for _, managed := range agent.Manages {
					if !strings.Contains(markdown, managed) {
						t.Errorf("Rendered markdown missing managed agent: %s", managed)
					}
				}
			}
		})
	}
}

func TestRenderAgentMarkdown_Formatting(t *testing.T) {
	agent := GetAgentByName("Product Manager")
	if agent == nil {
		t.Fatal("Product Manager agent not found")
	}

	markdown, err := RenderAgentMarkdown(agent)
	if err != nil {
		t.Fatalf("RenderAgentMarkdown() error = %v", err)
	}

	// Check responsibilities are formatted as list items
	lines := strings.Split(markdown, "\n")
	inResponsibilities := false
	foundListItems := false

	for _, line := range lines {
		if strings.Contains(line, "## Responsibilities") {
			inResponsibilities = true
			continue
		}
		if inResponsibilities {
			if strings.HasPrefix(line, "- ") {
				foundListItems = true
			}
			if strings.HasPrefix(line, "## ") {
				// Next section, stop checking
				break
			}
		}
	}

	if !foundListItems {
		t.Error("Responsibilities should be formatted as list items (starting with '- ')")
	}
}
