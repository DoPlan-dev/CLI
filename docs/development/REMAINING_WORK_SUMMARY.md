# Remaining Work from next-2-beta.md

Based on analysis of the codebase and `next-2-beta.md`, here's what still needs to be implemented.

## 📊 Overall Status

- **v0.0.18-beta:** ~80% Complete
- **v0.0.19-beta:** ~5% Complete

---

## 🚀 v0.0.18-beta - Remaining Work

### ✅ Phase 1: Architecture Setup - **COMPLETE**
- ✅ Project structure exists
- ✅ Configuration system (Viper)
- ✅ State management (.doplan/state.json)
- ✅ Dashboard JSON format
- ✅ Lipgloss theme system

### ✅ Phase 2: Smart Root Command & Context Detection - **COMPLETE**
- ✅ Context detection system (`internal/context/detector.go`)
- ✅ Smart root command behavior (`cmd/doplan/main.go`)
- ✅ Dashboard aliases (".", "dash", "d")
- ✅ New project wizard (`internal/wizard/new_project.go`) - **Just updated with header, templates, GitHub flow**
- ✅ Adopt project wizard (`internal/wizard/adopt_project.go`)
- ✅ Project analyzer (`internal/context/analyzer.go`)

### ⚠️ Phase 3: Dashboard Supercharge - **90% Complete**

**What's Done:**
- ✅ Dashboard JSON format exists
- ✅ Dashboard generator (`internal/dashboard/loader.go`, `internal/generators/dashboard.go`)
- ✅ TUI reads JSON (`internal/tui/screens/dashboard.go`)
- ✅ Basic sparklines (`internal/dashboard/sparkline.go`)

**What's Missing:**
- ❌ Enhanced sparklines with 14-day history and trend colors
- ❌ Auto-refresh every 30 seconds
- ❌ Activity feed with icons and time ago formatting
- ❌ Auto-update triggers on progress/git operations

**Files to Update:**
- `internal/dashboard/sparkline.go` - Add 14-day history and color coding
- `internal/tui/screens/dashboard.go` - Add auto-refresh and activity feed
- `internal/dashboard/updater.go` - Ensure triggers work

---

### ⚠️ Phase 4: Project-First Documentation - **60% Complete**

**What's Done:**
- ✅ CONTEXT.md generator exists (`internal/generators/context.go`)
- ✅ README.md generator exists (`internal/generators/readme.go`)

**What's Missing:**

1. **CONTEXT.md Structure Update**
   - ❌ Project Overview section (description, audience, features)
   - ❌ Technology Stack (Frontend/Backend/Services breakdown)
   - ❌ Project-Specific Documentation links
   - ❌ Development Guidelines section
   - ❌ DoPlan Resources in collapsible `<details>` section (currently at top)

2. **README.md Restructure**
   - ❌ Project-first content (currently DoPlan-focused)
   - ❌ DoPlan info moved to collapsible `<details>` at bottom
   - ❌ Project structure showing `##-phase-name/##-feature-name` format
   - ❌ Link to RAKD.md for environment variables

**Files to Update:**
- `internal/generators/context.go` - Restructure to match plan
- `internal/generators/readme.go` - Restructure to be project-first

**Estimated Effort:** 1-2 days

---

### ⚠️ Phase 5: GitHub & IDE Integration - **90% Complete**

**What's Done:**
- ✅ GitHub validator exists (`internal/github/validator.go`)
- ✅ IDE integration logic (`internal/integration/setup.go`)
- ✅ Cursor integration
- ✅ VS Code integration
- ✅ Generic/Other integration guides
- ✅ GitHub setup in new project wizard (just updated)

**What's Missing:**

1. **GitHub Requirement Enforcement**
   - ❌ Block protected actions without repo (discuss, plan, implement, feature)
   - ❌ Standalone GitHub setup wizard TUI screen (`internal/tui/screens/github_setup.go`)
   - ❌ Show beautiful TUI error when repo missing
   - ❌ Offer to launch GitHub setup wizard

2. **GitHub Badge on Dashboard**
   - ❌ Repository name as clickable badge
   - ❌ Display commit count and last commit time
   - ❌ Style with lipgloss (rounded border, primary color)

**Files to Create/Update:**
- `internal/tui/screens/github_setup.go` - Standalone wizard
- `internal/tui/screens/dashboard.go` - Add GitHub badge
- `internal/commands/*.go` - Add GitHub checks before protected actions

**Estimated Effort:** 2-3 days

---

### ⚠️ Phase 6: Foundational Polish - **85% Complete**

**What's Done:**
- ✅ Lipgloss theme system (`pkg/theme/`)
- ✅ Error handling (`pkg/errors/`, `internal/error/`)
- ✅ Logging (`pkg/logger/`)
- ✅ Animations (`pkg/animations/`)
- ✅ Tool installer (`pkg/tools/installer.go`)

