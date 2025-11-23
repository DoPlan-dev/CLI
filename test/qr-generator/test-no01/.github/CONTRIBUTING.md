# Contributing to QR Code Generator API

Thank you for your interest in contributing! This document provides guidelines and instructions for contributing to the project.

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on constructive feedback
- Respect different viewpoints and experiences

## Getting Started

1. **Fork the repository**
   ```bash
   git clone https://github.com/DoPlan-dev/test-no01.git
   cd test-no01
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Create a branch**
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

4. **Make your changes**
   - Follow the project's code style
   - Write clear, self-documenting code
   - Add comments for complex logic
   - Update documentation as needed

5. **Test your changes**
   ```bash
   npm run lint
   npm run build
   npm test  # when tests are available
   ```

6. **Commit your changes**
   ```bash
   git commit -m "feat: add new feature"
   ```
   Follow [Conventional Commits](https://www.conventionalcommits.org/) format:
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation
   - `style:` for formatting
   - `refactor:` for code refactoring
   - `test:` for tests
   - `chore:` for maintenance

7. **Push and create a Pull Request**
   ```bash
   git push origin feature/your-feature-name
   ```

## Development Workflow

### Branch Strategy
- `main` - Production-ready code
- `develop` - Development branch
- `feature/*` - New features
- `fix/*` - Bug fixes
- `docs/*` - Documentation updates

### Pull Request Process

1. **Update your branch**
   ```bash
   git fetch origin
   git rebase origin/develop
   ```

2. **Ensure all checks pass**
   - Linting passes
   - Type checking passes
   - Build succeeds
   - Tests pass (when available)

3. **Write a clear PR description**
   - Use the PR template
   - Describe what and why
   - Link related issues
   - Add screenshots if applicable

4. **Request review**
   - Assign reviewers
   - Respond to feedback promptly
   - Make requested changes

## Coding Standards

### TypeScript
- Use TypeScript strict mode
- Define types for all functions
- Avoid `any` type
- Use meaningful variable names

### Code Style
- Follow ESLint rules
- Use Prettier for formatting
- Keep functions small and focused
- Write self-documenting code

### Testing
- Write tests for new features
- Maintain or improve test coverage
- Test edge cases
- Update tests when changing functionality

## Project Structure

```
test-no01/
├── app/              # Next.js app directory
├── components/       # React components
├── lib/              # Utility functions and services
├── types/            # TypeScript type definitions
├── tests/            # Test files
└── public/           # Static assets
```

## Reporting Issues

- Use the issue templates
- Provide clear descriptions
- Include steps to reproduce
- Add screenshots if applicable
- Specify environment details

## Feature Requests

- Use the feature request template
- Explain the problem it solves
- Provide use cases
- Consider alternatives

## Questions?

- Open a discussion
- Create a question issue
- Check existing documentation

Thank you for contributing! 🎉

