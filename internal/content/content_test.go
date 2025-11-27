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
	files, err := ReadCommandDir("core")
	if err != nil {
		t.Fatalf("ReadCommandDir() error = %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected core command directory to contain files")
	}

	data, err := ReadCommandFile("core/build.md")
	if err != nil {
		t.Fatalf("ReadCommandFile() error = %v", err)
	}
	if !strings.Contains(string(data), "/build") {
		t.Errorf("expected build command markdown to mention /build")
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
