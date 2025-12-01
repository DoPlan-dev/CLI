# Antigravity Project Configuration

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
