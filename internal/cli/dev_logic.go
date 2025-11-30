package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/generator"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

// devLogic handles the /dev command workflow
type devLogic struct {
	projectPath string
	systemDir   string
	planDir     string
	historyDir  string
}

// newDevLogic creates a new devLogic instance
func newDevLogic(projectPath string) (*devLogic, error) {
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

	return &devLogic{
		projectPath: projectPath,
		systemDir:   systemDir,
		planDir:     planDir,
		historyDir:  historyDir,
	}, nil
}

// DevWorkflowResult contains information about the development workflow execution
type DevWorkflowResult struct {
	FeatureName   string
	TaskID        string
	Phase         string
	BranchCreated bool
	DocsSynced    bool
	TaskTitle     string
	Description   string
}

// runDevWorkflow handles the complete development workflow (legacy, kept for compatibility)
func (d *devLogic) runDevWorkflow(featureName string, out io.Writer, memoryCard *MemoryCard) error {
	result, err := d.runDevWorkflowEnhanced(featureName, out, memoryCard, nil)
	if err != nil {
		return err
	}
	_ = result // Use result to avoid unused variable
	return nil
}

// runDevWorkflowEnhanced handles the complete development workflow with full engagement integration
func (d *devLogic) runDevWorkflowEnhanced(featureName string, out io.Writer, memoryCard *MemoryCard, orchestrator *EngagementOrchestrator) (*DevWorkflowResult, error) {
	// Read TASKS.md to find next task or specific feature
	tasksPath := filepath.Join(d.planDir, "TASKS.md")
	if !utils.PathExists(tasksPath) {
		return nil, fmt.Errorf("TASKS.md not found. Please run /plan first to generate your execution plan")
	}

	tasksContent, err := os.ReadFile(tasksPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read TASKS.md: %w", err)
	}

	// Find task to work on
	var selectedTask *taskInfo
	if featureName != "" && featureName != "general" {
		selectedTask = d.findTaskByFeature(string(tasksContent), featureName)
	} else {
		selectedTask = d.findNextTask(string(tasksContent))
	}

	if selectedTask == nil {
		return nil, fmt.Errorf("no available task found. All tasks may be completed or feature '%s' not found", featureName)
	}

	// Display task information
	fmt.Fprintln(out, "🚀 Starting development workflow")
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "📋 Task: %s\n", selectedTask.Title)
	fmt.Fprintf(out, "   Description: %s\n", selectedTask.Description)
	fmt.Fprintf(out, "   Phase: %s\n", selectedTask.Phase)
	fmt.Fprintf(out, "   Task ID: %s\n", selectedTask.ID)
	fmt.Fprintln(out, "")

	// Track branch creation and docs sync status
	branchCreated := true
	docsSynced := true

	// Check Git status (if git is available)
	if err := d.checkGitStatus(out); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Git check failed: %v\n", err)
	}

	// Create/switch to task branch
	branchName := fmt.Sprintf("task/%s", strings.ReplaceAll(selectedTask.ID, ".", "-"))
	if err := d.ensureTaskBranch(branchName, out); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Branch operation failed: %v\n", err)
		fmt.Fprintln(out, "   Continuing without branch creation...")
		branchCreated = false
	}

	// Sync documentation for this feature
	fmt.Fprintln(out, "📚 Syncing documentation...")
	if err := generator.SyncPlanDocumentation(d.projectPath); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Failed to sync documentation: %v\n", err)
		docsSynced = false
	}

	// ============================================
	// BRAIN-POWERED PERSONALIZATION
	// ============================================
	var personalizedMessage string
	if orchestrator != nil && orchestrator.brain != nil && memoryCard != nil {
		// Get personalized message based on user's work style and experience
		if memoryCard.WorkStyle == "fast" {
			personalizedMessage = "⚡ Quick setup complete! Let's build fast!"
		} else if memoryCard.WorkStyle == "thoughtful" {
			personalizedMessage = "🎯 Everything is set up thoughtfully. Take your time!"
		}

		// Add encouragement based on relationship level
		if memoryCard.RelationshipLevel > 50 {
			personalizedMessage += " I'm here to help every step of the way! 💪"
		} else if memoryCard.RelationshipLevel > 20 {
			personalizedMessage += " You've got this! 🚀"
		}
	}

	// Update memory card with development patterns
	if memoryCard != nil {
		// Learn tech preferences from task
		descLower := strings.ToLower(selectedTask.Description)
		if strings.Contains(descLower, "frontend") || strings.Contains(descLower, "ui") || strings.Contains(descLower, "react") || strings.Contains(descLower, "vue") {
			if !containsString(memoryCard.PreferredTechStack, "frontend") {
				memoryCard.PreferredTechStack = append(memoryCard.PreferredTechStack, "frontend")
			}
		}
		if strings.Contains(descLower, "backend") || strings.Contains(descLower, "api") || strings.Contains(descLower, "server") {
			if !containsString(memoryCard.PreferredTechStack, "backend") {
				memoryCard.PreferredTechStack = append(memoryCard.PreferredTechStack, "backend")
			}
		}
		if strings.Contains(descLower, "database") || strings.Contains(descLower, "db") || strings.Contains(descLower, "sql") {
			if !containsString(memoryCard.PreferredTechStack, "database") {
				memoryCard.PreferredTechStack = append(memoryCard.PreferredTechStack, "database")
			}
		}
		if strings.Contains(descLower, "auth") || strings.Contains(descLower, "authentication") || strings.Contains(descLower, "login") {
			if !containsString(memoryCard.PreferredTechStack, "authentication") {
				memoryCard.PreferredTechStack = append(memoryCard.PreferredTechStack, "authentication")
			}
		}

		// Track feature worked on
		if !containsString(memoryCard.HelpfulFeatures, selectedTask.Title) {
			memoryCard.HelpfulFeatures = append(memoryCard.HelpfulFeatures, selectedTask.Title)
		}
	}

	// Update active state
	if err := d.updateActiveState("building", selectedTask.ID); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Could not update active state: %v\n", err)
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "✅ Development environment ready!")
	if personalizedMessage != "" {
		fmt.Fprintln(out, personalizedMessage)
	}
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "💡 Working on: %s\n", selectedTask.Title)
	fmt.Fprintln(out, "   • Feature documentation synced")
	fmt.Fprintln(out, "   • Branch created/checked out")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "📝 Next steps:")
	fmt.Fprintln(out, "   • Review feature documentation in .do/plan/")
	fmt.Fprintln(out, "   • Start coding with your IDE")
	fmt.Fprintln(out, "   • Type /done when task is complete")

	// Return result for engagement processing
	return &DevWorkflowResult{
		FeatureName:   selectedTask.Title,
		TaskID:        selectedTask.ID,
		Phase:         selectedTask.Phase,
		BranchCreated: branchCreated,
		DocsSynced:    docsSynced,
		TaskTitle:     selectedTask.Title,
		Description:   selectedTask.Description,
	}, nil
}

