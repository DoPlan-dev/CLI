package cli

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	timetracker "github.com/DoPlan-dev/CLI/internal/time"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

func init() {
	rootCmd.AddCommand(newHeyCommand())
	rootCmd.AddCommand(newDoCommand())
}

func newHeyCommand() *cobra.Command {
	var projectPath string

	cmd := &cobra.Command{
		Use:   "/hey",
		Short: "Welcome, tutorial, and command introductions",
		Long: `The /hey command provides an interactive onboarding experience:
- First-time welcome and tutorial
- System overview and agent hierarchy
- Command walkthrough and test drive
- Personalized tips based on experience level`,
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := resolveProjectPath(projectPath)
			if err != nil {
				return err
			}

			// Fast path: Check if this is a new project (no memory card or first time)
			// For new projects, skip heavy engagement system initialization
			memoryCardPath := filepath.Join(absPath, ".do", "system", "memory_card.json")
			isNewProject := !utils.PathExists(memoryCardPath)

			var orchestrator *EngagementOrchestrator
			if !isNewProject {
				// Only initialize engagement system for existing projects
				var err error
				orchestrator, err = NewEngagementOrchestrator()
				if err != nil {
					// Log but don't fail - graceful degradation
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Could not initialize engagement system: %v\n", err)
				}
			}

			// Initialize time tracker (lightweight)
			tracker, err := timetracker.New(absPath)
			if err != nil {
				return err
			}

			// Start tracking onboarding phase
			metadata := map[string]string{
				"command_type":   "onboarding",
				"is_new_project": fmt.Sprintf("%v", isNewProject),
			}
			tracker.Start(absPath, "onboarding", "/hey", args, metadata)

			// For new projects, skip heavy engagement processing
			if !isNewProject && orchestrator != nil {
				// Process with engagement system (only for existing projects)
				context := map[string]interface{}{
					"command":    "/hey",
					"project":    absPath,
					"phase":      "onboarding",
					"first_time": false,
				}

				if err := orchestrator.ProcessCommandWithEngagement("/hey", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// TODO: Implement actual /hey command logic
			// This would involve:
			// 1. Check if first time user (.do/system/user_profile.json)
			// 2. Run interactive TUI tutorial
			// 3. Save user profile and reference materials
			// For now, we'll just print a message and track time
			fmt.Fprintln(cmd.OutOrStdout(), "👋 Welcome to DoPlan!")
			fmt.Fprintln(cmd.OutOrStdout(), "Starting onboarding tutorial...")
			fmt.Fprintln(cmd.OutOrStdout(), "(Tutorial implementation coming soon)")

			// Track interaction (lightweight for new projects)
			if orchestrator != nil {
				orchestrator.TrackInteraction("/hey", "", "Welcome message", 1.0, "positive")
			} else if !isNewProject {
				// For existing projects without orchestrator, still track basic interaction
				if mc, err := LoadMemoryCard(); err == nil {
					mc.LastCommand = "/hey"
					mc.LastCommandTime = time.Now()
					if mc.CommandUsage == nil {
						mc.CommandUsage = make(map[string]int)
					}
					mc.CommandUsage["/hey"]++
					SaveMemoryCard(mc)
				}
			}

			// Stop tracking with success
			if stopErr := tracker.Stop(true, nil); stopErr != nil {
				return stopErr
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")

	return cmd
}

func newDoCommand() *cobra.Command {
	var projectPath string
	var idea string

	cmd := &cobra.Command{
		Use:   "/do",
		Short: "Capture project idea, conduct meeting, and refine suggestions",
		Long: `The /do command orchestrates the complete project initiation process:
1. Ideation - Capture your project idea (iterative conversation)
2. Meeting - Conduct discovery meeting to understand requirements
3. Refining - Enhance and provide suggestions for improvement

Usage:
  /do                          # Interactive mode - prompts for idea
  /do "Build a todo app"       # Direct mode - provide idea as argument
  /do feature                   # Add idea about a single feature
  /do now                       # Fast track with detailed prompt/PRD (skips to planning)
  /do i'm lucky                 # Get iterative idea suggestions for learning/inspiration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := resolveProjectPath(projectPath)
			if err != nil {
				return err
			}

			// Initialize engagement orchestrator
			orchestrator, err := NewEngagementOrchestrator()
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Could not initialize engagement system: %v\n", err)
			}

			// Initialize time tracker
			tracker, err := timetracker.New(absPath)
			if err != nil {
				return err
			}

			// Load memory card (cached for performance)
			var memoryCard *MemoryCard
			if !shouldUseFastPath(absPath) {
				memoryCard, err = loadMemoryCardCached()
				if err != nil {
					memoryCard = nil
				}
			}

			// Initialize do logic
			do, err := newDoLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /do logic: %w", err)
			}

			// Capture idea from args or flag
			if idea == "" && len(args) > 0 {
				idea = args[0]
			}

			// ============================================
			// PHASE 1: IDEATION
			// ============================================
			DisplayPhaseProgress(cmd.OutOrStdout(), "Ideation", 1, 3)

			ideationMetadata := map[string]string{}
			if idea != "" {
				ideationMetadata["idea_provided"] = "true"
				if len(idea) > 100 {
					ideationMetadata["idea_preview"] = idea[:100] + "..."
				} else {
					ideationMetadata["idea_preview"] = idea
				}
			} else {
				ideationMetadata["idea_provided"] = "false"
			}

			tracker.Start(absPath, "ideation", "/do", args, ideationMetadata)

			// Process engagement before ideation (only for existing projects)
			ideationContext := map[string]interface{}{
				"command": "/do",
				"project": absPath,
				"phase":   "ideation",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do", ideationContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run ideation phase
			startTime := time.Now()
			capturedIdea, err := do.runIdeationPhase(idea, cmd.OutOrStdout(), memoryCard)
			ideationDuration := time.Since(startTime).Seconds()

			if err != nil {
				if stopErr := tracker.Stop(false, err); stopErr != nil {
					return fmt.Errorf("ideation failed: %w (also failed to stop tracker: %v)", err, stopErr)
				}
				return fmt.Errorf("ideation failed: %w", err)
			}

			// Process engagement after ideation
			if orchestrator != nil {
				ideationContext["success"] = true
				ideationContext["ideation_completed"] = true
				ideationContext["idea_captured"] = capturedIdea != ""
				if err := orchestrator.ProcessCommandWithEngagement("/do", ideationContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/do", idea, "Ideation phase completed", ideationDuration, "positive")
			}

			// Stop ideation phase
			if stopErr := tracker.Stop(true, nil); stopErr != nil {
				return stopErr
			}

			// Display phase completion with time
			fmt.Fprintf(cmd.OutOrStdout(), "\n✅ Ideation phase completed in %.1f seconds\n\n", ideationDuration)

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			// ============================================
			// PHASE 2: MEETING
			// ============================================
			DisplayPhaseProgress(cmd.OutOrStdout(), "Discovery Meeting", 2, 3)

			meetingMetadata := map[string]string{
				"meeting_type": "discovery",
			}
			if capturedIdea != "" {
				meetingMetadata["idea_provided"] = "true"
			} else {
				meetingMetadata["idea_provided"] = "false"
			}

			tracker.Start(absPath, "meeting", "/do", args, meetingMetadata)

			// Process engagement before meeting (only for existing projects)
			meetingContext := map[string]interface{}{
				"command":      "/do",
				"project":      absPath,
				"phase":        "meeting",
				"meeting_type": "discovery",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do", meetingContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run meeting phase
			startTime = time.Now()
			if err := do.runMeetingPhase(capturedIdea, cmd.OutOrStdout()); err != nil {
				if stopErr := tracker.Stop(false, err); stopErr != nil {
					return fmt.Errorf("meeting failed: %w (also failed to stop tracker: %v)", err, stopErr)
				}
				return fmt.Errorf("meeting failed: %w", err)
			}
			meetingDuration := time.Since(startTime).Seconds()

			// Process engagement after meeting (important for achievements/challenges)
			if orchestrator != nil {
				meetingContext["success"] = true
				meetingContext["meeting_completed"] = true
				meetingContext["project_started"] = true // First project milestone
				if err := orchestrator.ProcessCommandWithEngagement("/do", meetingContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/do", capturedIdea, "Discovery meeting completed", meetingDuration, "positive")
			}

			// Stop meeting phase
			if stopErr := tracker.Stop(true, nil); stopErr != nil {
				return stopErr
			}

			// Display phase completion with time
			fmt.Fprintf(cmd.OutOrStdout(), "\n✅ Discovery meeting completed in %.1f seconds\n\n", meetingDuration)

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			// ============================================
			// PHASE 3: REFINING
			// ============================================
			DisplayPhaseProgress(cmd.OutOrStdout(), "Refining & Suggestions", 3, 3)

			refiningMetadata := map[string]string{
				"refining_type": "enhancement_and_suggestions",
			}

			tracker.Start(absPath, "refining", "/do", args, refiningMetadata)

			// Process engagement before refining (only for existing projects)
			refiningContext := map[string]interface{}{
				"command": "/do",
				"project": absPath,
				"phase":   "refining",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do", refiningContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run refining phase
			startTime = time.Now()
			if err := do.runRefiningPhase(cmd.OutOrStdout()); err != nil {
				if stopErr := tracker.Stop(false, err); stopErr != nil {
					return fmt.Errorf("refining failed: %w (also failed to stop tracker: %v)", err, stopErr)
				}
				return fmt.Errorf("refining failed: %w", err)
			}
			refiningDuration := time.Since(startTime).Seconds()

			// Process engagement after refining
			if orchestrator != nil {
				refiningContext["success"] = true
				refiningContext["refining_completed"] = true
				refiningContext["project_initiation_complete"] = true
				if err := orchestrator.ProcessCommandWithEngagement("/do", refiningContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/do", "", "Refining phase completed", refiningDuration, "positive")
			}

			// Stop refining phase
			if stopErr := tracker.Stop(true, nil); stopErr != nil {
				return stopErr
			}

			// Display phase completion with time
			fmt.Fprintf(cmd.OutOrStdout(), "\n✅ Refining phase completed in %.1f seconds\n\n", refiningDuration)

			// Final engagement processing for complete workflow
			completeContext := map[string]interface{}{
				"command":              "/do",
				"project":              absPath,
				"phase":                "complete",
				"workflow_complete":    true,
				"all_phases_completed": true,
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do", completeContext, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintln(cmd.OutOrStdout(), "✅ Project initiation complete!")
			fmt.Fprintln(cmd.OutOrStdout(), "   ✓ Idea captured and saved")
			fmt.Fprintln(cmd.OutOrStdout(), "   ✓ Discovery meeting completed")
			fmt.Fprintln(cmd.OutOrStdout(), "   ✓ Refinements and suggestions generated")
			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintln(cmd.OutOrStdout(), "📋 Next steps:")
			fmt.Fprintln(cmd.OutOrStdout(), "   • Review IDEA.md and BRAINSTORM.md")
			fmt.Fprintln(cmd.OutOrStdout(), "   • Check REFINEMENTS.md for suggestions")
			fmt.Fprintln(cmd.OutOrStdout(), "   • Type /plan to generate your execution plan")

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")
	cmd.Flags().StringVar(&idea, "idea", "", "Project idea to capture")

	// Add subcommands
	cmd.AddCommand(newDoFeatureCommand())
	cmd.AddCommand(newDoNowCommand())
	cmd.AddCommand(newDoImLuckyCommand())

	return cmd
}

// newDoFeatureCommand creates the /do feature subcommand
func newDoFeatureCommand() *cobra.Command {
	var projectPath string

	cmd := &cobra.Command{
		Use:   "feature",
		Short: "Add idea about a single feature",
		Long:  `The /do feature command allows you to add or refine a single feature idea to your project.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := resolveProjectPath(projectPath)
			if err != nil {
				return err
			}

			// Initialize engagement orchestrator
			orchestrator, err := NewEngagementOrchestrator()
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Could not initialize engagement system: %v\n", err)
			}

			// Load memory card
			memoryCard, _ := LoadMemoryCard()

			// Initialize do logic
			do, err := newDoLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /do logic: %w", err)
			}

			// Process engagement
			context := map[string]interface{}{
				"command": "/do feature",
				"project": absPath,
				"phase":   "feature_ideation",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do feature", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run feature ideation
			startTime := time.Now()
			featureIdea, err := do.runFeatureIdeation(cmd.OutOrStdout(), memoryCard)
			duration := time.Since(startTime).Seconds()

			if err != nil {
				return err
			}

			// Process engagement after completion
			context["success"] = true
			context["feature_idea_added"] = true
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do feature", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/do feature", featureIdea, "Feature idea added", duration, "positive")
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✅ Feature idea added: %s\n", featureIdea)

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")
	return cmd
}

// newDoNowCommand creates the /do now subcommand
func newDoNowCommand() *cobra.Command {
	var projectPath string
	var prdFile string
	var prompt string

	cmd := &cobra.Command{
		Use:   "now",
		Short: "Fast track with detailed prompt/PRD (skips to planning)",
		Long: `The /do now command is for users who have very detailed prompts or PRD files.
It skips the discovery meeting and goes directly to planning.

Usage:
  /do now                                    # Interactive mode
  /do now --prompt "detailed prompt"         # Provide detailed prompt
  /do now --prd path/to/PRD.md              # Provide PRD file
  /do now --prompt "..." --prd path/to/PRD.md # Both prompt and PRD`,
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := resolveProjectPath(projectPath)
			if err != nil {
				return err
			}

			// Initialize engagement orchestrator
			orchestrator, err := NewEngagementOrchestrator()
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Could not initialize engagement system: %v\n", err)
			}

			// Initialize do logic
			do, err := newDoLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /do logic: %w", err)
			}

			// Process engagement
			context := map[string]interface{}{
				"command":    "/do now",
				"project":    absPath,
				"phase":      "fast_track",
				"has_prompt": prompt != "",
				"has_prd":    prdFile != "",
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do now", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run fast track
			startTime := time.Now()
			if err := do.runFastTrack(cmd.OutOrStdout(), prompt, prdFile); err != nil {
				return err
			}
			duration := time.Since(startTime).Seconds()

			// Process engagement after completion
			context["success"] = true
			context["fast_track_completed"] = true
			context["ready_for_planning"] = true
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do now", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/do now", prompt, "Fast track completed", duration, "positive")
			}

			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintln(cmd.OutOrStdout(), "✅ Fast track complete! Ready to plan.")
			fmt.Fprintln(cmd.OutOrStdout(), "   Type /plan to generate your execution plan")

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")
	cmd.Flags().StringVar(&prdFile, "prd", "", "Path to PRD file")
	cmd.Flags().StringVar(&prompt, "prompt", "", "Detailed prompt about the project")

	return cmd
}

// newDoImLuckyCommand creates the /do i'm lucky subcommand
func newDoImLuckyCommand() *cobra.Command {
	var projectPath string

	cmd := &cobra.Command{
		Use:     "i'm lucky",
		Aliases: []string{"im-lucky", "lucky"},
		Short:   "Get iterative idea suggestions for learning/inspiration",
		Long: `The /do i'm lucky command helps you discover new ideas through iterative suggestions.
The agent will suggest 2 ideas, you choose one, then it suggests 2 more based on your choice.
This continues until you find something you love. No meeting required - goes straight to planning.

This is perfect for:
- Learning new technologies
- Finding inspiration
- Exploring different project ideas
- Discovering what you're passionate about`,
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := resolveProjectPath(projectPath)
			if err != nil {
				return err
			}

			// Initialize engagement orchestrator
			orchestrator, err := NewEngagementOrchestrator()
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Could not initialize engagement system: %v\n", err)
			}

			// Load memory card
			memoryCard, _ := LoadMemoryCard()

			// Initialize do logic
			do, err := newDoLogic(absPath)
			if err != nil {
				return fmt.Errorf("failed to initialize /do logic: %w", err)
			}

			// Process engagement
			context := map[string]interface{}{
				"command":       "/do i'm lucky",
				"project":       absPath,
				"phase":         "lucky_mode",
				"learning_mode": true,
			}
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do i'm lucky", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
			}

			// Run lucky mode
			startTime := time.Now()
			selectedIdea, err := do.runLuckyMode(cmd.OutOrStdout(), memoryCard)
			duration := time.Since(startTime).Seconds()

			if err != nil {
				return err
			}

			// Process engagement after completion
			context["success"] = true
			context["lucky_mode_completed"] = true
			context["idea_selected"] = true
			context["ready_for_planning"] = true
			if orchestrator != nil {
				if err := orchestrator.ProcessCommandWithEngagement("/do i'm lucky", context, cmd.OutOrStdout()); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: Engagement processing failed: %v\n", err)
				}
				orchestrator.TrackInteraction("/do i'm lucky", selectedIdea, "Lucky mode completed", duration, "excited")
			}

			fmt.Fprintln(cmd.OutOrStdout(), "")
			fmt.Fprintf(cmd.OutOrStdout(), "🎉 Great choice! Selected idea: %s\n", selectedIdea)
			fmt.Fprintln(cmd.OutOrStdout(), "   Type /plan to generate your execution plan")

			// Silently update dashboard data
			_ = UpdateDashboardData(absPath)

			return nil
		},
	}

	cmd.Flags().StringVar(&projectPath, "project", ".", "Path to project root")
	return cmd
}
