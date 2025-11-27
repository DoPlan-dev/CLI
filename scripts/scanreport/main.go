package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/progress"
	"github.com/DoPlan-dev/CLI/internal/statehistory"
)

type ScanMetadata struct {
	File             string          `json:"file"`
	ScanDate         string          `json:"scan_date"`
	Project          string          `json:"project"`
	ReportType       string          `json:"report_type"`
	SummaryHash      string          `json:"summary_hash"`
	GeneratedAt      string          `json:"generated_at"`
	ExecutiveSummary []string        `json:"executive_summary"`
	Findings         []string        `json:"findings"`
	NextActions      []string        `json:"next_actions"`
	RecentFeedback   []FeedbackEntry `json:"recent_feedback,omitempty"`
}

type FeedbackEntry struct {
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Author    string `json:"author"`
	GitHubURL string `json:"github_url,omitempty"`
}

type presetConfig struct {
	Title    string
	Sections []string
}

type diffSections struct {
	Executive  string
	Findings   string
	Next       string
	Feedback   string
	State      string
	Progress   string
	Visuals    string
	Dependency string
}

var presetConfigs = map[string]presetConfig{
	"standard": {
		Title: "Standard Report",
		Sections: []string{
			"executive",
			"findings",
			"next",
			"feedback",
			"state",
			"progress",
			"visuals",
			"dependency",
		},
	},
	"exec": {
		Title: "Executive Summary",
		Sections: []string{
			"executive",
			"progress",
			"visuals",
			"state",
			"feedback",
		},
	},
	"detailed": {
		Title: "Detailed Report",
		Sections: []string{
			"executive",
			"visuals",
			"findings",
			"next",
			"feedback",
			"state",
			"progress",
			"dependency",
		},
	},
}

