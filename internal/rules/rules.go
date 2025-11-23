package rules

import (
	"embed"
	"io/fs"
)

//go:embed library/*
var embeddedRules embed.FS

// Library returns the embedded rules filesystem
func Library() fs.FS {
	return embeddedRules
}

// ReadFile reads a file from the embedded rules library
func ReadFile(name string) ([]byte, error) {
	return embeddedRules.ReadFile("library/" + name)
}

// ReadDir reads a directory from the embedded rules library
func ReadDir(name string) ([]fs.DirEntry, error) {
	return embeddedRules.ReadDir("library/" + name)
}

// WalkDir walks the embedded rules library directory tree
func WalkDir(root string, fn fs.WalkDirFunc) error {
	path := "library"
	if root != "" && root != "." {
		path = "library/" + root
	}
	return fs.WalkDir(embeddedRules, path, fn)
}

