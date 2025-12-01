# Command Quick Reference

All DoPlan commands in one place. Quick lookup for syntax, options, and examples.

---

## 🎯 Workflow Commands

### `/hey`

**Purpose**: Onboarding, tutorial, welcome

**Syntax**:
```bash
/hey
/hey --project ./path
```

**Options**:
- `--project <path>` - Specify project path

**Examples**:
```bash
/hey
/hey --project ./my-project
```

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

---

### `/dev`

**Purpose**: Start development workflow

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

---

## ⚙️ System Commands

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

---

## 🔄 Common Workflows

### Complete Workflow

```bash
/hey
/do
/plan
/dev                    # Auto-detects completion
/sys status
/sys engagement
```

### Fast-Track

```bash
/do now --prompt "..."
/plan
/dev                    # Auto-detects completion
```

### Feature Addition

```bash
/do feature
/dev --feature "feature-name"  # Auto-detects completion
```

---

## 💡 Command Tips

- **Use `--project`** - Specify project path
- **Check `/sys status`** - Regular progress checks
- **View `/sys engagement`** - Stay motivated
- **`/dev` auto-detects completion** - No need to manually mark tasks done
- **Use fast-track** - When you have detailed prompt

---

**Next**: [File Structure](./02-File-Structure.md)

