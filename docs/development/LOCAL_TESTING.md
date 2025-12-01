# Local Testing Guide

## Quick Start: Build and Run

### 1. Build the CLI

```bash
# Build for your current platform
make build

# Or manually:
go build -o doplan ./cmd/doplan
```

### 2. Run the CLI Interactively

```bash
# Run the wizard interactively
./doplan

# The wizard will prompt you for:
# - Project name
# - IDE selection
# - And generate the complete project
```

### 3. Test with a Temporary Directory

```bash
# Create a test directory
mkdir -p /tmp/doplan-test
cd /tmp/doplan-test

# Run the CLI
/path/to/GoPlan-CLI/doplan

# Or if installed in PATH:
doplan
```

## Automated Testing

### Run All Tests

```bash
# Run all tests
make test

# Or:
go test ./... -v
```

### Run End-to-End Tests

```bash
# Run all end-to-end tests
go test ./internal/generator -v -run TestEndToEnd

# Run specific E2E test
go test ./internal/generator -v -run TestEndToEnd_CompleteWizardFlow
```

### Run Tests with Coverage

```bash
# Generate coverage report
make test-coverage

# Or manually:
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

# Open coverage report (macOS)
open coverage.html
```

### Test Individual Components

```bash
# Test TUI wizard
go test ./internal/tui/... -v

# Test generators
go test ./internal/generator/... -v

# Test CLI commands
go test ./internal/cli/... -v

# Test rules library
go test ./internal/rules/... -v
```

## Manual Testing Scenarios

### 1. Test Full Project Generation

```bash
# Build the CLI
make build

# Create test directory
mkdir -p /tmp/test-doplan
cd /tmp/test-doplan

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "test-project"
# - IDE: Cursor
# - Wait for generation

# Verify generated project
cd test-project
ls -la
cat README.md
ls .cursor/agents/ | wc -l  # Should show 18 agents
ls .cursor/commands/ | wc -l  # Should show 19 commands
```

### 2. Test Mobile App Idea Generation

```bash
# Build the CLI
make build

# Create test directory for mobile app
mkdir -p /tmp/test-mobile-app
cd /tmp/test-mobile-app

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "mobile-fitness-tracker"
# - IDE: Cursor
# - Wait for generation

# Verify generated project
cd mobile-fitness-tracker
ls -la
cat README.md
ls .cursor/agents/ | wc -l  # Should show 18 agents
ls .cursor/commands/ | wc -l  # Should show 19 commands

# Verify mobile-specific structure
cat .plan/00_System/IDEA.md | grep -i "mobile\|app\|fitness"  # Check idea is captured
ls .plan/  # Verify plan structure exists
```

### 3. Test All IDE Configs

```bash
# Test each IDE one by one
for ide in "Cursor" "Claude Code" "Antigravity" "Windsurf" "Cline" "OpenCode"; do
  mkdir -p /tmp/test-$ide
  cd /tmp/test-$ide
  # Run wizard and select $ide
  # Verify IDE-specific configs were generated
done
```

### 4. Test Error Handling

```bash
# Test invalid project names
./doplan  # Try entering invalid names like "test project" or "test@project"

# Test in existing directory
mkdir -p /tmp/existing-project
cd /tmp/existing-project
touch some-file.txt
./doplan  # Try to generate project here (should handle gracefully)
```

### 5. Test Performance

```bash
# Time the generation
time ./doplan

# Should complete in < 5 seconds (currently ~40-50ms)
```

### 6. Test Binary Size

```bash
# Build and check size
make build
ls -lh doplan

# Should be < 15MB
```

### 7. Test Web Application Project

```bash
# Build the CLI
make build

# Create test directory for web app
mkdir -p /tmp/test-web-app
cd /tmp/test-web-app

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "ecommerce-platform"
# - IDE: Cursor
# - Wait for generation

# Verify generated project
cd ecommerce-platform
ls -la
cat README.md
ls .cursor/agents/ | wc -l  # Should show 18 agents
ls .cursor/commands/ | wc -l  # Should show 19 commands

# Verify web-specific files
ls src/  # Should have Next.js boilerplate
cat package.json 2>/dev/null || echo "Checking project structure"
ls .github/workflows/  # Should have CI/CD workflows
```

### 8. Test SaaS Project Generation

