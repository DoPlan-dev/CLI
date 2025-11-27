package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/content"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

// parseMarkdownWithFrontmatter parses a markdown file with YAML frontmatter
// Returns metadata map and markdown content
func parseMarkdownWithFrontmatter(data []byte) (map[string]interface{}, string, error) {
	content := string(data)

	// Match YAML frontmatter: ---\n...\n---
	frontmatterRegex := regexp.MustCompile(`^---\s*\n([\s\S]*?)\n---\s*\n([\s\S]*)$`)
	matches := frontmatterRegex.FindStringSubmatch(content)

	if len(matches) != 3 {
		return nil, "", fmt.Errorf("invalid markdown frontmatter format")
	}

	yamlContent := matches[1]
	markdownContent := matches[2]

	// Simple YAML parser for our specific use case
	// We'll parse the frontmatter manually since we know the structure
	metadata := make(map[string]interface{})

	lines := strings.Split(yamlContent, "\n")
	var currentKey string
	var currentList []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Handle list items
		if strings.HasPrefix(line, "- ") {
			if currentKey != "" {
				currentList = append(currentList, strings.TrimPrefix(line, "- "))
				metadata[currentKey] = currentList
			}
			continue
		}

		// Flush current list if any
		if currentList != nil && currentKey != "" {
			metadata[currentKey] = currentList
			currentList = nil
		}

		// Handle key-value pairs
		if idx := strings.Index(line, ":"); idx > 0 {
			currentKey = strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])

			// Remove quotes if present
			if (strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)) ||
				(strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`)) {
				value = value[1 : len(value)-1]
			}

			if value == "[]" || value == "" {
				currentList = []string{}
				metadata[currentKey] = currentList
			} else {
				metadata[currentKey] = value
				currentKey = ""
			}
		}
	}

	// Flush any remaining list
	if currentList != nil && currentKey != "" {
		metadata[currentKey] = currentList
	}

	return metadata, markdownContent, nil
}

// extractSystemPromptFromMarkdown extracts the System Prompt section from markdown
func extractSystemPromptFromMarkdown(markdown string) string {
	// Look for "## System Prompt" section
	regex := regexp.MustCompile(`(?s)## System Prompt\s*\n(.*?)(?:\n## |$)`)
	matches := regex.FindStringSubmatch(markdown)
	if len(matches) > 1 {
		content := strings.TrimSpace(matches[1])
		// Remove trailing ## if present
		content = regexp.MustCompile(`(?s)\n## .*$`).ReplaceAllString(content, "")
		return strings.TrimSpace(content)
	}
	return ""
}

// LoadAgentsFromFiles loads agents from markdown files with YAML frontmatter (proof-of-concept)
// This follows the same pattern as ExtractRules
func LoadAgentsFromFiles() ([]Agent, error) {
	var agents []Agent

	// Walk through embedded agent files
	err := content.WalkAgentDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking agents: %w", err)
		}

		// Skip directories, non-markdown files, and README files
		if d.IsDir() || filepath.Ext(path) != ".md" || filepath.Base(path) == "README.md" {
			return nil
		}

		// Path from WalkDir already includes "agents/" prefix, but ReadAgentFile adds it
		// So we need to strip it. Path format: "agents/category/file.md"
		relPath := path
		const agentsPrefix = "agents/"
		if len(path) > len(agentsPrefix) && path[:len(agentsPrefix)] == agentsPrefix {
			relPath = path[len(agentsPrefix):]
		}

		data, err := content.ReadAgentFile(relPath)
		if err != nil {
			return fmt.Errorf("failed to read agent file %s (original path: %s): %w", relPath, path, err)
		}

		// Parse markdown with frontmatter
		metadata, markdownContent, err := parseMarkdownWithFrontmatter(data)
		if err != nil {
			return fmt.Errorf("failed to parse agent markdown %s: %w", relPath, err)
		}

		// Extract system prompt from markdown
		systemPrompt := extractSystemPromptFromMarkdown(markdownContent)

		// Build Agent struct from metadata
		agent := Agent{
			Name:         getString(metadata, "name"),
			Role:         getString(metadata, "role"),
			SystemPrompt: systemPrompt,
			ReportsTo:    getString(metadata, "reportsTo"),
			FileName:     filepath.Base(relPath),
			Category:     getString(metadata, "category"),
		}

		// Parse manages list
		if manages, ok := metadata["manages"].([]string); ok {
			agent.Manages = manages
		} else if manages, ok := metadata["manages"].(string); ok && manages == "" {
			agent.Manages = []string{}
		}

		// Parse responsibilities list
		if responsibilities, ok := metadata["responsibilities"].([]string); ok {
			agent.Responsibilities = responsibilities
		}

		agents = append(agents, agent)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error loading agents from files: %w", err)
	}

	return agents, nil
}

