# Agents File-Based Storage - Proof of Concept

## Overview

This POC demonstrates externalizing agent definitions from hardcoded Go structs to JSON files, similar to the existing rules library pattern. This improves testability, maintainability, and generator coverage.

## Structure

```
internal/content/
└── agents/                    # Embedded agent definitions (JSON)
    ├── leadership/
    │   └── project_orchestrator.json
    ├── engineering/
    │   ├── engineering_lead.json
    │   ├── system_architect.json
    │   ├── frontend_lead.json
    │   ├── backend_lead.json
    │   ├── devops_engineer.json
    │   ├── security_lead.json
    │   └── performance_engineer.json
    ├── product/
    │   └── product_manager.json
    ├── design/
    │   ├── design_manager.json
    │   └── ui_ux_designer.json
    ├── quality/
    │   ├── qa_manager.json
    │   └── qa_engineer.json
    ├── release/
    │   ├── release_manager.json
    │   ├── release_captain.json
    │   └── growth_coach.json
    └── documentation/
        ├── documentation_lead.json
        └── documentation_writer.json
```

## Implementation

### Files Created

1. **`internal/content/content.go`**: Embed system for agents (similar to `internal/rules/rules.go`)
2. **`internal/generator/agents_filebased.go`**: Loader functions that read from JSON files
3. **`internal/generator/agents_filebased_test.go`**: Comprehensive test suite

### Key Functions

- `LoadAgentsFromFiles()`: Loads agents from embedded JSON files
- `LoadAgentsFromDirectory(dir)`: Loads agents from a directory (useful for testing)
- `ExtractAgents(targetDir)`: Extracts embedded agents to disk (for inspection/testing)
- `GetAllAgentsFileBased()`: Loads from files with fallback to hardcoded

## JSON Format

Each agent JSON file contains:
```json
{
  "name": "Project Orchestrator",
  "role": "CEO / Engineering Manager",
  "systemPrompt": "...",
  "reportsTo": "",
  "manages": ["Product Manager", "Engineering Lead"],
  "responsibilities": ["Ultimate decision maker", ...],
  "fileName": "project_orchestrator.md",
  "category": "leadership"
}
```

## Test Coverage

The POC includes comprehensive tests:

- ✅ Loading from embedded files
- ✅ Loading from directory
- ✅ Extracting to disk
- ✅ Error handling (invalid paths, invalid JSON)
- ✅ Content validation
- ✅ File content verification

**All tests pass** ✅

## Benefits

### 1. **Testability** (Major Coverage Gain)
- Can test file I/O operations
- Can test error paths (missing files, invalid JSON, permission errors)
- Can test parsing and validation logic
- Can mock filesystem operations

### 2. **Maintainability**
- Non-developers can edit agent definitions
- Content changes don't require code changes
- Easier versioning and diff tracking

### 3. **Generator Coverage Impact**
- **Current**: File I/O operations are hardcoded, limited test paths
- **After**: Testable file operations, error handling, parsing logic
- **Estimated gain**: +5-10% coverage for `AgentsGenerator`

## Integration

The POC maintains **backward compatibility**:
- Existing `GetAllAgents()` still works (hardcoded fallback)
- `GetAllAgentsFileBased()` tries files first, falls back if needed
- Can migrate gradually without breaking existing code

## Next Steps

1. **Update AgentsGenerator** to use `GetAllAgentsFileBased()` instead of `GetAllAgents()`
2. **Add error path tests** for file-based loading in generator tests
3. **Extend to Commands**: Similar structure for commands (JSON files)
4. **Extend to Tutorials**: Markdown files for command tutorials
5. **Extend to Templates**: Template files for agents, commands, documents

## Coverage Impact Estimate

| Component | Current | After POC | After Full Implementation |
|-----------|---------|-----------|---------------------------|
| AgentsGenerator | ~77.8% | ~85% | ~95% |
| File I/O Operations | 0% | ~90% | ~95% |
| Error Handling | ~60% | ~85% | ~95% |

## Example Usage

```go
// Load from embedded files
agents, err := LoadAgentsFromFiles()
if err != nil {
    // Fallback to hardcoded
    agents = GetAllAgents()
}

// Load from directory (testing)
agents, err := LoadAgentsFromDirectory("/path/to/agents")

// Extract for inspection
err := ExtractAgents("/tmp/agents")
```

## Migration Path

1. ✅ **Phase 1 (POC)**: File-based loading with fallback (this POC)
2. **Phase 2**: Update `AgentsGenerator.Generate()` to use file-based
3. **Phase 3**: Remove hardcoded `GetAllAgents()` (optional)
4. **Phase 4**: Apply same pattern to Commands, Tutorials, Templates

---

**Status**: ✅ POC Complete - All tests passing, ready for integration

