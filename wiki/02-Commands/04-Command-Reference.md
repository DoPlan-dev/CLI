# Command Reference

Complete reference for all DoPlan commands with syntax, options, and examples.

---

## Workflow Commands

### `/hey`

**Purpose**: Onboarding, tutorial, and welcome

**Syntax**:
```bash
/hey
/hey --project ./path
```

**Options**:
- `--project <path>` - Specify project path (default: current directory)

**Examples**:
```bash
/hey
/hey --project ./my-project
```

**Output**: Interactive tutorial or personalized greeting

---

### `/do`

**Purpose**: Capture idea and discovery

**Syntax**:
```bash
/do
/do feature
/do now [--prompt "text"] [--prd ./path/to/PRD.md]
/do i'm lucky
```

**Options**:
- `--project <path>` - Specify project path
- `--prompt <text>` - Fast-track with detailed prompt
- `--prd <path>` - Fast-track with PRD file

**Examples**:
```bash
/do
/do feature
/do now --prompt "Build a todo app with React"
/do now --prd ./requirements.md
/do i'm lucky
```

**Output**: Creates IDEA.md, BRAINSTORM.md, REFINEMENTS.md

---

### `/plan`

**Purpose**: Generate execution plan

**Syntax**:
```bash
/plan
/plan --project ./path
```

**Options**:
- `--project <path>` - Specify project path

**Examples**:
```bash
/plan
/plan --project ./my-project
```

**Output**: Creates TASKS.md and phase structure

---

### `/dev`

**Purpose**: Start development workflow and auto-complete the task (absorbs former `/done`)

**Syntax**:
```bash
/dev
/dev --feature <name>
/dev --project <path>
```

**Options**:
- `--feature <name>` - Feature name or task ID
- `--project <path>` - Specify project path

**Examples**:
```bash
/dev
/dev --feature "auth"
/dev --feature "1.2"
/dev --project ./my-project
```

**Output**: Sets up development environment, creates branch, tracks time, and automatically completes the task with commit/push when the agent decides it is done

---

## System Commands

### `/sys`

**Purpose**: System control panel

**Syntax**:
```bash
/sys
/sys <subcommand>
```

**Subcommands**:
- `engagement` - Engagement dashboard
- `role` - Role management
- `security` - Security settings
- `control` - System control

**Examples**:
```bash
/sys
/sys engagement
```

**Output**: System tree or subcommand output

---

### `/sys engagement`

**Purpose**: Engagement dashboard

**Syntax**:
```bash
/sys engagement
```

**Examples**:
```bash
/sys engagement
```

**Output**: Full engagement dashboard

---

### `/sys role`

**Purpose**: Role management

**Syntax**:
```bash
/sys role
/sys role tree
/sys role list
/sys role show <role>
/sys role assign <role>
```

**Examples**:
```bash
/sys role
/sys role tree
/sys role list
/sys role show developer
/sys role assign developer
```

**Output**: Role information or assignment confirmation

---

### `/sys security`

**Purpose**: Security settings

**Syntax**:
```bash
/sys security
/sys security status
/sys security test
/sys security release test
/sys security audit
```

**Examples**:
```bash
/sys security
/sys security status
/sys security test
/sys security release test
/sys security audit
```

**Output**: Security status or test results

---

### `/sys control`

**Purpose**: System control

**Syntax**:
```bash
/sys control
/sys control system on|off
/sys control agents on|off
/sys control roles on|off
```

**Examples**:
```bash
/sys control
/sys control system off
/sys control agents on
/sys control roles off
```

**Output**: Control confirmation or status

---

## Command Flags

### Common Flags

All commands support:
- `--project <path>` - Specify project path (default: current directory)

### Command-Specific Flags

- `/do now`:
  - `--prompt <text>` - Detailed prompt for fast-tracking
  - `--prd <path>` - PRD file path for fast-tracking

- `/dev`:
  - `--feature <name>` - Feature name or task ID

---

## Command Combinations

### Complete Workflow

```bash
/hey                    # Onboarding
/do                     # Capture idea
/plan                   # Generate plan
/dev                    # Start development (auto-completes task)
/sys engagement         # View achievements
```

### Fast-Track Workflow

```bash
/do now --prompt "..."  # Fast-track ideation
/plan                   # Generate plan
/dev                    # Start development (auto-completes task)
```

### Lucky Mode

```bash
/do i'm lucky           # Get AI suggestions
# → Choose idea
/plan                   # Generate plan
/dev                    # Start development (auto-completes task)
```

---

## Command Output Formats

### Success Output

Commands show:
- ✅ Success indicators
- 📋 Task information
- 🌿 Branch information
- 📚 Documentation status
- ⏱️ Duration (when applicable)
- 💡 Next steps

### Error Output

Commands show:
- ❌ Error indicators
- ⚠️ Warnings
- Error messages
- Suggestions for resolution

### Engagement Output

When achievements/challenges are earned:
- 🎉 Achievement notifications
- 🏆 Challenge completions
- 💰 Points earned
- ⭐ Rarity indicators

---

## Command Dependencies

### Required Files

- `/do` requires: (none - creates files)
- `/plan` requires: IDEA.md, BRAINSTORM.md
- `/dev` requires: TASKS.md

### Generated Files

- `/do` creates: IDEA.md, BRAINSTORM.md, REFINEMENTS.md
- `/plan` creates: TASKS.md, phase structure
- `/dev` creates: Git branch, updates active_state.json, auto-completes with commit/push and state snapshot

---

## Command Best Practices

### Workflow
1. Always start with `/hey` first time
2. Use `/do` for complete ideation
3. Run `/plan` after ideation
4. Use `/dev` to start tasks (auto-completes when finished)

### Efficiency
- Use `/do now` for fast-tracking
- Use `/do feature` for single features
- Use `/do i'm lucky` for inspiration
- Check `/sys engagement` for motivation

### Safety
- Verify branch before coding
- Check dependencies before completion
- Use `/sys control` carefully

---

## Troubleshooting

### Command Not Found
- Check installation: `doplan --version`
- Verify PATH includes DoPlan
- Try `npx @doplan-dev/cli` instead

### Missing Files
- `/plan` fails: Run `/do` first
- `/dev` fails: Run `/plan` first

### State Issues
- Check `.do/system/history/active_state.json`
- View state history: `.do/system/history/state-*.json`
- Use state snapshots for rollback

---

**Next**: [Engagement System](../03-Engagement-System/)

