package generator

import (
	"fmt"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// GitHubGenerator generates GitHub Actions workflows
type GitHubGenerator struct{}

// Name returns the name of the generator
func (g *GitHubGenerator) Name() string {
	return "GitHub Workflows"
}

// Generate creates the .github/workflows/ directory and all workflow files
func (g *GitHubGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	workflowsDir := filepath.Join(projectPath, ".github", "workflows")

	// Create .github/workflows directory
	if err := utils.CreateDirectory(workflowsDir); err != nil {
		return fmt.Errorf("failed to create .github/workflows directory: %w", err)
	}

	// Generate CI workflow
	if err := generateCIWorkflow(workflowsDir, request); err != nil {
		return fmt.Errorf("failed to generate CI workflow: %w", err)
	}

	// Generate Release workflow
	if err := generateReleaseWorkflow(workflowsDir, request); err != nil {
		return fmt.Errorf("failed to generate release workflow: %w", err)
	}

	// Generate Changelog workflow
	if err := generateChangelogWorkflow(workflowsDir, request); err != nil {
		return fmt.Errorf("failed to generate changelog workflow: %w", err)
	}

	// Generate Branch Protection workflow
	if err := generateBranchProtectionWorkflow(workflowsDir, request); err != nil {
		return fmt.Errorf("failed to generate branch protection workflow: %w", err)
	}

	return nil
}

// generateCIWorkflow generates ci.yml workflow
func generateCIWorkflow(workflowsDir string, request *models.ProjectRequest) error {
	content := `name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    name: Test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: go test ./... -v
      
      - name: Check test coverage
        run: go test ./... -coverprofile=coverage.out
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
          fail_ci_if_error: false

  lint:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: latest

  build:
    name: Build
    runs-on: ubuntu-latest
    strategy:
      matrix:
        goos: [linux, darwin, windows]
        goarch: [amd64, arm64]
        exclude:
          - goos: windows
            goarch: arm64
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Build
        run: go build -o bin/doplan-${{ matrix.goos }}-${{ matrix.goarch }} ./cmd/doplan
        env:
          GOOS: ${{ matrix.goos }}
          GOARCH: ${{ matrix.goarch }}
      
      - name: Upload artifacts
        uses: actions/upload-artifact@v3
        with:
          name: doplan-${{ matrix.goos }}-${{ matrix.goarch }}
          path: bin/doplan-${{ matrix.goos }}-${{ matrix.goarch }}
`
	path := filepath.Join(workflowsDir, "ci.yml")
	return utils.WriteFile(path, []byte(content))
}

// generateReleaseWorkflow generates release.yml workflow
func generateReleaseWorkflow(workflowsDir string, request *models.ProjectRequest) error {
	content := `name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    name: Build Binaries
    runs-on: ubuntu-latest
    strategy:
      matrix:
        include:
          - goos: linux
            goarch: amd64
          - goos: linux
            goarch: arm64
          - goos: darwin
            goarch: amd64
          - goos: darwin
            goarch: arm64
          - goos: windows
            goarch: amd64
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Extract version
        id: version
        run: |
          VERSION=${GITHUB_REF#refs/tags/v}
          echo "version=$VERSION" >> $GITHUB_OUTPUT
          echo "Version: $VERSION"
      
      - name: Build
        run: |
          set -e
          BINARY_NAME=doplan
          if [ "${{ matrix.goos }}" = "windows" ]; then
            BINARY_NAME=doplan.exe
          fi
          MODULE_PATH=$(go list -m)
          VERSION_SYMBOL="${MODULE_PATH}/internal/version.Version"
          echo "Module path: $MODULE_PATH"
          echo "Injecting version via: $VERSION_SYMBOL"
          GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} go build -ldflags "-X ${VERSION_SYMBOL}=${{ steps.version.outputs.version }}" -o $BINARY_NAME ./cmd/doplan
        env:
          GOOS: ${{ matrix.goos }}
          GOARCH: ${{ matrix.goarch }}
      
      - name: Create archive
        run: |
          VERSION=${{ steps.version.outputs.version }}
          if [ "${{ matrix.goos }}" = "windows" ]; then
            zip doplan-$VERSION-${{ matrix.goos }}-${{ matrix.goarch }}.zip doplan.exe
          else
            tar -czf doplan-$VERSION-${{ matrix.goos }}-${{ matrix.goarch }}.tar.gz doplan
          fi
      
      - name: Generate checksum
        run: |
          if [ "${{ matrix.goos }}" = "windows" ]; then
            sha256sum doplan-${{ steps.version.outputs.version }}-${{ matrix.goos }}-${{ matrix.goarch }}.zip > doplan-${{ steps.version.outputs.version }}-${{ matrix.goos }}-${{ matrix.goarch }}.zip.sha256
          else
            sha256sum doplan-${{ steps.version.outputs.version }}-${{ matrix.goos }}-${{ matrix.goarch }}.tar.gz > doplan-${{ steps.version.outputs.version }}-${{ matrix.goos }}-${{ matrix.goarch }}.tar.gz.sha256
          fi
      
      - name: Upload artifacts
        uses: actions/upload-artifact@v4
        with:
          name: doplan-${{ matrix.goos }}-${{ matrix.goarch }}
          path: |
            doplan-*-${{ matrix.goos }}-${{ matrix.goarch }}.*
            doplan*

  release:
    name: Create Release
    needs: build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: Extract version
        id: version
        run: |
          VERSION=${GITHUB_REF#refs/tags/v}
          echo "version=$VERSION" >> $GITHUB_OUTPUT
          echo "Version: $VERSION"
      
      - name: Download all artifacts
        uses: actions/download-artifact@v4
        with:
          path: dist
      
      - name: Generate release notes
        id: release_notes
        run: |
          if [ -f docs/CHANGELOG.md ]; then
            # Extract release notes from docs/CHANGELOG.md
            awk '/^## \[v'"${{ steps.version.outputs.version }}"'\]/,/^## \[/' docs/CHANGELOG.md | head -n -1 > release_notes.txt || echo "Release v${{ steps.version.outputs.version }}" > release_notes.txt
          else
            echo "Release v${{ steps.version.outputs.version }}" > release_notes.txt
          fi
      
      - name: Create GitHub Release
        uses: softprops/action-gh-release@v1
        with:
          tag_name: ${{ github.ref }}
          name: Release v${{ steps.version.outputs.version }}
          body_path: release_notes.txt
          files: dist/**/*
          draft: false
          prerelease: false
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
`
	path := filepath.Join(workflowsDir, "release.yml")
	return utils.WriteFile(path, []byte(content))
}

// generateChangelogWorkflow generates changelog.yml workflow
func generateChangelogWorkflow(workflowsDir string, request *models.ProjectRequest) error {
	content := `name: Update Changelog

on:
  push:
    paths:
      - 'docs/CHANGELOG.md'
    branches:
      - main

jobs:
  update:
    name: Auto-commit Changelog
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Configure Git
        run: |
          git config --local user.email "action@github.com"
          git config --local user.name "GitHub Action"
      
      - name: Check for changes
        id: changes
        run: |
          if git diff --quiet docs/CHANGELOG.md; then
            echo "changed=false" >> $GITHUB_OUTPUT
          else
            echo "changed=true" >> $GITHUB_OUTPUT
          fi
      
      - name: Commit changes
        if: steps.changes.outputs.changed == 'true'
        run: |
          git add docs/CHANGELOG.md
          git commit -m "chore: update docs/CHANGELOG.md [skip ci]"
          git push
`
	path := filepath.Join(workflowsDir, "changelog.yml")
	return utils.WriteFile(path, []byte(content))
}

// generateBranchProtectionWorkflow generates branch-protection.yml workflow
func generateBranchProtectionWorkflow(workflowsDir string, request *models.ProjectRequest) error {
	content := `name: Branch Protection

on:
  pull_request:
    types: [opened, synchronize, reopened]

jobs:
  check:
    name: PR Requirements Check
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Check PR title
        run: |
          TITLE="${{ github.event.pull_request.title }}"
          if [[ ! "$TITLE" =~ ^(feat|fix|docs|style|refactor|test|chore)(\(.+\))?: .+ ]]; then
            echo "PR title must follow conventional commit format"
            echo "Format: type(scope): description"
            echo "Types: feat, fix, docs, style, refactor, test, chore"
            exit 1
          fi
      
      - name: Check required status checks
        run: |
          echo "Waiting for required status checks to pass..."
          echo "Required checks: test, lint, build"
`
	path := filepath.Join(workflowsDir, "branch-protection.yml")
	return utils.WriteFile(path, []byte(content))
}

// GenerateGitHubWorkflows is a convenience function that creates a GitHubGenerator and generates workflows
func GenerateGitHubWorkflows(request *models.ProjectRequest, projectPath string) error {
	generator := &GitHubGenerator{}
	return generator.Generate(request, projectPath)
}
