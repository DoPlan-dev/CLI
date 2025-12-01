# The Complete DoPlan CLI Guide

This guide provides comprehensive documentation on using DoPlan CLI commands and creating projects with the AI agency system.

---

## 📚 Table of Contents

1. [DoPlan Commands Reference](#doplan-commands-reference)
2. [Complete Workflow](#complete-workflow)
3. [Command Usage Examples](#command-usage-examples)
4. [Best Practices](#best-practices)
5. [Troubleshooting](#troubleshooting)

---

## 🎯 DoPlan Commands Reference

### Core Commands

#### `/hey` - Welcome & Tutorial

**Purpose:** Interactive welcome experience and tutorial for new users

**Usage:**
```
/hey
```

**What it does:**
- Provides first-time welcome and tutorial
- Introduces the DoPlan system
- Explains the agent hierarchy
- Guides you through the workflow
- Sets up your development support mode

---

#### `/do` - Capture Idea & Discovery Meeting

**Purpose:** Captures your project idea and conducts a discovery meeting with the AI team

**Usage:**
```
/do
```

**What it does:**
- Captures your project idea
- Conducts adaptive discovery meeting
- Generates BRAINSTORM.md with refined ideas
- Adapts questions based on your experience level
- Supports multiple meeting speeds (Quick, Standard, Comprehensive, Deep Dive)

**Meeting Speeds:**
- 🚀 **Quick Start** (5-10 min) - Simple projects or when in a hurry
- ⚡ **Standard** (15-20 min) - Balanced depth for most projects
- 📋 **Comprehensive** (30-45 min) - Detailed planning for complex projects
- 🔍 **Deep Dive** (60+ min) - Complete exploration for enterprise solutions

---

#### `/plan` - Generate Planning Documents & Tasks

**Purpose:** Generates comprehensive planning documents and implementation tasks

**Usage:**
```
/plan
```

**What it generates:**
- **PRD** (Product Requirements Document) - `.do/00_System/PRD.md`
  - User stories, acceptance criteria, feature specifications
- **Architecture** - `.do/00_System/ARCHITECTURE.md`
  - System design, database schema, API structure, tech stack decisions
- **Design System** - `.do/00_System/DESIGN_SYSTEM.md`
  - UI/UX guidelines, component library, color scheme, typography
- **TASKS.md** - Detailed implementation tasks broken down by phases

**Task Organization:**
- **Phase 1: Foundation** - Core infrastructure
- **Phase 2: Core Features** - Main functionality
- **Phase 3: Enhancement** - Polish and optimization

Each task includes:
- Clear description
- Acceptance criteria
- Estimated complexity
- Dependencies on other tasks
- Implementation hints

---

#### `/dev` - Start Development

**Purpose:** Begins implementation of tasks with AI agent guidance

**Usage:**
```
/dev                    # Start first/next task
/dev <task-number>      # Start specific task
/dev <specific request> # Get help with implementation
```

**What it does:**
- Guides you through each step
- Generates code snippets
- Helps with implementation decisions
- Reviews your code
- Suggests best practices
- Helps debug issues
- Automatically detects task completion
- Creates Git branches automatically
- Manages project state

**Examples:**
```
/dev
```

```
/dev 3
```

```
/dev I need help setting up the Express server with TypeScript. Show me the complete server.ts file structure
```

```
/dev Generate the QR service using the qrcode library with support for PNG and SVG formats
```

```
/dev Review the qrController.ts file for best practices and error handling
```

**Auto-Completion Detection:**
- When `/dev` detects a task is complete, it will:
  - Show a summary of what was accomplished
  - Ask if you want to mark it as done
  - Auto-commit and push changes
  - Move to the next task

---

#### `/sys` - System Management

**Purpose:** System control panel for project management and operations

**Usage:**
```
/sys                    # Show system status
/sys status             # Detailed project status
/sys performance        # Performance metrics
/sys backup             # Backup project
/sys restore            # Restore from backup
/sys memory             # Memory card management
/sys state              # State management
/sys feedback           # Feedback system
/sys github             # GitHub integration
/sys security           # Security settings
/sys permissions         # File permissions
/sys access             # Access control
```

**What it does:**
- **Status**: Shows current project state, progress, and metrics
- **Performance**: Displays performance metrics and cache statistics
- **Backup**: Creates full project backup
- **Restore**: Restores project from backup
- **Memory**: Manages memory card (personalization data)
- **State**: Manages project state snapshots
- **Feedback**: Logs feedback and suggestions
- **GitHub**: Syncs KPIs, creates issues, manages milestones
- **Security**: Security audit and settings
- **Permissions**: Manages file permissions
- **Access**: Controls access to project files

---

## 🔄 Complete Workflow

### Phase 1: Onboarding (First Time Only)

```
1. /hey                  → Welcome and tutorial
```

### Phase 2: Planning (15-30 minutes)

```
2. /do                   → Capture idea and discovery meeting
3. /plan                 → Generate planning documents and tasks
```

### Phase 3: Development (Iterative)

```
4. /dev                  → Start coding (auto-detects completion)
5. Repeat /dev            → Continue through remaining tasks
```

### Phase 4: System Management (As Needed)

```
6. /sys status           → Check project status
7. /sys performance      → View performance metrics
8. /sys backup           → Backup project
```

---

## 💡 Command Usage Examples

### Example 1: Starting a New Project

```
/hey                     # First-time tutorial
/do                      # Capture your idea and conduct meeting
/plan                    # Generate all planning documents
/dev                     # Start coding first task
```

### Example 2: Getting Implementation Help

```
/dev I'm getting an error: "Cannot find module 'qrcode'". Help me fix the import and ensure the package is installed correctly.
```

### Example 3: Adding Features

```
/dev Add rate limiting to the /api/qr endpoint: max 100 requests per hour per IP address. Use express-rate-limit package.
```

### Example 4: Code Review

```
/dev Review my qrController.ts file. Check for:
- Error handling best practices
- Input validation completeness
- Response format consistency
- Security considerations
```

### Example 5: Refactoring

```
/dev Refactor the analytics service to use async/await instead of callbacks. Ensure database operations are properly handled.
```

### Example 6: Testing

```
/dev Create comprehensive tests for the QR generation service:
- Test PNG generation
- Test SVG generation
- Test different error correction levels
- Test invalid inputs
- Test error handling
```

### Example 7: Deployment

```
/dev Help me prepare for deployment to Vercel:
- Create vercel.json configuration
- Set up environment variables
- Configure build settings
- Add deployment scripts
```

### Example 8: System Management

```
/sys status              # Check current project state
/sys performance         # View performance metrics
/sys backup              # Create backup
/sys github info         # Sync GitHub KPIs
```

---

## 🎓 Best Practices

### 1. Start with `/hey`
For new users, always start with `/hey` to get the tutorial and understand the system.

### 2. Use `/do` for Discovery
Use `/do` to capture your idea and let the AI team conduct a thorough discovery meeting. Choose the meeting speed that matches your project complexity.

### 3. Plan Before Coding
Always run `/plan` after `/do` to generate comprehensive planning documents and task breakdowns.

### 4. Iterate with `/dev`
Use `/dev` to start coding. The system will automatically detect when tasks are complete and guide you to the next one.

### 5. Monitor with `/sys`
Regularly check `/sys status` to track progress and `/sys performance` to monitor system health.

### 6. Backup Regularly
Use `/sys backup` before major changes to ensure you can restore if needed.

---

## 🔧 Troubleshooting

### Agent Not Responding or Giving Generic Answers?

Make sure you've run `/do` first to provide context about your project. The discovery meeting sets up all the necessary context for the AI agents.

### Stuck on a Task?

Use `/dev` with a specific request:
```
/dev I'm stuck on task 3.2. The authentication service isn't working. Help me debug the issue.
```

### Need to Check Project State?

```
/sys status              # See current state and progress
```

### Version or State Issues?

```
/sys state               # Manage state snapshots
/sys restore             # Restore from backup if needed
```

### Code Not Working? Need Debugging Help?

```
/dev Debug this issue: The QR code endpoint returns 500 error. The error message is "TypeError: Cannot read property 'toString' of undefined". Review the qrService.ts file and fix the issue.
```

---

## 📋 Quick Reference Card

### Core Commands
- `/hey` - Welcome and tutorial
- `/do` - Capture idea and discovery meeting
- `/plan` - Generate planning documents and tasks
- `/dev` - Start development (auto-detects completion)
- `/sys` - System management

### Common `/sys` Subcommands
- `/sys status` - Project status
- `/sys performance` - Performance metrics
- `/sys backup` - Backup project
- `/sys restore` - Restore from backup
- `/sys github info` - Sync GitHub KPIs

---

## 🚀 Getting Started

1. **Create Project:**
   ```bash
   npx @doplan-dev/cli
   ```

2. **Open in Cursor:**
   ```bash
   cd your-project-name
   cursor .
   ```

3. **Start Workflow:**
   - Type `/hey` for first-time tutorial
   - Type `/do` to capture your idea
   - Type `/plan` to generate planning documents
   - Type `/dev` to start coding

---

## 📚 Additional Resources

- **Project Structure:** See generated project structure for examples
- **Agent Definitions:** Check `.cursor/agents/` for all 18 agent personas
- **Command Definitions:** See `.cursor/commands/` for command details
- **Rules Library:** Explore `.cursor/rules/library/` for tech stack rules

---

**Happy Building! 🎉**

*Generated by DoPlan CLI - Your AI Project Director*
