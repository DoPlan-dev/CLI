# System Commands

The `/sys` command provides access to DoPlan's system settings, controls, and information. It acts as a control panel for configuration and engagement systems.

---

## `/sys` - System Control Panel

### Overview

The `/sys` command is your gateway to system management. It provides a tree view of all available subcommands with suggestions.

### Usage

```bash
/sys                    # Show system tree
/sys engagement         # Engagement dashboard
/sys role               # Role management
/sys security           # Security settings
/sys control            # System control
```

### Tree View

When you run `/sys` without arguments:

```
DoPlan System Control Panel:
├── engagement    - View engagement dashboard
├── role          - Manage roles and permissions
│   ├── tree      - Show role hierarchy
│   ├── list      - List all available roles
│   ├── show <role> - Show details for a specific role
│   └── assign <role> - Assign a role to the current user
├── security      - Security settings and tests
│   ├── status    - Show current security status
│   ├── test      - Run standard security tests
│   ├── release test - Run pre-release security tests
│   └── audit     - Run a comprehensive security audit
└── control       - System control (agents, roles, global kill switch)
    ├── system on|off - Global kill switch (requires confirmation)
    ├── agents on|off - Enable/disable agents
    └── roles on|off - Enable/disable roles
```

---

## `/sys engagement` - Engagement Dashboard

### Overview

View comprehensive engagement statistics, achievements, challenges, and progress.

### Usage

```bash
/sys engagement
```

### What It Shows

1. **Total Score** - Your current points (0-100,000+)
2. **Achievements** - Count of achievements earned
3. **Challenges** - Count of challenges completed
4. **Relationship Level** - 0-100 relationship strength
5. **Engagement Score** - 0-100% engagement percentage
6. **Last Reward** - Time since last achievement/challenge
7. **Pending Rewards** - Rewards scheduled for release
8. **Next Milestones** - Hints for upcoming achievements

### Example Output

```
📊 DoPlan Engagement Dashboard
============================================================

  💰 Total Score: 1,250 points
  🏆 Achievements: 8
  🎯 Challenges: 2

  🤝 Relationship Level: 45/100 💪 Developing
  📈 Engagement: 80% 👍 Good!

  ⏰ Last Reward: 2 hours ago
  ⏳ Pending Rewards: 1 (coming soon!)

  🎯 Next Milestones:
     🎯 'On the Rise': Reach 250 points
     📋 'Planner': Use /plan command 25 times
     🚀 'Do It Again': Use /do command 50 times
```

### When to Use

- Check your progress
- See achievements earned
- View relationship level
- Check pending rewards
- Get hints for next achievements
- Stay motivated

### Engagement Features

- **Score Tracking** - See your total points
- **Achievement Gallery** - View all earned achievements
- **Challenge Progress** - See completed challenges
- **Relationship Metrics** - Understand your relationship with DoPlan
- **Next Steps** - Hints for upcoming achievements

---

## `/sys role` - Role Management

### Overview

Manage roles and permissions for your DoPlan system. Roles control what features and commands are available.

### Usage

```bash
/sys role                    # Show role tree
/sys role tree               # Show role hierarchy
/sys role list               # List all available roles
/sys role show <role>        # Show details for a role
/sys role assign <role>       # Assign a role to current user
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

### When to Use

- View available roles
- Understand permissions
- Assign appropriate role
- Check current role
- Manage access control

---

## `/sys security` - Security Settings

### Overview

Run security tests, audits, and checks to ensure your project's security.

### Usage

```bash
/sys security                # Show security status
/sys security status         # Show current security status
/sys security test           # Run standard security tests
/sys security release test   # Run pre-release security tests
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

### When to Use

- Before releases
- Regular security checks
- After adding dependencies
- Before deployment
- Security compliance

---

## `/sys control` - System Control

### Overview

Control DoPlan's system features: enable/disable agents, roles, or the entire system.

### Usage

```bash
/sys control                 # Show control options
/sys control system on|off    # Global kill switch
/sys control agents on|off    # Enable/disable agents
/sys control roles on|off     # Enable/disable roles
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

### When to Use

- Emergency system shutdown
- Testing without agents
- Debugging system issues
- Temporary feature disable
- System maintenance

---

## 💡 System Command Tips

### Best Practices
- Check `/sys engagement` regularly for motivation
- Use `/sys role` to manage permissions appropriately
- Run `/sys security release test` before releases
- Use `/sys control` only when necessary
- Review `/sys` tree to discover features

### Power User
- Customize roles for your workflow
- Run security audits regularly
- Monitor engagement metrics
- Use system control for testing
- Explore all subcommands

### Safety
- Global kill switch requires confirmation
- State is preserved when disabling
- Can always re-enable features
- No data loss from control commands

---

## 🔄 System Command Examples

### Check Engagement

```bash
/sys engagement
# → View full engagement dashboard
```

### Manage Roles

```bash
/sys role tree
# → See role hierarchy

/sys role assign developer
# → Assign developer role
```

### Security Check

```bash
/sys security test
# → Run standard security tests

/sys security release test
# → Pre-release security check
```

### System Control

```bash
/sys control agents off
# → Disable agents temporarily

/sys control system off
# → Disable entire system (with confirmation)
```

---

**Next**: [Command Reference](./04-Command-Reference.md)

