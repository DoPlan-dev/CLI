package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	timetracker "github.com/DoPlan-dev/CLI/internal/time"
)

func init() {
	rootCmd.AddCommand(newPlanCommand())
	rootCmd.AddCommand(newDevCommand())
}

func newPlanCommand() *cobra.Command {
	var projectPath string

	cmd := &cobra.Command{
		Use:   "/plan",
		Short: "Generate execution plan from IDEA.md and BRAINSTORM.md",
		Long: `The /plan command reads your project idea and discovery meeting results,
then generates a structured execution plan with phases and tasks.

It creates:
- TASKS.md with organized phases and tasks
- Phase directories (01-Foundation, 02-Core, etc.)
- Feature folders with templates for each task
- Documentation sync to docs/ directory`,
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

			// Initialize plan logic
			plan, err := newPlanLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /plan logic: %w", err)
			}

			// Start tracking
			metadata := map[string]string{}
			if memoryCard != nil {
				metadata["work_style"] = memoryCard.WorkStyle
				metadata["planning_style"] = map[string]string{
					"fast":       "quick",
					"thoughtful": "detailed",
				}[memoryCard.WorkStyle]
			}
			tracker.Start(absPath, "planning", "/plan", args, metadata)

			// Process with engagement system before execution (only for existing projects)
			context := map[string]interface{}{
				"command": "/plan",
				"project": absPath,
				"phase":   "planning",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/plan", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run plan generation
			if err := plan.runPlanGeneration(cmd.OutOrStdout(), memoryCard); err != nil {
				if stopErr := tracker.Stop(false, err); stopErr != nil {
					return fmt.Errorf("planning failed: %w (also failed to stop tracker: %v)", err, stopErr)
				}
				return fmt.Errorf("planning failed: %w", err)
			}

			// Process engagement after successful execution
			if orchestrator != nil {
				context["success"] = true
				context["planning_completed"] = true
				if err := orchestrator.ProcessCommandWithEngagement("/plan", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/plan", "", "Plan generation completed", 0, "positive")
			}

			// Stop tracking
			if stopErr := tracker.Stop(true, nil); stopErr != nil {
				return stopErr
			}

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			return nil
		},
	}
	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")

	return cmd
}

