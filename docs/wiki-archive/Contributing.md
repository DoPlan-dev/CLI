# Contributing Guide

Thank you for your interest in contributing to DoPlan CLI! This guide will help you get started.

---

## How to Contribute

There are many ways to contribute:

- **Report Bugs**: Open an issue describing the bug
- **Suggest Features**: Propose new features or improvements
- **Submit Code**: Fix bugs or add features via pull requests
- **Improve Documentation**: Help make our docs better
- **Answer Questions**: Help others in discussions

---

## Development Setup

### Prerequisites

- **Go** >= 1.23.0
- **Node.js** >= 14.0.0 (for npx wrapper)
- **Git** for version control
- **Make** (optional, for build commands)

### Getting Started

1. **Fork the Repository**

   Visit [GitHub](https://github.com/DoPlan-dev/CLI) and fork the repository.

2. **Clone Your Fork**

   ```bash
   git clone https://github.com/YOUR_USERNAME/CLI.git
   cd CLI
   ```

3. **Add Upstream Remote**

   ```bash
   git remote add upstream https://github.com/DoPlan-dev/CLI.git
   ```

4. **Install Dependencies**

   ```bash
   # Go dependencies are managed via go.mod
   go mod download
   ```

5. **Build the Project**

   ```bash
   make build
   # Or
   go build -o doplan ./cmd/doplan
   ```

6. **Run Tests**

   ```bash
   make test
   # Or
   go test ./...
   ```

---

## Code Style Guide

### Go Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` for formatting
- Use `golint` for linting
- Follow existing code patterns

### Formatting

```bash
# Format code
go fmt ./...

# Or use goimports
goimports -w .
```

### Linting

```bash
# Run linter
golangci-lint run
```

### Naming Conventions

- **Packages**: lowercase, single word
- **Exported functions**: PascalCase
- **Unexported functions**: camelCase
- **Constants**: PascalCase or UPPER_CASE
- **Variables**: camelCase

### Code Organization

- Keep functions focused and small
- Use meaningful names
- Add comments for exported functions
- Group related code together

---

## Testing Guidelines

### Writing Tests

- Write tests for all new features
- Write tests for bug fixes
- Aim for high test coverage
- Use table-driven tests when appropriate

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestName ./...
```

### Test Structure

```go
func TestFunctionName(t *testing.T) {
    // Arrange
    // Act
    // Assert
}
```

### Integration Tests

- Place integration tests in `*_integration_test.go` files
- Use build tags if needed: `//go:build integration`

---

## Documentation Guidelines

### Code Documentation

- Document all exported functions, types, and constants
- Use clear, concise comments
- Include examples when helpful

**Example**:
```go
// GenerateProject creates a new project structure with the given
// configuration. It returns an error if project generation fails.
//
// Example:
//   err := GenerateProject(&ProjectRequest{
//       Name: "my-project",
//       IDE:  "cursor",
//   })
func GenerateProject(request *ProjectRequest) error {
    // ...
}
```

### Markdown Documentation

- Use clear headings
- Include code examples
- Keep formatting consistent
- Update related docs when making changes

### Documentation Structure

- Keep docs in `docs/` directory
- Follow existing structure
- Update `docs/README.md` when adding new docs

---

## Pull Request Process

### Before Submitting

1. **Update Your Fork**

   ```bash
   git fetch upstream
   git checkout main
   git merge upstream/main
   ```

2. **Create a Branch**

   ```bash
   git checkout -b feature/your-feature-name
   # Or
   git checkout -b fix/your-bug-fix
   ```

3. **Make Your Changes**

   - Write code
   - Add tests
   - Update documentation
   - Follow code style

4. **Test Your Changes**

   ```bash
   go test ./...
   go build ./...
   ```

5. **Commit Your Changes**

   Use [Conventional Commits](https://www.conventionalcommits.org/):

   ```bash
   git commit -m "feat: add new feature"
   git commit -m "fix: fix bug in generator"
   git commit -m "docs: update installation guide"
   ```

### Submitting a Pull Request

1. **Push Your Branch**

   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create Pull Request**

   - Go to GitHub
   - Click "New Pull Request"
   - Select your branch
   - Fill out the PR template

3. **PR Description**

   Include:
   - What changes were made
   - Why the changes were made
   - How to test the changes
   - Related issues

4. **Wait for Review**

   - Address review comments
   - Make requested changes
   - Update PR as needed

### PR Checklist

- [ ] Code follows style guidelines
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] All tests pass
- [ ] No merge conflicts
- [ ] PR description complete

---

## Issue Reporting

### Before Reporting

1. **Search Existing Issues**

   Check if the issue already exists.

2. **Check Documentation**

   Verify it's not covered in docs.

3. **Reproduce the Issue**

   Ensure you can reproduce it consistently.

### Creating an Issue

Use the issue template and include:

- **Title**: Clear, descriptive title
- **Description**: Detailed description
- **Steps to Reproduce**: Step-by-step instructions
- **Expected Behavior**: What should happen
- **Actual Behavior**: What actually happens
- **Environment**: OS, Go version, etc.
- **Screenshots**: If applicable

### Bug Report Template

```markdown
**Description**
Brief description of the bug.

**Steps to Reproduce**
1. Step one
2. Step two
3. Step three

**Expected Behavior**
What should happen.

**Actual Behavior**
What actually happens.

**Environment**
- OS: macOS/Windows/Linux
- Go Version: 1.23.0
- DoPlan CLI Version: 1.0.0

**Additional Context**
Any other relevant information.
```

---

## Feature Requests

### Before Requesting

1. **Search Existing Issues**

   Check if the feature was already requested.

2. **Consider Alternatives**

   Is there a workaround?

3. **Think About Use Cases**

   Who would benefit? How would it be used?

### Creating a Feature Request

Include:

- **Title**: Clear feature name
- **Description**: Detailed description
- **Use Cases**: Who would use it and how
- **Alternatives**: Other solutions considered
- **Additional Context**: Any other relevant information

---

## Development Workflow

### Branch Strategy

- **main**: Production-ready code
- **develop**: Development branch
- **feature/**: New features
- **fix/**: Bug fixes
- **docs/**: Documentation updates

### Commit Messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation
- `style:` Formatting
- `refactor:` Code refactoring
- `test:` Tests
- `chore:` Maintenance

**Examples**:
```
feat: add support for new IDE
fix: resolve binary download issue
docs: update installation guide
```

### Code Review

- Be respectful and constructive
- Focus on code, not person
- Explain your reasoning
- Be open to feedback

---

## Contribution Examples

### Example 1: Fixing a Bug

```bash
# 1. Create branch
git checkout -b fix/binary-download-error

# 2. Fix the bug
# ... make changes ...

# 3. Add tests
# ... write tests ...

# 4. Test
go test ./...

# 5. Commit
git commit -m "fix: resolve binary download error on Windows"

# 6. Push
git push origin fix/binary-download-error

# 7. Create PR
```

### Example 2: Adding a Feature

```bash
# 1. Create branch
git checkout -b feature/add-new-ide-support

# 2. Implement feature
# ... write code ...

# 3. Add tests
# ... write tests ...

# 4. Update docs
# ... update documentation ...

# 5. Test
go test ./...
go build ./...

# 6. Commit
git commit -m "feat: add support for new IDE"

# 7. Push
git push origin feature/add-new-ide-support

# 8. Create PR
```

### Example 3: Improving Documentation

```bash
# 1. Create branch
git checkout -b docs/update-installation-guide

# 2. Update docs
# ... edit markdown files ...

# 3. Commit
git commit -m "docs: update installation guide with new steps"

# 4. Push
git push origin docs/update-installation-guide

# 5. Create PR
```

---

## Getting Help

### Resources

- **Documentation**: Check [docs/](https://github.com/DoPlan-dev/CLI/tree/main/docs)
- **Issues**: Search [GitHub Issues](https://github.com/DoPlan-dev/CLI/issues)
- **Discussions**: Join [GitHub Discussions](https://github.com/DoPlan-dev/CLI/discussions)

### Questions

- Open a discussion for questions
- Check existing issues first
- Be clear and specific

---

## Code of Conduct

We follow the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/). Please:

- Be respectful
- Be inclusive
- Be constructive
- Focus on what's best for the community

---

## Recognition

Contributors will be:
- Listed in CONTRIBUTORS.md
- Recognized in release notes
- Thanked in discussions

---

## Related Pages

- [Development Documentation](Development) - Development setup
- [Code of Conduct](Code-of-Conduct) - Community guidelines
- [FAQ](FAQ) - Frequently asked questions
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

