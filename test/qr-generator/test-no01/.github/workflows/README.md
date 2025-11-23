# GitHub Workflows

This directory contains GitHub Actions workflows for CI/CD and automation.

## Workflows

### `ci.yml` - Continuous Integration
- **Triggers**: Push to `main`/`develop`, Pull Requests
- **Jobs**:
  - Lint & Type Check
  - Build
  - Test (on PRs)
  - Security Audit
- **Purpose**: Ensure code quality before merging

### `deploy.yml` - Deployment
- **Triggers**: Push to `main`, Manual dispatch
- **Jobs**:
  - Deploy Preview (on PRs)
  - Deploy to Production (on `main`)
- **Purpose**: Automate deployments to Vercel

### `release.yml` - Release Management
- **Triggers**: Tag push (v*.*.*)
- **Jobs**:
  - Create GitHub Release
  - Generate Changelog
- **Purpose**: Automate release creation

### `stale.yml` - Stale Issue Management
- **Triggers**: Weekly (Monday), Manual dispatch
- **Purpose**: Mark and close stale issues/PRs

### `codeql.yml` - Security Analysis
- **Triggers**: Push, PR, Weekly schedule
- **Purpose**: CodeQL security scanning

## Setup Instructions

### 1. Enable GitHub Actions
- Go to repository Settings → Actions → General
- Enable "Allow all actions and reusable workflows"

### 2. Configure Secrets (for deployment)
Go to Settings → Secrets and variables → Actions, add:
- `VERCEL_TOKEN` - Vercel authentication token
- `VERCEL_ORG_ID` - Vercel organization ID
- `VERCEL_PROJECT_ID` - Vercel project ID

To get these:
```bash
# Install Vercel CLI
npm i -g vercel

# Login and link project
vercel login
vercel link

# Get project ID from .vercel/project.json
```

### 3. Configure Branch Protection
See `.github/BRANCH_PROTECTION.md` for detailed instructions.

### 4. Enable Dependabot
Dependabot is configured via `.github/dependabot.yml` and will automatically create PRs for dependency updates.

## Workflow Status

View workflow runs at: https://github.com/DoPlan-dev/test-no01/actions

## Troubleshooting

### Workflow fails on first run
- Ensure Node.js version matches local development
- Check that all required secrets are configured
- Verify package.json scripts are correct

### Deployment fails
- Verify Vercel secrets are set correctly
- Check Vercel project is linked
- Ensure build succeeds locally first

### Tests fail
- Run tests locally: `npm test`
- Check test configuration
- Verify test environment setup

