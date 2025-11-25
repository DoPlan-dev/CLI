package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/generator"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

func main() {
	projectPath := flag.String("project", ".", "Path to the project root")
	projectType := flag.String("type", "Fullstack", "Project type for boilerplate generation")
	ide := flag.String("ide", "Cursor", "Preferred IDE (used for metadata only)")
	flag.Parse()

	absProjectPath, err := filepath.Abs(*projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve project path: %v\n", err)
		os.Exit(1)
	}

	request := &models.ProjectRequest{
		ProjectName: filepath.Base(absProjectPath),
		IDE:         *ide,
		ProjectType: *projectType,
	}

	if err := generator.GenerateBoilerplate(request, absProjectPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate boilerplate: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✨ Boilerplate generated in %s\n", absProjectPath)
}
