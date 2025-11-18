package generators

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/config"
)

// ContextGenerator generates the tech stack context document
type ContextGenerator struct {
	projectRoot string
}

// NewContextGenerator creates a new context generator
func NewContextGenerator(projectRoot string) *ContextGenerator {
	return &ContextGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates the CONTEXT.md file with tech stack information
func (g *ContextGenerator) Generate() error {
	contextPath := filepath.Join(g.projectRoot, "CONTEXT.md")

	// Detect technologies from project files
	techStack := g.detectTechStack()

	// Get project name from state or config
	projectName := g.getProjectName()

	// Generate context document
	content := g.generateContextContent(techStack, projectName)

	return os.WriteFile(contextPath, []byte(content), 0644)
}

// getProjectName retrieves project name from state or config
func (g *ContextGenerator) getProjectName() string {
	// Try to get from state first
	cfgMgr := config.NewManager(g.projectRoot)
	state, err := cfgMgr.LoadState()
	if err == nil && state != nil && state.Idea != nil && state.Idea.Name != "" {
		return state.Idea.Name
	}

	// Try to get from config YAML
	cfg, err := cfgMgr.LoadConfig()
	if err == nil && cfg != nil {
		// Check YAML config for project name
		configPath := filepath.Join(g.projectRoot, ".doplan", "config.yaml")
		if data, err := os.ReadFile(configPath); err == nil {
			// Simple YAML parsing for project.name
			content := string(data)
			if strings.Contains(content, "project:") {
				lines := strings.Split(content, "\n")
				for i, line := range lines {
					if strings.Contains(line, "name:") && i > 0 && strings.Contains(lines[i-1], "project:") {
						parts := strings.Split(line, "name:")
						if len(parts) > 1 {
							name := strings.TrimSpace(parts[1])
							if name != "" && name != `""` {
								return name
							}
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

// TechStack represents the detected technology stack
type TechStack struct {
	Languages  []Technology
	Frameworks []Technology
	CLIs       []Technology
	Services   []Technology
	Databases  []Technology
	Tools      []Technology
}

// Technology represents a technology with its documentation
type Technology struct {
	Name        string
	Version     string
	DocsURL     string
	Description string
	Usage       string
}

// detectTechStack detects technologies from project files
func (g *ContextGenerator) detectTechStack() *TechStack {
	stack := &TechStack{
		Languages:  []Technology{},
		Frameworks: []Technology{},
		CLIs:       []Technology{},
		Services:   []Technology{},
		Databases:  []Technology{},
		Tools:      []Technology{},
	}

	// Detect Go
	if g.fileExists("go.mod") {
		version := g.extractGoVersion()
		stack.Languages = append(stack.Languages, Technology{
			Name:        "Go",
			Version:     version,
			DocsURL:     "https://go.dev/doc/",
			Description: "Go programming language",
			Usage:       "Backend CLI development",
		})
	}

	// Detect Node.js
	if g.fileExists("package.json") {
		version := g.extractNodeVersion()
		stack.Languages = append(stack.Languages, Technology{
			Name:        "JavaScript/TypeScript",
			Version:     version,
			DocsURL:     "https://developer.mozilla.org/en-US/docs/Web/JavaScript",
			Description: "JavaScript or TypeScript",
			Usage:       "Frontend or backend development",
		})
	}

	// Detect Python
	if g.fileExists("requirements.txt") || g.fileExists("pyproject.toml") {
		stack.Languages = append(stack.Languages, Technology{
			Name:        "Python",
			Version:     "3.x",
			DocsURL:     "https://docs.python.org/3/",
			Description: "Python programming language",
			Usage:       "Backend or scripting",
		})
	}

	// Detect Rust
	if g.fileExists("Cargo.toml") {
		stack.Languages = append(stack.Languages, Technology{
			Name:        "Rust",
			Version:     "latest",
			DocsURL:     "https://doc.rust-lang.org/",
			Description: "Rust programming language",
			Usage:       "Systems programming",
		})
	}

	// Detect Git
	if g.fileExists(".git/config") {
		stack.Tools = append(stack.Tools, Technology{
			Name:        "Git",
			Version:     "latest",
			DocsURL:     "https://git-scm.com/doc",
			Description: "Version control system",
			Usage:       "Source code management",
		})
	}

	// Detect GitHub
	if g.fileExists(".github") {
		stack.Services = append(stack.Services, Technology{
			Name:        "GitHub",
			Version:     "latest",
			DocsURL:     "https://docs.github.com/",
			Description: "Git hosting and collaboration platform",
			Usage:       "Version control, CI/CD, project management",
		})
	}

	// Detect Docker
	if g.fileExists("Dockerfile") || g.fileExists("docker-compose.yml") {
		stack.Tools = append(stack.Tools, Technology{
			Name:        "Docker",
			Version:     "latest",
			DocsURL:     "https://docs.docker.com/",
			Description: "Containerization platform",
			Usage:       "Application containerization",
		})
	}

	// Detect SQLite (common for local dev)
	if g.fileExists("*.db") || strings.Contains(g.projectRoot, "sqlite") {
		stack.Databases = append(stack.Databases, Technology{
			Name:        "SQLite",
			Version:     "3.x",
			DocsURL:     "https://www.sqlite.org/docs.html",
			Description: "Lightweight SQL database",
			Usage:       "Local development database",
		})
	}

	// Detect PostgreSQL (common indicators)
	if g.fileExists("docker-compose.yml") {
		content, _ := os.ReadFile(filepath.Join(g.projectRoot, "docker-compose.yml"))
		if strings.Contains(string(content), "postgres") {
			stack.Databases = append(stack.Databases, Technology{
				Name:        "PostgreSQL",
				Version:     "latest",
				DocsURL:     "https://www.postgresql.org/docs/",
				Description: "Open-source relational database",
				Usage:       "Production database",
			})
		}
	}

	// Detect MySQL
	if g.fileExists("docker-compose.yml") {
		content, _ := os.ReadFile(filepath.Join(g.projectRoot, "docker-compose.yml"))
		if strings.Contains(string(content), "mysql") {
			stack.Databases = append(stack.Databases, Technology{
				Name:        "MySQL",
				Version:     "latest",
				DocsURL:     "https://dev.mysql.com/doc/",
				Description: "Open-source relational database",
				Usage:       "Production database",
			})
		}
	}

	// Detect MongoDB
	if g.fileExists("docker-compose.yml") {
		content, _ := os.ReadFile(filepath.Join(g.projectRoot, "docker-compose.yml"))
		if strings.Contains(string(content), "mongo") {
			stack.Databases = append(stack.Databases, Technology{
				Name:        "MongoDB",
				Version:     "latest",
				DocsURL:     "https://docs.mongodb.com/",
				Description: "NoSQL document database",
				Usage:       "Document storage",
			})
		}
	}

	// Add DoPlan CLI
	stack.CLIs = append(stack.CLIs, Technology{
		Name:        "DoPlan CLI",
		Version:     "1.0.0",
		DocsURL:     "https://github.com/DoPlan-dev/docs",
		Description: "Development workflow automation CLI",
		Usage:       "Project planning and workflow management",
	})

	// Detect IDE-specific tools
	if g.fileExists(".cursor") {
		stack.Tools = append(stack.Tools, Technology{
			Name:        "Cursor IDE",
			Version:     "latest",
			DocsURL:     "https://cursor.com/docs",
			Description: "AI-powered code editor",
			Usage:       "Primary development environment",
		})
	}

	return stack
}

// fileExists checks if a file exists
func (g *ContextGenerator) fileExists(filename string) bool {
	path := filepath.Join(g.projectRoot, filename)
	_, err := os.Stat(path)
	return err == nil
}

// extractGoVersion extracts Go version from go.mod
func (g *ContextGenerator) extractGoVersion() string {
	path := filepath.Join(g.projectRoot, "go.mod")
	content, err := os.ReadFile(path)
	if err != nil {
		return "latest"
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "go ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1]
			}
		}
	}
	return "latest"
}

// extractNodeVersion extracts Node version from package.json
func (g *ContextGenerator) extractNodeVersion() string {
	path := filepath.Join(g.projectRoot, "package.json")
	content, err := os.ReadFile(path)
	if err != nil {
		return "latest"
	}

	var pkg map[string]interface{}
	if err := json.Unmarshal(content, &pkg); err != nil {
		return "latest"
	}

	if engines, ok := pkg["engines"].(map[string]interface{}); ok {
		if node, ok := engines["node"].(string); ok {
			return node
		}
	}
	return "latest"
}

// generateContextContent generates the CONTEXT.md content
func (g *ContextGenerator) generateContextContent(stack *TechStack, projectName string) string {
	backtick := "`"
	var sb strings.Builder

	// Header with project name
	sb.WriteString(fmt.Sprintf("# Project Context: %s\n\n", projectName))

	// Project Overview section
	sb.WriteString("## Project Overview\n\n")
	cfgMgr := config.NewManager(g.projectRoot)
	state, _ := cfgMgr.LoadState()
	hasContent := false
	if state != nil && state.Idea != nil {
		if state.Idea.Description != "" {
			sb.WriteString(fmt.Sprintf("- **Brief description:** %s\n", state.Idea.Description))
			hasContent = true
		}
		if len(state.Idea.TargetUsers) > 0 {
			sb.WriteString(fmt.Sprintf("- **Target audience:** %s\n", strings.Join(state.Idea.TargetUsers, ", ")))
			hasContent = true
		}
		if state.Idea.Solution != "" {
			sb.WriteString(fmt.Sprintf("- **Core features:** %s\n", state.Idea.Solution))
			hasContent = true
		}
	}
	if !hasContent {
		sb.WriteString("- Brief description: [To be filled]\n")
		sb.WriteString("- Target audience: [To be filled]\n")
		sb.WriteString("- Core features: [To be filled]\n")
	}
	sb.WriteString("\n")

	// Technology Stack section
	sb.WriteString("## Technology Stack\n\n")

	// Frontend
	frontendTechs := g.categorizeFrontend(stack)
	if len(frontendTechs) > 0 {
		sb.WriteString("### Frontend\n\n")
		for _, tech := range frontendTechs {
			sb.WriteString(fmt.Sprintf("- **%s:** %s", tech.Name, tech.DocsURL))
			if tech.Version != "" && tech.Version != "latest" {
				sb.WriteString(fmt.Sprintf(" (v%s)", tech.Version))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	// Backend
	backendTechs := g.categorizeBackend(stack)
	if len(backendTechs) > 0 {
		sb.WriteString("### Backend\n\n")
		for _, tech := range backendTechs {
			sb.WriteString(fmt.Sprintf("- **%s:** %s", tech.Name, tech.DocsURL))
			if tech.Version != "" && tech.Version != "latest" {
				sb.WriteString(fmt.Sprintf(" (v%s)", tech.Version))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	// Services & APIs
	serviceTechs := g.categorizeServices(stack)
	if len(serviceTechs) > 0 {
		sb.WriteString("### Services & APIs\n\n")
		for _, tech := range serviceTechs {
			sb.WriteString(fmt.Sprintf("- **%s:** %s", tech.Name, tech.DocsURL))
			// Check for SOPS directory
			sopsPath := filepath.Join(g.projectRoot, ".doplan", "SOPS", strings.ToLower(tech.Name))
			if g.fileExists(sopsPath) {
				sb.WriteString(fmt.Sprintf(" - [Setup Guide](./.doplan/SOPS/%s/)", strings.ToLower(tech.Name)))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	// Project-Specific Documentation
	sb.WriteString("## Project-Specific Documentation\n\n")
	docs := []struct {
		name string
		path string
	}{
		{"API Specification", "doplan/contracts/api-spec.json"},
		{"Data Models", "doplan/contracts/data-model.md"},
		{"Design System", "doplan/design/DPR.md"},
		{"Product Requirements", "doplan/PRD.md"},
	}
	for _, doc := range docs {
		if g.fileExists(doc.path) {
			sb.WriteString(fmt.Sprintf("- [%s](./%s)\n", doc.name, doc.path))
		}
	}
	if len(docs) == 0 || !g.fileExists("doplan/contracts/api-spec.json") {
		sb.WriteString("- [API Specification](./doplan/contracts/api-spec.json) *(to be created)*\n")
		sb.WriteString("- [Data Models](./doplan/contracts/data-model.md) *(to be created)*\n")
		sb.WriteString("- [Design System](./doplan/design/DPR.md) *(to be created)*\n")
	}
	sb.WriteString("\n")

	// Development Guidelines
	sb.WriteString("## Development Guidelines\n\n")
	sb.WriteString("- **Coding standards:** Follow project conventions\n")
	sb.WriteString("- **File naming conventions:** Use kebab-case for files, PascalCase for components\n")
	sb.WriteString("- **Component patterns:** [To be defined]\n")
	sb.WriteString("- **Testing approach:** [To be defined]\n")
	sb.WriteString("\n")

	// DoPlan Resources (collapsible)
	sb.WriteString("## DoPlan Resources\n\n")
	sb.WriteString("<details>\n")
	sb.WriteString("<summary>DoPlan CLI Documentation</summary>\n\n")
	sb.WriteString("This project uses DoPlan for workflow automation and project management.\n\n")
	sb.WriteString("### IDE Integration\n\n")
	sb.WriteString(fmt.Sprintf("- **Rules:** %s.doplan/ai/rules/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.doplan/ai/commands/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Agents:** %s.doplan/ai/agents/%s\n", backtick, backtick))
	sb.WriteString("\n")
	sb.WriteString("### Cursor IDE\n")
	sb.WriteString(fmt.Sprintf("- **Rules:** %s.cursor/rules/%s (symlinked from .doplan/ai/rules/)\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.cursor/commands/%s (symlinked from .doplan/ai/commands/)\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** This file (%sCONTEXT.md%s) is automatically loaded\n\n", backtick, backtick))
	sb.WriteString("### Other IDEs\n")
	sb.WriteString(fmt.Sprintf("- **VS Code:** See %s.doplan/guides/vscode_setup.md%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Generic:** See %s.doplan/guides/generic_ide_setup.md%s\n", backtick, backtick))
	sb.WriteString("\n")
	sb.WriteString("</details>\n\n")
	sb.WriteString("---\n\n")
	sb.WriteString("**Last Updated:** Auto-generated by DoPlan\n")
	sb.WriteString("**Note:** This file is automatically updated during `doplan install` and can be manually edited to add additional technologies.\n")

	return sb.String()
}

// categorizeFrontend identifies frontend technologies
func (g *ContextGenerator) categorizeFrontend(stack *TechStack) []Technology {
	var frontend []Technology
	frontendKeywords := []string{"react", "vue", "angular", "next", "nuxt", "svelte", "tailwind", "css", "html", "typescript", "javascript"}

	for _, lang := range stack.Languages {
		if strings.Contains(strings.ToLower(lang.Name), "javascript") || strings.Contains(strings.ToLower(lang.Name), "typescript") {
			frontend = append(frontend, lang)
		}
	}

	for _, fw := range stack.Frameworks {
		nameLower := strings.ToLower(fw.Name)
		for _, keyword := range frontendKeywords {
			if strings.Contains(nameLower, keyword) {
				frontend = append(frontend, fw)
				break
			}
		}
	}

	return frontend
}

// categorizeBackend identifies backend technologies
func (g *ContextGenerator) categorizeBackend(stack *TechStack) []Technology {
	var backend []Technology
	backendKeywords := []string{"express", "fastify", "koa", "django", "flask", "rails", "gin", "echo", "fiber"}

	for _, lang := range stack.Languages {
		if strings.Contains(strings.ToLower(lang.Name), "go") || strings.Contains(strings.ToLower(lang.Name), "python") || strings.Contains(strings.ToLower(lang.Name), "rust") {
			backend = append(backend, lang)
		}
	}

	for _, fw := range stack.Frameworks {
		nameLower := strings.ToLower(fw.Name)
		for _, keyword := range backendKeywords {
			if strings.Contains(nameLower, keyword) {
				backend = append(backend, fw)
				break
			}
		}
	}

	for _, db := range stack.Databases {
		backend = append(backend, db)
	}

	return backend
}

// categorizeServices identifies services and APIs
func (g *ContextGenerator) categorizeServices(stack *TechStack) []Technology {
	var services []Technology
	services = append(services, stack.Services...)
	return services
}
