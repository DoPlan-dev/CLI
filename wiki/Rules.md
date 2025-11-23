# Rules Library Documentation

Complete guide to the DoPlan CLI rules library - 1000+ embedded best practices for all major tech stacks.

---

## Rules Library Overview

The DoPlan CLI rules library is a comprehensive collection of **1000+ best practices** organized into **15 categories**. These rules guide AI agents to make informed decisions and follow industry standards.

### Key Features

- **1000+ Rules**: Comprehensive coverage of modern development practices
- **15 Categories**: Organized by domain and technology
- **Embedded**: All rules embedded in the DoPlan CLI binary
- **Extracted**: Automatically extracted to `.cursor/rules/library/` when you generate a project
- **Customizable**: Easy to modify and extend
- **Tech Stack Coverage**: Rules for all major languages, frameworks, and tools

---

## Rules Categories

### 01. Core Workflow

**Location**: `.cursor/rules/library/01-core-workflow/`

**Purpose**: Core development workflow rules and practices.

**Contents**:
- GitHub workflow automation
- Continuous GitHub workflows
- Command-agent integration
- Development workflow patterns

**Key Files**:
- `github-workflow-automation.md` - GitHub Actions automation rules
- `README.md` - Category overview

**Use Cases**:
- Setting up CI/CD
- Automating workflows
- Integrating commands with agents

---

### 02. AI Agents

**Location**: `.cursor/rules/library/02-ai-agents/`

**Purpose**: Rules for AI agent interaction and behavior.

**Contents**:
- Agent interaction patterns
- Agent coordination rules
- Agent best practices
- Agent customization guidelines

**Key Files**:
- `README.md` - Agent rules overview

**Use Cases**:
- Customizing agent behavior
- Defining agent interactions
- Setting agent guidelines

---

### 03. Languages

**Location**: `.cursor/rules/library/03-languages/`

**Purpose**: Language-specific best practices and conventions.

**Contents**:
- Go
- TypeScript
- JavaScript
- Python
- And more...

**Key Files**:
- `go.md` - Go best practices
- `typescript.md` - TypeScript best practices
- `javascript.md` - JavaScript best practices
- `python.md` - Python best practices
- `README.md` - Language rules overview

**Use Cases**:
- Language-specific coding standards
- Best practices for each language
- Code style guidelines

---

### 04. Frameworks

**Location**: `.cursor/rules/library/04-frameworks/`

**Purpose**: Framework-specific rules and best practices.

**Contents**:
- Next.js
- React
- Express
- And more...

**Key Files**:
- `nextjs.md` - Next.js best practices
- `react.md` - React best practices
- `express.md` - Express best practices
- `README.md` - Framework rules overview

**Use Cases**:
- Framework-specific patterns
- Framework best practices
- Framework conventions

---

### 05. UI Libraries

**Location**: `.cursor/rules/library/05-ui-libraries/`

**Purpose**: UI library rules and guidelines.

**Contents**:
- UI library best practices
- Component patterns
- Styling guidelines

**Key Files**:
- `README.md` - UI library rules overview

**Use Cases**:
- UI component development
- Styling guidelines
- UI patterns

---

### 06. Cloud Infrastructure

**Location**: `.cursor/rules/library/06-cloud-infrastructure/`

**Purpose**: Cloud platform rules and deployment practices.

**Contents**:
- Cloud platform best practices
- Deployment guidelines
- Infrastructure patterns

**Key Files**:
- `README.md` - Cloud infrastructure rules overview

**Use Cases**:
- Cloud deployment
- Infrastructure setup
- Cloud best practices

---

### 07. Databases

**Location**: `.cursor/rules/library/07-databases/`

**Purpose**: Database rules and best practices.

**Contents**:
- PostgreSQL
- MongoDB
- And more...

**Key Files**:
- `postgresql.md` - PostgreSQL best practices
- `mongodb.md` - MongoDB best practices
- `README.md` - Database rules overview

