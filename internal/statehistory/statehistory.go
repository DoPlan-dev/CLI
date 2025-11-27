package statehistory

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	// ErrInsufficientSnapshots indicates there are not enough history entries
	// to perform the requested action (e.g., diff needs at least two).
	ErrInsufficientSnapshots = errors.New("not enough state snapshots")
)

// ActiveState mirrors the structure inside .do/plan/active_state.json.
type ActiveState struct {
	Phase        string   `json:"phase"`
	ActiveTask   string   `json:"active_task"`
	ActiveBranch string   `json:"active_branch"`
	Completed    []string `json:"completed"`
	Locked       bool     `json:"locked"`
}

// Snapshot represents a captured copy of the active state with metadata.
type Snapshot struct {
	ID         string
	CapturedAt time.Time
	Reason     string
	Hash       string
	Path       string
	State      ActiveState
	Raw        json.RawMessage
}

// SnapshotSummary is a lightweight descriptor used in diffs and listings.
type SnapshotSummary struct {
	ID         string    `json:"id"`
	CapturedAt time.Time `json:"captured_at"`
	Reason     string    `json:"reason,omitempty"`
	Path       string    `json:"path"`
}

// FieldChange captures before/after values for a simple field.
type FieldChange struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Changed bool   `json:"changed"`
}

// StateDiff describes the delta between two snapshots.
type StateDiff struct {
	From             SnapshotSummary `json:"from"`
	To               SnapshotSummary `json:"to"`
	Phase            FieldChange     `json:"phase"`
	ActiveTask       FieldChange     `json:"active_task"`
	ActiveBranch     FieldChange     `json:"active_branch"`
	CompletedAdded   []string        `json:"completed_added"`
	CompletedRemoved []string        `json:"completed_removed"`
}

// SaveSnapshot captures the current active_state.json into history.
func SaveSnapshot(statePath, historyDir, reason, label string) (*Snapshot, error) {
	data, err := os.ReadFile(statePath)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state ActiveState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}
	if state.Completed == nil {
		state.Completed = []string{}
	}

	if err := os.MkdirAll(historyDir, 0o755); err != nil {
		return nil, fmt.Errorf("create history dir: %w", err)
	}

	now := time.Now().UTC()
	id := now.Format("20060102T150405Z")
	cleanLabel := sanitizeLabel(label)
	name := fmt.Sprintf("state-%s", id)
	if cleanLabel != "" {
		name += "-" + cleanLabel
	}
	name += ".json"

	payload := snapshotFile{
		ID:         id,
		CapturedAt: now.Format(time.RFC3339),
		Reason:     reason,
		Hash:       fmt.Sprintf("%x", sha256.Sum256(data)),
		State:      json.RawMessage(data),
	}
	content, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal snapshot: %w", err)
	}

	fullPath := filepath.Join(historyDir, name)
	if err := os.WriteFile(fullPath, content, 0o644); err != nil {
		return nil, fmt.Errorf("write snapshot: %w", err)
	}

	return &Snapshot{
		ID:         id,
		CapturedAt: now,
		Reason:     reason,
		Hash:       payload.Hash,
		Path:       fullPath,
		State:      state,
		Raw:        payload.State,
	}, nil
}

// ListSnapshots returns every snapshot (sorted ascending by time).
func ListSnapshots(historyDir string) ([]*Snapshot, error) {
	entries, err := os.ReadDir(historyDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []*Snapshot{}, nil
		}
		return nil, fmt.Errorf("read history dir: %w", err)
	}

	var paths []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, "state-") && strings.HasSuffix(name, ".json") {
			paths = append(paths, filepath.Join(historyDir, name))
		}
	}
	sort.Strings(paths)

	snapshots := make([]*Snapshot, 0, len(paths))
	for _, path := range paths {
		snap, err := loadSnapshot(path)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, snap)
	}
	return snapshots, nil
}

// LatestSnapshots returns up to count newest snapshots.
func LatestSnapshots(historyDir string, count int) ([]*Snapshot, error) {
	snaps, err := ListSnapshots(historyDir)
	if err != nil {
		return nil, err
	}
	if len(snaps) <= count {
		return snaps, nil
	}
	return snaps[len(snaps)-count:], nil
}

