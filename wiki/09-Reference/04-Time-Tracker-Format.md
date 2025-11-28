# Time Tracker Format

Understanding the time-tracker.jsonl format for time tracking data.

---

## 📊 File Location

```
.do/system/time-tracker.jsonl
```

---

## 📝 Format

JSONL (JSON Lines) - One entry per line

### Entry Structure

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
  "status": "completed",
  "error": ""
}
```

### Fields

- **project** - Project path (string)
- **command** - Command name (string)
- **phase** - Workflow phase (string)
- **args** - Command arguments (array)
- **metadata** - Additional data (object)
- **started_at** - Start time (ISO 8601)
- **ended_at** - End time (ISO 8601)
- **duration_ms** - Duration in milliseconds (int64)
- **status** - "completed" or "failed" (string)
- **error** - Error message if failed (string)

---

## 💡 Usage Examples

### View Recent Entries

```bash
tail -5 .do/system/time-tracker.jsonl
```

### Parse with jq

```bash
# View all /dev commands
cat .do/system/time-tracker.jsonl | jq 'select(.command == "/dev")'

# Calculate total time
cat .do/system/time-tracker.jsonl | jq -s 'map(.duration_ms) | add'
```

---

**Next**: [Memory Card Schema](./05-Memory-Card-Schema.md)

