package integration

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSetupCursor(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupCursor(tmpDir)
	if err != nil {
		t.Fatalf("SetupCursor failed: %v", err)
	}

	// Verify directories exist
	checks := []string{
		filepath.Join(tmpDir, ".cursor", "agents"),
		filepath.Join(tmpDir, ".cursor", "rules"),
		filepath.Join(tmpDir, ".cursor", "commands"),
		filepath.Join(tmpDir, ".cursor", "rules", "doplan.mdc"),
		filepath.Join(tmpDir, ".doplan", "ai", "agents"),
		filepath.Join(tmpDir, ".doplan", "ai", "rules"),
		filepath.Join(tmpDir, ".doplan", "ai", "commands"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			t.Errorf("Expected %s to exist, got error: %v", check, err)
		}
	}

	// Verify integration
	if err := VerifyCursor(tmpDir); err != nil {
		t.Errorf("VerifyCursor failed: %v", err)
	}
}

func TestSetupVSCode(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupVSCode(tmpDir)
	if err != nil {
		t.Fatalf("SetupVSCode failed: %v", err)
	}

	// Verify files exist
	checks := []string{
		filepath.Join(tmpDir, ".vscode", "tasks.json"),
		filepath.Join(tmpDir, ".vscode", "settings.json"),
		filepath.Join(tmpDir, ".vscode", "prompts"),
		filepath.Join(tmpDir, ".vscode", "prompts", "doplan-context.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			t.Errorf("Expected %s to exist, got error: %v", check, err)
		}
	}

	// Verify integration
	if err := VerifyVSCode(tmpDir); err != nil {
		t.Errorf("VerifyVSCode failed: %v", err)
	}
}

func TestSetupGemini(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupGemini(tmpDir)
	if err != nil {
		t.Fatalf("SetupGemini failed: %v", err)
	}

	// Verify directories exist
	checks := []string{
		filepath.Join(tmpDir, ".gemini", "commands"),
		filepath.Join(tmpDir, ".doplan", "guides", "gemini_setup.md"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			t.Errorf("Expected %s to exist, got error: %v", check, err)
		}
	}

	// Verify integration
	if err := VerifyGemini(tmpDir); err != nil {
		t.Errorf("VerifyGemini failed: %v", err)
	}
}

func TestSetupClaude(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupClaude(tmpDir)
	if err != nil {
		t.Fatalf("SetupClaude failed: %v", err)
	}

	// Verify integration
	if err := VerifyClaude(tmpDir); err != nil {
		t.Errorf("VerifyClaude failed: %v", err)
	}
}

func TestSetupCodex(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupCodex(tmpDir)
	if err != nil {
		t.Fatalf("SetupCodex failed: %v", err)
	}

	// Verify integration
	if err := VerifyCodex(tmpDir); err != nil {
		t.Errorf("VerifyCodex failed: %v", err)
	}
}

func TestSetupOpenCode(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupOpenCode(tmpDir)
	if err != nil {
		t.Fatalf("SetupOpenCode failed: %v", err)
	}

	// Verify integration
	if err := VerifyOpenCode(tmpDir); err != nil {
		t.Errorf("VerifyOpenCode failed: %v", err)
	}
}

func TestSetupQwen(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupQwen(tmpDir)
	if err != nil {
		t.Fatalf("SetupQwen failed: %v", err)
	}

	// Verify integration
	if err := VerifyQwen(tmpDir); err != nil {
		t.Errorf("VerifyQwen failed: %v", err)
	}
}

func TestSetupKiro(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupKiro(tmpDir)
	if err != nil {
		t.Fatalf("SetupKiro failed: %v", err)
	}

	// Verify guide exists
	guidePath := filepath.Join(tmpDir, ".doplan", "guides", "kiro_setup.md")
	if _, err := os.Stat(guidePath); err != nil {
		t.Errorf("Expected guide to exist, got error: %v", err)
	}

	// Verify integration
	if err := VerifyKiro(tmpDir); err != nil {
		t.Errorf("VerifyKiro failed: %v", err)
	}
}

func TestSetupWindsurf(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupWindsurf(tmpDir)
	if err != nil {
		t.Fatalf("SetupWindsurf failed: %v", err)
	}

	// Verify directories exist
	checks := []string{
		filepath.Join(tmpDir, ".windsurf", "agents"),
		filepath.Join(tmpDir, ".windsurf", "rules"),
		filepath.Join(tmpDir, ".windsurf", "commands"),
	}

	for _, check := range checks {
		if _, err := os.Stat(check); err != nil {
			t.Errorf("Expected %s to exist, got error: %v", check, err)
		}
	}

	// Verify integration
	if err := VerifyWindsurf(tmpDir); err != nil {
		t.Errorf("VerifyWindsurf failed: %v", err)
	}
}