// LoadSnapshot fetches a snapshot by file name or ID.
func LoadSnapshot(historyDir, idOrFile string) (*Snapshot, error) {
	if idOrFile == "" {
		return nil, errors.New("snapshot identifier required")
	}

	candidate := idOrFile
	if !filepath.IsAbs(candidate) {
		candidate = filepath.Join(historyDir, idOrFile)
	}
	if strings.HasSuffix(candidate, ".json") {
		if _, err := os.Stat(candidate); err == nil {
			return loadSnapshot(candidate)
		}
	}

	// Try plain ID (without prefix) or partial name.
	idCandidate := filepath.Join(historyDir, fmt.Sprintf("state-%s.json", idOrFile))
	if _, err := os.Stat(idCandidate); err == nil {
		return loadSnapshot(idCandidate)
	}

	// Fallback: search for substring match.
	entries, err := os.ReadDir(historyDir)
	if err != nil {
		return nil, fmt.Errorf("read history dir: %w", err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.Contains(entry.Name(), idOrFile) {
			return loadSnapshot(filepath.Join(historyDir, entry.Name()))
		}
	}
	return nil, fmt.Errorf("snapshot %s not found", idOrFile)
}

// RestoreSnapshot overwrites active_state.json with the snapshot contents.
func RestoreSnapshot(statePath string, snapshot *Snapshot) error {
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, snapshot.Raw, "", "  "); err != nil {
		return fmt.Errorf("format state: %w", err)
	}
	if err := os.WriteFile(statePath, pretty.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write state: %w", err)
	}
	return nil
}

// ComputeDiff compares two snapshots (newer vs older).
func ComputeDiff(newer, older *Snapshot) StateDiff {
	diff := StateDiff{
		From: SnapshotSummary{
			ID:         older.ID,
			CapturedAt: older.CapturedAt,
			Reason:     older.Reason,
			Path:       older.Path,
		},
		To: SnapshotSummary{
			ID:         newer.ID,
			CapturedAt: newer.CapturedAt,
			Reason:     newer.Reason,
			Path:       newer.Path,
		},
		Phase: FieldChange{
			From: older.State.Phase,
			To:   newer.State.Phase,
		},
		ActiveTask: FieldChange{
			From: older.State.ActiveTask,
			To:   newer.State.ActiveTask,
		},
		ActiveBranch: FieldChange{
			From: older.State.ActiveBranch,
			To:   newer.State.ActiveBranch,
		},
	}

	diff.Phase.Changed = diff.Phase.From != diff.Phase.To
	diff.ActiveTask.Changed = diff.ActiveTask.From != diff.ActiveTask.To
	diff.ActiveBranch.Changed = diff.ActiveBranch.From != diff.ActiveBranch.To

	added, removed := diffSlices(newer.State.Completed, older.State.Completed)
	diff.CompletedAdded = added
	diff.CompletedRemoved = removed

	return diff
}

// FormatDiff renders a human-readable summary suitable for Markdown.
func FormatDiff(diff StateDiff) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Snapshots: %s (%s) → %s (%s)\n",
		diff.From.ID, diff.From.CapturedAt.Format(time.RFC3339), diff.To.ID, diff.To.CapturedAt.Format(time.RFC3339)))
	if diff.From.Reason != "" || diff.To.Reason != "" {
		b.WriteString("Reasons:\n")
		if diff.From.Reason != "" {
			b.WriteString(fmt.Sprintf("- From: %s\n", diff.From.Reason))
		}
		if diff.To.Reason != "" {
			b.WriteString(fmt.Sprintf("- To: %s\n", diff.To.Reason))
		}
	}
	if diff.Phase.Changed {
		b.WriteString(fmt.Sprintf("- Phase changed: %s → %s\n", fallback(diff.Phase.From, "(none)"), fallback(diff.Phase.To, "(none)")))
	}
	if diff.ActiveTask.Changed {
		b.WriteString(fmt.Sprintf("- Active task: %s → %s\n", fallback(diff.ActiveTask.From, "(none)"), fallback(diff.ActiveTask.To, "(none)")))
	}
	if diff.ActiveBranch.Changed {
		b.WriteString(fmt.Sprintf("- Active branch: %s → %s\n", fallback(diff.ActiveBranch.From, "(none)"), fallback(diff.ActiveBranch.To, "(none)")))
	}
	if len(diff.CompletedAdded) > 0 {
		b.WriteString("- Newly completed tasks:\n")
		for _, item := range diff.CompletedAdded {
			b.WriteString("  - " + item + "\n")
		}
	}
	if len(diff.CompletedRemoved) > 0 {
		b.WriteString("- Tasks removed from completed:\n")
		for _, item := range diff.CompletedRemoved {
			b.WriteString("  - " + item + "\n")
		}
	}
	if !diff.HasChanges() {
		b.WriteString("- No state changes detected between these snapshots.\n")
	}
	return b.String()
}

