# Integration

Connect DoPlan with other tools and services. Integrate with IDEs, CI/CD, analytics, and more.

---

## 🎯 Integration Overview

DoPlan integrates with:
- **IDEs** - 6 AI-powered IDEs supported
- **Git** - Automatic commits and pushes
- **GitHub** - PR suggestions, workflows
- **CI/CD** - Automated pipelines
- **Analytics** - Time tracking data

---

## 💻 IDE Integration

### Supported IDEs

DoPlan supports 6 AI-powered IDEs:
- **Cursor** - AI-powered editor
- **Claude Code** - Anthropic's IDE
- **Antigravity** - AI development
- **Windsurf** - AI coding
- **Cline** - AI assistant
- **OpenCode** - Open source IDE

### How It Works

- **Agent files** - Compatible with all IDEs
- **Command structure** - Works with IDE command palette
- **Documentation** - IDE-specific guides
- **Workflow** - Seamless integration

---

## 🔀 Git Integration

### Automatic Operations

- **Auto-commit** - Conventional format
- **Auto-push** - To remote
- **Branch management** - Create/checkout
- **PR suggestions** - Pull request creation

### Git Workflow

```
/dev
# → Branch created: task/2.1

[Code]

/done
# → Auto-commit: feat(task-2.1): complete user authentication
# → Auto-push: Pushed to task/2.1
# → PR suggestion
```

---

## 🐙 GitHub Integration

### PR Suggestions

After `/done`:
```
💡 Suggestion: Create a pull request?
   Run: gh pr create --title "feat: User Authentication" --body "Completes task 2.1"
```

### GitHub CLI

If `gh` CLI installed:
- PR creation suggested
- Automatic PR creation (future)
- Workflow integration
- CI/CD triggers

---

## 🔄 CI/CD Integration

### Automated Pipelines

DoPlan generates:
- **GitHub Actions** - Workflow files
- **CI configuration** - Test automation
- **Deployment scripts** - Automated deployment
- **Quality checks** - Code quality

### Integration Points

- **Git hooks** - Pre-commit, pre-push
- **CI triggers** - On push, PR
- **Deployment** - Automated releases
- **Testing** - Automated test runs

---

## 📊 Analytics Integration

### Time Tracking Data

Time data in JSONL format:
```
.do/system/time-tracker.jsonl
```

**Use for**:
- Productivity dashboards
- Time analysis
- Billing (if needed)
- Reporting

### Export Options

- **JSONL format** - Analytics ready
- **CSV export** - Spreadsheet analysis
- **API access** - Programmatic access (future)
- **Dashboard** - Visual analytics (future)

---

## 💡 Integration Tips

### IDE Setup

- **Install IDE** - Choose your preferred IDE
- **Configure agents** - Set up agent files
- **Use command palette** - Access DoPlan commands
- **Workflow integration** - Seamless development

### Git Workflow

- **Initialize repository** - `git init`
- **Set remote** - `git remote add origin ...`
- **Use DoPlan** - Automatic Git operations
- **Review commits** - Check commit messages

### CI/CD Setup

- **Review generated workflows** - GitHub Actions
- **Configure CI** - Test automation
- **Set up deployment** - Automated releases
- **Quality checks** - Code quality

---

**Next**: [Troubleshooting](./04-Troubleshooting.md)

