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
# Expected output: -rwxr-xr-x  1 user  staff  10065376 Dec  1 05:10 doplan
# If missing, the build step failed - check build logs

# Verify tar version
tar --version
# Expected output: bsdtar 3.5.3 - libarchive 3.7.4 ... (macOS)
# Or: tar (GNU tar) 1.34 (Linux)
# If command not found, tar is not installed

# Test tar manually
tar -czvf test.tar.gz doplan
# Expected output: a doplan
# If this fails, there's a problem with tar or the file

# Check disk space
df -h
# Ensure you have sufficient space (at least 50MB free)
# Look at the "Avail" column for available space

# Check permissions
ls -la
# Verify the binary is executable (x permission)
# If not, run: chmod +x doplan
```

### Example: Successful Local Test

When troubleshooting locally, you should see output like this:

```bash
$ rm -rf ~/.cache
$ tar --version
bsdtar 3.5.3 - libarchive 3.7.4 zlib/1.2.12 liblzma/5.4.3 bz2lib/1.0.8

$ ls -la doplan
-rwxr-xr-x  1 user  staff  10065376 Dec  1 05:10 doplan

$ tar -czvf test.tar.gz doplan
a doplan

$ ls -lh test.tar.gz
-rw-r--r--  1 user  staff   5.0M Dec  1 09:20 test.tar.gz
```

If all these steps succeed locally but fail in CI, the issue is likely:
- Different tar implementation (GNU tar vs bsdtar)
- Different working directory
- CI runner environment differences
- Build step not completing before archive step

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

### CI vs Local Differences

**macOS (local):**
- Uses `bsdtar` (libarchive-based)
- More lenient with some edge cases
- Example: `bsdtar 3.5.3 - libarchive 3.7.4`

**Linux (CI):**
- Uses `GNU tar`
- Stricter error handling
- Example: `tar (GNU tar) 1.34`

**Common CI Issues:**
- Binary not found: Check if build step actually created the file
- Path issues: CI runner might be in different directory
- Permissions: CI runners may have different default permissions
- Timing: Archive step might run before build completes (unlikely with `needs:`)

### Additional CI Debugging

Add these steps to your workflow for debugging:

```yaml
- name: Debug before archive
  run: |
    echo "Working directory: $(pwd)"
    echo "Files in directory:"
    ls -la
    echo "Binary exists:"
    test -f doplan && echo "YES" || echo "NO"
    echo "Binary size:"
    ls -lh doplan || echo "Binary not found"
    echo "Tar version:"
    tar --version
```

