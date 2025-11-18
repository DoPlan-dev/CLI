package rakd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Generator generates RAKD.md document
type Generator struct {
	projectRoot string
	data        *RAKDData
}

// NewGenerator creates a new RAKD generator
func NewGenerator(projectRoot string, data *RAKDData) *Generator {
	return &Generator{
		projectRoot: projectRoot,
		data:        data,
	}
}

// Generate creates the RAKD.md file
func (g *Generator) Generate() error {
	rakdPath := filepath.Join(g.projectRoot, "doplan", "RAKD.md")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(rakdPath), 0755); err != nil {
		return fmt.Errorf("failed to create doplan directory: %w", err)
	}

	content := g.generateRAKDContent()

	return os.WriteFile(rakdPath, []byte(content), 0644)
}

func (g *Generator) generateRAKDContent() string {
	var content strings.Builder

	// Header
	content.WriteString("# Required API Keys Document (RAKD)\n\n")
	content.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", g.data.LastUpdated))
	content.WriteString("This document tracks all API keys and service credentials required for this project.\n\n")
	content.WriteString("---\n\n")

	// Quick Status
	content.WriteString("## Quick Status\n\n")
	content.WriteString(g.generateQuickStatus())
	content.WriteString("\n\n")

	// Configured Services
	configured := g.getServicesByStatus(StatusConfigured)
	if len(configured) > 0 {
		content.WriteString("## ✅ Configured Services\n\n")
		content.WriteString(g.generateServiceSection(configured))
		content.WriteString("\n\n")
	}

	// Required (Missing) Services
	required := g.getServicesByStatus(StatusRequired)
	if len(required) > 0 {
		content.WriteString("## 🔴 Required (Missing)\n\n")
		content.WriteString("**⚠️ These services are required but not configured. Please set them up immediately.**\n\n")
		content.WriteString(g.generateServiceSection(required))
		content.WriteString("\n\n")
	}

	// Pending Services
	pending := g.getServicesByStatus(StatusPending)
	if len(pending) > 0 {
		content.WriteString("## 🟡 Pending Configuration\n\n")
		content.WriteString(g.generateServiceSection(pending))
		content.WriteString("\n\n")
	}

	// Optional Services
	optional := g.getServicesByStatus(StatusOptional)
	if len(optional) > 0 {
		content.WriteString("## 🔵 Optional Services\n\n")
		content.WriteString("These services are optional and can be configured if needed.\n\n")
		content.WriteString(g.generateServiceSection(optional))
		content.WriteString("\n\n")
	}

	// Invalid Services
	invalid := g.getServicesWithInvalidKeys()
	if len(invalid) > 0 {
		content.WriteString("## ⚠️ Invalid Configuration\n\n")
		content.WriteString("These services have keys configured but they appear to be invalid.\n\n")
		content.WriteString(g.generateServiceSection(invalid))
		content.WriteString("\n\n")
	}

	// Validation Results
	content.WriteString("## Validation Results\n\n")
	content.WriteString(g.generateValidationResults())
	content.WriteString("\n\n")

	// Quick Actions
	content.WriteString("## Quick Actions\n\n")
	content.WriteString(g.generateQuickActions())
	content.WriteString("\n\n")

	// Environment Variables Reference
	content.WriteString("## Environment Variables Reference\n\n")
	content.WriteString(g.generateEnvVarReference())
	content.WriteString("\n\n")

	return content.String()
}

func (g *Generator) generateQuickStatus() string {
	total := len(g.data.Services)
	progress := 0.0
	if total > 0 {
		progress = float64(g.data.ConfiguredCount) / float64(total) * 100
	}

	return fmt.Sprintf(`| Status | Count | Percentage |
|--------|-------|------------|
| ✅ Configured | %d | %.1f%% |
| 🔴 Required (Missing) | %d | %.1f%% |
| 🟡 Pending | %d | %.1f%% |
| 🔵 Optional | %d | %.1f%% |

**Overall Progress:** %.1f%% configured

`,
		g.data.ConfiguredCount, float64(g.data.ConfiguredCount)/float64(total)*100,
		g.data.RequiredCount, float64(g.data.RequiredCount)/float64(total)*100,
		g.data.PendingCount, float64(g.data.PendingCount)/float64(total)*100,
		g.data.OptionalCount, float64(g.data.OptionalCount)/float64(total)*100,
		progress)
}

func (g *Generator) generateServiceSection(services []Service) string {
	var content strings.Builder

	for _, service := range services {
		statusIcon := g.getStatusIcon(service.Status)
		detectedBadge := ""
		if service.Detected {
			detectedBadge = " 🔍 *Auto-detected*"
		}

		content.WriteString(fmt.Sprintf("### %s %s%s\n\n", statusIcon, service.Name, detectedBadge))
		content.WriteString(fmt.Sprintf("**Category:** %s  \n", strings.Title(service.Category)))
		content.WriteString(fmt.Sprintf("**Priority:** %s  \n", strings.Title(service.Priority)))
		content.WriteString(fmt.Sprintf("**Description:** %s\n\n", service.Description))

		// Keys table
		content.WriteString("| Key Name | Environment Variable | Status | Format |\n")
		content.WriteString("|----------|---------------------|--------|--------|\n")
		for _, key := range service.Keys {
			keyStatusIcon := g.getKeyStatusIcon(key.Status)
			format := key.Format
			if format == "" {
				format = "-"
			}
			content.WriteString(fmt.Sprintf("| %s | `%s` | %s %s | %s |\n",
				key.Name, key.EnvVar, keyStatusIcon, key.Status, format))
		}
		content.WriteString("\n")

		// Show errors if any
		for _, key := range service.Keys {
			if key.Error != "" {
				content.WriteString(fmt.Sprintf("⚠️ **%s:** %s\n\n", key.Name, key.Error))
			}
		}

		content.WriteString("---\n\n")
	}

	return content.String()
}

