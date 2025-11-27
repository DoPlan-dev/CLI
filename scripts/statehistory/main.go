package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/statehistory"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		usage(stdout)
		return 1
	}

	cmd := args[0]
	var err error
	switch cmd {
	case "snapshot":
		err = runSnapshot(args[1:], stdout)
	case "list":
		err = runList(args[1:], stdout)
	case "diff":
		err = runDiff(args[1:], stdout)
	case "restore":
		err = runRestore(args[1:], stdout)
	case "help", "-h", "--help":
		usage(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n\n", cmd)
		usage(stdout)
		return 1
	}

	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	return 0
}

func runSnapshot(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("snapshot", flag.ExitOnError)
	fs.SetOutput(io.Discard)
	statePath := fs.String("state", ".do/system/history/active_state.json", "Path to active_state.json")
	historyDir := fs.String("history", ".do/system/history", "Directory for history snapshots")
	reason := fs.String("reason", "", "Reason or trigger for this snapshot")
	label := fs.String("label", "", "Optional suffix appended to the file name")
	if err := fs.Parse(args); err != nil {
		return err
	}

	snap, err := statehistory.SaveSnapshot(*statePath, *historyDir, *reason, *label)
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, "Snapshot saved: %s (%s)\n", snap.Path, snap.ID)
	return nil
}

func runList(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	fs.SetOutput(io.Discard)
	historyDir := fs.String("history", ".do/system/history", "Directory for history snapshots")
	limit := fs.Int("limit", 10, "Maximum entries to display (0 = all)")
	jsonOut := fs.Bool("json", false, "Emit JSON instead of text")
	if err := fs.Parse(args); err != nil {
		return err
	}

	snaps, err := statehistory.ListSnapshots(*historyDir)
	if err != nil {
		return err
	}
	if *limit > 0 && len(snaps) > *limit {
		snaps = snaps[len(snaps)-*limit:]
	}

	if *jsonOut {
		payload := make([]statehistory.SnapshotSummary, 0, len(snaps))
		for _, snap := range snaps {
			payload = append(payload, summaryFromSnapshot(snap))
		}
		data, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(data))
		return nil
	}

	if len(snaps) == 0 {
		fmt.Fprintln(stdout, "No state snapshots found.")
		return nil
	}

	fmt.Fprintf(stdout, "Found %d snapshot(s):\n", len(snaps))
	for _, snap := range snaps {
		fmt.Fprintf(stdout, "- %s | %s | reason: %s | file: %s\n",
			snap.ID, snap.CapturedAt.Format(timeFmt()), valueOrDefault(snap.Reason, "(none)"), filepath.Base(snap.Path))
	}
	return nil
}

func runDiff(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("diff", flag.ExitOnError)
	fs.SetOutput(io.Discard)
	historyDir := fs.String("history", ".do/system/history", "Directory for history snapshots")
	fromID := fs.String("from", "", "Older snapshot ID or file")
	toID := fs.String("to", "", "Newer snapshot ID or file (defaults to latest)")
	jsonOut := fs.Bool("json", false, "Emit JSON instead of Markdown summary")
	if err := fs.Parse(args); err != nil {
		return err
	}

	snaps, err := statehistory.ListSnapshots(*historyDir)
	if err != nil {
		return err
	}
	if len(snaps) < 2 {
		return statehistory.ErrInsufficientSnapshots
	}

	var older, newer *statehistory.Snapshot

	switch {
	case *fromID != "" && *toID != "":
		older, err = findSnapshot(snaps, *fromID)
		if err != nil {
			return err
		}
		newer, err = findSnapshot(snaps, *toID)
		if err != nil {
			return err
		}
	case *fromID != "":
		older, err = findSnapshot(snaps, *fromID)
		if err != nil {
			return err
		}
		newer = snaps[len(snaps)-1]
	case *toID != "":
		newer, err = findSnapshot(snaps, *toID)
		if err != nil {
			return err
		}
		idx := indexOf(snaps, newer)
		if idx <= 0 {
			return fmt.Errorf("no older snapshot found before %s", newer.ID)
		}
		older = snaps[idx-1]
	default:
		older = snaps[len(snaps)-2]
		newer = snaps[len(snaps)-1]
	}

	if newer.CapturedAt.Before(older.CapturedAt) {
		older, newer = newer, older
	}

	diff := statehistory.ComputeDiff(newer, older)
	if *jsonOut {
		data, err := json.MarshalIndent(diff, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(data))
		return nil
	}

	fmt.Fprintln(stdout, statehistory.FormatDiff(diff))
	return nil
}

