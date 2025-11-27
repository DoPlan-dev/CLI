package cli

import (
	"bytes"
	"testing"
)

func TestRootCmd(t *testing.T) {
	rootCmd := GetRootCmd()

	if rootCmd.Use != AppName {
		t.Errorf("RootCmd.Use = %v, want %v", rootCmd.Use, AppName)
	}

	if rootCmd.Version != Version() {
		t.Errorf("RootCmd.Version = %v, want %v", rootCmd.Version, Version())
	}
}

func TestVersionFlag(t *testing.T) {
	rootCmd := GetRootCmd()
	rootCmd.SetArgs([]string{"--version"})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	output := buf.String()
	expected := AppName + " version " + Version() + "\n"
	if output != expected {
		t.Errorf("Version output = %q, want %q", output, expected)
	}
}

func TestHelpFlag(t *testing.T) {
	rootCmd := GetRootCmd()
	rootCmd.SetArgs([]string{"--help"})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	output := buf.String()
	if len(output) == 0 {
		t.Error("Help output should not be empty")
	}

	// Check that help contains expected text
	if !contains(output, "DoPlan CLI") {
		t.Error("Help output should contain 'DoPlan CLI'")
	}
}

func TestDefaultCommand(t *testing.T) {
	rootCmd := GetRootCmd()
	rootCmd.SetArgs([]string{})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	output := buf.String()
	if len(output) == 0 {
		t.Error("Default command output should not be empty")
	}

	// Check that output contains version info
	if !contains(output, "DoPlan CLI") {
		t.Error("Default command output should contain 'DoPlan CLI'")
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestGetIDECommand(t *testing.T) {
	testCases := []struct {
		ide  string
		want string
	}{
		{"Cursor", "cursor"},
		{"Claude Code", "claude"},
		{"Antigravity", "antigravity"},
		{"Windsurf", "windsurf"},
		{"Cline", "cline"},
		{"OpenCode", "opencode"},
		{"Unknown IDE", "code"}, // Default case
		{"", "code"},            // Empty string
		{"VS Code", "code"},     // Not in list
	}

	for _, tc := range testCases {
		t.Run(tc.ide, func(t *testing.T) {
			got := getIDECommand(tc.ide)
			if got != tc.want {
				t.Errorf("getIDECommand(%q) = %q, want %q", tc.ide, got, tc.want)
			}
		})
	}
}

func TestConstants(t *testing.T) {
	// Version is now a function that returns the version
	// Default version is "dev" if not set at build time
	version := Version()
	if version == "" {
		t.Error("Version() should not return empty string")
	}

	if AppName != "doplan" {
		t.Errorf("AppName = %q, want %q", AppName, "doplan")
	}
}
