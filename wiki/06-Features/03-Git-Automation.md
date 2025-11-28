# Git Automation

DoPlan automates Git operations so you can focus on coding. Auto-commit, auto-push, and PR suggestions - all handled automatically.

---

## 🎯 What is Git Automation?

Git Automation in DoPlan handles:
- **Auto-commit** - Conventional commit format automatically
- **Auto-push** - Push to remote without manual commands
- **Branch management** - Create/checkout branches automatically
- **PR suggestions** - Suggest pull request creation

---

## 🔄 Auto-Commit

### When It Happens

Auto-commit occurs when you run `/done`:
1. Task marked complete
2. Changes staged
3. Commit created automatically
4. Conventional format used

### Commit Format

Conventional commit format:
```
feat(task-2.1): complete user authentication
fix(task-3.2): resolve login bug
docs(task-1.1): update API documentation
```

### Format Breakdown

- **Type**: `feat`, `fix`, `docs`, `refactor`, `test`, etc.
- **Scope**: `task-[ID]` - Task identifier
- **Description**: Brief task description

### Examples

```
feat(task-2.1): complete user authentication
fix(task-3.2): resolve login validation bug
docs(task-1.1): update API documentation
refactor(task-4.3): improve code structure
test(task-5.1): add unit tests
```

---

## 🚀 Auto-Push

### When It Happens

Auto-push occurs after auto-commit:
1. Commit created
2. Push to remote automatically
3. Current branch pushed
4. No manual command needed

### Branch Handling

- **Task branches** - Pushed automatically
- **Main/master** - Protected (warns before commit)
- **Remote sync** - Automatic

### Process

```bash
/done
# → Task complete
# → Auto-commit: feat(task-2.1): complete user authentication
# → Auto-push: Pushed to task/2.1
```

---

## 🌿 Branch Management

### Automatic Branch Creation

When you run `/dev`:
1. Finds next task
2. Creates branch: `task/[ID]`
3. Checks out branch
4. Ready for development

### Branch Naming

Format: `task/[ID]`

Examples:
- `task/1.1` - Task 1.1
- `task/2.3` - Task 2.3
- `task/3.5` - Task 3.5

### Branch Lifecycle

1. **Created** - When `/dev` runs
2. **Active** - During development
3. **Committed** - When `/done` runs
4. **Pushed** - Auto-push to remote
5. **Merged** - Via PR (manual or suggested)

---

## 🔀 PR Suggestions

### When Suggested

PR suggestions appear after `/done`:
```
💡 Suggestion: Create a pull request?
   Run: gh pr create --title "feat: User Authentication" --body "Completes task 2.1"
```

### Requirements

- GitHub CLI (`gh`) installed
- Remote repository configured
- Branch pushed to remote

### PR Creation

Manual creation:
```bash
gh pr create --title "feat: User Authentication" --body "Completes task 2.1"
```

Or use suggested command from DoPlan.

---

## 🎯 Git Workflow Integration

### Standard Workflow

```
/dev
# → Branch created: task/2.1
# → Checked out

[Code]
# → Develop feature

/done
# → Auto-commit: feat(task-2.1): complete user authentication
# → Auto-push: Pushed to task/2.1
# → PR suggestion
```

### Branch Strategy

Follows best practices:
- **Feature branches** - One per task
- **Conventional commits** - Standard format
- **Remote sync** - Automatic
- **PR workflow** - Suggested

---

## 🔒 Safety Features

### Branch Verification

Before `/done`:
- Verifies you're on task branch
- Warns if on main/master
- Confirms before proceeding

### Dependency Checking

Before completion:
- Checks task dependencies
- Verifies prerequisites
- Blocks if dependencies incomplete

### Commit Safety

- Conventional format ensures consistency
- Task ID in scope for tracking
- Clear commit messages
- Git history quality

---

## 💡 Git Automation Tips

### Best Practices

1. **Always use `/done`** - Ensures auto-commit/push
2. **Review commits** - Check commit messages
3. **Use PR suggestions** - Create pull requests
4. **Verify branches** - Check you're on correct branch

### Workflow Optimization

- Let DoPlan handle Git operations
- Focus on coding, not Git commands
- Use PR workflow for code review
- Maintain clean Git history

### Advanced Usage

- Customize commit messages (if needed)
- Use PR templates
- Integrate with CI/CD
- Leverage branch protection

---

## 🚀 Git Integration Examples

### Example 1: Standard Task

```bash
/dev
# → Branch: task/2.1 created

[Code for 2 hours]

/done
# → Commit: feat(task-2.1): complete user authentication
# → Push: Pushed to task/2.1
# → PR suggestion shown
```

### Example 2: Multiple Tasks

```bash
/dev
# → Branch: task/2.1
[Code]
/done
# → Committed and pushed

/dev
# → Branch: task/2.2
[Code]
/done
# → Committed and pushed
```

### Example 3: With PR

```bash
/dev
[Code]
/done
# → Committed and pushed
# → PR suggestion

gh pr create --title "feat: User Authentication" --body "Completes task 2.1"
# → PR created
```

---

## 🎯 Git Automation Benefits

### For You

1. **No manual Git commands** - Everything automatic
2. **Consistent commits** - Conventional format
3. **Clean history** - Quality Git log
4. **Time saved** - Focus on coding
5. **Less errors** - Automated process

### For Team

1. **Standard workflow** - Consistent process
2. **Clear commits** - Easy to understand
3. **PR workflow** - Code review ready
4. **Branch strategy** - Best practices
5. **CI/CD ready** - Automated triggers

---

**Next**: [Learning Support](./04-Learning-Support.md)

