package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	timetracker "github.com/DoPlan-dev/CLI/internal/time"
)

func init() {
	rootCmd.AddCommand(newFinishedCommand())
}

func newFinishedCommand() *cobra.Command {
	var projectPath string

	cmd := &cobra.Command{
		Use:   "/done",
		Short: "Mark current task as complete and auto-commit/push",
		Long: `The /done command marks the current active task as complete, updates TASKS.md,
creates a state snapshot, auto-commits changes with conventional commit format, and pushes to remote.

It performs:
- Verifies active branch (warns if on main/master)
- Checks task dependencies
- Marks task complete in TASKS.md
- Updates active_state.json
- Creates state snapshot
- Auto-commits with conventional format
- Auto-pushes to remote branch
- Updates CHANGELOG.md (if significant)
- Suggests PR creation (if gh CLI available)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := resolveProjectPath(projectPath)
			if err != nil {
				return err
			}

			// Fast path: Only initialize engagement system for existing projects
			orchestrator, err := getOrCreateEngagementOrchestrator(absPath)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Could not initialize engagement system: %v\n", err)
			}

			// Load memory card (cached for performance)
			var memoryCard *MemoryCard
			if !shouldUseFastPath(absPath) {
				memoryCard, err = loadMemoryCardCached()
				if err != nil {
					memoryCard = nil
				}
			}

			// Initialize time tracker
			tracker, err := timetracker.New(absPath)
			if err != nil {
				return err
			}

			// Initialize finished logic
			finished, err := newFinishedLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /done logic: %w", err)
			}

			// Process with engagement system before execution (only for existing projects)
			context := map[string]interface{}{
				"command": "/done",
				"project": absPath,
				"phase":   "completion",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/done", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Start tracking
			metadata := map[string]string{}
			if memoryCard != nil {
				metadata["work_style"] = memoryCard.WorkStyle
			}
			tracker.Start(absPath, "completion", "/done", args, metadata)

			// Run finished workflow
			startTime := time.Now()
			result, err := finished.runFinishedWorkflow(cmd.OutOrStdout(), memoryCard, orchestrator)
			duration := time.Since(startTime).Seconds()

			if err != nil {
				if stopErr := tracker.Stop(false, err); stopErr != nil {
					return fmt.Errorf("task completion failed: %w (also failed to stop tracker: %v)", err, stopErr)
				}
				// Process engagement after failure
				if orchestrator != nil {
					context["success"] = false
					if err := orchestrator.ProcessCommandWithEngagement("/done", context, cmd.OutOrStdout()); err != nil {
						fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
					}
					orchestrator.TrackInteraction("/done", "", "Task completion failed", duration, "negative")
				}
				return fmt.Errorf("task completion failed: %w", err)
			}

			// Process engagement after successful execution
			if orchestrator != nil {
				context["success"] = true
				context["task_completed"] = true
				context["task_id"] = result.TaskID
				context["task_title"] = result.TaskTitle
				context["branch_name"] = result.BranchName
				context["committed"] = result.Committed
				context["pushed"] = result.Pushed
				context["completion_duration"] = duration
				context["task_duration_seconds"] = result.TaskDuration.Seconds()
				context["task_duration_minutes"] = result.TaskDuration.Minutes()

				if err := orchestrator.ProcessCommandWithEngagement("/done", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}

				// Track interaction
				response := fmt.Sprintf("Task %s (%s) marked complete", result.TaskID, result.TaskTitle)
				orchestrator.TrackInteraction("/done", result.TaskID, response, duration, "positive")

				// Update memory card
				if orchestrator.memoryCard != nil {
					orchestrator.memoryCard.CommandUsage["/done"]++
					orchestrator.memoryCard.CurrentPhase = "completed"
					orchestrator.memoryCard.LastCommand = "/done"
					orchestrator.memoryCard.LastCommandTime = time.Now()

					// Save memory card
					if err := SaveMemoryCard(orchestrator.memoryCard); err != nil {
						fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Failed to save memory card: %v\n", err)
					}
				}
			}

			// Stop tracking
			if stopErr := tracker.Stop(true, nil); stopErr != nil {
				return stopErr
			}

			// Display success message
			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintf(cmd.OutOrStdout(), "✅ Task %s marked complete!\n", result.TaskID)

			// Display task duration if available
			if result.TaskDuration > 0 {
				hours := int(result.TaskDuration.Hours())
				minutes := int(result.TaskDuration.Minutes()) % 60
				if hours > 0 {
					fmt.Fprintf(cmd.OutOrStdout(), "   ⏱️  Task duration: %dh %dm\n", hours, minutes)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "   ⏱️  Task duration: %dm\n", minutes)
				}
			}

			if result.Committed {
				fmt.Fprintf(cmd.OutOrStdout(), "   ✓ Changes committed\n")
			}
			if result.Pushed {
				fmt.Fprintf(cmd.OutOrStdout(), "   ✓ Changes pushed to %s\n", result.BranchName)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintln(cmd.OutOrStdout(), "💡 Next steps:")
			fmt.Fprintln(cmd.OutOrStdout(), "   • Type /dev to start the next task")
			fmt.Fprintln(cmd.OutOrStdout(), "   • Type /status to see overall progress")

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")

	return cmd
}
