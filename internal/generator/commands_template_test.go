package generator

import (
	"strings"
	"testing"
)

func TestRenderCommandMarkdown(t *testing.T) {
	cmd := &Command{
		Name:        "tell",
		Trigger:     "Exact match: /tell or /tell <idea>",
		Description: "Capture project idea",
		Action: `When user types /tell or /tell <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .plan/00_System/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /improve to brainstorm with the team."`,
		AgentInvolvement: []string{
			"Project Orchestrator",
			"Product Manager",
		},
		FilesRead: []string{},
		FilesModified: []string{
			".plan/00_System/IDEA.md",
			".plan/active_state.json",
		},
		Examples: []string{
			"/tell",
			"/tell Build a todo app",
		},
	}

	markdown, err := RenderCommandMarkdown(cmd)
	if err != nil {
		t.Fatalf("RenderCommandMarkdown() error = %v", err)
	}

	// Check required sections
	if !strings.Contains(markdown, "# /tell") {
		t.Error("Markdown should contain command title")
	}
	if !strings.Contains(markdown, "## Trigger") {
		t.Error("Markdown should contain Trigger section")
	}
	if !strings.Contains(markdown, "## Action") {
		t.Error("Markdown should contain Action section")
	}
	if !strings.Contains(markdown, "## Agent Involvement") {
		t.Error("Markdown should contain Agent Involvement section")
	}
	if !strings.Contains(markdown, "Project Orchestrator") {
		t.Error("Markdown should contain agent names")
	}
	if !strings.Contains(markdown, "## Examples") {
		t.Error("Markdown should contain Examples section when examples exist")
	}
	if !strings.Contains(markdown, "/tell") {
		t.Error("Markdown should contain example usage")
	}
	if !strings.Contains(markdown, "## Files Modified") {
		t.Error("Markdown should contain Files Modified section")
	}
	if !strings.Contains(markdown, ".plan/00_System/IDEA.md") {
		t.Error("Markdown should contain file paths")
	}
}

func TestRenderCommandMarkdown_NilCommand(t *testing.T) {
	_, err := RenderCommandMarkdown(nil)
	if err == nil {
		t.Errorf("expected error when passing nil to RenderCommandMarkdown, got nil")
	}
}

func TestRenderCommandMarkdown_AllCommands(t *testing.T) {
	commands := GetAllCommands()

	for _, cmd := range commands {
		t.Run(cmd.Name, func(t *testing.T) {
			markdown, err := RenderCommandMarkdown(&cmd)
			if err != nil {
				t.Fatalf("RenderCommandMarkdown() error = %v", err)
			}

			// Check required sections
			if !strings.Contains(markdown, "# /"+cmd.Name) {
				t.Errorf("Markdown should contain command title: # /%s", cmd.Name)
			}
			if !strings.Contains(markdown, "## Trigger") {
				t.Error("Markdown should contain Trigger section")
			}
			if !strings.Contains(markdown, "## Action") {
				t.Error("Markdown should contain Action section")
			}
			if !strings.Contains(markdown, "## Agent Involvement") {
				t.Error("Markdown should contain Agent Involvement section")
			}

			// Check optional sections
			if len(cmd.Examples) > 0 {
				if !strings.Contains(markdown, "## Examples") {
					t.Error("Markdown should contain Examples section when examples exist")
				}
			}

			if len(cmd.FilesRead) > 0 {
				if !strings.Contains(markdown, "## Files Read") {
					t.Error("Markdown should contain Files Read section when files are read")
				}
			}

			if len(cmd.FilesModified) > 0 {
				if !strings.Contains(markdown, "## Files Modified") {
					t.Error("Markdown should contain Files Modified section when files are modified")
				}
			}

			if cmd.GitHubAutomation != "" {
				if !strings.Contains(markdown, "## GitHub Automation") {
					t.Error("Markdown should contain GitHub Automation section when automation exists")
				}
			}

			// Check agent names are included
			for _, agent := range cmd.AgentInvolvement {
				if !strings.Contains(markdown, agent) {
					t.Errorf("Markdown should contain agent name: %s", agent)
				}
			}
		})
	}
}

