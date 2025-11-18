# DoPlan v0.0.18-beta Development Documentation

Welcome to the DoPlan v0.0.18-beta development documentation. This directory contains all planning, implementation, and reference documentation for the v0.0.18-beta release.

## 📚 Documentation Index

### 🎯 Getting Started

1. **[V0.0.18 Implementation Guide](V0.0.18_IMPLEMENTATION_GUIDE.md)** - **START HERE**
   - Master guide linking all documentation
   - Implementation roadmap
   - Quick links by task
   - Implementation checklist

2. **[Preparation Checklist](PREPARATION_CHECKLIST.md)**
   - Pre-implementation assessment
   - Architecture planning
   - Migration strategy preparation
   - Testing strategy preparation
   - Implementation readiness checklist

### 📋 Core Planning Documents

3. **[Migration Strategy](MIGRATION_STRATEGY.md)**
   - Current vs. new structure comparison
   - Step-by-step migration process
   - Safety measures and rollback procedures
   - Migration wizard flow
   - Troubleshooting guide

### 🛠️ Implementation Documents

4. **[TUI Wizard Flow](TUI_WIZARD_FLOW.md)**
   - Detailed screen mockups for all wizards
   - State machine diagrams
   - State transition logic
   - Keyboard shortcuts
   - Styling guide
   - Responsive design handling

5. **[IDE Integration Specifics](IDE_INTEGRATION_SPECIFICS.md)**
   - Step-by-step setup for each IDE:
     - Cursor
     - VS Code + Copilot
     - Kiro
     - Windsurf
     - Qoder
     - CLI tools (Gemini, Claude)
     - Generic/Other IDEs
   - Integration verification
   - Troubleshooting

6. **[Testing Scenarios](TESTING_SCENARIOS.md)**
   - Comprehensive test cases
   - Expected outcomes
   - Test environment setup
   - Automated test scripts
   - Performance testing

7. **[Error Handling](ERROR_HANDLING.md)**
   - Error code system
   - Error recovery strategies
   - User-friendly error messages
   - Error logging
   - Critical error scenarios

### 📖 Advanced Documentation

8. **[Advanced Usage](ADVANCED_USAGE.md)**
   - Advanced patterns and techniques
   - Development workflows
   - Common patterns
   - Debugging techniques
   - Performance optimization

9. **[Troubleshooting Guide](TROUBLESHOOTING_GUIDE.md)**
   - Diagnostic tools
   - Common issues and solutions
   - Advanced troubleshooting
   - Recovery procedures

10. **[Workflow Examples](WORKFLOW_EXAMPLES.md)**
    - Real-world workflow examples
    - Daily development routines
    - Sprint planning examples

## 🎯 Quick Start

### For New Developers

1. Read the [V0.0.18 Implementation Guide](V0.0.18_IMPLEMENTATION_GUIDE.md)
2. Complete the [Preparation Checklist](PREPARATION_CHECKLIST.md)
3. Review [TUI Wizard Flow](TUI_WIZARD_FLOW.md) for UI patterns
4. Study [Error Handling](ERROR_HANDLING.md) for error patterns
5. Set up test environment: `make setup`

### For Implementing Features

1. Review relevant documentation section
2. Check code templates in `templates/`
3. Use utility scripts from `scripts/`
4. Follow patterns from [Advanced Usage](ADVANCED_USAGE.md)
5. Write tests per [Testing Scenarios](TESTING_SCENARIOS.md)

### For Troubleshooting

1. Check [Troubleshooting Guide](TROUBLESHOOTING_GUIDE.md)
2. Review [Error Handling](ERROR_HANDLING.md) for error codes
3. Check [Testing Scenarios](TESTING_SCENARIOS.md) for similar cases
4. Review [Migration Strategy](MIGRATION_STRATEGY.md) for migration issues

## 🔧 Code Templates

Ready-to-use code templates are in `templates/`:

- `wizard_base.go` - Wizard base structure
- `context_detector.go` - Context detection
- `ide_integration_base.go` - IDE integration
- `error_handler.go` - Error handling
- `migration_detector.go` - Migration detection
- `migration_backup.go` - Backup management
- `migration_config.go` - Config migration
- `migration_folders.go` - Folder migration
- `migration_migrator.go` - Main migration orchestrator

## 🛠️ Utility Scripts

All utility scripts are in `scripts/` with comprehensive documentation in `scripts/README.md`.

**Quick commands:**
```bash
make setup          # Set up test environment
make test           # Run all tests
make validate       # Validate migration
make deps           # Check dependencies
make benchmark      # Run benchmarks
```

## 📊 Implementation Status

See [V0.0.18 Implementation Guide](V0.0.18_IMPLEMENTATION_GUIDE.md) for the implementation roadmap and checklist.

## 🔗 Related Resources

- [Main Release Plan](../next-2-beta.md) - Original feature plan
- [Architecture Documentation](../ARCHITECTURE.md) - System architecture
- [Contributing Guide](../../CONTRIBUTING.md) - Contribution guidelines

## 📝 Notes

- All code should follow existing DoPlan patterns
- Use Bubble Tea for TUI, Lipgloss for styling
- Follow error handling patterns in [Error Handling](ERROR_HANDLING.md)
- Test all scenarios from [Testing Scenarios](TESTING_SCENARIOS.md)
- Document all new features and changes

## 🆘 Getting Help

If you encounter issues:
1. Check [Error Handling](ERROR_HANDLING.md) for error codes
2. Review [Testing Scenarios](TESTING_SCENARIOS.md) for similar cases
3. Check [Troubleshooting Guide](TROUBLESHOOTING_GUIDE.md)
4. Review [Migration Strategy](MIGRATION_STRATEGY.md) for migration issues