```bash
# Build the CLI
make build

# Create test directory for SaaS
mkdir -p /tmp/test-saas
cd /tmp/test-saas

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "task-management-saas"
# - IDE: Claude Code
# - Wait for generation

# Verify generated project
cd task-management-saas
ls -la

# Verify all components
test -d .cursor/agents && echo "✅ Agents directory exists"
test -d .cursor/commands && echo "✅ Commands directory exists"
test -d .cursor/rules && echo "✅ Rules directory exists"
test -d .github/workflows && echo "✅ GitHub workflows exist"
test -f README.md && echo "✅ README exists"
```

### 9. Test Paid B2C Problem-Solving App

```bash
# Build the CLI
make build

# Create test directory for paid B2C service app
mkdir -p /tmp/test-paid-app
cd /tmp/test-paid-app

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "home-repair-connect"
# - IDE: Cursor
# - Wait for generation

# This app idea: Connects homeowners with verified local repair professionals
# Problem: Homeowners struggle to find trustworthy, available repair services
# Solution: Platform with verified professionals, instant booking, payment processing
# Monetization: Subscription fees from service providers + transaction fees from bookings

# Verify generated project
cd home-repair-connect
ls -la
cat README.md

# Verify all components
ls .cursor/agents/ | wc -l  # Should show 18 agents
ls .cursor/commands/ | wc -l  # Should show 19 commands

# Verify business-focused structure
cat .plan/00_System/IDEA.md | grep -i "problem\|solution\|payment\|subscription\|monetization" || echo "Checking idea documentation"
ls .plan/  # Verify plan structure exists

# Verify payment/subscription related files would be in generated structure
echo "✅ Project generated for paid B2C service app"
echo "✅ Ready to build marketplace connecting users with service providers"
```

### 10. Test Marketplace/Platform App (Two-Sided Market)

```bash
# Build the CLI
make build

# Create test directory for marketplace platform
mkdir -p /tmp/test-marketplace
cd /tmp/test-marketplace

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "skill-exchange-platform"
# - IDE: Cursor
# - Wait for generation

# This app idea: Platform where people teach skills to each other
# Problem: People want to learn new skills but traditional classes are expensive/inflexible
# Solution: Peer-to-peer skill exchange with video sessions, ratings, and payment
# Monetization: Platform takes commission from each transaction (10-15%)

# Verify generated project
cd skill-exchange-platform
ls -la

# Verify marketplace-specific structure
test -d .cursor/agents && echo "✅ Agents directory exists"
test -d .cursor/commands && echo "✅ Commands directory exists"
test -f README.md && echo "✅ README exists"

# Verify the idea captures the two-sided market model
cat .plan/00_System/IDEA.md | head -30
echo "✅ Marketplace platform project ready for development"
```

### 11. Test API Backend Project

```bash
# Build the CLI
make build

# Create test directory for API
mkdir -p /tmp/test-api
cd /tmp/test-api

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "rest-api-service"
# - IDE: Windsurf
# - Wait for generation

# Verify generated project
cd rest-api-service
ls -la

# Verify API-specific structure
cat .plan/00_System/IDEA.md | head -20
ls .cursor/agents/ | grep -i "backend\|api\|engineer" || echo "Checking agent names"
```

### 12. Test Different Project Name Formats

```bash
# Build the CLI
make build

# Test kebab-case
mkdir -p /tmp/test-naming
cd /tmp/test-naming
../../GoPlan-CLI/doplan  # Use "my-awesome-project"

# Test camelCase (if supported)
cd /tmp/test-naming
../../GoPlan-CLI/doplan  # Use "myAwesomeProject"

# Test with numbers
cd /tmp/test-naming
../../GoPlan-CLI/doplan  # Use "project-2024"

# Verify all generate correctly
```

### 13. Test All Generated Agents

```bash
# Build the CLI
make build

# Create test directory
mkdir -p /tmp/test-agents
cd /tmp/test-agents

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "agent-test"
# - IDE: Cursor
# - Wait for generation

# Verify all 18 agents exist
cd agent-test
AGENT_COUNT=$(ls .cursor/agents/*.md 2>/dev/null | wc -l)
echo "Found $AGENT_COUNT agents (expected 18)"

# Verify key agents exist
test -f .cursor/agents/project_orchestrator.md && echo "✅ Project Orchestrator"
test -f .cursor/agents/product_manager.md && echo "✅ Product Manager"
test -f .cursor/agents/engineering_lead.md && echo "✅ Engineering Lead"
test -f .cursor/agents/frontend_engineer.md && echo "✅ Frontend Engineer"
test -f .cursor/agents/backend_engineer.md && echo "✅ Backend Engineer"
test -f .cursor/agents/qa_engineer.md && echo "✅ QA Engineer"
```

### 14. Test All Generated Commands

