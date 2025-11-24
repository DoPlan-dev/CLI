package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/progress"
)

func main() {
	root := flag.String("root", ".", "Project root directory")
	jsonOut := flag.Bool("json", false, "Emit JSON instead of pretty text")
	includeDiffStruct := flag.Bool("diff-struct", false, "Include raw state diff JSON when available")
	flag.Parse()

	report, err := progress.Compute(*root)
	if err != nil {
		exitErr(err)
	}

	if !*includeDiffStruct {
		report.StateDiff = nil
	}

	if *jsonOut {
		data, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			exitErr(err)
		}
		fmt.Println(string(data))
		return
	}

	fmt.Println(progress.FormatPlain(report))
}

func exitErr(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
