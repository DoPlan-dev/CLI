# GitHub Workflow Automation

## Branching Strategy
- Use conventional branch naming: `feature/`, `bugfix/`, `hotfix/`, `release/`
- Main branch: `main` (production-ready code)
- Development branch: `develop` (integration branch)

## Commit Messages
- Use conventional commit format: `type(scope): description`
- Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
- Example: `feat(auth): add OAuth2 login`

## Pull Requests
- Require at least one approval before merge
- All CI checks must pass
- Link related issues
- Include description of changes

## Automation
- Auto-commit on `/finished` command
- Auto-push to current branch
- Update CHANGELOG.md on significant changes
- Trigger CI workflow on push

