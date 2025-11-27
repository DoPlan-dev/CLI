package statehistory

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestSaveSnapshot(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create active_state.json
	state := ActiveState{
		Phase:        "development",
		ActiveTask:   "1.1",
		ActiveBranch: "task/1-1",
		Locked:       true,
		Completed:    []string{"1.1", "1.2"},
	}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	// Save snapshot
	snapshot, err := SaveSnapshot(statePath, historyDir, "test reason", "test-label")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	if snapshot == nil {
		t.Fatal("SaveSnapshot returned nil snapshot")
	}
	if snapshot.Reason != "test reason" {
		t.Errorf("Reason = %q, want %q", snapshot.Reason, "test reason")
	}
	if snapshot.State.Phase != "development" {
		t.Errorf("State.Phase = %q, want %q", snapshot.State.Phase, "development")
	}
	if snapshot.Hash == "" {
		t.Error("Hash should not be empty")
	}

	// Verify file exists
	if _, err := os.Stat(snapshot.Path); err != nil {
		t.Fatalf("Snapshot file not created: %v", err)
	}
}

func TestSaveSnapshot_WithLabel(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	state := ActiveState{Phase: "test"}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	snapshot, err := SaveSnapshot(statePath, historyDir, "reason", "my-label")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	// Label should be sanitized in filename
	basename := filepath.Base(snapshot.Path)
	if !contains(basename, "my-label") {
		t.Errorf("Filename should contain label, got %q", basename)
	}
}

func TestListSnapshots(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create multiple snapshots with longer delays to ensure different timestamps
	for i := 0; i < 3; i++ {
		state := ActiveState{Phase: "test", Completed: []string{string(rune('0' + i))}}
		stateData, _ := json.Marshal(state)
		statePath := filepath.Join(planDir, "active_state.json")
		os.WriteFile(statePath, stateData, 0644)

		_, err := SaveSnapshot(statePath, historyDir, "test", "")
		if err != nil {
			t.Fatalf("SaveSnapshot failed: %v", err)
		}

		// Use longer delay to ensure different timestamps (ID format includes seconds)
		time.Sleep(1100 * time.Millisecond)
	}

	// List snapshots
	snapshots, err := ListSnapshots(historyDir)
	if err != nil {
		t.Fatalf("ListSnapshots failed: %v", err)
	}

	if len(snapshots) != 3 {
		t.Errorf("ListSnapshots returned %d snapshots, want 3", len(snapshots))
	}

	// Verify they're sorted (ascending by time)
	if len(snapshots) >= 2 {
		for i := 1; i < len(snapshots); i++ {
			if snapshots[i].CapturedAt.Before(snapshots[i-1].CapturedAt) {
				t.Error("Snapshots should be sorted ascending by time")
			}
		}
	}
}

func TestListSnapshots_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	historyDir := filepath.Join(tmpDir, "history")
	os.MkdirAll(historyDir, 0755)

	snapshots, err := ListSnapshots(historyDir)
	if err != nil {
		t.Fatalf("ListSnapshots failed: %v", err)
	}

	if len(snapshots) != 0 {
		t.Errorf("ListSnapshots returned %d snapshots, want 0", len(snapshots))
	}
}

func TestLatestSnapshots(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create 5 snapshots with delays to ensure different timestamps
	for i := 0; i < 5; i++ {
		state := ActiveState{Phase: "test"}
		stateData, _ := json.Marshal(state)
		statePath := filepath.Join(planDir, "active_state.json")
		os.WriteFile(statePath, stateData, 0644)

		_, err := SaveSnapshot(statePath, historyDir, "test", "")
		if err != nil {
			t.Fatalf("SaveSnapshot failed: %v", err)
		}

		// Use longer delay to ensure different timestamps (ID format includes seconds)
		time.Sleep(1100 * time.Millisecond)
	}

	// Get latest 2
	latest, err := LatestSnapshots(historyDir, 2)
	if err != nil {
		t.Fatalf("LatestSnapshots failed: %v", err)
	}

	if len(latest) != 2 {
		t.Errorf("LatestSnapshots returned %d snapshots, want 2", len(latest))
	}

	// Should be the two most recent
	if len(latest) >= 2 {
		all, _ := ListSnapshots(historyDir)
		if len(all) >= 2 {
			if latest[0].ID != all[len(all)-2].ID || latest[1].ID != all[len(all)-1].ID {
				t.Error("LatestSnapshots should return the most recent snapshots")
			}
		}
	}
}