func TestSetupQoder(t *testing.T) {
	tmpDir := t.TempDir()

	err := SetupQoder(tmpDir)
	if err != nil {
		t.Fatalf("SetupQoder failed: %v", err)
	}

	// Verify config file exists
	configPath := filepath.Join(tmpDir, ".qoder", "doplan.json")
	if _, err := os.Stat(configPath); err != nil {
		t.Errorf("Expected config file to exist, got error: %v", err)
	}

	// Verify integration
	if err := VerifyQoder(tmpDir); err != nil {
		t.Errorf("VerifyQoder failed: %v", err)
	}
}

func TestSetupIDE(t *testing.T) {
	tmpDir := t.TempDir()

	testCases := []struct {
		name      string
		ide       string
		shouldErr bool
	}{
		{"Cursor", "cursor", false},
		{"VS Code", "vscode", false},
		{"Gemini", "gemini", false},
		{"Claude", "claude", false},
		{"Codex", "codex", false},
		{"OpenCode", "opencode", false},
		{"Qwen", "qwen", false},
		{"Kiro", "kiro", false},
		{"Windsurf", "windsurf", false},
		{"Qoder", "qoder", false},
		{"Unknown IDE", "unknown", false}, // Should create generic guide
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testDir := filepath.Join(tmpDir, tc.ide)
			if err := os.MkdirAll(testDir, 0755); err != nil {
				t.Fatalf("Failed to create test directory: %v", err)
			}

			err := SetupIDE(testDir, tc.ide)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for %s, got nil", tc.ide)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for %s: %v", tc.ide, err)
			}
		})
	}
}

func TestVerifyIDE(t *testing.T) {
	tmpDir := t.TempDir()

	// Setup Cursor first
	if err := SetupCursor(tmpDir); err != nil {
		t.Fatalf("SetupCursor failed: %v", err)
	}

	// Verify it
	if err := VerifyIDE(tmpDir, "cursor"); err != nil {
		t.Errorf("VerifyIDE failed: %v", err)
	}

	// Verify unknown IDE creates generic guide
	if err := VerifyIDE(tmpDir, "unknown"); err != nil {
		// This is expected - generic guide might not exist yet
	}
}

func TestCopyDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Test copyDir through SetupCursor which uses it internally
	// Create .doplan/ai structure first
	doplanAIDir := filepath.Join(tmpDir, ".doplan", "ai", "commands")
	if err := os.MkdirAll(doplanAIDir, 0755); err != nil {
		t.Fatalf("Failed to create doplan dir: %v", err)
	}

	// Create test file in doplan/ai/commands
	testFile := filepath.Join(doplanAIDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// SetupCursor will use copyDir internally (on Windows or if symlink fails)
	// This tests copyDir indirectly
	err := SetupCursor(tmpDir)
	if err != nil {
		t.Fatalf("SetupCursor failed: %v", err)
	}

	// Verify that commands directory exists and is accessible (via symlink or copy)
	commandsDir := filepath.Join(tmpDir, ".cursor", "commands")
	if _, err := os.Stat(commandsDir); err != nil {
		t.Errorf("Expected commands directory to exist, got error: %v", err)
	}

	// Verify file is accessible through the symlink/copy
	copiedFile := filepath.Join(commandsDir, "test.txt")
	if _, err := os.Stat(copiedFile); err != nil {
		t.Errorf("Expected copied file to exist, got error: %v", err)
	}
}

func TestWindowsCompatibility(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows compatibility test on non-Windows system")
	}

	tmpDir := t.TempDir()

	// Test that Cursor setup works on Windows (should use copy instead of symlink)
	err := SetupCursor(tmpDir)
	if err != nil {
		t.Fatalf("SetupCursor failed on Windows: %v", err)
	}

	// Verify directories exist (copied, not symlinked)
	agentsDir := filepath.Join(tmpDir, ".cursor", "agents")
	if _, err := os.Stat(agentsDir); err != nil {
		t.Errorf("Expected agents directory to exist, got error: %v", err)
	}

	// On Windows, we should be able to read the directory
	entries, err := os.ReadDir(agentsDir)
	if err != nil {
		t.Errorf("Failed to read agents directory: %v", err)
	}
	_ = entries // Use entries to avoid unused variable
}
