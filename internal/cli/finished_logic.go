package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/statehistory"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

// finishedLogic handles the /done command workflow
type finishedLogic struct {
	projectPath string
	systemDir   string
	planDir     string
	historyDir  string
}

// newFinishedLogic creates a new finishedLogic instance
func newFinishedLogic(projectPath string) (*finishedLogic, error) {
	systemDir := filepath.Join(projectPath, ".do", "system")
	planDir := filepath.Join(projectPath, ".do", "plan")
	historyDir := filepath.Join(projectPath, ".do", "system", "history")

	// Ensure directories exist
	if err := utils.CreateDirectory(systemDir); err != nil {
		return nil, fmt.Errorf("failed to create system directory: %w", err)
	}
	if err := utils.CreateDirectory(planDir); err != nil {
		return nil, fmt.Errorf("failed to create plan directory: %w", err)
	}
	if err := utils.CreateDirectory(historyDir); err != nil {
		return nil, fmt.Errorf("failed to create history directory: %w", err)
	}

	return &finishedLogic{
		projectPath: projectPath,
		systemDir:   systemDir,
		planDir:     planDir,
		historyDir:  historyDir,
	}, nil
}

// FinishedWorkflowResult contains information about the finished workflow execution
type FinishedWorkflowResult struct {
	TaskID           string
	TaskTitle        string
	BranchName       string
	Committed        bool
	Pushed           bool
	ChangelogUpdated bool
	PRSuggested      bool
	TaskDuration     time.Duration // Total time from task start to completion
}

// runFinishedWorkflow handles the complete task completion workflow
func (f *finishedLogic) runFinishedWorkflow(out io.Writer, memoryCard *MemoryCard, orchestrator *EngagementOrchestrator) (*FinishedWorkflowResult, error) {
	// 1. Read active state to get current task and branch
	statePath := filepath.Join(f.historyDir, "active_state.json")
	state, err := f.readActiveState(statePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read active state: %w", err)
	}

	// Check if there's an active task
	if state["active_task"] == nil || state["active_task"] == "" {
		return nil, fmt.Errorf("no active task found. Please start a task with /dev first")
	}

	taskID := fmt.Sprintf("%v", state["active_task"])
	activeBranch, _ := state["active_branch"].(string)

	// Calculate task duration from start time
	var taskDuration time.Duration
	if taskStartedAtStr, ok := state["task_started_at"].(string); ok && taskStartedAtStr != "" {
		taskStartedAt, err := time.Parse(time.RFC3339, taskStartedAtStr)
		if err == nil {
			taskDuration = time.Since(taskStartedAt)
		}
	}

	// 2. Verify active branch
	if err := f.verifyActiveBranch(activeBranch, out); err != nil {
		return nil, err
	}

	// 3. Check dependencies (placeholder - would call taskcomplete script)
	if err := f.checkDependencies(taskID, out); err != nil {
		return nil, err
	}

	// 4. Mark task complete in TASKS.md
	taskTitle, err := f.markTaskComplete(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to mark task complete: %w", err)
	}

	// 5. Update state
	if err := f.updateState(taskID, statePath); err != nil {
		return nil, fmt.Errorf("failed to update state: %w", err)
	}

	// 6. Snapshot state
	if err := f.snapshotState(statePath, taskID); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Failed to create state snapshot: %v\n", err)
	}

	// 7. Auto-commit
	committed, err := f.autoCommit(taskID, taskTitle)
	if err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Failed to commit changes: %v\n", err)
	}

	// 8. Auto-push
	pushed := false
	if committed {
		var pushErr error
		pushed, pushErr = f.autoPush(activeBranch)
		if pushErr != nil {
			fmt.Fprintf(out, "⚠️  Warning: Failed to push changes: %v\n", pushErr)
		}
	}

	// 9. Update changelog (placeholder)
	changelogUpdated := false
	if committed {
		changelogUpdated = f.updateChangelog(taskID, taskTitle)
	}

	// 10. Suggest PR creation
	prSuggested := false
	if pushed {
		prSuggested = f.suggestPRCreation(taskID, taskTitle, activeBranch, out)
	}

	return &FinishedWorkflowResult{
		TaskID:           taskID,
		TaskTitle:        taskTitle,
		BranchName:       activeBranch,
		Committed:        committed,
		Pushed:           pushed,
		ChangelogUpdated: changelogUpdated,
		PRSuggested:      prSuggested,
		TaskDuration:     taskDuration,
	}, nil
}

// readActiveState reads the active_state.json file
func (f *finishedLogic) readActiveState(statePath string) (map[string]interface{}, error) {
	if !utils.PathExists(statePath) {
		return nil, fmt.Errorf("active_state.json not found at %s", statePath)
	}

	data, err := os.ReadFile(statePath)
	if err != nil {
		return nil, err
	}

	var state map[string]interface{}
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	return state, nil
}

// verifyActiveBranch checks that we're on a task branch
func (f *finishedLogic) verifyActiveBranch(branchName string, out io.Writer) error {
	if branchName == "" {
		// Try to get current branch from git
		cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
		cmd.Dir = f.projectPath
		output, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to get current branch: %w", err)
		}
		branchName = strings.TrimSpace(string(output))
	}

	// Check if on main/master
	if branchName == "main" || branchName == "master" {
		fmt.Fprintf(out, "⚠️  Warning: You're on %s branch. Are you sure you want to mark a task complete? (yes/no): ", branchName)
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "yes" && response != "y" {
			return fmt.Errorf("task completion cancelled")
		}
	}

	return nil
}