**What's Missing:**
- ❌ Full TUI audit for consistent styling (all screens)
- ❌ Enhanced sparklines with trend colors (green/amber/red)
- ❌ Ensure all headers use HeaderStyle
- ❌ Ensure all cards use CardStyle
- ❌ Ensure all buttons use ButtonStyle

**Files to Update:**
- All TUI screen files in `internal/tui/screens/`
- `internal/dashboard/sparkline.go` - Add color coding

**Estimated Effort:** 3-5 days

---

## 🚀 v0.0.19-beta - Remaining Work

### ❌ Phase 1: Unified TUI & AI Commands - **0% Complete**

**What Needs Implementation:**

1. **AI Command Definitions** (`.doplan/ai/commands/`)
   - ❌ `run.md` - Run dev server command
   - ❌ `undo.md` - Undo last action command
   - ❌ `deploy.md` - Deployment wizard command
   - ❌ `publish.md` - Package publishing command
   - ❌ `create.md` - New project wizard command
   - ❌ `security.md` - Security scan command
   - ❌ `fix.md` - Auto-fix command
   - ❌ `design.md` - Design system command
   - ❌ `keys.md` - API keys management command

2. **Backend Action Logic** (`internal/commands/`)
   - ❌ `run.go` - Auto-detect and run dev server
   - ❌ `undo.go` - Time-machine undo using state.json
   - ❌ `deploy.go` - Multi-platform deployment wizard
   - ❌ `publish.go` - Package publishing wizard
   - ❌ `security.go` - Comprehensive security scan
   - ❌ `fix.go` - AI-powered auto-fix

3. **Deployment System** (`internal/deployment/`)
   - ❌ Directory exists but empty
   - ❌ Support for: Vercel, Netlify, Railway, Render, Coolify, custom
   - ❌ Deployment wizard TUI

4. **Publisher System** (`internal/publisher/`)
   - ❌ Directory exists but empty
   - ❌ Support for: npm, Homebrew, Scoop, Winget
   - ❌ Publishing wizard TUI

5. **Security System** (`internal/security/`)
   - ❌ Directory exists but empty
   - ❌ npm audit, trufflehog, git-secrets, gosec, dive
   - ❌ Security scan TUI

6. **Fixer System** (`internal/fixer/`)
   - ❌ Directory exists but empty
   - ❌ AI-powered auto-fix logic
   - ❌ Fix wizard TUI

7. **Full TUI Menu Population**
   - ⚠️ Dashboard exists
   - ❌ All menu items: Run, Undo, Deploy, Publish, Security, Fix, etc.

**Estimated Effort:** 3-4 weeks

---

### ❌ Phase 2: Design System (DPR) Generation - **0% Complete**

**What Needs Implementation:**

1. **DPR Command & TUI** (`internal/dpr/`)
   - ❌ Directory exists but empty
   - ❌ `questionnaire.go` - Interactive 20-30 question TUI
   - ❌ `generator.go` - Generate DPR.md document
   - ❌ `tokens.go` - Generate design-tokens.json
   - ❌ `cursor_rules.go` - Generate design_rules.mdc

2. **Questionnaire Topics:**
   - ❌ Audience analysis
   - ❌ Emotional design
   - ❌ Style preferences
   - ❌ Colors
   - ❌ Typography
   - ❌ Layout
   - ❌ Components
   - ❌ Animation
   - ❌ References

3. **DPR.md Structure:**
   - ❌ Executive Summary
   - ❌ Audience Analysis
   - ❌ Design Principles
   - ❌ Visual Identity
   - ❌ Layout guidelines
   - ❌ Component Library
   - ❌ Animation guidelines
   - ❌ Wireframes
   - ❌ Accessibility

**Estimated Effort:** 2-3 weeks

---

### ❌ Phase 3: Secrets & API Keys (RAKD/SOPS) - **0% Complete**

**What Needs Implementation:**

1. **SOPS System** (`internal/sops/`)
   - ❌ Directory exists but empty
   - ❌ `generator.go` - Auto-generate service setup guides
   - ❌ `detector.go` - Auto-detect services from dependencies
   - ❌ Generate guides in `.doplan/SOPS/`:
     - authentication/
     - database/
     - payment/
     - storage/
     - email/
     - analytics/
     - ai/

2. **RAKD System** (`internal/rakd/`)
   - ❌ Directory exists but empty
   - ❌ `generator.go` - Generate RAKD.md
   - ❌ `detector.go` - Detect required API keys
   - ❌ `validator.go` - Validate API keys