// getString safely gets a string value from metadata map
func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}

// ExtractAgents extracts agent JSON files to a target directory (similar to ExtractRules)
func ExtractAgents(targetDir string) error {
	if targetDir == "" {
		return fmt.Errorf("target directory cannot be empty")
	}

	// Walk through embedded agents and extract them
	err := content.WalkAgentDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking agents: %w", err)
		}

		// Path from WalkDir includes "agents/" prefix, strip it for target path
		relPath := path
		const agentsPrefix = "agents/"
		if len(path) > len(agentsPrefix) && path[:len(agentsPrefix)] == agentsPrefix {
			relPath = path[len(agentsPrefix):]
		}

		targetPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			// Create directory
			if err := utils.CreateDirectory(targetPath); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
		} else {
			// Read file from embedded FS (relPath already has "agents/" stripped)
			fileData, err := content.ReadAgentFile(relPath)
			if err != nil {
				return fmt.Errorf("failed to read embedded file %s (original path: %s): %w", relPath, path, err)
			}

			// Write file to target
			if err := utils.WriteFile(targetPath, fileData); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error extracting agents: %w", err)
	}

	return nil
}

// GetAllAgents returns all agents, loading from files if available (with fallback)
func GetAllAgentsFileBased() ([]Agent, error) {
	// Try to load from embedded files first
	agents, err := LoadAgentsFromFiles()
	if err != nil {
		// Fallback to hardcoded (for backward compatibility)
		return GetAllAgents(), fmt.Errorf("failed to load from files, using hardcoded: %w", err)
	}

	if len(agents) == 0 {
		// Fallback if no agents loaded
		return GetAllAgents(), nil
	}

	return agents, nil
}

// LoadAgentsFromDirectory loads agents from a directory containing JSON files
// This is useful for testing and when agents are extracted to disk
func LoadAgentsFromDirectory(dir string) ([]Agent, error) {
	var agents []Agent

	if dir == "" {
		return nil, fmt.Errorf("directory cannot be empty")
	}

	// Walk directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, non-markdown files, and README files
		if info.IsDir() || filepath.Ext(path) != ".md" || filepath.Base(path) == "README.md" {
			return nil
		}

		// Read and parse markdown with frontmatter
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		metadata, markdownContent, err := parseMarkdownWithFrontmatter(data)
		if err != nil {
			return fmt.Errorf("failed to parse markdown in %s: %w", path, err)
		}

		// Extract system prompt from markdown
		systemPrompt := extractSystemPromptFromMarkdown(markdownContent)

		// Build Agent struct from metadata
		agent := Agent{
			Name:         getString(metadata, "name"),
			Role:         getString(metadata, "role"),
			SystemPrompt: systemPrompt,
			ReportsTo:    getString(metadata, "reportsTo"),
			FileName:     filepath.Base(path),
			Category:     getString(metadata, "category"),
		}

		// Parse manages list
		if manages, ok := metadata["manages"].([]string); ok {
			agent.Manages = manages
		}

		// Parse responsibilities list
		if responsibilities, ok := metadata["responsibilities"].([]string); ok {
			agent.Responsibilities = responsibilities
		}

		agents = append(agents, agent)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error loading agents from directory: %w", err)
	}

	return agents, nil
}
