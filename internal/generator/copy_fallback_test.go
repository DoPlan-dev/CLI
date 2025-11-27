package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyAgents(t *testing.T) {
	tmpDir := t.TempDir()

	// Create central agents directory with category folders
	centralAgentsDir := filepath.Join(tmpDir, "central", "agents")
	category1Dir := filepath.Join(centralAgentsDir, "leadership")
	category2Dir := filepath.Join(centralAgentsDir, "development")

	if err := os.MkdirAll(category1Dir, 0755); err != nil {
		t.Fatalf("Failed to create category1 dir: %v", err)
	}
	if err := os.MkdirAll(category2Dir, 0755); err != nil {
		t.Fatalf("Failed to create category2 dir: %v", err)
	}

	// Create agent files in categories
	agent1File := filepath.Join(category1Dir, "project_orchestrator.md")
	agent2File := filepath.Join(category2Dir, "frontend_lead.md")
	if err := os.WriteFile(agent1File, []byte("# Project Orchestrator"), 0644); err != nil {
		t.Fatalf("Failed to create agent1 file: %v", err)
	}
	if err := os.WriteFile(agent2File, []byte("# Frontend Lead"), 0644); err != nil {
		t.Fatalf("Failed to create agent2 file: %v", err)
	}

	// Create IDE agents directory
	ideAgentsDir := filepath.Join(tmpDir, "ide", "agents")
	if err := os.MkdirAll(ideAgentsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE agents dir: %v", err)
	}

	// Create sample agents list (not used in copyAgents but required for signature)
	agents := []Agent{
		{Name: "project_orchestrator", Category: "leadership"},
		{Name: "frontend_lead", Category: "development"},
	}

	// Test copyAgents
	if err := copyAgents(ideAgentsDir, centralAgentsDir, agents); err != nil {
		t.Fatalf("copyAgents() error = %v", err)
	}

	// Verify category folders were copied
	ideCategory1Dir := filepath.Join(ideAgentsDir, "leadership")
	ideCategory2Dir := filepath.Join(ideAgentsDir, "development")

	if _, err := os.Stat(ideCategory1Dir); os.IsNotExist(err) {
		t.Error("copyAgents() should copy category folders")
	}
	if _, err := os.Stat(ideCategory2Dir); os.IsNotExist(err) {
		t.Error("copyAgents() should copy all category folders")
	}

	// Verify agent files were copied
	ideAgent1File := filepath.Join(ideCategory1Dir, "project_orchestrator.md")
	ideAgent2File := filepath.Join(ideCategory2Dir, "frontend_lead.md")

	if _, err := os.Stat(ideAgent1File); os.IsNotExist(err) {
		t.Error("copyAgents() should copy agent files")
	}
	if _, err := os.Stat(ideAgent2File); os.IsNotExist(err) {
		t.Error("copyAgents() should copy all agent files")
	}

	// Verify file content
	content1, err := os.ReadFile(ideAgent1File)
	if err != nil {
		t.Fatalf("Failed to read copied agent1 file: %v", err)
	}
	if string(content1) != "# Project Orchestrator" {
		t.Error("copyAgents() should preserve file content")
	}
}

func TestCopyAgents_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	centralAgentsDir := filepath.Join(tmpDir, "central", "agents")
	if err := os.MkdirAll(centralAgentsDir, 0755); err != nil {
		t.Fatalf("Failed to create central dir: %v", err)
	}

	ideAgentsDir := filepath.Join(tmpDir, "ide", "agents")
	if err := os.MkdirAll(ideAgentsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Should not error on empty directory
	if err := copyAgents(ideAgentsDir, centralAgentsDir, []Agent{}); err != nil {
		t.Errorf("copyAgents() should handle empty directory, got error: %v", err)
	}
}

