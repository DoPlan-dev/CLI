package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/DoPlan-dev/CLI/internal/progress"
)

var (
	computeProgress = progress.Compute
	formatProgress  = progress.FormatPlain
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("progress", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	root := fs.String("root", ".", "Project root directory")
	jsonOut := fs.Bool("json", false, "Emit JSON instead of pretty text")
	includeDiffStruct := fs.Bool("diff-struct", false, "Include raw state diff JSON when available")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 2
	}

	report, err := computeProgress(*root)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}

	if !*includeDiffStruct {
		report.StateDiff = nil
	}

	if *jsonOut {
		data, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			fmt.Fprintf(stderr, "error: %v\n", err)
			return 1
		}
		fmt.Fprintln(stdout, string(data))
		return 0
	}

	fmt.Fprintln(stdout, formatProgress(report))
	return 0
}