var (
	outWriter                 io.Writer = os.Stdout
	errWriter                 io.Writer = os.Stderr
	summarizeStateHistoryFunc           = summarizeStateHistory
	summarizeProgressFunc               = summarizeProgress
	renderVisualsFunc                   = renderVisuals
	renderDependencyAuditFunc           = renderDependencyAudit
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	outWriter = stdout
	errWriter = stderr

	fs := flag.NewFlagSet("scanreport", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	projectDir := fs.String("project", ".", "Path to the project root containing .do/reports")
	presetFlag := fs.String("preset", "standard", "Report preset: standard, exec, or detailed")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(errWriter, "error: %v\n", err)
		return 2
	}

	presetValue := strings.ToLower(strings.TrimSpace(*presetFlag))
	userSetPreset := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "preset" {
			userSetPreset = true
		}
	})

	cfgFile := loadReportConfig(*projectDir)
	if cfgFile != nil && cfgFile.Preset != "" && !userSetPreset {
		presetValue = strings.ToLower(strings.TrimSpace(cfgFile.Preset))
	}

	baseCfg, ok := presetConfigs[presetValue]
	if !ok {
		fmt.Fprintf(errWriter, "unknown preset %q (valid: standard, exec, detailed)\n", presetValue)
		return 1
	}
	cfg := baseCfg
	if custom := normalizedSections(cfgFile); len(custom) > 0 {
		cfg.Sections = custom
	}

	reportsDir := filepath.Join(*projectDir, ".do", "reports")
	if _, err := os.Stat(reportsDir); err != nil {
		fmt.Fprintf(errWriter, "reports directory not found: %s\n", reportsDir)
		return 1
	}

	reportFiles, err := listReportFiles(reportsDir)
	if err != nil {
		fmt.Fprintf(errWriter, "failed to list reports: %v\n", err)
		return 1
	}
	if len(reportFiles) == 0 {
		fmt.Fprintln(outWriter, "No scan reports found.")
		return 0
	}

	var metadataFiles []string
	for _, file := range reportFiles {
		meta, err := parseReport(filepath.Join(reportsDir, file), *projectDir)
		if err != nil {
			fmt.Fprintf(errWriter, "failed to parse %s: %v\n", file, err)
			continue
		}
		metaPath := filepath.Join(reportsDir, strings.TrimSuffix(file, filepath.Ext(file))+".json")
		if err := writeMetadata(metaPath, meta); err != nil {
			fmt.Fprintf(errWriter, "failed to write metadata for %s: %v\n", file, err)
			continue
		}
		metadataFiles = append(metadataFiles, metaPath)
	}

	if len(reportFiles) < 2 {
		fmt.Fprintln(outWriter, "Report metadata generated. (Only one report present; skipping diff.)")
		return 0
	}

	latest := reportFiles[len(reportFiles)-1]
	previous := reportFiles[len(reportFiles)-2]

	latestMeta, err := loadMetadata(filepath.Join(reportsDir, strings.TrimSuffix(latest, filepath.Ext(latest))+".json"))
	if err != nil {
		fmt.Fprintf(errWriter, "failed to load metadata for %s: %v\n", latest, err)
		return 1
	}
	previousMeta, err := loadMetadata(filepath.Join(reportsDir, strings.TrimSuffix(previous, filepath.Ext(previous))+".json"))
	if err != nil {
		fmt.Fprintf(errWriter, "failed to load metadata for %s: %v\n", previous, err)
		return 1
	}

	stateSummary := summarizeStateHistoryFunc(*projectDir)
	progressSummary, progressReport := summarizeProgressFunc(*projectDir)
	visualSummary := renderVisualsFunc(progressReport)
	dependencyAudit := renderDependencyAuditFunc(*projectDir)

	sections := diffSections{
		Executive:  renderDiff(latestMeta.ExecutiveSummary, previousMeta.ExecutiveSummary),
		Findings:   renderDiff(latestMeta.Findings, previousMeta.Findings),
		Next:       renderDiff(latestMeta.NextActions, previousMeta.NextActions),
		Feedback:   renderFeedbackDiff(latestMeta.RecentFeedback, previousMeta.RecentFeedback),
		State:      stateSummary,
		Progress:   progressSummary,
		Visuals:    visualSummary,
		Dependency: dependencyAudit,
	}

	diffContent := buildDiff(latestMeta, previousMeta, sections, cfg)
	diffFile := filepath.Join(reportsDir, fmt.Sprintf("SCAN_DIFF_%s.md", sanitizeDate(latestMeta.ScanDate)))
	if err := os.WriteFile(diffFile, []byte(diffContent), 0644); err != nil {
		fmt.Fprintf(errWriter, "failed to write diff file: %v\n", err)
		return 1
	}

	fmt.Fprintf(outWriter, "Generated metadata for %d reports.\n", len(metadataFiles))
	fmt.Fprintf(outWriter, "Diff written to %s (current vs previous).\n", diffFile)
	return 0
}

func listReportFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var reports []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, "SCAN_REPORT_") && strings.HasSuffix(name, ".md") {
			reports = append(reports, name)
		}
	}
	sort.Strings(reports)
	return reports, nil
}

func parseReport(path string, projectRoot string) (*ScanMetadata, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	meta := &ScanMetadata{File: filepath.Base(path), GeneratedAt: time.Now().UTC().Format(time.RFC3339)}
	meta.SummaryHash = fmt.Sprintf("%x", sha256.Sum256(data))

	section := ""
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "**Scan Date**") {
			meta.ScanDate = extractValue(trimmed)
		} else if strings.HasPrefix(trimmed, "**Project**") {
			meta.Project = extractValue(trimmed)
		} else if strings.HasPrefix(trimmed, "**Report Type**") {
			meta.ReportType = extractValue(trimmed)
		}

		if strings.HasPrefix(trimmed, "## ") {
			section = strings.TrimSpace(strings.TrimPrefix(trimmed, "## "))
			continue
		}

		if strings.HasPrefix(trimmed, "-") {
			switch section {
			case "Executive Summary":
				meta.ExecutiveSummary = append(meta.ExecutiveSummary, trimmed)
			case "6. Findings & Risks":
				meta.Findings = append(meta.Findings, trimmed)
			case "7. Recommended Next Actions":
				meta.NextActions = append(meta.NextActions, trimmed)
			}
		}
	}

	if meta.ScanDate == "" {
		meta.ScanDate = time.Now().Format("2006-01-02")
	}

	meta.RecentFeedback = loadRecentFeedback(projectRoot, 5)

	return meta, nil
}

