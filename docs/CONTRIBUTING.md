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

### Rules

1. **Never add `.md` files to root** (except README.md and CHANGELOG.md)
2. **Organize by category** - Use subdirectories for related docs
3. **Update docs/README.md** - Add new docs to the index
4. **Keep root minimal** - Only essential project files

### Build Artifacts

Build artifacts should never be committed:
- `doplan` (binary)
- `coverage.out`
- `*.test` files
- Any generated files

These are already in `.gitignore`.

