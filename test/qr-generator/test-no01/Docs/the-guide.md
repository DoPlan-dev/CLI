# The Complete DoPlan CLI Guide

This guide provides comprehensive documentation on using DoPlan CLI commands and creating projects with the AI agency system.

---

## 📚 Table of Contents

1. [DoPlan Commands Reference](#doplan-commands-reference)
2. [Complete Workflow](#complete-workflow)
3. [QR Code Generator Project Prompt](#qr-code-generator-project-prompt)
4. [Command Usage Examples](#command-usage-examples)
5. [Best Practices](#best-practices)
6. [Troubleshooting](#troubleshooting)

---

## 🎯 DoPlan Commands Reference

### Planning Phase Commands

#### `/tell` - Capture Your Project Idea

**Purpose:** Captures your project idea and stores it in `.plan/00_System/IDEA.md`

**Usage:**
```
/tell <your detailed project description>
```

**What it does:**
- Creates or updates `IDEA.md` with your project description
- Provides context for all AI agents
- Sets the foundation for planning documents

**Example:**
```
/tell I want to build a task management app with React and Node.js
```

---

#### `/improve` - Brainstorm with AI Team

**Purpose:** Activates Product Manager, Engineering Lead, and Design Manager for brainstorming

**Usage:**
```
/improve
```

**Or with specific focus:**
```
/improve Focus on API design patterns and error handling strategies
```

**What it does:**
- Refines the feature set
- Suggests technical approaches
- Identifies potential challenges
- Recommends best practices
- Proposes alternative solutions
- Estimates complexity

**Output:** Results saved to `.plan/00_System/BRAINSTORM.md`

---

#### `/write` - Generate Planning Documents

**Purpose:** Generates comprehensive planning documents from your idea

**Usage:**
```
/write
```

**What it generates:**
- **PRD** (Product Requirements Document) - `.plan/00_System/PRD.md`
  - User stories, acceptance criteria, feature specifications
- **Architecture** - `.plan/00_System/ARCHITECTURE.md`
  - System design, database schema, API structure, tech stack decisions
- **Design System** - `.plan/00_System/DESIGN_SYSTEM.md`
  - UI/UX guidelines, component library, color scheme, typography

---

#### `/change` - Edit Planning Documents

**Purpose:** Update specific planning documents without regenerating everything

**Usage:**
```
/change prd <your update>
/change architecture <your update>
/change design_system <your update>
```

**Examples:**
```
/change prd Add requirement for API rate limiting: 100 requests per hour per IP address
```

```
/change architecture Use PostgreSQL instead of SQLite for better scalability, add connection pooling configuration
```

```
/change design_system Use Tailwind CSS for styling, color scheme: blue primary (#3B82F6), dark mode support
```

---

#### `/good` - Approve & Lock the Plan

**Purpose:** Locks the current plan and enables task generation

**Usage:**
```
/good
```

**What it does:**
- Locks the current plan (PRD, Architecture, Design System)
- Enables task generation
- Updates `.plan/active_state.json` to mark planning as complete
- Prepares the project for implementation phase

**Note:** After `/good`, you can still use `/change` commands, but the plan is considered "approved" for task generation.

---

### Implementation Phase Commands

#### `/plan` - Generate Implementation Tasks

**Purpose:** Creates detailed, actionable implementation tasks from the approved plan

**Usage:**
```
/plan
```

**What it generates:**
Creates `.plan/TASKS.md` with tasks broken down by:

- **Setup & Configuration**
  - Project initialization
  - Dependencies installation
  - TypeScript configuration
  - Environment setup

- **Core API Development**
  - Express server setup
  - Service layer implementation
  - API routes and controllers
  - Database setup and models

- **Frontend Development**
  - HTML/CSS/JS interface
  - Form handling
  - Data display
  - User interactions

- **Testing**
  - Unit tests
  - Integration tests
  - API endpoint tests

- **Deployment**
  - Build configuration
  - Deployment scripts
  - Environment variables

**Each task includes:**
- Clear description
- Acceptance criteria
- Estimated complexity
- Dependencies on other tasks
- Implementation hints

---

#### `/build` - Start Coding

**Purpose:** Begins implementation of tasks with AI agent guidance

**Usage:**
```
/build                    # Start first/next task
/build <task-number>      # Start specific task
/build <specific request> # Get help with implementation
```

**What it does:**
- Guides you through each step
- Generates code snippets
- Helps with implementation decisions
- Reviews your code
- Suggests best practices
- Helps debug issues

**Examples:**
```
/build
```

```
/build 3
```

```
/build I need help setting up the Express server with TypeScript. Show me the complete server.ts file structure
```

```
/build Generate the QR service using the qrcode library with support for PNG and SVG formats
```

```
/build Review the qrController.ts file for best practices and error handling
```

---

#### `/progress` - Check Progress

**Purpose:** Shows current project progress and task status

**Usage:**
```
/progress
```

**What it shows:**
- ✅ Completed tasks (with completion dates)
- 🔄 In-progress tasks (currently working on)
- ⏳ Remaining tasks (not started)
- 📊 Overall completion percentage
- 📈 Progress visualization

---

#### `/finished` - Mark Task Complete

**Purpose:** Marks a task as complete and updates project state

**Usage:**
```
/finished           # Mark current task complete
/finished <number>  # Mark specific task complete
```

**What it does:**
- Marks task as complete in `.plan/TASKS.md`
- Updates `.plan/active_state.json`
- Moves to the next task automatically
- Records completion timestamp

**Workflow:**
```
/finished
/build
```

This workflow: finish → build next, keeps you moving forward efficiently.

---

### Team & Context Commands

#### `/team` - Show Active Agents

**Purpose:** Displays which AI agents are currently active and their roles

**Usage:**
```
/team
```

**Shows:**
- List of all 18 AI agents
- Their current status (active/inactive)
- Their roles and responsibilities
- Hierarchy structure

---

#### `/load` - Inject Context

**Purpose:** Provides additional context to AI agents when they need more information

**Usage:**
```
/load <context information>
```

**Examples:**
```
/load Add context: We're using Vercel for deployment, need to configure environment variables for production
```

```
/load Context: We're building a QR Code Generator API. Current task is implementing the QR generation service. We're using Node.js, Express, TypeScript, and the qrcode npm package.
```

---

### Quality & Release Commands

#### `/ship` - Prepare for Release

**Purpose:** Prepares the project for release

**Usage:**
```
/ship
```

**What it does:**
- Generates changelog
- Updates version numbers
- Creates release notes
- Prepares deployment checklist
- Reviews release readiness

---

#### `/safe` - Security Audit

**Purpose:** Runs comprehensive security audit

**Usage:**
```
/safe
```

**What it checks:**
- Dependency vulnerabilities
- Code security issues
- API security best practices
- Authentication/authorization review
- Data encryption
- Input validation

---

#### `/cheap` - Cost Optimization

**Purpose:** Reviews and optimizes project costs

**Usage:**
```
/cheap
```

**What it analyzes:**
- Hosting costs analysis
- Database optimization opportunities
- API rate limiting strategies
- Resource usage optimization
- Cost-effective alternatives

---

## 🔄 Complete Workflow

### Phase 1: Planning (15-20 minutes)

```
1. /tell <your idea>           → Capture project idea
2. /improve                     → Brainstorm with AI team
3. /write                       → Generate PRD, Architecture, Design System
4. /change prd <update>         → Refine requirements (optional)
5. /change architecture <update> → Adjust technical design (optional)
6. /good                        → Approve plan and lock it
```

### Phase 2: Implementation (2-6 hours)

```
7. /plan                       → Generate implementation tasks
8. /build                       → Start first task (or /build <task-number>)
9. /progress                    → Check completion status
10. /finished                   → Mark task complete
11. /build                      → Continue to next task
12. Repeat steps 9-11           → Complete all tasks
```

### Phase 3: Quality & Release (1-2 hours)

```
13. /safe                       → Security audit
14. /cheap                      → Cost optimization review
15. /ship                       → Prepare for release
```

### Helper Commands (use anytime)

```
/team                           → See active AI agents
/load <context>                 → Inject context into agents
```

---

## 🎯 QR Code Generator Project Prompt

### Complete `/tell` Prompt for QR Code Generator API

Copy and paste this complete prompt when using `/tell` to create your QR Code Generator project:

```
/tell I want to build a QR Code Generator API micro SaaS. Here are the requirements:

**Core Features:**
- REST API endpoint that accepts text or URLs via POST request
- Generate QR codes in PNG and SVG formats
- Support customizable size (default 200x200 pixels)
- Support error correction levels (L, M, Q, H)
- Return base64 encoded image or direct file download
- Simple analytics tracking (generation count, timestamp)

**Technical Stack:**
- Backend: Node.js with Express framework
- Language: TypeScript for type safety
- Database: SQLite for MVP (easy to upgrade to PostgreSQL later)
- QR Library: qrcode npm package
- Image Processing: sharp for image manipulation if needed

**API Endpoints:**
1. POST /api/qr - Generate QR code
   - Request body: { text: string, size?: number, format?: 'png'|'svg', errorCorrection?: 'L'|'M'|'Q'|'H' }
   - Response: { qrCode: string (base64), format: string, size: number } or file download

2. GET /api/analytics - Get usage statistics
   - Response: { totalGenerations: number, recentActivity: array }

**Frontend:**
- Simple HTML/CSS/JavaScript interface
- Form to input text/URL
- Preview generated QR code
- Download button for PNG/SVG
- Display analytics

**MVP Scope:**
- Focus on core QR generation first
- Basic analytics (count only)
- Simple web interface
- No authentication required for MVP

**Future Enhancements (v2.0+):**
- Custom colors and styling
- Logo embedding in center
- Batch generation
- API keys for rate limiting
- Advanced analytics dashboard
- Custom domains
```

---

## 💡 Command Usage Examples

### Example 1: Debugging Help

```
/build I'm getting an error: "Cannot find module 'qrcode'". Help me fix the import and ensure the package is installed correctly.
```

### Example 2: Adding Features

```
/build Add rate limiting to the /api/qr endpoint: max 100 requests per hour per IP address. Use express-rate-limit package.
```

### Example 3: Code Review

```
/build Review my qrController.ts file. Check for:
- Error handling best practices
- Input validation completeness
- Response format consistency
- Security considerations
```

### Example 4: Refactoring

```
/build Refactor the analytics service to use async/await instead of callbacks. Ensure database operations are properly handled.
```

### Example 5: Testing

```
/build Create comprehensive tests for the QR generation service:
- Test PNG generation
- Test SVG generation
- Test different error correction levels
- Test invalid inputs
- Test error handling
```

### Example 6: Deployment

```
/build Help me prepare for deployment to Vercel:
- Create vercel.json configuration
- Set up environment variables
- Configure build settings
- Add deployment scripts
```

### Example 7: Focused Brainstorming

```
/improve Focus on API design patterns, error handling strategies, and database schema for analytics tracking
```

### Example 8: Architecture Changes

```
/change architecture We decided to use PostgreSQL instead of SQLite. Update the database schema and connection setup. Add migration scripts.
```

---

## 🎓 Best Practices

### 1. Start Simple
Focus on MVP first. Use:
```
/build Start with the simplest QR generation endpoint first
```

### 2. Iterate Frequently
Use `/change` commands to refine as you go:
```
/change architecture Add Redis caching for frequently generated QR codes
```

### 3. Test Early
Write tests alongside code:
```
/build After creating qrService.ts, immediately create test file with basic tests
```

### 4. Leverage AI Agents
Use `/team` to see which agents can help:
```
/team
```

Provide context when needed:
```
/load Context: We're building a QR Code Generator API. Current task is implementing the QR generation service. We're using Node.js, Express, TypeScript, and the qrcode npm package. The service needs to support PNG and SVG formats with customizable size and error correction levels.
```

### 5. Document Progress
Keep notes in `STANDUP.md`:
- Update daily with progress
- Note blockers
- Track decisions

### 6. Regular Progress Checks
```
/progress
```

Check progress regularly to stay on track.

---

## 🔧 Troubleshooting

### Agent Not Responding or Giving Generic Answers?

```
/load Context: We're building a QR Code Generator API. Current task is implementing the QR generation service. We're using Node.js, Express, TypeScript, and the qrcode npm package. The service needs to support PNG and SVG formats with customizable size and error correction levels.
```

### Stuck on a Task? Need Alternative Approaches?

```
/improve I'm having trouble with the QR code generation. The qrcode library is returning a buffer but I need base64. Suggest 3 different approaches to convert the buffer to base64 string.
```

### Need to Pivot or Change Direction?

```
/change architecture We decided to use PostgreSQL instead of SQLite. Update the database schema and connection setup. Add migration scripts.
```

### Version or State Issues?

```
/load Check .plan/active_state.json and verify the current project state. What tasks are marked as complete?
```

### Code Not Working? Need Debugging Help?

```
/build Debug this issue: The QR code endpoint returns 500 error. The error message is "TypeError: Cannot read property 'toString' of undefined". Review the qrService.ts file and fix the issue.
```

### Need to Understand Existing Code?

```
/load Explain how the current analytics service works. Show me the database schema and how data flows from API endpoint to database storage.
```

---

## 📋 Quick Reference Card

### Planning Commands
- `/tell <idea>` - Capture project idea
- `/improve` - Brainstorm
- `/write` - Generate planning docs
- `/change <doc> <update>` - Edit documents
- `/good` - Approve plan

### Implementation Commands
- `/plan` - Generate tasks
- `/build` - Start coding
- `/progress` - Check status
- `/finished` - Mark complete

### Helper Commands
- `/team` - Show agents
- `/load <context>` - Add context

### Quality Commands
- `/safe` - Security audit
- `/cheap` - Cost optimization
- `/ship` - Prepare release

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
   - Use `/tell` with the QR Code Generator prompt above
   - Follow the complete workflow
   - Build your MVP!

---

## 📚 Additional Resources

- **Project Structure:** See `test/qr-generator/test-no01/` for a complete example
- **Full Guide:** See `docs/APP_IDEAS.md` for detailed step-by-step instructions
- **Agent Definitions:** Check `.cursor/agents/` for all 18 agent personas
- **Command Definitions:** See `.cursor/commands/` for command details
- **Rules Library:** Explore `.cursor/rules/library/` for tech stack rules

---

**Happy Building! 🎉**

*Generated by DoPlan CLI - Your AI Project Director*

