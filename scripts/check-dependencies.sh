#!/bin/bash
# check-dependencies.sh - Check required dependencies

set -e

echo "Checking DoPlan dependencies..."
echo ""

MISSING=0

# Check Go
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo "✓ Go installed: $GO_VERSION"
    
    # Check version (need 1.24+)
    GO_MAJOR=$(echo "$GO_VERSION" | sed 's/go//' | cut -d. -f1)
    GO_MINOR=$(echo "$GO_VERSION" | sed 's/go//' | cut -d. -f2)
    if [ "$GO_MAJOR" -gt 1 ] || ([ "$GO_MAJOR" -eq 1 ] && [ "$GO_MINOR" -ge 24 ]); then
        echo "  ✓ Go version is compatible"
    else
        echo "  ❌ Go version too old (need 1.24+)"
        MISSING=$((MISSING + 1))
    fi
else
    echo "❌ Go not installed"
    MISSING=$((MISSING + 1))
fi

# Check Git
if command -v git &> /dev/null; then
    GIT_VERSION=$(git --version | awk '{print $3}')
    echo "✓ Git installed: $GIT_VERSION"
else
    echo "❌ Git not installed"
    MISSING=$((MISSING + 1))
fi

# Check GitHub CLI (optional but recommended)
if command -v gh &> /dev/null; then
    GH_VERSION=$(gh --version | head -1 | awk '{print $3}')
    echo "✓ GitHub CLI installed: $GH_VERSION"
else
    echo "⚠️  GitHub CLI not installed (optional but recommended)"
fi

# Check jq (optional, for JSON parsing)
if command -v jq &> /dev/null; then
    echo "✓ jq installed"
else
    echo "⚠️  jq not installed (optional, for JSON parsing)"
fi

# Check yq (optional, for YAML parsing)
if command -v yq &> /dev/null; then
    echo "✓ yq installed"
else
    echo "⚠️  yq not installed (optional, for YAML parsing)"
fi

echo ""
if [ $MISSING -eq 0 ]; then
    echo "✓ All required dependencies are installed"
    exit 0
else
    echo "❌ Missing $MISSING required dependency(ies)"
    exit 1
fi

