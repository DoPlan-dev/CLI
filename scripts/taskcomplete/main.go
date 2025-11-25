package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/statehistory"
)

type TaskInfo struct {
	ID          string
	Description string
	Dependencies []string
	Status      string
	LineStart   int
	LineEnd     int
}

func main() {
	taskID := flag.String("task", "", "Task ID to mark complete (e.g., 5.3)")
	projectPath := flag.String("project", ".", "Project root path")
	checkOnly := flag.Bool("check", false, "Only check dependencies, don't mark complete")
	flag.Parse()

	if *taskID == "" {
		fmt.Fprintf(os.Stderr, "Error: --task is required\n")
		os.Exit(1)
	}

	absPath, err := filepath.Abs(*projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to resolve project path: %v\n", err)
		os.Exit(1)
	}

	tasksPath := filepath.Join(absPath, ".plan", "TASKS.md")
	statePath := filepath.Join(absPath, ".plan", "active_state.json")

	// Load active state to get completed tasks
	state, err := loadActiveState(statePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to load active_state.json: %v\n", err)
		os.Exit(1)
	}

	// Parse TASKS.md to find the task
	task, err := findTask(tasksPath, *taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if task == nil {
		fmt.Fprintf(os.Stderr, "Error: task %s not found in TASKS.md\n", *taskID)
		os.Exit(1)
	}

	// Check dependencies
	if len(task.Dependencies) > 0 {
		missingDeps := checkDependencies(task.Dependencies, state.Completed)
		if len(missingDeps) > 0 {
			fmt.Fprintf(os.Stderr, "Error: Task %s cannot be completed. Missing dependencies: %s\n", *taskID, strings.Join(missingDeps, ", "))
			fmt.Fprintf(os.Stderr, "Please complete the following tasks first: %s\n", strings.Join(missingDeps, ", "))
			os.Exit(1)
		}
	}

	if *checkOnly {
		fmt.Printf("Task %s dependencies satisfied. Ready to complete.\n", *taskID)
		return
	}

	// Mark task as complete in TASKS.md
	if err := markTaskComplete(tasksPath, task); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to mark task complete: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %s marked as complete in TASKS.md\n", *taskID)
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

func findTask(tasksPath, taskID string) (*TaskInfo, error) {
	file, err := os.Open(tasksPath)
	if err != nil {
		return nil, fmt.Errorf("read TASKS.md: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	lineNum := 0
	var currentTask *TaskInfo
	inTask := false

	reTaskHeader := regexp.MustCompile(`^###\s+(\d+\.\d+)\s+(.+)$`)
	reDependencies := regexp.MustCompile(`^\*\*Dependencies\*\*:\s*(.+)$`)
	reStatus := regexp.MustCompile(`^\*\*Status\*\*:\s*(.+)$`)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		lineNum++

		trimmed := strings.TrimSpace(line)

		// Check for task header
		if matches := reTaskHeader.FindStringSubmatch(trimmed); len(matches) == 3 {
			// If we were tracking a task and found a new one, save the previous
			if currentTask != nil {
				currentTask.LineEnd = lineNum - 1
			}

			foundID := matches[1]
			if foundID == taskID {
				currentTask = &TaskInfo{
					ID:        foundID,
					LineStart: lineNum,
					Status:    "⏳ Pending", // default
				}
				inTask = true
			} else {
				inTask = false
				if currentTask != nil && currentTask.ID == taskID {
					// We found the end of our task
					currentTask.LineEnd = lineNum - 1
					break
				}
			}
			continue
		}

		if !inTask || currentTask == nil {
			continue
		}

		// Check for description
		if strings.HasPrefix(trimmed, "**Description**") {
			parts := strings.SplitN(trimmed, ":", 2)
			if len(parts) == 2 {
				currentTask.Description = strings.TrimSpace(parts[1])
			}
		}

		// Check for dependencies
		if matches := reDependencies.FindStringSubmatch(trimmed); len(matches) == 2 {
			depsStr := strings.TrimSpace(matches[1])
			if depsStr != "None" && depsStr != "" {
				// Parse dependencies (can be comma-separated or space-separated)
				deps := parseDependencies(depsStr)
				currentTask.Dependencies = deps
			}
		}

		// Check for status
		if matches := reStatus.FindStringSubmatch(trimmed); len(matches) == 2 {
			currentTask.Status = strings.TrimSpace(matches[1])
		}

		// Check if we've reached the next task (### header) or end of file
		if strings.HasPrefix(trimmed, "###") && currentTask.LineEnd == 0 {
			currentTask.LineEnd = lineNum - 1
		}
	}

	if currentTask != nil && currentTask.ID == taskID {
		if currentTask.LineEnd == 0 {
			currentTask.LineEnd = lineNum
		}
		return currentTask, nil
	}

	return nil, nil
}

func parseDependencies(depsStr string) []string {
	// Handle various formats: "1.1", "1.1, 1.2", "Phase 1", "None", etc.
	depsStr = strings.TrimSpace(depsStr)
	if depsStr == "None" || depsStr == "" {
		return []string{}
	}

	// Split by comma or space
	parts := strings.FieldsFunc(depsStr, func(r rune) bool {
		return r == ',' || r == ' '
	})

	var deps []string
	reTaskID := regexp.MustCompile(`^(\d+\.\d+)$`)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if reTaskID.MatchString(part) {
			deps = append(deps, part)
		}
	}

	return deps
}

func checkDependencies(dependencies, completed []string) []string {
	completedSet := make(map[string]bool)
	for _, id := range completed {
		completedSet[id] = true
	}

	var missing []string
	for _, dep := range dependencies {
		if !completedSet[dep] {
			missing = append(missing, dep)
		}
	}
	return missing
}

func markTaskComplete(tasksPath string, task *TaskInfo) error {
	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return fmt.Errorf("read TASKS.md: %w", err)
	}

	lines := strings.Split(string(data), "\n")

	// Update status line
	for i := task.LineStart - 1; i < len(lines) && i < task.LineEnd; i++ {
		line := lines[i]
		if strings.Contains(line, "**Status**") {
			lines[i] = regexp.MustCompile(`\*\*Status\*\*:\s*.+`).ReplaceAllString(line, "**Status**: ✅ Complete")
			break
		}
	}

	// Mark all checklist items as complete
	for i := task.LineStart - 1; i < len(lines) && i < task.LineEnd; i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "- [ ]") {
			lines[i] = strings.Replace(lines[i], "- [ ]", "- [x]", 1)
		}
	}

	output := strings.Join(lines, "\n")
	return os.WriteFile(tasksPath, []byte(output), 0644)
}

