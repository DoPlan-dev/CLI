package generators

import (
	"os"
	"path/filepath"
)

// TemplatesGenerator generates template files
type TemplatesGenerator struct {
	projectRoot string
}

// NewTemplatesGenerator creates a new templates generator
func NewTemplatesGenerator(projectRoot string) *TemplatesGenerator {
	return &TemplatesGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates all template files
func (g *TemplatesGenerator) Generate() error {
	templatesDir := filepath.Join(g.projectRoot, "doplan", "templates")

	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		return err
	}

	templates := map[string]string{
		"plan-template.md":   getPlanTemplate(),
		"design-template.md": getDesignTemplate(),
		"tasks-template.md":  getTasksTemplate(),
	}

	for filename, content := range templates {
		path := filepath.Join(templatesDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func getPlanTemplate() string {
	return `# Feature Plan: {{.Feature.Name}}

## Overview

{{.Feature.Description}}

## Objectives

{{range .Feature.Objectives}}
- [ ] {{.}}
{{else}}
No objectives defined yet.
{{end}}

## Requirements

{{range .Feature.Requirements}}
- {{.}}
{{else}}
No requirements defined yet.
{{end}}

## Dependencies

{{if .Feature.Dependencies}}
{{range .Feature.Dependencies}}
- {{.}}
{{end}}
{{else}}
No dependencies
{{end}}

## Timeline

- Start: {{default .Feature.StartDate "TBD"}}
- Target Completion: {{default .Feature.TargetDate "TBD"}}
- Duration: {{default .Feature.Duration "TBD"}}

## Progress

{{progressBar .Feature.Progress 20}}

**Status:** {{.Feature.Status}}

{{if hasBranch .Feature}}
**Branch:** {{.Feature.Branch}}
{{end}}

{{if hasPR .Feature}}
**Pull Request:** {{.Feature.PR.URL}}
{{end}}

## Success Criteria

{{range .Feature.Objectives}}
- [ ] {{.}}
{{else}}
Define success criteria based on objectives.
{{end}}

## Notes

Add your notes here...
`
}

func getDesignTemplate() string {
	return `# Feature Design: {{.Feature.Name}}

## Design Overview

{{if .Feature.DesignOverview}}
{{.Feature.DesignOverview}}
{{else}}
Describe the overall design approach for this feature.
{{end}}

## Architecture

{{if .Feature.Architecture}}
{{.Feature.Architecture}}
{{else}}
Describe the architectural decisions and patterns used.
{{end}}

## User Flow

{{if .Feature.UserFlow}}
{{.Feature.UserFlow}}
{{else}}
Describe the user interaction flow.
{{end}}

## Technical Specifications

{{if .Feature.TechnicalSpecs}}
{{.Feature.TechnicalSpecs}}
{{else}}
Add technical specifications here.
{{end}}

## API Endpoints

{{if .Project.apiEndpoints}}
{{.Project.apiEndpoints}}
{{else}}
List API endpoints if applicable:
- GET /api/endpoint
- POST /api/endpoint
{{end}}

## Database Schema

{{if .Project.databaseSchema}}
{{.Project.databaseSchema}}
{{else}}
Describe database changes if applicable.
{{end}}

## Security Considerations

{{if .Project.securityConsiderations}}
{{.Project.securityConsiderations}}
{{else}}
- Authentication requirements
- Authorization checks
- Data validation
- Input sanitization
{{end}}

## Performance Requirements

{{if .Project.performanceRequirements}}
{{.Project.performanceRequirements}}
{{else}}
- Response time targets
- Throughput requirements
- Resource constraints
{{end}}

## Integration Points

{{if .Feature.Dependencies}}
**Dependencies:**
{{range .Feature.Dependencies}}
- {{.}}
{{end}}
{{else}}
No external dependencies.
{{end}}
`
}

func getTasksTemplate() string {
	return `# Feature Tasks: {{.Feature.Name}}

## Task Breakdown

{{if .Feature.TaskPhases}}
{{range .Feature.TaskPhases}}
### {{.Name}}

{{range .Tasks}}
- {{if .Completed}}[x]{{else}}[ ]{{end}} {{.Name}}
{{end}}
{{end}}
{{else}}
### Phase 1: Setup
- [ ] Initialize feature branch
- [ ] Set up development environment
- [ ] Create basic structure

### Phase 2: Implementation
- [ ] Implement core functionality
- [ ] Add error handling
- [ ] Write unit tests

### Phase 3: Testing
- [ ] Run test suite
- [ ] Integration testing
- [ ] Performance testing

### Phase 4: Documentation
- [ ] Update API documentation
- [ ] Write user guide
- [ ] Update README
{{end}}

## Progress

{{progressBar .Feature.Progress 20}} ({{.Feature.Progress}}%)

{{if .Feature.Dependencies}}
## Dependencies

{{range .Feature.Dependencies}}
- {{.}}
{{end}}
{{end}}

## Notes

Add task-specific notes here...
`
}