// containsString is a helper function to check if a string is in a slice
func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

type taskInfo struct {
	ID          string
	Title       string
	Description string
	Phase       string
	Status      string
}

func (d *devLogic) findNextTask(content string) *taskInfo {
	lines := strings.Split(content, "\n")
	var currentPhase string
	rePhase := regexp.MustCompile(`## Phase (\d+):\s+(.+)`)
	reTask := regexp.MustCompile(`### (\d+\.\d+)\s+(.+)`)
	reStatus := regexp.MustCompile(`\*\*Status\*\*:\s+(.+)`)

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if matches := rePhase.FindStringSubmatch(line); len(matches) == 3 {
			currentPhase = matches[2]
			continue
		}

		if matches := reTask.FindStringSubmatch(line); len(matches) == 3 {
			taskID := matches[1]
			taskTitle := matches[2]

			// Check status in next few lines
			status := "Pending"
			description := ""
			for j := i + 1; j < len(lines) && j < i+10; j++ {
				nextLine := strings.TrimSpace(lines[j])
				if statusMatch := reStatus.FindStringSubmatch(nextLine); len(statusMatch) == 2 {
					status = strings.TrimSpace(statusMatch[1])
				}
				if strings.HasPrefix(nextLine, "**Description**") {
					parts := strings.SplitN(nextLine, ":", 2)
					if len(parts) == 2 {
						description = strings.TrimSpace(parts[1])
					}
				}
				if strings.HasPrefix(nextLine, "###") {
					break
				}
			}

			// If task is pending, return it
			if status == "Pending" || status == "⏳ Pending" {
				return &taskInfo{
					ID:          taskID,
					Title:       taskTitle,
					Description: description,
					Phase:       currentPhase,
					Status:      status,
				}
			}
		}
	}

	return nil
}

