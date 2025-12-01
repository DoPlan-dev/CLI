# CI/CD Build Troubleshooting

## TAR Exit Code 2

When builds fail with tar errors (exit code 2), follow these steps:

### Quick Fixes

1. **Clear cache**: `rm -rf ~/.cache`
   - Clears any corrupted cached files that might interfere with tar operations

2. **Verify tar works**: `tar --version`
   - Ensures tar is installed and accessible
   - Check if tar is in PATH

3. **Check file paths are valid**
   - Verify the binary file exists before archiving
   - Ensure no special characters or spaces in paths
   - Check file permissions

4. **Ensure no corrupted artifacts**
   - Remove any existing archive files before rebuilding
   - Check disk space availability
   - Verify file system integrity

5. **Rebuild with verbose output**
   - Use `tar -czvf` instead of `tar -czf` for verbose output
   - This shows exactly what tar is processing
   - Helps identify which file or path is causing issues

### Common Causes

- **Missing binary**: The `doplan` binary wasn't built successfully
- **Path issues**: Working directory changed or binary in unexpected location
- **Permissions**: Insufficient permissions to read binary or write archive
- **Disk space**: No space left on device
- **Corrupted files**: Previous build left corrupted state

### Debugging Steps

```bash
# Check if binary exists
ls -la doplan

# Verify tar version
tar --version

# Test tar manually
tar -czvf test.tar.gz doplan

# Check disk space
df -h

# Check permissions
ls -la
```

### Workflow-Specific

In the GitHub Actions workflow, the archive step:
- Creates `.tar.gz` for Linux/macOS builds
- Creates `.zip` for Windows builds
- Generates SHA256 checksums for all archives

If tar fails, check:
1. The build step completed successfully
2. The binary file exists in the expected location
3. The version variable is set correctly
4. No matrix variable issues (goos/goarch)

