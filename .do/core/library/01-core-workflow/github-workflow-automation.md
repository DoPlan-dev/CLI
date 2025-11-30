# GitHub Workflow Automation

## Overview
This rule defines automated GitHub workflows for version control, branching, committing, pushing, releases, and changelog management.

## Branching Strategy

### Main Branches
- **main** - Production-ready code
- **develop** - Integration branch for features

### Supporting Branches
- **feature/** - New features (e.g., `feature/user-auth`)
- **bugfix/** - Bug fixes (e.g., `bugfix/login-error`)
- **hotfix/** - Critical production fixes (e.g., `hotfix/security-patch`)
- **release/** - Release preparation (e.g., `release/v1.2.0`)

### Branch Naming Convention
- Use lowercase with hyphens
- Include type prefix: `feature/`, `bugfix/`, `hotfix/`, `release/`
- Descriptive names: `feature/user-authentication`, not `feature/auth`

## Commit Automation

### Conventional Commits
Always use [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Commit Types
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, missing semicolons, etc.)
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `chore`: Maintenance tasks
- `ci`: CI/CD changes
- `build`: Build system changes
- `revert`: Revert previous commit

### Automated Commit Process
1. **Pre-commit Checks:**
   - Lint code
   - Run tests
   - Check formatting
   - Validate commit message format

2. **Auto-staging:**
   - Stage modified files automatically
   - Exclude files in `.gitignore`
   - Ask for confirmation if many files changed

3. **Commit Message Generation:**
   - Analyze staged changes
   - Generate appropriate commit type
   - Create descriptive message
   - Include scope if applicable

4. **Post-commit:**
   - Push to remote branch
   - Create PR if on feature branch
   - Update changelog if needed

## Push Automation

### Automatic Push Rules
- **Feature branches:** Push automatically after commit
- **Main/develop:** Require PR approval before merge
- **Hotfix branches:** Push immediately, create PR for review

### Push Triggers
- After successful commit
- After pre-commit checks pass
- On `/build` command completion
- On `/finished` command

## Release Automation

### Release Workflow
1. **Version Bump:**
   - Update `package.json` version
   - Update `CHANGELOG.md`
   - Create git tag

2. **Release Branch:**
   - Create `release/vX.Y.Z` branch
   - Merge to `main` and `develop`
   - Tag release

3. **GitHub Release:**
   - Create GitHub release from tag
   - Generate release notes from changelog
   - Attach build artifacts
   - Publish release

### Release Types
- **Major (X.0.0):** Breaking changes
- **Minor (0.X.0):** New features, backward compatible
- **Patch (0.0.X):** Bug fixes, backward compatible

## Changelog Automation

### Changelog Format
Follow [Keep a Changelog](https://keepachangelog.com/) format:

```markdown
# Changelog

## [Unreleased]

## [1.2.0] - 2025-11-23

### Added
- New feature X
- New feature Y

### Changed
- Improved performance of Z

### Fixed
- Bug in A

### Security
- Security fix for B
```

### Changelog Updates
- **Automatic:** Update on every commit with conventional commit
- **Manual:** Use `/add-to-changelog` command
- **Release:** Generate release section from commits

### Changelog Generation
1. Parse commits since last release
2. Group by commit type
3. Format according to Keep a Changelog
4. Add to CHANGELOG.md
5. Commit changelog update

## GitHub Actions Workflows

### Required Workflows
1. **CI Workflow** - Run tests, lint, build on every push/PR
2. **Release Workflow** - Automate releases on version tag
3. **Changelog Workflow** - Update changelog on commits
4. **Branch Protection** - Enforce PR reviews for main/develop

### Workflow Triggers
- `push` - On push to any branch
- `pull_request` - On PR creation/update
- `release` - On release creation
- `workflow_dispatch` - Manual trigger

## Agent Responsibilities

### Release Captain
- Manages release process
- Creates release branches
- Coordinates version bumps
- Publishes releases

### Engineering Lead
- Reviews automated commits
- Approves PRs
- Manages branch protection
- Oversees CI/CD

### Documentation Lead
- Maintains changelog format
- Ensures release notes quality
- Documents workflow changes

## Best Practices

1. **Always use conventional commits**
2. **Never commit directly to main/develop**
3. **Create feature branches for all changes**
4. **Keep branches short-lived**
5. **Update changelog with every significant change**
6. **Tag releases immediately**
7. **Use semantic versioning**
8. **Automate everything possible**

## Integration with DoPlan Commands

- `/build` - Automatically commits and pushes after task completion
- `/finished` - Marks task done, commits changes
- `/ship` - Triggers release workflow
- `/add-to-changelog` - Manually add changelog entry

## References
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Keep a Changelog](https://keepachangelog.com/)
- [Semantic Versioning](https://semver.org/)
- [GitHub Flow](https://guides.github.com/introduction/flow/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