3. **Keys Management TUI**
   - ❌ TUI screen for key management
   - ❌ Show RAKD status
   - ❌ Validate all keys
   - ❌ Check for missing keys
   - ❌ Sync .env.example
   - ❌ Launch setup wizard for services
   - ❌ Test API connections

4. **Dashboard Widget**
   - ❌ API Keys Status card
   - ❌ Progress bar of configuration
   - ❌ Count of configured/pending/optional
   - ❌ Highlight high-priority missing keys

**Estimated Effort:** 2-3 weeks

---

### ❌ Phase 4: AI Agents System - **0% Complete**

**What Needs Implementation:**

1. **Agent Files** (`.doplan/ai/agents/`)
   - ❌ `README.md` - Comprehensive guide
   - ❌ `planner.agent.md` - Planner agent definition
   - ❌ `coder.agent.md` - Coder agent definition
   - ❌ `designer.agent.md` - Designer agent definition
   - ❌ `reviewer.agent.md` - Reviewer agent definition
   - ❌ `tester.agent.md` - Tester agent with Playwright
   - ❌ `devops.agent.md` - DevOps agent definition

2. **Workflow Rules** (`.doplan/ai/rules/`)
   - ❌ `workflow.mdc` - Perfect workflow sequence
   - ❌ `communication.mdc` - Agent interaction rules
   - ❌ `design_rules.mdc` - Design system rules (from Phase 2)

3. **Agent Generator** (`internal/agents/generator.go`)
   - ⚠️ `internal/generators/agents.go` exists but may not generate all agents
   - ❌ Need to generate all agent files during installation
   - ❌ Need to generate workflow and communication rules

4. **Tester Agent Features:**
   - ❌ Playwright (MCP) test execution
   - ❌ Screenshot capture to `.doplan/artifacts/screenshots/`
   - ❌ Visual regression checks
   - ❌ Bug reporting with screenshots

**Estimated Effort:** 3-4 weeks

---

### ❌ Phase 5: Workflow Guidance Engine - **0% Complete**

**What Needs Implementation:**

1. **Workflow Recommender** (`internal/workflow/`)
   - ❌ Directory exists but empty
   - ❌ `recommender.go` - `GetNextStep(lastAction string)` function
   - ❌ Map actions to recommended next steps
   - ❌ Follow workflow.mdc sequence

2. **TUI Integration**
   - ❌ "Recommended Next Step" box in TUI
   - ❌ Display after every successful action
   - ❌ Use SuccessStyle border (Lipgloss)
   - ❌ Show action and instructions

**Estimated Effort:** 1 week

---

## 📋 Priority Recommendations

### 🔴 Critical (Finish v0.0.18-beta) - 1-2 weeks

1. **Documentation Structure** (Phase 4)
   - Update CONTEXT.md generator
   - Update README.md generator
   - **Effort:** 1-2 days

2. **GitHub Enforcement** (Phase 5)
   - Add requirement checks
   - Create standalone setup wizard
   - Add GitHub badge to dashboard
   - **Effort:** 2-3 days

3. **TUI Polish** (Phase 6)
   - Audit all screens for consistent styling
   - Enhance sparklines
   - **Effort:** 3-5 days

### 🟡 High Priority (Start v0.0.19-beta) - 3-4 weeks

1. **Basic TUI Actions** (Phase 1)
   - Run Dev Server
   - Undo Last Action
   - AI command definitions
   - **Effort:** 1 week

2. **AI Agents** (Phase 4)
   - Generate all agent files
   - Create workflow rules
   - **Effort:** 2 weeks

### 🟢 Medium Priority (Complete v0.0.19-beta) - 8-10 weeks

1. **DPR System** (Phase 2) - 2-3 weeks
2. **RAKD/SOPS** (Phase 3) - 2-3 weeks
3. **Deployment System** (Phase 1) - 1-2 weeks
4. **Publisher System** (Phase 1) - 1 week
5. **Security System** (Phase 1) - 1 week
6. **Fixer System** (Phase 1) - 1 week
7. **Workflow Guidance** (Phase 5) - 1 week

---

## 🎯 Next Steps

1. **Complete v0.0.18-beta polish** (1-2 weeks)
   - Documentation updates
   - GitHub enforcement
   - TUI audit

2. **Start v0.0.19-beta Phase 1** (3-4 weeks)
   - Basic TUI actions (Run, Undo)
   - AI command definitions
   - Full TUI menu

3. **Continue with v0.0.19-beta** (8-10 weeks total)
   - DPR system
   - RAKD/SOPS
   - AI Agents
   - Workflow Guidance

---

## 📝 Notes

- Most infrastructure is in place
- v0.0.18-beta is close to completion (~80%)
- v0.0.19-beta requires significant new development (~5%)
- Focus on completing v0.0.18-beta before starting v0.0.19-beta
- Recent work: Updated new project wizard with header, templates, and GitHub flow

