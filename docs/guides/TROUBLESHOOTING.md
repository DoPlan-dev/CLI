# DoPlan CLI Troubleshooting Guide

## Commands Not Showing in Cursor

### Quick Fixes (Try These First)

1. **Restart Cursor** ⭐ Most Common Solution
   - Close Cursor completely
   - Reopen Cursor
   - Open the project folder
   - Commands should appear when you type `/`

2. **Verify Installation**
   ```bash
   ls -la .cursor/commands/
   ```
   Should show 8 JSON files

3. **Check JSON Validity**
   ```bash
   python3 -m json.tool .cursor/commands/discuss.json
   ```

4. **Reload Cursor Window**
   - `Cmd+Shift+P` (Mac) or `Ctrl+Shift+P` (Windows/Linux)
   - Type "Reload Window"
   - Select "Developer: Reload Window"

### Common Issues

#### Issue: Commands directory doesn't exist
**Solution:** Re-run installation
```bash
doplan install
```

#### Issue: JSON files are invalid
**Solution:** Check file format
```bash
# Each file should have: name, description, prompt
cat .cursor/commands/discuss.json
```

#### Issue: Cursor version too old
**Solution:** Update Cursor to latest version
- Cursor 0.30+ required for custom commands

#### Issue: Wrong project folder opened
**Solution:** Ensure you opened the project root folder
- Commands are project-specific
- Must open folder containing `.cursor/` directory

### Verification Steps

1. **Check files exist:**
   ```bash
   ls .cursor/commands/*.json | wc -l
   # Should output: 8
   ```

2. **Verify one command file:**
   ```bash
   cat .cursor/commands/discuss.json
   ```
   Should show valid JSON with `name`, `description`, `prompt`

3. **Test JSON validity:**
   ```bash
   for f in .cursor/commands/*.json; do
     python3 -m json.tool "$f" > /dev/null && echo "✓ $f" || echo "✗ $f"
   done
   ```

### Expected Command Format

Each command file should look like:
```json
{
  "name": "/Discuss",
  "description": "Start idea discussion and refinement",
  "prompt": "Start the DoPlan idea discussion workflow..."
}
```

### Still Not Working?

1. **Check Cursor Developer Tools:**
   - Help → Toggle Developer Tools
   - Check Console for errors

2. **Try manual test command:**
   ```bash
   echo '{"name":"/Test","description":"Test","prompt":"Hello"}' > .cursor/commands/test.json
   ```
   Restart Cursor and try `/Test`

3. **Check Cursor settings:**
   - Ensure custom commands are enabled
   - Check for any command-related settings

4. **Fresh reinstall:**
   ```bash
   rm -rf .cursor
   doplan install
   ```

### Getting Help

If commands still don't appear:
1. Note your Cursor version
2. Check if other custom commands work
3. Verify file permissions
4. Check Cursor logs for errors

## Other Common Issues

### Installation Fails
- Ensure you're in a git repository
- Check write permissions
- Verify binary exists: `which doplan`

### Dashboard Not Generating
- Check `doplan/` directory exists
- Verify write permissions
- Run: `doplan progress` to regenerate

### GitHub Sync Fails
- Ensure git is initialized
- Check GitHub CLI is installed (for PR features)
- Verify repository has a remote

