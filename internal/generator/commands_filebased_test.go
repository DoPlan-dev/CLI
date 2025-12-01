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
		if strings.EqualFold(cmd.Name, "dev") {
			found = true
			if !strings.Contains(cmd.Action, "/dev") && cmd.Trigger == "" {
				t.Errorf("expected dev command to include trigger/action")
			}
		}
	}
	if !found {
		t.Error("expected to find dev command in embedded commands")
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

	developingDir := filepath.Join(tmpDir, "developing")
	if _, err := os.Stat(developingDir); os.IsNotExist(err) {
		t.Fatalf("expected developing directory after extraction")
	}

	devCmd := filepath.Join(developingDir, "dev.md")
	if _, err := os.Stat(devCmd); os.IsNotExist(err) {
		t.Fatalf("expected dev command markdown after extraction")
	}
	data, err := os.ReadFile(devCmd)
	if err != nil {
		t.Fatalf("failed to read extracted command: %v", err)
	}
	if !strings.Contains(string(data), "/dev") {
		t.Errorf("expected dev command content to mention /dev")
	}
}