**Use Cases**:
- Database design
- Query optimization
- Database best practices

---

### 08. Testing

**Location**: `.cursor/rules/library/08-testing/`

**Purpose**: Testing framework rules and best practices.

**Contents**:
- Jest
- Vitest
- Go testing
- And more...

**Key Files**:
- `jest.md` - Jest best practices
- `vitest.md` - Vitest best practices
- `go-testing.md` - Go testing best practices
- `README.md` - Testing rules overview

**Use Cases**:
- Test writing
- Test organization
- Testing best practices

---

### 09. DevOps & CI/CD

**Location**: `.cursor/rules/library/09-devops-ci-cd/`

**Purpose**: DevOps and CI/CD rules and practices.

**Contents**:
- CI/CD best practices
- DevOps patterns
- Deployment strategies

**Key Files**:
- `README.md` - DevOps & CI/CD rules overview

**Use Cases**:
- CI/CD setup
- Deployment automation
- DevOps practices

---

### 10. Code Quality

**Location**: `.cursor/rules/library/10-code-quality/`

**Purpose**: Code quality rules and standards.

**Contents**:
- Linting rules
- Code style guidelines
- Quality standards

**Key Files**:
- `linting.md` - Linting rules
- `README.md` - Code quality rules overview

**Use Cases**:
- Code quality enforcement
- Linting configuration
- Code standards

---

### 11. Documentation

**Location**: `.cursor/rules/library/11-documentation/`

**Purpose**: Documentation rules and standards.

**Contents**:
- Documentation best practices
- Documentation standards
- Documentation patterns

**Key Files**:
- `README.md` - Documentation rules overview

**Use Cases**:
- Writing documentation
- Documentation standards
- Documentation patterns

---

### 12. Security

**Location**: `.cursor/rules/library/12-security/`

**Purpose**: Security rules and best practices.

**Contents**:
- Authentication
- Data encryption
- Security best practices

**Key Files**:
- `authentication.md` - Authentication rules
- `data-encryption.md` - Encryption rules
- `README.md` - Security rules overview

**Use Cases**:
- Security implementation
- Authentication setup
- Security best practices

---

### 13. Development Practices

**Location**: `.cursor/rules/library/13-development-practices/`

**Purpose**: General development practices and patterns.

**Contents**:
- Development patterns
- Best practices
- Development guidelines

**Key Files**:
- `README.md` - Development practices overview

**Use Cases**:
- Development workflows
- Best practices
- Development patterns

---

### 14. MCP Tools

**Location**: `.cursor/rules/library/14-mcp-tools/`

**Purpose**: MCP (Model Context Protocol) tools rules.

**Contents**:
- MCP tool integration
- MCP best practices

**Key Files**:
- `README.md` - MCP tools rules overview

**Use Cases**:
- MCP tool integration
- MCP best practices

---

### 15. Project-Specific

**Location**: `.cursor/rules/library/15-project-specific/`

**Purpose**: Project-specific rules and customizations.

**Contents**:
- Custom project rules
- Project-specific patterns

**Key Files**:
- `README.md` - Project-specific rules overview

**Use Cases**:
- Custom project rules
- Project-specific patterns

---

## Using Rules

### Automatic Rule Loading

When you generate a project, all rules are automatically extracted to `.cursor/rules/library/`. Agents automatically reference these rules when making decisions.

### Loading Rules Explicitly

Use the `/load` command to load specific rules into agent context:

```bash
/load @library/04-frameworks/nextjs.md
/load @library/07-databases/postgresql.md
/load @library/12-security/authentication.md
```

### Rule Reference in Commands

Rules are referenced in commands using the `@library/` prefix:

```bash
/load @library/03-languages/typescript.md
```

---

## Customizing Rules

### Modifying Existing Rules

1. Navigate to `.cursor/rules/library/`
2. Find the rule file you want to modify
3. Edit the markdown file
4. Save changes

