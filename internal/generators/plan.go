package generators

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/template"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// PlanGenerator generates phase and feature structure
type PlanGenerator struct {
	projectRoot string
	state       *models.State
}

// NewPlanGenerator creates a new plan generator
func NewPlanGenerator(projectRoot string, state *models.State) *PlanGenerator {
	return &PlanGenerator{
		projectRoot: projectRoot,
		state:       state,
	}
}

// Generate creates phase and feature directories with documents
func (g *PlanGenerator) Generate() error {
	// Ensure doplan directory exists even if no phases
	doplanDir := filepath.Join(g.projectRoot, "doplan")
	if err := os.MkdirAll(doplanDir, 0755); err != nil {
		return err
	}

	for i, phase := range g.state.Phases {
		phaseDir := filepath.Join(g.projectRoot, "doplan", fmt.Sprintf("%02d-phase", i+1))

		// Create phase directory
		if err := os.MkdirAll(phaseDir, 0755); err != nil {
			return err
		}

		// Generate phase-plan.md
		if err := g.generatePhasePlan(phaseDir, phase); err != nil {
			return err
		}

		// Generate phase-progress.json
		if err := g.generatePhaseProgress(phaseDir, phase); err != nil {
			return err
		}

		// Generate features
		for j, featureID := range phase.Features {
			feature := g.findFeature(featureID)
			if feature == nil {
				continue
			}

			featureDir := filepath.Join(phaseDir, fmt.Sprintf("%02d-Feature", j+1))
			if err := os.MkdirAll(featureDir, 0755); err != nil {
				return err
			}

			// Generate feature documents
			if err := g.generateFeaturePlan(featureDir, feature); err != nil {
				return err
			}
			if err := g.generateFeatureDesign(featureDir, feature); err != nil {
				return err
			}
			if err := g.generateFeatureTasks(featureDir, feature); err != nil {
				return err
			}
			if err := g.generateFeatureProgress(featureDir, feature); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *PlanGenerator) findFeature(featureID string) *models.Feature {
	for _, feature := range g.state.Features {
		if feature.ID == featureID {
			return &feature
		}
	}
	return nil
}

func (g *PlanGenerator) generatePhasePlan(phaseDir string, phase models.Phase) error {
	content := fmt.Sprintf(`# Phase Plan: %s

## Overview

%s

## Objectives

%s

## Features

%s

## Timeline

- Start: %s
- Target Completion: %s
- Duration: %s

## Status

**Status:** %s

## Notes

%s
`,
		phase.Name,
		phase.Description,
		g.formatList(phase.Objectives),
		g.formatFeatureList(phase.Features),
		g.getOrDefault(phase.StartDate, "TBD"),
		g.getOrDefault(phase.TargetDate, "TBD"),
		g.getOrDefault(phase.Duration, "TBD"),
		phase.Status,
		"",
	)

	path := filepath.Join(phaseDir, "phase-plan.md")
	return os.WriteFile(path, []byte(content), 0644)
}

func (g *PlanGenerator) generatePhaseProgress(phaseDir string, phase models.Phase) error {
	progress := map[string]interface{}{
		"phaseID":   phase.ID,
		"phaseName": phase.Name,
		"status":    phase.Status,
		"progress":  0,
		"features":  len(phase.Features),
	}

	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(phaseDir, "phase-progress.json")
	return os.WriteFile(path, data, 0644)
}

func (g *PlanGenerator) generateFeaturePlan(featureDir string, feature *models.Feature) error {
	// Try to use template first
	tmplMgr := template.NewManager(g.projectRoot)
	cfg, err := tmplMgr.LoadConfig()
	if err == nil {
		templateName := cfg.DefaultPlan
		if templateName == "" {
			templateName = "plan-template.md"
		}

		processor := template.NewProcessor(filepath.Join(g.projectRoot, "doplan", "templates"))
		phase := g.findPhaseForFeature(feature.Phase)

		cfgMgr := config.NewManager(g.projectRoot)
		configData, _ := cfgMgr.LoadConfig()

		data := template.TemplateData{
			Feature: feature,
			Phase:   phase,
			State:   g.state,
			Config:  configData,
			Project: make(map[string]interface{}),
		}

		content, err := processor.ProcessTemplate(templateName, data)
		if err == nil {
			path := filepath.Join(featureDir, "plan.md")
			return os.WriteFile(path, []byte(content), 0644)
		}
		// Fall through to default if template processing fails
	}

	// Fallback to default generation
	content := fmt.Sprintf(`# Feature Plan: %s

## Overview

%s

## Objectives

%s

## Requirements

%s

## Dependencies

%s

## Timeline

- Start: %s
- Target Completion: %s
- Duration: %s

## Success Criteria

%s

## Notes

%s
`,
		feature.Name,
		feature.Description,
		g.formatList(feature.Objectives),
		g.formatList(feature.Requirements),
		g.formatList(feature.Dependencies),
		g.getOrDefault(feature.StartDate, "TBD"),
		g.getOrDefault(feature.TargetDate, "TBD"),
		g.getOrDefault(feature.Duration, "TBD"),
		"",
		"",
	)

	path := filepath.Join(featureDir, "plan.md")
	return os.WriteFile(path, []byte(content), 0644)
}

func (g *PlanGenerator) generateFeatureDesign(featureDir string, feature *models.Feature) error {
	// Try to use template first
	tmplMgr := template.NewManager(g.projectRoot)
	cfg, err := tmplMgr.LoadConfig()
	if err == nil {
		templateName := cfg.DefaultDesign
		if templateName == "" {
			templateName = "design-template.md"
		}

		processor := template.NewProcessor(filepath.Join(g.projectRoot, "doplan", "templates"))
		phase := g.findPhaseForFeature(feature.Phase)

		cfgMgr := config.NewManager(g.projectRoot)
		configData, _ := cfgMgr.LoadConfig()

		data := template.TemplateData{
			Feature: feature,
			Phase:   phase,
			State:   g.state,
			Config:  configData,
			Project: make(map[string]interface{}),
		}

		content, err := processor.ProcessTemplate(templateName, data)
		if err == nil {
			path := filepath.Join(featureDir, "design.md")
			return os.WriteFile(path, []byte(content), 0644)
		}
		// Fall through to default if template processing fails
	}

	// Fallback to default generation
	content := fmt.Sprintf(`# Feature Design: %s

## Design Overview

%s

## Architecture

%s

## User Flow

%s

## Technical Specifications

%s

## API Endpoints

%s

## Database Schema

%s

## Security Considerations

%s

## Performance Requirements

%s
`,
		feature.Name,
		g.getOrDefault(feature.DesignOverview, "To be defined."),
		g.getOrDefault(feature.Architecture, "To be defined."),
		g.getOrDefault(feature.UserFlow, "To be defined."),
		g.getOrDefault(feature.TechnicalSpecs, "To be defined."),
		"To be defined.",
		"To be defined.",
		"To be defined.",
		"To be defined.",
	)

	path := filepath.Join(featureDir, "design.md")
	return os.WriteFile(path, []byte(content), 0644)
}

func (g *PlanGenerator) generateFeatureTasks(featureDir string, feature *models.Feature) error {
	// Try to use template first
	tmplMgr := template.NewManager(g.projectRoot)
	cfg, err := tmplMgr.LoadConfig()
	if err == nil {
		templateName := cfg.DefaultTasks
		if templateName == "" {
			templateName = "tasks-template.md"
		}

		processor := template.NewProcessor(filepath.Join(g.projectRoot, "doplan", "templates"))
		phase := g.findPhaseForFeature(feature.Phase)

		cfgMgr := config.NewManager(g.projectRoot)
		configData, _ := cfgMgr.LoadConfig()

		data := template.TemplateData{
			Feature: feature,
			Phase:   phase,
			State:   g.state,
			Config:  configData,
			Project: make(map[string]interface{}),
		}

		content, err := processor.ProcessTemplate(templateName, data)
		if err == nil {
			path := filepath.Join(featureDir, "tasks.md")
			return os.WriteFile(path, []byte(content), 0644)
		}
		// Fall through to default if template processing fails
	}

	// Fallback to default generation
	content := `# Feature Tasks: ` + feature.Name + `

## Task Breakdown

`

	for _, taskPhase := range feature.TaskPhases {
		content += fmt.Sprintf("### %s\n\n", taskPhase.Name)
		for _, task := range taskPhase.Tasks {
			checked := " "
			if task.Completed {
				checked = "x"
			}
			content += fmt.Sprintf("- [%s] %s\n", checked, task.Name)
		}
		content += "\n"
	}

	content += `## Dependencies

## Notes

`

	path := filepath.Join(featureDir, "tasks.md")
	return os.WriteFile(path, []byte(content), 0644)
}

func (g *PlanGenerator) findPhaseForFeature(phaseID string) *models.Phase {
	for _, phase := range g.state.Phases {
		if phase.ID == phaseID {
			return &phase
		}
	}
	return nil
}

func (g *PlanGenerator) generateFeatureProgress(featureDir string, feature *models.Feature) error {
	progress := map[string]interface{}{
		"featureID":   feature.ID,
		"featureName": feature.Name,
		"status":      feature.Status,
		"progress":    feature.Progress,
		"branch":      feature.Branch,
	}

	if feature.PR != nil {
		progress["pr"] = map[string]interface{}{
			"number": feature.PR.Number,
			"url":    feature.PR.URL,
			"status": feature.PR.Status,
		}
	}

	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(featureDir, "progress.json")
	return os.WriteFile(path, data, 0644)
}

func (g *PlanGenerator) formatList(items []string) string {
	if len(items) == 0 {
		return "- None"
	}
	var result strings.Builder
	for _, item := range items {
		result.WriteString(fmt.Sprintf("- %s\n", item))
	}
	return result.String()
}

func (g *PlanGenerator) formatFeatureList(featureIDs []string) string {
	if len(featureIDs) == 0 {
		return "- None"
	}
	var result strings.Builder
	for i, featureID := range featureIDs {
		feature := g.findFeature(featureID)
		if feature != nil {
			result.WriteString(fmt.Sprintf("%d. %s\n", i+1, feature.Name))
		}
	}
	return result.String()
}

func (g *PlanGenerator) getOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
