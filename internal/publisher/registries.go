package publisher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DetectPackageType detects the package type for publishing
func DetectPackageType(projectRoot string) (string, error) {
	// Check for package.json (npm)
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		return "npm", nil
	}

	// Check for go.mod (could be Homebrew)
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
		return "homebrew", nil
	}

	// Check for Cargo.toml (Rust/Cargo)
	if _, err := os.Stat(filepath.Join(projectRoot, "Cargo.toml")); err == nil {
		return "cargo", nil
	}

	return "unknown", fmt.Errorf("could not detect package type")
}

// PublishToNPM publishes to npm
func PublishToNPM(projectRoot string) error {
	// Check if npm is available
	if _, err := exec.LookPath("npm"); err != nil {
		return fmt.Errorf("npm not found. Please install Node.js")
	}

	fmt.Println("📦 Publishing to npm...")

	// Check if logged in
	cmd := exec.Command("npm", "whoami")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("not logged in to npm. Run 'npm login' first")
	}
	username := strings.TrimSpace(string(output))
	fmt.Printf("✅ Logged in as: %s\n", username)

	// Publish
	cmd = exec.Command("npm", "publish")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("npm publish failed: %w", err)
	}

	fmt.Println("✅ Successfully published to npm!")
	return nil
}

// PublishToHomebrew generates Homebrew formula
func PublishToHomebrew(projectRoot string) error {
	fmt.Println("🍺 Generating Homebrew formula...")

	// Check for go.mod
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err != nil {
		return fmt.Errorf("Homebrew publishing requires a Go project")
	}

	// Read go.mod to get module name
	goModPath := filepath.Join(projectRoot, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("failed to read go.mod: %w", err)
	}

	moduleName := extractModuleName(string(data))
	if moduleName == "" {
		return fmt.Errorf("could not determine module name from go.mod")
	}

	formulaPath := filepath.Join(projectRoot, fmt.Sprintf("%s.rb", moduleName))
	formula := generateHomebrewFormula(moduleName, projectRoot)

	if err := os.WriteFile(formulaPath, []byte(formula), 0644); err != nil {
		return fmt.Errorf("failed to write formula: %w", err)
	}

	fmt.Printf("✅ Generated Homebrew formula: %s\n", formulaPath)
	fmt.Println("💡 Next steps:")
	fmt.Println("   1. Create a GitHub release with the binary")
	fmt.Println("   2. Fork homebrew-core or create your own tap")
	fmt.Println("   3. Add the formula to your tap")
	fmt.Println("   4. Submit PR to homebrew-core (if applicable)")

	return nil
}

// PublishToScoop generates Scoop manifest
func PublishToScoop(projectRoot string) error {
	fmt.Println("🥄 Generating Scoop manifest...")

	manifestPath := filepath.Join(projectRoot, "scoop-manifest.json")
	manifest := generateScoopManifest(projectRoot)

	if err := os.WriteFile(manifestPath, []byte(manifest), 0644); err != nil {
		return fmt.Errorf("failed to write manifest: %w", err)
	}

	fmt.Printf("✅ Generated Scoop manifest: %s\n", manifestPath)
	fmt.Println("💡 Next steps:")
	fmt.Println("   1. Create a GitHub release with the binary")
	fmt.Println("   2. Fork scoop-bucket or create your own bucket")
	fmt.Println("   3. Add the manifest to your bucket")

	return nil
}

// PublishToWinget generates Winget manifest
func PublishToWinget(projectRoot string) error {
	fmt.Println("🪟 Generating Winget manifest...")

	manifestDir := filepath.Join(projectRoot, "winget-manifest")
	if err := os.MkdirAll(manifestDir, 0755); err != nil {
		return fmt.Errorf("failed to create manifest directory: %w", err)
	}

	manifest := generateWingetManifest(projectRoot)
	manifestPath := filepath.Join(manifestDir, "manifest.yaml")

	if err := os.WriteFile(manifestPath, []byte(manifest), 0644); err != nil {
		return fmt.Errorf("failed to write manifest: %w", err)
	}

	fmt.Printf("✅ Generated Winget manifest: %s\n", manifestPath)
	fmt.Println("💡 Next steps:")
	fmt.Println("   1. Create a GitHub release with the binary")
	fmt.Println("   2. Fork winget-pkgs or create your own repository")
	fmt.Println("   3. Add the manifest to your repository")
	fmt.Println("   4. Submit PR to winget-pkgs (if applicable)")

	return nil
}

func extractModuleName(goModContent string) string {
	lines := strings.Split(goModContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}

func generateHomebrewFormula(moduleName, projectRoot string) string {
	// Basic Homebrew formula template
	return fmt.Sprintf(`class %s < Formula
  desc "DoPlan CLI tool"
  homepage "https://github.com/DoPlan-dev/CLI"
  url "https://github.com/DoPlan-dev/CLI/archive/refs/tags/v0.0.19-beta.tar.gz"
  sha256 "PLACEHOLDER_SHA256"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"doplan", "./cmd/doplan"
  end

  test do
    system "#{bin}/doplan", "--version"
  end
end`, strings.Title(moduleName))
}

func generateScoopManifest(projectRoot string) string {
	return `{
  "version": "0.0.19-beta",
  "description": "DoPlan CLI tool",
  "homepage": "https://github.com/DoPlan-dev/CLI",
  "license": "MIT",
  "architecture": {
    "64bit": {
      "url": "https://github.com/DoPlan-dev/CLI/releases/download/v0.0.19-beta/doplan-windows-amd64.zip",
      "hash": "PLACEHOLDER_HASH"
    }
  },
  "bin": "doplan.exe",
  "checkver": "github",
  "autoupdate": {
    "architecture": {
      "64bit": {
        "url": "https://github.com/DoPlan-dev/CLI/releases/download/v$version/doplan-windows-amd64.zip"
      }
    }
  }
}`
}

func generateWingetManifest(projectRoot string) string {
	return `# yaml-language-server: $schema=https://aka.ms/winget-manifest.schema.2.0
PackageIdentifier: DoPlan.CLI
PackageVersion: 0.0.19-beta
PackageName: DoPlan CLI
Publisher: DoPlan
Description: DoPlan Project Workflow Manager CLI
License: MIT
Copyright: Copyright (c) 2024 DoPlan
Homepage: https://github.com/DoPlan-dev/CLI
Tags:
  - workflow
  - project-management
  - cli
Installers:
  - Architecture: x64
    InstallerType: zip
    InstallerUrl: https://github.com/DoPlan-dev/CLI/releases/download/v0.0.19-beta/doplan-windows-amd64.zip
    InstallerSha256: PLACEHOLDER_SHA256
ManifestType: version
ManifestVersion: 1.4.0`
}