**Example**: Modify TypeScript rules:
```bash
# Edit the file
.cursor/rules/library/03-languages/typescript.md

# Make your changes
# Save the file
```

### Adding New Rules

1. Navigate to the appropriate category directory
2. Create a new `.md` file
3. Write your rules following the existing format
4. Reference the rule using `/load`

**Example**: Add custom React rules:
```bash
# Create new file
.cursor/rules/library/04-frameworks/custom-react.md

# Write your rules
# Reference with /load @library/04-frameworks/custom-react.md
```

### Creating Custom Categories

1. Create a new directory in `.cursor/rules/library/`
2. Add a `README.md` explaining the category
3. Add your rule files
4. Reference rules using `/load`

---

## Creating Custom Rules

### Rule File Format

Rules are markdown files with the following structure:

```markdown
# Rule Title

## Overview
Brief description of the rule.

## Guidelines
- Guideline 1
- Guideline 2
- Guideline 3

## Examples

### Good Example
```code
// Good code example
```

### Bad Example
```code
// Bad code example
```

## Best Practices
- Practice 1
- Practice 2
```

### Rule Categories

Organize rules by:
- **Technology**: Language, framework, tool
- **Domain**: Security, testing, deployment
- **Pattern**: Design pattern, architecture pattern

### Rule Naming

Use descriptive names:
- `nextjs.md` - Next.js rules
- `postgresql.md` - PostgreSQL rules
- `authentication.md` - Authentication rules

---

## Rules Best Practices

### 1. Keep Rules Focused

Each rule file should focus on a specific topic:
- ✅ Good: `typescript.md` - TypeScript rules
- ❌ Bad: `everything.md` - All rules in one file

### 2. Use Clear Examples

Include both good and bad examples:
```markdown
### Good
```typescript
const user: User = { name: "John" };
```

### Bad
```typescript
const user = { name: "John" };
```
```

### 3. Keep Rules Updated

Update rules as best practices evolve:
- Review rules regularly
- Update for new versions
- Remove outdated practices

### 4. Document Rationale

Explain why rules exist:
```markdown
## Why This Rule Exists
This rule ensures type safety and prevents runtime errors.
```

### 5. Reference Standards

Link to official documentation:
```markdown
## References
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- [React Best Practices](https://react.dev/learn)
```

---

## Examples

### Example 1: Loading Framework Rules

```bash
# Load Next.js rules before building
/load @library/04-frameworks/nextjs.md
/build
```

### Example 2: Loading Database Rules

```bash
# Load PostgreSQL rules
/load @library/07-databases/postgresql.md
# Now agents will follow PostgreSQL best practices
```

### Example 3: Loading Security Rules

```bash
# Load authentication rules
/load @library/12-security/authentication.md
/safe  # Security audit
```

### Example 4: Custom Rule

Create `.cursor/rules/library/04-frameworks/custom-nextjs.md`:

```markdown
# Custom Next.js Rules

## Overview
Project-specific Next.js rules.

## Guidelines
- Use App Router exclusively
- Server components by default
- Client components only when needed
```

Then load it:
```bash
/load @library/04-frameworks/custom-nextjs.md
```

---

## Rule File Locations

Rules are organized in `.cursor/rules/library/`:

```
.cursor/rules/library/
├── 01-core-workflow/
├── 02-ai-agents/
├── 03-languages/
├── 04-frameworks/
├── 05-ui-libraries/
├── 06-cloud-infrastructure/
├── 07-databases/
├── 08-testing/
├── 09-devops-ci-cd/
├── 10-code-quality/
├── 11-documentation/
├── 12-security/
├── 13-development-practices/
├── 14-mcp-tools/
└── 15-project-specific/
```

---

## Related Pages

- [Commands Reference](Commands) - Using `/load` command
- [Agents Documentation](Agents) - How agents use rules
- [Customization Guide](Customization) - Customizing rules
- [Workflow Guide](Workflow) - Rules in workflow
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

