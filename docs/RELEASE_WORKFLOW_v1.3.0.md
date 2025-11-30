# Release Workflow for v1.3.0

## Current Status

- **Branch**: `v1.3.0`
- **Commits ahead of origin**: 4 commits
- **Tag**: `v1.3.0` (already pushed ✅)
- **Status**: Ready to push commits and merge to main

## Release Workflow Steps

### Step 1: Push Commits to v1.3.0 Branch

```bash
# Push all commits to the v1.3.0 branch
git push origin v1.3.0
```

This will push:
- `cb75b7e` - chore: Prepare v1.3.0 release
- `e494d5d` - fix: Add missing files for v1.3.0 release
- `1df4d99` - fix: Add missing internal/time package for v1.3.0
- `5aeb703` - fix: Add SyncPlanDocumentation function to plan.go for v1.3.0

### Step 2: Merge v1.3.0 to main

```bash
# Switch to main branch
git checkout main

# Pull latest changes
git pull origin main

# Merge v1.3.0 into main
git merge v1.3.0

# Push main branch
git push origin main
```

### Step 3: Verify Release

After merging, verify:
- [ ] GitHub Actions release workflow completed
- [ ] GitHub Release created at https://github.com/DoPlan-dev/CLI/releases/tag/v1.3.0
- [ ] Binaries uploaded for all platforms
- [ ] Release notes displayed correctly

## Alternative: Use Pull Request (Recommended)

If you prefer a safer approach with code review:

```bash
# Push v1.3.0 branch
git push origin v1.3.0

# Create a Pull Request from v1.3.0 to main on GitHub
# Then merge via GitHub UI after review
```

## Important Notes

1. **Tag is already pushed**: The v1.3.0 tag has been pushed, so the release workflow may have already started
2. **Modified files**: There are some modified files (Makefile, test files, wiki files) that should be committed if they're part of the release
3. **Deleted files**: Many old `.cursor/rules` and `.plan` files are marked as deleted - these are likely cleanup and should be committed

## Quick Commands

```bash
# Option 1: Direct push and merge (if you're confident)
git push origin v1.3.0
git checkout main
git pull origin main
git merge v1.3.0
git push origin main

# Option 2: Push and create PR (safer)
git push origin v1.3.0
# Then create PR on GitHub: v1.3.0 -> main
```

---

**Last Updated**: 2025-01-15

