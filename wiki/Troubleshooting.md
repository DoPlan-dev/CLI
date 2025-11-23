# Troubleshooting Guide

Common issues and solutions when using DoPlan CLI.

---

## Installation Problems

### Issue: "Command not found" or "doplan: command not found"

**Symptoms**: Terminal cannot find the `doplan` command after installation.

**Solutions**:

1. **Verify Installation**:
   ```bash
   # macOS/Linux
   which doplan
   
   # Windows
   where doplan
   ```

2. **Check PATH**:
   - Ensure the binary location is in your PATH
   - **macOS/Linux**: Add to `~/.bashrc`, `~/.zshrc`, or `~/.profile`
   - **Windows**: Add to System Environment Variables

3. **Restart Terminal**: Close and reopen your terminal after updating PATH

4. **Verify Binary Location**:
   ```bash
   # macOS/Linux
   ls -la /usr/local/bin/doplan
   
   # Or check your installation directory
   ```

5. **Reinstall**: If still not found, reinstall using one of the methods in the [Installation Guide](Installation)

---

### Issue: "Permission denied" (macOS/Linux)

**Symptoms**: Cannot execute the binary due to permission errors.

**Solutions**:

1. **Make Binary Executable**:
   ```bash
   chmod +x doplan
   ```

2. **If Installed to System Location**:
   ```bash
   sudo chmod +x /usr/local/bin/doplan
   ```

3. **Check File Permissions**:
   ```bash
   ls -la doplan
   # Should show: -rwxr-xr-x
   ```

---

### Issue: Wrong Architecture Binary Downloaded

**Symptoms**: Binary fails to run or shows architecture mismatch errors.

**Solutions**:

1. **Check Your Architecture**:
   ```bash
   # macOS/Linux
   uname -m
   # Should show: x86_64 (Intel) or arm64 (Apple Silicon/ARM)
   
   # Windows
   echo %PROCESSOR_ARCHITECTURE%
   ```

2. **Download Correct Binary**:
   - **macOS Intel**: `doplan-darwin-amd64`
   - **macOS Apple Silicon**: `doplan-darwin-arm64`
   - **Linux amd64**: `doplan-linux-amd64`
   - **Linux arm64**: `doplan-linux-arm64`
   - **Windows**: `doplan-windows-amd64.exe`

3. **Use npx**: `npx` automatically downloads the correct binary for your platform

---

### Issue: Homebrew Installation Fails

**Symptoms**: `brew install doplan` fails or tap not found.

**Solutions**:

1. **Update Homebrew**:
   ```bash
   brew update
   ```

2. **Add Tap**:
   ```bash
   brew tap doplan-dev/cli
   ```

3. **Check Tap**:
   ```bash
   brew tap list | grep doplan
   ```

4. **Try Installing Again**:
   ```bash
   brew install doplan
   ```

5. **Check for Errors**:
   ```bash
   brew install doplan --verbose
   ```

---

### Issue: Build from Source Fails

**Symptoms**: `go build` fails with errors.

**Solutions**:

1. **Check Go Version**:
   ```bash
   go version
   # Should be >= 1.23.0
   ```

2. **Update Go** if needed:
   ```bash
   # macOS
   brew upgrade go
   
   # Or download from https://golang.org/dl/
   ```

3. **Check Internet Connection**: Building requires downloading dependencies

4. **Clean Build**:
   ```bash
   go clean -cache
   go build -o doplan ./cmd/doplan
   ```

5. **Verbose Build**:
   ```bash
   go build -v -o doplan ./cmd/doplan
   ```

---

## Binary Not Found

### Issue: npx Cannot Find Binary

**Symptoms**: `npx @doplan-dev/cli` fails to download or find binary.

**Solutions**:

1. **Clear npm Cache**:
   ```bash
   npm cache clean --force
   ```

2. **Check Internet Connection**: First run requires downloading

3. **Try Again**:
   ```bash
   npx @doplan-dev/cli
   ```

4. **Check npm Version**:
   ```bash
   npm --version
   # Should be recent version
   ```

5. **Use Specific Version**:
   ```bash
   npx @doplan-dev/cli@latest
   ```

---

### Issue: Binary Downloaded but Not Executable

**Symptoms**: Binary exists but cannot be executed.

**Solutions**:

1. **Check Permissions**:
   ```bash
   ls -la doplan
   ```

2. **Make Executable**:
   ```bash
   chmod +x doplan
   ```

3. **Check File Type**:
   ```bash
   file doplan
   # Should show: Mach-O binary or ELF binary
   ```

