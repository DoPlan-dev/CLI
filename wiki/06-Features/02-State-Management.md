# State Management

DoPlan maintains complete state history with rollback capability. Track your project state, see what changed, and rollback if needed.

---

## 🎯 What is State Management?

State Management in DoPlan tracks:
- **Current state** - Where you are now
- **State history** - Complete history with snapshots
- **State deltas** - What changed between states
- **Rollback capability** - Restore previous states

---

## 📊 Active State

### Location

Current state stored in:
```
.do/system/history/active_state.json
```

### Structure

```json
{
  "phase": "building",
  "active_task": "2.1",
  "active_branch": "task/2.1",
  "completed": ["1.1", "1.2", "1.3"],
  "task_started_at": "2025-01-15T10:00:00Z",
  "locked": false
}
```

### Fields

- **phase** - Current workflow phase
- **active_task** - Currently active task ID
- **active_branch** - Current Git branch
- **completed** - Array of completed task IDs
- **task_started_at** - When current task started
- **locked** - Whether state is locked

---

## 📸 State Snapshots

### Location

State snapshots stored in:
```
.do/system/history/state-*.json
```

### Naming

Format: `state-[timestamp]-[reason].json`

Examples:
- `state-20250115T100000Z-build.json` - Before /dev
- `state-20250115T101530Z-done.json` - After /done
- `state-20250115T120000Z-status.json` - Status check

### When Created

Snapshots created:
- **Before `/dev`** - When starting task
- **After `/done`** - When completing task
- **On status checks** - Optional
- **Manually** - If needed

---

## 🔄 State Updates

### When State Updates

State updates when:
- **Starting task** (`/dev`) - Sets active_task, active_branch, task_started_at
- **Completing task** (`/done`) - Adds to completed, clears active_task/branch
- **Phase changes** - Updates phase
- **Manual updates** - If needed

### Update Process

1. **Read current state** - Load active_state.json
2. **Apply changes** - Update fields
3. **Create snapshot** - Save current state
4. **Write new state** - Save updated state

---

## 📊 State Deltas

### What Are State Deltas?

State deltas show what changed between states:
- Tasks completed
- Phase changes
- Branch changes
- New active tasks

### Viewing Deltas

```bash
/status
```

**Shows**:
```
State Delta (since last snapshot):
  • Task 1.3 completed
  • Phase: Foundation (no change)
  • Branch: task/1.3 → cleared
```

### Delta Information

- **Tasks completed** - New completed tasks
- **Phase changes** - Phase transitions
- **Branch changes** - Branch switches
- **Active task** - Current task changes

---

## 🔙 Rollback Capability

### How Rollback Works

1. **Find snapshot** - Locate desired state
2. **Read snapshot** - Load state data
3. **Restore state** - Write to active_state.json
4. **Update files** - Restore TASKS.md if needed

### When to Rollback

- **Mistake made** - Wrong task marked complete
- **State corruption** - State file issues
- **Experiment** - Try different approach
- **Recovery** - Restore from backup

### Rollback Process

```bash
# 1. Find snapshot
ls .do/system/history/state-*.json

# 2. Review snapshot
cat .do/system/history/state-20250115T100000Z-build.json

# 3. Restore (manual process)
# Copy snapshot to active_state.json
cp .do/system/history/state-20250115T100000Z-build.json \
   .do/system/history/active_state.json
```

---

## 📁 State Files

### Active State

**File**: `.do/system/history/active_state.json`
**Purpose**: Current project state
**Updates**: On every state change
**Format**: JSON

### State Snapshots

**Directory**: `.do/system/history/`
**Files**: `state-*.json`
**Purpose**: Historical states
**Format**: JSON
**Retention**: All snapshots kept

---

## 🎯 State Use Cases

### 1. Progress Tracking

**Use**: See where you are
**How**: Read active_state.json
**Benefit**: Know current status

### 2. History Review

**Use**: See what changed
**How**: Compare snapshots
**Benefit**: Understand evolution

### 3. Rollback

**Use**: Restore previous state
**How**: Copy snapshot to active
**Benefit**: Recovery capability

### 4. Analysis

**Use**: Analyze state changes
**How**: Process snapshot files
**Benefit**: Insights into workflow

---

## 💡 State Management Tips

### Regular Snapshots

- Snapshots created automatically
- No manual action needed
- All states preserved
- Full history available

### State Review

- Check `/status` regularly
- Review state deltas
- Understand changes
- Track progress

### Rollback Safety

- Snapshots are safe
- Can always restore
- No data loss
- Full recovery

### State Analysis

- Process snapshot files
- Analyze state changes
- Understand patterns
- Optimize workflow

---

## 🔒 State Safety

### Automatic Snapshots

- Created before major changes
- Preserved permanently
- No manual action needed
- Full history maintained

### State Integrity

- Validated on read
- Error handling
- Recovery mechanisms
- Safe operations

### Backup Capability

- Snapshots serve as backups
- Can restore any state
- Full project recovery
- Safety net

---

## 🚀 Advanced State Management

### Custom Snapshots

Create snapshots manually:
```bash
# Copy current state
cp .do/system/history/active_state.json \
   .do/system/history/state-$(date +%Y%m%dT%H%M%SZ)-custom.json
```

### State Comparison

Compare states:
```bash
# Compare two snapshots
diff .do/system/history/state-1.json \
     .do/system/history/state-2.json
```

### State Analysis

Analyze state history:
```bash
# Count snapshots
ls .do/system/history/state-*.json | wc -l

# List all states
ls -lt .do/system/history/state-*.json
```

---

**Next**: [Git Automation](./03-Git-Automation.md)

