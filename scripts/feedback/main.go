package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FeedbackEntry struct {
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Author    string `json:"author"`
	GitHubURL string `json:"github_url,omitempty"`
}

var now = time.Now

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("feedback", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	projectPath := fs.String("project", ".", "Path to project root")
	fbType := fs.String("type", "note", "Feedback type: bug | feature | question | note")
	title := fs.String("title", "", "Short title / summary")
	details := fs.String("details", "", "Detailed feedback text")
	author := fs.String("author", "anonymous", "Author or reporter name")
	github := fs.String("github", "", "Optional GitHub issue URL")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(stderr, "failed to parse flags: %v\n", err)
		return 2
	}

	if strings.TrimSpace(*title) == "" {
		fmt.Fprintln(stderr, "--title is required")
		return 1
	}

	entry := FeedbackEntry{
		Timestamp: now().UTC().Format(time.RFC3339),
		Type:      strings.ToLower(*fbType),
		Title:     strings.TrimSpace(*title),
		Details:   strings.TrimSpace(*details),
		Author:    strings.TrimSpace(*author),
		GitHubURL: strings.TrimSpace(*github),
	}

	historyDir := filepath.Join(*projectPath, "Docs", "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		fmt.Fprintf(stderr, "failed to create Docs/history: %v\n", err)
		return 1
	}

	markdownPath := filepath.Join(historyDir, "feedback.md")
	jsonPath := filepath.Join(historyDir, "feedback.json")

	if err := appendMarkdown(markdownPath, entry); err != nil {
		fmt.Fprintf(stderr, "failed to append markdown: %v\n", err)
		return 1
	}

	if err := appendJSON(jsonPath, entry); err != nil {
		fmt.Fprintf(stderr, "failed to append json: %v\n", err)
		return 1
	}

	fmt.Fprintf(stdout, "Feedback logged to %s and %s\n", markdownPath, jsonPath)
	return 0
}

func appendMarkdown(path string, entry FeedbackEntry) error {
	var builder strings.Builder
	builder.WriteString("\n---\n")
	builder.WriteString(fmt.Sprintf("### %s (%s)\n", entry.Title, strings.Title(entry.Type)))
	builder.WriteString(fmt.Sprintf("- **Reported**: %s UTC\n", entry.Timestamp))
	builder.WriteString(fmt.Sprintf("- **Author**: %s\n", entry.Author))
	if entry.GitHubURL != "" {
		builder.WriteString(fmt.Sprintf("- **GitHub**: %s\n", entry.GitHubURL))
	}
	if entry.Details != "" {
		builder.WriteString("\n")
		builder.WriteString(entry.Details)
		builder.WriteString("\n")
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(builder.String()); err != nil {
		return err
	}
	return nil
}

func appendJSON(path string, entry FeedbackEntry) error {
	var entries []FeedbackEntry
	if data, err := os.ReadFile(path); err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, &entries); err != nil {
			return err
		}
	}
	entries = append(entries, entry)
	encoded, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, encoded, 0644)
}