func TestRenderCommandMarkdown_OptionalSections(t *testing.T) {
	// Test command without examples
	cmdNoExamples := &Command{
		Name:             "test",
		Trigger:          "Exact match: /test",
		Description:      "Test command",
		Action:           "Test action",
		AgentInvolvement: []string{"Test Agent"},
		Examples:         []string{},
	}

	{
		markdown, err := RenderCommandMarkdown(cmdNoExamples)
		if err != nil {
			t.Fatalf("RenderCommandMarkdown() error = %v", err)
		}

		if strings.Contains(markdown, "## Examples") {
			t.Error("Markdown should not contain Examples section when no examples exist")
		}
	}

	// Test command without files read
	cmdNoFilesRead := &Command{
		Name:             "test2",
		Trigger:          "Exact match: /test2",
		Description:      "Test command 2",
		Action:           "Test action",
		AgentInvolvement: []string{"Test Agent"},
		FilesRead:        []string{},
	}

	{
		markdown, err := RenderCommandMarkdown(cmdNoFilesRead)
		if err != nil {
			t.Fatalf("RenderCommandMarkdown() error = %v", err)
		}

		if strings.Contains(markdown, "## Files Read") {
			t.Error("Markdown should not contain Files Read section when no files are read")
		}
	}

	// Test command without GitHub automation
	cmdNoGitHub := &Command{
		Name:             "test3",
		Trigger:          "Exact match: /test3",
		Description:      "Test command 3",
		Action:           "Test action",
		AgentInvolvement: []string{"Test Agent"},
		GitHubAutomation: "",
	}

	{
		markdown, err := RenderCommandMarkdown(cmdNoGitHub)
		if err != nil {
			t.Fatalf("RenderCommandMarkdown() error = %v", err)
		}

		if strings.Contains(markdown, "## GitHub Automation") {
			t.Error("Markdown should not contain GitHub Automation section when automation is empty")
		}
	}
}

func TestRenderCommandMarkdown_Formatting(t *testing.T) {
	cmd := &Command{
		Name:    "build",
		Trigger: "Exact match: /build or /build <task_id>",
		Action:  "Test action",
		AgentInvolvement: []string{
			"Engineering Lead",
			"Project Orchestrator",
		},
		Examples: []string{
			"/build",
			"/build 1.2",
		},
		FilesRead: []string{
			".plan/TASKS.md",
			".plan/active_state.json",
		},
		FilesModified: []string{
			".plan/active_state.json",
		},
		GitHubAutomation: "Test automation details",
	}

	markdown, err := RenderCommandMarkdown(cmd)
	if err != nil {
		t.Fatalf("RenderCommandMarkdown() error = %v", err)
	}

	// Check formatting
	lines := strings.Split(markdown, "\n")

	// Should start with title
	if !strings.HasPrefix(lines[0], "# /") {
		t.Error("Markdown should start with command title")
	}

	// Check section headers
	hasTrigger := false
	hasExamples := false
	hasAction := false
	hasAgentInvolvement := false
	hasFilesRead := false
	hasFilesModified := false
	hasGitHubAutomation := false

	for _, line := range lines {
		if strings.HasPrefix(line, "## Trigger") {
			hasTrigger = true
		}
		if strings.HasPrefix(line, "## Examples") {
			hasExamples = true
		}
		if strings.HasPrefix(line, "## Action") {
			hasAction = true
		}
		if strings.HasPrefix(line, "## Agent Involvement") {
			hasAgentInvolvement = true
		}
		if strings.HasPrefix(line, "## Files Read") {
			hasFilesRead = true
		}
		if strings.HasPrefix(line, "## Files Modified") {
			hasFilesModified = true
		}
		if strings.HasPrefix(line, "## GitHub Automation") {
			hasGitHubAutomation = true
		}
	}

	if !hasTrigger {
		t.Error("Markdown should contain Trigger section")
	}
	if !hasExamples {
		t.Error("Markdown should contain Examples section")
	}
	if !hasAction {
		t.Error("Markdown should contain Action section")
	}
	if !hasAgentInvolvement {
		t.Error("Markdown should contain Agent Involvement section")
	}
	if !hasFilesRead {
		t.Error("Markdown should contain Files Read section")
	}
	if !hasFilesModified {
		t.Error("Markdown should contain Files Modified section")
	}
	if !hasGitHubAutomation {
		t.Error("Markdown should contain GitHub Automation section")
	}
}
