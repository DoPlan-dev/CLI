package generators

// READMEGenerator generates README.md
type READMEGenerator struct {
	projectRoot string
}

// NewREADMEGenerator creates a new README generator
func NewREADMEGenerator(projectRoot string) *READMEGenerator {
	return &READMEGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates the README.md content
func (g *READMEGenerator) Generate() string {
	codeBlock := "```"
	backtick := "`"
	
	return `# DoPlan - Project Workflow Manager

DoPlan automates your project workflow from idea to deployment, combining **Spec-Kit** and **BMAD-METHOD** methodologies. It helps you create well-documented plans, manage features with Git branches, and track progress visually.

## 🚀 Quick Start

### 1. Verify Installation

DoPlan is already installed in this project! You can verify by checking:

- ✅ ` + backtick + `.cursor/commands/` + backtick + ` - Contains all DoPlan commands
- ✅ ` + backtick + `doplan/` + backtick + ` - Your project planning directory
- ✅ ` + backtick + `.cursor/config/doplan-config.json` + backtick + ` - Configuration file

### 2. View Dashboard

` + codeBlock + `bash
doplan dashboard
` + codeBlock + `

Or use the ` + backtick + `/Dashboard` + backtick + ` command in Cursor.

---

## 📋 Complete Workflow Guide

Follow these steps to plan and develop your project using DoPlan:

### **Step 1: Discuss Your Idea** 🗣️

Start by discussing and refining your project idea.

**In Cursor:**
` + codeBlock + `
/Discuss
` + codeBlock + `

**What happens:**
- Cursor will ask comprehensive questions about your idea
- It will suggest improvements and enhancements
- Help organize features into logical phases
- Recommend the best tech stack for your project
- Save everything to ` + backtick + `doplan/idea-notes.md` + backtick + ` and ` + backtick + `.cursor/config/doplan-state.json` + backtick + `

### **Step 2: Refine Your Idea** ✨

Enhance and improve your initial idea.

**In Cursor:**
` + codeBlock + `
/Refine
` + codeBlock + `

### **Step 3: Generate Documentation** 📄

Create comprehensive project documentation.

**In Cursor:**
` + codeBlock + `
/Generate
` + codeBlock + `

**What gets created:**
- ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
- ` + backtick + `doplan/structure.md` + backtick + ` - Project structure and architecture
- ` + backtick + `doplan/contracts/api-spec.json` + backtick + ` - API specification (OpenAPI/Swagger)
- ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models and schemas

### **Step 4: Create Your Plan** 📊

Generate the phase and feature structure.

**In Cursor:**
` + codeBlock + `
/Plan
` + codeBlock + `

**What gets created:**
` + codeBlock + `
doplan/
├── 01-phase/
│   ├── phase-plan.md
│   ├── phase-progress.json
│   ├── 01-Feature/
│   │   ├── plan.md
│   │   ├── design.md
│   │   ├── tasks.md
│   │   └── progress.json
│   └── 02-Feature/
│       └── ...
├── 02-phase/
│   └── ...
└── dashboard.md (updated)
` + codeBlock + `

### **Step 5: View Your Dashboard** 📈

See your project progress at a glance.

**In Cursor:**
` + codeBlock + `
/Dashboard
` + codeBlock + `

**Or in Terminal:**
` + codeBlock + `bash
doplan dashboard
` + codeBlock + `

### **Step 6: Start Implementing a Feature** 💻

Begin development on a specific feature.

**In Cursor:**
` + codeBlock + `
/Implement
` + codeBlock + `

**What happens automatically:**
1. ✅ Creates a Git branch: ` + backtick + `feature/01-phase-01-feature-name` + backtick + `
2. ✅ Commits initial feature files (plan.md, design.md, tasks.md)
3. ✅ Pushes branch to remote repository
4. ✅ Updates dashboard with branch information
5. ✅ Sets up the feature for development

### **Step 7: Develop the Feature** 🛠️

Follow the feature's documentation while developing.

**Workflow:**
1. Review the feature's ` + backtick + `plan.md` + backtick + `, ` + backtick + `design.md` + backtick + `, and ` + backtick + `tasks.md` + backtick + `
2. Implement tasks in order
3. Check off tasks in ` + backtick + `tasks.md` + backtick + ` as you complete them
4. Commit regularly: ` + backtick + `git commit -m "feat: implement [task description]"` + backtick + `
5. Push: ` + backtick + `git push origin feature/01-phase-01-feature-name` + backtick + `
6. Update progress: ` + backtick + `/Progress` + backtick + `

### **Step 8: Test Your Feature** 🧪

Test thoroughly before marking complete.

**Testing checklist:**
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed
- [ ] Edge cases handled
- [ ] Error handling works

### **Step 9: Fix Gaps and Issues** 🔧

Address any gaps or issues found during testing.

**Workflow:**
1. Identify the gap/issue
2. Fix it in your feature branch
3. Test the fix
4. Commit: ` + backtick + `git commit -m "fix: [description]"` + backtick + `
5. Push: ` + backtick + `git push origin feature/01-phase-01-feature-name` + backtick + `

### **Step 10: Complete the Feature** ✅

Mark the feature as complete and prepare for merge.

**Before completing:**
- [ ] All tasks in ` + backtick + `tasks.md` + backtick + ` are checked
- [ ] All tests pass
- [ ] Code is reviewed
- [ ] Documentation is updated
- [ ] No known bugs or issues

**Mark as complete:**
1. Update ` + backtick + `tasks.md` + backtick + ` - ensure all tasks are checked
2. Update ` + backtick + `progress.json` + backtick + ` - set ` + backtick + `"status": "complete"` + backtick + `
3. Commit final changes:
   ` + codeBlock + `bash
   git add .
   git commit -m "feat: complete [feature name]"
   git push origin feature/01-phase-01-feature-name
   ` + codeBlock + `

### **Step 11: Create Pull Request** 🔀

Create a PR to merge your feature into main.

**Automatically (if configured):**
- DoPlan can auto-create PR when feature is marked complete

**Manually:**
` + codeBlock + `bash
gh pr create --title "Feature: [Feature Name]" \
  --body "Implements [feature name] from Phase [X]"
` + codeBlock + `

### **Step 12: Review and Merge** 👀

Review the PR and merge when ready.

**After merge:**
1. Update local main branch:
   ` + codeBlock + `bash
   git checkout main
   git pull origin main
   ` + codeBlock + `

2. Update dashboard:
   ` + codeBlock + `
   /Progress
   ` + codeBlock + `

3. Start next feature:
   ` + codeBlock + `
   /Next
   ` + codeBlock + `

---

## 📚 Available Commands

### Cursor Slash Commands

| Command | Description |
|---------|-------------|
| ` + backtick + `/Discuss` + backtick + ` | Start idea discussion and refinement |
| ` + backtick + `/Refine` + backtick + ` | Enhance and improve your idea |
| ` + backtick + `/Generate` + backtick + ` | Generate PRD, Structure, and API contracts |
| ` + backtick + `/Plan` + backtick + ` | Generate phase and feature structure |
| ` + backtick + `/Dashboard` + backtick + ` | Show project dashboard with progress |
| ` + backtick + `/Implement` + backtick + ` | Start implementing a feature (creates branch) |
| ` + backtick + `/Next` + backtick + ` | Get recommendation for next action |
| ` + backtick + `/Progress` + backtick + ` | Update all progress tracking |

### CLI Commands

| Command | Description |
|---------|-------------|
| ` + backtick + `doplan` + backtick + ` | Show dashboard or installation menu |
| ` + backtick + `doplan install` + backtick + ` | Install/reinstall DoPlan |
| ` + backtick + `doplan dashboard` + backtick + ` | View project dashboard in terminal |
| ` + backtick + `doplan github` + backtick + ` | Update GitHub data (branches, commits, pushes) |
| ` + backtick + `doplan progress` + backtick + ` | Update all progress tracking |

---

## 📁 Project Structure

` + codeBlock + `
project-root/
├── .cursor/
│   ├── commands/          # DoPlan command definitions
│   ├── rules/             # Workflow rules and policies
│   └── config/            # Configuration and state
│       ├── doplan-config.json
│       └── doplan-state.json
├── doplan/                # Planning directory
│   ├── dashboard.md       # Visual progress dashboard
│   ├── PRD.md             # Product Requirements Document
│   ├── structure.md       # Project structure
│   ├── idea-notes.md      # Idea discussion notes
│   ├── contracts/         # API contracts
│   │   ├── api-spec.json
│   │   └── data-model.md
│   ├── templates/         # Reusable templates
│   │   ├── plan-template.md
│   │   ├── design-template.md
│   │   └── tasks-template.md
│   ├── 01-phase/          # Phase 1
│   │   ├── phase-plan.md
│   │   ├── phase-progress.json
│   │   ├── 01-Feature/
│   │   │   ├── plan.md
│   │   │   ├── design.md
│   │   │   ├── tasks.md
│   │   │   └── progress.json
│   │   └── 02-Feature/
│   │       └── ...
│   └── 02-phase/          # Phase 2
│       └── ...
└── README.md              # This file
` + codeBlock + `

---

## 🔄 Git Workflow

### Branch Naming Convention

- Format: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
- Example: ` + backtick + `feature/01-phase-01-user-authentication` + backtick + `
- Use kebab-case for feature names

### Typical Git Workflow

` + codeBlock + `bash
# 1. Start a new feature (automated by /Implement)
git checkout -b feature/01-phase-01-feature-name
git add doplan/01-phase/01-Feature/*
git commit -m "docs: add feature planning docs"
git push origin feature/01-phase-01-feature-name

# 2. Develop and commit regularly
git add .
git commit -m "feat: implement [task]"
git push origin feature/01-phase-01-feature-name

# 3. Fix issues
git add .
git commit -m "fix: [description]"
git push origin feature/01-phase-01-feature-name

# 4. After PR merge, update main
git checkout main
git pull origin main
` + codeBlock + `

---

## 🎯 Best Practices

### Planning Phase
- ✅ Be thorough in idea discussion
- ✅ Review and refine before generating docs
- ✅ Ensure PRD is complete before planning
- ✅ Break features into small, manageable tasks

### Development Phase
- ✅ Follow the feature's plan.md and design.md
- ✅ Check off tasks as you complete them
- ✅ Commit frequently with clear messages
- ✅ Keep feature branch updated with main
- ✅ Test as you develop, not just at the end

### Testing Phase
- ✅ Write tests before or alongside code
- ✅ Test edge cases and error scenarios
- ✅ Test integration with other features
- ✅ Fix issues before marking complete

---

*Generated by DoPlan - Project Workflow Manager*
`
}
