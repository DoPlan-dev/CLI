package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/generator"
)

var scaffoldPlanHierarchy = generator.ScaffoldPlanHierarchy

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("plan", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	projectPath := fs.String("project", ".", "Path to project root")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	absPath, err := filepath.Abs(*projectPath)
	if err != nil {
		fmt.Fprintf(stderr, "Error: failed to resolve project path: %v\n", err)
		return 1
	}

	if err := scaffoldPlanHierarchy(absPath); err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Fprintln(stdout, "Planning hierarchy scaffolded successfully!")
	fmt.Fprintln(stdout, "Phase folders created in .do/plan/")
	return 0
}
