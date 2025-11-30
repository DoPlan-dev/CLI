package cli

// GetAllChallengeDefinitions returns all challenge definitions
func GetAllChallengeDefinitions() []ChallengeDefinition {
	definitions := []ChallengeDefinition{}

	// Integration Challenges
	definitions = append(definitions, getIntegrationChallenges()...)

	// Database Challenges
	definitions = append(definitions, getDatabaseChallenges()...)

	// Deployment Challenges
	definitions = append(definitions, getDeploymentChallenges()...)

	// Testing Challenges
	definitions = append(definitions, getTestingChallenges()...)

	// Workflow Challenges
	definitions = append(definitions, getWorkflowChallenges()...)

	// Release Challenges
	definitions = append(definitions, getReleaseChallenges()...)

	// Performance Challenges
	definitions = append(definitions, getPerformanceChallenges()...)

	// Security Challenges
	definitions = append(definitions, getSecurityChallenges()...)

	return definitions
}

// Integration Challenges
func getIntegrationChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "first_api_integration",
			Title:       "API Integration Master",
			Description: "Generate and integrate your first API",
			Category:    "integration",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🔌",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if apiCreated, ok := context["api_created"].(bool); ok && apiCreated {
					return true
				}
				if integration, ok := context["integration"].(string); ok && integration == "api" {
					return true
				}
				return false
			},
		},
		{
			ID:          "api_integration_tested",
			Title:       "Integration Tested",
			Description: "Add API for integration and pass all tests",
			Category:    "integration",
			Points:      750,
			Rarity:      "epic",
			Icon:        "✅",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if apiTested, ok := context["api_tested"].(bool); ok && apiTested {
					if testsPassed, ok := context["tests_passed"].(bool); ok && testsPassed {
						return true
					}
				}
				return false
			},
		},
		{
			ID:          "third_party_integration",
			Title:       "Third-Party Connector",
			Description: "Integrate with a third-party service (Stripe, Auth0, etc.)",
			Category:    "integration",
			Points:      600,
			Rarity:      "epic",
			Icon:        "🔗",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if thirdParty, ok := context["third_party_integration"].(bool); ok && thirdParty {
					return true
				}
				return false
			},
		},
		{
			ID:          "webhook_setup",
			Title:       "Webhook Wizard",
			Description: "Set up your first webhook",
			Category:    "integration",
			Points:      400,
			Rarity:      "rare",
			Icon:        "🪝",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if webhook, ok := context["webhook_setup"].(bool); ok && webhook {
					return true
				}
				return false
			},
		},
	}
}

// Database Challenges
func getDatabaseChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "first_database_link",
			Title:       "Database Connected",
			Description: "Link your first database",
			Category:    "database",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🗄️",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if dbLinked, ok := context["database_linked"].(bool); ok && dbLinked {
					return true
				}
				if dbAction, ok := context["database_action"].(string); ok && dbAction == "linked" {
					return true
				}
				return false
			},
		},
		{
			ID:          "database_merge",
			Title:       "Database Merger",
			Description: "Successfully merge databases",
			Category:    "database",
			Points:      800,
			Rarity:      "epic",
			Icon:        "🔀",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if dbMerged, ok := context["database_merged"].(bool); ok && dbMerged {
					return true
				}
				if dbAction, ok := context["database_action"].(string); ok && dbAction == "merged" {
					return true
				}
				return false
			},
		},
		{
			ID:          "database_migration",
			Title:       "Migration Master",
			Description: "Create and run your first database migration",
			Category:    "database",
			Points:      400,
			Rarity:      "rare",
			Icon:        "📦",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if migration, ok := context["migration_created"].(bool); ok && migration {
					return true
				}
				return false
			},
		},
		{
			ID:          "database_backup",
			Title:       "Safety First",
			Description: "Set up automated database backups",
			Category:    "database",
			Points:      300,
			Rarity:      "rare",
			Icon:        "💾",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if backup, ok := context["backup_configured"].(bool); ok && backup {
					return true
				}
				return false
			},
		},
	}
}

