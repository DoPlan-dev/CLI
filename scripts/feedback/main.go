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

type FeedbackEntry struct {
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Author    string `json:"author"`
	GitHubURL string `json:"github_url,omitempty"`
}

func main() {
	projectPath := flag.String("project", ".", "Path to project root")
	fbType := flag.String("type", "note", "Feedback type: bug | feature | question | note")
	title := flag.String("title", "", "Short title / summary")
	details := flag.String("details", "", "Detailed feedback text")
	author := flag.String("author", "anonymous", "Author or reporter name")
	github := flag.String("github", "", "Optional GitHub issue URL")
	flag.Parse()

	if strings.TrimSpace(*title) == "" {
		fmt.Fprintln(os.Stderr, "--title is required")
		os.Exit(1)
	}

	entry := FeedbackEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Type:      strings.ToLower(*fbType),
		Title:     strings.TrimSpace(*title),
		Details:   strings.TrimSpace(*details),
		Author:    strings.TrimSpace(*author),
		GitHubURL: strings.TrimSpace(*github),
	}

	historyDir := filepath.Join(*projectPath, "Docs", "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create Docs/history: %v\n", err)
		os.Exit(1)
	}

	markdownPath := filepath.Join(historyDir, "feedback.md")
	jsonPath := filepath.Join(historyDir, "feedback.json")

	if err := appendMarkdown(markdownPath, entry); err != nil {
		fmt.Fprintf(os.Stderr, "failed to append markdown: %v\n", err)
		os.Exit(1)
	}

	if err := appendJSON(jsonPath, entry); err != nil {
		fmt.Fprintf(os.Stderr, "failed to append json: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Feedback logged to %s and %s\n", markdownPath, jsonPath)
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
