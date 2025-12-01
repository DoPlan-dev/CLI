package content

import (
	"io/fs"
	"strings"
	"testing"
)

func TestAgentsEmbeddedFS(t *testing.T) {
	files, err := ReadAgentDir("engineering")
	if err != nil {
		t.Fatalf("ReadAgentDir() error = %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected engineering directory to contain agent files")
	}

	data, err := ReadAgentFile("engineering/engineering_lead.md")
	if err != nil {
		t.Fatalf("ReadAgentFile() error = %v", err)
	}
	if !strings.Contains(string(data), "Engineering Lead") {
		t.Errorf("expected engineering lead content, got: %s", data[:50])
	}

	count := 0
	err = WalkAgentDir("engineering", func(path string, _ fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		count++
		return nil
	})
	if err != nil {
		t.Fatalf("WalkAgentDir() error = %v", err)
	}
	if count == 0 {
		t.Error("expected WalkAgentDir to traverse files")
	}
}

func TestCommandsEmbeddedFS(t *testing.T) {
	files, err := ReadCommandDir("developing")
	if err != nil {
		t.Fatalf("ReadCommandDir() error = %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected developing command directory to contain files")
	}

	data, err := ReadCommandFile("developing/dev.md")
	if err != nil {
		t.Fatalf("ReadCommandFile() error = %v", err)
	}
	if !strings.Contains(string(data), "/dev") {
		t.Errorf("expected dev command markdown to mention /dev")
	}
}

func TestTemplatesEmbeddedFS(t *testing.T) {
	files, err := ReadTemplateDir("agents")
	if err != nil {
		t.Fatalf("ReadTemplateDir() error = %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected agents template directory to contain templates")
	}

	data, err := ReadTemplateFile("agents/agent.md.tmpl")
	if err != nil {
		t.Fatalf("ReadTemplateFile() error = %v", err)
	}
	if !strings.Contains(string(data), "{{.Name}}") {
		t.Errorf("expected template to include {{.Name}}, got: %s", data[:60])
	}
}

func TestAgentsFS(t *testing.T) {
	fs := Agents()
	if fs == nil {
		t.Fatal("Agents() should not return nil")
	}

	// Verify we can read from the filesystem
	_, err := fs.Open("agents/engineering/engineering_lead.md")
	if err != nil {
		t.Errorf("Agents() filesystem should be accessible: %v", err)
	}
}

func TestCommandsFS(t *testing.T) {
	fs := Commands()
	if fs == nil {
		t.Fatal("Commands() should not return nil")
	}

	// Verify we can read from the filesystem
	_, err := fs.Open("commands/developing/dev.md")
	if err != nil {
		t.Errorf("Commands() filesystem should be accessible: %v", err)
	}
}

func TestTemplatesFS(t *testing.T) {
	fs := Templates()
	if fs == nil {
		t.Fatal("Templates() should not return nil")
	}

	// Verify we can read from the filesystem
	_, err := fs.Open("templates/agents/agent.md.tmpl")
	if err != nil {
		t.Errorf("Templates() filesystem should be accessible: %v", err)
	}
}

func TestWalkCommandDir(t *testing.T) {
	count := 0
	err := WalkCommandDir("", func(path string, _ fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		count++
		return nil
	})
	if err != nil {
		t.Fatalf("WalkCommandDir() error = %v", err)
	}
	if count == 0 {
		t.Error("expected WalkCommandDir to traverse files")
	}

	// Test with specific root
	count = 0
	err = WalkCommandDir("developing", func(path string, _ fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		count++
		return nil
	})
	if err != nil {
		t.Fatalf("WalkCommandDir('developing') error = %v", err)
	}
	if count == 0 {
		t.Error("expected WalkCommandDir('developing') to traverse files")
	}
}

func TestWalkTemplateDir(t *testing.T) {
	count := 0
	err := WalkTemplateDir("", func(path string, _ fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		count++
		return nil
	})
	if err != nil {
		t.Fatalf("WalkTemplateDir() error = %v", err)
	}
	if count == 0 {
		t.Error("expected WalkTemplateDir to traverse files")
	}

	// Test with specific root
	count = 0
	err = WalkTemplateDir("agents", func(path string, _ fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		count++
		return nil
	})
	if err != nil {
		t.Fatalf("WalkTemplateDir('agents') error = %v", err)
	}
	if count == 0 {
		t.Error("expected WalkTemplateDir('agents') to traverse files")
	}
}
