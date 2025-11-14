#!/bin/bash
# scripts/check-workflows.sh
# Check GitHub Actions workflow status

set -e

echo "🔍 Checking GitHub Actions Workflows"
echo ""

# Check if gh is installed
if ! command -v gh &> /dev/null; then
    echo "❌ GitHub CLI (gh) is not installed."
    echo "   Install it from: https://cli.github.com/"
    exit 1
fi

# Check if authenticated
if ! gh auth status &> /dev/null; then
    echo "❌ GitHub CLI is not authenticated."
    echo "   Run: gh auth login"
    exit 1
fi

echo "📋 Available Workflows:"
echo ""
gh workflow list

echo ""
echo "📊 Recent Workflow Runs:"
echo ""
gh run list --limit 10

echo ""
echo "🌐 View workflows in browser:"
echo "   https://github.com/DoPlan-dev/CLI/actions"
echo ""

# Check specific workflows
echo "🔍 Checking specific workflows:"
echo ""

WORKFLOWS=("test.yml" "lint.yml" "build.yml" "pr-checks.yml" "release.yml")

for workflow in "${WORKFLOWS[@]}"; do
    echo "Checking $workflow..."
    if gh run list --workflow="$workflow" --limit 1 &> /dev/null; then
        echo "  ✅ Workflow exists and has runs"
        gh run list --workflow="$workflow" --limit 1
    else
        echo "  ⚠️  Workflow exists but no runs yet (may trigger on next push)"
    fi
    echo ""
done

echo "✅ Workflow check complete!"
echo ""
echo "💡 Tips:"
echo "   - Workflows trigger automatically on push to master/main"
echo "   - Workflows also trigger on pull requests"
echo "   - Check the Actions tab for real-time status"
echo "   - Wait a few minutes after pushing for workflows to start"

