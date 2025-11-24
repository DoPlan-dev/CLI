# GitHub Wiki Documentation Plan

This document outlines the complete structure for the DoPlan CLI GitHub Wiki.

## 📋 Wiki Structure

### Home Page (Main Wiki Index)

**File**: `Home.md`

**Content**:
- Welcome message
- Quick links to all wiki pages
- Overview of DoPlan CLI
- Navigation guide
- Getting started quick link

---

## 🚀 Getting Started Section

### Installation

**File**: `Installation.md`

**Sections**:
1. Prerequisites
2. Quick Install (npx)
3. macOS Installation
   - Homebrew
   - Direct binary
   - Build from source
4. Windows Installation
   - Scoop
   - Direct binary
   - Build from source
5. Linux Installation
   - Direct binary
   - Package managers (Debian/Ubuntu, Arch)
   - Build from source
6. Docker Installation
7. Verification
8. Troubleshooting Installation Issues

### Quick Start

**File**: `Quick-Start.md`

**Sections**:
1. Your First Project (5-minute tutorial)
2. Step-by-step walkthrough
3. What Gets Generated
4. Next Steps
5. Common First-Time Questions

### First Project Tutorial

**File**: `First-Project-Tutorial.md`

**Sections**:
1. Creating a project
2. Understanding the structure
3. Using your first command (`/tell`)
4. Brainstorming (`/improve`)
5. Generating plans (`/write`)
6. Approving (`/good`)
7. Creating tasks (`/tasks`)
8. Building (`/build`)
9. Completing your first feature

---

## 📖 Usage Guides

### Commands

**File**: `Commands.md`

**Sections**:
1. Command Overview
2. Core Commands
   - `/tell` - Capture ideas
   - `/improve` - Brainstorm
   - `/write` - Generate documents
   - `/change` - Edit documents
   - `/good` - Approve plan
   - `/tasks` - Generate tasks
   - `/build` - Start coding
   - `/progress` - Track progress
   - `/finished` - Complete tasks
3. Team Commands
   - `/team` - Show agents
   - `/load` - Inject context
4. Specialized Commands
   - `/ship` - Release management
   - `/safe` - Security audit
   - `/cheap` - Cost optimization
5. Command Examples
6. Command Tips & Tricks

### Workflow

**File**: `Workflow.md`

**Sections**:
1. Complete Development Workflow
2. Planning Phase
3. Development Phase
4. Review Phase
5. Release Phase
6. Workflow Best Practices
7. Workflow Diagrams
8. Common Workflow Patterns

### Agents

**File**: `Agents.md`

**Sections**:
1. Agent System Overview
2. Agent Hierarchy
3. Individual Agent Roles
   - Project Orchestrator
   - Product Manager
   - Engineering Lead
   - System Architect
   - Frontend Lead
   - Backend Lead
   - DevOps Engineer
   - Security Lead
   - Design & UX Manager
   - UI/UX Designer
   - QA & Reliability Manager
   - QA Engineer
   - Release & Growth Manager
   - Release Captain
   - Growth Coach
   - Documentation Lead
   - Documentation Writer
   - Performance Engineer
4. How Agents Work Together
5. Customizing Agents
6. Agent Best Practices

### Rules Library

**File**: `Rules.md`

**Sections**:
1. Rules Library Overview
2. Rules Categories
   - Core Workflow
   - AI Agents
   - Languages
   - Frameworks
   - UI Libraries
   - Cloud Infrastructure
   - Databases
   - Testing
   - DevOps & CI/CD
   - Code Quality
   - Documentation
   - Security
   - Development Practices
   - MCP Tools
   - Project-Specific
3. Using Rules
4. Customizing Rules
5. Creating Custom Rules
6. Rules Best Practices

---

## 🎓 Advanced Topics

### Advanced Usage

**File**: `Advanced.md`

**Sections**:
1. Custom Project Types
2. Custom Agent Templates
3. Custom Command Definitions
4. Extending the Rules Library
5. CI/CD Integration
6. Team Workflows
7. Multi-Project Management
8. Performance Optimization

### Architecture

**File**: `Architecture.md`

**Sections**:
1. System Architecture Overview
2. Component Structure
3. Data Flow
4. File System Organization
5. Agent System Design
6. Rules System Design
7. Command System Design
8. TUI Architecture
9. Binary Distribution
10. Extension Points

### Customization

**File**: `Customization.md`

**Sections**:
1. Customizing Agents
2. Customizing Commands
3. Customizing Rules
4. Custom Project Templates
5. Custom IDE Configurations
6. Custom Workflows
7. Best Practices for Customization

### IDE Integration

**File**: `IDE-Integration.md`

**Sections**:
1. Supported IDEs
   - Cursor
   - Claude Code
   - Antigravity
   - Windsurf
   - Cline
   - OpenCode
2. IDE-Specific Setup
3. IDE Configuration Files
4. Troubleshooting IDE Issues
5. IDE Best Practices

---

## 🔧 Technical Reference

### Project Structure

**File**: `Project-Structure.md`

**Sections**:
1. Generated Project Layout
2. Directory Explanations
3. File Purposes
4. Configuration Files
5. Generated Files Reference
6. Customization Points

### Configuration

**File**: `Configuration.md`

