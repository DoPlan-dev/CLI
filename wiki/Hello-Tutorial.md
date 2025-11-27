# `/hello` Tutorial Guide

The `/hello` command is your optional onboarding walkthrough. Run it once to meet the agent hierarchy, capture your profile, and stash a personal quick-reference pack before you start `/tell`.

---

## Why `/hello` Matters

- **Personalized setup** – Captures your name, experience level, and preferred coaching style.
- **Context sharing** – Stores answers for later use by `/meeting`, `/write`, and other commands.
- **Reference kit** – Generates ready-to-read docs (`QUICK_REFERENCE.md`, `AGENT_HIERARCHY.md`, etc.).
- **Safe practice** – Lets you test-drive commands without modifying your project.

---

## When to Run It

| Scenario | Why it helps |
| --- | --- |
| First time using DoPlan CLI | Unlocks guided tutorial, agent overview, and personal profile files. |
| New team member onboarding | Gives an instant tour before they touch `/tell`. |
| Switching machines/IDEs | Rebuilds the reference pack if you wiped `.do/system/`. |
| Teaching sessions/demos | Use the practice flow to show commands without changing state. |

---

## Walkthrough Flow

1. **Warm Welcome & Profile Capture**  
   - Asks for your name (`user_name`) and experience (`user_experience_level`).  
   - Responses are reused later by `/meeting` and progress reports.

2. **System & Agent Overview**  
   - Explains what DoPlan automates.  
   - Prints the 18-agent hierarchy (Project Orchestrator, leads, QA, Growth, Docs, etc.).

3. **Support Mode Selection**  
   - Choose **Guided Mode** (check-ins after each `/finished`) or **Independent Mode** (silent partner).  
   - Saved as `development_support_mode`; you can change it later via `/hello` re-run or manual edit.

4. **Command Walkthrough**  
   - Highlights the core workflow: `/tell`, `/improve`, `/write`, `/change`, `/plan`, `/build`, `/content`.  
   - Explains what each command saves and how agents collaborate.

5. **Optional Test Drive**  
   - Practice `/tell`, `/meeting`, `/write`, or `/content` without touching real files.  
   - `/hello` narrates what *would* happen (IDEA saved, BRAINSTORM updates, documents generated) but discards the dry-run data.

6. **Personalized Tips**  
   - Shares guidance aligned with your experience level (e.g., reminders for beginners, autonomy notes for advanced users).

7. **Reference Pack Creation**  
   - Writes cheat sheets and hierarchy diagrams so you can revisit the tutorial without rerunning `/hello`.

8. **Next Steps Prompt**  
   - Ends by nudging you toward `/tell` to capture the real project idea.

---

## Files & Settings Created

| Artifact | Path | Description |
| --- | --- | --- |
| User profile | `.do/system/user_profile.json` | Stores `user_name`, `user_experience_level`, `development_support_mode`, and tutorial completion flags. |
| Quick reference | `.do/system/QUICK_REFERENCE.md` & `docs/references/QUICK_REFERENCE.md` | Cheat sheet summarizing commands and shortcuts. |
| Agent hierarchy | `docs/overview/AGENT_HIERARCHY.md` | Text-based tree of all 18 agents. |
| Command examples | `docs/references/COMMAND_EXAMPLES.md` | Sample prompts and expected outputs for core commands. |
| Tutorial notes | `docs/tutorials/TUTORIAL_NOTES.md` | Transcript of the onboarding session. |

> Already have these files? Re-running `/hello` overwrites them with your latest answers, which is useful when onboarding a new teammate on the same repo.

---

## Tips & Troubleshooting

- **Forgot your support mode?** Check `.do/system/user_profile.json` or rerun `/hello` to switch between guided and independent.
- **Need to demo commands safely?** Use the test-drive prompts so no real files change.
- **Multiple users on one repo?** Each person can copy `.do/system/QUICK_REFERENCE.md` into their IDE or rerun `/hello` after cloning.
- **Stuck during onboarding?** Delete `.do/system/user_profile.json` and run `/hello` again to restart the flow.

---

## Related Pages

- [Quick Start](Quick-Start) – Generate your first project.  
- [First Project Tutorial](First-Project-Tutorial) – Walkthrough after onboarding.  
- [Meeting Command](Meeting-Command) – Deep dive into the adaptive discovery session referenced during `/hello`.  
- [Commands Reference](Commands) – Full list of day-to-day commands.  
- [Workflow Guide](Workflow) – See how `/hello` feeds into the Capture & Align phase.

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

