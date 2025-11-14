# Show HN: DoPlan – Automate Your Project Workflow from Idea to Deployment

**Title:** Show HN: DoPlan – Automate your project workflow from idea to deployment

---

I've been building **DoPlan**, a CLI tool that automates project workflow from idea to deployment. It integrates with AI-powered IDEs (Cursor, Claude, Gemini, etc.) and helps you structure, document, and track your development projects.

## What It Does

**Key Features:**
- Break down ideas into phases and features automatically
- Generate PRD, API contracts, and project structure
- Visual dashboards with progress tracking
- GitHub automation (auto-branching, PR creation)
- Template system for customizing documents
- Checkpoint/Time Machine for project snapshots
- Fullscreen TUI with interactive dashboard

**What makes it different:**
- Works seamlessly with 6 different AI IDEs (Cursor, Claude CLI, Gemini CLI, Codex CLI, OpenCode, Qwen Code)
- Enforces best practices through automated rules
- Everything is generated and tracked automatically
- No manual project management overhead

## Built With

- **Go** (CLI)
- **Bubbletea** (TUI)
- Integrates with **GitHub CLI**

## Try It

```bash
# Install
brew install doplan  # or download from releases

# In your project
doplan install
# Select your IDE, then use commands like /Discuss, /Plan, /Dashboard
```

## Quick Demo

1. **Install in project:** `doplan install` (selects your IDE)
2. **In your IDE:** Type `/Discuss` to start refining your idea
3. **Generate docs:** `/Generate` creates PRD, API contracts, structure
4. **Create plan:** `/Plan` breaks down into phases and features
5. **View dashboard:** `/Dashboard` shows visual progress

## Links

- **GitHub:** https://github.com/DoPlan-dev/CLI
- **Demo:** [Add demo link/screenshot]
- **Docs:** https://doplan.dev (or GitHub Pages)

## What I'd Love Feedback On

- The workflow approach - does it match how you work?
- IDE integration experience - is it seamless?
- What features would be most useful?
- Any pain points in current project management?

## Technical Details

The tool uses Go's `text/template` for document generation, integrates with Git/GitHub for automation, and provides a fullscreen TUI built with Bubbletea. It generates markdown and HTML dashboards, manages project state in JSON, and supports custom templates.

Happy to answer questions and would love to hear your thoughts!

---

**Note:** This is ready to post on Hacker News. Just update the GitHub links and add a demo screenshot/video link.