func TestCopyAgents_NonExistentCentralDir(t *testing.T) {
	tmpDir := t.TempDir()

	centralAgentsDir := filepath.Join(tmpDir, "nonexistent", "agents")
	ideAgentsDir := filepath.Join(tmpDir, "ide", "agents")
	if err := os.MkdirAll(ideAgentsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Should return error for non-existent central directory
	if err := copyAgents(ideAgentsDir, centralAgentsDir, []Agent{}); err == nil {
		t.Error("copyAgents() should return error for non-existent central directory")
	}
}

func TestCopyCommands(t *testing.T) {
	tmpDir := t.TempDir()

	// Create central commands directory with category folders
	centralCommandsDir := filepath.Join(tmpDir, "central", "commands")
	category1Dir := filepath.Join(centralCommandsDir, "core")
	category2Dir := filepath.Join(centralCommandsDir, "tools")

	if err := os.MkdirAll(category1Dir, 0755); err != nil {
		t.Fatalf("Failed to create category1 dir: %v", err)
	}
	if err := os.MkdirAll(category2Dir, 0755); err != nil {
		t.Fatalf("Failed to create category2 dir: %v", err)
	}

	// Create command files in categories
	cmd1File := filepath.Join(category1Dir, "hello.md")
	cmd2File := filepath.Join(category2Dir, "github.md")
	if err := os.WriteFile(cmd1File, []byte("# Hello Command"), 0644); err != nil {
		t.Fatalf("Failed to create cmd1 file: %v", err)
	}
	if err := os.WriteFile(cmd2File, []byte("# GitHub Command"), 0644); err != nil {
		t.Fatalf("Failed to create cmd2 file: %v", err)
	}

	// Create IDE commands directory
	ideCommandsDir := filepath.Join(tmpDir, "ide", "commands")
	if err := os.MkdirAll(ideCommandsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE commands dir: %v", err)
	}

	// Create sample commands list
	commands := []Command{
		{Name: "hello", Category: "core"},
		{Name: "github", Category: "tools"},
	}

	// Test copyCommands
	if err := copyCommands(ideCommandsDir, centralCommandsDir, commands); err != nil {
		t.Fatalf("copyCommands() error = %v", err)
	}

	// Verify category folders were copied
	ideCategory1Dir := filepath.Join(ideCommandsDir, "core")
	ideCategory2Dir := filepath.Join(ideCommandsDir, "tools")

	if _, err := os.Stat(ideCategory1Dir); os.IsNotExist(err) {
		t.Error("copyCommands() should copy category folders")
	}
	if _, err := os.Stat(ideCategory2Dir); os.IsNotExist(err) {
		t.Error("copyCommands() should copy all category folders")
	}

	// Verify command files were copied
	ideCmd1File := filepath.Join(ideCategory1Dir, "hello.md")
	ideCmd2File := filepath.Join(ideCategory2Dir, "github.md")

	if _, err := os.Stat(ideCmd1File); os.IsNotExist(err) {
		t.Error("copyCommands() should copy command files")
	}
	if _, err := os.Stat(ideCmd2File); os.IsNotExist(err) {
		t.Error("copyCommands() should copy all command files")
	}

	// Verify file content
	content1, err := os.ReadFile(ideCmd1File)
	if err != nil {
		t.Fatalf("Failed to read copied cmd1 file: %v", err)
	}
	if string(content1) != "# Hello Command" {
		t.Error("copyCommands() should preserve file content")
	}
}

func TestCopyCommands_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	centralCommandsDir := filepath.Join(tmpDir, "central", "commands")
	if err := os.MkdirAll(centralCommandsDir, 0755); err != nil {
		t.Fatalf("Failed to create central dir: %v", err)
	}

	ideCommandsDir := filepath.Join(tmpDir, "ide", "commands")
	if err := os.MkdirAll(ideCommandsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Should not error on empty directory
	if err := copyCommands(ideCommandsDir, centralCommandsDir, []Command{}); err != nil {
		t.Errorf("copyCommands() should handle empty directory, got error: %v", err)
	}
}

func TestCopyCommands_NonExistentCentralDir(t *testing.T) {
	tmpDir := t.TempDir()

	centralCommandsDir := filepath.Join(tmpDir, "nonexistent", "commands")
	ideCommandsDir := filepath.Join(tmpDir, "ide", "commands")
	if err := os.MkdirAll(ideCommandsDir, 0755); err != nil {
		t.Fatalf("Failed to create IDE dir: %v", err)
	}

	// Should return error for non-existent central directory
	if err := copyCommands(ideCommandsDir, centralCommandsDir, []Command{}); err == nil {
		t.Error("copyCommands() should return error for non-existent central directory")
	}
}