**Sections**:
1. Configuration Files
2. Environment Variables
3. CLI Flags
4. Project Settings
5. Agent Configuration
6. Rules Configuration

### API Reference

**File**: `API-Reference.md`

**Sections**:
1. Command Line API
2. Generated Files API
3. Agent System API
4. Rules System API
5. Extension API

---

## 🐛 Troubleshooting

### Troubleshooting

**File**: `Troubleshooting.md`

**Sections**:
1. Common Issues
   - Installation Problems
   - Binary Not Found
   - Permission Issues
   - Network Issues
   - IDE Integration Issues
2. Error Messages
3. Debug Mode
4. Log Files
5. Getting Help
6. Reporting Issues

### FAQ

**File**: `FAQ.md`

**Sections**:
1. General Questions
2. Installation Questions
3. Usage Questions
4. Technical Questions
5. Troubleshooting Questions
6. Contributing Questions

---

## 🤝 Contributing

### Contributing

**File**: `Contributing.md`

**Sections**:
1. How to Contribute
2. Development Setup
3. Code Style Guide
4. Testing Guidelines
5. Documentation Guidelines
6. Pull Request Process
7. Issue Reporting
8. Feature Requests

### Development

**File**: `Development.md`

**Sections**:
1. Building from Source
2. Development Environment Setup
3. Project Structure (for contributors)
4. Running Tests
5. Debugging
6. Release Process
7. Development Workflow

### Code of Conduct

**File**: `Code-of-Conduct.md`

**Sections**:
1. Our Pledge
2. Our Standards
3. Enforcement
4. Reporting

---

## 📚 Additional Resources

### Examples

**File**: `Examples.md`

**Sections**:
1. Example Projects
2. Use Cases
3. Real-World Examples
4. Example Workflows
5. Example Customizations

### Best Practices

**File**: `Best-Practices.md`

**Sections**:
1. Project Organization
2. Agent Usage
3. Command Usage
4. Rules Management
5. Team Collaboration
6. Version Control
7. CI/CD Integration

### Migration Guide

**File**: `Migration-Guide.md`

**Sections**:
1. Upgrading DoPlan CLI
2. Migrating Projects
3. Breaking Changes
4. Compatibility

### Release Notes

**File**: `Release-Notes.md`

**Sections**:
1. Version History
2. Changelog
3. Upgrade Guides
4. Deprecation Notices

---

## 📝 Wiki Page Templates

### Standard Page Template

```markdown
# Page Title

Brief description of what this page covers.

## Overview

[Overview section]

## Section 1

[Content]

## Section 2

[Content]

## Related Pages

- [Link to related page 1](Related-Page-1)
- [Link to related page 2](Related-Page-2)

## See Also

- [Main Documentation](Home)
- [Quick Start](Quick-Start)
```

### Command Reference Template

```markdown
# Command Name

## Description

[Command description]

## Usage

\`\`\`bash
/command [options] [arguments]
\`\`\`

## Options

| Option | Description | Required |
|--------|-------------|----------|
| `--flag` | Flag description | No |

## Arguments

| Argument | Description | Required |
|----------|-------------|----------|
| `arg` | Argument description | Yes |

## Examples

### Example 1
\`\`\`bash
/command example
\`\`\`

## Related Commands

- [Related Command 1](Commands#related-command-1)
- [Related Command 2](Commands#related-command-2)

## See Also

- [Commands Overview](Commands)
- [Workflow Guide](Workflow)
```

---

## 🎯 Implementation Priority

### Phase 1: Essential Pages (Week 1)
1. Home
2. Installation
3. Quick Start
4. Commands
5. FAQ
6. Troubleshooting

### Phase 2: Core Documentation (Week 2)
1. Workflow
2. Agents
3. Rules Library
4. First Project Tutorial
5. Contributing

### Phase 3: Advanced Topics (Week 3)
1. Advanced Usage
2. Architecture
3. Customization
4. IDE Integration
5. Best Practices

### Phase 4: Reference & Examples (Week 4)
1. Project Structure
2. Configuration
3. API Reference
4. Examples
5. Migration Guide

---

## 📋 Content Guidelines

### Writing Style
- Clear and concise
- Beginner-friendly but professional
- Use examples liberally
- Include code snippets
- Add diagrams where helpful
- Use consistent formatting

### Code Examples
- Always include working examples
- Show both simple and advanced use cases
- Include expected output
- Add comments for clarity

### Visual Elements
- Use emojis sparingly for visual breaks
- Include diagrams for complex concepts
- Use tables for structured data
- Add screenshots for UI elements

### Links
- Link to related pages
- Link to external resources when relevant
- Keep links up to date
- Use descriptive link text

---

## ✅ Checklist for Each Wiki Page

- [ ] Clear title and description
- [ ] Table of contents (for long pages)
- [ ] Overview section
- [ ] Detailed content sections
- [ ] Code examples (where applicable)
- [ ] Related pages links
- [ ] See also section
- [ ] Proper formatting
- [ ] No broken links
- [ ] Reviewed for accuracy

---

## 🔄 Maintenance

### Regular Updates
- Review pages quarterly
- Update for new features
- Fix broken links
- Update examples
- Add new FAQs

### Version Control
- Keep wiki in sync with code
- Document breaking changes
- Maintain changelog
- Update migration guides

---

**Last Updated**: [Date]
**Maintained By**: Documentation Team




