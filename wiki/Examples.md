# Examples Documentation

Real-world examples and use cases for DoPlan CLI.

---

## Example Projects

### Example 1: Todo App

**Description**: Simple todo application with authentication and dark mode.

**Workflow**:
```bash
# 1. Generate project
npx @doplan-dev/cli
# Enter: todo-app
# Select: Cursor

# 2. Open project
cd todo-app
code .

# 3. Capture idea
/tell Build a todo app with authentication, dark mode, and offline support

# 4. Brainstorm
/improve

# 5. Generate plans
/write

# 6. Approve
/good

# 7. Generate tasks
/tasks

# 8. Build
/build
```

**Features**:
- Add, edit, delete todos
- Mark todos as complete
- Filter todos (all, active, completed)
- User authentication
- Dark mode
- Offline support

---

### Example 2: Blog Platform

**Description**: Full-featured blog platform with CMS.

**Workflow**:
```bash
# 1. Generate project
npx @doplan-dev/cli
# Enter: blog-platform
# Select: Claude Code

# 2. Capture idea
/tell Build a blog platform with:
- User authentication
- Post creation and editing
- Comments system
- Tag system
- Search functionality
- Admin dashboard

# 3. Follow workflow
/improve
/write
/good
/tasks
/build
```

**Features**:
- User authentication
- Post management
- Comments
- Tags
- Search
- Admin dashboard

---

### Example 3: E-commerce Platform

**Description**: E-commerce platform with payment integration.

**Workflow**:
```bash
# 1. Generate project
npx @doplan-dev/cli
# Enter: ecommerce-platform
# Select: Windsurf

# 2. Capture idea
/tell Build an e-commerce platform with:
- Product catalog
- Shopping cart
- Payment integration (Stripe)
- Order management
- User accounts
- Admin panel

# 3. Follow workflow
/improve
/write
/good
/tasks
/build
```

**Features**:
- Product catalog
- Shopping cart
- Payment processing
- Order management
- User accounts
- Admin panel

---

## Use Case Scenarios

### Scenario 1: Solo Developer Starting a Side Project

**Situation**: Developer wants to start a side project quickly.

**Solution**:
1. Run `npx @doplan-dev/cli`
2. Enter project name
3. Select IDE
4. Use `/tell` to capture idea
5. Follow workflow to generate project
6. Start building immediately

**Benefits**:
- No setup time
- Complete project structure
- AI agents ready to help
- Best practices included

---

### Scenario 2: Small Team Standardizing Workflow

**Situation**: Team wants consistent project structure.

**Solution**:
1. Use DoPlan CLI for all new projects
2. Customize agents and rules for team needs
3. Share customizations across team
4. Maintain consistency

**Benefits**:
- Consistent structure
- Standardized workflow
- Shared best practices
- Faster onboarding

---

### Scenario 3: Rapid Prototyping

**Situation**: Need to prototype quickly.

**Solution**:
1. Generate project with DoPlan CLI
2. Use `/tell` and `/write` for quick planning
3. Approve quickly with `/good`
4. Generate tasks and start building
5. Iterate rapidly

**Benefits**:
- Fast setup
- Quick planning
- Immediate development
- Rapid iteration

---

## Real-World Examples

### Example: SaaS Application

**Project**: Task management SaaS

**Tech Stack**: Next.js, PostgreSQL, Stripe

**Workflow**:
```bash
/tell Build a task management SaaS with:
- Team workspaces
- Task boards
- Time tracking
- Billing integration
- Real-time collaboration

/improve
/write
# Review and customize
/change prd Add mobile app support
/change architecture Use microservices
/good
/tasks
/build
```

**Result**: Complete SaaS application structure with all features planned.

---

### Example: API Service

**Project**: REST API service

**Tech Stack**: Go, PostgreSQL, Docker

**Workflow**:
```bash
/tell Build a REST API service for data processing with:
- Authentication
- Rate limiting
- Data validation
- Error handling
- API documentation

/improve
/write
/good
/tasks
/build
```

**Result**: Production-ready API structure.

---

## Example Workflows

### Workflow 1: Quick Start

```bash
# Minimal workflow for quick projects
npx @doplan-dev/cli
/tell [idea]
/write
/good
/tasks
/build
```

---

### Workflow 2: Thorough Planning

```bash
# Complete workflow for complex projects
npx @doplan-dev/cli
/tell [detailed idea]
/improve
/write
/change prd [refinements]
/change architecture [updates]
/change design [modifications]
/good
/tasks
/build
/progress
/finished
```

---

### Workflow 3: Iterative Development

```bash
# Iterative workflow
/tell [initial idea]
/write
/good
/tasks
/build
# ... develop ...
/finished
# ... iterate ...
/change prd [new features]
/good
/tasks
/build
```

---

## Example Customizations

### Customization 1: Custom Agent

**Goal**: Add domain-specific agent

**Steps**:
1. Create `.cursor/agents/domain_expert.md`
2. Define agent persona
3. Reference in commands if needed

**Example**:
```markdown
# Domain Expert

## Role
Domain-specific expert

## System Prompt
You are a domain expert specializing in [domain].
```

---

### Customization 2: Custom Command

**Goal**: Add project-specific command

**Steps**:
1. Create `.cursor/commands/custom.md`
2. Define trigger and action
3. Specify agent involvement

**Example**:
```markdown
# Custom Command

## Trigger
/custom-action

## Action
Performs custom action

## Agent Involvement
- Domain Expert
```

---

### Customization 3: Custom Rules

**Goal**: Add project-specific rules

**Steps**:
1. Create rule file in appropriate category
2. Define guidelines
3. Load with `/load`

**Example**:
```markdown
# Project-Specific Rule

## Overview
Project-specific guidelines

## Guidelines
- Guideline 1
- Guideline 2
```

---

## Before/After Comparisons

### Before DoPlan CLI

**Time to First Commit**: 2-4 hours
- Project setup: 30 minutes
- Structure creation: 30 minutes
- Configuration: 1 hour
- CI/CD setup: 1 hour
- Boilerplate: 30 minutes

**Issues**:
- Inconsistent structure
- Missing best practices
- Manual configuration
- No AI assistance

---

### After DoPlan CLI

**Time to First Commit**: 5 minutes
- Project generation: 5 seconds
- Planning: 2 minutes
- First commit: 3 minutes

**Benefits**:
- Consistent structure
- Best practices included
- Automated configuration
- AI agents ready

---

## Code Snippets

### Example: Using `/tell`

```
/tell Build a social media platform with:
- User profiles
- Posts and comments
- Real-time notifications
- Image uploads
- Search functionality
```

---

### Example: Using `/change`

```
/change prd Add video upload support
/change architecture Use CDN for media storage
/change design Add video player component
```

---

### Example: Using `/load`

```
/load @library/04-frameworks/nextjs.md
/load @library/07-databases/postgresql.md
/load @library/12-security/authentication.md
```

---

## Example Repositories

### Community Examples

- [Todo App Example](https://github.com/DoPlan-dev/examples-todo)
- [Blog Platform Example](https://github.com/DoPlan-dev/examples-blog)
- [E-commerce Example](https://github.com/DoPlan-dev/examples-ecommerce)

**Note**: Example repositories will be available soon.

---

## Related Pages

- [Quick Start](Quick-Start) - Getting started
- [First Project Tutorial](First-Project-Tutorial) - Step-by-step tutorial
- [Workflow Guide](Workflow) - Complete workflow
- [Customization Guide](Customization) - Customization examples
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

