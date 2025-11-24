package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BranchMatrix describes CI expectations per branch prefix
// Example: task/ -> { jobs: ["lint", "test"], required: ["lint", "test"] }
type BranchMatrix struct {
	Prefix   string   `json:"prefix"`
	Jobs     []string `json:"jobs"`
	Required []string `json:"required"`
}

type MatrixConfig struct {
	GeneratedAt string         `json:"generated_at"`
	Branches    []BranchMatrix `json:"branches"`
}

var (
	outputDir  = flag.String("out", ".github/workflows", "Directory for generated workflow")
	matrixFile = flag.String("matrix", "Docs/history/branch-matrix.json", "Matrix configuration file")
)

func main() {
	flag.Parse()

	matrix := loadMatrix(*matrixFile)
	workflow := renderWorkflow(matrix)
	path := filepath.Join(*outputDir, "task-branches.yml")
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create workflows dir: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(path, []byte(workflow), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write workflow: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Workflow generated: %s\n", path)
}

func loadMatrix(path string) MatrixConfig {
	data, err := os.ReadFile(path)
	if err != nil {
		return MatrixConfig{GeneratedAt: time.Now().UTC().Format(time.RFC3339), Branches: defaultMatrix()}
	}
	var cfg MatrixConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "invalid matrix json, falling back to defaults: %v\n", err)
		return MatrixConfig{GeneratedAt: time.Now().UTC().Format(time.RFC3339), Branches: defaultMatrix()}
	}
	return cfg
}

func defaultMatrix() []BranchMatrix {
	return []BranchMatrix{
		{
			Prefix:   "task/",
			Jobs:     []string{"lint", "test"},
			Required: []string{"lint", "test"},
		},
		{
			Prefix:   "feature/",
			Jobs:     []string{"lint", "test", "build"},
			Required: []string{"lint", "test"},
		},
		{
			Prefix:   "hotfix/",
			Jobs:     []string{"lint", "test", "build"},
			Required: []string{"lint", "test", "build"},
		},
	}
}

func renderWorkflow(cfg MatrixConfig) string {
	var builder strings.Builder
	builder.WriteString("name: Task Branch CI\n\n")
	builder.WriteString("on:\n  push:\n    branches:\n      - 'task/*'\n  workflow_dispatch:{}\n\n")
	builder.WriteString("jobs:\n")

	for _, branch := range cfg.Branches {
		jobName := sanitize(branch.Prefix)
		builder.WriteString(fmt.Sprintf("  %s:\n", jobName))
		builder.WriteString("    runs-on: ubuntu-latest\n")
		builder.WriteString("    if: startsWith(github.ref, 'refs/heads/" + branch.Prefix + "')\n")
		builder.WriteString("    steps:\n")
		builder.WriteString("      - uses: actions/checkout@v4\n")
		builder.WriteString("      - name: Set up Go\n        uses: actions/setup-go@v5\n        with:\n          go-version: '1.21'\n")
		for _, job := range branch.Jobs {
			builder.WriteString(renderJobStep(job))
		}
	}

	builder.WriteString("\n  summary:\n    runs-on: ubuntu-latest\n    needs:\n")
	for _, branch := range cfg.Branches {
		builder.WriteString("      - " + sanitize(branch.Prefix) + "\n")
	}
	builder.WriteString("    steps:\n      - name: Display required checks\n        run: |")
	builder.WriteString("\n          echo 'Required checks per branch:'\n")
	for _, branch := range cfg.Branches {
		builder.WriteString(fmt.Sprintf("          echo '%s -> %s'\n", branch.Prefix, strings.Join(branch.Required, ", ")))
	}
	builder.WriteString("\n")
	return builder.String()
}

func sanitize(prefix string) string {
	prefix = strings.TrimSuffix(prefix, "/")
	prefix = strings.ReplaceAll(prefix, "-", "_")
	prefix = strings.ReplaceAll(prefix, "/", "_")
	if prefix == "" {
		return "job"
	}
	return prefix
}

func renderJobStep(job string) string {
	switch job {
	case "lint":
		return "      - name: Run golangci-lint\n        uses: golangci/golangci-lint-action@v3\n"
	case "test":
		return "      - name: Run tests\n        run: go test ./...\n"
	case "build":
		return "      - name: Build CLI\n        run: go build ./cmd/doplan\n"
	case "publish":
		return "      - name: Publish Image\n        run: echo 'TODO: add publish step'\n"
	default:
		return "      - name: Custom step\n        run: echo 'Unknown job " + job + "'\n"
	}
}
