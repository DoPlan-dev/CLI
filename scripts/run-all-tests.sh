#!/bin/bash
# run-all-tests.sh - Run all test suites

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "Running DoPlan v0.0.18-beta Test Suite"
echo "======================================"
echo ""

cd "$PROJECT_ROOT"

# Setup test environment
echo "1. Setting up test environment..."
./scripts/setup-test-env.sh

# Run unit tests
echo ""
echo "2. Running unit tests..."
go test ./internal/migration/... -v || echo "⚠️  Migration tests failed"
go test ./internal/context/... -v || echo "⚠️  Context tests failed"
go test ./internal/wizard/... -v || echo "⚠️  Wizard tests failed"
go test ./internal/integration/... -v || echo "⚠️  Integration tests failed"
go test ./internal/tui/... -v || echo "⚠️  TUI tests failed"

# Run integration tests
echo ""
echo "3. Running integration tests..."
./scripts/test-migration.sh || echo "⚠️  Migration integration test failed"
./scripts/test-ide-integration.sh || echo "⚠️  IDE integration test failed"

echo ""
echo "Test suite complete!"