func extractValue(line string) string {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return strings.TrimSpace(line)
	}
	return strings.TrimSpace(parts[1])
}

func writeMetadata(path string, meta *ScanMetadata) error {
	content, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, content, 0644)
}

func loadRecentFeedback(projectRoot string, limit int) []FeedbackEntry {
	historyPath := filepath.Join(projectRoot, "Docs", "history", "feedback.json")
	data, err := os.ReadFile(historyPath)
	if err != nil || len(data) == 0 {
		return nil
	}
	var entries []FeedbackEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil
	}
	if limit > 0 && len(entries) > limit {
		entries = entries[len(entries)-limit:]
	}
	return entries
}

func loadMetadata(path string) (*ScanMetadata, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var meta ScanMetadata
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func buildDiff(current, previous *ScanMetadata, sections diffSections, cfg presetConfig) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("# %s (%s vs %s)\n\n", cfg.Title, current.ScanDate, previous.ScanDate))

	for _, section := range cfg.Sections {
		switch section {
		case "executive":
			renderSection(&builder, "Executive Summary Changes", sections.Executive)
		case "findings":
			renderSection(&builder, "Findings & Risks Changes", sections.Findings)
		case "next":
			renderSection(&builder, "Recommended Next Actions Changes", sections.Next)
		case "feedback":
			renderSection(&builder, "Feedback Changes", sections.Feedback)
		case "state":
			renderSection(&builder, "State Changes Since Last Snapshot", sections.State)
		case "progress":
			renderSection(&builder, "Overall Progress Snapshot", sections.Progress)
		case "visuals":
			renderSection(&builder, "Visual Summary", sections.Visuals)
		case "dependency":
			renderSection(&builder, "Dependency Audit", sections.Dependency)
		}
	}

	builder.WriteString("---\n")
	builder.WriteString(fmt.Sprintf("Generated on %s\n", time.Now().Format(time.RFC3339)))
	return builder.String()
}

func renderSection(builder *strings.Builder, title, content string) {
	builder.WriteString("## ")
	builder.WriteString(title)
	builder.WriteString("\n")
	if strings.TrimSpace(content) == "" {
		builder.WriteString("No data available.\n\n")
		return
	}
	builder.WriteString(content)
	if !strings.HasSuffix(content, "\n") {
		builder.WriteString("\n")
	}
	builder.WriteString("\n")
}

func renderDiff(current, previous []string) string {
	additions, removals := diffLists(current, previous)
	var b strings.Builder
	if len(additions) == 0 && len(removals) == 0 {
		b.WriteString("No changes detected.\n")
		return b.String()
	}
	if len(additions) > 0 {
		b.WriteString("### Added\n")
		for _, line := range additions {
			b.WriteString("- ")
			b.WriteString(line)
			b.WriteString("\n")
		}
	}
	if len(removals) > 0 {
		b.WriteString("### Removed\n")
		for _, line := range removals {
			b.WriteString("- ")
			b.WriteString(line)
			b.WriteString("\n")
		}
	}
	return b.String()
}

