# `/meeting` Command Guide

The `/meeting` command runs an adaptive discovery interview that captures everything your AI agency needs before `/write`. It tailors its depth, questions, and content planning to your project type and timeline.

---

## Overview

| Property | Details |
| --- | --- |
| Purpose | Capture project vision, audience, UX, content, marketing, and operational requirements. |
| Prerequisite | Run `/tell` so `.plan/00_System/IDEA.md` exists. |
| Output | `.plan/00_System/BRAINSTORM.md`, `.do/system/content/` directory, updated `active_state.json`. |
| Time | 5–60 minutes depending on the speed you choose. |

---

## Step-by-Step Flow

1. **Project Type Detection**  
   - Reads IDEA.md to determine whether you are building a website/agency project or a SaaS/startup product.  
   - Influences suggested speeds, question depth, and content defaults.

2. **Speed Selection**  
   - Offers four meeting speeds (Quick Start, Standard, Comprehensive, Deep Dive).  
   - Each speed maps to a set of phases (see table below).  
   - Suggested defaults are shown based on detected project type.

3. **GitHub & Automation Check**  
   - Asks if a repository already exists.  
   - If provided, validates the repo, detects existing `.github/workflows/`, and flags missing automation so `/github` or `/branchci` can fix it later.

4. **Content Planning Module**  
   - Asks whether you need AI-generated content.  
   - Lets you pick content categories (app pages, legal, blog, documentation, marketing, email, SEO).  
   - For each category, choose **Full AI**, **Keyword-guided**, or **Hybrid** generation.

5. **Phase-by-Phase Interview**  
   - Loads templates from `.do/core/brainstorm/phase-*.md`.  
   - Generates questions dynamically based on project type, speed, and your earlier answers.  
   - Waits for confirmation before moving to the next phase and probes for clarity when needed.

6. **Summary & Confirmation**  
   - Compiles everything using `CONFIRMATION_TEMPLATE.md`.  
   - Displays a review panel with options to save, revise a phase, add more info, or restart.

7. **Save & State Update**  
   - Writes the approved summary to `.plan/00_System/BRAINSTORM.md`.  
   - Creates `.do/system/content/` subfolders for each requested content type.  
   - Sets `phase: "brainstorm"` inside `.plan/history/active_state.json`.

8. **Next-Step Guidance**  
   - Reminds you to run `/write` (documents), `/content` (copy generation), and `/plan` (task synthesis).

---

## Meeting Speeds

| Speed | Duration | Phases Covered | Best For |
| --- | --- | --- | --- |
| Quick Start | ~5–10 min | 01 (Vision), 03 (Experience) | Simple websites, quick demos |
| Standard | ~15–20 min | 01, 02, 03, 06 | Company sites, agency work, MVPs |
| Comprehensive | ~30–45 min | All six phases (condensed) | Complex apps, growing startups |
| Deep Dive | 60+ min | All six phases (full) | SaaS, enterprise, regulated products |

> Phases: 01 Vision & Outcomes, 02 Audience & Differentiation, 03 Experience & Tech, 04 Content & SEO, 05 Marketing & Growth, 06 Delivery & Operations.

---

## Content Creation Matrix

| Project Type | Typical Content Types | Options per Type |
| --- | --- | --- |
| Website / Agency / Personal | App pages, legal, social media, blog, SEO | Full AI • Keyword-guided • Hybrid |
| SaaS / Startup | App pages, legal, social media, blog, documentation, marketing, email, SEO | Full AI • Keyword-guided • Hybrid |

Every selection records tone, audience, keywords, and delivery expectations so `/content` can pick the right template later.

---

## Files & Structures Produced

| Artifact | Location | Notes |
| --- | --- | --- |
| Meeting summary | `.plan/00_System/BRAINSTORM.md` | Organized by phases with checklists and follow-ups. |
| Content folders | `.do/system/content/` | Subdirectories (`app-pages/`, `legal/`, `blog/`, etc.) plus README files. |
| History entry | `.plan/history/state-*.json` | Snapshot showing phase advanced to `brainstorm`. |
| Session metadata | `.do/system/meeting_session.json` | Tracks speed, project type, GitHub info (used for resumes/resume). |

---

## Best Practices

1. **Pick the right speed** – Quick Start is tempting, but Comprehensive/Deep Dive prevent costly gaps for SaaS.  
2. **Be explicit about content** – Saying “yes” without selecting content types gives you empty folders; check the boxes you truly need.  
3. **Use revision mode** – If a phase answer feels thin, choose “Revise phase” instead of editing BRAINSTORM.md by hand.  
4. **Link GitHub early** – `/meeting` can flag missing workflows so `/github` or `/branchci` handle them before `/build`.  
5. **Review before saving** – Nothing is written until you press “Save it,” so take advantage of the confirmation screen.

---

## Troubleshooting

| Issue | Fix |
| --- | --- |
| Meeting keeps suggesting the wrong speed | Update IDEA.md with clearer keywords (e.g., “SaaS billing platform”) and rerun `/meeting`. |
| Content folders missing | Re-run `/meeting`, choose “Add information,” and enable the needed content categories. |
| Need to change a single phase after saving | Use `/change` on BRAINSTORM.md or rerun `/meeting` and pick “Revise phase”. |
| GitHub check fails | Confirm the repo URL, ensure you have permissions, or skip and set up workflows later with `/github ci`. |

---

## Workflow Placement

1. `/tell` – Capture the idea.  
2. `/meeting` – Run the adaptive discovery session (this page).  
3. `/write` – Generate PRD/Architecture/Design System based on IDEA + BRAINSTORM.  
4. `/good` – Approve the plan once documents look right.  
5. `/plan` – Turn the approved plan into TASKS.md.

---

## Related Pages

- [Commands Reference](Commands) – High-level registry of all commands.  
- [Workflow Guide](Workflow) – Shows where `/meeting` fits in the Capture & Align phase.  
- [Hello Tutorial](Hello-Tutorial) – Optional onboarding run before `/tell`.  
- [Project Structure](Project-Structure) – Understand where meeting outputs live in the repo.  
- [Quick Start](Quick-Start) – For generating your first project before running `/meeting`.

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

