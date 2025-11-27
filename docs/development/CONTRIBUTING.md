# Contributing Guidelines

## Documentation Organization

### Root Directory Rules

**Keep the root directory clean!** Only these files should be in the root:

- `README.md` - Project overview (required)
- `CHANGELOG.md` - Version history (required)
- `go.mod`, `go.sum` - Go dependencies
- `package.json`, `package-lock.json` - npm package config
- `Makefile` - Build commands
- Configuration files (`.gitignore`, `.npmignore`, etc.)

### Documentation Structure

All documentation should be organized in the `docs/` directory:

```
docs/
├── README.md                    # Documentation index
├── development/                 # Development guides
│   ├── BUILD.md
│   └── TESTING.md
├── release/                     # Release documentation
│   ├── RELEASE_CHECKLIST.md
│   ├── RELEASE_NOTES_*.md
│   └── SOCIAL_MEDIA_ANNOUNCEMENT.md
├── security/                    # Security documentation
│   └── SECURITY_AUDIT.md
└── [general docs]              # Other documentation
    ├── DOCUMENTATION_REVIEW.md
    ├── LAUNCH_CHECKLIST.md
    └── prompt.md
```

## 📝 Where to Create Documentation

### Development Documentation → `docs/development/`

Create here:
- **Build guides** → `docs/development/BUILD.md`
- **Testing documentation** → `docs/development/TESTING.md`
- **Development setup guides** → `docs/development/SETUP.md`
- **Architecture decisions** → `docs/development/ARCHITECTURE.md`
- **API documentation** → `docs/development/API.md`
- **Code style guides** → `docs/development/STYLE.md`
- **Debugging guides** → `docs/development/DEBUGGING.md`

**Examples:**
- How to build the project
- How to run tests
- How to set up development environment
- Code architecture and design decisions
- API reference documentation

### Release Documentation → `docs/release/`

Create here:
- **Release checklists** → `docs/release/RELEASE_CHECKLIST.md`
- **Release notes** → `docs/release/RELEASE_NOTES_v*.md`
- **Social media announcements** → `docs/release/SOCIAL_MEDIA_ANNOUNCEMENT.md`
- **Deployment guides** → `docs/release/DEPLOYMENT.md`
- **Version migration guides** → `docs/release/MIGRATION_v*.md`

**Examples:**
- Steps to create a release
- Release notes for each version
- Social media templates
- Deployment procedures
- Breaking changes and migration guides

### Security Documentation → `docs/security/`

Create here:
- **Security audits** → `docs/security/SECURITY_AUDIT.md`
- **Security policies** → `docs/security/POLICY.md`
- **Vulnerability reports** → `docs/security/VULNERABILITIES.md`
- **Security best practices** → `docs/security/BEST_PRACTICES.md`

**Examples:**
- Security audit reports
- Security policies and procedures
- Known vulnerabilities
- Security guidelines for contributors

### User Documentation → `docs/user/` (if needed)

Create here:
- **User guides** → `docs/user/GUIDE.md`
- **FAQ** → `docs/user/FAQ.md`
- **Troubleshooting** → `docs/user/TROUBLESHOOTING.md`
- **Examples** → `docs/user/EXAMPLES.md`

**Examples:**
- How to use the CLI
- Common questions
- Troubleshooting steps
- Usage examples

### General Documentation → `docs/`

Create here (root of docs/):
- **Project overview** → `docs/prompt.md`
- **Launch checklists** → `docs/LAUNCH_CHECKLIST.md`
- **Documentation reviews** → `docs/DOCUMENTATION_REVIEW.md`
- **Standup notes** → `docs/STANDUP.md`
- **Project planning** → `docs/PLANNING.md`

**Examples:**
- Project description/prompt
- Launch preparation
- Documentation audits
- Meeting notes
- Project planning documents

### Rules for Creating Documentation

1. **Check existing structure first** - Look for similar docs before creating new ones
2. **Use appropriate subdirectory** - Match the category (development/release/security)
3. **Create subdirectory if needed** - If a new category emerges, create it
4. **Update docs/README.md** - Always add new docs to the index
5. **Follow naming conventions**:
   - Use UPPERCASE for important docs: `BUILD.md`, `TESTING.md`
   - Use descriptive names: `RELEASE_NOTES_v1.0.0.md`
   - Use kebab-case for guides: `user-guide.md`, `api-reference.md`
6. **Never create `.md` files in root** (except README.md and CHANGELOG.md)
   - **Enforcement**: Run `./scripts/check-docs-organization.sh` before committing
   - **CI/CD**: Automated checks will block PRs with root-level `.md` files
   - **For Generated Projects**: All docs must live under `Docs/` (capital D) with canonical structure
   - **Rule Reference**: `.cursor/rules/library/11-documentation/docs-folder-structure.md`

### Quick Decision Tree

```
New documentation needed?
│
├─ Is it about building/testing/development?
│  └─ → docs/development/
│
├─ Is it about releases/deployment?
│  └─ → docs/release/
│
├─ Is it about security?
│  └─ → docs/security/
│
├─ Is it for end users?
│  └─ → docs/user/ (create if needed)
│
└─ Is it general project documentation?
   └─ → docs/ (root of docs/)
```

### Build Artifacts

Build artifacts should never be committed:
- `doplan` (binary)
- `coverage.out`
- `*.test` files
- Any generated files

These are already in `.gitignore`.