---

## Permission Issues

### Issue: Cannot Write to Project Directory

**Symptoms**: DoPlan CLI cannot create files or directories.

**Solutions**:

1. **Check Directory Permissions**:
   ```bash
   ls -la .
   ```

2. **Fix Permissions**:
   ```bash
   chmod u+w .
   ```

3. **Check Disk Space**:
   ```bash
   df -h .
   ```

4. **Run with Appropriate Permissions**: Ensure you have write access to the directory

---

### Issue: Cannot Access .cursor Directory

**Symptoms**: IDE cannot access or read `.cursor/` directory.

**Solutions**:

1. **Check Directory Exists**:
   ```bash
   ls -la .cursor
   ```

2. **Check Permissions**:
   ```bash
   ls -la .cursor/agents
   ls -la .cursor/commands
   ```

3. **Fix Permissions**:
   ```bash
   chmod -R u+r .cursor
   ```

4. **Verify IDE Access**: Ensure your IDE has read access to the project directory

---

## Network Issues

### Issue: Cannot Download Binary via npx

**Symptoms**: npx fails to download the binary.

**Solutions**:

1. **Check Internet Connection**:
   ```bash
   ping github.com
   ```

2. **Check npm Registry**:
   ```bash
   npm config get registry
   ```

3. **Try Direct Download**:
   - Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
   - Download binary manually

4. **Check Firewall/Proxy**: Ensure npm can access external resources

5. **Use Alternative Method**: Install via Homebrew, Scoop, or direct binary

---

### Issue: Slow Download

**Symptoms**: npx download is very slow.

**Solutions**:

1. **Check Network Speed**: Test your internet connection

2. **Use Alternative Method**: Download binary directly from GitHub Releases

3. **Check npm Registry**: Use a faster registry if needed

4. **Retry**: Sometimes network issues are temporary

---

## IDE Integration Issues

### Issue: IDE Doesn't Recognize Commands

**Symptoms**: Typing `/tell` or other commands doesn't work in IDE.

**Solutions**:

1. **Verify IDE Support**: Ensure you're using a supported IDE:
   - Cursor
   - Claude Code
   - Antigravity
   - Windsurf
   - Cline
   - OpenCode

2. **Check Configuration Files**: Verify IDE configuration files exist:
   ```bash
   ls -la .cursor
   ```

3. **Restart IDE**: Close and reopen your IDE

4. **Check IDE Settings**: Ensure AI features are enabled

5. **See IDE Integration Guide**: Check [IDE Integration](IDE-Integration) for specific setup

---

### Issue: Agents Not Loading

**Symptoms**: Agents don't appear or don't work.

**Solutions**:

1. **Check Agent Files**:
   ```bash
   ls -la .cursor/agents/
   # Should show 18 agent files
   ```

2. **Verify Agent Definitions**: Check that agent files contain valid markdown

3. **Check Command Files**:
   ```bash
   ls -la .cursor/commands/
   ```

4. **Restart IDE**: Sometimes IDE needs restart to load new files

5. **Check IDE Logs**: Look for errors in IDE console/logs

---

### Issue: Commands Not Working

**Symptoms**: Commands don't execute or produce errors.

**Solutions**:

1. **Check Command Files**:
   ```bash
   ls -la .cursor/commands/
   ```

2. **Verify Command Syntax**: Ensure commands are defined correctly

3. **Check Project State**:
   ```bash
   cat .plan/active_state.json
   ```

4. **Verify Dependencies**: Ensure required files exist (IDEA.md, PRD.md, etc.)

5. **Check IDE Compatibility**: Ensure your IDE version supports the command system

---

## Project Generation Issues

### Issue: Project Generation Fails

**Symptoms**: DoPlan CLI fails to generate project.

**Solutions**:

1. **Check Project Name**: 
   - Use lowercase letters, numbers, and hyphens
   - No spaces or special characters

2. **Check Directory Permissions**: Ensure you have write access

3. **Check Disk Space**: Ensure sufficient disk space

4. **Check for Existing Directory**: Ensure project name doesn't conflict

5. **Try Again**: Sometimes temporary issues occur

---

### Issue: Incomplete Project Generation

**Symptoms**: Some files or directories are missing.

**Solutions**:

1. **Check Generation Logs**: Look for error messages

2. **Verify All Files**:
   ```bash
   ls -la .cursor/agents/
   ls -la .cursor/commands/
   ls -la .plan/
   ```

3. **Regenerate**: Delete project and regenerate

4. **Check for Errors**: Look for permission or disk space issues

---