func newDevCommand() *cobra.Command {
	var projectPath string
	var feature string

	cmd := &cobra.Command{
		Use:   "/dev",
		Short: "Start development workflow for a feature",
		Long: `The /dev command initiates development for a specific feature or the next available task.
It prepares your development environment, creates/checks out the appropriate branch,
and syncs relevant documentation.

Usage:
  /dev                    # Start next available task
  /dev --feature "auth"   # Start specific feature by name
  /dev --feature "1.2"    # Start specific task by ID`,
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

			// Initialize dev logic
			dev, err := newDevLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /dev logic: %w", err)
			}

			if feature == "" {
				feature = "general"
			}

			// Start tracking (will be updated with task_id after dev workflow starts)
			meta := map[string]string{"feature": feature}
			if memoryCard != nil {
				meta["work_style"] = memoryCard.WorkStyle
				meta["experience_level"] = memoryCard.ExperienceLevel
			}
			tracker.Start(absPath, "development", "/dev", append([]string{feature}, args...), meta)

			// ============================================
			// BRAIN-POWERED PERSONALIZATION
			// ============================================
			var personalizedGreeting string
			if orchestrator != nil && orchestrator.brain != nil {
				// Get personalized greeting from brain
				personalizedGreeting = orchestrator.GetPersonalizedGreeting("development")
			}

			// Display personalized greeting if available
			if personalizedGreeting != "" {
				fmt.Fprintln(cmd.OutOrStdout(), personalizedGreeting)
				fmt.Fprintln(cmd.OutOrStdout(), "")
			}

			// Process with engagement system before execution (only for existing projects)
			context := map[string]interface{}{
				"command": "/dev",
				"project": absPath,
				"phase":   "development",
				"feature": feature,
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/dev", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run dev workflow (with enhanced brain integration)
			startTime := time.Now()
			devResult, err := dev.runDevWorkflowEnhanced(feature, cmd.OutOrStdout(), memoryCard, orchestrator)
			duration := time.Since(startTime).Seconds()

			if err != nil {
				if stopErr := tracker.Stop(false, err); stopErr != nil {
					return fmt.Errorf("development workflow failed: %w (also failed to stop tracker: %v)", err, stopErr)
				}
				return fmt.Errorf("development workflow failed: %w", err)
			}

			// Update tracker metadata with task_id if available
			if devResult.TaskID != "" {
				tracker.UpdateMetadata("task_id", devResult.TaskID)
			}
			if devResult.FeatureName != "" {
				tracker.UpdateMetadata("feature", devResult.FeatureName)
			}

			// ============================================
			// COMPREHENSIVE ENGAGEMENT PROCESSING
			// ============================================
			// Build rich context for achievements/challenges
			if orchestrator != nil {
				context["success"] = true
				context["development_started"] = true
				context["feature_name"] = devResult.FeatureName
				context["task_id"] = devResult.TaskID
				context["phase"] = devResult.Phase
				context["branch_created"] = devResult.BranchCreated
				context["docs_synced"] = devResult.DocsSynced
				context["development_duration"] = duration

				// Add feature and phase time tracking to context
				if devResult.FeatureName != "" {
					if featureTime, err := timetracker.GetFeatureTime(absPath, devResult.FeatureName); err == nil {
						context["feature_total_time"] = featureTime.Seconds()
					}
				}
				if phaseTime, err := timetracker.GetPhaseTime(absPath, "development"); err == nil {
					context["phase_total_time"] = phaseTime.Seconds()
				}

				// Detect development patterns for achievements
				if memoryCard != nil {
					context["dev_command_count"] = memoryCard.CommandUsage["/dev"]
					context["total_features_worked"] = len(memoryCard.HelpfulFeatures) + len(memoryCard.StruggledFeatures)
				}

				// Process engagement after successful execution
				if err := orchestrator.ProcessCommandWithEngagement("/dev", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}

				// Track interaction with rich data
				response := fmt.Sprintf("Started development on %s (Task %s)", devResult.FeatureName, devResult.TaskID)
				orchestrator.TrackInteraction("/dev", feature, response, duration, "positive")

				// Update memory card with development context
				if orchestrator.memoryCard != nil {
					orchestrator.memoryCard.CurrentProject = absPath
					orchestrator.memoryCard.CurrentPhase = "development"
					orchestrator.memoryCard.LastCommand = "/dev"
					orchestrator.memoryCard.LastCommandTime = time.Now()
					orchestrator.memoryCard.CommandUsage["/dev"]++

					// Learn from development patterns
					if devResult.FeatureName != "" {
						// Track feature worked on
						if !containsString(orchestrator.memoryCard.HelpfulFeatures, devResult.FeatureName) {
							orchestrator.memoryCard.HelpfulFeatures = append(orchestrator.memoryCard.HelpfulFeatures, devResult.FeatureName)
						}
					}

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

			// Display feature time if available
			if devResult.FeatureName != "" {
				if featureTime, err := timetracker.GetFeatureTime(absPath, devResult.FeatureName); err == nil && featureTime > 0 {
					hours := int(featureTime.Hours())
					minutes := int(featureTime.Minutes()) % 60
					if hours > 0 {
						fmt.Fprintf(cmd.OutOrStdout(), "\n⏱️  Total time on '%s': %dh %dm\n", devResult.FeatureName, hours, minutes)
					} else {
						fmt.Fprintf(cmd.OutOrStdout(), "\n⏱️  Total time on '%s': %dm\n", devResult.FeatureName, minutes)
					}
				}
			}

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")
	cmd.Flags().StringVar(&feature, "feature", "", "Feature name or task ID to develop")

	return cmd
}

func resolveProjectPath(path string) (string, error) {
	if path == "" {
		return "", errors.New("project path cannot be empty")
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("%s is not a directory", abs)
	}
	return abs, nil
}