func renderFeedbackDiff(current, previous []FeedbackEntry) string {
	added, removed := feedbackDiff(current, previous)
	var b strings.Builder
	if len(added) == 0 && len(removed) == 0 {
		b.WriteString("No feedback changes detected.\n")
		return b.String()
	}
	if len(added) > 0 {
		b.WriteString("### New Feedback\n")
		for _, entry := range added {
			b.WriteString(fmt.Sprintf("- **%s** (%s) by %s\n", entry.Title, strings.Title(entry.Type), entry.Author))
		}
	}
	if len(removed) > 0 {
		b.WriteString("### Resolved / Removed\n")
		for _, entry := range removed {
			b.WriteString(fmt.Sprintf("- **%s** (%s)\n", entry.Title, strings.Title(entry.Type)))
		}
	}
	return b.String()
}

func feedbackDiff(current, previous []FeedbackEntry) (added, removed []FeedbackEntry) {
	currMap := make(map[string]FeedbackEntry)
	prevMap := make(map[string]FeedbackEntry)
	for _, entry := range current {
		currMap[feedbackKey(entry)] = entry
	}
	for _, entry := range previous {
		prevMap[feedbackKey(entry)] = entry
	}
	for key, entry := range currMap {
		if _, exists := prevMap[key]; !exists {
			added = append(added, entry)
		}
	}
	for key, entry := range prevMap {
		if _, exists := currMap[key]; !exists {
			removed = append(removed, entry)
		}
	}
	return
}

func feedbackKey(entry FeedbackEntry) string {
	return fmt.Sprintf("%s|%s", entry.Timestamp, entry.Title)
}

func diffLists(current, previous []string) (added, removed []string) {
	currSet := make(map[string]struct{})
	prevSet := make(map[string]struct{})
	for _, line := range current {
		currSet[line] = struct{}{}
	}
	for _, line := range previous {
		prevSet[line] = struct{}{}
	}
	for line := range currSet {
		if _, exists := prevSet[line]; !exists {
			added = append(added, line)
		}
	}
	for line := range prevSet {
		if _, exists := currSet[line]; !exists {
			removed = append(removed, line)
		}
	}
	sort.Strings(added)
	sort.Strings(removed)
	return
}

func sanitizeDate(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return time.Now().Format("2006-01-02")
	}
	return strings.ReplaceAll(input, " ", "_")
}

func summarizeStateHistory(projectRoot string) string {
	historyDir := filepath.Join(projectRoot, ".do", "system", "history")
	diff, err := statehistory.LatestDiff(historyDir)
	if err != nil {
		if errors.Is(err, statehistory.ErrInsufficientSnapshots) {
			return "Not enough state snapshots have been recorded yet.\n"
		}
		return fmt.Sprintf("Unable to read state history (%v).\n", err)
	}
	return statehistory.FormatDiff(*diff)
}

func summarizeProgress(projectRoot string) (string, *progress.Report) {
	report, err := progress.Compute(projectRoot)
	if err != nil {
		return fmt.Sprintf("Unable to compute progress snapshot (%v).\n", err), nil
	}
	return progress.FormatMarkdown(report), report
}

func renderVisuals(report *progress.Report) string {
	if report == nil {
		return "Progress data unavailable.\n"
	}
	var b strings.Builder
	b.WriteString("### Completion\n")
	b.WriteString(renderProgressBar(report.Percentage, 30))
	b.WriteString(fmt.Sprintf(" %.1f%%\n\n", report.Percentage))

	b.WriteString("### Checklist\n")
	b.WriteString(fmt.Sprintf("- ✅ Completed: %d\n", report.Checklist.Completed))
	b.WriteString(fmt.Sprintf("- ⏳ Remaining: %d\n", report.Checklist.Pending))
	if report.Checklist.NextUp != "" {
		b.WriteString(fmt.Sprintf("- 🎯 Next Up: %s\n", report.Checklist.NextUp))
	}
	return b.String()
}