// HasChanges reports whether any tracked field differs.
func (d StateDiff) HasChanges() bool {
	return d.Phase.Changed || d.ActiveTask.Changed || d.ActiveBranch.Changed || len(d.CompletedAdded) > 0 || len(d.CompletedRemoved) > 0
}

// LatestDiffSummary loads the newest two snapshots and returns a formatted diff.
func LatestDiffSummary(historyDir string) (string, error) {
	snaps, err := LatestSnapshots(historyDir, 2)
	if err != nil {
		return "", err
	}
	if len(snaps) < 2 {
		return "", ErrInsufficientSnapshots
	}
	diff := ComputeDiff(snaps[1], snaps[0])
	return FormatDiff(diff), nil
}

// LatestDiff returns the diff struct for the newest two snapshots.
func LatestDiff(historyDir string) (*StateDiff, error) {
	snaps, err := LatestSnapshots(historyDir, 2)
	if err != nil {
		return nil, err
	}
	if len(snaps) < 2 {
		return nil, ErrInsufficientSnapshots
	}
	diff := ComputeDiff(snaps[1], snaps[0])
	return &diff, nil
}

func loadSnapshot(path string) (*Snapshot, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read snapshot: %w", err)
	}
	var payload snapshotFile
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("parse snapshot: %w", err)
	}
	captured, err := time.Parse(time.RFC3339, payload.CapturedAt)
	if err != nil {
		captured = time.Unix(0, 0).UTC()
	}
	var state ActiveState
	if err := json.Unmarshal(payload.State, &state); err != nil {
		return nil, fmt.Errorf("parse state payload: %w", err)
	}
	if state.Completed == nil {
		state.Completed = []string{}
	}
	return &Snapshot{
		ID:         payload.ID,
		CapturedAt: captured,
		Reason:     payload.Reason,
		Hash:       payload.Hash,
		Path:       path,
		State:      state,
		Raw:        payload.State,
	}, nil
}

func diffSlices(newer, older []string) (added, removed []string) {
	nset := make(map[string]struct{}, len(newer))
	oset := make(map[string]struct{}, len(older))
	for _, item := range newer {
		nset[item] = struct{}{}
	}
	for _, item := range older {
		oset[item] = struct{}{}
	}
	for item := range nset {
		if _, ok := oset[item]; !ok {
			added = append(added, item)
		}
	}
	for item := range oset {
		if _, ok := nset[item]; !ok {
			removed = append(removed, item)
		}
	}
	sort.Strings(added)
	sort.Strings(removed)
	return
}

func fallback(value, placeholder string) string {
	if strings.TrimSpace(value) == "" {
		return placeholder
	}
	return value
}

func sanitizeLabel(label string) string {
	label = strings.TrimSpace(label)
	if label == "" {
		return ""
	}
	label = strings.ToLower(label)
	label = strings.ReplaceAll(label, " ", "-")
	var filtered strings.Builder
	for _, r := range label {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			filtered.WriteRune(r)
		}
	}
	return filtered.String()
}

type snapshotFile struct {
	ID         string          `json:"id"`
	CapturedAt string          `json:"captured_at"`
	Reason     string          `json:"reason,omitempty"`
	Hash       string          `json:"hash"`
	State      json.RawMessage `json:"state"`
}
