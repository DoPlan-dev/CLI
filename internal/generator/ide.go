package generator

import (
	"fmt"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// IDEGenerator generates IDE-specific configuration files
type IDEGenerator struct{}

// Name returns the name of the generator
func (g *IDEGenerator) Name() string {
	return "IDE Configs"
}

// Generate creates IDE configuration files for all selected IDEs
func (g *IDEGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// Ensure IDEs list is populated (for backward compatibility)
	if len(request.IDEs) == 0 && request.IDE != "" {
		request.IDEs = []string{request.IDE}
	}

	// Generate config files for each selected IDE
	for _, ide := range request.IDEs {
		switch ide {
		case "Cursor":
			if err := generateCursorConfig(projectPath); err != nil {
				return fmt.Errorf("failed to generate Cursor config: %w", err)
			}
		case "Claude Code":
			if err := generateClaudeConfig(projectPath); err != nil {
				return fmt.Errorf("failed to generate Claude Code config: %w", err)
			}
		case "Antigravity":
			if err := generateAntigravityConfig(projectPath); err != nil {
				return fmt.Errorf("failed to generate Antigravity config: %w", err)
			}
		case "Windsurf":
			if err := generateWindsurfConfig(projectPath); err != nil {
				return fmt.Errorf("failed to generate Windsurf config: %w", err)
			}
		case "Cline":
			if err := generateClineConfig(projectPath); err != nil {
				return fmt.Errorf("failed to generate Cline config: %w", err)
			}
		case "OpenCode":
			if err := generateOpenCodeConfig(projectPath); err != nil {
				return fmt.Errorf("failed to generate OpenCode config: %w", err)
			}
		default:
			return fmt.Errorf("unsupported IDE: %s", ide)
		}
	}

	return nil
}

// generateCursorConfig generates .cursorrules file
func generateCursorConfig(projectPath string) error {
	content := `# DoPlan Project Configuration

## Overview
This project uses DoPlan's hierarchical AI agency structure with Cursor's advanced AI capabilities.

## Cursor Features
- **Agent Mode**: Use Ctrl+I to delegate complex tasks to the AI agent
- **Tab Autocomplete**: Multi-line code completions with context awareness
- **Inline Edit**: Select code and press Ctrl+K to describe changes
- **Ask Mode**: Use Ctrl+L for complex questions requiring multiple steps or deep codebase understanding
- **Codebase Understanding**: Cursor's retrieval model understands your entire codebase structure and relationships

## @Docs Feature
Add documentation sources that Cursor can reference automatically:

1. **Add documentation sources**:
   - Go to Settings → Features → Docs
   - Add directories like docs/, .do/, or specific markdown files
   - The AI will reference these when generating code

2. **Using @Docs in chat**:
   - Type @Docs to see available documentation sources
   - Use @filename or @folder to reference specific files
   - Cursor will automatically pull relevant context from your docs

3. **Documentation structure**:
   - Place project documentation in docs/ directory
   - Planning documents in .do/ directory
   - Rules and guidelines in .cursor/rules/ directory

4. **Keeping docs up to date**:
   - Ask Cursor: "Are any docs outdated based on recent changes?"
   - Use "Generate Cursor Rules" to create rules from code or conversations
   - Cursor can help identify and update stale documentation

## @Web Feature
Use @Web to enable web search in Cursor:
- Search the web for documentation, examples, and solutions
- Access up-to-date information from the internet
- Combine web results with your codebase context
- Useful for finding best practices and troubleshooting

## Model Context Protocol (MCP)
Cursor supports MCP for extended functionality:

1. **MCP for Documentation**:
   - Connect to internal documentation systems
   - Access project management tools
   - Integrate with knowledge bases
   - Pull context from external sources

2. **MCP Configuration**:
   - Create mcp.json in .cursor/ directory
   - Configure MCP servers and tools
   - CLI automatically detects and respects mcp.json
   - Same MCP servers work in both IDE and CLI

3. **Using MCP**:
   - Access tools and resources from MCP servers
   - Extend Cursor's capabilities with custom integrations
   - Connect to internal systems and APIs

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .do/core/agents/ (symlinked to .cursor/agents/)

Reference agents in chat using @agent-name (e.g., @frontend_lead, @backend_lead)

## Commands
All commands are defined in .do/core/commands/ (symlinked to .cursor/commands/). Type any command (e.g., /do, /plan) to activate it.

Available commands:
- /hey - Welcome, tutorial, and command introductions
- /do - Capture project idea, conduct meeting, and refine
- /plan - Generate documents & execution plan (subcommands: docs, content, everything, phases, next, etc.)
- /dev - Start coding with automatic completion detection
- /sys - System control panel (subcommands: status, optimize, performance, backup, restore, memory, state, feedback, github, security, permissions, access)

## Rules System

### Rules as Long-Term Memory
Rules serve as Cursor's long-term memory for your project:
- Capture project-specific patterns and conventions
- Remember architectural decisions and design choices
- Store domain knowledge and business logic
- Persist across conversations and sessions

### Rules Library
Stack-specific rules are organized in .cursor/rules/ directory. Each category folder (e.g., 01-core-workflow, 03-languages) is symlinked from .do/core/library/. These rules guide the AI's behavior for:
- Language-specific conventions (Go, JavaScript, TypeScript, Python)
- Framework best practices (React, Next.js, Express)
- Testing standards (Jest, Vitest, Go testing)
- Code quality and security guidelines

### Generate Cursor Rules
Create rules automatically from:
- **Code**: Extract patterns from existing codebase
- **Conversations**: Generate rules from chat discussions
- **Documentation**: Convert docs into actionable rules
- Use "Generate Cursor Rules" feature to create rules quickly

### Domain-Specific Rules (.mdc files)
Create custom rules files in .cursor/rules/ directory with .mdc extension:

` + "```markdown\n" + `---
description: Clear description of what the rule enforces
globs: path/to/files/*.ext, other/path/**/*
alwaysApply: true
---

Rule content here...
` + "```\n\n" + `

Example rule structure:
- **Frontmatter**: YAML metadata (description, globs, alwaysApply)
- **Content**: Markdown with clear guidelines
- **File references**: Use ` + "`[filename](mdc:path/to/file)`" + ` syntax
- **Code examples**: Include both DO and DON'T patterns

### Using Rules
- Rules in .cursorrules apply to the entire project
- Rules in .cursor/rules/*.mdc can target specific file patterns via globs
- Use alwaysApply: true for rules that should always be active
- Reference rules in chat: @rules or @library/path/to/rule.mdc

## Planning Process
For complex features and large changes, use Cursor's planning capabilities:

1. **Break down the task**: Ask Cursor to create a plan before implementing
   - "Create a plan to add user authentication"
   - "Plan the migration from REST to GraphQL"

2. **Review the plan**: Cursor will break down complex tasks into manageable steps
   - Review each step before execution
   - Modify the plan as needed

3. **Execute incrementally**: Implement one step at a time
   - Use Agent Mode (Ctrl+I) for each step
   - Review changes between steps

4. **Use Ask Mode for planning questions**: 
   - "What files will be affected by this change?"
   - "How should I structure this feature?"

## Documentation Maintenance
Keep your documentation up to date and AI-friendly:

1. **Keep docs current**: Update documentation as code changes
   - Cursor can help identify outdated docs
   - Ask: "Are any docs outdated based on recent changes?"

2. **LLM-optimized format**: Structure docs for both humans and AI:
   - Include concrete file references
   - Use clear section headings
   - Provide examples with file paths

3. **Use @Docs feature**: Add frequently referenced docs to @Docs
   - Architecture documentation
   - API specifications
   - Development workflows

4. **Timestamp tracking**: Include last-update timestamps in docs
   - Helps identify stale documentation
   - Format: <!-- Generated: YYYY-MM-DD HH:MM:SS UTC -->

## Large Codebases
For complex projects:
1. **Use Ask Mode (Ctrl+L)**: For questions requiring deep codebase understanding
2. **Leverage Rules**: Create comprehensive rules in .cursor/rules/ directory
3. **Documentation**: Add key docs to @Docs for automatic context
4. **Planning**: Use /plan command to break down complex tasks
5. **Agent Mode**: Use Ctrl+I for multi-file changes and complex refactoring

## Project State
Current project state is tracked in .do/active_state.json

## Workflow
1. **Welcome** (New Users): Type /hey for first-time tutorial and setup
2. **Add Documentation Sources**: Configure @Docs with your docs/, .do/, and README.md
3. **Reference Context**: Use @docs or @filename when asking questions
4. **Capture Ideas**: Type /do to capture your idea and conduct meeting
5. **Plan**: Type /plan to generate documents and execution plan
6. **Code**: Type /dev to start coding with Agent Mode (Ctrl+I) and automatic completion detection
7. **System**: Type /sys status to check progress and generate reports

For complex questions, use Ask Mode (Ctrl+L) instead of Composer (Ctrl+I).

## Context Management
- **@filename**: Reference specific files in chat
- **@folder**: Include entire directories
- **@agent-name**: Activate specific agent personalities
- **@docs**: Reference documentation sources
- **@rules**: Reference specific rule files
- **Ask Mode (Ctrl+L)**: Better for complex, multi-step questions
- **Agent Mode (Ctrl+I)**: Better for multi-file changes and refactoring

Cursor automatically understands your codebase structure and relationships through its retrieval model.

## Best Practices
- Keep .cursorrules focused on project-wide guidelines
- Create domain-specific rules in .cursor/rules/*.mdc files
- Add comprehensive documentation to @Docs
- Use Ask Mode for understanding, Agent Mode for implementing
- Reference rules and docs explicitly when asking complex questions

For full command list and detailed documentation, see README.md
`
	path := filepath.Join(projectPath, ".cursorrules")
	return utils.WriteFile(path, []byte(content))
}

// generateClaudeConfig generates CLAUDE.md file in docs/ directory
func generateClaudeConfig(projectPath string) error {
	content := `# Claude Code Project Configuration

## Overview
This project uses DoPlan's hierarchical AI agency structure with Claude Code, an agentic coding tool that lives in your terminal.

## Getting Started with Claude Code
Claude Code works directly from your terminal. From the project root:

` + "```bash\n" + `cd your-project
claude
` + "```\n\n" + `Claude Code can:
- Build features from descriptions in plain English
- Debug and fix issues automatically
- Navigate and understand your codebase
- Automate tedious tasks (linting, merge conflicts, release notes)
- Execute commands, edit files, and create commits directly

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .claude/agents/

Agents provide specialized knowledge for different aspects of development:
- Frontend/Backend leads for architecture
- Security lead for security best practices
- QA engineers for testing standards
- DevOps engineers for deployment automation

Reference agents in prompts: "Act as the frontend_lead and review this component"

## Commands
All commands are defined in .claude/commands/. These commands provide structured workflows:

- /hello - Welcome, tutorial, and command introductions (subcommands: goplan, meeting, plan, build, github)
- /tell - Capture your idea
- /meeting - Adaptive discovery meeting with progress tracking and timing
- /write - Generate documents & content (subcommands: plan, content, change, prd, architecture, design, etc.)
- /plan - Generate execution plan (subcommands: everything, phases, next, phase {no} tasks, phases tasks, all tasks)
- /build - Start coding with auto-completion detection
- /status - Show progress and reports (subcommands: report, full)
- /github - GitHub operations (subcommands: info, issue, milestone, ci, release)
- /state - State management (subcommands: snapshot, list, diff, restore)
- /feedback - Log feedback
- /security - Security review/audit (subcommands: review, audit, both)
- /permissions - Design RBAC system
- /optimize - Project optimization (subcommands: design, finance, performance)

## Rules
Stack-specific rules are organized in .claude/rules/library/ directory. Claude Code automatically references these rules when:
- Writing code in specific languages
- Using frameworks (React, Next.js, Express)
- Implementing tests (Jest, Vitest, Go testing)
- Following security and code quality guidelines

Rules are loaded automatically when Claude Code analyzes your project.

## Project State
Current project state is tracked in .do/active_state.json

Claude Code can read and update this state automatically during task execution.

## Advanced Features

### Sub-Agents
Claude Code supports sub-agents for specialized tasks:
- Delegate specific work to specialized agents
- Coordinate multiple agents for complex projects
- Each agent focuses on a specific domain or task
- Agents can collaborate and share context

### Skills
Extend Claude Code's capabilities with custom skills:
- Create reusable skills for common patterns
- Share skills across projects
- Build domain-specific expertise
- Integrate external tools and services

### Output Styles
Customize how Claude Code presents information:
- Control verbosity and detail level
- Format output for different contexts
- Configure default output preferences
- Adapt to your workflow preferences

### Hooks
Configure hooks to customize Claude Code's behavior:
- Pre-execution hooks for validation
- Post-execution hooks for cleanup
- Custom workflow triggers
- Integration with CI/CD pipelines

## Workflow Examples

### Build a Feature
` + "```bash\n" + `claude "Add user authentication with JWT tokens"
` + "```\n\n" + `### Debug an Issue
` + "```bash\n" + `claude -p "Fix the authentication bug: [paste error message]"
` + "```\n\n" + `### Review Code
` + "```bash\n" + `claude -p "Review this code for security vulnerabilities: @auth.ts"
` + "```\n\n" + `### Automate Tasks
` + "```bash\n" + `claude -p "If there are lint errors, fix them and raise a PR"
` + "```\n\n" + `### Common Workflows
- Feature development: From idea to implementation
- Bug fixing: Identify and resolve issues
- Code review: Automated code quality checks
- Refactoring: Improve code structure
- Documentation: Generate and update docs

## Unix Philosophy
Claude Code is composable and scriptable:
- Pipe output to Claude: ` + "`tail -f app.log | claude -p \"Alert me if errors appear\"`\n" + `- Use in CI/CD: ` + "`claude -p \"Generate release notes from git log\"`\n" + `- Chain commands: ` + "`claude \"Fix lint errors\" && claude \"Run tests\"`\n\n" + `## CLI Reference
Claude Code provides extensive CLI commands:
- Interactive mode: Just run 'claude'
- Prompt mode: 'claude -p "your prompt"'
- File mode: 'claude -f file.txt'
- Stream mode: For real-time processing
- Configuration: Customize via settings

For full CLI reference, see: https://code.claude.com/docs/en/cli-reference

## Privacy & Security
- Claude Code can use the Claude API or host on AWS/GCP
- Enterprise-grade security and compliance built-in
- Review Claude Code's settings for privacy options
- Configure authentication and access controls

For full command list and detailed documentation, see README.md

Learn more:
- Common Workflows: https://code.claude.com/docs/en/common-workflows
- Sub-Agents: https://code.claude.com/docs/en/sub-agents
- Skills: https://code.claude.com/docs/en/skills
- Output Styles: https://code.claude.com/docs/en/output-styles
- Hooks Guide: https://code.claude.com/docs/en/hooks-guide
`
	// Create docs directory if it doesn't exist
	docsDir := filepath.Join(projectPath, "docs")
	if err := utils.CreateDirectory(docsDir); err != nil {
		return fmt.Errorf("failed to create docs directory: %w", err)
	}

	path := filepath.Join(docsDir, "CLAUDE.md")
	return utils.WriteFile(path, []byte(content))
}

// generateAntigravityConfig generates .antigravity/config.md file
func generateAntigravityConfig(projectPath string) error {
	content := `# Antigravity Project Configuration

## Overview
This project uses DoPlan's hierarchical AI agency structure with Google Antigravity, an agent-first development platform.

## Antigravity Features
Antigravity is designed as an "agent-first" platform where AI agents are autonomous actors capable of:
- **Planning**: Creating task plans and implementation strategies
- **Executing**: Writing code, running tests, and making changes
- **Validating**: Checking code quality and test results
- **Iterating**: Refining solutions based on feedback

Key components:
- **Agent Manager**: Manage and configure autonomous agents
- **Editor**: Integrated code editor with AI assistance
- **Browser**: Agents can browse the web for research and documentation

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .antigravity/agents/

Agents work autonomously and can:
- Create task plans before implementation
- Generate implementation plans with clear steps
- Execute code changes across multiple files
- Validate their work through testing

Reference agents: "Use the frontend_lead agent to refactor this component"

## Commands
All commands are defined in .antigravity/commands/. These commands integrate with Antigravity's agent system:

- /hello - First-time welcome and tutorial experience
- /tell - Capture your idea (creates a task for agents)
- /meeting - Discovery meeting with adaptive speed options (agents generate improvement plans)
- /write - Generate documents (agents create documentation)
- /plan - Generate execution plan (agents create detailed plans)
- /build - Start coding (agents execute the plan)

## Rules
Stack-specific rules are organized in .antigravity/rules/library/ directory. Agents reference these rules when:
- Planning implementations
- Writing code in specific languages
- Following framework conventions
- Implementing security best practices
- Creating tests according to standards

## Project State
Current project state is tracked in .do/active_state.json

Agents read and update this state automatically during task execution.

## Terminal Execution Policy
Configure Antigravity's terminal execution policy:
- **Off**: Never auto-execute (except Allow list)
- **Auto**: Agent decides when to execute (default)
- **Turbo**: Always auto-execute (except Deny list)

Recommended: **Auto** - provides a good balance of autonomy and safety.

## Review Policy
Configure when agents request review:
- **Always Proceed**: Agent never asks for review
- **Agent Decides**: Agent asks when needed (recommended)
- **Request Review**: Agent always asks before proceeding

Recommended: **Agent Decides** - allows agents to work autonomously while checking in when appropriate.

## Usage Workflow

1. **Capture Ideas**: Use /do to describe what you want to build
2. **Plan**: Agents create a task plan and implementation strategy
3. **Review**: Review the agent's plan (if Review Policy requires it)
4. **Execute**: Agents implement the plan, making code changes
5. **Validate**: Agents run tests and validate their work
6. **Status**: Use /sys status to check progress and generate reports

## Agent-Driven Development
Agents can:
- Browse documentation automatically
- Research best practices from the web
- Generate comprehensive test coverage
- Refactor code while maintaining functionality
- Handle multi-file changes atomically

## Example Prompts
- "Generate unit tests for the OrderService with mock implementations"
- "Refactor the authentication module to use JWT tokens"
- "Add error handling to all API endpoints"
- "Create documentation for the API using OpenAPI spec"

For full command list and detailed documentation, see README.md

Learn more: https://antigravity.google/docs
`
	antigravityDir := filepath.Join(projectPath, ".antigravity")
	if err := utils.CreateDirectory(antigravityDir); err != nil {
		return fmt.Errorf("failed to create .antigravity directory: %w", err)
	}

	path := filepath.Join(antigravityDir, "config.md")
	return utils.WriteFile(path, []byte(content))
}

// generateWindsurfConfig generates .windsurf/config.md file and .windsurfrules file
func generateWindsurfConfig(projectPath string) error {
	// Create .windsurf directory for consistency
	windsurfDir := filepath.Join(projectPath, ".windsurf")
	if err := utils.CreateDirectory(windsurfDir); err != nil {
		return fmt.Errorf("failed to create .windsurf directory: %w", err)
	}

	// Generate config.md in .windsurf/ directory
	configContent := `# Windsurf Project Configuration

## Overview
This project uses DoPlan's hierarchical AI agency structure with Windsurf, an AI-powered IDE with advanced code generation and agent capabilities.

## Windsurf Features

### Command (Inline Code Generation)
- **Invoke**: Press Cmd/Ctrl+I for inline code generation and edits
- **No Premium Credits**: Command does NOT consume premium model credits
- **Edit Mode**: Highlight code before invoking to edit the selection
- **Generate Mode**: No selection generates code at cursor location
- **Accept/Reject**: Use code lens or shortcuts (Cmd/Ctrl+Enter or Cmd/Ctrl+Delete)
- **Terminal Command**: Use Cmd/Ctrl+I in terminal to generate CLI syntax from natural language prompts
- **Best Practices**: 
  - Works great for file-scoped, in-line changes
  - Simple prompts like "Fix this" or "Refactor" work thanks to context awareness
  - Specific prompts with clear objectives work even better

### Cascade (Agent Mode)
- **Autonomous Agents**: Complete complex, multi-step tasks
- **MCP Integration**: Use Model Context Protocol to connect external tools and services
- **Workflows**: Automated development workflows for common tasks
- **Memories**: Persistent context about project decisions and patterns

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .windsurf/agents/

Reference agents in chat: @agent-name (e.g., @frontend_lead, @backend_lead)

## Commands
All commands are defined in .windsurf/commands/. Type any command to activate it:

- /hello - Welcome, tutorial, and command introductions (subcommands: goplan, meeting, plan, build, github)
- /tell - Capture your idea
- /meeting - Adaptive discovery meeting with progress tracking and timing
- /write - Generate documents & content (subcommands: plan, content, change, prd, architecture, design, etc.)
- /plan - Generate execution plan (subcommands: everything, phases, next, phase {no} tasks, phases tasks, all tasks)
- /build - Start coding with auto-completion detection
- /status - Show progress and reports (subcommands: report, full)
- /github - GitHub operations (subcommands: info, issue, milestone, ci, release)
- /state - State management (subcommands: snapshot, list, diff, restore)
- /feedback - Log feedback
- /security - Security review/audit (subcommands: review, audit, both)
- /permissions - Design RBAC system
- /optimize - Project optimization (subcommands: design, finance, performance)

## Model Context Protocol (MCP)

### Adding MCP Plugins
1. **Plugin Store**: Click Plugins icon in Cascade panel or go to Windsurf Settings > Cascade > Plugins
2. **Official Plugins**: Show with blue checkmark (made by parent service company)
3. **Manual Addition**: Edit raw mcp_config.json file if plugin not in store
4. **Transport Types**: Supports stdio and http transports
5. **Refresh**: Press refresh button after adding new MCP plugin

### Configuring MCP Tools
- Cascade has a limit of 100 total tools accessible at any time
- Enable/disable tools at plugin level via Tools tab
- Manage from Windsurf Settings > Cascade > Manage plugins
- Enterprise users must manually enable MCP via settings

### MCP Configuration File
Location: mcp_config.json in project root or user config

Example configuration:
` + "```json\n" + `{
  "mcpServers": {
    "server-name": {
      "command": "command-path",
      "args": ["arg1", "arg2"],
      "env": {
        "ENV_VAR": "value"
      }
    }
  }
}
` + "```\n\n" + `

## Memories & Rules
Stack-specific rules are organized in .windsurf/rules/library/ directory.

Windsurf's Memories & Rules system:
- **Rules**: Project-specific guidelines and conventions stored in .windsurf/rules/
- **Memories**: Persistent context about project decisions, patterns, and history
- **Automatic Loading**: Memories and rules automatically loaded when working on relevant files
- **Language/Framework Rules**: Automatically apply based on file types
- **Context Awareness**: Windsurf remembers decisions and applies them consistently

### Using Memories
- Memories persist across sessions
- Capture important decisions and patterns
- Automatically referenced during code generation
- Can be created manually or through agent interactions

## Workflows
Windsurf supports automated workflows through Cascade:
- **Custom Workflows**: Create reusable workflows for common tasks
- **Workflow Automation**: Automate repetitive development tasks
- **Multi-step Tasks**: Break complex operations into manageable steps
- **Integration**: Workflows can leverage MCP tools and agents

## Project State
Current project state is tracked in .do/active_state.json

Windsurf can read and update this state during agent execution.

## Usage Workflow
1. **Welcome** (New Users): Type /hello for first-time tutorial
2. **Capture Ideas**: Type /tell to capture your idea
2. **Plan**: Cascade agents create a plan using Memories & Rules
3. **Review Plan**: Type /plan to review the execution plan
4. **Code Generation**: Use Command (Cmd/Ctrl+I) for inline edits or Cascade for complex tasks
5. **Agent Execution**: Type /build to start coding with agent assistance
6. **Reference Context**: Agents automatically reference rules and memories
7. **Status**: Type /status to check progress and generate reports

## Context Management
- **@filename**: Reference specific files in chat
- **@agent-name**: Activate agent personalities (e.g., @frontend_lead)
- **Memories**: Windsurf maintains persistent memory across conversations
- **Rules**: Automatically applied based on file types and project structure
- **MCP Tools**: Access external tools and services through MCP plugins

## Best Practices

### Using Command (Cmd/Ctrl+I)
- Keep prompts specific for better results
- Highlight code before invoking to edit existing code
- Use terminal Command for generating CLI commands
- Accept/reject changes immediately to maintain flow

### Using Cascade
- Leverage MCP plugins for extended capabilities
- Use workflows for repetitive tasks
- Let memories accumulate project knowledge over time
- Configure tool limits appropriately (100 tool max)

For full command list and detailed documentation, see README.md

Learn more: 
- Command: https://docs.windsurf.com/command/windsurf-overview
- Cascade MCP: https://docs.windsurf.com/windsurf/cascade/mcp
- Workflows: https://docs.windsurf.com/windsurf/cascade/workflows
- Memories: https://docs.windsurf.com/windsurf/cascade/memories
`
	configPath := filepath.Join(windsurfDir, "config.md")
	if err := utils.WriteFile(configPath, []byte(configContent)); err != nil {
		return fmt.Errorf("failed to write .windsurf/config.md: %w", err)
	}

	// Also generate .windsurfrules file for backward compatibility
	rulesContent := `# DoPlan Project Configuration

## Overview
This project uses DoPlan's hierarchical AI agency structure with Windsurf, an AI-powered IDE.

## Windsurf Features
- **Command (Cmd/Ctrl+I)**: Inline code generation and edits, no premium credits
- **Cascade**: Autonomous agents for complex, multi-step tasks
- **Memories & Rules**: Persistent context and project-specific guidelines
- **Workflows**: Automated development workflows
- **MCP Support**: Model Context Protocol integration for external tools

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .windsurf/agents/

Reference agents in chat: @agent-name (e.g., @frontend_lead, @backend_lead)

## Commands
All commands are defined in .windsurf/commands/. Type any command to activate it:

- /hello - Welcome, tutorial, and command introductions (subcommands: goplan, meeting, plan, build, github)
- /tell - Capture your idea
- /meeting - Adaptive discovery meeting with progress tracking and timing
- /write - Generate documents & content (subcommands: plan, content, change, prd, architecture, design, etc.)
- /plan - Generate execution plan (subcommands: everything, phases, next, phase {no} tasks, phases tasks, all tasks)
- /build - Start coding with auto-completion detection
- /status - Show progress and reports (subcommands: report, full)
- /github - GitHub operations (subcommands: info, issue, milestone, ci, release)
- /state - State management (subcommands: snapshot, list, diff, restore)
- /feedback - Log feedback
- /security - Security review/audit (subcommands: review, audit, both)
- /permissions - Design RBAC system
- /optimize - Project optimization (subcommands: design, finance, performance)

## Quick Reference
- **Command**: Cmd/Ctrl+I for inline edits (highlight code first to edit selection)
- **Terminal Command**: Cmd/Ctrl+I in terminal for CLI syntax generation
- **MCP Plugins**: Add via Plugin Store or edit mcp_config.json
- **Memories**: Persistent context automatically loaded during code generation

## Rules & Memories
Stack-specific rules are organized in .windsurf/rules/library/ directory.

Windsurf's Memories & Rules system:
- **Rules**: Project-specific guidelines and conventions
- **Memories**: Persistent context about project decisions and patterns
- Automatically loaded when working on relevant files
- Language and framework-specific rules apply automatically

## Project State
Current project state is tracked in .do/active_state.json

Windsurf can read and update this state during agent execution.

## Usage Workflow
1. Type /hello for first-time welcome and tutorial (new users)
1. Type /hello for first-time welcome and tutorial (new users)
2. Type /tell to capture your idea
3. Windsurf agents create a plan using Memories & Rules
4. Type /plan to review the execution plan
4. Use Command (Cmd/Ctrl+I) for inline edits or Cascade for complex tasks
5. Agents reference rules and memories automatically
6. Type /status to check progress and generate reports

## Context Management
- Use @filename to reference specific files
- Use @agent-name to activate agent personalities
- Windsurf maintains memory of previous conversations
- Rules are automatically applied based on file types

For full command list and detailed documentation, see README.md

Learn more: https://docs.windsurf.com
`
	rulesPath := filepath.Join(projectPath, ".windsurfrules")
	return utils.WriteFile(rulesPath, []byte(rulesContent))
}

// generateClineConfig generates .cline/config.md file
func generateClineConfig(projectPath string) error {
	content := `# Cline Project Configuration

## Overview
This project uses DoPlan's hierarchical AI agency structure with Cline, an open-source AI coding agent for VS Code.

## Cline Features
Cline is a comprehensive coding agent that:
- Understands entire codebases, not just single files
- Plans complex changes across multiple files
- Executes multi-step tasks autonomously
- Provides transparency and developer control
- Integrates directly into VS Code

Unlike autocomplete tools, Cline is a true coding agent capable of planning, executing, and validating complex engineering tasks.

## Commands & Shortcuts

### Code Commands
Cline provides various code commands for common operations:
- Generate code snippets
- Edit existing code
- Refactor code sections
- Fix linting issues

### Terminal Integration
- Execute terminal commands through Cline
- Run tests and scripts
- Manage dependencies
- Configure terminal pagers if needed

### Git Integration
- Stage and commit changes
- Create branches
- View git status and diff
- Manage git workflows

## Slash Commands

### Workflows
- **Quickstart**: Get started with Cline workflows
- **Best Practices**: Learn workflow optimization techniques
- Create reusable workflow patterns
- Automate repetitive development tasks

### Task Management
- **/new-task**: Create new tasks for Cline to work on
- **Understanding Tasks**: Learn how Cline manages tasks
- **Task Management**: View, update, and complete tasks
- Tasks are structured with clear objectives and steps

### Special Commands
- **/new-rule**: Create project-specific rules
- **/smol**: Use for smaller, focused tasks
- **/report-bug**: Report issues or problems
- **/deep-planning**: Request detailed, comprehensive planning

## Tasks

### Understanding Tasks
Tasks in Cline are structured units of work that:
- Have clear objectives
- Can span multiple files
- Include verification steps
- Track progress and completion

### Task Management
- View all active tasks
- Update task status
- Prioritize tasks
- Link related tasks together

## Model Context Protocol (MCP)

### Overview
Cline supports MCP for extending capabilities with external tools and services.

### Adding MCP Servers
- **From GitHub**: Add MCP servers directly from GitHub repositories
- **MCP Marketplace**: Browse and install pre-built MCP servers
- **Custom Servers**: Develop your own MCP servers using the development protocol

### Configuring MCP Servers
- Configure server endpoints and credentials
- Set up authentication
- Manage server connections
- Enable/disable specific tools

### Transport Mechanisms
- Support for different transport types
- **Remote Servers**: Connect to remote MCP servers
- Secure connections and authentication
- Network configuration options

## Advanced Features

### Hooks
Configure hooks to customize Cline's behavior:
- Pre-execution hooks
- Post-execution hooks
- Task completion hooks
- Custom workflow triggers

### Focus Chain
Use focus chain for:
- Multi-step reasoning
- Complex problem decomposition
- Sequential task execution
- Context-aware decision making

### Customization
- Disable terminal pagers for cleaner output
- Customize command behavior
- Configure output formats
- Adjust agent parameters

## Agent Hierarchy
This project uses a hierarchical AI agency structure. All agents are defined in .cline/agents/

Agents provide specialized expertise:
- Architecture and design patterns
- Language-specific best practices
- Security and performance guidelines
- Testing and quality assurance

Activate agents: "As the backend_lead, review this API design"

## Commands
All commands are defined in .cline/commands/. Use these commands in Cline chat:

- /hello - Welcome, tutorial, and command introductions (subcommands: goplan, meeting, plan, build, github)
- /tell - Capture your idea
- /meeting - Adaptive discovery meeting with progress tracking and timing
- /write - Generate documents & content (subcommands: plan, content, change, prd, architecture, design, etc.)
- /plan - Generate execution plan (subcommands: everything, phases, next, phase {no} tasks, phases tasks, all tasks)
- /build - Start coding with auto-completion detection
- /status - Show progress and reports (subcommands: report, full)
- /github - GitHub operations (subcommands: info, issue, milestone, ci, release)
- /state - State management (subcommands: snapshot, list, diff, restore)
- /feedback - Log feedback
- /security - Security review/audit (subcommands: review, audit, both)
- /permissions - Design RBAC system
- /optimize - Project optimization (subcommands: design, finance, performance)

## Rules
Stack-specific rules are organized in .cline/rules/library/ directory.

Cline references these rules when:
- Planning code changes
- Writing code in specific languages
- Following framework conventions
- Implementing security practices
- Creating tests

Create new rules using /new-rule command. Rules guide Cline's decision-making process and ensure consistency.

## Project State
Current project state is tracked in .do/active_state.json

Cline can read and update this state automatically during task execution.

## Context Management
Cline maintains context across:
- Entire codebase structure
- Previous conversation history
- File relationships and dependencies
- Current development state
- Active tasks and workflows

Use @filename to reference specific files in your prompts.

## Usage Workflow

### Quick Task Creation
1. Use /new-task to create a task
2. Describe what you want to accomplish
3. Cline will break it down into steps
4. Monitor progress as Cline executes

### Deep Planning
1. Use /deep-planning for complex features
2. Review the comprehensive plan
3. Approve or modify as needed
4. Let Cline execute step by step

### Workflows
1. Create reusable workflows for common patterns
2. Follow best practices for workflow design
3. Automate repetitive tasks
4. Iterate and improve workflows

### Execution
Cline will:
1. Make changes across multiple files
2. Update related components
3. Add necessary imports and dependencies
4. Maintain code consistency
5. Run tests and verify functionality

## Transparency & Control
Cline provides:
- Clear visibility into planned changes
- Step-by-step execution logs
- Ability to review before applying changes
- Full control over what gets executed
- Task progress tracking

## Example Prompts
- "Refactor the authentication module to support multiple providers"
- "Add comprehensive error handling to all API endpoints"
- "Create integration tests for the payment processing flow"
- "Optimize database queries in the user service"

## Best Practices

### Task Management
- Break large tasks into smaller, manageable pieces
- Use /smol for focused, quick tasks
- Use /deep-planning for complex features
- Review tasks regularly and update priorities

### Rules
- Create project-specific rules with /new-rule
- Keep rules focused and clear
- Update rules as project evolves
- Share rules across team members

### MCP Integration
- Start with marketplace servers
- Gradually add custom integrations
- Monitor server performance
- Keep credentials secure

For full command list and detailed documentation, see README.md

Learn more: https://docs.cline.bot
`
	clineDir := filepath.Join(projectPath, ".cline")
	if err := utils.CreateDirectory(clineDir); err != nil {
		return fmt.Errorf("failed to create .cline directory: %w", err)
	}

	path := filepath.Join(clineDir, "config.md")
	return utils.WriteFile(path, []byte(content))
}

// generateOpenCodeConfig generates opencode.json configuration file
func generateOpenCodeConfig(projectPath string) error {
	content := `{
  "$schema": "https://opencode.ai/config.json",
  
  // DoPlan Project Configuration
  // This project uses DoPlan's hierarchical AI agency structure with OpenCode
  
  // Instructions (Rules)
  // OpenCode will load these instruction files to guide AI behavior
  "instructions": [
    ".opencode/rules/library/**/*.md",
    "docs/README.md",
    ".do/**/*.md"
  ],
  
  // Permissions Configuration
  // Configure which tools require explicit approval
  "permission": {
    "edit": "ask",
    "bash": "ask",
    "write": "ask"
  },
  
  // Tools Configuration
  // Enable/disable specific tools available to the AI
  "tools": {
    "write": true,
    "edit": true,
    "bash": true,
    "read": true
  },
  
  // Agent Configuration
  // Custom agents are defined in .opencode/agents/ directory
  // Agents provide specialized knowledge for different development areas
  "agent": {
    "frontend_lead": {
      "description": "Frontend development specialist focused on UI/UX best practices",
      "tools": {
        "write": true,
        "edit": true
      }
    },
    "backend_lead": {
      "description": "Backend development specialist focused on API design and architecture",
      "tools": {
        "write": true,
        "edit": true
      }
    },
    "code-reviewer": {
      "description": "Code reviewer focused on security, performance, and maintainability",
      "tools": {
        "write": false,
        "edit": false,
        "read": true
      }
    }
  },
  
  // Commands Configuration
  // Custom commands are defined in .opencode/commands/ directory
  // These provide structured workflows for common tasks
  "command": {
    "tell": {
      "description": "Capture your project idea or requirement",
      "template": "Capture and document the following idea or requirement: $ARGUMENTS"
    },
    "improve": {
      "description": "Brainstorm improvements with the AI team",
      "template": "Brainstorm improvements for: $ARGUMENTS"
    },
    "write": {
      "description": "Generate documentation",
      "template": "Generate documentation for: $ARGUMENTS"
    },
    "plan": {
      "description": "Generate execution plan",
      "template": "Create an execution plan for: $ARGUMENTS"
    },
    "build": {
      "description": "Start coding implementation",
      "template": "Implement: $ARGUMENTS"
    }
  },
  
  // Sharing Configuration
  // Options: "manual" (default), "auto", "disabled"
  "share": "manual",
  
  // Autoupdate Configuration
  // true: auto-download updates, false: disable, "notify": notify only
  "autoupdate": true,
  
  // MCP Servers Configuration
  // Configure Model Context Protocol servers for extended capabilities
  // See: https://opencode.ai/docs/mcp-servers/
  "mcp": {},
  
  // Formatters Configuration
  // Configure code formatters for different file types
  "formatter": {
    "prettier": {
      "disabled": false
    }
  },
  
  // Instructions (Rules)
  // These paths point to instruction files that guide AI behavior
  // Supports glob patterns and file paths
  // Rules are automatically loaded from .opencode/rules/ directory
  // See: https://opencode.ai/docs/rules/
  
  // Tools
  // Available tools: write, edit, bash, read
  // Configure which tools are available to the AI
  // See: https://opencode.ai/docs/tools/
  
  // Agents
  // Custom agents can be defined in JSON or as markdown files in:
  // - ~/.config/opencode/agent/
  // - .opencode/agent/
  // See: https://opencode.ai/docs/agents/
  
  // Commands
  // Custom commands can be defined in JSON or as markdown files in:
  // - ~/.config/opencode/command/
  // - .opencode/command/
  // See: https://opencode.ai/docs/commands/
  
  // Permissions
  // Configure which operations require explicit approval
  // Options: "allow" (default), "ask", "deny"
  // See: https://opencode.ai/docs/permissions/
}
`
	path := filepath.Join(projectPath, "opencode.json")
	return utils.WriteFile(path, []byte(content))
}

// GenerateIDEConfigs is a convenience function that creates an IDEGenerator and generates IDE configs
func GenerateIDEConfigs(request *models.ProjectRequest, projectPath string) error {
	generator := &IDEGenerator{}
	return generator.Generate(request, projectPath)
}
