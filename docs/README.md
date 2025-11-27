# Documentation Structure

This directory contains all project documentation organized by category.

## Directory Structure

```
Docs/
├── design/              # Design documents and specifications
│   ├── HELLO_COMMAND_DESIGN.md
│   ├── MEETING_COMMAND_EXPLANATION.md
│   ├── MEETING_SYSTEM_EXPLAINED.md
│   ├── PROJECT_TYPES_COMPREHENSIVE.md
│   └── QUESTION_VARIATION_SYSTEM.md
│
├── development/        # Development guides and proposals
│   ├── BUILD.md
│   ├── ITERATIVE_WORKFLOW.md
│   ├── LOCAL_TESTING.md
│   ├── POC_AGENTS_README.md
│   ├── PROJECT_TRACKER.md
│   ├── REFACTOR_PROPOSAL.md
│   ├── SCAN_REPORT_TEMPLATES.md
│   └── TESTING.md
│
├── foundation/         # Core foundational documentation
│   └── the-guide.md
│
├── history/            # Project history and metadata
│   ├── beginner-mode-roadmap.md
│   ├── branch-matrix.json
│   └── github-meta.json
│
├── reference/          # Quick reference guides
│   ├── AGENT_HIERARCHY.md
│   ├── AGENT_HIERARCHY_CHAT_PREVIEW.md
│   ├── COMMAND_EXAMPLES.md
│   └── QUICK_REFERENCE.md
│
├── release/            # Release documentation
│   ├── post-launch-monitoring.md
│   ├── RELEASE_CHECKLIST.md
│   ├── RELEASE_NOTES_v1.0.0.md
│   └── SOCIAL_MEDIA_ANNOUNCEMENT.md
│
├── reports/            # Test reports and coverage
│   ├── CHANGES_REPORT_POST_1.2.0.md
│   ├── COVERAGE_90_TARGET_REPORT.md
│   ├── COVERAGE_FINAL_REPORT.md
│   ├── COVERAGE_IMPROVEMENT_REPORT.md
│   ├── FULL_TEST_REPORT.md
│   ├── TEST_REPORT.md
│   └── TEST_SUMMARY.md
│
├── security/           # Security documentation
│   └── (security audit documents)
│
└── features/           # Feature documentation
    └── (feature specifications)
```

## Categories

### Design (`design/`)
Design documents, command specifications, and system explanations.

### Development (`development/`)
Development guides, build instructions, testing procedures, and refactoring proposals.

### Foundation (`foundation/`)
Core foundational documentation that serves as the base reference for the project.

### History (`history/`)
Project history, roadmaps, and metadata files (JSON) used by automation tools.

### Reference (`reference/`)
Quick reference guides for agents, commands, and common workflows.

### Release (`release/`)
Release checklists, release notes, and launch-related documentation.

### Reports (`reports/`)
Test reports, coverage reports, and change reports generated during development.

### Security (`security/`)
Security audits and security-related documentation.

### Features (`features/`)
Feature specifications and documentation.

## Root-Level Files

Some documentation files remain at the root of `Docs/` for easy access:
- `README.md` - This file
- `CONTRIBUTING.md` - Contribution guidelines
- `APP_IDEAS.md` - Application ideas
- `LAUNCH_CHECKLIST.md` - Launch preparation checklist
- `STANDUP.md` - Standup meeting notes
- `TUTORIAL_NOTES.md` - Tutorial documentation
- `WIKI_PLAN.md` - Wiki maintenance plan
- `wiki-tasks.md` - Wiki-related tasks
- `prompt.md` - Prompt documentation
- `ascii.md` - ASCII art documentation
- `DOCUMENTATION_REVIEW.md` - Documentation review notes
- `FEEDBACK_AND_SUGGESTIONS_POST_1.2.0.md` - Feedback collection

## Notes

- This structure follows the canonical `Docs/` pattern used in generated projects
- All documentation from the old lowercase `docs/` folder has been consolidated here
- Build artifacts and generated files are excluded (see `.gitignore`)

