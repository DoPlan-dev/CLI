package progress

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/statehistory"
)

// TaskStats captures totals from TASKS.md checklists.
type TaskStats struct {
	Total     int      `json:"total"`
	Completed int      `json:"completed"`
	Pending   int      `json:"pending"`
	NextUp    string   `json:"next_up"`
	Sample    []string `json:"sample_pending"`
}

// Report is the aggregated project progress snapshot.
type Report struct {
	Phase            string                  `json:"phase"`
	ActiveTask       string                  `json:"active_task"`
	ActiveBranch     string                  `json:"active_branch"`
	Locked           bool                    `json:"locked"`
	CompletedIDs     []string                `json:"completed_ids"`
	Checklist        TaskStats               `json:"checklist"`
	Percentage       float64                 `json:"percent_complete"`
	StateDiffSummary string                  `json:"state_diff"`
	StateDiff        *statehistory.StateDiff `json:"state_diff_struct,omitempty"`
}

// Compute assembles a progress report for the supplied project root.
func Compute(root string) (*Report, error) {
	planDir := filepath.Join(root, ".plan")
	statePath := filepath.Join(planDir, "active_state.json")
	tasksPath := filepath.Join(planDir, "TASKS.md")
	historyDir := filepath.Join(planDir, "history")

	state, err := loadActiveState(statePath)
	if err != nil {
		return nil, err
	}
	stats, err := parseTasks(tasksPath)
	if err != nil {
		return nil, err
	}
	percent := 0.0
	if stats.Total > 0 {
		percent = float64(stats.Completed) / float64(stats.Total) * 100
	}
	diffSummary, diffStruct := loadStateDiff(historyDir)

	return &Report{
		Phase:            state.Phase,
		ActiveTask:       state.ActiveTask,
		ActiveBranch:     state.ActiveBranch,
		Locked:           state.Locked,
		CompletedIDs:     append([]string(nil), state.Completed...),
		Checklist:        stats,
		Percentage:       percent,
		StateDiffSummary: diffSummary,
		StateDiff:        diffStruct,
	}, nil
}

// FormatPlain renders a console-friendly snapshot.
func FormatPlain(r *Report) string {
	var b strings.Builder
	b.WriteString("📊 Project Progress\n")
	b.WriteString("===================\n")
	b.WriteString(fmt.Sprintf("Phase        : %s\n", fallback(r.Phase, "(unknown)")))
	b.WriteString(fmt.Sprintf("Active Task  : %s\n", fallback(r.ActiveTask, "(none)")))
	b.WriteString(fmt.Sprintf("Active Branch: %s\n", fallback(r.ActiveBranch, "(none)")))
	b.WriteString(fmt.Sprintf("Locked       : %t\n", r.Locked))
	b.WriteString(fmt.Sprintf("Checklist    : %d/%d complete (%.1f%%)\n", r.Checklist.Completed, r.Checklist.Total, r.Percentage))
	if r.Checklist.NextUp != "" {
		b.WriteString(fmt.Sprintf("Next Up      : %s\n", r.Checklist.NextUp))
	}
	if len(r.Checklist.Sample) > 0 {
		b.WriteString("Upcoming Tasks:\n")
		for _, task := range r.Checklist.Sample {
			b.WriteString("  - " + task + "\n")
		}
	}
	b.WriteString("\nState Delta (latest snapshots):\n")
	b.WriteString(r.StateDiffSummary)
	if !strings.HasSuffix(r.StateDiffSummary, "\n") {
		b.WriteString("\n")
	}
	return b.String()
}

// FormatMarkdown renders a compact markdown block suitable for reports.
func FormatMarkdown(r *Report) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("- Phase: %s\n", fallback(r.Phase, "(unknown)")))
	b.WriteString(fmt.Sprintf("- Active Task: %s\n", fallback(r.ActiveTask, "(none)")))
	b.WriteString(fmt.Sprintf("- Active Branch: %s\n", fallback(r.ActiveBranch, "(none)")))
	b.WriteString(fmt.Sprintf("- Locked: %t\n", r.Locked))
	b.WriteString(fmt.Sprintf("- Checklist: %d/%d complete (%.1f%%)\n", r.Checklist.Completed, r.Checklist.Total, r.Percentage))
	if r.Checklist.NextUp != "" {
		b.WriteString(fmt.Sprintf("- Next Up: %s\n", r.Checklist.NextUp))
	}
	if len(r.Checklist.Sample) > 0 {
		b.WriteString("- Upcoming Tasks:\n")
		for _, task := range r.Checklist.Sample {
			b.WriteString("  - " + task + "\n")
		}
	}
	b.WriteString("- State Delta:\n")
	indented := indentLines(r.StateDiffSummary, "  ")
	b.WriteString(indented)
	if !strings.HasSuffix(indented, "\n") {
		b.WriteString("\n")
	}
	return b.String()
}

func loadActiveState(path string) (*statehistory.ActiveState, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read active_state.json: %w", err)
	}
	var state statehistory.ActiveState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse active_state.json: %w", err)
	}
	return &state, nil
}

func parseTasks(path string) (TaskStats, error) {
	file, err := os.Open(path)
	if err != nil {
		return TaskStats{}, fmt.Errorf("read TASKS.md: %w", err)
	}
	defer file.Close()

	stats := TaskStats{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "-") {
			continue
		}
		lower := strings.ToLower(line)
		switch {
		case strings.HasPrefix(lower, "- [x]"):
			stats.Total++
			stats.Completed++
		case strings.HasPrefix(lower, "- [ ]"):
			stats.Total++
			stats.Pending++
			taskText := strings.TrimSpace(line[5:])
			if stats.NextUp == "" {
				stats.NextUp = taskText
			}
			if len(stats.Sample) < 3 {
				stats.Sample = append(stats.Sample, taskText)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return TaskStats{}, fmt.Errorf("scan TASKS.md: %w", err)
	}
	return stats, nil
}

func loadStateDiff(historyDir string) (string, *statehistory.StateDiff) {
	diff, err := statehistory.LatestDiff(historyDir)
	if err != nil {
		if err == statehistory.ErrInsufficientSnapshots {
			return "Not enough state snapshots have been recorded yet.", nil
		}
		return fmt.Sprintf("Unable to read state history (%v).", err), nil
	}
	summary := statehistory.FormatDiff(*diff)
	return summary, diff
}

func indentLines(input, prefix string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func fallback(value, placeholder string) string {
	if strings.TrimSpace(value) == "" {
		return placeholder
	}
	return value
}
