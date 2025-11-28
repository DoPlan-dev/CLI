# System Control

The `/sys` command provides complete control over DoPlan's systems. Customize agents, roles, security, and engagement to match your needs.

---

## 🎯 System Control Overview

The `/sys` command is your control panel:
- **Engagement** - View dashboard and statistics
- **Role** - Manage roles and permissions
- **Security** - Security settings and tests
- **Control** - Enable/disable system features

---

## 📊 `/sys engagement` - Engagement Dashboard

### Purpose

View comprehensive engagement statistics and progress.

### Usage

```bash
/sys engagement
```

### What It Shows

- **Total Score** - Your current points
- **Achievements** - Count of achievements earned
- **Challenges** - Count of challenges completed
- **Relationship Level** - 0-100 relationship strength
- **Engagement Score** - 0-100% engagement
- **Last Reward** - Time since last achievement/challenge
- **Pending Rewards** - Rewards scheduled for release
- **Next Milestones** - Hints for upcoming achievements

### When to Use

- Check your progress
- See achievements earned
- View relationship level
- Get motivated
- Check next milestones

---

## 👤 `/sys role` - Role Management

### Purpose

Manage roles and permissions for your DoPlan system.

### Usage

```bash
/sys role                    # Show role tree
/sys role tree              # Show role hierarchy
/sys role list              # List all available roles
/sys role show <role>       # Show details for a role
/sys role assign <role>     # Assign a role to current user
```

### Role Tree

```bash
/sys role tree
```

**Shows**:
- Role hierarchy
- Available roles
- Role descriptions
- Permissions per role

### List Roles

```bash
/sys role list
```

**Shows**:
- All available roles
- Role descriptions
- Current assigned role

### Show Role Details

```bash
/sys role show developer
```

**Shows**:
- Role name and description
- Permissions
- Available commands
- Restrictions

### Assign Role

```bash
/sys role assign developer
```

**What happens**:
- Assigns role to current user
- Updates permissions
- Enables/disables features based on role
- Updates memory card

---

## 🔒 `/sys security` - Security Settings

### Purpose

Run security tests, audits, and checks.

### Usage

```bash
/sys security                # Show security status
/sys security status         # Show current security status
/sys security test           # Run standard security tests
/sys security release test    # Run pre-release security tests
/sys security audit           # Run comprehensive security audit
```

### Security Status

```bash
/sys security status
```

**Shows**:
- Current security status
- Last test date
- Security score
- Known vulnerabilities
- Recommendations

### Standard Tests

```bash
/sys security test
```

**Runs**:
- Dependency vulnerability scan
- Code security checks
- Configuration security
- Basic security validation

### Pre-Release Tests

```bash
/sys security release test
```

**Runs**:
- All standard tests
- Production security checks
- Deployment security
- Release readiness validation

### Security Audit

```bash
/sys security audit
```

**Runs**:
- Comprehensive security review
- All security tests
- Deep vulnerability scan
- Security best practices check
- Detailed report generation

---

## ⚙️ `/sys control` - System Control

### Purpose

Control DoPlan's system features: enable/disable agents, roles, or the entire system.

### Usage

```bash
/sys control                 # Show control options
/sys control system on|off   # Global kill switch
/sys control agents on|off   # Enable/disable agents
/sys control roles on|off    # Enable/disable roles
```

### Global Kill Switch

```bash
/sys control system off
```

**What happens**:
- **Requires strong confirmation** - Multiple confirmations needed
- Disables entire DoPlan system
- All commands become unavailable
- State is preserved
- Can be re-enabled with `system on`

**⚠️ Warning**: This is a powerful command. Use with caution!

**Confirmation required**:
```
⚠️  WARNING: This will disable the entire DoPlan system!
    All commands will become unavailable.
    This action requires strong confirmation.
    
    Type 'DISABLE DOPLAN SYSTEM' to confirm:
```

### Control Agents

```bash
/sys control agents off
```

**What happens**:
- Disables all AI agents
- Commands still work
- Agent responses disabled
- Can be re-enabled with `agents on`

### Control Roles

```bash
/sys control roles off
```

**What happens**:
- Disables role system
- All features available
- Role restrictions removed
- Can be re-enabled with `roles on`

---

## 💡 System Control Tips

### Best Practices

- **Check engagement regularly** - Stay motivated
- **Use roles appropriately** - Manage permissions
- **Run security tests** - Before releases
- **Use control carefully** - Powerful commands
- **Review system tree** - Discover features

### Power User

- **Customize roles** - For your workflow
- **Run security audits** - Regular checks
- **Monitor engagement** - Track metrics
- **Use system control** - For testing
- **Explore all subcommands** - Full features

### Safety

- **Global kill switch requires confirmation** - Safety first
- **State is preserved** - When disabling
- **Can always re-enable** - No permanent changes
- **No data loss** - From control commands

---

**Next**: [Customization](./02-Customization.md)

