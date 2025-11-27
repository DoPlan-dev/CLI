package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRunFeedbackSuccess(t *testing.T) {
	tmp := t.TempDir()
	now = func() time.Time { return time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC) }
	defer func() { now = time.Now }()

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	code := run([]string{
		"--project", tmp,
		"--type", "bug",
		"--title", "Broken",
		"--details", "fix me",
		"--author", "QA",
	}, out, errOut)
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d: %s", code, errOut.String())
	}

	mdPath := filepath.Join(tmp, "Docs", "history", "feedback.md")
	jsonPath := filepath.Join(tmp, "Docs", "history", "feedback.json")
	if _, err := os.Stat(mdPath); err != nil {
		t.Fatalf("expected markdown file, got %v", err)
	}
	if _, err := os.Stat(jsonPath); err != nil {
		t.Fatalf("expected json file, got %v", err)
	}
}

func TestRunFeedbackRequiresTitle(t *testing.T) {
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	if code := run([]string{"--project", ".", "--title", "   "}, out, errOut); code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if errOut.Len() == 0 {
		t.Fatal("expected error output")
	}
}
