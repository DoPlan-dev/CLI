package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadCommandsFromFiles(t *testing.T) {
	commands, err := LoadCommandsFromFiles()
	if err != nil {
		t.Fatalf("LoadCommandsFromFiles() error = %v", err)
	}
	if len(commands) == 0 {
		t.Fatal("expected commands to be loaded from embedded files")
	}

	found := false
	for _, cmd := range commands {
		if strings.EqualFold(cmd.Name, "build") {
			found = true
			if !strings.Contains(cmd.Action, "/build") && cmd.Trigger == "" {
				t.Errorf("expected build command to include trigger/action")
			}
		}
	}
	if !found {
		t.Error("expected to find build command in embedded commands")
	}
}

func TestLoadCommandsFromDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	if err := ExtractCommands(tmpDir); err != nil {
		t.Fatalf("ExtractCommands() error = %v", err)
	}

	cmds, err := LoadCommandsFromDirectory(tmpDir)
	if err != nil {
		t.Fatalf("LoadCommandsFromDirectory() error = %v", err)
	}
	if len(cmds) == 0 {
		t.Fatal("expected commands to load from extracted directory")
	}
}

func TestLoadCommandsFromDirectory_InvalidMarkdown(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "invalid.md")
	if err := os.WriteFile(file, []byte("# invalid"), 0644); err != nil {
		t.Fatalf("failed to write invalid file: %v", err)
	}

	if _, err := LoadCommandsFromDirectory(tmpDir); err == nil {
		t.Fatal("expected error for invalid markdown")
	}
}

func TestExtractCommands(t *testing.T) {
	tmpDir := t.TempDir()
	if err := ExtractCommands(tmpDir); err != nil {
		t.Fatalf("ExtractCommands() error = %v", err)
	}

	coreDir := filepath.Join(tmpDir, "core")
	if _, err := os.Stat(coreDir); os.IsNotExist(err) {
		t.Fatalf("expected core directory after extraction")
	}

	buildCmd := filepath.Join(coreDir, "build.md")
	if _, err := os.Stat(buildCmd); os.IsNotExist(err) {
		t.Fatalf("expected build command markdown after extraction")
	}
	data, err := os.ReadFile(buildCmd)
	if err != nil {
		t.Fatalf("failed to read extracted command: %v", err)
	}
	if !strings.Contains(string(data), "/build") {
		t.Errorf("expected build command content to mention /build")
	}
}
