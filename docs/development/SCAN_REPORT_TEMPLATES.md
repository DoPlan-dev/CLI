## Scan Report Templates & Extensions

DoPlan ships with a flexible scan report engine powered by `scripts/scanreport/main.go`. Reports can be tailored per-project using presets, custom section ordering, and your own extensions.

### 1. Presets

Out of the box you can run:

```
/report                             # standard preset
/report --preset exec               # condensed executive summary
/report --preset detailed           # full detail + dependency audit
```

Presets control which sections render (executive summary, findings, progress, visuals, dependency audit, etc.). See `presetConfigs` in `scripts/scanreport/main.go` for the exact ordering.

### 2. Per-project Configuration (`.plan/reports/config.json`)

If you want defaults for a specific project, create `.plan/reports/config.json`:

```jsonc
{
  "preset": "exec",
  "sections": [
    "executive",
    "progress",
    "visuals",
    "state",
    "feedback"
  ]
}
```

- `preset` matches the CLI presets (`standard`, `exec`, `detailed`).  
- `sections` is optional; when provided it overrides the preset ordering. Valid entries:
  - `executive`
  - `findings`
  - `next`
  - `feedback`
  - `state`
  - `progress`
  - `visuals`
  - `dependency`

CLI flags always win over config values (e.g., `/report --preset detailed` ignores `preset` in config for that run).

### 3. Where to Customize

- **Visuals** come from `renderVisuals()` inside `scripts/scanreport/main.go`. Feel free to change the ASCII bars, add emoji badges, or pull metrics from other tools.
- **Dependency Audit** sources data via `collectPackageJSONDeps()` and `collectGoModDeps()`. Extend these functions (or add new ones) to support other ecosystems.
- **Progress Section** is shared via `internal/progress`. If you add fields (e.g., test coverage), expose them through the `progress.Report` struct and reference them in `summarizeProgress()`.

### 4. Adding New Sections

1. Extend the `validSections` map in `scripts/scanreport/main.go`.
2. Update `diffSections` to include the new data.
3. Populate the section in `main()` alongside the existing state/progress/visuals/dependency fields.
4. Handle rendering inside `buildDiff()` and `renderSection()`.

### 5. Tips

- Keep sections concise; link to deeper docs when needed.
- Regenerate reports after tweaking config to ensure markdown matches expectations.
- Commit `.plan/reports/config.json` so future scans stay consistent across teammates and CI.

