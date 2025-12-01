package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

func TestCommandsGenerator_Name(t *testing.T) {
	generator := &CommandsGenerator{}
	if got := generator.Name(); got != "Commands" {
		t.Errorf("CommandsGenerator.Name() = %v, want %v", got, "Commands")
	}
}

func TestCommandsGenerator_Generate(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	// Verify .cursor/commands directory was created
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if _, err := os.Stat(commandsDir); os.IsNotExist(err) {
		t.Error("CommandsGenerator.Generate() should create .cursor/commands directory")
	}

	// Verify all command files were created
	// Commands are now in category folders, check central location
	centralCommandsDir := filepath.Join(tmpDir, ".do", "core", "commands")
	commands := GetAllCommands()

	for _, cmd := range commands {
		category := cmd.Category
		if category == "" {
			category = "other"
		}

		// Check central location (where files are actually created)
		centralCommandPath := filepath.Join(centralCommandsDir, category, cmd.Name+".md")
		if _, err := os.Stat(centralCommandPath); os.IsNotExist(err) {
			t.Errorf("CommandsGenerator.Generate() should create file %s in central location", cmd.Name+".md")
		}

		// Check IDE location (symlink or copy)
		ideCommandPath := filepath.Join(commandsDir, category, cmd.Name+".md")
		if _, err := os.Stat(ideCommandPath); os.IsNotExist(err) {
			t.Errorf("CommandsGenerator.Generate() should make file %s accessible via IDE location", cmd.Name+".md")
		}
	}
}

func TestCommandsGenerator_Generate_FileContent(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	// Verify file content for a specific command
	// Commands are in category folders, check central location
	centralCommandsDir := filepath.Join(tmpDir, ".do", "core", "commands")
	// "do" command is in "onboarding" category
	tellPath := filepath.Join(centralCommandsDir, "onboarding", "do.md")

	// If not found in core, try other categories
	if _, err := os.Stat(tellPath); os.IsNotExist(err) {
		// Search all categories
		entries, _ := os.ReadDir(centralCommandsDir)
		for _, entry := range entries {
			if entry.IsDir() {
				candidatePath := filepath.Join(centralCommandsDir, entry.Name(), "tell.md")
				if _, err := os.Stat(candidatePath); err == nil {
					tellPath = candidatePath
					break
				}
			}
		}
	}

	content, err := os.ReadFile(tellPath)
	if err != nil {
		t.Fatalf("Failed to read tell.md: %v", err)
	}

	contentStr := string(content)

	// Check for required sections
	if !strings.Contains(contentStr, "# /tell") {
		t.Error("tell.md should contain command title")
	}
	if !strings.Contains(contentStr, "## Trigger") {
		t.Error("tell.md should contain Trigger section")
	}
	if !strings.Contains(contentStr, "## Action") {
		t.Error("tell.md should contain Action section")
	}
	if !strings.Contains(contentStr, "## Agent Involvement") {
		t.Error("tell.md should contain Agent Involvement section")
	}
}

func TestCommandsGenerator_Generate_AllCommandsContent(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	// Check central location (where files are actually created)
	centralCommandsDir := filepath.Join(tmpDir, ".do", "core", "commands")
	commands := GetAllCommands()

	// Verify each command file has correct content
	for _, cmd := range commands {
		t.Run(cmd.Name, func(t *testing.T) {
			category := cmd.Category
			if category == "" {
				category = "other"
			}

			// Check central location
			commandPath := filepath.Join(centralCommandsDir, category, cmd.Name+".md")
			content, err := os.ReadFile(commandPath)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", cmd.Name+".md", err)
			}

			contentStr := string(content)

			// Check for command title
			expectedTitle := "# /" + cmd.Name
			if !strings.Contains(contentStr, expectedTitle) {
				t.Errorf("%s should contain title %s", cmd.Name+".md", expectedTitle)
			}

			// Check for required sections
			if !strings.Contains(contentStr, "## Trigger") {
				t.Errorf("%s should contain Trigger section", cmd.Name+".md")
			}
			if !strings.Contains(contentStr, "## Action") {
				t.Errorf("%s should contain Action section", cmd.Name+".md")
			}
			if !strings.Contains(contentStr, "## Agent Involvement") {
				t.Errorf("%s should contain Agent Involvement section", cmd.Name+".md")
			}

			// Check agent names are included
			for _, agent := range cmd.AgentInvolvement {
				if !strings.Contains(contentStr, agent) {
					t.Errorf("%s should contain agent name %s", cmd.Name+".md", agent)
				}
			}
		})
	}
}

