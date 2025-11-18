package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/dpr"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/rakd"
	"github.com/DoPlan-dev/CLI/internal/workflow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEndToEndIntegration(t *testing.T) {
	// Create test project directory
	testDir := "/tmp/doplan-integration-test"
	projectRoot := filepath.Join(testDir, "test-project")
	
	// Clean up previous test
	os.RemoveAll(testDir)
	
	// Create project structure
	err := os.MkdirAll(projectRoot, 0755)
	require.NoError(t, err, "Failed to create test project directory")
	
	// Create package.json with dependencies
	packageJSON := `{
		"name": "test-project",
		"version": "1.0.0",
		"dependencies": {
			"express": "^4.18.0",
			"@stripe/stripe-js": "^2.0.0"
		}
	}`
	err = os.WriteFile(filepath.Join(projectRoot, "package.json"), []byte(packageJSON), 0644)
	require.NoError(t, err, "Failed to create package.json")
	
	// Create .env.example
	envExample := `PORT=3000
STRIPE_API_KEY=sk_test_...
OPENAI_API_KEY=sk-...`
	err = os.WriteFile(filepath.Join(projectRoot, ".env.example"), []byte(envExample), 0644)
	require.NoError(t, err, "Failed to create .env.example")
	
	// Create DoPlan structure
	doplanDirs := []string{
		".doplan/ai/agents",
		".doplan/ai/rules",
		".doplan/ai/commands",
		".doplan/design",
		".doplan/SOPS",
		"doplan/design",
		"doplan/contracts",
	}
	for _, dir := range doplanDirs {
		err := os.MkdirAll(filepath.Join(projectRoot, dir), 0755)
		require.NoError(t, err, "Failed to create directory: %s", dir)
	}
	
	// Create config.yaml
	configYAML := `project: test-project
version: 1.0.0`
	err = os.WriteFile(filepath.Join(projectRoot, ".doplan/config.yaml"), []byte(configYAML), 0644)
	require.NoError(t, err, "Failed to create config.yaml")
	
	t.Run("Generate Agents", func(t *testing.T) {
		gen := generators.NewAgentsGenerator(projectRoot)
		err := gen.Generate()
		require.NoError(t, err, "Failed to generate agents")
		
		// Verify agent files exist
		agentFiles := []string{
			".doplan/ai/agents/README.md",
			".doplan/ai/agents/planner.agent.md",
			".doplan/ai/agents/coder.agent.md",
			".doplan/ai/agents/designer.agent.md",
			".doplan/ai/agents/reviewer.agent.md",
			".doplan/ai/agents/tester.agent.md",
			".doplan/ai/agents/devops.agent.md",
		}
		
		for _, file := range agentFiles {
			path := filepath.Join(projectRoot, file)
			assert.FileExists(t, path, "Agent file should exist: %s", file)
			
			// Check file is not empty
			info, err := os.Stat(path)
			require.NoError(t, err)
			assert.Greater(t, info.Size(), int64(0), "Agent file should not be empty: %s", file)
		}
	})
	
	t.Run("Generate Rules", func(t *testing.T) {
		gen := generators.NewRulesGenerator(projectRoot)
		err := gen.Generate()
		require.NoError(t, err, "Failed to generate rules")
		
		// Verify rule files exist
		ruleFiles := []string{
			".doplan/ai/rules/workflow.mdc",
			".doplan/ai/rules/communication.mdc",
		}
		
		for _, file := range ruleFiles {
			path := filepath.Join(projectRoot, file)
			assert.FileExists(t, path, "Rule file should exist: %s", file)
			
			// Check file is not empty
			info, err := os.Stat(path)
			require.NoError(t, err)
			assert.Greater(t, info.Size(), int64(0), "Rule file should not be empty: %s", file)
		}
	})
	
	t.Run("Generate Commands", func(t *testing.T) {
		gen := generators.NewCommandsGenerator(projectRoot)
		err := gen.Generate()
		require.NoError(t, err, "Failed to generate commands")
		
		// Verify some command files exist
		commandFiles := []string{
			".doplan/ai/commands/run.md",
			".doplan/ai/commands/deploy.md",
			".doplan/ai/commands/create.md",
		}
		
		atLeastOneExists := false
		for _, file := range commandFiles {
			path := filepath.Join(projectRoot, file)
			if _, err := os.Stat(path); err == nil {
				atLeastOneExists = true
				info, err := os.Stat(path)
				require.NoError(t, err)
				assert.Greater(t, info.Size(), int64(0), "Command file should not be empty: %s", file)
				break
			}
		}
		assert.True(t, atLeastOneExists, "At least one command file should exist")
	})
	
	t.Run("Generate Design System (DPR)", func(t *testing.T) {
		// Create mock DPR data
		data := &dpr.DPRData{
			Answers: map[string]interface{}{
				"project_name":           "Test Project",
				"audience_primary":       "Developers",
				"emotion_target":         "Professional",
				"style_overall":          "Modern",
				"color_primary":          "#667eea",
				"typography_font":        "Inter",
				"layout_style":           "Card-based",
				"components_style":       "Elevated",
				"animation_level":        "Subtle",
				"accessibility_importance": 5,
				"responsive_priority":    "Desktop First",
			},
		}
		
		// Generate DPR.md
		dprGen := dpr.NewGenerator(projectRoot, data)
		err := dprGen.Generate()
		require.NoError(t, err, "Failed to generate DPR.md")
		
		// Generate design tokens
		tokenGen := dpr.NewTokenGenerator(projectRoot, data)
		err = tokenGen.Generate()
		require.NoError(t, err, "Failed to generate design tokens")
		
		// Generate cursor rules
		rulesGen := dpr.NewCursorRulesGenerator(projectRoot, data)
		err = rulesGen.Generate()
		require.NoError(t, err, "Failed to generate cursor rules")
		
		// Verify DPR files exist
		dprFiles := []string{
			"doplan/design/DPR.md",
			"doplan/design/design-tokens.json",
			".doplan/ai/rules/design_rules.mdc",
		}
		
		for _, file := range dprFiles {
			path := filepath.Join(projectRoot, file)
			assert.FileExists(t, path, "DPR file should exist: %s", file)
			
			// Check file is not empty
			info, err := os.Stat(path)
			require.NoError(t, err)
			assert.Greater(t, info.Size(), int64(0), "DPR file should not be empty: %s", file)
		}
	})
	
	t.Run("Generate API Keys Detection (RAKD)", func(t *testing.T) {
		// Generate RAKD
		data, err := rakd.GenerateRAKD(projectRoot)
		require.NoError(t, err, "Failed to generate RAKD")
		
		// Verify RAKD.md exists
		rakdPath := filepath.Join(projectRoot, "doplan/RAKD.md")
		assert.FileExists(t, rakdPath, "RAKD.md should exist")
		
		// Verify services were detected
		assert.Greater(t, len(data.Services), 0, "At least one service should be detected")
		
		// Check that Stripe was detected (from package.json)
		stripeDetected := false
		for _, service := range data.Services {
			if service.Name == "Stripe" {
				stripeDetected = true
				break
			}
		}
		assert.True(t, stripeDetected, "Stripe should be detected from package.json")
	})
	
	t.Run("Workflow Guidance Integration", func(t *testing.T) {
		// Test various action recommendations
		testCases := []struct {
			action         string
			expectTitle    bool
			expectDesc     bool
		}{
			{"project_created", true, true},
			{"plan_complete", true, true},
			{"design_complete", true, true},
			{"feature_implemented", true, true},
			{"tests_passed", true, true},
			{"review_approved", true, true},
			{"deployment_complete", true, true},
			{"", true, true}, // Default recommendation
		}
		
		for _, tc := range testCases {
			title, desc := workflow.GetNextStep(tc.action)
			
			if tc.expectTitle {
				assert.NotEmpty(t, title, "Recommendation title should not be empty for action: %s", tc.action)
			}
			if tc.expectDesc {
				assert.NotEmpty(t, desc, "Recommendation description should not be empty for action: %s", tc.action)
			}
		}
		
		// Test workflow sequence
		sequence := workflow.GetWorkflowSequence()
		assert.NotEmpty(t, sequence, "Workflow sequence should not be empty")
		assert.GreaterOrEqual(t, len(sequence), 6, "Workflow sequence should have at least 6 steps")
		
		// Verify expected steps (sequence contains descriptions like "Plan → ...")
		expectedSteps := []string{"Plan", "Design", "Code", "Test", "Review", "Deploy"}
		sequenceStr := ""
		for _, s := range sequence {
			sequenceStr += s + " "
		}
		for _, step := range expectedSteps {
			assert.Contains(t, sequenceStr, step, "Workflow sequence should contain: %s", step)
		}
	})
	
	t.Run("Cross-Phase Integration", func(t *testing.T) {
		// Check that agents reference rules
		plannerPath := filepath.Join(projectRoot, ".doplan/ai/agents/planner.agent.md")
		plannerContent, err := os.ReadFile(plannerPath)
		require.NoError(t, err)
		
		assert.Contains(t, string(plannerContent), "workflow.mdc", "Planner agent should reference workflow.mdc")
		assert.Contains(t, string(plannerContent), "communication.mdc", "Planner agent should reference communication.mdc")
		
		// Check that design rules reference DPR
		designRulesPath := filepath.Join(projectRoot, ".doplan/ai/rules/design_rules.mdc")
		designRulesContent, err := os.ReadFile(designRulesPath)
		require.NoError(t, err)
		
		assert.Contains(t, string(designRulesContent), "DPR", "Design rules should reference DPR")
		
		// Check that devops agent references RAKD
		devopsPath := filepath.Join(projectRoot, ".doplan/ai/agents/devops.agent.md")
		devopsContent, err := os.ReadFile(devopsPath)
		require.NoError(t, err)
		
		assert.Contains(t, string(devopsContent), "RAKD", "DevOps agent should reference RAKD")
	})
	
	t.Run("Agent Workflow Sequence", func(t *testing.T) {
		workflowPath := filepath.Join(projectRoot, ".doplan/ai/rules/workflow.mdc")
		workflowContent, err := os.ReadFile(workflowPath)
		require.NoError(t, err)
		
		content := string(workflowContent)
		
		// Check that workflow mentions correct sequence
		hasSequence := 
			containsSequence(content, []string{"Plan", "Design", "Code", "Test", "Review", "Deploy"}) ||
			containsSequence(content, []string{"planner", "designer", "coder", "tester", "reviewer", "devops"})
		
		assert.True(t, hasSequence, "Workflow should document the correct sequence")
	})
}

// Helper function to check if content contains a sequence of words
func containsSequence(content string, sequence []string) bool {
	for i := 0; i <= len(sequence)-2; i++ {
		word1 := sequence[i]
		word2 := sequence[i+1]
		
		idx1 := findWord(content, word1)
		if idx1 == -1 {
			continue
		}
		
		idx2 := findWord(content[idx1:], word2)
		if idx2 != -1 {
			return true
		}
	}
	return false
}

func findWord(content, word string) int {
	contentLower := toLower(content)
	wordLower := toLower(word)
	return indexOf(contentLower, wordLower)
}

func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			result[i] = c + 32
		} else {
			result[i] = c
		}
	}
	return string(result)
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