func renderProgressBar(percent float64, width int) string {
	if width <= 0 {
		width = 20
	}
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	filled := int((percent / 100) * float64(width))
	if filled > width {
		filled = width
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return "[" + bar + "]"
}

func renderDependencyAudit(projectRoot string) string {
	if summary := collectPackageJSONDeps(projectRoot); summary != "" {
		return summary
	}
	if summary := collectGoModDeps(projectRoot); summary != "" {
		return summary
	}
	return "No dependency manifest detected. Add package.json or go.mod to enable audit output.\n"
}

func collectPackageJSONDeps(projectRoot string) string {
	path := filepath.Join(projectRoot, "package.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	var manifest struct {
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		return ""
	}
	type dep struct {
		Name    string
		Version string
		Scope   string
	}
	var deps []dep
	for name, version := range manifest.Dependencies {
		deps = append(deps, dep{Name: name, Version: version, Scope: "runtime"})
	}
	for name, version := range manifest.DevDependencies {
		deps = append(deps, dep{Name: name, Version: version, Scope: "dev"})
	}
	if len(deps) == 0 {
		return ""
	}
	sort.Slice(deps, func(i, j int) bool {
		return deps[i].Name < deps[j].Name
	})
	if len(deps) > 10 {
		deps = deps[:10]
	}
	var b strings.Builder
	b.WriteString("Detected dependencies from package.json:\n")
	for _, d := range deps {
		b.WriteString(fmt.Sprintf("- %s %s (%s)\n", d.Name, d.Version, d.Scope))
	}
	if len(manifest.Dependencies)+len(manifest.DevDependencies) > len(deps) {
		b.WriteString("...additional dependencies omitted for brevity.\n")
	}
	return b.String()
}

func collectGoModDeps(projectRoot string) string {
	path := filepath.Join(projectRoot, "go.mod")
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	lines := strings.Split(string(data), "\n")
	type dep struct {
		Name    string
		Version string
	}
	var deps []dep
	inBlock := false
	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if strings.HasPrefix(line, "require (") {
			inBlock = true
			continue
		}
		if inBlock {
			if line == ")" {
				inBlock = false
				continue
			}
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				deps = append(deps, dep{Name: parts[0], Version: parts[1]})
			}
			continue
		}
		if strings.HasPrefix(line, "require ") {
			parts := strings.Fields(strings.TrimPrefix(line, "require"))
			if len(parts) >= 2 {
				deps = append(deps, dep{Name: parts[0], Version: parts[1]})
			}
		}
	}
	if len(deps) == 0 {
		return ""
	}
	sort.Slice(deps, func(i, j int) bool {
		return deps[i].Name < deps[j].Name
	})
	if len(deps) > 10 {
		deps = deps[:10]
	}
	var b strings.Builder
	b.WriteString("Detected dependencies from go.mod:\n")
	for _, d := range deps {
		b.WriteString(fmt.Sprintf("- %s %s\n", d.Name, d.Version))
	}
	return b.String()
}

type reportConfig struct {
	Preset   string   `json:"preset"`
	Sections []string `json:"sections"`
}

var validSections = map[string]bool{
	"executive":  true,
	"findings":   true,
	"next":       true,
	"feedback":   true,
	"state":      true,
	"progress":   true,
	"visuals":    true,
	"dependency": true,
}

func loadReportConfig(projectRoot string) *reportConfig {
	path := filepath.Join(projectRoot, ".do", "reports", "config.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var cfg reportConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Fprintf(errWriter, "warning: invalid report config: %v\n", err)
		return nil
	}
	return &cfg
}

func normalizedSections(cfg *reportConfig) []string {
	if cfg == nil || len(cfg.Sections) == 0 {
		return nil
	}
	seen := make(map[string]bool)
	var sections []string
	for _, raw := range cfg.Sections {
		sec := strings.ToLower(strings.TrimSpace(raw))
		if !validSections[sec] || seen[sec] {
			continue
		}
		seen[sec] = true
		sections = append(sections, sec)
	}
	return sections
}
