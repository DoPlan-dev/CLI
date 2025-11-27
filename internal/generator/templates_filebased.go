package generator

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/DoPlan-dev/CLI/internal/content"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

// LoadTemplate loads a template from embedded files
func LoadTemplate(name string) (*template.Template, error) {
	data, err := content.ReadTemplateFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read template file %s: %w", name, err)
	}

	tmpl, err := template.New(name).Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %w", name, err)
	}

	return tmpl, nil
}

// LoadTemplateFromFile loads a template from a file on disk
func LoadTemplateFromFile(path string) (*template.Template, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read template file %s: %w", path, err)
	}

	name := filepath.Base(path)
	tmpl, err := template.New(name).Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %w", path, err)
	}

	return tmpl, nil
}

// GetAgentTemplate returns the agent markdown template (tries file-based first, falls back to hardcoded)
func GetAgentTemplate() (string, error) {
	// Try to load from embedded files first
	data, err := content.ReadTemplateFile("agents/agent.md.tmpl")
	if err != nil {
		// Fallback to hardcoded template
		return agentTemplate, nil
	}

	return string(data), nil
}

// GetCommandTemplate returns the command markdown template (tries file-based first, falls back to hardcoded)
func GetCommandTemplate() (string, error) {
	// Try to load from embedded files first
	data, err := content.ReadTemplateFile("commands/command.md.tmpl")
	if err != nil {
		// Fallback to hardcoded template
		return commandTemplate, nil
	}

	return string(data), nil
}

// RenderAgentMarkdownFileBased renders an agent using file-based template
func RenderAgentMarkdownFileBased(agent *Agent) (string, error) {
	if agent == nil {
		return "", fmt.Errorf("agent cannot be nil")
	}

	templateStr, err := GetAgentTemplate()
	if err != nil {
		return "", fmt.Errorf("failed to get agent template: %w", err)
	}

	tmpl, err := template.New("agent").Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, agent); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// RenderCommandMarkdownFileBased renders a command using file-based template
func RenderCommandMarkdownFileBased(cmd *Command) (string, error) {
	if cmd == nil {
		return "", fmt.Errorf("command cannot be nil")
	}

	templateStr, err := GetCommandTemplate()
	if err != nil {
		return "", fmt.Errorf("failed to get command template: %w", err)
	}

	tmpl, err := template.New("command").Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cmd); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// LoadBrainstormPhaseTemplate loads a brainstorm phase template by phase number (01-06)
func LoadBrainstormPhaseTemplate(phaseNum string) (string, error) {
	// Map phase numbers to actual filenames
	phaseMap := map[string]string{
		"01": "documents/brainstorm/phase-01-vision.md",
		"02": "documents/brainstorm/phase-02-audience.md",
		"03": "documents/brainstorm/phase-03-experience.md",
		"04": "documents/brainstorm/phase-04-content.md",
		"05": "documents/brainstorm/phase-05-marketing.md",
		"06": "documents/brainstorm/phase-06-delivery.md",
	}

	templatePath, ok := phaseMap[phaseNum]
	if !ok {
		// Fallback to generic pattern
		templatePath = fmt.Sprintf("documents/brainstorm/phase-%s.md", phaseNum)
	}

	data, err := content.ReadTemplateFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to load phase template %s: %w", phaseNum, err)
	}

	return string(data), nil
}

// LoadBrainstormConfirmationTemplate loads the confirmation template
func LoadBrainstormConfirmationTemplate() (string, error) {
	data, err := content.ReadTemplateFile("documents/brainstorm/CONFIRMATION_TEMPLATE.md")
	if err != nil {
		return "", fmt.Errorf("failed to load confirmation template: %w", err)
	}

	return string(data), nil
}

// LoadBrainstormOutputTemplate loads the output template for BRAINSTORM.md
func LoadBrainstormOutputTemplate() (string, error) {
	data, err := content.ReadTemplateFile("documents/brainstorm/TEMPLATE_BRAINSTORM.md")
	if err != nil {
		return "", fmt.Errorf("failed to load brainstorm output template: %w", err)
	}

	return string(data), nil
}

// ExtractTemplates extracts all template files to a target directory
func ExtractTemplates(targetDir string) error {
	if targetDir == "" {
		return fmt.Errorf("target directory cannot be empty")
	}

	err := content.WalkTemplateDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking templates: %w", err)
		}

		// Path from WalkDir includes "templates/" prefix, strip it
		relPath := path
		const templatesPrefix = "templates/"
		if len(path) > len(templatesPrefix) && path[:len(templatesPrefix)] == templatesPrefix {
			relPath = path[len(templatesPrefix):]
		}

		targetPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			if err := utils.CreateDirectory(targetPath); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
		} else {
			fileData, err := content.ReadTemplateFile(relPath)
			if err != nil {
				return fmt.Errorf("failed to read embedded file %s: %w", relPath, err)
			}

			if err := utils.WriteFile(targetPath, fileData); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error extracting templates: %w", err)
	}

	return nil
}
