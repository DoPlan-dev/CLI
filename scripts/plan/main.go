package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/generator"
)

func main() {
	projectPath := flag.String("project", ".", "Path to project root")
	flag.Parse()

	absPath, err := filepath.Abs(*projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to resolve project path: %v\n", err)
		os.Exit(1)
	}

	if err := generator.ScaffoldPlanHierarchy(absPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Planning hierarchy scaffolded successfully!")
	fmt.Println("Phase folders created in .plan/")
}

