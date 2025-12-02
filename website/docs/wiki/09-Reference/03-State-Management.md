# State Management Reference

Complete reference for state management in DoPlan. Understanding active_state.json and state snapshots.

---

## 📊 Active State

### Location

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

- **phase** - Current workflow phase (`ideation`, `planning`, `building`)
- **active_task** - Currently active task ID (e.g., "2.1")
- **active_branch** - Current Git branch (e.g., "task/2.1")
- **completed** - Array of completed task IDs
- **task_started_at** - ISO 8601 timestamp when task started
- **locked** - Whether state is locked (boolean)

---

## 📸 State Snapshots

### Location

```
.do/system/history/state-*.json
```

### Naming Format

```
state-[timestamp]-[reason].json
```

**Examples**:
- `state-20250115T100000Z-build.json` - Before /dev
- `state-20250115T101530Z-done.json` - After /done
- `state-20250115T120000Z-status.json` - Status check

### Timestamp Format

ISO 8601: `YYYYMMDDTHHMMSSZ`

**Example**: `20250115T100000Z` = 2025-01-15 10:00:00 UTC

---

## 🔄 State Updates

### When State Updates

- **Starting task** (`/dev`) - Sets active_task, active_branch, task_started_at
- **Completing task** (`/done`) - Adds to completed, clears active_task/branch
- **Phase changes** - Updates phase
- **Manual updates** - If needed

### Update Process

1. Read current state
2. Apply changes
3. Create snapshot
4. Write new state

---

## 🔙 Rollback

### Restore from Snapshot

```bash
# 1. Find snapshot
ls -lt .do/system/history/state-*.json

# 2. Review snapshot
cat .do/system/history/state-20250115T100000Z-build.json

# 3. Restore
cp .do/system/history/state-20250115T100000Z-build.json \
   .do/system/history/active_state.json
```

---

**Next**: [Time Tracker Format](./04-Time-Tracker-Format.md)

