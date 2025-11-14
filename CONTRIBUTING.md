# Contributing to DoPlan CLI

Thank you for your interest in contributing to DoPlan CLI! This document provides guidelines and instructions for contributing.

## Development Setup

### Prerequisites

- **Go 1.21+** (for building from source)
- **Git** (for version control)
- **GitHub CLI** (`gh`) - Optional, for GitHub automation features
- **Make** - For running common tasks

### Getting Started

1. **Fork the repository**
   ```bash
   git clone https://github.com/DoPlan-dev/CLI.git
   cd cli
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the project**
   ```bash
   make build
   ```

4. **Run tests**
   ```bash
   make test
   ```

## Development Workflow

### Making Changes

1. **Create a branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Write code following Go conventions
   - Add tests for new functionality
   - Update documentation as needed

3. **Run tests and linting**
   ```bash
   make test
   make lint
   make test-coverage
   ```

4. **Commit your changes**
   - Use conventional commits format
   - Example: `feat: add new statistics feature`
   - Example: `fix: resolve dashboard display issue`

5. **Push and create a PR**
   ```bash
   git push origin feature/your-feature-name
   ```

## Code Standards

### Go Style Guide

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` for formatting (run `make fmt`)
- Run `go vet` for static analysis (run `make vet`)
- Keep functions focused and small
- Add comments for exported functions and types

### Testing

- Write unit tests for all new functionality
- Maintain or improve test coverage
- Use table-driven tests when appropriate
- Test error cases and edge cases

### Documentation

- Update README.md for user-facing changes
- Update code comments for API changes
- Add examples for new features
- Update CHANGELOG.md for significant changes

## Pull Request Process

### Before Submitting

1. **Run all tests**
   ```bash
   make test
   make test-coverage
   ```

2. **Check formatting**
   ```bash
   make fmt
   git diff --exit-code
   ```

3. **Run linting**
   ```bash
   make lint
   ```

4. **Verify the build**
   ```bash
   make build
   ```

### PR Checklist

- [ ] Tests pass locally
- [ ] Code is formatted (`make fmt`)
- [ ] Linting passes (`make lint`)
- [ ] Documentation updated
- [ ] CHANGELOG.md updated (if applicable)
- [ ] Commit messages follow conventional commits

### PR Description

Include:
- Description of changes
- Related issues (if any)
- Testing instructions
- Screenshots (if UI changes)

## Release Process

### Versioning

We follow [Semantic Versioning](https://semver.org/):
- **MAJOR** version for incompatible API changes
- **MINOR** version for new functionality (backwards compatible)
- **PATCH** version for bug fixes (backwards compatible)

### Creating a Release

1. **Update version**
   - Update CHANGELOG.md
   - Create a git tag: `git tag -a v1.0.0 -m "Release v1.0.0"`

2. **Push tag**
   ```bash
   git push origin v1.0.0
   ```

3. **GitHub Actions will automatically**
   - Build binaries for all platforms
   - Create GitHub release
   - Generate changelog
   - Create Homebrew PR (if tap repository exists)

See [docs/development/RELEASE.md](docs/development/RELEASE.md) for detailed release process.

## Testing

### Unit Tests

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific package
go test ./internal/commands/...
```

### Integration Tests

```bash
# Run integration tests
make test-scripts

# Run specific test script
./scripts/test-cli.sh
./scripts/test-install.sh
./scripts/test-integration.sh
```

### Manual Testing

1. Build the binary: `make build`
2. Test in a sample project
3. Verify all commands work as expected

## Code Review

### Review Process

1. PRs are reviewed by maintainers
2. Address review comments promptly
3. Make requested changes
4. Wait for approval before merging

### Review Criteria

- Code quality and style
- Test coverage
- Documentation completeness
- Backwards compatibility
- Performance implications

## Questions?

- Open an issue for questions or discussions
- Check existing issues and PRs
- Review documentation in `docs/` directory

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

