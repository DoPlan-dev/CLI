tforce # /access

## Trigger
Exact match: `/access` or `/access <scope>`

Examples:
- `/access` → Prepare .do/system, .do/plan, and docs/
- `/access docs` → Patch only the docs/ folder
- `/access .do/system` → Patch only system files

## Action
When the user types `/access [scope]`:

1. **Explain Purpose**:
   - "This command fixes permissions so DoPlan can write reference files (user profile, quick reference, tutorial docs). The CLI already applies it after scaffolding a new project; run this if you moved/cloned the repo or see a permission warning."
2. **Run Access Patch**:
   - Execute `npx --yes @doplan-dev/cli goplan access <scope>` (defaults to `all`).
   - Scope options: `all`, `.do/system`, `.do/plan`, `docs`.
3. **Report Results**:
   - Summarize which directories/files were created or updated.
   - If the script fails, surface stderr and suggest re-running with `--debug` (env `DEBUG=1`).
4. **Next Steps**:
   - Remind the user to run `/hello` again once the patch succeeds.

## Agent Involvement
- **Project Orchestrator** – guides the user through the patch
- **Documentation Lead** – confirms docs/ references are ready

## Files Modified
- `.do/system/**`
- `.do/plan/**`
- `docs/**`

## Notes
- This command never overwrites existing content; it only creates missing folders/files and fixes permissions.
- Recommend running it immediately after cloning or when the CLI reports insufficient permissions.

