# npm Publishing Guide

This document describes how to publish the DoPlan CLI to npm.

## Overview

The DoPlan CLI is published to npm as `doplan-cli`. The npm package includes a wrapper script that automatically downloads the platform-specific binary from GitHub releases.

## Package Structure

- `package.json` - npm package configuration
- `bin/doplan.js` - Wrapper script that downloads and executes the binary
- `scripts/postinstall.js` - Post-install script that downloads the binary
- `scripts/prepublish.js` - Pre-publish validation script
- `.npmrc` - npm configuration

## Publishing Process

### Automated Publishing (Recommended)

Publishing to npm is automated via GitHub Actions. When you push a tag matching `v*`:

1. The release workflow builds binaries for all platforms
2. Creates a GitHub release
3. Automatically publishes to npm with the same version

**Requirements:**
- `NPM_TOKEN` secret must be configured in GitHub repository settings
- Token must have publish permissions for the `doplan-cli` package

### Manual Publishing

If you need to publish manually:

```bash
# 1. Ensure you're logged in to npm
npm login

# 2. Update version in package.json (if needed)
npm version patch|minor|major

# 3. Run pre-publish validation
npm run prepublishOnly

# 4. Publish to npm
npm publish

# 5. Verify publication
npm view doplan-cli
```

## Setting up NPM_TOKEN

1. **Create an npm access token:**
   - Go to https://www.npmjs.com/settings/YOUR_USERNAME/tokens
   - Click "Generate New Token"
   - Select "Automation" token type
   - Copy the token

2. **Add to GitHub Secrets:**
   - Go to your repository settings
   - Navigate to Secrets and variables > Actions
   - Click "New repository secret"
   - Name: `NPM_TOKEN`
   - Value: Your npm token
   - Click "Add secret"

## Package Details

- **Package Name:** `doplan-cli`
- **Registry:** https://registry.npmjs.org/
- **Access:** Public
- **Scoped:** No (public package)

## How It Works

1. **Installation:** When users run `npm install -g doplan-cli`:
   - npm installs the package files
   - `postinstall.js` runs automatically
   - Downloads the platform-specific binary from GitHub releases
   - Stores binary in `bin/<platform>-<arch>/doplan`

2. **Execution:** When users run `doplan`:
   - `bin/doplan.js` wrapper script is executed
   - Script detects platform and architecture
   - If binary doesn't exist, downloads it from GitHub releases
   - Executes the binary with all arguments

## Version Management

- Version in `package.json` should match the Go release version
- GitHub Actions automatically syncs version from git tag
- Use semantic versioning (e.g., `1.0.0`, `1.0.1`, `1.1.0`)

## Troubleshooting

### Binary Download Fails

If the binary download fails during installation:
- Check GitHub releases exist for the version
- Verify network connectivity
- Users can manually download from GitHub releases

### Publishing Fails

If automated publishing fails:
- Check `NPM_TOKEN` secret is set correctly
- Verify token has publish permissions
- Check package name isn't already taken
- Ensure version doesn't already exist on npm

### Version Mismatch

If package.json version doesn't match release:
- GitHub Actions automatically updates it from git tag
- For manual publishing, ensure versions match

## Testing Before Publishing

```bash
# Test package locally
npm pack
tar -xzf doplan-cli-*.tgz
cd package
npm install

# Test the wrapper script
./bin/doplan.js --version

# Test installation simulation
npm run prepublishOnly
```

## Updating Package

To update the npm package:

1. Make changes to package files
2. Update version (automated in CI/CD)
3. Push git tag (triggers automated publish)
4. Or publish manually with `npm publish`

## Package Files Included

The following files are included in the npm package (defined in `package.json` `files` field):

- `bin/` - Wrapper scripts
- `scripts/` - Install and publish scripts
- `README.md` - Package documentation
- `LICENSE` - License file

Excluded files:
- Source code (Go files)
- Build artifacts
- Test files
- Development documentation

## Support

For issues with npm publishing:
- Check GitHub Actions workflow logs
- Verify npm token permissions
- Review npm package page: https://www.npmjs.com/package/doplan-cli

