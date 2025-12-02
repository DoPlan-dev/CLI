# Troubleshooting

Common issues and solutions. Fix problems quickly and get back to building.

---

## 🚨 Common Issues

### Command Not Found

**Problem**: `doplan: command not found`

**Solutions**:
1. **Check installation**:
   ```bash
   doplan --version
   ```

2. **Verify PATH**:
   ```bash
   echo $PATH
   ```

3. **Use npx**:
   ```bash
   npx @doplan-dev/cli
   ```

4. **Reinstall**:
   ```bash
   # macOS
   brew reinstall doplan
   
   # Or download from GitHub Releases
   ```

---

### Missing Files

**Problem**: `/plan` fails with "IDEA.md not found"

**Solutions**:
1. **Run `/do` first**:
   ```bash
   /do
   # → Creates IDEA.md and BRAINSTORM.md
   /plan
   ```

2. **Check file location**:
   ```bash
   ls .do/system/IDEA.md
   ```

3. **Verify project structure**:
   ```bash
   ls -la .do/system/
   ```

---

### State Issues

**Problem**: State file corrupted or missing

**Solutions**:
1. **Check active state**:
   ```bash
   cat .do/system/history/active_state.json
   ```

2. **Restore from snapshot**:
   ```bash
   # Find latest snapshot
   ls -lt .do/system/history/state-*.json
   
   # Restore
   cp .do/system/history/state-[timestamp].json \
      .do/system/history/active_state.json
   ```

3. **Reset state**:
   ```bash
   # Create new active state
   echo '{"phase":"ideation","completed":[]}' > \
        .do/system/history/active_state.json
   ```

---

### Git Issues

**Problem**: Auto-commit/push fails

**Solutions**:
1. **Check Git status**:
   ```bash
   git status
   ```

2. **Verify remote**:
   ```bash
   git remote -v
   ```

3. **Manual commit**:
   ```bash
   git add .
   git commit -m "feat(task-X.X): your message"
   git push
   ```

4. **Check permissions**:
   ```bash
   # Verify Git credentials
   git config --list
   ```

---

### Memory Card Issues

**Problem**: Personalization not working

**Solutions**:
1. **Check Memory Card**:
   ```bash
   cat ~/.doplan/memory_card.json
   ```

2. **Verify location**:
   ```bash
   ls -la ~/.doplan/
   ```

3. **Recreate if needed**:
   ```bash
   # Delete and restart
   rm ~/.doplan/memory_card.json
   /hey
   # → Creates new Memory Card
   ```

---

## 🔧 Advanced Troubleshooting

### Time Tracking Not Working

**Problem**: Time not being tracked

**Solutions**:
1. **Check time-tracker.jsonl**:
   ```bash
   cat .do/system/time-tracker.jsonl
   ```

2. **Verify directory**:
   ```bash
   ls -la .do/system/
   ```

3. **Check permissions**:
   ```bash
   ls -la .do/system/time-tracker.jsonl
   ```

---

### Engagement System Issues

**Problem**: Achievements not being awarded

**Solutions**:
1. **Check engagement**:
   ```bash
   /sys engagement
   ```

2. **Verify Memory Card**:
   ```bash
   cat ~/.doplan/memory_card.json | jq .achievements
   ```

3. **Check score**:
   ```bash
   cat ~/.doplan/memory_card.json | jq .score
   ```

---

### Branch Issues

**Problem**: Branch not created or wrong branch

**Solutions**:
1. **Check current branch**:
   ```bash
   git branch
   ```

2. **Verify active state**:
   ```bash
   cat .do/system/history/active_state.json | jq .active_branch
   ```

3. **Create branch manually**:
   ```bash
   git checkout -b task/2.1
   ```

---

## 💡 Troubleshooting Tips

### General

1. **Check error messages** - Read carefully
2. **Verify file locations** - Ensure files exist
3. **Check permissions** - Read/write access
4. **Review state** - Check active_state.json
5. **Use `/status`** - See current state

### Specific Issues

1. **Command issues** - Check installation
2. **File issues** - Verify project structure
3. **State issues** - Check state files
4. **Git issues** - Verify Git setup
5. **Memory Card** - Check home directory

---

## 🆘 Getting Help

### Documentation

- **Wiki** - Complete documentation
- **README** - Quick reference
- **Command help** - `doplan --help`

### Community

- **GitHub Issues** - Report problems
- **Discussions** - Ask questions
- **Contributing** - Help improve

---

**Next**: [Contributing](./05-Contributing.md)

