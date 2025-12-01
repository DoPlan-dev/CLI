package content

import (
	"embed"
	"io/fs"
)

//go:embed agents/*
var embeddedAgents embed.FS

//go:embed commands/**/*.md
var embeddedCommands embed.FS

//go:embed templates/agents/*.tmpl templates/commands/*.tmpl templates/documents/**/*.md
var embeddedTemplates embed.FS

// Agents returns the embedded agents filesystem
func Agents() fs.FS {
	return embeddedAgents
}

// ReadAgentFile reads an agent markdown file from the embedded content
func ReadAgentFile(name string) ([]byte, error) {
	return embeddedAgents.ReadFile("agents/" + name)
}

// ReadAgentDir reads a directory from the embedded agents
func ReadAgentDir(name string) ([]fs.DirEntry, error) {
	return embeddedAgents.ReadDir("agents/" + name)
}

// WalkAgentDir walks the embedded agents directory tree
func WalkAgentDir(root string, fn fs.WalkDirFunc) error {
	path := "agents"
	if root != "" && root != "." {
		path = "agents/" + root
	}
	return fs.WalkDir(embeddedAgents, path, fn)
}

// Commands returns the embedded commands filesystem
func Commands() fs.FS {
	return embeddedCommands
}

// ReadCommandFile reads a command markdown file from the embedded content
func ReadCommandFile(name string) ([]byte, error) {
	return embeddedCommands.ReadFile("commands/" + name)
}

// ReadCommandDir reads a directory from the embedded commands
func ReadCommandDir(name string) ([]fs.DirEntry, error) {
	return embeddedCommands.ReadDir("commands/" + name)
}

// WalkCommandDir walks the embedded commands directory tree
func WalkCommandDir(root string, fn fs.WalkDirFunc) error {
	path := "commands"
	if root != "" && root != "." {
		path = "commands/" + root
	}
	return fs.WalkDir(embeddedCommands, path, fn)
}

// Templates returns the embedded templates filesystem
func Templates() fs.FS {
	return embeddedTemplates
}

// ReadTemplateFile reads a template file from the embedded content
func ReadTemplateFile(name string) ([]byte, error) {
	return embeddedTemplates.ReadFile("templates/" + name)
}

// ReadTemplateDir reads a directory from the embedded templates
func ReadTemplateDir(name string) ([]fs.DirEntry, error) {
	return embeddedTemplates.ReadDir("templates/" + name)
}

// WalkTemplateDir walks the embedded templates directory tree
func WalkTemplateDir(root string, fn fs.WalkDirFunc) error {
	path := "templates"
	if root != "" && root != "." {
		path = "templates/" + root
	}
	return fs.WalkDir(embeddedTemplates, path, fn)
}
