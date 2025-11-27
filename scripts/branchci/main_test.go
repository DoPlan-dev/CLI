package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunBranchCI_WritesWorkflow(t *testing.T) {
	tmp := t.TempDir()
	matrixPath := filepath.Join(tmp, "matrix.json")
	if err := os.WriteFile(matrixPath, []byte(`{
		"generated_at":"2024-01-01T00:00:00Z",
		"branches":[{"prefix":"feat/","jobs":["lint"],"required":["lint"]}]
	}`), 0644); err != nil {
		t.Fatalf("failed to write matrix: %v", err)
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	code := run([]string{"--matrix", matrixPath, "--out", filepath.Join(tmp, "out")}, out, errOut)
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d (stderr=%s)", code, errOut.String())
	}
	workflowPath := filepath.Join(tmp, "out", "task-branches.yml")
	if _, err := os.Stat(workflowPath); err != nil {
		t.Fatalf("expected workflow file, got %v", err)
	}
	if out.Len() == 0 {
		t.Fatal("expected success output")
	}
}

func TestRunBranchCI_BadDir(t *testing.T) {
	tmp := t.TempDir()
	block := filepath.Join(tmp, "blocked")
	if err := os.WriteFile(block, []byte("file"), 0644); err != nil {
		t.Fatalf("failed to create blocking file: %v", err)
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	code := run([]string{"--out", block}, out, errOut)
	if code == 0 {
		t.Fatalf("expected non-zero exit code for invalid dir")
	}
	if errOut.Len() == 0 {
		t.Fatal("expected error output")
	}
}
