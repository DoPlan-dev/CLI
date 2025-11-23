# Frequently Asked Questions (FAQ)

Common questions about DoPlan CLI, organized by category.

---

## General Questions

### What is DoPlan CLI?

DoPlan CLI is a zero-install command-line tool that generates complete, production-ready project structures with a hierarchical AI agency system. It helps you bootstrap projects in seconds instead of hours.

### Who is DoPlan CLI for?

DoPlan CLI is perfect for:
- **Solo Developers** who want to focus on building, not configuration
- **Small Teams** looking to standardize their development workflow
- **Professionals** who need production-ready project structures from day one
- **Anyone** who wants to leverage AI agents for faster development

### How does DoPlan CLI work?

DoPlan CLI generates a complete project structure with:
- 18 specialized AI agents
- 1000+ embedded rules library
- Command system for interacting with agents
- Complete project structure and boilerplate
- GitHub workflows for CI/CD

### Is DoPlan CLI free?

Yes, DoPlan CLI is open source and free to use. See the [LICENSE](https://github.com/DoPlan-dev/CLI/blob/main/LICENSE) file for details.

### What makes DoPlan CLI different?

Key differentiators:
- **Zero-install**: Use via `npx` without installation
- **Offline-first**: Works completely offline after first run
- **Transparent**: All AI logic in markdown files
- **Complete**: Generates everything you need, not just boilerplate
- **IDE-agnostic**: Works with 6 AI-powered IDEs

### Do I need to know Go to use DoPlan CLI?

No! DoPlan CLI is written in Go, but you don't need to know Go to use it. It generates projects in any language/framework you choose.

---

## Installation Questions

### Do I need to install anything?

The easiest way is to use `npx` - no installation required:
```bash
npx @doplan-dev/cli
```

However, you can install it permanently if you prefer. See the [Installation Guide](Installation).

### What are the prerequisites?

- **Node.js** >= 14.0.0 (for npx wrapper)
- **Go** >= 1.23.0 (only if building from source)

### Which platforms are supported?

- macOS (Intel and Apple Silicon)
- Windows
- Linux (amd64 and arm64)
- Docker

### Can I use DoPlan CLI without Node.js?

If you install the binary directly or build from source, you don't need Node.js. However, the `npx` method requires Node.js.

### How do I update DoPlan CLI?

- **npx**: Always uses the latest version automatically
- **Homebrew**: `brew upgrade doplan`
- **Scoop**: `scoop update doplan`
- **Direct binary**: Download the latest release

See the [Installation Guide](Installation) for details.

---

## Usage Questions

### How do I get started?

1. Run `npx @doplan-dev/cli`
2. Enter your project name
3. Select your IDE
4. Open the project in your IDE
5. Type `/tell` to capture your idea

See the [Quick Start Guide](Quick-Start) for a complete tutorial.

### What commands are available?

Core commands:
- `/tell` - Capture idea
- `/improve` - Brainstorm
- `/write` - Generate plans
- `/build` - Start coding
- `/finished` - Complete task

See the [Commands Reference](Commands) for all commands.

### What is the typical workflow?

1. `/tell` - Capture your project idea
2. `/improve` - Brainstorm with the team
3. `/write` - Generate planning documents
4. `/good` - Approve the plan
5. `/tasks` - Generate implementation tasks
6. `/build` - Start coding
7. `/finished` - Complete tasks

See the [Workflow Guide](Workflow) for details.

### Can I skip steps in the workflow?

Yes, but following the workflow order gives the best results. You can:
- Edit documents directly
- Use `/change` to modify documents
- Build specific tasks with `/build <task_id>`

### How do I know which command to use?

Start with `/tell` and follow the workflow. The commands guide you through the process. See the [Commands Reference](Commands) for details.

### Can I use DoPlan CLI with existing projects?

DoPlan CLI is designed for new projects. However, you can:
- Generate a new project and copy structure
- Manually add agents and commands to existing projects
- Use the rules library in existing projects

---

## Technical Questions

### What IDEs are supported?

DoPlan CLI supports 6 AI-powered IDEs:
- Cursor
- Claude Code
- Antigravity
- Windsurf
- Cline
- OpenCode

See [IDE Integration](IDE-Integration) for setup guides.

### How do the AI agents work?

Each agent has a specific role and expertise. When you use commands, relevant agents are activated. All agent definitions are in `.cursor/agents/` as markdown files.

See the [Agents Documentation](Agents) for details.

### What is the rules library?

The rules library contains 1000+ best practices organized into 15 categories:
- Core Workflow
- Languages
- Frameworks
- Databases
- Testing
- CI/CD
- Security
- And more!

See the [Rules Library Documentation](Rules) for details.

### Can I customize agents and commands?

Yes! All agent definitions are in `.cursor/agents/` and commands are in `.cursor/commands/`. You can modify them to fit your needs.

See the [Customization Guide](Customization) for details.

### How does the command system work?

Commands are defined in `.cursor/commands/` as markdown files. Each command specifies:
- Trigger pattern
- Actions to perform
- Agents to activate
- Files to read/modify

See the [Commands Reference](Commands) for details.

### Does DoPlan CLI require internet?

Only for the first run (to download the binary via npx). After that, it works completely offline.

### How big is the generated project?

The generated project structure is minimal. The rules library is embedded but doesn't significantly increase project size. Most of the size comes from your actual source code.

---

## Troubleshooting Questions

### Command not found after installation

**Solution**: Ensure the binary is in your PATH. See [Troubleshooting](Troubleshooting) for details.

### npx download fails

**Solution**: 
1. Check your internet connection
2. Clear npm cache: `npm cache clean --force`
3. Try again

### Wrong architecture binary

**Solution**: Download the correct binary for your platform:
- macOS Intel: `doplan-darwin-amd64`
- macOS Apple Silicon: `doplan-darwin-arm64`
- Linux: `doplan-linux-amd64` or `doplan-linux-arm64`

### IDE doesn't recognize commands

**Solution**: 
1. Ensure you're using a supported IDE
2. Check IDE configuration files
3. See [IDE Integration](IDE-Integration) for setup

### Agents not working

**Solution**:
1. Check that agent files exist in `.cursor/agents/`
2. Verify command definitions in `.cursor/commands/`
3. Check project state in `.plan/active_state.json`

### Build command fails

**Solution**:
1. Check that tasks exist in `.plan/TASKS.md`
2. Verify project state in `.plan/active_state.json`
3. Ensure you've approved the plan with `/good`

See [Troubleshooting](Troubleshooting) for more solutions.

---

## Contributing Questions

### How can I contribute?

See the [Contributing Guide](Contributing) for details. You can:
- Report bugs
- Suggest features
- Submit pull requests
- Improve documentation

### Where do I report bugs?

Open an issue on [GitHub](https://github.com/DoPlan-dev/CLI/issues).

### How do I suggest features?

Open a feature request on [GitHub](https://github.com/DoPlan-dev/CLI/issues).

### Can I add new agents?

Yes! You can add custom agents to your project. See the [Customization Guide](Customization).

### Can I add new commands?

Yes! You can add custom commands to your project. See the [Customization Guide](Customization).

### How do I contribute to the rules library?

Rules are embedded in the DoPlan CLI binary. To contribute:
1. Fork the repository
2. Add your rules to `internal/rules/library/`
3. Submit a pull request

---

## Project-Specific Questions

### What project types are supported?

DoPlan CLI generates "Fullstack" projects by default. You can customize the generated structure to fit any project type.

### Can I use DoPlan CLI for mobile apps?

Yes! The generated structure can be customized for any project type, including mobile apps.

### Can I use DoPlan CLI for backend-only projects?

Yes! You can customize the generated structure to focus on backend development.

### What languages/frameworks are supported?

The rules library includes rules for:
- Languages: Go, TypeScript, JavaScript, Python, etc.
- Frameworks: Next.js, React, Express, etc.
- Databases: PostgreSQL, MongoDB, etc.

You can use any language/framework - the rules library provides best practices.

---

## Still Have Questions?

- Check the [Troubleshooting Guide](Troubleshooting) for common issues
- Review the [Commands Reference](Commands) for command details
- See the [Workflow Guide](Workflow) for workflow details
- Open an issue on [GitHub](https://github.com/DoPlan-dev/CLI/issues)

---

## Related Pages

- [Installation](Installation) - Installation guide
- [Quick Start](Quick-Start) - Getting started
- [Commands Reference](Commands) - All commands
- [Troubleshooting](Troubleshooting) - Common issues
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

