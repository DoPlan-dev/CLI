package generators

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	// Generate context document
	content := g.generateContextContent(techStack)

	return os.WriteFile(contextPath, []byte(content), 0644)
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
func (g *ContextGenerator) generateContextContent(stack *TechStack) string {
	backtick := "`"
	var sb strings.Builder

	sb.WriteString("# Project Technology Stack\n\n")
	sb.WriteString("This document provides a comprehensive overview of all technologies, frameworks, tools, and services used in this project.\n")
	sb.WriteString("It serves as a context file for IDEs and CLIs to understand the project's technical stack.\n\n")
	sb.WriteString("---\n\n")

	// Programming Languages
	if len(stack.Languages) > 0 {
		sb.WriteString("## Programming Languages\n\n")
		for _, lang := range stack.Languages {
			sb.WriteString(fmt.Sprintf("### %s\n\n", lang.Name))
			if lang.Version != "" {
				sb.WriteString(fmt.Sprintf("- **Version:** %s\n", lang.Version))
			}
			sb.WriteString(fmt.Sprintf("- **Documentation:** %s\n", lang.DocsURL))
			if lang.Description != "" {
				sb.WriteString(fmt.Sprintf("- **Description:** %s\n", lang.Description))
			}
			if lang.Usage != "" {
				sb.WriteString(fmt.Sprintf("- **Usage:** %s\n", lang.Usage))
			}
			sb.WriteString("\n")
		}
	}

	// Frameworks
	if len(stack.Frameworks) > 0 {
		sb.WriteString("## Frameworks\n\n")
		for _, fw := range stack.Frameworks {
			sb.WriteString(fmt.Sprintf("### %s\n\n", fw.Name))
			if fw.Version != "" {
				sb.WriteString(fmt.Sprintf("- **Version:** %s\n", fw.Version))
			}
			sb.WriteString(fmt.Sprintf("- **Documentation:** %s\n", fw.DocsURL))
			if fw.Description != "" {
				sb.WriteString(fmt.Sprintf("- **Description:** %s\n", fw.Description))
			}
			if fw.Usage != "" {
				sb.WriteString(fmt.Sprintf("- **Purpose:** %s\n", fw.Usage))
			}
			sb.WriteString("\n")
		}
	}

	// CLIs and Tools
	if len(stack.CLIs) > 0 {
		sb.WriteString("## CLIs and Development Tools\n\n")
		for _, cli := range stack.CLIs {
			sb.WriteString(fmt.Sprintf("### %s\n\n", cli.Name))
			if cli.Version != "" {
				sb.WriteString(fmt.Sprintf("- **Version:** %s\n", cli.Version))
			}
			sb.WriteString(fmt.Sprintf("- **Documentation:** %s\n", cli.DocsURL))
			if cli.Description != "" {
				sb.WriteString(fmt.Sprintf("- **Description:** %s\n", cli.Description))
			}
			if cli.Usage != "" {
				sb.WriteString(fmt.Sprintf("- **Usage:** %s\n", cli.Usage))
			}
			sb.WriteString("\n")
		}
	}

	// Services
	if len(stack.Services) > 0 {
		sb.WriteString("## Services and Platforms\n\n")
		for _, svc := range stack.Services {
			sb.WriteString(fmt.Sprintf("### %s\n\n", svc.Name))
			if svc.Version != "" {
				sb.WriteString(fmt.Sprintf("- **Version:** %s\n", svc.Version))
			}
			sb.WriteString(fmt.Sprintf("- **Documentation:** %s\n", svc.DocsURL))
			if svc.Description != "" {
				sb.WriteString(fmt.Sprintf("- **Description:** %s\n", svc.Description))
			}
			if svc.Usage != "" {
				sb.WriteString(fmt.Sprintf("- **Purpose:** %s\n", svc.Usage))
			}
			sb.WriteString("\n")
		}
	}

	// Databases
	if len(stack.Databases) > 0 {
		sb.WriteString("## Databases\n\n")
		for _, db := range stack.Databases {
			sb.WriteString(fmt.Sprintf("### %s\n\n", db.Name))
			if db.Version != "" {
				sb.WriteString(fmt.Sprintf("- **Version:** %s\n", db.Version))
			}
			sb.WriteString(fmt.Sprintf("- **Documentation:** %s\n", db.DocsURL))
			if db.Description != "" {
				sb.WriteString(fmt.Sprintf("- **Description:** %s\n", db.Description))
			}
			if db.Usage != "" {
				sb.WriteString(fmt.Sprintf("- **Usage:** %s\n", db.Usage))
			}
			sb.WriteString("\n")
		}
	}

	// Development Tools
	if len(stack.Tools) > 0 {
		sb.WriteString("## Development Tools\n\n")
		for _, tool := range stack.Tools {
			sb.WriteString(fmt.Sprintf("### %s\n\n", tool.Name))
			if tool.Version != "" {
				sb.WriteString(fmt.Sprintf("- **Version:** %s\n", tool.Version))
			}
			sb.WriteString(fmt.Sprintf("- **Documentation:** %s\n", tool.DocsURL))
			if tool.Description != "" {
				sb.WriteString(fmt.Sprintf("- **Description:** %s\n", tool.Description))
			}
			if tool.Usage != "" {
				sb.WriteString(fmt.Sprintf("- **Usage:** %s\n", tool.Usage))
			}
			sb.WriteString("\n")
		}
	}

	// IDE Integration Section
	sb.WriteString("## IDE/CLI Integration\n\n")
	sb.WriteString("This project uses DoPlan for workflow automation. The following IDEs/CLIs are supported:\n\n")
	sb.WriteString("### Cursor IDE\n")
	sb.WriteString(fmt.Sprintf("- **Rules:** %s.cursor/rules/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.cursor/commands/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** This file (%sCONTEXT.md%s) is automatically loaded\n\n", backtick, backtick))

	sb.WriteString("### Gemini CLI\n")
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.gemini/commands/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** Reference this file in command prompts\n\n"))

	sb.WriteString("### Claude Code\n")
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.claude/commands/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** Use %s@CONTEXT.md%s to reference this file\n\n", backtick, backtick))

	sb.WriteString("### Codex CLI\n")
	sb.WriteString(fmt.Sprintf("- **Prompts:** %s.codex/prompts/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** Reference this file in prompt templates\n\n"))

	sb.WriteString("### OpenCode\n")
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.opencode/command/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** Use %s@CONTEXT.md%s to reference this file\n\n", backtick, backtick))

	sb.WriteString("### Qwen Code\n")
	sb.WriteString(fmt.Sprintf("- **Commands:** %s.qwen/commands/%s\n", backtick, backtick))
	sb.WriteString(fmt.Sprintf("- **Context:** Reference this file in command prompts\n\n"))

	sb.WriteString("---\n\n")
	sb.WriteString("**Last Updated:** Auto-generated by DoPlan\n")
	sb.WriteString("**Note:** This file is automatically updated during `doplan install` and can be manually edited to add additional technologies.\n")

	return sb.String()
}