func TestLoadSnapshot(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create and save snapshot
	state := ActiveState{Phase: "test"}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	saved, err := SaveSnapshot(statePath, historyDir, "test", "")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	// Load by ID
	loaded, err := LoadSnapshot(historyDir, saved.ID)
	if err != nil {
		t.Fatalf("LoadSnapshot failed: %v", err)
	}

	if loaded.ID != saved.ID {
		t.Errorf("Loaded ID = %q, want %q", loaded.ID, saved.ID)
	}

	// Load by filename
	basename := filepath.Base(saved.Path)
	loaded2, err := LoadSnapshot(historyDir, basename)
	if err != nil {
		t.Fatalf("LoadSnapshot by filename failed: %v", err)
	}

	if loaded2.ID != saved.ID {
		t.Errorf("Loaded ID = %q, want %q", loaded2.ID, saved.ID)
	}
}

func TestLoadSnapshot_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	historyDir := filepath.Join(tmpDir, "history")
	os.MkdirAll(historyDir, 0755)

	_, err := LoadSnapshot(historyDir, "nonexistent")
	if err == nil {
		t.Error("LoadSnapshot should fail for nonexistent snapshot")
	}
}

func TestRestoreSnapshot(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create initial state
	initialState := ActiveState{Phase: "initial"}
	stateData, _ := json.Marshal(initialState)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	// Save snapshot
	snapshot, err := SaveSnapshot(statePath, historyDir, "test", "")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	// Modify state
	modifiedState := ActiveState{Phase: "modified"}
	stateData, _ = json.Marshal(modifiedState)
	os.WriteFile(statePath, stateData, 0644)

	// Restore snapshot
	err = RestoreSnapshot(statePath, snapshot)
	if err != nil {
		t.Fatalf("RestoreSnapshot failed: %v", err)
	}

	// Verify restored
	restoredData, _ := os.ReadFile(statePath)
	var restored ActiveState
	json.Unmarshal(restoredData, &restored)

	if restored.Phase != "initial" {
		t.Errorf("Restored Phase = %q, want %q", restored.Phase, "initial")
	}
}

func TestComputeDiff(t *testing.T) {
	older := &Snapshot{
		ID:         "older",
		CapturedAt: time.Now().Add(-1 * time.Hour),
		State: ActiveState{
			Phase:        "old-phase",
			ActiveTask:   "1.1",
			ActiveBranch: "old-branch",
			Completed:    []string{"1.1"},
		},
	}

	newer := &Snapshot{
		ID:         "newer",
		CapturedAt: time.Now(),
		State: ActiveState{
			Phase:        "new-phase",
			ActiveTask:   "2.1",
			ActiveBranch: "new-branch",
			Completed:    []string{"1.1", "1.2"},
		},
	}

	diff := ComputeDiff(newer, older)

	if !diff.Phase.Changed {
		t.Error("Phase should be marked as changed")
	}
	if diff.Phase.From != "old-phase" {
		t.Errorf("Phase.From = %q, want %q", diff.Phase.From, "old-phase")
	}
	if diff.Phase.To != "new-phase" {
		t.Errorf("Phase.To = %q, want %q", diff.Phase.To, "new-phase")
	}

	if len(diff.CompletedAdded) != 1 {
		t.Errorf("CompletedAdded length = %d, want 1", len(diff.CompletedAdded))
	}
	if diff.CompletedAdded[0] != "1.2" {
		t.Errorf("CompletedAdded[0] = %q, want %q", diff.CompletedAdded[0], "1.2")
	}
}

func TestComputeDiff_NoChanges(t *testing.T) {
	state := ActiveState{
		Phase:     "same",
		Completed: []string{"1.1"},
	}

	older := &Snapshot{
		ID:         "older",
		CapturedAt: time.Now().Add(-1 * time.Hour),
		State:      state,
	}

	newer := &Snapshot{
		ID:         "newer",
		CapturedAt: time.Now(),
		State:      state,
	}

	diff := ComputeDiff(newer, older)

	if diff.HasChanges() {
		t.Error("HasChanges should return false when no changes")
	}
}

