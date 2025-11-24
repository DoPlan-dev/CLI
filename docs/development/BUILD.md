# Build & Distribution Guide

## Building from Source

### Prerequisites
- Go 1.21 or later
- Git (for version detection)

### Quick Build

```bash
# Build for current platform
make build

# Or use Go directly
go build -o doplan ./cmd/doplan
```

### Build for All Platforms

```bash
# Using Make
make build-all

# Or using build script
bash scripts/build.sh

# On Windows
scripts\build.bat
```

### Build with Version

```bash
# Set version explicitly
VERSION=1.0.0 make build

# Or with Go
go build -ldflags "-X github.com/DoPlan-dev/CLI/internal/version.Version=1.0.0" -o doplan ./cmd/doplan
```

## Distribution

### GitHub Releases

The release workflow automatically:
1. Builds binaries for all platforms when a tag is pushed
2. Creates archives (tar.gz for Unix, zip for Windows)
3. Generates checksums
4. Uploads to GitHub Releases

**To create a release:**
```bash
# Tag the release
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

The GitHub Actions workflow will automatically:
- Build binaries for all platforms
- Create GitHub Release
- Upload binaries and checksums

### Supported Platforms

- **macOS**: darwin/amd64, darwin/arm64
- **Linux**: linux/amd64, linux/arm64
- **Windows**: windows/amd64

### Binary Downloads

Binaries are available from GitHub Releases:
```
https://github.com/DoPlan-dev/CLI/releases/download/v1.0.0/doplan-1.0.0-<platform>-<arch>.<ext>
```

### NPX Installation

Install via npm/npx:
```bash
# Install globally
npm install -g @doplan-dev/cli

# Or use npx (no installation needed)
npx @doplan-dev/cli
```

The npm package automatically downloads the correct binary for your platform.

## Versioning

Version is set at build time using Go's `-ldflags`:

```bash
go build -ldflags "-X github.com/DoPlan-dev/CLI/internal/version.Version=1.0.0" ./cmd/doplan
```

The version can be checked with:
```bash
doplan --version
```

## Checksums

All releases include SHA256 checksums:
- `doplan-<version>-<platform>-<arch>.tar.gz.sha256`
- `doplan-<version>-<platform>-<arch>.zip.sha256`

Verify downloads:
```bash
# Unix
shasum -a 256 -c doplan-1.0.0-linux-amd64.tar.gz.sha256

# Windows
certutil -hashfile doplan-1.0.0-windows-amd64.zip SHA256
```

## Makefile Targets

- `make build` - Build for current platform
- `make build-all` - Build for all platforms
- `make clean` - Clean build artifacts
- `make test` - Run tests
- `make test-coverage` - Run tests with coverage
- `make install` - Install to GOPATH/bin
- `make lint` - Run linter
- `make fmt` - Format code
- `make vet` - Run go vet
- `make version` - Show version

## Development

### Local Development Build
```bash
make build
./doplan
```

### Testing Build Scripts
```bash
# Test build script
bash scripts/build.sh

# Verify binaries
ls -lh dist/
```

## CI/CD

The project uses GitHub Actions for:
- **CI**: Tests and linting on every push
- **Release**: Automated builds and releases on tag push
- **Changelog**: Auto-commit CHANGELOG.md updates

See `.github/workflows/` for workflow definitions.