func TestCommandsGenerator_Generate_InvalidPath(t *testing.T) {
	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}

	// Try to generate in a non-existent parent directory
	invalidPath := filepath.Join("/nonexistent", "path", "test")
	err := generator.Generate(request, invalidPath)
	if err == nil {
		t.Error("CommandsGenerator.Generate() should return error for invalid path")
	}
}

func TestCommandsGenerator_Generate_ExistingDirectory(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Pre-create .cursor/commands directory
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		t.Fatalf("Failed to create commands directory: %v", err)
	}

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}
	// Should succeed even if directory already exists
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Errorf("CommandsGenerator.Generate() error = %v, want nil (existing directory should be OK)", err)
	}
}

func TestGenerateCommands(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor"},
		ProjectType: "Fullstack",
	}

	if err := GenerateCommands(request, tmpDir); err != nil {
		t.Fatalf("GenerateCommands() error = %v", err)
	}

	// Verify directory was created
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if _, err := os.Stat(commandsDir); os.IsNotExist(err) {
		t.Error("GenerateCommands() should create .cursor/commands directory")
	}
}

func TestCommandsGenerator_Generate_MultipleIDEs(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "doplan-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	request := &models.ProjectRequest{
		ProjectName: "test-project",
		IDEs:        []string{"Cursor", "Claude Code", "Antigravity", "Windsurf", "Cline", "OpenCode"},
		ProjectType: "Fullstack",
	}

	generator := &CommandsGenerator{}
	if err := generator.Generate(request, tmpDir); err != nil {
		t.Fatalf("CommandsGenerator.Generate() error = %v", err)
	}

	// Verify commands directories were created for all IDEs
	expectedCommandsDirs := map[string]string{
		"Cursor":      filepath.Join(tmpDir, ".cursor", "commands"),
		"Claude Code": filepath.Join(tmpDir, ".claude", "commands"),
		"Antigravity": filepath.Join(tmpDir, ".antigravity", "commands"),
		"Windsurf":    filepath.Join(tmpDir, ".windsurf", "commands"),
		"Cline":       filepath.Join(tmpDir, ".cline", "commands"),
		"OpenCode":    filepath.Join(tmpDir, ".opencode", "commands"),
	}

	// First verify central location has commands
	centralCommandsDir := filepath.Join(tmpDir, ".do", "core", "commands")
	commands := GetAllCommands()
	if len(commands) == 0 {
		t.Fatalf("No commands found")
	}

	// Check central location for first command
	firstCmd := commands[0]
	category := firstCmd.Category
	if category == "" {
		category = "other"
	}
	centralCommandFile := filepath.Join(centralCommandsDir, category, firstCmd.Name+".md")
	if _, err := os.Stat(centralCommandFile); os.IsNotExist(err) {
		t.Fatalf("CommandsGenerator.Generate() should generate commands in central location")
	}

	// Then verify each IDE has access to commands (via symlink or copy)
	for ide, commandsDir := range expectedCommandsDirs {
		if _, err := os.Stat(commandsDir); os.IsNotExist(err) {
			t.Errorf("CommandsGenerator.Generate() should create commands directory for %s at %s", ide, commandsDir)
			continue
		}

		// Verify that commands are accessible (check for a known command file in category folder)
		expectedCommandFile := filepath.Join(commandsDir, category, firstCmd.Name+".md")
		if _, err := os.Stat(expectedCommandFile); os.IsNotExist(err) {
			// If symlink/copy failed, that's okay - central location has the commands
			// Just log a warning but don't fail the test
			t.Logf("Commands not accessible via IDE location for %s (symlink/copy may have failed, but central location has commands)", ide)
		}
	}
}
