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

func TestNewPlanGenerator(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
	}

	gen := NewPlanGenerator(projectRoot, state)
	assert.NotNil(t, gen)
	assert.Equal(t, projectRoot, gen.projectRoot)
	assert.Equal(t, state, gen.state)
}

func TestPlanGenerator_Generate_EmptyState(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases:   []models.Phase{},
		Features: []models.Feature{},
	}

	gen := NewPlanGenerator(projectRoot, state)
	err := gen.Generate()
	require.NoError(t, err)

	// Should create doplan directory
	doplanDir := filepath.Join(projectRoot, "doplan")
	assert.DirExists(t, doplanDir)
}

func TestPlanGenerator_Generate_WithPhases(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:          "phase-1",
				Name:        "Phase 1",
				Description: "First phase",
				Objectives:  []string{"Objective 1", "Objective 2"},
				Features:    []string{"feat-1"},
			},
		},
		Features: []models.Feature{
			{
				ID:          "feat-1",
				Name:        "Feature 1",
				Description: "First feature",
				Objectives:  []string{"Feat Obj 1"},
			},
		},
	}

	gen := NewPlanGenerator(projectRoot, state)
	err := gen.Generate()
	require.NoError(t, err)

	// Check phase directory
	phaseDir := filepath.Join(projectRoot, "doplan", "01-phase")
	assert.DirExists(t, phaseDir)

	// Check phase-plan.md
	phasePlanPath := filepath.Join(phaseDir, "phase-plan.md")
	assert.FileExists(t, phasePlanPath)

	content, err := os.ReadFile(phasePlanPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Phase 1")
	assert.Contains(t, string(content), "First phase")

	// Check phase-progress.json
	phaseProgressPath := filepath.Join(phaseDir, "phase-progress.json")
	assert.FileExists(t, phaseProgressPath)

	// Check feature directory
	featureDir := filepath.Join(phaseDir, "01-Feature")
	assert.DirExists(t, featureDir)

	// Check feature files
	assert.FileExists(t, filepath.Join(featureDir, "plan.md"))
	assert.FileExists(t, filepath.Join(featureDir, "design.md"))
	assert.FileExists(t, filepath.Join(featureDir, "tasks.md"))
	assert.FileExists(t, filepath.Join(featureDir, "progress.json"))
}

func TestPlanGenerator_Generate_MultiplePhases(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:       "phase-1",
				Name:     "Phase 1",
				Features: []string{"feat-1"},
			},
			{
				ID:       "phase-2",
				Name:     "Phase 2",
				Features: []string{"feat-2"},
			},
		},
		Features: []models.Feature{
			{ID: "feat-1", Name: "Feature 1"},
			{ID: "feat-2", Name: "Feature 2"},
		},
	}

	gen := NewPlanGenerator(projectRoot, state)
	err := gen.Generate()
	require.NoError(t, err)

	// Check both phase directories
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "01-phase"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "02-phase"))
}

func TestPlanGenerator_Generate_MissingFeature(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Phases: []models.Phase{
			{
				ID:       "phase-1",
				Name:     "Phase 1",
				Features: []string{"nonexistent-feat"},
			},
		},
		Features: []models.Feature{},
	}

	gen := NewPlanGenerator(projectRoot, state)
	err := gen.Generate()
	// Should not error, just skip missing features
	require.NoError(t, err)
}

func TestPlanGenerator_findFeature(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{
		Features: []models.Feature{
			{ID: "feat-1", Name: "Feature 1"},
			{ID: "feat-2", Name: "Feature 2"},
		},
	}

	gen := NewPlanGenerator(projectRoot, state)

	feature := gen.findFeature("feat-1")
	assert.NotNil(t, feature)
	assert.Equal(t, "Feature 1", feature.Name)

	feature = gen.findFeature("nonexistent")
	assert.Nil(t, feature)
}

