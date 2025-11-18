package workflow

// GetNextStep returns a recommended next step based on the last completed action.
// It follows the perfect workflow sequence: Plan → Design → Code → Test → Review → Deploy
//
// Returns:
//   - title: The title/heading of the recommendation
//   - description: Detailed description of what to do next
func GetNextStep(lastAction string) (title string, description string) {
	// Map of actions to their recommended next steps
	recommendations := map[string]struct {
		title       string
		description string
	}{
		// Project creation and setup
		"project_created": {
			title:       "Start Planning",
			description: "Use @planner /Discuss to refine your idea, then /Plan to create the project structure. Or use the TUI menu: [p]lan",
		},
		"project_adopted": {
			title:       "Review and Plan",
			description: "Review the project structure, then use @planner /Plan to organize features into phases. Or use the TUI menu: [p]lan",
		},

		// Planning phase
		"plan_complete": {
			title:       "Create Design Specifications",
			description: "Use @designer /Design to create design specifications for your features. Or use the TUI menu: [d]esign",
		},
		"plan_phase_created": {
			title:       "Plan Features",
			description: "Continue planning features within this phase. Use @planner /Plan to add more features, or move to design with @designer /Design",
		},
		"plan_feature_created": {
			title:       "Design Feature",
			description: "Use @designer /Design to create design specifications for this feature. Or use the TUI menu: [d]esign",
		},

		// Design phase
		"design_complete": {
			title:       "Implement Feature",
			description: "Use @coder /Implement to start implementing the feature. Or use the TUI menu: [i]mplement (or select from menu)",
		},
		"design_created": {
			title:       "Review Design",
			description: "Review the design specifications, then use @coder /Implement to begin implementation. Or use the TUI menu: [i]mplement",
		},

		// Implementation phase
		"feature_implemented": {
			title:       "Run Tests",
			description: "Use @tester /Test to run tests and capture screenshots. Or use the TUI menu: [t]est (or security scan)",
		},
		"implementation_started": {
			title:       "Continue Implementation",
			description: "Continue implementing tasks from tasks.md. Update progress as you complete tasks. Use @tester /Test when ready.",
		},
		"task_completed": {
			title:       "Continue Tasks",
			description: "Check off completed tasks in tasks.md and continue with the next task. Use @tester /Test when all tasks are done.",
		},

		// Testing phase
		"tests_passed": {
			title:       "Review Code",
			description: "Use @reviewer /Review to review the code quality and ensure it meets standards. Or use the TUI menu: [r]eview",
		},
		"tests_failed": {
			title:       "Fix Issues",
			description: "Review test failures and fix the issues. Use @coder to address bugs, then run @tester /Test again.",
		},
		"screenshot_captured": {
			title:       "Review Code",
			description: "Screenshots captured! Use @reviewer /Review to review the implementation. Or use the TUI menu: [r]eview",
		},

		// Review phase
		"review_approved": {
			title:       "Deploy Feature",
			description: "Use @devops /Deploy to deploy the feature. Or use the TUI menu: [d]eploy",
		},
		"review_changes_requested": {
			title:       "Address Review Feedback",
			description: "Review the feedback and make necessary changes. Use @coder to implement fixes, then @tester /Test again.",
		},

		// Deployment phase
		"deployment_complete": {
			title:       "Plan Next Feature",
			description: "Great work! Plan the next feature using @planner /Plan, or check progress with the dashboard: [d]ashboard",
		},
		"deployment_started": {
			title:       "Monitor Deployment",
			description: "Monitor the deployment status. Check logs and verify the deployment is successful.",
		},

		// Document generation
		"prd_generated": {
			title:       "Create Project Plan",
			description: "PRD generated! Use @planner /Plan to create phases and features. Or use the TUI menu: [p]lan",
		},
		"documents_generated": {
			title:       "Start Planning",
			description: "Documents generated! Use @planner /Plan to organize features into phases. Or use the TUI menu: [p]lan",
		},

		// Design system
		"design_system_applied": {
			title:       "Design Features",
			description: "Design system ready! Use @designer /Design to create design specifications for your features. Or use the TUI menu: [d]esign",
		},

		// API keys
		"api_keys_configured": {
			title:       "Continue Development",
			description: "API keys configured! Continue with your current workflow. Check dashboard: [d]ashboard or plan next feature: [p]lan",
		},

		// Security
		"security_scan_complete": {
			title:       "Fix Security Issues",
			description: "Review security scan results. Use /fix to auto-fix issues, or address them manually. Then continue with your workflow.",
		},

		// Progress tracking
		"progress_updated": {
			title:       "Continue Work",
			description: "Progress updated! Check the dashboard to see overall progress: [d]ashboard, or continue with current tasks.",
		},

		// Idea discussion
		"idea_discussed": {
			title:       "Generate Documents",
			description: "Idea discussed! Use @planner /Generate to create PRD and contracts, or use the TUI menu: [g]enerate",
		},

		// Auto-fix
		"fix_complete": {
			title:       "Continue Development",
			description: "Issues fixed! Continue with your current workflow. Check progress: [p]rogress, or continue implementation.",
		},

		// Publishing
		"publish_started": {
			title:       "Monitor Publishing",
			description: "Publishing started! Monitor the publishing process. Check status or continue with other tasks.",
		},

		// Dev server
		"dev_server_started": {
			title:       "Continue Development",
			description: "Dev server running! Continue implementing features or run tests. Use the TUI menu for next steps.",
		},

		// Action undone
		"action_undone": {
			title:       "Review Changes",
			description: "Action undone! Review what was changed and decide on next steps. Check dashboard: [d]ashboard",
		},

		// Integration setup
		"integration_setup": {
			title:       "Start Planning",
			description: "IDE integration configured! Start planning your project with @planner /Plan, or use the TUI menu: [p]lan",
		},

		// Default fallback
		"": {
			title:       "Get Started",
			description: "Welcome to DoPlan! Start by planning your project with @planner /Plan, or use the TUI menu: [p]lan",
		},
	}

	// Get recommendation for the action
	rec, exists := recommendations[lastAction]
	if !exists {
		// Return default recommendation
		rec = recommendations[""]
	}

	return rec.title, rec.description
}

// GetWorkflowSequence returns the perfect workflow sequence as a list
func GetWorkflowSequence() []string {
	return []string{
		"Plan → Discuss idea, refine, generate PRD, create phases and features",
		"Design → Create design specifications following DPR",
		"Code → Implement features based on plans and designs",
		"Test → Run tests and capture screenshots",
		"Review → Review code quality and approve",
		"Deploy → Deploy to staging and production",
	}
}

// GetCurrentStep returns the current step in the workflow based on project state
func GetCurrentStep(hasPlan, hasDesign, hasImplementation, hasTests, hasReview, hasDeployment bool) string {
	if !hasPlan {
		return "Plan"
	}
	if !hasDesign {
		return "Design"
	}
	if !hasImplementation {
		return "Code"
	}
	if !hasTests {
		return "Test"
	}
	if !hasReview {
		return "Review"
	}
	if !hasDeployment {
		return "Deploy"
	}
	return "Complete"
}

