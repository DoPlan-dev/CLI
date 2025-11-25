package progress

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/statehistory"
)

func TestTaskStats(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected TaskStats
	}{
		{
			name: "mixed tasks",
			content: `# Tasks
- [x] Completed task 1
- [ ] Pending task 1
- [x] Completed task 2
- [ ] Pending task 2
- [ ] Pending task 3`,
			expected: TaskStats{
				Total:     5,
				Completed: 2,
				Pending:   3,
				NextUp:    "Pending task 1",
				Sample:    []string{"Pending task 1", "Pending task 2", "Pending task 3"},
			},
		},
		{
			name: "all completed",
			content: `# Tasks
- [x] Task 1
- [x] Task 2`,
			expected: TaskStats{
				Total:     2,
				Completed: 2,
				Pending:   0,
				NextUp:    "",
				Sample:    []string{},
			},
		},
		{
			name: "all pending",
			content: `# Tasks
- [ ] Task 1
- [ ] Task 2`,
			expected: TaskStats{
				Total:     2,
				Completed: 0,
				Pending:   2,
				NextUp:    "Task 1",
				Sample:    []string{"Task 1", "Task 2"},
			},
		},
		{
			name:     "empty file",
			content:  "",
			expected: TaskStats{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			tasksPath := filepath.Join(tmpDir, "TASKS.md")
			os.WriteFile(tasksPath, []byte(tt.content), 0644)

			stats, err := parseTasks(tasksPath)
			if err != nil {
				t.Fatalf("parseTasks failed: %v", err)
			}

			if stats.Total != tt.expected.Total {
				t.Errorf("Total = %d, want %d", stats.Total, tt.expected.Total)
			}
			if stats.Completed != tt.expected.Completed {
				t.Errorf("Completed = %d, want %d", stats.Completed, tt.expected.Completed)
			}
			if stats.Pending != tt.expected.Pending {
				t.Errorf("Pending = %d, want %d", stats.Pending, tt.expected.Pending)
			}
			if stats.NextUp != tt.expected.NextUp {
				t.Errorf("NextUp = %q, want %q", stats.NextUp, tt.expected.NextUp)
			}
			if len(stats.Sample) != len(tt.expected.Sample) {
				t.Errorf("Sample length = %d, want %d", len(stats.Sample), len(tt.expected.Sample))
			}
		})
	}
}

func TestFormatPlain(t *testing.T) {
	report := &Report{
		Phase:        "development",
		ActiveTask:   "1.1",
		ActiveBranch: "task/1-1",
		Locked:       true,
		CompletedIDs: []string{"1.1", "1.2"},
		Checklist: TaskStats{
			Total:     10,
			Completed: 5,
			Pending:   5,
			NextUp:    "Task 6",
			Sample:    []string{"Task 6", "Task 7", "Task 8"},
		},
		Percentage:       50.0,
		StateDiffSummary: "No changes",
	}

	output := FormatPlain(report)

	if !strings.Contains(output, "development") {
		t.Error("FormatPlain should include phase")
	}
	if !strings.Contains(output, "1.1") {
		t.Error("FormatPlain should include active task")
	}
	if !strings.Contains(output, "task/1-1") {
		t.Error("FormatPlain should include active branch")
	}
	if !strings.Contains(output, "50.0%") {
		t.Error("FormatPlain should include percentage")
	}
	if !strings.Contains(output, "5/10") {
		t.Error("FormatPlain should include checklist stats")
	}
}

