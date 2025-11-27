package main

import (
	"bytes"
	"errors"
	"testing"
	"time"
)

func resetGitMetaDeps() {
	detectGitFunc = detectGit
	extractSuccessMetricsFn = extractSuccessMetrics
	writeMetaFunc = writeMeta
	updateReadmeKPIsFunc = updateReadmeKPIs
	printIssueCommandFunc = printIssueCommand
	printMilestoneFunc = printMilestoneCommand
	timeNow = time.Now
}

func TestRunGithubMetaSuccess(t *testing.T) {
	defer resetGitMetaDeps()
	var savedMeta GitMeta
	detectGitFunc = func(string) (gitInfo, error) {
		return gitInfo{
			RemoteName:    "origin",
			RemoteURL:     "git@github.com:acme/repo.git",
			RepoSlug:      "acme/repo",
			DefaultBranch: "main",
		}, nil
	}
	extractSuccessMetricsFn = func(string) []string { return []string{"- KPI"} }
	writeMetaFunc = func(_ string, meta GitMeta) error {
		savedMeta = meta
		return nil
	}
	updateReadmeKPIsFunc = func(string, []string) error { return nil }
	printIssueCommandFunc = func(GitMeta, string, string) {}
	printMilestoneFunc = func(GitMeta, string, string) {}
	timeNow = func() time.Time { return time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC) }

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	code := run([]string{"--project", ".", "--sync-readme"}, out, errOut)
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d (stderr=%s)", code, errOut.String())
	}
	if savedMeta.RepoSlug != "acme/repo" || savedMeta.UpdatedAt == "" {
		t.Fatalf("expected meta to be saved, got %+v", savedMeta)
	}
	if out.Len() == 0 {
		t.Fatal("expected output")
	}
}

func TestRunGithubMetaDetectError(t *testing.T) {
	defer resetGitMetaDeps()
	detectGitFunc = func(string) (gitInfo, error) { return gitInfo{}, errors.New("boom") }
	writeMetaFunc = func(string, GitMeta) error { return nil }
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	code := run([]string{"--project", "."}, out, errOut)
	if code != 0 {
		t.Fatalf("expected exit 0 even on detect error, got %d", code)
	}
	if errOut.Len() == 0 {
		t.Fatal("expected warning output")
	}
}
