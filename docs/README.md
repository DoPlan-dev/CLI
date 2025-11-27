# Documentation Structure

This directory contains all project documentation organized by category.

## Directory Structure

```
docs/
├── design/              # Design documents and specifications
│   ├── ascii.md
│   ├── HELLO_COMMAND_DESIGN.md
│   ├── MEETING_COMMAND_EXPLANATION.md
│   ├── MEETING_SYSTEM_EXPLAINED.md
│   ├── PROJECT_TYPES_COMPREHENSIVE.md
│   └── QUESTION_VARIATION_SYSTEM.md
│
├── development/        # Development guides and proposals
│   ├── BUILD.md
│   ├── CONTRIBUTING.md
│   ├── DOCUMENTATION_REVIEW.md
│   ├── ITERATIVE_WORKFLOW.md
│   ├── LOCAL_TESTING.md
│   ├── POC_AGENTS_README.md
│   ├── PROJECT_TRACKER.md
│   ├── REFACTOR_PROPOSAL.md
│   ├── SCAN_REPORT_TEMPLATES.md
│   ├── STANDUP.md
│   ├── TESTING.md
│   ├── TUTORIAL_NOTES.md
│   ├── WIKI_PLAN.md
│   └── wiki-tasks.md
│
├── features/           # Feature documentation and ideas
│   └── APP_IDEAS.md
│
├── foundation/         # Core foundational documentation
│   ├── prompt.md
│   └── the-guide.md
│
├── history/            # Project history and metadata
│   ├── beginner-mode-roadmap.md
│   ├── branch-matrix.json
│   ├── github-meta.json
│   └── README_OLD.md
│
├── reference/          # Quick reference guides
│   ├── AGENT_HIERARCHY.md
│   ├── AGENT_HIERARCHY_CHAT_PREVIEW.md
│   ├── COMMAND_EXAMPLES.md
│   └── QUICK_REFERENCE.md
│
├── release/            # Release documentation
│   ├── LAUNCH_CHECKLIST.md
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
│   ├── FEEDBACK_AND_SUGGESTIONS_POST_1.2.0.md
│   ├── FULL_TEST_REPORT.md
│   ├── TEST_REPORT.md
│   └── TEST_SUMMARY.md
│
└── security/           # Security documentation
    └── (security audit documents)
```

## Categories

### Design (`design/`)
Design documents, command specifications, system explanations, and ASCII art assets.

### Development (`development/`)
Development guides, build instructions, testing procedures, refactoring proposals, contribution guidelines, documentation reviews, wiki planning, standup notes, and tutorial documentation.

### Features (`features/`)
Feature specifications, documentation, and application ideas.

### Foundation (`foundation/`)
Core foundational documentation that serves as the base reference for the project, including the core project prompt and guide.

### History (`history/`)
Project history, roadmaps, metadata files (JSON) used by automation tools, and archived documentation.

### Reference (`reference/`)
Quick reference guides for agents, commands, and common workflows.

### Release (`release/`)
Release checklists, release notes, launch checklists, and launch-related documentation.

### Reports (`reports/`)
Test reports, coverage reports, change reports, and feedback reports generated during development.

### Security (`security/`)
Security audits and security-related documentation.

## Notes

- This structure follows the canonical `docs/` pattern used in generated projects
- All documentation is organized into appropriate categories - no files remain in the root except `README.md`
- All documentation from the old lowercase `docs/` folder has been consolidated here
- Build artifacts and generated files are excluded (see `.gitignore`)