```bash
# Build the CLI
make build

# Create test directory
mkdir -p /tmp/test-commands
cd /tmp/test-commands

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "command-test"
# - IDE: Cursor
# - Wait for generation

# Verify all 19 commands exist
cd command-test
COMMAND_COUNT=$(ls .cursor/commands/*.md 2>/dev/null | wc -l)
echo "Found $COMMAND_COUNT commands (expected 19)"

# Verify key commands exist
test -f .cursor/commands/hey.md && echo "✅ /hey command"
test -f .cursor/commands/do.md && echo "✅ /do command"
test -f .cursor/commands/plan.md && echo "✅ /plan command"
test -f .cursor/commands/dev.md && echo "✅ /dev command"
test -f .cursor/commands/sys.md && echo "✅ /sys command"
```

### 15. Test Rules Library Extraction

```bash
# Build the CLI
make build

# Create test directory
mkdir -p /tmp/test-rules
cd /tmp/test-rules

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "rules-test"
# - IDE: Cursor
# - Wait for generation

# Verify rules library
cd rules-test
RULES_COUNT=$(find .cursor/rules/library -type f 2>/dev/null | wc -l)
echo "Found $RULES_COUNT rule files (expected 1000+)"

# Verify rules structure
test -d .cursor/rules/library && echo "✅ Rules library directory exists"
ls .cursor/rules/library/ | head -10  # Show first 10 rule categories
```

### 16. Test GitHub Workflows Generation

```bash
# Build the CLI
make build

# Create test directory
mkdir -p /tmp/test-workflows
cd /tmp/test-workflows

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "workflow-test"
# - IDE: Cursor
# - Wait for generation

# Verify GitHub workflows
cd workflow-test
WORKFLOW_COUNT=$(ls .github/workflows/*.yml 2>/dev/null | wc -l)
echo "Found $WORKFLOW_COUNT workflow files"

# Verify key workflows
test -f .github/workflows/ci.yml && echo "✅ CI workflow"
test -f .github/workflows/release.yml && echo "✅ Release workflow"
test -f .github/workflows/changelog.yml && echo "✅ Changelog workflow"
```

### 17. Test IDE-Specific Configurations

```bash
# Build the CLI
make build

# Test Cursor configuration
mkdir -p /tmp/test-cursor-config
cd /tmp/test-cursor-config
../../GoPlan-CLI/doplan  # Select Cursor
cd test-project
test -f docs/CLAUDE.md && echo "✅ IDE quick-start (docs/CLAUDE.md) exists"

# Test Claude Code configuration
mkdir -p /tmp/test-claude-config
cd /tmp/test-claude-config
../../GoPlan-CLI/doplan  # Select Claude Code
cd test-project
ls .cursor/rules  # Verify rules library is present

# Generate application code after planning (simulates the /build flow)
# Use whichever stack generator fits your project; for example:
npx create-next-app@latest .
test -f package.json && echo "✅ Boilerplate generated on demand"
```

### 18. Test Batch Project Generation

```bash
# Build the CLI
make build

# Generate multiple projects in sequence
for project in "project-1" "project-2" "project-3"; do
  mkdir -p /tmp/batch-test/$project
  cd /tmp/batch-test/$project
  ../../GoPlan-CLI/doplan  # Use $project as name, Cursor as IDE
  echo "✅ Generated $project"
done

# Verify all projects generated correctly
for project in "project-1" "project-2" "project-3"; do
  cd /tmp/batch-test/$project
  test -d .cursor && echo "✅ $project has .cursor directory"
  test -f README.md && echo "✅ $project has README"
done
```

### 19. Test Project Structure Completeness

```bash
# Build the CLI
make build

# Create test directory
mkdir -p /tmp/test-structure
cd /tmp/test-structure

# Run CLI
../../GoPlan-CLI/doplan

# Follow wizard:
# - Project name: "structure-test"
# - IDE: Cursor
# - Wait for generation

# Verify complete structure
cd structure-test

# Check all required directories
test -d .cursor && echo "✅ .cursor directory"
test -d .cursor/agents && echo "✅ .cursor/agents"
test -d .cursor/commands && echo "✅ .cursor/commands"
test -d .cursor/rules && echo "✅ .cursor/rules"
test -d .plan && echo "✅ .plan directory"
test -d .github && echo "✅ .github directory"
test -d .github/workflows && echo "✅ .github/workflows"
test -d src && echo "✅ src directory"

# Check required files
test -f README.md && echo "✅ README.md"
test -f .gitignore && echo "✅ .gitignore"
```