// Deployment Challenges
func getDeploymentChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "first_deployment",
			Title:       "First Launch",
			Description: "Deploy your project for the first time",
			Category:    "deployment",
			Points:      1000,
			Rarity:      "legendary",
			Icon:        "🚀",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if deployed, ok := context["deployed"].(bool); ok && deployed {
					return true
				}
				if phase, ok := context["phase"].(string); ok && phase == "deployed" {
					return true
				}
				return false
			},
		},
		{
			ID:          "production_deployment",
			Title:       "Production Ready",
			Description: "Deploy to production environment",
			Category:    "deployment",
			Points:      1500,
			Rarity:      "legendary",
			Icon:        "🌟",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if prodDeployed, ok := context["production_deployed"].(bool); ok && prodDeployed {
					return true
				}
				if env, ok := context["environment"].(string); ok && env == "production" {
					return true
				}
				return false
			},
		},
		{
			ID:          "ci_cd_setup",
			Title:       "Automation Master",
			Description: "Set up CI/CD pipeline",
			Category:    "deployment",
			Points:      600,
			Rarity:      "epic",
			Icon:        "⚙️",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if cicd, ok := context["cicd_setup"].(bool); ok && cicd {
					return true
				}
				return false
			},
		},
		{
			ID:          "docker_deployment",
			Title:       "Container Master",
			Description: "Deploy using Docker containers",
			Category:    "deployment",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🐳",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if docker, ok := context["docker_deployed"].(bool); ok && docker {
					return true
				}
				return false
			},
		},
		{
			ID:          "kubernetes_deployment",
			Title:       "Kubernetes Master",
			Description: "Deploy to Kubernetes",
			Category:    "deployment",
			Points:      1200,
			Rarity:      "legendary",
			Icon:        "☸️",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if k8s, ok := context["kubernetes_deployed"].(bool); ok && k8s {
					return true
				}
				return false
			},
		},
	}
}

// Testing Challenges
func getTestingChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "test_coverage_80",
			Title:       "Test Coverage Champion",
			Description: "Achieve 80% test coverage",
			Category:    "testing",
			Points:      600,
			Rarity:      "epic",
			Icon:        "📊",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if coverage, ok := context["test_coverage"].(float64); ok {
					return coverage >= 80.0 && coverage < 90.0
				}
				return false
			},
		},
		{
			ID:          "test_coverage_90",
			Title:       "Test Coverage Master",
			Description: "Achieve 90% test coverage",
			Category:    "testing",
			Points:      1000,
			Rarity:      "legendary",
			Icon:        "🏆",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if coverage, ok := context["test_coverage"].(float64); ok {
					return coverage >= 90.0
				}
				return false
			},
		},
		{
			ID:          "all_tests_passing",
			Title:       "Perfect Tests",
			Description: "Get all tests passing",
			Category:    "testing",
			Points:      400,
			Rarity:      "rare",
			Icon:        "✅",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if allPassing, ok := context["all_tests_passing"].(bool); ok && allPassing {
					return true
				}
				return false
			},
		},
		{
			ID:          "e2e_tests",
			Title:       "End-to-End Master",
			Description: "Set up and pass end-to-end tests",
			Category:    "testing",
			Points:      700,
			Rarity:      "epic",
			Icon:        "🔗",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if e2e, ok := context["e2e_tests_passing"].(bool); ok && e2e {
					return true
				}
				return false
			},
		},
		{
			ID:          "performance_tests",
			Title:       "Performance Tester",
			Description: "Set up and pass performance tests",
			Category:    "testing",
			Points:      500,
			Rarity:      "epic",
			Icon:        "⚡",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if perf, ok := context["performance_tests_passing"].(bool); ok && perf {
					return true
				}
				return false
			},
		},
	}
}

