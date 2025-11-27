# Refactoring Proposal: File-Based Storage for Content

## Overview
Move hardcoded content (meetings, hello tutorials, templates, agents, commands) from Go code to file-based storage similar to the rules library pattern. This will significantly improve testability, maintainability, and generator coverage.

## Current State

### Problems:
1. **Large string constants** in code (hello/meeting tutorials are 300+ lines)
2. **Hard to maintain** - content changes require code changes
3. **Hard to test** - file I/O operations not testable
4. **Low generator coverage** - template/parsing logic untestable
5. **No versioning** - can't track content changes independently

### Current Structure:
```
internal/generator/
  ├── agents.go         # GetAllAgents() - 342 lines of Agent structs
  ├── commands.go       # GetAllCommands() - 686+ lines of Command structs + huge Action strings
  ├── plan.go           # Template strings embedded in functions
  └── docs.go           # Template strings embedded in functions
```

## Proposed Structure

Following the `internal/rules/library/` pattern:

```
internal/
├── rules/
│   └── library/          # ✅ Already exists (good pattern)
│       ├── 01-core-workflow/
│       ├── 02-ai-agents/
│       └── ...
│
└── content/              # 🆕 New directory
    ├── agents/           # Agent definitions
    │   ├── leadership/
    │   │   ├── project_orchestrator.json
    │   │   └── ...
    │   ├── engineering/
    │   │   ├── engineering_lead.json
    │   │   └── ...
    │   └── ...
    │
    ├── commands/         # Command definitions
    │   ├── core/
    │   │   ├── hello.json
    │   │   ├── tell.json
    │   │   ├── meeting.json
    │   │   └── ...
    │   └── squad/
    │       └── ...
    │
    ├── tutorials/        # Tutorial content
    │   ├── hello/
    │   │   ├── main.md        # Main hello flow
    │   │   ├── meeting.md     # /hello meeting tutorial
    │   │   ├── plan.md        # /hello plan tutorial
    │   │   ├── build.md       # /hello build tutorial
    │   │   └── github.md      # /hello github tutorial
    │   └── meeting/
    │       ├── phase-01.md    # Phase 1 questions
    │       ├── phase-02.md    # Phase 2 questions
    │       ├── ...
    │       └── confirmation-template.md
    │
    └── templates/        # Document templates
        ├── agents/
        │   └── agent.md       # Agent markdown template
        ├── commands/
        │   └── command.md     # Command markdown template
        ├── documents/
        │   ├── idea.md
        │   ├── brainstorm.md
        │   ├── prd.md
        │   ├── architecture.md
        │   └── ...
        └── workflows/
            ├── ci.yml
            ├── release.yml
            └── ...
```

## Benefits for Generator Coverage

### Current Coverage Issues:
1. **GetAllAgents()** - 100% coverage but hardcoded data
2. **GetAllCommands()** - Contains huge Action strings (hard to test all paths)
3. **Template parsing** - Limited testability
4. **File I/O operations** - Not testable when content is in code

### After Refactoring:
1. **File reading/parsing** - Testable (file I/O operations)
2. **Template loading** - Testable (can test with mock files)
3. **Error handling** - Testable (missing files, invalid JSON, etc.)
4. **Content validation** - Testable (schema validation)
5. **Expected coverage increase: 82.5% → 90%+**

## Implementation Plan

### Phase 1: Create Content Structure
1. Create `internal/content/` directory
2. Move agent definitions to JSON files
3. Move command definitions to JSON files
4. Extract tutorials to markdown files
5. Extract templates to template files

### Phase 2: Update Generator Code
1. Replace `GetAllAgents()` with `LoadAgents(dir string)`
2. Replace `GetAllCommands()` with `LoadCommands(dir string)`
3. Add `LoadTutorial(name string)` function
4. Add `LoadTemplate(name string)` function
5. Update generators to use file-based loading

### Phase 3: Embed Content (like rules)
1. Use `//go:embed` to embed content directory
2. Extract content during generation (like ExtractRules)
3. Store in `.do/core/content/` (central location)
4. Create symlinks to IDE locations

### Phase 4: Testing
1. Test file loading functions
2. Test error handling (missing files, invalid JSON)
3. Test template rendering with file-based templates
4. Test content extraction
5. Expected coverage: **90%+**