func TestFormatMarkdown(t *testing.T) {
	report := &Report{
		Phase:        "development",
		ActiveTask:   "1.1",
		ActiveBranch: "task/1-1",
		Locked:       true,
		CompletedIDs: []string{"1.1", "1.2"},
		Checklist: TaskStats{
			Total:     10,
			Completed: 5,
			Pending:   5,
			NextUp:    "Task 6",
			Sample:    []string{"Task 6", "Task 7", "Task 8"},
		},
		Percentage:       50.0,
		StateDiffSummary: "No changes",
	}

	output := FormatMarkdown(report)

	if !strings.Contains(output, "- Phase: development") {
		t.Error("FormatMarkdown should include phase")
	}
	if !strings.Contains(output, "- Active Task: 1.1") {
		t.Error("FormatMarkdown should include active task")
	}
	if !strings.Contains(output, "- Checklist: 5/10") {
		t.Error("FormatMarkdown should include checklist stats")
	}
}

func TestCompute(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".plan")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")
	os.MkdirAll(historyDir, 0755)

	// Create active_state.json
	state := statehistory.ActiveState{
		Phase:        "development",
		ActiveTask:   "1.1",
		ActiveBranch: "task/1-1",
		Locked:       true,
		Completed:    []string{"1.1", "1.2"},
	}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	// Create TASKS.md
	tasksContent := `# Tasks
- [x] Task 1
- [x] Task 2
- [ ] Task 3
- [ ] Task 4`
	tasksPath := filepath.Join(planDir, "TASKS.md")
	os.WriteFile(tasksPath, []byte(tasksContent), 0644)

	// Compute report
	report, err := Compute(tmpDir)
	if err != nil {
		t.Fatalf("Compute failed: %v", err)
	}

	if report.Phase != "development" {
		t.Errorf("Phase = %q, want %q", report.Phase, "development")
	}
	if report.ActiveTask != "1.1" {
		t.Errorf("ActiveTask = %q, want %q", report.ActiveTask, "1.1")
	}
	if report.Checklist.Total != 4 {
		t.Errorf("Checklist.Total = %d, want %d", report.Checklist.Total, 4)
	}
	if report.Checklist.Completed != 2 {
		t.Errorf("Checklist.Completed = %d, want %d", report.Checklist.Completed, 2)
	}
	if report.Percentage != 50.0 {
		t.Errorf("Percentage = %f, want %f", report.Percentage, 50.0)
	}
}

func TestCompute_MissingFiles(t *testing.T) {
	tmpDir := t.TempDir()

	_, err := Compute(tmpDir)
	if err == nil {
		t.Error("Compute should fail when files are missing")
	}
}

func TestCompute_EmptyTasks(t *testing.T) {
	tmpDir := t.TempDir()
	planDir := filepath.Join(tmpDir, ".plan")
	os.MkdirAll(planDir, 0755)
	historyDir := filepath.Join(planDir, "history")
	os.MkdirAll(historyDir, 0755)

	// Create active_state.json
	state := statehistory.ActiveState{
		Phase:     "development",
		Completed: []string{},
	}
	stateData, _ := json.Marshal(state)
	statePath := filepath.Join(planDir, "active_state.json")
	os.WriteFile(statePath, stateData, 0644)

	// Create empty TASKS.md
	tasksPath := filepath.Join(planDir, "TASKS.md")
	os.WriteFile(tasksPath, []byte(""), 0644)

	// Compute report
	report, err := Compute(tmpDir)
	if err != nil {
		t.Fatalf("Compute failed: %v", err)
	}

	if report.Checklist.Total != 0 {
		t.Errorf("Checklist.Total = %d, want %d", report.Checklist.Total, 0)
	}
	if report.Percentage != 0.0 {
		t.Errorf("Percentage = %f, want %f", report.Percentage, 0.0)
	}
}

func TestFallback(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		placeholder string
		expected    string
	}{
		{"non-empty", "value", "placeholder", "value"},
		{"empty", "", "placeholder", "placeholder"},
		{"whitespace", "   ", "placeholder", "placeholder"},
		{"tab", "\t", "placeholder", "placeholder"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fallback(tt.value, tt.placeholder)
			if result != tt.expected {
				t.Errorf("fallback(%q, %q) = %q, want %q", tt.value, tt.placeholder, tt.expected, result)
			}
		})
	}
}

