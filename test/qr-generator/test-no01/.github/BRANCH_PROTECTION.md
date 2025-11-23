# Branch Protection Guidelines

This document outlines the recommended branch protection rules for this repository.

## Recommended Settings for `main` Branch

### Required Settings

1. **Require a pull request before merging**
   - ✅ Required
   - Minimum number of reviewers: 1
   - Dismiss stale pull request approvals when new commits are pushed: ✅

2. **Require status checks to pass before merging**
   - ✅ Required
   - Required status checks:
     - `lint-and-typecheck`
     - `build`
   - Require branches to be up to date before merging: ✅

3. **Require conversation resolution before merging**
   - ✅ Required

4. **Require signed commits**
   - Optional (recommended for production)

5. **Require linear history**
   - Optional (prevents merge commits)

6. **Include administrators**
   - ✅ Recommended (applies rules to admins too)

7. **Restrict who can push to matching branches**
   - Optional (for team repositories)

8. **Allow force pushes**
   - ❌ Disabled

9. **Allow deletions**
   - ❌ Disabled

## How to Configure

1. Go to repository Settings
2. Navigate to "Branches"
3. Click "Add rule" or edit existing rule for `main`
4. Configure the settings above
5. Save changes

## Alternative: Using GitHub API

You can also configure branch protection using the GitHub API or Terraform:

```bash
# Example using GitHub CLI
gh api repos/DoPlan-dev/test-no01/branches/main/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"contexts":["lint-and-typecheck","build"]}' \
  --field enforce_admins=true \
  --field required_pull_request_reviews='{"required_approving_review_count":1}' \
  --field restrictions=null
```

## Notes

- These settings help maintain code quality
- Adjust based on team size and workflow
- Consider using CODEOWNERS file for automatic reviewer assignment