## Example File Formats

### Agent JSON (agents/leadership/project_orchestrator.json)
```json
{
  "name": "Project Orchestrator",
  "role": "CEO / Engineering Manager",
  "systemPrompt": "You are the Project Orchestrator...",
  "reportsTo": "",
  "manages": [
    "Product Manager",
    "Engineering Lead",
    "Design & UX Manager",
    "QA & Reliability Manager",
    "Release & Growth Manager",
    "Documentation Lead"
  ],
  "responsibilities": [
    "Ultimate decision maker",
    "Resource allocation",
    "Team coordination",
    "Strategic vision",
    "Ensure project meets all success criteria"
  ],
  "fileName": "project_orchestrator.md",
  "category": "leadership"
}
```

### Command JSON (commands/core/hello.json)
```json
{
  "name": "hello",
  "category": "core",
  "trigger": "/hello [<subcommand>]",
  "description": "Welcome, tutorial, and command introductions",
  "action": "@tutorials/hello/main.md",
  "agentInvolvement": ["Documentation Lead"],
  "filesRead": [".do/system/user_profile.json"],
  "filesModified": [
    ".do/system/user_profile.json",
    ".do/system/QUICK_REFERENCE.md",
    "docs/references/QUICK_REFERENCE.md",
    "docs/overview/AGENT_HIERARCHY.md",
    "docs/references/COMMAND_EXAMPLES.md",
    "docs/tutorials/TUTORIAL_NOTES.md"
  ],
  "examples": ["/hello", "/hello meeting", "/hello plan"],
  "requirements": "First-time setup for new users"
}
```

### Tutorial Markdown (tutorials/hello/main.md)
```markdown
# Hello Tutorial - Main Flow

## Step 1: Check First Time
Read .do/system/user_profile.json. If "first_hello_completed" is true, show welcome back message.

## Step 2: Warm Greeting
Display: "Hello! 👋 I'm DoPlan..."

[... rest of tutorial content ...]
```

## Testing Improvements

### New Testable Functions:
```go
// Test file loading
func TestLoadAgents(t *testing.T)
func TestLoadCommands(t *testing.T)
func TestLoadTutorial(t *testing.T)
func TestLoadTemplate(t *testing.T)

// Test error handling
func TestLoadAgents_InvalidJSON(t *testing.T)
func TestLoadAgents_MissingFile(t *testing.T)
func TestLoadTutorial_NotFound(t *testing.T)

// Test content extraction (like ExtractRules)
func TestExtractContent(t *testing.T)
func TestExtractContent_ErrorPaths(t *testing.T)

// Test template rendering with file-based templates
func TestRenderAgentTemplate_FromFile(t *testing.T)
func TestRenderCommandTemplate_FromFile(t *testing.T)
```

## Migration Strategy

### Backward Compatibility:
1. Keep existing functions as wrappers
2. Gradually migrate to file-based loading
3. Add feature flag: `USE_FILE_BASED_CONTENT=true`
4. Test both paths initially

### Example Migration:
```go
// Old (keep for backward compatibility)
func GetAllAgents() []Agent {
    // Load from files if available, fallback to hardcoded
    if agents, err := LoadAgentsFromFiles(); err == nil {
        return agents
    }
    return getAllAgentsHardcoded() // Fallback
}

// New (preferred)
func LoadAgents(contentDir string) ([]Agent, error) {
    // Load from files, testable!
}
```

## Expected Coverage Impact

| Component | Current | After | Improvement |
|-----------|---------|-------|-------------|
| Agents Generator | 77.8% | 95%+ | +17.2% |
| Commands Generator | 77.8% | 95%+ | +17.2% |
| Template Rendering | 63.6% | 90%+ | +26.4% |
| Overall Generator | 82.5% | 90%+ | +7.5% |

## Next Steps

1. **Create content directory structure**
2. **Migrate one component at a time** (start with agents)
3. **Add tests for file loading**
4. **Measure coverage improvement**
5. **Repeat for commands, tutorials, templates**

This refactoring will significantly improve:
- ✅ Test coverage (file I/O becomes testable)
- ✅ Maintainability (content separate from code)
- ✅ Flexibility (easy to update content)
- ✅ Versioning (content changes tracked separately)
- ✅ External contributions (non-developers can update content)

