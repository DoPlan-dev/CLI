package context

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ProjectAnalysis contains analysis results for an existing project
type ProjectAnalysis struct {
	TechStack       []string         `json:"techStack"`
	ProjectFiles    []string         `json:"projectFiles"`
	Documentation   []string         `json:"documentation"`
	PotentialPhases []PotentialPhase `json:"potentialPhases"`
	TODOs           []string         `json:"todos"`
}

// PotentialPhase represents a potential phase detected from folder structure
type PotentialPhase struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Features []string `json:"features"`
}

// Analyzer analyzes an existing project
type Analyzer struct {
	projectRoot string
}

// NewAnalyzer creates a new project analyzer
func NewAnalyzer(projectRoot string) *Analyzer {
	return &Analyzer{
		projectRoot: projectRoot,
	}
}

// Analyze performs project analysis
func (a *Analyzer) Analyze() (*ProjectAnalysis, error) {
	analysis := &ProjectAnalysis{
		TechStack:       []string{},
		ProjectFiles:    []string{},
		Documentation:   []string{},
		PotentialPhases: []PotentialPhase{},
		TODOs:           []string{},
	}

	// Detect tech stack
	analysis.TechStack = a.detectTechStack()

	// Find project files
	analysis.ProjectFiles = a.findProjectFiles()

	// Find documentation
	analysis.Documentation = a.findDocumentation()

	// Identify potential phases/features
	analysis.PotentialPhases = a.identifyPotentialPhases()

	// Extract TODO comments from code
	analysis.TODOs = a.extractTODOs()

	return analysis, nil
}

// detectTechStack detects the technology stack from project files
func (a *Analyzer) detectTechStack() []string {
	var stack []string

	// Check for package.json (Node.js)
	if _, err := os.Stat(filepath.Join(a.projectRoot, "package.json")); err == nil {
		stack = append(stack, "Node.js")
		// Try to read and detect framework
		data, err := os.ReadFile(filepath.Join(a.projectRoot, "package.json"))
		if err == nil {
			var pkg map[string]interface{}
			if json.Unmarshal(data, &pkg) == nil {
				if deps, ok := pkg["dependencies"].(map[string]interface{}); ok {
					if _, hasReact := deps["react"]; hasReact {
						stack = append(stack, "React")
					}
					if _, hasNext := deps["next"]; hasNext {
						stack = append(stack, "Next.js")
					}
				}
			}
		}
	}

	// Check for go.mod (Go)
	if _, err := os.Stat(filepath.Join(a.projectRoot, "go.mod")); err == nil {
		stack = append(stack, "Go")
	}

	// Check for requirements.txt (Python)
	if _, err := os.Stat(filepath.Join(a.projectRoot, "requirements.txt")); err == nil {
		stack = append(stack, "Python")
	}

	// Check for Cargo.toml (Rust)
	if _, err := os.Stat(filepath.Join(a.projectRoot, "Cargo.toml")); err == nil {
		stack = append(stack, "Rust")
	}

	return stack
}

// findProjectFiles finds important project files
func (a *Analyzer) findProjectFiles() []string {
	var files []string

	commonFiles := []string{
		"package.json",
		"go.mod",
		"requirements.txt",
		"Cargo.toml",
		"README.md",
		"CONTRIBUTING.md",
		"LICENSE",
	}

	for _, file := range commonFiles {
		path := filepath.Join(a.projectRoot, file)
		if _, err := os.Stat(path); err == nil {
			files = append(files, file)
		}
	}

	return files
}

// findDocumentation finds documentation files
func (a *Analyzer) findDocumentation() []string {
	var docs []string

	docPatterns := []string{
		"*.md",
		"docs/**/*.md",
		"*.rst",
		"docs/**/*.rst",
	}

	for _, pattern := range docPatterns {
		matches, _ := filepath.Glob(filepath.Join(a.projectRoot, pattern))
		docs = append(docs, matches...)
	}

	return docs
}

// identifyPotentialPhases identifies potential phases from folder structure
func (a *Analyzer) identifyPotentialPhases() []PotentialPhase {
	var phases []PotentialPhase

	// Look for folders matching ##-slug-name pattern
	doplanDir := filepath.Join(a.projectRoot, "doplan")
	if _, err := os.Stat(doplanDir); err == nil {
		entries, err := os.ReadDir(doplanDir)
		if err == nil {
			phasePattern := regexp.MustCompile(`^(\d+)-(.+)$`)
			for _, entry := range entries {
				if entry.IsDir() {
					matches := phasePattern.FindStringSubmatch(entry.Name())
					if len(matches) > 0 {
						phasePath := filepath.Join(doplanDir, entry.Name())
						features := a.findFeaturesInPhase(phasePath)
						phases = append(phases, PotentialPhase{
							Name:     matches[2],
							Path:     phasePath,
							Features: features,
						})
					}
				}
			}
		}
	}

	return phases
}

// findFeaturesInPhase finds features within a phase directory
func (a *Analyzer) findFeaturesInPhase(phasePath string) []string {
	var features []string

	entries, err := os.ReadDir(phasePath)
	if err != nil {
		return features
	}

	featurePattern := regexp.MustCompile(`^(\d+)-(.+)$`)
	for _, entry := range entries {
		if entry.IsDir() {
			matches := featurePattern.FindStringSubmatch(entry.Name())
			if len(matches) > 0 {
				features = append(features, matches[2])
			}
		}
	}

	return features
}

// extractTODOs extracts TODO comments from code files
func (a *Analyzer) extractTODOs() []string {
	var todos []string

	// Basic TODO extraction; can be enhanced with more sophisticated parsing
	codeExtensions := []string{"*.go", "*.js", "*.ts", "*.py", "*.rs", "*.java"}

	for _, ext := range codeExtensions {
		matches, _ := filepath.Glob(filepath.Join(a.projectRoot, "**", ext))
		for _, file := range matches {
			data, err := os.ReadFile(file)
			if err == nil {
				lines := strings.Split(string(data), "\n")
				for _, line := range lines {
					if strings.Contains(strings.ToUpper(line), "TODO") {
						todos = append(todos, strings.TrimSpace(line))
					}
				}
			}
		}
	}

	return todos
}