func runRestore(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("restore", flag.ExitOnError)
	fs.SetOutput(io.Discard)
	historyDir := fs.String("history", ".do/system/history", "Directory for history snapshots")
	statePath := fs.String("state", ".do/system/history/active_state.json", "Path to active_state.json")
	fileID := fs.String("file", "", "Snapshot ID or file name to restore")
	confirm := fs.Bool("yes", false, "Confirm rollback (required)")
	capture := fs.Bool("snapshot", true, "Capture a new snapshot after restore")
	reason := fs.String("reason", "Rollback", "Reason for the post-restore snapshot")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if *fileID == "" {
		return fmt.Errorf("--file is required for restore")
	}
	if !*confirm {
		return fmt.Errorf("restore aborted: pass --yes to confirm the rollback")
	}

	snap, err := statehistory.LoadSnapshot(*historyDir, *fileID)
	if err != nil {
		return err
	}

	if err := statehistory.RestoreSnapshot(*statePath, snap); err != nil {
		return err
	}
	fmt.Fprintf(stdout, "State restored to snapshot %s (%s).\n", snap.ID, filepath.Base(snap.Path))

	if *capture {
		_, err = statehistory.SaveSnapshot(*statePath, *historyDir, fmt.Sprintf("Rollback to %s (%s)", snap.ID, *reason), "rollback")
		if err != nil {
			return fmt.Errorf("post-restore snapshot failed: %w", err)
		}
		fmt.Fprintln(stdout, "Post-restore snapshot captured.")
	}
	return nil
}

func usage(w io.Writer) {
	fmt.Fprintln(w, `Usage: statehistory <command> [options]

Commands:
  snapshot   Capture the current .do/plan/active_state.json as a history entry
  list       Show existing snapshots (latest first)
  diff       Compare two snapshots (default: latest vs previous)
  restore    Restore active_state.json from a snapshot (requires --yes)

Examples:
  go run scripts/statehistory/main.go snapshot --reason "after /finished"
  go run scripts/statehistory/main.go list --limit 5
  go run scripts/statehistory/main.go diff --json
  go run scripts/statehistory/main.go restore --file state-20251124T120000Z.json --yes`)
}

func summaryFromSnapshot(snap *statehistory.Snapshot) statehistory.SnapshotSummary {
	return statehistory.SnapshotSummary{
		ID:         snap.ID,
		CapturedAt: snap.CapturedAt,
		Reason:     snap.Reason,
		Path:       snap.Path,
	}
}

func valueOrDefault(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func timeFmt() string {
	return "2006-01-02 15:04:05"
}

func findSnapshot(snaps []*statehistory.Snapshot, identifier string) (*statehistory.Snapshot, error) {
	for _, snap := range snaps {
		if snap.ID == identifier || filepath.Base(snap.Path) == identifier || strings.Contains(filepath.Base(snap.Path), identifier) {
			return snap, nil
		}
	}
	return nil, fmt.Errorf("snapshot %s not found", identifier)
}

func indexOf(snaps []*statehistory.Snapshot, target *statehistory.Snapshot) int {
	for idx, snap := range snaps {
		if snap.Path == target.Path {
			return idx
		}
	}
	return -1
}
