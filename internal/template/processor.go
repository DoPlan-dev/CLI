package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// Processor handles template processing
type Processor struct {
	templatesDir string
	templates    map[string]*template.Template
}

// NewProcessor creates a new template processor
func NewProcessor(templatesDir string) *Processor {
	return &Processor{
		templatesDir: templatesDir,
		templates:    make(map[string]*template.Template),
	}
}

// TemplateData provides data to templates
type TemplateData struct {
	Feature *models.Feature
	Phase   *models.Phase
	State   *models.State
	Config  *models.Config
	Project map[string]interface{}
}

// LoadTemplate loads a template from file
func (p *Processor) LoadTemplate(name string) error {
	templatePath := filepath.Join(p.templatesDir, name)

	content, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template: %w", err)
	}

	tmpl, err := template.New(name).Funcs(p.getFuncMap()).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	p.templates[name] = tmpl
	return nil
}

// ProcessTemplate processes a template with data
func (p *Processor) ProcessTemplate(templateName string, data TemplateData) (string, error) {
	tmpl, exists := p.templates[templateName]
	if !exists {
		// Try to load it
		if err := p.LoadTemplate(templateName); err != nil {
			return "", fmt.Errorf("template not found: %s", templateName)
		}
		tmpl = p.templates[templateName]
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// getFuncMap provides template functions
func (p *Processor) getFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatList": func(items []string) string {
			if len(items) == 0 {
				return "None"
			}
			var result string
			for _, item := range items {
				result += fmt.Sprintf("- %s\n", item)
			}
			return result
		},
		"formatChecklist": func(items []string) string {
			if len(items) == 0 {
				return "None"
			}
			var result string
			for _, item := range items {
				result += fmt.Sprintf("- [ ] %s\n", item)
			}
			return result
		},
		"progressBar": func(progress int, width int) string {
			filled := (progress * width) / 100
			if filled > width {
				filled = width
			}
			empty := width - filled
			return fmt.Sprintf("[%s%s] %d%%",
				repeatString("█", filled),
				repeatString("░", empty),
				progress)
		},
		"default": func(value, defaultValue string) string {
			if value == "" {
				return defaultValue
			}
			return value
		},
		"hasBranch": func(feature *models.Feature) bool {
			return feature != nil && feature.Branch != ""
		},
		"hasPR": func(feature *models.Feature) bool {
			return feature != nil && feature.PR != nil && feature.PR.URL != ""
		},
		"isEmpty": func(s string) bool {
			return s == ""
		},
		"join": func(items []string, sep string) string {
			result := ""
			for i, item := range items {
				if i > 0 {
					result += sep
				}
				result += item
			}
			return result
		},
	}
}

func repeatString(s string, count int) string {
	if count <= 0 {
		return ""
	}
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

