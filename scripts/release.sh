#!/bin/bash
# scripts/release.sh
# Creates a new release with GoReleaser

set -e

VERSION=$1
if [ -z "$VERSION" ]; then
  echo "Usage: ./scripts/release.sh <version>"
  echo "Example: ./scripts/release.sh v1.0.0"
  exit 1
fi

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "❌ Invalid version format. Use semantic versioning: v1.0.0"
  exit 1
fi

echo "🚀 Creating release $VERSION"

# Check if tag exists
if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "❌ Tag $VERSION already exists"
  exit 1
fi

# Check if working directory is clean
if ! git diff-index --quiet HEAD --; then
  echo "❌ Working directory is not clean. Commit or stash changes first."
  exit 1
fi

# Run pre-release checks
echo "🔍 Running pre-release checks..."
./scripts/pre-release.sh

# Create and push tag
echo "📝 Creating tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION"

echo "📤 Pushing tag to remote..."
git push origin "$VERSION"

# Create release with goreleaser
echo "📦 Creating release with GoReleaser..."
if command -v goreleaser &> /dev/null; then
  goreleaser release --clean
else
  echo "⚠️  GoReleaser not found. Install with: brew install goreleaser"
  echo "   Or download from: https://goreleaser.com/install"
  exit 1
fi

echo ""
echo "✅ Release $VERSION created successfully!"
echo "📦 Binaries uploaded to GitHub Releases"
echo "🍺 Homebrew formula will be updated via PR (if configured)"
echo ""
echo "Next steps:"
echo "  1. Review the GitHub release"
echo "  2. Update Homebrew formula (if needed)"
echo "  3. Announce the release"