### 20. Test Offline Functionality

```bash
# Build the CLI
make build

# Generate a project (first run - may need network)
mkdir -p /tmp/test-offline
cd /tmp/test-offline
../../GoPlan-CLI/doplan  # Generate "offline-test"

# Disconnect network (or simulate offline)
# Then try to generate another project
mkdir -p /tmp/test-offline-2
cd /tmp/test-offline-2
../../GoPlan-CLI/doplan  # Should work offline

# Verify it generated successfully
cd offline-test-2
test -d .cursor && echo "✅ Works offline"
```

### 21. Test Project with Special Characters in Name

```bash
# Build the CLI
make build

# Test with hyphens (should work)
mkdir -p /tmp/test-special
cd /tmp/test-special
../../GoPlan-CLI/doplan  # Use "my-project-2024"

# Test with underscores (if supported)
cd /tmp/test-special
../../GoPlan-CLI/doplan  # Use "my_project_2024"

# Verify both generated correctly
```

### 22. Test Quick Verification Script

```bash
#!/bin/bash
# Quick verification script for any generated project

PROJECT_DIR=$1

if [ -z "$PROJECT_DIR" ]; then
  echo "Usage: $0 <project-directory>"
  exit 1
fi

cd "$PROJECT_DIR" || exit 1

echo "🔍 Verifying project structure..."
echo ""

# Count agents
AGENTS=$(ls .cursor/agents/*.md 2>/dev/null | wc -l | tr -d ' ')
echo "Agents: $AGENTS/18"

# Count commands
COMMANDS=$(ls .cursor/commands/*.md 2>/dev/null | wc -l | tr -d ' ')
echo "Commands: $COMMANDS/19"

# Count rules
RULES=$(find .cursor/rules/library -type f 2>/dev/null | wc -l | tr -d ' ')
echo "Rules: $RULES (expected 1000+)"

# Count workflows
WORKFLOWS=$(ls .github/workflows/*.yml 2>/dev/null | wc -l | tr -d ' ')
echo "Workflows: $WORKFLOWS"

# Check key files
echo ""
echo "Key files:"
test -f README.md && echo "  ✅ README.md" || echo "  ❌ README.md"
test -f .gitignore && echo "  ✅ .gitignore" || echo "  ❌ .gitignore"
test -d src && echo "  ✅ src/" || echo "  ❌ src/"

echo ""
echo "✅ Verification complete!"
```

## Testing with Go Run (No Build Required)

You can also test directly without building:

```bash
# Run the CLI directly
go run ./cmd/doplan/main.go

# Run with specific package
go run ./cmd/doplan
```

## Testing Scripts

### Run Progress Tool

```bash
# Test progress reporting
go run scripts/progress/main.go --root .

# Test with JSON output
go run scripts/progress/main.go --root . --json
```

### Run State History

```bash
# View state history
go run scripts/statehistory/main.go list

# View latest diff
go run scripts/statehistory/main.go diff
```

## Development Testing Workflow

1. **Make Changes**
   ```bash
   # Edit code
   vim internal/generator/plan.go
   ```

2. **Run Tests**
   ```bash
   # Quick test
   go test ./internal/generator/... -v -run TestPlan
   ```

3. **Build and Test Locally**
   ```bash
   # Build
   make build
   
   # Test manually
   ./doplan
   ```

4. **Verify Everything**
   ```bash
   # Run all tests
   make test
   
   # Check coverage
   make test-coverage
   
   # Format code
   make fmt
   
   # Lint
   make lint
   ```

## Common Test Commands

```bash
# Quick test cycle
make fmt && make vet && make test

# Full test suite
make fmt && make vet && make lint && make test && make test-coverage

# Test specific feature
go test ./internal/generator -v -run TestPlanGeneration

# Test with race detector
go test ./... -race

# Test with verbose output
go test ./... -v -args -test.v

# Benchmark tests
go test ./... -bench=. -benchmem
```

## Troubleshooting

### Tests Failing?

```bash
# Clean test cache
go clean -testcache

# Run tests again
go test ./... -v
```

### Build Errors?

```bash
# Clean build artifacts
make clean

# Update dependencies
go mod tidy

# Rebuild
make build
```

### Binary Too Large?

```bash
# Check what's included
go tool nm doplan | head -20

# Strip debug symbols (if needed)
strip doplan
ls -lh doplan
```

## CI/CD Testing

The same tests run in CI/CD:

```bash
# Run the same tests CI would run
make test
make lint
make vet
make check-docs

# Verify binary builds for all platforms
make build-all
```

