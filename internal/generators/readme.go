package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/config"
)

// READMEGenerator generates README.md
type READMEGenerator struct {
	projectRoot string
}

// NewREADMEGenerator creates a new README generator
func NewREADMEGenerator(projectRoot string) *READMEGenerator {
	return &READMEGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates the README.md content
func (g *READMEGenerator) Generate() string {
	codeBlock := "```"
	backtick := "`"

	// Get project information
	projectName := g.getProjectName()
	projectDescription := g.getProjectDescription()
	projectFeatures := g.getProjectFeatures()
	techStack := g.getTechStack()
	projectStructure := g.getProjectStructure()

	var sb strings.Builder

	// Project Header
	sb.WriteString(fmt.Sprintf("# %s\n\n", projectName))
	if projectDescription != "" {
		sb.WriteString(fmt.Sprintf("%s\n\n", projectDescription))
	} else {
		sb.WriteString("[Project tagline/description]\n\n")
	}

	// Quick Start
	sb.WriteString("## 🚀 Quick Start\n\n")
	sb.WriteString("[Installation and setup for THIS project]\n\n")
	if g.fileExists("package.json") {
		sb.WriteString(fmt.Sprintf("%sbash\nnpm install\nnpm run dev\n%s\n\n", codeBlock, codeBlock))
	} else if g.fileExists("go.mod") {
		sb.WriteString(fmt.Sprintf("%sbash\ngo mod download\ngo run .\n%s\n\n", codeBlock, codeBlock))
	} else {
		sb.WriteString(fmt.Sprintf("%sbash\n# Add installation instructions here\n%s\n\n", codeBlock, codeBlock))
	}

	// Features
	sb.WriteString("## 📋 Features\n\n")
	if len(projectFeatures) > 0 {
		for _, feature := range projectFeatures {
			sb.WriteString(fmt.Sprintf("- %s\n", feature))
		}
	} else {
		sb.WriteString("[Project features, not DoPlan features]\n")
	}
	sb.WriteString("\n")

	// Tech Stack
	sb.WriteString("## 🛠️ Tech Stack\n\n")
	if techStack != "" {
		sb.WriteString(techStack)
	} else {
		sb.WriteString("[Project technologies with links]\n")
	}
	sb.WriteString("\n")

	// Project Structure
	sb.WriteString("## 📁 Project Structure\n\n")
	if projectStructure != "" {
		sb.WriteString(projectStructure)
	} else {
		sb.WriteString(fmt.Sprintf("%s\n", g.generateDefaultStructure(codeBlock)))
	}
	sb.WriteString("\n")

	// Environment Variables
	sb.WriteString("## 🔑 Environment Variables\n\n")
	if g.fileExists(".doplan/RAKD.md") {
		sb.WriteString(fmt.Sprintf("See [RAKD.md](./doplan/RAKD.md) for required API keys and environment variables.\n\n"))
	} else {
		sb.WriteString("[Required .env variables - link to RAKD.md]\n\n")
	}

	// Documentation
	sb.WriteString("## 📚 Documentation\n\n")
	docs := []struct {
		name string
		path string
	}{
		{"Product Requirements", "doplan/PRD.md"},
		{"API Specification", "doplan/contracts/api-spec.json"},
		{"Design Guidelines", "doplan/design/DPR.md"},
		{"Development Progress", "doplan/dashboard.md"},
	}
	for _, doc := range docs {
		if g.fileExists(doc.path) {
			sb.WriteString(fmt.Sprintf("- [%s](./%s)\n", doc.name, doc.path))
		} else {
			sb.WriteString(fmt.Sprintf("- [%s](./%s) *(to be created)*\n", doc.name, doc.path))
		}
	}
	sb.WriteString("\n")

	// Contributing
	sb.WriteString("## 🤝 Contributing\n\n")
	sb.WriteString("[Project-specific guidelines]\n\n")

	// Separator
	sb.WriteString("---\n\n")

	// DoPlan Resources (collapsible)
	sb.WriteString("<details>\n")
	sb.WriteString("<summary>💼 Powered by DoPlan</summary>\n\n")
	sb.WriteString("This project uses DoPlan for workflow automation and project management.\n\n")
	sb.WriteString("### DoPlan Commands\n\n")
	sb.WriteString(fmt.Sprintf("- %s/Discuss%s - Refine ideas\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- %s/Generate%s - Create docs\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- %s/Plan%s - Structure phases\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- [View all commands](./.doplan/ai/commands/)\n\n"))
	sb.WriteString("### Quick Links\n\n")
	sb.WriteString(fmt.Sprintf("- Dashboard: Run %sdoplan%s or %sdoplan dashboard%s\n", backtick, backtick, backtick, backtick))
	sb.WriteString(fmt.Sprintf("- Progress: Run %sdoplan progress%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- GitHub Sync: Run %sdoplan github%s\n", backtick, backtick))
	sb.WriteString("\n")
	sb.WriteString("For more information, visit [DoPlan Documentation](https://github.com/DoPlan-dev/docs)\n\n")
	sb.WriteString("</details>\n")

	return sb.String()
}

// Helper methods
func (g *READMEGenerator) getProjectName() string {
	cfgMgr := config.NewManager(g.projectRoot)
	state, err := cfgMgr.LoadState()
	if err == nil && state != nil && state.Idea != nil && state.Idea.Name != "" {
		return state.Idea.Name
	}

	// Try config YAML
	configPath := filepath.Join(g.projectRoot, ".doplan", "config.yaml")
	if data, err := os.ReadFile(configPath); err == nil {
		content := string(data)
		if strings.Contains(content, "project:") {
			lines := strings.Split(content, "\n")
			for i, line := range lines {
				if strings.Contains(line, "name:") && i > 0 && strings.Contains(lines[i-1], "project:") {
					parts := strings.Split(line, "name:")
					if len(parts) > 1 {
						name := strings.TrimSpace(strings.Trim(parts[1], `"`))
						if name != "" {
							return name
						}
					}
				}
			}
		}
	}

	// Fallback to directory name
	base := filepath.Base(g.projectRoot)
	if base != "." && base != "/" {
		return base
	}

	return "Project"
}

func (g *READMEGenerator) getProjectDescription() string {
	cfgMgr := config.NewManager(g.projectRoot)
	state, err := cfgMgr.LoadState()
	if err == nil && state != nil && state.Idea != nil && state.Idea.Description != "" {
		return state.Idea.Description
	}
	return ""
}

func (g *READMEGenerator) getProjectFeatures() []string {
	cfgMgr := config.NewManager(g.projectRoot)
	state, err := cfgMgr.LoadState()
	if err == nil && state != nil {
		var features []string
		for _, phase := range state.Phases {
			for _, featureID := range phase.Features {
				for _, feature := range state.Features {
					if feature.ID == featureID {
						features = append(features, feature.Name)
					}
				}
			}
		}
		return features
	}
	return nil
}

func (g *READMEGenerator) getTechStack() string {
	// Detect from package.json or go.mod
	if g.fileExists("package.json") {
		return "- **Frontend:** JavaScript/TypeScript\n- **Package Manager:** npm/yarn/pnpm"
	}
	if g.fileExists("go.mod") {
		return "- **Language:** Go\n- **Module System:** Go Modules"
	}
	return ""
}

func (g *READMEGenerator) getProjectStructure() string {
	// Check if doplan directory exists and has phases
	doplanDir := filepath.Join(g.projectRoot, "doplan")
	if entries, err := os.ReadDir(doplanDir); err == nil {
		var phases []string
		for _, entry := range entries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "01-") || strings.HasPrefix(entry.Name(), "02-") {
				phases = append(phases, entry.Name())
			}
		}
		if len(phases) > 0 {
			return g.generateStructureFromPhases(phases)
		}
	}
	return ""
}

