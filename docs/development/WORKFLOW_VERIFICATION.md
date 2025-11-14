# Workflow Verification Guide

This guide helps you verify that CI/CD workflows are running correctly after pushing changes.

## Workflow Configuration

All workflows are configured to trigger on:
- **Push to `master` or `main` branches**
- **Pull requests** to `master` or `main` branches`

## Available Workflows

### 1. Test Workflow (`.github/workflows/test.yml`)
- **Triggers:** Push and PR
- **Jobs:**
  - Multi-platform testing (Linux, macOS, Windows)
  - Multiple Go versions (1.21, 1.22, 1.23)
  - Integration tests
  - Coverage reporting

### 2. Lint Workflow (`.github/workflows/lint.yml`)
- **Triggers:** Push and PR
- **Jobs:**
  - Format checking (`go fmt`)
  - Static analysis (`go vet`)
  - Go mod verification

### 3. Build Workflow (`.github/workflows/build.yml`)
- **Triggers:** Push and PR
- **Jobs:**
  - Multi-platform builds (Linux, macOS, Windows)
  - Multiple architectures (amd64, arm64)
  - Binary testing

### 4. PR Checks Workflow (`.github/workflows/pr-checks.yml`)
- **Triggers:** Pull requests only
- **Jobs:**
  - Comprehensive PR quality checks
  - Test execution
  - Format validation
  - Build verification

### 5. Release Workflow (`.github/workflows/release.yml`)
- **Triggers:** Tag push (v*)
- **Jobs:**
  - GoReleaser release
  - Homebrew formula generation
  - Release verification

## Verifying Workflows

### Method 1: GitHub Web Interface

1. **Navigate to Actions tab:**
   ```
   https://github.com/DoPlan-dev/CLI/actions
   ```

2. **Check workflow runs:**
   - Click on each workflow to see recent runs
   - Green checkmark = Success
   - Red X = Failed
   - Yellow circle = In progress

3. **View workflow logs:**
   - Click on a workflow run
   - Click on a job to see detailed logs
   - Check for any errors or warnings

### Method 2: GitHub CLI

1. **List all workflows:**
   ```bash
   gh workflow list
   ```

2. **List recent runs:**
   ```bash
   gh run list --limit 10
   ```

3. **Check specific workflow:**
   ```bash
   gh run list --workflow=test.yml --limit 5
   ```

4. **View workflow status:**
   ```bash
   gh run view <run-id>
   ```

5. **Watch workflow in real-time:**
   ```bash
   gh run watch <run-id>
   ```

### Method 3: Using Check Script

Run the check script:
```bash
./scripts/check-workflows.sh
```

This script will:
- List all available workflows
- Show recent workflow runs
- Check each workflow's status
- Provide tips and links

## Expected Behavior

### After Pushing to Master/Main

1. **Test workflow** should start automatically
   - Runs on 3 platforms (Linux, macOS, Windows)
   - Tests with Go 1.21, 1.22, 1.23
   - Duration: ~5-10 minutes

2. **Lint workflow** should start automatically
   - Checks formatting
   - Runs static analysis
   - Duration: ~2-3 minutes

3. **Build workflow** should start automatically
   - Builds for all platforms
   - Tests binaries
   - Duration: ~5-8 minutes

### After Creating a Pull Request

1. **PR Checks workflow** should start automatically
   - Runs all quality checks
   - Duration: ~3-5 minutes

2. **Test, Lint, and Build workflows** should also run
   - Same as push triggers
   - Results shown in PR status checks

### After Pushing a Tag

1. **Release workflow** should start automatically
   - Builds release binaries
   - Creates GitHub release
   - Generates Homebrew formula
   - Duration: ~5-10 minutes

## Troubleshooting

### Workflows Not Triggering

1. **Check branch name:**
   - Workflows trigger on `master` or `main`
   - Ensure you're pushing to the correct branch

2. **Check workflow files:**
   - Verify `.github/workflows/*.yml` files exist
   - Check YAML syntax is correct
   - Ensure trigger conditions are correct

3. **Check GitHub Actions:**
   - Verify Actions are enabled for the repository
   - Go to: Settings > Actions > General
   - Ensure "Allow all actions and reusable workflows" is enabled

4. **Check permissions:**
   - Verify workflow permissions in `.github/workflows/*.yml`
   - Ensure required permissions are granted

### Workflows Failing

1. **Check workflow logs:**
   - Click on failed workflow run
   - Review error messages
   - Check for missing dependencies or configuration

2. **Common issues:**
   - **Go version mismatch:** Check `go.mod` specifies correct version
   - **Missing dependencies:** Run `go mod download` locally
   - **Test failures:** Run tests locally to reproduce
   - **Format issues:** Run `make fmt` locally

3. **Fix and retry:**
   - Fix the issue locally
   - Commit and push changes
   - Workflows will automatically re-run

### Workflows Running Slowly

1. **Check queue:**
   - GitHub Actions has a queue for free accounts
   - Wait times can be 5-10 minutes during peak hours

2. **Optimize workflows:**
   - Reduce matrix builds if needed
   - Cache dependencies
   - Run fewer Go versions on non-Linux platforms

## Testing Workflows Locally

### Using act (Optional)

Install `act` to run workflows locally:
```bash
# Install act
brew install act  # macOS
# or
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Run a workflow
act -l
act push
```

### Manual Testing

Test workflow steps manually:
```bash
# Test workflow
make test
make lint
make build

# Check formatting
make fmt
git diff --exit-code

# Check vet
make vet
```

## Workflow Status Badges

Add badges to README to show workflow status:
```markdown
[![Test](https://github.com/DoPlan-dev/CLI/actions/workflows/test.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/test.yml)
[![Lint](https://github.com/DoPlan-dev/CLI/actions/workflows/lint.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/lint.yml)
[![Build](https://github.com/DoPlan-dev/CLI/actions/workflows/build.yml/badge.svg)](https://github.com/DoPlan-dev/CLI/actions/workflows/build.yml)
```

## Next Steps

1. **Monitor workflows:**
   - Check GitHub Actions tab regularly
   - Review workflow runs
   - Address any failures

2. **Optimize workflows:**
   - Add caching for dependencies
   - Reduce build matrix if needed
   - Add more test coverage

3. **Set up notifications:**
   - Configure email notifications for workflow failures
   - Set up Slack/Discord notifications (optional)

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Workflow Syntax](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions)
- [GitHub CLI Documentation](https://cli.github.com/manual/)
- [Actions Tab](https://github.com/DoPlan-dev/CLI/actions)

