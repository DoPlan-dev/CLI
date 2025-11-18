# Remaining Development Work - next-2-beta.md

Based on analysis of `next-2-beta.md` and the current codebase, here's what still needs development.

## 🎯 Quick Summary

- **v0.0.18-beta:** ~75% Complete - Mostly polish and documentation updates needed
- **v0.0.19-beta:** ~5% Complete - Not started, all features need implementation

---

## 📋 v0.0.18-beta - Remaining Work

### ⚠️ Phase 4: Project-First Documentation (40% Remaining)

**Current Status:** Generators exist but structure doesn't match plan exactly.

**What Needs Work:**

1. **CONTEXT.md Structure Update**
   - ✅ Generator exists: `internal/generators/context.go`
   - ❌ Structure needs update to match plan:
     - Project Overview section
     - Technology Stack (Frontend/Backend/Services)
     - Project-Specific Documentation links
     - Development Guidelines
     - DoPlan Resources in collapsible `<details>` section
   
2. **README.md Restructure**
   - ✅ Generator exists: `internal/generators/readme.go`
   - ❌ Structure needs update:
     - Project-first content (not DoPlan-first)
     - DoPlan info moved to collapsible `<details>` section at bottom
     - Project structure showing `##-phase-name/##-feature-name` format

**Files to Update:**
- `internal/generators/context.go`
- `internal/generators/readme.go`

---

### ⚠️ Phase 5: GitHub & IDE Integration (10% Remaining)

**Current Status:** Integration logic exists, but GitHub requirement enforcement needs work.

**What Needs Work:**

1. **GitHub Requirement Enforcement**
   - ⚠️ Validator exists but not enforced everywhere
   - ❌ Need to add checks before protected actions
   - ❌ Need standalone GitHub setup wizard TUI screen
   - ❌ Need to block actions without repo

2. **GitHub Badge on Dashboard**
   - ⚠️ May exist, needs verification
   - ❌ Should show repository name as clickable badge
   - ❌ Display commit count and last commit time

**Files to Update:**
- `internal/github/validator.go` (create if doesn't exist)
- `internal/tui/screens/github_setup.go` (create standalone wizard)
- `internal/tui/screens/dashboard.go` (add GitHub badge)

---

### ⚠️ Phase 6: Foundational Polish (15% Remaining)

**Current Status:** Core systems exist, needs full audit.

**What Needs Work:**

1. **Full TUI Audit**
   - ✅ Theme system exists: `pkg/theme/`
   - ❌ Need to audit ALL TUI screens for consistent styling
   - ❌ Ensure all headers use HeaderStyle
   - ❌ Ensure all cards use CardStyle
   - ❌ Ensure all buttons use ButtonStyle

2. **Enhanced Sparklines**
   - ⚠️ Basic sparkline exists: `internal/dashboard/sparkline.go`
   - ❌ Need enhanced visualization with trend colors
   - ❌ Need 14-day history tracking

**Files to Update:**
- All TUI screen files in `internal/tui/screens/`
- `internal/dashboard/sparkline.go`

---

## 📋 v0.0.19-beta - Complete Implementation Needed

### ❌ Phase 1: Unified TUI & AI Commands (0% Complete)

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
   - ❌ `create.go` - Template gallery (may exist as wizard)
   - ❌ `security.go` - Comprehensive security scan
   - ❌ `fix.go` - AI-powered auto-fix

3. **Deployment System** (`internal/deployment/`)
   - ❌ Directory exists but empty
   - ❌ Need support for: Vercel, Netlify, Railway, Render, Coolify, custom
   - ❌ Deployment wizard TUI

4. **Publisher System** (`internal/publisher/`)
   - ❌ Directory exists but empty
   - ❌ Need support for: npm, Homebrew, Scoop, Winget
   - ❌ Publishing wizard TUI

5. **Security System** (`internal/security/`)
   - ❌ Directory exists but empty
   - ❌ Need: npm audit, trufflehog, git-secrets, gosec, dive
   - ❌ Security scan TUI

6. **Fixer System** (`internal/fixer/`)
   - ❌ Directory exists but empty
   - ❌ Need AI-powered auto-fix logic
   - ❌ Fix wizard TUI

7. **Full TUI Menu Population**
   - ⚠️ Dashboard exists
   - ❌ Need all menu items: Run, Undo, Deploy, Publish, Security, Fix, etc.

**Estimated Effort:** 3-4 weeks

---

### ❌ Phase 2: Design System (DPR) Generation (0% Complete)

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

### ❌ Phase 3: Secrets & API Keys (RAKD/SOPS) (0% Complete)

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

### ❌ Phase 4: AI Agents System (0% Complete)

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

### ❌ Phase 5: Workflow Guidance Engine (0% Complete)

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

## 📊 Implementation Priority

### 🔴 Critical (Complete v0.0.18-beta)
1. **Documentation Structure** (Phase 4)
   - Update CONTEXT.md generator
   - Update README.md generator
   - **Effort:** 1-2 days

2. **GitHub Enforcement** (Phase 5)
   - Add requirement checks
   - Create setup wizard
   - **Effort:** 2-3 days

3. **TUI Polish** (Phase 6)
   - Audit all screens
   - Enhance sparklines
   - **Effort:** 3-5 days

### 🟡 High Priority (Start v0.0.19-beta)
1. **Basic TUI Actions** (Phase 1)
   - Run Dev Server
   - Undo Last Action
   - **Effort:** 1 week

2. **AI Agents** (Phase 4)
   - Generate all agent files
   - Create workflow rules
   - **Effort:** 2 weeks

3. **DPR System** (Phase 2)
   - Questionnaire TUI
   - Document generation
   - **Effort:** 2 weeks

### 🟢 Medium Priority (Complete v0.0.19-beta)
1. **Deployment System** (Phase 1)
2. **Publisher System** (Phase 1)
3. **Security System** (Phase 1)
4. **RAKD/SOPS** (Phase 3)
5. **Workflow Guidance** (Phase 5)

---

## 🎯 Recommended Next Steps

1. **Finish v0.0.18-beta polish** (1-2 weeks)
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
- v0.0.18-beta is close to completion
- v0.0.19-beta requires significant new development
- Focus on completing v0.0.18-beta before starting v0.0.19-beta

