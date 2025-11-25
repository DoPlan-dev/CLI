package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type GitMeta struct {
	RemoteName     string   `json:"remote_name"`
	RemoteURL      string   `json:"remote_url"`
	RepoSlug       string   `json:"repo_slug"`
	DefaultBranch  string   `json:"default_branch"`
	SuccessMetrics []string `json:"success_metrics"`
	UpdatedAt      string   `json:"updated_at"`
}

var (
	projectPath    = flag.String("project", ".", "Path to project root")
	syncReadme     = flag.Bool("sync-readme", false, "Update README KPI block")
	issueTitle     = flag.String("issue-title", "", "Issue title to compose gh command")
	issueBody      = flag.String("issue-body", "", "Issue body text")
	milestoneTitle = flag.String("milestone-title", "", "Milestone title to compose command")
	milestoneDue   = flag.String("milestone-due", "", "Optional milestone due date")
)

func main() {
	flag.Parse()

	gitInfo, err := detectGit(*projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Git detection failed: %v\n", err)
	}

	metrics := extractSuccessMetrics(filepath.Join(*projectPath, ".plan", "00_System", "PRD.md"))

	meta := GitMeta{
		RemoteName:     gitInfo.RemoteName,
		RemoteURL:      gitInfo.RemoteURL,
		RepoSlug:       gitInfo.RepoSlug,
		DefaultBranch:  gitInfo.DefaultBranch,
		SuccessMetrics: metrics,
		UpdatedAt:      time.Now().UTC().Format(time.RFC3339),
	}

	if err := writeMeta(*projectPath, meta); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to persist meta: %v\n", err)
	}

	if *syncReadme {
		if err := updateReadmeKPIs(*projectPath, metrics); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to sync README KPIs: %v\n", err)
		} else {
			fmt.Println("README KPI block updated.")
		}
	}

	if *issueTitle != "" {
		printIssueCommand(meta, *issueTitle, *issueBody)
	}

	if *milestoneTitle != "" {
		printMilestoneCommand(meta, *milestoneTitle, *milestoneDue)
	}

	fmt.Printf("GitHub metadata captured. Remote=%s (%s) DefaultBranch=%s\n", meta.RemoteName, meta.RepoSlug, meta.DefaultBranch)
}

type gitInfo struct {
	RemoteName    string
	RemoteURL     string
	RepoSlug      string
	DefaultBranch string
}

func detectGit(project string) (gitInfo, error) {
	info := gitInfo{RemoteName: "origin", DefaultBranch: "main"}

	remoteURL, err := runGit(project, "remote", "get-url", info.RemoteName)
	if err == nil {
		info.RemoteURL = strings.TrimSpace(remoteURL)
		info.RepoSlug = parseRepoSlug(info.RemoteURL)
	}

	ref, err := runGit(project, "symbolic-ref", "refs/remotes/"+info.RemoteName+"/HEAD")
	if err == nil {
		parts := strings.Split(strings.TrimSpace(ref), "/")
		info.DefaultBranch = parts[len(parts)-1]
	} else {
		// fallback: try config
		if branch, err := runGit(project, "config", "--get", "init.defaultBranch"); err == nil {
			info.DefaultBranch = strings.TrimSpace(branch)
		}
	}

	if info.RepoSlug == "" {
		return info, errors.New("unable to parse repo slug")
	}
	return info, nil
}

func parseRepoSlug(remote string) string {
	remote = strings.TrimSpace(remote)
	if strings.HasPrefix(remote, "git@") {
		remote = strings.TrimPrefix(remote, "git@")
		remote = strings.Replace(remote, ":", "/", 1)
	}
	remote = strings.TrimSuffix(remote, ".git")
	if strings.HasPrefix(remote, "https://") {
		remote = strings.TrimPrefix(remote, "https://")
	}
	if strings.HasPrefix(remote, "http://") {
		remote = strings.TrimPrefix(remote, "http://")
	}
	parts := strings.Split(remote, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-2] + "/" + parts[len(parts)-1]
	}
	return ""
}

func runGit(project string, args ...string) (string, error) {
	cmd := exec.Command("git", append([]string{"-C", project}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func extractSuccessMetrics(prdPath string) []string {
	data, err := os.ReadFile(prdPath)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(data), "\n")
	var metrics []string
	inSection := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "## Success Metrics") || strings.HasPrefix(trimmed, "### Success Metrics") {
			inSection = true
			continue
		}
		if inSection {
			if strings.HasPrefix(trimmed, "## ") && !strings.HasPrefix(trimmed, "## Success") {
				break
			}
			if strings.HasPrefix(trimmed, "-") {
				metrics = append(metrics, trimmed)
			}
		}
	}
	return metrics
}

func updateReadmeKPIs(project string, metrics []string) error {
	readmePath := filepath.Join(project, "README.md")
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}
	startMarker := "<!-- KPIS:START -->"
	endMarker := "<!-- KPIS:END -->"
	content := string(data)
	block := renderKPIBlock(metrics, startMarker, endMarker)

	if strings.Contains(content, startMarker) && strings.Contains(content, endMarker) {
		regex := regexp.MustCompile(regexp.QuoteMeta(startMarker) + "(?s).*?" + regexp.QuoteMeta(endMarker))
		content = regex.ReplaceAllString(content, block)
	} else {
		content += "\n## 📈 KPIs & Targets\n" + block + "\n"
	}

	return os.WriteFile(readmePath, []byte(content), 0644)
}

func renderKPIBlock(metrics []string, start, end string) string {
	var builder strings.Builder
	builder.WriteString(start)
	builder.WriteString("\n")
	if len(metrics) == 0 {
		builder.WriteString("_No success metrics detected. Update PRD.md to keep README in sync._\n")
	} else {
		for _, metric := range metrics {
			builder.WriteString(metric)
			builder.WriteString("\n")
		}
	}
	builder.WriteString(end)
	builder.WriteString("\n")
	return builder.String()
}

func writeMeta(project string, meta GitMeta) error {
	historyDir := filepath.Join(project, "Docs", "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		return err
	}
	path := filepath.Join(historyDir, "github-meta.json")
	payload, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, payload, 0644)
}

func printIssueCommand(meta GitMeta, title, body string) {
	if meta.RepoSlug == "" {
		fmt.Println("Unable to compose issue command (missing repo slug)")
		return
	}
	fmt.Println("Run the following command to create an issue:")
	builder := strings.Builder{}
	builder.WriteString("gh issue create --repo ")
	builder.WriteString(meta.RepoSlug)
	builder.WriteString(" --title \"")
	builder.WriteString(title)
	builder.WriteString("\"")
	if strings.TrimSpace(body) != "" {
		builder.WriteString(" --body \"")
		builder.WriteString(body)
		builder.WriteString("\"")
	}
	fmt.Println(builder.String())
}

func printMilestoneCommand(meta GitMeta, title, due string) {
	if meta.RepoSlug == "" {
		fmt.Println("Unable to compose milestone command (missing repo slug)")
		return
	}
	fmt.Println("Run the following command to create a milestone:")
	builder := strings.Builder{}
	builder.WriteString("gh api repos/")
	builder.WriteString(meta.RepoSlug)
	builder.WriteString("/milestones --method POST -f title=\"")
	builder.WriteString(title)
	builder.WriteString("\"")
	if strings.TrimSpace(due) != "" {
		builder.WriteString(" -f due_on=\"")
		builder.WriteString(due)
		builder.WriteString("\"")
	}
	fmt.Println(builder.String())
}
