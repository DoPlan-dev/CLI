# Release Process

This document describes the release process for DoPlan CLI.

## Overview

Releases are automated using GitHub Actions and GoReleaser. When a tag matching `v*` is pushed, the release workflow automatically:

1. Builds binaries for all platforms
2. Creates a GitHub release
3. Generates changelog
4. Creates Homebrew PR (if tap repository exists)
5. Uploads artifacts and checksums

## Versioning

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR** (v2.0.0): Incompatible API changes
- **MINOR** (v1.1.0): New functionality (backwards compatible)
- **PATCH** (v1.0.1): Bug fixes (backwards compatible)

## Release Steps

### 1. Pre-Release Checklist

- [ ] All tests passing
- [ ] Code formatted (`make fmt`)
- [ ] Linting passes (`make lint`)
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Version number determined

### 2. Update CHANGELOG.md

Add a new section for the release:

```markdown
## [1.0.1] - 2025-01-XX

### Fixed
- Bug fix description

### Changed
- Change description
```

### 3. Create Release Tag

```bash
# Create annotated tag
git tag -a v1.0.1 -m "Release v1.0.1"

# Push tag
git push origin v1.0.1
```

### 4. Automated Release

GitHub Actions will automatically:

1. **Build binaries** for:
   - Linux (amd64, arm64)
   - macOS (amd64, arm64)
   - Windows (amd64, arm64)

2. **Create GitHub release** with:
   - Release notes (auto-generated from changelog)
   - Binaries for all platforms
   - Checksums file
   - Debian and RPM packages

3. **Generate Homebrew formula** (if tap repository exists)
   - Creates PR to homebrew-doplan repository
   - Updates formula with new version and SHA256

### 5. Verify Release

After the workflow completes:

1. Check GitHub release: https://github.com/DoPlan-dev/CLI/releases
2. Verify all binaries are uploaded
3. Test downloading and installing a binary
4. Verify Homebrew PR (if applicable)

## Release Workflow

The release workflow (`.github/workflows/release.yml`) triggers on tag push:

```yaml
on:
  push:
    tags:
      - 'v*'
```

## GoReleaser Configuration

GoReleaser is configured in `.goreleaser.yml`:

- **Builds**: Multi-platform builds (Linux, macOS, Windows)
- **Archives**: Tarballs with README and LICENSE
- **Checksums**: SHA256 checksums file
- **Homebrew**: Automatic formula generation and PR creation
- **NFPMs**: Debian and RPM packages
- **Changelog**: Auto-generated from git commits

## Homebrew Release

### Setup (One-time)

1. **Create Homebrew tap repository**
   ```bash
   # Create repository: DoPlan-dev/homebrew-doplan
   # Initialize with Formula directory
   ```

2. **Configure GitHub token**
   - Create Personal Access Token with `repo` scope
   - Add as secret: `HOMEBREW_TAP_TOKEN`
   - Token needs write access to homebrew-doplan repository

### Automated Process

When a release is created:

1. GoReleaser generates formula
2. Creates PR to homebrew-doplan repository
3. PR includes updated formula with:
   - New version
   - SHA256 checksums
   - Download URLs

### Manual Process (if needed)

If automated PR creation fails:

1. Get formula from release artifacts
2. Create PR manually to homebrew-doplan
3. Update formula in `Formula/doplan.rb`

## Release Notes

Release notes are auto-generated from:

- Git commits since last tag
- CHANGELOG.md entries
- Conventional commits format

### Conventional Commits

Use conventional commits for better release notes:

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `chore:` - Maintenance tasks
- `refactor:` - Code refactoring
- `test:` - Test changes

## Troubleshooting

### Release Workflow Fails

1. Check workflow logs in GitHub Actions
2. Verify GoReleaser configuration
3. Check for missing secrets (GITHUB_TOKEN, HOMEBREW_TAP_TOKEN)
4. Verify tag format (must start with `v`)

### Homebrew PR Not Created

1. Verify `HOMEBREW_TAP_TOKEN` secret exists
2. Check token has write access to homebrew-doplan
3. Verify homebrew-doplan repository exists
4. Check GoReleaser logs for errors

### Binary Build Fails

1. Check Go version compatibility
2. Verify all dependencies are available
3. Check for platform-specific issues
4. Review build logs for errors

## Post-Release

### Announcement

- Update website (if applicable)
- Post on social media
- Notify users in issue tracker
- Update documentation

### Monitoring

- Monitor GitHub releases for issues
- Check Homebrew installs
- Review user feedback
- Address any critical issues

## Emergency Releases

For critical bug fixes:

1. Create hotfix branch from main
2. Apply fix
3. Update CHANGELOG.md
4. Create patch release tag
5. Follow normal release process

## Release Schedule

- **Major releases**: As needed (breaking changes)
- **Minor releases**: Monthly or as features are ready
- **Patch releases**: As needed (bug fixes)

## Resources

- [GoReleaser Documentation](https://goreleaser.com/)
- [Semantic Versioning](https://semver.org/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Actions](https://docs.github.com/en/actions)

