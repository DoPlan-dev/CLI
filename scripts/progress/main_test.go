package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/progress"
)

func TestRunProgressPlain(t *testing.T) {
	t.Cleanup(func() {
		computeProgress = progress.Compute
		formatProgress = progress.FormatPlain
	})

	computeProgress = func(root string) (*progress.Report, error) {
		return &progress.Report{Phase: "build"}, nil
	}
	formatProgress = func(*progress.Report) string {
		return "ok"
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	if code := run([]string{"--root", "."}, out, errOut); code != 0 {
		t.Fatalf("expected exit code 0, got %d (stderr=%s)", code, errOut.String())
	}
	if out.String() != "ok\n" {
		t.Fatalf("unexpected stdout: %s", out.String())
	}
}

func TestRunProgressJSON(t *testing.T) {
	t.Cleanup(func() {
		computeProgress = progress.Compute
		formatProgress = progress.FormatPlain
	})

	computeProgress = func(root string) (*progress.Report, error) {
		return &progress.Report{Phase: "idea"}, nil
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	if code := run([]string{"--root", ".", "--json", "--diff-struct"}, out, errOut); code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !bytes.Contains(out.Bytes(), []byte(`"Phase": "idea"`)) && !bytes.Contains(out.Bytes(), []byte(`"phase": "idea"`)) {
		t.Fatalf("expected JSON payload, got %s", out.String())
	}
}

func TestRunProgressError(t *testing.T) {
	t.Cleanup(func() {
		computeProgress = progress.Compute
		formatProgress = progress.FormatPlain
	})

	computeProgress = func(root string) (*progress.Report, error) {
		return nil, errors.New("boom")
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	if code := run([]string{"--root", "."}, out, errOut); code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if errOut.Len() == 0 {
		t.Fatal("expected error output")
	}
}