func (g *READMEGenerator) generateStructureFromPhases(phases []string) string {
	codeBlock := "```"
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", codeBlock))
	sb.WriteString("project-root/\n")
	sb.WriteString("├── src/                    # Source code\n")
	sb.WriteString("├── doplan/                 # Project planning\n")
	for i, phase := range phases {
		prefix := "├──"
		if i == len(phases)-1 {
			prefix = "└──"
		}
		sb.WriteString(fmt.Sprintf("%s %s/          # %s\n", prefix, phase, strings.ReplaceAll(phase, "-", " ")))
	}
	sb.WriteString(fmt.Sprintf("%s\n", codeBlock))
	return sb.String()
}

func (g *READMEGenerator) generateDefaultStructure(codeBlock string) string {
	return fmt.Sprintf(`%s
project-root/
├── src/                    # Source code
├── doplan/                 # Project planning (##-phase-name/##-feature-name structure)
│   ├── 01-phase-name/
│   │   └── 01-feature-name/
│   └── 02-phase-name/
│       └── 01-feature-name/
└── README.md
%s`, codeBlock, codeBlock)
}

func (g *READMEGenerator) fileExists(path string) bool {
	fullPath := filepath.Join(g.projectRoot, path)
	_, err := os.Stat(fullPath)
	return err == nil
}
