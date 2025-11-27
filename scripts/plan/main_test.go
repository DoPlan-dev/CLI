package main

import (
	"bytes"
	"errors"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/generator"
)

func TestRunPlanSuccess(t *testing.T) {
	t.Cleanup(func() { scaffoldPlanHierarchy = generator.ScaffoldPlanHierarchy })
	var capturedPath string
	scaffoldPlanHierarchy = func(path string) error {
		capturedPath = path
		return nil
	}

	tmp := t.TempDir()
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}

	if code := run([]string{"--project", tmp}, out, errOut); code != 0 {
		t.Fatalf("expected exit code 0, got %d (stderr=%s)", code, errOut.String())
	}
	if capturedPath != filepath.Clean(tmp) {
		t.Fatalf("expected scaffold path %s, got %s", tmp, capturedPath)
	}
	if out.Len() == 0 {
		t.Fatal("expected success output")
	}
}

func TestRunPlanFailure(t *testing.T) {
	t.Cleanup(func() { scaffoldPlanHierarchy = generator.ScaffoldPlanHierarchy })
	scaffoldPlanHierarchy = func(string) error {
		return errors.New("boom")
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}

	if code := run([]string{"--project", "."}, out, errOut); code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if errOut.Len() == 0 {
		t.Fatal("expected error output")
	}
}