// Workflow Challenges
func getWorkflowChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "best_github_workflow",
			Title:       "GitHub Workflow Master",
			Description: "Complete project with best GitHub branch workflow",
			Category:    "workflow",
			Points:      800,
			Rarity:      "epic",
			Icon:        "🌿",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if bestWorkflow, ok := context["best_github_workflow"].(bool); ok && bestWorkflow {
					return true
				}
				if workflow, ok := context["workflow_quality"].(string); ok && workflow == "best" {
					return true
				}
				return false
			},
		},
		{
			ID:          "conventional_commits",
			Title:       "Commit Master",
			Description: "Use conventional commits throughout project",
			Category:    "workflow",
			Points:      400,
			Rarity:      "rare",
			Icon:        "📝",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if conventional, ok := context["conventional_commits"].(bool); ok && conventional {
					return true
				}
				return false
			},
		},
		{
			ID:          "code_review_process",
			Title:       "Code Review Pro",
			Description: "Set up code review process",
			Category:    "workflow",
			Points:      500,
			Rarity:      "epic",
			Icon:        "👀",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if review, ok := context["code_review_setup"].(bool); ok && review {
					return true
				}
				return false
			},
		},
		{
			ID:          "automated_branching",
			Title:       "Branch Automation",
			Description: "Set up automated branch management",
			Category:    "workflow",
			Points:      600,
			Rarity:      "epic",
			Icon:        "🤖",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if automated, ok := context["branch_automation"].(bool); ok && automated {
					return true
				}
				return false
			},
		},
	}
}

// Release Challenges
func getReleaseChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "first_public_release",
			Title:       "Public Launch",
			Description: "Make your first public release",
			Category:    "release",
			Points:      2000,
			Rarity:      "legendary",
			Icon:        "🎉",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if publicRelease, ok := context["public_release"].(bool); ok && publicRelease {
					return true
				}
				if release, ok := context["release_type"].(string); ok && release == "public" {
					return true
				}
				return false
			},
		},
		{
			ID:          "version_1_0",
			Title:       "Version 1.0",
			Description: "Release version 1.0",
			Category:    "release",
			Points:      1500,
			Rarity:      "legendary",
			Icon:        "1️⃣",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if version, ok := context["version"].(string); ok && version == "1.0.0" {
					return true
				}
				return false
			},
		},
		{
			ID:          "release_notes",
			Title:       "Release Notes Pro",
			Description: "Create comprehensive release notes",
			Category:    "release",
			Points:      300,
			Rarity:      "rare",
			Icon:        "📋",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if notes, ok := context["release_notes_created"].(bool); ok && notes {
					return true
				}
				return false
			},
		},
		{
			ID:          "changelog_maintained",
			Title:       "Changelog Keeper",
			Description: "Maintain changelog throughout project",
			Category:    "release",
			Points:      400,
			Rarity:      "rare",
			Icon:        "📝",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if changelog, ok := context["changelog_maintained"].(bool); ok && changelog {
					return true
				}
				return false
			},
		},
	}
}

// Performance Challenges
func getPerformanceChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "performance_optimized",
			Title:       "Performance Optimizer",
			Description: "Optimize application performance",
			Category:    "performance",
			Points:      600,
			Rarity:      "epic",
			Icon:        "⚡",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if optimized, ok := context["performance_optimized"].(bool); ok && optimized {
					return true
				}
				return false
			},
		},
		{
			ID:          "lighthouse_90",
			Title:       "Lighthouse Master",
			Description: "Achieve 90+ Lighthouse score",
			Category:    "performance",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🏮",
			IsFirstTime: false,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if score, ok := context["lighthouse_score"].(float64); ok {
					return score >= 90.0
				}
				return false
			},
		},
	}
}

// Security Challenges
func getSecurityChallenges() []ChallengeDefinition {
	return []ChallengeDefinition{
		{
			ID:          "security_audit",
			Title:       "Security Auditor",
			Description: "Complete security audit",
			Category:    "security",
			Points:      700,
			Rarity:      "epic",
			Icon:        "🔒",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if audit, ok := context["security_audit"].(bool); ok && audit {
					return true
				}
				return false
			},
		},
		{
			ID:          "vulnerability_scan",
			Title:       "Vulnerability Scanner",
			Description: "Run and fix vulnerability scan",
			Category:    "security",
			Points:      500,
			Rarity:      "epic",
			Icon:        "🛡️",
			IsFirstTime: true,
			Condition: func(mc *MemoryCard, context map[string]interface{}) bool {
				if scan, ok := context["vulnerability_scan"].(bool); ok && scan {
					return true
				}
				return false
			},
		},
	}
}
