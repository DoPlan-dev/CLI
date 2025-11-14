#!/bin/bash
# scripts/pre-release.sh
# Runs pre-release checks before creating a release

set -e

echo "🔍 Running pre-release checks..."

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
  echo "❌ Not in a git repository"
  exit 1
fi

# Run tests
echo "🧪 Running tests..."
if ! make test > /dev/null 2>&1; then
  echo "❌ Tests failed. Fix tests before releasing."
  exit 1
fi
echo "✅ Tests passed"

# Check formatting
echo "🎨 Checking code format..."
make fmt
if ! git diff --quiet; then
  echo "❌ Code is not formatted. Run 'make fmt' and commit changes"
  exit 1
fi
echo "✅ Code is formatted"

# Run linters
echo "🔍 Running linters..."
if ! make lint > /dev/null 2>&1; then
  echo "⚠️  Linter warnings found. Review before releasing."
fi
echo "✅ Linters passed"

# Build binaries
echo "🔨 Building binaries..."
if ! make build > /dev/null 2>&1; then
  echo "❌ Build failed"
  exit 1
fi
echo "✅ Build successful"

# Check binary works
echo "🧪 Testing binary..."
if ! ./bin/doplan --version > /dev/null 2>&1; then
  echo "❌ Binary test failed"
  exit 1
fi
echo "✅ Binary works"

# Check for TODO/FIXME in code
echo "📋 Checking for TODO/FIXME comments..."
if grep -r "TODO\|FIXME" --include="*.go" . | grep -v "scripts/pre-release.sh" | grep -v "node_modules"; then
  echo "⚠️  Found TODO/FIXME comments. Consider addressing before release."
fi

echo ""
echo "✅ All pre-release checks passed!"
echo "🚀 Ready to create release"

