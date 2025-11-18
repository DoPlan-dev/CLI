#!/bin/bash
# setup-test-env.sh - Set up test environment with various project states

set -e

TEST_DIR="/tmp/doplan-test"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "Setting up DoPlan test environment..."

# Create test directory
mkdir -p "$TEST_DIR"

# 1. Empty folder
echo "Creating empty folder test project..."
EMPTY_DIR="$TEST_DIR/empty"
rm -rf "$EMPTY_DIR"
mkdir -p "$EMPTY_DIR"
echo "✓ Empty folder created: $EMPTY_DIR"

# 2. Existing code without DoPlan
echo "Creating existing code test project..."
EXISTING_DIR="$TEST_DIR/existing"
rm -rf "$EXISTING_DIR"
mkdir -p "$EXISTING_DIR"
cd "$EXISTING_DIR"

# Create a simple Next.js project structure
cat > package.json <<EOF
{
  "name": "test-project",
  "version": "1.0.0",
  "scripts": {
    "dev": "next dev"
  }
}
EOF

mkdir -p src pages components
echo "✓ Existing code project created: $EXISTING_DIR"

# 3. Old DoPlan structure (v0.0.17)
echo "Creating old DoPlan structure test project..."
OLD_DIR="$TEST_DIR/old"
rm -rf "$OLD_DIR"
mkdir -p "$OLD_DIR"
cd "$OLD_DIR"

# Create old config structure
mkdir -p .cursor/config
cat > .cursor/config/doplan-config.json <<EOF
{
  "ide": "cursor",
  "version": "0.0.17",
  "github": {
    "enabled": true,
    "autoBranch": true,
    "autoPR": true
  }
}
EOF

# Create old folder structure
mkdir -p doplan/01-phase doplan/02-phase
mkdir -p doplan/01-phase/01-Feature doplan/01-phase/02-Feature

# Create phase plan
cat > doplan/01-phase/phase-plan.md <<EOF
# Phase 1: User Authentication

This phase handles user authentication.
EOF

# Create feature plan
cat > doplan/01-phase/01-Feature/plan.md <<EOF
# Feature 1: Login with Email

This feature implements email-based login.
EOF

# Create progress files
cat > doplan/01-phase/phase-progress.json <<EOF
{
  "status": "in-progress",
  "progress": 50
}
EOF

cat > doplan/01-phase/01-Feature/progress.json <<EOF
{
  "status": "in-progress",
  "progress": 30
}
EOF

echo "✓ Old DoPlan structure created: $OLD_DIR"

# 4. New DoPlan structure (v0.0.18)
echo "Creating new DoPlan structure test project..."
NEW_DIR="$TEST_DIR/new"
rm -rf "$NEW_DIR"
mkdir -p "$NEW_DIR"
cd "$NEW_DIR"

# Create new config structure
mkdir -p .doplan/ai/{agents,rules,commands}
cat > .doplan/config.yaml <<EOF
project:
  name: "test-project"
  type: "web"
  version: "0.0.18"
  ide: "cursor"

github:
  repository: "user/test-project"
  enabled: true
  autoBranch: true
  autoPR: true

design:
  hasPreferences: false
  tokensPath: "doplan/design/design-tokens.json"

security:
  lastScan: null
  autoFix: false

apis:
  configured: []
  required: []

tui:
  theme: "default"
  animations: true
EOF

# Create new folder structure
mkdir -p doplan/01-user-authentication doplan/02-user-profile
mkdir -p doplan/01-user-authentication/01-login-with-email

# Create phase plan
cat > doplan/01-user-authentication/phase-plan.md <<EOF
# User Authentication

This phase handles user authentication.
EOF

# Create feature plan
cat > doplan/01-user-authentication/01-login-with-email/plan.md <<EOF
# Login with Email

This feature implements email-based login.
EOF

# Create progress files
cat > doplan/01-user-authentication/phase-progress.json <<EOF
{
  "status": "in-progress",
  "progress": 50
}
EOF

cat > doplan/01-user-authentication/01-login-with-email/progress.json <<EOF
{
  "status": "in-progress",
  "progress": 30
}
EOF

echo "✓ New DoPlan structure created: $NEW_DIR"

echo ""
echo "Test environment setup complete!"
echo ""
echo "Test projects:"
echo "  - Empty: $EMPTY_DIR"
echo "  - Existing code: $EXISTING_DIR"
echo "  - Old DoPlan: $OLD_DIR"
echo "  - New DoPlan: $NEW_DIR"
echo ""
echo "You can now run tests with:"
echo "  cd $EMPTY_DIR && doplan"
echo "  cd $EXISTING_DIR && doplan"
echo "  cd $OLD_DIR && doplan"
echo "  cd $NEW_DIR && doplan"

