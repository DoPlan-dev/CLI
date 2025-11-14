package generators

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRulesGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewRulesGenerator(projectRoot)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
}

func TestRulesGenerator_Generate(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	gen := NewRulesGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	// Check rules directory
	rulesDir := filepath.Join(projectRoot, ".cursor", "rules")
	assert.DirExists(t, rulesDir)

	// Check all rule files
	rules := []string{
		"workflow-rules.md",
		"github-rules.md",
		"command-rules.md",
		"branch-rules.md",
		"commit-rules.md",
		"docs-rules.md",
	}

	for _, rule := range rules {
		rulePath := filepath.Join(rulesDir, rule)
		assert.FileExists(t, rulePath)

		content, err := os.ReadFile(rulePath)
		require.NoError(t, err)
		assert.NotEmpty(t, content)
	}
}

func TestRulesGenerator_Generate_ExistingDirectory(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	rulesDir := filepath.Join(projectRoot, ".cursor", "rules")
	require.NoError(t, os.MkdirAll(rulesDir, 0755))

	gen := NewRulesGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)

	// Should still create files
	assert.FileExists(t, filepath.Join(rulesDir, "workflow-rules.md"))
}

func TestGenerateWorkflowRules(t *testing.T) {
	rules := generateWorkflowRules()
	assert.Contains(t, rules, "DoPlan Workflow Rules")
	assert.Contains(t, rules, "/Discuss")
	assert.Contains(t, rules, "/Refine")
	assert.Contains(t, rules, "/Generate")
	assert.Contains(t, rules, "/Plan")
	assert.Contains(t, rules, "/Implement")
}

func TestGenerateGitHubRules(t *testing.T) {
	rules := generateGitHubRules()
	assert.Contains(t, rules, "GitHub")
	assert.Contains(t, rules, "branch")
	assert.Contains(t, rules, "Pull Request")
}

func TestGenerateCommandRules(t *testing.T) {
	rules := generateCommandRules()
	assert.Contains(t, rules, "DoPlan Command Rules")
	assert.Contains(t, rules, "/Discuss Command")
	assert.Contains(t, rules, "/Refine Command")
	assert.Contains(t, rules, "/Generate Command")
}

func TestGenerateBranchRules(t *testing.T) {
	rules := generateBranchRules()
	assert.Contains(t, rules, "Branch")
	assert.Contains(t, rules, "feature/")
}

func TestGenerateCommitRules(t *testing.T) {
	rules := generateCommitRules()
	assert.Contains(t, rules, "Commit")
	assert.Contains(t, rules, "conventional")
}

func TestGenerateDocsRules(t *testing.T) {
	rules := generateDocsRules()
	assert.Contains(t, rules, "Documentation")
	assert.Contains(t, rules, "README.md")
}