func TestPlanGenerator_generatePhasePlan(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPlanGenerator(projectRoot, state)

	phaseDir := filepath.Join(projectRoot, "phase-test")
	require.NoError(t, os.MkdirAll(phaseDir, 0755))

	phase := models.Phase{
		Name:        "Test Phase",
		Description: "Test Description",
		Objectives:  []string{"Obj 1", "Obj 2"},
	}

	err := gen.generatePhasePlan(phaseDir, phase)
	require.NoError(t, err)

	planPath := filepath.Join(phaseDir, "phase-plan.md")
	assert.FileExists(t, planPath)

	content, err := os.ReadFile(planPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Test Phase")
	assert.Contains(t, string(content), "Test Description")
	assert.Contains(t, string(content), "Obj 1")
}

func TestPlanGenerator_generatePhaseProgress(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPlanGenerator(projectRoot, state)

	phaseDir := filepath.Join(projectRoot, "phase-test")
	require.NoError(t, os.MkdirAll(phaseDir, 0755))

	phase := models.Phase{
		ID:     "phase-1",
		Name:   "Test Phase",
		Status: "in-progress",
	}

	err := gen.generatePhaseProgress(phaseDir, phase)
	require.NoError(t, err)

	progressPath := filepath.Join(phaseDir, "phase-progress.json")
	assert.FileExists(t, progressPath)

	// Verify JSON is valid
	data, err := os.ReadFile(progressPath)
	require.NoError(t, err)
	assert.Contains(t, string(data), "phase-1")
	assert.Contains(t, string(data), "Test Phase")
	assert.Contains(t, string(data), "in-progress")
	// Progress is always set to 0 in generatePhaseProgress
	assert.Contains(t, string(data), "\"progress\": 0")
}

func TestPlanGenerator_generateFeaturePlan(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPlanGenerator(projectRoot, state)

	featureDir := filepath.Join(projectRoot, "feature-test")
	require.NoError(t, os.MkdirAll(featureDir, 0755))

	feature := models.Feature{
		Name:        "Test Feature",
		Description: "Test Description",
		Objectives:  []string{"Obj 1"},
	}

	err := gen.generateFeaturePlan(featureDir, &feature)
	require.NoError(t, err)

	planPath := filepath.Join(featureDir, "plan.md")
	assert.FileExists(t, planPath)

	content, err := os.ReadFile(planPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Test Feature")
	assert.Contains(t, string(content), "Test Description")
}

func TestPlanGenerator_generateFeatureDesign(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPlanGenerator(projectRoot, state)

	featureDir := filepath.Join(projectRoot, "feature-test")
	require.NoError(t, os.MkdirAll(featureDir, 0755))

	feature := models.Feature{
		Name:        "Test Feature",
		Description: "Test Description",
	}

	err := gen.generateFeatureDesign(featureDir, &feature)
	require.NoError(t, err)

	designPath := filepath.Join(featureDir, "design.md")
	assert.FileExists(t, designPath)
}

func TestPlanGenerator_generateFeatureTasks(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPlanGenerator(projectRoot, state)

	featureDir := filepath.Join(projectRoot, "feature-test")
	require.NoError(t, os.MkdirAll(featureDir, 0755))

	feature := models.Feature{
		Name: "Test Feature",
		TaskPhases: []models.TaskPhase{
			{
				Name: "Phase 1",
				Tasks: []models.Task{
					{Name: "Task 1", Completed: false},
				},
			},
		},
	}

	err := gen.generateFeatureTasks(featureDir, &feature)
	require.NoError(t, err)

	tasksPath := filepath.Join(featureDir, "tasks.md")
	assert.FileExists(t, tasksPath)

	content, err := os.ReadFile(tasksPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Task 1")
}

func TestPlanGenerator_generateFeatureProgress(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	state := &models.State{}
	gen := NewPlanGenerator(projectRoot, state)

	featureDir := filepath.Join(projectRoot, "feature-test")
	require.NoError(t, os.MkdirAll(featureDir, 0755))

	feature := models.Feature{
		ID:       "feat-1",
		Name:     "Test Feature",
		Progress: 75,
	}

	err := gen.generateFeatureProgress(featureDir, &feature)
	require.NoError(t, err)

	progressPath := filepath.Join(featureDir, "progress.json")
	assert.FileExists(t, progressPath)

	data, err := os.ReadFile(progressPath)
	require.NoError(t, err)
	assert.Contains(t, string(data), "feat-1")
	assert.Contains(t, string(data), "75")
}
