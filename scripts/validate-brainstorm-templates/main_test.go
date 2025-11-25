// +build ignore

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestValidateBrainstormTemplates(t *testing.T) {
	// This is an integration test that runs the script as a subprocess
	// since it's a main package that exits
	
	tmpDir := t.TempDir()
	
	// Create a minimal project structure with templates
	planDir := filepath.Join(tmpDir, ".plan", "templates", "brainstorm")
	os.MkdirAll(planDir, 0755)
	
	// Create required templates
	requiredTemplates := []string{
		"phase-01-vision.md",
		"phase-02-audience.md",
		"phase-03-experience.md",
		"phase-04-content.md",
		"phase-05-marketing.md",
		"phase-06-delivery.md",
		"CONFIRMATION_TEMPLATE.md",
		"TEMPLATE_BRAINSTORM.md",
		"README.md",
	}
	
	for _, template := range requiredTemplates {
		content := "# " + template + "\n\n- Test question"
		os.WriteFile(filepath.Join(planDir, template), []byte(content), 0644)
	}
	
	// Change to temp directory
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tmpDir)
	
	// Run the script
	cmd := exec.Command("go", "run", "../../scripts/validate-brainstorm-templates/main.go")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		t.Logf("Script output: %s", string(output))
		t.Fatalf("Script failed: %v", err)
	}
	
	// Check that output contains success indicators
	outputStr := string(output)
	if !contains(outputStr, "Found") && !contains(outputStr, "complete") {
		t.Errorf("Expected success output, got: %s", outputStr)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || 
		s[len(s)-len(substr):] == substr || 
		containsHelper(s, substr)))))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

