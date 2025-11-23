# GitHub Workflow Setup Guide

Complete guide to setting up GitHub workflows and repository configuration for `test-no01`.

## Quick Start Checklist

- [ ] Push all files to GitHub
- [ ] Enable GitHub Actions
- [ ] Configure repository secrets (for deployment)
- [ ] Set up branch protection rules
- [ ] Verify workflows run successfully

## Step-by-Step Setup

### 1. Push Files to GitHub

```bash
# Initialize git if not already done
git init

# Add remote (if not already added)
git remote add origin https://github.com/DoPlan-dev/test-no01.git

# Add all files
git add .

# Commit
git commit -m "chore: set up GitHub workflows and configuration"

# Push to main branch
git branch -M main
git push -u origin main
```

### 2. Enable GitHub Actions

1. Go to your repository: https://github.com/DoPlan-dev/test-no01
2. Navigate to **Settings** → **Actions** → **General**
3. Under "Workflow permissions", select:
   - ✅ **Read and write permissions**
   - ✅ **Allow GitHub Actions to create and approve pull requests**
4. Click **Save**

### 3. Configure Repository Secrets (for Vercel Deployment)

**Note**: Skip this step if you're not deploying to Vercel yet. The workflow will skip deployment but continue to work.

1. Go to **Settings** → **Secrets and variables** → **Actions**
2. Click **New repository secret**
3. Add the following secrets:

#### VERCEL_TOKEN
```bash
# Install Vercel CLI
npm i -g vercel

# Login
vercel login

# Get token from: https://vercel.com/account/tokens
# Or use: vercel whoami
```

#### VERCEL_ORG_ID and VERCEL_PROJECT_ID
```bash
# Link your project
vercel link

# Check .vercel/project.json
cat .vercel/project.json
# Copy orgId and projectId
```

### 4. Set Up Branch Protection

1. Go to **Settings** → **Branches**
2. Click **Add rule** or **Edit** next to `main` branch
3. Configure:
   - ✅ **Require a pull request before merging**
     - Required approving reviews: 1
     - ✅ Dismiss stale pull request approvals
   - ✅ **Require status checks to pass before merging**
     - Required checks: `lint-and-typecheck`, `build`
     - ✅ Require branches to be up to date
   - ✅ **Require conversation resolution before merging**
   - ✅ **Include administrators**
   - ❌ **Allow force pushes** (disabled)
   - ❌ **Allow deletions** (disabled)
4. Click **Save**

### 5. Verify Workflows

1. Go to **Actions** tab in your repository
2. You should see workflows running
3. Check that all jobs pass:
   - ✅ Lint & Type Check
   - ✅ Build
   - ⚠️ Deploy (will skip if secrets not configured - this is OK)

### 6. Test the Workflow

Create a test PR to verify everything works:

```bash
# Create a new branch
git checkout -b test/workflow-verification

# Make a small change
echo "# Test" >> TEST.md

# Commit and push
git add TEST.md
git commit -m "test: verify GitHub workflows"
git push origin test/workflow-verification
```

Then:
1. Go to GitHub and create a Pull Request
2. Check that CI workflow runs
3. Verify all checks pass
4. Merge the PR (or close it)

## Workflow Files Overview

### `.github/workflows/ci.yml`
- Runs on every push and PR
- Lints code, type checks, builds
- Runs security audit

### `.github/workflows/deploy.yml`
- Deploys preview builds on PRs
- Deploys to production on `main` branch
- Requires Vercel secrets

### `.github/workflows/release.yml`
- Creates GitHub releases when tags are pushed
- Generates changelog automatically

### `.github/workflows/stale.yml`
- Marks stale issues/PRs after 60 days
- Closes them after 7 more days of inactivity

### `.github/workflows/codeql.yml`
- Security scanning with CodeQL
- Runs on push, PR, and weekly schedule

## Issue Templates

When creating issues, you'll see templates for:
- 🐛 **Bug Report** - Report bugs
- ✨ **Feature Request** - Suggest features
- ❓ **Question** - Ask questions

## Pull Request Template

All PRs will automatically include a template with:
- Description
- Type of change
- Related issues
- Testing checklist
- Additional notes

## Dependabot

Dependabot is configured to:
- Check for updates weekly (Mondays)
- Create PRs for dependency updates
- Group minor/patch updates
- Auto-label PRs

## Troubleshooting

### Workflows not running
- Check that GitHub Actions is enabled in repository settings
- Verify `.github/workflows/` files are committed
- Check Actions tab for error messages

### CI fails
- Run `npm run lint` locally
- Run `npm run type-check` locally
- Run `npm run build` locally
- Fix any errors before pushing

### Deployment fails
- Verify Vercel secrets are set correctly
- Check Vercel project is linked
- Ensure build works locally first
- Note: Deployment will skip if secrets aren't configured (this is OK for now)

### Branch protection blocking
- Ensure required status checks pass
- Get required number of approvals
- Resolve all conversations
- Ensure branch is up to date

## Next Steps

1. ✅ Set up workflows (you're here!)
2. Start development following the tasks in `.plan/TASKS.md`
3. Create feature branches for new work
4. Use PRs for code review
5. Merge to `main` when ready

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Vercel Deployment Guide](https://vercel.com/docs)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Branch Protection Rules](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches)

---

**Need Help?** Open an issue with the `question` template!

