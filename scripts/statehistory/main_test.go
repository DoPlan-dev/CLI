package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/internal/statehistory"
)

func writeActiveState(t *testing.T, path string, completed []string) {
	t.Helper()
	state := statehistory.ActiveState{
		Phase:      "build",
		Completed:  completed,
		ActiveTask: "1.1",
	}
	data, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("marshal state: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("write state: %v", err)
	}
}

func TestSnapshotListDiffRestore(t *testing.T) {
	tmp := t.TempDir()
	statePath := filepath.Join(tmp, "active_state.json")
	historyDir := filepath.Join(tmp, "history")

	writeActiveState(t, statePath, []string{"1.1"})

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	if code := run([]string{"snapshot", "--state", statePath, "--history", historyDir, "--reason", "initial"}, out, errOut); code != 0 {
		t.Fatalf("snapshot exit=%d stderr=%s", code, errOut.String())
	}

	writeActiveState(t, statePath, []string{"1.1", "2.1"})
	time.Sleep(1100 * time.Millisecond)
	out.Reset()
	errOut.Reset()
	if code := run([]string{"snapshot", "--state", statePath, "--history", historyDir, "--reason", "second"}, out, errOut); code != 0 {
		t.Fatalf("second snapshot exit=%d stderr=%s", code, errOut.String())
	}

	out.Reset()
	errOut.Reset()
	if code := run([]string{"list", "--history", historyDir, "--json"}, out, errOut); code != 0 {
		t.Fatalf("list exit=%d stderr=%s", code, errOut.String())
	}
	if !bytes.Contains(out.Bytes(), []byte(`"id"`)) {
		t.Fatalf("expected json output, got %s", out.String())
	}

	out.Reset()
	errOut.Reset()
	if code := run([]string{"diff", "--history", historyDir}, out, errOut); code != 0 {
		t.Fatalf("diff exit=%d stderr=%s", code, errOut.String())
	}
	if out.Len() == 0 {
		t.Fatal("expected diff output")
	}

	snaps, err := statehistory.ListSnapshots(historyDir)
	if err != nil {
		t.Fatalf("list snapshots: %v", err)
	}
	if len(snaps) < 1 {
		t.Fatal("expected at least one snapshot")
	}
	target := snaps[0]

	writeActiveState(t, statePath, []string{})
	out.Reset()
	errOut.Reset()
	if code := run([]string{"restore", "--history", historyDir, "--state", statePath, "--file", target.ID, "--yes", "--snapshot=false"}, out, errOut); code != 0 {
		t.Fatalf("restore exit=%d stderr=%s", code, errOut.String())
	}
	data, err := os.ReadFile(statePath)
	if err != nil {
		t.Fatalf("read state: %v", err)
	}
	var restored statehistory.ActiveState
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(restored.Completed) == 0 {
		t.Fatalf("expected restored completed tasks, got %+v", restored)
	}
}
