---
name: access
category: tools
trigger: "/access [<scope>]"
description: "Beginner-friendly patch to fix .do/ and docs/ permissions"
agentInvolvement:
  - Project Orchestrator
  - Documentation Lead
filesModified:
  - ".do/system/**"
  - ".do/plan/**"
  - "docs/**"
examples:
  - "/access"
  - "/access .do/system"
---

When user types /access or /access <scope>:

1. **Explain Purpose**:
   - "This patch makes sure DoPlan can write reference docs and state files. The CLI already runs it right after scaffolding a project, but you can re-run it anytime (e.g., after moving/cloning the repo). It only creates missing folders/files and fixes permissions."

2. **Run Helper**:
   - Execute `npx --yes @doplan-dev/cli goplan access <scope>` (defaults to `all`).
   - Scope options:
     * `all` - patch `.do/system`, `.do/plan`, and `docs/`
     * `.do/system` or `system`
     * `.do/plan` or `plan`
     * `docs`

3. **Report Results**:
   - Show which folders/files were created or updated.
   - If script exits non-zero, surface stderr and prompt the user to re-run with `DEBUG=1` for verbose output.

4. **Next Steps**:
   - Encourage the user to run `/hello` again once the patch succeeds.
   - Mention that the patch is safe to re-run at any time.