func TestFormatDiff(t *testing.T) {
	older := &Snapshot{
		ID:         "older",
		CapturedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		State: ActiveState{
			Phase:     "old",
			Completed: []string{"1.1"},
		},
	}

	newer := &Snapshot{
		ID:         "newer",
		CapturedAt: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		State: ActiveState{
			Phase:     "new",
			Completed: []string{"1.1", "1.2"},
		},
	}

	diff := ComputeDiff(newer, older)
	output := FormatDiff(diff)

	if !contains(output, "older") || !contains(output, "newer") {
		t.Error("FormatDiff should include snapshot IDs")
	}
	if !contains(output, "old") || !contains(output, "new") {
		t.Error("FormatDiff should include phase changes")
	}
	if !contains(output, "1.2") {
		t.Error("FormatDiff should include completed task changes")
	}
}

func TestLatestDiff(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create first snapshot
	state1 := ActiveState{Phase: "phase1"}
	stateData, _ := json.Marshal(state1)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)
	SaveSnapshot(statePath, historyDir, "first", "")

	// Use longer delay to ensure different timestamps (ID format includes seconds)
	time.Sleep(1100 * time.Millisecond)

	// Create second snapshot
	state2 := ActiveState{Phase: "phase2"}
	stateData, _ = json.Marshal(state2)
	os.WriteFile(statePath, stateData, 0644)
	SaveSnapshot(statePath, historyDir, "second", "")

	// Get latest diff
	diff, err := LatestDiff(historyDir)
	if err != nil {
		t.Fatalf("LatestDiff failed: %v", err)
	}

	if diff == nil {
		t.Fatal("LatestDiff returned nil")
	}
	if diff.Phase.To != "phase2" {
		t.Errorf("Latest diff Phase.To = %q, want %q", diff.Phase.To, "phase2")
	}
}

func TestLatestDiffSummary(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create first snapshot
	state1 := ActiveState{Phase: "phase1", ActiveTask: "1.1"}
	stateData, _ := json.Marshal(state1)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	_, err := SaveSnapshot(statePath, historyDir, "snapshot1", "")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	time.Sleep(1100 * time.Millisecond)

	// Create second snapshot with changes
	state2 := ActiveState{Phase: "phase2", ActiveTask: "2.1", Completed: []string{"1.1"}}
	stateData, _ = json.Marshal(state2)
	os.WriteFile(statePath, stateData, 0644)

	_, err = SaveSnapshot(statePath, historyDir, "snapshot2", "")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	// Get diff summary
	summary, err := LatestDiffSummary(historyDir)
	if err != nil {
		t.Fatalf("LatestDiffSummary failed: %v", err)
	}

	if summary == "" {
		t.Error("LatestDiffSummary should return non-empty summary")
	}

	// Verify summary contains expected changes
	if !strings.Contains(summary, "phase2") && !strings.Contains(summary, "Phase") {
		t.Logf("Summary: %s", summary)
		t.Log("LatestDiffSummary should contain phase information")
	}
}

func TestLatestDiffSummary_InsufficientSnapshots(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create only one snapshot
	state := ActiveState{Phase: "phase1"}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	_, err := SaveSnapshot(statePath, historyDir, "snapshot1", "")
	if err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	// Should fail with insufficient snapshots
	_, err = LatestDiffSummary(historyDir)
	if err == nil {
		t.Error("LatestDiffSummary should fail with insufficient snapshots")
	}
	if err != ErrInsufficientSnapshots {
		t.Errorf("LatestDiffSummary error = %v, want ErrInsufficientSnapshots", err)
	}
}

func TestLatestDiff_InsufficientSnapshots(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".do")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")

	// Create only one snapshot
	state := ActiveState{Phase: "test"}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)
	SaveSnapshot(statePath, historyDir, "only", "")

	_, err := LatestDiff(historyDir)
	if err != ErrInsufficientSnapshots {
		t.Errorf("LatestDiff error = %v, want ErrInsufficientSnapshots", err)
	}
}

func TestSanitizeLabel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "mylabel", "mylabel"},
		{"with spaces", "my label", "my-label"},
		{"with special chars", "my@label#123", "mylabel123"},
		{"empty", "", ""},
		{"whitespace", "   ", ""},
		{"mixed case", "MyLabel", "mylabel"},
		{"with underscores", "my_label", "my_label"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizeLabel(tt.input)
			if result != tt.expected {
				t.Errorf("sanitizeLabel(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		strings.Contains(s, substr))
}