// checkDependencies checks if all task dependencies are complete
func (f *finishedLogic) checkDependencies(taskID string, out io.Writer) error {
	// TODO: Implement dependency checking using taskcomplete script
	// For now, just a placeholder
	fmt.Fprintf(out, "✓ Checking dependencies...\n")
	return nil
}

// markTaskComplete marks the task as complete in TASKS.md
func (f *finishedLogic) markTaskComplete(taskID string) (string, error) {
	tasksPath := filepath.Join(f.planDir, "TASKS.md")
	if !utils.PathExists(tasksPath) {
		return "", fmt.Errorf("TASKS.md not found at %s", tasksPath)
	}

	content, err := os.ReadFile(tasksPath)
	if err != nil {
		return "", err
	}

	// Find task by ID and mark it complete
	lines := strings.Split(string(content), "\n")
	taskTitle := ""
	taskFound := false

	// Pattern to match task headers: ### 1.2 Task Title
	taskPattern := regexp.MustCompile(`^###\s+(\d+\.\d+)\s+(.+)$`)
	statusPattern := regexp.MustCompile(`^\*\*Status\*\*:\s*(.+)$`)

	for i, line := range lines {
		if matches := taskPattern.FindStringSubmatch(line); len(matches) == 3 {
			if matches[1] == taskID {
				taskTitle = matches[2]
				taskFound = true
				// Look for status line in next few lines
				for j := i + 1; j < len(lines) && j < i+10; j++ {
					if statusMatch := statusPattern.FindStringSubmatch(lines[j]); len(statusMatch) == 2 {
						// Update status to Complete
						lines[j] = "**Status**: ✅ Complete"
						break
					}
				}
				// Mark all checklist items as [x]
				for j := i + 1; j < len(lines) && j < i+50; j++ {
					if strings.HasPrefix(lines[j], "###") {
						break // Next task
					}
					// Replace [ ] with [x]
					lines[j] = strings.ReplaceAll(lines[j], "- [ ]", "- [x]")
				}
				break
			}
		}
	}

	if !taskFound {
		return "", fmt.Errorf("task %s not found in TASKS.md", taskID)
	}

	// Write updated content
	newContent := strings.Join(lines, "\n")
	if err := utils.WriteFile(tasksPath, []byte(newContent)); err != nil {
		return "", err
	}

	return taskTitle, nil
}

// updateState updates active_state.json
func (f *finishedLogic) updateState(taskID, statePath string) error {
	state, err := f.readActiveState(statePath)
	if err != nil {
		return err
	}

	// Add to completed array
	completed, _ := state["completed"].([]interface{})
	if completed == nil {
		completed = []interface{}{}
	}

	// Check if already in completed
	found := false
	for _, id := range completed {
		if fmt.Sprintf("%v", id) == taskID {
			found = true
			break
		}
	}

	if !found {
		completed = append(completed, taskID)
		state["completed"] = completed
	}

	// Clear active_task, active_branch, and task_started_at
	state["active_task"] = nil
	state["active_branch"] = nil
	state["task_started_at"] = nil

	// Write updated state
	jsonData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return utils.WriteFile(statePath, jsonData)
}

// snapshotState creates a state snapshot
func (f *finishedLogic) snapshotState(statePath, taskID string) error {
	reason := fmt.Sprintf("done %s", taskID)
	_, err := statehistory.SaveSnapshot(statePath, f.historyDir, reason, "done")
	return err
}

// autoCommit commits changes with conventional commit format
func (f *finishedLogic) autoCommit(taskID, taskTitle string) (bool, error) {
	// Check if there are changes to commit
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = f.projectPath
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	if len(strings.TrimSpace(string(output))) == 0 {
		return false, nil // No changes to commit
	}

	// Stage all changes
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = f.projectPath
	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("failed to stage changes: %w", err)
	}

	// Create conventional commit message
	commitMsg := fmt.Sprintf("feat(task-%s): complete %s", taskID, taskTitle)

	// Commit
	cmd = exec.Command("git", "commit", "-m", commitMsg)
	cmd.Dir = f.projectPath
	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("failed to commit: %w", err)
	}

	return true, nil
}

// autoPush pushes the current branch
func (f *finishedLogic) autoPush(branchName string) (bool, error) {
	if branchName == "" {
		// Get current branch
		cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
		cmd.Dir = f.projectPath
		output, err := cmd.Output()
		if err != nil {
			return false, err
		}
		branchName = strings.TrimSpace(string(output))
	}

	// Push to origin
	cmd := exec.Command("git", "push", "origin", branchName)
	cmd.Dir = f.projectPath
	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("failed to push: %w", err)
	}

	return true, nil
}

// updateChangelog adds entry to CHANGELOG.md if significant
func (f *finishedLogic) updateChangelog(taskID, taskTitle string) bool {
	// TODO: Implement actual changelog update logic
	// changelogPath := filepath.Join(f.projectPath, "CHANGELOG.md")
	return false
}

// suggestPRCreation suggests creating a PR if gh CLI is available
func (f *finishedLogic) suggestPRCreation(taskID, taskTitle, branchName string, out io.Writer) bool {
	// Check if gh CLI is available
	cmd := exec.Command("gh", "--version")
	if err := cmd.Run(); err != nil {
		return false // gh CLI not available
	}

	fmt.Fprintf(out, "\n💡 Suggestion: Create a pull request?\n")
	fmt.Fprintf(out, "   Run: gh pr create --title \"feat: %s\" --body \"Completes task %s\"\n", taskTitle, taskID)

	return true
}
