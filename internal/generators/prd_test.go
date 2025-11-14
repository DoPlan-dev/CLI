package generators

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPRDGenerator_Generate(t *testing.T) {
	tmpDir := t.TempDir()

	state := &models.State{
		Idea: &models.Idea{
			Name:             "Test Project",
			Description:      "A test project",
			ProblemStatement: "Testing problem",
			Solution:         "Testing solution",
			TargetUsers:      []string{"Developers"},
			TechStack:        []string{"Go", "React"},
		},
	}

	gen := NewPRDGenerator(tmpDir, state)

	err := gen.Generate()
	require.NoError(t, err)

	prdPath := filepath.Join(tmpDir, "doplan", "PRD.md")
	assert.FileExists(t, prdPath)

	content, err := os.ReadFile(prdPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Test Project")
	assert.Contains(t, string(content), "Testing problem")
}

func TestPRDGenerator_GenerateEmpty(t *testing.T) {
	tmpDir := t.TempDir()

	state := &models.State{
		Idea: nil,
	}

	gen := NewPRDGenerator(tmpDir, state)

	err := gen.Generate()
	require.NoError(t, err)

	prdPath := filepath.Join(tmpDir, "doplan", "PRD.md")
	assert.FileExists(t, prdPath)

	content, err := os.ReadFile(prdPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Empty - Use /Discuss")
}

func TestPRDGenerator_getOrDefault(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	// Test with value
	result := gen.getOrDefault("test", "default")
	assert.Equal(t, "test", result)

	// Test with empty string
	result = gen.getOrDefault("", "default")
	assert.Equal(t, "default", result)
}

func TestPRDGenerator_formatList(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	// Test with items
	result := gen.formatList([]string{"item1", "item2"}, "empty")
	assert.Contains(t, result, "item1")
	assert.Contains(t, result, "item2")

	// Test with empty list
	result = gen.formatList([]string{}, "empty message")
	assert.Equal(t, "empty message", result)
}

func TestPRDGenerator_formatTechStack(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	// Test with stack
	result := gen.formatTechStack([]string{"Go", "React"})
	assert.Contains(t, result, "Go")
	assert.Contains(t, result, "React")

	// Test with empty stack
	result = gen.formatTechStack([]string{})
	assert.Contains(t, result, "To be determined")
}

func TestPRDGenerator_formatObjectives(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				Name:       "Phase 1",
				Objectives: []string{"Obj 1", "Obj 2"},
			},
		},
	}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatObjectives()
	assert.Contains(t, result, "Phase 1")
	assert.Contains(t, result, "Obj 1")
	assert.Contains(t, result, "Obj 2")

	// Test with no phases
	gen2 := NewPRDGenerator(projectRoot, &models.State{})
	result2 := gen2.formatObjectives()
	assert.Contains(t, result2, "No objectives defined")
}

func TestPRDGenerator_formatPersonas(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatPersonas()
	assert.Contains(t, result, "To be defined")
}

func TestPRDGenerator_formatFeatures(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Features: []models.Feature{
			{
				Name:        "Feature 1",
				Description: "Test feature",
				Objectives:  []string{"Obj 1"},
			},
		},
	}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatFeatures()
	assert.Contains(t, result, "Feature 1")
	assert.Contains(t, result, "Test feature")
	assert.Contains(t, result, "Obj 1")

	// Test with no features
	gen2 := NewPRDGenerator(projectRoot, &models.State{})
	result2 := gen2.formatFeatures()
	assert.Contains(t, result2, "No features defined")
}

func TestPRDGenerator_formatUserFlows(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Features: []models.Feature{
			{
				Name:     "Feature 1",
				UserFlow: "Test user flow",
			},
		},
	}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatUserFlows()
	assert.Contains(t, result, "Feature 1")
	assert.Contains(t, result, "Test user flow")

	// Test with no user flows
	gen2 := NewPRDGenerator(projectRoot, &models.State{
		Features: []models.Feature{
			{Name: "Feature 1"},
		},
	})
	result2 := gen2.formatUserFlows()
	assert.Contains(t, result2, "No user flows defined")
}

func TestPRDGenerator_formatTimeline(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				Name:       "Phase 1",
				StartDate:  "2025-01-01",
				TargetDate: "2025-01-31",
				Duration:   "1 month",
			},
		},
	}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatTimeline()
	assert.Contains(t, result, "Phase 1")
	assert.Contains(t, result, "2025-01-01")
	assert.Contains(t, result, "2025-01-31")
	assert.Contains(t, result, "1 month")

	// Test with no phases
	gen2 := NewPRDGenerator(projectRoot, &models.State{})
	result2 := gen2.formatTimeline()
	assert.Contains(t, result2, "No timeline defined")
}

func TestPRDGenerator_formatSuccessMetrics(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatSuccessMetrics()
	assert.Contains(t, result, "To be defined")
}

func TestPRDGenerator_formatRisks(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatRisks()
	assert.Contains(t, result, "To be defined")
}

func TestPRDGenerator_formatArchitecture(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatArchitecture()
	assert.Contains(t, result, "To be defined")
}

func TestPRDGenerator_formatInfrastructure(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPRDGenerator(projectRoot, state)

	result := gen.formatInfrastructure()
	assert.Contains(t, result, "To be defined")
}