func (g *Generator) generateValidationResults() string {
	var content strings.Builder

	totalKeys := 0
	validatedKeys := 0
	invalidKeys := 0
	missingKeys := 0

	for _, service := range g.data.Services {
		for _, key := range service.Keys {
			totalKeys++
			switch key.Status {
			case StatusConfigured:
				validatedKeys++
			case StatusInvalid:
				invalidKeys++
			case StatusRequired, StatusPending:
				if key.Required {
					missingKeys++
				}
			}
		}
	}

	content.WriteString(fmt.Sprintf("- **Total Keys:** %d\n", totalKeys))
	content.WriteString(fmt.Sprintf("- **✅ Validated:** %d\n", validatedKeys))
	content.WriteString(fmt.Sprintf("- **⚠️ Invalid:** %d\n", invalidKeys))
	content.WriteString(fmt.Sprintf("- **❌ Missing:** %d\n", missingKeys))

	if invalidKeys > 0 || missingKeys > 0 {
		content.WriteString("\n**⚠️ Action Required:** Please fix invalid or missing keys.\n")
	} else {
		content.WriteString("\n**✅ All keys are properly configured!**\n")
	}

	return content.String()
}

func (g *Generator) generateQuickActions() string {
	return `### Setup Missing Keys
1. Review the "Required (Missing)" section above
2. Follow the setup guides in .doplan/SOPS/ for each service
3. Add keys to your .env file
4. Run doplan keys validate to verify

### Validate All Keys
` + "```bash" + `
doplan keys validate
` + "```" + `

### Sync .env.example
` + "```bash" + `
doplan keys sync-env-example
` + "```" + `

### Test API Connections
` + "```bash" + `
doplan keys test
` + "```" + `

### Manage Keys (TUI)
` + "```bash" + `
doplan keys
` + "```" + `
`
}

func (g *Generator) generateEnvVarReference() string {
	var content strings.Builder

	envVars := make(map[string]APIKey)
	for _, service := range g.data.Services {
		for _, key := range service.Keys {
			envVars[key.EnvVar] = key
		}
	}

	content.WriteString("| Environment Variable | Service | Required | Description |\n")
	content.WriteString("|---------------------|---------|----------|-------------|\n")

	for envVar, key := range envVars {
		required := "No"
		if key.Required {
			required = "Yes"
		}
		content.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s |\n",
			envVar, key.Name, required, key.Description))
	}

	return content.String()
}

func (g *Generator) getServicesByStatus(status APIKeyStatus) []Service {
	var result []Service
	for _, service := range g.data.Services {
		if service.Status == status {
			result = append(result, service)
		}
	}
	return result
}

func (g *Generator) getServicesWithInvalidKeys() []Service {
	var result []Service
	for _, service := range g.data.Services {
		for _, key := range service.Keys {
			if key.Status == StatusInvalid {
				result = append(result, service)
				break
			}
		}
	}
	return result
}

func (g *Generator) getStatusIcon(status APIKeyStatus) string {
	switch status {
	case StatusConfigured:
		return "✅"
	case StatusRequired:
		return "🔴"
	case StatusPending:
		return "🟡"
	case StatusOptional:
		return "🔵"
	case StatusInvalid:
		return "⚠️"
	default:
		return "⚪"
	}
}

func (g *Generator) getKeyStatusIcon(status APIKeyStatus) string {
	return g.getStatusIcon(status)
}

// GenerateRAKD generates RAKD.md from detected services
func GenerateRAKD(projectRoot string) (*RAKDData, error) {
	detector := NewDetector(projectRoot)
	services, err := detector.DetectServices()
	if err != nil {
		return nil, fmt.Errorf("failed to detect services: %w", err)
	}

	validator := NewValidator(projectRoot)
	if err := validator.ValidateAll(services); err != nil {
		return nil, fmt.Errorf("failed to validate services: %w", err)
	}

	// Count by status
	data := &RAKDData{
		Services:    services,
		LastUpdated: time.Now().Format("January 2, 2006 at 3:04 PM"),
	}

	for _, service := range services {
		switch service.Status {
		case StatusConfigured:
			data.ConfiguredCount++
		case StatusPending:
			data.PendingCount++
		case StatusRequired:
			data.RequiredCount++
		case StatusOptional:
			data.OptionalCount++
		}
	}

	// Generate RAKD.md
	generator := NewGenerator(projectRoot, data)
	if err := generator.Generate(); err != nil {
		return nil, fmt.Errorf("failed to generate RAKD.md: %w", err)
	}

	return data, nil
}
