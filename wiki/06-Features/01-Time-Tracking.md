# Time Tracking

DoPlan automatically tracks time for every command, phase, and task. No manual logging needed - everything is tracked automatically!

---

## ⏱️ What is Time Tracking?

Time tracking in DoPlan is **completely automatic**. It tracks:
- Every command execution
- Each phase (ideation, planning, development)
- Complete task duration (from `/dev` start to auto-complete)
- Session duration
- Project time

**No manual action required!**

---

## 📊 What Gets Tracked

### Command Execution

Every command tracks:
- **Start time** - When command started
- **End time** - When command completed
- **Duration** - Total execution time
- **Status** - Completed or failed
- **Metadata** - Work style, experience level, etc.

### Task Duration

Complete task duration tracked:
- **Task start** - When `/dev` is run
- **Task end** - When `/dev` auto-completes
- **Total duration** - Full task time
- **Displayed** - Shown after completion

### Phase Tracking

Each phase tracked:
- **Ideation** - Time spent capturing idea
- **Planning** - Time spent generating plan
- **Development** - Time spent coding
- **Completion** - Time spent completing tasks

---

## 📁 Time Data Storage

### Location

Time data is stored in:
```
.do/system/time-tracker.jsonl
```

### Format (JSONL)

One entry per line in JSON format:

```json
{
  "project": "/path/to/project",
  "command": "/dev",
  "phase": "development",
  "args": ["auth"],
  "metadata": {
    "work_style": "fast",
    "experience_level": "intermediate"
  },
  "started_at": "2025-01-15T10:00:00Z",
  "ended_at": "2025-01-15T10:15:30Z",
  "duration_ms": 930000,
  "status": "completed"
}
```

### Why JSONL?

- **Append-only** - Easy to add new entries
- **Analytics-ready** - Perfect for dashboards
- **Human-readable** - Can view with text editor
- **Efficient** - No need to parse entire file

---

## 🎯 Time Tracking Examples

### Command Execution

**Example**: Running `/plan`

```json
{
  "project": "/Users/you/my-project",
  "command": "/plan",
  "phase": "planning",
  "started_at": "2025-01-15T10:00:00Z",
  "ended_at": "2025-01-15T10:01:30Z",
  "duration_ms": 90000,
  "status": "completed"
}
```

**Duration**: 1 minute 30 seconds

### Task Duration

**Example**: Task from `/dev` start to auto-complete

1. **Task Start** (`/dev`):
   - `task_started_at` stored in active_state.json
   - Time: 10:00 AM

2. **Development**:
   - You code for 2 hours

3. **Auto-Completion** (handled by `/dev`):
   - Duration calculated: 2 hours 15 minutes
   - Displayed: "⏱️ Task duration: 2h 15m"

**Total**: 2 hours 15 minutes tracked automatically

---

## 📊 Viewing Time Data

### View Recent Entries

```bash
# View last 5 entries
tail -5 .do/system/time-tracker.jsonl
```

### View All Entries

```bash
# View all entries
cat .do/system/time-tracker.jsonl
```

### Parse with jq (if available)

```bash
# View all /dev commands
cat .do/system/time-tracker.jsonl | jq 'select(.command == "/dev")'

# Calculate total time
cat .do/system/time-tracker.jsonl | jq -s 'map(.duration_ms) | add'
```

---

## 🎯 Time Tracking Features

### Automatic Tracking

- **No setup** - Works automatically
- **No configuration** - Just works
- **No manual logging** - Completely automatic
- **Comprehensive** - Tracks everything

### Metadata Included

Each entry includes:
- **Work style** - Fast or thoughtful
- **Experience level** - Beginner, intermediate, advanced
- **Command args** - Arguments passed
- **Project path** - Which project
- **Phase** - Current phase

### Status Tracking

Tracks success/failure:
- **Completed** - Command succeeded
- **Failed** - Command failed
- **Error message** - If failed

---

## 💡 Use Cases

### Time Analysis

Use time data to:
- Analyze productivity
- Identify time sinks
- Optimize workflow
- Track improvement

### Billing (if needed)

Time data can be used for:
- Client billing
- Project estimates
- Time reporting
- Productivity metrics

### Future Dashboards

Time data is ready for:
- Visual dashboards
- Analytics tools
- Reporting systems
- Progress tracking

---

## 🎯 Time Display

### After Task Completion

When `/dev` auto-completes a task:

```
✅ Task 2.1 marked complete!
   ⏱️  Task duration: 2h 15m
```

**Shows**:
- Hours and minutes format
- Readable display
- Accurate duration

### Duration Calculation

Duration is calculated from:
- `task_started_at` in active_state.json (when `/dev` was run)
- Current time (when `/dev` auto-completes)
- Difference = task duration

---

## 📊 Time Analytics

### Command Usage Time

Track time spent on:
- `/hey` - Onboarding time
- `/do` - Ideation time
- `/plan` - Planning time
- `/dev` - Development and auto-completion time

### Phase Time

Track time by phase:
- **Ideation** - Idea capture time
- **Planning** - Plan generation time
- **Development** - Coding time
- **Completion** - Auto-completion time

### Project Time

Track total project time:
- Sum of all command times
- Sum of all task durations
- Total development time
- Overall project duration

---

## 💡 Time Tracking Tips

### For Analysis

- **Export data** - Use JSONL for analysis
- **Parse with tools** - jq, Python, etc.
- **Create dashboards** - Visualize time data
- **Track trends** - See improvement over time

### For Productivity

- **Review regularly** - See where time goes
- **Identify patterns** - Find time sinks
- **Optimize workflow** - Improve efficiency
- **Set goals** - Time-based targets

### For Billing

- **Export entries** - Use for billing
- **Calculate totals** - Sum durations
- **Generate reports** - Time reports
- **Track by project** - Per-project billing

---

## 🚀 Future Enhancements

Time tracking data enables:
- **Visual dashboards** - Charts and graphs
- **Analytics tools** - Deep insights
- **Reporting systems** - Automated reports
- **Productivity metrics** - Performance tracking

---

**Next**: [State Management](./02-State-Management.md)

