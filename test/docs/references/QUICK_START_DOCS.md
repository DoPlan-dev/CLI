# Quick Start: Using @docs in Cursor

This guide helps you quickly set up and use Cursor's @docs feature with this project.

<!-- Generated: 2025-01-27 -->

## Setup (One-Time)

1. **Open Cursor Settings**
   - Go to Settings → Features → Docs

2. **Add Documentation Sources**
   Add these paths:
   - `docs/` - Entire documentation directory
   - `docs/overview/` - Project overview and guides
   - `docs/history/` - Changelog and historical records

3. **Verify Setup**
   - Type `@docs` in chat to see available sources
   - You should see the added directories listed

## Quick Usage Examples

### Reference Entire Documentation
```
@docs How do I get started with this project?
```

### Reference Specific Category
```
@docs/overview What commands are available?
```

### Reference Specific File
```
@docs/overview/README.md Show me the project structure
```

### Check Project History
```
@docs/history/CHANGELOG.md What changed in recent versions?
```

## Common Questions

### "How do I start working on this project?"
```
@docs/overview/README.md How do I get started?
```

### "What commands can I use?"
```
@docs/overview/README.md What commands are available?
```

### "Where should I add new documentation?"
```
@docs/README.md Where should I place new documentation?
```

### "What's the project structure?"
```
@docs/overview/README.md What is the project structure?
```

## Tips

1. **Use @docs first**: Start with `@docs` to get an overview
2. **Be specific**: Use `@docs/folder` or `@docs/file.md` for targeted questions
3. **Combine with commands**: Use @docs alongside project commands like `/plan`, `/build`
4. **Keep docs updated**: Ask "Are any docs outdated?" to check freshness

## Documentation Locations

- **Project Overview**: `docs/overview/README.md`
- **Documentation Index**: `docs/DOCS_INDEX.md`
- **Changelog**: `docs/history/CHANGELOG.md`
- **This Guide**: `docs/references/QUICK_START_DOCS.md`

---

**Need more help?** See [DOCS_INDEX.md](../DOCS_INDEX.md) for comprehensive documentation guide.

