package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/content"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

// parseCommandMarkdown parses a command markdown file with YAML frontmatter
func parseCommandMarkdown(data []byte) (*Command, error) {
	contentStr := string(data)

	// Split frontmatter and content
	parts := strings.SplitN(contentStr, "---\n", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown frontmatter format")
	}

	frontmatter := parts[1]
	markdownContent := parts[2]

	// Parse frontmatter (simple YAML parser for our use case)
	metadata := make(map[string]interface{})
	lines := strings.Split(frontmatter, "\n")
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

	// Extract Action from markdown (everything after frontmatter)
	action := strings.TrimSpace(markdownContent)

	// Build Command struct
	cmd := &Command{
		Name:        getString(metadata, "name"),
		Category:    getString(metadata, "category"),
		Trigger:     getString(metadata, "trigger"),
		Description: getString(metadata, "description"),
		Action:      action,
	}

	// Parse optional fields
	if agentInvolvement, ok := metadata["agentInvolvement"].([]string); ok {
		cmd.AgentInvolvement = agentInvolvement
	}
	if filesRead, ok := metadata["filesRead"].([]string); ok {
		cmd.FilesRead = filesRead
	}
	if filesModified, ok := metadata["filesModified"].([]string); ok {
		cmd.FilesModified = filesModified
	}
	if examples, ok := metadata["examples"].([]string); ok {
		cmd.Examples = examples
	}

	// Optional string fields
	cmd.GitHubAutomation = getString(metadata, "githubAutomation")
	cmd.Requirements = getString(metadata, "requirements")
	cmd.Notes = getString(metadata, "notes")
	cmd.Customize = getString(metadata, "customize")
	cmd.Options = getString(metadata, "options")
	cmd.OfflineSafety = getString(metadata, "offlineSafety")

	return cmd, nil
}

// LoadCommandsFromFiles loads commands from embedded markdown files
func LoadCommandsFromFiles() ([]Command, error) {
	var commands []Command

	err := content.WalkCommandDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking commands: %w", err)
		}

		// Skip directories, non-markdown files, and README files
		if d.IsDir() || filepath.Ext(path) != ".md" || filepath.Base(path) == "README.md" {
			return nil
		}

		// Path from WalkDir includes "commands/" prefix, strip it
		relPath := path
		const commandsPrefix = "commands/"
		if len(path) > len(commandsPrefix) && path[:len(commandsPrefix)] == commandsPrefix {
			relPath = path[len(commandsPrefix):]
		}

		data, err := content.ReadCommandFile(relPath)
		if err != nil {
			return fmt.Errorf("failed to read command file %s: %w", relPath, err)
		}

		cmd, err := parseCommandMarkdown(data)
		if err != nil {
			return fmt.Errorf("failed to parse command markdown %s: %w", relPath, err)
		}

		commands = append(commands, *cmd)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error loading commands from files: %w", err)
	}

	return commands, nil
}

// LoadCommandsFromDirectory loads commands from a directory containing markdown files
func LoadCommandsFromDirectory(dir string) ([]Command, error) {
	var commands []Command

	if dir == "" {
		return nil, fmt.Errorf("directory cannot be empty")
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, non-markdown files, and README files
		if info.IsDir() || filepath.Ext(path) != ".md" || filepath.Base(path) == "README.md" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		cmd, err := parseCommandMarkdown(data)
		if err != nil {
			return fmt.Errorf("failed to parse markdown in %s: %w", path, err)
		}

		commands = append(commands, *cmd)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error loading commands from directory: %w", err)
	}

	return commands, nil
}

// ExtractCommands extracts command markdown files to a target directory
func ExtractCommands(targetDir string) error {
	if targetDir == "" {
		return fmt.Errorf("target directory cannot be empty")
	}

	err := content.WalkCommandDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking commands: %w", err)
		}

		// Path from WalkDir includes "commands/" prefix, strip it
		relPath := path
		const commandsPrefix = "commands/"
		if len(path) > len(commandsPrefix) && path[:len(commandsPrefix)] == commandsPrefix {
			relPath = path[len(commandsPrefix):]
		}

		targetPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			if err := utils.CreateDirectory(targetPath); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
		} else {
			fileData, err := content.ReadCommandFile(relPath)
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
		return fmt.Errorf("error extracting commands: %w", err)
	}

	return nil
}

// GetAllCommandsFileBased returns all commands, loading from files if available (with fallback)
func GetAllCommandsFileBased() ([]Command, error) {
	// Try to load from embedded files first
	commands, err := LoadCommandsFromFiles()
	if err != nil {
		// Fallback to hardcoded
		return GetAllCommands(), fmt.Errorf("failed to load from files, using hardcoded: %w", err)
	}

	if len(commands) == 0 {
		// Fallback if no commands loaded
		return GetAllCommands(), nil
	}

	return commands, nil
}
