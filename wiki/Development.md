# Development Documentation

Complete guide for contributors to DoPlan CLI development.

---

## Building from Source

### Prerequisites

- **Go** >= 1.23.0
- **Git** (for version detection)
- **Make** (optional, for build commands)

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

# Or with Go (resolve module path dynamically to keep casing accurate)
MODULE_PATH=$(go list -m)
go build -ldflags "-X ${MODULE_PATH}/internal/version.Version=1.0.0" -o doplan ./cmd/doplan
```

---

## Development Environment Setup

### 1. Clone Repository

```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build Project

```bash
make build
```

### 4. Verify Installation

```bash
./doplan --version
```

---

## Project Structure for Contributors

### Directory Structure

```
CLI/
├── cmd/doplan/          # Main application entry point
├── internal/
│   ├── cli/             # CLI command definitions
│   ├── tui/             # Terminal UI (Bubbletea)
│   ├── generator/       # Project generation logic
│   └── rules/           # Rules library management
├── pkg/models/          # Data models
├── scripts/             # Build and utility scripts
├── docs/                # Documentation
└── test/                # Test files
```

### Key Files

- `cmd/doplan/main.go` - Application entry point
- `internal/cli/root.go` - CLI root command
- `internal/tui/wizard.go` - TUI wizard
- `internal/generator/generator.go` - Main generator
- `internal/rules/rules.go` - Rules extraction

---

## Running Tests

### Run All Tests

```bash
make test

# Or
go test ./...
```

### Run Tests with Coverage

```bash
make test-coverage

# Or
go test -cover ./...
```

### Run Specific Tests

```bash
go test -run TestName ./internal/generator
```

### Run Integration Tests

```bash
go test -tags=integration ./...
```

---

## Debugging

### Debug Mode

Enable debug logging:

```bash
export DOPLAN_DEBUG=1
./doplan
```

### Using Delve Debugger

```bash
# Install Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug
dlv debug ./cmd/doplan
```

### Print Debugging

Add debug prints:

```go
if os.Getenv("DOPLAN_DEBUG") == "1" {
    fmt.Printf("Debug: %v\n", value)
}
```

---

## Release Process

### 1. Update Version

Update version in:
- `internal/version/version.go`
- `package.json`

### 2. Update CHANGELOG

Add release notes to `CHANGELOG.md`

### 3. Create Release Tag

```bash
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### 4. GitHub Actions

The release workflow automatically:
- Builds binaries for all platforms
- Creates GitHub Release
- Uploads binaries and checksums

---

## Development Workflow

### 1. Create Feature Branch

```bash
git checkout -b feature/your-feature
```

### 2. Make Changes

- Write code
- Add tests
- Update documentation

### 3. Test Changes

```bash
make test
make build
```

### 4. Commit Changes

Use [Conventional Commits](https://www.conventionalcommits.org/):

```bash
git commit -m "feat: add new feature"
```

### 5. Push and Create PR

```bash
git push origin feature/your-feature
```

Create pull request on GitHub.

---

## Development Examples

### Example 1: Adding a New Agent

1. **Add Agent Definition**:
   ```go
   // In internal/generator/agents.go
   {
       Name: "New Agent",
       Role: "Agent role",
       // ...
   }
   ```

2. **Create Agent Template**:
   ```go
   // In internal/generator/agents.go
   func generateNewAgent() string {
       // Agent template
   }
   ```

3. **Test**:
   ```bash
   go test ./internal/generator
   ```

---

### Example 2: Adding a New Command

1. **Add Command Definition**:
   ```go
   // In internal/generator/commands.go
   {
       Name: "newcommand",
       Trigger: "/newcommand",
       // ...
   }
   ```

2. **Create Command Template**:
   ```go
   // In internal/generator/commands.go
   func generateNewCommand() string {
       // Command template
   }
   ```

3. **Test**:
   ```bash
   go test ./internal/generator
   ```

---

### Example 3: Adding a New Rule

1. **Add Rule File**:
   ```
   internal/rules/library/XX-category/new-rule.md
   ```

2. **Update Rules**:
   Rules are automatically embedded via `//go:embed`

3. **Test**:
   ```bash
   go test ./internal/rules
   ```

---

## Troubleshooting

### Issue: Build Fails

**Solution**:
1. Check Go version: `go version`
2. Update dependencies: `go mod download`
3. Clean build: `make clean && make build`

---

### Issue: Tests Fail

**Solution**:
1. Run tests individually: `go test -v ./internal/generator`
2. Check test files for issues
3. Verify test data

---

### Issue: Binary Too Large

**Solution**:
1. Check embedded files size
2. Optimize rules library
3. Use compression if needed

---

## Code Style

### Formatting

```bash
make fmt

# Or
go fmt ./...
```

### Linting

```bash
make lint

# Or
golangci-lint run
```

### Style Guidelines

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` for formatting
- Document exported functions
- Write tests for new code

---

## Related Pages

- [Contributing Guide](Contributing) - How to contribute
- [API Reference](API-Reference) - API documentation
- [Architecture](Architecture) - System architecture
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

