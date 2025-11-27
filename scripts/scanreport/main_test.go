package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DoPlan-dev/CLI/internal/progress"
)

func resetScanreportDeps() {
	summarizeStateHistoryFunc = summarizeStateHistory
	summarizeProgressFunc = summarizeProgress
	renderVisualsFunc = renderVisuals
	renderDependencyAuditFunc = renderDependencyAudit
}

func TestRunScanreportNoReports(t *testing.T) {
	defer resetScanreportDeps()

	tmp := t.TempDir()
	if err := os.MkdirAll(filepath.Join(tmp, ".do", "reports"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	code := run([]string{"--project", tmp}, out, errOut)
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d (stderr=%s)", code, errOut.String())
	}
	if !strings.Contains(out.String(), "No scan reports found") {
		t.Fatalf("expected no reports message, got %s", out.String())
	}
}

func TestRunScanreportGeneratesDiff(t *testing.T) {
	defer resetScanreportDeps()

	tmp := t.TempDir()
	reportsDir := filepath.Join(tmp, ".do", "reports")
	if err := os.MkdirAll(reportsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	reportContent := `**Scan Date**: 2025-01-0X
**Project**: Demo
**Report Type**: Standard

## Executive Summary
- Highlight

## 6. Findings & Risks
- Risk

## 7. Recommended Next Actions
- Next`

	if err := os.WriteFile(filepath.Join(reportsDir, "SCAN_REPORT_20250101.md"), []byte(strings.Replace(reportContent, "0X", "1", 1)), 0o644); err != nil {
		t.Fatalf("write report: %v", err)
	}
	if err := os.WriteFile(filepath.Join(reportsDir, "SCAN_REPORT_20250102.md"), []byte(strings.Replace(reportContent, "0X", "2", 1)), 0o644); err != nil {
		t.Fatalf("write report: %v", err)
	}

	summarizeStateHistoryFunc = func(string) string { return "State summary\n" }
	summarizeProgressFunc = func(string) (string, *progress.Report) {
		return "Progress summary\n", &progress.Report{}
	}
	renderVisualsFunc = func(*progress.Report) string { return "Visuals\n" }
	renderDependencyAuditFunc = func(string) string { return "Dependencies\n" }

	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	if code := run([]string{"--project", tmp}, out, errOut); code != 0 {
		t.Fatalf("expected exit code 0, got %d (stderr=%s)", code, errOut.String())
	}

	meta1 := filepath.Join(reportsDir, "SCAN_REPORT_20250101.json")
	meta2 := filepath.Join(reportsDir, "SCAN_REPORT_20250102.json")
	for _, path := range []string{meta1, meta2} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected metadata file %s: %v", path, err)
		}
	}

	files, err := filepath.Glob(filepath.Join(reportsDir, "SCAN_DIFF_*.md"))
	if err != nil || len(files) == 0 {
		t.Fatalf("expected diff file, got err=%v files=%v", err, files)
	}
}