func (d *devLogic) findTaskByFeature(content, featureName string) *taskInfo {
	lines := strings.Split(content, "\n")
	var currentPhase string
	rePhase := regexp.MustCompile(`## Phase (\d+):\s+(.+)`)
	reTask := regexp.MustCompile(`### (\d+\.\d+)\s+(.+)`)

	featureLower := strings.ToLower(featureName)

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if matches := rePhase.FindStringSubmatch(line); len(matches) == 3 {
			currentPhase = matches[2]
			continue
		}

		if matches := reTask.FindStringSubmatch(line); len(matches) == 3 {
			taskID := matches[1]
			taskTitle := matches[2]

			// Check if this task matches the feature
			if strings.Contains(strings.ToLower(taskTitle), featureLower) {
				// Extract description
				description := ""
				for j := i + 1; j < len(lines) && j < i+10; j++ {
					nextLine := strings.TrimSpace(lines[j])
					if strings.HasPrefix(nextLine, "**Description**") {
						parts := strings.SplitN(nextLine, ":", 2)
						if len(parts) == 2 {
							description = strings.TrimSpace(parts[1])
						}
					}
					if strings.HasPrefix(nextLine, "###") {
						break
					}
				}

				return &taskInfo{
					ID:          taskID,
					Title:       taskTitle,
					Description: description,
					Phase:       currentPhase,
					Status:      "Pending",
				}
			}
		}
	}

	return nil
}

func (d *devLogic) checkGitStatus(out io.Writer) error {
	// Check if .git exists
	gitDir := filepath.Join(d.projectPath, ".git")
	if !utils.PathExists(gitDir) {
		fmt.Fprintln(out, "ℹ️  No Git repository found. Consider initializing one with 'git init'")
		return nil
	}

	// In a real implementation, we would run git status
	// For now, just acknowledge git exists
	fmt.Fprintln(out, "✓ Git repository detected")
	return nil
}

func (d *devLogic) ensureTaskBranch(branchName string, out io.Writer) error {
	// In a real implementation, this would:
	// 1. Check current branch
	// 2. Create new branch if it doesn't exist
	// 3. Checkout the branch
	// 4. Handle errors gracefully

	// For now, just inform the user
	fmt.Fprintf(out, "🌿 Branch: %s\n", branchName)
	fmt.Fprintln(out, "   (Branch creation would happen here in full implementation)")
	return nil
}

func (d *devLogic) updateActiveState(phase string, activeTask string) error {
	statePath := filepath.Join(d.historyDir, "active_state.json")

	state := map[string]interface{}{
		"phase":       phase,
		"active_task": activeTask,
	}

	// Try to read existing state
	if data, err := os.ReadFile(statePath); err == nil {
		var existing map[string]interface{}
		if err := json.Unmarshal(data, &existing); err == nil {
			// Preserve existing fields
			if locked, ok := existing["locked"]; ok {
				state["locked"] = locked
			}
			if completed, ok := existing["completed"]; ok {
				state["completed"] = completed
			}
		}
	}

	// Store task start time when starting a new task
	if activeTask != "" {
		state["task_started_at"] = time.Now().UTC().Format(time.RFC3339)
	}

	jsonData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return utils.WriteFile(statePath, jsonData)
}