### Issue: Wrong Project Structure

**Symptoms**: Generated structure doesn't match expectations.

**Solutions**:

1. **Check DoPlan CLI Version**: Ensure you're using the latest version

2. **Verify Generation**: Check that generation completed successfully

3. **Customize**: You can modify the structure after generation

4. **Report Issue**: If structure is incorrect, [open an issue](https://github.com/DoPlan-dev/CLI/issues)

---

## Command-Specific Issues

### Issue: `/tell` Doesn't Save Idea

**Solutions**:

1. **Check File Permissions**: Ensure write access to `.plan/00_System/`

2. **Verify Directory Exists**:
   ```bash
   ls -la .plan/00_System/
   ```

3. **Check State File**:
   ```bash
   cat .plan/active_state.json
   ```

4. **Try Again**: Sometimes temporary issues occur

---

### Issue: `/write` Doesn't Generate Documents

**Solutions**:

1. **Check Prerequisites**: Ensure IDEA.md and BRAINSTORM.md exist

2. **Verify Files**:
   ```bash
   ls -la .plan/00_System/
   ```

3. **Check Permissions**: Ensure write access

4. **Try `/improve` First**: Ensure brainstorm is complete

---

### Issue: `/build` Fails

**Solutions**:

1. **Check Tasks Exist**:
   ```bash
   cat .plan/TASKS.md
   ```

2. **Verify Plan Approved**: Check `active_state.json` for `locked: true`

3. **Check Task ID**: Ensure task ID exists if using `/build <task_id>`

4. **Verify State**:
   ```bash
   cat .plan/active_state.json
   ```

---

### Issue: `/finished` Doesn't Commit

**Solutions**:

1. **Check Git Repository**: Ensure project is a git repository

2. **Check Git Configuration**: Ensure git user.name and user.email are set

3. **Check for Changes**: Ensure there are changes to commit

4. **Check Branch**: Ensure you're on a valid branch

5. **Manual Commit**: You can commit manually if needed

---

## Debug Mode

### Enabling Debug Mode

To get more detailed error information:

1. **Set Debug Environment Variable**:
   ```bash
   # macOS/Linux
   export DOPLAN_DEBUG=1
   doplan
   
   # Windows
   set DOPLAN_DEBUG=1
   doplan
   ```

2. **Check Logs**: Look for detailed error messages

---

## Log Files

### Log File Locations

DoPlan CLI may create log files in:

- **macOS/Linux**: `~/.doplan/logs/`
- **Windows**: `%APPDATA%\doplan\logs\`

### Checking Logs

```bash
# macOS/Linux
cat ~/.doplan/logs/doplan.log

# Windows
type %APPDATA%\doplan\logs\doplan.log
```

---

## Getting Help

### Resources

1. **Documentation**:
   - [FAQ](FAQ) - Common questions
   - [Commands Reference](Commands) - Command details
   - [Workflow Guide](Workflow) - Workflow details

2. **Community**:
   - [GitHub Issues](https://github.com/DoPlan-dev/CLI/issues) - Report bugs and ask questions
   - [GitHub Discussions](https://github.com/DoPlan-dev/CLI/discussions) - Community discussions

3. **Support**:
   - Check existing issues before creating new ones
   - Provide detailed information when reporting issues
   - Include error messages and logs

---

## Issue Reporting Guidelines

When reporting an issue, please include:

1. **DoPlan CLI Version**:
   ```bash
   doplan --version
   ```

2. **Platform Information**:
   ```bash
   # macOS/Linux
   uname -a
   
   # Windows
   systeminfo
   ```

3. **Error Messages**: Copy the full error message

4. **Steps to Reproduce**: Detailed steps to reproduce the issue

5. **Expected Behavior**: What you expected to happen

6. **Actual Behavior**: What actually happened

7. **Logs**: Include relevant log files (if available)

8. **Screenshots**: If applicable

---

## Common Error Messages

### "Project name is invalid"

**Solution**: Use lowercase letters, numbers, and hyphens only. No spaces or special characters.

### "IDE not supported"

**Solution**: Ensure you select a supported IDE from the list.

### "Cannot create directory"

**Solution**: Check permissions and disk space.

### "Binary not found"

**Solution**: Reinstall or use npx method.

### "Command failed"

**Solution**: Check prerequisites and file permissions.

---

## Related Pages

- [Installation](Installation) - Installation guide
- [FAQ](FAQ) - Frequently asked questions
- [Commands Reference](Commands) - Command details
- [IDE Integration](IDE-Integration) - IDE setup
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

